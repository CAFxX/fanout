package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int64, p1 int32, p2 int) uint8 {
    var v5 int64 = p0
    for fi6 := 0; fi6 < 20; fi6++ {
        switch (p2 + (fi6 & (fi6 % (((p2 - 1)) | 1)))) {
        case -9223372036854775808:
            if true {
                var v7 int64 = 9223372036854775807
                var v8 uint8 = 128
                _ = v7
                _ = v8
            }
        case 406:
            v9 := GenS1{F0: 2}
            _ = v9
        case 42:
            if false {
                var v10 uint16 = 0
                p0 -= (v5 - -1)
                p0 -= (v5 / ((p0) | 1))
                _ = v10
            }
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 193
}

func main() {
    var v1 bool = (!true)
    v2 := []string{(("hello" + "") + ("a" + "\n\t")), "foo bar"}
    var v3 int32 = -654
    v4 := make(map[int]string)
    if (42 <= -427) {
        _ = len(v4)
    }
    v11 := GenS1{F0: 42}
    v12 := GenS2{F0: (42 > -1), F1: true}
    if v1 {
        v13 := GenS1{F0: 42}
        if (!false) {
            v3 -= (v3 | -2147483648)
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 int = (fi14 / ((fi14) | 1))
                switch (fi14 / ((v15) | 1)) {
                case 258:
                    var v16 uint = 65535
                    _ = v16
                case 0:
                    var v17 string = ("a" + ("foo bar" + "foo bar"))
                    _ = v17
                default:
                    var v18 uint64 = 763
                    _ = v18
                }
                if v1 {
                    var v19 uint16 = 0
                    var v20 int16 = 217
                    var v21 int32 = v3
                    _ = v19
                    _ = v20
                    _ = v21
                }
                _ = fi14
                _ = v15
            }
            var v22 uint = 18446744073709551615
            _ = v22
        } else {
            for fi23 := 0; fi23 < 11; fi23++ {
                switch "" {
                case "foo bar":
                    if v1 {
                        v1 = true
                    } else {
                        v24 := GenS2{F0: (-2.25 <= float32(-2.25)), F1: v1}
                        v3 = (^(v3))
                        v25 := GenS3{F0: 65535, F1: ("hello" + "a")}
                        _ = v24
                        _ = v25
                    }
                default:
                    var v26 string = "a"
                    _ = v26
                }
                var v27 int8 = 127
                if (1 > 42) {
                    var v28 int32 = (v3 / ((v3) | 1))
                    delete(v4, (^(fi23)))
                    _ = v28
                } else {
                    if (255 != 2) {
                        switch "hello" {
                        case "":
                            v1 = v1
                        case "hello":
                            _ = len(v2)
                        case "\n\t":
                            var v29 uint32 = 2
                            _ = v29
                        }
                        var v30 uint16 = 1
                        _ = v30
                    } else {
                        var v31 uint64 = 656
                        fmt.Println(v3, fi23)
                        switch (-(fi23)) {
                        case -553:
                            var v32 bool = ((32767 < 786) && v1)
                            _ = v32
                        }
                        _ = v31
                    }
                    fmt.Println(v1, v1)
                    switch -1 {
                    case 128:
                        if (v27 > v27) {
                            fmt.Println(fi23, v27)
                            for fi33 := 0; fi33 < 2; fi33++ {
                                v27 += v27
                                fmt.Println(fi33, v1, fi23)
                                v34 := GenS4{F0: 32767, F1: 114, F2: (-(v3))}
                                switch fi33 {
                                case -1:
                                    var v35 bool = ((72 > 128) && (!v1))
                                    _ = v35
                                case 570:
                                    v27 += -1
                                case 128:
                                    fmt.Println(fi23, v1, v3)
                                }
                                var v36 uint8 = genFn0(1, v3, (fi23 & fi33))
                                _ = fi33
                                _ = v34
                                _ = v36
                            }
                        } else {
                            var v37 uint64 = 633
                            if v1 {
                                var v38 int64 = 1
                                var v39 int8 = -53
                                _ = v38
                                _ = v39
                            } else {
                                for fi40 := 0; fi40 < 4; fi40++ {
                                    v1 = (1e10 <= 1.5)
                                    _ = fi40
                                }
                                v3 -= -149
                            }
                            if (-856 == -1) {
                                v41 := GenS4{F0: 42, F1: (v37 << ((18446744073709551615) & 63)), F2: (v3 * 2147483647)}
                                fmt.Println(v3)
                                v3 = (v3 ^ v3)
                                _ = v41
                            } else {
                                var v42 uint64 = (v37 & 0)
                                if (!v1) {
                                    switch "" {
                                    case "a":
                                        if v1 {
                                            delete(v4, 106)
                                        } else {
                                            v3 = (v3 & 2147483647)
                                        }
                                    case "hello":
                                        for fi43 := 0; fi43 < 17; fi43++ {
                                            v3 -= v3
                                            var v44 int16 = 42
                                            _ = fi43
                                            _ = v44
                                        }
                                    default:
                                        v3 -= (v3 << ((4294967295) & 63))
                                    }
                                    v3 = (v3 % (((v3 << ((255) & 63))) | 1))
                                }
                                v1 = true
                                _ = v42
                            }
                            _ = v37
                        }
                    default:
                        var v45 float64 = 1.5
                        _ = v45
                    }
                }
                _ = fi23
                _ = v27
            }
        }
        for fi46 := 0; fi46 < 6; fi46++ {
            if v1 {
                fmt.Println(v3, v1)
                switch "" {
                case "foo bar":
                    v4[-2147483648] = ((("hello" + "\n\t") + "") + "\n\t")
                }
                var v47 int32 = -655
                _ = v47
            } else {
                var v48 string = "foo bar"
                _ = v48
            }
            switch "foo bar" {
            case "":
                var v49 int16 = 1
                _ = v49
            }
            _ = len(v2)
            _ = fi46
        }
        _ = v13
    } else {
        v1 = (1 != 806)
    }
    var v50 string = "\n\t"
    var v51 uint32 = 42
    v52 := GenS3{F0: (v51 * v51), F1: ""}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if v1 {
        v53 := GenS5{F0: 122}
        if false {
            if ((v3 ^ (v3 << ((139) & 63))) >= v3) {
                v51 = (v51 * (v51 << ((0) & 63)))
            }
            v3 -= 780
        } else {
            for fi54 := 0; fi54 < 2; fi54++ {
                fmt.Println(v50, fi54, v1)
                v55 := GenS1{F0: 42}
                v56 := GenS2{F0: false, F1: (float32(-2.25) <= 0.1)}
                _ = fi54
                _ = v55
                _ = v56
            }
            var v57 int64 = 1
            switch -452 {
            case -2147483648:
                var v58 uint16 = 42
                _ = v58
            }
            _ = v57
        }
        _ = v53
    } else {
        var v59 bool = false
        _ = v59
    }
    v51 += v51
    v60 := GenS4{F0: 748, F1: 1, F2: v3}
    var v61 uint = 0
    delete(v4, 32768)
    var v62 int8 = 100
    v63 := GenS4{F0: 0, F1: 177, F2: v3}
    switch v50 {
    case "hello":
        v61 += (v61 & v61)
    case "":
        switch "hello" {
        case "a":
            var v64 int32 = v3
            _ = v64
        }
    }
    v65 := GenS4{F0: uint16(v61), F1: 0, F2: (v3 | (v3 << ((v61) & 63)))}
    v66 := GenS5{F0: v62}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v52
    _ = v60
    _ = v61
    _ = v62
    _ = v63
    _ = v65
    _ = v66
}
