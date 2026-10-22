package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 uint64; F2 bool }
type GenS6 struct { F0 uint32; F1 bool }
type GenS7 struct { F0 uint16 }
func genFn0(p0 uint8) uint64 {
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return 1
}
func genFn1(p0 int64, p1 uint, p2 int16) int32 {
    if true {
        if true {
            for fi4 := 0; fi4 < 20; fi4++ {
                if true {
                    if ("\n\t" == "foo bar") {
                        p2 = (p2 ^ 1)
                        v5 := []int{fi4}
                        _ = v5
                    } else {
                        v6 := GenS1{F0: 268, F1: (p1 << ((p1) & 63))}
                        _ = v6
                    }
                }
                v7 := []int{fi4}
                v8 := GenS2{F0: 42}
                _ = fi4
                _ = v7
                _ = v8
            }
        } else {
            var v9 uint = p1
            fmt.Println(p0, p1)
            var v10 uint16 = 1
            _ = v10
            _ = v9
        }
        var v11 string = ("a" + "\n\t")
        _ = v11
    } else {
        var v12 int = -924
        _ = v12
    }
    if false {
        v13 := GenS2{F0: -1}
        _ = v13
    } else {
        v14 := GenS1{F0: 266, F1: (^(p1))}
        _ = v14
    }
    if (((p1 >= (p1 / ((p1) | 1))) || (!false)) || false) {
        var v15 float32 = 1.5
        _ = v15
    } else {
        v16 := []int{-9223372036854775808, 0, 127, -2147483648}
        v17 := GenS3{F0: p0, F1: 18446744073709551615}
        _ = v16
        _ = v17
    }
    p0 += p0
    if false {
        var v18 string = "foo bar"
        var v19 uint64 = 0
        _ = v18
        _ = v19
    } else {
        fmt.Println(p0)
        v20 := GenS4{F0: (p0 ^ -9223372036854775808)}
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 100
    v3 := []string{"foo bar", "hello", ("" + (("foo bar" + ("\n\t" + "a")) + "a"))}
    var v21 int8 = 1
    v22 := GenS3{F0: 42, F1: 567}
    v21 = v21
    if false {
        v23 := GenS5{F0: 255, F1: 0, F2: (0 <= 0)}
        var v24 uint32 = 65535
        _ = v23
        _ = v24
    }
    var v25 uint32 = 1
    v21 = v2
    fmt.Println(v2)
    v26 := GenS4{F0: 9223372036854775807}
    var v27 int16 = -559
    v28 := GenS5{F0: 269, F1: 789, F2: false}
    v29 := GenS6{F0: (v25 * 870), F1: true}
    v30 := GenS4{F0: -9223372036854775808}
    if true {
        if (214 == 0) {
            var v31 float64 = (((3.14159 / (-2.25 / 3.14159)) / 1.5) - (3.14159 + -2.25))
            _ = v31
        }
    } else {
        var v32 int32 = -1
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        for fi33 := 0; fi33 < 18; fi33++ {
            if ((-2.25 <= 1e10) || (("a" + (((("hello" + "hello") + "foo bar") + ("\n\t" + "")) + "")) < "foo bar")) {
                var v34 int32 = genFn1(-1, 384, 345)
                for fi35 := 0; fi35 < 4; fi35++ {
                    var v36 int16 = v27
                    _ = fi35
                    _ = v36
                }
                v37 := GenS5{F0: 2, F1: genFn0(127), F2: true}
                _ = v34
                _ = v37
            }
            var v38 int32 = (^(v32))
            if (!false) {
                if true {
                    var v39 bool = true
                    _ = v39
                } else {
                    for fi40 := 0; fi40 < 18; fi40++ {
                        var v41 uint8 = 0
                        var v42 uint64 = 1
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    for fi43 := 0; fi43 < 11; fi43++ {
                        var v44 uint64 = 4294967295
                        if true {
                            v45 := GenS4{F0: 9223372036854775807}
                            _ = v45
                        }
                        var v46 uint16 = 1
                        var v47 float32 = 3.14159
                        _ = fi43
                        _ = v44
                        _ = v46
                        _ = v47
                    }
                }
                v32 -= int32(fi33)
                var v48 uint32 = 4294967295
                _ = v48
            } else {
                v21 += v2
                for fi49 := 0; fi49 < 17; fi49++ {
                    var v50 float32 = 0.0
                    switch "foo bar" {
                    case "a":
                        v51 := GenS3{F0: 215, F1: 4294967295}
                        _ = v51
                    }
                    if false {
                        var v52 int64 = 42
                        var v53 string = "\n\t"
                        v50 = float32((0.0 / float64(((1e10 - (0.0 + -2.25)) / 1e10))))
                        _ = v52
                        _ = v53
                    }
                    _ = fi49
                    _ = v50
                }
                fmt.Println(fi33, v27)
            }
            _ = fi33
            _ = v38
        }
        _ = v32
    }
    var v54 uint64 = 42
    v55 := GenS7{F0: 2}
    if (!(!(!(!(-2147483648 >= 2147483647))))) {
        fmt.Println(v2, v54, v2)
    }
    for fi56 := 0; fi56 < 15; fi56++ {
        for fi57 := 0; fi57 < 16; fi57++ {
            if len(v3) > 0 {
                v3[0] = ("foo bar" + "")
            }
            v3 = append(v3, "a")
            _ = fi57
        }
        _ = fi56
    }
    if true {
        v25 += v25
        v2 += (-(v21))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v54
    _ = v55
}
