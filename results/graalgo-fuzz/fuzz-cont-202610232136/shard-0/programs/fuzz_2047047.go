package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 string) bool {
    p0 = "a"
    fmt.Println(p0)
    var v6 string = (p0 + p0)
    _ = p0
    _ = v6
    return false
}

func main() {
    var v1 int16 = 42
    var v2 int32 = 0
    var v3 bool = false
    var v4 string = "a"
    var v5 int8 = -80
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: (fi7 >> ((18446744073709551615) & 63)), F1: (v1 % (((v1 | v1)) | 1)), F2: 848}
        v9 := GenS1{F0: (fi7 + int(v1)), F1: (v1 % (((v1 % ((v1) | 1))) | 1)), F2: 42}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: int(v5), F1: v1, F2: 42}
    v5 = v5
    var v11 int = 127
    v11 += v11
    var v12 float64 = 1e10
    var v13 int = v11
    if false {
        v14 := GenS2{F0: 0}
        v3 = (0 < 32767)
        _ = v14
    }
    switch -2147483648 {
    case 0:
        v13 = (v13 ^ (v11 | 1))
    case -2147483648:
        var v15 uint = 0
        _ = v15
    default:
        switch v4 {
        case "":
            v16 := []int{140, (v13 >> ((1) & 63)), (v11 >> ((0) & 63)), v11}
            _ = v16
        case "foo bar":
            switch -20 {
            case 1:
                if (v3 || (v1 < (v1 >> ((520) & 63)))) {
                    var v17 int64 = -9223372036854775808
                    _ = v17
                } else {
                    switch "hello" {
                    case "a":
                        v18 := []int{}
                        _ = v18
                    default:
                        for fi19 := 0; fi19 < 17; fi19++ {
                            var v20 uint32 = 356
                            _ = fi19
                            _ = v20
                        }
                    }
                    var v21 uint8 = 127
                    var v22 uint32 = 42
                    _ = v21
                    _ = v22
                }
            default:
                v5 += int8(v1)
            }
        }
    }
    var v23 int = 32767
    var v24 uint16 = 255
    fmt.Println(v2)
    v25 := []int{v13, (v23 << ((65535) & 63)), 9223372036854775807, (v23 - v13)}
    var v26 float64 = v12
    v27 := GenS1{F0: (v13 + 32768), F1: (v1 * v1), F2: 0}
    v28 := GenS1{F0: v11, F1: (^(v1)), F2: v24}
    var v29 int32 = 2147483647
    fmt.Println(v3, v1, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
