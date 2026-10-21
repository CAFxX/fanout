package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int16 }
func genFn0(p0 string, p1 uint32) uint32 {
    p1 += (p1 % (((p1 / (((p1 | p1)) | 1))) | 1))
    fmt.Println(p0)
    switch -2147483648 {
    case 42:
        var v7 uint32 = (p1 + p1)
        _ = v7
    default:
        v8 := []int{}
        _ = v8
    }
    _ = p0
    _ = p1
    return (p1 + (p1 >> ((42) & 63)))
}
func genFn1(p0 uint32, p1 int8) int16 {
    fmt.Println(p0, p0)
    var v9 int32 = -2147483648
    switch 4294967295 {
    case -2147483648:
        if true {
            v10 := GenS1{F0: p1, F1: 128, F2: -32768}
            var v11 float64 = ((3.14159 + (1.5 * 0.0)) + 1e10)
            _ = v10
            _ = v11
        } else {
            p1 = p1
        }
    case 128:
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 bool = true
            switch "hello" {
            case "a":
                var v14 uint = 237
                _ = v14
            default:
                fmt.Println(v9)
            }
            var v15 uint16 = 200
            _ = fi12
            _ = v13
            _ = v15
        }
    case 42:
        v16 := []int{2147483647, -1, -93, -1}
        _ = v16
    }
    var v17 string = ("hello" + "a")
    var v18 int8 = (p1 & 0)
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    _ = v9
    return 42
}

func main() {
    var v1 string = ""
    v2 := []string{}
    v3 := []uint32{}
    var v4 int16 = -32768
    var v5 uint16 = 1
    var v6 int64 = 42
    v6 = v6
    var v19 float32 = 1e10
    if (-2.25 > 3.14159) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch "" {
        case "hello":
            if (!(-1 > -1)) {
                v4 = v4
            }
        case "\n\t":
            if false {
                fmt.Println(v6, v5, v1)
                _ = len(v3)
                v20 := GenS1{F0: 0, F1: 170, F2: v4}
                _ = v20
            } else {
                v2 = append(v2, "\n\t")
                var v21 uint8 = 1
                _ = v21
            }
        case "":
            switch 961 {
            case 0:
                switch -9223372036854775808 {
                case 501:
                    switch 32767 {
                    case -992:
                        v22 := GenS1{F0: -66, F1: 1, F2: (v4 + (v4 / (((v4 ^ (v4 | v4))) | 1)))}
                        _ = v22
                    }
                }
            case 77:
                v1 = ("foo bar" + v1)
            default:
                var v23 uint = 499
                _ = v23
            }
        }
        v5 = 2
    }
    var v24 uint64 = 18446744073709551615
    switch 32768 {
    case -1:
        v5 -= v5
    default:
        switch -2147483648 {
        case 1:
            for fi25 := 0; fi25 < 19; fi25++ {
                v1 = v1
                _ = fi25
            }
        default:
            var v26 uint8 = 56
            _ = v26
        }
    }
    var v27 uint32 = genFn0(v1, 2)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v28 uint8 = 0
    v28 -= v28
    v28 += (^(v28))
    var v29 uint8 = 42
    if len(v3) > 0 {
        v3[0] = genFn0(v1, v27)
    }
    if len(v2) > 0 {
        v2[0] = (("" + "a") + v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
