package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int8 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 uint64; F2 uint64 }
func genFn0(p0 uint16) int8 {
    p0 -= 42
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: 279, F1: fi6, F2: 42}
        _ = fi6
        _ = v7
    }
    if (127 >= 1) {
        if false {
            var v8 bool = true
            _ = v8
        } else {
            p0 += p0
        }
        v9 := []int{-142, 2147483647, -134}
        _ = v9
    } else {
        var v10 int = -9223372036854775808
        v11 := GenS1{F0: 65, F1: v10, F2: 1}
        var v12 uint = 18446744073709551615
        _ = v10
        _ = v11
        _ = v12
    }
    _ = p0
    return -1
}
func genFn1(p0 uint, p1 uint16, p2 uint8) uint8 {
    var v13 uint8 = (^(p2))
    v14 := []int{-9223372036854775808, -494, 0}
    var v15 string = ""
    v13 = v13
    switch v15 {
    case "":
        var v16 uint = p0
        _ = v16
    case "hello":
        switch v15 {
        case "hello":
            p0 = (p0 % (((p0 / ((p0) | 1))) | 1))
        case "":
            var v17 int64 = 42
            _ = v17
        default:
            v18 := GenS1{F0: (p0 * (p0 - p0)), F1: -2147483648, F2: genFn0((p1 / ((p1) | 1)))}
            _ = v18
        }
    case "a":
        if ((^(p1)) > (p1 & (p1 << ((2) & 63)))) {
            for fi19 := 0; fi19 < 1; fi19++ {
                if len(v14) > 0 {
                    v14[0] = (fi19 - fi19)
                }
                _ = fi19
            }
        } else {
            p2 = 128
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return uint8(p0)
}

func main() {
    var v1 uint32 = 323
    var v2 int64 = 1
    var v3 int8 = 127
    var v4 float32 = 3.14159
    var v5 int16 = -1
    var v20 float64 = -2.25
    v21 := GenS1{F0: 65535, F1: 4294967295, F2: (^(v3))}
    v22 := []int{-2147483648, 834, 32767, 4294967295}
    fmt.Println(v5, v4, v20)
    v23 := GenS2{F0: true, F1: (v5 | (v5 ^ (^(v5))))}
    v24 := GenS3{F0: 111}
    if true {
        for fi25 := 0; fi25 < 16; fi25++ {
            v26 := GenS1{F0: 42, F1: fi25, F2: v3}
            fmt.Println(v3)
            switch "foo bar" {
            case "a":
                var v27 uint64 = 18446744073709551615
                _ = v27
            case "hello":
                switch fi25 {
                case 32768:
                    v1 = (v1 + (^(v1)))
                }
            }
            _ = fi25
            _ = v26
        }
        v28 := GenS1{F0: 204, F1: 345, F2: v3}
        _ = v28
    }
    fmt.Println(v20)
    switch "foo bar" {
    case "hello":
        if (42 <= 18446744073709551615) {
            var v29 uint32 = v1
            _ = v29
        } else {
            v1 += v1
        }
    case "foo bar":
        v5 = (v5 & -1)
    }
    var v30 int16 = int16(v2)
    var v31 int8 = v3
    v22 = append(v22, 594)
    var v32 int8 = (v31 - v3)
    for fi33 := 0; fi33 < 15; fi33++ {
        if false {
            v32 += int8(v2)
            var v34 float32 = 0.1
            if len(v22) > 0 {
                v22[0] = fi33
            }
            _ = v34
        }
        var v35 float64 = -2.25
        _ = fi33
        _ = v35
    }
    if len(v22) > 0 {
        v22[0] = -145
    }
    for fi36 := 0; fi36 < 2; fi36++ {
        v37 := GenS4{F0: ("foo bar" != "\n\t"), F1: 0, F2: 42}
        _ = fi36
        _ = v37
    }
    if (4294967295 < 18446744073709551615) {
        v4 -= 0.0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
