package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int8; F1 int8; F2 uint }
func genFn0(p0 bool, p1 int16, p2 uint8) uint32 {
    v6 := GenS1{F0: 2}
    p0 = (p0 && (1 <= 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2147483647
}
func genFn1() uint8 {
    if false {
        var v7 int64 = -830
        var v8 uint16 = 1
        fmt.Println(v8)
        _ = v7
        _ = v8
    } else {
        var v9 int32 = 0
        if (!(!(!false))) {
            v9 = (v9 << ((536) & 63))
        } else {
            v10 := []int{2147483647, 2147483647, 9223372036854775807}
            _ = v10
        }
        var v11 int8 = 1
        _ = v11
        _ = v9
    }
    v12 := []int{32768, 856}
    v13 := GenS1{F0: 1}
    if true {
        var v14 int = -2147483648
        v14 = v14
        _ = v14
    }
    _ = v12
    _ = v13
    return 128
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]string)
    var v5 uint64 = 4294967295
    v15 := GenS2{F0: uint32(v5)}
    v16 := GenS1{F0: 1}
    if false {
        v17 := GenS3{F0: 666, F1: 1}
        _ = v17
    }
    _ = len(v2)
    var v18 float32 = 1e10
    v19 := GenS2{F0: 4294967295}
    if (-2.25 == 1.5) {
        v18 = ((v18 * v18) + float32((float64(0.0) * -2.25)))
    }
    v20 := []int{127, -9223372036854775808, 334}
    switch "hello" {
    case "foo bar":
        v5 = 2
    case "hello":
        for fi21 := 0; fi21 < 11; fi21++ {
            delete(v2, "")
            _ = fi21
        }
    case "a":
        switch 127 {
        case 4294967295:
            for fi22 := 0; fi22 < 4; fi22++ {
                for fi23 := 0; fi23 < 15; fi23++ {
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    _ = fi23
                }
                switch 32768 {
                case 702:
                    if (1 > 1000) {
                        var v24 int64 = 645
                        var v25 int16 = -1
                        _ = v24
                        _ = v25
                    } else {
                        if true {
                            v26 := GenS4{F0: -1, F1: 789, F2: v5}
                            v18 -= float32(0.0)
                            if true {
                                var v27 float64 = 1.5
                                if (-104 != 127) {
                                    for fi28 := 0; fi28 < 16; fi28++ {
                                        v27 = v27
                                        _ = fi28
                                    }
                                    v5 += v5
                                } else {
                                    switch ("hello" + "a") {
                                    case "":
                                        v20 = append(v20, fi22)
                                    case "a":
                                        v18 -= float32(v27)
                                    default:
                                        if len(v20) > 0 {
                                            v20[0] = fi22
                                        }
                                    }
                                }
                                var v29 int32 = 609
                                _ = v27
                                _ = v29
                            }
                            _ = v26
                        } else {
                            var v30 uint8 = 0
                            if false {
                                v31 := GenS5{F0: -122, F1: -128, F2: 1}
                                var v32 string = (("hello" + ("foo bar" + "hello")) + "")
                                _ = v31
                                _ = v32
                            }
                            fmt.Println(v30)
                            _ = v30
                        }
                        v33 := GenS3{F0: 1, F1: 42}
                        if ((600 <= 42) && true) {
                            switch ("" + "\n\t") {
                            case "":
                                v18 -= (v18 * v18)
                            }
                        }
                        _ = v33
                    }
                case 846:
                    v18 = (v18 * ((v18 + v18) + 0.1))
                default:
                    switch (-(fi22)) {
                    case 9223372036854775807:
                        v34 := GenS2{F0: 2147483648}
                        _ = v34
                    case 127:
                        v18 -= v18
                    case 128:
                        var v35 bool = false
                        _ = v35
                    default:
                        for fi36 := 0; fi36 < 10; fi36++ {
                            var v37 float32 = v18
                            _ = fi36
                            _ = v37
                        }
                    }
                }
                _ = fi22
            }
        case -595:
            var v38 uint8 = 255
            _ = v38
        default:
            var v39 int8 = 127
            _ = v39
        }
    }
    var v40 int = -286
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
