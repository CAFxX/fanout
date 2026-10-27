package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 bool; F1 uint8; F2 int16 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 int8 }
func genFn0() int {
    if false {
        if true {
            var v7 int16 = 960
            var v8 int32 = 765
            _ = v7
            _ = v8
        }
    } else {
        if true {
            var v9 int8 = 42
            v10 := GenS1{F0: 1, F1: 740, F2: 18446744073709551615}
            _ = v10
            _ = v9
        } else {
            var v11 uint32 = 2147483648
            v11 = 929
            _ = v11
        }
    }
    v12 := GenS2{F0: 314, F1: -2147483648}
    _ = v12
    return -867
}
func genFn1(p0 uint64, p1 string) uint32 {
    v13 := GenS2{F0: 255, F1: 953}
    v14 := GenS1{F0: -343, F1: 509, F2: p0}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 630
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    var v3 int16 = -968
    v4 := make(map[uint32]string)
    var v5 float64 = (0.1 + (0.1 - (float64(1e10) + (0.0 / 1.0))))
    v6 := make(map[string]string)
    v3 = v3
    v15 := GenS1{F0: (v3 & (^(v3))), F1: 42, F2: 1}
    var v16 int16 = -449
    for fi17 := 0; fi17 < 10; fi17++ {
        var v18 bool = (42 != 122)
        var v19 uint16 = 42
        v5 = (v5 + 1.5)
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(v16, v16, v16)
    if ("hello" > "foo bar") {
        v16 -= (v3 | (v3 & v16))
    } else {
        switch "a" {
        case "hello":
            if (255 >= 0) {
                if (4294967295 < 694) {
                    var v20 uint64 = 4294967295
                    fmt.Println(v5, v5)
                    _ = v20
                }
                var v21 uint32 = 801
                switch "foo bar" {
                case "\n\t":
                    if true {
                        for fi22 := 0; fi22 < 2; fi22++ {
                            v3 = 42
                            v23 := []int{genFn0(), (fi22 + fi22)}
                            _ = fi22
                            _ = v23
                        }
                        v21 -= v21
                        v24 := GenS1{F0: (v16 / ((0) | 1)), F1: 1, F2: 4294967295}
                        _ = v24
                    } else {
                        v25 := GenS3{F0: (genFn0() >= 1), F1: 2, F2: 32767}
                        _ = v25
                    }
                default:
                    v26 := GenS4{F0: -23, F1: 32767}
                    _ = v26
                }
                _ = v21
            } else {
                v16 += v16
            }
        }
        switch 94 {
        case 4294967295:
            v16 -= (v3 ^ (-(v16)))
        case -619:
            switch "\n\t" {
            case "\n\t":
                if false {
                    if true {
                        var v27 int64 = -1
                        fmt.Println(v16, v16)
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 10; fi28++ {
                        v29 := GenS2{F0: 2, F1: -2147483648}
                        _ = fi28
                        _ = v29
                    }
                }
            case "":
                v30 := []int{42, genFn0()}
                _ = v30
            default:
                if (true && (!true)) {
                    for fi31 := 0; fi31 < 3; fi31++ {
                        var v32 uint16 = 1
                        var v33 int8 = -89
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    v34 := GenS4{F0: -25, F1: (v16 ^ v16)}
                    _ = v34
                }
            }
        case -735:
            fmt.Println(v3, v16, v5)
        default:
            v16 += (v16 ^ v3)
        }
        v16 += v3
    }
    switch ("\n\t" + ("\n\t" + "a")) {
    case "\n\t":
        var v35 uint8 = 138
        _ = v35
    case "hello":
        v36 := v2[527]
        _ = v36
    }
    v5 += v5
    v37 := v1[""]
    fmt.Println(v5, v16, v37)
    v38 := GenS5{F0: -128}
    v16 += -1
    switch ("hello" + "\n\t") {
    case "foo bar":
        for fi39 := 0; fi39 < 2; fi39++ {
            var v40 uint16 = 894
            if true {
                fmt.Println(v3, fi39)
                for fi41 := 0; fi41 < 4; fi41++ {
                    v3 += (v16 & v3)
                    var v42 float64 = v5
                    _ = fi41
                    _ = v42
                }
            } else {
                switch "" {
                case "":
                    var v43 int64 = 36
                    _ = v43
                }
                fmt.Println(v5, fi39, v5)
                _ = len(v2)
            }
            _ = fi39
            _ = v40
        }
    default:
        var v44 uint8 = 127
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
