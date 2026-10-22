package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 bool }
type GenS2 struct { F0 int16; F1 uint; F2 uint16 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint64 }
func genFn0() int64 {
    for fi4 := 0; fi4 < 15; fi4++ {
        if (-9223372036854775808 >= -831) {
            var v5 int = (fi4 % (((-(fi4))) | 1))
            v5 += fi4
            _ = v5
        }
        _ = fi4
    }
    var v6 uint8 = 0
    v7 := GenS1{F0: 0, F1: true, F2: false}
    v8 := []int{128, 9223372036854775807, -417, -2147483648}
    _ = v6
    _ = v7
    _ = v8
    return -9223372036854775808
}
func genFn1(p0 uint8, p1 int16) int8 {
    v9 := GenS1{F0: 32768, F1: true, F2: (true && true)}
    var v10 uint32 = 2147483647
    var v11 uint16 = 42
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return -128
}

func main() {
    var v1 int = -2147483648
    var v2 int = v1
    var v3 bool = (!(-2147483648 != 49))
    switch (v1 ^ v2) {
    case 477:
        v12 := GenS1{F0: (-(v2)), F1: true, F2: (!(v3 || (("\n\t" + "hello") > "hello")))}
        _ = v12
    case 170:
        fmt.Println(v3, v1)
    }
    v13 := GenS1{F0: v2, F1: v3, F2: ((1.5 - 0.0) != ((1.5 / (1.5 / (3.14159 * (0.1 * 1.5)))) * -2.25))}
    var v14 int32 = 1
    fmt.Println(v2, v3, v3)
    v15 := []int{v1, v1}
    var v16 int8 = 127
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    v17 := GenS1{F0: (v1 >> ((790) & 63)), F1: (-2.25 < 0.0), F2: v3}
    _ = len(v15)
    switch "hello" {
    case "hello":
        var v18 int32 = (^(v14))
        _ = v18
    case "a":
        v2 = v2
    default:
        if v3 {
            var v19 uint8 = 127
            _ = v19
        } else {
            var v20 int16 = -32768
            if len(v15) > 0 {
                v15[0] = (v1 - v2)
            }
            if v3 {
                v1 = (v2 + v1)
                var v21 int64 = -9223372036854775808
                var v22 uint32 = 2147483647
                _ = v21
                _ = v22
            }
            _ = v20
        }
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        for fi24 := 0; fi24 < 6; fi24++ {
            switch v2 {
            case 32767:
                var v25 uint = 4294967295
                _ = v25
            case 127:
                if v3 {
                    if (false || (v3 && (1 >= 2))) {
                        v26 := GenS2{F0: -965, F1: 517, F2: 48}
                        v16 -= (v16 * v16)
                        var v27 bool = v3
                        _ = v26
                        _ = v27
                    }
                }
            default:
                var v28 int64 = 1
                _ = v28
            }
            switch (((("foo bar" + ("\n\t" + "hello")) + "\n\t") + "a") + ("hello" + "")) {
            case "":
                var v29 int64 = genFn0()
                _ = v29
            }
            if v3 {
                v3 = v3
                v16 -= v16
                if false {
                    fmt.Println(v3, v14)
                    v16 = (v16 ^ v16)
                }
            }
            v30 := GenS3{F0: 176, F1: genFn0()}
            switch "hello" {
            case "hello":
                var v31 int16 = 42
                _ = v31
            case "":
                if false {
                    v3 = v3
                    if v3 {
                        fmt.Println(fi24)
                        v32 := GenS2{F0: -1, F1: 1, F2: 255}
                        _ = v32
                    } else {
                        if len(v15) > 0 {
                            v15[0] = (fi23 / ((fi24) | 1))
                        }
                    }
                }
            }
            v33 := GenS4{F0: 2}
            _ = fi24
            _ = v30
            _ = v33
        }
        v2 = int(v16)
        _ = len(v15)
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
