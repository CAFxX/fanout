package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS3 struct { F0 bool; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS7 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint32) uint16 {
    var v5 uint32 = 2
    var v6 uint8 = 127
    switch "" {
    case "a":
        if false {
            v5 = p0
        }
    case "\n\t":
        v7 := GenS1{F0: 42, F1: 0}
        _ = v7
    case "":
        var v8 int = -1
        _ = v8
    default:
        if true {
            var v9 int = 9223372036854775807
            switch "hello" {
            case "hello":
                v10 := []int{v9, 1, 42}
                _ = v10
            default:
                fmt.Println(v5, v5, v5)
            }
            for fi11 := 0; fi11 < 10; fi11++ {
                var v12 uint16 = 524
                if (0 < 18446744073709551615) {
                    var v13 uint8 = v6
                    _ = v13
                } else {
                    fmt.Println(fi11, v5, v12)
                }
                _ = fi11
                _ = v12
            }
            _ = v9
        } else {
            v14 := []int{-663, -1}
            var v15 float64 = -2.25
            _ = v14
            _ = v15
        }
    }
    fmt.Println(v6)
    _ = p0
    _ = v5
    _ = v6
    return 255
}
func genFn1(p0 int, p1 int8) int32 {
    fmt.Println(p0, p0)
    switch "a" {
    case "foo bar":
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 int64 = -196
            _ = fi16
            _ = v17
        }
    case "hello":
        if false {
            var v18 string = "foo bar"
            _ = v18
        } else {
            v19 := GenS2{F0: (99 > 1), F1: 32767, F2: 174}
            p1 -= (p1 - (p1 + p1))
            _ = v19
        }
    default:
        var v20 uint8 = 2
        _ = v20
    }
    var v21 int64 = 1
    v21 = 42
    _ = p0
    _ = p1
    _ = v21
    return -1
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 241
    v3 := make(map[uint32]int)
    var v4 uint = 912
    var v22 int8 = 127
    v23 := GenS3{F0: (true || ("a" >= "hello")), F1: (^(v22)), F2: uint64(v4)}
    if false {
        v24 := GenS2{F0: ((true && true) || true), F1: 789, F2: genFn0(0)}
        _ = v24
    }
    var v25 uint8 = (v2 ^ v2)
    switch ("\n\t" + "") {
    case "foo bar":
        v4 += v4
    }
    v25 = v25
    v25 -= 2
    switch (("" + "") + (("hello" + ("foo bar" + ("foo bar" + "hello"))) + "")) {
    case "":
        var v26 float64 = 1.5
        _ = v26
    }
    fmt.Println(v2, v22)
    switch "" {
    case "foo bar":
        var v27 int32 = genFn1(4294967295, (v22 ^ v22))
        _ = v27
    default:
        fmt.Println(v4)
    }
    var v28 bool = ((1e10 * 3.14159) <= 1e10)
    v4 = uint(v2)
    v29 := []int{}
    for fi30 := 0; fi30 < 5; fi30++ {
        if true {
            var v31 uint64 = 18446744073709551615
            v25 -= v25
            _ = v31
        }
        var v32 uint32 = 2
        _ = fi30
        _ = v32
    }
    if v28 {
        for fi33 := 0; fi33 < 13; fi33++ {
            for fi34 := 0; fi34 < 10; fi34++ {
                switch ((("\n\t" + "foo bar") + ("a" + "a")) + "foo bar") {
                case "":
                    v35 := GenS1{F0: -1, F1: 1}
                    _ = v35
                case "a":
                    if true {
                        var v36 uint16 = 0
                        v29 = append(v29, -298)
                        if (true && (!v28)) {
                            v2 -= uint8(v4)
                            v37 := GenS1{F0: 9223372036854775807, F1: v2}
                            v2 = v2
                            _ = v37
                        } else {
                            v38 := GenS4{F0: 2, F1: v25, F2: 1}
                            if len(v29) > 0 {
                                v29[0] = fi34
                            }
                            v39 := GenS5{F0: (v2 ^ 127), F1: 42}
                            _ = v38
                            _ = v39
                        }
                        v29 = append(v29, fi33)
                        _ = v36
                    } else {
                        v40 := GenS2{F0: (float32(((0.1 - 1.5) + 1.5)) > (0.0 + (float32((float64(((((0.1 - ((3.14159 - -2.25) + 0.0)) * float32(0.1)) + 0.1) - ((((0.0 - 0.1) / -2.25) * float32(float64(0.0))) / 1.0))) * 3.14159)) / 1e10))), F1: 2, F2: 42}
                        if (42 != 2147483647) {
                            var v41 int32 = 0
                            if ((v41 << ((v4) & 63)) >= -2147483648) {
                                fmt.Println(v25, v41)
                                if true {
                                    var v42 int = 32767
                                    _ = v42
                                }
                                if v28 {
                                    switch "hello" {
                                    case "a":
                                        v41 -= 42
                                    case "\n\t":
                                        _ = len(v29)
                                    case "hello":
                                        if (v28 || (("foo bar" + "a") < "")) {
                                            var v43 uint64 = 881
                                            v44 := GenS6{F0: 843, F1: uint16(v43), F2: v22}
                                            _ = v43
                                            _ = v44
                                        }
                                    }
                                    if len(v29) > 0 {
                                        v29[0] = int(v4)
                                    }
                                    _ = len(v29)
                                } else {
                                    if len(v29) >= 2 {
                                        v29 = v29[1:len(v29)]
                                    }
                                    var v45 uint8 = 2
                                    _ = v45
                                }
                            }
                            v22 = 1
                            if ((v22 | (v22 & 100)) >= (^(v22))) {
                                var v46 int = fi33
                                var v47 int8 = int8(v41)
                                fmt.Println(fi33)
                                _ = v46
                                _ = v47
                            }
                            _ = v41
                        }
                        if false {
                            v22 = v22
                        }
                        _ = v40
                    }
                case "foo bar":
                    var v48 uint8 = (^(v25))
                    _ = v48
                }
                _ = fi34
            }
            var v49 int16 = -460
            v50 := GenS7{F0: (v2 >> ((v4) & 63)), F1: (-(v22))}
            _ = fi33
            _ = v49
            _ = v50
        }
        for fi51 := 0; fi51 < 20; fi51++ {
            var v52 int = -9223372036854775808
            _ = fi51
            _ = v52
        }
        var v53 float32 = 1e10
        _ = v53
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
