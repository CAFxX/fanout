package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int16 }
type GenS2 struct { F0 uint; F1 int64; F2 int8 }
type GenS3 struct { F0 int16; F1 int; F2 uint64 }

func main() {
    v1 := []string{"\n\t", "a", "", "foo bar"}
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    v4 := []string{"foo bar", "\n\t"}
    var v5 int8 = -112
    var v6 int16 = 32767
    var v7 bool = false
    for fi8 := 0; fi8 < 5; fi8++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi9 := 0; fi9 < 8; fi9++ {
            if (v7 || v7) {
                if true {
                    v5 -= (v5 >> ((93) & 63))
                    _ = len(v2)
                }
            } else {
                v7 = v7
                fmt.Println(fi8, fi8, fi9)
            }
            _ = fi9
        }
        if v7 {
            var v10 float32 = (3.14159 * -2.25)
            _ = v10
        } else {
            _ = len(v2)
        }
        _ = fi8
    }
    var v11 int64 = 42
    v7 = (-956 >= 0)
    if (1 >= 42) {
        for fi12 := 0; fi12 < 18; fi12++ {
            switch "foo bar" {
            case "hello":
                fmt.Println(v11, v5)
            case "":
                var v13 uint = 2
                _ = v13
            }
            _ = fi12
        }
        v6 -= v6
    } else {
        var v14 bool = v7
        var v15 string = "hello"
        for fi16 := 0; fi16 < 4; fi16++ {
            if ((0.0 * -2.25) < float32(0.1)) {
                v17 := GenS1{F0: "foo bar", F1: v11, F2: 279}
                var v18 uint8 = 1
                fmt.Println(v6)
                _ = v17
                _ = v18
            }
            _ = fi16
        }
        _ = v14
        _ = v15
    }
    v6 = (^(v6))
    v19 := GenS1{F0: "", F1: 42, F2: (v6 >> ((18446744073709551615) & 63))}
    switch "a" {
    case "foo bar":
        v7 = v7
    default:
        if (v7 || ("foo bar" == "hello")) {
            var v20 bool = true
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        v22 := GenS1{F0: "", F1: v11, F2: v6}
        _ = fi21
        _ = v22
    }
    if (1 == 485) {
        for fi23 := 0; fi23 < 15; fi23++ {
            var v24 float64 = (((1.5 - (1e10 - 1e10)) + 0.1) / (-2.25 + -2.25))
            var v25 int = (-(fi23))
            fmt.Println(v6)
            _ = fi23
            _ = v24
            _ = v25
        }
        var v26 float64 = 3.14159
        for fi27 := 0; fi27 < 17; fi27++ {
            _ = len(v1)
            if (fi27 != fi27) {
                if (61 < 926) {
                    for fi28 := 0; fi28 < 11; fi28++ {
                        v29 := GenS1{F0: ("hello" + ((("foo bar" + (("" + "a") + ("hello" + ("" + "foo bar")))) + ("\n\t" + "\n\t")) + "a")), F1: (^(v11)), F2: (-(v6))}
                        _ = fi28
                        _ = v29
                    }
                } else {
                    v5 = (v5 ^ (v5 ^ v5))
                    for fi30 := 0; fi30 < 16; fi30++ {
                        v2[0] = fi27
                        delete(v2, 360)
                        switch fi27 {
                        case -715:
                            v26 -= v26
                        default:
                            switch fi27 {
                            case 128:
                                v5 -= (v5 % (((-(v5))) | 1))
                            default:
                                var v31 string = "\n\t"
                                _ = v31
                            }
                        }
                        _ = fi30
                    }
                    v32 := GenS2{F0: 65535, F1: (v11 % (((v11 << ((4294967295) & 63))) | 1)), F2: v5}
                    _ = v32
                }
                v11 = v11
            }
            _ = fi27
        }
        _ = v26
    }
    v33 := GenS3{F0: (^(v6)), F1: -990, F2: 107}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
