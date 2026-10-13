package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS4 struct { F0 bool; F1 int32; F2 uint16 }
func genFn0(p0 uint64, p1 int8) int64 {
    switch -581 {
    case -281:
        fmt.Println(p0)
    case 640:
        v7 := []int{}
        _ = v7
    case 32768:
        var v8 int32 = 1
        _ = v8
    default:
        fmt.Println(p0, p1)
    }
    p1 -= (p1 ^ (p1 / ((p1) | 1)))
    var v9 bool = (((-115 != 32767) || (2147483647 >= 352)) && false)
    _ = p0
    _ = p1
    _ = v9
    return -1
}

func main() {
    var v1 int64 = -863
    var v2 int64 = v1
    v3 := []uint32{2147483648, 42, 2147483648}
    var v4 float64 = 1e10
    var v5 int32 = 1
    var v6 uint = 138
    for fi10 := 0; fi10 < 2; fi10++ {
        if (42 >= (v2 + genFn0(1, 108))) {
            if true {
                var v11 int32 = int32(fi10)
                for fi12 := 0; fi12 < 3; fi12++ {
                    var v13 string = "\n\t"
                    if true {
                        if len(v3) > 0 {
                            v3[0] = 2147483648
                        }
                        var v14 uint32 = 4294967295
                        var v15 int16 = 32767
                        _ = v14
                        _ = v15
                    } else {
                        v4 += 0.0
                    }
                    _ = fi12
                    _ = v13
                }
                if (true && ((("" + "a") > ("\n\t" + ("foo bar" + ("a" + "a")))) && (-999 != (-(fi10))))) {
                    var v16 uint64 = 58
                    v1 -= genFn0((v16 | (^(v16))), 100)
                    v17 := GenS1{F0: (v2 >> ((v6) & 63))}
                    _ = v16
                    _ = v17
                } else {
                    v4 = ((v4 / v4) - (v4 + v4))
                    var v18 uint64 = 18446744073709551615
                    for fi19 := 0; fi19 < 7; fi19++ {
                        v3 = append(v3, 0)
                        if (-60 < 42) {
                            v1 = (v1 % ((genFn0((v18 | 2), int8(v18))) | 1))
                        }
                        _ = fi19
                    }
                    _ = v18
                }
                _ = v11
            } else {
                if (false || (42 == 1)) {
                    v6 -= (v6 << ((0) & 63))
                } else {
                    var v20 string = "a"
                    _ = v20
                }
                v1 += (v1 - (v1 >> ((v6) & 63)))
                v6 -= (v6 / (((^(v6))) | 1))
            }
        }
        var v21 int16 = 1
        v5 -= v5
        _ = fi10
        _ = v21
    }
    v22 := GenS2{F0: 119}
    var v23 int = -2147483648
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 float32 = 1e10
        v3 = append(v3, 4294967295)
        _ = fi24
        _ = v25
    }
    if true {
        v26 := GenS3{F0: 14, F1: uint32(v6), F2: 665}
        fmt.Println(v6, v23, v4)
        var v27 int32 = (v5 ^ v5)
        _ = v26
        _ = v27
    }
    var v28 uint16 = 0
    if false {
        v5 += 42
    } else {
        for fi29 := 0; fi29 < 13; fi29++ {
            var v30 uint32 = 1
            var v31 int32 = v5
            v32 := GenS3{F0: 209, F1: 931, F2: 1}
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
        if true {
            var v33 uint = (v6 & uint(v28))
            var v34 float32 = 0.1
            fmt.Println(v4)
            _ = v33
            _ = v34
        }
    }
    v35 := GenS4{F0: (false && (255 != 249)), F1: (-(v5)), F2: (v28 | v28)}
    fmt.Println(v1, v23, v28)
    fmt.Println(v2)
    if len(v3) > 0 {
        v3[0] = uint32(v28)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
