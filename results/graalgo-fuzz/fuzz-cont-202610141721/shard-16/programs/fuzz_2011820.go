package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 uint; F2 int8 }
type GenS5 struct { F0 uint32; F1 uint; F2 int32 }

func main() {
    var v1 int16 = -32768
    var v2 uint8 = 127
    var v3 bool = (0 >= 2147483647)
    v4 := make(map[uint32]uint32)
    var v5 string = "hello"
    var v6 bool = true
    var v7 bool = false
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := GenS1{F0: (!(127 < 127))}
            if (597 >= 0) {
                var v11 int32 = 42
                var v12 int64 = 1
                _ = v11
                _ = v12
            } else {
                delete(v4, 1)
                if (2 < 65535) {
                    v5 = v5
                } else {
                    v13 := []int{(^(fi9)), fi9}
                    _ = v13
                }
                var v14 int16 = (v1 + (^(v1)))
                _ = v14
            }
            if v3 {
                if (372 == 18446744073709551615) {
                    if (false || (2147483648 < 257)) {
                        v15 := []int{}
                        _ = v15
                    } else {
                        if v6 {
                            v16 := GenS1{F0: ((0.0 * (3.14159 + 1.5)) >= (3.14159 / 0.1))}
                            _ = v16
                        }
                        if v7 {
                            v17 := GenS2{F0: "", F1: ((v5 + v5) + "\n\t")}
                            v6 = (v3 && v3)
                            fmt.Println(v2, fi9, fi9)
                            if ((v5 + "foo bar") < v5) {
                                if v3 {
                                    v18 := []int{fi8, (fi9 / ((fi9) | 1))}
                                    var v19 int32 = 1
                                    _ = v18
                                    _ = v19
                                } else {
                                    v20 := GenS3{F0: 42}
                                    v21 := GenS1{F0: v6}
                                    fmt.Println(fi9, v2)
                                    _ = v20
                                    _ = v21
                                }
                            }
                            v3 = true
                            _ = v17
                        }
                    }
                } else {
                    v22 := GenS1{F0: false}
                    v23 := v4[925]
                    var v24 int32 = 2147483647
                    _ = v22
                    _ = v23
                    _ = v24
                }
                var v25 int64 = 9223372036854775807
                v6 = v7
                var v26 int32 = -1
                _ = v25
                _ = v26
            } else {
                v7 = false
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    if ((v2 ^ (v2 + (v2 * 42))) <= (v2 - 42)) {
        fmt.Println(v5)
        switch (("" + (v5 + "")) + ("" + v5)) {
        case "":
            if false {
                v27 := GenS3{F0: 42}
                for fi28 := 0; fi28 < 6; fi28++ {
                    for fi29 := 0; fi29 < 14; fi29++ {
                        var v30 float32 = 0.0
                        v31 := []int{128}
                        v30 = v30
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    fmt.Println(fi28, v6)
                    _ = fi28
                }
                var v32 int8 = 42
                _ = v27
                _ = v32
            }
        default:
            v4[657] = 4294967295
        }
    }
    v4[1] = uint32(v2)
    v3 = ((0 == 42) && (false && (v6 || false)))
    fmt.Println(v1, v5)
    v33 := []int{-831, -664}
    _ = len(v33)
    var v34 bool = (false && true)
    for fi35 := 0; fi35 < 14; fi35++ {
        if v6 {
            v6 = v3
            for fi36 := 0; fi36 < 9; fi36++ {
                v4[84] = 122
                var v37 uint8 = (v2 ^ 226)
                switch fi36 {
                case 0:
                    switch (fi35 + (-(fi36))) {
                    case 9223372036854775807:
                        _ = len(v4)
                    }
                case 1:
                    var v38 int64 = -173
                    _ = v38
                case 90:
                    v39 := GenS2{F0: v5, F1: (((v5 + v5) + (((v5 + v5) + "foo bar") + "")) + "foo bar")}
                    _ = v39
                default:
                    fmt.Println(v34, fi36)
                }
                _ = fi36
                _ = v37
            }
            if (!false) {
                switch fi35 {
                case 32768:
                    v1 = v1
                case 357:
                    var v40 float64 = 0.0
                    _ = v40
                case 1:
                    for fi41 := 0; fi41 < 9; fi41++ {
                        _ = len(v4)
                        var v42 int16 = (v1 << ((199) & 63))
                        v33 = append(v33, -875)
                        v43 := GenS4{F0: (v5 + ("a" + v5)), F1: 42, F2: 0}
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                default:
                    for fi44 := 0; fi44 < 8; fi44++ {
                        _ = len(v4)
                        _ = fi44
                    }
                }
            }
        } else {
            for fi45 := 0; fi45 < 18; fi45++ {
                var v46 bool = (v7 || true)
                _ = fi45
                _ = v46
            }
        }
        fmt.Println(v6, v3, v6)
        var v47 int8 = 1
        _ = fi35
        _ = v47
    }
    if v6 {
        v48 := GenS1{F0: v34}
        _ = v48
    }
    var v49 int8 = 122
    for fi50 := 0; fi50 < 5; fi50++ {
        v5 = v5
        var v51 uint8 = 246
        v7 = v6
        _ = fi50
        _ = v51
    }
    if len(v33) >= 2 {
        v33 = v33[1:len(v33)]
    }
    if (1 > 255) {
        v33 = append(v33, 0)
    }
    for fi52 := 0; fi52 < 5; fi52++ {
        v53 := GenS4{F0: v5, F1: 18446744073709551615, F2: v49}
        _ = fi52
        _ = v53
    }
    for fi54 := 0; fi54 < 1; fi54++ {
        v55 := GenS5{F0: 2, F1: 0, F2: -1}
        for fi56 := 0; fi56 < 15; fi56++ {
            var v57 int64 = -842
            v1 -= -683
            _ = len(v4)
            _ = fi56
            _ = v57
        }
        _ = fi54
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
