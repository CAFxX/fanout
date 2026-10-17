package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 string }
type GenS2 struct { F0 int64; F1 uint64; F2 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 bool; F2 int8 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 int16 }
func genFn0(p0 int8, p1 uint8, p2 string) uint64 {
    p0 = (p0 & (p0 - p0))
    var v7 uint8 = p1
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 1
    var v3 string = "foo bar"
    var v4 uint16 = 484
    var v5 uint = 255
    var v6 float64 = 0.0
    v1[42] = 2147483647
    switch ("a" + (("foo bar" + "\n\t") + v3)) {
    case "a":
        fmt.Println(v6)
    case "hello":
        fmt.Println(v4, v5)
    case "\n\t":
        fmt.Println(v4, v6)
    default:
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := GenS1{F0: 244, F1: "hello", F2: ((v3 + "\n\t") + v3)}
            _ = fi8
            _ = v9
        }
    }
    if false {
        fmt.Println(v4)
    } else {
        var v10 int32 = v2
        fmt.Println(v4, v10, v6)
        _ = v10
    }
    var v11 string = (v3 + "\n\t")
    var v12 int = -540
    var v13 int8 = -116
    var v14 uint64 = 497
    var v15 float64 = v6
    if true {
        if ((((!false) && false) || (((^(v5)) != v5) || (!(0.1 == (v6 / v15))))) || (v11 > ("" + "foo bar"))) {
            v16 := GenS1{F0: -9223372036854775808, F1: (v11 + v11), F2: v11}
            for fi17 := 0; fi17 < 5; fi17++ {
                if false {
                    switch "a" {
                    case "":
                        if false {
                            v18 := []int{(-(fi17)), 251, (fi17 - (-(v12)))}
                            v13 += (-(v13))
                            _ = v18
                        } else {
                            v12 += 32767
                        }
                    }
                } else {
                    v19 := v1[2147483648]
                    v20 := GenS1{F0: 42, F1: (v11 + "\n\t"), F2: ("hello" + v3)}
                    switch fi17 {
                    case -9223372036854775808:
                        fmt.Println(v14, v4, v14)
                    case 32768:
                        v21 := GenS2{F0: 1, F1: 9, F2: v13}
                        _ = v21
                    case -1:
                        v3 = v11
                    }
                    _ = v19
                    _ = v20
                }
                switch v3 {
                case "":
                    v6 = float64((0.0 - 0.0))
                case "foo bar":
                    if ((v14 & 683) < uint64(v5)) {
                        if (!true) {
                            for fi22 := 0; fi22 < 2; fi22++ {
                                var v23 int = (v12 >> ((4294967295) & 63))
                                var v24 float64 = v6
                                v12 -= 1
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                            v12 += (fi17 >> ((593) & 63))
                        } else {
                            v1[0] = 2147483647
                            v25 := GenS3{F0: 9223372036854775807}
                            _ = v25
                        }
                        if true {
                            for fi26 := 0; fi26 < 3; fi26++ {
                                v1[460] = (^(fi26))
                                v13 += v13
                                fmt.Println(v5)
                                _ = fi26
                            }
                            if true {
                                var v27 uint16 = (v4 * 0)
                                v4 = v27
                                v11 = v3
                                _ = v27
                            }
                            fmt.Println(v4)
                        }
                        if false {
                            v28 := GenS4{F0: 1, F1: ((!true) || true), F2: -1}
                            v13 -= -109
                            for fi29 := 0; fi29 < 11; fi29++ {
                                var v30 int64 = -1
                                _ = len(v1)
                                if false {
                                    if true {
                                        switch "hello" {
                                        case "hello":
                                            if ((-(v30)) < -1) {
                                                fmt.Println(v14, v12, v2)
                                                var v31 float64 = v15
                                                v32 := GenS2{F0: 42, F1: v14, F2: v13}
                                                fmt.Println(v4, v31)
                                                _ = v31
                                                _ = v32
                                            }
                                        case "\n\t":
                                            if false {
                                                fmt.Println(v30)
                                            } else {
                                                v33 := GenS5{F0: fi29, F1: 974}
                                                _ = v33
                                            }
                                        }
                                    } else {
                                        var v34 uint64 = 948
                                        v34 -= uint64(v5)
                                        _ = v34
                                    }
                                    var v35 int64 = v30
                                    _ = v35
                                }
                                _ = fi29
                                _ = v30
                            }
                            _ = v28
                        }
                    }
                case "hello":
                    var v36 float32 = 0.0
                    _ = v36
                default:
                    v13 += (-(v13))
                }
                _ = fi17
            }
            _ = v16
        }
        v37 := GenS3{F0: 132}
        for fi38 := 0; fi38 < 3; fi38++ {
            if ("\n\t" == "foo bar") {
                for fi39 := 0; fi39 < 4; fi39++ {
                    v40 := GenS4{F0: 1, F1: true, F2: v13}
                    var v41 uint = 358
                    v11 = ("" + "\n\t")
                    _ = fi39
                    _ = v40
                    _ = v41
                }
            } else {
                v6 -= v6
            }
            _ = fi38
        }
        _ = v37
    }
    var v42 string = v3
    for fi43 := 0; fi43 < 17; fi43++ {
        switch (fi43 / ((-786) | 1)) {
        case 4294967295:
            fmt.Println(v5)
        case 128:
            v11 = v11
        case -290:
            if (v14 == (v14 << ((65535) & 63))) {
                fmt.Println(v42, v5)
                v44 := GenS4{F0: -367, F1: (!(true && false)), F2: 1}
                _ = v44
            }
        }
        v45 := GenS5{F0: v12, F1: uint32(v4)}
        v4 = v4
        _ = fi43
        _ = v45
    }
    if (!(true && false)) {
        v46 := GenS5{F0: v12, F1: 869}
        _ = v46
    } else {
        v47 := GenS6{F0: int16(v13)}
        _ = v47
    }
    _ = len(v1)
    var v48 string = (v42 + "")
    for fi49 := 0; fi49 < 11; fi49++ {
        v13 += 42
        _ = fi49
    }
    var v50 int8 = -57
    v42 = v42
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v48
    _ = v5
    _ = v50
    _ = v6
}
