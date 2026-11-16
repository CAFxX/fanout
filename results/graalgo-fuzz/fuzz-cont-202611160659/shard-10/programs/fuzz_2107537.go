package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint32 }
func genFn0() uint {
    for fi4 := 0; fi4 < 12; fi4++ {
        fmt.Println(fi4)
        for fi5 := 0; fi5 < 19; fi5++ {
            switch -143 {
            case 1:
                fmt.Println(fi4, fi4)
            case 9223372036854775807:
                var v6 uint16 = 0
                _ = v6
            case 32768:
                v7 := GenS1{F0: "hello"}
                _ = v7
            }
            switch "hello" {
            case "":
                var v8 float64 = 1.5
                _ = v8
            default:
                var v9 uint16 = 32767
                v9 = v9
                v10 := GenS2{F0: 1, F1: 255}
                var v11 int = (fi4 + fi4)
                var v12 uint16 = (v9 >> ((294) & 63))
                _ = v10
                _ = v11
                _ = v12
                _ = v9
            }
            _ = fi5
        }
        _ = fi4
    }
    var v13 uint8 = 0
    if true {
        for fi14 := 0; fi14 < 19; fi14++ {
            for fi15 := 0; fi15 < 6; fi15++ {
                v16 := GenS1{F0: ""}
                switch fi15 {
                case -9223372036854775808:
                    v13 = v13
                case 9223372036854775807:
                    if true {
                        if (0 <= 757) {
                            fmt.Println(fi14)
                            v13 -= v13
                        } else {
                            v17 := []int{fi14, (fi15 & fi15)}
                            v13 = v13
                            _ = v17
                        }
                        v13 = v13
                        v13 += v13
                        fmt.Println(fi15, fi15, fi14)
                    }
                default:
                    var v18 float64 = (0.1 / 0.1)
                    _ = v18
                }
                v19 := GenS1{F0: (((("hello" + "hello") + "a") + "hello") + "\n\t")}
                _ = fi15
                _ = v16
                _ = v19
            }
            _ = fi14
        }
        for fi20 := 0; fi20 < 15; fi20++ {
            if true {
                var v21 int8 = -128
                _ = v21
            } else {
                for fi22 := 0; fi22 < 14; fi22++ {
                    switch "\n\t" {
                    case "hello":
                        var v23 uint = 4294967295
                        _ = v23
                    case "a":
                        v13 = (v13 / ((182) | 1))
                    }
                    if false {
                        v13 += v13
                        v13 = v13
                        fmt.Println(fi22)
                    }
                    fmt.Println(fi20, fi22)
                    _ = fi22
                }
                v13 -= (^(v13))
            }
            v24 := GenS1{F0: (((("" + "") + ("\n\t" + ("" + "\n\t"))) + "") + ("a" + (("foo bar" + ((("foo bar" + (("foo bar" + "\n\t") + ("hello" + "foo bar"))) + "a") + ("hello" + "\n\t"))) + "\n\t")))}
            _ = fi20
            _ = v24
        }
    }
    v25 := GenS1{F0: "hello"}
    _ = v13
    _ = v25
    return 613
}
func genFn1(p0 uint32, p1 uint16, p2 uint16) uint16 {
    var v26 uint16 = (p2 + p2)
    p2 = p2
    v27 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    return v26
}

func main() {
    v1 := []string{"hello"}
    v2 := make(map[int]int)
    var v3 uint32 = 42
    var v28 int16 = -299
    switch 9223372036854775807 {
    case 618:
        v28 += (v28 % ((-1) | 1))
    default:
        v28 = (v28 * v28)
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 uint = genFn0()
        _ = fi29
        _ = v30
    }
    var v31 int32 = -1
    v2[32768] = 1
    var v32 float64 = 0.0
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        v34 := GenS3{F0: v28, F1: 0}
        if false {
            v32 += (0.0 / 1e10)
        } else {
            var v35 uint16 = 0
            v32 += (v32 / (v32 / v32))
            v35 = (v35 << ((2) & 63))
            _ = v35
        }
        for fi36 := 0; fi36 < 5; fi36++ {
            _ = len(v2)
            _ = fi36
        }
        _ = fi33
        _ = v34
    }
    v37 := GenS4{F0: v3}
    v2[-2147483648] = -746
    for fi38 := 0; fi38 < 19; fi38++ {
        for fi39 := 0; fi39 < 8; fi39++ {
            var v40 uint8 = 42
            var v41 float64 = v32
            var v42 bool = (v28 >= (v28 - (v28 ^ -32768)))
            _ = fi39
            _ = v40
            _ = v41
            _ = v42
        }
        v43 := GenS4{F0: v3}
        fmt.Println(v3, fi38)
        _ = fi38
        _ = v43
    }
    var v44 uint64 = 0
    v31 += v31
    for fi45 := 0; fi45 < 9; fi45++ {
        _ = len(v2)
        _ = fi45
    }
    if len(v1) > 0 {
        v1[0] = ("\n\t" + "hello")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v37
    _ = v44
}
