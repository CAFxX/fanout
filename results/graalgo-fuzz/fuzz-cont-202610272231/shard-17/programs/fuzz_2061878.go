package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint16; F1 bool }
func genFn0(p0 int16) bool {
    var v5 int8 = 67
    p0 = (p0 % (((p0 / ((p0) | 1))) | 1))
    v6 := GenS1{F0: (v5 * (v5 & -58)), F1: 688}
    v5 = (v5 / ((-128) | 1))
    _ = p0
    _ = v5
    _ = v6
    return (!(!(true && (1e10 < (float64(float32((1.5 + 0.0))) + 1.5)))))
}

func main() {
    var v1 int16 = -604
    var v2 string = "foo bar"
    var v3 uint32 = 439
    var v4 int = 117
    v3 -= (v3 / (((v3 | (v3 % ((4294967295) | 1)))) | 1))
    v1 = v1
    var v7 int16 = (^(v1))
    switch "hello" {
    case "a":
        fmt.Println(v2)
    case "foo bar":
        v1 -= v1
    case "hello":
        v8 := GenS2{F0: 926, F1: (!true)}
        _ = v8
    }
    var v9 int64 = 9223372036854775807
    if (v9 >= 9223372036854775807) {
        v4 = (^(v4))
    }
    v7 = v7
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 uint8 = 1
        _ = fi10
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
