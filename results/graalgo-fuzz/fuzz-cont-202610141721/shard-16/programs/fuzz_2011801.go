package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 int8; F2 int }
func genFn0(p0 uint16) bool {
    var v4 int16 = -551
    var v5 uint32 = 2
    v6 := GenS1{F0: true, F1: -86}
    v7 := GenS1{F0: true, F1: 100}
    switch 0 {
    case -1:
        v4 -= (v4 & -32768)
    default:
        var v8 string = "hello"
        _ = v8
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return false
}
func genFn1(p0 int16, p1 int64) int32 {
    var v9 float64 = 3.14159
    if ((214 <= 255) && false) {
        var v10 float32 = float32(0.0)
        var v11 string = "a"
        switch ((v11 + (v11 + v11)) + ("foo bar" + (("a" + v11) + v11))) {
        case "":
            var v12 float64 = (float64(v10) / ((v9 * ((-2.25 - (((float64(v10) - float64(v10)) / ((float64(v10) / v9) + v9)) + 1.5)) / v9)) + (float64(v10) + v9)))
            _ = v12
        default:
            v13 := []int{127, -1}
            _ = v13
        }
        _ = v10
        _ = v11
    } else {
        if false {
            switch ("hello" + ("" + ("hello" + "foo bar"))) {
            case "hello":
                v14 := []int{9, 0, -813, 1}
                _ = v14
            case "":
                v15 := []int{-2147483648, 163, -9223372036854775808, -1}
                _ = v15
            default:
                var v16 int32 = 1
                _ = v16
            }
            var v17 int = -2147483648
            _ = v17
        }
    }
    v18 := GenS2{F0: 1, F1: 209, F2: 0}
    v19 := []int{}
    switch -2147483648 {
    case -272:
        var v20 int32 = 861
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    _ = v9
    return -2147483648
}

func main() {
    var v1 uint8 = 68
    var v2 string = (("\n\t" + "a") + "\n\t")
    v3 := make(map[int]string)
    for fi21 := 0; fi21 < 20; fi21++ {
        v2 = v2
        v2 = v2
        _ = fi21
    }
    v22 := GenS1{F0: ((v1 / (((v1 ^ 1)) | 1)) <= v1), F1: -91}
    v23 := []int{666, 2147483647, 128}
    v24 := GenS2{F0: 0, F1: 455, F2: 0}
    v25 := GenS3{F0: -47}
    for fi26 := 0; fi26 < 5; fi26++ {
        switch fi26 {
        case 2147483647:
            v2 = (v2 + v2)
        default:
            v1 = (^(v1))
        }
        v3[fi26] = ("a" + ("\n\t" + (v2 + (v2 + (v2 + ((v2 + ("" + (((((v2 + "") + "\n\t") + ("" + "\n\t")) + "") + ((v2 + v2) + v2)))) + v2))))))
        v1 -= 128
        _ = fi26
    }
    fmt.Println(v1)
    if len(v23) > 0 {
        v23[0] = -91
    }
    if true {
        var v27 string = "a"
        var v28 uint16 = 603
        v29 := GenS3{F0: -128}
        _ = v27
        _ = v28
        _ = v29
    }
    switch -916 {
    case -1:
        var v30 string = v2
        _ = v30
    case -2147483648:
        for fi31 := 0; fi31 < 13; fi31++ {
            if false {
                v1 = v1
            } else {
                v1 = 255
                if false {
                    if (((uint16(v1) != 42) && false) && (!(42 > 669))) {
                        if len(v23) > 0 {
                            v23[0] = (-(fi31))
                        }
                    }
                    switch (v2 + v2) {
                    case "":
                        v32 := GenS2{F0: 2, F1: 0, F2: 284}
                        _ = v32
                    case "a":
                        v2 = v2
                    case "foo bar":
                        fmt.Println(v2, v2)
                    }
                }
            }
            for fi33 := 0; fi33 < 6; fi33++ {
                v2 = v2
                switch 128 {
                case 875:
                    switch v2 {
                    case "\n\t":
                        v34 := GenS4{F0: 278, F1: 42, F2: (-(fi31))}
                        _ = v34
                    }
                case -9223372036854775808:
                    v35 := GenS4{F0: 18446744073709551615, F1: 0, F2: -418}
                    _ = v35
                default:
                    v2 = v2
                }
                v36 := GenS1{F0: false, F1: 34}
                _ = fi33
                _ = v36
            }
            _ = fi31
        }
    }
    v2 = v2
    for fi37 := 0; fi37 < 17; fi37++ {
        v2 = "hello"
        v38 := GenS4{F0: 4294967295, F1: 68, F2: 2147483647}
        _ = fi37
        _ = v38
    }
    var v39 int16 = 1000
    for fi40 := 0; fi40 < 7; fi40++ {
        delete(v3, fi40)
        _ = fi40
    }
    for fi41 := 0; fi41 < 15; fi41++ {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        var v42 int = (-(fi41))
        _ = fi41
        _ = v42
    }
    var v43 float64 = 3.14159
    fmt.Println(v1, v1, v39)
    v39 = v39
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v39
    _ = v43
}
