package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int16 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int64; F1 int8; F2 int8 }
func genFn0(p0 int16) uint16 {
    for fi4 := 0; fi4 < 4; fi4++ {
        if (0 != 42) {
            switch ("foo bar" + (("hello" + "a") + "foo bar")) {
            case "hello":
                for fi5 := 0; fi5 < 16; fi5++ {
                    var v6 uint64 = 42
                    v7 := GenS1{F0: -423, F1: 127, F2: 32767}
                    var v8 int64 = 1
                    _ = fi5
                    _ = v6
                    _ = v7
                    _ = v8
                }
            }
            v9 := GenS1{F0: -1, F1: 2, F2: p0}
            if false {
                if false {
                    for fi10 := 0; fi10 < 10; fi10++ {
                        v11 := GenS1{F0: -2147483648, F1: 255, F2: -1}
                        var v12 uint64 = 42
                        var v13 int64 = 1
                        if true {
                            fmt.Println(fi10)
                            switch (fi10 % (((^(fi4))) | 1)) {
                            case -463:
                                switch "foo bar" {
                                case "":
                                    if false {
                                        v14 := GenS1{F0: 2147483647, F1: 96, F2: p0}
                                        p0 = (p0 | (-(p0)))
                                        if (v13 <= (-(v13))) {
                                            v12 = 1
                                            switch "a" {
                                            case "foo bar":
                                                switch fi10 {
                                                case 32767:
                                                    p0 -= int16(v13)
                                                case 1:
                                                    fmt.Println(v12, v12)
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v15 := GenS2{F0: (fi4 / ((fi10) | 1)), F1: v13}
                                                        _ = v15
                                                    case "hello":
                                                        v16 := GenS3{F0: v13, F1: 66, F2: 42}
                                                        fmt.Println(p0, v12, v12)
                                                        _ = v16
                                                    default:
                                                        if true {
                                                            var v17 string = (("a" + (("a" + "a") + ("foo bar" + "a"))) + "foo bar")
                                                            var v18 int32 = int32(v13)
                                                            fmt.Println(v12, v18)
                                                            _ = v17
                                                            _ = v18
                                                        }
                                                    }
                                                    v19 := []int{}
                                                    p0 = int16(v12)
                                                    _ = v19
                                                case 128:
                                                    p0 -= p0
                                                }
                                            case "hello":
                                                v12 = v12
                                            case "\n\t":
                                                v13 = v13
                                            }
                                            v20 := GenS2{F0: fi4, F1: (v13 % ((v13) | 1))}
                                            _ = v20
                                        }
                                        v13 -= -9223372036854775808
                                        fmt.Println(fi4, v12, fi4)
                                        _ = v14
                                    }
                                case "a":
                                    var v21 int64 = v13
                                    _ = v21
                                default:
                                    p0 = (p0 & p0)
                                }
                            case 1:
                                v13 += v13
                            }
                        }
                        _ = fi10
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    p0 = p0
                    v22 := []int{-9223372036854775808, (fi4 ^ fi4), (fi4 + 127), (fi4 & (fi4 * 2147483647))}
                    _ = v22
                }
            }
            _ = v9
        } else {
            p0 -= (-(p0))
        }
        var v23 int8 = 100
        switch "a" {
        case "foo bar":
            var v24 uint = 0
            _ = v24
        }
        _ = fi4
        _ = v23
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        var v26 int64 = 42
        _ = fi25
        _ = v26
    }
    v27 := []int{32767, -1, 32767, 32768}
    _ = p0
    _ = v27
    return 2
}
func genFn1() bool {
    v28 := GenS3{F0: 42, F1: 0, F2: 5}
    var v29 int64 = -1
    _ = v28
    _ = v29
    return (true || ((!((false || (701 >= 0)) || ((39 < -117) || ((((-2.25 + 1.5) - 3.14159) * (3.14159 * 0.0)) != -2.25)))) && (51 == v29)))
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 936
    var v3 int = 985
    fmt.Println(v2)
    v2 += 772
    for fi30 := 0; fi30 < 9; fi30++ {
        var v31 uint32 = v2
        var v32 uint32 = (v2 * (^(v31)))
        _ = fi30
        _ = v31
        _ = v32
    }
    v2 = (v2 + (v2 / ((1) | 1)))
    var v33 int = 32768
    v2 += 845
    fmt.Println(v33)
    v34 := GenS3{F0: 397, F1: 42, F2: 124}
    v35 := []int{42, v3, v33, 1}
    if len(v35) >= 2 {
        v35 = v35[1:len(v35)]
    }
    if (("foo bar" + (("foo bar" + "hello") + "a")) <= (("" + ("" + "")) + "\n\t")) {
        v33 = 2147483647
        v36 := GenS1{F0: 42, F1: 192, F2: -297}
        v3 += (v33 + (-(v3)))
        _ = v36
    } else {
        v37 := GenS2{F0: v33, F1: 0}
        var v38 uint16 = 2
        _ = v37
        _ = v38
    }
    v33 = 2147483647
    var v39 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v39
}
