package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint; F2 uint }
type GenS4 struct { F0 int16; F1 int8; F2 uint64 }
func genFn0(p0 int) int64 {
    p0 = (^(p0))
    v6 := []int{p0, p0, (p0 - (p0 ^ (p0 ^ 128)))}
    var v7 float32 = 0.1
    if len(v6) > 0 {
        v6[0] = p0
    }
    _ = p0
    _ = v6
    _ = v7
    return -182
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{42}
    v3 := []uint32{4294967295, 2}
    var v4 uint64 = 42
    v5 := []string{"", "foo bar"}
    if ((759 == 1) || false) {
        var v8 uint16 = 148
        _ = v8
    }
    v9 := GenS1{F0: "a"}
    v10 := GenS2{F0: 2147483648}
    switch (((("hello" + "hello") + "a") + "hello") + "hello") {
    case "hello":
        if true {
            v11 := GenS3{F0: "foo bar", F1: 0, F2: 341}
            for fi12 := 0; fi12 < 17; fi12++ {
                if (-9223372036854775808 < fi12) {
                    var v13 int8 = 1
                    for fi14 := 0; fi14 < 13; fi14++ {
                        v4 = 1
                        v4 += (^(v4))
                        fmt.Println(fi12, fi12, fi14)
                        _ = fi14
                    }
                    if false {
                        var v15 uint64 = (v4 | (v4 | v4))
                        for fi16 := 0; fi16 < 6; fi16++ {
                            var v17 int64 = 1
                            _ = fi16
                            _ = v17
                        }
                        v18 := GenS4{F0: 32767, F1: (v13 % ((v13) | 1)), F2: v4}
                        _ = v15
                        _ = v18
                    }
                    _ = v13
                }
                v19 := GenS4{F0: 1, F1: -128, F2: (v4 / (((v4 % (((v4 & v4)) | 1))) | 1))}
                var v20 int64 = 708
                _ = fi12
                _ = v19
                _ = v20
            }
            _ = v11
        } else {
            v4 = 4294967295
        }
    default:
        v21 := GenS4{F0: -32768, F1: 68, F2: (v4 - (v4 * 2))}
        _ = v21
    }
    var v22 int64 = 281
    delete(v1, 784)
    v23 := GenS1{F0: ("" + "")}
    v22 = 42
    var v24 int8 = 43
    delete(v1, uint32(v4))
    switch 32767 {
    case -470:
        switch 106 {
        case 32768:
            fmt.Println(v24)
        case 9223372036854775807:
            var v25 int32 = 42
            _ = v25
        }
    case -9223372036854775808:
        var v26 uint16 = 65535
        _ = v26
    case 4294967295:
        v3 = append(v3, 65535)
    }
    if (1000 > 1000) {
        v24 = (v24 << ((478) & 63))
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi27 := 0; fi27 < 16; fi27++ {
            var v28 float64 = -2.25
            _ = fi27
            _ = v28
        }
    }
    var v29 bool = (1.5 < float32(((0.0 * 1.5) - 0.0)))
    v30 := GenS3{F0: "", F1: 0, F2: 614}
    for fi31 := 0; fi31 < 8; fi31++ {
        v32 := GenS1{F0: "hello"}
        _ = fi31
        _ = v32
    }
    var v33 bool = (((v22 << ((671) & 63)) != -9223372036854775808) || v29)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v9
}
