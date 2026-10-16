package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int16; F1 int8; F2 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS7 struct { F0 int16; F1 int64; F2 uint64 }
func genFn0(p0 string) int16 {
    p0 = "hello"
    if (!false) {
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint64 = 865
            _ = fi7
            _ = v8
        }
    } else {
        var v9 int32 = -1
        _ = v9
    }
    p0 = "\n\t"
    if (!(0 <= 538)) {
        if (!true) {
            if false {
                var v10 int8 = 42
                var v11 float32 = (float32(3.14159) / 1.0)
                _ = v10
                _ = v11
            } else {
                var v12 string = p0
                v13 := []int{203, 1}
                _ = v12
                _ = v13
            }
            v14 := GenS1{F0: 65535, F1: 0}
            _ = v14
        } else {
            fmt.Println(p0)
        }
        for fi15 := 0; fi15 < 10; fi15++ {
            for fi16 := 0; fi16 < 6; fi16++ {
                if false {
                    p0 = "hello"
                    fmt.Println(fi16)
                }
                fmt.Println(fi16)
                _ = fi16
            }
            if (true || false) {
                if true {
                    var v17 int16 = -32768
                    p0 = ("" + "foo bar")
                    _ = v17
                } else {
                    if (false || ((1e10 - 1.5) <= 3.14159)) {
                        fmt.Println(fi15, fi15)
                    }
                    fmt.Println(fi15, p0)
                    if (-2.25 > ((0.0 - float32(((1e10 + 0.0) * -2.25))) + 0.1)) {
                        p0 = p0
                    }
                }
                v18 := GenS2{F0: 884, F1: -1}
                _ = v18
            } else {
                if (false && (((2 < 0) || ((p0 + p0) > "foo bar")) || ((42 == 42) || ((1e10 != 0.0) || (false && (!false)))))) {
                    var v19 bool = false
                    var v20 int = 32768
                    v20 -= v20
                    _ = v19
                    _ = v20
                }
                fmt.Println(fi15)
            }
            v21 := GenS2{F0: 0, F1: 9223372036854775807}
            _ = fi15
            _ = v21
        }
        fmt.Println(p0)
    }
    var v22 int64 = -1
    _ = p0
    _ = v22
    return 42
}
func genFn1() int16 {
    var v23 uint16 = 1
    v23 = 0
    var v24 uint64 = 18446744073709551615
    if (!(-9223372036854775808 != -9223372036854775808)) {
        var v25 float64 = 1e10
        _ = v25
    } else {
        v26 := []int{32768, -2147483648}
        _ = v26
    }
    _ = v23
    _ = v24
    return 42
}

func main() {
    v1 := []int{128}
    v2 := make(map[int]uint32)
    v3 := []string{}
    v4 := make(map[string]bool)
    v5 := make(map[uint32]bool)
    var v6 string = "a"
    for fi27 := 0; fi27 < 3; fi27++ {
        v28 := GenS2{F0: 42, F1: -652}
        var v29 uint8 = 128
        _ = fi27
        _ = v28
        _ = v29
    }
    if (3.14159 <= 3.14159) {
        for fi30 := 0; fi30 < 8; fi30++ {
            var v31 float64 = 0.0
            switch 1 {
            case 32767:
                switch (^(fi30)) {
                case 127:
                    for fi32 := 0; fi32 < 7; fi32++ {
                        v4[v6] = false
                        _ = fi32
                    }
                default:
                    switch fi30 {
                    case 733:
                        var v33 int32 = 1
                        _ = v33
                    default:
                        var v34 uint16 = 47
                        _ = v34
                    }
                }
            case -9223372036854775808:
                var v35 uint8 = 0
                _ = v35
            case 9223372036854775807:
                v6 = v6
            default:
                v31 -= (v31 + -2.25)
            }
            if false {
                fmt.Println(fi30, v6, v31)
            }
            _ = fi30
            _ = v31
        }
        var v36 string = ("foo bar" + (v6 + v6))
        fmt.Println(v6)
        _ = v36
    } else {
        switch 0 {
        case -737:
            v37 := GenS1{F0: 32767, F1: 42}
            _ = v37
        case -858:
            v6 = v6
        }
        v6 = (("" + ("a" + (("foo bar" + v6) + "hello"))) + (v6 + ("" + "foo bar")))
        var v38 int8 = 1
        _ = v38
    }
    if (!(32768 > 9223372036854775807)) {
        v39 := GenS3{F0: 0, F1: 42, F2: -110}
        switch 127 {
        case 0:
            fmt.Println(v6)
        case 197:
            if ("a" > (v6 + (((v6 + v6) + ("" + (v6 + "\n\t"))) + "\n\t"))) {
                if false {
                    _ = len(v2)
                } else {
                    delete(v2, -1)
                    v6 = (v6 + v6)
                    v6 = v6
                }
            }
        case 127:
            v40 := GenS3{F0: -32768, F1: 42, F2: 36}
            _ = v40
        }
        var v41 uint64 = 133
        _ = v39
        _ = v41
    }
    var v42 uint8 = 0
    for fi43 := 0; fi43 < 4; fi43++ {
        var v44 uint16 = 72
        if (-73 == 40) {
            v45 := GenS4{F0: v6}
            var v46 uint = 1
            if (3.14159 >= (float32(1e10) * 1.5)) {
                v47 := GenS5{F0: 8, F1: 42}
                _ = v47
            } else {
                fmt.Println(fi43, v6)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = v45
            _ = v46
        } else {
            if true {
                v48 := GenS1{F0: v44, F1: 0}
                v42 += 127
                v49 := GenS4{F0: v6}
                _ = v48
                _ = v49
            } else {
                v44 -= v44
            }
            v3 = append(v3, v6)
        }
        if false {
            var v50 bool = (32767 != genFn1())
            v50 = v50
            _ = v50
        } else {
            switch 470 {
            case 692:
                if ((fi43 * (fi43 ^ 128)) != fi43) {
                    for fi51 := 0; fi51 < 5; fi51++ {
                        v44 += v44
                        _ = fi51
                    }
                } else {
                    fmt.Println(fi43)
                    v42 -= 42
                    v52 := GenS6{F0: 65535, F1: 1, F2: 4294967295}
                    _ = v52
                }
            case 32767:
                v53 := GenS7{F0: 32767, F1: -506, F2: uint64(v44)}
                _ = v53
            }
        }
        _ = fi43
        _ = v44
    }
    var v54 int8 = -103
    switch 4294967295 {
    case 32768:
        v1 = append(v1, 2147483647)
    case 42:
        for fi55 := 0; fi55 < 5; fi55++ {
            fmt.Println(fi55, fi55)
            switch "\n\t" {
            case "hello":
                fmt.Println(v54)
            case "foo bar":
                var v56 uint16 = 32767
                _ = v56
            case "\n\t":
                var v57 uint16 = 255
                _ = v57
            }
            var v58 int16 = 0
            _ = fi55
            _ = v58
        }
    }
    if (((32 == 42) && (2147483647 == 42)) || false) {
        switch (v6 + v6) {
        case "\n\t":
            if (!false) {
                v59 := v5[4294967295]
                _ = v59
            }
        case "a":
            for fi60 := 0; fi60 < 18; fi60++ {
                fmt.Println(v6, v42, fi60)
                _ = fi60
            }
        default:
            switch "hello" {
            case "hello":
                if (!true) {
                    var v61 int16 = 1000
                    var v62 int32 = 383
                    if true {
                        v61 -= (-(v61))
                        v54 = v54
                        v42 -= (v42 >> ((0) & 63))
                    } else {
                        var v63 uint = 42
                        _ = v63
                    }
                    _ = v61
                    _ = v62
                } else {
                    for fi64 := 0; fi64 < 11; fi64++ {
                        for fi65 := 0; fi65 < 18; fi65++ {
                            fmt.Println(v6, v54, fi64)
                            fmt.Println(v42, fi65, v42)
                            _ = fi65
                        }
                        for fi66 := 0; fi66 < 11; fi66++ {
                            v67 := GenS3{F0: 32767, F1: v54, F2: (-(v54))}
                            if (false || true) {
                                v54 -= -18
                                var v68 float64 = 0.0
                                _ = v68
                            } else {
                                var v69 uint8 = v42
                                v6 = "\n\t"
                                var v70 int = fi66
                                _ = v69
                                _ = v70
                            }
                            fmt.Println(fi64, v42)
                            _ = fi66
                            _ = v67
                        }
                        _ = fi64
                    }
                    v71 := GenS5{F0: 4294967295, F1: 1000}
                    for fi72 := 0; fi72 < 14; fi72++ {
                        var v73 int32 = -2147483648
                        _ = fi72
                        _ = v73
                    }
                    _ = v71
                }
            }
        }
        switch ((v6 + v6) + v6) {
        case "a":
            var v74 uint32 = 629
            _ = v74
        case "hello":
            fmt.Println(v6, v54, v42)
        default:
            for fi75 := 0; fi75 < 1; fi75++ {
                if (674 < 18446744073709551615) {
                    if (117 < 0) {
                        v42 -= 2
                    } else {
                        if false {
                            v6 = v6
                            switch v6 {
                            case "foo bar":
                                v42 -= 255
                            case "\n\t":
                                _ = len(v3)
                            }
                            v6 = v6
                        }
                    }
                } else {
                    v54 = (^(v54))
                    fmt.Println(v54, v6)
                }
                _ = fi75
            }
        }
        v54 = -128
    }
    var v76 int = 32768
    var v77 float64 = ((-2.25 * 3.14159) / (3.14159 * -2.25))
    var v78 uint32 = 758
    var v79 int32 = 327
    v76 = v76
    v80 := GenS6{F0: 42, F1: 891, F2: 4294967295}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v54
    _ = v6
    _ = v76
    _ = v77
    _ = v78
    _ = v79
    _ = v80
}
