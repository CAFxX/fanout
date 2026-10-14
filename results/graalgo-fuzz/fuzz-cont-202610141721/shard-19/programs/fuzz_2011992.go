package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int }
type GenS5 struct { F0 uint16; F1 int64; F2 uint }
func genFn0() int8 {
    var v7 bool = (-9223372036854775808 < 42)
    if v7 {
        for fi8 := 0; fi8 < 4; fi8++ {
            if (0 != 42) {
                v9 := []int{fi8, 2147483647, fi8}
                _ = v9
            } else {
                for fi10 := 0; fi10 < 10; fi10++ {
                    switch (("foo bar" + (("foo bar" + (("\n\t" + ("\n\t" + "foo bar")) + ("" + ("" + (("a" + "hello") + ("" + "hello")))))) + "\n\t")) + "\n\t") {
                    case "foo bar":
                        var v11 string = "a"
                        _ = v11
                    case "":
                        fmt.Println(fi10, v7)
                    case "a":
                        v12 := GenS1{F0: -48, F1: 892, F2: fi10}
                        v7 = (-2147483648 >= fi10)
                        v7 = v7
                        _ = v12
                    }
                    if v7 {
                        var v13 uint = 65535
                        _ = v13
                    } else {
                        v14 := GenS2{F0: 934, F1: 4, F2: 2}
                        _ = v14
                    }
                    switch "a" {
                    case "foo bar":
                        fmt.Println(fi10, v7, fi8)
                    default:
                        fmt.Println(fi10)
                    }
                    _ = fi10
                }
                v15 := []int{}
                fmt.Println(v7)
                _ = v15
            }
            v16 := []int{fi8, fi8}
            _ = fi8
            _ = v16
        }
        v17 := []int{231}
        _ = v17
    }
    _ = v7
    return -1
}

func main() {
    var v1 uint32 = 2
    var v2 string = ""
    v3 := []int{127, 4294967295}
    v4 := make(map[string]int)
    var v5 int = 36
    var v6 int16 = -975
    if true {
        if ((float32(float64((1.5 - 1.5))) <= (3.14159 - (float32(0.0) - float32(1e10)))) && false) {
            var v18 uint16 = 644
            v2 = "hello"
            v1 = (v1 - v1)
            _ = v18
        } else {
            var v19 uint = 4294967295
            v19 -= 2
            v6 -= v6
            _ = v19
        }
        _ = len(v3)
    } else {
        if true {
            v4["foo bar"] = v5
            var v20 int64 = -9223372036854775808
            v1 -= (v1 >> ((65535) & 63))
            _ = v20
        }
        v6 = (v6 & -1)
        for fi21 := 0; fi21 < 6; fi21++ {
            var v22 uint32 = v1
            _ = fi21
            _ = v22
        }
    }
    v23 := GenS1{F0: -262, F1: 329, F2: (v5 | (v5 / ((v5) | 1)))}
    var v24 uint32 = v1
    v25 := GenS3{F0: uint16(v24), F1: 4294967295}
    var v26 int64 = 1
    for fi27 := 0; fi27 < 20; fi27++ {
        v3 = append(v3, -9223372036854775808)
        _ = fi27
    }
    fmt.Println(v1)
    v28 := GenS4{F0: 2, F1: 0, F2: 656}
    if true {
        v29 := GenS2{F0: 0, F1: 170, F2: 42}
        _ = v29
    } else {
        for fi30 := 0; fi30 < 8; fi30++ {
            v5 -= (v5 * 127)
            fmt.Println(v2, v6)
            v31 := GenS4{F0: 524, F1: (v24 << ((255) & 63)), F2: (v5 + int(v1))}
            _ = fi30
            _ = v31
        }
        for fi32 := 0; fi32 < 6; fi32++ {
            for fi33 := 0; fi33 < 8; fi33++ {
                v24 += v1
                v34 := GenS5{F0: 746, F1: int64(v24), F2: 1}
                _ = len(v3)
                _ = fi33
                _ = v34
            }
            if len(v3) > 0 {
                v3[0] = (v5 - (fi32 + 1))
            }
            _ = fi32
        }
    }
    var v35 int64 = v26
    v36 := GenS2{F0: 1, F1: 127, F2: 128}
    for fi37 := 0; fi37 < 15; fi37++ {
        if (((((1.5 / 3.14159) + float32(1.5)) + 0.0) - -2.25) == 0.1) {
            switch "hello" {
            case "foo bar":
                for fi38 := 0; fi38 < 2; fi38++ {
                    fmt.Println(v2, v24, v2)
                    fmt.Println(v35)
                    _ = fi38
                }
            case "a":
                v5 = v5
            case "hello":
                v26 -= 1
            }
            v39 := v4[v2]
            fmt.Println(v35, v26)
            _ = v39
        }
        v2 = ("\n\t" + v2)
        _ = fi37
    }
    v35 = v35
    v26 = v35
    var v40 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
