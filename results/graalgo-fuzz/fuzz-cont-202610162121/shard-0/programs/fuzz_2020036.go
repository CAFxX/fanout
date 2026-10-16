package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 int64; F2 uint64 }
func genFn0(p0 bool, p1 int16, p2 string) int16 {
    var v5 int16 = p1
    p0 = p0
    v5 = (^(v5))
    v6 := GenS1{F0: 4294967295, F1: -9223372036854775808, F2: 138}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return v5
}
func genFn1(p0 int, p1 uint8) bool {
    var v7 uint = 960
    var v8 int = (p0 << (((v7 / ((2) | 1))) & 63))
    fmt.Println(p0, p1, v8)
    v9 := GenS1{F0: (p0 >> ((v7) & 63)), F1: -894, F2: (p1 << ((148) & 63))}
    for fi10 := 0; fi10 < 2; fi10++ {
        switch p0 {
        case 2147483647:
            fmt.Println(v7, fi10)
        }
        var v11 bool = false
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 string = "foo bar"
    var v2 float64 = 0.0
    v3 := make(map[string]string)
    var v4 uint32 = 65535
    v2 += (v2 / 1.0)
    var v12 uint8 = 127
    v13 := GenS2{F0: ("foo bar" + (((v1 + ("foo bar" + "foo bar")) + "foo bar") + v1))}
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 string = v1
        fmt.Println(v15, fi14, v2)
        _ = fi14
        _ = v15
    }
    v1 = ""
    if (-1 <= 2147483647) {
        _ = len(v3)
        v16 := GenS2{F0: "a"}
        _ = v16
    }
    var v17 float32 = 1.5
    v18 := []int{128}
    v18 = append(v18, 128)
    v17 -= float32((v2 + 1e10))
    if len(v18) > 0 {
        v18[0] = -762
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v20 := GenS1{F0: fi19, F1: 9223372036854775807, F2: 255}
        _ = fi19
        _ = v20
    }
    if len(v18) > 0 {
        v18[0] = -2147483648
    }
    var v21 float32 = (v17 / (-2.25 / 0.1))
    if false {
        var v22 float32 = v21
        if true {
            if false {
                v2 += 0.1
            } else {
                for fi23 := 0; fi23 < 3; fi23++ {
                    _ = len(v18)
                    _ = fi23
                }
            }
        }
        var v24 uint16 = 872
        _ = v22
        _ = v24
    } else {
        for fi25 := 0; fi25 < 16; fi25++ {
            v1 = v1
            if true {
                switch "" {
                case "foo bar":
                    for fi26 := 0; fi26 < 11; fi26++ {
                        fmt.Println(v12, v2)
                        var v27 string = v1
                        _ = fi26
                        _ = v27
                    }
                default:
                    var v28 float64 = (v2 + (v2 / 1.5))
                    _ = v28
                }
                v29 := GenS1{F0: (fi25 >> ((2) & 63)), F1: 42, F2: (v12 - v12)}
                var v30 uint16 = uint16(v12)
                _ = v29
                _ = v30
            } else {
                v31 := GenS3{F0: 32768, F1: 9223372036854775807, F2: 433}
                var v32 int = fi25
                _ = v31
                _ = v32
            }
            _ = fi25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
