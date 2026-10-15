package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint64) int32 {
    p0 += (p0 % ((p0) | 1))
    if (p0 < p0) {
        var v5 uint8 = 107
        p0 += (p0 << ((0) & 63))
        for fi6 := 0; fi6 < 20; fi6++ {
            v7 := []int{fi6, fi6, 42}
            _ = fi6
            _ = v7
        }
        _ = v5
    } else {
        p0 = p0
    }
    p0 -= (^(p0))
    v8 := []int{-605, 32767, 128}
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := GenS1{F0: 1, F1: 19, F2: 810}
        if (!(!true)) {
            p0 += (p0 * 204)
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v8
    return 42
}
func genFn1(p0 int, p1 uint32) uint {
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := []int{}
        for fi13 := 0; fi13 < 7; fi13++ {
            fmt.Println(p0, fi11, p0)
            fmt.Println(p1, fi13)
            _ = fi13
        }
        var v14 int32 = 215
        _ = fi11
        _ = v12
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        p1 -= (p1 / ((p1) | 1))
        switch fi15 {
        case -779:
            p1 = p1
        }
        _ = fi15
    }
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[string]uint32)
    var v3 string = (("\n\t" + (("hello" + "") + ((("hello" + ("\n\t" + ("foo bar" + "hello"))) + "a") + "hello"))) + "hello")
    v4 := make(map[uint32]string)
    v16 := v2["foo bar"]
    if false {
        v17 := []int{0, 42, 2147483647, -9223372036854775808}
        var v18 float32 = 3.14159
        _ = v17
        _ = v18
    }
    v1 = (v1 << ((2) & 63))
    v4[v16] = "a"
    for fi19 := 0; fi19 < 20; fi19++ {
        if ((^(v16)) < v16) {
            v3 = v3
            v20 := GenS1{F0: 255, F1: -1, F2: (v16 % (((^(v16))) | 1))}
            fmt.Println(v3, fi19, v3)
            _ = v20
        }
        _ = fi19
    }
    if (0.0 > (((0.0 + 1e10) / -2.25) - 3.14159)) {
        v16 += 1
        for fi21 := 0; fi21 < 2; fi21++ {
            v3 = (((((v3 + "hello") + v3) + v3) + "\n\t") + (("hello" + v3) + "foo bar"))
            var v22 string = v3
            for fi23 := 0; fi23 < 18; fi23++ {
                fmt.Println(v22, v16, v1)
                if true {
                    v16 = 4294967295
                    v1 = 255
                }
                if (true || (2 <= 4294967295)) {
                    fmt.Println(v16, fi21)
                    v1 = (^(v1))
                }
                var v24 uint32 = (v16 << ((2) & 63))
                _ = fi23
                _ = v24
            }
            _ = fi21
            _ = v22
        }
        delete(v2, v3)
    }
    if (128 == 1) {
        if false {
            var v25 int = 2147483647
            v26 := GenS1{F0: 65535, F1: -9223372036854775808, F2: v16}
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 7; fi27++ {
            fmt.Println(v1, fi27, v16)
            _ = fi27
        }
    }
    if (true && false) {
        v28 := GenS2{F0: v1, F1: 42}
        _ = v28
    }
    v29 := v2["\n\t"]
    v30 := GenS1{F0: 2, F1: 1, F2: 2}
    var v31 int = -2147483648
    delete(v4, v16)
    var v32 float32 = (1.5 + ((0.1 / float32(-2.25)) - float32(((0.0 * 1.5) * 0.0))))
    var v33 uint = 180
    v16 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
