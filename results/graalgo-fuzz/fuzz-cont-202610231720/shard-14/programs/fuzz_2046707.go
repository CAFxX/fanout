package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int32 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 int64; F2 uint8 }
type GenS5 struct { F0 int16; F1 int32 }

func main() {
    v1 := []string{("" + "hello"), ("foo bar" + (("\n\t" + "") + "a"))}
    v2 := make(map[int]uint32)
    var v3 uint = 2
    var v4 int = 722
    v5 := []string{}
    var v6 bool = true
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            if ((!(!v6)) && v6) {
                var v9 int64 = -9223372036854775808
                switch "" {
                case "":
                    fmt.Println(v3, fi7, fi7)
                case "hello":
                    _ = len(v5)
                case "foo bar":
                    delete(v2, fi8)
                default:
                    fmt.Println(fi8, v6)
                }
                v10 := v2[(fi7 * v4)]
                _ = v10
                _ = v9
            }
            switch (^(fi8)) {
            case -1:
                if v6 {
                    _ = len(v5)
                    v11 := GenS1{F0: 4294967295, F1: "hello", F2: 555}
                    v4 += -899
                    var v12 int32 = 0
                    v3 = v3
                    _ = v11
                    _ = v12
                } else {
                    fmt.Println(fi8, fi7)
                    var v13 uint = v3
                    _ = v13
                }
            case 728:
                var v14 uint = (v3 ^ 65535)
                _ = v14
            case 752:
                v4 = 32768
            default:
                v15 := GenS1{F0: 2, F1: "\n\t", F2: 0}
                _ = v15
            }
            _ = fi8
        }
        v6 = v6
        v5 = append(v5, "hello")
        _ = fi7
    }
    v16 := GenS1{F0: 915, F1: "hello", F2: 2147483647}
    for fi17 := 0; fi17 < 10; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            v6 = false
            var v19 bool = (!(("hello" + "\n\t") != "hello"))
            v19 = ((v3 << (((v3 << ((v3) & 63))) & 63)) <= v3)
            _ = fi18
            _ = v19
        }
        v4 -= (v4 / ((fi17) | 1))
        _ = fi17
    }
    var v20 uint = (v3 & (v3 % (((v3 >> (((v3 * (v3 ^ (^(v3))))) & 63))) | 1)))
    if v6 {
        if (!v6) {
            var v21 float32 = (0.0 + 3.14159)
            v20 -= (v3 | (^(v3)))
            v4 -= v4
            _ = v21
        }
        if true {
            var v22 int8 = -74
            _ = v22
        } else {
            if (-917 != 1) {
                var v23 int16 = -1
                _ = v23
            } else {
                v4 = v4
                if len(v5) > 0 {
                    v5[0] = "hello"
                }
                v6 = ((((((v4 - 1) <= v4) || v6) || (!(-128 > -16))) || v6) || v6)
            }
            _ = len(v1)
            v4 += (-(v4))
        }
    } else {
        for fi24 := 0; fi24 < 16; fi24++ {
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 uint16 = 24
                switch 42 {
                case 463:
                    v1 = append(v1, "hello")
                case 32768:
                    v27 := GenS1{F0: 1, F1: "hello", F2: 2147483647}
                    _ = v27
                case 1:
                    if true {
                        v28 := GenS2{F0: 65, F1: -270}
                        v1 = append(v1, "hello")
                        _ = v28
                    }
                    v29 := GenS2{F0: 1, F1: -2147483648}
                    var v30 uint64 = 4294967295
                    _ = v29
                    _ = v30
                }
                fmt.Println(fi25)
                _ = fi25
                _ = v26
            }
            _ = fi24
        }
        fmt.Println(v20, v20)
        v31 := GenS2{F0: 255, F1: -43}
        _ = v31
    }
    v6 = v6
    v3 = 1
    fmt.Println(v3)
    fmt.Println(v4, v20, v3)
    switch ("foo bar" + "a") {
    case "foo bar":
        var v32 uint16 = 2
        _ = v32
    }
    var v33 int = (^(v4))
    switch ("hello" + ("hello" + "hello")) {
    case "":
        switch v4 {
        case -789:
            var v34 int16 = 71
            _ = v34
        case -1:
            v35 := GenS2{F0: 2, F1: -2147483648}
            _ = v35
        case -208:
            if ((float64((-2.25 / ((1e10 - 0.1) - 0.0))) == 1.5) || (!true)) {
                v36 := v2[(-(v33))]
                var v37 int32 = 2147483647
                for fi38 := 0; fi38 < 7; fi38++ {
                    v39 := GenS3{F0: "foo bar"}
                    fmt.Println(v4)
                    _ = fi38
                    _ = v39
                }
                _ = v36
                _ = v37
            } else {
                v3 = v3
            }
        default:
            v40 := GenS4{F0: "", F1: 0, F2: 173}
            _ = v40
        }
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        fmt.Println(v6)
        if ((!v6) || (127 >= -84)) {
            if (100 < -1) {
                for fi42 := 0; fi42 < 4; fi42++ {
                    var v43 string = "\n\t"
                    var v44 uint = v20
                    v44 += v20
                    _ = fi42
                    _ = v43
                    _ = v44
                }
            } else {
                v45 := GenS5{F0: 1, F1: 1}
                _ = v45
            }
            delete(v2, fi41)
        } else {
            v46 := GenS3{F0: "hello"}
            _ = v46
        }
        if false {
            var v47 float32 = 1e10
            _ = v47
        }
        _ = fi41
    }
    _ = len(v5)
    v4 -= 630
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
