package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 string; F2 uint8 }
type GenS4 struct { F0 int; F1 int32; F2 int32 }
func genFn0(p0 bool, p1 uint64) int32 {
    p1 = (p1 * (p1 + p1))
    var v6 uint32 = 111
    if (42 <= -518) {
        for fi7 := 0; fi7 < 5; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 int16 = 32767
                var v10 uint = 65535
                var v11 uint8 = 0
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            var v12 int8 = -128
            _ = fi7
            _ = v12
        }
        v6 = (v6 << ((638) & 63))
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := []int{4294967295}
        if (!((0.0 * 3.14159) > (1.5 - 1.5))) {
            for fi15 := 0; fi15 < 14; fi15++ {
                switch fi15 {
                case 157:
                    var v16 bool = false
                    _ = v16
                case -758:
                    if (66 < -42) {
                        if len(v14) > 0 {
                            v14[0] = (fi13 % (((^(fi15))) | 1))
                        }
                        switch (fi15 - (fi13 & fi15)) {
                        case 0:
                            var v17 uint16 = 40
                            _ = v17
                        case -1:
                            var v18 float32 = 0.1
                            fmt.Println(fi15, p1)
                            fmt.Println(p1)
                            _ = v18
                        case 9223372036854775807:
                            if len(v14) > 0 {
                                v14[0] = 9223372036854775807
                            }
                        default:
                            v19 := GenS1{F0: 255, F1: (^(fi15)), F2: ""}
                            _ = v19
                        }
                        var v20 float32 = 0.1
                        _ = v20
                    } else {
                        v14 = append(v14, 0)
                        if (4294967295 > (fi13 % ((42) | 1))) {
                            var v21 uint = 1
                            var v22 uint16 = 0
                            _ = v21
                            _ = v22
                        } else {
                            v23 := GenS2{F0: (^(p1))}
                            _ = v23
                        }
                    }
                }
                _ = len(v14)
                var v24 int16 = 32767
                if (true && p0) {
                    var v25 uint16 = 0
                    _ = v25
                }
                _ = fi15
                _ = v24
            }
            var v26 uint64 = (p1 << ((65535) & 63))
            _ = v26
        }
        var v27 string = ("\n\t" + "")
        _ = fi13
        _ = v14
        _ = v27
    }
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    var v1 uint = 42
    var v2 float64 = (-2.25 / 1.0)
    var v3 uint32 = 976
    v4 := []uint32{v3}
    var v5 int16 = 0
    switch "\n\t" {
    case "\n\t":
        var v28 int = -578
        _ = v28
    }
    if ((true && (false || false)) && false) {
        v3 -= (v3 - v3)
        v3 = v3
    } else {
        var v29 int32 = -2147483648
        for fi30 := 0; fi30 < 17; fi30++ {
            for fi31 := 0; fi31 < 17; fi31++ {
                v29 += v29
                var v32 int32 = 42
                _ = fi31
                _ = v32
            }
            if (-1 > -96) {
                if false {
                    fmt.Println(fi30, v1, v5)
                } else {
                    v2 = (float64(1.5) / (v2 / v2))
                }
            }
            fmt.Println(v29)
            _ = fi30
        }
        _ = v29
    }
    var v33 uint = v1
    v34 := GenS2{F0: 2}
    var v35 int32 = -2147483648
    v3 = 283
    for fi36 := 0; fi36 < 13; fi36++ {
        if (!(!(!(135 < 2)))) {
            for fi37 := 0; fi37 < 8; fi37++ {
                if true {
                    v3 += v3
                } else {
                    if len(v4) > 0 {
                        v4[0] = (^(v3))
                    }
                }
                v38 := GenS3{F0: (fi36 & 128), F1: (("a" + "\n\t") + "\n\t"), F2: 127}
                _ = fi37
                _ = v38
            }
            if len(v4) > 0 {
                v4[0] = v3
            }
        } else {
            v33 += (v1 & (v1 / ((v33) | 1)))
            switch int(v35) {
            case 0:
                var v39 uint64 = 2
                _ = v39
            case 2147483647:
                _ = len(v4)
            case -2147483648:
                for fi40 := 0; fi40 < 5; fi40++ {
                    switch fi40 {
                    case 247:
                        v35 = (v35 + (v35 ^ (v35 - (^(v35)))))
                    case 127:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    default:
                        if true {
                            fmt.Println(v3, v33)
                            v33 = v1
                            v41 := GenS4{F0: 42, F1: (-(v35)), F2: genFn0(false, 568)}
                            _ = v41
                        }
                        v1 = 1
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    if false {
                        v42 := GenS3{F0: -1, F1: "hello", F2: 0}
                        switch ((("a" + (("\n\t" + "a") + ("hello" + (("" + "\n\t") + "hello")))) + "a") + "foo bar") {
                        case "foo bar":
                            fmt.Println(v35)
                        case "hello":
                            v43 := GenS3{F0: -493, F1: ("" + (("a" + "\n\t") + "foo bar")), F2: 2}
                            _ = v43
                        case "\n\t":
                            var v44 float64 = -2.25
                            _ = v44
                        default:
                            var v45 uint16 = 2
                            _ = v45
                        }
                        _ = len(v4)
                        _ = v42
                    }
                    v46 := GenS3{F0: (-(fi40)), F1: "\n\t", F2: 165}
                    _ = fi40
                    _ = v46
                }
            default:
                var v47 uint64 = 42
                _ = v47
            }
            var v48 bool = false
            _ = v48
        }
        var v49 float32 = 1.5
        v50 := GenS3{F0: fi36, F1: ("\n\t" + ""), F2: 59}
        _ = fi36
        _ = v49
        _ = v50
    }
    for fi51 := 0; fi51 < 6; fi51++ {
        v52 := GenS3{F0: -391, F1: "a", F2: 255}
        var v53 int16 = v5
        _ = fi51
        _ = v52
        _ = v53
    }
    var v54 string = ""
    var v55 int64 = -9223372036854775808
    _ = len(v4)
    fmt.Println(v35)
    if (!true) {
        var v56 int16 = -500
        v3 = v3
        switch "a" {
        case "":
            var v57 int64 = -9223372036854775808
            _ = v57
        default:
            v5 -= v56
        }
        _ = v56
    }
    v4 = append(v4, v3)
    v58 := GenS1{F0: 135, F1: -383, F2: "\n\t"}
    switch "foo bar" {
    case "hello":
        if false {
            v4 = append(v4, 2147483648)
            var v59 uint16 = 32767
            if len(v4) > 0 {
                v4[0] = v3
            }
            _ = v59
        }
    }
    var v60 uint = v33
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v58
    _ = v60
}
