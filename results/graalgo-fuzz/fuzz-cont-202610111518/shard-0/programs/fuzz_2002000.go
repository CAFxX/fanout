package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 string; F1 uint16 }

func main() {
    var v1 float32 = 1.5
    var v2 int16 = -32768
    var v3 uint8 = 255
    v4 := make(map[int]uint32)
    var v5 uint8 = (v3 & v3)
    v6 := GenS1{F0: true}
    var v7 int64 = 42
    if true {
        for fi8 := 0; fi8 < 3; fi8++ {
            v3 += (v3 & v5)
            var v9 uint64 = 760
            _ = fi8
            _ = v9
        }
    } else {
        var v10 bool = ((v7 < (-(v7))) && false)
        switch 823 {
        case 32767:
            fmt.Println(v2, v10)
        }
        _ = v10
    }
    switch "" {
    case "":
        v5 -= (^(v3))
    case "\n\t":
        for fi11 := 0; fi11 < 4; fi11++ {
            v12 := GenS1{F0: (((v2 | (^(v2))) > 621) && (42 < 429))}
            v13 := GenS1{F0: ((v2 != (v2 * v2)) && false)}
            switch fi11 {
            case -2147483648:
                var v14 float64 = (0.0 * (0.1 * float64(((1.5 - 0.0) / (3.14159 / v1)))))
                _ = v14
            case -749:
                v15 := []int{}
                _ = v15
            }
            _ = fi11
            _ = v12
            _ = v13
        }
    case "foo bar":
        v16 := GenS2{F0: 1, F1: (^(v5))}
        _ = v16
    default:
        v17 := GenS3{F0: v5, F1: 2}
        _ = v17
    }
    fmt.Println(v2, v3)
    fmt.Println(v3, v5)
    if (v1 != v1) {
        v18 := v4[2147483647]
        var v19 uint = 255
        _ = v18
        _ = v19
    } else {
        fmt.Println(v3, v7, v7)
    }
    if (v7 == (^(v7))) {
        var v20 int = -501
        var v21 uint32 = 1
        _ = v20
        _ = v21
    }
    _ = len(v4)
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := []int{9223372036854775807, (fi22 << ((82) & 63)), (fi22 + fi22)}
        _ = fi22
        _ = v23
    }
    var v24 uint8 = (v5 + v5)
    v25 := []int{}
    v1 += (1e10 * float32(0.1))
    var v26 bool = false
    v2 += -1
    switch 127 {
    case 32767:
        for fi27 := 0; fi27 < 17; fi27++ {
            switch "a" {
            case "a":
                v7 = v7
            default:
                v1 = (3.14159 - v1)
            }
            var v28 int = (fi27 << ((65535) & 63))
            for fi29 := 0; fi29 < 20; fi29++ {
                v4[fi27] = 2147483648
                v4[fi29] = 1
                _ = fi29
            }
            _ = fi27
            _ = v28
        }
    case 0:
        var v30 int = -1
        _ = v30
    default:
        v1 = v1
    }
    switch ("a" + (("hello" + "foo bar") + ("foo bar" + "foo bar"))) {
    case "\n\t":
        for fi31 := 0; fi31 < 7; fi31++ {
            if (false || (v26 && v26)) {
                if len(v25) > 0 {
                    v25[0] = 42
                }
                v5 += 1
                v32 := v4[-964]
                _ = v32
            } else {
                if (4294967295 > 1) {
                    var v33 int16 = v2
                    switch "a" {
                    case "a":
                        _ = len(v25)
                    case "\n\t":
                        for fi34 := 0; fi34 < 5; fi34++ {
                            var v35 uint8 = v5
                            var v36 int8 = 100
                            v37 := GenS2{F0: 2, F1: v3}
                            v38 := GenS4{F0: "a", F1: 366}
                            v24 = v5
                            fmt.Println(fi34, fi34)
                            _ = fi34
                            _ = v35
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                    }
                    switch (("" + "hello") + "hello") {
                    case "\n\t":
                        v39 := GenS2{F0: 2147483648, F1: v5}
                        _ = v39
                    case "hello":
                        if ((^(v7)) != int64(v2)) {
                            for fi40 := 0; fi40 < 7; fi40++ {
                                if (-9223372036854775808 < v7) {
                                    var v41 uint = 4294967295
                                    _ = v41
                                } else {
                                    var v42 int = fi31
                                    var v43 uint32 = 2147483648
                                    _ = v42
                                    _ = v43
                                }
                                fmt.Println(v5, v3)
                                fmt.Println(v3)
                                _ = fi40
                            }
                        }
                    case "foo bar":
                        if false {
                            var v44 uint32 = 959
                            _ = v44
                        }
                    }
                    _ = v33
                }
                for fi45 := 0; fi45 < 1; fi45++ {
                    v46 := GenS4{F0: "", F1: 0}
                    if true {
                        v25 = append(v25, fi31)
                    } else {
                        _ = len(v4)
                    }
                    _ = fi45
                    _ = v46
                }
                if true {
                    v47 := v4[fi31]
                    var v48 int16 = 925
                    _ = v47
                    _ = v48
                }
            }
            v49 := GenS1{F0: v26}
            _ = fi31
            _ = v49
        }
    case "foo bar":
        v50 := GenS1{F0: v26}
        _ = v50
    case "":
        v1 = (((v1 * v1) - v1) + float32(-2.25))
    default:
        v5 = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
