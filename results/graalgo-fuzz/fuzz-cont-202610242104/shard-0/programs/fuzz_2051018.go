package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 int64 }
func genFn0(p0 int16, p1 uint8, p2 uint8) uint64 {
    if (-2.25 > 0.0) {
        switch "\n\t" {
        case "\n\t":
            v4 := GenS1{F0: -1}
            _ = v4
        default:
            var v5 uint = 18446744073709551615
            _ = v5
        }
        p2 = p2
    } else {
        if true {
            var v6 float64 = float64(1.5)
            v6 += v6
            v6 += v6
            _ = v6
        }
    }
    v7 := []int{128, 439, 42, 764}
    if (32 != 4294967295) {
        v8 := GenS1{F0: 42}
        _ = v8
    } else {
        fmt.Println(p1)
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 4294967295
}

func main() {
    var v1 int16 = -1
    v2 := make(map[string]string)
    var v3 float64 = -2.25
    var v9 int64 = 9223372036854775807
    v10 := GenS1{F0: 126}
    var v11 float32 = (1.5 + (0.0 / (1e10 * (3.14159 + -2.25))))
    delete(v2, "a")
    v12 := []int{32767}
    v9 = 9223372036854775807
    var v13 float64 = (v3 / v3)
    v9 -= 0
    fmt.Println(v11, v13, v13)
    for fi14 := 0; fi14 < 19; fi14++ {
        if true {
            v15 := GenS2{F0: 955, F1: (v1 ^ 1000)}
            switch "foo bar" {
            case "hello":
                fmt.Println(v1)
            }
            if false {
                _ = len(v12)
            } else {
                fmt.Println(v9, v13, v11)
            }
            _ = v15
        } else {
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 uint32 = 945
                var v18 bool = false
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = fi16
                _ = v17
                _ = v18
            }
            v19 := GenS1{F0: 22}
            _ = v19
        }
        var v20 float32 = (v11 / v11)
        if false {
            delete(v2, "foo bar")
        } else {
            v13 = v13
        }
        _ = fi14
        _ = v20
    }
    var v21 bool = true
    fmt.Println(v11, v13)
    switch "\n\t" {
    case "foo bar":
        var v22 uint32 = 331
        _ = v22
    case "hello":
        for fi23 := 0; fi23 < 6; fi23++ {
            fmt.Println(v3, fi23, v13)
            v24 := GenS2{F0: 65535, F1: v1}
            var v25 uint32 = 901
            _ = fi23
            _ = v24
            _ = v25
        }
    case "a":
        var v26 uint16 = 65535
        _ = v26
    default:
        v27 := GenS1{F0: -44}
        _ = v27
    }
    v13 -= v13
    var v28 string = "hello"
    if len(v12) > 0 {
        v12[0] = 170
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 int = (fi29 % ((fi29) | 1))
        v2[("foo bar" + "\n\t")] = v28
        switch "\n\t" {
        case "foo bar":
            var v31 float32 = 1.5
            _ = v31
        case "hello":
            for fi32 := 0; fi32 < 17; fi32++ {
                if (v21 && v21) {
                    if len(v12) > 0 {
                        v12[0] = (fi29 + v30)
                    }
                    v33 := GenS3{F0: 2}
                    fmt.Println(v13, v13)
                    _ = v33
                } else {
                    v11 = ((v11 * v11) + v11)
                    fmt.Println(fi29)
                    v21 = v21
                    v2[v28] = v28
                    v30 += (^(fi29))
                }
                if (18446744073709551615 != 2) {
                    fmt.Println(v21, fi29, v11)
                }
                if (v21 && v21) {
                    if (-310 > v9) {
                        var v34 uint = 255
                        if true {
                            var v35 uint = 2
                            _ = v35
                        } else {
                            switch (v28 + "hello") {
                            case "\n\t":
                                var v36 bool = (false || v21)
                                _ = v36
                            case "hello":
                                v37 := GenS3{F0: 102}
                                _ = v37
                            case "a":
                                v38 := GenS4{F0: 4294967295, F1: v9}
                                _ = v38
                            }
                        }
                        var v39 uint64 = genFn0(42, 176, 128)
                        v21 = (355 == 1)
                        _ = v34
                        _ = v39
                    }
                    v2[("hello" + v28)] = "foo bar"
                    v28 = ""
                    var v40 float32 = float32(((v3 - ((v13 * (v3 - v13)) * v3)) + (v3 - (v3 / (0.0 / (0.1 * (v3 * v13)))))))
                    var v41 int8 = 52
                    _ = v40
                    _ = v41
                }
                _ = fi32
            }
        case "":
            var v42 uint = 858
            _ = v42
        }
        _ = fi29
        _ = v30
    }
    switch -9223372036854775808 {
    case 128:
        var v43 int32 = 944
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v9
}
