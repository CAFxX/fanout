package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int8; F1 int64; F2 string }
func genFn0(p0 int8) uint8 {
    var v6 float32 = -2.25
    if true {
        v7 := []int{}
        _ = v7
    } else {
        p0 = p0
    }
    _ = p0
    _ = v6
    return 255
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[string]int)
    var v3 uint16 = 1
    var v4 uint16 = 32767
    var v5 uint32 = v1
    v3 = v3
    for fi8 := 0; fi8 < 14; fi8++ {
        if (100 <= (fi8 & (^(fi8)))) {
            v3 += v4
            switch "" {
            case "a":
                v9 := GenS1{F0: 42}
                _ = v9
            }
        }
        switch (((("\n\t" + (("foo bar" + ("foo bar" + "foo bar")) + "hello")) + (("foo bar" + "foo bar") + "\n\t")) + "hello") + ("a" + "hello")) {
        case "foo bar":
            v4 = v3
        default:
            fmt.Println(v1, v5, v1)
        }
        var v10 uint32 = (v5 << ((2) & 63))
        _ = fi8
        _ = v10
    }
    var v11 float32 = float32(-2.25)
    if true {
        for fi12 := 0; fi12 < 16; fi12++ {
            v3 = v3
            _ = fi12
        }
        v3 += v3
    }
    if (!true) {
        v4 -= v3
        v3 = (^(v3))
    } else {
        v13 := GenS2{F0: 2147483648, F1: 32767}
        var v14 uint64 = 42
        v15 := GenS2{F0: 62, F1: 512}
        _ = v13
        _ = v14
        _ = v15
    }
    if (int(v1) != 2147483647) {
        if (0 != 733) {
            for fi16 := 0; fi16 < 12; fi16++ {
                for fi17 := 0; fi17 < 8; fi17++ {
                    if (2147483647 <= (v5 << ((1) & 63))) {
                        if true {
                            if false {
                                var v18 bool = (!true)
                                fmt.Println(v18, v4, v18)
                                _ = v18
                            }
                            v19 := GenS1{F0: 559}
                            var v20 bool = (-2147483648 != 42)
                            _ = v19
                            _ = v20
                        } else {
                            v1 -= (v1 << ((1) & 63))
                            fmt.Println(v5, fi17, v3)
                            v21 := []int{fi16, -2147483648}
                            _ = v21
                        }
                    } else {
                        var v22 float32 = v11
                        _ = len(v2)
                        var v23 string = ("hello" + ("hello" + ""))
                        _ = v22
                        _ = v23
                    }
                    fmt.Println(fi17, fi17, fi17)
                    var v24 int16 = 1000
                    v11 -= v11
                    _ = fi17
                    _ = v24
                }
                v25 := []int{}
                _ = fi16
                _ = v25
            }
            v11 = (float32((1.5 / float64(1.5))) / 1.5)
        }
        if (genFn0(0) == genFn0(-128)) {
            v26 := []int{-2147483648}
            switch "hello" {
            case "":
                if (((!true) || false) || false) {
                    fmt.Println(v11)
                } else {
                    _ = len(v2)
                    v11 = ((((v11 * v11) / (v11 - (float32((0.1 * 0.0)) / float32((0.1 + (((((-2.25 - 1e10) / 3.14159) * 0.0) / 0.1) / ((((1e10 - (1.5 * 0.0)) * ((float64((v11 + 1e10)) + (1.5 + 1e10)) * 0.0)) - ((0.0 * float64(0.0)) * 3.14159)) * 0.1))))))) - (float32(-2.25) * float32(-2.25))) / v11)
                }
            case "foo bar":
                var v27 int = 2147483647
                _ = v27
            default:
                v11 -= ((v11 / (v11 / (float32(3.14159) + v11))) / (((v11 + v11) - (v11 * 0.0)) * -2.25))
            }
            v28 := GenS3{F0: 1, F1: -475, F2: ""}
            _ = v26
            _ = v28
        } else {
            v29 := []int{}
            fmt.Println(v3, v3)
            v4 -= v3
            _ = v29
        }
        switch 783 {
        case -862:
            v30 := GenS3{F0: 1, F1: 42, F2: "foo bar"}
            _ = v30
        case 128:
            for fi31 := 0; fi31 < 7; fi31++ {
                switch (("" + "a") + "foo bar") {
                case "a":
                    v32 := []int{int(v5), (^(fi31))}
                    _ = v32
                case "\n\t":
                    v11 -= 0.0
                case "hello":
                    var v33 int8 = -49
                    _ = v33
                default:
                    v34 := GenS3{F0: 0, F1: 610, F2: "\n\t"}
                    _ = v34
                }
                _ = fi31
            }
        case -903:
            v11 -= v11
        }
    } else {
        v4 += (v4 | 1)
        var v35 uint8 = 128
        _ = v35
    }
    v36 := []int{128, 32767}
    if false {
        var v37 int16 = -961
        _ = v37
    } else {
        v11 += (v11 * (v11 + 0.1))
    }
    fmt.Println(v5, v5)
    v1 -= v1
    v38 := GenS1{F0: 0}
    fmt.Println(v11, v1)
    v4 = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v4
    _ = v5
}
