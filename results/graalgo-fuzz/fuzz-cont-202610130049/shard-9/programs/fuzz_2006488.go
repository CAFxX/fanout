package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 string; F2 int8 }
type GenS3 struct { F0 int64; F1 string }
func genFn0(p0 uint16, p1 int32) string {
    var v4 uint32 = 1
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 int32 = (-(p1))
        v7 := GenS1{F0: "foo bar"}
        if ((fi5 / ((fi5) | 1)) != -541) {
            fmt.Println(v6)
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 uint64 = 0
                var v10 int16 = 83
                v6 += p1
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    if false {
        v11 := []int{127, 128, -769, 1}
        v11 = append(v11, 741)
        _ = v11
    } else {
        for fi12 := 0; fi12 < 13; fi12++ {
            for fi13 := 0; fi13 < 4; fi13++ {
                v14 := []int{(fi12 ^ (fi12 - fi13)), (fi13 / ((int(v4)) | 1)), (fi13 + int(p1)), (-(fi13))}
                var v15 int = (fi12 & fi13)
                v15 += fi13
                _ = fi13
                _ = v14
                _ = v15
            }
            var v16 int64 = 0
            if false {
                v4 += (v4 * (v4 / ((2147483647) | 1)))
            } else {
                var v17 string = ("hello" + ("" + "a"))
                _ = v17
            }
            _ = fi12
            _ = v16
        }
        var v18 uint16 = 611
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v4
    return ""
}
func genFn1(p0 string) int32 {
    fmt.Println(p0)
    fmt.Println(p0)
    v19 := GenS2{F0: 2, F1: p0, F2: 127}
    for fi20 := 0; fi20 < 10; fi20++ {
        var v21 uint32 = 766
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 uint8 = 133
        _ = fi22
        _ = v23
    }
    _ = p0
    _ = v19
    return 121
}

func main() {
    v1 := []int{57}
    var v2 uint32 = 65535
    var v3 uint16 = 65535
    var v24 string = "foo bar"
    fmt.Println(v3, v2)
    fmt.Println(v3)
    v3 -= (v3 - 65535)
    v24 = "foo bar"
    switch v24 {
    case "a":
        var v25 bool = true
        _ = v25
    case "foo bar":
        v3 -= uint16(v2)
    case "":
        switch 72 {
        case 128:
            for fi26 := 0; fi26 < 11; fi26++ {
                var v27 uint = 18446744073709551615
                switch v24 {
                case "":
                    fmt.Println(v2, v27)
                case "hello":
                    switch fi26 {
                    case 1:
                        if (-9223372036854775808 != 146) {
                            var v28 int = 32768
                            for fi29 := 0; fi29 < 3; fi29++ {
                                if false {
                                    var v30 uint8 = 250
                                    if true {
                                        fmt.Println(v2, fi26)
                                        var v31 int16 = 1000
                                        if len(v1) > 0 {
                                            v1[0] = v28
                                        }
                                        v32 := GenS1{F0: v24}
                                        v28 = 127
                                        switch (v24 + v24) {
                                        case "foo bar":
                                            v1 = append(v1, fi26)
                                        default:
                                            var v33 uint32 = v2
                                            _ = v33
                                        }
                                        v1 = append(v1, (v28 & (v28 ^ 32768)))
                                        _ = v31
                                        _ = v32
                                    }
                                    if false {
                                        if len(v1) > 0 {
                                            v1[0] = fi26
                                        }
                                        v2 += 2
                                        switch v24 {
                                        case "foo bar":
                                            fmt.Println(v3, v3)
                                        case "\n\t":
                                            v34 := GenS3{F0: -662, F1: v24}
                                            _ = v34
                                        case "":
                                            if (true || false) {
                                                fmt.Println(v24, v3, v3)
                                            }
                                        }
                                    }
                                    _ = v30
                                }
                                _ = fi29
                            }
                            _ = v28
                        }
                    case 32767:
                        v35 := GenS1{F0: "\n\t"}
                        _ = v35
                    case 335:
                        v3 -= v3
                    }
                default:
                    v3 = 0
                }
                _ = fi26
                _ = v27
            }
        case 4294967295:
            for fi36 := 0; fi36 < 3; fi36++ {
                v1 = append(v1, 634)
                var v37 uint = 119
                var v38 int = fi36
                _ = fi36
                _ = v37
                _ = v38
            }
        default:
            var v39 int32 = -2147483648
            _ = v39
        }
    }
    v40 := GenS3{F0: 128, F1: v24}
    v3 -= 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v40
}
