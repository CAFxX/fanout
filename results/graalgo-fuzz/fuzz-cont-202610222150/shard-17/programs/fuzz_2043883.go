package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 int16; F2 int16 }
type GenS5 struct { F0 uint64; F1 uint }

func main() {
    var v1 int = 2147483647
    v2 := make(map[uint32]bool)
    var v3 uint = 42
    v4 := []int{4294967295}
    var v5 uint32 = 65535
    v6 := make(map[int]string)
    v7 := GenS1{F0: "hello", F1: (v3 >> (((v3 & 5)) & 63)), F2: 1}
    fmt.Println(v1, v1, v1)
    var v8 int16 = -32768
    var v9 int = (v1 + v1)
    v2[v5] = (-9223372036854775808 <= 0)
    var v10 bool = true
    if len(v4) > 0 {
        v4[0] = 128
    }
    fmt.Println(v9, v9, v9)
    v8 = (v8 << (((v3 % ((796) | 1))) & 63))
    v3 = (v3 - 0)
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            switch "foo bar" {
            case "hello":
                v10 = true
            case "a":
                var v13 int16 = -1
                _ = v13
            default:
                var v14 int32 = -1
                v14 -= v14
                _ = v14
            }
            _ = fi12
        }
        v2[v5] = v10
        _ = fi11
    }
    if false {
        for fi15 := 0; fi15 < 6; fi15++ {
            fmt.Println(v9, v5)
            _ = fi15
        }
        v4 = append(v4, (v9 * (v9 | v1)))
        v9 += -308
    }
    switch "foo bar" {
    case "\n\t":
        v1 = 42
    case "foo bar":
        var v16 uint64 = 18446744073709551615
        _ = v16
    case "hello":
        for fi17 := 0; fi17 < 16; fi17++ {
            if ((v10 && (2 < 4294967295)) && (v8 > v8)) {
                var v18 uint64 = 0
                v19 := GenS2{F0: (v8 % ((v8) | 1))}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v18
                _ = v19
            }
            if (v10 || v10) {
                if false {
                    v20 := GenS1{F0: "hello", F1: v3, F2: 141}
                    _ = v20
                } else {
                    v21 := v2[v5]
                    _ = v21
                }
                for fi22 := 0; fi22 < 1; fi22++ {
                    if (true || ("hello" > "foo bar")) {
                        if len(v4) > 0 {
                            v4[0] = -9223372036854775808
                        }
                    } else {
                        fmt.Println(fi17, v10, fi22)
                        v23 := GenS3{F0: 898}
                        v24 := GenS1{F0: "foo bar", F1: (v3 * v3), F2: 42}
                        _ = v23
                        _ = v24
                    }
                    var v25 bool = false
                    var v26 int8 = 0
                    _ = fi22
                    _ = v25
                    _ = v26
                }
                v27 := GenS4{F0: -9223372036854775808, F1: (v8 | (v8 / ((v8) | 1))), F2: (^(v8))}
                _ = v27
            }
            _ = fi17
        }
    default:
        if (v10 || v10) {
            v4 = append(v4, (v9 * (v1 - (-(v1)))))
            v8 = v8
        } else {
            v4 = append(v4, v9)
        }
    }
    v28 := GenS5{F0: 2, F1: (v3 & (^(v3)))}
    v3 = v3
    var v29 string = ""
    v30 := GenS1{F0: "\n\t", F1: v3, F2: int32(v5)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
