package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int8; F2 int32 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 uint; F1 uint32; F2 uint64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 int32 = -2147483648
    if (0 != 214) {
        var v4 uint16 = 65535
        _ = v4
    } else {
        fmt.Println(v3)
        v5 := GenS1{F0: 0, F1: 127, F2: 2}
        _ = v5
    }
    v6 := GenS2{F0: -128}
    v2[(("foo bar" + "hello") + (("hello" + "a") + ""))] = ""
    v7 := GenS3{F0: "a", F1: 100, F2: 1}
    var v8 float32 = (3.14159 / 1.0)
    v9 := GenS1{F0: v3, F1: 73, F2: 0}
    v10 := []int{42, -2147483648}
    for fi11 := 0; fi11 < 10; fi11++ {
        if ((true && true) && false) {
            if (false && true) {
                var v12 bool = false
                var v13 int64 = 42
                _ = v12
                _ = v13
            } else {
                switch ("a" + "") {
                case "":
                    _ = len(v2)
                case "\n\t":
                    var v14 uint = 18446744073709551615
                    _ = v14
                default:
                    var v15 float64 = 3.14159
                    _ = v15
                }
                if len(v10) > 0 {
                    v10[0] = -2147483648
                }
                var v16 uint32 = 2147483648
                _ = v16
            }
        } else {
            for fi17 := 0; fi17 < 6; fi17++ {
                v18 := GenS4{F0: (18446744073709551615 == 607), F1: ("foo bar" + "a")}
                _ = fi17
                _ = v18
            }
            fmt.Println(v8)
            v8 += (((v8 + v8) - ((1e10 - 0.0) * v8)) / ((v8 - (float32((1e10 - (((0.1 / float64((float32(3.14159) - ((((float32(float64(float32(float64(1.5)))) - ((3.14159 + v8) / v8)) - v8) * (v8 - v8)) / float32(-2.25))))) / (0.1 / 1.0)) / -2.25))) - v8)) * v8))
        }
        v19 := GenS4{F0: true, F1: "\n\t"}
        if false {
            if (128 <= 128) {
                if ((("hello" + "hello") < "\n\t") && (-128 > 100)) {
                    delete(v2, "hello")
                    var v20 float64 = 0.0
                    v21 := GenS5{F0: 2, F1: 42, F2: 0}
                    _ = v20
                    _ = v21
                }
                fmt.Println(fi11, fi11, v3)
            }
        }
        _ = fi11
        _ = v19
    }
    fmt.Println(v8)
    fmt.Println(v3)
    var v22 bool = (32767 < 255)
    v23 := v2[("a" + "hello")]
    if ((v3 ^ v3) != v3) {
        v8 -= float32((0.1 / (1e10 * (((float64((v8 + v8)) + ((1e10 / 1e10) + 3.14159)) / float64(float32((1e10 * 1e10)))) - -2.25))))
    } else {
        fmt.Println(v8)
    }
    if (!true) {
        fmt.Println(v22, v3, v23)
        var v24 string = ("" + ("\n\t" + (((v23 + (v23 + v23)) + v23) + v23)))
        fmt.Println(v3, v22)
        _ = v24
    }
    v3 -= (^(v3))
    var v25 uint16 = 42
    var v26 int64 = 0
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
