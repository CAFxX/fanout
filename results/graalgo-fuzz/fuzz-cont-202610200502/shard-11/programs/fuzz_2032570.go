package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int8; F2 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 int16; F1 uint8 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 int32, p1 int64, p2 uint16) uint64 {
    v5 := GenS1{F0: 1}
    v6 := GenS2{F0: int16(p0)}
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            if false {
                var v9 bool = true
                switch "" {
                case "hello":
                    switch fi7 {
                    case -812:
                        p2 = p2
                    }
                }
                _ = v9
            }
            p0 = (p0 >> ((18446744073709551615) & 63))
            _ = fi8
        }
        for fi10 := 0; fi10 < 10; fi10++ {
            p0 = 277
            var v11 uint16 = 32767
            var v12 float32 = (3.14159 * (((((((0.0 - (0.1 + (0.0 * (1.5 + 1.5)))) + 3.14159) + 3.14159) * 3.14159) * 1e10) * (1e10 / 1.5)) / 0.1))
            v12 = (v12 / 1.0)
            v13 := []int{}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 856
}

func main() {
    var v1 int64 = -219
    v2 := []string{"\n\t", "a", ("a" + "foo bar")}
    v3 := []uint32{}
    var v4 uint32 = 1
    if true {
        var v14 int = -2147483648
        fmt.Println(v4)
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 uint64 = 1
        v17 := GenS3{F0: int64(v4), F1: 1, F2: (fi15 * (fi15 >> ((0) & 63)))}
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        fmt.Println(fi18, v4)
        _ = fi18
    }
    var v19 int64 = (^(v1))
    switch ((("\n\t" + "a") + "a") + "a") {
    case "\n\t":
        if (float32(1e10) == (0.1 - -2.25)) {
            for fi20 := 0; fi20 < 11; fi20++ {
                for fi21 := 0; fi21 < 2; fi21++ {
                    switch ("foo bar" + ("hello" + "hello")) {
                    case "hello":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        var v22 float32 = 1.5
                        _ = v22
                    }
                    _ = fi21
                }
                fmt.Println(v1, fi20)
                switch (^(fi20)) {
                case 32767:
                    for fi23 := 0; fi23 < 7; fi23++ {
                        var v24 uint32 = v4
                        fmt.Println(v4, fi23)
                        _ = fi23
                        _ = v24
                    }
                case -366:
                    if true {
                        var v25 bool = true
                        if v25 {
                            if len(v3) > 0 {
                                v3[0] = 496
                            }
                            var v26 uint8 = 187
                            if v25 {
                                for fi27 := 0; fi27 < 17; fi27++ {
                                    v28 := GenS3{F0: (v1 + (v1 / ((580) | 1))), F1: 117, F2: fi20}
                                    v26 -= (v26 * v26)
                                    _ = fi27
                                    _ = v28
                                }
                                v29 := GenS4{F0: (v26 >> ((1) & 63))}
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v29
                            }
                            _ = v26
                        } else {
                            switch fi20 {
                            case 0:
                                if (uint64(v4) != genFn0(581, -85, 2)) {
                                    if v25 {
                                        var v30 float64 = 1.5
                                        _ = v30
                                    } else {
                                        v31 := GenS4{F0: 128}
                                        v2 = append(v2, "")
                                        v32 := GenS4{F0: 127}
                                        _ = v31
                                        _ = v32
                                    }
                                }
                            case -2147483648:
                                var v33 float64 = ((0.1 - 0.1) + 3.14159)
                                _ = v33
                            case 4294967295:
                                if v25 {
                                    if (((("\n\t" + ("foo bar" + "a")) <= "a") && true) && (203 != 6)) {
                                        for fi34 := 0; fi34 < 1; fi34++ {
                                            switch fi20 {
                                            case 42:
                                                var v35 int8 = 100
                                                _ = v35
                                            case 32767:
                                                v3 = append(v3, (v4 - (v4 << ((126) & 63))))
                                            }
                                            var v36 int32 = 0
                                            var v37 float64 = -2.25
                                            fmt.Println(v37, v1, fi34)
                                            var v38 uint = 720
                                            fmt.Println(v38, fi20)
                                            _ = fi34
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                        }
                                    } else {
                                        v25 = (-254 > 0)
                                    }
                                } else {
                                    var v39 float32 = 0.0
                                    v25 = (42 != -2147483648)
                                    if (-517 <= -1) {
                                        for fi40 := 0; fi40 < 4; fi40++ {
                                            v2 = append(v2, "a")
                                            fmt.Println(v19)
                                            _ = fi40
                                        }
                                        v4 -= v4
                                        var v41 float64 = (-2.25 * (((3.14159 + -2.25) - 0.1) + (0.1 - float64(0.0))))
                                        _ = v41
                                    } else {
                                        v42 := GenS2{F0: -1}
                                        _ = v42
                                    }
                                    _ = v39
                                }
                            }
                            fmt.Println(fi20, v4, v19)
                            for fi43 := 0; fi43 < 11; fi43++ {
                                if len(v2) > 0 {
                                    v2[0] = "foo bar"
                                }
                                _ = fi43
                            }
                        }
                        var v44 int = 128
                        _ = v25
                        _ = v44
                    }
                case 42:
                    if (true && ((!true) && (256 >= 65535))) {
                        var v45 string = ("\n\t" + "\n\t")
                        _ = v45
                    } else {
                        var v46 uint16 = 2
                        v2 = append(v2, "\n\t")
                        for fi47 := 0; fi47 < 1; fi47++ {
                            v48 := GenS3{F0: v1, F1: 100, F2: int(v1)}
                            v49 := GenS4{F0: 89}
                            _ = fi47
                            _ = v48
                            _ = v49
                        }
                        _ = v46
                    }
                }
                _ = fi20
            }
        }
    case "":
        if false {
            var v50 float64 = 3.14159
            _ = v50
        } else {
            var v51 string = ("\n\t" + "")
            for fi52 := 0; fi52 < 8; fi52++ {
                if len(v2) > 0 {
                    v2[0] = "a"
                }
                switch v51 {
                case "\n\t":
                    for fi53 := 0; fi53 < 4; fi53++ {
                        v54 := GenS5{F0: 0, F1: 531}
                        v4 -= v4
                        switch 727 {
                        case -2147483648:
                            switch -237 {
                            case -2147483648:
                                v2 = append(v2, (v51 + ((((v51 + (("" + "a") + ("\n\t" + (((((v51 + "hello") + v51) + "\n\t") + v51) + ("" + (v51 + ((v51 + v51) + v51))))))) + "") + "hello") + ("foo bar" + (v51 + "foo bar")))))
                            default:
                                var v55 uint16 = 65535
                                _ = v55
                            }
                        case 4294967295:
                            var v56 uint16 = 2
                            _ = v56
                        }
                        v57 := GenS6{F0: -415, F1: 25}
                        _ = fi53
                        _ = v54
                        _ = v57
                    }
                case "foo bar":
                    switch v51 {
                    case "\n\t":
                        v4 = (^(v4))
                    default:
                        for fi58 := 0; fi58 < 7; fi58++ {
                            v1 += -1
                            _ = fi58
                        }
                    }
                default:
                    v59 := GenS7{F0: 239}
                    _ = v59
                }
                _ = fi52
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v51
        }
    case "a":
        var v60 uint64 = 169
        _ = v60
    default:
        _ = len(v3)
    }
    var v61 int16 = -32768
    var v62 uint8 = 128
    fmt.Println(v62, v19)
    var v63 int16 = int16(v4)
    var v64 int64 = v1
    if true {
        var v65 int8 = -128
        _ = v65
    }
    switch -874 {
    case -612:
        v66 := GenS1{F0: (v62 % (((^(v62))) | 1))}
        _ = v66
    case -681:
        for fi67 := 0; fi67 < 1; fi67++ {
            fmt.Println(v61)
            fmt.Println(v4)
            var v68 int16 = 0
            _ = fi67
            _ = v68
        }
    default:
        v19 -= -9223372036854775808
    }
    var v69 bool = true
    var v70 float32 = 0.0
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v69
    _ = v70
}
