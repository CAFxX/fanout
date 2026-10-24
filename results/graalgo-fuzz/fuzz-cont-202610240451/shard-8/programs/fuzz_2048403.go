package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint8; F1 uint; F2 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 32767
    var v3 int16 = 1000
    var v4 float64 = -2.25
    var v5 bool = true
    var v6 int8 = 1
    var v7 int16 = (^(v2))
    v8 := GenS1{F0: 39}
    _ = len(v1)
    var v9 string = ("hello" + "hello")
    if (-105 >= 32767) {
        v2 -= -426
        var v10 int = -1
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := v1[v9]
            v5 = (!((v12 & (v12 << ((18446744073709551615) & 63))) < v12))
            var v13 uint32 = v12
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v10
    }
    if v5 {
        var v14 string = "hello"
        _ = v14
    } else {
        if false {
            var v15 int16 = 42
            _ = v15
        }
        v9 = v9
        fmt.Println(v6, v4, v9)
    }
    v16 := []int{0, 127}
    var v17 int64 = 9223372036854775807
    if true {
        var v18 uint32 = 879
        for fi19 := 0; fi19 < 12; fi19++ {
            v17 -= v17
            var v20 float32 = 3.14159
            for fi21 := 0; fi21 < 17; fi21++ {
                v22 := GenS1{F0: 255}
                _ = fi21
                _ = v22
            }
            _ = fi19
            _ = v20
        }
        if false {
            if (255 < 1) {
                if ((v5 || v5) && (304 < 263)) {
                    if (!false) {
                        switch (v9 + "foo bar") {
                        case "\n\t":
                            var v23 string = v9
                            _ = v23
                        }
                    }
                    if v5 {
                        if (802 == 18446744073709551615) {
                            if true {
                                v7 = (v2 % (((v2 >> ((255) & 63))) | 1))
                            }
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                            for fi24 := 0; fi24 < 9; fi24++ {
                                v25 := GenS1{F0: 129}
                                v1["foo bar"] = v18
                                _ = fi24
                                _ = v25
                            }
                        }
                    } else {
                        var v26 uint16 = 65535
                        var v27 string = v9
                        _ = v26
                        _ = v27
                    }
                } else {
                    fmt.Println(v3, v18, v7)
                    var v28 uint8 = 42
                    var v29 int8 = (v6 - v6)
                    _ = v28
                    _ = v29
                }
            }
            var v30 float32 = -2.25
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            _ = v30
        } else {
            v9 = "hello"
        }
        _ = v18
    } else {
        _ = len(v16)
        if v5 {
            for fi31 := 0; fi31 < 8; fi31++ {
                v7 += v7
                if false {
                    v9 = v9
                    if v5 {
                        fmt.Println(v17, v3, fi31)
                        if ((v17 ^ (v17 + 0)) < (^(v17))) {
                            if len(v16) > 0 {
                                v16[0] = -2147483648
                            }
                            v9 = (("hello" + (v9 + v9)) + (("\n\t" + "\n\t") + v9))
                            switch (v9 + ("hello" + v9)) {
                            case "a":
                                if v5 {
                                    for fi32 := 0; fi32 < 1; fi32++ {
                                        v33 := GenS1{F0: 128}
                                        var v34 int16 = v7
                                        v6 = v6
                                        if false {
                                            var v35 int64 = v17
                                            _ = v35
                                        }
                                        _ = fi32
                                        _ = v33
                                        _ = v34
                                    }
                                    var v36 float32 = 0.1
                                    _ = v36
                                } else {
                                    fmt.Println(v5)
                                    for fi37 := 0; fi37 < 7; fi37++ {
                                        var v38 uint = 1
                                        v39 := GenS2{F0: v3}
                                        switch fi37 {
                                        case 9223372036854775807:
                                            if v5 {
                                                switch "hello" {
                                                case "foo bar":
                                                    var v40 int32 = -637
                                                    _ = v40
                                                case "":
                                                    v41 := GenS3{F0: 299}
                                                    v42 := GenS3{F0: v38}
                                                    _ = v41
                                                    _ = v42
                                                default:
                                                    v2 = (v7 << (((v38 | (v38 * v38))) & 63))
                                                }
                                            } else {
                                                var v43 float32 = 1.5
                                                if (1000 < 65535) {
                                                    var v44 uint32 = 2147483648
                                                    if (!v5) {
                                                        if true {
                                                            fmt.Println(fi31)
                                                            if (!v5) {
                                                                var v45 float32 = v43
                                                                v46 := GenS3{F0: v38}
                                                                var v47 uint64 = 18446744073709551615
                                                                _ = v45
                                                                _ = v46
                                                                _ = v47
                                                            }
                                                        } else {
                                                            if true {
                                                                delete(v1, (v9 + v9))
                                                                if (int32(v38) != -2147483648) {
                                                                    v2 = int16(v38)
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        if false {
                                                                            v5 = v5
                                                                        }
                                                                    }
                                                                } else {
                                                                    fmt.Println(v7, v2)
                                                                    var v48 uint16 = 65535
                                                                    var v49 bool = ((v4 / -2.25) <= (v4 * 3.14159))
                                                                    _ = v48
                                                                    _ = v49
                                                                }
                                                                var v50 float32 = float32(v4)
                                                                var v51 int64 = v17
                                                                switch (-(fi37)) {
                                                                case -86:
                                                                    v43 = v43
                                                                case 1:
                                                                    fmt.Println(fi37)
                                                                case 42:
                                                                    v7 -= -1
                                                                default:
                                                                    if (2147483647 > -1) {
                                                                        switch fi37 {
                                                                        case 32768:
                                                                            v52 := GenS4{F0: 1, F1: uint8(v44)}
                                                                            _ = v52
                                                                        }
                                                                        v9 = ""
                                                                        v53 := GenS5{F0: 0, F1: 4294967295, F2: 2147483647}
                                                                        _ = v53
                                                                    } else {
                                                                        fmt.Println(v9)
                                                                        var v54 int16 = v7
                                                                        var v55 uint64 = 1
                                                                        _ = v54
                                                                        _ = v55
                                                                    }
                                                                }
                                                                var v56 string = ""
                                                                v50 += float32(v4)
                                                                _ = v50
                                                                _ = v51
                                                                _ = v56
                                                            }
                                                            fmt.Println(v2, v6)
                                                            v3 += (^(v2))
                                                            fmt.Println(v4, v38, v6)
                                                        }
                                                        v17 -= (v17 ^ 91)
                                                        fmt.Println(v17, v7)
                                                    }
                                                    if v5 {
                                                        var v57 uint32 = 4294967295
                                                        var v58 string = v9
                                                        v57 = 939
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                    _ = v44
                                                }
                                                _ = v43
                                            }
                                        case 682:
                                            var v59 uint64 = 18446744073709551615
                                            _ = v59
                                        case 128:
                                            var v60 int64 = -1
                                            _ = v60
                                        default:
                                            v5 = (v5 && (v5 || true))
                                        }
                                        _ = len(v16)
                                        _ = fi37
                                        _ = v38
                                        _ = v39
                                    }
                                }
                            default:
                                v4 = -2.25
                            }
                        } else {
                            if (128 != (fi31 / (((fi31 | fi31)) | 1))) {
                                _ = len(v1)
                                var v61 bool = (65535 != 532)
                                v2 += 32767
                                _ = v61
                            } else {
                                v62 := GenS1{F0: 127}
                                _ = len(v1)
                                _ = v62
                            }
                            for fi63 := 0; fi63 < 4; fi63++ {
                                var v64 int16 = 0
                                delete(v1, "\n\t")
                                _ = fi63
                                _ = v64
                            }
                            for fi65 := 0; fi65 < 7; fi65++ {
                                var v66 int8 = -68
                                v7 -= v2
                                _ = len(v1)
                                _ = fi65
                                _ = v66
                            }
                        }
                        for fi67 := 0; fi67 < 8; fi67++ {
                            v2 += v3
                            _ = fi67
                        }
                    }
                } else {
                    if v5 {
                        v7 = (v2 ^ v2)
                        var v68 int = -966
                        _ = v68
                    }
                }
                if (!v5) {
                    var v69 uint64 = 18446744073709551615
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    _ = v69
                }
                _ = fi31
            }
            var v70 int8 = v6
            _ = v70
        } else {
            for fi71 := 0; fi71 < 9; fi71++ {
                for fi72 := 0; fi72 < 14; fi72++ {
                    v16 = append(v16, fi72)
                    _ = fi72
                }
                var v73 string = ("\n\t" + v9)
                v74 := GenS2{F0: (v3 + v2)}
                _ = fi71
                _ = v73
                _ = v74
            }
        }
    }
    for fi75 := 0; fi75 < 10; fi75++ {
        var v76 int = fi75
        v4 = float64(3.14159)
        _ = fi75
        _ = v76
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
