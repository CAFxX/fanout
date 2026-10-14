package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int16; F2 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint32; F1 uint8 }
func genFn0(p0 int32, p1 int8, p2 int32) int32 {
    var v4 int16 = 0
    var v5 int16 = 32767
    var v6 int16 = (v5 << ((65535) & 63))
    var v7 bool = (!(0 >= 392))
    var v8 int32 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (^(p2))
}
func genFn1(p0 uint32) int64 {
    var v9 uint16 = 2
    v9 -= 683
    v10 := []int{1, 42, -1, -910}
    var v11 uint8 = 0
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    var v3 uint32 = 1
    v3 += v3
    v12 := GenS1{F0: -742}
    switch "foo bar" {
    case "hello":
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 string = "a"
            var v15 string = (v14 + ((v14 + (("a" + "a") + v14)) + ("foo bar" + v14)))
            v16 := []int{(^(fi13)), (fi13 * (fi13 / (((-(fi13))) | 1)))}
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
    case "":
        if false {
            if true {
                _ = len(v1)
                if (701 >= 0) {
                    for fi17 := 0; fi17 < 17; fi17++ {
                        var v18 uint = 50
                        switch (fi17 + 4294967295) {
                        case -501:
                            var v19 float32 = 1.5
                            _ = v19
                        case 128:
                            var v20 int32 = -2147483648
                            _ = v20
                        case 4294967295:
                            v18 = v18
                        default:
                            var v21 uint32 = 65535
                            _ = v21
                        }
                        _ = fi17
                        _ = v18
                    }
                    if (((1 > 0) || (((!(42 >= 0)) && (65535 != 0)) || (255 != 87))) && false) {
                        if false {
                            fmt.Println(v3)
                        } else {
                            fmt.Println(v3)
                            if (255 != 255) {
                                var v22 float64 = (0.1 * 3.14159)
                                _ = len(v1)
                                _ = v22
                            }
                        }
                    }
                } else {
                    if (-2147483648 >= genFn0(42, -128, 0)) {
                        v23 := GenS2{F0: 100, F1: -37, F2: false}
                        _ = v23
                    } else {
                        delete(v2, -436)
                        delete(v1, 2147483648)
                    }
                }
                var v24 uint64 = 867
                _ = v24
            } else {
                v3 -= v3
                v25 := []int{9223372036854775807}
                var v26 int16 = 730
                _ = v25
                _ = v26
            }
        }
    }
    _ = len(v1)
    if false {
        v27 := GenS2{F0: -1, F1: 1000, F2: (0 == 128)}
        v28 := GenS2{F0: 42, F1: 0, F2: false}
        if (559 != 42) {
            var v29 int16 = -364
            v29 = v29
            _ = v29
        } else {
            var v30 uint32 = 0
            fmt.Println(v3)
            _ = v30
        }
        _ = v27
        _ = v28
    } else {
        delete(v1, 4294967295)
        if (-83 < -1) {
            v31 := v2[0]
            _ = v31
        }
        for fi32 := 0; fi32 < 6; fi32++ {
            for fi33 := 0; fi33 < 8; fi33++ {
                v3 -= v3
                if true {
                    var v34 uint64 = 42
                    var v35 uint64 = (v34 >> ((65535) & 63))
                    v36 := []int{fi33}
                    _ = v34
                    _ = v35
                    _ = v36
                } else {
                    delete(v2, fi32)
                    v3 = (v3 / ((v3) | 1))
                    if true {
                        fmt.Println(fi32)
                        if (true || (0.1 < 1.5)) {
                            if false {
                                var v37 uint64 = 42
                                var v38 bool = false
                                _ = v37
                                _ = v38
                            }
                            _ = len(v2)
                        }
                        v3 = 2147483648
                    } else {
                        fmt.Println(fi33, fi32)
                    }
                    v3 += v3
                }
                var v39 int8 = 100
                switch (fi33 - fi33) {
                case -210:
                    switch 127 {
                    case -341:
                        if false {
                            var v40 uint = 928
                            v41 := GenS1{F0: 820}
                            _ = v40
                            _ = v41
                        } else {
                            _ = len(v2)
                            var v42 uint = 18446744073709551615
                            _ = v42
                        }
                    }
                case 2147483647:
                    v43 := GenS3{F0: "\n\t"}
                    _ = v43
                case 841:
                    var v44 uint = 933
                    _ = v44
                default:
                    var v45 float32 = -2.25
                    _ = v45
                }
                if false {
                    switch (fi33 ^ -2147483648) {
                    case 0:
                        v46 := GenS4{F0: true, F1: 1, F2: 845}
                        _ = v46
                    }
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        v47 := []int{483}
                        _ = v47
                    case "":
                        if ((false && true) && (false || (!((!false) && (!true))))) {
                            var v48 int = fi32
                            v48 -= fi32
                            v39 += v39
                            _ = v48
                        } else {
                            v3 = v3
                            var v49 int64 = 0
                            _ = v49
                        }
                        v50 := []int{-2147483648, fi32, fi32}
                        _ = v50
                    }
                }
                if (0 <= 51) {
                    v3 = v3
                    var v51 int16 = 1
                    var v52 float64 = 1.5
                    _ = v51
                    _ = v52
                } else {
                    var v53 int = (fi32 * (^(fi33)))
                    v54 := GenS5{F0: v3, F1: 253}
                    _ = v53
                    _ = v54
                }
                _ = fi33
                _ = v39
            }
            _ = fi32
        }
    }
    v55 := GenS1{F0: -1}
    v56 := []int{9223372036854775807}
    if false {
        var v57 int8 = 42
        var v58 int64 = -9223372036854775808
        _ = v57
        _ = v58
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v56:", len(v56), v56)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v55
    _ = v56
}
