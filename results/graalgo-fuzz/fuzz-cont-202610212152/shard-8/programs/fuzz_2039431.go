package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint; F1 string; F2 int32 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 uint64 }
func genFn0() string {
    switch "a" {
    case "hello":
        var v7 string = ""
        _ = v7
    case "a":
        var v8 uint32 = 346
        _ = v8
    default:
        var v9 uint32 = 65535
        _ = v9
    }
    v10 := []int{127, 32767}
    _ = v10
    return ((("foo bar" + "") + "") + "")
}
func genFn1(p0 int, p1 uint64, p2 int16) int8 {
    p0 -= p0
    v11 := GenS1{F0: 885, F1: p2, F2: 1}
    p1 -= 2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 120
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{}
    var v3 float64 = 0.1
    v4 := make(map[string]uint32)
    var v5 uint64 = 4294967295
    var v6 int = 128
    if false {
        v6 -= v6
        if (!(0 == 16)) {
            var v12 float64 = v3
            _ = v12
        }
        var v13 int32 = 719
        _ = v13
    } else {
        v14 := v1[(-(v6))]
        switch "" {
        case "hello":
            fmt.Println(v14, v5, v6)
        case "\n\t":
            fmt.Println(v5, v5, v3)
        default:
            fmt.Println(v14, v6)
        }
        v3 += v3
        _ = v14
    }
    var v15 int32 = 42
    v3 -= (1e10 + float64(((((-2.25 + 1.5) / 1.5) * 0.0) + (1.5 / 1e10))))
    if (false && true) {
        v15 = v15
        var v16 int8 = 18
        _ = v16
    } else {
        if ((v15 + v15) >= 42) {
            for fi17 := 0; fi17 < 11; fi17++ {
                v6 += 32768
                for fi18 := 0; fi18 < 5; fi18++ {
                    delete(v1, (fi17 - (fi17 - fi17)))
                    v6 = -1
                    fmt.Println(v15)
                    _ = fi18
                }
                var v19 int32 = (v15 | v15)
                _ = fi17
                _ = v19
            }
        } else {
            for fi20 := 0; fi20 < 19; fi20++ {
                v3 += v3
                v21 := GenS1{F0: 816, F1: 42, F2: 4294967295}
                if true {
                    var v22 int = v6
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v6 = v6
            } else {
                for fi23 := 0; fi23 < 15; fi23++ {
                    v15 = (v15 & -35)
                    _ = fi23
                }
            }
        }
    }
    var v24 float64 = ((((v3 - v3) / v3) + (v3 * v3)) - v3)
    v3 = 0.1
    for fi25 := 0; fi25 < 11; fi25++ {
        v3 = v24
        _ = fi25
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v15 += 2147483647
    v15 = v15
    var v26 string = "\n\t"
    if (!(!false)) {
        var v27 string = ((v26 + v26) + (v26 + ""))
        var v28 float32 = 0.1
        _ = v27
        _ = v28
    } else {
        var v29 int8 = -5
        switch (v26 + v26) {
        case "":
            switch int(v15) {
            case 127:
                v15 += 2147483647
            case 422:
                if true {
                    v30 := GenS2{F0: 4294967295, F1: "", F2: 1}
                    v2 = append(v2, "")
                    switch (v6 << ((uint(v5)) & 63)) {
                    case 1:
                        if (3.14159 < 1.5) {
                            v15 = v15
                            v29 = 100
                            for fi31 := 0; fi31 < 7; fi31++ {
                                var v32 uint = 18446744073709551615
                                v3 += -2.25
                                _ = fi31
                                _ = v32
                            }
                        } else {
                            var v33 int16 = -760
                            var v34 float32 = (0.0 + (float32(3.14159) / 0.1))
                            for fi35 := 0; fi35 < 10; fi35++ {
                                var v36 uint = 355
                                _ = fi35
                                _ = v36
                            }
                            _ = v33
                            _ = v34
                        }
                    case 2147483647:
                        v37 := GenS3{F0: 1, F1: v5}
                        _ = v37
                    }
                    _ = v30
                }
            case -370:
                var v38 uint16 = 290
                _ = v38
            default:
                _ = len(v4)
            }
        case "hello":
            v2 = append(v2, v26)
        case "foo bar":
            for fi39 := 0; fi39 < 13; fi39++ {
                var v40 int8 = 127
                switch v26 {
                case "hello":
                    if false {
                        v1[fi39] = (^(v6))
                    } else {
                        switch (v26 + ("" + "\n\t")) {
                        case "a":
                            v3 = v3
                        case "foo bar":
                            if (32767 <= 2) {
                                switch v26 {
                                case "hello":
                                    v6 = 1
                                case "foo bar":
                                    if (!false) {
                                        var v41 int8 = (-(v40))
                                        if true {
                                            var v42 int32 = 42
                                            _ = v42
                                        }
                                        v3 = v3
                                        _ = v41
                                    }
                                case "":
                                    v43 := GenS3{F0: 255, F1: v5}
                                    _ = v43
                                }
                                var v44 int8 = (v40 >> ((42) & 63))
                                _ = len(v4)
                                _ = v44
                            }
                        case "hello":
                            for fi45 := 0; fi45 < 19; fi45++ {
                                v46 := GenS2{F0: 65535, F1: "a", F2: -560}
                                v2 = append(v2, (v26 + ""))
                                if ((-884 < 2147483647) && true) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = (v26 + ("a" + v26))
                                    }
                                    if true {
                                        v3 += 0.0
                                        fmt.Println(v26, v26)
                                        var v47 int8 = 42
                                        _ = v47
                                    } else {
                                        v29 += genFn1((v6 & fi39), v5, -1)
                                    }
                                } else {
                                    v48 := GenS4{F0: v5}
                                    v49 := GenS3{F0: 2, F1: 674}
                                    _ = v48
                                    _ = v49
                                }
                                _ = fi45
                                _ = v46
                            }
                        }
                        fmt.Println(v29, v5)
                    }
                case "":
                    fmt.Println(v40)
                case "a":
                    v2 = append(v2, (v26 + ""))
                }
                v6 = v6
                _ = fi39
                _ = v40
            }
        }
        if ((1 > 647) || true) {
            fmt.Println(v26)
        }
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
