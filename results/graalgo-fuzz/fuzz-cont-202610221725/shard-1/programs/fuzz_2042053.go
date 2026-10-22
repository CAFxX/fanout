package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int8 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 string; F1 int16; F2 int8 }
type GenS4 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint32; F1 uint64 }

func main() {
    v1 := []int{1, 308, 2147483647, -9223372036854775808}
    var v2 uint = 581
    v3 := make(map[string]string)
    var v4 bool = true
    switch 915 {
    case 82:
        if (!v4) {
            var v5 uint64 = 0
            v4 = false
            var v6 int16 = 0
            _ = v5
            _ = v6
        }
    default:
        var v7 uint16 = 32767
        _ = v7
    }
    v8 := GenS1{F0: 380, F1: -1, F2: 100}
    v9 := GenS1{F0: 32767, F1: -9223372036854775808, F2: 127}
    v10 := GenS1{F0: -1, F1: 9223372036854775807, F2: 100}
    v4 = (((1.5 * 0.1) - -2.25) > 0.0)
    v4 = v4
    if (("a" < (("foo bar" + ("" + "\n\t")) + "hello")) || (v4 || v4)) {
        var v11 uint32 = 4294967295
        for fi12 := 0; fi12 < 11; fi12++ {
            fmt.Println(v4)
            var v13 uint32 = v11
            _ = fi12
            _ = v13
        }
        _ = v11
    } else {
        v4 = v4
        var v14 string = "hello"
        for fi15 := 0; fi15 < 18; fi15++ {
            v4 = v4
            if ((18446744073709551615 > 1) || (!(-2.25 < 0.0))) {
                _ = len(v1)
                for fi16 := 0; fi16 < 7; fi16++ {
                    if (-1 != 121) {
                        v17 := GenS2{F0: 1, F1: (v2 * v2)}
                        _ = v17
                    }
                    var v18 float64 = 0.1
                    _ = fi16
                    _ = v18
                }
                v19 := GenS1{F0: (fi15 / ((1) | 1)), F1: -1, F2: 71}
                _ = v19
            } else {
                fmt.Println(v14)
            }
            _ = fi15
        }
        _ = v14
    }
    if v4 {
        var v20 uint64 = 887
        if true {
            for fi21 := 0; fi21 < 4; fi21++ {
                var v22 int16 = 32767
                switch fi21 {
                case -147:
                    v2 = 2
                case 775:
                    var v23 int8 = 100
                    _ = v23
                default:
                    for fi24 := 0; fi24 < 10; fi24++ {
                        switch (^(fi21)) {
                        case 42:
                            var v25 uint64 = (v20 & v20)
                            _ = v25
                        default:
                            switch fi24 {
                            case 128:
                                if (125 >= -128) {
                                    var v26 uint32 = 2147483647
                                    var v27 float64 = (3.14159 + 0.1)
                                    _ = v26
                                    _ = v27
                                } else {
                                    var v28 string = ("" + ("" + ""))
                                    if false {
                                        v29 := v3[(v28 + v28)]
                                        _ = v29
                                    } else {
                                        var v30 int32 = 334
                                        _ = v30
                                    }
                                    _ = v28
                                }
                            case 11:
                                v31 := GenS3{F0: "\n\t", F1: int16(v20), F2: 0}
                                _ = v31
                            case -373:
                                v32 := v3["foo bar"]
                                _ = v32
                            }
                        }
                        if false {
                            v2 = (v2 >> (((v2 / ((v2) | 1))) & 63))
                            v33 := GenS3{F0: "", F1: v22, F2: 127}
                            v34 := v3["a"]
                            _ = v33
                            _ = v34
                        } else {
                            v35 := GenS4{F0: false, F1: 141, F2: (^(v20))}
                            var v36 uint32 = 2147483647
                            var v37 uint = (v2 & v2)
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        _ = fi24
                    }
                }
                v38 := GenS4{F0: false, F1: 2, F2: 4294967295}
                _ = fi21
                _ = v22
                _ = v38
            }
            v39 := GenS4{F0: true, F1: 2, F2: v20}
            fmt.Println(v4, v2, v20)
            _ = v39
        } else {
            if (v4 || (v4 || v4)) {
                _ = len(v3)
                v4 = (1e10 > 3.14159)
            }
        }
        _ = v20
    }
    v40 := GenS5{F0: 1, F1: 423}
    fmt.Println(v4)
    v2 = (^(v2))
    for fi41 := 0; fi41 < 15; fi41++ {
        v2 -= v2
        _ = fi41
    }
    var v42 uint8 = 117
    if (!true) {
        v42 += v42
        var v43 float32 = 0.1
        for fi44 := 0; fi44 < 10; fi44++ {
            v42 = v42
            _ = fi44
        }
        _ = v43
    }
    for fi45 := 0; fi45 < 7; fi45++ {
        v42 = (v42 % ((v42) | 1))
        var v46 uint = v2
        v46 = v2
        _ = fi45
        _ = v46
    }
    v1 = append(v1, -742)
    v42 += v42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v8
    _ = v9
}
