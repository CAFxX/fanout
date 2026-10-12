package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 uint }

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]int)
    var v3 float32 = (3.14159 * float32(float64(1.5)))
    var v4 float32 = v3
    var v5 uint64 = 6
    v6 := make(map[int]uint32)
    for fi7 := 0; fi7 < 14; fi7++ {
        v3 = v3
        v8 := GenS1{F0: "\n\t"}
        var v9 int = (fi7 - 42)
        _ = fi7
        _ = v8
        _ = v9
    }
    v1 = v1
    v5 += v5
    if ((!false) && (true && true)) {
        var v10 float32 = (v4 + float32(1e10))
        v10 = 0.1
        _ = v10
    } else {
        var v11 uint16 = 32767
        if (-1 != -9223372036854775808) {
            switch 9223372036854775807 {
            case 1:
                var v12 uint16 = 1
                _ = v12
            case 689:
                v1 += v1
            }
            v1 += v1
            var v13 int16 = 1000
            _ = v13
        } else {
            if ((v3 * (-2.25 * (((((((float32(0.0) + (v4 / ((3.14159 - 0.1) - v3))) + float32((1e10 * float64(1.5)))) - v3) - v3) - ((v3 * float32((((1e10 / 1.5) / (3.14159 + 0.0)) + 1e10))) + v3)) - (v3 + v4)) * v4))) <= v4) {
                if ((v4 <= ((v4 * ((((1.5 - v3) + (((v3 * (float32(1e10) - (float32(1.5) - (3.14159 / (v3 - (((v3 + (float32(-2.25) - float32(1e10))) + 0.1) * -2.25)))))) / float32(1.5)) - (1.5 / v4))) - -2.25) / (v3 / v3))) * v3)) || ((("\n\t" + ("hello" + "hello")) != ("foo bar" + "hello")) && (42 >= 237))) {
                    fmt.Println(v5, v5, v11)
                } else {
                    if false {
                        if (("hello" + "foo bar") <= "foo bar") {
                            v1 = v1
                        }
                    } else {
                        fmt.Println(v11)
                        for fi14 := 0; fi14 < 4; fi14++ {
                            v4 = v4
                            _ = fi14
                        }
                    }
                    if (687 > 0) {
                        fmt.Println(v1, v3, v11)
                        v15 := GenS2{F0: 42}
                        _ = v15
                    } else {
                        var v16 float64 = -2.25
                        _ = v16
                    }
                }
            }
            for fi17 := 0; fi17 < 7; fi17++ {
                v18 := GenS3{F0: 65535}
                v4 = (v4 + v4)
                _ = fi17
                _ = v18
            }
        }
        _ = v11
    }
    fmt.Println(v4)
    switch "foo bar" {
    case "a":
        v19 := GenS4{F0: 255}
        _ = v19
    default:
        fmt.Println(v1, v1, v3)
    }
    v20 := GenS5{F0: -9223372036854775808, F1: ""}
    v3 = (3.14159 - v4)
    if false {
        for fi21 := 0; fi21 < 3; fi21++ {
            delete(v6, fi21)
            _ = fi21
        }
        v22 := GenS3{F0: 0}
        v23 := GenS3{F0: 4294967295}
        _ = v22
        _ = v23
    } else {
        v24 := []int{-648, 9223372036854775807}
        _ = v24
    }
    var v25 int = 32768
    var v26 uint64 = v5
    for fi27 := 0; fi27 < 6; fi27++ {
        v26 -= v5
        _ = fi27
    }
    v28 := GenS6{F0: uint(v26)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
