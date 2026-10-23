package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 uint; F2 string }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 uint; F1 uint8 }

func main() {
    v1 := []int{821, 804}
    v2 := []string{"\n\t", "foo bar", ("foo bar" + "a"), (("\n\t" + "foo bar") + "a")}
    var v3 uint = 2
    if false {
        var v4 float32 = 0.0
        v5 := GenS1{F0: 0}
        for fi6 := 0; fi6 < 1; fi6++ {
            for fi7 := 0; fi7 < 1; fi7++ {
                if (!(!((^(v3)) == (v3 / (((^(v3))) | 1))))) {
                    var v8 float64 = ((3.14159 + 0.0) / -2.25)
                    _ = v8
                } else {
                    v9 := GenS1{F0: 459}
                    fmt.Println(fi7, fi6, v4)
                    switch ("" + "hello") {
                    case "\n\t":
                        if true {
                            var v10 int8 = 13
                            switch ("hello" + "foo bar") {
                            case "foo bar":
                                v11 := GenS1{F0: 0}
                                _ = v11
                            }
                            _ = v10
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v12 int = (fi7 >> (((v3 * v3)) & 63))
                        _ = v12
                    default:
                        v4 += (1.5 + (((float32(0.0) - (v4 - 3.14159)) * v4) * (float32(((1e10 / 1.0) / (1e10 + 0.0))) / (0.0 * ((3.14159 / 0.1) + v4)))))
                    }
                    _ = v9
                }
                fmt.Println(fi7, fi7, v4)
                _ = fi7
            }
            var v13 uint64 = 219
            _ = fi6
            _ = v13
        }
        _ = v4
        _ = v5
    } else {
        _ = len(v1)
    }
    var v14 int64 = -9223372036854775808
    var v15 float32 = 3.14159
    for fi16 := 0; fi16 < 11; fi16++ {
        v2 = append(v2, ("\n\t" + ("" + "\n\t")))
        v15 = float32(float64(v15))
        _ = fi16
    }
    if ((-995 > 688) && (42 == 1)) {
        v14 = -9223372036854775808
        if false {
            switch 127 {
            case -834:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                v17 := GenS2{F0: 42}
                _ = v17
            }
            var v18 int64 = (v14 | v14)
            _ = v18
        } else {
            v15 = v15
            v15 -= (v15 + -2.25)
            for fi19 := 0; fi19 < 19; fi19++ {
                for fi20 := 0; fi20 < 9; fi20++ {
                    var v21 uint32 = 787
                    _ = fi20
                    _ = v21
                }
                v15 += -2.25
                _ = fi19
            }
        }
        var v22 uint64 = 487
        _ = v22
    } else {
        v3 += 1
    }
    v23 := GenS1{F0: 2}
    if (float32(1.5) > ((1e10 / float32((float64(v15) * 1.5))) * 3.14159)) {
        if (false || false) {
            v3 = (v3 >> (((v3 % ((v3) | 1))) & 63))
            var v24 uint64 = 18446744073709551615
            _ = v24
        }
        var v25 bool = false
        for fi26 := 0; fi26 < 9; fi26++ {
            for fi27 := 0; fi27 < 8; fi27++ {
                v28 := GenS3{F0: true, F1: v3, F2: "\n\t"}
                _ = fi27
                _ = v28
            }
            var v29 string = "a"
            _ = len(v2)
            _ = fi26
            _ = v29
        }
        _ = v25
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        if true {
            v31 := GenS3{F0: true, F1: (v3 * 715), F2: "\n\t"}
            _ = v31
        } else {
            v15 -= -2.25
            v32 := GenS3{F0: false, F1: 255, F2: "a"}
            if (("foo bar" + "foo bar") == "foo bar") {
                switch (fi30 * fi30) {
                case -9223372036854775808:
                    if true {
                        for fi33 := 0; fi33 < 17; fi33++ {
                            _ = len(v1)
                            _ = fi33
                        }
                    } else {
                        v14 = (v14 * (v14 * (v14 / ((v14) | 1))))
                        v34 := GenS3{F0: (800 > (v14 >> ((v3) & 63))), F1: 980, F2: ""}
                        _ = v34
                    }
                }
            }
            _ = v32
        }
        for fi35 := 0; fi35 < 1; fi35++ {
            fmt.Println(v3, fi30, v14)
            var v36 int64 = (^(v14))
            _ = len(v2)
            var v37 string = "a"
            _ = fi35
            _ = v36
            _ = v37
        }
        if len(v1) > 0 {
            v1[0] = fi30
        }
        _ = fi30
    }
    switch 0 {
    case 387:
        v1 = append(v1, -2147483648)
    case 4294967295:
        fmt.Println(v3)
    }
    var v38 int64 = v14
    v39 := GenS4{F0: "foo bar", F1: -241}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        if (-9223372036854775808 == 82) {
            var v40 int = 128
            v40 = (v40 << ((v3) & 63))
            v41 := GenS3{F0: false, F1: (v3 & v3), F2: ("\n\t" + ("" + ""))}
            _ = v40
            _ = v41
        }
    } else {
        if true {
            var v42 uint8 = 42
            _ = v42
        } else {
            v15 -= ((0.0 - 0.1) / (float32((0.1 * ((0.1 * (0.1 / (1.5 - -2.25))) + (-2.25 / 3.14159)))) / float32(1e10)))
        }
        v43 := GenS4{F0: "a", F1: -722}
        _ = v43
    }
    if false {
        switch 1 {
        case 0:
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
        }
    } else {
        if true {
            v38 += 42
            for fi44 := 0; fi44 < 4; fi44++ {
                var v45 string = "\n\t"
                switch v45 {
                case "foo bar":
                    v45 = v45
                default:
                    v46 := GenS5{F0: 2, F1: 255}
                    _ = v46
                }
                var v47 int8 = 75
                _ = fi44
                _ = v45
                _ = v47
            }
            if len(v2) > 0 {
                v2[0] = ((("\n\t" + "a") + "a") + "foo bar")
            }
        }
    }
    v15 = (float32((0.0 / (float64(v15) / float64(v15)))) / float32((1.5 + 0.0)))
    for fi48 := 0; fi48 < 15; fi48++ {
        v49 := GenS5{F0: 18446744073709551615, F1: 1}
        _ = fi48
        _ = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v38
    _ = v39
}
