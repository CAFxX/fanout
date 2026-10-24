package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint32, p1 uint8, p2 uint32) int {
    if false {
        var v7 int8 = -128
        switch 221 {
        case 2147483647:
            for fi8 := 0; fi8 < 5; fi8++ {
                switch "" {
                case "":
                    for fi9 := 0; fi9 < 9; fi9++ {
                        switch "foo bar" {
                        case "\n\t":
                            v10 := []int{fi8, fi9}
                            _ = v10
                        case "a":
                            fmt.Println(p2)
                        case "hello":
                            if false {
                                fmt.Println(p1)
                            }
                        }
                        if (fi9 < fi9) {
                            fmt.Println(p0)
                            var v11 int8 = (v7 ^ int8(p1))
                            var v12 float32 = (float32(float64(float32(((3.14159 * (1.5 + 0.0)) * 0.0)))) + 3.14159)
                            var v13 uint = 2
                            if false {
                                if (98 != 873) {
                                    var v14 uint16 = 1
                                    var v15 uint64 = 0
                                    _ = v14
                                    _ = v15
                                }
                                p1 = 127
                            }
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                        if (1 >= 42) {
                            switch fi8 {
                            case -1:
                                fmt.Println(fi8, p0, p1)
                            case 9223372036854775807:
                                var v16 uint32 = 981
                                _ = v16
                            }
                            v17 := GenS1{F0: 2, F1: 81, F2: 1}
                            v18 := []int{128}
                            if (3.14159 >= 1.5) {
                                var v19 int = 1
                                var v20 uint8 = (p1 & (p1 >> ((0) & 63)))
                                _ = v19
                                _ = v20
                            }
                            _ = v17
                            _ = v18
                        } else {
                            p0 += p0
                            p1 += 0
                        }
                        _ = fi9
                    }
                case "a":
                    v21 := []int{128, -9223372036854775808, fi8, fi8}
                    _ = v21
                case "hello":
                    for fi22 := 0; fi22 < 15; fi22++ {
                        fmt.Println(fi22, p2, p2)
                        _ = fi22
                    }
                }
                _ = fi8
            }
        case 0:
            v23 := []int{32767, -2147483648, -9223372036854775808, 4294967295}
            _ = v23
        case 32767:
            if ((p0 + p2) == p0) {
                switch ("hello" + ("foo bar" + ("" + "a"))) {
                case "\n\t":
                    for fi24 := 0; fi24 < 7; fi24++ {
                        v25 := GenS1{F0: uint16(p2), F1: 141, F2: 616}
                        _ = fi24
                        _ = v25
                    }
                }
                p2 = (p0 / ((0) | 1))
            }
        }
        _ = v7
    }
    var v26 int32 = -665
    v26 += v26
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return -31
}
func genFn1(p0 uint32, p1 uint8) uint8 {
    var v27 float32 = -2.25
    v27 = v27
    p1 = p1
    for fi28 := 0; fi28 < 2; fi28++ {
        v29 := []int{(fi28 | genFn0(0, p1, p0)), genFn0(p0, (p1 & p1), p0), (^(fi28))}
        var v30 int16 = 826
        var v31 int64 = -1
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    _ = p0
    _ = p1
    _ = v27
    return 2
}

func main() {
    v1 := make(map[string]int)
    var v2 string = "foo bar"
    var v3 string = v2
    var v4 bool = true
    var v5 uint = 122
    var v6 uint = 932
    v2 = (v3 + v2)
    v32 := []int{4294967295, 1, 0}
    var v33 int32 = 42
    v34 := GenS1{F0: 0, F1: 878, F2: 2}
    v35 := GenS1{F0: 2, F1: 2, F2: 216}
    v36 := GenS1{F0: 2, F1: 1, F2: 0}
    for fi37 := 0; fi37 < 13; fi37++ {
        if v4 {
            v38 := GenS1{F0: 1, F1: 863, F2: 32767}
            v39 := GenS1{F0: 718, F1: 911, F2: 32767}
            _ = v38
            _ = v39
        }
        _ = fi37
    }
    delete(v1, v3)
    if (-1 <= -9223372036854775808) {
        var v40 string = v3
        v33 = (v33 + (v33 << (((v6 % ((v5) | 1))) & 63)))
        _ = v40
    } else {
        if len(v32) >= 2 {
            v32 = v32[1:len(v32)]
        }
        v33 = 2147483647
    }
    var v41 int16 = -32768
    v4 = v4
    v3 = v2
    v42 := GenS2{F0: 42}
    v41 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
