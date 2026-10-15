package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int16 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 string; F1 uint; F2 string }
type GenS4 struct { F0 string; F1 uint32 }
func genFn0() uint16 {
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: 197, F1: 255, F2: 32767}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 1, F1: 936, F2: -919}
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 uint = 255
        _ = fi9
        _ = v10
    }
    if false {
        if true {
            if false {
                if true {
                    if true {
                        v11 := []int{32767}
                        _ = v11
                    } else {
                        var v12 int16 = 1
                        v12 -= (v12 >> ((975) & 63))
                        switch "foo bar" {
                        case "hello":
                            if true {
                                v13 := []int{}
                                var v14 uint32 = 602
                                var v15 float64 = 3.14159
                                _ = v13
                                _ = v14
                                _ = v15
                            } else {
                                for fi16 := 0; fi16 < 20; fi16++ {
                                    v17 := []int{}
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                    for fi18 := 0; fi18 < 1; fi18++ {
                                        if (-927 < 32767) {
                                            var v19 float32 = 0.1
                                            _ = v19
                                        }
                                        _ = fi18
                                    }
                                    _ = fi16
                                    _ = v17
                                }
                            }
                        case "foo bar":
                            v20 := []int{32768, 127, 9223372036854775807}
                            _ = v20
                        case "a":
                            v21 := GenS1{F0: -283, F1: 65535, F2: (v12 >> ((255) & 63))}
                            _ = v21
                        default:
                            var v22 uint8 = 42
                            _ = v22
                        }
                        _ = v12
                    }
                    var v23 int16 = -32768
                    _ = v23
                } else {
                    var v24 float64 = float64(0.1)
                    _ = v24
                }
                var v25 float32 = float32(0.1)
                v26 := GenS1{F0: -9223372036854775808, F1: 65535, F2: -32768}
                _ = v25
                _ = v26
            }
        }
        for fi27 := 0; fi27 < 7; fi27++ {
            switch fi27 {
            case -1:
                var v28 float32 = (float32(0.0) + float32(1.5))
                _ = v28
            case -45:
                var v29 uint8 = 0
                _ = v29
            }
            if (-644 != 1000) {
                var v30 uint16 = 42
                v31 := []int{}
                _ = v30
                _ = v31
            } else {
                if ((127 > 42) || (127 <= 100)) {
                    v32 := GenS1{F0: -440, F1: 2, F2: -1}
                    for fi33 := 0; fi33 < 9; fi33++ {
                        v34 := []int{fi33, -9223372036854775808}
                        _ = fi33
                        _ = v34
                    }
                    _ = v32
                } else {
                    var v35 int8 = 127
                    var v36 int32 = -873
                    _ = v35
                    _ = v36
                }
                var v37 uint32 = 138
                _ = v37
            }
            fmt.Println(fi27)
            _ = fi27
        }
    }
    if (true && (114 < 255)) {
        var v38 uint = 4294967295
        _ = v38
    }
    _ = v8
    return 464
}
func genFn1() int32 {
    var v39 int16 = 774
    switch -832 {
    case -9223372036854775808:
        var v40 bool = (128 > 42)
        _ = v40
    case -383:
        var v41 uint64 = 71
        _ = v41
    default:
        for fi42 := 0; fi42 < 8; fi42++ {
            var v43 uint32 = 0
            if true {
                v43 += 65535
                v43 = (^(v43))
            }
            _ = fi42
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 3; fi44++ {
        v45 := GenS2{F0: 294, F1: 167}
        if ((0 <= -112) || (0 != 4294967295)) {
            fmt.Println(fi44, fi44)
            fmt.Println(fi44)
            var v46 float32 = (-2.25 + -2.25)
            _ = v46
        }
        _ = fi44
        _ = v45
    }
    var v47 string = "hello"
    v39 -= 1000
    _ = v39
    _ = v47
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]uint32)
    v3 := []string{"foo bar", "a", ("\n\t" + "\n\t"), (("foo bar" + "hello") + "a")}
    var v4 uint8 = 42
    v5 := []uint32{4294967295, 42, 0, 0}
    var v48 int32 = -2147483648
    var v49 int8 = 105
    var v50 uint = 255
    v51 := GenS3{F0: "a", F1: uint(v4), F2: "a"}
    switch "a" {
    case "foo bar":
        for fi52 := 0; fi52 < 1; fi52++ {
            v4 = (v4 % ((v4) | 1))
            fmt.Println(fi52, fi52)
            _ = fi52
        }
    }
    var v53 uint = v50
    if (0 <= 0) {
        v50 += 255
        if true {
            v50 = 18446744073709551615
            if false {
                v54 := GenS3{F0: "hello", F1: (v53 * (v53 * (v53 - (v50 * v50)))), F2: ""}
                _ = v54
            }
        }
    }
    switch "" {
    case "foo bar":
        if true {
            for fi55 := 0; fi55 < 11; fi55++ {
                switch "a" {
                case "hello":
                    if len(v5) > 0 {
                        v5[0] = 2147483647
                    }
                case "foo bar":
                    var v56 string = "a"
                    _ = v56
                }
                _ = fi55
            }
        } else {
            var v57 uint16 = 75
            _ = v57
        }
    case "hello":
        for fi58 := 0; fi58 < 1; fi58++ {
            if (18446744073709551615 == 1) {
                v59 := GenS1{F0: -9223372036854775808, F1: 0, F2: -32768}
                _ = v59
            } else {
                if (169 >= 2147483647) {
                    v60 := GenS1{F0: 170, F1: v53, F2: -1}
                    _ = v60
                } else {
                    fmt.Println(v49, v49, v48)
                    v61 := GenS4{F0: "\n\t", F1: 65535}
                    _ = v61
                }
                v62 := GenS3{F0: "", F1: (v53 - 0), F2: ((("hello" + "\n\t") + "") + (("foo bar" + "a") + "\n\t"))}
                _ = v62
            }
            _ = fi58
        }
    case "":
        switch "a" {
        case "\n\t":
            for fi63 := 0; fi63 < 7; fi63++ {
                for fi64 := 0; fi64 < 3; fi64++ {
                    v65 := GenS4{F0: (("" + ((((((("hello" + "a") + "foo bar") + "a") + "") + "") + (("hello" + "hello") + "")) + ("foo bar" + "hello"))) + ""), F1: 42}
                    _ = fi64
                    _ = v65
                }
                fmt.Println(v48)
                _ = fi63
            }
        case "hello":
            if (!(4294967295 >= 0)) {
                var v66 int16 = 1000
                if (42 < 541) {
                    var v67 bool = true
                    _ = v67
                } else {
                    fmt.Println(v50, v49)
                    var v68 uint64 = 42
                    switch int(v68) {
                    case -2147483648:
                        switch 2147483647 {
                        case -2147483648:
                            fmt.Println(v68, v50, v49)
                        }
                    case 579:
                        switch "hello" {
                        case "\n\t":
                            var v69 string = ""
                            _ = v69
                        case "":
                            var v70 float64 = 0.1
                            _ = v70
                        case "foo bar":
                            v68 -= v68
                        default:
                            var v71 bool = (int8(v53) <= -32)
                            _ = v71
                        }
                    default:
                        var v72 float64 = 1.5
                        _ = v72
                    }
                    _ = v68
                }
                _ = v66
            }
        case "foo bar":
            var v73 bool = true
            _ = v73
        default:
            if true {
                var v74 int64 = 0
                v53 -= (^(v50))
                _ = v74
            }
        }
    default:
        if true {
            if len(v5) > 0 {
                v5[0] = 65535
            }
            fmt.Println(v53)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v53
}
