package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 string; F1 bool; F2 int8 }
func genFn0(p0 int32, p1 int8, p2 uint64) int32 {
    v4 := GenS1{F0: "foo bar"}
    if true {
        v5 := []int{9223372036854775807, 32768, -961, 32768}
        var v6 float32 = ((0.0 - (-2.25 * 1.5)) - 1.5)
        _ = v5
        _ = v6
    } else {
        fmt.Println(p2)
        var v7 uint8 = 2
        for fi8 := 0; fi8 < 20; fi8++ {
            switch fi8 {
            case 726:
                v9 := GenS1{F0: "a"}
                _ = v9
            case 32768:
                switch int(p2) {
                case 32767:
                    if true {
                        for fi10 := 0; fi10 < 18; fi10++ {
                            v11 := GenS1{F0: "a"}
                            _ = fi10
                            _ = v11
                        }
                        if true {
                            v12 := []int{(fi8 * fi8)}
                            v13 := GenS2{F0: "\n\t"}
                            _ = v12
                            _ = v13
                        }
                    }
                case 0:
                    v14 := []int{}
                    _ = v14
                case 127:
                    p1 -= (p1 | p1)
                default:
                    p2 -= (p2 + 547)
                }
            case 9223372036854775807:
                v15 := []int{(fi8 << ((0) & 63))}
                _ = v15
            default:
                p1 += 64
            }
            _ = fi8
        }
        _ = v7
    }
    var v16 uint16 = 42
    if false {
        p1 = (p1 ^ -1)
    } else {
        if (((p1 <= (p1 << ((42) & 63))) && (false || (false || ("foo bar" == "\n\t")))) && true) {
            fmt.Println(p0, p0, v16)
            fmt.Println(p0)
        } else {
            var v17 uint32 = 407
            v18 := GenS2{F0: "hello"}
            _ = v17
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v4
    return p0
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]uint32)
    var v3 float32 = 0.1
    v1 = (v1 << ((0) & 63))
    v19 := []int{128, 1}
    var v20 int64 = 1
    for fi21 := 0; fi21 < 1; fi21++ {
        if (!true) {
            var v22 uint32 = v1
            v23 := GenS1{F0: "hello"}
            _ = v22
            _ = v23
        } else {
            _ = len(v2)
            v24 := GenS1{F0: "a"}
            v25 := v2[(v1 << ((18446744073709551615) & 63))]
            _ = v24
            _ = v25
        }
        _ = fi21
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        v27 := GenS3{F0: 42}
        v1 = v1
        if true {
            if true {
                v19 = append(v19, (fi26 & fi26))
                for fi28 := 0; fi28 < 7; fi28++ {
                    var v29 float32 = ((v3 / (1e10 + 0.1)) * (v3 + -2.25))
                    v30 := GenS4{F0: 4294967295, F1: -941}
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
        }
        _ = fi26
        _ = v27
    }
    for fi31 := 0; fi31 < 10; fi31++ {
        switch fi31 {
        case 1:
            var v32 uint8 = 183
            _ = v32
        default:
            for fi33 := 0; fi33 < 3; fi33++ {
                var v34 bool = true
                switch fi31 {
                case -9223372036854775808:
                    v35 := GenS2{F0: ("foo bar" + "hello")}
                    _ = v35
                case -312:
                    if (v34 && (!(false || (332 < 18446744073709551615)))) {
                        fmt.Println(v1, v34)
                        switch (fi33 / ((fi31) | 1)) {
                        case 128:
                            v20 = 329
                        case 32767:
                            if ((!(true || v34)) && (v3 >= (v3 * v3))) {
                                v36 := GenS5{F0: "hello", F1: v34, F2: 127}
                                _ = v36
                            } else {
                                v1 += v1
                            }
                        default:
                            if (v20 == 0) {
                                var v37 int = 0
                                if v34 {
                                    v34 = v34
                                    v37 = 32768
                                }
                                var v38 bool = v34
                                _ = v37
                                _ = v38
                            } else {
                                _ = len(v19)
                                v34 = (4294967295 < 65535)
                            }
                        }
                        var v39 uint32 = v1
                        _ = v39
                    }
                    v2[(v1 - v1)] = 4294967295
                    v3 += float32(3.14159)
                }
                v40 := GenS3{F0: 18446744073709551615}
                _ = fi33
                _ = v34
                _ = v40
            }
        }
        v2[(v1 >> ((889) & 63))] = 2
        v41 := GenS1{F0: ("hello" + ((("" + "\n\t") + ("a" + "a")) + ""))}
        _ = fi31
        _ = v41
    }
    fmt.Println(v20)
    if true {
        var v42 uint8 = 1
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
