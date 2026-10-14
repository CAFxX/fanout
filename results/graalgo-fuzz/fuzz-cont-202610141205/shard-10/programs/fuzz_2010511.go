package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint16 }
func genFn0() int32 {
    var v6 float64 = (0.1 * ((0.1 - 3.14159) * ((1e10 * -2.25) * 1e10)))
    v6 += v6
    fmt.Println(v6)
    _ = v6
    return -2147483648
}
func genFn1(p0 bool) int16 {
    v7 := GenS1{F0: 42}
    v8 := []int{-9223372036854775808, 127, 0, 0}
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint32 = 919
    var v3 int8 = -49
    v4 := make(map[string]int)
    var v5 int32 = (v1 << ((18446744073709551615) & 63))
    v4[(("" + "a") + (("hello" + ("\n\t" + "")) + "\n\t"))] = 107
    v9 := GenS2{F0: false, F1: 959}
    switch ("" + ((("" + ((("foo bar" + ("" + "hello")) + "a") + ("a" + ((("foo bar" + "foo bar") + "foo bar") + "\n\t")))) + "") + "")) {
    case "hello":
        if true {
            v10 := GenS1{F0: v3}
            _ = v10
        } else {
            v11 := []int{32767}
            v5 = v1
            v12 := GenS1{F0: v3}
            _ = v11
            _ = v12
        }
    }
    v13 := GenS3{F0: -9223372036854775808, F1: true}
    v14 := GenS3{F0: 9223372036854775807, F1: (840 <= 2)}
    v1 -= v5
    if (974 >= uint(v2)) {
        var v15 uint16 = 0
        v5 -= v1
        v16 := GenS4{F0: 1, F1: (v15 | v15), F2: 2}
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 9; fi17++ {
            v18 := GenS2{F0: true, F1: 0}
            v5 = int32(fi17)
            _ = fi17
            _ = v18
        }
    }
    if ((!false) || (330 != 18446744073709551615)) {
        v19 := []int{}
        var v20 uint32 = 42
        v5 -= v5
        _ = v19
        _ = v20
    }
    if (!(3.14159 <= 3.14159)) {
        v1 = v1
    }
    v3 -= -19
    var v21 int64 = 0
    v3 -= v3
    if true {
        v22 := GenS1{F0: v3}
        var v23 string = ""
        _ = v22
        _ = v23
    }
    v5 = v5
    var v24 uint = 145
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
