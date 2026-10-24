package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 uint; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 uint8 }
func genFn0() bool {
    var v6 bool = (3.14159 > 3.14159)
    if false {
        v7 := GenS1{F0: 32767}
        _ = v7
    } else {
        if (223 >= 34) {
            switch 132 {
            case -9223372036854775808:
                v6 = (42 <= -2147483648)
            case 4294967295:
                v8 := GenS1{F0: 1000}
                _ = v8
            case -2147483648:
                fmt.Println(v6)
            }
            if v6 {
                var v9 float64 = 0.0
                var v10 uint64 = 700
                _ = v10
                _ = v9
            }
            v6 = false
        } else {
            v11 := GenS2{F0: -1, F1: 640, F2: 1}
            fmt.Println(v6)
            _ = v11
        }
        switch "\n\t" {
        case "":
            v12 := GenS1{F0: -642}
            _ = v12
        case "a":
            var v13 int8 = 1
            _ = v13
        default:
            v6 = (1 <= 42)
        }
    }
    v6 = true
    v14 := GenS2{F0: 110, F1: 229, F2: 1}
    _ = v14
    _ = v6
    return (!(892 < 255))
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]int)
    var v3 uint16 = 29
    var v4 int16 = 615
    var v5 int16 = (v4 - (-(v4)))
    if (!true) {
        if true {
            v3 = (v3 + v3)
            v15 := GenS3{F0: 230}
            var v16 bool = false
            _ = v15
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 9; fi17++ {
            v1 = (v1 ^ 14)
            _ = fi17
        }
    }
    v5 = -32768
    fmt.Println(v4, v3)
    switch 4294967295 {
    case 2147483647:
        fmt.Println(v4, v4)
    case -99:
        var v18 bool = false
        _ = v18
    case -2147483648:
        v19 := GenS3{F0: -9223372036854775808}
        _ = v19
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        v4 = v4
        var v21 bool = true
        switch fi20 {
        case -9223372036854775808:
            v22 := GenS3{F0: int64(v5)}
            _ = v22
        }
        _ = fi20
        _ = v21
    }
    var v23 int8 = -121
    if true {
        var v24 int8 = v23
        for fi25 := 0; fi25 < 9; fi25++ {
            for fi26 := 0; fi26 < 4; fi26++ {
                v27 := []int{-1, (fi26 >> ((uint(v3)) & 63))}
                v23 = 42
                if (528 < -2147483648) {
                    var v28 int64 = -9223372036854775808
                    var v29 int32 = 42
                    _ = v28
                    _ = v29
                } else {
                    if false {
                        fmt.Println(fi26, v4, v1)
                    }
                }
                if (1.5 < 0.0) {
                    v4 -= -259
                    v30 := GenS4{F0: -32768, F1: v1}
                    _ = v30
                }
                _ = fi26
                _ = v27
            }
            if (1 != 0) {
                v31 := GenS4{F0: 0, F1: 66}
                var v32 int64 = -9223372036854775808
                _ = v31
                _ = v32
            }
            _ = fi25
        }
        _ = v24
    }
    var v33 uint = 65535
    fmt.Println(v23, v1, v1)
    switch 9223372036854775807 {
    case 127:
        var v34 uint = v33
        _ = v34
    default:
        if (!(2 == 18446744073709551615)) {
            var v35 int32 = 2147483647
            _ = v35
        } else {
            var v36 int16 = -1
            var v37 bool = (true || (-128 > (v23 | v23)))
            v3 = (v3 ^ (v3 % ((v3) | 1)))
            _ = v36
            _ = v37
        }
    }
    v33 += 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
