package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint16 }

func main() {
    var v1 uint16 = 228
    var v2 uint64 = 1
    v3 := make(map[string]string)
    var v4 uint16 = (v1 * (v1 << ((169) & 63)))
    v5 := make(map[string]bool)
    if ((("hello" + "a") + "\n\t") <= (((("" + "foo bar") + "foo bar") + "hello") + "foo bar")) {
        for fi6 := 0; fi6 < 14; fi6++ {
            v1 += 65535
            var v7 uint = 1
            var v8 int8 = -1
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: 0, F1: ""}
        _ = v9
    } else {
        v10 := GenS1{F0: 32767, F1: (("hello" + "\n\t") + ("" + ("" + "a")))}
        var v11 uint64 = (v2 >> ((4294967295) & 63))
        _ = v10
        _ = v11
    }
    if false {
        v4 += 1
    }
    switch ("foo bar" + "a") {
    case "":
        if false {
            var v12 uint8 = 255
            switch int(v2) {
            case 9223372036854775807:
                v2 = (v2 | v2)
            default:
                var v13 int8 = -128
                _ = v13
            }
            _ = v12
        }
    default:
        v14 := []int{174, 4294967295, 0}
        _ = v14
    }
    v15 := GenS2{F0: -2147483648}
    v16 := GenS2{F0: 0}
    switch -9223372036854775808 {
    case 4294967295:
        v1 += v1
    }
    v17 := []int{32768, -812}
    for fi18 := 0; fi18 < 16; fi18++ {
        if (0.0 != (float64((0.1 / float32((((1.5 + 1e10) / (1e10 * 1e10)) / (float64(((float32(-2.25) * 1e10) / 1e10)) + float64((3.14159 + 3.14159))))))) * ((0.0 * (-2.25 * 3.14159)) - (-2.25 * float64(-2.25))))) {
            if len(v17) > 0 {
                v17[0] = (fi18 / (((^(fi18))) | 1))
            }
        }
        _ = fi18
    }
    var v19 float32 = (float32((1.5 + (3.14159 * 1.5))) * 0.0)
    if (((v19 - v19) * 3.14159) <= (v19 / 1.0)) {
        v2 = v2
    } else {
        var v20 int32 = 0
        _ = v20
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        if true {
            switch "foo bar" {
            case "":
                v4 = (v1 >> ((1) & 63))
            case "\n\t":
                fmt.Println(v1, v4, v1)
            default:
                for fi22 := 0; fi22 < 2; fi22++ {
                    if false {
                        v4 += 0
                    } else {
                        v23 := GenS1{F0: 905, F1: "hello"}
                        _ = v23
                    }
                    var v24 uint = 2
                    switch fi22 {
                    case -957:
                        var v25 int8 = 100
                        v26 := GenS2{F0: 0}
                        var v27 int8 = 0
                        _ = v25
                        _ = v26
                        _ = v27
                    default:
                        v19 -= float32(0.1)
                    }
                    _ = fi22
                    _ = v24
                }
            }
        } else {
            v28 := GenS1{F0: v4, F1: "\n\t"}
            _ = v28
        }
        for fi29 := 0; fi29 < 18; fi29++ {
            v30 := GenS2{F0: -2147483648}
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            var v31 uint8 = 2
            _ = fi29
            _ = v30
            _ = v31
        }
        switch fi21 {
        case 128:
            _ = len(v5)
        case 32768:
            v19 = 0.1
        case -299:
            switch 127 {
            case 128:
                fmt.Println(v4)
            default:
                v32 := GenS3{F0: 405, F1: v1}
                _ = v32
            }
        }
        _ = fi21
    }
    switch 32767 {
    case -9223372036854775808:
        _ = len(v3)
    case 127:
        v2 = 58
    }
    var v33 uint64 = (v2 / (((v2 << ((65535) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
