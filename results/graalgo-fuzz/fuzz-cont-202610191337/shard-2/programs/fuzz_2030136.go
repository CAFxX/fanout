package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 string; F2 int32 }
func genFn0(p0 uint32, p1 uint32, p2 int16) uint {
    var v5 uint64 = 4294967295
    v5 = v5
    var v6 int = 869
    if false {
        v6 += (v6 & (v6 ^ (v6 >> ((924) & 63))))
        v7 := GenS1{F0: 42, F1: true, F2: 2}
        if (-956 < 2147483647) {
            var v8 int8 = 33
            _ = v8
        }
        _ = v7
    }
    fmt.Println(v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 673
}
func genFn1(p0 int16, p1 uint, p2 int16) int8 {
    var v9 int16 = (p0 & p2)
    v10 := GenS1{F0: -742, F1: (1 < 1), F2: p1}
    p0 += 32767
    if false {
        p0 = (^(p2))
        v11 := []int{0, -697, 4294967295}
        v12 := GenS1{F0: -2147483648, F1: true, F2: p1}
        _ = v11
        _ = v12
    }
    p0 += 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 1
}

func main() {
    v1 := []int{32768}
    v2 := []uint32{2147483647, 320}
    v3 := []string{"", ("hello" + "foo bar")}
    var v4 bool = true
    var v13 float32 = 0.1
    var v14 float64 = 1e10
    var v15 float32 = (v13 * (1.5 - (v13 / v13)))
    fmt.Println(v14, v4, v4)
    v14 -= v14
    fmt.Println(v14)
    for fi16 := 0; fi16 < 18; fi16++ {
        if true {
            v17 := GenS1{F0: 42, F1: v4, F2: genFn0(230, 686, 1000)}
            var v18 int8 = 100
            _ = v17
            _ = v18
        }
        v14 = v14
        _ = fi16
    }
    v13 = ((v13 * ((-2.25 + -2.25) / v15)) + (v15 / float32((v14 + v14))))
    var v19 int64 = 9223372036854775807
    switch "hello" {
    case "":
        var v20 int = 1
        _ = v20
    case "hello":
        switch "\n\t" {
        case "foo bar":
            if v4 {
                for fi21 := 0; fi21 < 18; fi21++ {
                    if v4 {
                        if len(v1) > 0 {
                            v1[0] = (^(fi21))
                        }
                        if true {
                            for fi22 := 0; fi22 < 7; fi22++ {
                                var v23 string = (("" + "hello") + "hello")
                                v19 = v19
                                _ = fi22
                                _ = v23
                            }
                            var v24 float32 = v13
                            v25 := GenS1{F0: 2147483647, F1: (v15 == v13), F2: 1}
                            _ = v24
                            _ = v25
                        } else {
                            fmt.Println(v14, v14)
                        }
                        fmt.Println(v19, v4, v13)
                    }
                    for fi26 := 0; fi26 < 15; fi26++ {
                        v27 := GenS2{F0: 65535}
                        if ((fi26 + 128) <= (^(fi26))) {
                            var v28 int64 = (v19 % ((v19) | 1))
                            v29 := GenS3{F0: -32768}
                            fmt.Println(v19, v14)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch 843 {
                            case 545:
                                v1 = append(v1, fi26)
                            case -1:
                                v28 += -73
                            case 162:
                                var v30 uint64 = 2
                                _ = v30
                            }
                            switch (fi21 & 128) {
                            case 32768:
                                v31 := GenS4{F0: -1, F1: "foo bar", F2: 2147483647}
                                _ = v31
                            case -773:
                                fmt.Println(v13)
                            case 127:
                                v13 += (1e10 * float32(v14))
                            }
                            _ = v28
                            _ = v29
                        }
                        var v32 int16 = 42
                        _ = fi26
                        _ = v27
                        _ = v32
                    }
                    if v4 {
                        var v33 int = fi21
                        var v34 float64 = 1e10
                        var v35 int64 = v19
                        _ = v33
                        _ = v34
                        _ = v35
                    } else {
                        var v36 string = "\n\t"
                        _ = v36
                    }
                    _ = fi21
                }
                for fi37 := 0; fi37 < 17; fi37++ {
                    v15 += (float32((float64(float32(1e10)) / v14)) + ((v13 - 0.0) * v15))
                    _ = fi37
                }
            } else {
                for fi38 := 0; fi38 < 19; fi38++ {
                    var v39 float32 = (v13 * v13)
                    switch ("\n\t" + "a") {
                    case "":
                        v4 = (!(true || v4))
                    }
                    _ = fi38
                    _ = v39
                }
                if (false || (v4 || false)) {
                    var v40 int = -38
                    _ = v40
                }
            }
        case "":
            var v41 int8 = 127
            _ = v41
        case "hello":
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
        default:
            if v4 {
                var v42 uint = 2
                _ = v42
            } else {
                v43 := GenS2{F0: 462}
                v15 = 1e10
                _ = v43
            }
        }
    case "foo bar":
        if (!(1.5 >= v14)) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
