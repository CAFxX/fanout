package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int16; F2 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 int8; F2 uint64 }
func genFn0(p0 int32, p1 bool, p2 string) int64 {
    p0 = p0
    var v4 float64 = (0.1 + (3.14159 / -2.25))
    switch "foo bar" {
    case "":
        v5 := GenS1{F0: 0, F1: 336}
        _ = v5
    }
    if ((p1 || p1) && true) {
        p0 = 1
        v6 := []int{168, 9223372036854775807, 127, -1}
        _ = v6
    } else {
        var v7 uint16 = 1
        p0 = p0
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 uint32 = 2147483647
            var v10 uint64 = 2
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    }
    v11 := GenS1{F0: -32768, F1: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v4
    return 42
}
func genFn1() int8 {
    v12 := []int{-2147483648, 32768, 703, 32768}
    switch 92 {
    case 2147483647:
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 int8 = 100
            _ = fi13
            _ = v14
        }
    case -2147483648:
        switch (("" + "\n\t") + "") {
        case "\n\t":
            for fi15 := 0; fi15 < 16; fi15++ {
                if false {
                    var v16 float32 = -2.25
                    v17 := GenS1{F0: 313, F1: 2147483648}
                    _ = v16
                    _ = v17
                } else {
                    fmt.Println(fi15)
                    for fi18 := 0; fi18 < 4; fi18++ {
                        v12 = append(v12, fi18)
                        switch -753 {
                        case 9223372036854775807:
                            var v19 int64 = 42
                            _ = v19
                        default:
                            var v20 float64 = 1e10
                            _ = v20
                        }
                        _ = fi18
                    }
                    var v21 float64 = 1e10
                    _ = v21
                }
                _ = fi15
            }
        case "hello":
            var v22 uint = 694
            _ = v22
        case "":
            var v23 int8 = 100
            _ = v23
        }
    }
    var v24 uint = 0
    for fi25 := 0; fi25 < 16; fi25++ {
        var v26 int16 = -1
        var v27 int8 = 127
        var v28 uint32 = 2147483647
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    _ = v12
    _ = v24
    return -91
}

func main() {
    v1 := []string{"a"}
    var v2 float32 = 0.0
    var v3 uint32 = 42
    switch "" {
    case "\n\t":
        var v29 string = "\n\t"
        _ = v29
    case "hello":
        for fi30 := 0; fi30 < 17; fi30++ {
            if true {
                fmt.Println(v3, fi30)
                v2 += v2
                fmt.Println(v2, v2, v2)
            } else {
                for fi31 := 0; fi31 < 5; fi31++ {
                    v32 := GenS2{F0: 127, F1: -729, F2: 42}
                    fmt.Println(fi30, fi31)
                    switch fi30 {
                    case 32767:
                        fmt.Println(v2, fi30)
                    case 470:
                        switch ("" + ((("a" + "foo bar") + "") + ("a" + ""))) {
                        case "":
                            var v33 int8 = 47
                            _ = v33
                        case "a":
                            v34 := GenS3{F0: 2}
                            _ = v34
                        default:
                            var v35 float64 = float64(0.1)
                            v36 := GenS1{F0: 1, F1: (v3 ^ 0)}
                            switch fi30 {
                            case 1:
                                fmt.Println(v2, fi31, fi30)
                            }
                            _ = v35
                            _ = v36
                        }
                    default:
                        v2 = ((v2 * float32(((1.5 * (0.0 / float64(0.1))) * (float64(v2) / -2.25)))) + (v2 / v2))
                    }
                    _ = fi31
                    _ = v32
                }
                fmt.Println(fi30, v3, fi30)
            }
            if (true && false) {
                v2 -= (v2 - (v2 * 0.0))
            } else {
                v1 = append(v1, "foo bar")
                fmt.Println(v2)
            }
            _ = fi30
        }
    case "":
        if false {
            var v37 uint64 = 4294967295
            _ = v37
        } else {
            v2 += v2
        }
    default:
        v2 = v2
    }
    switch -2147483648 {
    case 651:
        var v38 string = "hello"
        _ = v38
    default:
        var v39 uint32 = (v3 / ((1) | 1))
        _ = v39
    }
    fmt.Println(v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v40 int16 = 32767
    var v41 bool = (-1 == 128)
    switch ("foo bar" + ("hello" + "a")) {
    case "":
        v2 -= float32(0.0)
    case "a":
        var v42 uint32 = (v3 - v3)
        _ = v42
    default:
        var v43 uint = 1
        _ = v43
    }
    var v44 bool = v41
    var v45 uint32 = 1
    if (42 != 18446744073709551615) {
        for fi46 := 0; fi46 < 1; fi46++ {
            v47 := GenS4{F0: v45, F1: 43, F2: 965}
            for fi48 := 0; fi48 < 9; fi48++ {
                var v49 bool = true
                _ = fi48
                _ = v49
            }
            _ = fi46
            _ = v47
        }
    }
    fmt.Println(v3, v3)
    fmt.Println(v3, v41)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v44
    _ = v45
}
