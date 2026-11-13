package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int64 }
func genFn0() uint8 {
    switch "" {
    case "hello":
        v7 := GenS1{F0: 2147483648}
        _ = v7
    }
    v8 := []int{-257}
    if true {
        if true {
            v9 := GenS2{F0: 65535}
            _ = v9
        }
    } else {
        var v10 int8 = -128
        _ = v10
    }
    _ = v8
    return 225
}
func genFn1(p0 int16, p1 string, p2 int16) int {
    v11 := []int{-9223372036854775808, -587, 32768, 0}
    v12 := GenS1{F0: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return -1
}

func main() {
    var v1 int32 = -2147483648
    var v2 string = ("hello" + "foo bar")
    v3 := make(map[string]bool)
    var v4 uint8 = 0
    var v5 uint64 = 0
    var v6 float64 = ((-2.25 + float64(0.0)) + ((3.14159 * ((1.5 * ((0.0 / 1.5) + 0.1)) * 0.1)) + -2.25))
    v5 -= (^(v5))
    fmt.Println(v1, v5, v1)
    v5 -= (v5 << ((255) & 63))
    var v13 uint = 4294967295
    var v14 float64 = 3.14159
    v15 := []int{int(v5), 9223372036854775807, 9223372036854775807, -9223372036854775808}
    var v16 int8 = 1
    v16 = v16
    var v17 uint32 = 59
    var v18 uint32 = (v17 / ((uint32(v4)) | 1))
    v19 := GenS1{F0: v17}
    v20 := GenS3{F0: (v18 % (((v17 % (((v17 & v17)) | 1))) | 1)), F1: (v2 + v2)}
    switch v2 {
    case "\n\t":
        v3[(((("hello" + v2) + ((v2 + "a") + (v2 + v2))) + ((v2 + "\n\t") + v2)) + (v2 + ((v2 + "foo bar") + ("foo bar" + v2))))] = true
    case "":
        if (float32(1.5) > 0.1) {
            v14 = float64(((0.1 * 3.14159) + (1.5 / ((float32(float64(3.14159)) + -2.25) + (3.14159 / 3.14159)))))
            v13 += v13
        } else {
            for fi21 := 0; fi21 < 5; fi21++ {
                if false {
                    fmt.Println(fi21)
                } else {
                    if false {
                        var v22 int = fi21
                        switch "hello" {
                        case "a":
                            v17 = v18
                        case "":
                            v23 := GenS4{F0: -182}
                            _ = v23
                        }
                        _ = v22
                    }
                }
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = fi21
            }
        }
    default:
        if false {
            _ = len(v3)
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
        } else {
            v24 := GenS2{F0: 0}
            fmt.Println(v2)
            _ = v24
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
