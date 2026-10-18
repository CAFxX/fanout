package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int; F1 uint16; F2 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16; F1 bool }
func genFn0(p0 int16, p1 bool, p2 uint64) uint64 {
    p1 = false
    var v5 int = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}

func main() {
    var v1 int16 = -32768
    var v2 int = -726
    var v3 float32 = 0.1
    v4 := []uint32{2147483648}
    var v6 string = ("" + "")
    var v7 int32 = 311
    switch 1 {
    case 42:
        if len(v4) > 0 {
            v4[0] = 112
        }
    case 9223372036854775807:
        v8 := GenS1{F0: 18446744073709551615}
        _ = v8
    case -2147483648:
        v6 = ""
    default:
        for fi9 := 0; fi9 < 16; fi9++ {
            v4 = append(v4, 2147483648)
            _ = fi9
        }
    }
    if false {
        var v10 bool = (4294967295 < 2147483647)
        if (!v10) {
            v11 := GenS2{F0: 255, F1: v10}
            for fi12 := 0; fi12 < 7; fi12++ {
                v13 := GenS3{F0: (v2 & v2), F1: 2, F2: 867}
                var v14 uint32 = 2147483648
                if (!v10) {
                    var v15 uint64 = 778
                    v16 := GenS2{F0: 18446744073709551615, F1: (!true)}
                    for fi17 := 0; fi17 < 17; fi17++ {
                        v18 := GenS1{F0: 279}
                        if false {
                            if (785 <= fi12) {
                                v6 = ((v6 + v6) + "foo bar")
                            }
                        } else {
                            fmt.Println(fi17)
                            var v19 int64 = -9223372036854775808
                            _ = v19
                        }
                        _ = fi17
                        _ = v18
                    }
                    _ = v15
                    _ = v16
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            var v20 float64 = ((3.14159 / 1e10) * 0.0)
            _ = v11
            _ = v20
        } else {
            var v21 string = v6
            for fi22 := 0; fi22 < 9; fi22++ {
                v23 := GenS4{F0: v10}
                for fi24 := 0; fi24 < 2; fi24++ {
                    if true {
                        fmt.Println(fi22, v1, v7)
                    }
                    _ = fi24
                }
                _ = fi22
                _ = v23
            }
            _ = v21
        }
        _ = v10
    } else {
        if len(v4) > 0 {
            v4[0] = 1
        }
    }
    var v25 int16 = (v1 | -1)
    var v26 int64 = -982
    v26 = 593
    if true {
        for fi27 := 0; fi27 < 4; fi27++ {
            for fi28 := 0; fi28 < 15; fi28++ {
                v26 -= v26
                var v29 bool = true
                v30 := GenS4{F0: v29}
                var v31 string = "foo bar"
                if ((v3 - (-2.25 * v3)) < (-2.25 * v3)) {
                    fmt.Println(fi27)
                    if len(v4) > 0 {
                        v4[0] = 2
                    }
                } else {
                    v25 = 0
                    v4 = append(v4, 0)
                    if v29 {
                        switch 42 {
                        case 42:
                            v31 = v6
                        case 1:
                            if v29 {
                                v6 = ""
                                var v32 int64 = (v26 | v26)
                                _ = v32
                            } else {
                                v33 := GenS3{F0: v2, F1: 0, F2: 65535}
                                switch v31 {
                                case "":
                                    v25 = v25
                                case "\n\t":
                                    var v34 int16 = 0
                                    _ = v34
                                case "a":
                                    var v35 int = fi27
                                    _ = v35
                                default:
                                    var v36 int32 = 771
                                    if (v26 >= v26) {
                                        v37 := GenS3{F0: int(v25), F1: 400, F2: 42}
                                        var v38 uint32 = 624
                                        _ = v37
                                        _ = v38
                                    } else {
                                        var v39 float32 = (v3 + (v3 + v3))
                                        _ = v39
                                    }
                                    _ = v36
                                }
                                v40 := GenS4{F0: v29}
                                _ = v33
                                _ = v40
                            }
                        case 172:
                            v41 := GenS5{F0: (v1 << ((385) & 63)), F1: (1 >= 2147483647)}
                            _ = v41
                        }
                    }
                }
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            _ = fi27
        }
    }
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    v4 = append(v4, 1)
    fmt.Println(v3, v1)
    var v42 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v42
    _ = v6
    _ = v7
}
