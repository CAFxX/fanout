package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
func genFn0(p0 int16, p1 uint8, p2 int) int32 {
    var v7 uint16 = 2
    if true {
        if (0 == 0) {
            p2 = p2
        } else {
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 int = (p2 >> ((0) & 63))
                for fi10 := 0; fi10 < 7; fi10++ {
                    fmt.Println(fi8, v9, fi8)
                    v7 += (v7 * v7)
                    var v11 uint64 = 587
                    var v12 bool = ((p1 + 255) >= (p1 - 42))
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v13 := GenS1{F0: 2}
                _ = fi8
                _ = v13
                _ = v9
            }
        }
        var v14 uint64 = 18446744073709551615
        fmt.Println(v14, p2, p1)
        _ = v14
    }
    v15 := GenS1{F0: 1}
    if true {
        v16 := []int{1, (p2 >> ((4294967295) & 63)), p2}
        p0 = (p0 % (((p0 >> ((18446744073709551615) & 63))) | 1))
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v7
    return 400
}

func main() {
    var v1 uint = 4294967295
    var v2 int32 = 1
    v3 := []uint32{1, 2, 65535}
    var v4 uint64 = 0
    v5 := make(map[int]bool)
    var v6 uint32 = 0
    for fi17 := 0; fi17 < 5; fi17++ {
        v6 = v6
        for fi18 := 0; fi18 < 6; fi18++ {
            if len(v3) > 0 {
                v3[0] = v6
            }
            fmt.Println(v2)
            var v19 int8 = 100
            _ = fi18
            _ = v19
        }
        _ = len(v3)
        _ = fi17
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        if true {
            var v21 int8 = 100
            v22 := GenS1{F0: v1}
            if false {
                if false {
                    if ((v4 + (v4 << ((v1) & 63))) > (v4 >> ((v1) & 63))) {
                        delete(v5, 2147483647)
                        v2 = (v2 * (v2 / (((v2 ^ -1)) | 1)))
                        v23 := GenS1{F0: v1}
                        _ = v23
                    } else {
                        v4 = 2
                        var v24 float32 = 0.1
                        _ = v24
                    }
                }
            }
            _ = v21
            _ = v22
        } else {
            for fi25 := 0; fi25 < 17; fi25++ {
                switch 964 {
                case 9223372036854775807:
                    var v26 uint = v1
                    _ = v26
                case 32768:
                    v1 = (v1 & (^(v1)))
                case 2147483647:
                    var v27 float32 = (0.1 - (0.1 + 0.1))
                    _ = v27
                default:
                    _ = len(v5)
                }
                _ = fi25
            }
        }
        for fi28 := 0; fi28 < 14; fi28++ {
            v29 := GenS1{F0: v1}
            var v30 uint64 = v4
            _ = fi28
            _ = v29
            _ = v30
        }
        v4 = (v4 * (v4 | v4))
        _ = fi20
    }
    fmt.Println(v6, v2)
    fmt.Println(v4, v6)
    v2 += v2
    if ((v6 < (^(v6))) || (!(v2 <= (v2 ^ genFn0(683, 212, -372))))) {
        var v31 int32 = (-(v2))
        fmt.Println(v31, v6, v2)
        _ = v31
    } else {
        v1 += (v1 + (^(v1)))
        for fi32 := 0; fi32 < 14; fi32++ {
            var v33 float64 = -2.25
            _ = fi32
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        v35 := GenS1{F0: (v1 ^ (^(v1)))}
        v2 += v2
        v5[632] = true
        _ = fi34
        _ = v35
    }
    v36 := GenS2{F0: ("foo bar" + "foo bar")}
    var v37 int = 9223372036854775807
    var v38 int64 = -9223372036854775808
    v6 -= (^(v6))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
