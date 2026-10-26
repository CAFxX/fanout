package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int64, p1 int64) uint32 {
    var v6 float32 = float32((float64(0.1) - -2.25))
    v7 := []int{32767, 127, 9223372036854775807}
    var v8 uint64 = 828
    var v9 uint8 = 246
    var v10 float32 = 0.1
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint64, p1 bool) uint8 {
    v11 := []int{}
    switch 127 {
    case -101:
        var v12 uint = 18446744073709551615
        _ = v12
    case -927:
        p0 = (p0 | (p0 * p0))
    case 1:
        var v13 int = -224
        _ = v13
    }
    if (!(-728 >= -57)) {
        var v14 int16 = 0
        _ = v14
    } else {
        for fi15 := 0; fi15 < 1; fi15++ {
            switch (fi15 << ((1) & 63)) {
            case 32768:
                v16 := GenS1{F0: ((!p1) && false), F1: 0, F2: 2}
                _ = v16
            default:
                for fi17 := 0; fi17 < 9; fi17++ {
                    p1 = false
                    switch (fi17 + -9223372036854775808) {
                    case -696:
                        var v18 float64 = 0.0
                        _ = v18
                    case -379:
                        var v19 bool = p1
                        _ = v19
                    }
                    fmt.Println(p1)
                    _ = fi17
                }
            }
            if ("hello" < (("" + "") + ("a" + "a"))) {
                for fi20 := 0; fi20 < 20; fi20++ {
                    if true {
                        v21 := GenS2{F0: 427}
                        if (!(4294967295 <= 587)) {
                            var v22 int64 = 0
                            p0 -= 2
                            switch "foo bar" {
                            case "":
                                var v23 int8 = 100
                                _ = v23
                            case "hello":
                                _ = len(v11)
                            case "a":
                                _ = len(v11)
                            }
                            _ = v22
                        }
                        _ = v21
                    }
                    v24 := GenS2{F0: -1}
                    _ = fi20
                    _ = v24
                }
                _ = len(v11)
                if false {
                    var v25 float64 = 1.5
                    var v26 uint64 = 18446744073709551615
                    _ = v25
                    _ = v26
                } else {
                    if ((!p1) && p1) {
                        var v27 float64 = (0.0 - -2.25)
                        _ = len(v11)
                        _ = v27
                    } else {
                        p1 = false
                    }
                }
            }
            _ = fi15
        }
    }
    if (((float64(-2.25) + (0.0 / ((((float64(1e10) - ((-2.25 / 1.0) + (3.14159 - 1e10))) * 1e10) / -2.25) - (((-2.25 / 3.14159) + 3.14159) + 0.0)))) - -2.25) <= (3.14159 * (-2.25 * 3.14159))) {
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = v11
    return 42
}

func main() {
    var v1 uint64 = 268
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]string)
    v4 := []int{-9223372036854775808, 2147483647, 32767}
    v5 := []int{221, -660, 42}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        v1 = v2
        for fi28 := 0; fi28 < 16; fi28++ {
            var v29 uint = 4294967295
            switch fi28 {
            case 4294967295:
                if true {
                    var v30 bool = (1 > v29)
                    _ = v30
                } else {
                    switch "\n\t" {
                    case "a":
                        var v31 int16 = -45
                        _ = v31
                    case "foo bar":
                        if (-1 < int64(v1)) {
                            var v32 uint64 = v1
                            _ = v32
                        }
                    case "":
                        switch 9223372036854775807 {
                        case 1:
                            v33 := GenS1{F0: (65535 >= 1), F1: 32767, F2: (v29 / (((^(v29))) | 1))}
                            _ = v33
                        case 4294967295:
                            var v34 int = fi28
                            _ = v34
                        case 2147483647:
                            if ((3.14159 + float32((3.14159 * 1e10))) <= (0.0 - 0.1)) {
                                for fi35 := 0; fi35 < 10; fi35++ {
                                    switch (("foo bar" + "foo bar") + "\n\t") {
                                    case "foo bar":
                                        v1 -= v2
                                    default:
                                        switch 32768 {
                                        case 42:
                                            v2 -= v2
                                        default:
                                            v36 := GenS3{F0: 42}
                                            _ = v36
                                        }
                                    }
                                    _ = fi35
                                }
                                v37 := GenS1{F0: false, F1: 32767, F2: (v29 / ((v29) | 1))}
                                _ = v37
                            }
                        default:
                            var v38 string = ("hello" + "foo bar")
                            _ = v38
                        }
                    default:
                        v39 := v3[2]
                        _ = v39
                    }
                }
            case 0:
                v1 += v1
            default:
                var v40 uint = (v29 >> (((v29 - (v29 ^ v29))) & 63))
                _ = v40
            }
            _ = fi28
            _ = v29
        }
        var v41 float64 = (0.0 - 3.14159)
        _ = v41
    }
    v3[927] = "a"
    v42 := GenS1{F0: false, F1: 0, F2: 1}
    var v43 uint16 = 6
    var v44 uint16 = (v43 >> ((0) & 63))
    v45 := v3[genFn0(-1, 0)]
    var v46 float32 = float32(((3.14159 * (0.1 + 0.0)) + 0.0))
    v2 = (^(v1))
    v5 = append(v5, 32767)
    v47 := GenS1{F0: false, F1: 42, F2: 255}
    for fi48 := 0; fi48 < 10; fi48++ {
        fmt.Println(v2, v2)
        if (!(!(0 == 1))) {
            var v49 float32 = v46
            v46 -= v49
            for fi50 := 0; fi50 < 14; fi50++ {
                v1 = (v2 - v2)
                _ = fi50
            }
            _ = v49
        } else {
            var v51 int16 = -959
            _ = v51
        }
        _ = fi48
    }
    fmt.Println(v46, v44, v44)
    var v52 int = 32768
    for fi53 := 0; fi53 < 7; fi53++ {
        v54 := GenS3{F0: -855}
        if true {
            var v55 float32 = v46
            v56 := GenS2{F0: 2147483647}
            _ = v55
            _ = v56
        } else {
            fmt.Println(v44, fi53, v45)
            v57 := GenS1{F0: ((!true) && (true || false)), F1: 93, F2: 863}
            _ = v57
        }
        delete(v3, 0)
        _ = fi53
        _ = v54
    }
    switch v45 {
    case "foo bar":
        _ = len(v3)
    case "hello":
        v2 -= 0
    case "":
        var v58 float32 = (((v46 + float32((1.5 - 0.0))) + ((v46 * 1e10) + v46)) - (((1e10 + v46) * (v46 - (0.1 + v46))) / v46))
        _ = v58
    }
    var v59 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v52
    _ = v59
}
