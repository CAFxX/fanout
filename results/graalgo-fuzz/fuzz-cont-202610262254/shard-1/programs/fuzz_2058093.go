package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint32; F1 int16; F2 int }
type GenS3 struct { F0 uint64; F1 int; F2 int32 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint }
func genFn0(p0 uint8, p1 uint8, p2 uint64) bool {
    v6 := []int{-17, -9223372036854775808}
    v7 := GenS1{F0: 0, F1: 1}
    v6 = append(v6, -2147483648)
    for fi8 := 0; fi8 < 1; fi8++ {
        _ = len(v6)
        if false {
            var v9 uint32 = 672
            _ = v9
        } else {
            switch 0 {
            case 32767:
                v6 = append(v6, fi8)
            case 0:
                var v10 uint32 = 0
                _ = v10
            default:
                v6 = append(v6, fi8)
            }
            var v11 float32 = 0.0
            _ = v11
        }
        fmt.Println(p2, p2)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (!(true && true))
}

func main() {
    var v1 string = "hello"
    v2 := []string{"", v1, ("" + (("foo bar" + v1) + v1))}
    var v3 uint64 = 1
    var v4 uint32 = 928
    v5 := make(map[string]string)
    if false {
        v12 := GenS2{F0: 2, F1: 1, F2: 0}
        v13 := GenS3{F0: v3, F1: 976, F2: -1}
        _ = v12
        _ = v13
    } else {
        v14 := GenS4{F0: (v1 + v1), F1: 42}
        _ = v14
    }
    v3 += 18446744073709551615
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS5{F0: (v3 + v3), F1: 2}
        _ = len(v2)
        var v17 uint32 = v4
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := v5[v1]
    v19 := GenS6{F0: -887}
    var v20 int64 = 9223372036854775807
    for fi21 := 0; fi21 < 19; fi21++ {
        for fi22 := 0; fi22 < 12; fi22++ {
            var v23 int64 = v20
            _ = len(v2)
            v5[v18] = "foo bar"
            v23 -= v20
            v2 = append(v2, "foo bar")
            _ = fi22
            _ = v23
        }
        var v24 uint8 = 1
        if ((v4 << ((1) & 63)) <= v4) {
            v25 := GenS1{F0: int8(v20), F1: 42}
            _ = v25
        } else {
            switch ((v1 + (("foo bar" + v18) + v18)) + v18) {
            case "a":
                v3 -= 85
            default:
                var v26 float32 = (float32((1.5 - (0.0 + 0.1))) + (-2.25 - (0.1 / -2.25)))
                _ = v26
            }
            fmt.Println(fi21)
        }
        _ = fi21
        _ = v24
    }
    v3 += (^(v3))
    var v27 int32 = 42
    for fi28 := 0; fi28 < 13; fi28++ {
        var v29 int = 127
        fmt.Println(v3, v3, v4)
        if false {
            v30 := GenS7{F0: 669}
            _ = v30
        } else {
            v31 := GenS2{F0: 173, F1: 1, F2: (-(fi28))}
            for fi32 := 0; fi32 < 4; fi32++ {
                var v33 float64 = -2.25
                _ = fi32
                _ = v33
            }
            if ((^(v4)) >= v4) {
                var v34 int64 = v20
                for fi35 := 0; fi35 < 5; fi35++ {
                    var v36 float64 = (float64(0.0) / (-2.25 - 1e10))
                    v36 = float64(1e10)
                    v27 = v27
                    _ = fi35
                    _ = v36
                }
                var v37 uint = 18446744073709551615
                _ = v34
                _ = v37
            } else {
                fmt.Println(v27, v20)
                switch "a" {
                case "\n\t":
                    var v38 float64 = 1e10
                    _ = v38
                case "a":
                    switch v29 {
                    case 42:
                        v18 = v18
                    case -9223372036854775808:
                        if (817 >= 183) {
                            switch v29 {
                            case 437:
                                var v39 int16 = -1
                                _ = v39
                            case 0:
                                if true {
                                    fmt.Println(v3, v29, v3)
                                    switch (v29 ^ fi28) {
                                    case -9223372036854775808:
                                        for fi40 := 0; fi40 < 17; fi40++ {
                                            var v41 uint = 659
                                            if true {
                                                fmt.Println(v3)
                                            } else {
                                                _ = len(v2)
                                                switch (fi40 | fi28) {
                                                case -2147483648:
                                                    v42 := v5[v1]
                                                    _ = v42
                                                default:
                                                    if false {
                                                        var v43 int32 = -1
                                                        _ = v43
                                                    }
                                                }
                                            }
                                            _ = fi40
                                            _ = v41
                                        }
                                    case 32768:
                                        var v44 int8 = int8(v29)
                                        _ = v44
                                    default:
                                        fmt.Println(fi28, v4, v29)
                                    }
                                }
                            case -414:
                                v45 := GenS2{F0: (v4 - 65535), F1: -407, F2: fi28}
                                _ = v45
                            default:
                                for fi46 := 0; fi46 < 16; fi46++ {
                                    if true {
                                        _ = len(v2)
                                        var v47 float32 = 3.14159
                                        if len(v2) > 0 {
                                            v2[0] = v18
                                        }
                                        _ = v47
                                    }
                                    _ = fi46
                                }
                            }
                        } else {
                            switch v29 {
                            case -548:
                                fmt.Println(v29, v27, v3)
                            case -181:
                                var v48 uint = 0
                                _ = v48
                            case 42:
                                if true {
                                    v4 = 304
                                    v29 += v29
                                    if ((!true) || false) {
                                        v18 = "\n\t"
                                        v49 := GenS3{F0: v3, F1: fi28, F2: v27}
                                        v1 = v1
                                        _ = v49
                                    } else {
                                        fmt.Println(fi28, v1, fi28)
                                    }
                                }
                            }
                        }
                    default:
                        if false {
                            v3 -= v3
                            for fi50 := 0; fi50 < 6; fi50++ {
                                v4 -= v4
                                if true {
                                    v27 -= 42
                                }
                                var v51 uint16 = 42
                                _ = fi50
                                _ = v51
                            }
                        }
                    }
                default:
                    for fi52 := 0; fi52 < 10; fi52++ {
                        var v53 int = fi28
                        v54 := GenS7{F0: 377}
                        _ = fi52
                        _ = v53
                        _ = v54
                    }
                }
                for fi55 := 0; fi55 < 10; fi55++ {
                    fmt.Println(fi28)
                    delete(v5, ("\n\t" + (v18 + (v18 + ("\n\t" + "\n\t")))))
                    fmt.Println(v20)
                    _ = fi55
                }
            }
            _ = v31
        }
        _ = fi28
        _ = v29
    }
    if (0 != v4) {
        v4 -= (v4 << ((4294967295) & 63))
        v1 = v18
        fmt.Println(v18, v18)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
