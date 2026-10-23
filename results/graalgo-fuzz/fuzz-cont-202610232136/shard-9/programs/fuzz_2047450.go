package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint16; F2 string }
type GenS5 struct { F0 int8; F1 uint16; F2 int16 }
func genFn0() int32 {
    switch "" {
    case "foo bar":
        var v6 int = 127
        _ = v6
    default:
        var v7 float32 = 0.1
        _ = v7
    }
    var v8 int32 = 0
    _ = v8
    return v8
}
func genFn1(p0 uint32, p1 uint64, p2 uint32) int32 {
    if false {
        var v9 uint64 = 2
        var v10 uint = 584
        var v11 uint = v10
        _ = v10
        _ = v11
        _ = v9
    } else {
        var v12 float32 = 0.1
        switch ("a" + ("foo bar" + "hello")) {
        case "\n\t":
            var v13 int = -974
            _ = v13
        case "":
            var v14 uint16 = 171
            _ = v14
        }
        var v15 uint8 = 2
        _ = v12
        _ = v15
    }
    p2 -= (p2 | (p2 | p0))
    fmt.Println(p0, p2, p0)
    var v16 float32 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return -2147483648
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 3.14159
    var v3 uint = 1
    var v4 float32 = float32(0.0)
    var v5 float32 = ((v4 - (v4 * float32(3.14159))) - v2)
    fmt.Println(v4)
    var v17 uint = v3
    v17 += (v3 << ((v3) & 63))
    delete(v1, ("\n\t" + "hello"))
    var v18 uint64 = 790
    switch "" {
    case "foo bar":
        v3 = 65535
    }
    fmt.Println(v5)
    if (0.1 >= (-2.25 + float64(v5))) {
        fmt.Println(v5, v2, v5)
        if ("hello" < "foo bar") {
            v17 = 2
            var v19 float32 = v2
            var v20 int16 = -668
            _ = v19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        v4 -= float32(3.14159)
        switch fi21 {
        case 32767:
            var v22 uint = 1
            _ = v22
        case -413:
            v4 = v5
        }
        _ = fi21
    }
    if (v18 >= (v18 + v18)) {
        v23 := []int{2147483647}
        var v24 float64 = 0.0
        fmt.Println(v5)
        _ = v23
        _ = v24
    } else {
        v2 = v2
        v25 := GenS1{F0: -1, F1: "foo bar", F2: v18}
        if true {
            fmt.Println(v2, v4, v5)
            var v26 string = ("hello" + "foo bar")
            if (100 > -128) {
                if false {
                    v27 := []int{}
                    switch v26 {
                    case "a":
                        if true {
                            v26 = "\n\t"
                        }
                    case "":
                        if false {
                            switch 128 {
                            case -1:
                                for fi28 := 0; fi28 < 9; fi28++ {
                                    v3 += v3
                                    v26 = v26
                                    _ = fi28
                                }
                            case 569:
                                switch 2147483647 {
                                case 9223372036854775807:
                                    v4 -= v5
                                }
                            }
                        } else {
                            switch 32767 {
                            case 0:
                                fmt.Println(v17, v5)
                            case 9223372036854775807:
                                v26 = v26
                            default:
                                switch 1 {
                                case 455:
                                    if true {
                                        var v29 uint = (v3 >> ((v3) & 63))
                                        if (!true) {
                                            if len(v27) >= 2 {
                                                v27 = v27[1:len(v27)]
                                            }
                                            var v30 uint16 = 417
                                            for fi31 := 0; fi31 < 6; fi31++ {
                                                v32 := GenS2{F0: fi31}
                                                _ = fi31
                                                _ = v32
                                            }
                                            _ = v30
                                        }
                                        fmt.Println(v2, v2)
                                        _ = v29
                                    }
                                default:
                                    v18 -= (v18 | (v18 ^ 187))
                                }
                            }
                            fmt.Println(v18, v18)
                        }
                    case "hello":
                        v26 = v26
                    default:
                        for fi33 := 0; fi33 < 15; fi33++ {
                            v34 := GenS3{F0: false}
                            switch (fi33 << ((v17) & 63)) {
                            case -2147483648:
                                if false {
                                    var v35 int32 = 2147483647
                                    var v36 string = ("\n\t" + "a")
                                    if true {
                                        for fi37 := 0; fi37 < 12; fi37++ {
                                            v1[(v36 + "")] = 431
                                            var v38 int64 = int64(v35)
                                            v1[v36] = 1
                                            _ = fi37
                                            _ = v38
                                        }
                                        v5 = ((float32((1.5 + 1.5)) * 0.0) * 0.1)
                                        v39 := v1[v36]
                                        _ = v39
                                    } else {
                                        v17 += (v3 / (((^(v17))) | 1))
                                        switch (fi33 - fi33) {
                                        case 32767:
                                            var v40 int16 = -1
                                            _ = v40
                                        case 32768:
                                            var v41 uint32 = 2147483648
                                            _ = v41
                                        default:
                                            if false {
                                                switch (fi33 + (fi33 - (fi33 - 2147483647))) {
                                                case 4294967295:
                                                    var v42 int64 = 0
                                                    _ = v42
                                                case 32768:
                                                    v17 -= (v17 | (v3 + (^(v17))))
                                                default:
                                                    if (!false) {
                                                        var v43 uint16 = uint16(v18)
                                                        fmt.Println(v4, v3)
                                                        _ = v43
                                                    } else {
                                                        if (-1 >= 32767) {
                                                            v4 = (v2 + float32(float64(((((0.0 * ((1.5 - v2) + v5)) + float32((1.5 - ((0.0 - 1e10) + -2.25)))) * v5) / (v5 + v5)))))
                                                        } else {
                                                            v26 = v26
                                                        }
                                                    }
                                                }
                                                v3 -= (v3 & 255)
                                                v44 := v1[(v26 + ((("foo bar" + ("foo bar" + (v36 + v36))) + v36) + v36))]
                                                _ = v44
                                            }
                                        }
                                        v26 = v26
                                    }
                                    _ = v35
                                    _ = v36
                                } else {
                                    var v45 uint8 = 2
                                    _ = v45
                                }
                            case 32767:
                                var v46 uint64 = (v18 << (((^(v3))) & 63))
                                _ = v46
                            case 42:
                                v17 = v3
                            default:
                                fmt.Println(v26, v2)
                            }
                            var v47 int8 = -128
                            _ = fi33
                            _ = v34
                            _ = v47
                        }
                    }
                    _ = v27
                } else {
                    v5 -= 0.0
                    v48 := GenS1{F0: -200, F1: v26, F2: 42}
                    _ = v48
                }
                if (127 != -727) {
                    fmt.Println(v5)
                    var v49 int = -39
                    v50 := GenS4{F0: (-680 != 0), F1: 1, F2: v26}
                    _ = v49
                    _ = v50
                }
                if (128 == 1) {
                    switch (v26 + (v26 + (v26 + v26))) {
                    case "a":
                        var v51 string = "a"
                        _ = v51
                    case "hello":
                        switch "foo bar" {
                        case "a":
                            v5 += (0.1 * (v5 / v5))
                        case "\n\t":
                            if (v2 != 1e10) {
                                for fi52 := 0; fi52 < 4; fi52++ {
                                    var v53 uint16 = 639
                                    switch "foo bar" {
                                    case "":
                                        var v54 int8 = 43
                                        _ = v54
                                    case "\n\t":
                                        if true {
                                            for fi55 := 0; fi55 < 9; fi55++ {
                                                v56 := GenS1{F0: 1, F1: ("hello" + v26), F2: v18}
                                                v57 := GenS5{F0: -128, F1: (v53 % ((v53) | 1)), F2: 42}
                                                _ = fi55
                                                _ = v56
                                                _ = v57
                                            }
                                        }
                                    case "foo bar":
                                        fmt.Println(v26, v26, v53)
                                    }
                                    _ = fi52
                                    _ = v53
                                }
                                if (!false) {
                                    var v58 uint64 = v18
                                    for fi59 := 0; fi59 < 7; fi59++ {
                                        fmt.Println(fi59, fi59, v2)
                                        v60 := GenS4{F0: (-9223372036854775808 != -535), F1: 1, F2: v26}
                                        v61 := []int{fi59, (fi59 % ((fi59) | 1)), 0, (^(fi59))}
                                        _ = fi59
                                        _ = v60
                                        _ = v61
                                    }
                                    _ = v58
                                }
                            }
                        case "":
                            var v62 bool = false
                            _ = v62
                        }
                    case "\n\t":
                        switch 2147483647 {
                        case -2147483648:
                            v3 += uint(v18)
                        case 138:
                            var v63 uint32 = 693
                            _ = v63
                        case 242:
                            v64 := GenS1{F0: -890, F1: ((v26 + ("\n\t" + "\n\t")) + ((v26 + "foo bar") + "hello")), F2: v18}
                            _ = v64
                        }
                    }
                    var v65 int16 = 1
                    _ = v65
                }
            } else {
                v5 = 0.1
            }
            _ = v26
        }
        _ = v25
    }
    v18 += (^(v18))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
