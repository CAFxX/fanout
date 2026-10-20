package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int; F2 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 string; F2 uint8 }
type GenS5 struct { F0 int16; F1 uint }
func genFn0() int16 {
    if true {
        var v6 int = -655
        if (127 <= -128) {
            v7 := GenS1{F0: 0}
            fmt.Println(v6)
            _ = v7
        } else {
            v8 := GenS2{F0: -926, F1: v6, F2: 2}
            _ = v8
        }
        _ = v6
    }
    var v9 uint16 = 14
    _ = v9
    return int16(v9)
}
func genFn1() int16 {
    switch -2147483648 {
    case 2147483647:
        if (!((42 > 42) && ((!(-2147483648 == 1)) && false))) {
            v10 := []int{2147483647}
            for fi11 := 0; fi11 < 19; fi11++ {
                if (!("hello" < "foo bar")) {
                    var v12 int32 = -809
                    fmt.Println(fi11, fi11)
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    _ = v12
                } else {
                    fmt.Println(fi11)
                    var v13 int8 = 1
                    var v14 bool = true
                    _ = v13
                    _ = v14
                }
                _ = len(v10)
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                _ = fi11
            }
            _ = v10
        } else {
            var v15 uint16 = 1
            v16 := GenS3{F0: 127}
            _ = v15
            _ = v16
        }
    default:
        v17 := GenS2{F0: 779, F1: 32767, F2: 574}
        _ = v17
    }
    if (("hello" + "") <= "\n\t") {
        v18 := GenS4{F0: 0, F1: "", F2: 128}
        _ = v18
    }
    if false {
        switch 128 {
        case 440:
            var v19 float32 = (0.1 + 3.14159)
            _ = v19
        case -926:
            var v20 uint32 = 0
            _ = v20
        case 127:
            v21 := GenS1{F0: 175}
            _ = v21
        }
    } else {
        var v22 bool = false
        _ = v22
    }
    return 0
}

func main() {
    var v1 uint16 = 32767
    v2 := []int{128}
    v3 := []int{242}
    v4 := []int{245, -2147483648, 0, -2147483648}
    var v5 int8 = 88
    if (42 <= 0) {
        v23 := GenS5{F0: 585, F1: 255}
        _ = v23
    }
    fmt.Println(v1)
    switch "a" {
    case "a":
        var v24 int16 = -278
        _ = v24
    default:
        v1 += 65535
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        v5 = v5
        _ = fi25
    }
    var v26 int16 = 461
    if (42 <= 2) {
        v5 = 45
    } else {
        v26 -= int16(v5)
    }
    v5 = (v5 | (v5 % (((v5 + (v5 + (v5 >> ((255) & 63))))) | 1)))
    if (!(1 < 255)) {
        var v27 uint16 = (^(v1))
        v3 = append(v3, 2147483647)
        var v28 uint64 = 4294967295
        _ = v27
        _ = v28
    }
    v26 = v26
    var v29 int = -156
    if false {
        var v30 uint64 = 0
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
