package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int }
type GenS2 struct { F0 int8; F1 uint64; F2 int8 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 bool; F1 string; F2 uint16 }
func genFn0(p0 int16) int {
    fmt.Println(p0)
    v5 := GenS1{F0: 0, F1: (p0 - p0), F2: 1}
    var v6 int64 = -1
    fmt.Println(p0, v6)
    v6 -= (v6 & 42)
    _ = p0
    _ = v5
    _ = v6
    return 128
}
func genFn1(p0 int16) uint64 {
    v7 := []int{420, 563, -1}
    p0 = 0
    p0 -= p0
    _ = p0
    _ = v7
    return 91
}

func main() {
    v1 := []string{("" + "hello"), "", "a"}
    var v2 int8 = 1
    v3 := []uint32{}
    v4 := make(map[uint32]int)
    _ = len(v4)
    if true {
        v8 := GenS1{F0: 1, F1: 0, F2: genFn0(42)}
        if (v2 >= v2) {
            var v9 uint32 = 2147483647
            _ = v9
        } else {
            fmt.Println(v2)
            for fi10 := 0; fi10 < 19; fi10++ {
                v2 -= 0
                _ = fi10
            }
            var v11 int64 = -1
            _ = v11
        }
        _ = v8
    }
    switch genFn0(-32768) {
    case 42:
        var v12 uint8 = 1
        _ = v12
    default:
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 float32 = ((0.0 + 0.1) + 1e10)
            if ("" <= ("\n\t" + (("hello" + "a") + ("" + (("foo bar" + "\n\t") + "\n\t"))))) {
                fmt.Println(fi13, fi13, v14)
                for fi15 := 0; fi15 < 15; fi15++ {
                    if (false || (902 == 381)) {
                        var v16 int16 = 0
                        fmt.Println(fi13)
                        _ = v16
                    }
                    _ = fi15
                }
            } else {
                v17 := GenS1{F0: 2147483647, F1: 1, F2: -244}
                _ = v17
            }
            _ = fi13
            _ = v14
        }
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        switch (fi18 & fi18) {
        case 32767:
            switch 9223372036854775807 {
            case 32768:
                v19 := GenS2{F0: (v2 ^ (v2 ^ (-(v2)))), F1: 42, F2: 42}
                _ = v19
            case 128:
                if true {
                    v20 := GenS2{F0: (-(v2)), F1: 4294967295, F2: v2}
                    _ = v20
                }
            case 42:
                for fi21 := 0; fi21 < 5; fi21++ {
                    if ((!false) || false) {
                        var v22 int64 = 0
                        v23 := GenS1{F0: 34, F1: -32768, F2: -1}
                        var v24 uint32 = 0
                        v3 = append(v3, 42)
                        v25 := GenS3{F0: -1, F1: 609}
                        _ = v22
                        _ = v23
                        _ = v24
                        _ = v25
                    } else {
                        var v26 uint16 = 561
                        switch "foo bar" {
                        case "":
                            var v27 uint = 4294967295
                            _ = v27
                        default:
                            var v28 float32 = (float32((float64(0.1) * 0.1)) + -2.25)
                            _ = v28
                        }
                        _ = v26
                    }
                    if true {
                        v2 -= v2
                        v29 := GenS3{F0: (v2 & (v2 | v2)), F1: 32767}
                        switch (-(fi21)) {
                        case 127:
                            var v30 float64 = ((1e10 - (1e10 / float64(3.14159))) * (0.0 - float64(1e10)))
                            _ = v30
                        case 9223372036854775807:
                            var v31 int8 = v2
                            _ = v31
                        case -1:
                            var v32 uint32 = 4294967295
                            _ = v32
                        default:
                            v2 = v2
                        }
                        _ = v29
                    } else {
                        switch -216 {
                        case 4294967295:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        var v33 uint8 = 127
                        v3 = append(v3, 210)
                        _ = v33
                    }
                    v4[1] = (fi18 % ((fi18) | 1))
                    _ = fi21
                }
            default:
                for fi34 := 0; fi34 < 17; fi34++ {
                    v4[2147483648] = (fi18 % ((fi34) | 1))
                    var v35 int8 = (v2 ^ (v2 >> ((4294967295) & 63)))
                    _ = fi34
                    _ = v35
                }
            }
        }
        _ = fi18
    }
    var v36 int8 = v2
    var v37 int16 = 201
    var v38 int = 128
    for fi39 := 0; fi39 < 3; fi39++ {
        if (!false) {
            switch "\n\t" {
            case "":
                var v40 uint = 575
                _ = v40
            case "\n\t":
                for fi41 := 0; fi41 < 4; fi41++ {
                    switch "hello" {
                    case "a":
                        v37 += (v37 - 1000)
                    case "":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case "\n\t":
                        v36 = (v2 + (v36 >> ((65535) & 63)))
                    default:
                        var v42 float64 = ((-2.25 - -2.25) + (3.14159 * (float64(1.5) / 1.5)))
                        var v43 float64 = 0.1
                        _ = v42
                        _ = v43
                    }
                    switch "\n\t" {
                    case "":
                        if len(v3) > 0 {
                            v3[0] = 4294967295
                        }
                    case "\n\t":
                        switch "\n\t" {
                        case "":
                            var v44 int = (fi39 ^ (v38 ^ 755))
                            _ = v44
                        case "foo bar":
                            v38 += fi41
                        }
                    case "foo bar":
                        v38 = (fi39 << ((225) & 63))
                    default:
                        v2 = -128
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi41
                }
            case "hello":
                switch ("hello" + ("foo bar" + (("\n\t" + ((("" + "\n\t") + "\n\t") + "foo bar")) + ("hello" + "foo bar")))) {
                case "hello":
                    fmt.Println(v2, fi39, v37)
                case "a":
                    v37 = (v37 % ((v37) | 1))
                }
            }
            v38 = (^(v38))
        } else {
            if ((true || true) || false) {
                var v45 uint32 = 1
                _ = v45
            }
            v4[65535] = v38
        }
        _ = len(v4)
        if true {
            v37 = 0
            if (true || (127 < 191)) {
                switch "hello" {
                case "a":
                    if ((1.5 / ((float64((1.5 + (3.14159 - -2.25))) + (float64(0.1) / 3.14159)) * 0.1)) < 1.5) {
                        switch ("foo bar" + "hello") {
                        case "foo bar":
                            if ((255 < 0) && (1 < -9223372036854775808)) {
                                var v46 uint = 42
                                if (((true && (((-2.25 / 1.0) - (float64(3.14159) + 0.0)) != ((((3.14159 * 0.0) * 1e10) / 0.1) + (((float64((((float32((3.14159 + ((1e10 + 3.14159) - float64(1e10)))) / -2.25) - -2.25) * 3.14159)) - 1.5) * -2.25) + 1.5)))) && true) || true) {
                                    v3 = append(v3, 2)
                                    var v47 int8 = (v2 << (((v46 + (v46 * 0))) & 63))
                                    _ = v47
                                }
                                _ = v46
                            }
                        case "\n\t":
                            if (2 > 255) {
                                fmt.Println(v2, v2, v37)
                                v38 -= v38
                                var v48 int8 = v36
                                _ = v48
                            }
                        case "a":
                            v4[1] = v38
                        default:
                            var v49 float32 = float32(float64(3.14159))
                            _ = v49
                        }
                        switch "" {
                        case "\n\t":
                            _ = len(v3)
                        case "":
                            v50 := v4[42]
                            _ = v50
                        default:
                            v36 -= v36
                        }
                    }
                case "foo bar":
                    if ((-(v2)) >= (^(v36))) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v51 int32 = 0
                        var v52 uint64 = 42
                        _ = v51
                        _ = v52
                    } else {
                        v37 += (v37 * (v37 | (v37 >> ((583) & 63))))
                    }
                }
            }
            for fi53 := 0; fi53 < 19; fi53++ {
                v54 := GenS4{F0: ((1 >= 2147483647) || ("hello" < (("" + ("foo bar" + ("a" + "foo bar"))) + "\n\t"))), F1: "foo bar", F2: 131}
                _ = fi53
                _ = v54
            }
        } else {
            v2 = (v2 & (^(v2)))
        }
        _ = fi39
    }
    var v55 uint16 = 1
    var v56 float32 = 1e10
    for fi57 := 0; fi57 < 4; fi57++ {
        v55 = (v55 & (v55 % ((v55) | 1)))
        _ = len(v4)
        _ = fi57
    }
    v38 = 42
    _ = len(v1)
    var v58 int = 128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v55
    _ = v56
    _ = v58
}
