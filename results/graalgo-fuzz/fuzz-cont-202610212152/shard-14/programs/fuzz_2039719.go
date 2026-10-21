package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
func genFn0(p0 int32, p1 int32) uint64 {
    p1 += (p0 >> ((2) & 63))
    for fi6 := 0; fi6 < 13; fi6++ {
        if false {
            p1 = (p0 - (p0 ^ p1))
            var v7 uint32 = 664
            _ = v7
        }
        _ = fi6
    }
    if true {
        var v8 uint64 = 42
        fmt.Println(v8, p1)
        v9 := GenS1{F0: (v8 << ((0) & 63))}
        _ = v8
        _ = v9
    }
    if (65 >= 18446744073709551615) {
        v10 := []int{-512}
        _ = v10
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 18446744073709551615
}
func genFn1(p0 uint64) int32 {
    var v11 uint = 2
    v11 += v11
    v12 := GenS2{F0: 9223372036854775807}
    _ = p0
    _ = v11
    _ = v12
    return 42
}

func main() {
    var v1 bool = false
    v2 := []int{1, 0, 1, 32768}
    var v3 int64 = 42
    v4 := []int{}
    var v5 uint8 = 0
    var v13 int16 = -1
    if len(v4) > 0 {
        v4[0] = -1
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 32768 {
    case -2147483648:
        var v14 int64 = 0
        _ = v14
    case 226:
        var v15 float64 = (1.5 * (1.5 - float64(-2.25)))
        _ = v15
    default:
        if (("" + "foo bar") >= ("\n\t" + "")) {
            v2 = append(v2, 128)
        }
    }
    fmt.Println(v13)
    v16 := GenS1{F0: 2}
    fmt.Println(v3)
    v1 = v1
    v5 = 151
    fmt.Println(v5, v13)
    var v17 int64 = -804
    if v1 {
        var v18 int8 = -128
        _ = v18
    }
    switch int(v13) {
    case 908:
        if (42 < 127) {
            if len(v2) > 0 {
                v2[0] = 127
            }
            var v19 int16 = v13
            _ = v19
        } else {
            v17 = 302
            var v20 int = 1
            v13 -= (v13 >> ((255) & 63))
            _ = v20
        }
    case -184:
        switch ("a" + ("\n\t" + "a")) {
        case "foo bar":
            var v21 float64 = (-2.25 * -2.25)
            _ = v21
        case "a":
            fmt.Println(v1, v17, v5)
        case "":
            if false {
                var v22 float32 = 3.14159
                if false {
                    if len(v2) > 0 {
                        v2[0] = -881
                    }
                    if len(v4) > 0 {
                        v4[0] = 4294967295
                    }
                }
                var v23 int64 = 42
                _ = v22
                _ = v23
            }
        default:
            v1 = ((127 >= 127) || (!v1))
        }
    }
    v5 = (v5 ^ (v5 * (^(v5))))
    v13 += v13
    if (98 < 1) {
        _ = len(v4)
        var v24 float64 = 0.0
        if len(v4) > 0 {
            v4[0] = 32767
        }
        _ = v24
    }
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
