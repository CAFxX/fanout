package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint16 }
func genFn0(p0 string) uint {
    p0 = p0
    v7 := []int{32767, 2147483647, 377}
    p0 = p0
    _ = p0
    _ = v7
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 2
    v3 := []int{}
    var v4 int8 = -5
    var v5 int64 = 433
    var v6 bool = (false && (!true))
    v6 = true
    for fi8 := 0; fi8 < 5; fi8++ {
        _ = len(v3)
        _ = fi8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        v6 = v6
        _ = fi9
    }
    v6 = v6
    fmt.Println(v4)
    var v10 float32 = (1e10 - 1.5)
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 uint8 = 255
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: "hello", F1: v5}
    var v14 int32 = 42
    if false {
        v15 := GenS1{F0: ("" + "foo bar"), F1: v5}
        if (!false) {
            switch -1 {
            case 32768:
                for fi16 := 0; fi16 < 1; fi16++ {
                    for fi17 := 0; fi17 < 14; fi17++ {
                        fmt.Println(v4, v5)
                        v18 := GenS1{F0: "hello", F1: (v5 | (v5 >> ((4294967295) & 63)))}
                        _ = fi17
                        _ = v18
                    }
                    v19 := GenS1{F0: "hello", F1: int64(v14)}
                    switch (^(fi16)) {
                    case 1:
                        for fi20 := 0; fi20 < 1; fi20++ {
                            var v21 uint16 = 466
                            _ = fi20
                            _ = v21
                        }
                    case -369:
                        if v6 {
                            switch "hello" {
                            case "":
                                v2 = 32767
                            default:
                                v22 := GenS2{F0: 401}
                                _ = v22
                            }
                        } else {
                            switch -1 {
                            case 127:
                                var v23 int8 = -40
                                _ = v23
                            }
                            for fi24 := 0; fi24 < 3; fi24++ {
                                switch ("hello" + "") {
                                case "":
                                    var v25 uint16 = v2
                                    _ = v25
                                default:
                                    _ = len(v1)
                                }
                                _ = fi24
                            }
                            v26 := GenS2{F0: 255}
                            _ = v26
                        }
                    }
                    _ = fi16
                    _ = v19
                }
            default:
                var v27 uint16 = v2
                _ = v27
            }
        } else {
            if (!v6) {
                var v28 uint = 0
                if ((true || (v2 == 6)) && (21 <= v28)) {
                    fmt.Println(v6)
                    for fi29 := 0; fi29 < 2; fi29++ {
                        switch ((("hello" + ("foo bar" + "")) + "\n\t") + "hello") {
                        case "":
                            switch 9223372036854775807 {
                            case -498:
                                _ = len(v3)
                            case 304:
                                if ((v14 | 42) != v14) {
                                    var v30 uint8 = 1
                                    _ = v30
                                }
                            default:
                                var v31 bool = v6
                                _ = v31
                            }
                        }
                        _ = fi29
                    }
                    for fi32 := 0; fi32 < 17; fi32++ {
                        var v33 uint8 = 127
                        var v34 bool = v6
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                }
                var v35 int = int(v14)
                _ = v28
                _ = v35
            } else {
                v14 += -1
                var v36 int8 = v4
                _ = v36
            }
            var v37 bool = (0 > 777)
            v38 := GenS3{F0: 42, F1: v2}
            _ = v37
            _ = v38
        }
        var v39 uint16 = v2
        _ = v15
        _ = v39
    }
    v40 := GenS2{F0: genFn0(("a" + ("\n\t" + "")))}
    v1[("" + "foo bar")] = v6
    fmt.Println(v10)
    if (((!((v4 < (-(v4))) || (!((float32(0.1) - v10) == v10)))) || (-335 <= -444)) || v6) {
        v41 := v1["foo bar"]
        var v42 int32 = -2147483648
        var v43 int32 = 42
        _ = v41
        _ = v42
        _ = v43
    } else {
        for fi44 := 0; fi44 < 20; fi44++ {
            var v45 float32 = ((v10 + ((v10 * ((1e10 * (0.0 + (v10 - (-2.25 / 1.0)))) / (((v10 * (0.1 * (v10 / -2.25))) / v10) / 3.14159))) / v10)) + v10)
            v4 -= (^(v4))
            var v46 int32 = (v14 >> ((255) & 63))
            _ = fi44
            _ = v45
            _ = v46
        }
        v6 = (!true)
    }
    v47 := GenS2{F0: 194}
    v4 = (v4 + 0)
    switch "a" {
    case "\n\t":
        switch "hello" {
        case "hello":
            if ("a" == ("\n\t" + "foo bar")) {
                v5 = v5
                fmt.Println(v6, v6, v10)
            }
        default:
            if (v10 < 0.1) {
                var v48 uint16 = v2
                if (v6 || false) {
                    var v49 int32 = v14
                    v50 := v1["a"]
                    _ = v49
                    _ = v50
                }
                _ = v48
            } else {
                for fi51 := 0; fi51 < 5; fi51++ {
                    for fi52 := 0; fi52 < 20; fi52++ {
                        fmt.Println(v10, v10)
                        var v53 bool = ((v2 ^ 779) > (v2 * v2))
                        fmt.Println(v14)
                        _ = fi52
                        _ = v53
                    }
                    for fi54 := 0; fi54 < 3; fi54++ {
                        v55 := v1["a"]
                        var v56 uint8 = 25
                        v1["\n\t"] = (v55 || v55)
                        _ = fi54
                        _ = v55
                        _ = v56
                    }
                    _ = fi51
                }
                fmt.Println(v2)
            }
        }
    case "hello":
        if (true && ("foo bar" >= "a")) {
            if (!(!(v2 <= v2))) {
                v4 += (v4 ^ 42)
                for fi57 := 0; fi57 < 18; fi57++ {
                    v14 -= (v14 * (v14 - (-(v14))))
                    _ = fi57
                }
            }
            v5 = v5
            fmt.Println(v14)
        }
    default:
        v58 := GenS2{F0: 2}
        _ = v58
    }
    for fi59 := 0; fi59 < 7; fi59++ {
        fmt.Println(v5, v6)
        if v6 {
            var v60 int64 = 9223372036854775807
            if v6 {
                v61 := GenS1{F0: "a", F1: (^(v60))}
                v62 := GenS1{F0: "foo bar", F1: -551}
                _ = v61
                _ = v62
            } else {
                if len(v3) > 0 {
                    v3[0] = 895
                }
                v6 = (!v6)
            }
            switch fi59 {
            case 2147483647:
                fmt.Println(v5, v4)
            }
            _ = v60
        } else {
            var v63 int32 = v14
            _ = v63
        }
        switch "foo bar" {
        case "foo bar":
            v3 = append(v3, fi59)
        case "hello":
            fmt.Println(v14)
        case "a":
            var v64 uint8 = 42
            _ = v64
        default:
            fmt.Println(v6, v5)
        }
        _ = fi59
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v47
    _ = v5
    _ = v6
}
