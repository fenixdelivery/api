#!/bin/bash

# Directorio raíz del proyecto
PROJECT_DIR="./"

# Comando para formatear archivos .go
GO_FMT_CMD="go fmt"

# Buscar todos los archivos .go dentro del directorio del proyecto
find "$PROJECT_DIR" -type f -name "*.go" -print0 | while IFS= read -r $0 file; do
  echo "Formateando: $file"
  "$GO_FMT_CMD" "$file"
  if [ $? -ne 0 ]; then
    echo "Error al formatear: $file"
  fi
done

echo "Formateo de archivos .go completado."