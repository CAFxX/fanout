package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 string }
type GenS2 struct { F0 int64; F1 int8; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 uint8; F2 int64 }
func genFn0() uint16 {
    var v7 uint32 = 820
    v7 = (^(v7))
    v8 := []int{-1, 426, 4294967295}
    var v9 uint8 = 2
    v10 := GenS1{F0: v7, F1: "a", F2: "foo bar"}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 65535
}
func genFn1() bool {
    var v11 float64 = float64(3.14159)
    for fi12 := 0; fi12 < 2; fi12++ {
        v11 -= float64(((-2.25 / 1e10) - (1e10 / 1e10)))
        _ = fi12
    }
    switch 127 {
    case 127:
        switch 32768 {
        case -1:
            for fi13 := 0; fi13 < 14; fi13++ {
                for fi14 := 0; fi14 < 18; fi14++ {
                    v11 = float64((float32((v11 + v11)) * (0.1 * 0.0)))
                    _ = fi14
                }
                fmt.Println(v11)
                var v15 uint = 183
                _ = fi13
                _ = v15
            }
        }
    case 924:
        for fi16 := 0; fi16 < 8; fi16++ {
            v17 := []int{fi16, 157}
            v11 = float64(float32(v11))
            v18 := GenS1{F0: 65535, F1: "foo bar", F2: (("" + "") + ((("a" + "") + "a") + ("" + ("a" + ("" + "hello")))))}
            _ = fi16
            _ = v17
            _ = v18
        }
    case 4294967295:
        if true {
            for fi19 := 0; fi19 < 16; fi19++ {
                var v20 uint = 1
                _ = fi19
                _ = v20
            }
            if true {
                var v21 bool = (((v11 - ((float64(0.0) / -2.25) / v11)) <= v11) && (-32768 <= 1000))
                var v22 uint32 = 2147483647
                _ = v21
                _ = v22
            }
            switch -9223372036854775808 {
            case 786:
                for fi23 := 0; fi23 < 19; fi23++ {
                    v24 := GenS1{F0: 65535, F1: "", F2: "hello"}
                    var v25 uint = 255
                    fmt.Println(fi23)
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            case -592:
                var v26 int = 128
                _ = v26
            case 9223372036854775807:
                fmt.Println(v11)
            default:
                var v27 uint8 = 127
                _ = v27
            }
        } else {
            fmt.Println(v11)
        }
    default:
        fmt.Println(v11)
    }
    v28 := GenS2{F0: -1, F1: -48, F2: 9223372036854775807}
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 int16 = -571
        _ = fi29
        _ = v30
    }
    _ = v11
    _ = v28
    return (1000 < 785)
}

func main() {
    v1 := []int{1, -900}
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    v4 := []string{("\n\t" + "\n\t")}
    var v5 int = -2147483648
    v6 := []string{"\n\t", "\n\t"}
    for fi31 := 0; fi31 < 8; fi31++ {
        if ((32767 == 721) && false) {
            v32 := GenS3{F0: 32767}
            switch (fi31 | 32767) {
            case 1:
                var v33 float32 = (3.14159 * 0.0)
                _ = v33
            default:
                v5 = 32768
            }
            _ = v32
        } else {
            v5 -= fi31
            v5 += (v5 - (fi31 + v5))
        }
        _ = fi31
    }
    var v34 uint16 = genFn0()
    var v35 bool = false
    v36 := v3[2147483648]
    var v37 int16 = 0
    if len(v6) > 0 {
        v6[0] = "hello"
    }
    v37 = v37
    fmt.Println(v37)
    if v35 {
        for fi38 := 0; fi38 < 14; fi38++ {
            for fi39 := 0; fi39 < 16; fi39++ {
                v36 -= (v36 + v36)
                if v35 {
                    var v40 uint32 = v36
                    _ = v40
                } else {
                    v35 = (872 >= 0)
                    fmt.Println(v34, fi38, v36)
                    if len(v6) > 0 {
                        v6[0] = ("" + "foo bar")
                    }
                    v4 = append(v4, ("\n\t" + "a"))
                }
                _ = fi39
            }
            for fi41 := 0; fi41 < 18; fi41++ {
                v42 := GenS2{F0: 1, F1: -128, F2: -312}
                fmt.Println(fi41, v5, fi41)
                _ = fi41
                _ = v42
            }
            fmt.Println(v36, v35, v37)
            _ = fi38
        }
        v37 = (-(v37))
    }
    var v43 string = ""
    fmt.Println(v35, v43)
    switch (v5 | v5) {
    case 0:
        if v35 {
            v35 = false
            v44 := GenS1{F0: (v36 >> ((0) & 63)), F1: v43, F2: v43}
            var v45 int = int(v34)
            _ = v44
            _ = v45
        }
    case 1:
        var v46 bool = ((!v35) && v35)
        _ = v46
    case -40:
        if (((0 != 1) && v35) && v35) {
            fmt.Println(v36)
        }
    default:
        switch ((("foo bar" + "a") + v43) + (("foo bar" + ((v43 + "hello") + "foo bar")) + ("a" + v43))) {
        case "hello":
            _ = len(v3)
        case "a":
            switch 9223372036854775807 {
            case -579:
                v47 := GenS4{F0: v34, F1: 42, F2: 467}
                _ = v47
            case 478:
                v48 := v2[v36]
                _ = v48
            case -342:
                v49 := GenS3{F0: v37}
                _ = v49
            }
        case "foo bar":
            var v50 int16 = -32768
            _ = v50
        default:
            var v51 string = v43
            _ = v51
        }
    }
    if len(v1) > 0 {
        v1[0] = v5
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (9223372036854775807 < 9223372036854775807) {
        for fi52 := 0; fi52 < 6; fi52++ {
            if v35 {
                v53 := GenS1{F0: (v36 << ((24) & 63)), F1: v43, F2: ("hello" + v43)}
                for fi54 := 0; fi54 < 4; fi54++ {
                    var v55 int64 = 9223372036854775807
                    if (("\n\t" == v43) || (!(-128 > 100))) {
                        v3[1] = (v36 << ((0) & 63))
                    } else {
                        if (42 >= 1) {
                            v36 -= 2147483647
                        } else {
                            if true {
                                v56 := GenS4{F0: (v34 ^ (v34 >> ((255) & 63))), F1: 42, F2: v55}
                                _ = v56
                            } else {
                                var v57 bool = (2147483647 != (v36 * 428))
                                if false {
                                    fmt.Println(fi54)
                                    v57 = ((fi52 >= 1) || v35)
                                }
                                var v58 int16 = v37
                                v59 := GenS1{F0: v36, F1: v43, F2: "a"}
                                v35 = (1 < 42)
                                v34 += uint16(v36)
                                _ = v57
                                _ = v58
                                _ = v59
                            }
                            fmt.Println(v5, fi54)
                        }
                    }
                    fmt.Println(v37)
                    _ = fi54
                    _ = v55
                }
                _ = v53
            }
            v36 += 1
            _ = fi52
        }
    }
    switch 42 {
    case 893:
        switch (v5 + v5) {
        case -9223372036854775808:
            if v35 {
                v36 = 4294967295
            } else {
                v35 = true
                switch "a" {
                case "a":
                    v37 = (v37 - v37)
                default:
                    switch 1 {
                    case 1:
                        fmt.Println(v37)
                    case -655:
                        fmt.Println(v37)
                    case -2147483648:
                        var v60 uint = uint(v34)
                        _ = v60
                    }
                }
                if true {
                    v61 := GenS1{F0: 949, F1: v43, F2: v43}
                    var v62 float32 = 1e10
                    switch ((v43 + ("foo bar" + v43)) + ("foo bar" + v43)) {
                    case "\n\t":
                        v34 = (v34 / (((v34 / ((genFn0()) | 1))) | 1))
                    case "":
                        v35 = v35
                    }
                    _ = v61
                    _ = v62
                } else {
                    v43 = ((v43 + v43) + v43)
                    var v63 int16 = 0
                    _ = v63
                }
            }
        case 32767:
            if ((v43 + "foo bar") > (v43 + v43)) {
                for fi64 := 0; fi64 < 5; fi64++ {
                    var v65 float32 = 0.1
                    var v66 string = v43
                    _ = fi64
                    _ = v65
                    _ = v66
                }
                v35 = false
                if (genFn0() >= 457) {
                    var v67 float64 = 0.0
                    _ = v67
                } else {
                    v6 = append(v6, (v43 + ""))
                    if ("a" != v43) {
                        var v68 int8 = 100
                        _ = v68
                    }
                    v37 = v37
                }
            }
        case -2147483648:
            var v69 float64 = 3.14159
            _ = v69
        default:
            v70 := GenS4{F0: (v34 & genFn0()), F1: 128, F2: 42}
            _ = v70
        }
    case 1:
        var v71 uint64 = 4294967295
        _ = v71
    case -422:
        var v72 string = (v43 + v43)
        _ = v72
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
