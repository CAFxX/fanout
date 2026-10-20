package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int64; F2 string }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int8 }
func genFn0() uint8 {
    v7 := GenS1{F0: -97}
    var v8 uint = 255
    for fi9 := 0; fi9 < 20; fi9++ {
        v8 -= (v8 << ((v8) & 63))
        v10 := GenS1{F0: 1}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 uint64 = 1
        _ = fi11
        _ = v12
    }
    var v13 string = "\n\t"
    _ = v13
    _ = v7
    _ = v8
    return 148
}
func genFn1(p0 uint64, p1 int64, p2 uint16) int16 {
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint32 = 2147483647
        v16 := GenS2{F0: (v15 + 286), F1: p1, F2: "a"}
        if true {
            if false {
                p1 -= -9223372036854775808
            } else {
                fmt.Println(fi14, v15, p2)
                for fi17 := 0; fi17 < 4; fi17++ {
                    v18 := GenS2{F0: v15, F1: 9223372036854775807, F2: ""}
                    v19 := GenS3{F0: (^(v15)), F1: (p0 & p0), F2: (v15 & (v15 % ((0) | 1)))}
                    switch "a" {
                    case "foo bar":
                        fmt.Println(fi17, fi14, fi14)
                    case "":
                        var v20 string = (("foo bar" + ("" + "\n\t")) + ("hello" + ("hello" + ("hello" + ("hello" + "")))))
                        _ = v20
                    case "a":
                        if false {
                            v21 := GenS1{F0: int8(v15)}
                            if false {
                                var v22 float64 = -2.25
                                var v23 uint8 = 255
                                switch 1 {
                                case 127:
                                    fmt.Println(fi17, v23)
                                case 32768:
                                    fmt.Println(p1, p0, v23)
                                case -138:
                                    var v24 string = "a"
                                    _ = v24
                                }
                                var v25 int8 = 127
                                v26 := []int{}
                                var v27 string = "a"
                                _ = v22
                                _ = v23
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                            _ = v21
                        }
                        fmt.Println(p0, fi14)
                        v28 := []int{32767}
                        _ = v28
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
        } else {
            for fi29 := 0; fi29 < 18; fi29++ {
                var v30 float64 = float64(-2.25)
                v31 := GenS4{F0: 1}
                switch fi14 {
                case -2147483648:
                    v15 -= v15
                }
                _ = fi29
                _ = v30
                _ = v31
            }
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    var v32 int32 = -764
    switch "a" {
    case "\n\t":
        switch -2147483648 {
        case 4294967295:
            p1 += (p1 | p1)
        default:
            v33 := []int{567, -510}
            _ = v33
        }
    case "foo bar":
        if (!false) {
            v34 := []int{4294967295, 128}
            _ = v34
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    return -461
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    var v4 string = "\n\t"
    v5 := make(map[int]uint32)
    var v6 uint8 = 2
    for fi35 := 0; fi35 < 14; fi35++ {
        v4 = v4
        var v36 int = fi35
        var v37 bool = true
        _ = fi35
        _ = v36
        _ = v37
    }
    v38 := []int{}
    v4 = v4
    v4 = v4
    _ = len(v38)
    delete(v3, "a")
    v4 = ("a" + ("a" + "a"))
    v4 = ("\n\t" + ("a" + (v4 + "foo bar")))
    v1[-9223372036854775808] = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
