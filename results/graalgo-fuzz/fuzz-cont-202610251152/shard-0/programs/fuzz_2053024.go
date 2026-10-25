package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 string) int64 {
    switch "hello" {
    case "hello":
        v6 := []int{-9223372036854775808, -829, 864}
        _ = v6
    case "a":
        v7 := []int{128, 32768}
        _ = v7
    default:
        p0 = ("" + p0)
    }
    switch "hello" {
    case "a":
        var v8 int64 = -9223372036854775808
        _ = v8
    default:
        var v9 float64 = (-2.25 + 0.1)
        _ = v9
    }
    var v10 int64 = 1
    fmt.Println(v10, p0)
    _ = p0
    _ = v10
    return v10
}

func main() {
    var v1 uint64 = 834
    var v2 uint = 109
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]int)
    var v5 uint32 = 2147483648
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := v4["foo bar"]
        _ = fi11
        _ = v12
    }
    v13 := []int{723, 46, -2147483648, 422}
    var v14 uint8 = 1
    fmt.Println(v5, v5)
    if (127 == 4294967295) {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        v5 -= uint32(v2)
        fmt.Println(v1, v1)
    } else {
        v15 := GenS1{F0: -1, F1: 1}
        v2 += 4294967295
        var v16 int = 716
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        v2 -= (v2 * 4294967295)
        v2 = 149
        _ = fi17
    }
    fmt.Println(v14)
    v14 += 134
    if true {
        fmt.Println(v2, v2)
    } else {
        v1 = (v1 >> (((v2 & uint(v1))) & 63))
        var v18 bool = true
        _ = v18
    }
    v13 = append(v13, 2147483647)
    switch ((("a" + "a") + ((("" + "hello") + "a") + "foo bar")) + "hello") {
    case "a":
        for fi19 := 0; fi19 < 16; fi19++ {
            var v20 int = 4294967295
            for fi21 := 0; fi21 < 19; fi21++ {
                var v22 uint16 = 336
                if ((!(((v22 & (v22 >> (((v2 ^ v2)) & 63))) <= 0) && false)) && (false || (-1 < -128))) {
                    var v23 int = fi21
                    v24 := GenS1{F0: -915, F1: v20}
                    _ = v23
                    _ = v24
                } else {
                    if false {
                        if len(v13) > 0 {
                            v13[0] = fi19
                        }
                        var v25 int = v20
                        _ = v25
                    } else {
                        var v26 int8 = 0
                        var v27 string = ("foo bar" + (("foo bar" + "") + "a"))
                        var v28 float32 = 3.14159
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    var v29 int = (-(v20))
                    _ = v29
                }
                _ = fi21
                _ = v22
            }
            for fi30 := 0; fi30 < 15; fi30++ {
                v14 += 2
                var v31 string = "\n\t"
                v32 := GenS2{F0: (v14 << ((v2) & 63)), F1: (v1 << ((18446744073709551615) & 63))}
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi19
            _ = v20
        }
    case "hello":
        v5 -= (v5 & uint32(v2))
    default:
        delete(v4, "a")
    }
    v2 += v2
    var v33 uint16 = 65535
    var v34 uint64 = v1
    v35 := GenS3{F0: 1}
    var v36 bool = (0.1 != float64((1e10 - 0.0)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
