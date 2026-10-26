package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = (0.1 * 0.1)
    var v3 uint32 = 2147483648
    v4 := []int{42, -875, -738, 2147483647}
    for fi5 := 0; fi5 < 6; fi5++ {
        fmt.Println(fi5, fi5, v3)
        v6 := GenS1{F0: (("" + "hello") + "foo bar")}
        var v7 uint = 499
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        if (((v3 % ((v3) | 1)) >= (v3 >> ((42) & 63))) && true) {
            v9 := GenS1{F0: (("foo bar" + "\n\t") + "foo bar")}
            v3 += (v3 * v3)
            _ = v9
        }
        _ = fi8
    }
    if (false || (1000 != -32768)) {
        delete(v1, 519)
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := GenS1{F0: "hello"}
            if true {
                for fi12 := 0; fi12 < 7; fi12++ {
                    v3 = v3
                    fmt.Println(fi10, fi10)
                    if false {
                        v1[128] = (("foo bar" + "\n\t") + "foo bar")
                        var v13 uint16 = 1
                        _ = v13
                    }
                    var v14 float32 = v2
                    var v15 float32 = (v14 * ((((0.0 - v14) * float32(-2.25)) - 0.1) + v2))
                    v2 = (v14 * float32(3.14159))
                    _ = fi12
                    _ = v14
                    _ = v15
                }
                if ((-(fi10)) >= fi10) {
                    fmt.Println(fi10)
                    fmt.Println(v3, fi10, v3)
                } else {
                    switch -433 {
                    case -842:
                        v16 := GenS1{F0: ("" + "")}
                        _ = v16
                    default:
                        v3 -= v3
                    }
                }
            }
            _ = fi10
            _ = v11
        }
        var v17 int = -1
        _ = v17
    }
    v18 := GenS1{F0: ""}
    if (!true) {
        for fi19 := 0; fi19 < 15; fi19++ {
            for fi20 := 0; fi20 < 9; fi20++ {
                v3 += 0
                v4 = append(v4, fi20)
                _ = fi20
            }
            var v21 int64 = -1
            _ = len(v4)
            _ = fi19
            _ = v21
        }
    } else {
        for fi22 := 0; fi22 < 11; fi22++ {
            v3 = v3
            var v23 float64 = ((0.1 - ((0.1 + float64(float32(1e10))) * float64(v2))) + 0.0)
            _ = fi22
            _ = v23
        }
        v3 = v3
        var v24 int8 = 0
        _ = v24
    }
    switch 42 {
    case 32768:
        v25 := GenS1{F0: "foo bar"}
        _ = v25
    }
    switch (("hello" + "foo bar") + "\n\t") {
    case "a":
        v26 := GenS2{F0: 0, F1: 42}
        _ = v26
    }
    v1[315] = "a"
    var v27 int64 = 42
    var v28 uint32 = v3
    switch "a" {
    case "":
        if true {
            for fi29 := 0; fi29 < 9; fi29++ {
                var v30 int64 = (v27 / ((v27) | 1))
                var v31 uint16 = 2
                v3 = v3
                _ = fi29
                _ = v30
                _ = v31
            }
            switch 9223372036854775807 {
            case 4294967295:
                if false {
                    fmt.Println(v27, v27)
                }
            default:
                for fi32 := 0; fi32 < 15; fi32++ {
                    var v33 int16 = -1
                    _ = fi32
                    _ = v33
                }
            }
        } else {
            for fi34 := 0; fi34 < 18; fi34++ {
                if true {
                    var v35 bool = ((4294967295 < 1) || true)
                    var v36 uint = 408
                    fmt.Println(v3)
                    _ = v35
                    _ = v36
                } else {
                    if true {
                        fmt.Println(fi34, fi34, v28)
                        var v37 uint64 = 390
                        delete(v1, (fi34 * 127))
                        _ = v37
                    } else {
                        fmt.Println(v27, fi34)
                        if (2147483647 <= 42) {
                            v38 := GenS2{F0: -1, F1: 668}
                            v3 = v28
                            _ = v38
                        } else {
                            v39 := GenS1{F0: "a"}
                            v1[fi34] = "foo bar"
                            var v40 uint64 = 775
                            _ = v39
                            _ = v40
                        }
                    }
                    var v41 int64 = v27
                    _ = v41
                }
                v42 := v1[128]
                v43 := GenS3{F0: ("" + (v42 + (v42 + "\n\t")))}
                _ = fi34
                _ = v42
                _ = v43
            }
            var v44 uint = 4294967295
            switch 2147483647 {
            case 9223372036854775807:
                v45 := GenS2{F0: 0, F1: 468}
                _ = v45
            default:
                v28 += v3
            }
            _ = v44
        }
    case "foo bar":
        var v46 uint64 = 2
        _ = v46
    case "a":
        v2 -= (v2 * 1.5)
    default:
        _ = len(v4)
    }
    for fi47 := 0; fi47 < 9; fi47++ {
        var v48 bool = ((v28 & 106) != (v3 & (v28 & v3)))
        _ = fi47
        _ = v48
    }
    v28 -= 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
