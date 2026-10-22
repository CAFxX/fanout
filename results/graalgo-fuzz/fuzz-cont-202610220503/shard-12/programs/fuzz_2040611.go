package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 string; F2 int32 }
type GenS5 struct { F0 int8; F1 int64; F2 bool }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    var v3 int64 = -559
    var v4 int16 = 42
    var v5 string = "hello"
    v6 := GenS1{F0: 715, F1: 42}
    v3 = (-(v3))
    switch "hello" {
    case "hello":
        var v7 int64 = (v3 | -9223372036854775808)
        _ = v7
    case "a":
        v8 := GenS2{F0: 212, F1: v4}
        _ = v8
    default:
        if (42 >= 889) {
            v3 = (v3 + (v3 << ((4294967295) & 63)))
            var v9 int8 = 100
            _ = v9
        }
    }
    var v10 uint32 = 2147483647
    v3 = v3
    for fi11 := 0; fi11 < 20; fi11++ {
        if (!false) {
            if (v5 > "\n\t") {
                if (413 <= v4) {
                    v10 += (^(v10))
                } else {
                    var v12 float64 = 1.5
                    _ = v12
                }
                v3 = -9223372036854775808
            } else {
                v13 := GenS1{F0: 64, F1: 65535}
                _ = v13
            }
            switch "a" {
            case "":
                v14 := v2[(("" + "foo bar") + v5)]
                _ = v14
            case "a":
                v15 := []int{9223372036854775807, fi11, fi11}
                _ = v15
            case "foo bar":
                fmt.Println(v5, v4)
            default:
                for fi16 := 0; fi16 < 4; fi16++ {
                    var v17 float64 = (1e10 * 0.0)
                    if (739 < 952) {
                        if false {
                            v5 = ""
                            v18 := GenS1{F0: 0, F1: 0}
                            v19 := GenS3{F0: 42}
                            _ = v18
                            _ = v19
                        }
                        if true {
                            switch v5 {
                            case "a":
                                var v20 uint16 = 2
                                _ = v20
                            default:
                                v21 := GenS1{F0: 1, F1: 255}
                                _ = v21
                            }
                        } else {
                            var v22 string = v5
                            v3 -= 1
                            _ = v22
                        }
                    }
                    v17 += (1e10 / v17)
                    fmt.Println(fi16)
                    _ = fi16
                    _ = v17
                }
            }
        }
        for fi23 := 0; fi23 < 3; fi23++ {
            if (!true) {
                if (42 > 0) {
                    v3 += v3
                    v5 = (v5 + "a")
                }
            }
            v4 = 0
            _ = fi23
        }
        _ = fi11
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        v4 = 0
        _ = fi24
    }
    fmt.Println(v3)
    for fi25 := 0; fi25 < 13; fi25++ {
        var v26 uint16 = uint16(v10)
        _ = fi25
        _ = v26
    }
    v4 -= 231
    v27 := GenS4{F0: v4, F1: v5, F2: 1}
    var v28 uint64 = 0
    switch 943 {
    case 0:
        v28 += (v28 % (((v28 * (v28 + (v28 >> ((129) & 63))))) | 1))
    case -717:
        for fi29 := 0; fi29 < 6; fi29++ {
            var v30 int = fi29
            var v31 string = (v5 + "")
            _ = fi29
            _ = v30
            _ = v31
        }
    }
    if (0.0 <= (((float32((-2.25 * 0.0)) + 0.1) / 1.0) - ((1e10 * 1e10) + 1e10))) {
        var v32 string = v5
        v3 = v3
        _ = v32
    } else {
        v33 := GenS5{F0: -1, F1: (^(v3)), F2: (true || (!false))}
        for fi34 := 0; fi34 < 9; fi34++ {
            var v35 int8 = int8(v4)
            switch v5 {
            case "":
                v35 -= v35
            case "a":
                if (42 != 1) {
                    v10 = 42
                    switch ((v5 + "hello") + "hello") {
                    case "a":
                        v36 := []int{(fi34 | (fi34 * (fi34 ^ -416))), (fi34 ^ int(v3)), (fi34 * (fi34 & 1)), -723}
                        _ = v36
                    default:
                        for fi37 := 0; fi37 < 20; fi37++ {
                            if (0.0 < 3.14159) {
                                fmt.Println(v5, v35, v35)
                                v5 = (v5 + ("" + "foo bar"))
                                v5 = (("\n\t" + "hello") + v5)
                            } else {
                                v28 = (v28 ^ v28)
                                var v38 uint16 = 0
                                _ = v38
                            }
                            var v39 string = v5
                            var v40 uint16 = uint16(v10)
                            v40 = 65535
                            var v41 int8 = (v35 - v35)
                            v3 += int64(v4)
                            _ = len(v1)
                            _ = fi37
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                    }
                }
            }
            _ = fi34
            _ = v35
        }
        var v42 float64 = 0.0
        _ = v33
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
