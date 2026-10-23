package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int8 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 int8; F1 int }

func main() {
    var v1 uint = 42
    v2 := make(map[string]uint32)
    v3 := make(map[string]uint32)
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := []int{fi4, fi4, (fi4 - (fi4 + fi4)), fi4}
        var v6 uint = (v1 * v1)
        if true {
            fmt.Println(v6, fi4)
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        switch "foo bar" {
        case "":
            var v8 int64 = -114
            _ = v8
        case "a":
            var v9 uint16 = uint16(v1)
            _ = v9
        case "\n\t":
            v1 = 65535
        default:
            if false {
                for fi10 := 0; fi10 < 8; fi10++ {
                    v11 := []int{fi7, (fi10 % ((4294967295) | 1))}
                    _ = fi10
                    _ = v11
                }
                var v12 string = ("a" + (("hello" + (("hello" + "\n\t") + "")) + ("foo bar" + "")))
                _ = v12
            } else {
                for fi13 := 0; fi13 < 4; fi13++ {
                    if (951 > 65535) {
                        var v14 int32 = 0
                        _ = v14
                    } else {
                        fmt.Println(v1, v1)
                    }
                    v15 := []int{(fi7 + (-(fi13))), 4294967295, 127, fi13}
                    _ = fi13
                    _ = v15
                }
                var v16 uint8 = 128
                if (!true) {
                    v17 := GenS1{F0: 82, F1: 215, F2: 42}
                    _ = v17
                } else {
                    switch (fi7 * (fi7 >> (((v1 | 4294967295)) & 63))) {
                    case 219:
                        var v18 uint32 = 851
                        _ = v18
                    case -718:
                        if false {
                            v19 := GenS2{F0: -2147483648, F1: v1}
                            v20 := GenS1{F0: 0, F1: -9223372036854775808, F2: -128}
                            var v21 int16 = -585
                            _ = v19
                            _ = v20
                            _ = v21
                        } else {
                            for fi22 := 0; fi22 < 4; fi22++ {
                                v16 -= v16
                                var v23 int32 = -2147483648
                                _ = fi22
                                _ = v23
                            }
                            var v24 uint = v1
                            v25 := GenS3{F0: v1, F1: -1}
                            _ = v24
                            _ = v25
                        }
                    default:
                        var v26 uint16 = uint16(v1)
                        _ = v26
                    }
                }
                _ = v16
            }
        }
        var v27 int8 = 100
        v27 += v27
        _ = fi7
        _ = v27
    }
    if false {
        var v28 int = 42
        for fi29 := 0; fi29 < 1; fi29++ {
            fmt.Println(fi29)
            _ = fi29
        }
        switch v28 {
        case 4294967295:
            v30 := []int{(-(v28))}
            _ = v30
        case -857:
            v31 := GenS4{F0: "", F1: (!((2 < 2) || (float64(1e10) >= ((((-2.25 + float64(0.0)) / 3.14159) + 1e10) * (float64(1e10) + 0.1)))))}
            _ = v31
        case 654:
            v28 -= (v28 - 4294967295)
        }
        _ = v28
    } else {
        if false {
            var v32 int32 = 42
            _ = v32
        }
        var v33 float32 = float32((float64(float32(((1.5 - (((3.14159 / -2.25) + 0.0) / (1.5 * 3.14159))) * 1.5))) - (1e10 * ((0.1 + 1.5) - -2.25))))
        _ = v33
    }
    switch ((("" + ("foo bar" + "")) + "a") + "a") {
    case "foo bar":
        v1 = (^(v1))
    }
    v34 := v3[("foo bar" + "foo bar")]
    v34 += (v34 * 65535)
    var v35 int64 = 0
    fmt.Println(v1, v35, v35)
    var v36 int32 = 42
    v36 = (v36 ^ v36)
    for fi37 := 0; fi37 < 7; fi37++ {
        v38 := GenS5{F0: 0, F1: fi37}
        v1 -= (v1 % (((v1 >> ((v1) & 63))) | 1))
        _ = fi37
        _ = v38
    }
    var v39 int8 = -128
    if (!false) {
        v35 = int64(v34)
        for fi40 := 0; fi40 < 19; fi40++ {
            fmt.Println(v39, v35)
            _ = fi40
        }
        v39 = (^(v39))
    } else {
        v39 -= (v39 | v39)
        var v41 int = -9223372036854775808
        for fi42 := 0; fi42 < 16; fi42++ {
            v34 -= v34
            v43 := GenS4{F0: (("\n\t" + "a") + (("" + "\n\t") + "")), F1: (!true)}
            _ = fi42
            _ = v43
        }
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v39
}
