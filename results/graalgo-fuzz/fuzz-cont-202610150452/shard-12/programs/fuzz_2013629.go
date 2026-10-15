package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint; F2 int32 }
type GenS3 struct { F0 uint16; F1 int32; F2 int32 }
func genFn0(p0 int16) int {
    fmt.Println(p0)
    v6 := []int{32767, -252}
    fmt.Println(p0)
    _ = p0
    _ = v6
    return -1
}
func genFn1(p0 int16, p1 uint8, p2 int16) uint64 {
    var v7 uint64 = 4294967295
    v8 := []int{-1}
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return v7
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    v3 := []string{("" + "\n\t"), "hello", "\n\t", "foo bar"}
    var v4 uint8 = 2
    v5 := []int{}
    if false {
        for fi9 := 0; fi9 < 3; fi9++ {
            v4 = 127
            _ = fi9
        }
        v4 = v4
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := GenS1{F0: 2}
            var v12 float64 = 0.0
            for fi13 := 0; fi13 < 18; fi13++ {
                if len(v5) > 0 {
                    v5[0] = 9223372036854775807
                }
                var v14 uint8 = 127
                _ = len(v5)
                switch genFn0(-863) {
                case 42:
                    v14 = 0
                }
                v15 := GenS2{F0: 790, F1: 42, F2: -693}
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    fmt.Println(v4)
    if true {
        v4 = (v4 << ((255) & 63))
    } else {
        var v16 bool = false
        _ = v16
    }
    v4 = v4
    if false {
        var v17 float64 = (float64(1e10) + 0.1)
        v18 := GenS3{F0: 32767, F1: 0, F2: 453}
        _ = v17
        _ = v18
    }
    switch 0 {
    case 1:
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    case 376:
        v19 := GenS2{F0: 576, F1: 65535, F2: 42}
        _ = v19
    default:
        for fi20 := 0; fi20 < 13; fi20++ {
            switch (("\n\t" + "hello") + "foo bar") {
            case "\n\t":
                switch (fi20 | (fi20 | (fi20 >> ((2) & 63)))) {
                case 416:
                    if true {
                        v5 = append(v5, (fi20 << ((65535) & 63)))
                        v4 = (v4 << ((634) & 63))
                    }
                case -9223372036854775808:
                    v4 = 128
                case 0:
                    var v21 uint32 = 42
                    _ = v21
                default:
                    if true {
                        var v22 uint8 = v4
                        var v23 int = (fi20 + (fi20 | (fi20 | (fi20 - (fi20 & fi20)))))
                        _ = v22
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 12; fi24++ {
                            delete(v1, 510)
                            var v25 uint8 = (v4 ^ 0)
                            if false {
                                v26 := GenS1{F0: 1}
                                _ = v26
                            } else {
                                v25 -= 138
                                delete(v2, "hello")
                            }
                            v2[("foo bar" + "\n\t")] = ("" + ("" + "a"))
                            if ((false && (!(0 != 646))) && (1 >= 4294967295)) {
                                v27 := GenS3{F0: 255, F1: 2147483647, F2: -589}
                                fmt.Println(fi20, fi24, v25)
                                fmt.Println(fi20, v25, fi20)
                                _ = v27
                            }
                            switch fi24 {
                            case -939:
                                var v28 uint64 = 42
                                _ = v28
                            case 307:
                                var v29 float64 = float64(1e10)
                                _ = v29
                            }
                            _ = fi24
                            _ = v25
                        }
                        if false {
                            v4 += v4
                        } else {
                            v4 = (v4 + v4)
                            _ = len(v2)
                        }
                    }
                }
            case "hello":
                var v30 float64 = -2.25
                _ = v30
            }
            _ = fi20
        }
    }
    if len(v3) > 0 {
        v3[0] = ("\n\t" + "\n\t")
    }
    for fi31 := 0; fi31 < 7; fi31++ {
        v4 = 64
        _ = fi31
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        v4 = (^(v4))
        var v33 uint16 = 65535
        _ = fi32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
