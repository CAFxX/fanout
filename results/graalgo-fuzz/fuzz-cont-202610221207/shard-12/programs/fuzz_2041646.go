package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 string; F1 uint32 }
func genFn0(p0 uint8, p1 string) int64 {
    switch p1 {
    case "\n\t":
        v7 := []int{766, 32768}
        _ = v7
    case "hello":
        var v8 uint = 395
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        p0 = (p0 & p0)
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 string = p1
            var v12 float64 = float64(1e10)
            var v13 uint32 = 982
            if ((fi9 ^ 32768) == 0) {
                v12 -= v12
                var v14 uint = 2
                _ = v14
            }
            v15 := GenS1{F0: 973, F1: 65535}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
            _ = v15
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    return 9223372036854775807
}
func genFn1(p0 bool) int32 {
    var v16 bool = false
    switch 9223372036854775807 {
    case 653:
        fmt.Println(p0, v16)
    case 128:
        var v17 float64 = -2.25
        _ = v17
    default:
        for fi18 := 0; fi18 < 20; fi18++ {
            for fi19 := 0; fi19 < 7; fi19++ {
                if true {
                    if ((1e10 >= 0.0) && v16) {
                        if false {
                            p0 = (65535 >= 1)
                            if (!(!(1 < -128))) {
                                p0 = (!(127 > 1))
                            }
                            p0 = (((42 >= 967) || (2 == 294)) && (!v16))
                            v16 = p0
                            fmt.Println(fi18, fi19, fi19)
                        }
                    }
                }
                v20 := []int{(^(fi18)), fi18, fi19}
                var v21 int = 128
                var v22 uint = 974
                var v23 int16 = -32768
                if (1.5 == 0.1) {
                    v22 += v22
                }
                switch (((("" + (("" + "hello") + "hello")) + "\n\t") + "foo bar") + "\n\t") {
                case "hello":
                    fmt.Println(fi18, fi18)
                case "foo bar":
                    if p0 {
                        if (int(v22) < v21) {
                            fmt.Println(v16, v16, p0)
                            if true {
                                fmt.Println(fi19, v21)
                                switch (v21 >> ((688) & 63)) {
                                case 128:
                                    p0 = v16
                                }
                            }
                            v21 = (fi19 * (fi19 + 9223372036854775807))
                            v16 = (!p0)
                        }
                    }
                default:
                    switch (fi19 & v21) {
                    case 1:
                        var v24 int64 = 42
                        _ = v24
                    }
                }
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
                _ = v23
            }
            _ = fi18
        }
    }
    _ = p0
    _ = v16
    return -2147483648
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = float32(1e10)
    v3 := make(map[string]bool)
    v4 := make(map[uint32]bool)
    var v5 uint64 = 803
    var v6 uint64 = v5
    if true {
        var v25 float32 = v2
        _ = v25
    } else {
        v6 = v5
        v2 = v2
        fmt.Println(v6, v2, v5)
    }
    fmt.Println(v5, v6, v6)
    v26 := GenS2{F0: 45, F1: 42}
    for fi27 := 0; fi27 < 14; fi27++ {
        var v28 int16 = -32768
        _ = fi27
        _ = v28
    }
    v5 -= v5
    var v29 uint64 = (v6 & v5)
    switch 4294967295 {
    case 20:
        for fi30 := 0; fi30 < 17; fi30++ {
            var v31 int8 = 100
            v31 += (v31 % (((v31 / ((v31) | 1))) | 1))
            v32 := []int{(fi30 & fi30), fi30}
            _ = fi30
            _ = v31
            _ = v32
        }
    case 127:
        for fi33 := 0; fi33 < 13; fi33++ {
            var v34 uint = 80
            var v35 int32 = 1
            if true {
                v36 := GenS3{F0: "", F1: 4294967295}
                v37 := GenS1{F0: 42, F1: 168}
                _ = v36
                _ = v37
            }
            _ = fi33
            _ = v34
            _ = v35
        }
    case 286:
        var v38 uint = 843
        _ = v38
    default:
        v39 := GenS2{F0: 128, F1: 1}
        _ = v39
    }
    fmt.Println(v2, v5, v6)
    fmt.Println(v5, v29, v6)
    if true {
        fmt.Println(v5, v5)
        v40 := GenS1{F0: 32767, F1: 364}
        _ = v40
    } else {
        var v41 uint = 2
        v42 := GenS3{F0: "\n\t", F1: 282}
        v43 := GenS1{F0: 137, F1: 65535}
        _ = v41
        _ = v42
        _ = v43
    }
    var v44 int = 4294967295
    for fi45 := 0; fi45 < 12; fi45++ {
        fmt.Println(v5)
        _ = fi45
    }
    v44 += (^(v44))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
