package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint32; F2 string }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int16; F1 uint32; F2 bool }

func main() {
    var v1 uint = 65535
    var v2 float64 = -2.25
    v3 := []uint32{435}
    var v4 int32 = 1
    var v5 uint = v1
    var v6 int64 = 1
    for fi7 := 0; fi7 < 19; fi7++ {
        if (!((fi7 & int(v5)) < (fi7 + -2147483648))) {
            var v8 uint16 = 2
            _ = v8
        }
        v5 += 2
        v6 = (v6 | v6)
        _ = fi7
    }
    fmt.Println(v2)
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 int = fi9
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS1{F0: (true && true)}
        if (-128 == 0) {
            var v13 float64 = v2
            if len(v3) > 0 {
                v3[0] = 783
            }
            var v14 int32 = v4
            _ = v13
            _ = v14
        } else {
            var v15 int = (fi11 / ((-29) | 1))
            _ = v15
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v5, v5, v6)
    if (1e10 == v2) {
        v16 := GenS1{F0: (32767 <= 82)}
        switch "foo bar" {
        case "a":
            if (32767 > 2) {
                for fi17 := 0; fi17 < 17; fi17++ {
                    var v18 string = "\n\t"
                    var v19 int32 = v4
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                for fi20 := 0; fi20 < 13; fi20++ {
                    var v21 uint8 = 1
                    fmt.Println(v4, v1)
                    var v22 int16 = -312
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                var v23 uint16 = uint16(v5)
                _ = v23
            } else {
                var v24 float32 = ((((-2.25 / (3.14159 / 1.0)) / 1.5) + 0.0) / 1.5)
                v25 := GenS2{F0: 1, F1: 659, F2: "foo bar"}
                var v26 float32 = 3.14159
                _ = v24
                _ = v25
                _ = v26
            }
        default:
            for fi27 := 0; fi27 < 17; fi27++ {
                v2 -= v2
                _ = fi27
            }
        }
        _ = v16
    }
    fmt.Println(v1, v2, v5)
    _ = len(v3)
    _ = len(v3)
    for fi28 := 0; fi28 < 20; fi28++ {
        var v29 int16 = -32768
        _ = len(v3)
        var v30 int16 = (v29 - v29)
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        for fi32 := 0; fi32 < 19; fi32++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v33 int64 = v6
            var v34 int8 = -79
            v1 = v5
            if true {
                v1 += (v1 / (((v5 % (((v5 ^ (v1 & 4294967295))) | 1))) | 1))
                v35 := GenS3{F0: 24, F1: 255}
                if false {
                    v5 -= (v5 << ((42) & 63))
                }
                _ = v35
            } else {
                var v36 uint = (v1 | (v1 - v5))
                var v37 float32 = -2.25
                _ = v36
                _ = v37
            }
            if true {
                fmt.Println(v5, v2, v1)
            }
            _ = fi32
            _ = v33
            _ = v34
        }
        switch (fi31 << (((v1 * 2)) & 63)) {
        case -940:
            fmt.Println(v5, v6, v4)
        case 4294967295:
            if (true || true) {
                v4 = v4
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        default:
            if (false || false) {
                var v38 float64 = v2
                for fi39 := 0; fi39 < 13; fi39++ {
                    var v40 string = ("foo bar" + "foo bar")
                    v2 = v38
                    if true {
                        if true {
                            v6 = v6
                            var v41 uint8 = 1
                            _ = v41
                        }
                        v6 = v6
                        if (false || (!(!(!(false || true))))) {
                            var v42 string = (v40 + "")
                            if (250 == 0) {
                                if true {
                                    v42 = v40
                                    v6 += (v6 * (v6 & v6))
                                }
                                v40 = v42
                                v43 := GenS4{F0: (v6 % ((v6) | 1)), F1: 2}
                                v38 = float64(0.1)
                                _ = v43
                            }
                            _ = v42
                        }
                        if false {
                            if true {
                                v3 = append(v3, 214)
                            }
                            if false {
                                v40 = (("foo bar" + "") + (v40 + v40))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    } else {
                        v44 := GenS4{F0: (v6 | (v6 << (((^(v1))) & 63))), F1: 65535}
                        v45 := GenS2{F0: int8(v6), F1: 497, F2: (v40 + v40)}
                        if (true && true) {
                            if (!((v6 ^ int64(v4)) <= v6)) {
                                v6 = (v6 << ((2) & 63))
                            } else {
                                switch ((v40 + "hello") + v40) {
                                case "a":
                                    v1 += (v5 % ((75) | 1))
                                case "hello":
                                    v46 := GenS3{F0: 128, F1: 42}
                                    _ = v46
                                }
                                v3 = append(v3, 2147483647)
                            }
                            if (false || true) {
                                v4 = -181
                                var v47 float32 = 0.0
                                _ = v47
                            } else {
                                var v48 int16 = -524
                                v48 += (-(v48))
                                _ = v48
                            }
                            if true {
                                v4 -= (v4 & v4)
                                v49 := GenS1{F0: true}
                                _ = v49
                            } else {
                                v1 += 290
                                var v50 int16 = 0
                                v51 := GenS5{F0: (int32(v5) < v4)}
                                if ((v1 - (v5 ^ (v5 ^ (v5 * v1)))) > (v5 | v1)) {
                                    var v52 bool = false
                                    var v53 string = v40
                                    v40 = v53
                                    _ = v52
                                    _ = v53
                                } else {
                                    v6 += v6
                                    switch (fi39 & fi39) {
                                    case 127:
                                        var v54 int32 = (-(v4))
                                        _ = v54
                                    default:
                                        fmt.Println(v1, v6, fi31)
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                switch v40 {
                                case "":
                                    _ = len(v3)
                                }
                                _ = v50
                                _ = v51
                            }
                            v5 += 42
                        }
                        _ = v44
                        _ = v45
                    }
                    _ = fi39
                    _ = v40
                }
                _ = v38
            }
        }
        for fi55 := 0; fi55 < 15; fi55++ {
            var v56 int64 = 0
            v57 := GenS6{F0: -32768}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v58 int8 = -1
            _ = fi55
            _ = v56
            _ = v57
            _ = v58
        }
        _ = fi31
    }
    v59 := GenS7{F0: 693, F1: 380, F2: (1 >= v6)}
    v60 := GenS4{F0: v6, F1: 48}
    fmt.Println(v4)
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
}
