package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 string }
type GenS2 struct { F0 string; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int64; F2 uint8 }
type GenS6 struct { F0 uint; F1 uint64 }
type GenS7 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS8 struct { F0 uint64 }
func genFn0(p0 uint8, p1 uint, p2 int8) uint16 {
    v5 := GenS1{F0: 42, F1: -9223372036854775808, F2: "\n\t"}
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 float64 = 0.1
        fmt.Println(p0)
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 1, F1: 42, F2: (("a" + "foo bar") + ("foo bar" + ("foo bar" + "")))}
    switch 973 {
    case 127:
        fmt.Println(p2, p2, p0)
    default:
        var v9 uint8 = p0
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return 255
}
func genFn1(p0 int8, p1 int16) bool {
    var v10 uint8 = 1
    fmt.Println(p1, p1, p1)
    var v11 float32 = -2.25
    for fi12 := 0; fi12 < 2; fi12++ {
        if true {
            var v13 uint32 = 65535
            _ = v13
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return false
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{}
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    v14 := GenS2{F0: ("a" + ("foo bar" + "hello")), F1: 0, F2: 0}
    var v15 bool = (112 > -2147483648)
    v16 := GenS3{F0: 206, F1: (4294967295 < 18446744073709551615)}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        v18 := GenS2{F0: "", F1: 338, F2: 255}
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: 32767, F1: 42, F2: "\n\t"}
    if v15 {
        var v20 uint32 = 2147483647
        var v21 int8 = 42
        _ = v20
        _ = v21
    } else {
        if (-675 < 2147483647) {
            if (!v15) {
                for fi22 := 0; fi22 < 3; fi22++ {
                    for fi23 := 0; fi23 < 3; fi23++ {
                        var v24 bool = (true && v15)
                        switch ("foo bar" + "a") {
                        case "":
                            fmt.Println(fi23, v24)
                        case "a":
                            v25 := v1[(fi22 | 9223372036854775807)]
                            _ = v25
                        case "foo bar":
                            switch (("\n\t" + "hello") + "foo bar") {
                            case "":
                                fmt.Println(fi22, fi22)
                            case "foo bar":
                                if (!(v24 && (!(fi22 <= -2147483648)))) {
                                    fmt.Println(v15)
                                    if (v24 || v15) {
                                        var v26 uint16 = genFn0(42, 235, 127)
                                        _ = v26
                                    } else {
                                        v15 = (1 == genFn0(42, 834, 100))
                                    }
                                } else {
                                    v27 := GenS3{F0: 364, F1: (1.5 > 0.0)}
                                    _ = v27
                                }
                            case "a":
                                fmt.Println(v24, fi22, v15)
                            }
                            if v15 {
                                v24 = false
                                switch "a" {
                                case "foo bar":
                                    if ((347 <= 42) || false) {
                                        var v28 float64 = 3.14159
                                        fmt.Println(v24)
                                        var v29 float64 = v28
                                        _ = v28
                                        _ = v29
                                    }
                                }
                            }
                        }
                        var v30 int64 = 0
                        _ = fi23
                        _ = v24
                        _ = v30
                    }
                    if v15 {
                        switch -1 {
                        case 32767:
                            v15 = (2147483647 != 1)
                        default:
                            if (2 >= 786) {
                                switch fi22 {
                                case 2147483647:
                                    var v31 int32 = -2147483648
                                    _ = v31
                                case -77:
                                    v2 = append(v2, 88)
                                default:
                                    fmt.Println(v15, v15)
                                }
                            } else {
                                if v15 {
                                    v15 = true
                                    fmt.Println(v15, v15)
                                    switch "\n\t" {
                                    case "foo bar":
                                        v15 = false
                                    }
                                } else {
                                    fmt.Println(fi22, v15)
                                }
                            }
                        }
                        var v32 uint64 = 243
                        _ = v32
                    } else {
                        for fi33 := 0; fi33 < 17; fi33++ {
                            fmt.Println(v15, v15, v15)
                            v34 := v1[fi33]
                            if false {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v35 uint16 = 497
                                _ = v35
                            } else {
                                v34 = fi22
                            }
                            _ = fi33
                            _ = v34
                        }
                        if ((!v15) || true) {
                            var v36 uint32 = 2147483648
                            switch (("foo bar" + "") + ("a" + "a")) {
                            case "\n\t":
                                v15 = v15
                            }
                            _ = v36
                        } else {
                            var v37 int16 = 42
                            _ = v37
                        }
                    }
                    _ = fi22
                }
                v3["hello"] = (!false)
            } else {
                switch 4294967295 {
                case 42:
                    v38 := GenS4{F0: v15}
                    _ = v38
                default:
                    v39 := GenS3{F0: 0, F1: v15}
                    _ = v39
                }
                if false {
                    switch 128 {
                    case 2147483647:
                        var v40 int8 = -63
                        _ = v40
                    case -748:
                        fmt.Println(v15)
                    }
                    v41 := GenS4{F0: true}
                    _ = v41
                }
            }
            v2 = append(v2, 2)
        } else {
            if (0 >= 412) {
                if (1 < 42) {
                    for fi42 := 0; fi42 < 12; fi42++ {
                        fmt.Println(v15, fi42)
                        fmt.Println(v15, fi42)
                        _ = fi42
                    }
                    v43 := GenS5{F0: 1, F1: -603, F2: 125}
                    _ = v43
                }
            }
            v15 = v15
        }
        v44 := GenS6{F0: 65535, F1: 390}
        v45 := GenS7{F0: 128, F1: -801, F2: 32767}
        _ = v44
        _ = v45
    }
    var v46 uint8 = 255
    var v47 float64 = float64(3.14159)
    var v48 int16 = -25
    delete(v3, "")
    if v15 {
        var v49 float32 = float32(v47)
        for fi50 := 0; fi50 < 15; fi50++ {
            fmt.Println(v48)
            v51 := GenS8{F0: 1}
            _ = fi50
            _ = v51
        }
        var v52 int = -9223372036854775808
        _ = v49
        _ = v52
    } else {
        for fi53 := 0; fi53 < 2; fi53++ {
            fmt.Println(v47, v15)
            _ = fi53
        }
        v2 = append(v2, 4294967295)
    }
    fmt.Println(v48, v46)
    v15 = (false && (v15 || (!v15)))
    for fi54 := 0; fi54 < 15; fi54++ {
        for fi55 := 0; fi55 < 4; fi55++ {
            var v56 uint64 = 42
            _ = fi55
            _ = v56
        }
        v57 := v4[fi54]
        _ = fi54
        _ = v57
    }
    fmt.Println(v47)
    if v15 {
        v58 := GenS4{F0: false}
        if len(v2) > 0 {
            v2[0] = 627
        }
        var v59 int = -2147483648
        _ = v58
        _ = v59
    } else {
        v4[0] = "hello"
        v46 += (^(v46))
        switch ("hello" + "") {
        case "":
            if (!(!((^(v46)) != v46))) {
                fmt.Println(v47, v46, v48)
            }
        case "a":
            var v60 uint = 1
            _ = v60
        default:
            v48 -= 27
        }
    }
    var v61 int64 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v61
}
