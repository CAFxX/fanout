package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }

func main() {
    var v1 float64 = float64(1.5)
    v2 := make(map[int]string)
    v3 := []int{-449}
    v1 -= -2.25
    var v4 int64 = -675
    delete(v2, 0)
    v1 -= ((v1 / v1) / 1.5)
    v1 -= float64(3.14159)
    v1 = ((v1 / v1) - (0.1 / v1))
    var v5 int16 = 1000
    v2[32768] = (("a" + "") + "")
    var v6 uint = 615
    v4 += int64(v6)
    v4 = (v4 & (v4 | v4))
    if true {
        if len(v3) > 0 {
            v3[0] = 42
        }
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := v2[fi7]
            var v9 uint16 = 98
            v4 -= (-(v4))
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: int8(v6), F1: 0}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 9; fi11++ {
            for fi12 := 0; fi12 < 6; fi12++ {
                if len(v3) > 0 {
                    v3[0] = fi12
                }
                fmt.Println(v6, fi12, v4)
                v1 += v1
                _ = fi12
            }
            switch "foo bar" {
            case "":
                v13 := GenS1{F0: -10, F1: 0}
                _ = v13
            case "\n\t":
                v4 += (-(v4))
            }
            _ = fi11
        }
        v14 := GenS1{F0: 99, F1: 834}
        _ = v14
    }
    switch 897 {
    case -1:
        var v15 uint8 = 128
        _ = v15
    case -810:
        if (!(true && (0 != 42))) {
            var v16 int = 0
            var v17 uint64 = 1
            _ = len(v3)
            _ = v16
            _ = v17
        } else {
            var v18 float64 = 0.0
            _ = v18
        }
    case 1:
        fmt.Println(v6, v1)
    }
    fmt.Println(v1, v6)
    if false {
        for fi19 := 0; fi19 < 6; fi19++ {
            if len(v3) > 0 {
                v3[0] = 363
            }
            if true {
                var v20 float64 = float64((((float32(v1) * 0.0) + 1e10) - -2.25))
                _ = v20
            } else {
                for fi21 := 0; fi21 < 16; fi21++ {
                    var v22 float64 = 1e10
                    _ = fi21
                    _ = v22
                }
                v5 = (v5 << ((v6) & 63))
                if false {
                    var v23 uint16 = 32767
                    _ = v23
                }
            }
            _ = fi19
        }
    }
    v5 += -98
    var v24 uint16 = 840
    v4 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
