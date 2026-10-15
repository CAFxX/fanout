package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint32; F1 int64; F2 string }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 uint16) uint {
    var v5 int16 = -957
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint16 = p0
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := []int{(fi6 % (((fi8 << ((2) & 63))) | 1))}
            if len(v9) > 0 {
                v9[0] = fi6
            }
            switch (("foo bar" + "") + "") {
            case "a":
                if false {
                    var v10 uint = 0
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    v11 := GenS1{F0: "\n\t", F1: "hello"}
                    switch fi8 {
                    case -1:
                        var v12 int = (fi6 << (((v10 + v10)) & 63))
                        _ = v12
                    case 631:
                        v9 = append(v9, fi6)
                    case 174:
                        v13 := GenS1{F0: "a", F1: "hello"}
                        _ = v13
                    }
                    if false {
                        v10 = 2
                        v7 += (p0 % ((uint16(v10)) | 1))
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    } else {
                        v10 -= (v10 >> ((v10) & 63))
                    }
                    var v14 int32 = 2147483647
                    v15 := GenS2{F0: uint32(v10), F1: 0, F2: (("\n\t" + "a") + "")}
                    var v16 int16 = -1
                    if false {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    } else {
                        v17 := GenS2{F0: 130, F1: int64(fi6), F2: "foo bar"}
                        v14 = -2147483648
                        switch "" {
                        case "\n\t":
                            p0 = (^(p0))
                        }
                        _ = v17
                    }
                    _ = v10
                    _ = v11
                    _ = v14
                    _ = v15
                    _ = v16
                }
            case "":
                var v18 uint8 = 0
                _ = v18
            case "\n\t":
                fmt.Println(p0, v5)
            }
            _ = fi8
            _ = v9
        }
        fmt.Println(fi6)
        _ = fi6
        _ = v7
    }
    var v19 int8 = -1
    v5 += (v5 / ((v5) | 1))
    switch (("hello" + "") + "") {
    case "foo bar":
        p0 += (p0 & p0)
    case "\n\t":
        v20 := []int{725, 32768, -366}
        _ = v20
    }
    _ = p0
    _ = v19
    _ = v5
    return 255
}
func genFn1(p0 int32, p1 uint64) uint {
    p1 = 613
    if (65535 <= 2) {
        p1 += p1
    } else {
        switch 127 {
        case 32768:
            for fi21 := 0; fi21 < 4; fi21++ {
                var v22 uint16 = 242
                _ = fi21
                _ = v22
            }
        case 9223372036854775807:
            p0 = (p0 ^ 2147483647)
        }
        for fi23 := 0; fi23 < 3; fi23++ {
            if false {
                fmt.Println(p1)
                for fi24 := 0; fi24 < 8; fi24++ {
                    p1 -= (^(p1))
                    if (true && ((4294967295 != p1) || true)) {
                        if (1 >= 73) {
                            switch ("\n\t" + ("\n\t" + "foo bar")) {
                            case "hello":
                                var v25 uint32 = 65535
                                _ = v25
                            case "a":
                                v26 := GenS3{F0: 32767, F1: 0}
                                _ = v26
                            default:
                                fmt.Println(fi23)
                            }
                        } else {
                            p0 += (p0 + p0)
                        }
                        var v27 int = (fi24 + (fi24 | fi24))
                        var v28 bool = (!(1 < int64(p1)))
                        if v28 {
                            var v29 int64 = 42
                            _ = v29
                        }
                        v30 := []int{(fi24 % (((fi23 - v27)) | 1)), (fi24 + (v27 << ((403) & 63)))}
                        _ = v27
                        _ = v28
                        _ = v30
                    }
                    _ = fi24
                }
                var v31 int16 = -32768
                _ = v31
            } else {
                p0 = 1
                for fi32 := 0; fi32 < 9; fi32++ {
                    p1 = p1
                    if (("foo bar" + "a") == "") {
                        var v33 uint32 = 4294967295
                        _ = v33
                    }
                    v34 := GenS4{F0: fi23}
                    p1 += p1
                    _ = fi32
                    _ = v34
                }
                if true {
                    p1 += 303
                }
            }
            switch "foo bar" {
            case "foo bar":
                fmt.Println(p1)
            }
            _ = fi23
        }
    }
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = (123 <= 72)
    v3 := make(map[uint32]int)
    v4 := []int{9223372036854775807, -2147483648, 930, 101}
    v35 := GenS4{F0: 0}
    var v36 uint32 = 747
    if (!(690 >= 2)) {
        v2 = (v36 == (v36 ^ 4294967295))
    }
    var v37 int32 = -2147483648
    for fi38 := 0; fi38 < 16; fi38++ {
        var v39 int16 = -818
        _ = fi38
        _ = v39
    }
    if v2 {
        var v40 int8 = -73
        _ = v40
    } else {
        if len(v4) > 0 {
            v4[0] = 127
        }
        var v41 uint16 = 480
        var v42 bool = (100 < 127)
        _ = v41
        _ = v42
    }
    v37 = -1
    var v43 uint = 0
    v36 = v36
    if false {
        v44 := GenS3{F0: 18, F1: 1}
        var v45 uint = genFn1((v37 + 2147483647), uint64(v43))
        v37 = (v37 >> ((743) & 63))
        _ = v44
        _ = v45
    }
    if ((1000 < 299) && (true && true)) {
        var v46 uint8 = 177
        switch 266 {
        case 1:
            var v47 int8 = 42
            _ = v47
        case 127:
            for fi48 := 0; fi48 < 8; fi48++ {
                v46 = 0
                _ = fi48
            }
        case 2147483647:
            v36 = (v36 + v36)
        }
        v49 := GenS2{F0: (^(v36)), F1: -9223372036854775808, F2: "foo bar"}
        _ = v46
        _ = v49
    }
    v36 = v36
    if false {
        for fi50 := 0; fi50 < 3; fi50++ {
            v51 := GenS4{F0: 128}
            var v52 bool = (v43 <= v43)
            _ = fi50
            _ = v51
            _ = v52
        }
        switch 42 {
        case 711:
            v53 := GenS5{F0: -152}
            _ = v53
        default:
            for fi54 := 0; fi54 < 13; fi54++ {
                v43 = 1
                fmt.Println(v37, v37)
                for fi55 := 0; fi55 < 6; fi55++ {
                    switch "hello" {
                    case "foo bar":
                        v43 = (^(v43))
                    case "":
                        if (1 == v43) {
                            var v56 uint32 = 399
                            switch ("" + "foo bar") {
                            case "":
                                var v57 int8 = -41
                                _ = v57
                            default:
                                switch fi54 {
                                case -1:
                                    var v58 int = 1
                                    _ = v58
                                case 127:
                                    fmt.Println(v37)
                                case 128:
                                    var v59 int32 = v37
                                    _ = v59
                                }
                            }
                            _ = v56
                        } else {
                            var v60 int64 = -990
                            _ = v60
                        }
                    case "a":
                        v61 := GenS6{F0: v36, F1: v36}
                        _ = v61
                    default:
                        v36 -= v36
                    }
                    _ = fi55
                }
                _ = fi54
            }
        }
        var v62 float64 = 3.14159
        _ = v62
    }
    var v63 int8 = 11
    for fi64 := 0; fi64 < 20; fi64++ {
        if v2 {
            fmt.Println(fi64)
            var v65 float32 = -2.25
            if len(v4) > 0 {
                v4[0] = (fi64 ^ (fi64 | 0))
            }
            _ = v65
        } else {
            var v66 uint32 = v36
            v63 = (^(v63))
            _ = v66
        }
        _ = fi64
    }
    switch -9223372036854775808 {
    case 127:
        v37 -= 2147483647
    }
    v67 := GenS7{F0: 232}
    fmt.Println(v37)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v43
    _ = v63
    _ = v67
}
