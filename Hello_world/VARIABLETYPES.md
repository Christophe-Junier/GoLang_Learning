# GoLang Variable Types

All type of variable usable with go

## INTEGER

### unsigned
  uint8 	 8 bits 	0 à 255
  byte 	   8 bits 	0 à 255
  uint16 	16 bits 	0 à 65535
  uint32 	32 bits 	0 à 4294967295
  uint64 	64 bits 	0 à 18446744073709551615
  uint 	  32/64 bits

### signed
  int8 	     8 bits 	-128 à 127
  int16 	  16 bits 	-32768 à 32767
  int32 	  32 bits 	-2147483648 à 2147483647
  rune      32 bits 	-2147483648 à 2147483647
  int64 	  64 bits 	-9223372036854775808 à 9223372036854775807
  float32 	32 bits 	Environ 7 chiffres décimaux
  float64 	64 bits 	Environ 16 chiffres décimaux
  int       32/64 bits

## BOOLEAN

  bool

## STRING

  string

# How to use variable in Go
  variable name can't have spacing 'ex: var bad named integer int'
  variable name can't have accents 'ex: var bad_named_intégèr int'
  variable name can't start by a number 'ex: var 1_bad_named_integer int'

exemple of variable declaration:

```go
  var signed_integer int16
```
