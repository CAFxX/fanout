package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
func genFn0(p0 int16, p1 uint16, p2 bool) uint8 {
    v7 := GenS1{F0: 255}
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            if (fi8 > 210) {
                if p2 {
                    if p2 {
                        p0 = (p0 / ((p0) | 1))
                        if p2 {
                            v10 := []int{-818, 0, fi8, (-(fi9))}
                            var v11 uint8 = 140
                            if ((p1 << ((42) & 63)) <= p1) {
                                switch (fi8 - fi8) {
                                case 128:
                                    v12 := GenS1{F0: (^(v11))}
                                    _ = v12
                                case 4294967295:
                                    p1 = p1
                                }
                                if p2 {
                                    v11 -= (v11 >> ((65535) & 63))
                                    v10 = append(v10, -2147483648)
                                    v13 := GenS1{F0: 2}
                                    _ = v13
                                }
                            }
                            switch fi9 {
                            case 241:
                                v14 := GenS2{F0: ""}
                                _ = v14
                            case 2147483647:
                                fmt.Println(p0, p2)
                            case 73:
                                v15 := GenS2{F0: "\n\t"}
                                _ = v15
                            }
                            fmt.Println(fi8)
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            _ = v10
                            _ = v11
                        } else {
                            v16 := []int{}
                            _ = v16
                        }
                        if ("" > ("a" + "hello")) {
                            v17 := []int{}
                            _ = v17
                        } else {
                            switch fi9 {
                            case -716:
                                v18 := GenS2{F0: "hello"}
                                _ = v18
                            case 4294967295:
                                var v19 int32 = -936
                                switch (fi8 + 0) {
                                case -182:
                                    if (p2 && p2) {
                                        var v20 uint32 = 360
                                        fmt.Println(p0, p1)
                                        _ = v20
                                    }
                                case 128:
                                    var v21 int = fi8
                                    _ = v21
                                default:
                                    var v22 float32 = (((3.14159 - (0.0 / 0.1)) - 0.0) - 0.1)
                                    _ = v22
                                }
                                _ = v19
                            case -559:
                                p0 = p0
                            default:
                                p2 = ((1.5 / -2.25) != float32(float64(-2.25)))
                            }
                            if false {
                                var v23 uint32 = 805
                                v23 += (v23 ^ (^(v23)))
                                _ = v23
                            }
                            p0 = (p0 >> ((1) & 63))
                        }
                    }
                    switch 32768 {
                    case -1:
                        v24 := GenS3{F0: -1}
                        _ = v24
                    }
                    v25 := GenS1{F0: 255}
                    fmt.Println(fi8, fi9, fi9)
                    _ = v25
                }
                fmt.Println(p2)
            }
            _ = fi9
        }
        _ = fi8
    }
    p2 = false
    var v26 uint = 4294967295
    v27 := []int{32768, -928, 128, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    _ = v7
    return 128
}
func genFn1(p0 bool) string {
    p0 = (!(p0 && p0))
    v28 := []int{0, -395}
    _ = p0
    _ = v28
    return ((("a" + "foo bar") + "\n\t") + ("foo bar" + "hello"))
}

func main() {
    var v1 int16 = 0
    v2 := make(map[string]uint32)
    var v3 bool = (!true)
    var v4 uint32 = 4294967295
    v5 := make(map[int]bool)
    v6 := make(map[int]uint32)
    for fi29 := 0; fi29 < 11; fi29++ {
        _ = len(v6)
        if (v4 > (v4 * v4)) {
            v1 = 32767
        } else {
            fmt.Println(v1)
            if v3 {
                v30 := GenS1{F0: 212}
                fmt.Println(v3, v4, v1)
                var v31 uint8 = genFn0((v1 / (((-(v1))) | 1)), 255, (("hello" + "a") < ("foo bar" + "foo bar")))
                _ = v30
                _ = v31
            }
            var v32 bool = (42 > 728)
            _ = v32
        }
        v3 = false
        _ = fi29
    }
    v33 := GenS2{F0: ""}
    for fi34 := 0; fi34 < 2; fi34++ {
        _ = len(v2)
        v3 = ((false || v3) || (!(!false)))
        v1 += v1
        _ = fi34
    }
    delete(v6, 374)
    for fi35 := 0; fi35 < 13; fi35++ {
        var v36 uint16 = 39
        var v37 uint16 = v36
        _ = fi35
        _ = v36
        _ = v37
    }
    v4 += v4
    v3 = false
    for fi38 := 0; fi38 < 9; fi38++ {
        v1 = (v1 | (v1 >> ((18446744073709551615) & 63)))
        v39 := v6[fi38]
        _ = fi38
        _ = v39
    }
    v40 := GenS3{F0: 9223372036854775807}
    var v41 float32 = float32(((0.0 + (1.5 * 1e10)) + (0.0 * -2.25)))
    v42 := GenS1{F0: 255}
    var v43 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
