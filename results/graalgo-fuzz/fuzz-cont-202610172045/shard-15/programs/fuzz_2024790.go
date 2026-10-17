package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int8 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 string; F1 int32; F2 uint32 }
type GenS4 struct { F0 int64; F1 int64; F2 uint }
func genFn0(p0 int16) int {
    var v5 int8 = -128
    for fi6 := 0; fi6 < 19; fi6++ {
        p0 = p0
        switch fi6 {
        case -526:
            if (!(-647 != -9223372036854775808)) {
                v5 = (v5 << ((1) & 63))
                var v7 bool = true
                _ = v7
            }
        }
        v8 := GenS1{F0: 32767, F1: 445, F2: v5}
        _ = fi6
        _ = v8
    }
    v5 += (v5 - v5)
    for fi9 := 0; fi9 < 7; fi9++ {
        switch int(p0) {
        case -602:
            var v10 int64 = -314
            _ = v10
        }
        for fi11 := 0; fi11 < 6; fi11++ {
            p0 = p0
            var v12 int32 = -372
            v13 := []int{(fi11 % ((-127) | 1)), fi11}
            _ = fi11
            _ = v12
            _ = v13
        }
        v5 = v5
        _ = fi9
    }
    _ = p0
    _ = v5
    return -564
}
func genFn1(p0 uint64) uint32 {
    if true {
        fmt.Println(p0)
    } else {
        var v14 int64 = -383
        v15 := []int{genFn0(793)}
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: 124, F1: 674}
    p0 -= (p0 * 832)
    fmt.Println(p0)
    var v17 uint8 = 0
    _ = p0
    _ = v16
    _ = v17
    return 42
}

func main() {
    var v1 string = "foo bar"
    var v2 uint = 2
    v3 := []string{((((v1 + (v1 + ((("\n\t" + (((v1 + v1) + (v1 + "\n\t")) + "")) + (v1 + (v1 + v1))) + v1))) + ("hello" + v1)) + v1) + (v1 + (("hello" + ((v1 + v1) + "foo bar")) + "foo bar"))), ("hello" + ((v1 + (v1 + v1)) + "hello"))}
    var v4 string = ""
    var v18 uint32 = 2147483648
    v19 := GenS3{F0: "\n\t", F1: 42, F2: (v18 >> ((237) & 63))}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 int8 = -128
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi20
        _ = v21
    }
    fmt.Println(v1, v1, v1)
    var v22 int32 = 1
    var v23 uint8 = 42
    v24 := GenS4{F0: int64(v22), F1: 42, F2: v2}
    var v25 uint64 = 691
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    fmt.Println(v25, v23, v1)
    var v26 float32 = 3.14159
    _ = len(v3)
    var v27 uint = (v2 * v2)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
