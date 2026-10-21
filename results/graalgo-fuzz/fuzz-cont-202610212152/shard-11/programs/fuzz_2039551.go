package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int8; F1 string }

func main() {
    var v1 int8 = 62
    var v2 float64 = (((((float64(0.0) + 0.0) - 1e10) / (1e10 / 1.0)) / 3.14159) - (-2.25 - (3.14159 + -2.25)))
    v3 := make(map[string]int)
    v1 = (v1 | (v1 + 100))
    v2 += 3.14159
    v4 := GenS1{F0: -2147483648}
    v2 -= 0.0
    v3[(("\n\t" + ("foo bar" + "\n\t")) + "")] = 0
    if true {
        switch "a" {
        case "a":
            var v5 string = ""
            _ = v5
        case "":
            delete(v3, "a")
        default:
            for fi6 := 0; fi6 < 3; fi6++ {
                var v7 uint = 255
                v2 = v2
                _ = fi6
                _ = v7
            }
        }
        var v8 int64 = 1
        _ = v8
    }
    v9 := GenS2{F0: 32767, F1: 32768}
    var v10 uint = 1
    v11 := GenS3{F0: 255, F1: 385}
    v12 := GenS3{F0: 255, F1: -107}
    fmt.Println(v1, v1, v10)
    for fi13 := 0; fi13 < 6; fi13++ {
        switch (-(fi13)) {
        case 139:
            var v14 int64 = 42
            _ = v14
        case 1:
            fmt.Println(v2, v1)
        default:
            v1 += v1
        }
        var v15 int64 = 42
        var v16 int64 = (v15 & (v15 + (v15 | v15)))
        _ = fi13
        _ = v15
        _ = v16
    }
    if (32767 > -32768) {
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 float64 = (((v2 + (v2 * float64(((1.5 * (((0.1 - (float32(v2) / 1e10)) + float32(v2)) - 1.5)) - (((1e10 + (1e10 / 0.1)) * 3.14159) / -2.25))))) * (v2 / 1e10)) / v2)
            switch fi17 {
            case 32767:
                var v19 float32 = -2.25
                _ = v19
            case 664:
                var v20 uint8 = 42
                _ = v20
            case -636:
                v21 := []int{}
                _ = v21
            }
            v22 := []int{(fi17 % ((4294967295) | 1)), fi17, fi17}
            _ = fi17
            _ = v18
            _ = v22
        }
        v1 = v1
        switch ((("" + "hello") + "hello") + ("foo bar" + "hello")) {
        case "foo bar":
            for fi23 := 0; fi23 < 5; fi23++ {
                fmt.Println(fi23, v2, v2)
                var v24 uint8 = 42
                v10 -= v10
                _ = fi23
                _ = v24
            }
        case "hello":
            v25 := GenS4{F0: v1, F1: "foo bar"}
            _ = v25
        case "a":
            for fi26 := 0; fi26 < 18; fi26++ {
                switch fi26 {
                case 515:
                    for fi27 := 0; fi27 < 17; fi27++ {
                        v28 := GenS1{F0: fi26}
                        _ = fi27
                        _ = v28
                    }
                case 9223372036854775807:
                    var v29 uint64 = 18446744073709551615
                    _ = v29
                case 42:
                    v10 = v10
                }
                _ = fi26
            }
        default:
            if false {
                if ((100 == v1) || (true || true)) {
                    v1 += 0
                    v2 += (v2 / float64(0.1))
                }
            }
        }
    } else {
        v1 -= 100
    }
    if ((255 >= 212) || false) {
        for fi30 := 0; fi30 < 13; fi30++ {
            for fi31 := 0; fi31 < 4; fi31++ {
                var v32 uint16 = 0
                v33 := v3["\n\t"]
                _ = fi31
                _ = v32
                _ = v33
            }
            _ = fi30
        }
    }
    fmt.Println(v2, v10, v1)
    var v34 uint = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v9
}
