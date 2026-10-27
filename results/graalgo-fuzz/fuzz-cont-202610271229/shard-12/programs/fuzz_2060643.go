package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint16; F1 int32; F2 int8 }
type GenS3 struct { F0 uint8; F1 int32 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 1e10
    var v3 uint16 = 78
    var v4 string = "a"
    var v5 uint8 = 1
    v1[32767] = v4
    var v6 float32 = float32((1e10 - ((1e10 - -2.25) * 1.5)))
    if (!(((v4 + "foo bar") < (v4 + v4)) || true)) {
        switch 1 {
        case 127:
            switch -582 {
            case -460:
                var v7 int64 = 0
                _ = v7
            case 32767:
                for fi8 := 0; fi8 < 5; fi8++ {
                    delete(v1, (fi8 % ((2147483647) | 1)))
                    var v9 uint = 18446744073709551615
                    var v10 int32 = 2147483647
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            default:
                var v11 string = ("foo bar" + (v4 + v4))
                _ = v11
            }
        case 32767:
            switch -9223372036854775808 {
            case -2147483648:
                v12 := GenS1{F0: 280, F1: -32768}
                _ = v12
            case 9223372036854775807:
                for fi13 := 0; fi13 < 15; fi13++ {
                    v6 = ((1.5 * 1.5) / (1e10 / (0.0 * ((v2 + float32((3.14159 - 1e10))) * v6))))
                    _ = fi13
                }
            case 678:
                if (-2147483648 >= 42) {
                    fmt.Println(v6, v3, v5)
                    var v14 uint64 = 2
                    v6 += ((((v2 - v2) + (float32(3.14159) / v6)) - v2) + float32(float64(float32(float64(float32((1.5 * float64(((v6 / v2) * (v6 / (v2 - (((0.0 * -2.25) + ((((3.14159 / v2) - 0.0) * ((3.14159 - (v6 + 0.1)) * v6)) + 1.5)) * (v6 - 0.0)))))))))))))
                    _ = v14
                }
            default:
                if true {
                    v15 := []int{}
                    _ = v15
                } else {
                    v4 = ""
                    v6 = 1.5
                    _ = len(v1)
                }
            }
        case -9223372036854775808:
            switch "foo bar" {
            case "\n\t":
                if false {
                    v3 += 509
                    fmt.Println(v2, v2)
                    v2 -= v2
                }
            }
        default:
            switch "foo bar" {
            case "":
                if (false && (18446744073709551615 <= 0)) {
                    fmt.Println(v4, v5)
                    v2 -= (0.1 - v2)
                }
            case "a":
                switch "" {
                case "foo bar":
                    v16 := GenS1{F0: 18446744073709551615, F1: -170}
                    _ = v16
                case "hello":
                    for fi17 := 0; fi17 < 15; fi17++ {
                        var v18 float64 = (((0.0 + ((0.1 - 3.14159) - ((1.5 * 1.5) * float64(float32(1e10))))) - 0.0) + float64(v6))
                        fmt.Println(v6, v4)
                        fmt.Println(fi17, v18)
                        _ = fi17
                        _ = v18
                    }
                }
            }
        }
        v1[-2147483648] = ""
    } else {
        v19 := GenS1{F0: 2, F1: -1}
        _ = v19
    }
    var v20 uint64 = 18446744073709551615
    v21 := GenS2{F0: v3, F1: -1, F2: int8(v20)}
    for fi22 := 0; fi22 < 5; fi22++ {
        fmt.Println(v2, v4)
        var v23 uint64 = (^(v20))
        _ = fi22
        _ = v23
    }
    fmt.Println(v6)
    if false {
        var v24 bool = (-9223372036854775808 >= -917)
        v25 := GenS3{F0: 128, F1: -1}
        _ = v24
        _ = v25
    } else {
        v6 -= (1e10 * (1e10 - 0.0))
    }
    if (false || (true && (!((((v4 + "\n\t") + v4) != (("\n\t" + (((("foo bar" + v4) + "") + (((v4 + "\n\t") + "foo bar") + ((("hello" + "") + "") + "a"))) + v4)) + v4)) || (0 < -941))))) {
        if true {
            v20 = v20
        }
        var v26 uint = 1
        if false {
            v2 -= (((v2 - (1e10 - v6)) - float32(3.14159)) - v2)
            for fi27 := 0; fi27 < 12; fi27++ {
                if (v6 != (v6 / 3.14159)) {
                    fmt.Println(v6, v2)
                }
                _ = fi27
            }
            var v28 uint8 = (v5 - 87)
            _ = v28
        }
        _ = v26
    } else {
        var v29 int64 = 0
        fmt.Println(v3, v2, v20)
        _ = v29
    }
    v30 := []int{-1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
