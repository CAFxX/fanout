package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int8 }
func genFn0(p0 int16, p1 string) int32 {
    if (0 <= 65535) {
        var v5 int = -2147483648
        _ = v5
    }
    p1 = p1
    switch "\n\t" {
    case "hello":
        if true {
            v6 := []int{}
            _ = v6
        }
    case "\n\t":
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    return 2147483647
}
func genFn1(p0 int64, p1 uint16) uint64 {
    fmt.Println(p1, p1)
    for fi7 := 0; fi7 < 9; fi7++ {
        if (true || (false && false)) {
            p1 += p1
            var v8 int = fi7
            if false {
                p0 += (^(p0))
                if true {
                    fmt.Println(p1)
                    var v9 float32 = (-2.25 * float32(1.5))
                    var v10 float64 = 0.1
                    _ = v10
                    _ = v9
                }
                switch (v8 & (v8 >> ((148) & 63))) {
                case -937:
                    if false {
                        if true {
                            if false {
                                for fi11 := 0; fi11 < 11; fi11++ {
                                    if false {
                                        v12 := []int{(-(v8)), v8, (v8 ^ 42)}
                                        p1 -= (p1 / ((304) | 1))
                                        v13 := GenS1{F0: "a", F1: 737}
                                        _ = v12
                                        _ = v13
                                    }
                                    _ = fi11
                                }
                                var v14 int32 = -1
                                _ = v14
                            }
                            var v15 float32 = 1.5
                            p1 = (p1 + p1)
                            _ = v15
                        }
                        if true {
                            var v16 int8 = -14
                            _ = v16
                        } else {
                            v17 := []int{(v8 >> ((868) & 63)), fi7, (v8 / ((598) | 1))}
                            _ = v17
                        }
                    } else {
                        v18 := []int{2147483647}
                        _ = v18
                    }
                case 1:
                    for fi19 := 0; fi19 < 4; fi19++ {
                        v20 := GenS1{F0: "", F1: 0}
                        v21 := []int{(fi19 % (((-(v8))) | 1)), fi7, (v8 - (v8 << ((1) & 63))), (-(fi7))}
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                        _ = len(v21)
                        var v22 bool = false
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                case -702:
                    p1 += 1
                default:
                    v8 -= 128
                }
            }
            _ = v8
        } else {
            var v23 bool = (false || false)
            var v24 int8 = 127
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 5; fi25++ {
            v26 := []int{fi7, -1, fi7, fi7}
            p1 += p1
            p0 = (p0 ^ (p0 / ((0) | 1)))
            _ = fi25
            _ = v26
        }
        p0 = (p0 + (p0 >> ((443) & 63)))
        _ = fi7
    }
    p0 -= (p0 * 332)
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    var v4 int = -1
    v4 = (v4 / ((0) | 1))
    for fi27 := 0; fi27 < 3; fi27++ {
        if false {
            v28 := GenS2{F0: 1, F1: 127}
            v29 := []int{}
            _ = v28
            _ = v29
        }
        v4 = (fi27 << ((440) & 63))
        _ = fi27
    }
    if true {
        fmt.Println(v4)
        v30 := []int{(v4 | v4), v4, (v4 % (((v4 - v4)) | 1)), -401}
        v4 += v4
        _ = v30
    } else {
        var v31 uint32 = 42
        for fi32 := 0; fi32 < 15; fi32++ {
            v33 := []int{fi32, v4, v4}
            v34 := GenS2{F0: 599, F1: 226}
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = v31
    }
    v4 -= 720
    v4 += -2147483648
    if (811 > 2147483647) {
        fmt.Println(v4)
        var v35 int = (v4 / ((444) | 1))
        _ = v35
    } else {
        v4 += v4
        v4 = v4
    }
    v4 -= v4
    if (!false) {
        v36 := GenS3{F0: 138, F1: 0}
        _ = v36
    } else {
        v4 -= (v4 >> ((0) & 63))
        v4 -= (v4 >> ((99) & 63))
        for fi37 := 0; fi37 < 4; fi37++ {
            v38 := GenS1{F0: "foo bar", F1: 0}
            _ = fi37
            _ = v38
        }
    }
    for fi39 := 0; fi39 < 17; fi39++ {
        v40 := GenS4{F0: 100}
        if ((fi39 * fi39) < (fi39 | 574)) {
            v4 += 32768
            v4 = 32767
            v4 -= (v4 * (fi39 & (v4 * 609)))
        }
        _ = fi39
        _ = v40
    }
    v4 -= (v4 / ((-867) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
