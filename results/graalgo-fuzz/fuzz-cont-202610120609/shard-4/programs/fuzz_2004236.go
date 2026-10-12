package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 uint32; F2 bool }
type GenS5 struct { F0 int; F1 uint64; F2 uint16 }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 uint; F1 bool; F2 int16 }
type GenS8 struct { F0 int64; F1 int64; F2 uint }

func main() {
    var v1 int = 9223372036854775807
    var v2 uint16 = 473
    var v3 string = "hello"
    v4 := make(map[string]string)
    var v5 uint = 1
    v2 += v2
    var v6 uint64 = 18446744073709551615
    if (!(-1 > -508)) {
        var v7 int16 = 405
        v8 := GenS1{F0: (v2 << (((v5 * 65535)) & 63))}
        for fi9 := 0; fi9 < 16; fi9++ {
            if false {
                v10 := GenS2{F0: (^(v2)), F1: 1}
                var v11 uint64 = uint64(v2)
                _ = v10
                _ = v11
            } else {
                var v12 string = "foo bar"
                v13 := []int{127, 267, fi9, (fi9 * v1)}
                switch v12 {
                case "a":
                    switch v12 {
                    case "hello":
                        v14 := GenS3{F0: 563}
                        _ = v14
                    case "\n\t":
                        if ((1e10 - 3.14159) > 3.14159) {
                            for fi15 := 0; fi15 < 12; fi15++ {
                                v16 := GenS2{F0: 755, F1: 42}
                                v17 := GenS4{F0: v7, F1: 4294967295, F2: (v5 < v5)}
                                if len(v13) > 0 {
                                    v13[0] = (fi15 % (((v1 | (fi9 >> ((255) & 63)))) | 1))
                                }
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                        } else {
                            var v18 uint64 = v6
                            for fi19 := 0; fi19 < 2; fi19++ {
                                fmt.Println(fi9, v12, v5)
                                fmt.Println(fi9, v2, v3)
                                v20 := GenS5{F0: 127, F1: v18, F2: v2}
                                _ = fi19
                                _ = v20
                            }
                            _ = v18
                        }
                    default:
                        v21 := GenS2{F0: v2, F1: -26}
                        _ = v21
                    }
                }
                _ = v12
                _ = v13
            }
            if false {
                var v22 int16 = (v7 ^ v7)
                for fi23 := 0; fi23 < 2; fi23++ {
                    v24 := GenS4{F0: v7, F1: 2147483647, F2: (-2147483648 >= 0)}
                    fmt.Println(v7, v3)
                    v5 -= (v5 | (^(v5)))
                    _ = fi23
                    _ = v24
                }
                for fi25 := 0; fi25 < 9; fi25++ {
                    v26 := []int{fi25, v1}
                    _ = fi25
                    _ = v26
                }
                _ = v22
            }
            v27 := []int{732, -2147483648, (fi9 - (fi9 | v1))}
            _ = fi9
            _ = v27
        }
        _ = v7
        _ = v8
    } else {
        var v28 int64 = int64(v1)
        _ = v28
    }
    var v29 string = v3
    v3 = "\n\t"
    fmt.Println(v3, v5)
    switch v1 {
    case 368:
        fmt.Println(v2, v29)
    case 32767:
        switch v1 {
        case 109:
            v30 := GenS6{F0: 0, F1: 1}
            _ = v30
        case 0:
            var v31 int = (v1 ^ 42)
            _ = v31
        case 848:
            v4["foo bar"] = ""
        default:
            var v32 uint32 = 42
            _ = v32
        }
    case 4294967295:
        if (173 != 936) {
            if ((v5 * v5) < (v5 >> ((65535) & 63))) {
                var v33 float64 = 3.14159
                if false {
                    for fi34 := 0; fi34 < 7; fi34++ {
                        switch ((("hello" + v3) + ((v29 + v29) + v3)) + ("\n\t" + ((v29 + v3) + "foo bar"))) {
                        case "hello":
                            var v35 int = (v1 << (((v5 - v5)) & 63))
                            _ = v35
                        default:
                            if false {
                                fmt.Println(v2)
                            } else {
                                fmt.Println(v33, fi34, fi34)
                            }
                        }
                        v36 := GenS2{F0: (v2 + (v2 & v2)), F1: 127}
                        _ = fi34
                        _ = v36
                    }
                    if true {
                        if (!(((v29 + v29) + "\n\t") == v3)) {
                            v37 := []int{}
                            v38 := GenS3{F0: 0}
                            _ = v37
                            _ = v38
                        }
                    } else {
                        v6 += (v6 >> (((v5 + 4)) & 63))
                        if true {
                            if ((true && (v5 > (^(v5)))) || false) {
                                if true {
                                    v1 -= 9223372036854775807
                                }
                            }
                            for fi39 := 0; fi39 < 10; fi39++ {
                                switch (v29 + (v3 + v29)) {
                                case "\n\t":
                                    var v40 int32 = 0
                                    _ = v40
                                case "hello":
                                    v41 := []int{(fi39 % (((v1 & 32767)) | 1)), v1}
                                    _ = v41
                                }
                                _ = fi39
                            }
                            var v42 float64 = 1.5
                            _ = v42
                        } else {
                            for fi43 := 0; fi43 < 5; fi43++ {
                                var v44 float64 = v33
                                _ = fi43
                                _ = v44
                            }
                        }
                    }
                } else {
                    v2 = (v2 & (v2 >> ((1) & 63)))
                    v29 = ("hello" + v3)
                    var v45 uint64 = 42
                    _ = v45
                }
                _ = v33
            } else {
                v6 = v6
                for fi46 := 0; fi46 < 9; fi46++ {
                    var v47 uint = v5
                    _ = fi46
                    _ = v47
                }
            }
            if (42 == 0) {
                v1 -= (v1 / (((v1 % (((-(v1))) | 1))) | 1))
            }
            for fi48 := 0; fi48 < 12; fi48++ {
                if false {
                    switch int(v5) {
                    case -1:
                        v49 := GenS7{F0: v5, F1: false, F2: -35}
                        _ = v49
                    case 0:
                        var v50 uint = v5
                        _ = v50
                    case 4294967295:
                        for fi51 := 0; fi51 < 12; fi51++ {
                            switch (^(v1)) {
                            case -429:
                                fmt.Println(fi48)
                            }
                            var v52 bool = false
                            v53 := GenS6{F0: 101, F1: 9223372036854775807}
                            v29 = v3
                            switch (v3 + v29) {
                            case "":
                                var v54 uint64 = (^(v6))
                                _ = v54
                            case "a":
                                v1 = (fi51 * fi51)
                            }
                            var v55 int16 = -32768
                            _ = fi51
                            _ = v52
                            _ = v53
                            _ = v55
                        }
                    default:
                        v1 = v1
                    }
                } else {
                    delete(v4, v29)
                }
                var v56 float32 = 1e10
                _ = fi48
                _ = v56
            }
        }
    }
    for fi57 := 0; fi57 < 2; fi57++ {
        fmt.Println(v1, fi57, v29)
        _ = fi57
    }
    for fi58 := 0; fi58 < 15; fi58++ {
        fmt.Println(v6, v1, v2)
        var v59 uint16 = v2
        _ = fi58
        _ = v59
    }
    if true {
        switch "a" {
        case "foo bar":
            v60 := GenS4{F0: 1, F1: 42, F2: true}
            _ = v60
        case "":
            var v61 int8 = -100
            _ = v61
        default:
            for fi62 := 0; fi62 < 1; fi62++ {
                v6 = (v6 + 2)
                _ = fi62
            }
        }
    } else {
        var v63 float32 = -2.25
        var v64 uint64 = 638
        if (1e10 >= v63) {
            var v65 int16 = -32768
            _ = len(v4)
            _ = v65
        } else {
            if false {
                var v66 int16 = 1000
                if true {
                    var v67 float64 = 1e10
                    var v68 float64 = 0.1
                    _ = v67
                    _ = v68
                }
                v69 := GenS8{F0: -1, F1: 9223372036854775807, F2: 2}
                _ = v66
                _ = v69
            } else {
                switch 0 {
                case 127:
                    var v70 int8 = 42
                    _ = v70
                }
                v71 := GenS1{F0: (v2 / ((v2) | 1))}
                _ = v71
            }
        }
        _ = v63
        _ = v64
    }
    if true {
        for fi72 := 0; fi72 < 14; fi72++ {
            v3 = v29
            v73 := []int{(v1 ^ (-(fi72))), (-(v1)), (^(v1))}
            _ = fi72
            _ = v73
        }
    } else {
        switch "" {
        case "":
            fmt.Println(v1, v3, v5)
        }
        switch v3 {
        case "foo bar":
            v74 := GenS3{F0: (v2 >> (((^(v5))) & 63))}
            _ = v74
        case "a":
            fmt.Println(v3, v2)
        case "":
            for fi75 := 0; fi75 < 15; fi75++ {
                for fi76 := 0; fi76 < 6; fi76++ {
                    var v77 string = (("a" + v3) + (v29 + v29))
                    _ = fi76
                    _ = v77
                }
                fmt.Println(v3, v5, v3)
                _ = fi75
            }
        default:
            if (!(868 != 42)) {
                v1 += 633
                if (!true) {
                    v6 = v6
                    v78 := GenS6{F0: 2, F1: 9223372036854775807}
                    _ = v78
                } else {
                    v1 -= v1
                }
                var v79 uint16 = 1
                _ = v79
            } else {
                for fi80 := 0; fi80 < 12; fi80++ {
                    var v81 int64 = 9223372036854775807
                    delete(v4, (v29 + v29))
                    _ = fi80
                    _ = v81
                }
            }
        }
    }
    v29 = v3
    for fi82 := 0; fi82 < 16; fi82++ {
        switch ("\n\t" + v29) {
        case "\n\t":
            var v83 float32 = (1.5 + -2.25)
            _ = v83
        case "a":
            v6 = (v6 | uint64(v2))
        default:
            var v84 uint16 = 1
            _ = v84
        }
        _ = fi82
    }
    v5 -= (v5 | 65535)
    var v85 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v85
}
