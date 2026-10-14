package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 int }
func genFn0(p0 int, p1 uint64, p2 bool) uint {
    switch ("\n\t" + "\n\t") {
    case "hello":
        var v7 int8 = 127
        _ = v7
    default:
        v8 := GenS1{F0: 327}
        _ = v8
    }
    if true {
        var v9 float64 = float64(1e10)
        v10 := []int{}
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := []int{32767, p0, (p0 / (((p0 & (^(fi11)))) | 1)), 128}
            _ = fi11
            _ = v12
        }
        var v13 uint16 = 0
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        p0 -= (p0 ^ 1)
        v15 := []int{}
        _ = fi14
        _ = v15
    }
    if p2 {
        switch "\n\t" {
        case "foo bar":
            p0 -= (p0 << ((65535) & 63))
        case "hello":
            switch p0 {
            case 4294967295:
                var v16 string = "\n\t"
                _ = v16
            default:
                v17 := GenS2{F0: 577}
                _ = v17
            }
        case "\n\t":
            var v18 int32 = 0
            _ = v18
        default:
            var v19 uint8 = 2
            _ = v19
        }
    } else {
        if ((!((p1 & (^(p1))) == (p1 << ((18446744073709551615) & 63)))) || p2) {
            var v20 uint16 = 255
            _ = v20
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 int = 4294967295
            _ = fi21
            _ = v22
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1() uint64 {
    var v23 float64 = float64(0.1)
    v23 -= float64(-2.25)
    v24 := []int{}
    v23 += 1e10
    _ = v23
    _ = v24
    return 42
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    var v3 uint8 = 128
    var v4 uint16 = 65535
    var v5 uint16 = 1
    v6 := []string{}
    v6 = append(v6, "foo bar")
    v1[("foo bar" + "")] = "hello"
    if ((!true) && false) {
        v25 := v1[(("\n\t" + ("hello" + "hello")) + ("foo bar" + "\n\t"))]
        _ = v25
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 int = 33
            switch "hello" {
            case "foo bar":
                var v29 bool = (!false)
                _ = v29
            }
            var v30 int16 = 1
            _ = fi27
            _ = v28
            _ = v30
        }
        fmt.Println(v3, v4, v3)
        _ = fi26
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        v32 := GenS1{F0: int32(v3)}
        for fi33 := 0; fi33 < 17; fi33++ {
            var v34 int16 = 1000
            _ = fi33
            _ = v34
        }
        switch "" {
        case "a":
            if false {
                switch "hello" {
                case "":
                    v35 := GenS2{F0: 42}
                    _ = v35
                case "\n\t":
                    v36 := GenS3{F0: "\n\t"}
                    _ = v36
                case "hello":
                    v3 += v3
                }
                v6 = append(v6, "foo bar")
                v37 := GenS4{F0: 737, F1: 818}
                _ = v37
            } else {
                var v38 string = "hello"
                var v39 uint32 = 2
                fmt.Println(v38, v38, fi31)
                _ = v38
                _ = v39
            }
        case "hello":
            for fi40 := 0; fi40 < 10; fi40++ {
                v4 = uint16(v3)
                v5 += 42
                _ = fi40
            }
        case "\n\t":
            if (18446744073709551615 > 0) {
                v6 = append(v6, "")
                for fi41 := 0; fi41 < 10; fi41++ {
                    fmt.Println(v5, v5)
                    v42 := GenS5{F0: (!true), F1: -931}
                    fmt.Println(v3, v5)
                    _ = fi41
                    _ = v42
                }
                fmt.Println(fi31, fi31)
            } else {
                var v43 int8 = -128
                var v44 uint32 = 589
                var v45 uint16 = v4
                _ = v43
                _ = v44
                _ = v45
            }
        }
        _ = fi31
        _ = v32
    }
    for fi46 := 0; fi46 < 9; fi46++ {
        fmt.Println(fi46)
        for fi47 := 0; fi47 < 1; fi47++ {
            if len(v2) > 0 {
                v2[0] = "a"
            }
            if (!false) {
                var v48 uint32 = 969
                fmt.Println(v4, v48)
                _ = v48
            }
            _ = fi47
        }
        for fi49 := 0; fi49 < 5; fi49++ {
            v50 := GenS3{F0: "a"}
            var v51 float64 = 0.0
            _ = fi49
            _ = v50
            _ = v51
        }
        _ = fi46
    }
    var v52 float64 = 1e10
    var v53 uint64 = genFn1()
    for fi54 := 0; fi54 < 16; fi54++ {
        var v55 string = ("\n\t" + "\n\t")
        _ = fi54
        _ = v55
    }
    var v56 int8 = 108
    v57 := GenS6{F0: 0}
    switch 42 {
    case 127:
        var v58 float64 = (1e10 / v52)
        _ = v58
    case 789:
        if (1 > (v53 ^ v53)) {
            switch "hello" {
            case "a":
                for fi59 := 0; fi59 < 2; fi59++ {
                    v53 = (v53 & genFn1())
                    if (true || false) {
                        v3 -= 1
                        fmt.Println(fi59, v4, v5)
                        fmt.Println(v52, v52)
                    } else {
                        fmt.Println(v5, v4)
                        var v60 uint = 255
                        _ = v60
                    }
                    var v61 int32 = 1
                    _ = fi59
                    _ = v61
                }
            default:
                v62 := GenS6{F0: 9223372036854775807}
                _ = v62
            }
            switch -584 {
            case 32767:
                v1[""] = ("foo bar" + ("" + "\n\t"))
            case 1:
                v4 += 255
            default:
                if (false && (!(v52 == 1.5))) {
                    v4 += v4
                    switch (((("foo bar" + (("a" + ("\n\t" + "foo bar")) + "hello")) + "a") + "\n\t") + ("" + "a")) {
                    case "a":
                        switch 0 {
                        case -622:
                            v63 := GenS2{F0: int32(v5)}
                            _ = v63
                        }
                    case "\n\t":
                        var v64 float64 = (float64((-2.25 - float32(v52))) + v52)
                        _ = v64
                    case "hello":
                        v3 -= 107
                    default:
                        v52 = 0.1
                    }
                    var v65 int16 = 0
                    _ = v65
                } else {
                    fmt.Println(v5, v53, v52)
                    for fi66 := 0; fi66 < 12; fi66++ {
                        if (!(-9223372036854775808 > -9223372036854775808)) {
                            if ("hello" < "a") {
                                var v67 int = (fi66 & fi66)
                                _ = v67
                            }
                            switch fi66 {
                            case 4294967295:
                                var v68 uint64 = uint64(v3)
                                _ = v68
                            case -9223372036854775808:
                                var v69 int64 = -84
                                _ = v69
                            default:
                                switch (fi66 / ((fi66) | 1)) {
                                case -1:
                                    var v70 int = 1
                                    _ = v70
                                case 128:
                                    v71 := GenS3{F0: "\n\t"}
                                    _ = v71
                                }
                            }
                            for fi72 := 0; fi72 < 3; fi72++ {
                                v4 += (v4 << ((984) & 63))
                                v73 := GenS1{F0: 690}
                                _ = fi72
                                _ = v73
                            }
                        }
                        v52 = 0.1
                        var v74 int = (fi66 * (fi66 ^ int(v56)))
                        _ = fi66
                        _ = v74
                    }
                }
            }
            if (false && true) {
                v3 = (v3 | 42)
                v5 -= v4
                if true {
                    fmt.Println(v53, v5)
                }
            }
        } else {
            v75 := v1[((("" + "foo bar") + "hello") + "\n\t")]
            var v76 int8 = v56
            _ = v75
            _ = v76
        }
    case -1:
        v53 += 691
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v77 int8 = (v56 ^ (v56 / ((v56) | 1)))
    for fi78 := 0; fi78 < 9; fi78++ {
        for fi79 := 0; fi79 < 15; fi79++ {
            v77 = (-(v56))
            var v80 uint64 = v53
            _ = fi79
            _ = v80
        }
        var v81 uint8 = v3
        _ = fi78
        _ = v81
    }
    v82 := GenS2{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v56
    _ = v57
    _ = v6
    _ = v77
    _ = v82
}
