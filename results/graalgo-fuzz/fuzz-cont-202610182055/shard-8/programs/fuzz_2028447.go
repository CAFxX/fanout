package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int; F1 uint; F2 uint8 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 int64; F1 uint; F2 int8 }
func genFn0(p0 bool) uint16 {
    p0 = (!(-1 > 32768))
    v5 := GenS1{F0: 1, F1: 0}
    if ((568 < 42) && (-1 != 754)) {
        if false {
            for fi6 := 0; fi6 < 7; fi6++ {
                for fi7 := 0; fi7 < 11; fi7++ {
                    var v8 uint8 = 128
                    var v9 uint = 654
                    if p0 {
                        if (!p0) {
                            var v10 int = 42
                            v11 := []int{(fi7 % (((fi7 / ((127) | 1))) | 1)), fi7}
                            _ = v10
                            _ = v11
                        } else {
                            if (0.0 != ((1e10 - ((1.5 + (3.14159 + 1.5)) / (1.5 + (3.14159 * (-2.25 + 1e10))))) * (1.5 * float64(1e10)))) {
                                switch ((("\n\t" + "foo bar") + "foo bar") + "hello") {
                                case "\n\t":
                                    var v12 int8 = 127
                                    _ = v12
                                default:
                                    v13 := GenS2{F0: fi6, F1: v9, F2: v8}
                                    _ = v13
                                }
                                v14 := GenS1{F0: v9, F1: 1}
                                v15 := []int{}
                                v9 -= (^(v9))
                                _ = v14
                                _ = v15
                            }
                            fmt.Println(fi6, v8)
                        }
                        fmt.Println(v8)
                    } else {
                        var v16 int = (-(fi7))
                        _ = v16
                    }
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                if p0 {
                    switch fi6 {
                    case 942:
                        v17 := []int{fi6, fi6}
                        _ = v17
                    default:
                        fmt.Println(p0)
                    }
                } else {
                    p0 = (1.5 == -2.25)
                    for fi18 := 0; fi18 < 17; fi18++ {
                        p0 = true
                        var v19 uint8 = 42
                        _ = fi18
                        _ = v19
                    }
                    v20 := GenS1{F0: 18446744073709551615, F1: 594}
                    _ = v20
                }
                p0 = p0
                _ = fi6
            }
            p0 = p0
        } else {
            var v21 bool = (875 != 278)
            v22 := []int{9223372036854775807}
            _ = v21
            _ = v22
        }
        var v23 uint32 = 796
        _ = v23
    }
    if false {
        for fi24 := 0; fi24 < 3; fi24++ {
            for fi25 := 0; fi25 < 14; fi25++ {
                switch (fi24 | -1) {
                case 32767:
                    var v26 uint16 = 0
                    _ = v26
                default:
                    p0 = true
                }
                if p0 {
                    v27 := []int{fi25, (fi25 % ((fi24) | 1)), (-(fi24)), -458}
                    var v28 string = "hello"
                    _ = v27
                    _ = v28
                } else {
                    v29 := []int{-9223372036854775808}
                    v30 := GenS1{F0: 2, F1: 4294967295}
                    _ = v29
                    _ = v30
                }
                var v31 uint64 = 4294967295
                _ = fi25
                _ = v31
            }
            var v32 int32 = 1
            _ = fi24
            _ = v32
        }
        v33 := GenS1{F0: 65535, F1: 985}
        _ = v33
    }
    _ = p0
    _ = v5
    return 42
}
func genFn1(p0 int32, p1 uint32) int64 {
    p0 += p0
    p1 += (p1 >> ((0) & 63))
    p0 = p0
    _ = p0
    _ = p1
    return -622
}

func main() {
    var v1 string = "a"
    var v2 int = 32767
    v3 := make(map[int]string)
    var v4 uint8 = 0
    v4 = (v4 / ((v4) | 1))
    v34 := GenS2{F0: (v2 % ((1) | 1)), F1: 381, F2: 2}
    switch ("" + v1) {
    case "foo bar":
        switch v2 {
        case 32767:
            if true {
                for fi35 := 0; fi35 < 8; fi35++ {
                    fmt.Println(v1)
                    v2 = v2
                    _ = fi35
                }
                if (false && false) {
                    fmt.Println(v2, v1, v1)
                    var v36 uint16 = 289
                    var v37 float64 = 0.1
                    _ = v36
                    _ = v37
                }
            } else {
                v38 := GenS3{F0: ("a" + "hello"), F1: 657}
                fmt.Println(v2)
                _ = v38
            }
        }
    case "hello":
        fmt.Println(v1, v4, v1)
    }
    var v39 uint16 = 1
    v4 -= v4
    for fi40 := 0; fi40 < 17; fi40++ {
        if ((^(v2)) == (fi40 * (v2 - (v2 * fi40)))) {
            switch (fi40 + 1) {
            case 2147483647:
                v39 = genFn0(false)
            case 371:
                v1 = v1
            default:
                fmt.Println(v2, v4, v1)
            }
            _ = len(v3)
        } else {
            v41 := GenS4{F0: v4, F1: ((v1 + "foo bar") + (v1 + v1))}
            var v42 uint16 = (v39 >> ((65535) & 63))
            var v43 uint16 = v39
            _ = v41
            _ = v42
            _ = v43
        }
        v44 := GenS5{F0: 0, F1: 0, F2: -67}
        _ = fi40
        _ = v44
    }
    switch ("foo bar" + (("foo bar" + v1) + v1)) {
    case "a":
        var v45 bool = (!true)
        _ = v45
    default:
        switch v1 {
        case "foo bar":
            v46 := GenS4{F0: 0, F1: v1}
            _ = v46
        case "\n\t":
            fmt.Println(v4, v1, v2)
        case "hello":
            _ = len(v3)
        default:
            fmt.Println(v39)
        }
    }
    v47 := GenS1{F0: 65535, F1: 42}
    var v48 uint32 = uint32(v4)
    var v49 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v39
    _ = v4
    _ = v47
    _ = v48
    _ = v49
}
