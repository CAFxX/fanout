package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint32 }
func genFn0(p0 uint16) int64 {
    switch "foo bar" {
    case "foo bar":
        if ((p0 / (((p0 * (p0 & 32767))) | 1)) != 2) {
            v5 := []int{}
            for fi6 := 0; fi6 < 9; fi6++ {
                switch "\n\t" {
                case "hello":
                    var v7 int8 = -1
                    _ = v7
                }
                _ = fi6
            }
            var v8 uint64 = 4294967295
            _ = v5
            _ = v8
        } else {
            var v9 uint16 = 2
            v9 += (v9 & p0)
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 int8 = -57
                _ = fi10
                _ = v11
            }
            _ = v9
        }
    case "\n\t":
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 uint32 = 4294967295
            v14 := []int{-1, (fi12 >> ((559) & 63))}
            _ = fi12
            _ = v13
            _ = v14
        }
    default:
        fmt.Println(p0)
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        p0 -= p0
        v16 := []int{(fi15 / ((fi15) | 1)), 32768, fi15}
        if len(v16) > 0 {
            v16[0] = fi15
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(p0)
    _ = p0
    return 1
}
func genFn1() int32 {
    if (3.14159 >= (-2.25 / 1.0)) {
        var v17 bool = (1000 == -798)
        v18 := []int{-51, 42}
        _ = v17
        _ = v18
    }
    var v19 int8 = -128
    switch -414 {
    case 128:
        if false {
            fmt.Println(v19)
            var v20 bool = false
            _ = v20
        }
    default:
        for fi21 := 0; fi21 < 5; fi21++ {
            v22 := []int{115, (^(fi21)), fi21, (^(fi21))}
            _ = fi21
            _ = v22
        }
    }
    _ = v19
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = -2147483648
    var v3 float64 = (3.14159 + 1.5)
    var v4 uint32 = 2
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 float64 = v3
        for fi25 := 0; fi25 < 4; fi25++ {
            v24 = float64(3.14159)
            if true {
                if (1 > 255) {
                    v3 -= (v24 - (((-2.25 + v3) / v24) + ((0.1 - v24) / v24)))
                    if false {
                        fmt.Println(v24, v24, v3)
                        if ((0.0 + 0.1) != 3.14159) {
                            v26 := GenS1{F0: 60, F1: 42}
                            var v27 float32 = 3.14159
                            v4 -= v4
                            _ = v26
                            _ = v27
                        } else {
                            v24 = (1e10 + v24)
                        }
                    }
                    fmt.Println(v3)
                    switch "foo bar" {
                    case "":
                        v28 := []int{(fi23 << ((1) & 63)), (^(fi25))}
                        _ = v28
                    }
                    if true {
                        switch fi25 {
                        case -781:
                            var v29 uint = 842
                            _ = v29
                        }
                    } else {
                        v30 := GenS2{F0: 2, F1: true}
                        v4 = v4
                        v4 = 447
                        _ = v30
                    }
                } else {
                    v3 = v24
                    var v31 bool = false
                    _ = v31
                }
                switch (fi25 % (((fi23 ^ int(v2))) | 1)) {
                case 334:
                    var v32 int = 32767
                    _ = v32
                }
            } else {
                fmt.Println(v24, v24)
                if false {
                    var v33 int32 = v2
                    v34 := GenS1{F0: -1, F1: 48}
                    delete(v1, "")
                    v35 := GenS2{F0: 4294967295, F1: (-128 != 0)}
                    v36 := GenS3{F0: -32768}
                    _ = v33
                    _ = v34
                    _ = v35
                    _ = v36
                }
            }
            _ = fi25
        }
        _ = fi23
        _ = v24
    }
    fmt.Println(v3, v2)
    v4 += (v4 * (v4 & 542))
    switch 32767 {
    case -2147483648:
        var v37 int32 = (^(v2))
        _ = v37
    case 32768:
        var v38 uint = uint(v4)
        _ = v38
    }
    switch ("" + ("" + "hello")) {
    case "hello":
        if true {
            v3 = (v3 - 1e10)
        }
    case "a":
        for fi39 := 0; fi39 < 16; fi39++ {
            v2 = v2
            _ = fi39
        }
    }
    for fi40 := 0; fi40 < 19; fi40++ {
        var v41 int8 = 20
        _ = fi40
        _ = v41
    }
    var v42 int = 128
    if true {
        var v43 bool = (v42 == (v42 / ((v42) | 1)))
        delete(v1, "hello")
        v44 := GenS4{F0: v4, F1: 9223372036854775807, F2: v4}
        _ = v43
        _ = v44
    }
    switch (v42 * 963) {
    case 9223372036854775807:
        fmt.Println(v42)
    case 32767:
        v3 += float64((0.0 * 1e10))
    case 1:
        v45 := GenS2{F0: 891, F1: true}
        _ = v45
    }
    v2 += (v2 / ((genFn1()) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
}
