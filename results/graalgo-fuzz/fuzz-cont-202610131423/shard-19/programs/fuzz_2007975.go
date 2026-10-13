package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 bool }
func genFn0() int32 {
    var v4 uint16 = 32767
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := GenS1{F0: 42, F1: 32767}
        for fi7 := 0; fi7 < 10; fi7++ {
            if true {
                if (true && (!(true && false))) {
                    fmt.Println(fi7)
                    var v8 int32 = 0
                    _ = v8
                }
                v9 := GenS2{F0: 2147483648, F1: -1}
                v10 := GenS3{F0: "a", F1: 0}
                _ = v10
                _ = v9
            }
            var v11 uint16 = 1
            var v12 uint8 = 204
            var v13 uint = 2
            _ = fi7
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi5
        _ = v6
    }
    switch "hello" {
    case "":
        switch -368 {
        case 32767:
            switch (("foo bar" + ("" + "a")) + "hello") {
            case "hello":
                if (1 == 819) {
                    if false {
                        v4 = 92
                        v14 := GenS4{F0: 2, F1: -1}
                        _ = v14
                    } else {
                        v15 := []int{2147483647, -2147483648}
                        for fi16 := 0; fi16 < 5; fi16++ {
                            v4 += 65535
                            for fi17 := 0; fi17 < 12; fi17++ {
                                switch fi17 {
                                case 127:
                                    v4 += (^(v4))
                                case 309:
                                    v4 -= v4
                                case -1:
                                    var v18 uint8 = 2
                                    _ = v18
                                }
                                switch (^(fi17)) {
                                case -9223372036854775808:
                                    v19 := GenS5{F0: 42}
                                    _ = v19
                                case -860:
                                    v20 := GenS2{F0: 4294967295, F1: 127}
                                    _ = v20
                                case 0:
                                    var v21 string = ("hello" + ("foo bar" + "hello"))
                                    _ = v21
                                }
                                _ = fi17
                            }
                            var v22 string = ("a" + "foo bar")
                            _ = fi16
                            _ = v22
                        }
                        _ = v15
                    }
                } else {
                    v4 = v4
                    fmt.Println(v4)
                }
            case "a":
                v23 := GenS1{F0: 1, F1: v4}
                _ = v23
            case "foo bar":
                for fi24 := 0; fi24 < 8; fi24++ {
                    v4 = v4
                    _ = fi24
                }
            }
        case 32768:
            if (0 > 239) {
                if false {
                    fmt.Println(v4)
                } else {
                    v25 := []int{1, 42, 1, -666}
                    for fi26 := 0; fi26 < 11; fi26++ {
                        switch fi26 {
                        case 128:
                            var v27 uint = 4294967295
                            _ = v27
                        case -507:
                            v4 -= v4
                        default:
                            v28 := GenS3{F0: "\n\t", F1: 127}
                            _ = v28
                        }
                        v4 += 65535
                        _ = fi26
                    }
                    _ = v25
                }
                v4 += v4
                if true {
                    fmt.Println(v4)
                    if false {
                        if ((1 <= -800) && false) {
                            fmt.Println(v4)
                            v4 += (v4 ^ v4)
                        } else {
                            var v29 int64 = 334
                            _ = v29
                        }
                    }
                } else {
                    if (93 != -80) {
                        var v30 int16 = 751
                        _ = v30
                    } else {
                        v4 -= v4
                        v4 += (v4 - v4)
                    }
                    var v31 int64 = -9223372036854775808
                    _ = v31
                }
            } else {
                if (false || false) {
                    v4 = v4
                    v4 = v4
                } else {
                    var v32 uint32 = 42
                    _ = v32
                }
                var v33 float32 = 0.0
                _ = v33
            }
        default:
            if false {
                var v34 string = "hello"
                v35 := GenS2{F0: 0, F1: 2}
                _ = v34
                _ = v35
            }
        }
    case "foo bar":
        var v36 uint = 1
        _ = v36
    case "hello":
        v37 := []int{2147483647, 914, 127}
        _ = v37
    default:
        for fi38 := 0; fi38 < 5; fi38++ {
            v39 := []int{(fi38 / ((fi38) | 1)), fi38, (fi38 - (fi38 * fi38)), -821}
            v40 := GenS4{F0: 2, F1: 42}
            var v41 uint64 = 233
            _ = fi38
            _ = v39
            _ = v40
            _ = v41
        }
    }
    _ = v4
    return -507
}

func main() {
    v1 := []string{"\n\t", (("a" + "\n\t") + "")}
    v2 := []string{"foo bar"}
    var v3 uint32 = 0
    var v42 uint16 = 1
    var v43 uint64 = 0
    if (132 <= 0) {
        switch 42 {
        case -340:
            v3 = (v3 << ((255) & 63))
        case -9223372036854775808:
            v2 = append(v2, "\n\t")
        case -422:
            var v44 int64 = 42
            _ = v44
        default:
            for fi45 := 0; fi45 < 17; fi45++ {
                if (genFn0() >= -2147483648) {
                    var v46 int8 = -45
                    _ = v46
                }
                v2 = append(v2, "a")
                _ = fi45
            }
        }
        var v47 uint16 = (v42 ^ (v42 + v42))
        v48 := GenS5{F0: 46}
        _ = v47
        _ = v48
    } else {
        var v49 uint = 2
        _ = v49
    }
    v42 += v42
    for fi50 := 0; fi50 < 12; fi50++ {
        switch "" {
        case "a":
            var v51 int = (fi50 ^ 163)
            _ = v51
        }
        switch "" {
        case "foo bar":
            v43 = 18446744073709551615
        case "":
            fmt.Println(v3)
        case "a":
            if true {
                switch 42 {
                case -904:
                    var v52 int = (-(fi50))
                    _ = v52
                case 9223372036854775807:
                    v42 = (v42 << ((255) & 63))
                case -1:
                    v53 := GenS5{F0: -128}
                    _ = v53
                default:
                    for fi54 := 0; fi54 < 2; fi54++ {
                        if (-2147483648 > -73) {
                            if len(v2) > 0 {
                                v2[0] = "hello"
                            }
                            if (("foo bar" + "\n\t") >= "\n\t") {
                                if (1 < -9223372036854775808) {
                                    v55 := GenS6{F0: 42, F1: false}
                                    var v56 int16 = int16(v42)
                                    _ = v55
                                    _ = v56
                                }
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        _ = fi54
                    }
                }
                v43 -= 2
                v42 = v42
            } else {
                for fi57 := 0; fi57 < 18; fi57++ {
                    v3 -= v3
                    _ = fi57
                }
                fmt.Println(v43)
            }
        default:
            var v58 float32 = 0.1
            _ = v58
        }
        _ = fi50
    }
    switch (("hello" + "foo bar") + "foo bar") {
    case "a":
        if (((v42 / (((v42 >> ((0) & 63))) | 1)) == (v42 - v42)) || true) {
            if len(v2) > 0 {
                v2[0] = ("foo bar" + "foo bar")
            }
            var v59 bool = false
            _ = v59
        }
    case "foo bar":
        fmt.Println(v3)
    case "hello":
        fmt.Println(v42, v3, v42)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "a" {
    case "\n\t":
        var v60 int32 = 42
        _ = v60
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = ((("\n\t" + "") + "foo bar") + ("a" + ("a" + ("hello" + (("\n\t" + ((("foo bar" + "hello") + ("\n\t" + ("foo bar" + "\n\t"))) + ((("hello" + "hello") + "") + ("a" + "")))) + "foo bar")))))
        }
    case "":
        var v61 int = 32768
        _ = v61
    }
    var v62 string = ("a" + ((("a" + (("hello" + (("a" + ("a" + "\n\t")) + ("foo bar" + "\n\t"))) + "hello")) + "hello") + ("" + "hello")))
    _ = len(v2)
    v63 := GenS4{F0: (v43 << ((21) & 63)), F1: 42}
    switch 32768 {
    case -1:
        for fi64 := 0; fi64 < 8; fi64++ {
            for fi65 := 0; fi65 < 18; fi65++ {
                if len(v2) > 0 {
                    v2[0] = v62
                }
                _ = fi65
            }
            _ = fi64
        }
    default:
        var v66 int64 = 9223372036854775807
        _ = v66
    }
    var v67 uint8 = 128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v62:", v62)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
    _ = v62
    _ = v63
    _ = v67
}
