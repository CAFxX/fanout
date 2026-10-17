package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int16 }

func main() {
    v1 := []int{}
    var v2 int16 = -32768
    var v3 uint16 = 1
    v4 := make(map[uint32]string)
    for fi5 := 0; fi5 < 3; fi5++ {
        fmt.Println(fi5, v3)
        var v6 bool = false
        _ = fi5
        _ = v6
    }
    if false {
        v7 := GenS1{F0: 1}
        var v8 uint = 65535
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 0}
        if ((float32(1.5) * (1e10 - 1.5)) < 1e10) {
            v10 := GenS2{F0: 948, F1: v2}
            _ = v10
        }
        _ = v9
    }
    v3 += (v3 >> ((65535) & 63))
    for fi11 := 0; fi11 < 1; fi11++ {
        if (0 >= 0) {
            fmt.Println(fi11)
            for fi12 := 0; fi12 < 1; fi12++ {
                v13 := GenS2{F0: 32767, F1: 42}
                _ = fi12
                _ = v13
            }
            v2 = v2
        } else {
            var v14 uint = 0
            switch (^(fi11)) {
            case 127:
                _ = len(v4)
            case -142:
                switch -9223372036854775808 {
                case 4294967295:
                    v3 = (v3 * (v3 >> (((^(v14))) & 63)))
                default:
                    v1 = append(v1, (fi11 + (fi11 / ((fi11) | 1))))
                }
            }
            v15 := GenS2{F0: (v3 ^ v3), F1: v2}
            _ = v14
            _ = v15
        }
        v2 = (v2 | v2)
        _ = fi11
    }
    v2 = (v2 * (v2 & (-(v2))))
    fmt.Println(v2)
    switch 42 {
    case 128:
        _ = len(v4)
    case -1:
        var v16 int32 = -719
        _ = v16
    default:
        var v17 uint16 = 42
        _ = v17
    }
    fmt.Println(v3, v2)
    v3 += (^(v3))
    v18 := GenS1{F0: 255}
    for fi19 := 0; fi19 < 4; fi19++ {
        for fi20 := 0; fi20 < 17; fi20++ {
            delete(v4, 4294967295)
            switch -1 {
            case 0:
                v2 -= v2
            case 389:
                if (("\n\t" + "foo bar") <= "") {
                    var v21 int8 = 100
                    _ = v21
                }
            default:
                if true {
                    var v22 bool = true
                    _ = v22
                }
            }
            switch fi20 {
            case 32767:
                fmt.Println(fi20)
            }
            _ = fi20
        }
        v23 := GenS1{F0: 128}
        _ = fi19
        _ = v23
    }
    if (0 > 4294967295) {
        delete(v4, 914)
        v2 = (v2 + (v2 & (^(v2))))
        var v24 float64 = 3.14159
        _ = v24
    } else {
        v2 += (v2 << ((0) & 63))
        var v25 bool = ((float32(float64(-2.25)) * 1.5) != (((3.14159 * (0.0 * ((0.0 / 1.0) * float32(1.5)))) - -2.25) / float32(float64(1e10))))
        if v25 {
            v4[65535] = "a"
            for fi26 := 0; fi26 < 16; fi26++ {
                var v27 string = "hello"
                if (-354 <= 1) {
                    v27 = "\n\t"
                    var v28 float64 = ((1e10 / -2.25) - 0.1)
                    _ = v28
                }
                v29 := GenS2{F0: 42, F1: (v2 | v2)}
                _ = fi26
                _ = v27
                _ = v29
            }
            fmt.Println(v25, v3)
        }
        _ = v25
    }
    if (!(!((-317 > 1) || false))) {
        v30 := GenS1{F0: 127}
        v31 := GenS2{F0: 804, F1: v2}
        _ = v30
        _ = v31
    }
    var v32 int = 128
    fmt.Println(v3)
    var v33 uint8 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
