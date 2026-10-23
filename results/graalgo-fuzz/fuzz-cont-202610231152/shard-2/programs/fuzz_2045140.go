package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
func genFn0(p0 bool) int8 {
    p0 = true
    if false {
        fmt.Println(p0)
        p0 = (false && (0.1 >= -2.25))
        fmt.Println(p0)
    }
    p0 = p0
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 float32 = ((-2.25 - 1e10) / ((1e10 + 3.14159) + 0.1))
        _ = fi7
        _ = v8
    }
    var v9 bool = false
    _ = p0
    _ = v9
    return 47
}
func genFn1(p0 string, p1 uint32) uint32 {
    v10 := GenS1{F0: 376, F1: 0}
    p1 += (p1 / ((p1) | 1))
    _ = p0
    _ = p1
    _ = v10
    return p1
}

func main() {
    var v1 int64 = 205
    var v2 string = ""
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]bool)
    var v5 int = 473
    v6 := []int{32767, v5, (v5 >> ((0) & 63))}
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 uint16 = 644
        delete(v3, 946)
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 uint16 = 32767
        _ = fi13
        _ = v14
    }
    var v15 int16 = 32767
    var v16 int64 = v1
    v15 += v15
    v17 := v4[65535]
    for fi18 := 0; fi18 < 14; fi18++ {
        fmt.Println(v16)
        if len(v6) > 0 {
            v6[0] = v5
        }
        _ = fi18
    }
    v1 += (v1 % ((v16) | 1))
    if v17 {
        if (v17 && v17) {
            if (65535 != 173) {
                fmt.Println(v17, v16, v5)
            }
            if (18446744073709551615 > 2) {
                if v17 {
                    var v19 int16 = (v15 ^ 42)
                    _ = v19
                } else {
                    v17 = (float64(3.14159) <= 3.14159)
                }
                var v20 int = (^(v5))
                v15 = (v15 >> ((760) & 63))
                _ = v20
            }
        }
        switch v2 {
        case "hello":
            v21 := v4[488]
            _ = v21
        case "\n\t":
            if (!v17) {
                v2 = ("hello" + "\n\t")
                v22 := GenS1{F0: 42, F1: 409}
                _ = v22
            }
        }
        if true {
            if v17 {
                var v23 int64 = (v1 << ((4294967295) & 63))
                v15 -= v15
                for fi24 := 0; fi24 < 17; fi24++ {
                    if (("\n\t" + v2) <= "") {
                        v17 = ((0.1 - 0.0) <= (1.5 - 1.5))
                        v17 = v17
                    }
                    for fi25 := 0; fi25 < 14; fi25++ {
                        fmt.Println(v2, v17)
                        _ = len(v3)
                        _ = fi25
                    }
                    v26 := GenS2{F0: v2}
                    _ = fi24
                    _ = v26
                }
                _ = v23
            }
        } else {
            v5 -= 4294967295
            var v27 float64 = float64(0.1)
            if (716 == 997) {
                fmt.Println(v17, v16)
                var v28 bool = false
                v29 := GenS1{F0: 0, F1: 255}
                _ = v28
                _ = v29
            }
            _ = v27
        }
    }
    v17 = ((2 == 1) || true)
    v16 = (v1 + 9223372036854775807)
    delete(v4, 2147483647)
    _ = len(v3)
    for fi30 := 0; fi30 < 5; fi30++ {
        for fi31 := 0; fi31 < 2; fi31++ {
            v16 += v1
            _ = fi31
        }
        _ = fi30
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        var v33 int64 = (^(v1))
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        fmt.Println(v17)
        v35 := GenS3{F0: 229}
        _ = fi34
        _ = v35
    }
    v5 += v5
    for fi36 := 0; fi36 < 8; fi36++ {
        v2 = (v2 + "a")
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
