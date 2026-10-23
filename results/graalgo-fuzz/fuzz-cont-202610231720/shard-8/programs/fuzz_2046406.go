package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint64; F1 int; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint8, p1 uint) int32 {
    v5 := GenS1{F0: "hello", F1: -147}
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 int = fi6
        var v8 bool = (1000 == 0)
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 string = "\n\t"
            _ = fi9
            _ = v10
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    return -1
}
func genFn1(p0 uint, p1 uint32) uint8 {
    v11 := []int{}
    v12 := GenS1{F0: "a", F1: 1}
    p1 -= (p1 * (p1 << (((p0 - uint(p1))) & 63)))
    if ((((((1e10 + 1.5) - 3.14159) * -2.25) - (1e10 * 0.0)) != (0.0 * -2.25)) && (uint8(p0) < 42)) {
        if ((2 == 708) && false) {
            switch -9223372036854775808 {
            case 32768:
                if (32767 < 508) {
                    var v13 uint32 = 2147483648
                    v13 -= (v13 & (v13 + p1))
                    _ = v13
                }
            case 0:
                var v14 int32 = 42
                _ = v14
            }
            fmt.Println(p0)
        } else {
            for fi15 := 0; fi15 < 12; fi15++ {
                for fi16 := 0; fi16 < 1; fi16++ {
                    if true {
                        fmt.Println(p1, p1, p1)
                    }
                    _ = len(v11)
                    _ = fi16
                }
                v17 := GenS1{F0: "a", F1: -1}
                _ = fi15
                _ = v17
            }
            var v18 string = ("foo bar" + "\n\t")
            v19 := GenS2{F0: 204, F1: 32768, F2: 128}
            _ = v18
            _ = v19
        }
    }
    var v20 uint = (^(p0))
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v20
    return 0
}

func main() {
    var v1 bool = true
    var v2 int = 1
    v3 := make(map[string]string)
    var v4 int = 0
    v4 = v4
    v3[("hello" + ("hello" + "foo bar"))] = ("foo bar" + "a")
    var v21 int16 = -1
    for fi22 := 0; fi22 < 1; fi22++ {
        switch 4294967295 {
        case 9223372036854775807:
            for fi23 := 0; fi23 < 17; fi23++ {
                if v1 {
                    fmt.Println(fi23)
                    v4 += v4
                    if (3.14159 < -2.25) {
                        if (975 < 834) {
                            if (!(!("hello" <= ""))) {
                                _ = len(v3)
                                v21 = v21
                            }
                        }
                        var v24 int8 = -87
                        _ = v24
                    } else {
                        var v25 uint32 = 2
                        _ = v25
                    }
                    var v26 uint8 = 121
                    _ = v26
                }
                _ = fi23
            }
        default:
            if ((v1 && ((-128 == 109) || v1)) || v1) {
                v27 := GenS2{F0: 42, F1: 2147483647, F2: 128}
                var v28 uint16 = 0
                v29 := GenS2{F0: 2, F1: v4, F2: 253}
                _ = v27
                _ = v28
                _ = v29
            } else {
                v30 := v3["hello"]
                for fi31 := 0; fi31 < 6; fi31++ {
                    var v32 int16 = v21
                    _ = fi31
                    _ = v32
                }
                _ = v30
            }
        }
        if (v1 && true) {
            var v33 uint8 = 255
            switch ("" + "foo bar") {
            case "a":
                fmt.Println(fi22, v2)
            case "hello":
                v34 := GenS2{F0: 2, F1: v4, F2: v33}
                _ = v34
            default:
                v33 = genFn1(65535, 2147483648)
            }
            _ = v33
        } else {
            fmt.Println(v21, v21, v1)
        }
        v35 := GenS1{F0: (("hello" + ("foo bar" + ("foo bar" + ""))) + "a"), F1: 42}
        _ = fi22
        _ = v35
    }
    if v1 {
        var v36 float32 = 1.5
        _ = v36
    }
    if true {
        switch "hello" {
        case "":
            v2 += (v4 % ((v2) | 1))
        case "\n\t":
            var v37 int16 = (v21 ^ 935)
            _ = v37
        default:
            fmt.Println(v4, v2)
        }
        switch ("\n\t" + ("" + "a")) {
        case "a":
            fmt.Println(v4, v2)
        case "foo bar":
            v38 := GenS3{F0: 0, F1: 2147483647}
            _ = v38
        case "\n\t":
            fmt.Println(v2)
        }
        v39 := v3[((("\n\t" + "hello") + "\n\t") + "a")]
        _ = v39
    } else {
        var v40 int16 = v21
        if v1 {
            var v41 uint16 = 42
            var v42 float64 = 1e10
            fmt.Println(v42, v41, v21)
            _ = v41
            _ = v42
        }
        fmt.Println(v1, v1)
        _ = v40
    }
    v43 := []int{v2, (v2 & (v4 & (v4 | -1)))}
    delete(v3, ((((("" + "") + "") + ("foo bar" + ("hello" + ("" + ("" + ""))))) + ((("hello" + ("hello" + "")) + "foo bar") + "a")) + (("hello" + "hello") + "foo bar")))
    if v1 {
        v44 := GenS2{F0: 42, F1: -350, F2: 127}
        v21 -= v21
        _ = v44
    }
    fmt.Println(v4, v2)
    var v45 string = ""
    if (874 > genFn0(167, 4294967295)) {
        for fi46 := 0; fi46 < 17; fi46++ {
            delete(v3, ((v45 + v45) + v45))
            v21 = v21
            _ = fi46
        }
        v47 := GenS3{F0: 42, F1: 422}
        _ = v47
    } else {
        for fi48 := 0; fi48 < 9; fi48++ {
            v45 = ("" + (((((v45 + ("foo bar" + v45)) + "foo bar") + ("hello" + v45)) + v45) + "\n\t"))
            v43 = append(v43, (-(v2)))
            v21 = 35
            _ = fi48
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v43
    _ = v45
}
