package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 bool }
type GenS2 struct { F0 int; F1 bool; F2 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 uint16; F1 int64; F2 int }
func genFn0(p0 int, p1 int) uint8 {
    p0 += (p1 & 32768)
    p0 += (^(p1))
    _ = p0
    _ = p1
    return 1
}
func genFn1() int {
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint8 = 0
        v9 := []int{fi7, 767, fi7}
        var v10 uint8 = v8
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS1{F0: genFn0((-(fi11)), fi11), F1: 200, F2: true}
        for fi13 := 0; fi13 < 15; fi13++ {
            v14 := GenS1{F0: 224, F1: 127, F2: (true || (!((false || ("" <= ("\n\t" + "hello"))) || false)))}
            v15 := GenS2{F0: fi11, F1: true, F2: (("" + "foo bar") + ("foo bar" + "foo bar"))}
            var v16 int32 = 2147483647
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi11
        _ = v12
    }
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 1.5
    var v3 string = ""
    var v4 float32 = (float32(-2.25) + 1.5)
    v5 := []uint32{4294967295, 4294967295}
    v6 := []uint32{0, 250}
    v3 = v3
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 int16 = 1
        _ = fi17
        _ = v18
    }
    switch -206 {
    case 0:
        v1[875] = v3
    case -485:
        switch (v3 + v3) {
        case "a":
            switch 9223372036854775807 {
            case 128:
                if (2 == 1) {
                    fmt.Println(v4)
                    v19 := GenS3{F0: -1}
                    _ = v19
                } else {
                    if (1 == 1) {
                        v20 := GenS4{F0: 42, F1: 0}
                        _ = v20
                    }
                }
            case -330:
                for fi21 := 0; fi21 < 19; fi21++ {
                    var v22 int64 = 415
                    _ = fi21
                    _ = v22
                }
            }
        default:
            fmt.Println(v4, v4)
        }
    case 220:
        v23 := GenS3{F0: 896}
        _ = v23
    }
    if true {
        fmt.Println(v3, v3)
    }
    v4 = (v4 + v2)
    var v24 bool = false
    v4 -= (1e10 + 0.1)
    var v25 uint8 = 110
    var v26 uint16 = 608
    v26 = (v26 % ((v26) | 1))
    v1[42] = ((("foo bar" + ("hello" + (v3 + (("a" + v3) + v3)))) + v3) + "\n\t")
    var v27 int8 = 127
    if v24 {
        if true {
            for fi28 := 0; fi28 < 17; fi28++ {
                if v24 {
                    v25 = (^(v25))
                    if false {
                        v4 = -2.25
                        _ = len(v1)
                        v3 = "hello"
                    } else {
                        if (uint8(v26) > v25) {
                            if (!((65535 > 2147483648) || v24)) {
                                var v29 int8 = int8(fi28)
                                var v30 string = v3
                                switch ("" + "hello") {
                                case "":
                                    v31 := GenS1{F0: v25, F1: (v25 % (((v25 >> ((938) & 63))) | 1)), F2: v24}
                                    _ = v31
                                case "hello":
                                    v26 += (v26 - 0)
                                default:
                                    v32 := GenS4{F0: 65535, F1: 1}
                                    _ = v32
                                }
                                _ = v29
                                _ = v30
                            } else {
                                v3 = v3
                            }
                            for fi33 := 0; fi33 < 9; fi33++ {
                                if false {
                                    fmt.Println(fi33, v26)
                                }
                                v25 = (v25 + v25)
                                var v34 string = v3
                                v34 = (("hello" + v34) + v34)
                                v27 += -39
                                _ = fi33
                                _ = v34
                            }
                            if (65535 <= 0) {
                                fmt.Println(v27)
                                if (v24 || false) {
                                    var v35 int = (fi28 >> ((253) & 63))
                                    _ = v35
                                } else {
                                    if len(v6) > 0 {
                                        v6[0] = 4294967295
                                    }
                                }
                                if v24 {
                                    v4 -= v4
                                    var v36 uint32 = 638
                                    _ = v36
                                } else {
                                    switch v3 {
                                    case "":
                                        if true {
                                            v25 = v25
                                            v37 := GenS4{F0: 163, F1: 42}
                                            switch fi28 {
                                            case 42:
                                                v38 := GenS5{F0: (v26 % ((v26) | 1)), F1: -9223372036854775808, F2: (fi28 & (fi28 + fi28))}
                                                _ = v38
                                            case 938:
                                                var v39 uint64 = 18446744073709551615
                                                _ = v39
                                            case 2147483647:
                                                var v40 uint = 53
                                                _ = v40
                                            }
                                            _ = v37
                                        } else {
                                            if v24 {
                                                if v24 {
                                                    var v41 int32 = 0
                                                    v27 = (v27 | v27)
                                                    _ = len(v1)
                                                    _ = v41
                                                }
                                            } else {
                                                var v42 int64 = -150
                                                switch genFn1() {
                                                case 2147483647:
                                                    if len(v5) > 0 {
                                                        v5[0] = 65535
                                                    }
                                                case -224:
                                                    v43 := GenS4{F0: 42, F1: 4294967295}
                                                    _ = v43
                                                case 1:
                                                    var v44 string = "\n\t"
                                                    _ = v44
                                                default:
                                                    v27 = (^(v27))
                                                }
                                                if ((v42 * 441) < v42) {
                                                    var v45 uint32 = 65535
                                                    var v46 uint = uint(v25)
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    v47 := GenS3{F0: -464}
                                                    _ = v47
                                                }
                                                _ = v42
                                            }
                                            fmt.Println(v2, fi28, fi28)
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
                _ = fi28
            }
        } else {
            var v48 int8 = v27
            v49 := GenS5{F0: (v26 >> ((2) & 63)), F1: 42, F2: 127}
            for fi50 := 0; fi50 < 6; fi50++ {
                for fi51 := 0; fi51 < 9; fi51++ {
                    fmt.Println(v4, v26)
                    fmt.Println(v25)
                    var v52 float32 = v4
                    _ = fi51
                    _ = v52
                }
                var v53 int8 = (v27 - (v27 ^ (v48 << ((367) & 63))))
                if false {
                    for fi54 := 0; fi54 < 1; fi54++ {
                        v25 += v25
                        var v55 uint64 = 18446744073709551615
                        _ = fi54
                        _ = v55
                    }
                } else {
                    delete(v1, 4294967295)
                    _ = len(v5)
                }
                _ = fi50
                _ = v53
            }
            _ = v48
            _ = v49
        }
        fmt.Println(v27)
    } else {
        if (901 > 82) {
            fmt.Println(v3, v2)
        }
        var v56 uint = 642
        _ = v56
    }
    if v24 {
        v57 := GenS3{F0: 1000}
        v27 -= v27
        _ = len(v5)
        _ = v57
    }
    fmt.Println(v25, v26)
    var v58 int64 = 1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (42 == 18446744073709551615) {
        var v59 int = -2147483648
        fmt.Println(v25, v24, v4)
        for fi60 := 0; fi60 < 3; fi60++ {
            for fi61 := 0; fi61 < 13; fi61++ {
                fmt.Println(v24)
                fmt.Println(v25, v58)
                _ = fi61
            }
            _ = fi60
        }
        _ = v59
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
}
