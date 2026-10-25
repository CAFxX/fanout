package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int64 }
type GenS2 struct { F0 int32; F1 int8; F2 uint }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint }

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 835
    var v3 uint = 42
    var v4 bool = false
    v5 := GenS1{F0: v2, F1: ("foo bar" + ("a" + "a")), F2: -9223372036854775808}
    v2 = 255
    v2 = (^(v2))
    var v6 int8 = -1
    var v7 int = -9223372036854775808
    switch (v7 % ((v7) | 1)) {
    case -2147483648:
        if true {
            var v8 float64 = -2.25
            _ = v8
        } else {
            v9 := []int{(v7 - v7), v7}
            v6 = int8(v2)
            _ = v9
        }
    case 908:
        if ((0.0 * (0.0 / 1.5)) <= 3.14159) {
            v10 := GenS2{F0: -712, F1: v6, F2: (v3 | 432)}
            switch ("hello" + "a") {
            case "":
                fmt.Println(v6)
            case "a":
                v11 := []int{(v7 << (((v3 + 65535)) & 63)), (v7 - v7)}
                _ = v11
            default:
                var v12 bool = (!true)
                _ = v12
            }
            _ = v10
        }
    case 84:
        for fi13 := 0; fi13 < 20; fi13++ {
            v2 = v2
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 float32 = float32(((0.1 + 3.14159) / ((3.14159 - -2.25) - (1.5 - 1.5))))
                v15 -= (1.5 * float32((-2.25 + 3.14159)))
                v16 := []int{}
                _ = fi14
                _ = v15
                _ = v16
            }
            v17 := GenS2{F0: -2147483648, F1: (v6 >> ((v3) & 63)), F2: (v3 & (^(v3)))}
            _ = fi13
            _ = v17
        }
    }
    if v4 {
        if v4 {
            for fi18 := 0; fi18 < 1; fi18++ {
                for fi19 := 0; fi19 < 11; fi19++ {
                    v4 = v4
                    _ = fi19
                }
                _ = fi18
            }
        }
    } else {
        v20 := GenS3{F0: v3, F1: v2}
        _ = v20
    }
    v21 := GenS4{F0: (v3 & 255)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
