package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
func genFn0(p0 string) int8 {
    switch 127 {
    case 1:
        p0 = p0
    case -9223372036854775808:
        var v6 int16 = 113
        _ = v6
    case 0:
        v7 := GenS1{F0: -2147483648}
        _ = v7
    default:
        switch p0 {
        case "foo bar":
            p0 = p0
        case "\n\t":
            if true {
                p0 = p0
            } else {
                var v8 bool = true
                switch -9223372036854775808 {
                case 127:
                    v8 = false
                case 2147483647:
                    var v9 float64 = -2.25
                    _ = v9
                default:
                    fmt.Println(v8, v8)
                }
                v8 = false
                _ = v8
            }
        case "a":
            v10 := GenS1{F0: -629}
            _ = v10
        default:
            for fi11 := 0; fi11 < 2; fi11++ {
                for fi12 := 0; fi12 < 1; fi12++ {
                    p0 = p0
                    if (((true || (371 >= 410)) || true) && (4294967295 >= 4294967295)) {
                        if ((!(true && true)) && true) {
                            fmt.Println(fi12, fi11, fi12)
                        } else {
                            p0 = (p0 + p0)
                            v13 := []int{(fi11 - 1), 737}
                            fmt.Println(p0)
                            _ = v13
                        }
                        p0 = p0
                        var v14 int32 = 0
                        fmt.Println(p0, p0)
                        p0 = p0
                        _ = v14
                    } else {
                        if (0 != 1) {
                            v15 := []int{2147483647, (fi12 * fi12), -175}
                            var v16 int32 = 1
                            _ = v15
                            _ = v16
                        }
                        var v17 uint = 42
                        if (4294967295 > 331) {
                            v18 := []int{fi11, (fi11 ^ 4294967295), fi12}
                            fmt.Println(v17)
                            _ = v18
                        }
                        _ = v17
                    }
                    fmt.Println(p0)
                    _ = fi12
                }
                p0 = p0
                if true {
                    p0 = p0
                    var v19 int8 = -103
                    p0 = (p0 + "hello")
                    _ = v19
                } else {
                    v20 := []int{}
                    p0 = p0
                    _ = v20
                }
                _ = fi11
            }
        }
    }
    fmt.Println(p0)
    v21 := GenS1{F0: 2147483647}
    _ = p0
    _ = v21
    return -23
}
func genFn1(p0 int32, p1 int16, p2 uint) int32 {
    v22 := []int{}
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := GenS1{F0: fi23}
        for fi25 := 0; fi25 < 2; fi25++ {
            p2 = 42
            _ = fi25
        }
        _ = fi23
        _ = v24
    }
    var v26 float32 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v26
    return (p0 >> (((p2 | 42)) & 63))
}

func main() {
    v1 := []int{-1, 32767, 32768, -920}
    var v2 uint16 = 0
    var v3 uint8 = 47
    var v4 int8 = -1
    v5 := make(map[string]uint32)
    v27 := GenS2{F0: 2147483647}
    if (true || false) {
        var v28 uint32 = 391
        v29 := GenS3{F0: false}
        _ = v28
        _ = v29
    }
    v3 = v3
    var v30 uint32 = 643
    var v31 float32 = float32(1e10)
    v31 += 3.14159
    switch "\n\t" {
    case "\n\t":
        v4 = (v4 << ((659) & 63))
    }
    var v32 int8 = (v4 - (-(v4)))
    var v33 string = "foo bar"
    if len(v1) > 0 {
        v1[0] = -2147483648
    }
    v34 := GenS1{F0: 601}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
