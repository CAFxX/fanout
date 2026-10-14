package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint32; F1 int16 }

func main() {
    var v1 int32 = -1
    v2 := []int{}
    v3 := make(map[string]string)
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 128
    v6 := make(map[uint32]int)
    var v7 float32 = (0.0 - (3.14159 / ((1.5 / (1.5 - float32((3.14159 + (-2.25 - 1e10))))) / (1e10 / 3.14159))))
    v7 -= 0.1
    var v8 uint8 = (v5 ^ v5)
    v8 = (v8 | 166)
    if (!true) {
        if true {
            switch ("" + "a") {
            case "\n\t":
                var v9 float32 = 0.0
                _ = v9
            }
            v7 += (((1e10 / 1e10) - float32(-2.25)) * v7)
        }
        if (!true) {
            v2 = append(v2, 32767)
        } else {
            if ((v7 - (1.5 / ((v7 - float32((0.1 + float64((float32((3.14159 * -2.25)) + ((1.5 + 1e10) * (float32(0.0) + -2.25))))))) + v7))) >= v7) {
                v10 := GenS1{F0: 255, F1: 127, F2: 32767}
                fmt.Println(v1, v1, v7)
                var v11 float32 = 1e10
                _ = v10
                _ = v11
            } else {
                for fi12 := 0; fi12 < 2; fi12++ {
                    var v13 uint16 = 49
                    if (true || false) {
                        v5 -= (v5 / ((v8) | 1))
                        if false {
                            v1 += v1
                            for fi14 := 0; fi14 < 13; fi14++ {
                                v7 += float32((((1.5 + ((((1e10 / 3.14159) + 1e10) + (1e10 / float64(((3.14159 + 1e10) + v7)))) + 3.14159)) + 0.0) + (1.5 - 0.0)))
                                v4[0] = 1
                                _ = fi14
                            }
                        }
                    }
                    v13 -= v13
                    _ = fi12
                    _ = v13
                }
                v7 += v7
            }
        }
    } else {
        v8 -= (v8 ^ 0)
        v15 := GenS2{F0: 2147483648, F1: 576}
        var v16 uint64 = 1
        _ = v15
        _ = v16
    }
    switch (("foo bar" + ("" + "foo bar")) + "a") {
    case "a":
        v7 -= ((v7 * 0.1) - v7)
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        if true {
            var v18 int64 = -1
            fmt.Println(v1, fi17)
            _ = v18
        }
        _ = fi17
    }
    if false {
        v5 = 255
        switch 4294967295 {
        case 2147483647:
            var v19 uint64 = 433
            _ = v19
        default:
            fmt.Println(v1, v1, v8)
        }
    } else {
        if (true || (!true)) {
            fmt.Println(v7)
        }
    }
    v7 -= (1.5 + v7)
    v8 -= (v8 & 1)
    var v20 float32 = 0.0
    if (((!(!(!(796 >= 1)))) || (0.1 < v7)) || true) {
        for fi21 := 0; fi21 < 8; fi21++ {
            fmt.Println(v7, v20, v5)
            _ = fi21
        }
        var v22 uint8 = v8
        _ = v22
    }
    switch ("a" + "hello") {
    case "\n\t":
        if (uint16(v8) == 2) {
            v7 = float32(3.14159)
        } else {
            v23 := GenS2{F0: 2, F1: -32768}
            _ = v23
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
