package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint64; F2 bool }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS4 struct { F0 string; F1 string }

func main() {
    v1 := []int{}
    var v2 uint = 18446744073709551615
    var v3 string = ("\n\t" + "foo bar")
    var v4 uint8 = 67
    var v5 bool = true
    var v6 uint8 = (v4 + (v4 - v4))
    fmt.Println(v5, v2, v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 = true
    if (!false) {
        if (42 != 255) {
            v7 := GenS1{F0: (0 > 0)}
            var v8 int = -744
            _ = v7
            _ = v8
        }
        v9 := GenS2{F0: -76, F1: 811, F2: false}
        v6 -= v4
        _ = v9
    }
    var v10 float32 = 3.14159
    if (v5 || (v5 || (2 != 228))) {
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := GenS3{F0: (^(v4)), F1: 658, F2: 1}
            v13 := GenS3{F0: (v4 & v6), F1: 4294967295, F2: 4294967295}
            _ = fi11
            _ = v12
            _ = v13
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v14 := GenS4{F0: "foo bar", F1: (v3 + "a")}
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        if false {
            var v16 uint = v2
            _ = v16
        } else {
            v3 = v3
            fmt.Println(v2, v2, v4)
            if (!true) {
                _ = len(v1)
                v10 = v10
                if (v5 || (-968 >= 880)) {
                    v4 = (v6 / ((v4) | 1))
                    v1 = append(v1, 1)
                    var v17 float32 = v10
                    _ = v17
                } else {
                    if v5 {
                        var v18 uint32 = 829
                        var v19 uint16 = 2
                        _ = v18
                        _ = v19
                    } else {
                        var v20 uint32 = 1
                        _ = v20
                    }
                }
            }
        }
        _ = fi15
    }
    v21 := GenS1{F0: false}
    v22 := GenS2{F0: 100, F1: 42, F2: (-823 <= -1)}
    fmt.Println(v10)
    if (65535 == 2147483647) {
        fmt.Println(v6)
        var v23 int64 = 42
        _ = v23
    }
    if v5 {
        _ = len(v1)
        if v5 {
            switch v3 {
            case "hello":
                for fi24 := 0; fi24 < 11; fi24++ {
                    var v25 string = v3
                    var v26 int32 = -1
                    for fi27 := 0; fi27 < 17; fi27++ {
                        v3 = ("\n\t" + (("foo bar" + v25) + "a"))
                        _ = fi27
                    }
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            case "foo bar":
                v2 -= 42
            default:
                v5 = true
            }
        } else {
            switch v3 {
            case "hello":
                fmt.Println(v4, v5)
            case "":
                fmt.Println(v2)
            case "a":
                fmt.Println(v4, v6, v6)
            }
            for fi28 := 0; fi28 < 15; fi28++ {
                v6 = v4
                _ = fi28
            }
            if (0 == 0) {
                v10 -= (((v10 * (((v10 / (v10 + v10)) + v10) + v10)) + 1e10) - v10)
                if false {
                    var v29 bool = v5
                    v2 = (v2 * v2)
                    _ = v29
                }
            }
        }
    } else {
        v10 += 1e10
        v6 = (^(v6))
        var v30 int16 = 885
        _ = v30
    }
    var v31 string = v3
    v1 = append(v1, 32767)
    var v32 int = 4294967295
    if ((^(v32)) > (-(v32))) {
        v10 += v10
        fmt.Println(v32, v4)
    }
    var v33 int8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
