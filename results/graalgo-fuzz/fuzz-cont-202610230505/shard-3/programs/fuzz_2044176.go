package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint8 }
func genFn0(p0 int32, p1 uint16, p2 int64) uint16 {
    for fi6 := 0; fi6 < 6; fi6++ {
        p2 = (p2 | 118)
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: (^(p2)), F1: 42, F2: (p1 ^ 32767)}
            var v9 uint8 = 1
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    var v10 string = ("a" + "foo bar")
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p1
}

func main() {
    var v1 int = 1
    v2 := make(map[uint32]bool)
    v3 := []string{"", "foo bar", (("" + "hello") + "hello"), ((("foo bar" + "a") + "hello") + (("" + "a") + "hello"))}
    var v4 int16 = 42
    v5 := []uint32{272}
    v11 := GenS1{F0: 9223372036854775807, F1: -106, F2: 42}
    v12 := GenS1{F0: -613, F1: 100, F2: 952}
    var v13 int = (-(v1))
    v14 := GenS1{F0: -776, F1: 83, F2: 436}
    _ = len(v3)
    var v15 uint32 = 43
    switch "" {
    case "a":
        fmt.Println(v13, v15, v4)
    default:
        if false {
            v16 := GenS1{F0: 9223372036854775807, F1: 1, F2: 1}
            var v17 uint64 = 18446744073709551615
            var v18 int64 = 0
            _ = v16
            _ = v17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        v1 -= (v1 / (((^(v1))) | 1))
        if (2147483647 >= 42) {
            var v20 float64 = float64(-2.25)
            _ = v20
        }
        if len(v3) > 0 {
            v3[0] = "a"
        }
        _ = fi19
    }
    v4 = v4
    for fi21 := 0; fi21 < 20; fi21++ {
        switch "\n\t" {
        case "foo bar":
            switch v13 {
            case -9223372036854775808:
                for fi22 := 0; fi22 < 18; fi22++ {
                    var v23 float64 = 0.1
                    _ = fi22
                    _ = v23
                }
            case -706:
                var v24 float32 = 0.0
                _ = v24
            default:
                for fi25 := 0; fi25 < 20; fi25++ {
                    v15 -= v15
                    _ = fi25
                }
            }
        case "hello":
            var v26 int = 4294967295
            _ = v26
        case "":
            if true {
                var v27 int32 = -692
                fmt.Println(v15, v15, v15)
                v28 := GenS1{F0: -605, F1: -91, F2: 42}
                _ = v27
                _ = v28
            }
        }
        var v29 int64 = 50
        if (42 <= -1) {
            for fi30 := 0; fi30 < 15; fi30++ {
                _ = len(v3)
                if false {
                    v31 := GenS1{F0: 0, F1: -1, F2: 32767}
                    switch (fi21 % ((fi30) | 1)) {
                    case 32768:
                        switch "foo bar" {
                        case "":
                            switch "\n\t" {
                            case "":
                                var v32 int64 = (v29 | 1)
                                fmt.Println(fi30)
                                v33 := v2[(^(v15))]
                                _ = v32
                                _ = v33
                            default:
                                if false {
                                    v3 = append(v3, ("hello" + "foo bar"))
                                }
                            }
                        case "\n\t":
                            v34 := GenS1{F0: (v29 | (v29 % (((v29 - (v29 ^ v29))) | 1))), F1: 100, F2: 32767}
                            _ = v34
                        default:
                            if true {
                                if false {
                                    v15 = 42
                                }
                                v15 += v15
                                var v35 string = ""
                                _ = v35
                            } else {
                                switch (v1 + (-(v13))) {
                                case 1:
                                    fmt.Println(fi21, fi30)
                                }
                                var v36 bool = true
                                _ = v36
                            }
                        }
                    case 1:
                        v29 = -9223372036854775808
                    case 42:
                        var v37 int32 = 972
                        _ = v37
                    }
                    var v38 bool = (1 < (v29 + v29))
                    v29 += v29
                    _ = v31
                    _ = v38
                }
                v4 = (^(v4))
                _ = fi30
            }
            fmt.Println(v13)
            fmt.Println(v13, v13)
        }
        _ = fi21
        _ = v29
    }
    if false {
        v1 += (v1 * -503)
        if (!false) {
            v13 = (v13 ^ v13)
            var v39 float32 = 1e10
            _ = v39
        }
        var v40 uint8 = 39
        _ = v40
    }
    v13 = (v13 * (v1 * 554))
    if (((("foo bar" + "\n\t") + "foo bar") + "\n\t") > ("foo bar" + "")) {
        v41 := GenS2{F0: false, F1: 127}
        _ = v41
    }
    var v42 string = "\n\t"
    var v43 uint64 = 18446744073709551615
    v13 = -131
    for fi44 := 0; fi44 < 13; fi44++ {
        v43 += 173
        _ = fi44
    }
    v15 += 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
