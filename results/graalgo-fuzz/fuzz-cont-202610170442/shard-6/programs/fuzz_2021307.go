package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint32; F1 uint; F2 int16 }
func genFn0(p0 uint16, p1 int8, p2 int) uint {
    v7 := []int{32768, p2, p2, (p2 * int(p0))}
    v8 := GenS1{F0: 4294967295, F1: "foo bar", F2: -882}
    v9 := GenS2{F0: false}
    var v10 int = (p2 % ((p2) | 1))
    p0 = (p0 * 0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 65535
}

func main() {
    v1 := []int{4294967295, 42}
    v2 := make(map[int]int)
    var v3 int16 = -32768
    v4 := make(map[uint32]uint32)
    var v5 float32 = float32(1.5)
    v6 := []int{-124, 1, -1}
    if true {
        for fi11 := 0; fi11 < 20; fi11++ {
            if false {
                if false {
                    for fi12 := 0; fi12 < 3; fi12++ {
                        if len(v6) > 0 {
                            v6[0] = -154
                        }
                        _ = fi12
                    }
                    v4[33] = 1
                    v13 := GenS3{F0: 275, F1: -13}
                    _ = v13
                } else {
                    if false {
                        v3 = (v3 - -94)
                    } else {
                        v3 = v3
                        if (true || (860 < 18446744073709551615)) {
                            switch fi11 {
                            case -1:
                                var v14 uint8 = 128
                                _ = v14
                            default:
                                if len(v1) > 0 {
                                    v1[0] = (fi11 | fi11)
                                }
                            }
                            v5 -= (float32((0.1 / float64(float32(0.1)))) + 0.1)
                            switch ("\n\t" + "foo bar") {
                            case "hello":
                                for fi15 := 0; fi15 < 17; fi15++ {
                                    var v16 bool = false
                                    _ = fi15
                                    _ = v16
                                }
                            case "a":
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            default:
                                fmt.Println(fi11, v5)
                            }
                        } else {
                            fmt.Println(fi11)
                            fmt.Println(fi11, v3)
                            fmt.Println(v3, v5)
                        }
                    }
                    v3 = v3
                }
            } else {
                if false {
                    if (v3 != (v3 | (v3 >> ((4294967295) & 63)))) {
                        v5 = (v5 * 1.5)
                    } else {
                        if false {
                            var v17 uint16 = 2
                            fmt.Println(v17, fi11, v17)
                            _ = v17
                        } else {
                            v5 += v5
                            v18 := GenS4{F0: 781, F1: 181, F2: v3}
                            _ = v18
                        }
                        v19 := GenS3{F0: 281, F1: 42}
                        _ = v19
                    }
                    if len(v1) > 0 {
                        v1[0] = fi11
                    }
                } else {
                    var v20 bool = (42 != 124)
                    _ = v20
                }
                v3 = v3
                var v21 string = "foo bar"
                _ = v21
            }
            _ = fi11
        }
    }
    fmt.Println(v3)
    v22 := GenS1{F0: 441, F1: "foo bar", F2: 339}
    if (42 <= 1) {
        if (!false) {
            var v23 int64 = 1
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v24 float32 = float32(float64(v5))
            _ = v23
            _ = v24
        }
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    var v25 uint16 = 155
    var v26 int32 = -2147483648
    v5 = v5
    v5 -= v5
    switch "foo bar" {
    case "hello":
        v27 := GenS3{F0: 65535, F1: 42}
        _ = v27
    case "foo bar":
        _ = len(v4)
    case "a":
        v6 = append(v6, -51)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
