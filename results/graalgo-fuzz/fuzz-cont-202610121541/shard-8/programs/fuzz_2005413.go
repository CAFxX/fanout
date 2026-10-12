package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int32 }
type GenS2 struct { F0 int; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint32; F1 string; F2 uint32 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS5 struct { F0 int16; F1 uint8; F2 string }
type GenS6 struct { F0 bool; F1 int64; F2 int32 }
type GenS7 struct { F0 string; F1 bool }
func genFn0() uint32 {
    if (!((-188 <= 32767) && false)) {
        var v6 float32 = 0.0
        v6 += 3.14159
        v6 = ((v6 - (v6 * float32(0.0))) * (float32((1e10 - (3.14159 / float64((v6 + v6))))) / v6))
        _ = v6
    } else {
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := []int{}
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 uint32 = 0
                if true {
                    v11 := GenS1{F0: -1, F1: "foo bar", F2: 42}
                    if true {
                        var v12 int16 = 32767
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        _ = v12
                    }
                    v13 := GenS2{F0: -594, F1: (0 < -2147483648), F2: 42}
                    _ = v11
                    _ = v13
                }
                _ = fi9
                _ = v10
            }
            for fi14 := 0; fi14 < 9; fi14++ {
                var v15 string = "foo bar"
                _ = fi14
                _ = v15
            }
            _ = fi7
            _ = v8
        }
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        if (-9223372036854775808 == 1) {
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 float64 = 3.14159
                if true {
                    v18 -= float64(1.5)
                    v18 = 1.5
                }
                var v19 bool = false
                _ = fi17
                _ = v18
                _ = v19
            }
            v20 := GenS1{F0: 1, F1: "hello", F2: 0}
            _ = v20
        } else {
            v21 := GenS3{F0: 465, F1: "", F2: 1}
            _ = v21
        }
        v22 := GenS4{F0: 127, F1: 465, F2: 721}
        _ = fi16
        _ = v22
    }
    var v23 string = "foo bar"
    for fi24 := 0; fi24 < 2; fi24++ {
        var v25 int64 = -9223372036854775808
        var v26 string = "\n\t"
        if false {
            switch (fi24 - 128) {
            case -1:
                var v27 int64 = 0
                _ = v27
            case -254:
                v23 = "hello"
            }
        } else {
            v28 := GenS5{F0: 1000, F1: 2, F2: ""}
            var v29 float32 = ((0.0 + (0.1 / float32((3.14159 * 3.14159)))) * (float32((((1.5 / 1.5) - 1.5) / -2.25)) - (0.0 * (3.14159 / 1.5))))
            _ = v28
            _ = v29
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    var v30 int8 = 100
    _ = v23
    _ = v30
    return 1
}

func main() {
    var v1 float64 = (0.0 / 3.14159)
    var v2 uint8 = 149
    v3 := make(map[int]uint32)
    v4 := []int{127}
    v5 := make(map[int]bool)
    v1 = v1
    if (443 >= 42) {
        if ((!((v1 / (v1 / v1)) >= 0.1)) || false) {
            v31 := GenS6{F0: false, F1: 42, F2: 741}
            _ = v31
        } else {
            var v32 uint8 = (v2 ^ v2)
            _ = v32
        }
        v33 := v3[-868]
        if (0 <= v33) {
            fmt.Println(v33, v2)
        }
        _ = v33
    } else {
        v2 += (v2 + (v2 | 128))
    }
    var v34 int16 = 32767
    var v35 int64 = -9223372036854775808
    v34 = 42
    v2 += v2
    var v36 int32 = 617
    if ((1.5 > 3.14159) && (int(v36) != -1)) {
        if (268 != 555) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v2 += 127
        }
        _ = len(v3)
        switch 1 {
        case 128:
            if (-32768 <= (^(v34))) {
                var v37 uint64 = 646
                var v38 uint = 153
                v39 := GenS6{F0: (2 != (^(v2))), F1: (v35 & v35), F2: v36}
                _ = v37
                _ = v38
                _ = v39
            } else {
                fmt.Println(v34, v36)
                for fi40 := 0; fi40 < 8; fi40++ {
                    var v41 int16 = (-(v34))
                    _ = fi40
                    _ = v41
                }
                switch "\n\t" {
                case "":
                    v2 += v2
                case "\n\t":
                    v36 -= (^(v36))
                }
            }
        case 42:
            fmt.Println(v36, v34, v2)
        }
    } else {
        v42 := GenS5{F0: (v34 % ((v34) | 1)), F1: (v2 & (^(v2))), F2: "hello"}
        v34 -= v34
        _ = v42
    }
    switch 42 {
    case 933:
        v34 = v34
    case 113:
        if len(v4) > 0 {
            v4[0] = 127
        }
    case -2147483648:
        v3[int(v34)] = 273
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v34 = 32767
    var v43 string = "a"
    v44 := GenS7{F0: ("hello" + ("\n\t" + v43)), F1: (v35 != (v35 ^ (v35 * (v35 | (v35 ^ 1)))))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
