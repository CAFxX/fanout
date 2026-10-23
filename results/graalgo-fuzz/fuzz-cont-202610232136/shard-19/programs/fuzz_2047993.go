package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint16; F1 string; F2 int64 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint8; F1 bool }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    var v4 int = -397
    v4 -= v4
    fmt.Println(v4)
    var v5 bool = true
    var v6 int8 = -1
    v6 += (v6 << ((4294967295) & 63))
    _ = len(v2)
    v6 -= 0
    var v7 int = v4
    switch ("foo bar" + "") {
    case "\n\t":
        var v8 uint32 = 1
        _ = v8
    default:
        v5 = (128 <= 55)
    }
    switch "a" {
    case "":
        if (!((1000 < 0) && v5)) {
            if v5 {
                var v9 uint16 = 255
                _ = v9
            }
        }
    }
    switch ("a" + "foo bar") {
    case "hello":
        for fi10 := 0; fi10 < 7; fi10++ {
            v5 = false
            _ = fi10
        }
    case "\n\t":
        fmt.Println(v4, v7)
    case "foo bar":
        var v11 uint16 = 0
        _ = v11
    default:
        v6 -= (v6 + v6)
    }
    if v5 {
        if v5 {
            v12 := []int{911}
            switch "" {
            case "hello":
                if (!(2 >= 42)) {
                    for fi13 := 0; fi13 < 12; fi13++ {
                        for fi14 := 0; fi14 < 17; fi14++ {
                            _ = len(v1)
                            _ = fi14
                        }
                        if true {
                            for fi15 := 0; fi15 < 9; fi15++ {
                                var v16 int32 = 91
                                _ = fi15
                                _ = v16
                            }
                        }
                        _ = len(v12)
                        _ = fi13
                    }
                    v17 := GenS1{F0: (("" + ("foo bar" + "a")) + "")}
                    _ = v17
                }
            default:
                var v18 int16 = 284
                _ = v18
            }
            v4 = v7
            _ = v12
        } else {
            _ = len(v3)
            switch ("a" + "") {
            case "a":
                for fi19 := 0; fi19 < 15; fi19++ {
                    var v20 int64 = -1
                    _ = fi19
                    _ = v20
                }
            case "hello":
                v6 = (v6 + 1)
            default:
                for fi21 := 0; fi21 < 11; fi21++ {
                    v22 := GenS2{F0: 32767, F1: 42}
                    var v23 int64 = -9223372036854775808
                    if (v5 || (true && v5)) {
                        var v24 uint = 65535
                        _ = v24
                    } else {
                        var v25 int16 = 85
                        if v5 {
                            v5 = true
                            for fi26 := 0; fi26 < 3; fi26++ {
                                var v27 uint64 = 42
                                _ = fi26
                                _ = v27
                            }
                            v25 = -32768
                        }
                        var v28 bool = false
                        _ = v25
                        _ = v28
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            }
            for fi29 := 0; fi29 < 18; fi29++ {
                var v30 uint16 = 255
                v31 := v3["\n\t"]
                if true {
                    for fi32 := 0; fi32 < 17; fi32++ {
                        var v33 int32 = 1
                        if true {
                            if v5 {
                                v34 := GenS1{F0: ""}
                                var v35 int = (v4 % (((v4 | (-(v4)))) | 1))
                                v35 = (v4 << ((42) & 63))
                                if (!(-746 == -9223372036854775808)) {
                                    var v36 uint = 4294967295
                                    var v37 uint = (v36 >> (((v36 | v36)) & 63))
                                    v31 -= 32767
                                    _ = v36
                                    _ = v37
                                } else {
                                    var v38 float64 = 1e10
                                    v39 := v2[2147483647]
                                    _ = v38
                                    _ = v39
                                }
                                _ = v34
                                _ = v35
                            } else {
                                switch (("hello" + ("foo bar" + ("\n\t" + ""))) + "\n\t") {
                                case "hello":
                                    v40 := GenS3{F0: (v30 & (v30 % (((v30 ^ v30)) | 1))), F1: ("\n\t" + (("" + "\n\t") + "a")), F2: -9223372036854775808}
                                    _ = v40
                                }
                                switch "hello" {
                                case "\n\t":
                                    v4 = 2
                                case "":
                                    v3["foo bar"] = v31
                                case "foo bar":
                                    switch ("hello" + "a") {
                                    case "hello":
                                        if ((("\n\t" < ("a" + "hello")) && v5) || (!true)) {
                                            fmt.Println(v30)
                                        }
                                    case "a":
                                        if (0 < 65535) {
                                            fmt.Println(v30, v31)
                                            if v5 {
                                                v41 := GenS3{F0: v30, F1: "", F2: -1}
                                                v33 += v33
                                                _ = v41
                                            } else {
                                                v42 := GenS3{F0: 42, F1: "hello", F2: 0}
                                                _ = v42
                                            }
                                        } else {
                                            v2[2147483647] = (v7 | v31)
                                        }
                                        var v43 float32 = 1e10
                                        _ = v43
                                    default:
                                        if v5 {
                                            fmt.Println(v31, v30, v5)
                                            var v44 uint16 = (v30 ^ (v30 - (v30 & (v30 % ((v30) | 1)))))
                                            if ((0 > 0) && (1 >= v33)) {
                                                switch "a" {
                                                case "":
                                                    fmt.Println(fi29, v31)
                                                default:
                                                    v6 = (v6 >> ((1) & 63))
                                                }
                                                v5 = false
                                                var v45 uint32 = 65535
                                                var v46 uint64 = 384
                                                _ = v45
                                                _ = v46
                                            } else {
                                                switch fi29 {
                                                case 1:
                                                    var v47 string = ("\n\t" + "\n\t")
                                                    _ = v47
                                                case 292:
                                                    v6 -= (v6 & (v6 / (((v6 << ((42) & 63))) | 1)))
                                                    v31 += 127
                                                    v4 += 128
                                                }
                                            }
                                            _ = v44
                                        }
                                    }
                                }
                                v48 := v1[0]
                                _ = v48
                            }
                        }
                        _ = fi32
                        _ = v33
                    }
                    v49 := []int{42, v7}
                    var v50 uint8 = 1
                    _ = v49
                    _ = v50
                } else {
                    v51 := GenS1{F0: ("hello" + "hello")}
                    for fi52 := 0; fi52 < 7; fi52++ {
                        var v53 string = "a"
                        v5 = false
                        v54 := []int{(fi52 & (v31 * (fi52 * (fi29 + v7)))), (v31 << ((4294967295) & 63)), (v4 | (v31 << ((82) & 63))), v7}
                        _ = fi52
                        _ = v53
                        _ = v54
                    }
                    _ = v51
                }
                _ = fi29
                _ = v30
                _ = v31
            }
        }
        if (!v5) {
            v55 := GenS1{F0: "\n\t"}
            v56 := GenS4{F0: 32767, F1: 613, F2: 2147483648}
            v1[1] = "foo bar"
            _ = v55
            _ = v56
        } else {
            v57 := []int{(v7 | (-(v7))), (^(v7))}
            var v58 uint8 = 127
            _ = v57
            _ = v58
        }
        switch v4 {
        case -169:
            fmt.Println(v6)
        case 2147483647:
            v59 := []int{(v7 << ((1) & 63)), v4, v4}
            _ = v59
        case 4294967295:
            v60 := GenS2{F0: 0, F1: 996}
            _ = v60
        default:
            if v5 {
                var v61 float32 = 0.0
                if v5 {
                    if v5 {
                        fmt.Println(v61, v5, v61)
                        if v5 {
                            for fi62 := 0; fi62 < 18; fi62++ {
                                v6 -= 42
                                switch -400 {
                                case -2147483648:
                                    var v63 int8 = (v6 | -5)
                                    _ = v63
                                case 1:
                                    for fi64 := 0; fi64 < 16; fi64++ {
                                        v5 = true
                                        _ = fi64
                                    }
                                case 32767:
                                    if v5 {
                                        v4 = v7
                                        for fi65 := 0; fi65 < 11; fi65++ {
                                            var v66 uint8 = 216
                                            var v67 string = ((("hello" + "hello") + (("\n\t" + ("hello" + "a")) + "\n\t")) + "")
                                            v68 := []int{(fi65 * 245)}
                                            v4 += v7
                                            var v69 uint8 = 128
                                            _ = fi65
                                            _ = v66
                                            _ = v67
                                            _ = v68
                                            _ = v69
                                        }
                                    } else {
                                        for fi70 := 0; fi70 < 12; fi70++ {
                                            v71 := GenS5{F0: 42, F1: v5}
                                            _ = fi70
                                            _ = v71
                                        }
                                    }
                                default:
                                    for fi72 := 0; fi72 < 6; fi72++ {
                                        if v5 {
                                            switch fi62 {
                                            case 32768:
                                                var v73 uint = 114
                                                _ = v73
                                            }
                                            var v74 string = ""
                                            _ = v74
                                        } else {
                                            var v75 float32 = ((v61 - v61) / float32(1.5))
                                            var v76 int16 = -762
                                            if v5 {
                                                v77 := v3[("hello" + "hello")]
                                                fmt.Println(v77, v7)
                                                fmt.Println(v77, v6, fi72)
                                                _ = v77
                                            } else {
                                                var v78 int16 = v76
                                                v1[2] = "foo bar"
                                                var v79 uint8 = 2
                                                _ = v78
                                                _ = v79
                                            }
                                            _ = v75
                                            _ = v76
                                        }
                                        _ = fi72
                                    }
                                }
                                v61 += v61
                                _ = fi62
                            }
                            v80 := GenS4{F0: -921, F1: 2147483647, F2: 1}
                            _ = v80
                        }
                    }
                    var v81 int32 = -1
                    _ = v81
                }
                _ = v61
            } else {
                v82 := []int{v4, v4, 4294967295}
                if v5 {
                    for fi83 := 0; fi83 < 20; fi83++ {
                        v4 -= (v7 % ((int(v6)) | 1))
                        var v84 int8 = (v6 * v6)
                        fmt.Println(v7, v5)
                        _ = fi83
                        _ = v84
                    }
                    v5 = (216 == 767)
                    fmt.Println(v7, v7)
                } else {
                    v5 = v5
                }
                _ = v82
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
