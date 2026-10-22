package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint64; F1 int64; F2 int8 }
type GenS3 struct { F0 int8; F1 int64; F2 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 uint64; F2 int16 }
func genFn0(p0 int64, p1 int32, p2 bool) uint8 {
    switch 1 {
    case 128:
        var v5 uint8 = 127
        _ = v5
    default:
        var v6 int64 = p0
        _ = v6
    }
    v7 := GenS1{F0: 42, F1: (^(p1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 255
}
func genFn1() uint32 {
    var v8 int = -9223372036854775808
    v8 -= 876
    var v9 float64 = 0.0
    if ((v8 & (v8 * 32767)) <= (v8 | (-(v8)))) {
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := GenS2{F0: 846, F1: -1, F2: -128}
            if (!(!(!true))) {
                v12 := GenS3{F0: -128, F1: 768, F2: 571}
                for fi13 := 0; fi13 < 5; fi13++ {
                    var v14 int8 = 42
                    var v15 uint64 = 0
                    if false {
                        fmt.Println(fi13, v15, v9)
                        v15 = v15
                        if (-935 < 0) {
                            v8 = v8
                            v15 -= 1
                            var v16 int64 = -1
                            var v17 uint16 = 32767
                            v18 := []int{32768, (-(fi10)), int(v16), 4294967295}
                            _ = v16
                            _ = v17
                            _ = v18
                        } else {
                            v14 -= (v14 ^ -1)
                            v15 += v15
                        }
                        var v19 float64 = 0.1
                        var v20 int32 = 1
                        v21 := []int{v8, (fi13 & v8)}
                        if (v15 == v15) {
                            v22 := GenS1{F0: (v14 * v14), F1: v20}
                            v21 = append(v21, fi13)
                            _ = v22
                        } else {
                            var v23 int32 = v20
                            v19 -= v19
                            _ = len(v21)
                            _ = v23
                        }
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                        var v24 int64 = -1
                        v14 += v14
                        _ = v19
                        _ = v20
                        _ = v21
                        _ = v24
                    } else {
                        var v25 uint = 1
                        v26 := GenS2{F0: 1, F1: -665, F2: int8(v8)}
                        var v27 uint8 = uint8(v25)
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v9 -= (float64(3.14159) * (((1.5 + float64((1.5 + 1e10))) - v9) + v9))
                _ = v12
            } else {
                v28 := []int{9223372036854775807, (fi10 ^ -399), -9223372036854775808, 727}
                _ = v28
            }
            v29 := []int{}
            _ = fi10
            _ = v11
            _ = v29
        }
        var v30 uint16 = 255
        _ = v30
    }
    var v31 string = ("hello" + ("a" + "foo bar"))
    _ = v31
    _ = v8
    _ = v9
    return 387
}

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    var v4 int = v1
    fmt.Println(v1)
    if (18446744073709551615 < 0) {
        _ = len(v3)
        for fi32 := 0; fi32 < 20; fi32++ {
            var v33 uint32 = 16
            if false {
                delete(v2, 32768)
                if false {
                    switch 32768 {
                    case -228:
                        var v34 int8 = 1
                        _ = v34
                    }
                    switch ("foo bar" + "") {
                    case "":
                        v33 = (v33 - (v33 / ((2147483648) | 1)))
                    case "foo bar":
                        if ((-2.25 == ((((float32(1.5) * float32(((0.0 / float64((0.1 + (-2.25 + 1.5)))) + 0.0))) + 0.1) / 1.0) + (3.14159 - 1.5))) || (!(true && false))) {
                            for fi35 := 0; fi35 < 19; fi35++ {
                                fmt.Println(v4, v4)
                                v2[(fi32 | 4294967295)] = "a"
                                _ = fi35
                            }
                            var v36 int64 = -495
                            _ = v36
                        }
                    default:
                        var v37 uint8 = 127
                        _ = v37
                    }
                    var v38 int32 = 428
                    _ = v38
                } else {
                    v1 += (^(v1))
                }
                v39 := GenS3{F0: -65, F1: -900, F2: -72}
                _ = v39
            } else {
                var v40 string = ("" + "\n\t")
                for fi41 := 0; fi41 < 12; fi41++ {
                    var v42 uint32 = (v33 + 65535)
                    _ = fi41
                    _ = v42
                }
                _ = v40
            }
            if false {
                if (!((v33 + (^(v33))) == (v33 / ((v33) | 1)))) {
                    if ((-(fi32)) == (v1 << ((uint(v33)) & 63))) {
                        v33 += v33
                        var v43 uint = 2
                        _ = v43
                    } else {
                        v44 := v3[""]
                        _ = v44
                    }
                }
                for fi45 := 0; fi45 < 10; fi45++ {
                    var v46 string = "hello"
                    v46 = v46
                    _ = fi45
                    _ = v46
                }
            }
            _ = fi32
            _ = v33
        }
    } else {
        v1 = -761
    }
    v4 += 0
    var v47 int16 = 0
    var v48 uint = 132
    v48 += 18446744073709551615
    if (1 > 255) {
        for fi49 := 0; fi49 < 3; fi49++ {
            switch "" {
            case "foo bar":
                v50 := GenS4{F0: false}
                _ = v50
            default:
                v51 := GenS5{F0: "\n\t", F1: 696, F2: (v47 << ((v48) & 63))}
                _ = v51
            }
            v3["hello"] = 651
            var v52 uint16 = 255
            _ = fi49
            _ = v52
        }
        switch (((("hello" + ("a" + (("foo bar" + "foo bar") + "\n\t"))) + "hello") + ("" + ("hello" + "\n\t"))) + "hello") {
        case "a":
            fmt.Println(v4, v4)
        case "hello":
            if true {
                switch ("a" + "hello") {
                case "":
                    var v53 float64 = -2.25
                    _ = v53
                case "a":
                    fmt.Println(v48, v4, v48)
                }
                if (64 < 1) {
                    v54 := GenS1{F0: 4, F1: 0}
                    fmt.Println(v4, v1)
                    if (("foo bar" > "") || true) {
                        v48 += v48
                        v1 = 2147483647
                        if false {
                            var v55 uint64 = 668
                            switch v4 {
                            case 749:
                                v56 := GenS2{F0: v55, F1: 1, F2: 0}
                                _ = v56
                            }
                            var v57 int = v4
                            _ = v55
                            _ = v57
                        }
                    }
                    _ = v54
                } else {
                    fmt.Println(v1)
                    fmt.Println(v48)
                }
            } else {
                v1 = (v4 | (v4 | (v4 + -589)))
                fmt.Println(v1, v4)
            }
        case "\n\t":
            v48 = v48
        }
    }
    for fi58 := 0; fi58 < 4; fi58++ {
        var v59 int8 = 42
        v47 = v47
        _ = fi58
        _ = v59
    }
    var v60 float32 = (1.5 - 0.0)
    for fi61 := 0; fi61 < 13; fi61++ {
        v47 += 1000
        var v62 int = (fi61 | -1)
        var v63 uint16 = 749
        _ = fi61
        _ = v62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v60
}
