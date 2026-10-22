package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 int32; F1 uint64; F2 uint16 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 42
    var v3 int64 = 9223372036854775807
    var v4 float32 = 3.14159
    var v5 uint8 = (^(v2))
    v6 := make(map[int]uint32)
    if ((-2.25 / 1.0) <= ((1.5 - 3.14159) + 1.5)) {
        for fi7 := 0; fi7 < 4; fi7++ {
            for fi8 := 0; fi8 < 11; fi8++ {
                var v9 int32 = -2147483648
                var v10 int = -2147483648
                v11 := GenS1{F0: -1, F1: 1}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            v3 -= int64(v5)
            if false {
                v3 += (v3 | (v3 ^ (v3 << ((65535) & 63))))
            }
            _ = fi7
        }
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 string = "a"
        fmt.Println(v4, v13, fi12)
        _ = fi12
        _ = v13
    }
    if ((-2.25 - ((v4 / (v4 - ((((3.14159 + float32(3.14159)) / (float32(1.5) * ((float32((-2.25 + float64(0.0))) + float32(float64(v4))) * (v4 * -2.25)))) - 3.14159) - v4))) / v4)) != 1e10) {
        switch ("" + "hello") {
        case "":
            delete(v6, -9223372036854775808)
        default:
            if ("hello" <= ("foo bar" + ("" + ("a" + "hello")))) {
                fmt.Println(v5, v4)
                if (2 < v5) {
                    var v14 int8 = 100
                    fmt.Println(v14, v4)
                    _ = v14
                }
                if (v2 <= (v5 - 131)) {
                    if (42 <= 0) {
                        for fi15 := 0; fi15 < 15; fi15++ {
                            fmt.Println(v5, v5, v5)
                            _ = len(v6)
                            _ = fi15
                        }
                        for fi16 := 0; fi16 < 12; fi16++ {
                            v17 := []int{(fi16 << ((18446744073709551615) & 63))}
                            _ = fi16
                            _ = v17
                        }
                    }
                }
            } else {
                _ = len(v1)
                var v18 float32 = -2.25
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 float32 = 0.1
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 15; fi21++ {
            v4 += v4
            if (!(false || false)) {
                var v22 bool = (!false)
                var v23 uint8 = (v2 << ((16) & 63))
                v24 := GenS2{F0: 32767}
                _ = v22
                _ = v23
                _ = v24
            } else {
                v25 := []int{(fi21 % ((fi21) | 1)), fi21}
                if true {
                    for fi26 := 0; fi26 < 9; fi26++ {
                        if false {
                            fmt.Println(v3, v4, v4)
                        }
                        _ = fi26
                    }
                    v3 = v3
                    var v27 int32 = -1
                    _ = v27
                }
                var v28 float64 = 0.1
                _ = v25
                _ = v28
            }
            _ = fi21
        }
    } else {
        var v29 float64 = -2.25
        v30 := GenS3{F0: "a", F1: 0}
        fmt.Println(v4)
        _ = v29
        _ = v30
    }
    v31 := GenS4{F0: 2147483647, F1: 2147483647}
    for fi32 := 0; fi32 < 15; fi32++ {
        v33 := GenS5{F0: 2147483647, F1: 4294967295, F2: 628}
        _ = fi32
        _ = v33
    }
    var v34 int8 = 127
    v4 = (v4 + 0.0)
    fmt.Println(v34, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
