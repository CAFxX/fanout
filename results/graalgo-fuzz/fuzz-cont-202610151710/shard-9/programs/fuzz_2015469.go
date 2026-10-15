package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 uint; F1 int16; F2 bool }
type GenS5 struct { F0 uint64; F1 int32 }
func genFn0(p0 int64, p1 string) int64 {
    for fi7 := 0; fi7 < 2; fi7++ {
        if (42 != 1) {
            v8 := []int{(-(fi7))}
            _ = v8
        }
        _ = fi7
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 int8 = 1
    v2 := make(map[uint32]int)
    v3 := []uint32{268, 0}
    v4 := []int{int(v1), 4294967295, -9223372036854775808}
    var v5 int64 = -9223372036854775808
    v6 := make(map[string]int)
    v5 = genFn0((^(v5)), "a")
    v9 := GenS1{F0: 32767, F1: v1}
    v1 += int8(v5)
    if (9223372036854775807 > -425) {
        v1 = (v1 * -1)
    }
    var v10 uint = 11
    var v11 uint16 = 2
    var v12 int = int(v10)
    if (v12 == 233) {
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 int32 = -2147483648
            _ = len(v3)
            _ = fi13
            _ = v14
        }
    } else {
        v5 -= genFn0(42, "\n\t")
        if true {
            if true {
                fmt.Println(v10)
                for fi15 := 0; fi15 < 18; fi15++ {
                    v16 := GenS2{F0: (v5 & (v5 % ((v5) | 1))), F1: 1}
                    switch v12 {
                    case -2147483648:
                        var v17 float64 = (3.14159 - 0.1)
                        _ = v17
                    case 9223372036854775807:
                        var v18 bool = (751 >= 42)
                        _ = v18
                    }
                    if (!true) {
                        v10 -= uint(v11)
                        for fi19 := 0; fi19 < 8; fi19++ {
                            fmt.Println(v10, fi15)
                            v20 := GenS3{F0: 157, F1: true}
                            v2[430] = int(v1)
                            switch "\n\t" {
                            case "hello":
                                var v21 uint = (^(v10))
                                _ = v21
                            case "\n\t":
                                if false {
                                    v1 += 19
                                } else {
                                    v22 := GenS4{F0: (v10 >> (((v10 & (^(v10)))) & 63)), F1: 934, F2: false}
                                    _ = v22
                                }
                            case "a":
                                v23 := v6[(("foo bar" + "hello") + "a")]
                                _ = v23
                            }
                            _ = fi19
                            _ = v20
                        }
                    } else {
                        v5 -= int64(fi15)
                        var v24 int = (^(v12))
                        _ = v24
                    }
                    _ = fi15
                    _ = v16
                }
            } else {
                if (1e10 != (((-2.25 + 0.0) - 1e10) / (3.14159 / 1.5))) {
                    v25 := GenS3{F0: 4294967295, F1: (2 > 166)}
                    var v26 int64 = genFn0((v5 * (v5 * (v5 ^ 1))), "hello")
                    _ = v25
                    _ = v26
                } else {
                    fmt.Println(v11, v1)
                    v12 = (v12 >> ((65535) & 63))
                }
            }
            if (0 <= 4294967295) {
                v11 -= 42
            } else {
                var v27 bool = false
                _ = v27
            }
        } else {
            switch (^(v12)) {
            case 448:
                v28 := GenS5{F0: 2, F1: 0}
                _ = v28
            }
            var v29 float32 = 0.0
            var v30 uint = (v10 << ((v10) & 63))
            _ = v29
            _ = v30
        }
    }
    v31 := GenS2{F0: (v5 % (((v5 - (-(v5)))) | 1)), F1: 65535}
    var v32 uint16 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
