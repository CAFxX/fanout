package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32 }
func genFn0() uint64 {
    var v6 int64 = 42
    switch 297 {
    case -584:
        v7 := []int{128, 268}
        _ = v7
    default:
        v8 := []int{}
        _ = v8
    }
    if (((((((3.14159 + 0.0) - ((1e10 / (-2.25 - 0.0)) - float32(-2.25))) + 1e10) - 1e10) + 0.0) <= 1e10) || false) {
        v9 := GenS1{F0: 42}
        v10 := []int{127, -1, 4294967295, -764}
        _ = v10
        _ = v9
    } else {
        v11 := []int{-9223372036854775808}
        _ = v11
    }
    var v12 float64 = (1.5 + 1e10)
    fmt.Println(v6, v6)
    _ = v12
    _ = v6
    return 133
}

func main() {
    v1 := []uint32{}
    var v2 float64 = 1.5
    var v3 string = "foo bar"
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]bool)
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 float32 = float32(v2)
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: 231}
    switch v3 {
    case "hello":
        var v16 string = ""
        _ = v16
    case "a":
        v17 := GenS1{F0: -2147483648}
        _ = v17
    case "foo bar":
        _ = len(v1)
    default:
        var v18 uint32 = 0
        _ = v18
    }
    v2 = 1.5
    switch ("" + "hello") {
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch "\n\t" {
    case "a":
        if len(v1) > 0 {
            v1[0] = 1
        }
    case "\n\t":
        v19 := GenS1{F0: -9223372036854775808}
        _ = v19
    case "foo bar":
        var v20 int32 = -2147483648
        _ = v20
    }
    var v21 uint16 = 65535
    _ = len(v1)
    switch 42 {
    case -2147483648:
        var v22 int = 32767
        _ = v22
    case -901:
        v3 = v3
    }
    if false {
        switch v3 {
        case "hello":
            v3 = (v3 + (v3 + v3))
        case "":
            v2 += v2
        }
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        if false {
            v24 := GenS2{F0: false, F1: 32768}
            var v25 uint8 = 115
            if true {
                for fi26 := 0; fi26 < 14; fi26++ {
                    if (!(((42 == -9223372036854775808) && (true || true)) || (42 == 42))) {
                        v27 := GenS1{F0: (fi23 << ((257) & 63))}
                        var v28 int16 = 32767
                        _ = v27
                        _ = v28
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v29 int = -9223372036854775808
                        _ = v29
                    }
                    _ = fi26
                }
                if (706 < 1000) {
                    v30 := v4[969]
                    var v31 bool = true
                    _ = v30
                    _ = v31
                } else {
                    if false {
                        if (false && (!false)) {
                            for fi32 := 0; fi32 < 19; fi32++ {
                                fmt.Println(fi23)
                                v3 = ("a" + v3)
                                _ = fi32
                            }
                            for fi33 := 0; fi33 < 19; fi33++ {
                                v34 := GenS2{F0: (!false), F1: int(v21)}
                                _ = fi33
                                _ = v34
                            }
                        } else {
                            for fi35 := 0; fi35 < 19; fi35++ {
                                v36 := GenS1{F0: fi35}
                                fmt.Println(v25, v3)
                                var v37 int64 = 582
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                            v2 -= v2
                            if ((4294967295 <= 42) || (!false)) {
                                fmt.Println(v3, v21)
                                fmt.Println(v3)
                            } else {
                                v38 := GenS3{F0: -986}
                                _ = v38
                            }
                        }
                        v2 -= 0.1
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    v25 += v25
                    v1 = append(v1, uint32(v25))
                }
                var v39 uint = 86
                _ = v39
            } else {
                switch (fi23 << ((0) & 63)) {
                case 189:
                    var v40 bool = false
                    _ = v40
                case -70:
                    switch "foo bar" {
                    case "a":
                        v41 := GenS1{F0: int(v25)}
                        _ = v41
                    case "hello":
                        var v42 uint8 = 2
                        _ = v42
                    }
                }
                v2 += float64(3.14159)
            }
            _ = v24
            _ = v25
        } else {
            v2 -= (v2 / v2)
            v43 := GenS4{F0: 0}
            v21 = (v21 % ((255) | 1))
            _ = v43
        }
        _ = fi23
    }
    switch 9223372036854775807 {
    case 0:
        for fi44 := 0; fi44 < 6; fi44++ {
            fmt.Println(v3, v2, v21)
            if len(v1) > 0 {
                v1[0] = 65535
            }
            _ = fi44
        }
    }
    _ = len(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
