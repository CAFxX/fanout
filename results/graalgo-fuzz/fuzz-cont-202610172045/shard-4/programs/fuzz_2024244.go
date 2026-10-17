package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint64 }
func genFn0() uint64 {
    var v7 int = 32768
    switch "a" {
    case "hello":
        switch v7 {
        case 821:
            var v8 string = "foo bar"
            _ = v8
        default:
            v7 = 124
        }
    }
    _ = v7
    return 42
}
func genFn1(p0 uint, p1 string, p2 int64) string {
    if (32767 > 50) {
        var v9 uint = p0
        _ = v9
    } else {
        v10 := []int{32767}
        switch p1 {
        case "a":
            p0 = (p0 / ((p0) | 1))
        default:
            var v11 int = -1
            _ = v11
        }
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = v10
    }
    if true {
        for fi12 := 0; fi12 < 3; fi12++ {
            if true {
                for fi13 := 0; fi13 < 4; fi13++ {
                    var v14 int64 = p2
                    var v15 float32 = 0.1
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            }
            switch (^(fi12)) {
            case 2147483647:
                for fi16 := 0; fi16 < 18; fi16++ {
                    v17 := []int{}
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi12
        }
        var v18 uint16 = 1
        _ = v18
    } else {
        v19 := GenS1{F0: "foo bar", F1: genFn0()}
        for fi20 := 0; fi20 < 7; fi20++ {
            if ((!false) && true) {
                v21 := []int{fi20, fi20}
                _ = v21
            } else {
                v22 := GenS1{F0: p1, F1: 18446744073709551615}
                var v23 float32 = 0.0
                _ = v22
                _ = v23
            }
            var v24 int = (fi20 - fi20)
            _ = fi20
            _ = v24
        }
        p2 += p2
        _ = v19
    }
    switch 4294967295 {
    case -372:
        switch 2147483647 {
        case -158:
            var v25 bool = true
            _ = v25
        case 2147483647:
            switch 9223372036854775807 {
            case 1:
                v26 := []int{298, 128, -9223372036854775808}
                _ = v26
            case 128:
                v27 := []int{}
                _ = v27
            }
        default:
            for fi28 := 0; fi28 < 15; fi28++ {
                v29 := GenS1{F0: p1, F1: 18446744073709551615}
                _ = fi28
                _ = v29
            }
        }
    case 128:
        var v30 uint8 = 1
        _ = v30
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 uint64 = 622
    v2 := make(map[uint32]bool)
    var v3 int64 = -338
    var v4 int16 = 1
    v5 := make(map[string]string)
    var v6 uint16 = 1
    var v31 int64 = (v3 << ((4294967295) & 63))
    fmt.Println(v1, v31)
    var v32 int16 = (v4 << ((18446744073709551615) & 63))
    var v33 uint = 42
    v34 := []int{32768, 9223372036854775807, 128, -333}
    switch "hello" {
    case "\n\t":
        var v35 float32 = 1e10
        _ = v35
    case "":
        switch -678 {
        case 4294967295:
            var v36 uint64 = v1
            _ = v36
        default:
            fmt.Println(v4, v33, v4)
        }
    default:
        if (false || false) {
            if false {
                switch "\n\t" {
                case "":
                    switch (("hello" + "hello") + "foo bar") {
                    case "":
                        v37 := GenS1{F0: "hello", F1: (v1 & genFn0())}
                        _ = v37
                    case "\n\t":
                        v31 -= (v31 << ((v33) & 63))
                    case "hello":
                        if false {
                            var v38 uint64 = 18446744073709551615
                            _ = v38
                        }
                    }
                case "a":
                    var v39 int64 = (v31 + v31)
                    _ = v39
                case "hello":
                    if (!false) {
                        v40 := GenS2{F0: (v1 << ((v33) & 63)), F1: v6, F2: 0}
                        v31 -= 1
                        _ = v40
                    } else {
                        v32 += 42
                        for fi41 := 0; fi41 < 15; fi41++ {
                            v4 = int16(v1)
                            v32 = (v4 << (((v33 << ((v33) & 63))) & 63))
                            for fi42 := 0; fi42 < 18; fi42++ {
                                var v43 int8 = 42
                                if ((("a" + "foo bar") != "hello") || true) {
                                    var v44 int8 = (v43 << (((v33 + uint(v6))) & 63))
                                    _ = v44
                                } else {
                                    var v45 int64 = v3
                                    v31 = v31
                                    _ = v45
                                }
                                _ = fi42
                                _ = v43
                            }
                            _ = fi41
                        }
                        v46 := GenS2{F0: (v1 >> ((v33) & 63)), F1: (v6 >> ((v33) & 63)), F2: 2147483647}
                        _ = v46
                    }
                default:
                    for fi47 := 0; fi47 < 11; fi47++ {
                        if (!(false || false)) {
                            v6 += (v6 - (v6 & 581))
                            if false {
                                if len(v34) > 0 {
                                    v34[0] = fi47
                                }
                            } else {
                                for fi48 := 0; fi48 < 11; fi48++ {
                                    v49 := GenS3{F0: 42, F1: 0}
                                    v50 := GenS1{F0: "hello", F1: (v1 & genFn0())}
                                    v31 -= (v31 - int64(v4))
                                    _ = fi48
                                    _ = v49
                                    _ = v50
                                }
                                var v51 uint = (v33 + v33)
                                switch (((("a" + ((("hello" + "hello") + ("hello" + "\n\t")) + "hello")) + ("a" + ("\n\t" + ""))) + "\n\t") + "foo bar") {
                                case "":
                                    for fi52 := 0; fi52 < 8; fi52++ {
                                        delete(v2, 65535)
                                        _ = fi52
                                    }
                                case "\n\t":
                                    v51 = (v33 * (v51 & (v33 >> ((v33) & 63))))
                                case "foo bar":
                                    for fi53 := 0; fi53 < 14; fi53++ {
                                        v31 -= v3
                                        var v54 int64 = v3
                                        var v55 bool = false
                                        _ = fi53
                                        _ = v54
                                        _ = v55
                                    }
                                }
                                _ = v51
                            }
                            fmt.Println(v31, v1, v32)
                        } else {
                            var v56 float32 = (0.0 * 0.1)
                            switch (fi47 | (fi47 / (((fi47 | (fi47 - int(v33)))) | 1))) {
                            case 750:
                                v56 += 1e10
                            case 0:
                                if true {
                                    var v57 uint8 = 42
                                    v58 := GenS2{F0: (v1 % ((2) | 1)), F1: v6, F2: 2}
                                    for fi59 := 0; fi59 < 1; fi59++ {
                                        var v60 int16 = 32767
                                        fmt.Println(v3, v4)
                                        _ = fi59
                                        _ = v60
                                    }
                                    _ = v57
                                    _ = v58
                                }
                            }
                            _ = v56
                        }
                        _ = fi47
                    }
                }
            }
            switch 32768 {
            case -352:
                for fi61 := 0; fi61 < 7; fi61++ {
                    for fi62 := 0; fi62 < 6; fi62++ {
                        switch fi62 {
                        case -986:
                            var v63 uint32 = 2147483647
                            _ = v63
                        case 0:
                            fmt.Println(v1, fi61, v32)
                        default:
                            _ = len(v34)
                        }
                        var v64 float64 = ((1e10 * float64(float32(((1e10 / -2.25) - ((0.0 * (1e10 - 3.14159)) + (((-2.25 * (3.14159 / 1.5)) + 0.0) + 0.1)))))) / -2.25)
                        var v65 bool = false
                        _ = fi62
                        _ = v64
                        _ = v65
                    }
                    _ = fi61
                }
            case 346:
                if len(v34) >= 2 {
                    v34 = v34[1:len(v34)]
                }
            }
        }
    }
    switch ((("\n\t" + "") + ("a" + "a")) + "a") {
    case "foo bar":
        v66 := GenS4{F0: v1}
        _ = v66
    }
    for fi67 := 0; fi67 < 12; fi67++ {
        v5["a"] = ""
        v3 = (v31 * -9223372036854775808)
        var v68 int64 = int64(v6)
        _ = fi67
        _ = v68
    }
    switch 4294967295 {
    case -1:
        var v69 int32 = 571
        _ = v69
    default:
        delete(v5, "a")
    }
    var v70 int8 = 100
    switch (("foo bar" + "foo bar") + ("\n\t" + "foo bar")) {
    case "hello":
        var v71 uint = 4294967295
        _ = v71
    case "a":
        var v72 int64 = (^(v3))
        _ = v72
    case "":
        var v73 bool = true
        _ = v73
    }
    for fi74 := 0; fi74 < 16; fi74++ {
        v1 = v1
        _ = fi74
    }
    for fi75 := 0; fi75 < 5; fi75++ {
        if len(v34) > 0 {
            v34[0] = fi75
        }
        v4 = v4
        var v76 int8 = v70
        _ = fi75
        _ = v76
    }
    v3 -= v3
    var v77 int32 = -777
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v77
}
