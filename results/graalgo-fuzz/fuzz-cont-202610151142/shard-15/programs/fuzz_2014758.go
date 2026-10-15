package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]bool)
    var v3 uint32 = 658
    var v4 float32 = float32(0.0)
    v5 := []uint32{v3, v3, (v3 * 148), 2}
    v3 = 4294967295
    v6 := GenS1{F0: (^(v3)), F1: 0}
    fmt.Println(v4)
    var v7 uint64 = 1
    if (true || false) {
        if (false || true) {
            for fi8 := 0; fi8 < 19; fi8++ {
                if (42 <= 1) {
                    if true {
                        v4 -= v4
                        v7 -= v7
                    }
                } else {
                    for fi9 := 0; fi9 < 2; fi9++ {
                        if false {
                            var v10 float32 = 0.0
                            _ = v10
                        }
                        if (false && ("foo bar" >= ("hello" + ("\n\t" + "\n\t")))) {
                            var v11 float32 = v4
                            fmt.Println(v7, v4, fi9)
                            v11 = v11
                            _ = v11
                        } else {
                            var v12 float32 = 1.5
                            _ = v12
                        }
                        v3 = 4294967295
                        _ = fi9
                    }
                    var v13 uint = 18446744073709551615
                    v2[(fi8 + -582)] = (false || (-1 > 52))
                    _ = v13
                }
                switch "\n\t" {
                case "":
                    if false {
                        fmt.Println(v4, v7)
                    }
                case "foo bar":
                    if len(v5) > 0 {
                        v5[0] = uint32(v7)
                    }
                case "\n\t":
                    var v14 uint32 = v3
                    _ = v14
                default:
                    var v15 uint16 = 417
                    _ = v15
                }
                _ = fi8
            }
            var v16 uint32 = 2147483648
            v17 := GenS1{F0: 42, F1: v7}
            _ = v16
            _ = v17
        } else {
            var v18 bool = (1 == 65535)
            v18 = v18
            _ = v18
        }
        v7 += 18446744073709551615
    } else {
        if true {
            v7 -= (v7 & 548)
            delete(v2, -2147483648)
        }
        _ = len(v5)
    }
    switch ("hello" + ("a" + "")) {
    case "":
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 uint8 = 1
            _ = fi19
            _ = v20
        }
    case "a":
        _ = len(v2)
    }
    v21 := GenS2{F0: 42}
    v7 += (v7 ^ v7)
    v3 -= 4294967295
    var v22 int = -866
    v4 = float32(3.14159)
    switch "\n\t" {
    case "hello":
        fmt.Println(v3, v22, v22)
    case "\n\t":
        if ((v3 >> ((18446744073709551615) & 63)) != v3) {
            v3 = (v3 & v3)
            v3 = 1
        }
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
