package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 string; F1 int32; F2 int8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint16, p1 string) int8 {
    v6 := GenS1{F0: p0, F1: p0}
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := []int{-1, (fi7 / ((32767) | 1)), fi7, fi7}
        var v9 int64 = 42
        switch fi7 {
        case -1:
            p1 = p1
        case -2147483648:
            for fi10 := 0; fi10 < 4; fi10++ {
                switch (fi10 & -9223372036854775808) {
                case -2147483648:
                    v11 := GenS1{F0: (p0 / ((p0) | 1)), F1: (p0 | (p0 + (p0 << ((1) & 63))))}
                    _ = v11
                default:
                    var v12 int8 = -1
                    _ = v12
                }
                fmt.Println(p1)
                if true {
                    fmt.Println(fi7, fi10, fi10)
                    v13 := GenS2{F0: ("foo bar" + p1), F1: 42, F2: 0}
                    if len(v8) > 0 {
                        v8[0] = fi7
                    }
                    _ = v13
                }
                _ = fi10
            }
        case 128:
            var v14 int64 = 459
            _ = v14
        default:
            switch 582 {
            case 929:
                if (!false) {
                    var v15 int32 = 909
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 16; fi16++ {
                        switch ("" + "foo bar") {
                        case "":
                            var v17 string = "foo bar"
                            _ = v17
                        case "hello":
                            if len(v8) > 0 {
                                v8[0] = fi7
                            }
                            p1 = p1
                        default:
                            p0 = (^(p0))
                        }
                        var v18 uint16 = (p0 + 42)
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        _ = fi16
                        _ = v18
                    }
                    v9 -= v9
                    fmt.Println(v9, p1)
                }
            case 127:
                p0 = 2
            case -9223372036854775808:
                p0 -= 42
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return 127
}
func genFn1(p0 uint16, p1 int16) uint32 {
    v19 := []int{}
    if true {
        p1 = p1
        var v20 uint16 = 309
        v19 = append(v19, -2147483648)
        _ = v20
    }
    if true {
        var v21 int16 = p1
        var v22 int8 = -47
        _ = v21
        _ = v22
    }
    if false {
        p1 -= p1
    }
    _ = p0
    _ = p1
    _ = v19
    return 2147483647
}

func main() {
    var v1 float32 = 1.5
    v2 := []uint32{179, 990}
    v3 := []int{}
    v4 := make(map[int]uint32)
    v5 := []uint32{}
    v1 -= 1.5
    v1 -= v1
    v23 := GenS1{F0: 1, F1: 65535}
    var v24 bool = ("" < ("" + (("\n\t" + "a") + "\n\t")))
    v1 = v1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v25 int = 961
    var v26 int8 = 100
    if false {
        for fi27 := 0; fi27 < 2; fi27++ {
            var v28 float64 = (1.5 / 1.5)
            if (4294967295 <= 1) {
                var v29 uint16 = 65535
                fmt.Println(v29, v1)
                var v30 int8 = (v26 & genFn0(v29, "\n\t"))
                _ = v29
                _ = v30
            }
            _ = fi27
            _ = v28
        }
        if (((680 == 354) && (!v24)) || v24) {
            v31 := GenS1{F0: 65535, F1: 2}
            _ = v31
        }
    }
    var v32 uint64 = 42
    v24 = v24
    v33 := GenS3{F0: 42}
    v1 = 1e10
    var v34 bool = v24
    if (("a" + "foo bar") >= "a") {
        if len(v5) > 0 {
            v5[0] = genFn1(65535, 482)
        }
    } else {
        var v35 bool = false
        fmt.Println(v34)
        switch ("a" + (("a" + "\n\t") + "")) {
        case "a":
            _ = len(v3)
        }
        _ = v35
    }
    for fi36 := 0; fi36 < 4; fi36++ {
        fmt.Println(v34, v24)
        var v37 int64 = 1
        _ = fi36
        _ = v37
    }
    v4[(^(v25))] = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
