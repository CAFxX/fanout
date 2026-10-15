package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint; F1 uint16; F2 int64 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int32 }
func genFn0(p0 int32) bool {
    v5 := GenS1{F0: 0}
    v6 := []int{2147483647, 121}
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int = (fi7 / (((fi7 ^ (fi7 % ((fi7) | 1)))) | 1))
        p0 -= (p0 >> ((501) & 63))
        _ = fi7
        _ = v8
    }
    switch -9223372036854775808 {
    case 372:
        p0 = p0
    case 0:
        var v9 uint16 = 32767
        _ = v9
    case 52:
        if (0 > 2) {
            if true {
                v6 = append(v6, 32767)
                p0 = (p0 / ((p0) | 1))
                for fi10 := 0; fi10 < 2; fi10++ {
                    var v11 int16 = 0
                    v12 := GenS1{F0: -1}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            } else {
                v6 = append(v6, 127)
                p0 = -849
            }
        }
    default:
        p0 -= 291
    }
    var v13 uint = 0
    _ = p0
    _ = v13
    _ = v5
    _ = v6
    return (!(!true))
}
func genFn1(p0 int32, p1 uint64, p2 uint64) int32 {
    fmt.Println(p0, p2)
    var v14 int32 = -2147483648
    for fi15 := 0; fi15 < 5; fi15++ {
        if (1 < 423) {
            if (!(-849 == -748)) {
                v16 := GenS1{F0: 0}
                var v17 float32 = 3.14159
                for fi18 := 0; fi18 < 3; fi18++ {
                    var v19 int32 = 2147483647
                    if ((((-2.25 / ((1.5 / float64((v17 * 3.14159))) + (((((3.14159 / 1e10) / 1.0) * 1.5) * -2.25) - 0.1))) + 1e10) - (3.14159 / float64((v17 + float32(1e10))))) < 1.5) {
                        v20 := GenS2{F0: 0, F1: 439}
                        fmt.Println(fi15, p1)
                        p0 -= p0
                        _ = v20
                    }
                    var v21 int = (fi15 * -1)
                    _ = fi18
                    _ = v19
                    _ = v21
                }
                _ = v16
                _ = v17
            } else {
                if true {
                    for fi22 := 0; fi22 < 6; fi22++ {
                        p1 += 298
                        _ = fi22
                    }
                } else {
                    v23 := []int{(fi15 & fi15), fi15, fi15}
                    p0 = (v14 % ((int32(fi15)) | 1))
                    v14 = (p0 << ((255) & 63))
                    _ = v23
                }
                var v24 float64 = (1.5 - (1e10 * (0.0 + 1.5)))
                if (false || (false || (547 >= 0))) {
                    if true {
                        v25 := []int{}
                        fmt.Println(p2, v14)
                        p0 = v14
                        _ = v25
                    }
                } else {
                    var v26 uint64 = 4294967295
                    var v27 uint64 = (p1 / ((v26) | 1))
                    fmt.Println(p2, p2)
                    _ = v26
                    _ = v27
                }
                _ = v24
            }
            p0 -= v14
        }
        v28 := GenS3{F0: 65535, F1: 255, F2: -9223372036854775808}
        for fi29 := 0; fi29 < 10; fi29++ {
            v30 := []int{(^(fi15)), (fi29 ^ (fi15 * 9223372036854775807)), fi29}
            _ = fi29
            _ = v30
        }
        _ = fi15
        _ = v28
    }
    fmt.Println(v14)
    for fi31 := 0; fi31 < 12; fi31++ {
        if false {
            p2 += (p2 | p1)
            v32 := GenS2{F0: 625, F1: -1}
            _ = v32
        } else {
            fmt.Println(v14)
            var v33 uint = 87
            v34 := GenS4{F0: -128, F1: 159}
            _ = v33
            _ = v34
        }
        _ = fi31
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return p0
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{-9223372036854775808, 128, 158}
    var v3 int = -9223372036854775808
    var v4 uint16 = 65535
    fmt.Println(v3, v4)
    if len(v2) > 0 {
        v2[0] = (v3 + (v3 >> ((42) & 63)))
    }
    if false {
        v4 = v4
        v35 := GenS3{F0: 1, F1: v4, F2: -562}
        _ = v35
    }
    if (3.14159 > (1.5 + (-2.25 + 3.14159))) {
        v4 -= (v4 / ((v4) | 1))
    }
    switch ((("a" + "") + "\n\t") + "\n\t") {
    case "a":
        _ = len(v2)
    case "\n\t":
        v36 := GenS5{F0: 29, F1: -775}
        _ = v36
    case "hello":
        if (false || false) {
            var v37 float64 = 1.5
            _ = v37
        }
    }
    switch (^(v3)) {
    case 32767:
        var v38 int32 = -149
        _ = v38
    }
    v1[-9223372036854775808] = (((1e10 * 0.0) != ((0.0 * (0.0 / float64((0.1 / (-2.25 + ((3.14159 / 0.1) + -2.25)))))) - -2.25)) && true)
    fmt.Println(v4, v4)
    var v39 int32 = 42
    var v40 int = (^(v3))
    v40 += v40
    v41 := GenS2{F0: 4294967295, F1: 674}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
