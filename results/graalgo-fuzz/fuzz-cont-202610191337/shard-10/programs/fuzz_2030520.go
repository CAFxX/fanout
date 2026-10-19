package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int16 }
func genFn0(p0 int16, p1 string, p2 int8) int {
    p1 = ((p1 + p1) + "")
    if false {
        p1 = (p1 + "\n\t")
    }
    v4 := GenS1{F0: "", F1: 42, F2: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -93
}
func genFn1(p0 int, p1 uint16) int32 {
    p1 = p1
    p0 = 42
    fmt.Println(p0)
    var v5 int8 = -1
    _ = p0
    _ = p1
    _ = v5
    return -2147483648
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 0
    var v3 uint = 4294967295
    v2 += (v2 % (((v2 & v2)) | 1))
    var v6 float64 = 0.0
    v7 := []int{32767, 0, 407}
    var v8 uint16 = 344
    v8 -= 65535
    var v9 uint8 = v2
    var v10 uint64 = 1
    for fi11 := 0; fi11 < 1; fi11++ {
        if (-1 >= 0) {
            if (v10 == (v10 >> ((v3) & 63))) {
                delete(v1, "hello")
                if (!true) {
                    fmt.Println(v6)
                }
            }
            var v12 int16 = -1
            _ = v12
        } else {
            fmt.Println(fi11)
            switch int(v9) {
            case 1:
                var v13 int16 = -32768
                _ = v13
            default:
                v3 = 0
            }
            if (1e10 == 0.1) {
                v14 := GenS1{F0: ("a" + "foo bar"), F1: 1, F2: -32768}
                for fi15 := 0; fi15 < 18; fi15++ {
                    var v16 uint8 = v9
                    if true {
                        var v17 uint8 = (v16 & (v16 - v2))
                        _ = v17
                    }
                    if false {
                        switch "foo bar" {
                        case "a":
                            fmt.Println(fi11)
                        case "foo bar":
                            var v18 string = (("\n\t" + "a") + "")
                            _ = v18
                        default:
                            v8 = v8
                        }
                        var v19 uint32 = 2147483648
                        var v20 int16 = -32768
                        _ = v19
                        _ = v20
                    }
                    fmt.Println(fi11, v6)
                    fmt.Println(v6)
                    _ = fi15
                    _ = v16
                }
                v6 -= v6
                _ = v14
            }
        }
        v9 = (v2 & v2)
        var v21 uint8 = (v9 << (((v3 << ((v3) & 63))) & 63))
        _ = fi11
        _ = v21
    }
    var v22 int8 = 0
    if (genFn1(-1, (v8 ^ (v8 ^ (v8 - (v8 | uint16(v10)))))) != -2147483648) {
        v6 = ((v6 * ((-2.25 / (v6 - v6)) * ((v6 / ((v6 - v6) * 0.1)) / float64(1.5)))) * ((v6 * v6) / v6))
        if true {
            _ = len(v7)
            if (!((v3 % ((uint(v8)) | 1)) == v3)) {
                fmt.Println(v22, v6)
            } else {
                v2 -= uint8(v3)
            }
        }
    }
    v23 := GenS1{F0: "foo bar", F1: -896, F2: 1000}
    var v24 int8 = v22
    if ((444 < int16(v2)) || ((42 == -1) && ("foo bar" == ""))) {
        fmt.Println(v3, v9)
        if (!true) {
            var v25 int32 = 2147483647
            _ = v25
        }
    } else {
        var v26 int = 32767
        _ = v26
    }
    var v27 float64 = v6
    if false {
        fmt.Println(v2)
        v24 = v22
        var v28 int = 289
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
