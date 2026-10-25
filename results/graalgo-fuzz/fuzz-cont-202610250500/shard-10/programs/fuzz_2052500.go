package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 string }
func genFn0() uint {
    var v5 float64 = ((3.14159 * (((-2.25 - float64(0.0)) + 1e10) / 0.1)) - 3.14159)
    if false {
        fmt.Println(v5)
        v5 -= (v5 * ((3.14159 + v5) / v5))
        var v6 float32 = float32(0.1)
        _ = v6
    } else {
        fmt.Println(v5)
        var v7 string = ""
        _ = v7
    }
    switch "" {
    case "":
        fmt.Println(v5)
    case "hello":
        if false {
            if false {
                fmt.Println(v5)
            }
            var v8 int64 = -9223372036854775808
            _ = v8
        } else {
            fmt.Println(v5)
        }
    }
    if (-2147483648 == 0) {
        for fi9 := 0; fi9 < 9; fi9++ {
            v5 = v5
            if (42 <= 1) {
                v5 = 0.0
                for fi10 := 0; fi10 < 16; fi10++ {
                    fmt.Println(v5)
                    v11 := []int{(fi9 / ((2147483647) | 1)), fi9}
                    var v12 uint64 = 42
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v13 := []int{(fi9 + fi9), (-(fi9)), -359, fi9}
                _ = v13
            }
            _ = fi9
        }
        v14 := GenS1{F0: 863, F1: 28}
        _ = v14
    } else {
        var v15 float32 = 0.1
        var v16 uint64 = 18446744073709551615
        _ = v15
        _ = v16
    }
    v5 = (v5 - 3.14159)
    _ = v5
    return 42
}

func main() {
    var v1 uint32 = 42
    var v2 uint64 = 0
    v3 := []int{}
    var v4 int32 = 800
    var v17 bool = (!(1 >= (v1 | (v1 % ((v1) | 1)))))
    fmt.Println(v17)
    fmt.Println(v4, v17)
    if len(v3) > 0 {
        v3[0] = 128
    }
    var v18 uint8 = 166
    if (true && (!(v17 || false))) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v4 = (v4 * int32(v1))
        _ = len(v3)
    } else {
        v18 = v18
    }
    v4 -= int32(v18)
    var v19 uint32 = (v1 << ((734) & 63))
    for fi20 := 0; fi20 < 4; fi20++ {
        if (v2 != v2) {
            var v21 int32 = (v4 ^ (-(v4)))
            _ = v21
        } else {
            var v22 string = "foo bar"
            var v23 float64 = ((((-2.25 + ((0.1 * 0.1) - 1e10)) / (3.14159 / 1.5)) / ((((3.14159 / 0.1) / 1.0) - float64((1e10 * (0.1 - 1e10)))) / -2.25)) + -2.25)
            _ = v22
            _ = v23
        }
        if v17 {
            if (((false || ((0 > -1) && (((!false) && ((v4 & (v4 ^ (v4 % (((v4 + -1)) | 1)))) != v4)) && true))) || (!(!(!("" <= "hello"))))) || ((v1 | (v19 >> ((57) & 63))) > v1)) {
                switch (fi20 % ((fi20) | 1)) {
                case 2147483647:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            } else {
                var v24 float32 = 1e10
                _ = v24
            }
            var v25 uint16 = 601
            _ = v25
        }
        _ = fi20
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v3 = append(v3, -9223372036854775808)
    _ = len(v3)
    var v26 int16 = -32768
    v27 := GenS2{F0: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
