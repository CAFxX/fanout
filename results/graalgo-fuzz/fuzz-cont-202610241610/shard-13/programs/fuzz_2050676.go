package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint64 }
func genFn0(p0 int8, p1 int16) uint8 {
    if ((1e10 + float32((0.1 * (-2.25 + 0.1)))) < -2.25) {
        v5 := []int{2147483647, 128, 128}
        p1 += p1
        p1 += p1
        _ = v5
    } else {
        p1 = p1
        if (true && (1e10 <= 0.1)) {
            switch 32767 {
            case 0:
                var v6 bool = true
                _ = v6
            case 32768:
                var v7 uint8 = 127
                _ = v7
            }
        } else {
            v8 := []int{2147483647, 1, 2147483647}
            p1 += (-(p1))
            if true {
                switch 1 {
                case -2147483648:
                    v9 := GenS1{F0: 32767, F1: 1}
                    _ = v9
                case 73:
                    fmt.Println(p0)
                default:
                    _ = len(v8)
                }
                if (true || (("" + "foo bar") < "a")) {
                    var v10 float64 = 1e10
                    for fi11 := 0; fi11 < 6; fi11++ {
                        v12 := GenS1{F0: 402, F1: 1}
                        v10 = v10
                        _ = fi11
                        _ = v12
                    }
                    _ = v10
                } else {
                    var v13 uint16 = 32767
                    _ = v13
                }
                var v14 bool = (!false)
                _ = v14
            } else {
                for fi15 := 0; fi15 < 20; fi15++ {
                    for fi16 := 0; fi16 < 16; fi16++ {
                        switch fi16 {
                        case 42:
                            var v17 uint8 = 154
                            _ = v17
                        case 1:
                            v8 = append(v8, 0)
                        default:
                            p1 += (p1 >> ((18446744073709551615) & 63))
                        }
                        _ = fi16
                    }
                    var v18 float32 = 0.1
                    _ = fi15
                    _ = v18
                }
            }
            _ = v8
        }
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        v20 := GenS1{F0: 65535, F1: 42}
        var v21 int16 = (p1 + 1000)
        if ((1e10 * 3.14159) <= 0.1) {
            if (p0 >= 100) {
                p0 = (p0 & (p0 | (p0 / ((p0) | 1))))
            }
            for fi22 := 0; fi22 < 3; fi22++ {
                p0 += p0
                _ = fi22
            }
        } else {
            for fi23 := 0; fi23 < 5; fi23++ {
                v24 := []int{}
                if (((((1e10 - (0.1 - 3.14159)) * -2.25) + ((0.0 - 1.5) + (-2.25 - 0.0))) - 3.14159) != 3.14159) {
                    p1 = (v21 << ((339) & 63))
                    v25 := GenS1{F0: 32767, F1: 1}
                    var v26 int8 = p0
                    _ = v25
                    _ = v26
                } else {
                    v27 := GenS1{F0: 0, F1: 941}
                    v28 := GenS1{F0: 42, F1: 42}
                    var v29 uint64 = 1
                    _ = v27
                    _ = v28
                    _ = v29
                }
                if false {
                    switch (("foo bar" + "\n\t") + "\n\t") {
                    case "":
                        var v30 int16 = v21
                        _ = v30
                    case "\n\t":
                        if len(v24) > 0 {
                            v24[0] = fi19
                        }
                    default:
                        v31 := GenS1{F0: 32767, F1: 42}
                        _ = v31
                    }
                    v21 -= p1
                } else {
                    v32 := GenS1{F0: 2, F1: 42}
                    v33 := GenS1{F0: 42, F1: 840}
                    v34 := GenS1{F0: 413, F1: 375}
                    _ = v32
                    _ = v33
                    _ = v34
                }
                _ = fi23
                _ = v24
            }
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    if (42 > 9223372036854775807) {
        var v35 int64 = -9223372036854775808
        _ = v35
    }
    _ = p0
    _ = p1
    return 234
}

func main() {
    var v1 bool = false
    v2 := make(map[string]string)
    var v3 uint8 = 127
    var v4 uint8 = 255
    v36 := GenS2{F0: int32(v3), F1: 1}
    var v37 int8 = -128
    var v38 uint64 = 2
    v4 = 177
    fmt.Println(v38, v1)
    v1 = v1
    v39 := []int{-2147483648, -362, 1, -454}
    for fi40 := 0; fi40 < 16; fi40++ {
        _ = len(v39)
        v41 := GenS2{F0: -2147483648, F1: v38}
        _ = fi40
        _ = v41
    }
    if v1 {
        var v42 string = ("\n\t" + "hello")
        v2[((v42 + v42) + v42)] = ("" + ((v42 + v42) + v42))
        _ = v42
    }
    _ = len(v39)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
