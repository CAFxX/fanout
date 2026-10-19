package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 bool }
func genFn0(p0 int) int16 {
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            switch fi8 {
            case 0:
                var v9 string = ("foo bar" + ("foo bar" + "foo bar"))
                _ = v9
            }
            v10 := []int{-2147483648, 128, fi7, (fi7 & (p0 ^ fi8))}
            _ = fi8
            _ = v10
        }
        _ = fi7
    }
    if true {
        var v11 float64 = 0.1
        switch "hello" {
        case "\n\t":
            var v12 bool = (-55 == 0)
            _ = v12
        case "foo bar":
            if false {
                p0 -= p0
                switch p0 {
                case -221:
                    p0 -= (p0 | (^(p0)))
                default:
                    for fi13 := 0; fi13 < 19; fi13++ {
                        v14 := GenS1{F0: true}
                        _ = fi13
                        _ = v14
                    }
                }
                p0 += 848
            }
        default:
            p0 -= p0
        }
        v11 += v11
        _ = v11
    } else {
        if ((3.14159 * ((float32((0.1 + ((0.1 + 1.5) * 3.14159))) * 1.5) / 1e10)) <= (0.0 - 0.0)) {
            v15 := []int{p0, -2}
            var v16 string = "hello"
            _ = v15
            _ = v16
        } else {
            if true {
                fmt.Println(p0)
                if (1 != 1) {
                    fmt.Println(p0)
                } else {
                    v17 := []int{p0}
                    _ = v17
                }
                v18 := GenS2{F0: 2, F1: 0}
                _ = v18
            } else {
                for fi19 := 0; fi19 < 9; fi19++ {
                    p0 = (-(p0))
                    v20 := GenS3{F0: -612}
                    _ = fi19
                    _ = v20
                }
                if (false && ((true && (255 <= 514)) || false)) {
                    v21 := GenS4{F0: 4294967295}
                    _ = v21
                }
                for fi22 := 0; fi22 < 14; fi22++ {
                    var v23 int64 = -306
                    for fi24 := 0; fi24 < 13; fi24++ {
                        if (((18446744073709551615 <= 42) || (true && (1.5 < (1.5 * 1e10)))) && ((1.5 / (1.5 / (1e10 / 1e10))) >= 1e10)) {
                            v23 += 1
                            v25 := []int{}
                            _ = v25
                        }
                        v23 = (^(v23))
                        v26 := []int{p0, 286, p0}
                        _ = fi24
                        _ = v26
                    }
                    _ = fi22
                    _ = v23
                }
            }
        }
        var v27 uint = 1
        _ = v27
    }
    if (!true) {
        var v28 int16 = 32767
        _ = v28
    }
    p0 = -8
    _ = p0
    return -353
}

func main() {
    var v1 bool = (189 >= 255)
    var v2 uint32 = 65535
    var v3 int32 = -2147483648
    var v4 string = "hello"
    v5 := make(map[string]string)
    var v6 uint64 = 2
    v29 := GenS5{F0: 127, F1: 100}
    if (!v1) {
        var v30 float64 = float64(1e10)
        _ = v30
    }
    v31 := GenS3{F0: -2147483648}
    var v32 uint16 = 65535
    var v33 int32 = 834
    v3 -= (v33 - (v3 ^ v3))
    var v34 int16 = -736
    var v35 float64 = ((3.14159 - 1.5) * 3.14159)
    var v36 string = ((v4 + v4) + v4)
    v37 := GenS4{F0: (v6 / (((^(v6))) | 1))}
    var v38 float32 = ((float32(0.1) + -2.25) + ((1.5 + ((1e10 + (-2.25 * -2.25)) / (1.5 * 1.5))) + (float32(0.0) * 1.5)))
    for fi39 := 0; fi39 < 10; fi39++ {
        delete(v5, "a")
        switch ("a" + (v36 + v4)) {
        case "\n\t":
            v35 += 1e10
        default:
            if v1 {
                var v40 int32 = (v3 & v33)
                var v41 int8 = 1
                _ = v40
                _ = v41
            } else {
                var v42 int8 = 127
                var v43 bool = false
                v44 := GenS6{F0: v43}
                _ = v42
                _ = v43
                _ = v44
            }
        }
        _ = fi39
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
