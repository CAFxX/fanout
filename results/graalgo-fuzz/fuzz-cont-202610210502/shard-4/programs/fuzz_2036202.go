package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 uint16; F2 uint64 }
type GenS5 struct { F0 string; F1 int8; F2 uint64 }

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -2147483648
    var v3 uint32 = 4294967295
    var v4 float64 = 0.1
    var v5 float32 = ((0.0 / 1.0) / 0.1)
    var v6 float32 = v5
    fmt.Println(v6, v5)
    v7 := GenS1{F0: ""}
    var v8 bool = (!true)
    var v9 uint8 = 117
    var v10 int32 = 2147483647
    v11 := []int{747, -280, 779}
    if true {
        for fi12 := 0; fi12 < 11; fi12++ {
            _ = len(v1)
            v4 -= v4
            _ = fi12
        }
        if false {
            var v13 float32 = v5
            for fi14 := 0; fi14 < 2; fi14++ {
                if v8 {
                    switch "hello" {
                    case "a":
                        switch (fi14 + fi14) {
                        case 42:
                            v4 += (1e10 + (v4 + ((v4 / v4) * 1.5)))
                        case -24:
                            for fi15 := 0; fi15 < 3; fi15++ {
                                v8 = (42 <= int64(v2))
                                var v16 int32 = v10
                                if false {
                                    v4 += v4
                                    v17 := GenS1{F0: "\n\t"}
                                    v9 += 124
                                    _ = v17
                                }
                                v6 = 0.1
                                var v18 int64 = -9223372036854775808
                                v19 := GenS1{F0: "foo bar"}
                                if ((^(v3)) <= v3) {
                                    v20 := GenS2{F0: 42}
                                    v3 += v3
                                    v10 = v16
                                    fmt.Println(v4, v9, fi14)
                                    if true {
                                        switch "hello" {
                                        case "foo bar":
                                            var v21 uint64 = 2
                                            _ = v21
                                        default:
                                            v22 := GenS3{F0: v16}
                                            _ = v22
                                        }
                                    } else {
                                        var v23 uint8 = (v9 >> ((856) & 63))
                                        switch (fi14 * fi14) {
                                        case 1:
                                            v24 := GenS4{F0: "a", F1: 704, F2: 0}
                                            _ = v24
                                        case -1:
                                            v25 := v1["a"]
                                            _ = v25
                                        case -9223372036854775808:
                                            v16 = v16
                                        default:
                                            fmt.Println(v16, v8)
                                        }
                                        _ = v23
                                    }
                                    _ = v20
                                }
                                fmt.Println(v16, v16)
                                v2 = (v10 - v10)
                                _ = fi15
                                _ = v16
                                _ = v18
                                _ = v19
                            }
                        default:
                            var v26 uint16 = 255
                            _ = v26
                        }
                    case "\n\t":
                        var v27 int64 = -858
                        _ = v27
                    case "":
                        if (!true) {
                            var v28 uint8 = (v9 ^ 42)
                            v8 = (!(-2.25 > v13))
                            fmt.Println(v8, v10, v10)
                            _ = v28
                        }
                    default:
                        v29 := GenS2{F0: v3}
                        _ = v29
                    }
                } else {
                    var v30 uint32 = v3
                    if v8 {
                        v31 := GenS2{F0: v30}
                        fmt.Println(v8, v9, fi14)
                        for fi32 := 0; fi32 < 3; fi32++ {
                            if v8 {
                                v9 = 1
                                v33 := GenS3{F0: v2}
                                v1[("\n\t" + "\n\t")] = (4294967295 < 0)
                                _ = v33
                            }
                            v34 := v1[(("foo bar" + (("hello" + "") + ("foo bar" + "\n\t"))) + ("foo bar" + ("" + "foo bar")))]
                            v35 := v1[""]
                            _ = fi32
                            _ = v34
                            _ = v35
                        }
                        _ = v31
                    }
                    switch (fi14 % ((2147483647) | 1)) {
                    case -1:
                        var v36 int16 = 42
                        _ = v36
                    default:
                        v13 -= (v6 - v5)
                    }
                    _ = v30
                }
                _ = fi14
            }
            v37 := GenS4{F0: "hello", F1: 65535, F2: 421}
            _ = v13
            _ = v37
        }
        v4 += 0.1
    } else {
        var v38 uint16 = 0
        _ = v38
    }
    var v39 int8 = -114
    if (-32768 <= 401) {
        v40 := v1["hello"]
        _ = v40
    } else {
        v41 := GenS2{F0: (v3 ^ v3)}
        v9 -= v9
        _ = v41
    }
    if ((32767 >= 0) || v8) {
        if len(v11) > 0 {
            v11[0] = 2147483647
        }
    }
    for fi42 := 0; fi42 < 20; fi42++ {
        v43 := GenS5{F0: ("" + "\n\t"), F1: v39, F2: 812}
        if ((!false) && true) {
            v44 := v1["foo bar"]
            var v45 uint8 = (v9 % (((v9 & v9)) | 1))
            _ = v44
            _ = v45
        }
        _ = fi42
        _ = v43
    }
    v46 := GenS3{F0: (v2 ^ (v2 + (v10 + v2)))}
    var v47 float32 = v6
    if len(v11) > 0 {
        v11[0] = -9223372036854775808
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
