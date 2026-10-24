package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 int32; F1 uint32; F2 bool }
type GenS7 struct { F0 bool; F1 int32; F2 int32 }
type GenS8 struct { F0 bool; F1 int16 }
func genFn0(p0 bool) uint64 {
    fmt.Println(p0)
    var v7 int16 = -1
    v8 := []int{32768, 567, 377}
    v8 = append(v8, 32767)
    for fi9 := 0; fi9 < 20; fi9++ {
        if (130 > 128) {
            for fi10 := 0; fi10 < 4; fi10++ {
                switch "" {
                case "foo bar":
                    var v11 float64 = 1.5
                    _ = v11
                case "a":
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                case "\n\t":
                    v12 := GenS1{F0: -9223372036854775808, F1: -895}
                    _ = v12
                }
                v13 := GenS2{F0: (18446744073709551615 < 905), F1: (!("" > ""))}
                if (42 != -2147483648) {
                    fmt.Println(p0, fi10)
                    v14 := GenS1{F0: -9223372036854775808, F1: 0}
                    _ = v14
                }
                _ = fi10
                _ = v13
            }
            v15 := GenS3{F0: true, F1: -102}
            _ = v15
        } else {
            p0 = p0
            var v16 uint = 0
            _ = len(v8)
            _ = v16
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 float32 = 0.0
            var v19 uint32 = 1
            if (0 < 871) {
                if (("a" + "") >= "foo bar") {
                    v19 = v19
                    var v20 bool = false
                    v21 := GenS1{F0: -1, F1: 869}
                    _ = v20
                    _ = v21
                }
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        for fi22 := 0; fi22 < 9; fi22++ {
            v23 := GenS4{F0: (!(p0 && (0 != 326))), F1: 128}
            var v24 string = "a"
            v24 = v24
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi9
    }
    _ = p0
    _ = v7
    _ = v8
    return 21
}
func genFn1(p0 uint64, p1 uint32, p2 int8) uint16 {
    for fi25 := 0; fi25 < 12; fi25++ {
        v26 := []int{fi25, fi25, (fi25 >> ((65535) & 63))}
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        v28 := GenS5{F0: p1, F1: 2}
        var v29 uint64 = (p0 % ((0) | 1))
        if (((32767 != 795) && (1000 != -1)) || false) {
            for fi30 := 0; fi30 < 13; fi30++ {
                var v31 int32 = 0
                if (!("hello" == "a")) {
                    if (true && false) {
                        var v32 int8 = p2
                        var v33 int64 = -41
                        v34 := GenS6{F0: -2147483648, F1: uint32(v29), F2: true}
                        _ = v32
                        _ = v33
                        _ = v34
                    } else {
                        var v35 bool = false
                        fmt.Println(p0, v35, fi27)
                        var v36 int64 = 0
                        _ = v35
                        _ = v36
                    }
                } else {
                    var v37 int16 = -32768
                    _ = v37
                }
                if (1 == 18446744073709551615) {
                    p1 -= 0
                }
                if false {
                    v38 := GenS7{F0: (!true), F1: (^(v31)), F2: v31}
                    if (128 > 2) {
                        var v39 uint = 589
                        _ = v39
                    } else {
                        v31 -= (v31 + v31)
                        v29 = v29
                        var v40 int = (-(fi30))
                        _ = v40
                    }
                    switch (^(fi27)) {
                    case 1:
                        v41 := GenS8{F0: (763 > 709), F1: 32767}
                        _ = v41
                    case -1:
                        v31 -= -989
                        switch "\n\t" {
                        case "hello":
                            var v42 int32 = v31
                            v42 = (v42 >> ((1) & 63))
                            _ = v42
                        case "a":
                            var v43 uint32 = p1
                            _ = v43
                        case "foo bar":
                            var v44 uint16 = 255
                            _ = v44
                        }
                        p1 -= p1
                    case -2147483648:
                        var v45 uint64 = 18446744073709551615
                        _ = v45
                    }
                    _ = v38
                }
                switch 128 {
                case 141:
                    var v46 uint8 = 255
                    _ = v46
                default:
                    p0 -= (v29 | (p0 - genFn0(true)))
                }
                _ = fi30
                _ = v31
            }
        } else {
            v29 = genFn0(false)
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    if false {
        if false {
            var v47 string = "a"
            var v48 uint = 0
            _ = v47
            _ = v48
        } else {
            p0 += genFn0(false)
            if false {
                v49 := GenS5{F0: p1, F1: 177}
                p2 = (p2 * (p2 >> ((18446744073709551615) & 63)))
                fmt.Println(p1, p0, p1)
                _ = v49
            }
        }
        switch "\n\t" {
        case "foo bar":
            var v50 int16 = 0
            _ = v50
        case "hello":
            var v51 uint8 = 0
            _ = v51
        default:
            var v52 int = -1
            _ = v52
        }
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]bool)
    var v3 int64 = 903
    var v4 int8 = 0
    var v5 int16 = 42
    var v6 float32 = float32((1.5 + 0.0))
    var v53 string = "\n\t"
    for fi54 := 0; fi54 < 10; fi54++ {
        var v55 float64 = ((0.0 * 3.14159) - 0.1)
        if true {
            fmt.Println(v55, v6)
        }
        v56 := GenS7{F0: false, F1: 42, F2: -520}
        _ = fi54
        _ = v55
        _ = v56
    }
    var v57 uint64 = 42
    delete(v2, v53)
    switch v53 {
    case "foo bar":
        fmt.Println(v53, v6)
    case "a":
        for fi58 := 0; fi58 < 12; fi58++ {
            var v59 int64 = (v3 - v3)
            _ = fi58
            _ = v59
        }
    }
    v5 += 841
    var v60 float32 = (v6 / (float32((float64(v6) * 1.5)) / float32(1.5)))
    fmt.Println(v57, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v57
    _ = v6
    _ = v60
}
