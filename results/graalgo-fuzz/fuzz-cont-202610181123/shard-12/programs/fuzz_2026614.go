package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 uint16; F1 uint; F2 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 int64; F1 int; F2 uint }

func main() {
    var v1 int = -9223372036854775808
    var v2 uint32 = 0
    var v3 int16 = -1
    v4 := make(map[string]string)
    v5 := []uint32{v2, (v2 + 376), 2147483647, (v2 & 2147483648)}
    v6 := make(map[int]int)
    for fi7 := 0; fi7 < 6; fi7++ {
        v3 -= v3
        if len(v5) > 0 {
            v5[0] = (v2 << ((255) & 63))
        }
        fmt.Println(v2)
        _ = fi7
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v8 := GenS1{F0: true, F1: (v2 * (v2 | (^(v2)))), F2: 893}
    var v9 uint64 = 0
    if ((1 >= 0) || true) {
        var v10 int64 = -926
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 int32 = 202
            for fi13 := 0; fi13 < 15; fi13++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                switch 32767 {
                case 32767:
                    v1 -= v1
                case 303:
                    v6[(fi11 + v1)] = (fi13 & fi13)
                case 127:
                    fmt.Println(v3, v1)
                }
                _ = fi13
            }
            if true {
                v14 := GenS1{F0: true, F1: (v2 | v2), F2: 658}
                _ = v14
            } else {
                switch v1 {
                case 128:
                    switch "\n\t" {
                    case "\n\t":
                        v2 = 42
                    case "hello":
                        v3 += (v3 - (v3 - 1))
                    case "foo bar":
                        v15 := GenS2{F0: false, F1: 0}
                        _ = v15
                    }
                case 127:
                    delete(v4, ((((("\n\t" + "\n\t") + "a") + "") + ((("a" + "foo bar") + "\n\t") + "a")) + "\n\t"))
                case -2147483648:
                    if true {
                        for fi16 := 0; fi16 < 17; fi16++ {
                            switch (v1 + fi11) {
                            case -2147483648:
                                fmt.Println(v10)
                            case 979:
                                v17 := v4["hello"]
                                v18 := GenS1{F0: false, F1: (v2 * 1), F2: 612}
                                _ = v17
                                _ = v18
                            }
                            var v19 uint32 = 112
                            var v20 uint16 = uint16(v2)
                            _ = fi16
                            _ = v19
                            _ = v20
                        }
                        for fi21 := 0; fi21 < 10; fi21++ {
                            var v22 int8 = -124
                            var v23 float32 = 1e10
                            if (!false) {
                                var v24 uint8 = 42
                                v6[32767] = 128
                                v1 += fi21
                                v25 := GenS3{F0: int16(fi11), F1: false}
                                v24 = v24
                                if ((1.5 * v23) >= 1.5) {
                                    v26 := GenS4{F0: "", F1: (^(v9))}
                                    if (18446744073709551615 == 1) {
                                        v22 = (-(v22))
                                    } else {
                                        v27 := GenS3{F0: int16(v12), F1: true}
                                        _ = v27
                                    }
                                    v12 = -2147483648
                                    _ = v26
                                }
                                var v28 float32 = (v23 / v23)
                                switch (-(v1)) {
                                case 32768:
                                    if len(v5) > 0 {
                                        v5[0] = 781
                                    }
                                case 42:
                                    var v29 int64 = v10
                                    var v30 int64 = int64(v3)
                                    _ = v29
                                    _ = v30
                                }
                                if (v12 != v12) {
                                    var v31 int16 = (^(v3))
                                    if ((!(false || false)) || (42 != (v12 / (((-(v12))) | 1)))) {
                                        fmt.Println(v23, v10)
                                        _ = len(v5)
                                        switch fi11 {
                                        case 2147483647:
                                            v32 := GenS5{F0: 0, F1: 4294967295, F2: -782}
                                            _ = v32
                                        case 1:
                                            if ((!(!(2 < (v24 >> ((23) & 63))))) || (!(("foo bar" + "") >= ""))) {
                                                v5 = append(v5, (v2 | v2))
                                                var v33 int32 = v12
                                                _ = v33
                                            } else {
                                                v3 = -751
                                            }
                                        default:
                                            v10 = (v10 * -1)
                                        }
                                    }
                                    var v34 uint = 42
                                    _ = v31
                                    _ = v34
                                }
                                _ = v24
                                _ = v25
                                _ = v28
                            }
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                    }
                default:
                    v35 := GenS6{F0: 0}
                    _ = v35
                }
            }
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        if true {
            var v36 int64 = 9223372036854775807
            fmt.Println(v1)
            v37 := GenS1{F0: true, F1: v2, F2: 255}
            _ = v36
            _ = v37
        }
        fmt.Println(v1, v9, v9)
        v38 := GenS7{F0: (^(v1)), F1: 430}
        _ = v38
    }
    var v39 int64 = int64(v1)
    switch "\n\t" {
    case "foo bar":
        var v40 uint64 = v9
        _ = v40
    default:
        _ = len(v4)
    }
    fmt.Println(v3)
    var v41 int8 = 1
    for fi42 := 0; fi42 < 14; fi42++ {
        var v43 uint16 = 460
        _ = fi42
        _ = v43
    }
    if false {
        _ = len(v5)
        v41 -= v41
        v44 := GenS6{F0: 1}
        _ = v44
    }
    fmt.Println(v3)
    var v45 bool = (!(837 >= 1))
    v46 := GenS3{F0: (v3 * -32768), F1: true}
    v47 := GenS8{F0: (v39 << ((42) & 63)), F1: (v1 - v1), F2: uint(v2)}
    v5 = append(v5, 42)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v41
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
