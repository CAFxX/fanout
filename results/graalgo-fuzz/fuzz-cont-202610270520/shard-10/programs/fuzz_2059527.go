package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint8 }
func genFn0(p0 uint8) uint {
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(p0, fi7)
        var v8 float64 = 0.0
        var v9 uint16 = 65535
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 += (p0 * (p0 + (p0 * 120)))
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 bool = true
            switch fi11 {
            case 127:
                var v13 int64 = 1
                _ = v13
            default:
                var v14 int32 = 42
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        fmt.Println(p0, p0)
        _ = fi10
    }
    var v15 float64 = (1e10 + 0.0)
    _ = p0
    _ = v15
    return 612
}

func main() {
    var v1 uint16 = 32767
    var v2 uint64 = 132
    var v3 int16 = 32767
    var v4 uint8 = 98
    v5 := make(map[string]string)
    v6 := []string{}
    if (int8(v3) == -91) {
        v1 -= 911
    }
    switch 32768 {
    case 42:
        var v16 string = "foo bar"
        _ = v16
    case 128:
        v6 = append(v6, "\n\t")
    case 127:
        if (!true) {
            fmt.Println(v1)
            for fi17 := 0; fi17 < 1; fi17++ {
                if true {
                    var v18 int32 = -43
                    var v19 uint16 = (^(v1))
                    var v20 uint16 = (v19 + v19)
                    _ = v18
                    _ = v19
                    _ = v20
                }
                _ = fi17
            }
        } else {
            var v21 uint64 = v2
            delete(v5, "\n\t")
            _ = len(v6)
            _ = v21
        }
    }
    var v22 uint16 = v1
    var v23 uint32 = 1
    for fi24 := 0; fi24 < 13; fi24++ {
        if ((^(v22)) == 2) {
            var v25 float32 = -2.25
            var v26 uint64 = 798
            _ = v25
            _ = v26
        } else {
            v23 = v23
        }
        if (false && (fi24 < (-(fi24)))) {
            if (!false) {
                v27 := GenS1{F0: 0, F1: uint(v23), F2: v4}
                _ = v27
            } else {
                var v28 string = ""
                _ = v28
            }
            for fi29 := 0; fi29 < 16; fi29++ {
                if (!(((0.0 / 1.5) - (0.0 / 1.0)) <= 1.5)) {
                    v2 -= v2
                    var v30 int = -8
                    var v31 uint32 = 718
                    _ = v30
                    _ = v31
                }
                _ = fi29
            }
        } else {
            if (int64(v23) <= 42) {
                v32 := GenS1{F0: -1, F1: genFn0(70), F2: 42}
                switch fi24 {
                case -787:
                    _ = len(v6)
                case -60:
                    var v33 int64 = -201
                    _ = v33
                default:
                    var v34 bool = (1 != genFn0(127))
                    _ = v34
                }
                for fi35 := 0; fi35 < 4; fi35++ {
                    var v36 int16 = v3
                    var v37 uint32 = 759
                    if (846 >= 9223372036854775807) {
                        v38 := GenS1{F0: 42, F1: 65535, F2: (v4 & 34)}
                        fmt.Println(v23)
                        _ = v38
                    } else {
                        v1 = v1
                    }
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                _ = v32
            } else {
                fmt.Println(v1, v4)
                var v39 uint = 1
                _ = v39
            }
        }
        _ = fi24
    }
    v3 = (v3 / ((v3) | 1))
    fmt.Println(v23)
    for fi40 := 0; fi40 < 9; fi40++ {
        v22 = v1
        v22 += 1
        _ = fi40
    }
    var v41 int = 4294967295
    for fi42 := 0; fi42 < 7; fi42++ {
        v23 += (^(v23))
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
