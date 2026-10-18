package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 string, p1 string) uint {
    var v5 uint32 = 2
    if false {
        var v6 float64 = 1e10
        _ = v6
    }
    p1 = "foo bar"
    p0 = p1
    if true {
        var v7 uint8 = 255
        v5 += 4294967295
        v8 := GenS1{F0: int8(v7), F1: v5, F2: 703}
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    return 2
}

func main() {
    var v1 int32 = 2147483647
    var v2 int8 = 127
    v3 := make(map[int]int)
    var v4 int8 = 0
    var v9 string = ("hello" + "a")
    switch 42 {
    case 566:
        switch v9 {
        case "foo bar":
            if false {
                if (42 < 0) {
                    if (false && true) {
                        var v10 int = 1
                        var v11 bool = (((((v9 + v9) + v9) + "") + (v9 + v9)) > "hello")
                        var v12 int64 = 9223372036854775807
                        _ = v10
                        _ = v11
                        _ = v12
                    } else {
                        var v13 int64 = -9223372036854775808
                        v14 := GenS1{F0: v2, F1: 42, F2: 42}
                        if true {
                            v15 := GenS2{F0: 1}
                            switch 843 {
                            case 127:
                                var v16 int32 = 2147483647
                                _ = v16
                            default:
                                v9 = v9
                            }
                            _ = v15
                        }
                        _ = v13
                        _ = v14
                    }
                    var v17 float32 = (((0.0 / 1.5) + 3.14159) * 3.14159)
                    _ = v17
                }
                fmt.Println(v1, v1, v1)
                fmt.Println(v2, v2, v4)
            }
        default:
            switch -2147483648 {
            case -9223372036854775808:
                if (true && (1e10 >= (0.1 + 3.14159))) {
                    var v18 uint16 = 65535
                    _ = v18
                }
            default:
                if (!true) {
                    v1 = 2147483647
                    v19 := GenS1{F0: (v4 - v4), F1: 2, F2: 32767}
                    _ = v19
                }
            }
        }
    default:
        _ = len(v3)
    }
    if ((128 != 2) && true) {
        v9 = v9
    }
    if true {
        delete(v3, 42)
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 uint64 = 18446744073709551615
            v21 = 18446744073709551615
            _ = fi20
            _ = v21
        }
        v1 -= (v1 * (^(v1)))
    } else {
        fmt.Println(v1, v9, v2)
        v22 := GenS1{F0: (v4 << ((0) & 63)), F1: 2147483648, F2: 2}
        _ = v22
    }
    v9 = ((v9 + (("" + (v9 + "foo bar")) + (v9 + ("a" + "\n\t")))) + v9)
    fmt.Println(v9)
    var v23 uint8 = 128
    var v24 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
