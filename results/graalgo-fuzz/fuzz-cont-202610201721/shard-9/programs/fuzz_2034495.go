package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint }
func genFn0(p0 int8) uint32 {
    if (-1 > 0) {
        var v5 int = -629
        if (0 <= 762) {
            var v6 uint16 = 0
            v5 = 9223372036854775807
            _ = v6
        } else {
            var v7 int16 = -32768
            if true {
                v8 := GenS1{F0: 39, F1: 2, F2: 18446744073709551615}
                _ = v8
            }
            var v9 uint = 255
            _ = v7
            _ = v9
        }
        v10 := []int{506, 585, (v5 * (v5 * (v5 - v5))), 32767}
        _ = v10
        _ = v5
    } else {
        var v11 int16 = 32767
        p0 = p0
        p0 += (p0 ^ (p0 | p0))
        _ = v11
    }
    var v12 float32 = ((((3.14159 + (0.1 * 1.5)) * -2.25) + float32((1.5 * 1.5))) * 0.0)
    v13 := []int{-9223372036854775808, 0, -9223372036854775808}
    if (!((((1.5 - -2.25) + (1e10 - 1.5)) / (float64((v12 * ((-2.25 - (((float32(1e10) - v12) / v12) - (float32(0.1) + v12))) + (0.1 / (v12 - ((v12 / v12) - (0.1 * float32(1e10)))))))) - 0.1)) == 3.14159)) {
        p0 = 42
    }
    var v14 float32 = 3.14159
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 0
}
func genFn1(p0 int32, p1 uint16, p2 uint) uint32 {
    var v15 float32 = (0.1 + (((float32(float64((0.0 * 1e10))) - (1e10 * (-2.25 - (((0.1 * (0.1 + 0.1)) * ((1e10 - 3.14159) + ((0.0 * float32(1e10)) / -2.25))) - 3.14159)))) + (0.1 + float32(0.0))) / 3.14159))
    v16 := []int{128}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 65535
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]string)
    v3 := make(map[uint32]bool)
    v4 := []int{127, 174, 32768, -882}
    if (!(100 == 84)) {
        switch -1 {
        case 4294967295:
            v1 -= (v1 % (((v1 % (((v1 / ((2147483648) | 1))) | 1))) | 1))
        default:
            switch "hello" {
            case "foo bar":
                v1 = (^(v1))
            case "":
                if (!false) {
                    for fi17 := 0; fi17 < 20; fi17++ {
                        v1 = (v1 % ((2) | 1))
                        _ = fi17
                    }
                } else {
                    v1 += v1
                    var v18 bool = ("foo bar" == "hello")
                    _ = v18
                }
            case "hello":
                v19 := GenS1{F0: 127, F1: 1, F2: 4294967295}
                _ = v19
            }
        }
        if false {
            if len(v4) > 0 {
                v4[0] = 1
            }
        } else {
            v1 = (v1 - 42)
            v1 += 65535
        }
    }
    if (0 == 129) {
        fmt.Println(v1)
        var v20 int32 = -2147483648
        _ = v20
    }
    if ((-9223372036854775808 == -9223372036854775808) && false) {
        v1 = 0
    }
    if false {
        var v21 int32 = 1
        for fi22 := 0; fi22 < 18; fi22++ {
            v21 += v21
            switch int(v21) {
            case 127:
                v23 := GenS1{F0: 1, F1: 647, F2: 2}
                _ = v23
            case -266:
                v1 -= v1
            case -1:
                var v24 uint32 = v1
                _ = v24
            }
            _ = fi22
        }
        v25 := GenS1{F0: 2, F1: 32767, F2: 255}
        _ = v21
        _ = v25
    }
    v1 = v1
    v1 = 0
    v1 = 0
    for fi26 := 0; fi26 < 14; fi26++ {
        v1 += v1
        switch fi26 {
        case -208:
            if true {
                fmt.Println(v1)
            }
        case 128:
            v1 -= (^(v1))
        }
        _ = fi26
    }
    switch 55 {
    case 32767:
        var v27 int32 = 1
        _ = v27
    case 0:
        v3[2] = false
    case 1:
        v1 += (v1 / ((v1) | 1))
    default:
        switch (("foo bar" + "\n\t") + "a") {
        case "foo bar":
            var v28 bool = false
            _ = v28
        default:
            var v29 string = "hello"
            _ = v29
        }
    }
    v30 := GenS1{F0: 128, F1: 32767, F2: 255}
    for fi31 := 0; fi31 < 18; fi31++ {
        if (4294967295 == 4294967295) {
            delete(v2, "hello")
            var v32 uint16 = 477
            _ = v32
        }
        fmt.Println(fi31, v1)
        _ = fi31
    }
    v33 := GenS1{F0: 0, F1: 42, F2: 350}
    v1 -= 2147483647
    switch "" {
    case "a":
        for fi34 := 0; fi34 < 19; fi34++ {
            v1 += 2
            v1 = (v1 % ((v1) | 1))
            _ = fi34
        }
    default:
        if false {
            var v35 int8 = -46
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v4
}
