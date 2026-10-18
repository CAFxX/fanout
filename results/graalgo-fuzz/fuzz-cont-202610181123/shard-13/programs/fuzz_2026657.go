package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 string; F1 uint16 }
func genFn0(p0 uint16, p1 int64, p2 uint64) uint32 {
    if (p2 == (p2 / (((p2 << ((911) & 63))) | 1))) {
        if false {
            fmt.Println(p1, p0, p0)
            var v7 string = "foo bar"
            var v8 float32 = 0.0
            _ = v7
            _ = v8
        }
        if true {
            v9 := GenS1{F0: "foo bar", F1: (p1 | (p1 - (p1 / ((p1) | 1)))), F2: 107}
            _ = v9
        } else {
            var v10 uint8 = 0
            v10 -= v10
            _ = v10
        }
        p0 += (p0 - (p0 ^ (^(p0))))
    } else {
        fmt.Println(p1)
        var v11 float64 = (0.1 * -2.25)
        _ = v11
    }
    fmt.Println(p1, p0, p2)
    p0 = p0
    if true {
        v12 := GenS1{F0: (("foo bar" + "foo bar") + "a"), F1: (-(p1)), F2: 986}
        fmt.Println(p2, p2, p1)
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 uint = 4294967295
            _ = fi13
            _ = v14
        }
        _ = v12
    } else {
        v15 := GenS1{F0: ("\n\t" + ""), F1: 843, F2: 2}
        v16 := GenS2{F0: 1000, F1: 255}
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    var v1 uint = 1
    v2 := make(map[uint32]int)
    v3 := []string{"hello"}
    var v4 float32 = 0.1
    v5 := make(map[uint32]int)
    var v6 int64 = -9223372036854775808
    fmt.Println(v1, v6, v4)
    var v17 uint8 = 80
    for fi18 := 0; fi18 < 12; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            if true {
                _ = len(v3)
                delete(v2, 376)
                if true {
                    var v20 uint = v1
                    _ = v20
                }
            } else {
                v17 = (v17 & 1)
            }
            _ = fi19
        }
        var v21 uint8 = (v17 & v17)
        delete(v2, 802)
        _ = fi18
        _ = v21
    }
    v6 += 0
    switch 176 {
    case -9223372036854775808:
        var v22 int8 = 71
        _ = v22
    default:
        v2[700] = 0
    }
    fmt.Println(v1)
    fmt.Println(v1)
    if ((1 != -77) && false) {
        for fi23 := 0; fi23 < 3; fi23++ {
            v6 += v6
            switch "\n\t" {
            case "a":
                v4 = v4
            case "foo bar":
                fmt.Println(v17, v6, v4)
            }
            v5[0] = fi23
            _ = fi23
        }
    }
    if true {
        v24 := GenS3{F0: ("foo bar" + ""), F1: 0}
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
