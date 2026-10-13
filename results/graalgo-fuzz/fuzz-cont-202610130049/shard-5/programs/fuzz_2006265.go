package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 bool }
type GenS2 struct { F0 int16; F1 int32; F2 uint32 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint16, p1 bool, p2 uint) int16 {
    v5 := []int{}
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := GenS1{F0: p0, F1: (("a" + "\n\t") + ("\n\t" + "a")), F2: p1}
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 float32 = 0.1
            p2 = (^(p2))
            _ = fi9
            _ = v10
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 32767
}

func main() {
    var v1 float32 = 1.5
    var v2 int32 = -1
    var v3 int8 = -128
    v4 := []int{2147483647}
    var v11 uint64 = 4294967295
    v12 := GenS1{F0: 32767, F1: "\n\t", F2: true}
    v11 = (v11 + 305)
    var v13 uint32 = 2147483647
    if true {
        fmt.Println(v3, v13, v3)
    } else {
        v14 := GenS1{F0: 65535, F1: "hello", F2: false}
        _ = v14
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 uint64 = (v11 >> ((2) & 63))
        v1 = 1e10
        switch fi15 {
        case -2147483648:
            v11 = (^(v11))
        case 540:
            var v17 int16 = 42
            _ = v17
        case 4294967295:
            _ = len(v4)
        }
        _ = fi15
        _ = v16
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "\n\t" {
    case "a":
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 string = "foo bar"
            _ = fi18
            _ = v19
        }
    case "":
        for fi20 := 0; fi20 < 10; fi20++ {
            if (!false) {
                switch "a" {
                case "a":
                    for fi21 := 0; fi21 < 18; fi21++ {
                        fmt.Println(v3, v1)
                        _ = fi21
                    }
                case "\n\t":
                    for fi22 := 0; fi22 < 3; fi22++ {
                        var v23 string = "a"
                        v23 = v23
                        _ = fi22
                        _ = v23
                    }
                case "":
                    switch -596 {
                    case -481:
                        if (false || (79 >= (v3 * v3))) {
                            v24 := GenS1{F0: 489, F1: "foo bar", F2: true}
                            if ((!(54 < v3)) || (!(255 == 2))) {
                                v25 := GenS1{F0: 2, F1: "\n\t", F2: (true || false)}
                                v1 = float32(-2.25)
                                _ = v25
                            }
                            if ((true && (!(!false))) || (v2 != v2)) {
                                if true {
                                    if len(v4) > 0 {
                                        v4[0] = fi20
                                    }
                                    v26 := GenS1{F0: 65535, F1: ("foo bar" + (("\n\t" + "") + "foo bar")), F2: (69 == (-(v3)))}
                                    _ = v26
                                }
                                v27 := GenS1{F0: 65535, F1: ("hello" + "hello"), F2: ((v11 | v11) == v11)}
                                _ = v27
                            }
                            _ = v24
                        } else {
                            switch "" {
                            case "\n\t":
                                var v28 int = (fi20 ^ fi20)
                                _ = v28
                            case "hello":
                                var v29 uint = 0
                                _ = v29
                            }
                        }
                    default:
                        v3 = (v3 % (((v3 + v3)) | 1))
                    }
                }
                for fi30 := 0; fi30 < 10; fi30++ {
                    var v31 int64 = -9223372036854775808
                    if len(v4) > 0 {
                        v4[0] = fi30
                    }
                    v32 := GenS1{F0: 1, F1: "", F2: false}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                v33 := GenS1{F0: 197, F1: (("" + "\n\t") + (("hello" + ((("foo bar" + "hello") + "a") + ("\n\t" + "a"))) + "foo bar")), F2: (!(!((1 != -1) || true)))}
                _ = v33
            } else {
                if true {
                    v34 := GenS1{F0: 335, F1: "a", F2: (528 < 32767)}
                    v4 = append(v4, -767)
                    _ = v34
                } else {
                    fmt.Println(v1)
                    v35 := GenS1{F0: 76, F1: "\n\t", F2: (!true)}
                    _ = v35
                }
                if (1.5 != 3.14159) {
                    v1 -= v1
                    fmt.Println(v13, v1, v3)
                    var v36 int64 = -9223372036854775808
                    _ = v36
                }
                for fi37 := 0; fi37 < 17; fi37++ {
                    v38 := GenS1{F0: 155, F1: "a", F2: (42 <= -2147483648)}
                    _ = fi37
                    _ = v38
                }
            }
            if true {
                for fi39 := 0; fi39 < 9; fi39++ {
                    if (42 <= 0) {
                        switch 42 {
                        case -1:
                            v40 := GenS1{F0: 1, F1: "a", F2: false}
                            _ = v40
                        }
                        var v41 float64 = 3.14159
                        fmt.Println(fi39, v11, fi20)
                        var v42 string = "hello"
                        _ = v41
                        _ = v42
                    }
                    v13 += (v13 ^ v13)
                    if (v3 >= v3) {
                        switch (fi39 >> ((65535) & 63)) {
                        case 1:
                            var v43 int16 = 32767
                            _ = v43
                        case -1:
                            var v44 uint = 4294967295
                            _ = v44
                        case -265:
                            fmt.Println(fi20, fi39, fi20)
                        }
                        var v45 uint = 255
                        _ = v45
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(v1, v1)
                        switch (fi39 / (((fi39 << ((4294967295) & 63))) | 1)) {
                        case 9223372036854775807:
                            switch 128 {
                            case -980:
                                v4 = append(v4, fi20)
                            case 32767:
                                var v46 int64 = 42
                                _ = v46
                            default:
                                v47 := GenS2{F0: 293, F1: (-(v2)), F2: v13}
                                _ = v47
                            }
                        case -2147483648:
                            v48 := GenS3{F0: 32767, F1: -1}
                            _ = v48
                        }
                    }
                    _ = fi39
                }
                if true {
                    v49 := GenS4{F0: v2}
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v50 int8 = (v3 & v3)
                    _ = v49
                    _ = v50
                }
                v51 := GenS1{F0: 883, F1: "foo bar", F2: false}
                _ = v51
            }
            v2 += -1
            _ = fi20
        }
    case "\n\t":
        v11 = v11
    }
    if (!(false && (!(238 < 128)))) {
        v52 := GenS1{F0: 42, F1: "\n\t", F2: true}
        _ = v52
    }
    for fi53 := 0; fi53 < 13; fi53++ {
        if true {
            for fi54 := 0; fi54 < 2; fi54++ {
                _ = len(v4)
                v13 -= (^(v13))
                var v55 int32 = v2
                _ = fi54
                _ = v55
            }
            v56 := GenS3{F0: -794, F1: 42}
            _ = v56
        } else {
            var v57 string = ("\n\t" + "\n\t")
            _ = v57
        }
        _ = fi53
    }
    fmt.Println(v1, v11)
    var v58 uint32 = (v13 + v13)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v58
}
