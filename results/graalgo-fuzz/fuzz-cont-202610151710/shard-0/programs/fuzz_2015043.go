package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint; F2 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16 }
func genFn0() int32 {
    var v7 uint32 = 2147483648
    v7 = v7
    v7 = (v7 ^ 693)
    _ = v7
    return -655
}
func genFn1(p0 int32, p1 uint32, p2 uint64) int {
    var v8 uint64 = p2
    v8 = 2
    switch -9223372036854775808 {
    case 119:
        if true {
            for fi9 := 0; fi9 < 11; fi9++ {
                v8 += p2
                _ = fi9
            }
            fmt.Println(p1, p0, p0)
        } else {
            var v10 string = "\n\t"
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 uint = 568
        _ = fi11
        _ = v12
    }
    v13 := []int{2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v8
    return 127
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 1
    var v3 uint8 = 214
    v4 := make(map[string]uint32)
    v5 := []uint32{42, 213, 4294967295, 2}
    var v6 float32 = 3.14159
    if (v6 >= float32((float64(1.5) + 0.1))) {
        v2 = (v2 + (v2 | v2))
        var v14 int8 = -128
        for fi15 := 0; fi15 < 12; fi15++ {
            _ = len(v5)
            for fi16 := 0; fi16 < 19; fi16++ {
                v17 := v4[(("\n\t" + "hello") + ("foo bar" + ("" + "hello")))]
                var v18 int64 = 0
                _ = fi16
                _ = v17
                _ = v18
            }
            fmt.Println(v6, v14)
            _ = fi15
        }
        _ = v14
    } else {
        v19 := GenS1{F0: -2147483648}
        if len(v5) > 0 {
            v5[0] = 754
        }
        _ = v19
    }
    v20 := GenS1{F0: 0}
    v2 += (v2 | genFn1(-1, 42, 0))
    v6 = 1e10
    var v21 uint64 = 1
    for fi22 := 0; fi22 < 12; fi22++ {
        switch "hello" {
        case "hello":
            v5 = append(v5, 65535)
        default:
            _ = len(v5)
        }
        v2 = v2
        _ = fi22
    }
    v23 := GenS1{F0: 1}
    if (-2147483648 <= 0) {
        v3 += 42
        var v24 string = "\n\t"
        _ = v24
    }
    if (255 < 752) {
        var v25 int = genFn1(2147483647, 65535, 18446744073709551615)
        for fi26 := 0; fi26 < 5; fi26++ {
            fmt.Println(v21)
            _ = fi26
        }
        v27 := GenS1{F0: 2147483647}
        _ = v25
        _ = v27
    } else {
        v28 := GenS2{F0: 0}
        for fi29 := 0; fi29 < 18; fi29++ {
            _ = len(v4)
            for fi30 := 0; fi30 < 12; fi30++ {
                var v31 float32 = 0.0
                v31 -= v6
                if true {
                    var v32 int8 = -1
                    v2 += 1
                    _ = v32
                }
                _ = fi30
                _ = v31
            }
            _ = fi29
        }
        _ = v28
    }
    switch "a" {
    case "\n\t":
        for fi33 := 0; fi33 < 1; fi33++ {
            if true {
                for fi34 := 0; fi34 < 17; fi34++ {
                    v2 += v2
                    if ((v3 >= v3) || true) {
                        switch 127 {
                        case -2147483648:
                            v21 = v21
                        default:
                            v6 = 0.0
                        }
                        var v35 int8 = -1
                        if ((!false) && (0.0 != float64(v6))) {
                            var v36 bool = false
                            if (v3 != uint8(v21)) {
                                var v37 int64 = 987
                                fmt.Println(v37, v6)
                                _ = v37
                            } else {
                                var v38 bool = (0 <= 0)
                                if true {
                                    v4["hello"] = 2147483648
                                    if v36 {
                                        v21 = (v21 ^ v21)
                                    }
                                    v39 := GenS1{F0: 1}
                                    _ = v39
                                }
                                switch (fi33 << ((406) & 63)) {
                                case 2147483647:
                                    v40 := GenS1{F0: 0}
                                    _ = v40
                                }
                                var v41 uint8 = (v3 >> ((65535) & 63))
                                if false {
                                    var v42 string = ("" + "")
                                    switch v42 {
                                    case "\n\t":
                                        v43 := v1[1]
                                        _ = v43
                                    case "hello":
                                        v44 := GenS3{F0: v36, F1: 42, F2: int32(v3)}
                                        _ = v44
                                    }
                                    _ = v42
                                } else {
                                    if v36 {
                                        fmt.Println(v38)
                                        v38 = v36
                                        fmt.Println(v6, fi34)
                                    }
                                    v45 := GenS4{F0: (int64(v21) > 0)}
                                    _ = v45
                                }
                                _ = v38
                                _ = v41
                            }
                            _ = v36
                        } else {
                            delete(v1, 4294967295)
                            if false {
                                v46 := GenS5{F0: 422}
                                _ = v46
                            } else {
                                var v47 string = "\n\t"
                                _ = v47
                            }
                            var v48 bool = false
                            _ = v48
                        }
                        delete(v4, ("a" + ("a" + "a")))
                        _ = v35
                    }
                    _ = fi34
                }
                v1[2147483647] = (fi33 / ((504) | 1))
            } else {
                v3 -= (^(v3))
                v49 := GenS4{F0: (false && false)}
                _ = v49
            }
            _ = fi33
        }
    }
    fmt.Println(v3, v6)
    v2 = (v2 << ((1) & 63))
    v50 := GenS1{F0: 835}
    if ("hello" <= "foo bar") {
        v2 = v2
        _ = len(v5)
        fmt.Println(v3)
    }
    var v51 int64 = 9223372036854775807
    var v52 uint8 = (v3 / (((v3 | (v3 / ((1) | 1)))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}
