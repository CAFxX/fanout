package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int8 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 uint16 }
func genFn0() uint16 {
    if false {
        var v6 uint64 = 4294967295
        _ = v6
    } else {
        v7 := GenS1{F0: -128, F1: 1, F2: -128}
        if false {
            if (1 < 660) {
                switch ("foo bar" + "a") {
                case "foo bar":
                    if true {
                        var v8 float64 = 1.5
                        _ = v8
                    } else {
                        if false {
                            v9 := []int{458, -1}
                            for fi10 := 0; fi10 < 9; fi10++ {
                                if true {
                                    for fi11 := 0; fi11 < 9; fi11++ {
                                        switch 32767 {
                                        case 1:
                                            var v12 int64 = 9223372036854775807
                                            _ = v12
                                        case -639:
                                            var v13 uint64 = 876
                                            _ = v13
                                        case 128:
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                        }
                                        _ = fi11
                                    }
                                    var v14 uint = 1
                                    v14 -= 1
                                    _ = v14
                                } else {
                                    var v15 uint16 = 1
                                    var v16 string = ("" + "\n\t")
                                    _ = v15
                                    _ = v16
                                }
                                for fi17 := 0; fi17 < 6; fi17++ {
                                    _ = len(v9)
                                    var v18 uint8 = 244
                                    v18 = (v18 - (v18 ^ 127))
                                    _ = fi17
                                    _ = v18
                                }
                                for fi19 := 0; fi19 < 1; fi19++ {
                                    fmt.Println(fi10, fi10)
                                    var v20 uint64 = 1
                                    v20 -= v20
                                    _ = fi19
                                    _ = v20
                                }
                                _ = fi10
                            }
                            if true {
                                var v21 int16 = -1
                                var v22 int16 = 1000
                                var v23 float64 = 1e10
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            _ = v9
                        } else {
                            v24 := []int{32767, -9223372036854775808, 523}
                            v25 := GenS1{F0: -128, F1: 32768, F2: 100}
                            v26 := GenS1{F0: -52, F1: 990, F2: 42}
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    }
                }
                v27 := GenS1{F0: -1, F1: 4294967295, F2: -66}
                v28 := []int{-1, 374}
                _ = v27
                _ = v28
            }
        }
        var v29 bool = (729 > 0)
        _ = v29
        _ = v7
    }
    v30 := []int{}
    _ = v30
    return 42
}

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    v3 := []int{128}
    var v4 int = 42
    var v5 bool = (((!(float32(0.0) >= float32(float64(0.1)))) || (4294967295 != 2)) || ("hello" < ("foo bar" + "\n\t")))
    fmt.Println(v5, v4, v4)
    v5 = v5
    var v31 uint32 = 2147483648
    if (1.5 <= 3.14159) {
        if v1 {
            for fi32 := 0; fi32 < 3; fi32++ {
                for fi33 := 0; fi33 < 15; fi33++ {
                    fmt.Println(v1)
                    v5 = v5
                    if v1 {
                        delete(v2, (^(fi32)))
                    } else {
                        if v5 {
                            fmt.Println(v4)
                        } else {
                            v34 := GenS1{F0: -1, F1: (fi33 >> ((519) & 63)), F2: 42}
                            var v35 uint = 143
                            _ = v34
                            _ = v35
                        }
                    }
                    _ = fi33
                }
                _ = fi32
            }
        } else {
            v31 += v31
            var v36 int = 2147483647
            v37 := GenS1{F0: 42, F1: 32768, F2: 1}
            _ = v36
            _ = v37
        }
        v38 := GenS2{F0: 128, F1: ("hello" + "a")}
        fmt.Println(v4, v1, v1)
        _ = v38
    }
    var v39 int8 = -101
    v5 = v1
    v1 = v5
    v31 = v31
    for fi40 := 0; fi40 < 19; fi40++ {
        v31 = (v31 * v31)
        _ = fi40
    }
    for fi41 := 0; fi41 < 10; fi41++ {
        v42 := GenS1{F0: v39, F1: (fi41 / ((0) | 1)), F2: -1}
        var v43 uint = 1
        _ = fi41
        _ = v42
        _ = v43
    }
    v44 := GenS3{F0: (v31 - v31)}
    for fi45 := 0; fi45 < 12; fi45++ {
        var v46 int32 = 822
        _ = fi45
        _ = v46
    }
    if v1 {
        v47 := GenS2{F0: 0, F1: "foo bar"}
        v48 := GenS4{F0: 191, F1: (^(v4))}
        v49 := GenS5{F0: 1}
        _ = v47
        _ = v48
        _ = v49
    } else {
        var v50 float64 = 1e10
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v44
    _ = v5
}
