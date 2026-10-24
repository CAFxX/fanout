package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 int32; F1 uint }

func main() {
    var v1 bool = true
    v2 := make(map[int]int)
    var v3 int8 = -72
    if v1 {
        var v4 int64 = 42
        _ = v4
    }
    var v5 int64 = -9223372036854775808
    for fi6 := 0; fi6 < 1; fi6++ {
        v3 -= (v3 * v3)
        if (3.14159 <= float64(-2.25)) {
            var v7 float64 = (0.1 + 0.0)
            _ = v7
        }
        fmt.Println(v5, v3)
        _ = fi6
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        switch "a" {
        case "\n\t":
            if v1 {
                switch fi8 {
                case 1:
                    var v9 bool = false
                    _ = v9
                default:
                    for fi10 := 0; fi10 < 11; fi10++ {
                        fmt.Println(fi10, v3)
                        if (v3 != v3) {
                            v1 = v1
                            var v11 uint32 = 42
                            _ = v11
                        }
                        if ((0.0 == float32(1.5)) && (1 != 18446744073709551615)) {
                            v12 := GenS1{F0: 1, F1: "hello"}
                            _ = v12
                        }
                        _ = fi10
                    }
                }
                if false {
                    var v13 int64 = v5
                    v14 := []int{(fi8 ^ fi8), fi8, -9223372036854775808, fi8}
                    var v15 float64 = (float64(3.14159) + 1e10)
                    _ = v13
                    _ = v14
                    _ = v15
                } else {
                    if false {
                        fmt.Println(fi8)
                        for fi16 := 0; fi16 < 8; fi16++ {
                            var v17 int64 = -664
                            v18 := GenS1{F0: 672, F1: (("\n\t" + ("" + "\n\t")) + ("a" + ("a" + "a")))}
                            if ("foo bar" == "a") {
                                var v19 float32 = (1.5 - (0.0 - float32((1.5 * 0.0))))
                                _ = len(v2)
                                _ = v19
                            }
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        v3 = 42
                    } else {
                        v20 := GenS1{F0: 0, F1: ((("foo bar" + ("hello" + "a")) + "") + "foo bar")}
                        _ = v20
                    }
                    v3 = v3
                }
            } else {
                v3 += v3
                v3 = v3
                v5 += (-(v5))
            }
        case "a":
            fmt.Println(v5, v3)
        }
        var v21 bool = v1
        _ = fi8
        _ = v21
    }
    v22 := v2[9223372036854775807]
    fmt.Println(v3, v22)
    v22 = v22
    if v1 {
        switch ("hello" + "hello") {
        case "a":
            if v1 {
                for fi23 := 0; fi23 < 7; fi23++ {
                    v1 = true
                    _ = fi23
                }
            }
        case "":
            if v1 {
                for fi24 := 0; fi24 < 4; fi24++ {
                    if v1 {
                        v3 += v3
                    } else {
                        fmt.Println(fi24, v3)
                        v3 += (v3 >> ((745) & 63))
                        v5 += v5
                    }
                    if v1 {
                        v25 := GenS2{F0: 65535}
                        v26 := GenS3{F0: 32767}
                        if v1 {
                            v27 := GenS4{F0: 174, F1: 1}
                            v28 := GenS1{F0: 0, F1: "\n\t"}
                            var v29 int = (-(fi24))
                            _ = v27
                            _ = v28
                            _ = v29
                        } else {
                            fmt.Println(v3, v3, v3)
                            var v30 bool = false
                            var v31 uint8 = 127
                            _ = v30
                            _ = v31
                        }
                        _ = v25
                        _ = v26
                    } else {
                        v32 := GenS2{F0: 2}
                        _ = v32
                    }
                    fmt.Println(v5)
                    _ = fi24
                }
                v22 += v22
                fmt.Println(v22)
            }
        case "foo bar":
            if (v1 && (((((((float32(float64(float32(float64((0.1 - ((float32(0.1) - (-2.25 - 0.0)) + 0.1)))))) + (-2.25 * float32((0.1 - (1.5 - float64((1.5 / 1.5))))))) * float32(1.5)) * -2.25) - 0.1) - float32(float64((0.0 + 0.1)))) + ((0.0 * float32(float64(3.14159))) - 1.5)) > (-2.25 * -2.25))) {
                v5 = v5
            }
        default:
            var v33 int8 = v3
            _ = v33
        }
        v34 := GenS2{F0: 784}
        _ = v34
    }
    v2[128] = (v22 ^ v22)
    var v35 uint16 = 0
    v36 := GenS5{F0: -438, F1: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v35
    _ = v36
    _ = v5
}
