package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 int64; F2 uint32 }
func genFn0(p0 uint, p1 bool, p2 bool) uint8 {
    var v4 uint32 = 4294967295
    if p1 {
        if p2 {
            var v5 int64 = -9223372036854775808
            _ = v5
        } else {
            for fi6 := 0; fi6 < 9; fi6++ {
                var v7 float32 = 1e10
                _ = fi6
                _ = v7
            }
            if (127 >= 121) {
                var v8 uint16 = uint16(p0)
                _ = v8
            } else {
                fmt.Println(v4, p0, p2)
            }
            for fi9 := 0; fi9 < 16; fi9++ {
                v10 := GenS1{F0: 70, F1: "hello", F2: (p0 / ((p0) | 1))}
                if (633 >= fi9) {
                    v11 := GenS1{F0: 0, F1: ("\n\t" + ("a" + "hello")), F2: p0}
                    v12 := []int{(fi9 % ((-1) | 1))}
                    _ = v11
                    _ = v12
                } else {
                    var v13 float64 = 0.1
                    switch "hello" {
                    case "":
                        var v14 int16 = -32768
                        _ = v14
                    case "a":
                        fmt.Println(p1, v4)
                    case "foo bar":
                        switch fi9 {
                        case 691:
                            switch "hello" {
                            case "foo bar":
                                p2 = (0 >= 2)
                            case "":
                                v15 := GenS2{F0: -25}
                                _ = v15
                            }
                        case 206:
                            switch (-(fi9)) {
                            case -606:
                                p2 = (1 <= 0)
                            case 9223372036854775807:
                                for fi16 := 0; fi16 < 4; fi16++ {
                                    v17 := []int{(fi9 & (fi16 ^ fi16)), (fi9 + 32767), (^(fi16))}
                                    _ = fi16
                                    _ = v17
                                }
                            case -521:
                                v18 := GenS2{F0: -1}
                                _ = v18
                            }
                        case 128:
                            for fi19 := 0; fi19 < 17; fi19++ {
                                p2 = (-744 <= 0)
                                _ = fi19
                            }
                        }
                    }
                    _ = v13
                }
                _ = fi9
                _ = v10
            }
        }
        if (-2.25 < 1e10) {
            var v20 float64 = 0.1
            _ = v20
        }
        if (-434 > 42) {
            if ((v4 ^ 572) == 2147483647) {
                if (((3.14159 != 1e10) || (-2.25 <= (1e10 + 0.1))) || p1) {
                    var v21 uint64 = 2
                    _ = v21
                } else {
                    v4 -= (v4 >> (((^(p0))) & 63))
                    if ((p0 << ((1) & 63)) < (p0 * 1)) {
                        for fi22 := 0; fi22 < 19; fi22++ {
                            fmt.Println(p0)
                            var v23 float64 = (float64(3.14159) - (0.1 / ((float64(-2.25) * 3.14159) * 3.14159)))
                            fmt.Println(fi22, p0)
                            _ = fi22
                            _ = v23
                        }
                        var v24 uint = 617
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 2; fi25++ {
                            fmt.Println(p0)
                            _ = fi25
                        }
                        var v26 string = ""
                        var v27 uint8 = 255
                        _ = v26
                        _ = v27
                    }
                    if p1 {
                        v28 := []int{-2147483648}
                        _ = v28
                    }
                }
                switch "" {
                case "hello":
                    if (!p1) {
                        fmt.Println(p1, p2)
                        p0 = (p0 / ((439) | 1))
                    } else {
                        var v29 int16 = -32768
                        v30 := []int{127}
                        _ = v29
                        _ = v30
                    }
                case "\n\t":
                    p2 = p1
                case "a":
                    fmt.Println(p1)
                default:
                    v31 := []int{-1, -1, -451, -1}
                    _ = v31
                }
            } else {
                var v32 uint = p0
                var v33 uint64 = 2
                if (v4 > v4) {
                    v34 := GenS3{F0: -1}
                    _ = v34
                }
                _ = v32
                _ = v33
            }
            var v35 float32 = 1.5
            fmt.Println(v4, p1, p2)
            _ = v35
        } else {
            v36 := GenS1{F0: 97, F1: "a", F2: (p0 % ((p0) | 1))}
            _ = v36
        }
    } else {
        v4 = 894
        var v37 int = -517
        _ = v37
    }
    v4 = v4
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = -2147483648
    var v3 int64 = 0
    var v38 int8 = -47
    for fi39 := 0; fi39 < 20; fi39++ {
        if true {
            _ = len(v1)
        } else {
            for fi40 := 0; fi40 < 2; fi40++ {
                if true {
                    var v41 int = fi39
                    v41 = -2147483648
                    v41 = int(v3)
                    _ = len(v1)
                    _ = v41
                } else {
                    v42 := GenS3{F0: int16(fi40)}
                    switch "" {
                    case "hello":
                        var v43 int64 = v3
                        _ = v43
                    case "a":
                        var v44 int = fi39
                        _ = v44
                    default:
                        var v45 int64 = (-(v3))
                        _ = v45
                    }
                    var v46 int8 = v38
                    _ = v42
                    _ = v46
                }
                v2 = v2
                switch fi40 {
                case 32768:
                    var v47 uint16 = 57
                    _ = v47
                case 2147483647:
                    if (float32(float64(1e10)) < 3.14159) {
                        v48 := GenS1{F0: 1, F1: "", F2: 4294967295}
                        v2 = -2147483648
                        fmt.Println(fi39, v2, fi39)
                        _ = v48
                    }
                case 1:
                    v2 = 0
                }
                _ = fi40
            }
        }
        if (true || (-2.25 <= (0.0 / (float32(1.5) / ((3.14159 * 0.0) - -2.25))))) {
            var v49 float32 = 1.5
            _ = v49
        }
        _ = fi39
    }
    v50 := GenS2{F0: (v2 << ((0) & 63))}
    v3 = (v3 & (v3 ^ -9223372036854775808))
    fmt.Println(v38, v2, v3)
    var v51 uint = 4294967295
    v38 -= v38
    for fi52 := 0; fi52 < 2; fi52++ {
        v53 := GenS4{F0: "\n\t", F1: (v3 << ((v51) & 63)), F2: 608}
        _ = fi52
        _ = v53
    }
    v38 = (v38 % ((v38) | 1))
    fmt.Println(v2, v51)
    fmt.Println(v3, v3, v51)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v50
    _ = v51
}
