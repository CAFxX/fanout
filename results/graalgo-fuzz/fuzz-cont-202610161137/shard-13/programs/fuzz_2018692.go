package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 int64; F1 int32; F2 string }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint; F1 uint8 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint8; F1 int32 }
func genFn0() uint32 {
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := []int{fi5, fi5, (fi5 ^ fi5), 32768}
        fmt.Println(fi5)
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 int32 = 0
            fmt.Println(v8, fi5)
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    if (42 < -1) {
        v9 := []int{0, 0}
        _ = v9
    }
    return 65535
}

func main() {
    var v1 int8 = -46
    v2 := []int{9223372036854775807, 42}
    var v3 uint8 = 210
    v4 := make(map[uint32]bool)
    for fi10 := 0; fi10 < 11; fi10++ {
        switch ("a" + "") {
        case "\n\t":
            v1 += (v1 & v1)
        case "":
            v1 = 0
        case "foo bar":
            if true {
                v3 += 255
                fmt.Println(v3, v3, fi10)
                for fi11 := 0; fi11 < 10; fi11++ {
                    fmt.Println(fi11)
                    _ = fi11
                }
            }
        default:
            switch (fi10 % ((fi10) | 1)) {
            case -437:
                if true {
                    v3 = v3
                    switch fi10 {
                    case 32768:
                        if (false && (false || true)) {
                            var v12 uint16 = 341
                            var v13 string = "hello"
                            _ = v12
                            _ = v13
                        }
                    case 42:
                        v1 = v1
                    }
                } else {
                    v1 = (v1 >> ((18446744073709551615) & 63))
                }
            case 127:
                _ = len(v2)
            default:
                v3 = (v3 * 0)
            }
        }
        switch ("a" + "hello") {
        case "":
            if (2 < 255) {
                _ = len(v4)
                v14 := GenS1{F0: 1}
                for fi15 := 0; fi15 < 11; fi15++ {
                    v2 = append(v2, 1)
                    v16 := v4[0]
                    v1 = (-(v1))
                    _ = fi15
                    _ = v16
                }
                _ = v14
            } else {
                var v17 uint = uint(v3)
                v18 := GenS2{F0: v1}
                for fi19 := 0; fi19 < 18; fi19++ {
                    v17 = (v17 - (v17 << (((v17 >> (((v17 / (((v17 ^ v17)) | 1))) & 63))) & 63)))
                    _ = fi19
                }
                _ = v17
                _ = v18
            }
        case "\n\t":
            for fi20 := 0; fi20 < 5; fi20++ {
                v3 -= (^(v3))
                _ = fi20
            }
        }
        _ = fi10
    }
    if ((float64(0.1) / (1e10 - 0.1)) < 0.1) {
        for fi21 := 0; fi21 < 5; fi21++ {
            switch "\n\t" {
            case "hello":
                v3 -= v3
            default:
                switch fi21 {
                case 467:
                    v3 += (v3 + 1)
                case -634:
                    v3 += 255
                case 9223372036854775807:
                    v3 = (v3 & (v3 * v3))
                default:
                    var v22 int8 = 4
                    _ = v22
                }
            }
            _ = fi21
        }
        v23 := GenS3{F0: true, F1: genFn0(), F2: 1}
        var v24 uint8 = v3
        _ = v23
        _ = v24
    }
    v1 = 1
    for fi25 := 0; fi25 < 5; fi25++ {
        var v26 int8 = (v1 >> ((440) & 63))
        for fi27 := 0; fi27 < 15; fi27++ {
            v1 = (v26 * (-(v26)))
            v28 := GenS4{F0: v3, F1: 0}
            var v29 float64 = 3.14159
            if (false || false) {
                v30 := GenS5{F0: 255, F1: 417}
                v1 -= 127
                _ = v30
            } else {
                v29 = ((v29 * v29) / (0.1 - v29))
                v31 := GenS6{F0: -813, F1: -284, F2: "a"}
                _ = v31
            }
            switch fi25 {
            case 2147483647:
                switch ("foo bar" + "foo bar") {
                case "a":
                    v32 := GenS7{F0: 1}
                    _ = v32
                default:
                    var v33 int8 = -1
                    _ = v33
                }
            case 947:
                v26 += v26
            case 128:
                v34 := GenS8{F0: v26}
                _ = v34
            default:
                var v35 uint64 = 0
                _ = v35
            }
            if true {
                v29 = v29
                v36 := GenS9{F0: 42, F1: (v3 + 42)}
                fmt.Println(v29, fi25, v3)
                _ = v36
            } else {
                v37 := v4[2147483647]
                _ = v37
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi25
        _ = v26
    }
    var v38 int64 = -1
    switch 4294967295 {
    case -304:
        var v39 uint16 = 0
        _ = v39
    case 32767:
        v40 := GenS8{F0: 100}
        _ = v40
    case 9223372036854775807:
        v2 = append(v2, 252)
    }
    fmt.Println(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 -= (^(v3))
    v3 -= (v3 - (v3 >> ((65535) & 63)))
    if (((1e10 - 1.5) / 1.0) <= 1e10) {
        if true {
            v1 = v1
            v41 := GenS10{F0: 181}
            _ = v41
        }
        for fi42 := 0; fi42 < 3; fi42++ {
            var v43 int16 = 243
            var v44 string = "\n\t"
            _ = fi42
            _ = v43
            _ = v44
        }
        v45 := GenS11{F0: 42}
        _ = v45
    } else {
        for fi46 := 0; fi46 < 3; fi46++ {
            var v47 string = "foo bar"
            var v48 float32 = -2.25
            var v49 float32 = float32(float64(v48))
            _ = fi46
            _ = v47
            _ = v48
            _ = v49
        }
        v50 := GenS12{F0: (v38 * 584)}
        _ = v50
    }
    for fi51 := 0; fi51 < 3; fi51++ {
        for fi52 := 0; fi52 < 19; fi52++ {
            v1 = v1
            v38 -= -1
            _ = fi52
        }
        for fi53 := 0; fi53 < 5; fi53++ {
            switch "foo bar" {
            case "foo bar":
                switch ("\n\t" + ("\n\t" + "")) {
                case "foo bar":
                    var v54 bool = (1e10 != 3.14159)
                    _ = v54
                case "":
                    v55 := v4[65535]
                    _ = v55
                case "a":
                    if false {
                        var v56 float32 = 0.1
                        if false {
                            v38 = (v38 ^ 42)
                            v57 := GenS9{F0: 4294967295, F1: (v3 << ((22) & 63))}
                            v3 = (v3 - (v3 >> ((155) & 63)))
                            _ = v57
                        } else {
                            var v58 int = (fi53 ^ 128)
                            v59 := GenS9{F0: 36, F1: v3}
                            _ = v58
                            _ = v59
                        }
                        var v60 int32 = int32(v1)
                        _ = v56
                        _ = v60
                    }
                    var v61 int16 = int16(v38)
                    switch fi51 {
                    case -9223372036854775808:
                        fmt.Println(fi51)
                        v38 -= (v38 + v38)
                        v2 = append(v2, (fi53 / ((fi53) | 1)))
                    }
                    _ = v61
                default:
                    v3 += v3
                }
            default:
                v62 := GenS13{F0: 127, F1: 1}
                switch fi51 {
                case 32768:
                    v63 := GenS6{F0: (^(v38)), F1: 42, F2: ("foo bar" + "\n\t")}
                    _ = v63
                case 42:
                    fmt.Println(fi51, v3)
                default:
                    var v64 int32 = 1
                    _ = v64
                }
                v38 += v38
                _ = v62
            }
            _ = fi53
        }
        for fi65 := 0; fi65 < 9; fi65++ {
            fmt.Println(v3, fi51)
            var v66 int16 = -87
            v3 = (v3 / ((v3) | 1))
            var v67 float64 = 1.5
            _ = fi65
            _ = v66
            _ = v67
        }
        _ = fi51
    }
    fmt.Println(v3)
    if (-32768 != 42) {
        v1 = v1
        var v68 float64 = 1.5
        _ = v68
    }
    fmt.Println(v1, v38)
    var v69 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v69
}
