package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int64 }
func genFn0(p0 uint64) uint {
    if true {
        p0 -= p0
        p0 = (p0 ^ 42)
        fmt.Println(p0)
    }
    fmt.Println(p0)
    if (!(32767 != 32767)) {
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := []int{fi6, 0}
            var v8 float32 = (1e10 * ((0.0 + 1e10) * 1e10))
            v8 += v8
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 int32 = 2147483647
        var v10 int16 = 1
        _ = v10
        _ = v9
    }
    v11 := []int{485, 4294967295, 127}
    switch 9223372036854775807 {
    case 32767:
        if ((!(((0.1 + (((-2.25 / ((float32((-2.25 + 0.0)) - 0.1) - -2.25)) + 1e10) + (1.5 * 3.14159))) < 0.1) || false)) && (!false)) {
            p0 = (p0 % ((p0) | 1))
        }
    case 0:
        if true {
            if (!(-2147483648 >= 1)) {
                v12 := GenS1{F0: 255}
                fmt.Println(p0)
                var v13 int64 = 861
                _ = v12
                _ = v13
            } else {
                switch "foo bar" {
                case "":
                    p0 = (p0 | p0)
                case "foo bar":
                    for fi14 := 0; fi14 < 9; fi14++ {
                        var v15 float32 = 0.0
                        p0 = (p0 - 18446744073709551615)
                        if len(v11) > 0 {
                            v11[0] = -930
                        }
                        _ = fi14
                        _ = v15
                    }
                case "a":
                    if false {
                        switch "foo bar" {
                        case "hello":
                            var v16 float32 = (0.1 * (3.14159 * -2.25))
                            _ = v16
                        case "\n\t":
                            var v17 uint64 = 1
                            _ = v17
                        case "a":
                            fmt.Println(p0)
                        }
                        if false {
                            v18 := GenS2{F0: ((("\n\t" + "\n\t") + ((("" + ("\n\t" + "hello")) + "foo bar") + "")) + "hello")}
                            _ = v18
                        }
                    }
                default:
                    for fi19 := 0; fi19 < 13; fi19++ {
                        if len(v11) > 0 {
                            v11[0] = -657
                        }
                        _ = fi19
                    }
                }
                var v20 int16 = 0
                _ = v20
            }
        } else {
            var v21 float64 = 1.5
            v11 = append(v11, 1)
            _ = v21
        }
    case 128:
        p0 = (^(p0))
    }
    _ = p0
    _ = v11
    return 1
}
func genFn1() int64 {
    for fi22 := 0; fi22 < 8; fi22++ {
        if (127 != 0) {
            var v23 uint8 = 2
            var v24 string = ""
            v24 = v24
            _ = v23
            _ = v24
        } else {
            v25 := GenS2{F0: "hello"}
            v26 := []int{fi22, (fi22 << ((255) & 63)), (fi22 << ((2) & 63))}
            var v27 string = "\n\t"
            _ = v25
            _ = v26
            _ = v27
        }
        var v28 float32 = -2.25
        v29 := []int{(-(fi22)), fi22}
        _ = fi22
        _ = v28
        _ = v29
    }
    var v30 uint = 184
    _ = v30
    return 42
}

func main() {
    var v1 int8 = -41
    v2 := []uint32{65535}
    var v3 float32 = (float32(0.0) / float32(3.14159))
    v4 := make(map[string]int)
    v5 := make(map[uint32]int)
    for fi31 := 0; fi31 < 7; fi31++ {
        if false {
            var v32 float32 = v3
            v33 := GenS3{F0: (!false)}
            _ = v32
            _ = v33
        } else {
            var v34 int32 = -2147483648
            _ = v34
        }
        _ = fi31
    }
    for fi35 := 0; fi35 < 4; fi35++ {
        var v36 float32 = ((0.0 - 1e10) * float32(0.1))
        for fi37 := 0; fi37 < 12; fi37++ {
            var v38 int8 = (v1 << ((255) & 63))
            if ((!false) && false) {
                var v39 float32 = 0.0
                _ = v39
            } else {
                var v40 int8 = v1
                if (true || true) {
                    fmt.Println(fi37, v38)
                    switch (fi37 >> ((695) & 63)) {
                    case -9223372036854775808:
                        var v41 uint32 = 42
                        _ = v41
                    }
                    v42 := GenS4{F0: "\n\t", F1: 0}
                    _ = v42
                }
                if true {
                    fmt.Println(v36, v38, v36)
                    var v43 int64 = 718
                    fmt.Println(fi37, v3, v36)
                    _ = v43
                }
                _ = v40
            }
            var v44 uint16 = 1
            var v45 int32 = 716
            _ = fi37
            _ = v38
            _ = v44
            _ = v45
        }
        _ = fi35
        _ = v36
    }
    v46 := GenS3{F0: ((!(!(!true))) || (291 >= 1))}
    if (0 != 2147483648) {
        v47 := GenS5{F0: 2}
        v1 = -1
        _ = v47
    } else {
        var v48 uint32 = 4294967295
        var v49 string = "a"
        v50 := GenS6{F0: 344}
        _ = v48
        _ = v49
        _ = v50
    }
    fmt.Println(v3, v1)
    _ = len(v2)
    delete(v4, "")
    var v51 float64 = float64(((float32(0.0) * 0.1) / 0.1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v51
}
