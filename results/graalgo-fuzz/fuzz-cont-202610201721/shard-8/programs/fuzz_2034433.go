package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int32 }
type GenS2 struct { F0 uint; F1 int; F2 uint64 }
type GenS3 struct { F0 bool; F1 string }

func main() {
    var v1 int = 0
    v2 := []uint32{42, 65535, 2147483647}
    v3 := []string{(("foo bar" + "\n\t") + ("" + "foo bar")), ("hello" + ""), ("a" + "foo bar")}
    var v4 uint8 = 201
    var v5 float32 = ((1e10 * 1.5) - ((1.5 * 1e10) / 0.1))
    _ = len(v3)
    v6 := GenS1{F0: (312 != 4294967295), F1: 127, F2: -670}
    _ = len(v2)
    v7 := GenS2{F0: 4294967295, F1: (v1 + (v1 >> ((104) & 63))), F2: 1}
    v4 -= 1
    var v8 int32 = 262
    var v9 uint16 = 42
    var v10 bool = (true && (v8 > 784))
    _ = len(v2)
    switch 127 {
    case 9223372036854775807:
        if len(v3) > 0 {
            v3[0] = ("\n\t" + "hello")
        }
    case 32768:
        v11 := GenS3{F0: false, F1: "a"}
        _ = v11
    default:
        v10 = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
