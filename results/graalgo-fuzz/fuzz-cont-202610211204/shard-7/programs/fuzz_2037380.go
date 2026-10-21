package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int) int32 {
    var v6 int = (p0 >> ((776) & 63))
    if false {
        if (!(0 > 0)) {
            v6 = (v6 >> ((255) & 63))
            v7 := []int{}
            _ = v7
        } else {
            for fi8 := 0; fi8 < 4; fi8++ {
                for fi9 := 0; fi9 < 17; fi9++ {
                    var v10 int = fi8
                    v11 := []int{(v10 & p0)}
                    p0 = 0
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                switch 4294967295 {
                case -9223372036854775808:
                    var v12 int8 = 42
                    _ = v12
                case 1:
                    p0 += (fi8 - v6)
                case -457:
                    var v13 string = "hello"
                    _ = v13
                }
                _ = fi8
            }
            var v14 int8 = 0
            _ = v14
        }
        var v15 float64 = 0.0
        var v16 int32 = 2147483647
        _ = v15
        _ = v16
    } else {
        v17 := []int{v6, 127, (^(p0))}
        if false {
            v18 := GenS1{F0: (65535 == 42)}
            v19 := GenS1{F0: false}
            var v20 string = "\n\t"
            _ = v18
            _ = v19
            _ = v20
        }
        _ = v17
    }
    _ = p0
    _ = v6
    return 2147483647
}

func main() {
    v1 := []uint32{373, 13, 0, 190}
    var v2 uint32 = 2
    var v3 uint = 4294967295
    var v4 int8 = 1
    var v5 int16 = -32768
    var v21 uint16 = 42
    v5 = -32768
    if (!false) {
        v1 = append(v1, (v2 ^ 2147483648))
        v4 = (^(v4))
    } else {
        if len(v1) > 0 {
            v1[0] = 1
        }
        v21 += v21
    }
    fmt.Println(v2, v21, v4)
    if (true && (false && (157 > 127))) {
        var v22 float32 = 3.14159
        var v23 int = 128
        var v24 int32 = 0
        _ = v22
        _ = v23
        _ = v24
    }
    if len(v1) > 0 {
        v1[0] = 2147483648
    }
    var v25 uint8 = 42
    switch ("" + "") {
    case "\n\t":
        var v26 float64 = 0.1
        _ = v26
    case "a":
        var v27 string = "foo bar"
        _ = v27
    case "":
        var v28 int = -840
        _ = v28
    default:
        v4 -= 1
    }
    if ((v5 ^ (v5 >> ((42) & 63))) >= 32767) {
        v3 = (v3 << ((34) & 63))
    }
    v25 += v25
    v21 = (^(v21))
    switch 127 {
    case 0:
        v1 = append(v1, uint32(v21))
    default:
        var v29 int64 = -9223372036854775808
        _ = v29
    }
    var v30 uint32 = (^(v2))
    v31 := GenS1{F0: (("\n\t" <= ("a" + ("" + "a"))) || true)}
    var v32 int16 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
