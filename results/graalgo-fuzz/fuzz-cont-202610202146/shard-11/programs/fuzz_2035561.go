package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 bool; F1 int64 }
func genFn0(p0 int, p1 int16, p2 uint8) int {
    p0 += p0
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 uint32 = 0
        for fi8 := 0; fi8 < 5; fi8++ {
            p0 += (p0 >> ((977) & 63))
            _ = fi8
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            fmt.Println(fi9, p0)
            var v10 int = fi6
            p2 += (p2 * p2)
            _ = fi9
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        p2 = (p2 << ((1) & 63))
        var v12 uint64 = 18446744073709551615
        _ = fi11
        _ = v12
    }
    v13 := []int{42, (-(p0))}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return p0
}
func genFn1() uint {
    var v14 int16 = -5
    v15 := []int{-352}
    v14 = (-(v14))
    v16 := GenS1{F0: "\n\t", F1: (((("" + ("a" + "")) + ((((("" + "") + (("" + "") + ("hello" + "foo bar"))) + "foo bar") + "hello") + ("" + ("" + "")))) + "foo bar") >= "a"), F2: -1}
    _ = v14
    _ = v15
    _ = v16
    return 65535
}

func main() {
    var v1 int16 = 1000
    var v2 uint64 = 0
    v3 := []string{("hello" + "a"), ""}
    var v4 int16 = (-(v1))
    v5 := make(map[string]bool)
    if false {
        var v17 float32 = 0.1
        if false {
            for fi18 := 0; fi18 < 3; fi18++ {
                var v19 bool = (genFn0((fi18 << ((255) & 63)), v4, 2) < genFn0(0, (v4 | (v4 << ((255) & 63))), 137))
                var v20 bool = ((!v19) || (2 != 65535))
                switch ("a" + "foo bar") {
                case "":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "foo bar":
                    if true {
                        v21 := v5["foo bar"]
                        var v22 float32 = ((((v17 + (float32(float64((v17 + 1.5))) - (0.0 - (float32(1.5) * 1.5)))) * float32((1e10 / 3.14159))) + (0.1 / v17)) * (1.5 * v17))
                        _ = v21
                        _ = v22
                    } else {
                        if ((0.0 - 1e10) < 0.0) {
                            var v23 float64 = (1.5 - (-2.25 + float64(((v17 + 1e10) + (0.0 / (v17 - (-2.25 + float32(((1.5 + -2.25) / 1.5)))))))))
                            _ = v23
                        } else {
                            if true {
                                fmt.Println(fi18, v19, v20)
                                for fi24 := 0; fi24 < 14; fi24++ {
                                    var v25 uint16 = 1
                                    v26 := GenS2{F0: 65535}
                                    v17 = ((1.5 * ((float32(1.5) - (0.0 / v17)) * ((1e10 / float32(3.14159)) * 3.14159))) * -2.25)
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                            } else {
                                fmt.Println(v20, v19, v17)
                            }
                            var v27 uint = 2
                            _ = v27
                        }
                        fmt.Println(v4)
                    }
                case "hello":
                    v20 = v20
                default:
                    v28 := GenS2{F0: 785}
                    _ = v28
                }
                _ = fi18
                _ = v19
                _ = v20
            }
            var v29 uint = 65535
            for fi30 := 0; fi30 < 16; fi30++ {
                v17 += v17
                _ = fi30
            }
            _ = v29
        } else {
            if (912 == 32767) {
                v31 := GenS2{F0: 4294967295}
                v4 = (v4 + v4)
                _ = v31
            }
            var v32 int32 = 621
            _ = v32
        }
        switch ("" + ("" + "")) {
        case "hello":
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
        case "foo bar":
            v33 := GenS2{F0: 950}
            _ = v33
        case "\n\t":
            if false {
                v34 := GenS1{F0: "", F1: ((!true) || true), F2: 0}
                if true {
                    var v35 int64 = 1
                    switch 0 {
                    case 9223372036854775807:
                        if false {
                            if (9223372036854775807 <= 2147483647) {
                                fmt.Println(v4, v2, v2)
                                var v36 uint16 = 2
                                _ = v36
                            }
                            var v37 uint32 = 602
                            var v38 int64 = -1
                            _ = v37
                            _ = v38
                        } else {
                            switch 42 {
                            case -2147483648:
                                v39 := GenS3{F0: 1, F1: v35}
                                _ = v39
                            case -279:
                                v1 = v4
                            case 0:
                                v4 -= v1
                            default:
                                var v40 uint32 = 42
                                _ = v40
                            }
                            switch ("a" + "hello") {
                            case "foo bar":
                                switch "hello" {
                                case "\n\t":
                                    fmt.Println(v17, v17, v4)
                                }
                            case "":
                                v41 := GenS4{F0: false, F1: v35}
                                _ = v41
                            default:
                                v2 += 130
                            }
                            v17 = float32(1.5)
                        }
                    case -1:
                        for fi42 := 0; fi42 < 9; fi42++ {
                            switch fi42 {
                            case 32768:
                                v4 = -1
                            case -2147483648:
                                var v43 string = ""
                                _ = v43
                            }
                            if (!(!false)) {
                                v2 = (v2 << ((42) & 63))
                                var v44 float64 = 1.5
                                _ = v44
                            }
                            _ = fi42
                        }
                    default:
                        v45 := v5[(("" + (("" + ("a" + (("hello" + "hello") + ("hello" + ((("foo bar" + "foo bar") + (("foo bar" + "") + "\n\t")) + "\n\t"))))) + "hello")) + (("a" + "\n\t") + "a"))]
                        _ = v45
                    }
                    v4 += v4
                    _ = v35
                } else {
                    for fi46 := 0; fi46 < 7; fi46++ {
                        var v47 float32 = float32((0.1 * 0.1))
                        _ = fi46
                        _ = v47
                    }
                }
                var v48 int8 = 100
                _ = v34
                _ = v48
            } else {
                v1 = (-(v4))
            }
        }
        _ = v17
    }
    v5[""] = (960 >= 42)
    v3 = append(v3, "hello")
    for fi49 := 0; fi49 < 20; fi49++ {
        v50 := GenS2{F0: 4294967295}
        _ = fi49
        _ = v50
    }
    v4 += (^(v4))
    v3 = append(v3, "\n\t")
    var v51 bool = (((genFn0(1, v1, 228) == -2147483648) || (1.5 <= 3.14159)) && true)
    var v52 string = ("hello" + "foo bar")
    var v53 bool = false
    _ = len(v3)
    if (((152 <= 0) && (4294967295 != 42)) && v51) {
        var v54 uint32 = 968
        v55 := v5[("a" + "hello")]
        _ = v54
        _ = v55
    }
    var v56 bool = v53
    fmt.Println(v53)
    for fi57 := 0; fi57 < 14; fi57++ {
        for fi58 := 0; fi58 < 19; fi58++ {
            v59 := GenS1{F0: v52, F1: v53, F2: -2147483648}
            if false {
                if (!(0.1 == 0.1)) {
                    if len(v3) > 0 {
                        v3[0] = (v52 + "a")
                    }
                    if (v56 && false) {
                        fmt.Println(v52)
                        v2 -= v2
                        var v60 uint8 = 103
                        _ = v60
                    } else {
                        v53 = v51
                        fmt.Println(fi57)
                        v56 = true
                    }
                } else {
                    var v61 float32 = (float32(1e10) - 0.0)
                    v61 = float32(1e10)
                    if (-2147483648 != -2147483648) {
                        var v62 uint = 4294967295
                        v52 = v52
                        var v63 int16 = v4
                        _ = v62
                        _ = v63
                    }
                    _ = v61
                }
                var v64 int16 = (v4 * (v4 + (v4 + (^(v4)))))
                _ = v64
            } else {
                var v65 int16 = (v1 << ((4294967295) & 63))
                _ = v65
            }
            if (fi57 > (^(fi57))) {
                v4 = v4
                var v66 int32 = 992
                _ = v66
            }
            _ = fi58
            _ = v59
        }
        v51 = v56
        if true {
            v1 += (v1 + v1)
        } else {
            var v67 int16 = (^(v4))
            switch v52 {
            case "foo bar":
                switch "hello" {
                case "foo bar":
                    for fi68 := 0; fi68 < 15; fi68++ {
                        fmt.Println(v4)
                        var v69 int8 = 100
                        v69 = -25
                        _ = fi68
                        _ = v69
                    }
                case "hello":
                    switch (fi57 & genFn0(42, (v4 ^ 0), 128)) {
                    case 9223372036854775807:
                        var v70 float64 = (0.0 - (1.5 / float64(3.14159)))
                        _ = v70
                    case -240:
                        var v71 uint16 = 133
                        _ = v71
                    }
                case "":
                    var v72 string = ("\n\t" + (v52 + "foo bar"))
                    _ = v72
                default:
                    var v73 float32 = (0.0 - 0.1)
                    _ = v73
                }
            case "a":
                v1 = 0
            }
            _ = v67
        }
        _ = fi57
    }
    switch v52 {
    case "a":
        fmt.Println(v1, v4, v1)
    case "\n\t":
        _ = len(v3)
    case "foo bar":
        fmt.Println(v2, v52, v4)
    default:
        _ = len(v5)
    }
    v56 = v51
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v56
}
