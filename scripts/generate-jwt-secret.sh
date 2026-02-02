#!/bin/bash

BYTES=${1:-32}
ENV_FILE=".env"

# Create temporary Go file
TMP_GO_FILE=$(mktemp /tmp/gen_jwt_secret_XXXX.go)

cat <<EOF > $TMP_GO_FILE
package main
import (
    "crypto/rand"
    "encoding/base64"
    "fmt"
)
func main() {
    b := make([]byte, $BYTES)
    _, err := rand.Read(b)
    if err != nil {
        panic(err)
    }
    fmt.Println(base64.StdEncoding.EncodeToString(b))
}
EOF

# Run temporary Go file
SECRET=$(go run "$TMP_GO_FILE")

# Remove temporary Go file
rm -f "$TMP_GO_FILE"

echo "🔑 Generated JWT secret: $SECRET"

# Ensure .env exists
if [ ! -f "$ENV_FILE" ]; then
    touch "$ENV_FILE"
fi

# Add or update JWT_SECRET
if grep -q "^JWT_SECRET=" "$ENV_FILE"; then
    # Update existing value
    sed -i "s|^JWT_SECRET=.*|JWT_SECRET=$SECRET|" "$ENV_FILE"
    echo "✅ Updated JWT_SECRET in $ENV_FILE"
else
    # Ensure file ends with newline
    [ -s "$ENV_FILE" ] && echo "" >> "$ENV_FILE"
    echo "JWT_SECRET=$SECRET" >> "$ENV_FILE"
    echo "✅ Added JWT_SECRET to $ENV_FILE"
fi

