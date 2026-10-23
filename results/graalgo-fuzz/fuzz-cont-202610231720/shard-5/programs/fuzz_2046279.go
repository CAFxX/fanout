package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 uint16 }

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -635
    var v3 int64 = 262
    for fi4 := 0; fi4 < 20; fi4++ {
        v2 = 0
        var v5 float64 = (-2.25 + 1.5)
        if false {
            v6 := []int{fi4, fi4}
            v7 := GenS1{F0: (v2 + v2)}
            _ = v6
            _ = v7
        } else {
            v2 += v2
            switch (^(fi4)) {
            case 42:
                if (0 != 396) {
                    var v8 int64 = v3
                    _ = v8
                }
            case -1:
                v2 -= (v2 | (v2 >> ((65535) & 63)))
            case -246:
                if ((v2 % (((v2 - v2)) | 1)) > (v2 / ((v2) | 1))) {
                    v9 := GenS1{F0: -1}
                    _ = v9
                }
            default:
                switch (fi4 | (fi4 - fi4)) {
                case 0:
                    var v10 bool = true
                    _ = v10
                case 305:
                    var v11 uint64 = 1
                    _ = v11
                }
            }
        }
        _ = fi4
        _ = v5
    }
    if false {
        v12 := v1[9223372036854775807]
        for fi13 := 0; fi13 < 4; fi13++ {
            v14 := GenS1{F0: v2}
            var v15 uint64 = 4294967295
            v12 = true
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = v12
    } else {
        v16 := GenS2{F0: 9223372036854775807, F1: ("hello" + "")}
        v2 -= v2
        _ = v16
    }
    delete(v1, 1)
    fmt.Println(v2)
    for fi17 := 0; fi17 < 8; fi17++ {
        for fi18 := 0; fi18 < 14; fi18++ {
            var v19 uint16 = 950
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    fmt.Println(v2, v3)
    v20 := []int{2147483647, 9223372036854775807}
    v21 := GenS2{F0: 2147483647, F1: "a"}
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v3, v2)
    }
    var v22 uint = 156
    var v23 int = int(v22)
    for fi24 := 0; fi24 < 16; fi24++ {
        for fi25 := 0; fi25 < 2; fi25++ {
            if false {
                fmt.Println(v23, fi24)
                var v26 uint16 = 32767
                v26 = v26
                if false {
                    v2 -= 0
                    var v27 int8 = -128
                    fmt.Println(v2, fi25)
                    _ = v27
                }
                var v28 int64 = (v3 * (v3 / ((v3) | 1)))
                _ = v26
                _ = v28
            }
            v23 += (fi24 + v23)
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            _ = fi25
        }
        for fi29 := 0; fi29 < 17; fi29++ {
            v23 -= fi29
            if false {
                if ((v3 & -942) > v3) {
                    v2 = 42
                    v30 := GenS3{F0: 18446744073709551615, F1: v22}
                    v22 = 0
                    var v31 int64 = -9223372036854775808
                    _ = v30
                    _ = v31
                } else {
                    switch (("a" + (("foo bar" + "a") + "hello")) + (("hello" + "foo bar") + ("a" + ("hello" + "foo bar")))) {
                    case "foo bar":
                        switch v23 {
                        case -612:
                            if (0 == 42) {
                                var v32 int64 = (v3 ^ (v3 | int64(v2)))
                                _ = v32
                            } else {
                                fmt.Println(v3, fi24, fi24)
                                if (!((1e10 + 1.5) != (float64(0.0) * 0.1))) {
                                    switch v23 {
                                    case -1:
                                        v2 += (v2 / (((-(v2))) | 1))
                                    case 9223372036854775807:
                                        var v33 uint16 = 861
                                        _ = v33
                                    }
                                    v23 = (fi29 ^ fi29)
                                } else {
                                    var v34 int64 = (v3 + (v3 << ((18446744073709551615) & 63)))
                                    _ = v34
                                }
                            }
                            var v35 uint8 = 255
                            if false {
                                v1[fi29] = (42 != 0)
                                switch 194 {
                                case 32767:
                                    v2 = v2
                                case -627:
                                    v36 := GenS4{F0: 2, F1: 75}
                                    _ = v36
                                }
                                v22 = 4294967295
                            }
                            _ = v35
                        case 2147483647:
                            if ((v22 << ((v22) & 63)) >= (v22 & 18446744073709551615)) {
                                fmt.Println(v2, fi29, v22)
                                switch "foo bar" {
                                case "\n\t":
                                    if (2147483647 > 65535) {
                                        v22 += (^(v22))
                                        v37 := GenS2{F0: (-(fi29)), F1: "\n\t"}
                                        v38 := GenS4{F0: 352, F1: 100}
                                        _ = v37
                                        _ = v38
                                    }
                                case "foo bar":
                                    var v39 bool = (-2.25 >= 0.0)
                                    _ = v39
                                default:
                                    v23 = fi29
                                }
                                v40 := GenS4{F0: 203, F1: 42}
                                _ = v40
                            } else {
                                delete(v1, int(v3))
                                if false {
                                    switch ("foo bar" + ("hello" + "\n\t")) {
                                    case "a":
                                        var v41 int16 = 0
                                        _ = v41
                                    case "hello":
                                        fmt.Println(v3)
                                        v42 := GenS5{F0: 1}
                                        _ = v42
                                    case "foo bar":
                                        if false {
                                            v20 = append(v20, fi24)
                                        } else {
                                            var v43 string = "foo bar"
                                            _ = v43
                                        }
                                    default:
                                        v23 += 126
                                    }
                                    v22 = v22
                                    v2 = (v2 + v2)
                                }
                            }
                        }
                    case "a":
                        var v44 int64 = v3
                        _ = v44
                    case "hello":
                        switch "hello" {
                        case "foo bar":
                            v45 := GenS3{F0: 943, F1: 42}
                            _ = v45
                        case "a":
                            v2 += 42
                        case "\n\t":
                            switch fi29 {
                            case 4294967295:
                                v22 = v22
                            }
                        default:
                            if (!false) {
                                if false {
                                    if ((1e10 >= 0.0) && (v22 <= v22)) {
                                        if len(v20) > 0 {
                                            v20[0] = (fi29 | (fi24 >> ((4294967295) & 63)))
                                        }
                                        var v46 uint = v22
                                        if false {
                                            switch ((("hello" + "a") + ("" + "")) + "\n\t") {
                                            case "foo bar":
                                                fmt.Println(v2, v22)
                                            case "a":
                                                v47 := GenS1{F0: v2}
                                                _ = v47
                                            }
                                        } else {
                                            v46 -= v46
                                        }
                                        _ = v46
                                    }
                                    v20 = append(v20, 1)
                                    var v48 int32 = (^(v2))
                                    _ = v48
                                }
                                if (1 >= v3) {
                                    var v49 int32 = v2
                                    switch fi24 {
                                    case 206:
                                        var v50 float64 = -2.25
                                        _ = v50
                                    default:
                                        v22 = 42
                                    }
                                    _ = v49
                                } else {
                                    var v51 int = fi29
                                    if false {
                                        v2 += v2
                                        if true {
                                            var v52 int32 = v2
                                            v51 = (^(fi29))
                                            var v53 int32 = 42
                                            _ = v52
                                            _ = v53
                                        } else {
                                            var v54 string = "\n\t"
                                            _ = v54
                                        }
                                        var v55 int32 = (v2 << ((v22) & 63))
                                        _ = v55
                                    } else {
                                        var v56 int32 = v2
                                        _ = v56
                                    }
                                    _ = v51
                                }
                            } else {
                                v57 := GenS1{F0: v2}
                                _ = v57
                            }
                        }
                    }
                    fmt.Println(v23)
                }
                v58 := GenS3{F0: 228, F1: (v22 % ((1) | 1))}
                v59 := GenS5{F0: 2}
                _ = v58
                _ = v59
            } else {
                fmt.Println(fi24)
            }
            var v60 uint64 = 18446744073709551615
            _ = fi29
            _ = v60
        }
        _ = fi24
    }
    var v61 int8 = 3
    var v62 float64 = 0.0
    switch "\n\t" {
    case "hello":
        var v63 uint = 68
        _ = v63
    default:
        if (32767 == 955) {
            v62 += (v62 + float64(1e10))
        }
    }
    v64 := GenS4{F0: 0, F1: v61}
    for fi65 := 0; fi65 < 8; fi65++ {
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            for fi66 := 0; fi66 < 16; fi66++ {
                if len(v20) > 0 {
                    v20[0] = (fi65 / ((32767) | 1))
                }
                var v67 uint8 = 2
                if (65535 > v22) {
                    v68 := GenS4{F0: 0, F1: (-(v61))}
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    _ = v68
                }
                _ = fi66
                _ = v67
            }
        case "\n\t":
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
        }
        _ = fi65
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v61
    _ = v62
    _ = v64
}
