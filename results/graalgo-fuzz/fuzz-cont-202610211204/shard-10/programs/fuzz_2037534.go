package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 int16, p1 int64) int32 {
    v6 := []int{975, 42}
    if (!false) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        p0 = -1
        v7 := GenS1{F0: 271}
        _ = v7
    } else {
        if false {
            v8 := GenS1{F0: 42}
            var v9 float64 = -2.25
            var v10 int32 = -1
            _ = v10
            _ = v8
            _ = v9
        }
    }
    fmt.Println(p1, p0)
    _ = len(v6)
    var v11 int16 = p0
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return 0
}
func genFn1() uint8 {
    for fi12 := 0; fi12 < 2; fi12++ {
        v13 := GenS1{F0: 18446744073709551615}
        if (false && true) {
            switch "hello" {
            case "foo bar":
                for fi14 := 0; fi14 < 7; fi14++ {
                    if (!false) {
                        var v15 uint8 = 127
                        var v16 uint16 = 255
                        _ = v15
                        _ = v16
                    } else {
                        fmt.Println(fi14)
                        var v17 uint = 255
                        _ = v17
                    }
                    var v18 uint64 = 2
                    v18 -= v18
                    _ = fi14
                    _ = v18
                }
            case "":
                var v19 uint8 = 77
                _ = v19
            case "a":
                if false {
                    var v20 float32 = 3.14159
                    fmt.Println(fi12, fi12)
                    _ = v20
                }
            }
            var v21 int8 = 1
            v22 := GenS2{F0: fi12, F1: 42}
            _ = v21
            _ = v22
        } else {
            for fi23 := 0; fi23 < 12; fi23++ {
                var v24 uint16 = 255
                _ = fi23
                _ = v24
            }
            var v25 float32 = float32((3.14159 / 1.0))
            v25 += v25
            _ = v25
        }
        var v26 string = ""
        _ = fi12
        _ = v13
        _ = v26
    }
    var v27 uint16 = 255
    for fi28 := 0; fi28 < 20; fi28++ {
        v27 -= (v27 / ((42) | 1))
        _ = fi28
    }
    v27 = v27
    v29 := GenS3{F0: 48, F1: 100}
    _ = v27
    _ = v29
    return 2
}

func main() {
    var v1 int8 = 1
    var v2 uint8 = 42
    var v3 int64 = 42
    var v4 float64 = 1.5
    var v5 uint64 = 4294967295
    var v30 string = "a"
    v1 -= 1
    for fi31 := 0; fi31 < 2; fi31++ {
        var v32 uint32 = 143
        _ = fi31
        _ = v32
    }
    v2 = genFn1()
    for fi33 := 0; fi33 < 3; fi33++ {
        v2 -= 2
        _ = fi33
    }
    var v34 uint = uint(v5)
    for fi35 := 0; fi35 < 16; fi35++ {
        v36 := GenS2{F0: (fi35 | fi35), F1: v2}
        v3 = v3
        _ = fi35
        _ = v36
    }
    if true {
        v37 := GenS1{F0: v34}
        _ = v37
    } else {
        v1 -= v1
    }
    if true {
        var v38 int64 = v3
        _ = v38
    } else {
        v39 := []int{2147483647}
        _ = v39
    }
    v40 := GenS3{F0: (v2 + 255), F1: v1}
    v1 -= (v1 % ((v1) | 1))
    v41 := []int{-1, 2147483647}
    if false {
        v30 = ("\n\t" + v30)
    } else {
        var v42 string = (((v30 + "") + (v30 + ((v30 + (("hello" + "\n\t") + v30)) + ("hello" + "hello")))) + v30)
        switch v42 {
        case "":
            v43 := GenS2{F0: -771, F1: (v2 / (((v2 | v2)) | 1))}
            _ = v43
        case "a":
            v42 = v42
        case "hello":
            v4 = (v4 - (float64((float32(v4) + (-2.25 * float32((v4 + float64((((float32(((((v4 * v4) + (-2.25 * v4)) - v4) * 0.1)) * 1.5) + ((0.1 * (-2.25 / (1e10 / 1.0))) + (float32(((v4 + float64((float32(((v4 * ((v4 - 0.1) - v4)) / 3.14159)) * ((3.14159 - float32(v4)) + 0.1)))) * 0.1)) - (-2.25 * ((3.14159 * 1.5) + (-2.25 - 1e10)))))) / (0.0 - ((float32(float64(-2.25)) * 0.1) - (-2.25 / -2.25)))))))))) + ((v4 + v4) / v4)))
        default:
            var v44 int8 = (v1 - (^(v1)))
            _ = v44
        }
        _ = v42
    }
    v45 := GenS4{F0: 87}
    v46 := GenS1{F0: (v34 / ((2) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v40
    _ = v41
    _ = v45
    _ = v46
    _ = v5
}
