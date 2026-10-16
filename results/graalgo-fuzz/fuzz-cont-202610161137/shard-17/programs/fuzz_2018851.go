package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint8, p1 uint16) int8 {
    p1 -= 819
    var v5 uint8 = uint8(p1)
    _ = p0
    _ = p1
    _ = v5
    return -1
}
func genFn1() int8 {
    var v6 string = "a"
    if (false || true) {
        fmt.Println(v6)
    }
    v7 := []int{-9223372036854775808, 1, -1, 4294967295}
    v6 = "a"
    v8 := GenS1{F0: (-805 == 0)}
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"a", "\n\t"}
    v3 := make(map[int]string)
    v4 := make(map[int]string)
    v9 := GenS2{F0: "foo bar", F1: 255, F2: 0}
    if ((true || true) || false) {
        var v10 int = -2147483648
        _ = v10
    }
    if (-9223372036854775808 < -9223372036854775808) {
        var v11 string = "\n\t"
        if ((784 < -1) || (((1 == 0) || false) && false)) {
            v11 = (v11 + v11)
        }
        _ = v11
    } else {
        switch 0 {
        case -651:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 2147483647:
            switch "foo bar" {
            case "a":
                var v12 uint16 = 65535
                _ = v12
            case "":
                if false {
                    var v13 uint32 = 60
                    var v14 string = "foo bar"
                    _ = v13
                    _ = v14
                } else {
                    var v15 uint16 = 2
                    _ = v15
                }
            case "foo bar":
                if (42 != 814) {
                    v2 = append(v2, "a")
                    var v16 uint = 65535
                    _ = v16
                }
            }
        default:
            var v17 int32 = 0
            _ = v17
        }
        if (!((float64((1e10 * 0.0)) == 1.5) && (!(-874 == 0)))) {
            var v18 int16 = -32768
            var v19 bool = true
            if v19 {
                fmt.Println(v19, v18)
                v20 := v1[2]
                v20 = ((v20 + v20) + ((v20 + v20) + (v20 + ((v20 + (("\n\t" + v20) + v20)) + "hello"))))
                _ = v20
            }
            _ = v18
            _ = v19
        } else {
            var v21 int64 = 555
            if false {
                v22 := GenS1{F0: ((0.0 - 0.0) > 1e10)}
                v23 := GenS2{F0: "a", F1: 4294967295, F2: 2}
                v21 = v21
                _ = v22
                _ = v23
            } else {
                v21 -= (v21 << ((0) & 63))
                for fi24 := 0; fi24 < 5; fi24++ {
                    for fi25 := 0; fi25 < 9; fi25++ {
                        v26 := v4[(fi24 * int(v21))]
                        _ = fi25
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 19; fi27++ {
                        var v28 bool = ("a" != "hello")
                        if (0.0 <= 3.14159) {
                            if false {
                                var v29 int32 = -1
                                _ = v29
                            }
                            switch (fi24 | fi24) {
                            case 611:
                                if true {
                                    switch (("" + "") + (("\n\t" + "hello") + ("\n\t" + ("foo bar" + ("a" + "foo bar"))))) {
                                    case "":
                                        v30 := GenS1{F0: (!((v28 || (!v28)) || true))}
                                        _ = v30
                                    }
                                    var v31 uint8 = 31
                                    _ = v31
                                }
                            case 651:
                                if v28 {
                                    var v32 int16 = -321
                                    _ = v32
                                } else {
                                    if true {
                                        switch fi27 {
                                        case 739:
                                            if ((!(!v28)) || false) {
                                                switch 4294967295 {
                                                case -665:
                                                    fmt.Println(v21)
                                                case 32767:
                                                    switch "" {
                                                    case "a":
                                                        var v33 string = "a"
                                                        _ = v33
                                                    case "":
                                                        var v34 uint8 = 2
                                                        _ = v34
                                                    case "hello":
                                                        v28 = v28
                                                    }
                                                }
                                                if (754 <= 4294967295) {
                                                    switch (fi27 * (-(fi24))) {
                                                    case 32768:
                                                        if false {
                                                            v28 = (1 != 466)
                                                        }
                                                    case 4294967295:
                                                        var v35 int64 = (v21 << ((4294967295) & 63))
                                                        _ = v35
                                                    case 32767:
                                                        delete(v3, (fi24 | (fi24 % ((int(v21)) | 1))))
                                                    }
                                                    var v36 uint = 628
                                                    switch 32767 {
                                                    case -2147483648:
                                                        switch ("hello" + ("a" + "\n\t")) {
                                                        case "foo bar":
                                                            var v37 uint8 = 128
                                                            _ = v37
                                                        case "a":
                                                            var v38 int32 = -893
                                                            _ = v38
                                                        case "\n\t":
                                                            if v28 {
                                                                if v28 {
                                                                    if (723 <= -114) {
                                                                        var v39 uint8 = 41
                                                                        fmt.Println(v39, v28)
                                                                        fmt.Println(fi24)
                                                                        _ = v39
                                                                    }
                                                                    _ = len(v3)
                                                                }
                                                                v40 := GenS2{F0: "foo bar", F1: (v36 & v36), F2: 127}
                                                                if (true && v28) {
                                                                    delete(v3, fi27)
                                                                }
                                                                var v41 int = (fi27 >> (((v36 * 828)) & 63))
                                                                v42 := GenS3{F0: 771, F1: 1}
                                                                _ = v40
                                                                _ = v41
                                                                _ = v42
                                                            } else {
                                                                var v43 bool = true
                                                                v44 := GenS3{F0: uint16(v36), F1: 0}
                                                                _ = v43
                                                                _ = v44
                                                            }
                                                        default:
                                                            if (-2147483648 < 2147483647) {
                                                                v36 += 18446744073709551615
                                                            } else {
                                                                v45 := GenS3{F0: 964, F1: 2}
                                                                _ = v45
                                                            }
                                                        }
                                                    default:
                                                        v46 := GenS3{F0: 255, F1: 1}
                                                        _ = v46
                                                    }
                                                    _ = v36
                                                }
                                                switch (fi27 / ((fi24) | 1)) {
                                                case 532:
                                                    if false {
                                                        if (1 > 340) {
                                                            v21 -= v21
                                                            fmt.Println(v21, v28, v21)
                                                        }
                                                    }
                                                case 2147483647:
                                                    var v47 uint = 18446744073709551615
                                                    _ = v47
                                                case 127:
                                                    var v48 float64 = 3.14159
                                                    _ = v48
                                                default:
                                                    fmt.Println(fi27)
                                                }
                                                v49 := GenS4{F0: 0, F1: 702}
                                                _ = len(v4)
                                                _ = v49
                                            } else {
                                                var v50 uint32 = 249
                                                fmt.Println(fi27, v28, fi27)
                                                var v51 uint8 = 207
                                                _ = v50
                                                _ = v51
                                            }
                                        case 608:
                                            _ = len(v4)
                                        }
                                        v28 = (false && (-1 > fi27))
                                        v28 = v28
                                    } else {
                                        var v52 int32 = 626
                                        var v53 uint16 = 933
                                        _ = v52
                                        _ = v53
                                    }
                                    v21 += (v21 - (v21 / (((v21 ^ (v21 * v21))) | 1)))
                                    if (2147483647 > 42) {
                                        if true {
                                            var v54 float32 = 0.0
                                            _ = v54
                                        }
                                        if v28 {
                                            v21 = 9223372036854775807
                                            v2 = append(v2, "\n\t")
                                            switch "\n\t" {
                                            case "hello":
                                                switch fi24 {
                                                case 128:
                                                    var v55 int16 = -532
                                                    _ = v55
                                                case -292:
                                                    v56 := v1[562]
                                                    var v57 int32 = 2147483647
                                                    var v58 uint8 = 42
                                                    _ = v56
                                                    _ = v57
                                                    _ = v58
                                                case 2147483647:
                                                    var v59 int32 = -2147483648
                                                    _ = v59
                                                }
                                            case "a":
                                                fmt.Println(v21, v28, v28)
                                            case "\n\t":
                                                v60 := GenS3{F0: 490, F1: 81}
                                                _ = v60
                                            default:
                                                v28 = false
                                                switch (fi27 >> ((0) & 63)) {
                                                case 0:
                                                    if ("hello" < ("" + ("" + ("hello" + "")))) {
                                                        v61 := GenS3{F0: 682, F1: 128}
                                                        v62 := GenS1{F0: (!v28)}
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                case 589:
                                                    var v63 float32 = (((0.0 / (-2.25 - ((((1.5 * 1.5) - 0.1) / 1e10) - 0.0))) / 3.14159) / 1e10)
                                                    _ = v63
                                                }
                                            }
                                        }
                                        v21 += (v21 * v21)
                                    }
                                }
                            case 166:
                                switch 530 {
                                case 0:
                                    fmt.Println(v21, fi27)
                                }
                            }
                        }
                        v64 := GenS3{F0: 65535, F1: 127}
                        _ = fi27
                        _ = v28
                        _ = v64
                    }
                    for fi65 := 0; fi65 < 3; fi65++ {
                        v66 := v1[65535]
                        v66 = v66
                        var v67 float64 = (0.1 / 1.0)
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                    _ = fi24
                }
            }
            _ = v21
        }
        switch -867 {
        case 2147483647:
            delete(v1, 219)
        default:
            v1[2147483648] = "foo bar"
        }
    }
    var v68 float32 = float32(0.0)
    fmt.Println(v68)
    var v69 uint = 0
    v4[32768] = (("\n\t" + (("hello" + "hello") + "\n\t")) + (("\n\t" + "") + ""))
    v68 = v68
    for fi70 := 0; fi70 < 16; fi70++ {
        _ = len(v2)
        for fi71 := 0; fi71 < 15; fi71++ {
            fmt.Println(v69, fi70)
            v72 := GenS2{F0: "foo bar", F1: v69, F2: 255}
            v73 := v4[fi71]
            _ = fi71
            _ = v72
            _ = v73
        }
        _ = fi70
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v68
    _ = v69
    _ = v9
}
