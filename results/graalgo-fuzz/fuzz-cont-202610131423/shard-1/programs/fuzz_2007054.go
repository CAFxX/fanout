package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint64 }
func genFn0() string {
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 string = ((("hello" + "hello") + "foo bar") + "foo bar")
        _ = fi5
        _ = v6
    }
    var v7 int = 4294967295
    v8 := GenS1{F0: 65535}
    switch 42 {
    case -30:
        for fi9 := 0; fi9 < 10; fi9++ {
            v7 = (fi9 / (((v7 - 982)) | 1))
            _ = fi9
        }
    case 0:
        fmt.Println(v7)
    case 127:
        fmt.Println(v7)
    default:
        var v10 float64 = 1.5
        _ = v10
    }
    _ = v7
    _ = v8
    return "\n\t"
}

func main() {
    v1 := []int{}
    v2 := make(map[int]string)
    var v3 float32 = 0.0
    var v4 bool = true
    if v4 {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 int32 = 1
            if (535 == 0) {
                var v13 uint8 = 2
                fmt.Println(fi11, v4)
                _ = v13
            } else {
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    switch (("foo bar" + (("hello" + "foo bar") + "")) + "a") {
                    case "\n\t":
                        v4 = v4
                    default:
                        var v14 uint8 = 2
                        _ = v14
                    }
                }
            }
            fmt.Println(v12)
            _ = fi11
            _ = v12
        }
    }
    v4 = (float32(0.0) == (1.5 + v3))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    v15 := GenS1{F0: 396}
    v16 := GenS1{F0: 137}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if v4 {
        v17 := GenS2{F0: -9223372036854775808, F1: 0, F2: 662}
        fmt.Println(v3, v4)
        _ = v17
    }
    v3 -= ((float32(((((((((0.0 - 1e10) + float64((float32(float64(1.5)) * v3))) / (1.5 - 1e10)) * 1e10) - (1.5 + 1e10)) + (0.1 - ((0.1 * (3.14159 * float64((((float32(((3.14159 - ((1e10 - (float64(v3) + float64(((0.0 / v3) * v3)))) / 1e10)) - -2.25)) * v3) + (v3 - v3)) / v3)))) - (float64((v3 + float32(-2.25))) * ((1.5 / 1.0) * 1.5))))) * 1.5) * ((-2.25 + 1.5) - float64(v3)))) * float32(float64((v3 * (v3 / v3))))) - v3)
    v4 = (13 == 0)
    var v18 uint32 = 0
    if v4 {
        v3 += float32(float64((-2.25 * (v3 / v3))))
        _ = len(v2)
        switch 32768 {
        case 32768:
            if ((!false) && ((251 > 0) || true)) {
                if (42 < v18) {
                    if (32767 == -1) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v19 int8 = -18
                        _ = v19
                    }
                    var v20 int = 128
                    if (995 < -458) {
                        v21 := GenS3{F0: 42}
                        v22 := GenS2{F0: 0, F1: 494, F2: 0}
                        _ = v21
                        _ = v22
                    }
                    _ = v20
                }
            }
        case 1000:
            v4 = false
        case 42:
            switch -9223372036854775808 {
            case 42:
                var v23 int32 = 42
                _ = v23
            case -1:
                for fi24 := 0; fi24 < 11; fi24++ {
                    if ((uint(v18) > 255) && (!false)) {
                        fmt.Println(v18, fi24, v3)
                    } else {
                        v3 = (v3 * 1.5)
                        delete(v2, (fi24 - (^(fi24))))
                        v1 = append(v1, -2147483648)
                    }
                    _ = fi24
                }
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
