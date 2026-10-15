package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint16; F1 bool; F2 int8 }
func genFn0(p0 int64) uint16 {
    for fi6 := 0; fi6 < 14; fi6++ {
        if false {
            v7 := []int{128, fi6}
            switch 0 {
            case 32767:
                fmt.Println(fi6, fi6)
            default:
                var v8 uint64 = 2
                _ = v8
            }
            v7 = append(v7, fi6)
            _ = v7
        }
        _ = fi6
    }
    p0 += p0
    var v9 int8 = 42
    var v10 int = 431
    _ = p0
    _ = v10
    _ = v9
    return 0
}
func genFn1(p0 uint16, p1 uint32) int16 {
    var v11 int32 = 422
    var v12 int64 = -1
    fmt.Println(v11, v12, p1)
    if false {
        v12 += (v12 & v12)
        v13 := []int{-2147483648, 9223372036854775807, 2147483647, -2147483648}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 706
}

func main() {
    var v1 int8 = 35
    v2 := []uint32{943, 204}
    v3 := make(map[uint32]string)
    var v4 uint8 = 0
    var v5 int = 128
    var v14 int16 = 1
    var v15 uint16 = 255
    if true {
        switch 42 {
        case 1:
            v2 = append(v2, 2)
        case 2147483647:
            for fi16 := 0; fi16 < 11; fi16++ {
                var v17 float32 = float32(1e10)
                var v18 uint64 = 4294967295
                _ = fi16
                _ = v17
                _ = v18
            }
        case 42:
            delete(v3, 4294967295)
        default:
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 uint16 = v15
                _ = fi19
                _ = v20
            }
        }
        v5 = (v5 << ((1) & 63))
    }
    fmt.Println(v15)
    v21 := v3[uint32(v15)]
    fmt.Println(v5, v15)
    v22 := GenS1{F0: v14, F1: 1}
    var v23 int64 = 0
    switch 541 {
    case 0:
        v1 = -49
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        for fi25 := 0; fi25 < 1; fi25++ {
            v26 := v3[4294967295]
            fmt.Println(v5, v15, v21)
            fmt.Println(v26)
            v26 = "hello"
            delete(v3, 42)
            v21 = v26
            v27 := GenS2{F0: (v15 - (v15 >> ((4294967295) & 63))), F1: (v1 > (v1 + (^(v1)))), F2: v1}
            if true {
                var v28 uint8 = v4
                _ = v28
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        switch "\n\t" {
        case "hello":
            v15 = v15
        case "\n\t":
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        case "":
            var v29 uint = 2
            _ = v29
        default:
            v2 = append(v2, 4294967295)
        }
        _ = fi24
    }
    var v30 string = ((v21 + "") + "foo bar")
    fmt.Println(v23, v5)
    v15 = (v15 / (((v15 / ((v15) | 1))) | 1))
    _ = len(v3)
    v31 := GenS2{F0: v15, F1: ((!true) || false), F2: int8(v15)}
    for fi32 := 0; fi32 < 9; fi32++ {
        var v33 int64 = (v23 | v23)
        switch (-(fi32)) {
        case -2147483648:
            var v34 uint16 = 32767
            _ = v34
        }
        if (v23 >= (^(v23))) {
            v35 := GenS1{F0: 42, F1: 449}
            fmt.Println(v15, v15)
            _ = v35
        }
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
