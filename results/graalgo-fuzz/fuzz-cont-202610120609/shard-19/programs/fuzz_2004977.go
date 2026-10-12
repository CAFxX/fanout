package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int8; F1 int; F2 uint }
func genFn0() int {
    for fi5 := 0; fi5 < 7; fi5++ {
        if (0 <= 482) {
            switch (fi5 - fi5) {
            case 32767:
                switch 128 {
                case 1:
                    var v6 uint16 = 493
                    _ = v6
                case 127:
                    v7 := []int{}
                    _ = v7
                }
            case 127:
                v8 := []int{fi5, 127, 314}
                _ = v8
            default:
                var v9 float64 = 0.0
                _ = v9
            }
        } else {
            for fi10 := 0; fi10 < 12; fi10++ {
                var v11 uint = 255
                _ = fi10
                _ = v11
            }
            var v12 uint32 = 2
            _ = v12
        }
        var v13 string = ""
        if ((1e10 + 0.0) == (3.14159 / -2.25)) {
            var v14 uint16 = 65535
            _ = v14
        }
        _ = fi5
        _ = v13
    }
    var v15 uint8 = 0
    _ = v15
    return 0
}

func main() {
    var v1 float64 = 0.0
    var v2 uint = 18446744073709551615
    v3 := make(map[string]string)
    var v4 int16 = 0
    var v16 uint16 = 2
    _ = len(v3)
    fmt.Println(v4)
    for fi17 := 0; fi17 < 17; fi17++ {
        v18 := GenS1{F0: "", F1: (v2 ^ 255)}
        if false {
            var v19 int8 = -128
            _ = v19
        }
        var v20 float32 = (float32((0.1 * (v1 / float64(1e10)))) * ((1.5 - (0.0 - -2.25)) * ((0.0 + 1e10) * float32(v1))))
        _ = fi17
        _ = v18
        _ = v20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 uint = 566
        _ = fi21
        _ = v22
    }
    if false {
        fmt.Println(v1)
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 string = "a"
        v16 = (v16 / (((v16 >> ((v2) & 63))) | 1))
        _ = fi23
        _ = v24
    }
    v2 = (v2 - 1)
    v2 += v2
    for fi25 := 0; fi25 < 6; fi25++ {
        switch "" {
        case "\n\t":
            fmt.Println(v1, fi25)
        case "hello":
            v26 := GenS2{F0: 1, F1: fi25, F2: 18446744073709551615}
            _ = v26
        case "foo bar":
            switch (fi25 * genFn0()) {
            case 4294967295:
                v16 = 255
            }
        default:
            var v27 uint = (v2 % (((v2 * (v2 >> ((v2) & 63)))) | 1))
            _ = v27
        }
        _ = fi25
    }
    v1 = v1
    if (18446744073709551615 == 0) {
        v1 = (v1 / 3.14159)
        v2 = v2
    }
    v4 += (v4 - (v4 * v4))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
