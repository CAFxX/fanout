package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 string }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint16; F2 string }
type GenS5 struct { F0 int64; F1 int32; F2 int8 }
type GenS6 struct { F0 int; F1 uint8; F2 bool }
type GenS7 struct { F0 uint64; F1 string; F2 uint32 }
type GenS8 struct { F0 bool; F1 int64 }
func genFn0(p0 int, p1 bool, p2 uint8) int {
    var v6 int16 = 621
    p1 = false
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p0
}

func main() {
    var v1 int32 = -801
    v2 := []string{"foo bar", "hello", ((("a" + "foo bar") + "") + "\n\t")}
    var v3 float32 = (-2.25 / 1.5)
    v4 := []int{42, -190, -352, 42}
    var v5 uint32 = 592
    v7 := GenS1{F0: v5, F1: 687, F2: "foo bar"}
    if (1 < -1) {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint64 = 665
            for fi10 := 0; fi10 < 15; fi10++ {
                _ = len(v4)
                _ = fi10
            }
            v4 = append(v4, genFn0(genFn0(fi8, false, 0), false, 89))
            _ = fi8
            _ = v9
        }
        if false {
            _ = len(v2)
            var v11 uint32 = v5
            v4 = append(v4, -1)
            _ = v11
        }
    } else {
        if len(v2) > 0 {
            v2[0] = ("" + "hello")
        }
        v12 := GenS2{F0: 247, F1: 624}
        v5 = (v5 & 632)
        _ = v12
    }
    if (uint(v5) == 69) {
        fmt.Println(v1, v3)
        var v13 string = "foo bar"
        _ = v13
    } else {
        var v14 uint16 = 0
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 int = (fi15 & 0)
        var v17 uint = 711
        if len(v4) > 0 {
            v4[0] = fi15
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    v4 = append(v4, genFn0(-1, false, 0))
    fmt.Println(v5)
    for fi18 := 0; fi18 < 7; fi18++ {
        if len(v4) > 0 {
            v4[0] = fi18
        }
        var v19 uint16 = 65535
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: 1, F1: 1, F2: ""}
    for fi21 := 0; fi21 < 19; fi21++ {
        if false {
            if true {
                for fi22 := 0; fi22 < 20; fi22++ {
                    _ = len(v4)
                    if false {
                        var v23 uint8 = 255
                        _ = v23
                    } else {
                        if false {
                            v2 = append(v2, "")
                            var v24 int = (fi21 * fi22)
                            _ = v24
                        }
                        var v25 int8 = 100
                        fmt.Println(v3)
                        _ = v25
                    }
                    switch ("" + "hello") {
                    case "a":
                        var v26 int32 = (v1 | v1)
                        _ = v26
                    case "hello":
                        v27 := GenS1{F0: v5, F1: 118, F2: ("a" + "\n\t")}
                        _ = v27
                    default:
                        v2 = append(v2, "foo bar")
                    }
                    _ = fi22
                }
                var v28 bool = (66 <= 1)
                var v29 uint8 = 42
                _ = v28
                _ = v29
            }
            v3 -= 3.14159
            v30 := GenS2{F0: 493, F1: 255}
            _ = v30
        }
        for fi31 := 0; fi31 < 11; fi31++ {
            if ((v5 + v5) == 0) {
                v32 := GenS3{F0: (0 == -128)}
                _ = v32
            }
            _ = fi31
        }
        _ = fi21
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 float64 = 3.14159
        v35 := GenS4{F0: (false || false), F1: 624, F2: ("foo bar" + (("foo bar" + "foo bar") + "foo bar"))}
        if true {
            v36 := GenS5{F0: -1, F1: v1, F2: -1}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v37 := GenS6{F0: (fi33 - (fi33 | genFn0((fi33 - genFn0(128, false, 0)), true, 127))), F1: 128, F2: true}
            _ = v36
            _ = v37
        } else {
            v38 := GenS7{F0: 42, F1: "foo bar", F2: v5}
            v39 := GenS8{F0: (2 == 4294967295), F1: 1}
            _ = v38
            _ = v39
        }
        _ = fi33
        _ = v34
        _ = v35
    }
    if (1 != 18446744073709551615) {
        v3 = 1.5
    }
    var v40 uint16 = 0
    var v41 string = (("" + ("\n\t" + "")) + ("" + "hello"))
    var v42 int16 = 1
    v40 = (v40 * v40)
    switch 564 {
    case 32767:
        _ = len(v4)
    case 127:
        var v43 int8 = -1
        _ = v43
    default:
        fmt.Println(v42)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v7
}
