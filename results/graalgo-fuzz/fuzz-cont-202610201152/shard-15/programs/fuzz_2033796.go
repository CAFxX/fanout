package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 uint64; F2 int8 }
func genFn0(p0 uint64) uint64 {
    v4 := GenS1{F0: 2, F1: 236, F2: p0}
    v5 := []int{-1, 124, 128}
    for fi6 := 0; fi6 < 3; fi6++ {
        fmt.Println(p0, fi6)
        var v7 int32 = -2147483648
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v4
    _ = v5
    return (^(p0))
}
func genFn1(p0 bool, p1 uint16) int32 {
    if (2147483648 != 65535) {
        v8 := []int{}
        if (!p0) {
            var v9 float64 = 1.5
            switch "foo bar" {
            case "\n\t":
                for fi10 := 0; fi10 < 14; fi10++ {
                    v11 := GenS1{F0: 65535, F1: 42, F2: 1}
                    p1 = (p1 << ((65535) & 63))
                    var v12 string = (("\n\t" + ("" + "\n\t")) + "")
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case "a":
                v13 := GenS1{F0: 488, F1: 139, F2: 201}
                _ = v13
            case "hello":
                fmt.Println(p1, p0)
            }
            _ = v9
        } else {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            fmt.Println(p0, p1)
        }
        _ = v8
    } else {
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := GenS2{F0: 4294967295}
            p0 = (!(-2.25 != 1.5))
            _ = fi14
            _ = v15
        }
    }
    p0 = (!(false || true))
    _ = p0
    _ = p1
    return -937
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]bool)
    var v3 uint = 65535
    var v16 int64 = -814
    switch "foo bar" {
    case "":
        var v17 int64 = v16
        _ = v17
    case "a":
        v18 := GenS3{F0: false}
        _ = v18
    }
    if (("a" >= "foo bar") && true) {
        v19 := GenS3{F0: (genFn1(false, 0) < 42)}
        v16 = (v16 + v16)
        _ = v19
    }
    var v20 int16 = 42
    fmt.Println(v20)
    var v21 uint64 = genFn0(4294967295)
    var v22 uint8 = v1
    var v23 uint = 155
    v3 = 4294967295
    v20 = (-(v20))
    switch "foo bar" {
    case "foo bar":
        if (true && (v16 == (v16 - 1))) {
            v21 += 18446744073709551615
            switch "foo bar" {
            case "hello":
                fmt.Println(v23)
            case "foo bar":
                var v24 int8 = int8(v23)
                _ = v24
            case "\n\t":
                var v25 int32 = -2147483648
                _ = v25
            default:
                v26 := GenS4{F0: (v23 - v23), F1: 4294967295, F2: 127}
                _ = v26
            }
        } else {
            var v27 uint16 = 2
            v21 = (v21 | uint64(v22))
            _ = v27
        }
    case "hello":
        if (674 != 65535) {
            var v28 bool = ((^(v1)) == (v1 >> ((v23) & 63)))
            v29 := []int{1, int(v22)}
            _ = v28
            _ = v29
        } else {
            var v30 bool = true
            _ = v30
        }
    }
    if false {
        if (2147483648 > 2) {
            fmt.Println(v23, v3, v21)
        }
    } else {
        v21 = (v21 - v21)
        switch "a" {
        case "hello":
            var v31 int16 = 42
            _ = v31
        case "a":
            fmt.Println(v21)
        case "foo bar":
            v2[-9223372036854775808] = (0.1 > (3.14159 * (1.5 * 3.14159)))
        }
    }
    var v32 int8 = -128
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v32
}
