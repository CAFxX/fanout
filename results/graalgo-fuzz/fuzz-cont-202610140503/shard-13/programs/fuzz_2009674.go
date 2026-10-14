package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }

func main() {
    var v1 int16 = -1
    var v2 uint = 1
    v3 := make(map[int]int)
    var v4 uint64 = 0
    var v5 int64 = -9223372036854775808
    v5 += (v5 - (v5 >> (((v2 + (v2 >> ((1) & 63)))) & 63)))
    switch -181 {
    case 42:
        if (-2147483648 <= 0) {
            var v6 bool = false
            _ = v6
        } else {
            var v7 int32 = -1
            var v8 int16 = (v1 % ((v1) | 1))
            switch 232 {
            case -1:
                var v9 uint16 = 572
                _ = v9
            case 42:
                v10 := GenS1{F0: v2}
                _ = v10
            case 127:
                v4 += v4
            }
            _ = v7
            _ = v8
        }
    case 422:
        fmt.Println(v1)
    }
    v11 := []int{-866}
    v12 := v3[9223372036854775807]
    fmt.Println(v4, v4, v5)
    v12 -= (v12 << ((v2) & 63))
    fmt.Println(v12, v4)
    var v13 uint64 = 728
    var v14 float32 = ((3.14159 / 3.14159) - ((((((float32(-2.25) - 0.1) * 1.5) / 1.0) + 0.1) - 1.5) - -2.25))
    v3[(v12 << ((1) & 63))] = -916
    if true {
        if (58 > -128) {
            for fi15 := 0; fi15 < 18; fi15++ {
                delete(v3, 32768)
                _ = fi15
            }
            for fi16 := 0; fi16 < 10; fi16++ {
                v13 += (v4 % ((v13) | 1))
                _ = fi16
            }
        }
    } else {
        if (v4 >= v4) {
            fmt.Println(v13)
        }
        v12 = int(v4)
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    }
    if (false && (false && (-143 <= 2147483647))) {
        switch "foo bar" {
        case "foo bar":
            v4 = v13
        case "":
            if true {
                fmt.Println(v5)
                for fi17 := 0; fi17 < 9; fi17++ {
                    if false {
                        var v18 uint = (v2 & v2)
                        _ = v18
                    } else {
                        for fi19 := 0; fi19 < 18; fi19++ {
                            v20 := GenS1{F0: (v2 >> ((1) & 63))}
                            switch int(v2) {
                            case -2147483648:
                                v3[v12] = (v12 & (^(fi17)))
                            case -1:
                                v21 := GenS1{F0: v2}
                                _ = v21
                            }
                            _ = fi19
                            _ = v20
                        }
                        for fi22 := 0; fi22 < 20; fi22++ {
                            v1 += v1
                            v23 := GenS1{F0: (v2 / ((0) | 1))}
                            v4 += (v13 | (v4 | 0))
                            _ = fi22
                            _ = v23
                        }
                    }
                    v24 := GenS2{F0: (^(v2))}
                    _ = fi17
                    _ = v24
                }
                v13 = (v13 ^ 381)
            }
        default:
            if (0.1 <= -2.25) {
                v13 = (v4 / ((v4) | 1))
                var v25 uint8 = 8
                var v26 float64 = (0.0 / 3.14159)
                _ = v25
                _ = v26
            } else {
                if (0.0 <= -2.25) {
                    switch "a" {
                    case "a":
                        var v27 uint16 = 32767
                        _ = v27
                    }
                    v28 := v3[(v12 / (((-(v12))) | 1))]
                    var v29 uint = (v2 % ((v2) | 1))
                    _ = v28
                    _ = v29
                } else {
                    fmt.Println(v1, v14)
                    var v30 uint64 = 2
                    _ = v30
                }
                v5 += (v5 + (v5 / ((v5) | 1)))
            }
        }
    } else {
        var v31 int64 = (v5 % ((987) | 1))
        v12 += v12
        fmt.Println(v4)
        _ = v31
    }
    var v32 int16 = v1
    var v33 int = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
