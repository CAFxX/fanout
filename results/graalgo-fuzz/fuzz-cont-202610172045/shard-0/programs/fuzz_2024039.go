package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint16; F1 uint; F2 int64 }
type GenS4 struct { F0 int16; F1 string; F2 int16 }
type GenS5 struct { F0 int32; F1 int64 }
func genFn0(p0 uint) uint {
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 uint64 = 18446744073709551615
        for fi7 := 0; fi7 < 15; fi7++ {
            v6 -= (v6 >> (((p0 & 79)) & 63))
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    if (!(((2 != 42) && (42 >= 978)) || ("hello" != ("foo bar" + "foo bar")))) {
        switch 32768 {
        case 32768:
            v8 := GenS1{F0: 65535, F1: 346}
            _ = v8
        case 9223372036854775807:
            p0 = 42
        default:
            p0 += (p0 - p0)
        }
        switch 128 {
        case -9223372036854775808:
            var v9 int32 = 2147483647
            _ = v9
        default:
            v10 := GenS1{F0: (p0 / (((^(p0))) | 1)), F1: 1}
            _ = v10
        }
        switch -602 {
        case -9223372036854775808:
            if ((!(false && ("" >= "\n\t"))) && true) {
                var v11 uint = 1
                v11 -= (^(p0))
                _ = v11
            } else {
                v12 := GenS2{F0: 42, F1: (^(p0))}
                _ = v12
            }
        default:
            p0 += (p0 + p0)
        }
    }
    if (true || true) {
        v13 := []int{32767, 1}
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            p0 = 65535
            switch "" {
            case "":
                var v16 bool = (-32768 > 1000)
                _ = v16
            }
            _ = fi15
        }
        var v17 int32 = 0
        _ = fi14
        _ = v17
    }
    p0 += 255
    _ = p0
    return p0
}
func genFn1(p0 uint32, p1 uint, p2 uint16) int32 {
    for fi18 := 0; fi18 < 4; fi18++ {
        v19 := GenS1{F0: 236, F1: -2147483648}
        v20 := GenS2{F0: uint32(p1), F1: p1}
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        for fi22 := 0; fi22 < 6; fi22++ {
            v23 := GenS3{F0: (p2 ^ 0), F1: 4294967295, F2: 1}
            var v24 float32 = ((float32((((0.1 / 1e10) + (3.14159 + float64(3.14159))) + 0.1)) + (float32(3.14159) / 1.5)) + float32(((1.5 / (0.0 + (-2.25 - 0.0))) * 0.1)))
            p0 += p0
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi21
    }
    var v25 uint8 = 43
    v26 := []int{-1}
    var v27 int8 = 30
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v26
    _ = v27
    return -1
}

func main() {
    v1 := []string{"", ""}
    var v2 int64 = 0
    var v3 int16 = 1000
    var v4 float64 = (0.0 * 0.1)
    var v28 uint16 = 798
    v4 = (v4 * float64(-2.25))
    v29 := GenS3{F0: v28, F1: 2, F2: (v2 | -9223372036854775808)}
    v30 := GenS4{F0: (v3 - -32768), F1: "a", F2: -491}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        if true {
            var v31 int16 = v3
            switch "hello" {
            case "hello":
                var v32 int8 = -128
                _ = v32
            case "":
                fmt.Println(v4)
            case "foo bar":
                if ((127 >= -1) || (!false)) {
                    fmt.Println(v4)
                    if (false || false) {
                        for fi33 := 0; fi33 < 18; fi33++ {
                            for fi34 := 0; fi34 < 3; fi34++ {
                                fmt.Println(fi33, fi34)
                                v3 += -1
                                _ = fi34
                            }
                            for fi35 := 0; fi35 < 11; fi35++ {
                                if (!false) {
                                    fmt.Println(v3, v31)
                                } else {
                                    var v36 float32 = 0.1
                                    _ = len(v1)
                                    _ = v36
                                }
                                _ = fi35
                            }
                            if ((!false) && (-32768 <= (v3 >> ((95) & 63)))) {
                                v3 = 32767
                                for fi37 := 0; fi37 < 7; fi37++ {
                                    v4 += 1e10
                                    _ = fi37
                                }
                                if false {
                                    var v38 float64 = 0.0
                                    var v39 uint16 = v28
                                    _ = v38
                                    _ = v39
                                }
                            }
                            _ = fi33
                        }
                    } else {
                        v40 := GenS5{F0: 2147483647, F1: (v2 << ((genFn0(255)) & 63))}
                        _ = v40
                    }
                    v2 = 9223372036854775807
                }
            }
            _ = v31
        }
        switch -2147483648 {
        case -695:
            v2 -= 976
        case 0:
            var v41 bool = true
            _ = v41
        default:
            var v42 uint8 = 1
            _ = v42
        }
        var v43 uint = 4294967295
        _ = v43
    }
    if true {
        var v44 string = ("" + ("a" + ("foo bar" + ("hello" + ("hello" + "a")))))
        _ = v44
    }
    var v45 uint16 = 32767
    var v46 uint64 = 411
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v45
    _ = v46
}
