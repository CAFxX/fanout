package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int16; F2 uint8 }
func genFn0() uint {
    var v4 int16 = 0
    switch (("" + "foo bar") + "hello") {
    case "hello":
        for fi5 := 0; fi5 < 17; fi5++ {
            fmt.Println(fi5)
            var v6 uint16 = 42
            for fi7 := 0; fi7 < 3; fi7++ {
                v6 -= (v6 % ((v6) | 1))
                switch fi7 {
                case 805:
                    switch "a" {
                    case "\n\t":
                        var v8 int64 = -612
                        _ = v8
                    }
                    var v9 int8 = 127
                    v4 -= (v4 >> ((304) & 63))
                    _ = v9
                case 127:
                    v10 := []int{fi7, (fi7 + 32768)}
                    _ = v10
                }
                switch (fi7 ^ fi5) {
                case 32768:
                    var v11 int16 = (v4 >> ((4294967295) & 63))
                    _ = v11
                case 232:
                    if false {
                        if true {
                            v12 := []int{}
                            var v13 int = -9223372036854775808
                            _ = v12
                            _ = v13
                        } else {
                            var v14 int8 = 0
                            _ = v14
                        }
                        v15 := GenS1{F0: 1}
                        v4 -= (v4 & (^(v4)))
                        v16 := []int{42, 9223372036854775807, (fi5 ^ fi5), fi5}
                        var v17 uint64 = 1
                        _ = v15
                        _ = v16
                        _ = v17
                    } else {
                        var v18 int64 = 42
                        var v19 int = 42
                        _ = v18
                        _ = v19
                    }
                case 128:
                    v20 := []int{fi7, (fi5 & 2147483647)}
                    _ = v20
                }
                _ = fi7
            }
            _ = fi5
            _ = v6
        }
    default:
        var v21 float32 = float32((1.5 / 1.5))
        _ = v21
    }
    switch 4294967295 {
    case -545:
        v4 = (v4 % ((-176) | 1))
    default:
        var v22 int64 = 1
        _ = v22
    }
    v4 = (v4 >> ((42) & 63))
    _ = v4
    return 0
}

func main() {
    var v1 uint32 = 65535
    v2 := []string{(("\n\t" + ("foo bar" + "")) + "hello")}
    var v3 int16 = 1
    for fi23 := 0; fi23 < 16; fi23++ {
        v3 -= v3
        var v24 int32 = 0
        _ = fi23
        _ = v24
    }
    _ = len(v2)
    fmt.Println(v3)
    fmt.Println(v3, v1)
    switch "hello" {
    case "a":
        fmt.Println(v1, v3)
    case "\n\t":
        var v25 uint64 = 310
        _ = v25
    }
    var v26 uint8 = 0
    var v27 int8 = 1
    fmt.Println(v3, v3, v3)
    var v28 float32 = (float32((float64(-2.25) / 1.5)) - ((((1e10 * (3.14159 * (((3.14159 / 1.0) - (float32((((1.5 * 1.5) * ((-2.25 + 1.5) * 1.5)) * (-2.25 - float64((3.14159 - 0.0))))) - (((((3.14159 - 1.5) + 0.1) * float32(1.5)) / 0.1) - (((3.14159 * 3.14159) + (-2.25 / float32(0.1))) * -2.25)))) - (1.5 * ((1.5 * 1e10) * (0.1 + float32(-2.25))))))) / (1.5 / 1.0)) - 0.1) / 1.0))
    for fi29 := 0; fi29 < 4; fi29++ {
        switch 274 {
        case -567:
            v30 := GenS1{F0: 2147483647}
            _ = v30
        case 600:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 42:
            if false {
                for fi31 := 0; fi31 < 8; fi31++ {
                    v26 = (v26 ^ 255)
                    fmt.Println(fi31)
                    _ = fi31
                }
                var v32 uint8 = 103
                _ = v32
            }
        default:
            v33 := GenS2{F0: 1, F1: (^(v3)), F2: (v26 - (v26 - v26))}
            _ = v33
        }
        v26 += (v26 / (((v26 - v26)) | 1))
        _ = fi29
    }
    var v34 string = (("\n\t" + "") + "hello")
    for fi35 := 0; fi35 < 11; fi35++ {
        v1 = (^(v1))
        v2 = append(v2, v34)
        _ = fi35
    }
    fmt.Println(v1, v27, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v34
}
