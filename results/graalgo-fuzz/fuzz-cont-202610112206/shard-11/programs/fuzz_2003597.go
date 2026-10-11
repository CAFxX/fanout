package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 bool }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int8) uint16 {
    var v5 int8 = p0
    switch -918 {
    case 930:
        fmt.Println(v5)
    case -2147483648:
        var v6 float64 = 0.0
        _ = v6
    default:
        fmt.Println(v5, v5)
    }
    fmt.Println(v5, p0)
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint16 = 65535
        v9 := GenS1{F0: 0, F1: "\n\t", F2: (((v8 * v8) == 255) && ((!true) && (352 != 0)))}
        var v10 uint8 = 0
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 uint64 = 822
    _ = p0
    _ = v11
    _ = v5
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = (!true)
    var v3 string = ("" + (("\n\t" + "a") + "a"))
    var v4 string = v3
    v3 = v4
    v12 := []int{2147483647, -542, 1}
    v13 := GenS2{F0: 0}
    switch v4 {
    case "":
        if len(v12) > 0 {
            v12[0] = 127
        }
    }
    v14 := GenS1{F0: 798, F1: v4, F2: v2}
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS2{F0: 42}
        v17 := GenS1{F0: 1, F1: v4, F2: false}
        _ = fi15
        _ = v16
        _ = v17
    }
    if ((!v2) && (!(((1e10 - (0.0 - 3.14159)) - 1.5) <= (0.0 - float64((1.5 + 1e10)))))) {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        if (42 <= 649) {
            var v18 int32 = 0
            if v2 {
                if (42 < -55) {
                    for fi19 := 0; fi19 < 8; fi19++ {
                        var v20 uint16 = 1
                        fmt.Println(v18, v20)
                        var v21 uint = 0
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                switch -9223372036854775808 {
                case -659:
                    v12 = append(v12, 32768)
                case -2147483648:
                    var v22 float64 = 1.5
                    _ = v22
                case -593:
                    var v23 uint = 65535
                    _ = v23
                default:
                    v2 = (!(255 != 498))
                }
            } else {
                var v24 uint16 = 855
                var v25 float64 = ((-2.25 * 3.14159) + 1e10)
                v25 -= (v25 + (1e10 / v25))
                _ = v24
                _ = v25
            }
            switch v3 {
            case "\n\t":
                var v26 float64 = 1.5
                _ = v26
            case "foo bar":
                var v27 int16 = 638
                _ = v27
            case "a":
                for fi28 := 0; fi28 < 6; fi28++ {
                    v2 = v2
                    _ = fi28
                }
            }
            _ = v18
        }
    }
    if v2 {
        var v29 int64 = 221
        if false {
            delete(v1, -395)
            var v30 uint = 65535
            _ = v30
        }
        v29 = (-(v29))
        _ = v29
    } else {
        var v31 int32 = 0
        fmt.Println(v2)
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
