package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint8 }
func genFn0(p0 uint16, p1 uint64, p2 bool) int {
    var v7 int16 = -1
    if (!(1 == 2)) {
        if (32767 > 828) {
            for fi8 := 0; fi8 < 6; fi8++ {
                for fi9 := 0; fi9 < 2; fi9++ {
                    var v10 uint = 65535
                    if (-2147483648 == -1) {
                        v7 = (v7 / (((v7 << ((v10) & 63))) | 1))
                    } else {
                        v11 := []int{}
                        var v12 int64 = 178
                        v7 -= 216
                        _ = v11
                        _ = v12
                    }
                    var v13 int64 = -9223372036854775808
                    _ = fi9
                    _ = v10
                    _ = v13
                }
                _ = fi8
            }
        } else {
            var v14 uint8 = 127
            v15 := GenS1{F0: 14, F1: 18446744073709551615, F2: 2}
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 9; fi16++ {
            p2 = (42 == 1)
            v7 += (v7 - v7)
            _ = fi16
        }
        if ((v7 | (v7 % ((-1) | 1))) < 42) {
            var v17 uint16 = p0
            var v18 float64 = float64(3.14159)
            _ = v17
            _ = v18
        } else {
            fmt.Println(v7, p1)
        }
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        switch fi19 {
        case -562:
            if p2 {
                v7 -= (v7 & v7)
            }
        }
        var v20 uint64 = (^(p1))
        _ = fi19
        _ = v20
    }
    p0 -= (p0 & p0)
    v21 := []int{32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v7
    return 32767
}
func genFn1(p0 uint64, p1 int64, p2 uint) uint8 {
    v22 := GenS1{F0: 2, F1: 170, F2: 951}
    for fi23 := 0; fi23 < 4; fi23++ {
        if (-118 <= -32768) {
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 int64 = (p1 & (p1 - p1))
                p0 += 1
                _ = fi24
                _ = v25
            }
            v26 := []int{fi23, genFn0(0, p0, (uint32(p0) <= 1)), genFn0(1, 4294967295, ((("foo bar" + "\n\t") + "a") <= "hello"))}
            p0 += 443
            _ = v26
        } else {
            var v27 float32 = 0.0
            for fi28 := 0; fi28 < 7; fi28++ {
                switch genFn0(254, p0, (((0.0 / 1e10) - 3.14159) >= 1e10)) {
                case 9223372036854775807:
                    p0 -= uint64(p2)
                }
                p0 = (p0 << ((1) & 63))
                v29 := []int{(fi23 / ((908) | 1)), (fi23 ^ (^(fi23)))}
                _ = fi28
                _ = v29
            }
            p0 -= (p0 / ((18446744073709551615) | 1))
            _ = v27
        }
        fmt.Println(fi23, p1, p2)
        _ = fi23
    }
    if false {
        v30 := GenS1{F0: 127, F1: p2, F2: 2}
        _ = v30
    }
    var v31 uint = (p2 ^ p2)
    v32 := []int{680, 4294967295, int(p1), -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v31
    _ = v32
    return 128
}

func main() {
    var v1 string = "hello"
    v2 := []int{}
    var v3 float64 = float64(float32(-2.25))
    var v4 int64 = -1
    var v5 float32 = 3.14159
    var v6 int32 = 1
    v4 = v4
    v2 = append(v2, 4294967295)
    var v33 float32 = 0.0
    v34 := GenS2{F0: 65535, F1: 42, F2: 2}
    v4 = (v4 | (^(v4)))
    v35 := GenS2{F0: 378, F1: 42, F2: 218}
    var v36 uint8 = 2
    v4 = (v4 % ((v4) | 1))
    var v37 uint16 = 255
    var v38 uint16 = v37
    fmt.Println(v36, v38)
    v39 := GenS1{F0: v36, F1: 2, F2: 42}
    var v40 string = ("" + ("" + ("foo bar" + "hello")))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
