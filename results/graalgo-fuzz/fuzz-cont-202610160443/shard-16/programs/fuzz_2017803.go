package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 string; F1 uint; F2 int16 }
type GenS4 struct { F0 bool; F1 int; F2 string }
func genFn0(p0 uint16) uint {
    v4 := []int{9223372036854775807, 1, 127, 32768}
    v5 := GenS1{F0: 103}
    if true {
        for fi6 := 0; fi6 < 4; fi6++ {
            v7 := GenS1{F0: 47}
            _ = fi6
            _ = v7
        }
        v8 := GenS2{F0: 1, F1: 0}
        _ = v8
    }
    v9 := GenS3{F0: "", F1: 4294967295, F2: 42}
    _ = p0
    _ = v4
    _ = v5
    _ = v9
    return 4294967295
}

func main() {
    var v1 float64 = (0.0 - (((0.1 / (-2.25 - (float64(0.1) - (((1e10 + ((1.5 - 1.5) + 3.14159)) + ((((1e10 + -2.25) / (float64(1.5) - 3.14159)) * (3.14159 / 3.14159)) / 1.0)) + (float64(float32(((float64(0.0) + float64(float32(float64(0.0)))) - 1.5))) / 1.0))))) * 1e10) - (0.1 / 3.14159)))
    var v2 int16 = -32768
    var v3 float64 = v1
    v1 -= v1
    v10 := GenS4{F0: false, F1: 127, F2: (("foo bar" + "hello") + (("a" + "hello") + ("hello" + ("hello" + "foo bar"))))}
    if (!(v1 == ((v3 - (v3 - v1)) / v3))) {
        v1 -= (0.0 + (1e10 + 0.1))
        var v11 uint8 = 42
        var v12 float32 = 3.14159
        _ = v11
        _ = v12
    }
    v2 += (^(v2))
    fmt.Println(v2, v1)
    var v13 uint8 = 42
    for fi14 := 0; fi14 < 5; fi14++ {
        for fi15 := 0; fi15 < 11; fi15++ {
            v16 := []int{}
            v16 = append(v16, fi14)
            var v17 uint32 = 832
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    v18 := GenS4{F0: ((65535 > 4294967295) || false), F1: -201, F2: "hello"}
    for fi19 := 0; fi19 < 12; fi19++ {
        v2 -= -769
        v2 += v2
        if (423 <= 42) {
            if false {
                v1 += ((float64(0.0) * (1.5 * ((0.0 + v1) * (float64(3.14159) / (((v1 + ((v1 + ((((v3 * ((((v1 / 3.14159) + (v1 + v1)) - v1) / v3)) * v3) / (((-2.25 / v1) * (v1 * v3)) / v3)) - -2.25)) * 0.0)) / ((((v1 * v1) + (v1 / 1.5)) / -2.25) - ((-2.25 / 1.0) / ((1.5 / (3.14159 + ((((0.1 - float64(0.0)) + 0.0) + v1) + ((v1 / 1.0) * float64(1e10))))) - (3.14159 - v1))))) + float64(-2.25)))))) - (0.0 * ((float64((float32(float64(-2.25)) * float32(0.1))) + v1) - v3)))
            }
        }
        _ = fi19
    }
    v2 += int16(v13)
    v20 := GenS2{F0: -74, F1: -1}
    for fi21 := 0; fi21 < 20; fi21++ {
        v22 := []int{int(v13)}
        if ((0 != 65535) || false) {
            v2 += v2
        } else {
            var v23 int = (-(fi21))
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    v13 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v18
    _ = v2
    _ = v20
    _ = v3
}
