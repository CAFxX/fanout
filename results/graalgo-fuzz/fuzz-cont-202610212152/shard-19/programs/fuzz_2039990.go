package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 bool; F2 int16 }
type GenS6 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS7 struct { F0 int64; F1 bool; F2 int16 }
type GenS8 struct { F0 int16; F1 uint16 }
type GenS9 struct { F0 uint16; F1 int32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    var v3 bool = (42 <= 42)
    v3 = false
    v4 := GenS1{F0: "hello"}
    var v5 int16 = 536
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 int8 = -128
        var v8 uint = 1
        if false {
            if (fi6 == (fi6 - (fi6 - (fi6 - fi6)))) {
                fmt.Println(v5)
                for fi9 := 0; fi9 < 13; fi9++ {
                    v10 := GenS2{F0: "", F1: 2, F2: 0}
                    v3 = v3
                    _ = fi9
                    _ = v10
                }
            } else {
                switch "hello" {
                case "\n\t":
                    v11 := []int{(-(fi6)), fi6}
                    _ = v11
                case "":
                    v12 := []int{}
                    _ = v12
                }
                for fi13 := 0; fi13 < 13; fi13++ {
                    switch fi13 {
                    case -2147483648:
                        v8 += v8
                    }
                    var v14 int = (fi6 >> ((v8) & 63))
                    v8 -= (^(v8))
                    _ = fi13
                    _ = v14
                }
            }
            if (v7 < (v7 * (-(v7)))) {
                if true {
                    fmt.Println(v8)
                    if (127 < 42) {
                        for fi15 := 0; fi15 < 20; fi15++ {
                            if v3 {
                                v16 := GenS1{F0: "\n\t"}
                                v17 := GenS2{F0: ("hello" + "\n\t"), F1: 42, F2: 41}
                                _ = v16
                                _ = v17
                            } else {
                                switch fi15 {
                                case 675:
                                    if (2147483647 < 2147483647) {
                                        delete(v1, "hello")
                                        v5 += (^(v5))
                                        var v18 uint = v8
                                        var v19 int = fi15
                                        fmt.Println(fi6, v5)
                                        _ = v18
                                        _ = v19
                                    }
                                case 9223372036854775807:
                                    var v20 int64 = 1
                                    _ = v20
                                }
                            }
                            v21 := v2[""]
                            _ = fi15
                            _ = v21
                        }
                    } else {
                        switch fi6 {
                        case 42:
                            var v22 float32 = 3.14159
                            _ = v22
                        case 127:
                            fmt.Println(v3, v3, v7)
                        default:
                            fmt.Println(v8)
                        }
                        if v3 {
                            if (1.5 <= 0.1) {
                                for fi23 := 0; fi23 < 15; fi23++ {
                                    switch (fi6 << ((18446744073709551615) & 63)) {
                                    case -1:
                                        v24 := GenS3{F0: 753}
                                        _ = v24
                                    case 127:
                                        v7 = v7
                                    }
                                    v7 += v7
                                    var v25 int8 = (v7 ^ v7)
                                    _ = fi23
                                    _ = v25
                                }
                                v26 := []int{}
                                v8 -= v8
                                _ = v26
                            }
                            var v27 uint64 = 795
                            _ = v27
                        } else {
                            fmt.Println(v3)
                            v5 += (v5 & (v5 * 891))
                            v28 := []int{(fi6 & 1), (fi6 ^ fi6), fi6, -9223372036854775808}
                            _ = v28
                        }
                        if (v3 || (!("\n\t" <= "foo bar"))) {
                            var v29 uint64 = 929
                            fmt.Println(fi6, fi6, v8)
                            _ = v29
                        }
                    }
                } else {
                    var v30 uint64 = 509
                    _ = v30
                }
                v31 := []int{fi6, fi6}
                _ = v31
            } else {
                v3 = (317 != -1)
            }
            var v32 int64 = 0
            _ = v32
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    if (1 < 42) {
        v5 = v5
    }
    fmt.Println(v5, v3)
    v5 += 32767
    var v33 bool = v3
    switch "\n\t" {
    case "":
        var v34 uint8 = 180
        _ = v34
    }
    v35 := GenS4{F0: (-2147483648 < 0)}
    fmt.Println(v5, v3, v3)
    switch "foo bar" {
    case "foo bar":
        for fi36 := 0; fi36 < 18; fi36++ {
            for fi37 := 0; fi37 < 11; fi37++ {
                v38 := GenS5{F0: 915, F1: v33, F2: v5}
                _ = fi37
                _ = v38
            }
            for fi39 := 0; fi39 < 12; fi39++ {
                v40 := GenS6{F0: v3, F1: 2, F2: 35}
                fmt.Println(v5, v3, v33)
                v3 = v3
                if (-1 != 0) {
                    v41 := v1["hello"]
                    v5 -= v5
                    _ = v41
                } else {
                    var v42 int16 = v5
                    v43 := []int{32767, fi39}
                    switch ((("foo bar" + "\n\t") + "\n\t") + "hello") {
                    case "hello":
                        if (0 != 1) {
                            v44 := GenS7{F0: 42, F1: v3, F2: (v5 ^ (^(v5)))}
                            if v33 {
                                _ = len(v2)
                                v5 += 48
                                fmt.Println(v33, v42, v33)
                                _ = len(v2)
                                v42 -= (v42 * 615)
                            }
                            _ = v44
                        } else {
                            var v45 float64 = ((3.14159 + 3.14159) * (-2.25 / 3.14159))
                            switch ("a" + "hello") {
                            case "hello":
                                v46 := v1[""]
                                _ = v46
                            default:
                                fmt.Println(v3)
                            }
                            var v47 uint16 = 1
                            var v48 uint = 328
                            var v49 float32 = 1e10
                            _ = v45
                            _ = v47
                            _ = v48
                            _ = v49
                        }
                    case "a":
                        v50 := GenS8{F0: (v42 + (v5 / (((v42 ^ (-(v5)))) | 1))), F1: 65535}
                        _ = v50
                    case "":
                        _ = len(v1)
                    }
                    _ = v42
                    _ = v43
                }
                v51 := GenS9{F0: 255, F1: 42}
                _ = fi39
                _ = v40
                _ = v51
            }
            _ = fi36
        }
    case "a":
        switch 4294967295 {
        case 635:
            fmt.Println(v3, v3)
        case 562:
            switch 299 {
            case 2147483647:
                var v52 float32 = float32(1.5)
                _ = v52
            default:
                v33 = v33
            }
        }
    }
    var v53 uint = 0
    for fi54 := 0; fi54 < 6; fi54++ {
        var v55 string = "\n\t"
        v3 = v33
        v56 := []int{(-(fi54)), 903}
        _ = fi54
        _ = v55
        _ = v56
    }
    if true {
        v5 += (v5 - 32767)
        v57 := []int{338, -348}
        v3 = v33
        _ = v57
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v4
    _ = v5
    _ = v53
}
