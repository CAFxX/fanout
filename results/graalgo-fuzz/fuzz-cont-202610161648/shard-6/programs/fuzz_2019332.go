package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int8; F1 uint16; F2 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 uint }
func genFn0(p0 int16, p1 uint64, p2 uint16) int64 {
    for fi5 := 0; fi5 < 18; fi5++ {
        fmt.Println(p2)
        _ = fi5
    }
    p1 = p1
    var v6 int8 = 48
    fmt.Println(p2, p1, v6)
    v6 -= v6
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}
func genFn1() uint64 {
    var v7 uint64 = 42
    if true {
        if true {
            var v8 bool = (((2 <= 65535) && (((^(v7)) < (v7 % ((v7) | 1))) && false)) || false)
            fmt.Println(v8)
            _ = v8
        }
    } else {
        v9 := GenS1{F0: 0, F1: 18446744073709551615}
        var v10 bool = true
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        if (1 != 4294967295) {
            if false {
                v12 := GenS1{F0: 0, F1: 4294967295}
                _ = v12
            }
        } else {
            var v13 int = 4294967295
            for fi14 := 0; fi14 < 9; fi14++ {
                var v15 uint = 42
                var v16 uint16 = 752
                switch (v13 - -1) {
                case -193:
                    var v17 int64 = 164
                    _ = v17
                default:
                    v18 := []int{}
                    var v19 uint16 = v16
                    _ = v18
                    _ = v19
                }
                var v20 bool = true
                _ = fi14
                _ = v15
                _ = v16
                _ = v20
            }
            v21 := []int{v13, v13}
            _ = v13
            _ = v21
        }
        switch "\n\t" {
        case "foo bar":
            for fi22 := 0; fi22 < 14; fi22++ {
                v23 := GenS2{F0: 0, F1: 545, F2: -1}
                _ = fi22
                _ = v23
            }
        default:
            var v24 int64 = 9223372036854775807
            _ = v24
        }
        _ = fi11
    }
    if (v7 >= 148) {
        var v25 uint32 = 4294967295
        var v26 uint16 = 272
        _ = v25
        _ = v26
    }
    _ = v7
    return v7
}

func main() {
    var v1 int32 = 2147483647
    var v2 int = -2147483648
    var v3 uint = 4294967295
    v4 := make(map[uint32]string)
    var v27 float64 = (0.0 - 0.1)
    if false {
        v2 = v2
        v1 -= -1
        var v28 uint64 = 18446744073709551615
        _ = v28
    }
    v29 := GenS1{F0: 551, F1: 42}
    for fi30 := 0; fi30 < 16; fi30++ {
        var v31 int64 = -1
        switch "\n\t" {
        case "":
            v32 := GenS2{F0: -31, F1: 255, F2: v1}
            _ = v32
        case "foo bar":
            v2 = v2
        case "a":
            var v33 uint16 = 1
            _ = v33
        default:
            switch ("hello" + "\n\t") {
            case "a":
                fmt.Println(v3, v27)
            }
        }
        _ = fi30
        _ = v31
    }
    var v34 bool = false
    v35 := GenS1{F0: genFn0(-793, 18446744073709551615, uint16(v3)), F1: 1}
    var v36 uint64 = 534
    if ((136 == 1) || v34) {
        v37 := v4[0]
        fmt.Println(v27, v3, v2)
        _ = v37
    } else {
        v38 := GenS2{F0: -95, F1: 2, F2: 2147483647}
        switch v2 {
        case 947:
            v39 := GenS1{F0: genFn0(-32768, 192, 750), F1: (^(v3))}
            _ = v39
        case 0:
            if (-4 == 75) {
                v34 = ((v1 + v1) == 2147483647)
                v27 += (-2.25 / 1e10)
            }
        default:
            if (false && v34) {
                v1 -= (v1 << ((v3) & 63))
            }
        }
        v2 += (v2 | 128)
        _ = v38
    }
    for fi40 := 0; fi40 < 6; fi40++ {
        var v41 uint64 = v36
        v41 += genFn1()
        v2 -= fi40
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 17; fi42++ {
        var v43 bool = false
        if v43 {
            v44 := GenS3{F0: fi42}
            _ = v44
        } else {
            var v45 string = "foo bar"
            v46 := GenS4{F0: 60, F1: 9223372036854775807, F2: 127}
            var v47 uint16 = 142
            _ = v45
            _ = v46
            _ = v47
        }
        var v48 int32 = 1
        _ = fi42
        _ = v43
        _ = v48
    }
    v49 := v4[0]
    v49 = (v49 + v49)
    for fi50 := 0; fi50 < 20; fi50++ {
        if true {
            v49 = ((v49 + (("hello" + v49) + v49)) + "")
            fmt.Println(v1)
        } else {
            fmt.Println(v27, v3, v1)
            fmt.Println(v34, v1, v34)
        }
        switch v49 {
        case "foo bar":
            v36 = 0
        }
        if (v34 && v34) {
            switch fi50 {
            case 387:
                for fi51 := 0; fi51 < 10; fi51++ {
                    var v52 uint8 = 127
                    v53 := GenS5{F0: 1}
                    v1 = 107
                    _ = fi51
                    _ = v52
                    _ = v53
                }
            case -1:
                var v54 uint32 = 346
                _ = v54
            }
            fmt.Println(v34, v27, v2)
            if (!((v3 - 2) > 4294967295)) {
                if (v27 >= v27) {
                    if v34 {
                        if (v3 != (v3 + v3)) {
                            var v55 uint16 = 0
                            var v56 bool = v34
                            v57 := []int{v2, -418, fi50}
                            _ = v55
                            _ = v56
                            _ = v57
                        } else {
                            for fi58 := 0; fi58 < 10; fi58++ {
                                v59 := GenS6{F0: 32767, F1: (v1 - v1)}
                                switch (fi50 + 1) {
                                case 42:
                                    v27 -= float64(0.1)
                                case -2147483648:
                                    var v60 int = (fi58 & (fi58 / ((v2) | 1)))
                                    _ = v60
                                case 127:
                                    var v61 uint64 = v36
                                    _ = v61
                                }
                                _ = fi58
                                _ = v59
                            }
                            for fi62 := 0; fi62 < 14; fi62++ {
                                if v34 {
                                    v27 = 3.14159
                                } else {
                                    fmt.Println(v27, v27)
                                    var v63 int64 = 9223372036854775807
                                    _ = v63
                                }
                                v64 := []int{-384, (fi50 / ((fi62) | 1)), fi62}
                                _ = fi62
                                _ = v64
                            }
                            v65 := GenS7{F0: 250}
                            _ = v65
                        }
                    }
                } else {
                    v49 = "\n\t"
                    switch ("\n\t" + (v49 + v49)) {
                    case "":
                        _ = len(v4)
                    case "foo bar":
                        var v66 uint8 = 220
                        _ = v66
                    default:
                        v1 = v1
                    }
                }
                v49 = v49
                v67 := GenS3{F0: 2147483647}
                _ = v67
            }
        } else {
            v36 = (v36 - 18446744073709551615)
            switch v49 {
            case "a":
                v68 := GenS6{F0: 32767, F1: v1}
                _ = v68
            case "hello":
                v69 := v4[2147483648]
                _ = v69
            default:
                v2 = fi50
            }
            fmt.Println(v3, v49, v36)
        }
        _ = fi50
    }
    for fi70 := 0; fi70 < 17; fi70++ {
        v71 := GenS7{F0: v3}
        v34 = (!(3 < 42))
        _ = fi70
        _ = v71
    }
    var v72 uint64 = (v36 >> ((v3) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v49
    _ = v72
}
