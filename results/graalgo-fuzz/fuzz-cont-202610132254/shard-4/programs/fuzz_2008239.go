package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int64; F2 bool }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{}
    var v3 float64 = 3.14159
    var v4 int64 = 628
    v5 := make(map[int]bool)
    var v6 int64 = v4
    v4 += (v4 + (v4 * (v6 << ((969) & 63))))
    v3 = v3
    v7 := GenS1{F0: 710}
    if false {
        v3 += (1e10 - (((((-2.25 * v3) * -2.25) - 1e10) * -2.25) - v3))
        v3 += -2.25
        if (65535 == 0) {
            if true {
                v6 -= 42
            } else {
                var v8 int32 = -1
                _ = v8
            }
            if (753 < 2) {
                if false {
                    v9 := GenS1{F0: 2}
                    _ = v9
                }
                v10 := v5[128]
                _ = v10
            }
        }
    } else {
        var v11 int32 = 42
        _ = v11
    }
    v6 += 1
    _ = len(v2)
    var v12 float32 = 3.14159
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 bool = false
        v15 := GenS1{F0: 1}
        switch ("hello" + (("\n\t" + ("foo bar" + "a")) + "hello")) {
        case "a":
            var v16 uint8 = 127
            _ = v16
        default:
            v17 := GenS1{F0: 0}
            _ = v17
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v3 += 0.0
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        v4 += 9223372036854775807
        var v19 float64 = ((v3 - 0.1) * (v3 / -2.25))
        _ = fi18
        _ = v19
    }
    switch 926 {
    case -1:
        switch "\n\t" {
        case "a":
            var v20 int64 = v4
            _ = v20
        default:
            var v21 float64 = -2.25
            _ = v21
        }
    case 1:
        if (0 > 127) {
            v5[4294967295] = (((v4 >= v6) && true) || (!true))
        } else {
            if false {
                var v22 int64 = 0
                var v23 float32 = ((v12 / v12) + 0.0)
                _ = v22
                _ = v23
            }
        }
    default:
        var v24 int32 = 42
        _ = v24
    }
    var v25 int64 = v4
    if false {
        v5[32768] = (!(-1 < 360))
        var v26 int64 = (-(v6))
        for fi27 := 0; fi27 < 20; fi27++ {
            for fi28 := 0; fi28 < 3; fi28++ {
                delete(v1, 139)
                v6 += (^(v4))
                v29 := GenS2{F0: 2147483647}
                _ = fi28
                _ = v29
            }
            _ = fi27
        }
        _ = v26
    } else {
        v4 = 42
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        fmt.Println(v25, v4, v3)
        switch "a" {
        case "":
            v25 += v4
        case "\n\t":
            v31 := GenS3{F0: v25, F1: 42, F2: (!true)}
            _ = v31
        case "hello":
            var v32 int64 = (v4 & v25)
            _ = v32
        default:
            var v33 int = (fi30 << ((2) & 63))
            _ = v33
        }
        fmt.Println(v3, v12)
        _ = fi30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
