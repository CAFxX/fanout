package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int64, p1 int, p2 uint64) uint8 {
    v5 := GenS1{F0: -1}
    v6 := GenS2{F0: 924, F1: (p2 | p2), F2: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 235
}
func genFn1(p0 uint) uint32 {
    fmt.Println(p0)
    if true {
        for fi7 := 0; fi7 < 19; fi7++ {
            p0 = (p0 ^ p0)
            fmt.Println(fi7, fi7)
            _ = fi7
        }
    }
    var v8 int16 = 32767
    v9 := GenS2{F0: 556, F1: 134, F2: 0}
    _ = p0
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 uint8 = 255
    var v2 bool = (float32(0.0) == 1.5)
    var v3 uint64 = uint64(v1)
    var v4 uint64 = (v3 * 135)
    v2 = ((0 == 255) && (genFn1(18446744073709551615) > 836))
    v10 := GenS1{F0: 1}
    if ("hello" < ("\n\t" + "foo bar")) {
        var v11 string = "hello"
        _ = v11
    }
    var v12 uint32 = 276
    v3 += v3
    for fi13 := 0; fi13 < 5; fi13++ {
        if ((1 > 42) || ((!(false && true)) || (!v2))) {
            if (float64(1e10) > 1.5) {
                for fi14 := 0; fi14 < 13; fi14++ {
                    v3 -= 4294967295
                    v12 = genFn1(65535)
                    _ = fi14
                }
            } else {
                v4 = uint64(v12)
                var v15 float32 = float32(0.1)
                _ = v15
            }
            var v16 bool = (1 != 624)
            v17 := GenS2{F0: (v12 + v12), F1: v3, F2: genFn1(uint(v3))}
            _ = v16
            _ = v17
        } else {
            var v18 uint16 = 77
            v19 := GenS2{F0: 2147483647, F1: (v3 | (v3 * 4294967295)), F2: (v12 >> ((608) & 63))}
            fmt.Println(v3)
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 uint16 = 2
            v22 := []int{(fi20 / ((fi13) | 1)), 128, fi13}
            v23 := GenS1{F0: -624}
            if v2 {
                switch ("hello" + "a") {
                case "\n\t":
                    var v24 int16 = int16(fi20)
                    fmt.Println(v3, v24, fi13)
                    _ = v24
                default:
                    _ = len(v22)
                }
            }
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        v25 := []int{}
        _ = fi13
        _ = v25
    }
    v3 += v4
    if (3.14159 < 1e10) {
        if v2 {
            var v26 uint64 = (v4 << ((4294967295) & 63))
            _ = v26
        }
        v4 -= uint64(v1)
        if v2 {
            if true {
                fmt.Println(v4)
                v27 := GenS1{F0: 42}
                for fi28 := 0; fi28 < 14; fi28++ {
                    var v29 int = fi28
                    if false {
                        var v30 int32 = 1
                        v31 := GenS1{F0: (v30 & (v30 >> ((4294967295) & 63)))}
                        _ = v30
                        _ = v31
                    }
                    fmt.Println(v4, v4, v4)
                    _ = fi28
                    _ = v29
                }
                _ = v27
            } else {
                v32 := GenS3{F0: 32767}
                for fi33 := 0; fi33 < 13; fi33++ {
                    v1 += v1
                    _ = fi33
                }
                var v34 int8 = 1
                _ = v32
                _ = v34
            }
        }
    } else {
        var v35 float64 = 3.14159
        var v36 uint64 = (v4 * v3)
        var v37 int16 = 151
        _ = v35
        _ = v36
        _ = v37
    }
    v12 += uint32(v3)
    v38 := GenS4{F0: (v4 % (((v3 & v3)) | 1)), F1: 4294967295, F2: (^(v4))}
    v39 := GenS5{F0: uint8(v12)}
    var v40 int16 = 1
    v40 += (v40 | v40)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
