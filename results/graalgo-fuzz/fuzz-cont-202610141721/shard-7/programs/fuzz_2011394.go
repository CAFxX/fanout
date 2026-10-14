package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int }
func genFn0(p0 uint16, p1 int16) uint {
    var v4 uint = 255
    p1 = (p1 & p1)
    _ = p0
    _ = p1
    _ = v4
    return (v4 * v4)
}
func genFn1(p0 uint8, p1 int64, p2 uint32) uint8 {
    v5 := []int{}
    switch -9223372036854775808 {
    case -901:
        var v6 int16 = 846
        _ = v6
    }
    v7 := GenS1{F0: 42, F1: 1, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return (p0 ^ (p0 ^ 0))
}

func main() {
    var v1 uint16 = 397
    var v2 int8 = 42
    var v3 uint = 255
    var v8 uint64 = uint64(v3)
    v2 = (v2 + v2)
    v8 += v8
    var v9 uint = v3
    if (42 > 2147483647) {
        if true {
            for fi10 := 0; fi10 < 4; fi10++ {
                fmt.Println(v1, v3)
                v11 := GenS1{F0: uint(v1), F1: v9, F2: (fi10 / ((fi10) | 1))}
                for fi12 := 0; fi12 < 6; fi12++ {
                    v13 := []int{fi10}
                    var v14 int16 = 249
                    v15 := GenS1{F0: 2, F1: genFn0((v1 - v1), -32768), F2: 128}
                    v3 = v9
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                _ = fi10
                _ = v11
            }
        } else {
            var v16 uint32 = uint32(v8)
            fmt.Println(v3)
            _ = v16
        }
        if true {
            switch ("a" + "hello") {
            case "foo bar":
                if (-1 == -1) {
                    v9 = (v3 + uint(v8))
                    switch ("" + "foo bar") {
                    case "":
                        fmt.Println(v3)
                    case "\n\t":
                        var v17 uint64 = (v8 ^ (^(v8)))
                        _ = v17
                    case "hello":
                        v2 -= v2
                    }
                    for fi18 := 0; fi18 < 16; fi18++ {
                        fmt.Println(v1)
                        v19 := []int{32768}
                        _ = fi18
                        _ = v19
                    }
                } else {
                    for fi20 := 0; fi20 < 9; fi20++ {
                        v1 = v1
                        v8 = v8
                        _ = fi20
                    }
                    if true {
                        var v21 int32 = 42
                        var v22 int32 = v21
                        _ = v21
                        _ = v22
                    }
                }
            default:
                switch ("a" + "\n\t") {
                case "foo bar":
                    v23 := []int{633, -2147483648, -2147483648}
                    _ = v23
                }
            }
            fmt.Println(v9, v9, v3)
            v3 = uint(v8)
        }
        fmt.Println(v8, v3, v1)
    } else {
        v9 = (v9 | 1)
        v24 := []int{2147483647, 107, -2147483648}
        _ = v24
    }
    if (v1 <= (v1 + 732)) {
        var v25 uint8 = 128
        _ = v25
    }
    switch ("hello" + "\n\t") {
    case "hello":
        switch 2147483647 {
        case 127:
            var v26 uint64 = 347
            _ = v26
        case 128:
            fmt.Println(v3)
        case 32767:
            fmt.Println(v3)
        default:
            v9 = (v9 << ((0) & 63))
        }
    }
    var v27 int8 = 87
    var v28 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v8
    _ = v9
}
