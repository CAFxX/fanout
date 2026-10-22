package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int64; F1 int64; F2 int32 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 bool; F1 int8; F2 int16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32; F1 int16; F2 int16 }
func genFn0(p0 bool, p1 int16) uint64 {
    var v6 int = 128
    v7 := GenS1{F0: 2, F1: (p1 / (((p1 / ((p1) | 1))) | 1)), F2: 947}
    for fi8 := 0; fi8 < 5; fi8++ {
        switch (fi8 | 2147483647) {
        case 128:
            if (209 != 127) {
                if (("\n\t" < (("\n\t" + (("\n\t" + "\n\t") + "foo bar")) + "hello")) || (128 != 0)) {
                    v9 := GenS2{F0: 67, F1: 65535}
                    _ = v9
                }
            }
        case 127:
            v6 -= (fi8 << ((2) & 63))
        }
        p1 -= p1
        v6 = 42
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1() int64 {
    v10 := GenS3{F0: -1, F1: -417}
    var v11 int = 128
    v11 += -2147483648
    _ = v10
    _ = v11
    return -417
}

func main() {
    var v1 int32 = 0
    v2 := make(map[uint32]uint32)
    v3 := []uint32{760}
    var v4 uint16 = 42
    var v5 uint16 = (v4 ^ (v4 / ((v4) | 1)))
    if true {
        for fi12 := 0; fi12 < 2; fi12++ {
            switch "foo bar" {
            case "\n\t":
                v5 = (v4 << ((1) & 63))
            case "a":
                v13 := GenS2{F0: 42, F1: 302}
                _ = v13
            case "hello":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            switch (-(fi12)) {
            case 326:
                for fi14 := 0; fi14 < 19; fi14++ {
                    if true {
                        v15 := GenS1{F0: genFn0(true, 42), F1: -699, F2: 18446744073709551615}
                        var v16 float64 = 0.1
                        v16 -= (v16 * 0.1)
                        _ = v15
                        _ = v16
                    } else {
                        if (true && false) {
                            var v17 int16 = -32768
                            v1 = (v1 - v1)
                            _ = v17
                        }
                        v18 := GenS4{F0: -244, F1: 0, F2: (v1 & v1)}
                        _ = v18
                    }
                    _ = fi14
                }
            case 42:
                fmt.Println(v1, fi12, v4)
            }
            for fi19 := 0; fi19 < 6; fi19++ {
                if false {
                    v20 := GenS5{F0: 42, F1: (fi19 >> ((42) & 63))}
                    v5 = v4
                    _ = v20
                }
                if ((v1 - (-(v1))) >= v1) {
                    fmt.Println(v1, v5)
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                    if false {
                        fmt.Println(fi19, v1, v1)
                        v21 := v2[65535]
                        var v22 int64 = 581
                        _ = v21
                        _ = v22
                    } else {
                        v1 -= 260
                        var v23 float64 = 0.0
                        fmt.Println(v1, fi12)
                        _ = v23
                    }
                    v4 = 42
                    switch fi12 {
                    case 32767:
                        var v24 int8 = int8(fi19)
                        _ = v24
                    case 127:
                        fmt.Println(fi12)
                    }
                    if (370 <= 42) {
                        var v25 uint16 = v4
                        if (9223372036854775807 == fi12) {
                            v1 -= v1
                            v25 += (^(v4))
                        }
                        if false {
                            v26 := GenS3{F0: v1, F1: -9223372036854775808}
                            _ = v26
                        }
                        _ = v25
                    }
                } else {
                    v4 = (v5 | (^(v4)))
                    switch "" {
                    case "hello":
                        v5 = v5
                    case "a":
                        fmt.Println(v1, v4)
                    case "\n\t":
                        var v27 string = "hello"
                        _ = v27
                    default:
                        v28 := GenS5{F0: fi19, F1: (fi12 ^ (^(fi19)))}
                        _ = v28
                    }
                    v4 += (v5 * (v5 | (v5 << ((709) & 63))))
                }
                _ = fi19
            }
            _ = fi12
        }
    }
    v29 := GenS6{F0: false, F1: 127, F2: 1}
    var v30 string = "hello"
    fmt.Println(v1, v4, v30)
    if false {
        v30 = v30
        v31 := GenS7{F0: 128}
        v5 = 65535
        _ = v31
    } else {
        for fi32 := 0; fi32 < 14; fi32++ {
            var v33 float64 = 3.14159
            _ = fi32
            _ = v33
        }
    }
    v30 = "a"
    for fi34 := 0; fi34 < 4; fi34++ {
        for fi35 := 0; fi35 < 13; fi35++ {
            switch (fi34 | (fi35 & (fi34 * 592))) {
            case 0:
                if true {
                    if len(v3) > 0 {
                        v3[0] = 2
                    }
                    switch 4294967295 {
                    case 0:
                        if ("foo bar" == ("a" + v30)) {
                            v5 -= (v4 / (((v5 | 255)) | 1))
                            if len(v3) > 0 {
                                v3[0] = 0
                            }
                            v5 = (v4 >> ((255) & 63))
                            var v36 int64 = 14
                            _ = v36
                        } else {
                            var v37 int64 = -245
                            if true {
                                var v38 float64 = 0.1
                                _ = v38
                            }
                            _ = v37
                        }
                    }
                } else {
                    v39 := v2[0]
                    var v40 float64 = 0.1
                    var v41 float64 = v40
                    switch fi34 {
                    case -84:
                        switch v30 {
                        case "\n\t":
                            var v42 bool = true
                            _ = v42
                        case "":
                            switch "hello" {
                            case "":
                                if true {
                                    var v43 int8 = -4
                                    _ = v43
                                }
                            }
                        default:
                            v40 = ((((-2.25 - (v40 - 0.1)) + float64(0.1)) / 1.5) * float64(1.5))
                        }
                    }
                    _ = v39
                    _ = v40
                    _ = v41
                }
                v1 = v1
                v44 := GenS2{F0: -9223372036854775808, F1: 0}
                _ = v44
            case 127:
                v3 = append(v3, 2147483647)
            case -837:
                _ = len(v2)
            default:
                if (945 >= 0) {
                    var v45 float32 = 1e10
                    _ = v45
                } else {
                    if true {
                        if true {
                            v46 := GenS8{F0: 141, F1: 32767, F2: -1}
                            fmt.Println(fi34, v4)
                            _ = v46
                        } else {
                            v47 := v2[2147483647]
                            _ = v47
                        }
                        v30 = ("hello" + (v30 + ""))
                    } else {
                        var v48 float32 = (float32((0.0 * 3.14159)) + (-2.25 - (0.0 * float32((0.0 - (0.1 / 1.5))))))
                        fmt.Println(v1)
                        _ = v48
                    }
                    if true {
                        if true {
                            var v49 bool = true
                            _ = v49
                        }
                        v1 = -834
                    } else {
                        v50 := GenS2{F0: -274, F1: 65535}
                        var v51 float32 = float32((1e10 * 3.14159))
                        _ = v50
                        _ = v51
                    }
                    if len(v3) > 0 {
                        v3[0] = 53
                    }
                    var v52 uint64 = 2
                    var v53 int16 = 1000
                    if len(v3) > 0 {
                        v3[0] = 2147483647
                    }
                    v54 := GenS2{F0: 9223372036854775807, F1: 0}
                    delete(v2, 542)
                    _ = v52
                    _ = v53
                    _ = v54
                }
            }
            v30 = v30
            var v55 float64 = 3.14159
            _ = fi35
            _ = v55
        }
        _ = fi34
    }
    for fi56 := 0; fi56 < 4; fi56++ {
        _ = len(v2)
        _ = fi56
    }
    fmt.Println(v5, v1, v5)
    if (v1 == -409) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch (v30 + v30) {
        case "foo bar":
            var v57 uint8 = 127
            _ = v57
        case "a":
            var v58 int64 = 42
            _ = v58
        case "":
            var v59 bool = (true && false)
            _ = v59
        }
    } else {
        var v60 uint16 = 1
        if true {
            delete(v2, 4294967295)
        } else {
            delete(v2, 0)
            v60 = 65535
        }
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
