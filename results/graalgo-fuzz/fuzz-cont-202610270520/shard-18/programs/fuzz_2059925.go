package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint32) int16 {
    switch 42 {
    case 9223372036854775807:
        switch 42 {
        case 2147483647:
            p0 += p0
        case 120:
            var v7 int = -319
            _ = v7
        default:
            p0 += p0
        }
    case 368:
        var v8 uint8 = 0
        _ = v8
    }
    v9 := GenS1{F0: 1}
    for fi10 := 0; fi10 < 4; fi10++ {
        p0 += (p0 / (((p0 / ((p0) | 1))) | 1))
        for fi11 := 0; fi11 < 1; fi11++ {
            fmt.Println(fi10)
            _ = fi11
        }
        var v12 uint32 = (p0 ^ p0)
        _ = fi10
        _ = v12
    }
    v13 := []int{32767, 32768, 32768}
    _ = p0
    _ = v13
    _ = v9
    return -254
}

func main() {
    var v1 int32 = 42
    var v2 uint8 = 42
    v3 := make(map[string]bool)
    v4 := []uint32{}
    v5 := make(map[string]int)
    var v6 int64 = -9223372036854775808
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 uint8 = (v2 | (v2 ^ (^(v2))))
        v15 -= (^(v15))
        switch ("foo bar" + "foo bar") {
        case "hello":
            if (938 < 42) {
                delete(v3, ("foo bar" + ("" + (("\n\t" + ("foo bar" + "\n\t")) + ("hello" + "foo bar")))))
                if (4294967295 <= 65535) {
                    v16 := GenS2{F0: 2}
                    var v17 uint8 = 103
                    _ = v16
                    _ = v17
                }
            }
        case "a":
            v15 -= (^(v15))
        default:
            v2 = (v15 / ((153) | 1))
        }
        _ = fi14
        _ = v15
    }
    if ((v2 + v2) != v2) {
        if true {
            delete(v3, ("" + "hello"))
            switch 42 {
            case 32768:
                for fi18 := 0; fi18 < 7; fi18++ {
                    var v19 string = "foo bar"
                    for fi20 := 0; fi20 < 12; fi20++ {
                        var v21 uint8 = 255
                        v19 = ("a" + (v19 + ""))
                        v22 := GenS3{F0: genFn0(0), F1: genFn0(1)}
                        v23 := GenS3{F0: -32768, F1: int16(v1)}
                        v21 = (v2 << ((273) & 63))
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 15; fi24++ {
                        var v25 uint16 = 42
                        switch "foo bar" {
                        case "a":
                            fmt.Println(v6, v1, fi24)
                        case "hello":
                            v25 += (v25 >> ((65535) & 63))
                        case "foo bar":
                            var v26 int = (fi24 & fi24)
                            _ = v26
                        default:
                            var v27 uint16 = (v25 & v25)
                            _ = v27
                        }
                        if false {
                            if (835 == (fi24 >> ((1) & 63))) {
                                var v28 float64 = 3.14159
                                if true {
                                    v29 := GenS4{F0: v19}
                                    v30 := GenS2{F0: 65535}
                                    v25 -= (v25 % ((255) | 1))
                                    _ = v29
                                    _ = v30
                                }
                                _ = v28
                            }
                        } else {
                            var v31 int8 = -128
                            if len(v4) > 0 {
                                v4[0] = 2147483648
                            }
                            v32 := GenS4{F0: v19}
                            switch v19 {
                            case "hello":
                                v33 := GenS3{F0: int16(v25), F1: -32768}
                                _ = v33
                            default:
                                var v34 uint = 65535
                                _ = v34
                            }
                            var v35 float64 = 1.5
                            _ = v31
                            _ = v32
                            _ = v35
                        }
                        _ = fi24
                        _ = v25
                    }
                    _ = fi18
                    _ = v19
                }
            default:
                v36 := GenS4{F0: (((("\n\t" + "foo bar") + "\n\t") + "hello") + "\n\t")}
                _ = v36
            }
            v37 := GenS4{F0: ("a" + (("a" + (("\n\t" + "a") + "foo bar")) + "a"))}
            _ = v37
        } else {
            fmt.Println(v1, v6)
            if (-9223372036854775808 >= (v6 & (v6 ^ (v6 << ((435) & 63))))) {
                if true {
                    v38 := v5["foo bar"]
                    switch (v38 & (v38 | (v38 - 9223372036854775807))) {
                    case -530:
                        v1 -= (-(v1))
                    }
                    v3["foo bar"] = false
                    _ = v38
                } else {
                    v2 = 128
                    v39 := GenS1{F0: 255}
                    _ = v39
                }
            } else {
                var v40 uint32 = 4294967295
                _ = v40
            }
        }
    } else {
        if (!true) {
            var v41 uint8 = (v2 >> ((813) & 63))
            if (!false) {
                var v42 uint64 = 486
                v6 = (v6 - v6)
                v1 = (-(v1))
                _ = v42
            } else {
                var v43 int32 = (-(v1))
                var v44 int16 = -32768
                _ = v43
                _ = v44
            }
            _ = v41
        } else {
            v2 += v2
            var v45 float32 = 0.0
            _ = v45
        }
        var v46 float32 = 3.14159
        _ = v46
    }
    v47 := GenS1{F0: 2}
    for fi48 := 0; fi48 < 10; fi48++ {
        for fi49 := 0; fi49 < 11; fi49++ {
            v50 := GenS5{F0: (2 <= 998)}
            fmt.Println(v2, v6, fi49)
            var v51 string = "foo bar"
            var v52 int16 = -32768
            var v53 uint8 = (v2 % ((v2) | 1))
            _ = fi49
            _ = v50
            _ = v51
            _ = v52
            _ = v53
        }
        _ = fi48
    }
    var v54 float32 = 0.1
    var v55 uint64 = 42
    var v56 uint64 = v55
    v2 = (v2 | 255)
    var v57 uint8 = 127
    switch int(v6) {
    case 9223372036854775807:
        if (false || (1000 < genFn0(1))) {
            v57 = (v57 / ((uint8(v56)) | 1))
        } else {
            switch 32767 {
            case 1:
                v58 := v5["hello"]
                _ = v58
            default:
                switch "hello" {
                case "a":
                    v56 = 18446744073709551615
                }
            }
            delete(v5, "\n\t")
            if ((v57 % ((v57) | 1)) >= (v57 | v57)) {
                for fi59 := 0; fi59 < 16; fi59++ {
                    var v60 uint16 = 300
                    var v61 int16 = 427
                    _ = fi59
                    _ = v60
                    _ = v61
                }
                v54 += 1e10
                v56 -= v56
            } else {
                for fi62 := 0; fi62 < 10; fi62++ {
                    for fi63 := 0; fi63 < 18; fi63++ {
                        var v64 bool = ("" != (((("a" + ("hello" + "a")) + "hello") + "hello") + ""))
                        if (137 > (v57 / ((129) | 1))) {
                            fmt.Println(v57)
                        } else {
                            v65 := GenS3{F0: genFn0(2147483648), F1: 1000}
                            fmt.Println(v55, v6, fi62)
                            if v64 {
                                var v66 int = fi63
                                v4 = append(v4, 42)
                                _ = v66
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v54 = (1.5 - v54)
                            }
                            var v67 uint64 = v55
                            v4 = append(v4, 4294967295)
                            var v68 uint16 = 255
                            _ = v65
                            _ = v67
                            _ = v68
                        }
                        v69 := GenS5{F0: (!v64)}
                        _ = fi63
                        _ = v64
                        _ = v69
                    }
                    _ = fi62
                }
                v57 += v57
                switch ("foo bar" + "a") {
                case "hello":
                    switch 2147483647 {
                    case 693:
                        var v70 int64 = (v6 & 1)
                        _ = v70
                    case -9223372036854775808:
                        _ = len(v3)
                    case -683:
                        delete(v5, "foo bar")
                    default:
                        v2 -= (v57 / (((v57 & (v2 - (v57 ^ 132)))) | 1))
                    }
                case "a":
                    v1 = -609
                }
            }
        }
    case 128:
        v71 := GenS6{F0: 0}
        _ = v71
    }
    v56 -= v56
    var v72 uint8 = 1
    var v73 int8 = 1
    for fi74 := 0; fi74 < 3; fi74++ {
        var v75 int = fi74
        _ = fi74
        _ = v75
    }
    var v76 uint32 = 2
    v1 = -511
    v77 := GenS6{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v6
    _ = v72
    _ = v73
    _ = v76
    _ = v77
}
