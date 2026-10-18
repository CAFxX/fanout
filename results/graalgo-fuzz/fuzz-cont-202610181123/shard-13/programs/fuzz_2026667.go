package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int; F2 uint64 }
type GenS3 struct { F0 int8; F1 uint8 }
func genFn0() uint64 {
    var v7 string = "\n\t"
    for fi8 := 0; fi8 < 13; fi8++ {
        if (1 >= -9223372036854775808) {
            for fi9 := 0; fi9 < 16; fi9++ {
                if true {
                    if true {
                        v7 = v7
                    }
                    var v10 int64 = 163
                    var v11 bool = (0 == 128)
                    if (-1 != (v10 << ((2) & 63))) {
                        v12 := []int{(fi9 | int(v10))}
                        v13 := GenS1{F0: 637}
                        v14 := GenS1{F0: 42}
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    _ = v10
                    _ = v11
                }
                if false {
                    var v15 int16 = 387
                    _ = v15
                }
                v16 := []int{(fi9 ^ fi8)}
                _ = fi9
                _ = v16
            }
            if (-475 < 0) {
                var v17 uint = 35
                v17 = (v17 - (v17 | (v17 / ((794) | 1))))
                _ = v17
            }
            if true {
                for fi18 := 0; fi18 < 1; fi18++ {
                    var v19 uint64 = 42
                    v19 = (v19 << ((2) & 63))
                    _ = fi18
                    _ = v19
                }
            }
        }
        v7 = v7
        _ = fi8
    }
    fmt.Println(v7)
    v20 := GenS1{F0: 1}
    v21 := []int{128, -9223372036854775808}
    _ = v20
    _ = v21
    _ = v7
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = -221
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    var v5 float32 = ((0.1 * 1e10) * 0.0)
    var v6 uint = 2
    var v22 uint64 = 2
    v23 := GenS2{F0: v22, F1: -759, F2: v22}
    var v24 uint32 = 2
    if (((("foo bar" + (("" + "foo bar") + "\n\t")) + "a") + "\n\t") != "a") {
        for fi25 := 0; fi25 < 15; fi25++ {
            v26 := GenS1{F0: 2}
            v27 := []int{(fi25 % (((-(v2))) | 1))}
            for fi28 := 0; fi28 < 3; fi28++ {
                v29 := GenS2{F0: 2, F1: 4294967295, F2: (v22 & (^(v22)))}
                if true {
                    var v30 int = fi28
                    var v31 uint64 = (v22 << (((v6 ^ (v6 / ((18446744073709551615) | 1)))) & 63))
                    fmt.Println(v22, fi25)
                    _ = v30
                    _ = v31
                }
                _ = fi28
                _ = v29
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        var v32 uint = (v6 << ((324) & 63))
        _ = v32
    } else {
        var v33 int64 = 0
        v33 -= (v33 - v33)
        _ = v33
    }
    var v34 uint64 = v22
    v35 := v3[862]
    switch "" {
    case "foo bar":
        var v36 int8 = 42
        _ = v36
    default:
        fmt.Println(v6, v6)
    }
    if false {
        v37 := GenS2{F0: 42, F1: v2, F2: (v34 * v34)}
        for fi38 := 0; fi38 < 3; fi38++ {
            for fi39 := 0; fi39 < 2; fi39++ {
                v40 := GenS1{F0: 194}
                v41 := []int{v2, (v35 & -367), (^(fi38)), (v35 - v2)}
                _ = fi39
                _ = v40
                _ = v41
            }
            v42 := GenS3{F0: -6, F1: 2}
            switch 32767 {
            case 32767:
                v24 += 2
            default:
                for fi43 := 0; fi43 < 11; fi43++ {
                    v2 = v2
                    v2 = 42
                    var v44 uint8 = uint8(v24)
                    _ = fi43
                    _ = v44
                }
            }
            _ = fi38
            _ = v42
        }
        _ = v37
    } else {
        v35 = (v35 * (v35 | 0))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
