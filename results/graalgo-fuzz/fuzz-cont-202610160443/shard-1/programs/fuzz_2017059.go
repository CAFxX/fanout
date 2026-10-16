package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 int }
func genFn0(p0 string, p1 int32, p2 uint) string {
    p2 = p2
    v4 := GenS1{F0: (42 <= 32767), F1: false, F2: p0}
    var v5 string = (p0 + (((p0 + p0) + p0) + "a"))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p0
}

func main() {
    var v1 int16 = -32768
    var v2 int8 = 42
    var v3 uint32 = 2
    if (3.14159 == -2.25) {
        if (781 < 42) {
            var v6 int32 = -1
            v2 = (v2 ^ 100)
            _ = v6
        } else {
            v1 = -286
        }
    } else {
        v3 += v3
        var v7 uint32 = 178
        var v8 uint32 = v3
        _ = v7
        _ = v8
    }
    var v9 string = (("" + "\n\t") + "foo bar")
    v10 := []int{}
    v11 := GenS1{F0: true, F1: (108 != 68), F2: v9}
    if len(v10) > 0 {
        v10[0] = -2147483648
    }
    v2 = v2
    if (881 != 612) {
        v12 := GenS1{F0: false, F1: true, F2: ""}
        _ = v12
    } else {
        v13 := GenS1{F0: (42 < 105), F1: true, F2: v9}
        _ = v13
    }
    var v14 uint = 0
    if (v14 <= 0) {
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := GenS1{F0: (989 >= (v14 | v14)), F1: false, F2: ("\n\t" + ("foo bar" + "foo bar"))}
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 2; fi17++ {
            if (4294967295 <= 1) {
                fmt.Println(v1)
                v18 := GenS1{F0: (42 < 1), F1: (float32(0.1) >= 1.5), F2: ((v9 + v9) + v9)}
                for fi19 := 0; fi19 < 19; fi19++ {
                    if (!true) {
                        fmt.Println(v2)
                    }
                    v20 := GenS1{F0: (true || true), F1: (((1.5 / 3.14159) - (0.0 + ((((1.5 / (((0.0 * -2.25) * 1.5) - 3.14159)) + (0.1 / 1.5)) * 1e10) / 1.0))) < (0.1 + 1e10)), F2: v9}
                    _ = fi19
                    _ = v20
                }
                _ = v18
            } else {
                if ((v3 * v3) >= 314) {
                    v21 := GenS1{F0: false, F1: false, F2: "hello"}
                    var v22 float32 = -2.25
                    if true {
                        if (true || true) {
                            v3 = 1
                        }
                        if true {
                            if false {
                                switch 32767 {
                                case 216:
                                    v10 = append(v10, fi17)
                                default:
                                    v14 -= (v14 ^ v14)
                                }
                                var v23 uint16 = 600
                                _ = v23
                            }
                        }
                    }
                    _ = v21
                    _ = v22
                } else {
                    v3 = (v3 >> (((v14 / ((304) | 1))) & 63))
                    var v24 string = v9
                    v1 -= -936
                    _ = v24
                }
                var v25 int16 = v1
                var v26 uint32 = (v3 + 725)
                _ = v25
                _ = v26
            }
            v27 := GenS2{F0: v9}
            _ = fi17
            _ = v27
        }
        if (42 > 2) {
            switch (v9 + "hello") {
            case "\n\t":
                v2 = (v2 << (((v14 / ((v14) | 1))) & 63))
            case "a":
                var v28 float32 = (1e10 + (float32((0.0 + 3.14159)) + 0.0))
                _ = v28
            case "foo bar":
                fmt.Println(v2)
            default:
                for fi29 := 0; fi29 < 2; fi29++ {
                    var v30 int64 = int64(fi29)
                    _ = fi29
                    _ = v30
                }
            }
        }
    }
    if (-2.25 != -2.25) {
        if len(v10) > 0 {
            v10[0] = 2147483647
        }
    }
    var v31 uint16 = 646
    if (!(-9223372036854775808 <= -183)) {
        for fi32 := 0; fi32 < 7; fi32++ {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            v33 := GenS3{F0: 128}
            _ = fi32
            _ = v33
        }
        switch 4294967295 {
        case -1:
            var v34 string = v9
            _ = v34
        case 484:
            var v35 int64 = -1
            _ = v35
        default:
            v9 = ((v9 + "") + v9)
        }
    } else {
        switch "foo bar" {
        case "foo bar":
            v14 = (v14 & v14)
        case "\n\t":
            v9 = (v9 + (v9 + v9))
        case "hello":
            v10 = append(v10, -9223372036854775808)
        default:
            for fi36 := 0; fi36 < 19; fi36++ {
                for fi37 := 0; fi37 < 17; fi37++ {
                    v31 = v31
                    _ = fi37
                }
                _ = fi36
            }
        }
        v9 = ("" + "")
        v9 = "\n\t"
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        var v39 float32 = 3.14159
        _ = fi38
        _ = v39
    }
    if (v9 != "a") {
        v2 = (v2 >> (((^(v14))) & 63))
        v40 := GenS3{F0: 42}
        _ = len(v10)
        _ = v40
    } else {
        for fi41 := 0; fi41 < 7; fi41++ {
            var v42 int64 = 794
            _ = fi41
            _ = v42
        }
    }
    v43 := GenS4{F0: -1, F1: 32768}
    var v44 float64 = 1e10
    v44 = ((float64(-2.25) + (v44 * v44)) / v44)
    var v45 int16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v43
    _ = v44
    _ = v45
    _ = v9
}
