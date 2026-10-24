package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 string; F2 uint64 }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 int32, p1 uint) int64 {
    v7 := []int{655, -2147483648}
    if true {
        for fi8 := 0; fi8 < 1; fi8++ {
            for fi9 := 0; fi9 < 10; fi9++ {
                switch (fi8 % (((fi8 - fi8)) | 1)) {
                case -9223372036854775808:
                    var v10 string = "hello"
                    _ = v10
                case 128:
                    p1 -= p1
                default:
                    p0 = p0
                }
                var v11 int64 = 9223372036854775807
                var v12 uint8 = 0
                _ = fi9
                _ = v11
                _ = v12
            }
            _ = fi8
        }
    } else {
        if true {
            p0 += p0
            v13 := GenS1{F0: p0}
            _ = v13
        } else {
            v14 := GenS2{F0: (p0 % (((p0 / ((-1) | 1))) | 1))}
            _ = v14
        }
        v15 := GenS1{F0: (p0 & p0)}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v7
    return -9223372036854775808
}
func genFn1() uint32 {
    v16 := []int{1, -529, 858, 1}
    if (0 >= 65535) {
        var v17 float64 = float64((((-2.25 + 3.14159) / (float32(0.1) - -2.25)) + 1e10))
        _ = v17
    } else {
        v16 = append(v16, 341)
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    }
    if len(v16) > 0 {
        v16[0] = 42
    }
    _ = v16
    return 2147483648
}

func main() {
    var v1 float64 = (float64(1.5) / 1e10)
    var v2 int64 = 0
    v3 := make(map[uint32]bool)
    v4 := make(map[string]uint32)
    var v5 int16 = 1
    var v6 int16 = (v5 ^ 500)
    if false {
        for fi18 := 0; fi18 < 19; fi18++ {
            delete(v4, ("" + "a"))
            _ = fi18
        }
        v6 = 612
        if (!(-2.25 != (0.1 - (0.1 / float32(((0.1 + 0.0) + v1)))))) {
            v19 := GenS3{F0: 127}
            for fi20 := 0; fi20 < 14; fi20++ {
                if true {
                    if true {
                        if false {
                            fmt.Println(fi20, v2)
                        } else {
                            v2 = genFn0(0, 897)
                            fmt.Println(fi20, fi20)
                        }
                        if false {
                            if ("foo bar" <= "\n\t") {
                                v21 := GenS4{F0: 214}
                                v2 = 136
                                _ = v21
                            }
                        }
                        if (!false) {
                            for fi22 := 0; fi22 < 11; fi22++ {
                                v23 := GenS2{F0: 1}
                                _ = fi22
                                _ = v23
                            }
                        } else {
                            var v24 uint = 2
                            var v25 int16 = v5
                            v26 := GenS5{F0: 2, F1: (("\n\t" + "") + ("" + "\n\t")), F2: 0}
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    } else {
                        switch "" {
                        case "":
                            var v27 float32 = (1.5 * (1.5 / 0.1))
                            _ = v27
                        case "foo bar":
                            v5 += v6
                        }
                        v6 -= v5
                    }
                    if (false || true) {
                        v5 = (v5 - v6)
                        var v28 string = (("" + "\n\t") + "a")
                        _ = v28
                    }
                }
                var v29 uint16 = 258
                v30 := GenS6{F0: "a", F1: 283}
                _ = fi20
                _ = v29
                _ = v30
            }
            _ = v19
        }
    }
    v6 += 177
    var v31 int16 = (-(v6))
    var v32 float32 = 3.14159
    v33 := v4[("" + "")]
    fmt.Println(v31, v33)
    var v34 int64 = (v2 | v2)
    if false {
        var v35 string = ("a" + "foo bar")
        _ = v35
    }
    var v36 int8 = 100
    v37 := GenS1{F0: 0}
    for fi38 := 0; fi38 < 9; fi38++ {
        v31 = 32767
        v39 := GenS7{F0: 0}
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
