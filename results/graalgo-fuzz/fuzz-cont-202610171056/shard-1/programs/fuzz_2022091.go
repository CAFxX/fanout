package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint }
func genFn0(p0 bool, p1 int16, p2 uint8) uint64 {
    switch -1 {
    case 9223372036854775807:
        p2 = (p2 + (p2 % (((p2 >> ((255) & 63))) | 1)))
    }
    var v5 float64 = ((((-2.25 - 0.0) - ((3.14159 / (1e10 * -2.25)) * float64(1.5))) + 3.14159) / (-2.25 / float64(-2.25)))
    var v6 string = "foo bar"
    for fi7 := 0; fi7 < 18; fi7++ {
        if (100 > -1) {
            v8 := []int{}
            if p0 {
                v8 = append(v8, (fi7 * (-(fi7))))
            } else {
                v9 := GenS1{F0: 18446744073709551615}
                var v10 int8 = -71
                var v11 uint = 0
                _ = v10
                _ = v11
                _ = v9
            }
            _ = v8
        } else {
            for fi12 := 0; fi12 < 18; fi12++ {
                var v13 int32 = 670
                var v14 int8 = 1
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        _ = fi7
    }
    var v15 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v5
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 float64 = (3.14159 * -2.25)
    v2 := make(map[string]uint32)
    v3 := []int{128, 823, 2147483647}
    var v4 int8 = -128
    var v16 int64 = 0
    var v17 uint = 562
    var v18 int = 127
    var v19 uint64 = 4294967295
    v20 := GenS2{F0: "foo bar"}
    var v21 int32 = 42
    var v22 int32 = v21
    var v23 float64 = 0.1
    v2["foo bar"] = 717
    for fi24 := 0; fi24 < 5; fi24++ {
        if false {
            v22 += (v21 << ((27) & 63))
            v1 -= float64(0.1)
        } else {
            switch 264 {
            case 0:
                delete(v2, "\n\t")
            case 9223372036854775807:
                switch "" {
                case "hello":
                    v25 := GenS2{F0: (("foo bar" + "\n\t") + "hello")}
                    _ = v25
                default:
                    if true {
                        for fi26 := 0; fi26 < 18; fi26++ {
                            var v27 int64 = 0
                            v28 := GenS3{F0: 42, F1: 976}
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        v29 := GenS4{F0: v4, F1: 94}
                        if len(v3) > 0 {
                            v3[0] = 42
                        }
                        _ = v29
                    } else {
                        if false {
                            switch fi24 {
                            case -438:
                                v30 := GenS5{F0: "", F1: (v22 / (((v22 * -657)) | 1))}
                                _ = v30
                            }
                            fmt.Println(v18, v18)
                            var v31 string = "foo bar"
                            _ = v31
                        }
                        v32 := v2["\n\t"]
                        if true {
                            if false {
                                var v33 int64 = (v16 - v16)
                                _ = v33
                            } else {
                                var v34 bool = true
                                var v35 int16 = 32767
                                _ = v34
                                _ = v35
                            }
                        }
                        _ = v32
                    }
                }
            }
            fmt.Println(v23, v23, fi24)
        }
        _ = fi24
    }
    var v36 uint64 = (v19 ^ v19)
    switch "hello" {
    case "hello":
        for fi37 := 0; fi37 < 14; fi37++ {
            if (v21 == 42) {
                if (v21 == v22) {
                    for fi38 := 0; fi38 < 9; fi38++ {
                        var v39 uint64 = (^(v36))
                        _ = fi38
                        _ = v39
                    }
                    if false {
                        v22 -= v22
                        v36 += genFn0(true, -32768, 1)
                        _ = len(v3)
                    } else {
                        v21 = (v22 * (^(v22)))
                        for fi40 := 0; fi40 < 12; fi40++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi40
                        }
                        delete(v2, ("\n\t" + "\n\t"))
                    }
                }
            } else {
                for fi41 := 0; fi41 < 3; fi41++ {
                    switch ("a" + "") {
                    case "":
                        v23 += (v1 + -2.25)
                    default:
                        v4 = (^(v4))
                    }
                    v4 -= 42
                    var v42 float32 = 1e10
                    _ = fi41
                    _ = v42
                }
                switch (fi37 ^ fi37) {
                case -2147483648:
                    v17 -= 140
                case 42:
                    for fi43 := 0; fi43 < 15; fi43++ {
                        switch "\n\t" {
                        case "foo bar":
                            _ = len(v3)
                        case "a":
                            fmt.Println(v17, v19)
                        case "hello":
                            v44 := GenS4{F0: 81, F1: 2}
                            v4 = 125
                            _ = v44
                        }
                        var v45 int = (fi37 - fi37)
                        _ = fi43
                        _ = v45
                    }
                default:
                    fmt.Println(v16)
                }
            }
            delete(v2, "a")
            v46 := GenS4{F0: (v4 * v4), F1: 2}
            _ = fi37
            _ = v46
        }
    case "":
        switch "\n\t" {
        case "":
            switch 0 {
            case 32768:
                v47 := GenS6{F0: -1}
                _ = v47
            case -1:
                if ((v1 >= v1) || false) {
                    switch (v18 << (((v17 / ((v17) | 1))) & 63)) {
                    case 4294967295:
                        var v48 int = (v18 - v18)
                        _ = v48
                    }
                } else {
                    v49 := GenS7{F0: 18446744073709551615}
                    _ = v49
                }
            default:
                v19 -= (v19 % ((v19) | 1))
            }
        default:
            v1 -= (((float64((3.14159 / -2.25)) + (1.5 * (v1 + v1))) + -2.25) * float64(3.14159))
        }
    }
    v4 = (v4 % ((49) | 1))
    if (genFn0(true, -93, 255) <= v19) {
        switch "\n\t" {
        case "":
            if (-924 == (v16 - 1)) {
                var v50 uint16 = 2
                var v51 float32 = 0.0
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v50
                _ = v51
            } else {
                if false {
                    for fi52 := 0; fi52 < 2; fi52++ {
                        var v53 float64 = (-2.25 + v1)
                        _ = len(v2)
                        fmt.Println(v16)
                        _ = fi52
                        _ = v53
                    }
                } else {
                    var v54 int16 = -470
                    switch 128 {
                    case -9223372036854775808:
                        var v55 string = "foo bar"
                        _ = v55
                    case -662:
                        v56 := GenS8{F0: v17}
                        _ = v56
                    }
                    _ = v54
                }
                if false {
                    _ = len(v3)
                    switch v18 {
                    case 32768:
                        v22 -= (^(v22))
                    case -22:
                        var v57 string = "hello"
                        _ = v57
                    default:
                        if false {
                            var v58 uint32 = 805
                            v59 := v2[""]
                            v59 = v58
                            _ = v58
                            _ = v59
                        } else {
                            var v60 uint32 = 617
                            _ = v60
                        }
                    }
                    v36 += (v36 >> ((v17) & 63))
                } else {
                    var v61 uint64 = v19
                    v61 = v61
                    _ = v61
                }
            }
        }
        fmt.Println(v22, v36)
        v62 := GenS5{F0: (("\n\t" + ((("" + "\n\t") + "foo bar") + ("" + "hello"))) + "foo bar"), F1: v22}
        _ = v62
    }
    v36 -= 0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v36
    _ = v4
}
