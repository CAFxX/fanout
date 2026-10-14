package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint }
func genFn0(p0 int8) uint64 {
    var v6 uint = 2
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{362, (fi7 | fi7), (fi7 - fi7)}
        if false {
            v6 -= 181
        } else {
            if (203 < 65535) {
                p0 -= (p0 * (p0 << ((42) & 63)))
                if true {
                    var v9 uint8 = 1
                    _ = v9
                }
                fmt.Println(fi7, v6, p0)
            }
        }
        if (true || (true || (0 <= 0))) {
            if (0 > -995) {
                switch (fi7 + (fi7 + 32767)) {
                case 258:
                    v6 = 65535
                case 937:
                    v6 = v6
                case 4294967295:
                    v8 = append(v8, fi7)
                default:
                    v8 = append(v8, -9223372036854775808)
                }
            } else {
                var v10 float32 = 0.1
                var v11 float32 = 1e10
                p0 -= (p0 << ((v6) & 63))
                _ = v10
                _ = v11
            }
            if len(v8) > 0 {
                v8[0] = fi7
            }
        }
        _ = fi7
        _ = v8
    }
    v12 := GenS1{F0: false, F1: -555}
    _ = p0
    _ = v12
    _ = v6
    return 42
}
func genFn1(p0 string) int {
    v13 := []int{384, 42, -376, 406}
    v14 := GenS2{F0: 0}
    v15 := GenS2{F0: 42}
    _ = p0
    _ = v13
    _ = v14
    _ = v15
    return -720
}

func main() {
    var v1 int = -569
    var v2 float64 = float64(float32(3.14159))
    v3 := make(map[int]int)
    v4 := []int{}
    var v5 uint32 = 0
    if false {
        var v16 int32 = 1
        var v17 float64 = float64(0.0)
        v3[(v1 - (v1 / ((v1) | 1)))] = genFn1(("hello" + ("\n\t" + (((("\n\t" + ("foo bar" + "hello")) + ("\n\t" + "foo bar")) + "\n\t") + "hello"))))
        _ = v16
        _ = v17
    }
    v2 = ((v2 - 0.0) + v2)
    switch 0 {
    case 1:
        v1 -= (v1 ^ v1)
    default:
        if (0 > uint8(v5)) {
            v2 -= (v2 - 3.14159)
        } else {
            if (!(78 != (v5 ^ 0))) {
                for fi18 := 0; fi18 < 16; fi18++ {
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v5 -= (^(v5))
                        var v20 float32 = 3.14159
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
                var v21 uint16 = 2
                _ = v21
            } else {
                var v22 uint8 = 223
                switch v1 {
                case 687:
                    delete(v3, (v1 >> ((2) & 63)))
                }
                _ = v22
            }
            v23 := GenS2{F0: 290}
            _ = v23
        }
    }
    if false {
        v24 := GenS2{F0: 32767}
        v2 += v2
        v1 = 982
        _ = v24
    } else {
        switch (v1 - 720) {
        case 128:
            v25 := GenS2{F0: 1}
            _ = v25
        case 32768:
            v2 = v2
        }
    }
    var v26 uint64 = 1
    switch v1 {
    case 173:
        var v27 uint = 4294967295
        _ = v27
    case 0:
        var v28 float32 = 0.1
        _ = v28
    default:
        if false {
            v29 := GenS1{F0: false, F1: -915}
            v4 = append(v4, v1)
            _ = v29
        } else {
            var v30 uint8 = 132
            _ = v30
        }
    }
    v4 = append(v4, genFn1(("\n\t" + ("hello" + "hello"))))
    v2 = 3.14159
    v31 := GenS1{F0: ((v5 + v5) <= (v5 ^ (^(v5)))), F1: -2147483648}
    v32 := GenS3{F0: (^(v26)), F1: 1}
    fmt.Println(v26)
    for fi33 := 0; fi33 < 4; fi33++ {
        fmt.Println(v2)
        for fi34 := 0; fi34 < 5; fi34++ {
            v3[(fi33 + v1)] = (v1 * v1)
            var v35 uint16 = 952
            _ = fi34
            _ = v35
        }
        v1 -= 9223372036854775807
        _ = fi33
    }
    v26 = v26
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
