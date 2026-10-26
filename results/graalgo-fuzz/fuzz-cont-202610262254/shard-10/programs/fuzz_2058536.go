package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS2 struct { F0 string; F1 string; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 int16; F2 int }
func genFn0(p0 int16, p1 uint32, p2 uint64) int64 {
    v6 := GenS1{F0: 65535, F1: 42, F2: 1}
    v7 := []int{0, 112}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -1
}
func genFn1() bool {
    var v8 int = 479
    v9 := GenS1{F0: 4294967295, F1: 42, F2: 2}
    _ = v8
    _ = v9
    return ((!false) && (!true))
}

func main() {
    v1 := []uint32{0, 2, 2147483648, 2147483648}
    v2 := []int{115, -2147483648}
    v3 := make(map[int]int)
    var v4 uint16 = 69
    v5 := make(map[uint32]int)
    v4 = 32767
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS2{F0: "", F1: "foo bar", F2: 42}
        if (false || (true && (140 < 1))) {
            if (!false) {
                var v12 float32 = 3.14159
                for fi13 := 0; fi13 < 19; fi13++ {
                    if len(v2) > 0 {
                        v2[0] = (fi13 ^ (fi10 * (fi13 - (fi13 / ((9223372036854775807) | 1)))))
                    }
                    if true {
                        v4 += (v4 & 1)
                        v4 = 108
                    } else {
                        var v14 float64 = 3.14159
                        if (((v14 - (v14 - float64(0.0))) / 1.5) <= (float64((0.0 / -2.25)) * (-2.25 / v14))) {
                            delete(v3, fi13)
                        } else {
                            v15 := GenS1{F0: 65535, F1: 65535, F2: 2147483648}
                            switch ("" + "\n\t") {
                            case "":
                                if true {
                                    if len(v2) > 0 {
                                        v2[0] = 32768
                                    }
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if false {
                                        v12 += v12
                                    }
                                }
                            default:
                                fmt.Println(v14)
                                if ((!false) || (-552 >= 42)) {
                                    v12 = ((0.0 - v12) * 1.5)
                                    if len(v2) > 0 {
                                        v2[0] = (fi10 - (fi10 & 32767))
                                    }
                                    delete(v5, 42)
                                } else {
                                    v4 -= (^(v4))
                                    var v16 bool = false
                                    _ = v16
                                }
                                fmt.Println(fi13, v12, v12)
                                v17 := GenS3{F0: 0}
                                _ = v17
                            }
                            v12 += 0.1
                            _ = v15
                        }
                        var v18 float64 = v14
                        _ = v14
                        _ = v18
                    }
                    _ = fi13
                }
                for fi19 := 0; fi19 < 19; fi19++ {
                    if (v12 < 0.0) {
                        fmt.Println(fi10, v12, v12)
                        var v20 bool = true
                        v4 += (v4 & 42)
                        _ = v20
                    }
                    fmt.Println(fi19)
                    _ = fi19
                }
                _ = v12
            }
        }
        _ = fi10
        _ = v11
    }
    v21 := GenS4{F0: 512, F1: 32767, F2: 32767}
    if (false || true) {
        if ((255 <= 160) || (true && (!(!(true || true))))) {
            if false {
                v4 = (v4 ^ (v4 & v4))
                var v22 uint64 = 4294967295
                _ = v22
            }
            for fi23 := 0; fi23 < 8; fi23++ {
                fmt.Println(v4, v4)
                fmt.Println(fi23, v4)
                _ = fi23
            }
            v4 += v4
        }
    }
    v4 = (^(v4))
    if (false || false) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v4 = 32767
        v24 := GenS2{F0: "", F1: "", F2: 1}
        _ = v24
    } else {
        for fi25 := 0; fi25 < 9; fi25++ {
            v26 := v3[(fi25 ^ fi25)]
            _ = fi25
            _ = v26
        }
    }
    v4 = v4
    v4 -= v4
    v4 -= 676
    fmt.Println(v4)
    v1 = append(v1, 0)
    v27 := GenS2{F0: "\n\t", F1: ("hello" + ("" + "\n\t")), F2: 0}
    v4 = 42
    var v28 int32 = 967
    if false {
        v4 = 557
        v28 = (v28 - (^(v28)))
        for fi29 := 0; fi29 < 8; fi29++ {
            _ = len(v3)
            if len(v1) > 0 {
                v1[0] = 65535
            }
            _ = fi29
        }
    } else {
        v30 := GenS4{F0: 231, F1: -32768, F2: 32767}
        fmt.Println(v28, v28)
        for fi31 := 0; fi31 < 19; fi31++ {
            v4 = 752
            v4 += (v4 << ((246) & 63))
            switch 32767 {
            case -9223372036854775808:
                v28 = v28
            }
            _ = fi31
        }
        _ = v30
    }
    v3[905] = 32768
    v28 = v28
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
