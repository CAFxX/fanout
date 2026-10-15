package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int16; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int16; F1 uint64; F2 int8 }
func genFn0() int32 {
    for fi4 := 0; fi4 < 11; fi4++ {
        if (266 > 828) {
            var v5 int8 = -80
            _ = v5
        }
        for fi6 := 0; fi6 < 11; fi6++ {
            if (78 <= 244) {
                var v7 uint64 = 1
                _ = v7
            }
            _ = fi6
        }
        _ = fi4
    }
    v8 := GenS1{F0: 364}
    _ = v8
    return -1
}
func genFn1(p0 uint16, p1 int64) uint32 {
    p1 -= (p1 >> ((18446744073709551615) & 63))
    switch ((((("\n\t" + "hello") + "foo bar") + ("" + "hello")) + "foo bar") + "a") {
    case "\n\t":
        p0 = p0
    case "":
        p1 += p1
    default:
        v9 := GenS1{F0: 1}
        _ = v9
    }
    var v10 uint = 255
    _ = p0
    _ = p1
    _ = v10
    return 717
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 255
    var v3 uint = 251
    if true {
        fmt.Println(v3, v3)
    }
    if (!false) {
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 float32 = -2.25
            var v13 bool = (((((!true) && (!true)) && ((("foo bar" + "hello") + "\n\t") > "foo bar")) && (v12 >= v12)) && ((!((-9223372036854775808 > -179) || true)) || true))
            _ = fi11
            _ = v12
            _ = v13
        }
        var v14 float64 = (0.1 - float64(((1.5 / 1.0) - 0.0)))
        _ = v14
    } else {
        switch 32768 {
        case 9223372036854775807:
            v3 = 484
        case -9223372036854775808:
            v2 -= v2
        default:
            for fi15 := 0; fi15 < 6; fi15++ {
                if false {
                    var v16 int32 = 42
                    if true {
                        var v17 bool = ((!false) || (-1 >= 89))
                        _ = v17
                    } else {
                        for fi18 := 0; fi18 < 12; fi18++ {
                            v19 := []int{(fi15 | fi15), -918, fi15}
                            var v20 uint64 = 1
                            var v21 float64 = 1e10
                            _ = fi18
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                    }
                    if false {
                        var v22 int8 = 1
                        v23 := GenS2{F0: -726, F1: 32767, F2: 921}
                        _ = v22
                        _ = v23
                    }
                    _ = v16
                }
                var v24 float32 = -2.25
                switch ("" + ("foo bar" + "hello")) {
                case "\n\t":
                    v25 := v1[("" + "\n\t")]
                    _ = v25
                case "hello":
                    if true {
                        switch fi15 {
                        case -1:
                            v24 -= (v24 / 1.5)
                        case -505:
                            switch (-(fi15)) {
                            case 42:
                                fmt.Println(v3, fi15)
                            default:
                                v2 -= (^(v2))
                            }
                        }
                    } else {
                        var v26 int64 = -115
                        _ = v26
                    }
                }
                _ = fi15
                _ = v24
            }
        }
        if false {
            v27 := []int{32768}
            _ = v27
        }
        fmt.Println(v3, v2)
    }
    v28 := v1["\n\t"]
    for fi29 := 0; fi29 < 4; fi29++ {
        var v30 uint16 = 664
        for fi31 := 0; fi31 < 16; fi31++ {
            var v32 int16 = -1
            if v28 {
                v30 += 1
                v30 = (v30 & (v30 >> (((v3 * 420)) & 63)))
                v30 = v30
            }
            delete(v1, "\n\t")
            var v33 float64 = (0.1 * (1e10 - 0.1))
            v32 = (v32 & (v32 + int16(v3)))
            v3 = v3
            v34 := GenS3{F0: v32, F1: -122, F2: v2}
            if ((v3 + (^(v3))) != (v3 >> (((v3 ^ 18446744073709551615)) & 63))) {
                var v35 uint32 = 2147483648
                switch ((("\n\t" + "") + ("\n\t" + "hello")) + "a") {
                case "\n\t":
                    v36 := GenS4{F0: uint(v35), F1: -2147483648}
                    _ = v36
                case "a":
                    if (((fi29 - fi29) > (fi29 ^ fi31)) && v28) {
                        fmt.Println(v32, v28)
                        v37 := GenS3{F0: v32, F1: 0, F2: v2}
                        _ = v37
                    }
                default:
                    v38 := GenS1{F0: (^(v32))}
                    _ = v38
                }
                var v39 int32 = 42
                v40 := GenS5{F0: 731}
                _ = v35
                _ = v39
                _ = v40
            }
            v30 -= (v30 | v30)
            v41 := []int{(fi31 - fi29), (fi29 << (((v3 % ((v3) | 1))) & 63)), fi31}
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
            _ = v41
        }
        _ = fi29
        _ = v30
    }
    var v42 float64 = 3.14159
    v43 := GenS1{F0: 127}
    v44 := GenS6{F0: 303, F1: 4294967295, F2: -80}
    var v45 int16 = -1
    v46 := GenS5{F0: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
