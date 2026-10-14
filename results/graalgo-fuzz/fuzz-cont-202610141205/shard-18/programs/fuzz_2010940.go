package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
type GenS3 struct { F0 int16; F1 int64; F2 uint64 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint8) uint16 {
    switch -837 {
    case 32768:
        var v7 uint = 65535
        _ = v7
    case 127:
        var v8 bool = true
        _ = v8
    case -2147483648:
        switch ("" + "\n\t") {
        case "\n\t":
            v9 := []int{-9223372036854775808, 139, -2147483648, -9223372036854775808}
            _ = v9
        default:
            v10 := GenS1{F0: -620, F1: 255}
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := GenS1{F0: (-(fi11)), F1: 246}
        p0 = (^(p0))
        _ = fi11
        _ = v12
    }
    if false {
        p0 -= (p0 & 255)
        v13 := []int{}
        fmt.Println(p0)
        _ = v13
    }
    v14 := []int{302, 999, -711, 128}
    _ = p0
    _ = v14
    return 0
}

func main() {
    v1 := []uint32{42, 0}
    var v2 float32 = 3.14159
    var v3 string = "hello"
    var v4 int16 = -506
    v5 := make(map[int]string)
    var v6 int = 32767
    for fi15 := 0; fi15 < 16; fi15++ {
        if (!(18446744073709551615 >= 0)) {
            switch v6 {
            case 42:
                var v16 uint16 = 32767
                _ = v16
            }
            v4 = (v4 - (-(v4)))
            fmt.Println(fi15, v4)
        }
        var v17 uint8 = 127
        var v18 uint8 = (v17 << ((908) & 63))
        _ = fi15
        _ = v17
        _ = v18
    }
    v1 = append(v1, 4294967295)
    var v19 float32 = v2
    for fi20 := 0; fi20 < 13; fi20++ {
        v2 += v2
        fmt.Println(v6, fi20)
        fmt.Println(v19, fi20, fi20)
        _ = fi20
    }
    v4 -= -32768
    _ = len(v1)
    var v21 uint64 = 831
    if ((124 > 255) || false) {
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 uint8 = 16
            switch v3 {
            case "\n\t":
                var v24 int32 = 42
                _ = v24
            case "hello":
                for fi25 := 0; fi25 < 5; fi25++ {
                    fmt.Println(v3, v3)
                    if false {
                        var v26 uint32 = 0
                        if true {
                            v5[(v6 ^ -1)] = "foo bar"
                            if true {
                                if (0 >= genFn0(1)) {
                                    v27 := GenS2{F0: (v3 + v3), F1: v23, F2: false}
                                    fmt.Println(fi22, v6, fi25)
                                    _ = v27
                                } else {
                                    var v28 int8 = 50
                                    _ = v28
                                }
                                var v29 float64 = 1e10
                                v3 = v3
                                _ = v29
                            } else {
                                var v30 uint8 = v23
                                var v31 int32 = -2147483648
                                if false {
                                    switch v3 {
                                    case "":
                                        var v32 uint32 = (^(v26))
                                        _ = v32
                                    case "hello":
                                        fmt.Println(v30)
                                    default:
                                        v33 := v5[(v6 ^ 127)]
                                        _ = v33
                                    }
                                    _ = len(v5)
                                    var v34 float32 = v2
                                    _ = v34
                                }
                                var v35 int8 = 0
                                _ = v30
                                _ = v31
                                _ = v35
                            }
                        } else {
                            if false {
                                fmt.Println(fi25, v6)
                                fmt.Println(v19, v6, v23)
                            } else {
                                switch "foo bar" {
                                case "\n\t":
                                    v1 = append(v1, 2147483648)
                                case "hello":
                                    v19 = 1.5
                                }
                            }
                            if false {
                                var v36 float32 = float32(1.5)
                                _ = v36
                            } else {
                                _ = len(v5)
                                _ = len(v5)
                            }
                        }
                        v21 -= (v21 | v21)
                        _ = v26
                    } else {
                        v5[fi22] = (v3 + (v3 + (v3 + ("hello" + "\n\t"))))
                        var v37 string = ("" + v3)
                        v21 += v21
                        _ = v37
                    }
                    _ = fi25
                }
            case "a":
                var v38 uint8 = 128
                _ = v38
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi22
            _ = v23
        }
        for fi39 := 0; fi39 < 5; fi39++ {
            if true {
                for fi40 := 0; fi40 < 18; fi40++ {
                    if false {
                        var v41 int16 = -300
                        _ = v41
                    }
                    if (v3 != v3) {
                        v4 -= int16(v21)
                        v42 := GenS3{F0: (v4 * (v4 << ((502) & 63))), F1: -784, F2: v21}
                        v6 = fi39
                        _ = v42
                    } else {
                        v21 = (v21 ^ 456)
                        var v43 int32 = 1
                        _ = len(v5)
                        _ = v43
                    }
                    var v44 uint32 = 42
                    _ = fi40
                    _ = v44
                }
                for fi45 := 0; fi45 < 13; fi45++ {
                    fmt.Println(fi45, v21, v19)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v46 string = v3
                    _ = fi45
                    _ = v46
                }
            }
            switch -2147483648 {
            case -770:
                if (1e10 > 3.14159) {
                    if (((-122 == 100) || true) && true) {
                        if true {
                            v47 := GenS2{F0: "hello", F1: 127, F2: false}
                            switch v3 {
                            case "hello":
                                v2 -= v19
                            }
                            v4 += -996
                            _ = v47
                        } else {
                            _ = len(v1)
                            switch (v6 / ((v6) | 1)) {
                            case 0:
                                v21 = (^(v21))
                            default:
                                switch (fi39 & fi39) {
                                case 127:
                                    var v48 uint64 = 0
                                    _ = v48
                                case -9223372036854775808:
                                    switch fi39 {
                                    case 32767:
                                        v5[(fi39 - (v6 << ((2) & 63)))] = (("foo bar" + "foo bar") + v3)
                                    case 4294967295:
                                        var v49 bool = false
                                        _ = v49
                                    case 259:
                                        delete(v5, (v6 % ((2147483647) | 1)))
                                    }
                                }
                            }
                        }
                        var v50 uint8 = 42
                        _ = v50
                    } else {
                        var v51 int32 = 605
                        v52 := GenS4{F0: v51}
                        _ = v51
                        _ = v52
                    }
                } else {
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v5[(v6 * -2147483648)] = "a"
                    } else {
                        v4 = (v4 + v4)
                        var v53 string = v3
                        fmt.Println(v6, v6)
                        _ = v53
                    }
                    v54 := v5[1]
                    v6 = (fi39 & 4294967295)
                    _ = v54
                }
            }
            _ = fi39
        }
    }
    fmt.Println(v2)
    switch "" {
    case "\n\t":
        v19 = (1.5 - ((v19 - float32((float64(0.0) - ((0.0 / float64((1e10 / v2))) - 1e10)))) * ((1.5 / v19) + (0.0 * -2.25))))
    case "foo bar":
        for fi55 := 0; fi55 < 18; fi55++ {
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
            var v56 uint8 = 2
            _ = fi55
            _ = v56
        }
    }
    if false {
        var v57 uint64 = (v21 * 18446744073709551615)
        if true {
            delete(v5, (v6 << ((42) & 63)))
            v5[v6] = ("hello" + v3)
            v57 += v57
        }
        if true {
            var v58 string = (((v3 + "foo bar") + "") + "foo bar")
            for fi59 := 0; fi59 < 6; fi59++ {
                v5[v6] = v3
                _ = fi59
            }
            _ = v58
        }
        _ = v57
    }
    var v60 int32 = 42
    v21 += 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
}
