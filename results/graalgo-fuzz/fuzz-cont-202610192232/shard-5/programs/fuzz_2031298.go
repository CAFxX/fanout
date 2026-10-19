package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int; F2 bool }
func genFn0(p0 int16, p1 uint, p2 int32) int {
    v7 := []int{4294967295}
    var v8 uint8 = 221
    var v9 int32 = (p2 / ((p2) | 1))
    var v10 uint16 = 0
    v11 := GenS1{F0: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 130
}

func main() {
    var v1 bool = false
    var v2 int = -2147483648
    v3 := []string{"foo bar", "hello", "foo bar", ("\n\t" + ("\n\t" + "hello"))}
    var v4 uint16 = 2
    v5 := []int{(v2 - v2), -9223372036854775808, (v2 / ((v2) | 1)), v2}
    var v6 uint64 = uint64(v4)
    v12 := GenS2{F0: 0, F1: v2, F2: v1}
    var v13 float64 = (-2.25 - 1.5)
    v2 = (v2 & v2)
    v13 += 0.1
    var v14 int32 = 0
    var v15 bool = (2 != 660)
    _ = len(v3)
    switch int(v6) {
    case -2147483648:
        fmt.Println(v14, v2, v15)
    case 32768:
        switch "\n\t" {
        case "\n\t":
            if v1 {
                for fi16 := 0; fi16 < 16; fi16++ {
                    v17 := GenS1{F0: 350}
                    v1 = (!(!(false || v1)))
                    if ("hello" >= "a") {
                        var v18 int16 = -32768
                        _ = v18
                    }
                    _ = fi16
                    _ = v17
                }
                for fi19 := 0; fi19 < 16; fi19++ {
                    if len(v5) > 0 {
                        v5[0] = v2
                    }
                    var v20 int32 = int32(v2)
                    _ = fi19
                    _ = v20
                }
                if true {
                    fmt.Println(v13, v15, v15)
                }
            }
        case "":
            v3 = append(v3, "a")
        default:
            v4 -= 42
        }
    case 128:
        if true {
            var v21 bool = v15
            _ = v21
        } else {
            var v22 int16 = -32768
            fmt.Println(v4, v13)
            _ = v22
        }
    default:
        var v23 int32 = (v14 * -2147483648)
        _ = v23
    }
    var v24 int64 = 0
    var v25 int16 = 32767
    var v26 uint64 = (^(v6))
    v25 += v25
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
