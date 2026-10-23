package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int8 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 bool; F1 string; F2 string }
func genFn0(p0 uint32) uint8 {
    v6 := GenS1{F0: -1, F1: 226, F2: 127}
    for fi7 := 0; fi7 < 7; fi7++ {
        p0 -= (p0 & p0)
        for fi8 := 0; fi8 < 20; fi8++ {
            p0 -= p0
            _ = fi8
        }
        switch (fi7 ^ fi7) {
        case 32767:
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 uint64 = 18446744073709551615
                _ = fi9
                _ = v10
            }
        case 4294967295:
            p0 = (p0 / ((p0) | 1))
        }
        _ = fi7
    }
    if true {
        v11 := []int{0, 567, 127}
        _ = v11
    } else {
        p0 = p0
    }
    p0 -= p0
    var v12 uint = 4294967295
    _ = p0
    _ = v12
    _ = v6
    return 0
}

func main() {
    var v1 int8 = 42
    v2 := make(map[string]int)
    var v3 uint16 = 42
    var v4 int16 = 0
    var v5 int8 = (v1 - (v1 * 121))
    var v13 float64 = 0.0
    v13 += v13
    for fi14 := 0; fi14 < 1; fi14++ {
        if ((v4 - 0) == -32768) {
            v13 = v13
            v13 += (1e10 + (v13 / v13))
        } else {
            var v15 string = "hello"
            v16 := GenS1{F0: 42, F1: 366, F2: (v5 / ((v5) | 1))}
            if true {
                switch v15 {
                case "foo bar":
                    if ((v15 + "hello") >= ("a" + "foo bar")) {
                        var v17 float32 = -2.25
                        _ = v17
                    }
                case "":
                    v4 = v4
                case "hello":
                    delete(v2, v15)
                default:
                    for fi18 := 0; fi18 < 15; fi18++ {
                        var v19 uint32 = 2
                        v20 := GenS1{F0: 899, F1: 1, F2: (v1 % ((39) | 1))}
                        if (true && (-2.25 != -2.25)) {
                            if (2147483647 == -778) {
                                if (-1 != -9223372036854775808) {
                                    v1 = v1
                                    v21 := GenS1{F0: 0, F1: v3, F2: v5}
                                    v4 -= 1000
                                    _ = v21
                                }
                            }
                            v22 := []int{fi14, fi18, fi18, (fi14 % ((fi18) | 1))}
                            v1 = v5
                            _ = v22
                        } else {
                            var v23 int64 = 9223372036854775807
                            _ = v23
                        }
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                }
                for fi24 := 0; fi24 < 17; fi24++ {
                    v3 = v3
                    switch v15 {
                    case "foo bar":
                        v1 = (v5 ^ (v1 & v1))
                    case "hello":
                        var v25 uint16 = (v3 + 135)
                        _ = v25
                    case "":
                        if (325 < 2147483647) {
                            v26 := GenS1{F0: 0, F1: v3, F2: v5}
                            var v27 uint64 = 1
                            _ = v26
                            _ = v27
                        } else {
                            var v28 uint32 = 65535
                            if ((!false) && false) {
                                v29 := GenS1{F0: 889, F1: (^(v3)), F2: (v5 ^ 100)}
                                _ = v29
                            } else {
                                fmt.Println(v1)
                                v30 := []int{(fi24 << ((65535) & 63)), fi14}
                                v31 := GenS1{F0: 42, F1: (^(v3)), F2: v5}
                                _ = v30
                                _ = v31
                            }
                            v32 := GenS1{F0: 42, F1: (v3 ^ (v3 ^ 788)), F2: v1}
                            fmt.Println(fi14, fi24)
                            _ = v28
                            _ = v32
                        }
                    default:
                        v33 := GenS1{F0: 42, F1: (v3 % ((v3) | 1)), F2: (v5 % (((^(v5))) | 1))}
                        _ = v33
                    }
                    _ = fi24
                }
            } else {
                var v34 int64 = 78
                _ = v34
            }
            _ = v15
            _ = v16
        }
        v35 := GenS1{F0: 313, F1: v3, F2: (v5 | v1)}
        if ((v4 <= 32767) && (true || true)) {
            v36 := []int{fi14}
            v13 += v13
            _ = v36
        }
        _ = fi14
        _ = v35
    }
    v1 = (v1 << ((466) & 63))
    var v37 uint = 698
    v13 += 1.5
    v13 = float64(-2.25)
    for fi38 := 0; fi38 < 6; fi38++ {
        v39 := GenS1{F0: 867, F1: (v3 >> ((1) & 63)), F2: (v1 % ((0) | 1))}
        for fi40 := 0; fi40 < 8; fi40++ {
            delete(v2, ("a" + "a"))
            _ = fi40
        }
        _ = fi38
        _ = v39
    }
    v13 -= 3.14159
    var v41 int64 = 392
    if (v37 == 1) {
        _ = len(v2)
        if (559 < 1) {
            var v42 string = "hello"
            v43 := v2["hello"]
            _ = v42
            _ = v43
        }
        for fi44 := 0; fi44 < 17; fi44++ {
            for fi45 := 0; fi45 < 10; fi45++ {
                v3 = (v3 ^ 603)
                _ = len(v2)
                v46 := GenS2{F0: ("\n\t" + ""), F1: 0}
                _ = fi45
                _ = v46
            }
            var v47 float32 = 1.5
            _ = fi44
            _ = v47
        }
    } else {
        var v48 float64 = float64((-2.25 - 0.1))
        for fi49 := 0; fi49 < 5; fi49++ {
            switch fi49 {
            case 9223372036854775807:
                var v50 int8 = (v5 ^ v5)
                _ = v50
            case 2147483647:
                v48 = v48
            case 32768:
                switch ((("foo bar" + "") + "a") + "hello") {
                case "hello":
                    v51 := GenS2{F0: "a", F1: (fi49 * (fi49 | 42))}
                    _ = v51
                case "":
                    var v52 uint16 = 32767
                    _ = v52
                case "a":
                    var v53 int32 = int32(v3)
                    _ = v53
                default:
                    if (228 >= (v37 ^ 255)) {
                        v54 := GenS1{F0: 1, F1: v3, F2: v1}
                        v55 := GenS1{F0: 0, F1: 761, F2: (v1 / ((v1) | 1))}
                        _ = v54
                        _ = v55
                    }
                }
            default:
                var v56 int8 = v1
                _ = v56
            }
            _ = len(v2)
            _ = fi49
        }
        var v57 int = 32768
        _ = v48
        _ = v57
    }
    v58 := []int{}
    for fi59 := 0; fi59 < 18; fi59++ {
        switch "" {
        case "foo bar":
            var v60 uint32 = 0
            _ = v60
        case "":
            fmt.Println(v13, v5, v1)
        case "hello":
            if (true && (uint16(v37) > 1)) {
                v5 = v1
            } else {
                if ((false && false) || true) {
                    fmt.Println(v37, fi59, v1)
                    if len(v58) >= 2 {
                        v58 = v58[1:len(v58)]
                    }
                }
                v61 := GenS3{F0: (((((v13 * v13) / (v13 - v13)) <= v13) || (228 != 784)) && true), F1: "\n\t", F2: "a"}
                var v62 string = (("a" + "") + "\n\t")
                _ = v61
                _ = v62
            }
        }
        _ = fi59
    }
    if false {
        switch ("" + "a") {
        case "":
            fmt.Println(v1, v4, v3)
        case "\n\t":
            v58 = append(v58, 127)
        case "foo bar":
            for fi63 := 0; fi63 < 13; fi63++ {
                v41 = v41
                _ = fi63
            }
        }
        var v64 string = "a"
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v58:", len(v58), v58)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v5
    _ = v58
}
