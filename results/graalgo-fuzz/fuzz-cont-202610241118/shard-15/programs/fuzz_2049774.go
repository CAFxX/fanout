package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint32; F1 string; F2 int16 }
func genFn0(p0 uint8) int8 {
    v4 := GenS1{F0: 511}
    v5 := GenS1{F0: 2}
    if false {
        v6 := GenS2{F0: true, F1: 458}
        v7 := []int{32767, -1}
        v8 := GenS1{F0: 28}
        _ = v6
        _ = v7
        _ = v8
    } else {
        var v9 int64 = 42
        fmt.Println(v9, v9)
        _ = v9
    }
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 uint8, p1 string, p2 bool) int32 {
    p1 = "foo bar"
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 int16 = -1
        _ = fi10
        _ = v11
    }
    var v12 int64 = -9223372036854775808
    var v13 bool = (1e10 <= (1.5 * 1e10))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 850
}

func main() {
    v1 := []string{}
    var v2 uint16 = 284
    v3 := make(map[uint32]string)
    switch ("" + "hello") {
    case "":
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 float32 = (0.0 - 3.14159)
            if true {
                if false {
                    var v16 int16 = -584
                    _ = v16
                } else {
                    fmt.Println(v2, fi14, v2)
                }
                fmt.Println(v2, fi14)
                var v17 bool = (v2 != (v2 ^ (v2 << ((893) & 63))))
                _ = v17
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi14
            _ = v15
        }
    case "\n\t":
        for fi18 := 0; fi18 < 11; fi18++ {
            v19 := GenS3{F0: 500, F1: 32767}
            _ = fi18
            _ = v19
        }
    default:
        v2 += (v2 - (v2 / ((v2) | 1)))
    }
    v2 -= 255
    var v20 uint64 = 42
    fmt.Println(v20, v20)
    var v21 int = 920
    v2 = 42
    _ = len(v3)
    v22 := GenS3{F0: 18446744073709551615, F1: 1000}
    v20 = (v20 | v20)
    _ = len(v3)
    v23 := GenS4{F0: uint32(v2), F1: ("foo bar" + (("\n\t" + ("a" + "a")) + ((("a" + "hello") + "") + "foo bar"))), F2: 1}
    switch v21 {
    case -2147483648:
        v21 = v21
    case -9223372036854775808:
        for fi24 := 0; fi24 < 8; fi24++ {
            if false {
                var v25 int16 = -502
                _ = v25
            }
            for fi26 := 0; fi26 < 4; fi26++ {
                if true {
                    v27 := GenS1{F0: (v2 << ((401) & 63))}
                    switch "" {
                    case "\n\t":
                        v3[213] = "foo bar"
                    case "foo bar":
                        v28 := GenS1{F0: v2}
                        _ = v28
                    }
                    var v29 uint32 = 644
                    _ = v27
                    _ = v29
                } else {
                    fmt.Println(v20)
                    _ = len(v3)
                }
                var v30 uint64 = (v20 / ((1) | 1))
                var v31 int32 = -1
                _ = fi26
                _ = v30
                _ = v31
            }
            if ((v2 - v2) >= v2) {
                v21 = 658
                v2 += v2
                var v32 uint64 = v20
                _ = v32
            } else {
                v33 := GenS4{F0: 0, F1: ("" + "hello"), F2: 675}
                v34 := v3[0]
                fmt.Println(v34, v34)
                _ = v33
                _ = v34
            }
            _ = fi24
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v20 = v20
    v20 = 715
    for fi35 := 0; fi35 < 15; fi35++ {
        if true {
            var v36 float64 = 0.1
            _ = v36
        } else {
            fmt.Println(v2, fi35, fi35)
        }
        v37 := v3[946]
        var v38 uint16 = (^(v2))
        _ = fi35
        _ = v37
        _ = v38
    }
    v21 = 811
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
