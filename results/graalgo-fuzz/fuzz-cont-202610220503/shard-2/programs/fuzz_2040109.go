package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 int8; F1 string; F2 int16 }
func genFn0(p0 uint64) int8 {
    if (127 >= 42) {
        for fi7 := 0; fi7 < 2; fi7++ {
            fmt.Println(p0)
            for fi8 := 0; fi8 < 5; fi8++ {
                if true {
                    p0 = p0
                }
                _ = fi8
            }
            for fi9 := 0; fi9 < 14; fi9++ {
                if false {
                    if true {
                        var v10 string = "foo bar"
                        _ = v10
                    }
                    switch "" {
                    case "":
                        p0 += 1
                    case "a":
                        p0 -= (^(p0))
                    case "hello":
                        v11 := GenS1{F0: 128, F1: 0}
                        _ = v11
                    default:
                        p0 -= (p0 + p0)
                    }
                }
                var v12 int = fi9
                _ = fi9
                _ = v12
            }
            _ = fi7
        }
    }
    v13 := GenS2{F0: 727}
    p0 = 290
    switch 42 {
    case 722:
        v14 := []int{-9223372036854775808}
        _ = v14
    }
    _ = p0
    _ = v13
    return -128
}
func genFn1(p0 bool) int16 {
    var v15 int = 4294967295
    var v16 uint64 = 4294967295
    var v17 float64 = (0.0 + 0.1)
    switch "\n\t" {
    case "":
        switch v15 {
        case 171:
            v18 := GenS1{F0: 0, F1: 0}
            _ = v18
        case 2147483647:
            switch (("foo bar" + "a") + "") {
            case "hello":
                for fi19 := 0; fi19 < 2; fi19++ {
                    var v20 int32 = 42
                    if ((v16 <= v16) || p0) {
                        v21 := GenS3{F0: 9223372036854775807}
                        v22 := GenS1{F0: 1, F1: 156}
                        var v23 uint32 = 2
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    _ = fi19
                    _ = v20
                }
            case "foo bar":
                if p0 {
                    v24 := []int{}
                    v25 := GenS4{F0: uint(v16), F1: 32768}
                    _ = v24
                    _ = v25
                }
            case "a":
                if p0 {
                    var v26 uint16 = 720
                    _ = v26
                }
            default:
                v27 := []int{-99, 356}
                _ = v27
            }
        case -1:
            if (v15 == (v15 / ((v15) | 1))) {
                if false {
                    v28 := []int{(-(v15))}
                    for fi29 := 0; fi29 < 20; fi29++ {
                        var v30 int64 = 9223372036854775807
                        _ = fi29
                        _ = v30
                    }
                    _ = v28
                } else {
                    v31 := GenS5{F0: -572, F1: v15}
                    var v32 uint8 = 2
                    v17 = -2.25
                    _ = v31
                    _ = v32
                }
            }
        }
    case "foo bar":
        v15 += -1
    }
    var v33 int16 = -678
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    _ = v33
    return v33
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{65535, 42, 2147483647}
    var v3 int8 = 100
    v4 := make(map[string]int)
    var v5 bool = false
    v6 := make(map[uint32]bool)
    v5 = v5
    var v34 uint16 = 42
    v35 := GenS2{F0: 0}
    var v36 int32 = -2147483648
    _ = len(v1)
    switch ("a" + ((("\n\t" + "") + ("\n\t" + "hello")) + (("foo bar" + "a") + "hello"))) {
    case "":
        for fi37 := 0; fi37 < 17; fi37++ {
            v38 := v6[78]
            var v39 uint16 = (v34 + (v34 >> ((73) & 63)))
            _ = fi37
            _ = v38
            _ = v39
        }
    }
    v5 = v5
    if (v5 || false) {
        if v5 {
            if (1 <= 128) {
                switch "\n\t" {
                case "foo bar":
                    for fi40 := 0; fi40 < 3; fi40++ {
                        for fi41 := 0; fi41 < 17; fi41++ {
                            var v42 float64 = 1.5
                            _ = fi41
                            _ = v42
                        }
                        _ = fi40
                    }
                case "\n\t":
                    if (2147483647 == 2147483647) {
                        v3 = (v3 | (^(v3)))
                        for fi43 := 0; fi43 < 18; fi43++ {
                            if false {
                                v44 := v4[("\n\t" + "a")]
                                var v45 int32 = 0
                                fmt.Println(v5, v3)
                                _ = v44
                                _ = v45
                            } else {
                                var v46 uint64 = 4294967295
                                _ = v46
                            }
                            _ = fi43
                        }
                        v47 := GenS6{F0: 42, F1: 2147483647}
                        _ = v47
                    } else {
                        var v48 uint64 = 2
                        var v49 int64 = 0
                        _ = v48
                        _ = v49
                    }
                case "hello":
                    v5 = v5
                default:
                    v50 := GenS7{F0: -531, F1: genFn0(1)}
                    _ = v50
                }
                v5 = (4294967295 <= 2)
            } else {
                v51 := GenS5{F0: 0, F1: 1}
                _ = v51
            }
            switch 32768 {
            case -658:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 0:
                v52 := GenS8{F0: -107, F1: "foo bar", F2: 0}
                _ = v52
            }
        }
        if v5 {
            switch "a" {
            case "":
                v34 += 42
            case "a":
                if (!(v5 || false)) {
                    var v53 int64 = 9223372036854775807
                    _ = v53
                }
            }
            var v54 uint16 = (v34 + (v34 & (v34 + (v34 + 255))))
            _ = v54
        }
    } else {
        var v55 uint64 = 18446744073709551615
        for fi56 := 0; fi56 < 6; fi56++ {
            switch ("\n\t" + ("" + (("foo bar" + ("hello" + "foo bar")) + "foo bar"))) {
            case "foo bar":
                switch fi56 {
                case 9223372036854775807:
                    v55 = (v55 & 1)
                }
            case "hello":
                fmt.Println(v5, v34)
            case "":
                var v57 uint32 = 823
                _ = v57
            }
            fmt.Println(fi56)
            if true {
                v58 := GenS2{F0: (-(v36))}
                var v59 float64 = 0.0
                for fi60 := 0; fi60 < 5; fi60++ {
                    var v61 uint32 = 2147483648
                    v59 = (float64(-2.25) + v59)
                    _ = fi60
                    _ = v61
                }
                _ = v58
                _ = v59
            }
            _ = fi56
        }
        fmt.Println(v34)
        _ = v55
    }
    v36 = (v36 % (((^(v36))) | 1))
    fmt.Println(v36)
    var v62 bool = ((v36 * v36) <= v36)
    _ = len(v2)
    fmt.Println(v3, v62)
    var v63 float32 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v63
}
