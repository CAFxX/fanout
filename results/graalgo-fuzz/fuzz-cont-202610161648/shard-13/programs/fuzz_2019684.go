package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int32 }
func genFn0() bool {
    var v4 int64 = -961
    v4 = -1
    var v5 uint64 = 42
    _ = v4
    _ = v5
    return (-987 <= -1)
}
func genFn1(p0 int32, p1 uint) int32 {
    v6 := []int{-943, 1, 301}
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int64 = 650
        var v9 uint64 = 4294967295
        _ = fi7
        _ = v8
        _ = v9
    }
    if (583 != 0) {
        for fi10 := 0; fi10 < 15; fi10++ {
            if false {
                var v11 uint = (^(p1))
                _ = v11
            }
            fmt.Println(fi10, p0)
            _ = fi10
        }
        v12 := GenS1{F0: p1, F1: 255}
        _ = v12
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 int64 = 9223372036854775807
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v6
    return p0
}

func main() {
    var v1 int64 = 8
    v2 := []int{319}
    v3 := []uint32{0, 2147483648}
    var v15 uint64 = 666
    v3 = append(v3, 42)
    if ((v1 << ((584) & 63)) <= v1) {
        for fi16 := 0; fi16 < 3; fi16++ {
            v17 := GenS1{F0: 0, F1: 127}
            fmt.Println(fi16, v15)
            _ = fi16
            _ = v17
        }
        if true {
            v3 = append(v3, 2147483647)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v2) > 0 {
                v2[0] = 127
            }
        }
    } else {
        fmt.Println(v15, v15)
        switch 9223372036854775807 {
        case -501:
            var v18 float32 = (((1e10 - float32(1.5)) / 1.5) + 1e10)
            _ = v18
        }
        fmt.Println(v15)
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        v2 = append(v2, -563)
        v15 += (v15 >> ((1) & 63))
        v15 = 0
        _ = fi19
    }
    if (false && false) {
        v3 = append(v3, 2147483647)
    }
    var v20 int16 = -1
    var v21 int64 = v1
    v15 = v15
    if true {
        var v22 uint32 = 4294967295
        _ = v22
    } else {
        var v23 int8 = 63
        var v24 int = 32767
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        if true {
            if (!(v20 >= v20)) {
                if (!(-2147483648 == 633)) {
                    v26 := GenS2{F0: 11, F1: v20}
                    v3 = append(v3, 398)
                    var v27 bool = true
                    _ = v26
                    _ = v27
                } else {
                    var v28 uint8 = 1
                    var v29 uint = 65535
                    var v30 float64 = (0.1 * ((((-2.25 + 1e10) - 0.1) - 3.14159) / 1e10))
                    _ = v28
                    _ = v29
                    _ = v30
                }
            }
            fmt.Println(fi25)
        }
        switch "foo bar" {
        case "a":
            v1 = (v21 << ((65535) & 63))
        case "hello":
            var v31 float32 = ((3.14159 + 3.14159) / -2.25)
            _ = v31
        case "":
            v21 -= v21
        }
        v32 := GenS3{F0: genFn1(1, 18446744073709551615)}
        _ = fi25
        _ = v32
    }
    var v33 int = 1
    v34 := GenS1{F0: 65535, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
    _ = v34
}
