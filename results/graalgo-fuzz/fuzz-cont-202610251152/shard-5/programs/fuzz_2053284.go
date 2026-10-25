package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int64 }
func genFn0() uint8 {
    v6 := GenS1{F0: 1, F1: 32768, F2: 1}
    var v7 int32 = 644
    fmt.Println(v7)
    v8 := []int{216}
    _ = v6
    _ = v7
    _ = v8
    return 154
}
func genFn1(p0 bool) int16 {
    v9 := []int{-1, 32767}
    v10 := GenS1{F0: 9223372036854775807, F1: -2147483648, F2: -1}
    for fi11 := 0; fi11 < 16; fi11++ {
        for fi12 := 0; fi12 < 1; fi12++ {
            fmt.Println(fi11, fi12)
            if p0 {
                var v13 int64 = -855
                p0 = (645 < 255)
                if p0 {
                    v13 -= v13
                    var v14 uint32 = 982
                    switch "" {
                    case "":
                        if (v14 != (v14 - (^(v14)))) {
                            switch ("hello" + "a") {
                            case "":
                                var v15 uint64 = 42
                                _ = v15
                            }
                        } else {
                            v13 -= -539
                            var v16 int64 = (v13 & v13)
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            _ = v16
                        }
                    case "a":
                        v9 = append(v9, (fi12 ^ (fi12 & fi11)))
                    case "foo bar":
                        switch ("hello" + "\n\t") {
                        case "foo bar":
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                        }
                    default:
                        switch (fi12 + fi12) {
                        case 32768:
                            var v17 uint64 = 2
                            _ = v17
                        case 4294967295:
                            var v18 int64 = (^(v13))
                            _ = v18
                        case 127:
                            if len(v9) > 0 {
                                v9[0] = fi11
                            }
                        default:
                            if (0.0 > 3.14159) {
                                var v19 string = "a"
                                _ = v19
                            }
                        }
                    }
                    _ = v14
                }
                _ = v13
            }
            _ = fi12
        }
        p0 = (p0 && (!false))
        _ = fi11
    }
    p0 = ((-1 > -9223372036854775808) || p0)
    if p0 {
        var v20 uint64 = 4294967295
        for fi21 := 0; fi21 < 11; fi21++ {
            for fi22 := 0; fi22 < 16; fi22++ {
                switch "foo bar" {
                case "a":
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                case "hello":
                    fmt.Println(p0, p0, p0)
                }
                v20 += (v20 / ((18446744073709551615) | 1))
                _ = fi22
            }
            var v23 int8 = 127
            _ = len(v9)
            _ = fi21
            _ = v23
        }
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = v20
    } else {
        p0 = (!(!false))
    }
    _ = p0
    _ = v10
    _ = v9
    return -96
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint32 = 238
    var v3 uint = 192
    v4 := make(map[uint32]uint32)
    var v5 uint = (v3 * (v3 % (((v3 / (((v3 / ((v3) | 1))) | 1))) | 1)))
    var v24 int = int(v5)
    if (v24 != (v24 % (((^(v24))) | 1))) {
        v3 += (^(v3))
    }
    fmt.Println(v2)
    fmt.Println(v5, v24)
    v25 := GenS1{F0: 9223372036854775807, F1: (^(v24)), F2: 1}
    fmt.Println(v1, v5)
    v26 := GenS1{F0: 1, F1: v24, F2: -9223372036854775808}
    v27 := GenS1{F0: 42, F1: int(v3), F2: -816}
    if false {
        v28 := []int{(v24 ^ (v24 % (((v24 >> (((v5 & (v5 + v5))) & 63))) | 1))), v24, -252}
        _ = v28
    } else {
        for fi29 := 0; fi29 < 15; fi29++ {
            for fi30 := 0; fi30 < 6; fi30++ {
                v2 = v2
                _ = fi30
            }
            var v31 float32 = 3.14159
            _ = fi29
            _ = v31
        }
        v32 := GenS1{F0: 794, F1: 286, F2: 1}
        _ = v32
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        v5 += uint(v2)
        _ = fi33
    }
    fmt.Println(v1, v3, v5)
    fmt.Println(v1)
    var v34 uint32 = 2147483647
    fmt.Println(v24)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
