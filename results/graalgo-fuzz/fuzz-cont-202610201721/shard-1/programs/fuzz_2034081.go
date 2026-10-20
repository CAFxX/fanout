package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int) uint8 {
    fmt.Println(p0)
    var v4 int = 704
    v5 := GenS1{F0: (p0 % (((p0 << ((1) & 63))) | 1))}
    if ((0.0 - 1e10) >= ((((3.14159 * 1e10) - (1.5 - (3.14159 + ((1.5 - 0.0) / 0.1)))) / (1e10 - 1.5)) / float64(((1e10 + ((0.0 - (1.5 * 0.1)) - (3.14159 - ((1.5 / 1e10) - 3.14159)))) * (1.5 / float32(0.1)))))) {
        var v6 int32 = 42
        fmt.Println(v6, v4)
        _ = v6
    }
    v7 := []int{(p0 & (v4 - (p0 | p0))), (v4 / (((p0 | -9223372036854775808)) | 1))}
    _ = p0
    _ = v4
    _ = v5
    _ = v7
    return 42
}
func genFn1(p0 int64, p1 uint8) int64 {
    p0 -= p0
    var v8 int16 = -1
    p0 = 836
    v9 := []int{}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return p0
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 18446744073709551615
    var v3 string = "hello"
    if false {
        v3 = v3
        v10 := GenS1{F0: 106}
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 uint8 = 127
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    var v13 int16 = 1
    v3 = "\n\t"
    switch "\n\t" {
    case "hello":
        delete(v1, "hello")
    }
    var v14 string = "foo bar"
    v14 = v3
    v3 = (v3 + (((v3 + v14) + v14) + v3))
    if false {
        if false {
            var v15 uint32 = 2
            _ = v15
        } else {
            if ((0.0 < 3.14159) && (-551 >= -1)) {
                for fi16 := 0; fi16 < 19; fi16++ {
                    v13 = v13
                    _ = fi16
                }
                var v17 int8 = -128
                v18 := GenS2{F0: (v17 * v17), F1: 0}
                _ = v17
                _ = v18
            } else {
                var v19 uint16 = 42
                v13 -= (v13 - -1)
                if false {
                    v3 = ""
                    var v20 int32 = 42
                    v21 := GenS3{F0: (^(v2)), F1: 42, F2: 671}
                    _ = v20
                    _ = v21
                }
                _ = v19
            }
            v22 := []int{}
            _ = v22
        }
        fmt.Println(v3, v2)
    }
    if (-2147483648 > 32768) {
        v23 := []int{}
        v14 = "hello"
        _ = v23
    } else {
        v13 = (v13 % ((v13) | 1))
        _ = len(v1)
        var v24 int32 = 317
        _ = v24
    }
    fmt.Println(v2, v2)
    if false {
        switch 127 {
        case 32767:
            var v25 uint = 2
            _ = v25
        case -9223372036854775808:
            if (!(true && (1 >= 0))) {
                var v26 uint64 = v2
                _ = len(v1)
                var v27 bool = (!(true && (false || true)))
                _ = v26
                _ = v27
            } else {
                for fi28 := 0; fi28 < 18; fi28++ {
                    v29 := []int{128, fi28}
                    var v30 uint16 = 65535
                    v14 = v3
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                v2 += (v2 >> ((90) & 63))
                var v31 uint = 1
                _ = v31
            }
        }
    }
    switch v3 {
    case "":
        for fi32 := 0; fi32 < 7; fi32++ {
            v33 := GenS4{F0: 0}
            if true {
                var v34 uint32 = 2147483647
                var v35 bool = ("\n\t" == v14)
                _ = v34
                _ = v35
            } else {
                v36 := GenS4{F0: 42}
                _ = v36
            }
            _ = fi32
            _ = v33
        }
    case "hello":
        var v37 string = v14
        _ = v37
    case "a":
        v38 := GenS5{F0: 667, F1: 42}
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
