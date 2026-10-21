package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 int; F2 int32 }
type GenS6 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS7 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint64, p1 bool, p2 uint32) uint8 {
    var v7 uint64 = p0
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int16 = 1
        v10 := []int{fi8, (^(fi8)), 128, (fi8 - fi8)}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 205
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 1
    var v3 int = -1
    var v4 uint32 = 0
    v5 := make(map[uint32]int)
    v6 := make(map[string]int)
    v2 = 4294967295
    v11 := []int{(v3 & 32768), v3, v3, (v3 << (((v2 << ((v2) & 63))) & 63))}
    v12 := GenS1{F0: (v2 | 4294967295)}
    v13 := GenS2{F0: 0, F1: 72, F2: 534}
    var v14 uint = 0
    v15 := GenS1{F0: uint(v4)}
    for fi16 := 0; fi16 < 3; fi16++ {
        if (v3 <= (-(v3))) {
            v1["foo bar"] = -1
            v17 := GenS3{F0: 255, F1: 4294967295, F2: -1}
            v18 := GenS4{F0: 1}
            _ = v17
            _ = v18
        } else {
            if false {
                v14 -= 2
                var v19 int64 = 241
                _ = v19
            } else {
                if len(v11) > 0 {
                    v11[0] = (fi16 << ((65535) & 63))
                }
                v14 -= 2
                fmt.Println(v14, fi16, v3)
            }
            var v20 float32 = 3.14159
            for fi21 := 0; fi21 < 20; fi21++ {
                switch (-(fi21)) {
                case -9223372036854775808:
                    if false {
                        var v22 int = (-(v3))
                        switch ("" + ("foo bar" + ("a" + "foo bar"))) {
                        case "hello":
                            fmt.Println(fi16, v4, v3)
                        case "":
                            v22 += -2147483648
                            fmt.Println(v14)
                            if (float64(v20) > (3.14159 - 1e10)) {
                                var v23 float32 = v20
                                _ = v23
                            }
                            v2 = v2
                            var v24 uint32 = (^(v4))
                            _ = v24
                        case "\n\t":
                            switch 4294967295 {
                            case -1:
                                if (("a" + (("hello" + "foo bar") + "hello")) != "hello") {
                                    if (0.0 <= 1.5) {
                                        if false {
                                            if len(v11) > 0 {
                                                v11[0] = fi21
                                            }
                                        }
                                        var v25 string = ((("" + ("\n\t" + ("hello" + "foo bar"))) + (("\n\t" + "a") + "foo bar")) + ("hello" + ("" + (("foo bar" + "\n\t") + (("hello" + "") + ("\n\t" + ("\n\t" + (("\n\t" + "\n\t") + "foo bar"))))))))
                                        v1[v25] = int(v4)
                                        _ = v25
                                    } else {
                                        var v26 int64 = 903
                                        _ = v26
                                    }
                                    var v27 uint32 = v4
                                    _ = v27
                                } else {
                                    v28 := GenS5{F0: 42, F1: (-(v22)), F2: -1}
                                    if (-2.25 != ((((3.14159 - 0.0) + 1.5) - 1.5) - 0.1)) {
                                        if len(v11) > 0 {
                                            v11[0] = (v22 ^ 127)
                                        }
                                    } else {
                                        var v29 int64 = 0
                                        v2 += v14
                                        v4 = v4
                                        v30 := v6[("" + "foo bar")]
                                        if false {
                                            v2 -= 42
                                            v31 := GenS2{F0: 2, F1: 1, F2: 456}
                                            _ = v31
                                        }
                                        _ = v29
                                        _ = v30
                                    }
                                    _ = v28
                                }
                            case 519:
                                v3 -= fi21
                            case 300:
                                v3 -= (v3 / ((v3) | 1))
                            }
                        }
                        _ = v22
                    } else {
                        fmt.Println(v4, v14, v2)
                        var v32 float32 = (((v20 / 1.0) + 0.1) + (((v20 * v20) / v20) / v20))
                        if ((-276 != int64(v2)) && (!false)) {
                            if false {
                                var v33 float64 = -2.25
                                fmt.Println(v2, fi21)
                                _ = v33
                            }
                            var v34 int = (fi16 ^ fi16)
                            v35 := GenS1{F0: (v2 * (v14 * 255))}
                            _ = v34
                            _ = v35
                        }
                        _ = v32
                    }
                    v36 := GenS3{F0: 509, F1: 49, F2: 2147483647}
                    _ = v36
                default:
                    switch ("\n\t" + (("\n\t" + "hello") + (("" + "") + ("hello" + ("hello" + ("foo bar" + ("\n\t" + "hello"))))))) {
                    case "hello":
                        var v37 int64 = 0
                        _ = v37
                    case "":
                        v4 -= 4294967295
                    }
                }
                if (42 <= 42) {
                    fmt.Println(v14, v20)
                    v38 := v1["a"]
                    _ = v38
                } else {
                    switch "hello" {
                    case "":
                        v14 -= (v2 << (((^(v2))) & 63))
                    case "hello":
                        var v39 string = ("" + "hello")
                        _ = v39
                    case "foo bar":
                        v40 := GenS6{F0: -550, F1: 1, F2: 475}
                        _ = v40
                    }
                    switch (fi21 + (fi21 | v3)) {
                    case 32768:
                        var v41 uint = v14
                        _ = v41
                    case 42:
                        if ((false && true) || ((fi16 <= (fi16 / ((v3) | 1))) || false)) {
                            v3 -= 128
                        }
                    case 128:
                        v42 := GenS1{F0: (v14 | 4294967295)}
                        _ = v42
                    default:
                        fmt.Println(fi21, v14, v2)
                    }
                    var v43 uint = (v14 ^ v14)
                    _ = v43
                }
                switch (v3 + (fi16 / (((-(fi21))) | 1))) {
                case 1:
                    var v44 uint64 = 490
                    _ = v44
                case 847:
                    v20 = v20
                case 42:
                    v14 = (v14 + (^(v14)))
                }
                var v45 float64 = 1e10
                _ = fi21
                _ = v45
            }
            _ = v20
        }
        for fi46 := 0; fi46 < 9; fi46++ {
            v47 := v5[(v4 & (v4 ^ (v4 | (v4 ^ v4))))]
            v2 -= (v2 << (((v14 * (v2 ^ v2))) & 63))
            if false {
                var v48 uint16 = 789
                delete(v6, "")
                v11 = append(v11, 389)
                _ = v48
            }
            _ = fi46
            _ = v47
        }
        _ = fi16
    }
    var v49 uint = v2
    for fi50 := 0; fi50 < 4; fi50++ {
        switch "\n\t" {
        case "":
            var v51 int32 = -28
            _ = v51
        case "hello":
            var v52 bool = (75 >= 100)
            _ = v52
        case "a":
            var v53 bool = true
            _ = v53
        default:
            v54 := GenS7{F0: -818, F1: v4}
            _ = v54
        }
        _ = fi50
    }
    fmt.Println(v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
