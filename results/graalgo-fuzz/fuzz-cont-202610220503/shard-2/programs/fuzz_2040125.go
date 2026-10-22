package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint16; F2 int }
type GenS4 struct { F0 string }
func genFn0(p0 uint8) int64 {
    v7 := []int{}
    if len(v7) > 0 {
        v7[0] = 227
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        if len(v7) > 0 {
            v7[0] = (fi8 * fi8)
        }
        if len(v7) > 0 {
            v7[0] = (fi8 % (((fi8 >> ((65535) & 63))) | 1))
        }
        fmt.Println(fi8)
        _ = fi8
    }
    _ = p0
    _ = v7
    return 0
}
func genFn1() bool {
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: 0, F1: 2147483647}
        for fi11 := 0; fi11 < 12; fi11++ {
            if (1 > 0) {
                var v12 uint8 = 128
                switch -705 {
                case 9223372036854775807:
                    v13 := []int{-879, (fi9 >> ((1) & 63))}
                    _ = v13
                }
                v14 := []int{(^(fi11)), (fi9 / ((0) | 1))}
                _ = v12
                _ = v14
            }
            if (-2.25 >= (float32(0.0) * float32(float64(((1e10 / float32(-2.25)) / 1e10))))) {
                var v15 float64 = 0.1
                if (!false) {
                    switch fi9 {
                    case 32768:
                        v15 = v15
                    case 127:
                        var v16 uint = 65535
                        _ = v16
                    case 4294967295:
                        var v17 int8 = -1
                        _ = v17
                    }
                    fmt.Println(v15)
                    fmt.Println(fi9, fi9)
                    if ((1 < 32767) && true) {
                        var v18 int8 = -128
                        v19 := GenS2{F0: 535, F1: 42}
                        _ = v18
                        _ = v19
                    }
                    var v20 string = ("a" + "foo bar")
                    _ = v20
                }
                if ((42 < -99) && (true || false)) {
                    var v21 string = "hello"
                    _ = v21
                }
                _ = v15
            }
            var v22 bool = false
            v23 := GenS1{F0: 48, F1: 906}
            _ = fi11
            _ = v22
            _ = v23
        }
        v24 := []int{fi9, (^(fi9)), (fi9 << ((962) & 63))}
        _ = fi9
        _ = v10
        _ = v24
    }
    v25 := []int{-955, 2147483647, -405}
    if false {
        var v26 float64 = 3.14159
        _ = v26
    }
    var v27 float64 = float64(((0.0 / (0.0 - -2.25)) - 0.1))
    _ = v25
    _ = v27
    return (0.0 != v27)
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    var v3 int = 4294967295
    var v4 string = "\n\t"
    v5 := []uint32{2, 65535, 42, 863}
    v6 := make(map[string]string)
    for fi28 := 0; fi28 < 2; fi28++ {
        switch -788 {
        case 128:
            for fi29 := 0; fi29 < 12; fi29++ {
                _ = len(v6)
                _ = len(v5)
                if true {
                    var v30 bool = true
                    v3 += fi28
                    var v31 string = v4
                    var v32 int64 = -249
                    v33 := GenS1{F0: 255, F1: 100}
                    delete(v6, v4)
                    var v34 uint = 993
                    var v35 float64 = (0.0 - -2.25)
                    _ = v30
                    _ = v31
                    _ = v32
                    _ = v33
                    _ = v34
                    _ = v35
                }
                var v36 float64 = -2.25
                _ = fi29
                _ = v36
            }
        case -878:
            v37 := GenS3{F0: int8(fi28), F1: 42, F2: (fi28 + -1)}
            _ = v37
        case -737:
            fmt.Println(v4, v3, v4)
        }
        for fi38 := 0; fi38 < 1; fi38++ {
            if (18446744073709551615 > 4294967295) {
                if false {
                    var v39 uint8 = 1
                    v40 := GenS3{F0: 0, F1: 845, F2: 2147483647}
                    if true {
                        if false {
                            fmt.Println(v39, fi38, v39)
                        }
                    }
                    _ = v39
                    _ = v40
                }
            } else {
                if false {
                    switch fi38 {
                    case -1:
                        var v41 uint64 = 18446744073709551615
                        _ = v41
                    case 127:
                        v4 = (((v4 + "foo bar") + "foo bar") + "\n\t")
                    }
                    fmt.Println(v4, fi38)
                    fmt.Println(v4, v3, fi28)
                } else {
                    v42 := GenS2{F0: 1, F1: 128}
                    _ = v42
                }
                var v43 string = v4
                switch fi38 {
                case 127:
                    v44 := GenS1{F0: 0, F1: 42}
                    var v45 string = v4
                    var v46 float32 = (0.0 + 1e10)
                    _ = v44
                    _ = v45
                    _ = v46
                }
                _ = v43
            }
            var v47 uint64 = 18446744073709551615
            _ = fi38
            _ = v47
        }
        _ = fi28
    }
    var v48 uint8 = 0
    var v49 int = (v3 + v3)
    v48 += v48
    switch v4 {
    case "":
        var v50 int32 = 655
        _ = v50
    case "a":
        v49 -= 0
    case "\n\t":
        var v51 uint8 = 127
        _ = v51
    }
    for fi52 := 0; fi52 < 19; fi52++ {
        switch ((v4 + "hello") + v4) {
        case "foo bar":
            v53 := GenS4{F0: ("foo bar" + "\n\t")}
            _ = v53
        default:
            for fi54 := 0; fi54 < 17; fi54++ {
                if (-2147483648 < (v49 * fi52)) {
                    v48 += (v48 ^ v48)
                    var v55 uint64 = 222
                    v3 -= -17
                    _ = v55
                } else {
                    switch fi52 {
                    case 705:
                        var v56 bool = false
                        _ = v56
                    case 32767:
                        if false {
                            switch v4 {
                            case "":
                                if false {
                                    if true {
                                        v57 := GenS3{F0: 0, F1: 65535, F2: fi54}
                                        fmt.Println(v3, v49)
                                        if true {
                                            switch "hello" {
                                            case "foo bar":
                                                if false {
                                                    switch fi54 {
                                                    case 42:
                                                        fmt.Println(v3, fi54, v48)
                                                    case -2147483648:
                                                        var v58 int32 = 1
                                                        _ = v58
                                                    default:
                                                        v4 = (v4 + ((v4 + (v4 + (v4 + "foo bar"))) + v4))
                                                    }
                                                    v3 += 4294967295
                                                    if (true && true) {
                                                        v4 = "a"
                                                        v48 = v48
                                                    }
                                                    v4 = ""
                                                }
                                            }
                                            if ((42 != 787) || (v48 < (^(v48)))) {
                                                fmt.Println(v48, v48)
                                                fmt.Println(v48, v4)
                                            }
                                            v48 -= v48
                                        } else {
                                            if (fi54 < 478) {
                                                var v59 bool = (561 < 4294967295)
                                                _ = v59
                                            }
                                        }
                                        _ = v57
                                    } else {
                                        v60 := GenS4{F0: v4}
                                        var v61 int32 = -2147483648
                                        switch ("a" + v4) {
                                        case "foo bar":
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        case "hello":
                                            var v62 int8 = int8(v48)
                                            _ = len(v5)
                                            _ = v62
                                        case "\n\t":
                                            var v63 int = -306
                                            _ = v63
                                        default:
                                            if true {
                                                switch ("a" + (v4 + "")) {
                                                case "hello":
                                                    _ = len(v1)
                                                default:
                                                    v3 = (v3 << ((111) & 63))
                                                }
                                                v49 += (v3 % ((fi52) | 1))
                                                var v64 int = fi52
                                                _ = v64
                                            }
                                        }
                                        _ = v60
                                        _ = v61
                                    }
                                    var v65 int8 = -58
                                    var v66 int64 = genFn0(255)
                                    _ = v65
                                    _ = v66
                                } else {
                                    fmt.Println(fi52)
                                    var v67 uint8 = (v48 % (((v48 % ((v48) | 1))) | 1))
                                    fmt.Println(v4)
                                    _ = v67
                                }
                            default:
                                if true {
                                    v5 = append(v5, 908)
                                    if (3.14159 == (1e10 * ((1e10 - ((0.1 - 0.1) - 1e10)) / 1.0))) {
                                        _ = len(v5)
                                    } else {
                                        if len(v5) > 0 {
                                            v5[0] = 2147483647
                                        }
                                        v48 = v48
                                        if false {
                                            var v68 bool = true
                                            fmt.Println(fi54, v3, v48)
                                            _ = v68
                                        }
                                    }
                                    switch (("hello" + ("" + "")) + v4) {
                                    case "":
                                        v4 = v4
                                    }
                                }
                            }
                            var v69 float32 = ((0.0 + 0.0) + 3.14159)
                            _ = v69
                        }
                    default:
                        fmt.Println(fi52, fi54)
                    }
                }
                _ = fi54
            }
        }
        _ = fi52
    }
    var v70 int8 = -69
    switch (v49 - 4294967295) {
    case 9223372036854775807:
        var v71 string = (v4 + (v4 + v4))
        _ = v71
    case 42:
        var v72 int8 = (v70 / ((v70) | 1))
        _ = v72
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v70
}
