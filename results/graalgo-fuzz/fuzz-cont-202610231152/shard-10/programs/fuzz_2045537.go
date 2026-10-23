package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int8; F1 int8; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint64 }
func genFn0(p0 uint32, p1 uint16, p2 int32) int8 {
    var v7 string = "\n\t"
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 float32 = (float32(((1.5 / 1.0) / ((3.14159 + 1.5) / 3.14159))) / float32(3.14159))
        p1 += (p1 ^ (p1 << ((212) & 63)))
        p2 = p2
        _ = fi8
        _ = v9
    }
    switch 32768 {
    case 251:
        if true {
            v10 := GenS1{F0: 1, F1: 1}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 int32 = -1
                v12 = (v12 ^ -1)
                var v13 float64 = 3.14159
                _ = fi11
                _ = v12
                _ = v13
            }
            v14 := []int{0}
            _ = v14
        }
    }
    v15 := []int{770, 172, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v7
    return int8(p2)
}
func genFn1() bool {
    v16 := []int{577, 128, -2147483648}
    switch 42 {
    case -247:
        var v17 int64 = 1
        _ = v17
    case 629:
        switch "foo bar" {
        case "":
            var v18 int8 = genFn0(386, 255, 635)
            _ = v18
        case "foo bar":
            if (("foo bar" + ("" + "\n\t")) <= "\n\t") {
                var v19 int32 = 42
                v20 := GenS2{F0: 42, F1: 85, F2: 951}
                if true {
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                }
                _ = v19
                _ = v20
            } else {
                var v21 uint8 = 54
                v22 := GenS3{F0: 328}
                _ = v21
                _ = v22
            }
        }
    }
    v23 := GenS3{F0: 18446744073709551615}
    _ = v16
    _ = v23
    return (!(42 > 951))
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    var v3 uint = 4294967295
    var v4 bool = ((((!true) && false) && false) && true)
    v5 := make(map[string]bool)
    var v6 int32 = -1
    if false {
        v24 := GenS3{F0: 114}
        v3 = 1
        fmt.Println(v4, v4, v3)
        _ = v24
    }
    var v25 uint32 = 714
    v26 := []int{2147483647, 127, 1}
    v6 = v6
    switch 32767 {
    case -620:
        var v27 uint = v3
        _ = v27
    case 128:
        switch 32 {
        case 481:
            v28 := GenS3{F0: v3}
            _ = v28
        }
    default:
        var v29 string = "\n\t"
        _ = v29
    }
    v6 -= v6
    if (v4 || (!v4)) {
        v30 := GenS4{F0: 0}
        var v31 float64 = float64(float32((1e10 * 1.5)))
        _ = v30
        _ = v31
    }
    if ((!(v4 || (v4 && (!v4)))) || (9223372036854775807 <= 32768)) {
        var v32 bool = (!v4)
        _ = len(v1)
        v32 = ("\n\t" >= ((("" + "\n\t") + "a") + "\n\t"))
        _ = v32
    } else {
        var v33 uint32 = (v25 / (((v25 << (((v3 + (^(v3)))) & 63))) | 1))
        v34 := GenS5{F0: 127, F1: 0}
        _ = v33
        _ = v34
    }
    fmt.Println(v4, v3, v6)
    v4 = (!v4)
    if (((float32(-2.25) * 1.5) - float32((((float64(((0.0 * 1.5) + float32(((3.14159 * 0.1) * ((0.1 - (1.5 * 0.0)) / 1.0))))) - float64((3.14159 / 1.0))) * (0.0 * 0.0)) * (0.0 / 0.1)))) == 1.5) {
        for fi35 := 0; fi35 < 3; fi35++ {
            v4 = (!v4)
            v36 := GenS4{F0: 2}
            _ = fi35
            _ = v36
        }
        v3 += 1
        delete(v2, "hello")
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
