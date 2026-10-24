package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint32 }

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = -9223372036854775808
    var v3 int32 = int32(v2)
    fmt.Println(v2, v2)
    var v4 bool = ((!false) || false)
    switch 490 {
    case 4294967295:
        var v5 bool = true
        _ = v5
    case 42:
        switch 32767 {
        case -9223372036854775808:
            v6 := GenS1{F0: 65535}
            _ = v6
        case 222:
            switch ("" + "a") {
            case "\n\t":
                v2 = (-(v2))
            case "":
                v3 -= 2147483647
            case "hello":
                var v7 int = 9223372036854775807
                _ = v7
            }
        }
    case -360:
        var v8 uint = 65535
        _ = v8
    default:
        if (1 <= 42) {
            v2 = int64(v3)
            var v9 float32 = ((1.5 * (-2.25 * 1.5)) / 1e10)
            _ = v9
        }
    }
    var v10 float32 = (-2.25 - (float32(1e10) - 0.1))
    if v4 {
        delete(v1, "foo bar")
        v11 := GenS1{F0: 18446744073709551615}
        switch (("" + ("hello" + "a")) + "") {
        case "":
            if v4 {
                v12 := []int{295, -9223372036854775808, 128, 787}
                _ = v12
            }
        }
        _ = v11
    } else {
        v4 = false
        var v13 int = 128
        var v14 float64 = 0.1
        _ = v13
        _ = v14
    }
    var v15 int8 = 0
    v16 := GenS1{F0: 42}
    for fi17 := 0; fi17 < 15; fi17++ {
        switch 32767 {
        case 1:
            var v18 uint = 4294967295
            _ = v18
        case 32768:
            v19 := []int{fi17}
            _ = v19
        }
        v3 = (v3 / (((v3 / ((273) | 1))) | 1))
        _ = fi17
    }
    v20 := v1[("a" + (("hello" + "a") + ("" + ("" + ("foo bar" + "")))))]
    switch ("" + "a") {
    case "\n\t":
        var v21 int16 = 649
        _ = v21
    case "a":
        for fi22 := 0; fi22 < 20; fi22++ {
            for fi23 := 0; fi23 < 1; fi23++ {
                _ = len(v1)
                var v24 int8 = (v15 + v15)
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    case "hello":
        if (42 <= -272) {
            if (!((0.1 / float64(v10)) > ((-2.25 - float64(v10)) + ((0.0 - 3.14159) / 3.14159)))) {
                v4 = (65535 > 983)
            } else {
                v3 = (v3 % ((-1) | 1))
                v25 := GenS2{F0: (v15 * 39), F1: 65535}
                v26 := GenS1{F0: 4294967295}
                _ = v25
                _ = v26
            }
        }
    default:
        var v27 float64 = (0.1 - ((0.0 * 0.0) * float64((1e10 / float32(3.14159)))))
        _ = v27
    }
    v28 := GenS2{F0: (v15 % (((^(v15))) | 1)), F1: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
}
