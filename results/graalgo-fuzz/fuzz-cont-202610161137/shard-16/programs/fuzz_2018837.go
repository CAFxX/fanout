package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int32; F1 int16 }
func genFn0() uint16 {
    var v4 uint32 = 139
    var v5 float32 = 1.5
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: (^(v4)), F1: 196, F2: v4}
        _ = fi6
        _ = v7
    }
    v8 := []int{539, 4294967295, -9223372036854775808, 32768}
    _ = v4
    _ = v5
    _ = v8
    return 1
}
func genFn1() int8 {
    var v9 int16 = 32767
    v10 := []int{127, 1, 42, 32767}
    _ = v10
    _ = v9
    return -128
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[int]int)
    v3 := make(map[string]uint32)
    v11 := []int{}
    v12 := GenS1{F0: 413, F1: 193, F2: 493}
    v3[("foo bar" + ("a" + ("" + "\n\t")))] = 0
    v1 = v1
    if (true || false) {
        if true {
            v1 = v1
            var v13 float32 = ((((float32(((0.1 + 3.14159) * 0.1)) + (0.1 + 1e10)) * -2.25) * 3.14159) + float32(1.5))
            _ = v13
        }
        v14 := GenS1{F0: 2, F1: 1, F2: 2}
        _ = v14
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        fmt.Println(v1, fi15)
        v1 += (^(v1))
        _ = fi15
    }
    v16 := GenS2{F0: 1, F1: -1}
    v1 = v1
    v1 = (v1 % ((v1) | 1))
    for fi17 := 0; fi17 < 12; fi17++ {
        v1 = 2
        var v18 bool = true
        _ = fi17
        _ = v18
    }
    var v19 bool = true
    if len(v11) > 0 {
        v11[0] = 1
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        v21 := GenS1{F0: 65535, F1: 255, F2: 2147483647}
        for fi22 := 0; fi22 < 5; fi22++ {
            var v23 uint = 255
            if (192 > 5) {
                var v24 uint32 = 2147483648
                var v25 bool = false
                v26 := GenS1{F0: (v24 * v24), F1: 128, F2: (^(v24))}
                _ = v24
                _ = v25
                _ = v26
            } else {
                var v27 int32 = 888
                var v28 bool = (-2.25 > -2.25)
                _ = v27
                _ = v28
            }
            if true {
                fmt.Println(fi20, v23)
                v29 := GenS2{F0: 144, F1: -629}
                if (fi20 != fi22) {
                    fmt.Println(v23, fi20)
                    v30 := v3[(("foo bar" + ((("" + "hello") + ("foo bar" + "")) + ("foo bar" + "a"))) + "a")]
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    _ = v30
                } else {
                    var v31 int16 = 0
                    _ = v31
                }
                _ = v29
            } else {
                delete(v2, (fi20 >> ((v23) & 63)))
            }
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    var v32 float32 = 0.0
    v2[335] = 42
    fmt.Println(v1, v32)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v32
}
