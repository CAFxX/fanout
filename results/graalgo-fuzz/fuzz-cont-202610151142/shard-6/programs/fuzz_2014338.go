package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint) uint16 {
    fmt.Println(p0)
    p0 += p0
    var v4 uint16 = 2
    _ = p0
    _ = v4
    return 255
}
func genFn1(p0 int, p1 uint) uint16 {
    p1 -= (^(p1))
    v5 := []int{(p0 | 417), p0}
    v6 := GenS1{F0: -1, F1: false, F2: (^(p1))}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 650
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    v3 := []int{-2147483648}
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            fmt.Println(fi8)
            if (0 <= int32(fi7)) {
                var v9 int16 = -1
                var v10 float32 = 1.5
                switch fi8 {
                case 128:
                    v10 = (v10 + ((float32(1e10) * 3.14159) / ((3.14159 / (v10 + v10)) * (0.1 / 1.0))))
                    v11 := GenS2{F0: 1, F1: 0}
                    fmt.Println(fi7, v9)
                    _ = v11
                default:
                    fmt.Println(v10)
                }
                _ = v10
                _ = v9
            } else {
                var v12 int16 = -674
                v12 = (v12 & -225)
                _ = v12
            }
            _ = len(v2)
            _ = fi8
        }
        if false {
            switch fi7 {
            case -322:
                for fi13 := 0; fi13 < 11; fi13++ {
                    switch (("a" + "foo bar") + "") {
                    case "foo bar":
                        v14 := GenS2{F0: 375, F1: 1}
                        _ = v14
                    case "":
                        var v15 int = (fi7 | (^(fi7)))
                        _ = v15
                    }
                    var v16 int8 = -17
                    _ = fi13
                    _ = v16
                }
            case 4294967295:
                var v17 int64 = 0
                _ = v17
            default:
                var v18 int = -312
                _ = v18
            }
        } else {
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 uint32 = 286
                if (487 != 482) {
                    _ = len(v2)
                    var v21 int8 = 119
                    _ = v21
                } else {
                    if true {
                        switch "foo bar" {
                        case "":
                            if false {
                                switch 361 {
                                case 2147483647:
                                    v20 -= v20
                                case 128:
                                    v20 = v20
                                case 9223372036854775807:
                                    if false {
                                        v20 = v20
                                        v20 = v20
                                    }
                                    v20 += v20
                                    var v22 uint16 = genFn1((fi7 - fi7), 93)
                                    _ = v22
                                default:
                                    v23 := GenS3{F0: 52, F1: 1}
                                    _ = v23
                                }
                                delete(v2, (-(fi19)))
                            }
                        case "foo bar":
                            if true {
                                v20 = (^(v20))
                            }
                        }
                        var v24 int64 = -1
                        _ = v24
                    } else {
                        if ((true && true) || false) {
                            delete(v1, 2147483648)
                            if false {
                                switch fi7 {
                                case 32767:
                                    var v25 uint16 = 42
                                    v20 += (v20 - v20)
                                    v26 := GenS1{F0: 100, F1: (true && (18446744073709551615 >= 4294967295)), F2: 42}
                                    _ = v25
                                    _ = v26
                                case 965:
                                    fmt.Println(v20)
                                case 9223372036854775807:
                                    v27 := GenS1{F0: 0, F1: (1 > 0), F2: 974}
                                    _ = v27
                                }
                            }
                            switch (fi7 >> ((42) & 63)) {
                            case -1:
                                delete(v2, (fi7 ^ (fi19 - (fi19 / ((9223372036854775807) | 1)))))
                            case 4294967295:
                                if ((fi7 / ((684) | 1)) == 51) {
                                    if (false && (true || ((v20 + (v20 & 2)) > (v20 >> ((167) & 63))))) {
                                        _ = len(v3)
                                        var v28 uint64 = 0
                                        _ = v28
                                    } else {
                                        v20 = (v20 | (v20 % ((v20) | 1)))
                                    }
                                    fmt.Println(fi7, v20, fi19)
                                }
                            case -9223372036854775808:
                                var v29 int16 = 32767
                                _ = v29
                            }
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v30 uint = 0
                            _ = v30
                        }
                    }
                }
                v20 = (v20 | v20)
                _ = fi19
                _ = v20
            }
        }
        if false {
            for fi31 := 0; fi31 < 17; fi31++ {
                var v32 float32 = 3.14159
                _ = fi31
                _ = v32
            }
        } else {
            if true {
                fmt.Println(fi7)
                switch ((("hello" + ("foo bar" + ((("" + "\n\t") + "a") + "hello"))) + "hello") + ("hello" + "foo bar")) {
                case "\n\t":
                    var v33 float64 = 1e10
                    _ = v33
                default:
                    fmt.Println(fi7)
                }
            }
            fmt.Println(fi7)
        }
        _ = fi7
    }
    v34 := GenS3{F0: -96, F1: 307}
    var v35 uint16 = 1
    v36 := GenS1{F0: -99, F1: (-36 <= -1), F2: 0}
    var v37 int = 32768
    switch 9223372036854775807 {
    case -80:
        v38 := GenS1{F0: 19, F1: false, F2: 255}
        _ = v38
    case 492:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 0:
        for fi39 := 0; fi39 < 16; fi39++ {
            var v40 string = "a"
            v41 := GenS4{F0: 0, F1: 42, F2: 2147483648}
            _ = fi39
            _ = v40
            _ = v41
        }
    default:
        fmt.Println(v35, v37)
    }
    if true {
        if ((1e10 != (1e10 * 0.1)) || ((v37 | -816) < v37)) {
            switch 42 {
            case 1:
                if false {
                    for fi42 := 0; fi42 < 12; fi42++ {
                        var v43 bool = (-2.25 < -2.25)
                        var v44 float32 = 0.0
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                    v1[2147483648] = (-(v37))
                    v35 -= (^(v35))
                } else {
                    var v45 int = v37
                    v35 = 32767
                    _ = v45
                }
            case 9223372036854775807:
                v46 := v1[363]
                _ = v46
            case -1:
                var v47 int32 = 1
                _ = v47
            }
            var v48 uint32 = 65535
            _ = v48
        }
    }
    v35 = genFn0(4294967295)
    if ((v35 - v35) != (v35 & (v35 | v35))) {
        fmt.Println(v37, v37)
        v35 = v35
    }
    var v49 int = v37
    fmt.Println(v49, v35)
    fmt.Println(v35, v37)
    if ("hello" >= "a") {
        v50 := v1[65535]
        _ = v50
    } else {
        if (!false) {
            switch (v49 & v37) {
            case -624:
                fmt.Println(v49, v35)
            case 1:
                v51 := GenS5{F0: 721}
                _ = v51
            }
        } else {
            v52 := GenS1{F0: -1, F1: false, F2: 2}
            v49 += v49
            _ = v52
        }
        v49 = 9223372036854775807
        for fi53 := 0; fi53 < 16; fi53++ {
            var v54 int32 = 42
            _ = fi53
            _ = v54
        }
    }
    if false {
        v37 = v37
        v3 = append(v3, (v37 + (v49 | (v37 - v49))))
    }
    var v55 bool = ((v35 | (v35 % ((genFn1((^(v37)), 2)) | 1))) > v35)
    v49 = v49
    v1[2147483648] = (v49 * 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v49:", v49)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v49
    _ = v55
}
