package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint32; F1 int16 }

func main() {
    var v1 int32 = 1
    var v2 uint32 = 2147483647
    var v3 float32 = ((3.14159 + (3.14159 - -2.25)) - 0.1)
    v4 := []string{"a", "foo bar"}
    v5 := []int{0, 144, 2147483647}
    v6 := make(map[int]int)
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int8 = -60
        v9 := GenS1{F0: true, F1: (v1 ^ (^(v1)))}
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v3, v1, v1)
    v1 -= v1
    v1 -= -953
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            if false {
                if true {
                    v12 := GenS1{F0: (!true), F1: -2147483648}
                    fmt.Println(v2, fi10, v2)
                    v1 += (v1 * -941)
                    _ = v12
                }
                var v13 uint8 = 42
                switch "hello" {
                case "foo bar":
                    var v14 uint64 = 0
                    if ((("a" + "\n\t") < "\n\t") && ((255 >= 255) && false)) {
                        var v15 uint16 = 42
                        _ = v15
                    }
                    _ = v14
                default:
                    fmt.Println(v3)
                }
                _ = v13
            }
            _ = len(v6)
            _ = fi11
        }
        v2 = (v2 / ((717) | 1))
        _ = fi10
    }
    var v16 uint16 = 387
    fmt.Println(v2, v16)
    fmt.Println(v2, v2, v16)
    var v17 float64 = (0.0 * 0.0)
    v5 = append(v5, 9223372036854775807)
    var v18 uint8 = 0
    fmt.Println(v18, v3)
    if (false || (v3 <= (v3 * (v3 - (float32((((v17 / 1.0) / ((v17 + ((float64(v3) - 1e10) + 1e10)) + 3.14159)) - (v17 / v17))) * v3))))) {
        for fi19 := 0; fi19 < 18; fi19++ {
            v3 -= (v3 * v3)
            fmt.Println(v1)
            v2 -= 4294967295
            _ = fi19
        }
        for fi20 := 0; fi20 < 11; fi20++ {
            for fi21 := 0; fi21 < 13; fi21++ {
                if true {
                    switch fi21 {
                    case 42:
                        if true {
                            v3 = (float32((((v17 + v17) + v17) * 0.0)) - v3)
                        }
                    }
                    v22 := GenS1{F0: true, F1: 0}
                    if false {
                        var v23 float32 = v3
                        if ((42 <= 42) || true) {
                            if (557 == int64(v18)) {
                                fmt.Println(v16, v2, fi20)
                                v24 := GenS1{F0: true, F1: (v1 | (v1 / ((v1) | 1)))}
                                v25 := GenS1{F0: true, F1: v1}
                                _ = v24
                                _ = v25
                            }
                            v23 -= (1e10 + v23)
                        }
                        var v26 uint8 = 176
                        switch ("a" + "") {
                        case "foo bar":
                            var v27 bool = true
                            _ = v27
                        case "":
                            v6[(fi20 * (fi20 - fi20))] = (-(fi20))
                            v28 := GenS2{F0: (fi20 >> ((0) & 63)), F1: -9223372036854775808, F2: 0}
                            _ = v28
                        case "a":
                            if len(v5) > 0 {
                                v5[0] = (fi21 & -1)
                            }
                        }
                        v2 -= v2
                        _ = v23
                        _ = v26
                    } else {
                        var v29 bool = false
                        _ = v29
                    }
                    var v30 float32 = v3
                    _ = v22
                    _ = v30
                }
                _ = fi21
            }
            if true {
                v31 := v6[fi20]
                fmt.Println(v16, v31)
                for fi32 := 0; fi32 < 3; fi32++ {
                    fmt.Println(fi20, v2)
                    _ = fi32
                }
                _ = v31
            } else {
                switch (("\n\t" + "hello") + "foo bar") {
                case "hello":
                    if ((("foo bar" < "") || false) && ((2 != (v16 ^ v16)) && false)) {
                        for fi33 := 0; fi33 < 4; fi33++ {
                            v16 = v16
                            if (1.5 > 1.5) {
                                var v34 int8 = -99
                                var v35 uint64 = 0
                                var v36 uint = 18446744073709551615
                                _ = v34
                                _ = v35
                                _ = v36
                            } else {
                                v37 := GenS2{F0: 1, F1: 1, F2: 0}
                                v38 := GenS1{F0: false, F1: v1}
                                _ = v37
                                _ = v38
                            }
                            _ = fi33
                        }
                    }
                case "foo bar":
                    v16 += (v16 % (((^(v16))) | 1))
                }
                for fi39 := 0; fi39 < 13; fi39++ {
                    var v40 uint64 = 559
                    _ = fi39
                    _ = v40
                }
                v41 := GenS3{F0: 0, F1: 351, F2: 4294967295}
                _ = v41
            }
            var v42 uint = 708
            _ = fi20
            _ = v42
        }
        v3 += v3
    } else {
        var v43 uint8 = (v18 << ((255) & 63))
        _ = v43
    }
    switch 9223372036854775807 {
    case 128:
        var v44 uint32 = (v2 >> ((625) & 63))
        _ = v44
    case 32768:
        var v45 uint32 = v2
        _ = v45
    case 2147483647:
        v46 := GenS4{F0: 2, F1: 42}
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
