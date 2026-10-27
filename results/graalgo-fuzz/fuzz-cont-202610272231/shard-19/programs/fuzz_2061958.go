package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 int16; F2 int16 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint64, p1 uint) string {
    var v7 float32 = 3.14159
    p1 = (p1 & p1)
    fmt.Println(p1)
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := GenS1{F0: 65535, F1: "", F2: 1}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return ("" + "\n\t")
}

func main() {
    var v1 uint32 = 4294967295
    var v2 float32 = 3.14159
    var v3 int = 4294967295
    var v4 uint64 = 0
    var v5 int32 = -369
    var v6 string = ""
    var v10 int32 = 837
    v11 := GenS1{F0: 0, F1: "a", F2: 2}
    v2 += ((float32(1.5) - (0.1 - 0.0)) * v2)
    if true {
        if (int8(v1) > 1) {
            for fi12 := 0; fi12 < 2; fi12++ {
                v13 := GenS2{F0: -128, F1: v4, F2: -420}
                switch v3 {
                case 0:
                    v14 := []int{fi12, 42, fi12, (^(v3))}
                    _ = v14
                case 506:
                    for fi15 := 0; fi15 < 9; fi15++ {
                        switch v6 {
                        case "\n\t":
                            v16 := GenS3{F0: 42}
                            _ = v16
                        case "":
                            v5 = (v10 + 2147483647)
                        case "foo bar":
                            v17 := GenS4{F0: v6, F1: 0, F2: -643}
                            _ = v17
                        default:
                            var v18 int = fi12
                            _ = v18
                        }
                        v1 = 4294967295
                        _ = fi15
                    }
                }
                if true {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        v2 = float32(3.14159)
                        _ = fi19
                    }
                    for fi20 := 0; fi20 < 18; fi20++ {
                        var v21 uint32 = 2147483648
                        var v22 int8 = 127
                        var v23 uint32 = 2147483648
                        if false {
                            if (964 <= 729) {
                                var v24 uint64 = v4
                                _ = v24
                            }
                            v25 := []int{(-(v3)), fi20, fi12, 137}
                            v25 = append(v25, fi12)
                            v26 := GenS4{F0: (v6 + (((((("hello" + (v6 + (v6 + ((v6 + v6) + "")))) + ("" + v6)) + ((v6 + v6) + "hello")) + ("a" + (v6 + "hello"))) + v6) + (v6 + v6))), F1: -1, F2: 1000}
                            var v27 uint16 = 33
                            if (false || false) {
                                fmt.Println(fi20)
                                v1 = (^(v1))
                            } else {
                                var v28 uint32 = (v21 + v23)
                                _ = v28
                            }
                            _ = len(v25)
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        var v29 float64 = 1.5
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                        _ = v29
                    }
                }
                _ = fi12
                _ = v13
            }
            fmt.Println(v5, v5)
        } else {
            for fi30 := 0; fi30 < 18; fi30++ {
                v31 := []int{(fi30 | v3), (fi30 * (v3 % ((fi30) | 1)))}
                _ = fi30
                _ = v31
            }
            switch v3 {
            case -9223372036854775808:
                if true {
                    var v32 string = v6
                    var v33 uint = 2
                    var v34 uint32 = (^(v1))
                    _ = v32
                    _ = v33
                    _ = v34
                }
            case 405:
                for fi35 := 0; fi35 < 5; fi35++ {
                    var v36 uint = 2
                    fmt.Println(v3)
                    _ = fi35
                    _ = v36
                }
            }
        }
        var v37 string = v6
        _ = v37
    }
    var v38 float32 = v2
    if (!false) {
        for fi39 := 0; fi39 < 18; fi39++ {
            fmt.Println(v4, fi39, v38)
            v40 := []int{(v3 >> ((4294967295) & 63)), (^(v3)), 732}
            if (-32768 <= 42) {
                var v41 int8 = -66
                v3 = (v3 << ((1) & 63))
                _ = v41
            }
            _ = fi39
            _ = v40
        }
        v42 := []int{}
        var v43 uint = 956
        _ = v42
        _ = v43
    } else {
        v10 = 293
        fmt.Println(v4, v2)
    }
    for fi44 := 0; fi44 < 6; fi44++ {
        v45 := []int{v3, fi44, (-(fi44)), (fi44 & 2147483647)}
        var v46 int64 = -398
        switch (v3 - fi44) {
        case 0:
            v10 = v10
        }
        _ = fi44
        _ = v45
        _ = v46
    }
    var v47 uint8 = 2
    v1 += (^(v1))
    v48 := []int{}
    v4 = v4
    v49 := GenS5{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
