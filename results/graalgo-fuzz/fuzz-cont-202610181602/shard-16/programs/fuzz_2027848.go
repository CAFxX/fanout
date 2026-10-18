package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 bool; F2 int }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 uint16 }
func genFn0() int64 {
    var v4 bool = (42 != 1)
    for fi5 := 0; fi5 < 11; fi5++ {
        fmt.Println(fi5, fi5)
        var v6 int32 = 42
        _ = fi5
        _ = v6
    }
    var v7 int = 148
    v7 -= v7
    _ = v4
    _ = v7
    return 9223372036854775807
}
func genFn1() uint64 {
    v8 := GenS1{F0: -128}
    var v9 uint16 = 865
    v10 := []int{9223372036854775807, 127}
    if ("hello" < "a") {
        v11 := GenS1{F0: 42}
        for fi12 := 0; fi12 < 6; fi12++ {
            fmt.Println(fi12)
            var v13 int32 = 2147483647
            _ = fi12
            _ = v13
        }
        fmt.Println(v9)
        _ = v11
    } else {
        v9 += 42
        var v14 int = 2147483647
        _ = v14
    }
    _ = v10
    _ = v8
    _ = v9
    return 586
}

func main() {
    var v1 uint16 = 1
    var v2 bool = false
    v3 := make(map[uint32]uint32)
    v3[459] = 65535
    if true {
        v15 := GenS1{F0: 42}
        for fi16 := 0; fi16 < 2; fi16++ {
            v17 := GenS2{F0: 42}
            _ = fi16
            _ = v17
        }
        var v18 int16 = 42
        _ = v15
        _ = v18
    }
    v19 := v3[0]
    fmt.Println(v1)
    var v20 uint64 = 1
    var v21 uint64 = genFn1()
    if true {
        v20 -= genFn1()
        switch "foo bar" {
        case "\n\t":
            for fi22 := 0; fi22 < 4; fi22++ {
                for fi23 := 0; fi23 < 12; fi23++ {
                    var v24 uint64 = genFn1()
                    _ = fi23
                    _ = v24
                }
                var v25 int32 = -1
                switch -90 {
                case 32767:
                    var v26 int64 = 0
                    _ = v26
                case 4294967295:
                    v27 := GenS3{F0: "", F1: true, F2: (fi22 & 9223372036854775807)}
                    _ = v27
                case -788:
                    for fi28 := 0; fi28 < 13; fi28++ {
                        if (v25 > -1) {
                            var v29 int8 = 127
                            _ = v29
                        } else {
                            if v2 {
                                _ = len(v3)
                                v3[(v19 << ((0) & 63))] = v19
                                var v30 string = "hello"
                                _ = v30
                            }
                            var v31 uint16 = (v1 / ((1) | 1))
                            _ = v31
                        }
                        switch fi28 {
                        case -746:
                            v1 -= (v1 & 32767)
                        default:
                            switch int(v1) {
                            case 934:
                                v32 := GenS4{F0: ("\n\t" + ("a" + "")), F1: (v25 / ((v25) | 1))}
                                _ = v32
                            case 2147483647:
                                if (!v2) {
                                    var v33 int8 = -2
                                    var v34 float64 = 0.1
                                    v25 += v25
                                    _ = len(v3)
                                    _ = v33
                                    _ = v34
                                } else {
                                    if v2 {
                                        if true {
                                            v21 -= v21
                                            v35 := []int{128, (fi22 ^ 2147483647)}
                                            _ = v35
                                        } else {
                                            v36 := GenS5{F0: 469}
                                            _ = v36
                                        }
                                    }
                                    fmt.Println(v20)
                                    v37 := GenS4{F0: "\n\t", F1: v25}
                                    _ = v37
                                }
                            }
                        }
                        _ = fi28
                    }
                }
                _ = fi22
                _ = v25
            }
        default:
            fmt.Println(v1)
        }
        var v38 int32 = 612
        _ = v38
    }
    v1 -= 32767
    var v39 uint = 72
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v39
}
