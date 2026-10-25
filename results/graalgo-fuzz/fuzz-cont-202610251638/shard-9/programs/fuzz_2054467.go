package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 uint64; F2 int }
type GenS7 struct { F0 int32; F1 bool }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int32; F1 string; F2 uint64 }
func genFn0() uint8 {
    v5 := GenS1{F0: 1, F1: 4294967295, F2: 18}
    var v6 uint64 = 4294967295
    v6 = v6
    v7 := []int{-1, 2147483647, 2147483647, -1}
    _ = v5
    _ = v6
    _ = v7
    return 127
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 1000
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 32767
    v8 := v1["\n\t"]
    if true {
        var v9 uint = 167
        if true {
            fmt.Println(v9, v2)
            v10 := GenS2{F0: 116}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 16; fi11++ {
                var v12 int = 2147483647
                var v13 string = "\n\t"
                _ = fi11
                _ = v12
                _ = v13
            }
            if true {
                for fi14 := 0; fi14 < 15; fi14++ {
                    v8 = ((("" + v8) + "foo bar") + v8)
                    v8 = (v8 + v8)
                    var v15 uint8 = 127
                    _ = fi14
                    _ = v15
                }
                if (v4 > (^(v4))) {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        v17 := []int{}
                        v18 := GenS2{F0: (fi16 * (fi16 & -12))}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    v19 := GenS1{F0: (v9 >> ((v9) & 63)), F1: (v9 * v9), F2: 1}
                    v20 := GenS3{F0: 2, F1: uint16(v9), F2: 255}
                    _ = v19
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 15; fi21++ {
                        v8 = v8
                        v9 -= (v9 & uint(v4))
                        _ = fi21
                    }
                }
            }
            var v22 int16 = v2
            _ = v22
        }
        var v23 uint64 = 1
        _ = v23
        _ = v9
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        fmt.Println(v8)
        fmt.Println(v8, v8, v2)
        _ = fi24
    }
    var v25 bool = (-1 > 558)
    if true {
        var v26 int16 = -960
        _ = v26
    } else {
        var v27 uint16 = (v4 * v4)
        v28 := GenS3{F0: 6, F1: 32767, F2: 128}
        if (uint32(v4) > 2147483647) {
            if ((1.5 - 3.14159) > 0.1) {
                v29 := GenS1{F0: 458, F1: 320, F2: -833}
                _ = v29
            }
            switch -144 {
            case -9223372036854775808:
                var v30 int8 = 42
                _ = v30
            case 128:
                v31 := GenS4{F0: int16(v27)}
                _ = v31
            case 606:
                for fi32 := 0; fi32 < 18; fi32++ {
                    if (true || (!v25)) {
                        v33 := []int{1, fi32}
                        _ = v33
                    } else {
                        v34 := []int{(fi32 & fi32), 32767}
                        switch (fi32 | 649) {
                        case -1:
                            var v35 int16 = -105
                            _ = v35
                        default:
                            var v36 uint = 0
                            _ = v36
                        }
                        _ = v34
                    }
                    var v37 bool = true
                    _ = fi32
                    _ = v37
                }
            default:
                var v38 float64 = (3.14159 / (float64((1e10 - -2.25)) * ((1.5 + 3.14159) - -2.25)))
                _ = v38
            }
            switch 756 {
            case -1:
                v25 = (("a" != ("hello" + "foo bar")) && (v25 || true))
            }
        } else {
            v25 = v25
            switch 32768 {
            case 1:
                for fi39 := 0; fi39 < 11; fi39++ {
                    if ((v4 / ((v4) | 1)) < 255) {
                        fmt.Println(v4, v2)
                    } else {
                        v2 -= 552
                        v40 := []int{fi39, fi39}
                        _ = v40
                    }
                    var v41 uint8 = 127
                    _ = fi39
                    _ = v41
                }
            case -273:
                if true {
                    v1[v8] = "\n\t"
                } else {
                    var v42 uint8 = 161
                    fmt.Println(v8)
                    _ = v42
                }
            case 127:
                v4 -= v4
            }
        }
        _ = v27
        _ = v28
    }
    switch 32768 {
    case -678:
        for fi43 := 0; fi43 < 3; fi43++ {
            var v44 string = ("foo bar" + v8)
            v8 = (v44 + "\n\t")
            _ = fi43
            _ = v44
        }
    case 842:
        var v45 bool = v25
        _ = v45
    case 9223372036854775807:
        v46 := GenS4{F0: -1}
        _ = v46
    default:
        v47 := GenS5{F0: 0}
        _ = v47
    }
    var v48 uint32 = 2147483647
    switch "" {
    case "a":
        var v49 string = (v8 + ("a" + ((v8 + v8) + v8)))
        _ = v49
    }
    switch -53 {
    case -878:
        v25 = v25
    case 1:
        switch "\n\t" {
        case "hello":
            fmt.Println(v8)
        default:
            var v50 uint32 = v48
            _ = v50
        }
    case -741:
        for fi51 := 0; fi51 < 9; fi51++ {
            if v25 {
                switch -720 {
                case -1:
                    var v52 float32 = 1e10
                    _ = v52
                default:
                    var v53 float64 = 0.1
                    _ = v53
                }
                v54 := GenS5{F0: (v2 << ((416) & 63))}
                var v55 uint32 = (v48 & (v48 / ((v48) | 1)))
                _ = v54
                _ = v55
            } else {
                fmt.Println(v25, v4, v2)
                if (155 != 42) {
                    if true {
                        for fi56 := 0; fi56 < 16; fi56++ {
                            switch 4294967295 {
                            case 127:
                                v48 += v48
                            }
                            _ = fi56
                        }
                        fmt.Println(v2, v8)
                        var v57 int16 = (-(v2))
                        _ = v57
                    } else {
                        v58 := GenS6{F0: v2, F1: 430, F2: (-(fi51))}
                        v2 = 356
                        _ = v58
                    }
                } else {
                    for fi59 := 0; fi59 < 18; fi59++ {
                        fmt.Println(v25, v48, v48)
                        var v60 uint32 = 4294967295
                        _ = fi59
                        _ = v60
                    }
                }
                for fi61 := 0; fi61 < 3; fi61++ {
                    _ = len(v1)
                    _ = fi61
                }
            }
            for fi62 := 0; fi62 < 8; fi62++ {
                v8 = v8
                v63 := GenS7{F0: -2147483648, F1: v25}
                switch 32767 {
                case -901:
                    v8 = v8
                }
                if false {
                    var v64 float64 = -2.25
                    var v65 bool = v25
                    _ = v64
                    _ = v65
                } else {
                    v66 := GenS8{F0: -1}
                    v67 := []int{(fi51 << ((255) & 63))}
                    var v68 uint8 = 255
                    v8 = "foo bar"
                    _ = v66
                    _ = v67
                    _ = v68
                }
                if false {
                    v69 := GenS8{F0: 382}
                    fmt.Println(fi51, v4, fi62)
                    switch "" {
                    case "a":
                        if (v2 == (-(v2))) {
                            var v70 uint32 = v48
                            _ = v70
                        } else {
                            fmt.Println(v25, fi62, v8)
                            v4 += 42
                            var v71 int8 = 127
                            _ = v71
                        }
                    default:
                        v4 = v4
                        var v72 int16 = (v2 ^ 1)
                        _ = v72
                    }
                    _ = v69
                } else {
                    v4 = v4
                    v2 -= v2
                    v73 := GenS4{F0: -473}
                    _ = v73
                }
                v4 = v4
                v25 = (4294967295 != uint(v48))
                var v74 uint64 = 4294967295
                delete(v1, v8)
                _ = fi62
                _ = v63
                _ = v74
            }
            if false {
                v75 := GenS1{F0: 0, F1: 0, F2: 1}
                _ = v75
            }
            _ = fi51
        }
    default:
        v76 := GenS9{F0: int32(v2), F1: "\n\t", F2: 2}
        _ = v76
    }
    delete(v3, v48)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v48
    _ = v8
}
