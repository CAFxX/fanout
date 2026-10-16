package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint, p2 string) int32 {
    v5 := GenS1{F0: 1, F1: 255, F2: 100}
    fmt.Println(p1, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}
func genFn1(p0 int64, p1 int, p2 string) int16 {
    var v6 bool = (true && (true || (65535 <= 0)))
    p0 -= -595
    var v7 uint32 = 785
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int64 = (p0 >> ((0) & 63))
        fmt.Println(v9, v9, v9)
        v6 = false
        _ = fi8
        _ = v9
    }
    v10 := []int{p1, (-(p1)), p1, p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v7
    return 1000
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"\n\t", "foo bar", ("a" + ""), ("" + "\n\t")}
    var v3 int32 = 42
    v4 := []string{}
    if (true && (0 == 492)) {
        v11 := GenS1{F0: 0, F1: 42, F2: -35}
        _ = v11
    }
    v12 := GenS2{F0: 32767, F1: 466, F2: genFn1(50, 42, (("" + ("a" + "")) + "hello"))}
    fmt.Println(v3)
    var v13 int = 618
    switch "foo bar" {
    case "hello":
        var v14 string = "hello"
        _ = v14
    case "":
        _ = len(v1)
    case "a":
        switch v13 {
        case 127:
            fmt.Println(v3)
        default:
            v13 -= v13
        }
    default:
        fmt.Println(v3, v3)
    }
    var v15 uint64 = 42
    fmt.Println(v3)
    var v16 string = ""
    for fi17 := 0; fi17 < 10; fi17++ {
        switch v16 {
        case "foo bar":
            if false {
                if ("a" > v16) {
                    delete(v1, 65535)
                }
                v13 = -528
                fmt.Println(fi17)
            } else {
                var v18 float64 = 1e10
                v19 := GenS1{F0: 32767, F1: 0, F2: 42}
                _ = v18
                _ = v19
            }
        case "a":
            for fi20 := 0; fi20 < 6; fi20++ {
                v3 -= v3
                switch v13 {
                case 391:
                    var v21 int16 = -32768
                    _ = v21
                case 127:
                    v15 = v15
                case 578:
                    if ((true || true) || (-2.25 >= 3.14159)) {
                        fmt.Println(v16, fi17, v3)
                    } else {
                        switch (v13 * fi17) {
                        case 9223372036854775807:
                            v22 := GenS1{F0: -1, F1: 977, F2: 1}
                            _ = v22
                        default:
                            v4 = append(v4, v16)
                            if (1 >= 437) {
                                delete(v1, 779)
                                v13 = fi20
                                v13 += (fi20 + -689)
                            } else {
                                var v23 int64 = 1
                                v13 = 128
                                _ = v23
                            }
                            v15 = v15
                        }
                        fmt.Println(v15, v13)
                        v24 := GenS2{F0: 65535, F1: 42, F2: 1000}
                        _ = v24
                    }
                }
                _ = fi20
            }
        case "\n\t":
            switch fi17 {
            case 32768:
                for fi25 := 0; fi25 < 15; fi25++ {
                    v13 -= (v13 >> ((2) & 63))
                    switch v16 {
                    case "\n\t":
                        var v26 uint32 = 849
                        _ = v26
                    case "":
                        if true {
                            if false {
                                var v27 bool = (!(-32768 < 32767))
                                _ = v27
                            }
                            if true {
                                v16 = v16
                                v28 := GenS2{F0: 1, F1: 2, F2: -32768}
                                v15 = 649
                                _ = v28
                            } else {
                                var v29 int8 = -1
                                v30 := v1[1]
                                v16 = v16
                                _ = v29
                                _ = v30
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            if false {
                                if false {
                                    v31 := GenS1{F0: -32768, F1: 18446744073709551615, F2: 127}
                                    _ = len(v1)
                                    switch "a" {
                                    case "\n\t":
                                        v13 += 0
                                    default:
                                        switch -9223372036854775808 {
                                        case 667:
                                            v32 := GenS2{F0: 0, F1: 329, F2: 42}
                                            _ = v32
                                        default:
                                            v16 = v16
                                        }
                                    }
                                    _ = v31
                                } else {
                                    v16 = "a"
                                }
                                if true {
                                    var v33 int8 = -128
                                    v34 := GenS2{F0: 32767, F1: 923, F2: 42}
                                    _ = len(v2)
                                    _ = v33
                                    _ = v34
                                } else {
                                    var v35 uint64 = (v15 * 18446744073709551615)
                                    _ = v35
                                }
                            } else {
                                var v36 int16 = 1000
                                _ = len(v4)
                                v36 = (^(v36))
                                var v37 uint = 976
                                _ = v36
                                _ = v37
                            }
                        }
                    case "a":
                        v15 = v15
                    default:
                        v38 := GenS1{F0: 42, F1: 507, F2: 100}
                        _ = v38
                    }
                    _ = fi25
                }
            }
        }
        if (true || true) {
            v39 := GenS3{F0: -904}
            _ = v39
        } else {
            for fi40 := 0; fi40 < 20; fi40++ {
                switch "\n\t" {
                case "foo bar":
                    var v41 uint16 = 600
                    _ = v41
                case "\n\t":
                    v16 = (v16 + v16)
                case "":
                    v13 = (^(fi17))
                default:
                    var v42 int = v13
                    _ = v42
                }
                v3 -= (v3 << ((255) & 63))
                _ = fi40
            }
            var v43 string = ("\n\t" + (("a" + v16) + v16))
            _ = v43
        }
        _ = fi17
    }
    if false {
        fmt.Println(v16, v13, v15)
        var v44 int16 = -1
        if ((!true) && (!false)) {
            for fi45 := 0; fi45 < 17; fi45++ {
                switch v13 {
                case 128:
                    for fi46 := 0; fi46 < 6; fi46++ {
                        if false {
                            v47 := GenS3{F0: (^(v3))}
                            var v48 uint32 = 1
                            v13 = (^(v13))
                            _ = v47
                            _ = v48
                        } else {
                            if true {
                                v16 = "foo bar"
                                var v49 int16 = (^(v44))
                                _ = v49
                            } else {
                                var v50 int32 = -836
                                var v51 uint8 = 0
                                v52 := GenS4{F0: 498}
                                _ = v50
                                _ = v51
                                _ = v52
                            }
                        }
                        v3 += v3
                        v53 := GenS5{F0: (v44 * v44), F1: ((("foo bar" + "\n\t") + (v16 + "")) + "foo bar")}
                        _ = fi46
                        _ = v53
                    }
                case 127:
                    v13 -= 1
                case 9223372036854775807:
                    var v54 float32 = -2.25
                    _ = v54
                default:
                    var v55 int16 = v44
                    _ = v55
                }
                for fi56 := 0; fi56 < 19; fi56++ {
                    v3 = (v3 << ((2) & 63))
                    _ = fi56
                }
                var v57 uint16 = 938
                _ = fi45
                _ = v57
            }
            v16 = ((v16 + "a") + v16)
        } else {
            v16 = (v16 + "hello")
        }
        _ = v44
    } else {
        v15 = (^(v15))
        v58 := GenS6{F0: 1}
        _ = v58
    }
    var v59 int8 = -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v59
}
