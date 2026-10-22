package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS7 struct { F0 int16; F1 uint32; F2 int32 }
type GenS8 struct { F0 bool }

func main() {
    var v1 uint = 42
    v2 := []string{"hello"}
    v3 := []int{-93, -2147483648, 2147483647}
    var v4 uint64 = 2
    var v5 int8 = -1
    if (!true) {
        var v6 int8 = v5
        v7 := GenS1{F0: "", F1: v4, F2: -9223372036854775808}
        if (-2.25 > 1.5) {
            if (-865 != 1) {
                v1 += (^(v1))
                v2 = append(v2, "foo bar")
                fmt.Println(v6, v4)
            } else {
                v8 := GenS2{F0: v1, F1: 2, F2: (v4 - (v4 * (^(v4))))}
                v1 -= (v1 & (^(v1)))
                v9 := GenS1{F0: (("" + "a") + "foo bar"), F1: v4, F2: 966}
                _ = v8
                _ = v9
            }
            if false {
                v3 = append(v3, 32767)
                v5 = v6
                if false {
                    v3 = append(v3, -9223372036854775808)
                    for fi10 := 0; fi10 < 5; fi10++ {
                        for fi11 := 0; fi11 < 7; fi11++ {
                            var v12 int64 = -217
                            var v13 int32 = 1
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        if true {
                            switch "hello" {
                            case "\n\t":
                                fmt.Println(v4, v5, v4)
                            case "foo bar":
                                for fi14 := 0; fi14 < 6; fi14++ {
                                    if (((fi14 >> ((846) & 63)) > fi10) && false) {
                                        var v15 float32 = 1.5
                                        _ = v15
                                    }
                                    switch (^(fi14)) {
                                    case 42:
                                        v16 := GenS2{F0: v1, F1: 742, F2: (v4 & v4)}
                                        _ = v16
                                    case -1:
                                        fmt.Println(fi14, fi10, v6)
                                    case 295:
                                        var v17 uint = 2
                                        _ = v17
                                    default:
                                        v4 += 18446744073709551615
                                    }
                                    _ = fi14
                                }
                            case "hello":
                                if len(v3) > 0 {
                                    v3[0] = fi10
                                }
                            default:
                                v3 = append(v3, (^(fi10)))
                            }
                            v18 := GenS3{F0: -1}
                            v1 += v1
                            _ = v18
                        } else {
                            switch fi10 {
                            case 0:
                                v19 := GenS4{F0: 443}
                                _ = v19
                            default:
                                v20 := GenS5{F0: "foo bar"}
                                _ = v20
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v21 := GenS3{F0: -1}
                            _ = v21
                        }
                        v4 = v4
                        _ = fi10
                    }
                } else {
                    var v22 uint8 = 255
                    var v23 uint = v1
                    for fi24 := 0; fi24 < 5; fi24++ {
                        for fi25 := 0; fi25 < 6; fi25++ {
                            v26 := GenS3{F0: int16(v22)}
                            v2 = append(v2, "")
                            v1 = v23
                            _ = fi25
                            _ = v26
                        }
                        v2 = append(v2, ("\n\t" + "foo bar"))
                        if false {
                            v27 := GenS6{F0: 42, F1: uint16(v22), F2: (^(v22))}
                            _ = v27
                        } else {
                            v28 := GenS5{F0: ("a" + (("hello" + "hello") + ("\n\t" + "")))}
                            _ = v28
                        }
                        _ = fi24
                    }
                    _ = v22
                    _ = v23
                }
            }
        } else {
            if (!(-810 == -216)) {
                for fi29 := 0; fi29 < 20; fi29++ {
                    var v30 uint8 = 0
                    fmt.Println(v4)
                    _ = fi29
                    _ = v30
                }
                fmt.Println(v1, v4)
                v31 := GenS4{F0: 42}
                _ = v31
            } else {
                v4 = v4
                v6 = (-(v5))
                if len(v2) > 0 {
                    v2[0] = ""
                }
            }
        }
        _ = v6
        _ = v7
    } else {
        v5 = (v5 / ((v5) | 1))
        if false {
            if false {
                var v32 string = ("" + "foo bar")
                v1 += (v1 >> ((v1) & 63))
                fmt.Println(v32, v32, v5)
                _ = v32
            } else {
                if len(v2) > 0 {
                    v2[0] = ""
                }
            }
            for fi33 := 0; fi33 < 2; fi33++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi33
            }
            var v34 int64 = -230
            _ = v34
        } else {
            v5 = 100
            fmt.Println(v1, v5, v1)
        }
        v1 += uint(v4)
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi35 := 0; fi35 < 19; fi35++ {
            var v36 float32 = 0.0
            _ = fi35
            _ = v36
        }
        v5 = (-(v5))
    } else {
        fmt.Println(v5)
        if len(v3) > 0 {
            v3[0] = 972
        }
    }
    v37 := GenS1{F0: "\n\t", F1: 0, F2: 0}
    if true {
        v2 = append(v2, "\n\t")
        v5 -= (^(v5))
    } else {
        var v38 uint = v1
        if true {
            if (0 <= -288) {
                v5 = v5
            }
            v5 = 42
            var v39 int = 127
            _ = v39
        }
        for fi40 := 0; fi40 < 13; fi40++ {
            switch fi40 {
            case -2147483648:
                if (127 <= 127) {
                    v41 := GenS7{F0: 1000, F1: 2147483647, F2: -1}
                    var v42 uint16 = 0
                    _ = v41
                    _ = v42
                } else {
                    v43 := GenS3{F0: 556}
                    var v44 int64 = -613
                    _ = v43
                    _ = v44
                }
            case -991:
                v45 := GenS1{F0: "foo bar", F1: 4294967295, F2: -301}
                _ = v45
            case 32768:
                var v46 int32 = -2147483648
                _ = v46
            default:
                var v47 bool = false
                _ = v47
            }
            v48 := GenS7{F0: 1000, F1: 42, F2: 1}
            for fi49 := 0; fi49 < 15; fi49++ {
                var v50 int16 = -1
                var v51 int8 = v5
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi49
                _ = v50
                _ = v51
            }
            _ = fi40
            _ = v48
        }
        _ = v38
    }
    if true {
        v52 := GenS8{F0: (false || false)}
        _ = v52
    }
    v53 := GenS7{F0: -32768, F1: 2, F2: 1}
    v4 -= (v4 & v4)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v54 bool = (!false)
    var v55 uint16 = 2
    v56 := GenS4{F0: v55}
    var v57 int64 = 958
    v1 = (v1 % ((v1) | 1))
    var v58 bool = (!(0.1 > float64((float32(1.5) * 1e10))))
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
}
