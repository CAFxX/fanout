package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64 }
func genFn0(p0 int16) uint8 {
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(p0, fi7)
        var v8 uint = 65535
        v8 -= v8
        _ = fi7
        _ = v8
    }
    switch (("" + "") + "hello") {
    case "":
        p0 = p0
    }
    v9 := []int{173}
    _ = p0
    _ = v9
    return 255
}

func main() {
    var v1 uint16 = 32767
    var v2 int64 = -256
    var v3 bool = (((!true) && true) || true)
    var v4 uint16 = 65535
    var v5 int8 = -1
    v6 := []uint32{}
    v6 = append(v6, 65535)
    v1 = 2
    for fi10 := 0; fi10 < 10; fi10++ {
        if true {
            v11 := GenS1{F0: fi10, F1: "\n\t"}
            if true {
                var v12 int8 = (v5 | -128)
                if (v3 || v3) {
                    if (v3 || (1e10 == -2.25)) {
                        var v13 int8 = (v5 << ((489) & 63))
                        switch fi10 {
                        case -1:
                            if (!(v2 > int64(v4))) {
                                for fi14 := 0; fi14 < 3; fi14++ {
                                    var v15 uint = 255
                                    var v16 int32 = -586
                                    var v17 int32 = -1
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                }
                            }
                        default:
                            for fi18 := 0; fi18 < 10; fi18++ {
                                switch fi10 {
                                case -9223372036854775808:
                                    v6 = append(v6, 65535)
                                case -220:
                                    _ = len(v6)
                                case 372:
                                    var v19 int32 = 42
                                    _ = v19
                                default:
                                    var v20 int16 = 1
                                    _ = v20
                                }
                                switch "" {
                                case "\n\t":
                                    var v21 int = (^(fi18))
                                    _ = v21
                                }
                                var v22 int16 = -1
                                var v23 int64 = (v2 | v2)
                                _ = fi18
                                _ = v22
                                _ = v23
                            }
                        }
                        _ = v13
                    } else {
                        if len(v6) > 0 {
                            v6[0] = 582
                        }
                    }
                    switch 2147483647 {
                    case 9223372036854775807:
                        fmt.Println(v1, v2)
                    case 4294967295:
                        if len(v6) > 0 {
                            v6[0] = 4294967295
                        }
                    }
                }
                _ = len(v6)
                _ = v12
            } else {
                for fi24 := 0; fi24 < 8; fi24++ {
                    var v25 uint8 = 39
                    v6 = append(v6, 2)
                    var v26 float32 = (0.0 - 1.5)
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                for fi27 := 0; fi27 < 11; fi27++ {
                    switch "foo bar" {
                    case "":
                        if (1 >= v2) {
                            var v28 int32 = -1
                            switch 307 {
                            case -9223372036854775808:
                                v28 = 0
                                var v29 uint64 = 637
                                v30 := GenS2{F0: fi10, F1: (v5 ^ v5)}
                                _ = v29
                                _ = v30
                            }
                            v1 -= (v4 / (((v1 / (((v1 - (v4 / (((v4 / ((v1) | 1))) | 1)))) | 1))) | 1))
                            _ = v28
                        }
                        fmt.Println(v2, fi10)
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    v4 -= 0
                    v3 = v3
                    v2 = v2
                    v31 := GenS2{F0: (fi27 | fi27), F1: (v5 + v5)}
                    _ = fi27
                    _ = v31
                }
            }
            _ = v11
        } else {
            switch fi10 {
            case 42:
                var v32 int64 = (v2 % (((^(v2))) | 1))
                _ = v32
            case -1:
                if (!(!(((v5 ^ (v5 % ((42) | 1))) < v5) || v3))) {
                    if (1 != 0) {
                        v3 = false
                    } else {
                        v33 := GenS3{F0: "foo bar"}
                        if v3 {
                            var v34 uint8 = 112
                            v5 = v5
                            v4 -= 853
                            _ = v34
                        }
                        _ = v33
                    }
                    v35 := GenS2{F0: 128, F1: (^(v5))}
                    _ = v35
                }
            default:
                if len(v6) > 0 {
                    v6[0] = 1
                }
            }
            var v36 float32 = ((((((float32(1.5) / (1e10 / 1e10)) + (((0.1 - 1e10) / float32(1e10)) * 0.0)) * float32((-2.25 / (0.0 - 1e10)))) * 0.0) - (0.1 + 1.5)) + float32(0.1))
            _ = v36
        }
        var v37 float64 = ((0.1 + ((((3.14159 - 0.0) / 1.5) * (float64(3.14159) / float64(1e10))) - 1e10)) / (((0.1 - 1.5) - (-2.25 / 3.14159)) + -2.25))
        switch (("a" + "a") + "a") {
        case "hello":
            var v38 uint8 = 2
            _ = v38
        case "\n\t":
            for fi39 := 0; fi39 < 4; fi39++ {
                var v40 int16 = 762
                _ = fi39
                _ = v40
            }
        }
        _ = fi10
        _ = v37
    }
    if v3 {
        if true {
            switch "hello" {
            case "\n\t":
                v41 := GenS4{F0: (v2 | 9223372036854775807)}
                _ = v41
            case "a":
                var v42 uint32 = 65535
                _ = v42
            case "":
                switch -9223372036854775808 {
                case 42:
                    fmt.Println(v5)
                default:
                    switch "\n\t" {
                    case "":
                        fmt.Println(v2)
                    default:
                        v6 = append(v6, 475)
                    }
                }
            }
            if (!v3) {
                var v43 int = 1
                v44 := GenS3{F0: ""}
                _ = v43
                _ = v44
            } else {
                var v45 float32 = (-2.25 - float32(3.14159))
                switch "hello" {
                case "":
                    for fi46 := 0; fi46 < 14; fi46++ {
                        var v47 float32 = v45
                        _ = fi46
                        _ = v47
                    }
                }
                _ = v45
            }
            fmt.Println(v1)
        } else {
            var v48 float32 = 0.0
            switch ("\n\t" + ("\n\t" + "a")) {
            case "hello":
                switch (("a" + "hello") + ("hello" + ("hello" + "foo bar"))) {
                case "hello":
                    var v49 uint = 848
                    _ = v49
                case "\n\t":
                    v2 = (v2 ^ v2)
                default:
                    for fi50 := 0; fi50 < 4; fi50++ {
                        v48 -= -2.25
                        _ = fi50
                    }
                }
            }
            _ = v48
        }
        fmt.Println(v4)
        v51 := GenS5{F0: 853}
        _ = v51
    }
    v3 = (42 <= -983)
    v5 = v5
    v52 := GenS2{F0: 0, F1: int8(v1)}
    v5 += (v5 >> ((134) & 63))
    if len(v6) > 0 {
        v6[0] = 0
    }
    if v3 {
        var v53 int16 = 501
        var v54 int64 = v2
        _ = v53
        _ = v54
    }
    fmt.Println(v3)
    v55 := GenS4{F0: v2}
    var v56 int64 = (-(v2))
    v3 = (250 < 2)
    for fi57 := 0; fi57 < 4; fi57++ {
        var v58 int32 = -1
        _ = fi57
        _ = v58
    }
    var v59 bool = ("\n\t" < (("a" + "a") + ""))
    var v60 uint16 = 65535
    var v61 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v55
    _ = v56
    _ = v59
    _ = v6
    _ = v60
    _ = v61
}
