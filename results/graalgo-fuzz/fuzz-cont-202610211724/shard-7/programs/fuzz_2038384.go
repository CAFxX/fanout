package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int8; F2 string }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int) bool {
    var v4 uint32 = 65535
    v4 = 0
    _ = p0
    _ = v4
    return true
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    if (!true) {
        var v5 uint = 1
        _ = v5
    } else {
        v6 := GenS1{F0: "foo bar", F1: 128, F2: -327}
        var v7 uint8 = 2
        var v8 int32 = -699
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 float32 = 1.5
    switch 32767 {
    case -416:
        if true {
            v10 := v2[1]
            v9 -= (((v9 * (v9 + ((v9 * 1e10) * ((v9 / ((v9 * (v9 * v9)) - (((3.14159 - float32(float64(v9))) - 0.1) * ((v9 + v9) / (v9 / (0.1 / (1.5 + v9))))))) / v9)))) + float32(1.5)) / v9)
            var v11 int16 = 42
            _ = v10
            _ = v11
        } else {
            if (true && (("foo bar" + "hello") == "")) {
                v9 = (3.14159 * ((v9 + (v9 - ((v9 / (float32(1.5) / (-2.25 / v9))) / v9))) - (float32(3.14159) - 1e10)))
                switch (("" + "") + "foo bar") {
                case "a":
                    v12 := []int{32768}
                    _ = v12
                }
            } else {
                for fi13 := 0; fi13 < 12; fi13++ {
                    switch fi13 {
                    case 9223372036854775807:
                        var v14 float32 = (0.0 * 1e10)
                        _ = v14
                    case 304:
                        if true {
                            v9 = float32((-2.25 * 0.0))
                            v9 += v9
                        } else {
                            if (false || (0.0 < v9)) {
                                _ = len(v2)
                                for fi15 := 0; fi15 < 18; fi15++ {
                                    v16 := GenS2{F0: 127}
                                    _ = fi15
                                    _ = v16
                                }
                                v17 := v3[fi13]
                                _ = v17
                            }
                            v9 = ((float32(1e10) - (float32(-2.25) * v9)) - (v9 / (v9 + (0.1 * 0.0))))
                            v9 = v9
                        }
                    }
                    _ = fi13
                }
            }
            v9 -= v9
        }
    case 128:
        switch -9223372036854775808 {
        case 468:
            v9 = v9
        case -367:
            if true {
                if (!(109 != 109)) {
                    v18 := []int{213, 2147483647, 127, 42}
                    _ = v18
                }
            } else {
                for fi19 := 0; fi19 < 15; fi19++ {
                    v20 := []int{(fi19 + fi19)}
                    _ = fi19
                    _ = v20
                }
            }
        case 705:
            fmt.Println(v9)
        default:
            for fi21 := 0; fi21 < 11; fi21++ {
                fmt.Println(fi21, v9)
                for fi22 := 0; fi22 < 14; fi22++ {
                    fmt.Println(v9, fi22, fi21)
                    v23 := GenS3{F0: 33, F1: 1, F2: ("foo bar" + "hello")}
                    v9 = float32(1e10)
                    switch "foo bar" {
                    case "hello":
                        v9 -= v9
                    case "a":
                        v24 := GenS3{F0: 81, F1: 100, F2: ("\n\t" + (("a" + ("a" + "\n\t")) + ("hello" + ("\n\t" + "hello"))))}
                        _ = v24
                    }
                    _ = fi22
                    _ = v23
                }
                _ = fi21
            }
        }
    case 2147483647:
        for fi25 := 0; fi25 < 11; fi25++ {
            fmt.Println(fi25, fi25)
            var v26 int32 = -1
            _ = fi25
            _ = v26
        }
    default:
        for fi27 := 0; fi27 < 16; fi27++ {
            v28 := []int{(fi27 / (((fi27 & (fi27 + fi27))) | 1)), (fi27 >> ((327) & 63))}
            if false {
                v9 += ((((v9 * 0.0) * (float32((1.5 / 0.1)) + v9)) / (1.5 / 1.5)) / (-2.25 / 1.5))
                var v29 int = fi27
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
                _ = v29
            }
            v30 := GenS4{F0: 2147483647}
            _ = fi27
            _ = v28
            _ = v30
        }
    }
    v9 += 1e10
    for fi31 := 0; fi31 < 13; fi31++ {
        v2[(fi31 >> ((71) & 63))] = 2147483648
        _ = fi31
    }
    var v32 float64 = 1.5
    for fi33 := 0; fi33 < 2; fi33++ {
        fmt.Println(fi33)
        v34 := []int{(fi33 ^ (-(fi33))), 566, fi33}
        v32 -= v32
        _ = fi33
        _ = v34
    }
    var v35 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v9
}
