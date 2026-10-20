package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint64 }
func genFn0(p0 bool, p1 uint32, p2 int64) bool {
    var v7 uint32 = p1
    var v8 int32 = 699
    var v9 int8 = 1
    var v10 string = (("" + ((("a" + "foo bar") + "\n\t") + "a")) + ("hello" + ""))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (!(!false))
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]int)
    var v3 uint = 18446744073709551615
    var v4 uint32 = 844
    v5 := make(map[int]string)
    var v6 uint8 = 42
    switch 42 {
    case 2147483647:
        v11 := v5[387]
        _ = v11
    case 9223372036854775807:
        var v12 float64 = 3.14159
        _ = v12
    default:
        switch "foo bar" {
        case "a":
            var v13 float64 = ((0.0 + 3.14159) / -2.25)
            _ = v13
        case "foo bar":
            var v14 int16 = 32767
            _ = v14
        default:
            v1 = v1
        }
    }
    fmt.Println(v4)
    fmt.Println(v3)
    switch -1 {
    case -386:
        var v15 int32 = -2147483648
        _ = v15
    default:
        var v16 bool = (1000 < -269)
        _ = v16
    }
    fmt.Println(v3)
    if true {
        fmt.Println(v6, v1)
        v17 := []int{0, 32767, 9223372036854775807}
        var v18 uint8 = v6
        _ = v17
        _ = v18
    } else {
        v4 += v4
        if ((-100 != 127) || false) {
            var v19 uint = v3
            var v20 int8 = 127
            switch 1 {
            case 1:
                v6 = (v6 & 127)
            default:
                var v21 uint16 = 905
                _ = v21
            }
            _ = v19
            _ = v20
        }
        fmt.Println(v1)
    }
    v4 -= (v4 - uint32(v3))
    if (uint8(v1) <= (v6 % ((42) | 1))) {
        v22 := GenS1{F0: 628, F1: (v1 >> ((uint(v4)) & 63)), F2: (v1 | (v1 << ((v3) & 63)))}
        _ = v22
    }
    var v23 int16 = 1000
    v23 += (v23 & v23)
    var v24 int32 = 53
    if false {
        if ((v1 <= v1) || (!false)) {
            v25 := []int{-9223372036854775808, -486, 0}
            _ = v25
        } else {
            for fi26 := 0; fi26 < 9; fi26++ {
                var v27 int16 = (^(v23))
                v27 -= 42
                switch "\n\t" {
                case "foo bar":
                    delete(v2, (fi26 - (fi26 / (((-(fi26))) | 1))))
                case "":
                    switch "hello" {
                    case "":
                        if true {
                            var v28 int32 = int32(v27)
                            _ = v28
                        }
                    }
                }
                _ = fi26
                _ = v27
            }
            v1 -= v1
        }
        v23 -= (v23 & (^(v23)))
        v3 -= (v3 >> ((uint(v1)) & 63))
    } else {
        switch "" {
        case "hello":
            v29 := []int{1}
            _ = v29
        case "a":
            var v30 int32 = 568
            _ = v30
        default:
            v24 = v24
        }
    }
    v6 += v6
    var v31 uint32 = 1
    v23 += v23
    fmt.Println(v6, v24, v31)
    for fi32 := 0; fi32 < 7; fi32++ {
        v4 = 42
        _ = fi32
    }
    v33 := GenS1{F0: 468, F1: (v1 / ((18446744073709551615) | 1)), F2: v1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
