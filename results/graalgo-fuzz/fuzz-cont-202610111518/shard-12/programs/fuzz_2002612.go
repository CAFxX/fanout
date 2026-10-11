package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint8; F1 int64; F2 int }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint8) bool {
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return true
}
func genFn1() int {
    var v4 int32 = 1
    v5 := []int{128, 128, 550, 32767}
    v5 = append(v5, 142)
    for fi6 := 0; fi6 < 3; fi6++ {
        v4 = -2147483648
        _ = fi6
    }
    _ = v4
    _ = v5
    return -299
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[string]string)
    var v3 int = 1
    fmt.Println(v1, v3)
    for fi7 := 0; fi7 < 1; fi7++ {
        fmt.Println(fi7)
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 int16 = -32768
            if true {
                v10 := []int{127, (v3 >> ((2) & 63)), fi7, (fi8 >> ((763) & 63))}
                v1 = v1
                v1 = v1
                _ = v10
            } else {
                v11 := []int{(v3 ^ 32768), fi8, (fi7 * (v3 + genFn1())), (fi8 / (((v3 + (fi8 % (((fi8 & genFn1())) | 1)))) | 1))}
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                v12 := GenS1{F0: 0, F1: 2147483648, F2: 255}
                _ = v11
                _ = v12
            }
            v3 += fi7
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v13 float32 = 0.0
    var v14 int64 = -1
    for fi15 := 0; fi15 < 9; fi15++ {
        v3 += (-(fi15))
        if (false || true) {
            v16 := GenS2{F0: 566, F1: -9223372036854775808}
            v13 += 0.0
            v17 := []int{genFn1(), (v3 - v3), v3, fi15}
            _ = v16
            _ = v17
        } else {
            if ((v1 + v1) != v1) {
                if false {
                    if false {
                        v2[(v1 + "\n\t")] = (v1 + ((v1 + v1) + v1))
                    } else {
                        v18 := GenS1{F0: 65535, F1: 42, F2: 2}
                        _ = v18
                    }
                } else {
                    v14 -= (v14 % (((-(v14))) | 1))
                }
                var v19 uint32 = 2147483648
                _ = v19
            }
            v14 = (v14 % ((v14) | 1))
        }
        delete(v2, v1)
        _ = fi15
    }
    v20 := GenS3{F0: 0, F1: (^(v14)), F2: v3}
    var v21 int64 = v14
    v1 = v1
    for fi22 := 0; fi22 < 5; fi22++ {
        v14 += -465
        var v23 string = "\n\t"
        switch -192 {
        case 1:
            for fi24 := 0; fi24 < 13; fi24++ {
                delete(v2, v23)
                v25 := GenS4{F0: 42}
                delete(v2, v1)
                _ = fi24
                _ = v25
            }
        case -2147483648:
            v14 = 9
        default:
            v14 = v14
        }
        _ = fi22
        _ = v23
    }
    if (true || false) {
        v13 += v13
    } else {
        var v26 uint = 639
        _ = v26
    }
    var v27 int8 = 127
    fmt.Println(v1, v3)
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 bool = false
        v14 = (v21 - 1)
        for fi30 := 0; fi30 < 8; fi30++ {
            v27 = v27
            _ = fi30
        }
        _ = fi28
        _ = v29
    }
    var v31 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v31
}
