package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 uint64; F2 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 108
    var v3 string = "foo bar"
    var v4 uint64 = 848
    var v5 string = v3
    v6 := []uint32{}
    var v7 string = ((v5 + "foo bar") + "hello")
    if (2 == 234) {
        _ = len(v6)
    }
    if ((!((v2 ^ (v2 | (v2 / ((99) | 1)))) < (v2 + v2))) && true) {
        fmt.Println(v7)
        v1["a"] = -792
    } else {
        v4 += (v4 + 1)
        v7 = ("foo bar" + v7)
        if (!(((967 == v4) && (-435 == -1)) && (65535 < 513))) {
            if true {
                v3 = v7
                var v8 uint8 = 2
                var v9 uint32 = 65535
                _ = v8
                _ = v9
            } else {
                switch 127 {
                case -2147483648:
                    v10 := GenS1{F0: 0}
                    _ = v10
                case -1:
                    v2 = 127
                case 693:
                    fmt.Println(v5, v7, v4)
                }
            }
        } else {
            switch 862 {
            case -1:
                v11 := GenS1{F0: 1}
                _ = v11
            case 406:
                for fi12 := 0; fi12 < 8; fi12++ {
                    if ((fi12 - (fi12 << ((65535) & 63))) <= fi12) {
                        var v13 bool = ((0 == 832) || (fi12 != (fi12 | (fi12 % ((fi12) | 1)))))
                        if true {
                            v3 = v3
                        } else {
                            v4 += v4
                            v14 := GenS1{F0: 649}
                            _ = v14
                        }
                        _ = v13
                    }
                    v5 = "foo bar"
                    v4 = (v4 & 4294967295)
                    _ = fi12
                }
            default:
                v15 := GenS2{F0: -519, F1: 2, F2: 2}
                _ = v15
            }
        }
    }
    var v16 uint8 = 127
    v17 := GenS1{F0: 530}
    var v18 int16 = 1
    switch "a" {
    case "":
        v2 = 74
    case "a":
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 string = (v7 + ((v3 + v3) + v7))
            v18 -= 1
            _ = fi19
            _ = v20
        }
    }
    var v21 int8 = (v2 % (((-(v2))) | 1))
    v22 := GenS2{F0: v18, F1: 704, F2: 2}
    if len(v6) > 0 {
        v6[0] = 985
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 float64 = float64(3.14159)
        switch "foo bar" {
        case "hello":
            if true {
                var v25 uint64 = 0
                fmt.Println(v4)
                for fi26 := 0; fi26 < 13; fi26++ {
                    fmt.Println(v18, v5)
                    _ = fi26
                }
                _ = v25
            } else {
                v27 := GenS1{F0: (v18 / ((v18) | 1))}
                v28 := GenS3{F0: fi23}
                var v29 int16 = (v18 % ((v18) | 1))
                _ = v27
                _ = v28
                _ = v29
            }
        case "\n\t":
            var v30 float64 = (float64(0.1) - (float64(3.14159) * 3.14159))
            _ = v30
        case "a":
            var v31 float32 = 0.1
            _ = v31
        default:
            v32 := GenS4{F0: 18446744073709551615, F1: (v4 ^ (v4 | (v4 % ((4294967295) | 1)))), F2: v2}
            _ = v32
        }
        var v33 uint16 = 0
        _ = fi23
        _ = v24
        _ = v33
    }
    if true {
        if false {
            if len(v6) > 0 {
                v6[0] = 1
            }
            for fi34 := 0; fi34 < 17; fi34++ {
                for fi35 := 0; fi35 < 16; fi35++ {
                    if false {
                        v7 = ("hello" + ((v3 + ((v7 + v7) + (v7 + "a"))) + (("foo bar" + v7) + "\n\t")))
                        v3 = v7
                        var v36 int64 = 1
                        _ = v36
                    } else {
                        v2 -= 127
                    }
                    _ = fi35
                }
                for fi37 := 0; fi37 < 19; fi37++ {
                    var v38 uint32 = uint32(v16)
                    _ = len(v1)
                    _ = fi37
                    _ = v38
                }
                for fi39 := 0; fi39 < 11; fi39++ {
                    var v40 float64 = -2.25
                    var v41 int16 = v18
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                _ = fi34
            }
            _ = len(v6)
        } else {
            for fi42 := 0; fi42 < 6; fi42++ {
                if len(v6) > 0 {
                    v6[0] = 2147483648
                }
                for fi43 := 0; fi43 < 9; fi43++ {
                    var v44 int16 = 0
                    _ = fi43
                    _ = v44
                }
                var v45 float32 = 1.5
                _ = fi42
                _ = v45
            }
        }
    } else {
        delete(v1, "a")
        switch "" {
        case "\n\t":
            var v46 string = v7
            _ = v46
        case "foo bar":
            var v47 uint16 = 2
            _ = v47
        case "a":
            v2 += (v2 & (v21 | v21))
        }
        if false {
            switch 128 {
            case -9223372036854775808:
                var v48 int16 = (v18 | v18)
                _ = v48
            case 9223372036854775807:
                var v49 uint32 = 4294967295
                _ = v49
            case -418:
                v21 = v2
            }
            var v50 int = 1
            _ = v50
        } else {
            for fi51 := 0; fi51 < 15; fi51++ {
                var v52 bool = (715 != 4294967295)
                delete(v1, v5)
                _ = fi51
                _ = v52
            }
            fmt.Println(v4, v16)
        }
    }
    switch v5 {
    case "\n\t":
        for fi53 := 0; fi53 < 4; fi53++ {
            if false {
                for fi54 := 0; fi54 < 1; fi54++ {
                    v21 += v2
                    v4 = (^(v4))
                    _ = fi54
                }
                var v55 int8 = int8(v16)
                var v56 int16 = int16(v4)
                _ = v55
                _ = v56
            } else {
                for fi57 := 0; fi57 < 13; fi57++ {
                    v1[""] = (fi53 | fi57)
                    var v58 int16 = (v18 ^ -800)
                    v16 -= v16
                    if true {
                        var v59 int16 = v18
                        var v60 bool = true
                        v21 = v2
                        _ = v59
                        _ = v60
                    }
                    if false {
                        var v61 int32 = -2147483648
                        _ = v61
                    }
                    _ = fi57
                    _ = v58
                }
                v3 = v3
                for fi62 := 0; fi62 < 4; fi62++ {
                    var v63 uint64 = (^(v4))
                    v64 := GenS3{F0: fi62}
                    switch fi62 {
                    case 563:
                        switch "foo bar" {
                        case "\n\t":
                            _ = len(v1)
                        default:
                            var v65 int16 = (v18 + v18)
                            _ = v65
                        }
                        var v66 int8 = (v21 & 0)
                        _ = v66
                    case -988:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    case -2147483648:
                        fmt.Println(v3, v3, v18)
                    default:
                        var v67 string = ""
                        _ = v67
                    }
                    _ = fi62
                    _ = v63
                    _ = v64
                }
            }
            var v68 uint = 1
            switch (((v5 + "foo bar") + v7) + v5) {
            case "a":
                var v69 uint8 = 0
                _ = v69
            default:
                var v70 uint64 = v4
                _ = v70
            }
            _ = fi53
            _ = v68
        }
    case "foo bar":
        if (-1 <= v21) {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v71 uint = 1
            v72 := GenS2{F0: (v18 + v18), F1: 42, F2: 0}
            _ = v71
            _ = v72
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
