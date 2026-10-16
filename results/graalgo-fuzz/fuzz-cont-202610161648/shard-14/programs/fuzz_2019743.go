package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 uint = 961
    var v5 uint32 = 2
    v5 = 102
    v2[581] = ""
    var v6 float64 = 1.5
    switch ("" + ("foo bar" + "")) {
    case "":
        v7 := GenS1{F0: 2147483647, F1: ("" + "foo bar")}
        _ = v7
    }
    v3[1] = ("" + "foo bar")
    fmt.Println(v6)
    var v8 float32 = -2.25
    v9 := []int{42, 32767}
    var v10 int8 = 0
    var v11 float32 = (-2.25 * v8)
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 uint16 = 0
        _ = fi12
        _ = v13
    }
    v5 += 249
    if true {
        switch 999 {
        case 127:
            for fi14 := 0; fi14 < 9; fi14++ {
                v8 -= float32((float64(v11) - 1e10))
                v11 -= (-2.25 * (v8 + (-2.25 - 1e10)))
                var v15 string = (("a" + ("a" + (((("\n\t" + ("a" + "a")) + ((("" + "\n\t") + "foo bar") + ("" + "hello"))) + "a") + "foo bar"))) + ((("\n\t" + "\n\t") + "foo bar") + ("\n\t" + ("foo bar" + (("\n\t" + (("hello" + "foo bar") + "")) + (("\n\t" + (("a" + "") + "foo bar")) + ("hello" + "foo bar")))))))
                _ = fi14
                _ = v15
            }
        case 32767:
            if (0.0 >= v8) {
                var v16 int8 = (v10 - (v10 & (v10 ^ (v10 >> (((v4 >> (((v4 & (v4 - uint(v5)))) & 63))) & 63)))))
                v17 := GenS1{F0: (v5 * (v5 * 1)), F1: "foo bar"}
                if (0.1 != 1e10) {
                    v4 += v4
                    var v18 int32 = 256
                    var v19 int = 32768
                    _ = v18
                    _ = v19
                }
                _ = v16
                _ = v17
            } else {
                var v20 int8 = v10
                _ = v20
            }
        }
    } else {
        v6 -= ((v6 * 1.5) / 1e10)
        switch -474 {
        case 2147483647:
            for fi21 := 0; fi21 < 9; fi21++ {
                if (!(!((v5 + v5) < (v5 >> ((v4) & 63))))) {
                    for fi22 := 0; fi22 < 8; fi22++ {
                        v23 := GenS1{F0: (^(v5)), F1: ""}
                        _ = fi22
                        _ = v23
                    }
                }
                fmt.Println(v10, v11, v11)
                for fi24 := 0; fi24 < 14; fi24++ {
                    v10 += 1
                    if false {
                        if true {
                            _ = len(v9)
                            switch "foo bar" {
                            case "\n\t":
                                if (!false) {
                                    v25 := GenS2{F0: 4294967295, F1: 127, F2: v4}
                                    var v26 float64 = v6
                                    switch (fi21 * (fi24 - fi24)) {
                                    case -338:
                                        var v27 int8 = -128
                                        _ = v27
                                    }
                                    _ = v25
                                    _ = v26
                                } else {
                                    var v28 int64 = -289
                                    fmt.Println(fi21, v8)
                                    v29 := GenS1{F0: v5, F1: "a"}
                                    _ = v28
                                    _ = v29
                                }
                            case "hello":
                                _ = len(v9)
                            case "":
                                switch 127 {
                                case 127:
                                    v30 := GenS2{F0: 18446744073709551615, F1: uint8(v4), F2: 658}
                                    _ = v30
                                case 1:
                                    fmt.Println(v6, fi21, v6)
                                case -556:
                                    fmt.Println(v5)
                                default:
                                    v31 := GenS2{F0: 870, F1: 1, F2: v4}
                                    switch fi24 {
                                    case 249:
                                        v32 := GenS3{F0: 634}
                                        _ = v32
                                    case 31:
                                        var v33 int16 = 516
                                        _ = v33
                                    }
                                    v1[v5] = 0
                                    _ = v31
                                }
                            }
                            _ = len(v1)
                            fmt.Println(fi21, fi24, v10)
                            if false {
                                if (0.1 <= (((v6 / (v6 - ((v6 / v6) - (((v6 - v6) + 0.0) + 3.14159)))) + -2.25) - v6)) {
                                    v5 -= v5
                                } else {
                                    fmt.Println(v4)
                                    var v34 float32 = v11
                                    _ = v34
                                }
                                v6 = (0.0 - (v6 * float64((-2.25 + v11))))
                            } else {
                                v11 += 1.5
                                switch "\n\t" {
                                case "\n\t":
                                    if (true && false) {
                                        v35 := GenS1{F0: (v5 >> (((v4 >> ((v4) & 63))) & 63)), F1: "hello"}
                                        if ((^(v4)) > v4) {
                                            var v36 uint64 = 0
                                            _ = v36
                                        }
                                        _ = v35
                                    }
                                }
                            }
                        }
                        fmt.Println(v5, v10)
                    } else {
                        var v37 uint8 = 1
                        _ = v37
                    }
                    if true {
                        v38 := GenS1{F0: v5, F1: "a"}
                        if len(v9) > 0 {
                            v9[0] = (fi24 >> (((v4 / ((v4) | 1))) & 63))
                        }
                        _ = v38
                    }
                    _ = fi24
                }
                _ = fi21
            }
        default:
            var v39 int = 1
            _ = v39
        }
        if (!false) {
            var v40 int8 = (v10 ^ v10)
            var v41 string = "a"
            v5 -= v5
            _ = v40
            _ = v41
        }
    }
    if (!(false || (!(!true)))) {
        v8 = v11
    } else {
        if false {
            if false {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                if ((!(!(4294967295 <= 42))) || false) {
                    fmt.Println(v10)
                    var v42 float32 = -2.25
                    fmt.Println(v10, v42, v42)
                    _ = v42
                } else {
                    for fi43 := 0; fi43 < 20; fi43++ {
                        if true {
                            var v44 float64 = (0.0 - 1.5)
                            if ((v8 + (1.5 + v11)) == 1.5) {
                                for fi45 := 0; fi45 < 5; fi45++ {
                                    _ = len(v3)
                                    fmt.Println(v8)
                                    v44 = (0.0 / float64((v11 - (((v11 * 0.1) - v11) / (v8 + 0.1)))))
                                    _ = fi45
                                }
                                var v46 uint16 = 255
                                if (1 <= 0) {
                                    for fi47 := 0; fi47 < 5; fi47++ {
                                        switch int(v46) {
                                        case 42:
                                            if (128 != 255) {
                                                var v48 uint8 = uint8(v46)
                                                var v49 uint16 = 860
                                                _ = v48
                                                _ = v49
                                            } else {
                                                v50 := GenS1{F0: (v5 << (((v4 + 2)) & 63)), F1: ("" + ("" + "foo bar"))}
                                                v11 = v8
                                                v46 -= v46
                                                _ = v50
                                            }
                                        default:
                                            var v51 int32 = 176
                                            _ = v51
                                        }
                                        if false {
                                            var v52 uint32 = (v5 | 0)
                                            var v53 float32 = v8
                                            if true {
                                                var v54 int16 = 411
                                                var v55 int16 = (v54 << (((v4 ^ v4)) & 63))
                                                if true {
                                                    if (v10 >= v10) {
                                                        if false {
                                                            v4 += v4
                                                        }
                                                    } else {
                                                        fmt.Println(fi47, v11)
                                                    }
                                                    var v56 uint16 = (v46 | v46)
                                                    _ = v56
                                                }
                                                _ = v54
                                                _ = v55
                                            } else {
                                                var v57 int8 = 100
                                                v58 := v1[(v5 | (v52 >> (((v4 % ((v4) | 1))) & 63)))]
                                                var v59 int16 = 1000
                                                v60 := GenS1{F0: (v58 << ((v4) & 63)), F1: ""}
                                                if true {
                                                    v61 := GenS1{F0: (v52 ^ 42), F1: "foo bar"}
                                                    v62 := GenS2{F0: 18446744073709551615, F1: 1, F2: uint(v46)}
                                                    if (v58 <= 748) {
                                                        switch "a" {
                                                        case "a":
                                                            if true {
                                                                v44 -= 1.5
                                                                fmt.Println(v53, v44)
                                                            }
                                                        case "hello":
                                                            _ = len(v9)
                                                        default:
                                                            var v63 int64 = -1
                                                            _ = v63
                                                        }
                                                        v64 := GenS2{F0: 0, F1: uint8(v4), F2: v4}
                                                        _ = v64
                                                    }
                                                    _ = v61
                                                    _ = v62
                                                }
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                                _ = v57
                                                _ = v58
                                                _ = v59
                                                _ = v60
                                            }
                                            switch (fi43 >> (((v4 >> (((v4 << (((v4 & v4)) & 63))) & 63))) & 63)) {
                                            case -2147483648:
                                                var v65 int32 = int32(v4)
                                                _ = v65
                                            }
                                            _ = v52
                                            _ = v53
                                        }
                                        v66 := GenS2{F0: 4294967295, F1: 128, F2: v4}
                                        v8 += 1e10
                                        var v67 float64 = v44
                                        _ = fi47
                                        _ = v66
                                        _ = v67
                                    }
                                    if (0 != 42) {
                                        v68 := GenS3{F0: 438}
                                        _ = v68
                                    } else {
                                        for fi69 := 0; fi69 < 2; fi69++ {
                                            var v70 int64 = 190
                                            v71 := GenS1{F0: v5, F1: ("foo bar" + "\n\t")}
                                            v8 += v8
                                            _ = fi69
                                            _ = v70
                                            _ = v71
                                        }
                                        var v72 uint8 = 98
                                        _ = v72
                                    }
                                    var v73 uint16 = (v46 << (((v4 | (v4 >> (((^(v4))) & 63)))) & 63))
                                    _ = v73
                                } else {
                                    v8 = v8
                                }
                                _ = v46
                            }
                            var v74 string = "\n\t"
                            _ = v44
                            _ = v74
                        }
                        _ = fi43
                    }
                    fmt.Println(v8, v8)
                }
                v6 = v6
            } else {
                v6 += (0.1 - 0.0)
            }
            delete(v1, 947)
            v4 = v4
        }
    }
    v75 := GenS4{F0: -867}
    var v76 uint = (v4 & uint(v5))
    v77 := GenS4{F0: 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v76
    _ = v77
    _ = v8
    _ = v9
}
