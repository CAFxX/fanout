package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 uint }

func main() {
    var v1 uint = 65535
    var v2 int8 = 1
    v3 := make(map[uint32]bool)
    v4 := []string{}
    var v5 uint64 = 2
    var v6 uint = (v1 << ((255) & 63))
    v7 := v3[0]
    v8 := GenS1{F0: -447}
    v6 -= 255
    for fi9 := 0; fi9 < 8; fi9++ {
        switch (fi9 << ((v6) & 63)) {
        case 32768:
            for fi10 := 0; fi10 < 15; fi10++ {
                v1 += (^(v1))
                _ = len(v4)
                v11 := GenS1{F0: 2147483647}
                _ = fi10
                _ = v11
            }
        default:
            v12 := GenS2{F0: 236, F1: (v5 | (v5 | 42))}
            _ = v12
        }
        v2 -= (-(v2))
        if v7 {
            v5 -= v5
        }
        _ = fi9
    }
    switch 4294967295 {
    case -22:
        v2 = (-(v2))
    case -1:
        var v13 uint8 = 2
        _ = v13
    case -497:
        v14 := v3[42]
        _ = v14
    }
    switch ("foo bar" + ("" + "a")) {
    case "hello":
        for fi15 := 0; fi15 < 8; fi15++ {
            for fi16 := 0; fi16 < 1; fi16++ {
                v7 = false
                switch fi15 {
                case -942:
                    v17 := v3[635]
                    _ = v17
                case -9223372036854775808:
                    switch (^(fi16)) {
                    case -1:
                        v18 := GenS2{F0: 2, F1: (^(v5))}
                        if ("" < ("a" + "hello")) {
                            v19 := GenS3{F0: 42, F1: "hello"}
                            if (((v7 && (v7 || (v7 || v7))) || (((0.0 - 0.1) - 3.14159) >= 1.5)) && true) {
                                v4 = append(v4, "")
                                v5 -= v5
                                var v20 int16 = 1
                                _ = v20
                            } else {
                                v2 = (v2 >> (((^(v6))) & 63))
                                switch (^(fi15)) {
                                case 4294967295:
                                    var v21 uint16 = 171
                                    _ = v21
                                case -9223372036854775808:
                                    v22 := GenS1{F0: -2147483648}
                                    _ = v22
                                case -415:
                                    if (0.0 < 1e10) {
                                        v5 = uint64(v6)
                                        var v23 float64 = ((1e10 * ((1e10 / float64(1e10)) + 1.5)) * 0.1)
                                        _ = v23
                                    }
                                    if (!(true || false)) {
                                        var v24 int32 = 660
                                        v25 := GenS4{F0: -9223372036854775808, F1: fi16}
                                        _ = v24
                                        _ = v25
                                    } else {
                                        v6 += v6
                                    }
                                default:
                                    v2 += (v2 + (-(v2)))
                                }
                                var v26 bool = v7
                                _ = v26
                            }
                            _ = v19
                        }
                        v3[49] = ("hello" == "a")
                        _ = v18
                    default:
                        _ = len(v4)
                    }
                default:
                    if true {
                        var v27 int = (fi16 * -2147483648)
                        _ = v27
                    } else {
                        _ = len(v4)
                        v4 = append(v4, (("hello" + "foo bar") + ""))
                        fmt.Println(v5, fi16)
                    }
                    v28 := GenS5{F0: 42, F1: 2, F2: v5}
                    v29 := GenS6{F0: true, F1: 32767}
                    var v30 int32 = 2147483647
                    _ = v28
                    _ = v29
                    _ = v30
                }
                _ = fi16
            }
            var v31 bool = (((float32(1e10) + float32((0.0 / (3.14159 / 1.0)))) - (1e10 / (3.14159 * float32((3.14159 + (-2.25 / (1.5 / (float64(0.1) / 1.0)))))))) != (0.0 * (1.5 - 0.1)))
            var v32 int64 = -96
            _ = fi15
            _ = v31
            _ = v32
        }
    case "\n\t":
        var v33 uint32 = 1
        _ = v33
    case "":
        var v34 bool = v7
        _ = v34
    default:
        var v35 uint32 = 938
        _ = v35
    }
    v36 := GenS7{F0: (v6 & (v6 ^ 349))}
    for fi37 := 0; fi37 < 4; fi37++ {
        v38 := v3[861]
        switch "\n\t" {
        case "foo bar":
            switch fi37 {
            case 32768:
                fmt.Println(v2, v5)
            case 11:
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
            default:
                var v39 uint = 42
                _ = v39
            }
        case "hello":
            var v40 bool = v7
            _ = v40
        default:
            v41 := v3[15]
            _ = v41
        }
        for fi42 := 0; fi42 < 6; fi42++ {
            if (!v7) {
                v43 := GenS4{F0: 0, F1: fi42}
                _ = len(v3)
                v1 = (v6 << ((v1) & 63))
                _ = v43
            }
            if true {
                if v38 {
                    v38 = false
                }
            } else {
                var v44 string = ""
                var v45 uint8 = 54
                _ = v44
                _ = v45
            }
            _ = len(v3)
            _ = fi42
        }
        _ = fi37
        _ = v38
    }
    v46 := GenS6{F0: true, F1: 2}
    v7 = (0.0 > 0.0)
    var v47 int = -1
    var v48 int64 = 9223372036854775807
    var v49 bool = (v7 || (v2 < (v2 * -1)))
    v47 += (v47 >> ((22) & 63))
    for fi50 := 0; fi50 < 17; fi50++ {
        fmt.Println(v47)
        var v51 int16 = 889
        if v49 {
            v2 = 0
            delete(v3, 42)
        }
        _ = fi50
        _ = v51
    }
    for fi52 := 0; fi52 < 5; fi52++ {
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
        switch "\n\t" {
        case "\n\t":
            v1 = v6
        case "foo bar":
            var v53 bool = v49
            _ = v53
        default:
            v54 := GenS1{F0: (v47 & fi52)}
            _ = v54
        }
        _ = fi52
    }
    v2 -= (v2 % ((100) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
