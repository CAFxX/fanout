package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 bool }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int64; F1 uint; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint; F2 int }
func genFn0(p0 string) int {
    for fi4 := 0; fi4 < 4; fi4++ {
        v5 := []int{-514}
        if (42 < 42) {
            var v6 int = fi4
            if false {
                v5 = append(v5, -201)
                p0 = (((("hello" + "\n\t") + (p0 + "foo bar")) + "") + "")
            }
            var v7 string = ("hello" + "foo bar")
            _ = v6
            _ = v7
        } else {
            var v8 int64 = -234
            if len(v5) > 0 {
                v5[0] = (^(fi4))
            }
            v9 := GenS1{F0: true, F1: 1000, F2: true}
            _ = v8
            _ = v9
        }
        if len(v5) > 0 {
            v5[0] = (fi4 + 435)
        }
        _ = fi4
        _ = v5
    }
    p0 = p0
    switch -1 {
    case 4294967295:
        switch 339 {
        case -671:
            if true {
                fmt.Println(p0)
                fmt.Println(p0)
            }
        case 169:
            var v10 float64 = 0.0
            _ = v10
        default:
            if (42 <= 61) {
                if true {
                    if ((255 != 65535) || (-32768 < 42)) {
                        var v11 string = p0
                        p0 = "\n\t"
                        _ = v11
                    }
                    fmt.Println(p0)
                } else {
                    if true {
                        for fi12 := 0; fi12 < 9; fi12++ {
                            for fi13 := 0; fi13 < 2; fi13++ {
                                p0 = p0
                                v14 := GenS2{F0: 18446744073709551615, F1: 0}
                                var v15 string = ("" + ("" + p0))
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                            switch p0 {
                            case "a":
                                p0 = (((p0 + p0) + p0) + p0)
                            case "\n\t":
                                var v16 int = fi12
                                _ = v16
                            default:
                                var v17 uint16 = 136
                                _ = v17
                            }
                            _ = fi12
                        }
                    } else {
                        if true {
                            var v18 int32 = -2147483648
                            var v19 float32 = 3.14159
                            fmt.Println(p0)
                            _ = v18
                            _ = v19
                        }
                        v20 := GenS3{F0: -25, F1: 2, F2: 0}
                        _ = v20
                    }
                }
            } else {
                if false {
                    fmt.Println(p0)
                    var v21 uint64 = 4294967295
                    v22 := []int{}
                    _ = v21
                    _ = v22
                } else {
                    var v23 bool = false
                    switch p0 {
                    case "a":
                        v24 := GenS4{F0: 32767, F1: 0, F2: 2147483647}
                        _ = v24
                    case "foo bar":
                        v23 = ((float64(-2.25) / 1e10) < float64(0.0))
                    case "":
                        v25 := []int{-9223372036854775808, 2147483647, 42, 0}
                        _ = v25
                    }
                    _ = v23
                }
            }
        }
    default:
        v26 := GenS3{F0: 9223372036854775807, F1: 1, F2: -1}
        _ = v26
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 uint32 = 4294967295
        if (false && false) {
            var v29 uint32 = 0
            _ = v29
        }
        var v30 int64 = 1
        _ = fi27
        _ = v28
        _ = v30
    }
    if (!true) {
        p0 = "\n\t"
    } else {
        var v31 uint64 = 1
        _ = v31
    }
    _ = p0
    return 4294967295
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[string]uint32)
    var v3 bool = false
    if v3 {
        if (!(v3 || false)) {
            for fi32 := 0; fi32 < 13; fi32++ {
                fmt.Println(fi32)
                fmt.Println(fi32, v1)
                _ = fi32
            }
        }
        var v33 bool = false
        _ = v33
    }
    v34 := GenS2{F0: 4294967295, F1: 100}
    for fi35 := 0; fi35 < 8; fi35++ {
        fmt.Println(v1, v3, v3)
        _ = fi35
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        var v37 int32 = -2147483648
        var v38 float32 = -2.25
        _ = fi36
        _ = v37
        _ = v38
    }
    v1 += 32767
    for fi39 := 0; fi39 < 15; fi39++ {
        var v40 int64 = 427
        v1 = (v1 << ((0) & 63))
        var v41 int = (fi39 % ((1) | 1))
        _ = fi39
        _ = v40
        _ = v41
    }
    v42 := GenS2{F0: 65535, F1: 108}
    var v43 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v42
    _ = v43
}
