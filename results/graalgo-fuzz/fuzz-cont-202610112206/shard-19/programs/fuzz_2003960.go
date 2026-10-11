package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int32; F1 uint8; F2 int32 }
func genFn0(p0 uint8, p1 uint32, p2 int8) uint16 {
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(p2, p0, p0)
        p1 = (^(p1))
        v8 := GenS1{F0: (-3 < p2), F1: 0}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: ((p0 | p0) == p0), F1: uint(p1)}
    p0 -= (p0 * p0)
    var v10 int16 = -32768
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return uint16(p1)
}

func main() {
    var v1 int32 = 1
    var v2 string = ""
    var v3 uint8 = 2
    var v4 uint32 = 300
    var v5 float64 = ((-2.25 * float64(0.0)) / -2.25)
    var v6 uint8 = 178
    if true {
        v5 -= float64(0.1)
    }
    v1 -= 1
    var v12 int16 = -1
    v13 := []int{188, -1}
    var v14 uint = 0
    var v15 uint32 = v4
    var v16 uint8 = v3
    v17 := GenS2{F0: (v1 % (((v1 & -2147483648)) | 1)), F1: 1, F2: v1}
    if true {
        switch -9223372036854775808 {
        case 906:
            switch v2 {
            case "a":
                switch -1 {
                case 2147483647:
                    switch "a" {
                    case "":
                        var v18 bool = true
                        _ = v18
                    case "foo bar":
                        v3 -= (^(v16))
                    default:
                        fmt.Println(v16, v5, v2)
                    }
                }
            }
        case -1:
            var v19 int8 = 0
            _ = v19
        }
        var v20 bool = (v5 <= (v5 / ((v5 / 0.1) + ((float64(1e10) - (v5 * 1e10)) - float64(3.14159)))))
        switch v2 {
        case "\n\t":
            fmt.Println(v6, v15, v6)
        case "":
            var v21 int64 = -1
            _ = v21
        default:
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        }
        _ = v20
    } else {
        var v22 float32 = float32(v5)
        _ = v22
    }
    var v23 uint64 = uint64(v3)
    var v24 int32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
