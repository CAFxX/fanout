package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int16 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int64 }
func genFn0(p0 bool, p1 uint32, p2 int32) uint32 {
    p2 -= p2
    switch -1 {
    case 128:
        p2 = p2
    case 32768:
        v6 := GenS1{F0: -32768, F1: 1, F2: 277}
        _ = v6
    case 1:
        if true {
            for fi7 := 0; fi7 < 12; fi7++ {
                v8 := []int{}
                p0 = (-183 >= 1)
                v9 := GenS2{F0: 118, F1: -680}
                _ = fi7
                _ = v8
                _ = v9
            }
            var v10 float32 = 0.1
            switch 32768 {
            case -2147483648:
                for fi11 := 0; fi11 < 5; fi11++ {
                    if p0 {
                        v12 := GenS2{F0: 127, F1: (^(fi11))}
                        _ = v12
                    } else {
                        switch (fi11 * (fi11 / ((976) | 1))) {
                        case 163:
                            if (p0 && false) {
                                var v13 float32 = v10
                                fmt.Println(p2, v10, fi11)
                                _ = v13
                            }
                        default:
                            v14 := []int{fi11, (fi11 + int(p1))}
                            _ = v14
                        }
                        p2 = p2
                        v15 := []int{(fi11 + (fi11 - (fi11 + fi11))), (fi11 ^ (fi11 & (fi11 * fi11))), fi11, fi11}
                        _ = v15
                    }
                    if true {
                        p1 = (^(p1))
                        for fi16 := 0; fi16 < 14; fi16++ {
                            if p0 {
                                v17 := GenS3{F0: 9223372036854775807}
                                _ = v17
                            }
                            fmt.Println(fi11, p2)
                            var v18 int64 = -185
                            _ = fi16
                            _ = v18
                        }
                        v19 := []int{(fi11 & fi11), fi11, 127, fi11}
                        _ = v19
                    } else {
                        var v20 int16 = 1
                        _ = v20
                    }
                    p0 = false
                    _ = fi11
                }
            case -3:
                v10 += v10
            default:
                fmt.Println(p2)
            }
            _ = v10
        }
    }
    v21 := []int{-833, 4294967295, -128, -790}
    if ("" >= "foo bar") {
        p2 = (-(p2))
        if ((float32(3.14159) != 3.14159) && (!true)) {
            switch 154 {
            case 1:
                if (p2 >= -503) {
                    var v22 int32 = p2
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                    _ = v22
                } else {
                    v23 := GenS2{F0: 127, F1: 2147483647}
                    _ = v23
                }
            }
            switch 127 {
            case 127:
                p0 = (!false)
            case 9223372036854775807:
                switch 0 {
                case 988:
                    v21 = append(v21, -1)
                }
            default:
                v24 := GenS1{F0: -985, F1: 18446744073709551615, F2: -32768}
                _ = v24
            }
        } else {
            p0 = p0
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return 65535
}
func genFn1() int32 {
    var v25 int = -2
    v26 := []int{}
    var v27 bool = (42 > 2147483647)
    _ = v25
    _ = v26
    _ = v27
    return 2147483647
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]int)
    var v3 string = ""
    v4 := []uint32{376, 1, 42, 4294967295}
    v5 := make(map[string]uint32)
    v28 := GenS1{F0: 0, F1: 0, F2: -625}
    v3 = ""
    for fi29 := 0; fi29 < 3; fi29++ {
        v3 = (v3 + "foo bar")
        _ = fi29
    }
    v30 := GenS1{F0: 523, F1: 18446744073709551615, F2: -1}
    for fi31 := 0; fi31 < 19; fi31++ {
        var v32 uint = 0
        v3 = v3
        var v33 int16 = 26
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println(v3)
    v34 := GenS2{F0: 100, F1: 0}
    var v35 uint64 = 132
    if (true || false) {
        v35 += (v35 << ((4294967295) & 63))
    } else {
        v36 := v5[(((v3 + v3) + (v3 + v3)) + v3)]
        var v37 int32 = 1
        _ = v36
        _ = v37
    }
    v38 := v1[("a" + (v3 + ((v3 + (((v3 + "a") + ((("hello" + v3) + v3) + "a")) + ("" + "a"))) + v3)))]
    v39 := GenS2{F0: -128, F1: 869}
    var v40 string = (v3 + v3)
    v41 := v2[0]
    v3 = (v40 + (v40 + v3))
    fmt.Println(v35)
    _ = len(v4)
    for fi42 := 0; fi42 < 13; fi42++ {
        if len(v4) > 0 {
            v4[0] = v38
        }
        v43 := GenS2{F0: 0, F1: (^(fi42))}
        _ = fi42
        _ = v43
    }
    if len(v4) > 0 {
        v4[0] = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
