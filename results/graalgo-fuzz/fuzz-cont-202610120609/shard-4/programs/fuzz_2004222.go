package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int8 }
func genFn0(p0 uint32) int16 {
    var v7 uint32 = (p0 << ((18446744073709551615) & 63))
    if (true || (true || true)) {
        switch "a" {
        case "\n\t":
            v8 := []int{-2147483648, 42}
            _ = v8
        default:
            v9 := GenS1{F0: 128}
            _ = v9
        }
        if true {
            fmt.Println(p0)
            if true {
                v10 := GenS1{F0: 128}
                switch ("hello" + "a") {
                case "foo bar":
                    for fi11 := 0; fi11 < 10; fi11++ {
                        if ((-2.25 / (3.14159 + -2.25)) == ((0.1 - ((1e10 * 1.5) - -2.25)) * -2.25)) {
                            p0 -= v7
                            for fi12 := 0; fi12 < 20; fi12++ {
                                if (0.1 > (3.14159 * 0.0)) {
                                    var v13 int16 = 0
                                    var v14 float32 = (-2.25 / 1.0)
                                    fmt.Println(v14, p0)
                                    _ = v13
                                    _ = v14
                                } else {
                                    v15 := []int{}
                                    _ = v15
                                }
                                v16 := []int{(fi11 % (((fi12 >> ((4294967295) & 63))) | 1))}
                                if (false && (-2.25 >= 1.5)) {
                                    fmt.Println(fi12)
                                    if (147 >= 804) {
                                        var v17 float32 = 1e10
                                        fmt.Println(p0, v7)
                                        v18 := GenS1{F0: 127}
                                        _ = v17
                                        _ = v18
                                    } else {
                                        var v19 uint8 = 1
                                        v7 += v7
                                        p0 += 495
                                        _ = v19
                                    }
                                } else {
                                    var v20 int8 = -71
                                    _ = v20
                                }
                                _ = fi12
                                _ = v16
                            }
                        } else {
                            fmt.Println(v7)
                        }
                        if false {
                            v7 += p0
                            var v21 int16 = -1
                            v22 := GenS2{F0: 32767, F1: int8(v21)}
                            _ = v21
                            _ = v22
                        }
                        _ = fi11
                    }
                case "":
                    for fi23 := 0; fi23 < 2; fi23++ {
                        v24 := GenS2{F0: 148, F1: -128}
                        var v25 uint = 4294967295
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                }
                var v26 int8 = -128
                _ = v10
                _ = v26
            }
        } else {
            fmt.Println(v7)
        }
        v27 := GenS1{F0: 162}
        _ = v27
    }
    var v28 int32 = 444
    var v29 float64 = (0.0 + float64(1e10))
    var v30 uint32 = p0
    _ = p0
    _ = v28
    _ = v29
    _ = v30
    _ = v7
    return 1
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[int]string)
    var v3 bool = (378 != 255)
    v4 := make(map[string]int)
    var v5 float64 = (((v1 * 0.1) * v1) - v1)
    var v6 string = ""
    v6 = v6
    var v31 int64 = 368
    var v32 string = v6
    v33 := GenS2{F0: 686, F1: 1}
    v32 = (v32 + v32)
    v1 -= v1
    _ = len(v4)
    for fi34 := 0; fi34 < 16; fi34++ {
        v1 = v1
        v1 += 1.5
        _ = fi34
    }
    v35 := []int{1, -1, -9223372036854775808, 42}
    for fi36 := 0; fi36 < 16; fi36++ {
        var v37 bool = v3
        var v38 string = ("a" + ("\n\t" + v32))
        fmt.Println(v38, v3, v31)
        _ = fi36
        _ = v37
        _ = v38
    }
    var v39 float32 = (0.0 + ((0.1 / 1.0) + ((((0.0 * (0.1 + 0.0)) * float32((((1.5 + v5) - v1) * 0.1))) / 1.5) / 1e10)))
    v32 = v32
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
