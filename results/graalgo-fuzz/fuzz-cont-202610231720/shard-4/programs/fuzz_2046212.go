package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint16; F1 int16 }

func main() {
    var v1 string = ""
    v2 := make(map[string]uint32)
    var v3 uint8 = 255
    var v4 int8 = 91
    var v5 int8 = v4
    var v6 float32 = ((1e10 * ((((0.1 - 3.14159) / (0.0 - (-2.25 * 1.5))) + 0.1) * 1e10)) - 1.5)
    if true {
        v6 += (float32(3.14159) / -2.25)
    }
    if false {
        for fi7 := 0; fi7 < 3; fi7++ {
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 int32 = 0
                var v10 uint16 = 960
                var v11 int16 = 32767
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            _ = fi7
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            v6 = ((-2.25 + (float32(float64(float32((((3.14159 - -2.25) - ((-2.25 / 0.1) - ((1.5 * (float64(v6) + 0.0)) * 0.1))) * 0.0)))) - float32((1e10 / 0.1)))) + v6)
            _ = fi12
        }
        switch v1 {
        case "\n\t":
            for fi13 := 0; fi13 < 15; fi13++ {
                if (uint32(v3) != 65535) {
                    v4 = (-(v5))
                    v14 := GenS1{F0: -1, F1: 32767, F2: fi13}
                    _ = v14
                }
                switch (v1 + v1) {
                case "a":
                    if true {
                        if ("a" >= ((v1 + "a") + v1)) {
                            var v15 uint64 = 4294967295
                            v16 := []int{fi13, (fi13 ^ fi13)}
                            fmt.Println(v3, v3, v5)
                            _ = v15
                            _ = v16
                        }
                        v17 := GenS1{F0: 1, F1: 42, F2: (fi13 - (-(fi13)))}
                        fmt.Println(v5, v5)
                        _ = v17
                    }
                }
                if (false && (0 == 42)) {
                    v5 = -128
                    for fi18 := 0; fi18 < 3; fi18++ {
                        var v19 bool = ((v6 * (v6 * (v6 - (0.1 - v6)))) == (v6 - v6))
                        v20 := GenS2{F0: 28, F1: (false && ((-1 < -1) && (v19 || (!v19))))}
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    v21 := GenS2{F0: -162, F1: false}
                    _ = v21
                }
                _ = fi13
            }
        case "a":
            v5 -= 0
        case "foo bar":
            fmt.Println(v6)
        default:
            v3 -= v3
        }
    }
    if true {
        fmt.Println(v1)
    }
    if true {
        var v22 int16 = -1
        if (!(836 >= 32767)) {
            v1 = v1
        }
        _ = v22
    }
    fmt.Println(v6)
    for fi23 := 0; fi23 < 9; fi23++ {
        v4 = (v4 % ((v4) | 1))
        _ = fi23
    }
    v24 := []int{-814, 1, 2147483647}
    v6 = (1.5 + (v6 * 3.14159))
    fmt.Println(v1, v6, v5)
    if false {
        v3 = (v3 - 109)
        v25 := GenS3{F0: 255, F1: -952}
        _ = v25
    } else {
        for fi26 := 0; fi26 < 12; fi26++ {
            v6 = float32(-2.25)
            _ = fi26
        }
    }
    switch -2147483648 {
    case -978:
        if (((v6 + (v6 * (0.1 / float32(3.14159)))) + ((v6 - v6) / (v6 + v6))) >= v6) {
            v4 -= 100
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
        }
    case -678:
        v27 := v2[v1]
        _ = v27
    default:
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
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
