package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() string {
    switch "\n\t" {
    case "":
        var v5 string = "foo bar"
        _ = v5
    default:
        var v6 int = 4294967295
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        if true {
            switch "foo bar" {
            case "":
                v8 := []int{}
                _ = v8
            case "\n\t":
                var v9 float64 = 3.14159
                _ = v9
            case "hello":
                for fi10 := 0; fi10 < 3; fi10++ {
                    fmt.Println(fi7, fi7)
                    _ = fi10
                }
            default:
                var v11 uint32 = 2
                _ = v11
            }
            var v12 float64 = -2.25
            v13 := GenS1{F0: 42}
            _ = v12
            _ = v13
        } else {
            var v14 string = "a"
            for fi15 := 0; fi15 < 14; fi15++ {
                switch 0 {
                case 9223372036854775807:
                    var v16 int8 = -128
                    _ = v16
                default:
                    var v17 int = (fi7 ^ (fi15 << ((1) & 63)))
                    v18 := []int{fi7, fi7, 128, (^(fi15))}
                    v14 = "foo bar"
                    _ = v17
                    _ = v18
                }
                var v19 int32 = -70
                var v20 uint8 = 128
                _ = fi15
                _ = v19
                _ = v20
            }
            _ = v14
        }
        _ = fi7
    }
    v21 := GenS1{F0: 693}
    _ = v21
    return ""
}
func genFn1() int {
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 float32 = (0.1 / 3.14159)
        fmt.Println(fi22)
        _ = fi22
        _ = v23
    }
    var v24 uint8 = 128
    fmt.Println(v24)
    v24 = 1
    _ = v24
    return 332
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    var v3 uint32 = 754
    v4 := make(map[string]string)
    v25 := GenS1{F0: 42}
    v3 = v3
    v3 -= 583
    v26 := GenS1{F0: 1}
    for fi27 := 0; fi27 < 6; fi27++ {
        var v28 int32 = 1
        _ = fi27
        _ = v28
    }
    switch ("" + ("a" + "hello")) {
    case "foo bar":
        for fi29 := 0; fi29 < 6; fi29++ {
            fmt.Println(v3, fi29)
            for fi30 := 0; fi30 < 2; fi30++ {
                v3 = 715
                fmt.Println(fi29, v3, fi29)
                _ = fi30
            }
            _ = fi29
        }
    case "a":
        switch 32767 {
        case 365:
            if (32767 > 255) {
                var v31 uint8 = 0
                var v32 int32 = 42
                if (false || false) {
                    for fi33 := 0; fi33 < 3; fi33++ {
                        var v34 int8 = 42
                        v34 = 26
                        _ = fi33
                        _ = v34
                    }
                    v35 := v2[genFn1()]
                    _ = v35
                }
                _ = v31
                _ = v32
            } else {
                v36 := []int{90, 2147483647, 9223372036854775807, 128}
                _ = v36
            }
        case -105:
            v3 += (v3 % ((v3) | 1))
        case 127:
            switch 2 {
            case -128:
                v3 = (v3 % (((v3 + (^(v3)))) | 1))
            case 2147483647:
                v3 += v3
            case 42:
                var v37 int16 = -681
                _ = v37
            }
        }
    }
    fmt.Println(v3)
    var v38 uint32 = 2147483647
    switch -899 {
    case 128:
        v38 += 1
    case -613:
        v38 -= 65535
    }
    fmt.Println(v3)
    var v39 float64 = 0.0
    var v40 uint16 = uint16(v3)
    v3 -= v3
    v41 := GenS1{F0: 536}
    var v42 float32 = (-2.25 + (((0.1 * (0.0 * -2.25)) / (((1e10 * (-2.25 * 3.14159)) - -2.25) * (0.0 + -2.25))) * (3.14159 - 1e10)))
    v42 = v42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
