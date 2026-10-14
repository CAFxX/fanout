package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 bool; F2 int64 }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint64; F1 uint; F2 bool }

func main() {
    var v1 int8 = 42
    v2 := make(map[string]string)
    v3 := []uint32{589, 468, 0, 4294967295}
    v4 := []string{("foo bar" + "foo bar"), "a", "\n\t", ""}
    var v5 bool = true
    var v6 uint64 = 4294967295
    v6 -= v6
    v5 = v5
    switch "hello" {
    case "hello":
        v5 = (324 != 994)
    case "foo bar":
        var v7 string = "hello"
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 1; fi8++ {
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 int16 = -1
                v11 := GenS1{F0: 88}
                _ = len(v3)
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
    default:
        v6 = v6
    }
    var v12 int32 = -2147483648
    v13 := GenS1{F0: 98}
    switch 9223372036854775807 {
    case -1:
        var v14 float64 = (1.5 - (1.5 - (3.14159 + 1.5)))
        _ = v14
    case -9223372036854775808:
        v15 := GenS2{F0: 850, F1: 9223372036854775807}
        _ = v15
    case 32768:
        for fi16 := 0; fi16 < 14; fi16++ {
            fmt.Println(v5, v5, v1)
            if v5 {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v17 := GenS1{F0: 128}
                _ = v17
            } else {
                v5 = (42 < 2)
            }
            _ = fi16
        }
    }
    if (true || (v1 < v1)) {
        if v5 {
            var v18 uint8 = 87
            var v19 int8 = 1
            _ = v18
            _ = v19
        }
        v20 := GenS2{F0: 255, F1: 1}
        fmt.Println(v5)
        _ = v20
    }
    var v21 int16 = 32767
    v5 = v5
    switch 134 {
    case 32767:
        var v22 int16 = (v21 / ((v21) | 1))
        _ = v22
    case 42:
        v21 += v21
    default:
        if v5 {
            switch "a" {
            case "a":
                v5 = (!v5)
            default:
                fmt.Println(v1, v12, v6)
            }
        } else {
            v23 := GenS2{F0: 255, F1: -856}
            _ = v23
        }
    }
    v1 = v1
    for fi24 := 0; fi24 < 1; fi24++ {
        v25 := GenS1{F0: 1}
        switch ("\n\t" + (("a" + "a") + ("foo bar" + ("" + "foo bar")))) {
        case "\n\t":
            v6 += 18446744073709551615
        case "foo bar":
            if (33 < 127) {
                if (!(42 < (-(v1)))) {
                    var v26 int = (fi24 / (((fi24 - 2147483647)) | 1))
                    var v27 uint8 = 184
                    _ = v26
                    _ = v27
                } else {
                    fmt.Println(v12)
                    v12 = (v12 - v12)
                }
                var v28 int64 = 0
                if len(v4) > 0 {
                    v4[0] = ("a" + "\n\t")
                }
                _ = v28
            } else {
                v5 = (277 <= 65535)
                var v29 bool = (-301 > (v21 ^ (v21 ^ 32767)))
                for fi30 := 0; fi30 < 19; fi30++ {
                    fmt.Println(v1, v21, v29)
                    v12 = v12
                    _ = fi30
                }
                _ = v29
            }
        case "hello":
            switch (fi24 ^ 65) {
            case 32768:
                var v31 uint16 = 2
                _ = v31
            case 9223372036854775807:
                for fi32 := 0; fi32 < 19; fi32++ {
                    v12 += (v12 + (-(v12)))
                    _ = fi32
                }
            default:
                if (!(589 == 255)) {
                    delete(v2, "hello")
                } else {
                    v6 -= v6
                    for fi33 := 0; fi33 < 4; fi33++ {
                        v1 = -31
                        if ((fi24 + 2147483647) >= 4294967295) {
                            switch ("foo bar" + (("hello" + ("hello" + ("hello" + "hello"))) + ("foo bar" + ("foo bar" + "\n\t")))) {
                            case "foo bar":
                                if ("" != "a") {
                                    v2[(("foo bar" + "foo bar") + "foo bar")] = "foo bar"
                                    v1 = (^(v1))
                                }
                            case "":
                                var v34 int16 = (-(v21))
                                _ = v34
                            default:
                                v3 = append(v3, 4294967295)
                            }
                            var v35 uint64 = v6
                            _ = v35
                        }
                        _ = fi33
                    }
                    v1 -= v1
                }
            }
        }
        for fi36 := 0; fi36 < 6; fi36++ {
            if (fi24 >= (fi36 % ((32768) | 1))) {
                v12 = (v12 >> ((1) & 63))
                v6 -= 1
                switch 2147483647 {
                case 606:
                    fmt.Println(fi36)
                default:
                    if (v5 && v5) {
                        switch fi24 {
                        case 762:
                            var v37 uint16 = 42
                            _ = v37
                        }
                        switch "" {
                        case "a":
                            switch (fi24 >> ((1) & 63)) {
                            case 1:
                                _ = len(v2)
                            default:
                                if (!true) {
                                    v38 := GenS1{F0: 127}
                                    var v39 float32 = (1.5 * ((3.14159 - 0.0) + (0.1 * 0.1)))
                                    _ = v38
                                    _ = v39
                                } else {
                                    v40 := GenS3{F0: 228}
                                    _ = v40
                                }
                            }
                        case "\n\t":
                            var v41 float32 = (0.1 * 0.1)
                            _ = v41
                        case "":
                            v12 += v12
                            fmt.Println(v5, v21)
                            v42 := GenS4{F0: 173, F1: v5, F2: 42}
                            _ = v42
                        }
                        switch (-(fi36)) {
                        case -676:
                            var v43 float64 = 0.0
                            _ = v43
                        default:
                            switch fi36 {
                            case -437:
                                var v44 int64 = -1
                                _ = v44
                            }
                        }
                    } else {
                        if v5 {
                            var v45 int64 = 9223372036854775807
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v46 := GenS4{F0: 65535, F1: v5, F2: v45}
                            _ = v45
                            _ = v46
                        }
                        if false {
                            fmt.Println(v5, v1)
                        } else {
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v47 float32 = 1e10
                                switch "hello" {
                                case "hello":
                                    switch fi24 {
                                    case 128:
                                        if v5 {
                                            var v48 uint64 = 9
                                            var v49 int32 = (v12 * v12)
                                            fmt.Println(v5, fi24, v5)
                                            _ = v48
                                            _ = v49
                                        } else {
                                            var v50 int16 = (v21 + v21)
                                            v6 -= v6
                                            _ = v50
                                        }
                                    }
                                case "":
                                    v51 := GenS5{F0: v1, F1: 32767}
                                    _ = v51
                                case "\n\t":
                                    var v52 float64 = 1e10
                                    _ = v52
                                default:
                                    v3 = append(v3, 0)
                                    var v53 int32 = (^(v12))
                                    v3 = append(v3, 801)
                                    _ = v53
                                }
                                fmt.Println(v21, v12, v1)
                                if false {
                                    var v54 bool = v5
                                    _ = v54
                                } else {
                                    if v5 {
                                        v55 := GenS3{F0: 2}
                                        _ = v55
                                    } else {
                                        if false {
                                            var v56 bool = ((fi24 << ((942) & 63)) >= -114)
                                            _ = v56
                                        }
                                        var v57 uint32 = 2147483648
                                        _ = v57
                                    }
                                    var v58 string = ("a" + "\n\t")
                                    _ = v58
                                }
                                var v59 uint16 = 85
                                v1 -= v1
                                _ = v47
                                _ = v59
                            } else {
                                var v60 uint8 = 42
                                fmt.Println(fi24, v12)
                                v3 = append(v3, 4294967295)
                                _ = v60
                            }
                            var v61 int = 9223372036854775807
                            v6 = 42
                            _ = v61
                        }
                        delete(v2, (("a" + "\n\t") + ("foo bar" + ("hello" + "\n\t"))))
                        var v62 float32 = (float32(3.14159) - 0.1)
                        _ = v62
                    }
                }
            } else {
                v63 := GenS6{F0: 703, F1: 1, F2: v5}
                var v64 string = "\n\t"
                _ = v63
                _ = v64
            }
            v1 += v1
            v65 := GenS2{F0: 0, F1: -67}
            _ = fi36
            _ = v65
        }
        _ = fi24
        _ = v25
    }
    fmt.Println(v21, v1)
    var v66 int32 = (^(v12))
    for fi67 := 0; fi67 < 5; fi67++ {
        fmt.Println(v1, v1, v5)
        _ = fi67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
}
