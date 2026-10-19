package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 bool }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8; F1 uint16; F2 int8 }
func genFn0() uint8 {
    var v6 float64 = 1.5
    v6 = (v6 - 1.5)
    var v7 uint = 255
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 int16 = 42
        v10 := GenS1{F0: 143, F1: (v9 * (v9 / (((^(v9))) | 1))), F2: false}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int16 = 42
    _ = v11
    _ = v6
    _ = v7
    return 42
}
func genFn1() int8 {
    var v12 uint32 = 65535
    for fi13 := 0; fi13 < 2; fi13++ {
        if ((true && false) || false) {
            v12 = 869
            v14 := GenS1{F0: genFn0(), F1: int16(fi13), F2: (!false)}
            v12 += v12
            _ = v14
        }
        _ = fi13
    }
    _ = v12
    return 100
}

func main() {
    var v1 float32 = 1e10
    v2 := []int{42, 4294967295, 128}
    v3 := make(map[uint32]string)
    var v4 uint8 = 126
    v5 := []uint32{}
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 int16 = 0
        if true {
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 uint32 = 864
                if true {
                    var v19 int16 = v16
                    if len(v2) > 0 {
                        v2[0] = fi15
                    }
                    _ = v19
                } else {
                    var v20 int = fi17
                    _ = v20
                }
                _ = len(v2)
                v21 := v3[1]
                fmt.Println(fi15, fi15, v4)
                _ = fi17
                _ = v18
                _ = v21
            }
            v1 += (v1 * v1)
        }
        _ = fi15
        _ = v16
    }
    v1 += v1
    if false {
        v4 += 5
    } else {
        var v22 int32 = -1
        if true {
            if false {
                v4 = 0
                for fi23 := 0; fi23 < 14; fi23++ {
                    v24 := GenS2{F0: 979, F1: -588}
                    _ = fi23
                    _ = v24
                }
                v22 = v22
            } else {
                v25 := GenS3{F0: 980, F1: 12}
                fmt.Println(v22, v22)
                _ = v25
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v26 int = 32767
            for fi27 := 0; fi27 < 9; fi27++ {
                for fi28 := 0; fi28 < 14; fi28++ {
                    v4 = v4
                    var v29 bool = (true || (4294967295 > 2))
                    _ = fi28
                    _ = v29
                }
                _ = fi27
            }
            _ = v26
        }
        v30 := GenS4{F0: genFn0(), F1: -32768}
        _ = v22
        _ = v30
    }
    for fi31 := 0; fi31 < 1; fi31++ {
        switch 469 {
        case 398:
            v32 := GenS5{F0: (v4 - (v4 % ((v4) | 1)))}
            _ = v32
        case 32768:
            fmt.Println(v4, v4, v4)
        case 904:
            var v33 uint32 = 10
            _ = v33
        }
        if false {
            var v34 int32 = 42
            switch fi31 {
            case 569:
                _ = len(v2)
            case 420:
                for fi35 := 0; fi35 < 20; fi35++ {
                    if false {
                        if (!false) {
                            v36 := GenS1{F0: v4, F1: -32768, F2: (32767 >= 0)}
                            var v37 uint64 = 2
                            _ = v36
                            _ = v37
                        } else {
                            v4 -= genFn0()
                        }
                    }
                    if len(v2) > 0 {
                        v2[0] = 127
                    }
                    v34 -= (v34 | v34)
                    _ = fi35
                }
            default:
                v34 -= v34
            }
            for fi38 := 0; fi38 < 17; fi38++ {
                v4 = (v4 >> ((978) & 63))
                var v39 string = (("foo bar" + ("\n\t" + "\n\t")) + "\n\t")
                v40 := v3[4294967295]
                v34 = v34
                _ = fi38
                _ = v39
                _ = v40
            }
            _ = v34
        }
        var v41 float64 = 1e10
        _ = fi31
        _ = v41
    }
    v42 := GenS1{F0: v4, F1: 1, F2: true}
    v4 = v4
    _ = len(v5)
    var v43 int32 = -1
    if true {
        fmt.Println(v4)
        v4 -= v4
    }
    v4 = (v4 | v4)
    v44 := GenS6{F0: genFn1(), F1: 0, F2: -1}
    var v45 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
