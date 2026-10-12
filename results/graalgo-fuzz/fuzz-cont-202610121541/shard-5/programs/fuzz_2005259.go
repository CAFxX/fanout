package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint64 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint32; F1 string }

func main() {
    var v1 int64 = -1
    var v2 float64 = 0.1
    v3 := []int{-1}
    var v4 float64 = 3.14159
    var v5 int8 = 100
    var v6 uint32 = 1
    switch 4294967295 {
    case 42:
        switch ("hello" + "\n\t") {
        case "foo bar":
            if (-128 == v5) {
                _ = len(v3)
                _ = len(v3)
                v2 -= v4
            }
        default:
            if false {
                if true {
                    if ((-2.25 + 0.0) < (1e10 / 0.1)) {
                        for fi7 := 0; fi7 < 17; fi7++ {
                            v8 := GenS1{F0: 0, F1: "", F2: 553}
                            for fi9 := 0; fi9 < 12; fi9++ {
                                v10 := GenS1{F0: 1, F1: "hello", F2: 18446744073709551615}
                                switch (("foo bar" + "hello") + "foo bar") {
                                case "a":
                                    var v11 int32 = 316
                                    _ = v11
                                case "foo bar":
                                    _ = len(v3)
                                }
                                var v12 int32 = 2147483647
                                _ = fi9
                                _ = v10
                                _ = v12
                            }
                            for fi13 := 0; fi13 < 11; fi13++ {
                                var v14 bool = false
                                var v15 int8 = (v5 % ((v5) | 1))
                                if len(v3) > 0 {
                                    v3[0] = fi13
                                }
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                            _ = fi7
                            _ = v8
                        }
                    }
                    if ((float64(-2.25) / 0.1) <= v4) {
                        for fi16 := 0; fi16 < 12; fi16++ {
                            if (1.5 < ((-2.25 / 1.0) + 1.5)) {
                                v1 = 1
                                _ = len(v3)
                                v17 := GenS2{F0: 2, F1: "a"}
                                _ = v17
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v6 -= (v6 ^ (v6 % ((4294967295) | 1)))
                            }
                            fmt.Println(v6, v6)
                            fmt.Println(v5, v6)
                            _ = fi16
                        }
                        var v18 int8 = v5
                        v2 = v4
                        _ = v18
                    }
                }
            }
        }
    case -793:
        var v19 uint64 = 18446744073709551615
        _ = v19
    case -1:
        var v20 uint64 = 842
        _ = v20
    default:
        var v21 int64 = (^(v1))
        _ = v21
    }
    v22 := GenS3{F0: 305, F1: "\n\t"}
    if true {
        var v23 float64 = v2
        v4 = 3.14159
        _ = len(v3)
        _ = v23
    }
    if len(v3) > 0 {
        v3[0] = -463
    }
    v24 := GenS1{F0: 42, F1: "a", F2: 4294967295}
    v25 := GenS3{F0: 42, F1: "a"}
    v26 := GenS1{F0: 42, F1: "foo bar", F2: 4294967295}
    switch "" {
    case "":
        var v27 int = 32768
        _ = v27
    }
    var v28 int64 = v1
    _ = len(v3)
    v28 = (-(v1))
    v29 := GenS1{F0: 42, F1: "foo bar", F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
