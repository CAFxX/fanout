package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint64; F2 int }
type GenS3 struct { F0 string; F1 uint8; F2 uint64 }
func genFn0(p0 int) int32 {
    v4 := GenS1{F0: 2, F1: 100, F2: 42}
    if (false || (42 == 42)) {
        var v5 uint = 42
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := []int{(p0 / ((127) | 1))}
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    p0 = (p0 << ((856) & 63))
    _ = p0
    _ = v4
    return -1
}

func main() {
    var v1 string = "a"
    var v2 int32 = 2147483647
    var v3 int64 = 42
    var v8 uint8 = 211
    v3 += v3
    v9 := GenS1{F0: 18446744073709551615, F1: 42, F2: v8}
    v8 = 136
    switch -9223372036854775808 {
    case -9223372036854775808:
        var v10 uint8 = v8
        _ = v10
    case 844:
        if true {
            var v11 int = 32767
            _ = v11
        }
    default:
        switch (v1 + v1) {
        case "foo bar":
            if false {
                v12 := []int{4294967295, -154, -128}
                for fi13 := 0; fi13 < 4; fi13++ {
                    v3 -= v3
                    _ = fi13
                }
                _ = v12
            } else {
                v2 = v2
                var v14 bool = true
                fmt.Println(v3, v2, v2)
                _ = v14
            }
        case "\n\t":
            v15 := []int{1, 1, 798}
            _ = v15
        case "hello":
            v16 := GenS2{F0: v8, F1: 42, F2: 4294967295}
            _ = v16
        default:
            fmt.Println(v3, v2)
        }
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        v2 = genFn0(9223372036854775807)
        _ = fi17
    }
    var v18 uint8 = (v8 << ((1) & 63))
    switch v1 {
    case "":
        v8 = (v18 % ((v18) | 1))
    case "foo bar":
        var v19 int8 = -128
        _ = v19
    case "\n\t":
        for fi20 := 0; fi20 < 6; fi20++ {
            v21 := []int{fi20, 32768, 4294967295}
            v22 := GenS3{F0: ("\n\t" + "\n\t"), F1: v8, F2: 18446744073709551615}
            _ = fi20
            _ = v21
            _ = v22
        }
    default:
        var v23 int = -2147483648
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
