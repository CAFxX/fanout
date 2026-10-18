package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint32 }
func genFn0(p0 uint16, p1 uint, p2 uint) string {
    if true {
        if false {
            for fi7 := 0; fi7 < 3; fi7++ {
                v8 := []int{(fi7 & -9223372036854775808)}
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = fi7
                _ = v8
            }
            p2 -= (p1 ^ p2)
        }
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return ("foo bar" + "")
}
func genFn1(p0 int64, p1 uint, p2 int64) int8 {
    v9 := []int{int(p2)}
    v10 := GenS1{F0: -1, F1: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -128
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{2}
    var v3 uint = 1
    var v4 bool = true
    var v5 int16 = -32768
    var v6 bool = false
    fmt.Println(v3, v6)
    if (((0.0 - 1e10) / 1e10) == 1e10) {
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 uint32 = 42
            if v4 {
                for fi13 := 0; fi13 < 3; fi13++ {
                    _ = len(v2)
                    var v14 int64 = 428
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if false {
                        switch "hello" {
                        case "\n\t":
                            v15 := GenS2{F0: -1, F1: v5}
                            _ = v15
                        default:
                            v16 := GenS1{F0: 2147483647, F1: 2}
                            _ = v16
                        }
                        switch "" {
                        case "a":
                            v14 = v14
                        case "foo bar":
                            v4 = v6
                        default:
                            v6 = (-113 >= 93)
                        }
                        fmt.Println(v5, v14, fi13)
                    }
                    if len(v2) > 0 {
                        v2[0] = (v12 & (v12 << ((v3) & 63)))
                    }
                    _ = fi13
                    _ = v14
                }
            }
            _ = fi11
            _ = v12
        }
    } else {
        for fi17 := 0; fi17 < 11; fi17++ {
            for fi18 := 0; fi18 < 20; fi18++ {
                fmt.Println(fi18, fi18, fi17)
                _ = fi18
            }
            v1[2] = v4
            _ = fi17
        }
        v5 -= v5
        v19 := GenS3{F0: v3, F1: 2, F2: 2}
        _ = v19
    }
    if len(v2) > 0 {
        v2[0] = 578
    }
    fmt.Println(v6)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v20 := GenS4{F0: v5, F1: 65535, F2: 753}
    v3 = (v3 + (v3 ^ (^(v3))))
    v21 := GenS3{F0: (^(v3)), F1: 4294967295, F2: 127}
    v6 = (-1 != -1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
