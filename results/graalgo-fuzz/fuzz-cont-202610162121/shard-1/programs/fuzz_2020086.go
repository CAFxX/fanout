package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 string; F1 bool; F2 bool }
type GenS5 struct { F0 int64 }

func main() {
    var v1 int32 = 1
    var v2 bool = false
    var v3 int64 = -1
    var v4 string = "foo bar"
    v5 := []int{4294967295, 9223372036854775807, 0}
    for fi6 := 0; fi6 < 4; fi6++ {
        if (0.1 == (1e10 + 1.5)) {
            if ((0.0 + 3.14159) != 1e10) {
                v7 := GenS1{F0: int8(fi6), F1: -128}
                if true {
                    if (!(v2 && v2)) {
                        var v8 uint16 = 130
                        _ = v8
                    } else {
                        for fi9 := 0; fi9 < 9; fi9++ {
                            if (-10 >= 100) {
                                v4 = ""
                            } else {
                                v10 := GenS2{F0: (fi9 & (^(fi6))), F1: 32767}
                                var v11 bool = v2
                                _ = v10
                                _ = v11
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi9
                        }
                        for fi12 := 0; fi12 < 9; fi12++ {
                            v2 = v2
                            if (104 <= 212) {
                                var v13 int32 = -13
                                var v14 string = (v4 + v4)
                                _ = v13
                                _ = v14
                            }
                            _ = fi12
                        }
                        fmt.Println(v3, v4)
                    }
                    v2 = ((v2 || v2) && (9223372036854775807 >= fi6))
                    for fi15 := 0; fi15 < 10; fi15++ {
                        var v16 bool = (!((v3 - -293) >= (-(v3))))
                        v17 := GenS3{F0: 4294967295, F1: (v2 || v16)}
                        if (58 < 0) {
                            if v2 {
                                fmt.Println(v4, v1, fi6)
                            } else {
                                v18 := GenS4{F0: v4, F1: false, F2: (878 > 2)}
                                if v16 {
                                    if len(v5) > 0 {
                                        v5[0] = 0
                                    }
                                    v1 = (v1 * v1)
                                    v3 = (v3 << ((18446744073709551615) & 63))
                                    v19 := GenS1{F0: 42, F1: 100}
                                    _ = v19
                                } else {
                                    var v20 uint16 = 255
                                    fmt.Println(fi6, v20, v3)
                                    _ = v20
                                }
                                v16 = v16
                                _ = v18
                            }
                            var v21 uint16 = 32767
                            _ = v21
                        }
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                } else {
                    var v22 int64 = v3
                    v4 = "hello"
                    _ = v22
                }
                _ = v7
            } else {
                fmt.Println(v1, fi6, v3)
                switch fi6 {
                case 9223372036854775807:
                    for fi23 := 0; fi23 < 4; fi23++ {
                        fmt.Println(v2, v1, fi23)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        fmt.Println(v2, v3)
                        _ = fi23
                    }
                }
            }
            v24 := GenS2{F0: fi6, F1: 42}
            if v2 {
                fmt.Println(v1, v4)
                _ = len(v5)
                if v2 {
                    v1 = v1
                }
            } else {
                if v2 {
                    _ = len(v5)
                }
                v4 = (((v4 + v4) + ("a" + ("foo bar" + v4))) + "foo bar")
                var v25 bool = v2
                _ = v25
            }
            _ = v24
        }
        _ = fi6
    }
    var v26 string = "foo bar"
    fmt.Println(v1, v4)
    var v27 bool = (!v2)
    v28 := GenS5{F0: (v3 + v3)}
    fmt.Println(v3, v27, v3)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if false {
        var v29 int8 = -128
        var v30 string = v26
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v29
        _ = v30
    }
    if (-32768 >= 85) {
        v2 = true
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(v26, v4)
    } else {
        var v31 int32 = v1
        switch 665 {
        case -812:
            fmt.Println(v2)
        case -298:
            var v32 bool = v27
            _ = v32
        }
        _ = v31
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        v34 := GenS3{F0: 0, F1: (!(!v2))}
        v35 := GenS1{F0: -97, F1: 42}
        _ = fi33
        _ = v34
        _ = v35
    }
    if v27 {
        var v36 int64 = v3
        var v37 uint64 = 1
        switch v4 {
        case "a":
            switch v4 {
            case "\n\t":
                v36 = v3
            case "":
                switch v26 {
                case "\n\t":
                    v38 := GenS2{F0: 128, F1: 0}
                    _ = v38
                }
            case "hello":
                for fi39 := 0; fi39 < 11; fi39++ {
                    v40 := GenS1{F0: 51, F1: 0}
                    fmt.Println(v4, v4)
                    for fi41 := 0; fi41 < 12; fi41++ {
                        var v42 int = fi39
                        _ = fi41
                        _ = v42
                    }
                    _ = fi39
                    _ = v40
                }
            default:
                var v43 uint16 = 125
                _ = v43
            }
        case "\n\t":
            _ = len(v5)
        case "foo bar":
            if v2 {
                for fi44 := 0; fi44 < 3; fi44++ {
                    v1 += v1
                    v45 := GenS2{F0: fi44, F1: 743}
                    var v46 float64 = 1e10
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            }
        }
        _ = v36
        _ = v37
    } else {
        switch -185 {
        case 118:
            if v27 {
                for fi47 := 0; fi47 < 15; fi47++ {
                    v5 = append(v5, 0)
                    _ = fi47
                }
                if (!((!v2) && (v27 && ((-(v1)) <= -1)))) {
                    v27 = (3.14159 < 0.1)
                    v48 := GenS4{F0: "\n\t", F1: true, F2: false}
                    _ = v48
                }
                switch (v26 + ("a" + "")) {
                case "a":
                    if true {
                        var v49 int = -9223372036854775808
                        var v50 bool = ((v49 ^ 128) < -9223372036854775808)
                        _ = v49
                        _ = v50
                    } else {
                        if v2 {
                            v51 := GenS3{F0: 691, F1: (922 <= 16)}
                            v27 = v2
                            _ = v51
                        }
                    }
                default:
                    var v52 uint8 = 128
                    _ = v52
                }
            } else {
                switch v4 {
                case "hello":
                    for fi53 := 0; fi53 < 13; fi53++ {
                        if (1 <= 127) {
                            for fi54 := 0; fi54 < 20; fi54++ {
                                if true {
                                    if v27 {
                                        _ = len(v5)
                                    }
                                    if v2 {
                                        fmt.Println(v26, v4)
                                        fmt.Println(v4, v27, v27)
                                        var v55 uint16 = 65535
                                        _ = v55
                                    }
                                } else {
                                    fmt.Println(v26, fi54, fi54)
                                    var v56 int8 = 1
                                    _ = v56
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v2 = v27
                                _ = fi54
                            }
                        }
                        _ = len(v5)
                        v3 -= -9223372036854775808
                        _ = fi53
                    }
                }
                v57 := GenS2{F0: 0, F1: 42}
                _ = v57
            }
        case -1:
            v3 += v3
        }
    }
    for fi58 := 0; fi58 < 16; fi58++ {
        var v59 string = (v26 + ((v4 + (v26 + "")) + v26))
        _ = fi58
        _ = v59
    }
    if (!(1 != -128)) {
        var v60 uint16 = 1
        fmt.Println(v2)
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
