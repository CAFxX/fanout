package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int; F1 string; F2 uint64 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 string) uint32 {
    fmt.Println(p0)
    v6 := GenS1{F0: -590, F1: true}
    if true {
        if ((42 < 42) && (135 != 42)) {
            p0 = (("\n\t" + ("\n\t" + "foo bar")) + "")
            switch ((p0 + (p0 + (p0 + ""))) + p0) {
            case "hello":
                for fi7 := 0; fi7 < 15; fi7++ {
                    p0 = ("" + "")
                    v8 := GenS2{F0: (fi7 >> ((42) & 63)), F1: "a", F2: 42}
                    _ = fi7
                    _ = v8
                }
            }
        }
        if (-903 < 1) {
            p0 = p0
        }
        var v9 uint64 = 182
        _ = v9
    } else {
        if (!true) {
            for fi10 := 0; fi10 < 17; fi10++ {
                var v11 int8 = 17
                var v12 bool = ((fi10 | 994) < (fi10 << ((1) & 63)))
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        v13 := []int{-1, 1, 127}
        p0 = ("foo bar" + p0)
        _ = v13
    }
    var v14 uint32 = 275
    _ = p0
    _ = v14
    _ = v6
    return v14
}
func genFn1(p0 uint64, p1 int8, p2 int8) uint32 {
    switch ((("a" + ("hello" + "a")) + "hello") + "a") {
    case "hello":
        fmt.Println(p1, p1, p2)
    case "\n\t":
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := []int{(fi15 ^ fi15), (fi15 * fi15)}
            if len(v16) > 0 {
                v16[0] = 0
            }
            var v17 uint8 = 53
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        p2 -= (p1 % ((0) | 1))
        var v19 uint8 = 13
        if (-1 != 0) {
            fmt.Println(fi18)
        }
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    v1 := []string{"a", "a"}
    var v2 float64 = ((0.0 + 0.0) * 3.14159)
    v3 := make(map[string]int)
    var v4 uint32 = 744
    v5 := []string{"", ("" + "a")}
    if false {
        v20 := GenS1{F0: 1, F1: true}
        _ = v20
    }
    v2 -= -2.25
    if (2147483647 <= 2147483648) {
        v2 = (v2 + (1.5 + v2))
        v2 -= float64(0.0)
    }
    v4 += (v4 / (((^(v4))) | 1))
    v21 := GenS3{F0: 0, F1: (!((float32((3.14159 * 0.0)) - (1e10 + 1.5)) > float32((v2 + (0.1 - v2)))))}
    switch 0 {
    case -1:
        if false {
            for fi22 := 0; fi22 < 17; fi22++ {
                for fi23 := 0; fi23 < 6; fi23++ {
                    v24 := GenS2{F0: -9223372036854775808, F1: (("foo bar" + (((("foo bar" + "foo bar") + "\n\t") + "a") + "a")) + "foo bar"), F2: 594}
                    if false {
                        var v25 uint16 = 32767
                        if (true || (1.5 < (1.5 / 3.14159))) {
                            v25 = v25
                            v1 = append(v1, ("" + "hello"))
                        }
                        _ = v25
                    } else {
                        switch "a" {
                        case "foo bar":
                            if (-117 > -7) {
                                var v26 uint32 = (v4 & v4)
                                var v27 string = ("\n\t" + ("a" + ""))
                                if (!false) {
                                    var v28 int64 = 9223372036854775807
                                    var v29 uint = 506
                                    _ = v28
                                    _ = v29
                                } else {
                                    switch (fi23 + fi23) {
                                    case -9223372036854775808:
                                        fmt.Println(v2, v27, fi23)
                                    case -341:
                                        var v30 uint16 = 32767
                                        _ = v30
                                    case 42:
                                        _ = len(v3)
                                        var v31 float64 = (v2 * v2)
                                        v3[v27] = 0
                                        _ = v31
                                    default:
                                        v2 = v2
                                    }
                                }
                                _ = v26
                                _ = v27
                            } else {
                                v2 -= (0.1 / 1e10)
                                switch 2147483647 {
                                case -1:
                                    fmt.Println(fi23, v4)
                                case 127:
                                    if (!(true && (215 < 1))) {
                                        v4 -= genFn0((("" + "\n\t") + ("hello" + (("a" + "a") + "foo bar"))))
                                    } else {
                                        var v32 int64 = 9223372036854775807
                                        v3["hello"] = (fi23 ^ (fi23 % (((fi23 * (fi22 & fi22))) | 1)))
                                        _ = len(v1)
                                        _ = v32
                                    }
                                case 128:
                                    v2 = v2
                                default:
                                    if false {
                                        v33 := v3[("" + "\n\t")]
                                        var v34 int16 = 32767
                                        var v35 int64 = 1
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                }
                                if false {
                                    v4 = 2147483647
                                } else {
                                    v36 := GenS3{F0: 0, F1: (-669 != -321)}
                                    _ = v36
                                }
                            }
                        case "hello":
                            v37 := GenS4{F0: 0, F1: uint64(v4)}
                            _ = v37
                        case "":
                            v5 = append(v5, "\n\t")
                        }
                        delete(v3, "")
                    }
                    _ = fi23
                    _ = v24
                }
                _ = fi22
            }
            if (2 <= 0) {
                var v38 float32 = ((1.5 * ((0.1 + 1.5) - 0.1)) * (((1e10 + (-2.25 / 3.14159)) - 3.14159) / float32(((1.5 + 0.0) * -2.25))))
                _ = v38
            }
            for fi39 := 0; fi39 < 11; fi39++ {
                v40 := GenS5{F0: 42, F1: 563}
                _ = fi39
                _ = v40
            }
        } else {
            for fi41 := 0; fi41 < 8; fi41++ {
                v2 += (1e10 - (v2 * v2))
                switch -583 {
                case 127:
                    v42 := GenS6{F0: 591}
                    _ = v42
                default:
                    v43 := v3[(("\n\t" + "foo bar") + "a")]
                    _ = v43
                }
                if ((!(1 < 0)) || (true && (65535 >= 1))) {
                    var v44 uint32 = v4
                    var v45 string = ((("foo bar" + "hello") + (("\n\t" + ("\n\t" + "foo bar")) + "foo bar")) + ("" + ("\n\t" + "foo bar")))
                    _ = v44
                    _ = v45
                } else {
                    var v46 string = ("\n\t" + "\n\t")
                    var v47 string = ("foo bar" + v46)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v46
                    _ = v47
                }
                _ = fi41
            }
            var v48 float32 = -2.25
            _ = v48
        }
    default:
        v3["a"] = 42
    }
    v4 += (v4 % ((828) | 1))
    var v49 int8 = 42
    v4 -= v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v49
    _ = v5
}
