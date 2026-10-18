package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int32; F2 int32 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 string; F1 int8; F2 int }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 int }
func genFn0(p0 int32, p1 int, p2 uint64) int64 {
    p2 -= 4294967295
    if false {
        fmt.Println(p2, p2)
    } else {
        switch "hello" {
        case "\n\t":
            for fi7 := 0; fi7 < 4; fi7++ {
                p1 = (fi7 >> ((1) & 63))
                _ = fi7
            }
        case "hello":
            switch ((("hello" + ("" + "")) + ("hello" + "hello")) + ("a" + "\n\t")) {
            case "foo bar":
                p1 += 9223372036854775807
            default:
                var v8 bool = true
                _ = v8
            }
        default:
            v9 := []int{}
            _ = v9
        }
    }
    v10 := []int{p1, p1, 548, p1}
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := GenS1{F0: -2147483648}
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: "a", F1: (p0 ^ (p0 >> ((42) & 63))), F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v13
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    v4 := []int{32767, 150, 401}
    var v5 int8 = -1
    v6 := []int{-9223372036854775808, -651}
    switch "a" {
    case "hello":
        v5 = v5
    case "foo bar":
        var v14 int32 = -1
        _ = v14
    default:
        switch -777 {
        case 4294967295:
            switch ((("foo bar" + "foo bar") + "foo bar") + ((("" + "a") + "foo bar") + ("\n\t" + ("" + "\n\t")))) {
            case "\n\t":
                v15 := v1[688]
                _ = v15
            }
        case 793:
            v5 -= 42
        default:
            for fi16 := 0; fi16 < 13; fi16++ {
                fmt.Println(v5, fi16)
                _ = fi16
            }
        }
    }
    fmt.Println(v5)
    for fi17 := 0; fi17 < 7; fi17++ {
        for fi18 := 0; fi18 < 4; fi18++ {
            if false {
                v5 = 0
                if (!(0 != 65535)) {
                    fmt.Println(fi17)
                    v5 = v5
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                fmt.Println(fi17, v5)
            } else {
                v19 := v2["\n\t"]
                v5 = (v5 | v5)
                _ = v19
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v20 int = (fi18 << ((18446744073709551615) & 63))
            _ = fi18
            _ = v20
        }
        delete(v3, "a")
        switch (fi17 << ((0) & 63)) {
        case -2147483648:
            fmt.Println(fi17, fi17)
        case 128:
            v5 += v5
        case -183:
            switch "\n\t" {
            case "a":
                fmt.Println(v5)
            default:
                switch "" {
                case "a":
                    for fi21 := 0; fi21 < 9; fi21++ {
                        v5 = 0
                        var v22 uint = 44
                        var v23 int16 = 1000
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                case "foo bar":
                    if (1 != 255) {
                        fmt.Println(fi17, v5)
                        if true {
                            switch fi17 {
                            case 0:
                                switch (fi17 % (((fi17 + fi17)) | 1)) {
                                case 32767:
                                    v24 := GenS2{F0: "\n\t", F1: -1, F2: -2147483648}
                                    _ = v24
                                case -373:
                                    var v25 string = "a"
                                    _ = v25
                                case -1:
                                    if true {
                                        var v26 int16 = -32768
                                        v27 := GenS2{F0: "a", F1: -1, F2: -1}
                                        fmt.Println(fi17, fi17)
                                        _ = v26
                                        _ = v27
                                    }
                                default:
                                    delete(v3, "")
                                }
                            case 32767:
                                for fi28 := 0; fi28 < 8; fi28++ {
                                    if (!false) {
                                        var v29 float64 = (0.1 - 1e10)
                                        var v30 int64 = -657
                                        if ((^(v30)) >= -963) {
                                            var v31 float32 = 0.1
                                            v32 := GenS3{F0: 8, F1: 157}
                                            fmt.Println(fi17)
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            if (1 <= 42) {
                                                v33 := GenS4{F0: "hello", F1: (v5 * v5), F2: (fi17 | (^(fi17)))}
                                                v30 += (v30 | v30)
                                                var v34 uint16 = 391
                                                _ = v33
                                                _ = v34
                                            } else {
                                                _ = len(v6)
                                                v31 = ((1e10 * float32(((0.1 + v29) * ((float64(-2.25) * (v29 * v29)) - float64((v31 + ((0.1 + -2.25) - ((((v31 - 0.1) - -2.25) / v31) * 3.14159)))))))) / v31)
                                                v35 := GenS4{F0: ("hello" + ("\n\t" + "a")), F1: v5, F2: (^(fi17))}
                                                _ = v35
                                            }
                                            _ = v31
                                            _ = v32
                                        } else {
                                            var v36 uint8 = 1
                                            var v37 int = -2147483648
                                            fmt.Println(v29, v37)
                                            _ = v36
                                            _ = v37
                                        }
                                        _ = v29
                                        _ = v30
                                    } else {
                                        switch 4294967295 {
                                        case 42:
                                            if (!(1 >= 2147483648)) {
                                                v38 := GenS5{F0: 255, F1: (v5 >= v5)}
                                                if false {
                                                    delete(v1, 2)
                                                } else {
                                                    v5 += (v5 >> ((65535) & 63))
                                                    if (((fi17 / (((fi28 * fi28)) | 1)) <= fi28) && (1e10 <= (0.1 / 1.0))) {
                                                        var v39 float64 = ((0.1 + float64(0.1)) - -2.25)
                                                        if false {
                                                            var v40 float32 = (-2.25 / (float32(v39) + 1e10))
                                                            switch "foo bar" {
                                                            case "hello":
                                                                v40 = 0.1
                                                            case "a":
                                                                delete(v3, "\n\t")
                                                            case "\n\t":
                                                                var v41 uint8 = 128
                                                                _ = v41
                                                            }
                                                            _ = v40
                                                        }
                                                        switch "foo bar" {
                                                        case "hello":
                                                            if false {
                                                                var v42 int64 = 42
                                                                v43 := GenS4{F0: ("" + "a"), F1: v5, F2: (^(fi28))}
                                                                _ = v42
                                                                _ = v43
                                                            } else {
                                                                v44 := GenS3{F0: (fi17 * (fi17 * fi28)), F1: 42}
                                                                var v45 int8 = (v5 & v5)
                                                                var v46 int = (fi17 & fi17)
                                                                var v47 int8 = (^(v5))
                                                                _ = v44
                                                                _ = v45
                                                                _ = v46
                                                                _ = v47
                                                            }
                                                        case "foo bar":
                                                            switch fi28 {
                                                            case 42:
                                                                v48 := GenS6{F0: (fi17 & fi28)}
                                                                _ = v48
                                                            case -424:
                                                                var v49 int8 = v5
                                                                v39 -= (-2.25 - (((((1e10 / v39) / v39) / v39) + v39) + (v39 / (((((v39 / 1.0) + v39) - v39) + ((3.14159 / v39) - v39)) + float64(3.14159)))))
                                                                _ = v49
                                                            case 473:
                                                                v4 = append(v4, 1)
                                                            default:
                                                                var v50 int = (-(fi28))
                                                                _ = v50
                                                            }
                                                        case "\n\t":
                                                            v5 -= (v5 / (((v5 << ((255) & 63))) | 1))
                                                        default:
                                                            if (!(922 >= 42)) {
                                                                v51 := GenS1{F0: (-(fi17))}
                                                                _ = v51
                                                            } else {
                                                                v52 := GenS3{F0: -383, F1: 836}
                                                                fmt.Println(v39, fi28)
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                _ = v52
                                                            }
                                                        }
                                                        _ = v39
                                                    } else {
                                                        switch ("a" + "\n\t") {
                                                        case "hello":
                                                            v5 += v5
                                                        case "a":
                                                            v53 := v3[("" + (((("a" + "a") + "a") + "\n\t") + ""))]
                                                            _ = v53
                                                        }
                                                    }
                                                }
                                                _ = v38
                                            } else {
                                                var v54 int = -2147483648
                                                v54 = fi17
                                                v55 := v2[(("\n\t" + "a") + "a")]
                                                v54 -= 128
                                                _ = v54
                                                _ = v55
                                            }
                                        case 32768:
                                            switch (fi17 & fi28) {
                                            case 0:
                                                if (42 >= 799) {
                                                    v5 += 111
                                                    v5 += (v5 - (^(v5)))
                                                } else {
                                                    v56 := GenS1{F0: (fi17 | fi17)}
                                                    _ = len(v6)
                                                    v5 = -85
                                                    _ = v56
                                                }
                                            case 127:
                                                fmt.Println(v5)
                                            default:
                                                v5 = 42
                                            }
                                        default:
                                            if false {
                                                v57 := v3[""]
                                                _ = v57
                                            }
                                        }
                                    }
                                    _ = fi28
                                }
                            case 704:
                                for fi58 := 0; fi58 < 19; fi58++ {
                                    var v59 uint = 488
                                    fmt.Println(fi17, v5)
                                    v59 += 751
                                    _ = fi58
                                    _ = v59
                                }
                            }
                            var v60 uint8 = 0
                            _ = v60
                        } else {
                            v5 += (v5 << ((42) & 63))
                            if len(v4) > 0 {
                                v4[0] = (-(fi17))
                            }
                            var v61 uint32 = 42
                            _ = v61
                        }
                        if false {
                            switch (fi17 << ((1) & 63)) {
                            case 127:
                                v5 -= (v5 - v5)
                            }
                            fmt.Println(v5)
                            var v62 float32 = float32((0.0 - 0.1))
                            _ = v62
                        }
                    }
                case "\n\t":
                    v5 -= v5
                }
            }
        default:
            v5 = v5
        }
        _ = fi17
    }
    if true {
        var v63 uint32 = 2147483647
        _ = v63
    } else {
        for fi64 := 0; fi64 < 19; fi64++ {
            v5 += v5
            var v65 int16 = 1
            _ = fi64
            _ = v65
        }
        var v66 uint = 4294967295
        var v67 float32 = float32(0.0)
        _ = v66
        _ = v67
    }
    if true {
        if true {
            fmt.Println(v5)
        }
        _ = len(v1)
        switch 4294967295 {
        case -1:
            var v68 int32 = 42
            _ = v68
        case 128:
            _ = len(v1)
        default:
            fmt.Println(v5)
        }
    } else {
        fmt.Println(v5)
    }
    v69 := GenS5{F0: 2, F1: (42 >= 9223372036854775807)}
    v70 := GenS5{F0: 255, F1: (!(32767 != 255))}
    v2["foo bar"] = 4294967295
    if false {
        delete(v2, "foo bar")
    }
    v71 := GenS2{F0: "hello", F1: 42, F2: 656}
    switch "\n\t" {
    case "a":
        delete(v3, ("\n\t" + ("foo bar" + "foo bar")))
    }
    for fi72 := 0; fi72 < 5; fi72++ {
        v5 = 127
        _ = fi72
    }
    v73 := v3["a"]
    v74 := GenS6{F0: (v73 ^ 4294967295)}
    v73 -= v73
    var v75 int64 = 42
    for fi76 := 0; fi76 < 6; fi76++ {
        delete(v2, (("\n\t" + "a") + (("\n\t" + "") + "\n\t")))
        _ = fi76
    }
    switch ("" + "hello") {
    case "":
        v5 = (v5 % ((-62) | 1))
    case "a":
        if ("" < "foo bar") {
            for fi77 := 0; fi77 < 5; fi77++ {
                var v78 float64 = (3.14159 * (1e10 + 0.1))
                _ = fi77
                _ = v78
            }
            v75 += v75
            var v79 uint = 42
            _ = v79
        } else {
            v5 = (^(v5))
            var v80 uint8 = 0
            var v81 int = 9223372036854775807
            _ = v80
            _ = v81
        }
    case "hello":
        v75 += (v75 + v75)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v73:", v73)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v70
    _ = v71
    _ = v73
    _ = v74
    _ = v75
}
