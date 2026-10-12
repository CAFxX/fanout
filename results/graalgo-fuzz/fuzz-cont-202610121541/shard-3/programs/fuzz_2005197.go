package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int16; F1 int16; F2 int64 }
type GenS3 struct { F0 int8; F1 int; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 uint16; F1 uint32; F2 bool }
func genFn0(p0 uint8) uint8 {
    p0 = 128
    v4 := GenS1{F0: 2147483647, F1: -1}
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS2{F0: -723, F1: 32767, F2: -1}
        p0 = (p0 / ((p0) | 1))
        for fi7 := 0; fi7 < 1; fi7++ {
            fmt.Println(p0, p0, p0)
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    p0 = (p0 ^ p0)
    if (("\n\t" >= ("a" + ("foo bar" + "\n\t"))) && true) {
        v8 := GenS3{F0: -128, F1: 0, F2: (("a" + "a") + "a")}
        _ = v8
    } else {
        switch ("foo bar" + "\n\t") {
        case "hello":
            if false {
                v9 := GenS2{F0: 912, F1: -450, F2: 1}
                _ = v9
            } else {
                if ((float64((3.14159 * (float32(((0.1 - 0.0) - 0.0)) + -2.25))) / (1.5 * (0.1 * 0.1))) > 0.0) {
                    p0 = p0
                    p0 += (p0 - (p0 + p0))
                } else {
                    p0 -= p0
                    var v10 uint = 497
                    _ = v10
                }
            }
        case "foo bar":
            switch 636 {
            case 759:
                var v11 int = -9223372036854775808
                _ = v11
            default:
                var v12 uint16 = 538
                _ = v12
            }
        case "a":
            fmt.Println(p0)
        default:
            p0 = p0
        }
        var v13 int16 = 42
        _ = v13
    }
    _ = p0
    _ = v4
    return p0
}

func main() {
    var v1 uint64 = 1
    var v2 int8 = -1
    v3 := make(map[string]bool)
    v14 := []int{42, 874, 32767}
    v15 := GenS4{F0: 32767}
    _ = len(v14)
    var v16 int64 = -943
    for fi17 := 0; fi17 < 3; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            fmt.Println(v1)
            if true {
                var v19 int64 = v16
                if true {
                    if (!(1 <= 42)) {
                        if true {
                            if (189 > 49) {
                                fmt.Println(fi18, fi18, v19)
                            } else {
                                var v20 int32 = -785
                                _ = v20
                            }
                            _ = len(v14)
                        }
                        v16 = (v19 % ((-643) | 1))
                        v19 += v19
                    }
                    if false {
                        fmt.Println(v16, fi18)
                        v21 := GenS1{F0: -1, F1: 51}
                        if true {
                            v2 -= v2
                            if (0.0 == (3.14159 + 1.5)) {
                                var v22 uint32 = 4294967295
                                _ = v22
                            }
                            if (550 <= 0) {
                                var v23 int8 = -75
                                v24 := v3["foo bar"]
                                delete(v3, ("foo bar" + "\n\t"))
                                fmt.Println(fi18, v24)
                                _ = v23
                                _ = v24
                            } else {
                                v19 -= v19
                                if false {
                                    fmt.Println(fi17, fi17)
                                }
                            }
                            fmt.Println(fi17, v2, v16)
                            fmt.Println(fi17, v19, fi18)
                            var v25 uint32 = 65535
                            var v26 uint64 = v1
                            var v27 string = ("foo bar" + "a")
                            _ = v25
                            _ = v26
                            _ = v27
                        } else {
                            fmt.Println(v19, v19)
                            v2 = v2
                            var v28 float32 = 1.5
                            _ = v28
                        }
                        _ = v21
                    }
                    var v29 uint = 255
                    fmt.Println(fi17, v1)
                    var v30 int64 = (v19 % ((42) | 1))
                    _ = v29
                    _ = v30
                }
                _ = v19
            } else {
                var v31 int16 = 1
                v32 := GenS4{F0: (-(v31))}
                if false {
                    var v33 int32 = 0
                    _ = v33
                } else {
                    var v34 uint8 = genFn0(genFn0(128))
                    if false {
                        var v35 float64 = (-2.25 - 1.5)
                        var v36 uint16 = 77
                        delete(v3, "foo bar")
                        _ = v35
                        _ = v36
                    } else {
                        var v37 uint16 = 2
                        v16 = (v16 & int64(v37))
                        _ = v37
                    }
                    _ = v34
                }
                if len(v14) > 0 {
                    v14[0] = (fi17 ^ 32767)
                }
                fmt.Println(v16, v16)
                _ = v31
                _ = v32
            }
            var v38 int32 = -1
            switch fi17 {
            case 9223372036854775807:
                v1 -= v1
            }
            _ = fi18
            _ = v38
        }
        _ = fi17
    }
    var v39 uint16 = 961
    fmt.Println(v39)
    for fi40 := 0; fi40 < 6; fi40++ {
        switch "" {
        case "\n\t":
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        case "":
            var v41 float32 = 0.1
            _ = v41
        case "foo bar":
            if false {
                if (-1 < -1) {
                    var v42 float32 = (0.0 * ((0.0 + 3.14159) - float32(0.0)))
                    for fi43 := 0; fi43 < 19; fi43++ {
                        fmt.Println(v16)
                        fmt.Println(v1, fi40)
                        _ = fi43
                    }
                    _ = v42
                }
            } else {
                if true {
                    var v44 int64 = (v16 | 9223372036854775807)
                    v39 = v39
                    if (0.0 < (-2.25 - 1e10)) {
                        fmt.Println(v2, fi40, v44)
                    } else {
                        v45 := GenS5{F0: "foo bar"}
                        var v46 int8 = (v2 << ((1) & 63))
                        v47 := GenS6{F0: 1, F1: 1}
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    _ = v44
                }
            }
        }
        v48 := GenS7{F0: 172, F1: false}
        var v49 string = ""
        _ = fi40
        _ = v48
        _ = v49
    }
    if false {
        switch "a" {
        case "\n\t":
            if true {
                v16 -= v16
            } else {
                var v50 uint = 42
                fmt.Println(v16, v2)
                v51 := GenS1{F0: -1, F1: v16}
                _ = v50
                _ = v51
            }
        case "hello":
            var v52 uint = 0
            _ = v52
        case "foo bar":
            var v53 int16 = 0
            _ = v53
        default:
            if false {
                for fi54 := 0; fi54 < 4; fi54++ {
                    switch fi54 {
                    case 4294967295:
                        fmt.Println(v39, v1, v2)
                    case 32767:
                        var v55 int64 = v16
                        _ = v55
                    case 2147483647:
                        if false {
                            v2 = v2
                            var v56 float64 = (1e10 + 0.1)
                            _ = v56
                        } else {
                            fmt.Println(fi54, v39, fi54)
                            var v57 bool = (0.1 != ((float64((0.0 / 1.0)) + ((1e10 - (0.0 * 1.5)) + 1e10)) / float64((((float32(1e10) - -2.25) - 0.0) - 0.1))))
                            v58 := GenS1{F0: 42, F1: (^(v16))}
                            _ = v57
                            _ = v58
                        }
                    }
                    _ = fi54
                }
                var v59 uint64 = v1
                v14 = append(v14, -9223372036854775808)
                _ = v59
            }
        }
        _ = len(v3)
        _ = len(v3)
    }
    switch 1 {
    case 42:
        v60 := GenS8{F0: v39, F1: 1, F2: (-32768 == 104)}
        _ = v60
    case 32768:
        fmt.Println(v39, v16)
    default:
        v2 = (v2 ^ (v2 ^ v2))
    }
    var v61 float32 = ((0.1 / float32(-2.25)) + -2.25)
    v62 := GenS7{F0: 17, F1: (true || true)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v39
    _ = v61
    _ = v62
}
