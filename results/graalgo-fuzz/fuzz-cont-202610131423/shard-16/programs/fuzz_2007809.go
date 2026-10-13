package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint8 }
func genFn0(p0 int16, p1 int16) int64 {
    p1 = 854
    var v6 uint16 = 0
    switch 528 {
    case 128:
        if true {
            for fi7 := 0; fi7 < 19; fi7++ {
                v6 += (^(v6))
                p1 = int16(fi7)
                p0 += p0
                _ = fi7
            }
            v8 := GenS1{F0: 42}
            _ = v8
        } else {
            switch 1 {
            case 42:
                p0 = 1
            }
            if (1 == 1) {
                for fi9 := 0; fi9 < 11; fi9++ {
                    fmt.Println(p0, p0, v6)
                    _ = fi9
                }
                p1 += (p1 - (p1 % (((p1 - p1)) | 1)))
                v6 = (v6 ^ (v6 / ((32767) | 1)))
            }
            v10 := GenS1{F0: 1}
            _ = v10
        }
    case 42:
        var v11 int16 = p0
        _ = v11
    case -2147483648:
        v12 := []int{32767, 988, 32768}
        _ = v12
    default:
        fmt.Println(p1)
    }
    var v13 int16 = (p0 >> ((0) & 63))
    _ = p0
    _ = p1
    _ = v13
    _ = v6
    return -379
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 2
    v3 := make(map[uint32]int)
    var v4 uint64 = 823
    var v5 int16 = 0
    v2 = v2
    var v14 uint64 = 1
    var v15 int8 = -128
    v16 := []int{}
    for fi17 := 0; fi17 < 4; fi17++ {
        v4 -= (v14 - v14)
        fmt.Println(v4, v2)
        _ = fi17
    }
    _ = len(v16)
    var v18 int8 = 1
    v5 = (v5 | v5)
    v18 = int8(v5)
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 uint16 = 133
        var v21 uint8 = 127
        if ((1e10 - (((1.5 - 0.1) / -2.25) + (float64(0.0) * (float64((1e10 / (-2.25 + 1e10))) / (-2.25 / ((-2.25 + -2.25) * -2.25)))))) != 3.14159) {
            if true {
                var v22 int16 = -620
                v21 = (v21 << ((343) & 63))
                _ = v22
            } else {
                _ = len(v16)
                var v23 uint8 = 126
                v16 = append(v16, (fi19 << (((v2 % (((v2 % (((v2 << ((v2) & 63))) | 1))) | 1))) & 63)))
                _ = v23
            }
        } else {
            v24 := GenS1{F0: 709}
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            switch ("" + ((("foo bar" + "a") + ("a" + (("hello" + (("hello" + ("foo bar" + ("" + "foo bar"))) + "hello")) + "\n\t"))) + "")) {
            case "a":
                v21 += v21
            default:
                var v25 float32 = 0.0
                _ = v25
            }
            _ = v24
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    if false {
        for fi26 := 0; fi26 < 2; fi26++ {
            v14 -= v4
            fmt.Println(fi26, v15)
            _ = fi26
        }
    } else {
        for fi27 := 0; fi27 < 11; fi27++ {
            var v28 int16 = -32768
            var v29 uint8 = 42
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    v30 := GenS2{F0: v18, F1: 1}
    var v31 uint8 = 0
    v4 = (v14 >> ((v2) & 63))
    if (1 < 619) {
        if len(v16) > 0 {
            v16[0] = -2147483648
        }
    }
    var v32 uint16 = 141
    var v33 int64 = -1
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
