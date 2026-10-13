package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int; F1 uint64; F2 int16 }
type GenS4 struct { F0 int16; F1 uint32 }

func main() {
    var v1 uint32 = 65535
    v2 := []uint32{v1, v1, 450, (v1 | v1)}
    v3 := []uint32{65535, v1, 0}
    var v4 uint8 = 130
    var v5 int64 = 1
    if false {
        v6 := GenS1{F0: uint32(v4)}
        fmt.Println(v4, v5, v4)
        _ = v6
    } else {
        if true {
            v4 = v4
            fmt.Println(v5, v4)
        } else {
            v5 = -799
        }
    }
    fmt.Println(v4, v4, v1)
    v4 -= v4
    var v7 float64 = (1.5 - ((-2.25 + 1e10) * 0.1))
    _ = len(v2)
    fmt.Println(v1, v1, v4)
    if (v5 <= 0) {
        fmt.Println(v4, v7)
        v7 += float64(1.5)
        if true {
            fmt.Println(v1)
            v5 += 0
        }
    }
    switch 2147483647 {
    case 431:
        v8 := GenS1{F0: v1}
        _ = v8
    case 32767:
        if len(v2) > 0 {
            v2[0] = v1
        }
    case 42:
        fmt.Println(v7, v4)
    }
    v7 += v7
    var v9 uint64 = 18446744073709551615
    var v10 float64 = v7
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        var v11 int8 = 109
        _ = v11
    } else {
        for fi12 := 0; fi12 < 7; fi12++ {
            fmt.Println(fi12, v9)
            if (2 != uint16(v9)) {
                v13 := GenS2{F0: 32767, F1: 42}
                if len(v2) > 0 {
                    v2[0] = (v1 ^ (v1 + v1))
                }
                switch fi12 {
                case 32767:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case -2147483648:
                    if (fi12 <= -9223372036854775808) {
                        switch (fi12 >> ((4294967295) & 63)) {
                        case -2147483648:
                            for fi14 := 0; fi14 < 3; fi14++ {
                                var v15 float64 = v7
                                switch ("" + "foo bar") {
                                case "foo bar":
                                    var v16 bool = (true || true)
                                    var v17 uint16 = 686
                                    if (!((fi12 > (-(fi14))) && v16)) {
                                        v18 := GenS1{F0: uint32(v17)}
                                        switch "a" {
                                        case "a":
                                            v4 = 0
                                        case "hello":
                                            if (!true) {
                                                v15 -= v10
                                                v2 = append(v2, 2)
                                            }
                                        case "foo bar":
                                            fmt.Println(v4, fi14)
                                        default:
                                            if ((v5 >= (v5 << ((364) & 63))) && (v5 == -177)) {
                                                switch ((("a" + ("\n\t" + "hello")) + "a") + ("" + "foo bar")) {
                                                case "":
                                                    v19 := GenS3{F0: (fi12 >> ((582) & 63)), F1: 1, F2: 1000}
                                                    _ = v19
                                                case "\n\t":
                                                    var v20 string = "a"
                                                    _ = v20
                                                case "foo bar":
                                                    v21 := GenS4{F0: 0, F1: (v1 << ((uint(v9)) & 63))}
                                                    _ = v21
                                                }
                                                v22 := GenS4{F0: 30, F1: v1}
                                                v9 += v9
                                                _ = v22
                                            }
                                        }
                                        var v23 bool = v16
                                        _ = v18
                                        _ = v23
                                    }
                                    _ = v16
                                    _ = v17
                                default:
                                    v9 += (v9 & v9)
                                }
                                v5 = v5
                                _ = fi14
                                _ = v15
                            }
                        case 128:
                            var v24 bool = (fi12 <= fi12)
                            _ = v24
                        case 32767:
                            var v25 uint64 = 2
                            _ = v25
                        default:
                            var v26 float32 = 0.0
                            _ = v26
                        }
                        var v27 uint8 = (v4 % ((v4) | 1))
                        v10 = (v7 / (v10 + v10))
                        _ = v27
                    } else {
                        v9 = 771
                        v7 = v10
                        v28 := GenS4{F0: int16(v1), F1: (v1 & 42)}
                        _ = v28
                    }
                case 4294967295:
                    var v29 uint32 = v1
                    _ = v29
                }
                _ = v13
            } else {
                fmt.Println(v7, fi12, v4)
            }
            v4 = v4
            _ = fi12
        }
        for fi30 := 0; fi30 < 12; fi30++ {
            for fi31 := 0; fi31 < 20; fi31++ {
                var v32 uint16 = 509
                v32 += v32
                _ = fi31
                _ = v32
            }
            var v33 uint32 = v1
            v34 := GenS4{F0: -654, F1: v33}
            _ = fi30
            _ = v33
            _ = v34
        }
        var v35 bool = true
        _ = v35
    }
    var v36 uint32 = (v1 & 0)
    switch 127 {
    case 42:
        fmt.Println(v10, v4, v7)
    case 156:
        v37 := GenS2{F0: -566, F1: 0}
        _ = v37
    }
    v38 := GenS1{F0: (^(v1))}
    var v39 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
