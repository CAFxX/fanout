package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 int; F1 uint32; F2 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32 }

func main() {
    var v1 uint32 = 65535
    var v2 uint64 = 18446744073709551615
    var v3 int64 = 478
    v4 := make(map[uint32]int)
    v5 := []uint32{v1, (v1 + 65535), (v1 / ((592) | 1))}
    var v6 bool = false
    v2 = v2
    v2 -= 882
    v7 := GenS1{F0: 2147483647, F1: 65535}
    var v8 uint = 1
    if v6 {
        if (!(!((0.1 - 1e10) != 0.1))) {
            var v9 int32 = 0
            if true {
                var v10 bool = (1.5 != -2.25)
                if false {
                    var v11 int32 = (v9 % ((42) | 1))
                    var v12 uint8 = 173
                    var v13 int8 = -128
                    _ = v11
                    _ = v12
                    _ = v13
                }
                _ = v10
            } else {
                var v14 bool = v6
                if true {
                    if (v14 || (v2 > 860)) {
                        if (v3 > (v3 + 176)) {
                            for fi15 := 0; fi15 < 10; fi15++ {
                                var v16 int = (-(fi15))
                                v9 -= v9
                                var v17 int16 = 32767
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                            if (int8(v8) > 43) {
                                if (((v8 / ((113) | 1)) <= v8) || ("a" > "hello")) {
                                    v18 := v4[1]
                                    switch ("\n\t" + ((("foo bar" + "foo bar") + "") + "hello")) {
                                    case "foo bar":
                                        if v14 {
                                            var v19 float32 = (float32((3.14159 * 3.14159)) / ((0.1 * 0.1) / 1e10))
                                            switch "foo bar" {
                                            case "hello":
                                                if v14 {
                                                    if ("" != ("" + "a")) {
                                                        for fi20 := 0; fi20 < 17; fi20++ {
                                                            for fi21 := 0; fi21 < 6; fi21++ {
                                                                fmt.Println(fi21, v2)
                                                                _ = fi21
                                                            }
                                                            if (-2.25 == ((3.14159 * 1e10) / (0.1 / 1.0))) {
                                                                v22 := GenS2{F0: (v18 * (fi20 & fi20)), F1: (v18 ^ fi20)}
                                                                var v23 int16 = 32767
                                                                _ = v22
                                                                _ = v23
                                                            }
                                                            var v24 int8 = 1
                                                            _ = fi20
                                                            _ = v24
                                                        }
                                                        v25 := GenS2{F0: 0, F1: 314}
                                                        _ = v25
                                                    } else {
                                                        v18 -= 0
                                                    }
                                                    v26 := GenS1{F0: v18, F1: (v8 | 1)}
                                                    if false {
                                                        var v27 int = 2147483647
                                                        if (-19 < 1) {
                                                            v28 := GenS3{F0: -128, F1: v2}
                                                            v29 := GenS2{F0: -749, F1: v18}
                                                            _ = v28
                                                            _ = v29
                                                        }
                                                        var v30 uint64 = v2
                                                        _ = v27
                                                        _ = v30
                                                    } else {
                                                        v31 := GenS4{F0: -737, F1: (v1 >> (((v8 | v8)) & 63)), F2: -434}
                                                        for fi32 := 0; fi32 < 10; fi32++ {
                                                            v33 := GenS4{F0: 42, F1: (v1 << ((v8) & 63)), F2: 1}
                                                            fmt.Println(v3, v8, v8)
                                                            _ = fi32
                                                            _ = v33
                                                        }
                                                        _ = v31
                                                    }
                                                    _ = v26
                                                }
                                            case "\n\t":
                                                for fi34 := 0; fi34 < 11; fi34++ {
                                                    for fi35 := 0; fi35 < 20; fi35++ {
                                                        switch ("a" + (("\n\t" + ("\n\t" + "a")) + "foo bar")) {
                                                        case "foo bar":
                                                            var v36 uint32 = 4294967295
                                                            _ = v36
                                                        case "\n\t":
                                                            fmt.Println(v19, v14, v3)
                                                        default:
                                                            v4[v1] = (v18 - -388)
                                                        }
                                                        switch ("a" + ("foo bar" + "foo bar")) {
                                                        case "a":
                                                            v37 := GenS2{F0: fi35, F1: -381}
                                                            _ = v37
                                                        }
                                                        _ = fi35
                                                    }
                                                    _ = fi34
                                                }
                                            case "a":
                                                v18 = 0
                                            }
                                            v19 = (((float32(0.0) + float32(0.0)) - (v19 + -2.25)) * 1e10)
                                            _ = v19
                                        } else {
                                            _ = len(v4)
                                        }
                                    case "a":
                                        var v38 int = (-(v18))
                                        _ = v38
                                    }
                                    if (!((v3 << ((v8) & 63)) <= (^(v3)))) {
                                        v9 -= v9
                                    }
                                    _ = v18
                                } else {
                                    var v39 uint8 = uint8(v2)
                                    _ = v39
                                }
                            } else {
                                fmt.Println(v1)
                                v8 += (v8 + (v8 | 1))
                                for fi40 := 0; fi40 < 3; fi40++ {
                                    var v41 string = "hello"
                                    var v42 uint = uint(v2)
                                    v5 = append(v5, 172)
                                    _ = fi40
                                    _ = v41
                                    _ = v42
                                }
                            }
                        } else {
                            for fi43 := 0; fi43 < 6; fi43++ {
                                v14 = v6
                                for fi44 := 0; fi44 < 6; fi44++ {
                                    var v45 int16 = 32767
                                    _ = fi44
                                    _ = v45
                                }
                                var v46 int32 = (v9 + int32(v2))
                                _ = fi43
                                _ = v46
                            }
                        }
                        var v47 uint64 = 558
                        _ = v47
                    } else {
                        var v48 float32 = float32((float64(-2.25) * 1e10))
                        v14 = v14
                        _ = v48
                    }
                } else {
                    v1 = v1
                    var v49 uint = v8
                    for fi50 := 0; fi50 < 13; fi50++ {
                        fmt.Println(v6)
                        _ = fi50
                    }
                    _ = v49
                }
                _ = v14
            }
            _ = v9
        } else {
            v3 = 0
            v3 = v3
            for fi51 := 0; fi51 < 17; fi51++ {
                v2 = 1
                var v52 uint = 65535
                _ = fi51
                _ = v52
            }
        }
    }
    fmt.Println(v6, v3)
    fmt.Println(v1, v2, v8)
    v6 = (false || ((v1 | (v1 | v1)) == (v1 >> (((v8 % (((v8 | 174)) | 1))) & 63))))
    v53 := GenS5{F0: 231}
    switch "foo bar" {
    case "\n\t":
        var v54 uint64 = 42
        _ = v54
    }
    v55 := GenS6{F0: 1}
    _ = len(v4)
    var v56 int16 = -974
    v3 = (v3 + (v3 | (v3 * int64(v56))))
    fmt.Println(v1, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v55
    _ = v56
    _ = v6
    _ = v7
    _ = v8
}
