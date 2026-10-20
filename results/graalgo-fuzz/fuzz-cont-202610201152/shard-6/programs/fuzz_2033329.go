package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int64 }
type GenS2 struct { F0 int32; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 int32 }
func genFn0(p0 uint64, p1 int8, p2 uint) int8 {
    for fi6 := 0; fi6 < 2; fi6++ {
        fmt.Println(p1)
        _ = fi6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: 42, F1: p2, F2: 9223372036854775807}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 uint16 = 886
    var v2 int64 = -1
    var v3 float64 = (1.5 + 1.5)
    var v4 uint32 = 4294967295
    var v5 string = ""
    v1 -= v1
    var v9 float32 = (0.1 - 1.5)
    switch int(v2) {
    case 127:
        switch -912 {
        case 42:
            v10 := GenS1{F0: v2, F1: 18, F2: -199}
            _ = v10
        default:
            if true {
                switch (("\n\t" + ("hello" + (("\n\t" + v5) + v5))) + ("foo bar" + v5)) {
                case "hello":
                    if ((815 <= v2) || (42 != 210)) {
                        if false {
                            var v11 int32 = 778
                            v4 += 222
                            _ = v11
                        } else {
                            v12 := GenS1{F0: (v2 / (((v2 & v2)) | 1)), F1: 42, F2: (v2 * (^(v2)))}
                            var v13 uint16 = (v1 % (((v1 & (v1 ^ (v1 / (((v1 >> ((18446744073709551615) & 63))) | 1))))) | 1))
                            var v14 uint64 = 4294967295
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                    }
                case "a":
                    for fi15 := 0; fi15 < 10; fi15++ {
                        if false {
                            fmt.Println(v5, fi15, v3)
                            for fi16 := 0; fi16 < 13; fi16++ {
                                if (v4 < v4) {
                                    switch (fi16 / ((fi16) | 1)) {
                                    case -2147483648:
                                        if false {
                                            v2 -= (v2 >> ((0) & 63))
                                        } else {
                                            v17 := []int{fi15}
                                            if ((!true) && true) {
                                                var v18 bool = ((v3 - (0.1 + v3)) < 1.5)
                                                _ = v18
                                            }
                                            _ = v17
                                        }
                                    }
                                    v1 -= v1
                                }
                                v19 := []int{}
                                _ = fi16
                                _ = v19
                            }
                            v20 := GenS2{F0: 449, F1: uint16(v4), F2: 0}
                            _ = v20
                        }
                        var v21 bool = (((v2 / ((v2) | 1)) <= (v2 / ((v2) | 1))) || (fi15 <= (fi15 * fi15)))
                        fmt.Println(v2, v2, v21)
                        _ = fi15
                        _ = v21
                    }
                case "":
                    var v22 float64 = 0.1
                    _ = v22
                default:
                    v23 := []int{}
                    _ = v23
                }
                v24 := GenS3{F0: 127}
                var v25 int8 = -128
                _ = v24
                _ = v25
            }
        }
    default:
        v26 := []int{-85, 32767}
        _ = v26
    }
    if (32767 <= 1000) {
        for fi27 := 0; fi27 < 7; fi27++ {
            v28 := []int{}
            v29 := GenS1{F0: 9223372036854775807, F1: 4294967295, F2: int64(fi27)}
            if len(v28) > 0 {
                v28[0] = (fi27 ^ fi27)
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        v30 := GenS3{F0: 166}
        _ = v30
    }
    if (v4 < v4) {
        if false {
            fmt.Println(v3, v1, v4)
            v31 := GenS1{F0: (v2 * int64(v4)), F1: 42, F2: (v2 << ((142) & 63))}
            var v32 float32 = v9
            _ = v31
            _ = v32
        }
        switch v5 {
        case "":
            if true {
                for fi33 := 0; fi33 < 12; fi33++ {
                    var v34 bool = (v1 >= v1)
                    var v35 uint8 = 234
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            }
        case "a":
            var v36 float32 = float32(float64(v9))
            _ = v36
        case "\n\t":
            v37 := GenS4{F0: 255, F1: 551}
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 1; fi38++ {
        v2 += (v2 - 42)
        _ = fi38
    }
    var v39 uint8 = 142
    v2 = (v2 | 0)
    if (1 != 128) {
        v1 = v1
    }
    var v40 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v9
}
