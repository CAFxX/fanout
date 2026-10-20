package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint }
type GenS2 struct { F0 int64; F1 uint64; F2 bool }

func main() {
    var v1 float32 = float32((1e10 * (1e10 - float64(((1e10 - 0.1) - 0.0)))))
    v2 := []string{}
    v3 := []int{-9223372036854775808, 9223372036854775807, 32767}
    var v4 float32 = float32((0.0 - 3.14159))
    if (-2147483648 >= 0) {
        v1 += (((1e10 - (v4 / 0.1)) - 1.5) + (-2.25 + float32(0.0)))
        v4 += ((-2.25 / (v1 / (v1 - (float32(3.14159) * (float32(1.5) * 3.14159))))) / v4)
        if false {
            var v5 int8 = 100
            v4 = (float32((0.1 - ((float64(1.5) + 3.14159) * ((0.1 / float64(v1)) - 1e10)))) + (-2.25 - v4))
            var v6 bool = (!false)
            _ = v5
            _ = v6
        } else {
            v7 := GenS1{F0: 124, F1: 42, F2: 1}
            _ = v7
        }
    } else {
        v1 -= v1
    }
    var v8 int64 = 42
    if true {
        fmt.Println(v1)
    }
    v1 -= v4
    v8 += (^(v8))
    if len(v2) > 0 {
        v2[0] = ""
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        if (true && true) {
            _ = len(v3)
            fmt.Println(v1, v1)
            switch ("\n\t" + "hello") {
            case "\n\t":
                fmt.Println(fi9, fi9)
            case "a":
                for fi10 := 0; fi10 < 7; fi10++ {
                    fmt.Println(v4, fi10, v1)
                    var v11 uint = 18446744073709551615
                    _ = fi10
                    _ = v11
                }
            }
        } else {
            for fi12 := 0; fi12 < 10; fi12++ {
                v13 := GenS2{F0: 1, F1: 4294967295, F2: (true || false)}
                var v14 bool = (625 >= (fi9 + 989))
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        switch (fi9 - fi9) {
        case 1:
            v15 := GenS1{F0: 76, F1: 2147483647, F2: 503}
            _ = v15
        case -399:
            v3 = append(v3, fi9)
        case 9223372036854775807:
            var v16 uint32 = 0
            _ = v16
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            switch fi17 {
            case 42:
                if (!true) {
                    var v18 uint16 = 590
                    var v19 int = 128
                    _ = v18
                    _ = v19
                }
            case -383:
                v8 = v8
            case 991:
                v4 -= 1e10
            }
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                fmt.Println(v1, v4, fi9)
                var v20 bool = true
                v4 -= 3.14159
                _ = v20
            }
            _ = fi17
        }
        _ = fi9
    }
    v1 = (v1 + (((v4 / v4) + (v1 * (v4 + (float32((((float64(v1) / 1.0) / 1.0) + (((3.14159 * float64(0.0)) * 0.1) + -2.25))) - v1)))) - 1.5))
    fmt.Println(v8)
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 uint64 = 18446744073709551615
        for fi23 := 0; fi23 < 18; fi23++ {
            v4 += (1.5 * v4)
            v24 := GenS1{F0: 0, F1: -2147483648, F2: 4294967295}
            v25 := GenS1{F0: 127, F1: 0, F2: 852}
            var v26 string = "hello"
            var v27 int16 = 42
            switch (^(fi23)) {
            case 32767:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case 0:
                v28 := GenS1{F0: 128, F1: -2147483648, F2: 0}
                _ = v28
            default:
                if false {
                    v29 := GenS2{F0: (v8 + (v8 * v8)), F1: v22, F2: ((v27 / (((-(v27))) | 1)) == 206)}
                    _ = v29
                }
            }
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
            _ = v27
        }
        var v30 uint8 = 42
        _ = fi21
        _ = v22
        _ = v30
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
