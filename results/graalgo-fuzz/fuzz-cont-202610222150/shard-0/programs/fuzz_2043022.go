package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS4 struct { F0 int32; F1 int32; F2 int16 }
type GenS5 struct { F0 int; F1 int8 }
func genFn0() string {
    for fi7 := 0; fi7 < 19; fi7++ {
        if false {
            v8 := GenS1{F0: 29, F1: -1}
            _ = v8
        } else {
            var v9 uint8 = 255
            _ = v9
        }
        var v10 int32 = -706
        _ = fi7
        _ = v10
    }
    switch (((("\n\t" + ("" + "\n\t")) + "hello") + "\n\t") + ("\n\t" + "hello")) {
    case "\n\t":
        if (true || (true && false)) {
            var v11 int64 = 0
            var v12 uint64 = 933
            if false {
                switch -9223372036854775808 {
                case 0:
                    for fi13 := 0; fi13 < 5; fi13++ {
                        for fi14 := 0; fi14 < 10; fi14++ {
                            v11 -= (^(v11))
                            v15 := GenS2{F0: 168}
                            _ = fi14
                            _ = v15
                        }
                        _ = fi13
                    }
                default:
                    v16 := GenS3{F0: (v12 / (((v12 & v12)) | 1)), F1: (0 > 128), F2: (v12 ^ v12)}
                    _ = v16
                }
                v11 = (v11 - v11)
                var v17 uint = 0
                _ = v17
            }
            _ = v11
            _ = v12
        } else {
            v18 := GenS3{F0: 710, F1: (-1 != -311), F2: 2}
            var v19 uint8 = 42
            v20 := []int{-2147483648, -9223372036854775808, -495, -1}
            _ = v18
            _ = v19
            _ = v20
        }
    case "hello":
        for fi21 := 0; fi21 < 8; fi21++ {
            v22 := GenS1{F0: -1, F1: -122}
            _ = fi21
            _ = v22
        }
    }
    if (!(!(2 >= 1))) {
        if true {
            v23 := GenS3{F0: 0, F1: (float64(float32(((0.0 * 0.0) - float64(-2.25)))) != 1.5), F2: 461}
            _ = v23
        }
        var v24 string = "a"
        fmt.Println(v24)
        _ = v24
    }
    var v25 int64 = 1
    v26 := GenS4{F0: 1000, F1: -1, F2: 214}
    _ = v25
    _ = v26
    return "hello"
}
func genFn1() uint64 {
    v27 := GenS5{F0: 9223372036854775807, F1: 100}
    var v28 uint64 = 18446744073709551615
    v29 := []int{}
    _ = v27
    _ = v28
    _ = v29
    return v28
}

func main() {
    var v1 int64 = -1
    var v2 int32 = -2147483648
    v3 := make(map[uint32]int)
    v4 := []uint32{}
    var v5 int64 = (v1 << ((42) & 63))
    v6 := []string{""}
    for fi30 := 0; fi30 < 19; fi30++ {
        switch "" {
        case "\n\t":
            for fi31 := 0; fi31 < 15; fi31++ {
                var v32 string = "a"
                v5 = (v1 | v5)
                _ = fi31
                _ = v32
            }
        default:
            v2 = (v2 / (((v2 / (((v2 + -1)) | 1))) | 1))
        }
        var v33 uint64 = 920
        v2 = (v2 << ((661) & 63))
        _ = fi30
        _ = v33
    }
    var v34 float32 = (1.5 * -2.25)
    v35 := GenS2{F0: 32767}
    delete(v3, 2147483648)
    if false {
        _ = len(v3)
        v2 = v2
        for fi36 := 0; fi36 < 18; fi36++ {
            v5 = 1
            delete(v3, 1)
            _ = fi36
        }
    } else {
        v2 -= v2
    }
    v5 += 1
    v5 = (v1 << ((1) & 63))
    v37 := GenS4{F0: v2, F1: v2, F2: 42}
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
