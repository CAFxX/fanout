package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint8 }
func genFn0(p0 uint, p1 int16) uint8 {
    if true {
        var v6 int8 = 1
        switch "foo bar" {
        case "":
            v7 := []int{-963, 1, 32768, 32767}
            _ = v7
        default:
            fmt.Println(v6, p1, p0)
        }
        p1 = p1
        _ = v6
    }
    p0 = (p0 << ((4294967295) & 63))
    v8 := []int{-2147483648}
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v8
    return 42
}
func genFn1(p0 bool) uint {
    v9 := []int{}
    p0 = (((0.0 + (3.14159 + (float64((0.0 * float32(1e10))) - -2.25))) + 0.0) != 1e10)
    var v10 int8 = -99
    _ = p0
    _ = v10
    _ = v9
    return 232
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = -1
    v3 := make(map[int]string)
    v4 := []uint32{310}
    var v5 int = 0
    v2 = (-(v2))
    _ = len(v3)
    for fi11 := 0; fi11 < 4; fi11++ {
        switch v5 {
        case -2147483648:
            var v12 int32 = -825
            _ = v12
        case 42:
            v5 = 0
        case -9223372036854775808:
            switch (("hello" + (("hello" + "a") + ("\n\t" + ""))) + "a") {
            case "":
                var v13 uint32 = 440
                _ = v13
            case "foo bar":
                v14 := GenS1{F0: 6, F1: "foo bar", F2: v5}
                _ = v14
            case "\n\t":
                if true {
                    if (!false) {
                        v15 := GenS1{F0: 2, F1: "foo bar", F2: int(v2)}
                        fmt.Println(v5, fi11, fi11)
                        if true {
                            switch 9223372036854775807 {
                            case 1:
                                var v16 float32 = (-2.25 * 0.0)
                                _ = v16
                            case 127:
                                var v17 int8 = 45
                                _ = v17
                            }
                        } else {
                            v5 -= 0
                        }
                        _ = v15
                    }
                }
            default:
                switch v5 {
                case 128:
                    v5 -= v5
                case 9223372036854775807:
                    var v18 string = "hello"
                    _ = v18
                case 0:
                    _ = len(v4)
                default:
                    _ = len(v3)
                }
            }
        }
        v19 := GenS1{F0: 2, F1: "hello", F2: fi11}
        _ = fi11
        _ = v19
    }
    fmt.Println(v5, v5)
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 bool = true
        v4 = append(v4, 182)
        _ = fi20
        _ = v21
    }
    var v22 int = v5
    switch "a" {
    case "foo bar":
        var v23 int32 = -385
        _ = v23
    case "hello":
        if false {
            if len(v4) > 0 {
                v4[0] = 2
            }
        } else {
            var v24 int8 = -99
            v5 += v22
            _ = v24
        }
    case "a":
        var v25 uint16 = 255
        _ = v25
    default:
        if false {
            v22 = v22
            var v26 int8 = 42
            _ = v26
        } else {
            for fi27 := 0; fi27 < 9; fi27++ {
                v28 := GenS1{F0: 2, F1: "foo bar", F2: v5}
                if false {
                    var v29 uint32 = 42
                    if (!((-1 < -1) && true)) {
                        v5 = -2147483648
                        v30 := v3[-9223372036854775808]
                        var v31 uint32 = 2
                        _ = v30
                        _ = v31
                    }
                    v32 := GenS2{F0: uint(v29), F1: -93}
                    _ = v29
                    _ = v32
                }
                _ = fi27
                _ = v28
            }
            v33 := v3[(v22 / (((-(v5))) | 1))]
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        fmt.Println(v5)
        v22 = (-(fi34))
        _ = fi34
    }
    var v35 string = "\n\t"
    fmt.Println(v5, v2, v22)
    v36 := GenS3{F0: (v5 - 589)}
    v37 := GenS3{F0: (v5 & (v5 & (v22 >> ((2) & 63))))}
    for fi38 := 0; fi38 < 10; fi38++ {
        var v39 int64 = 1
        _ = fi38
        _ = v39
    }
    for fi40 := 0; fi40 < 14; fi40++ {
        v35 = v35
        _ = fi40
    }
    v41 := GenS1{F0: 229, F1: "a", F2: v22}
    var v42 float64 = 1e10
    switch "a" {
    case "a":
        v43 := GenS4{F0: genFn0(726, 605), F1: 0, F2: 138}
        _ = v43
    case "":
        v44 := GenS1{F0: 2, F1: "a", F2: 9223372036854775807}
        _ = v44
    case "hello":
        var v45 bool = false
        _ = v45
    }
    var v46 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v46
    _ = v5
}
