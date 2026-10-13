package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 string; F2 uint8 }
type GenS4 struct { F0 string }
func genFn0(p0 int16, p1 uint16, p2 int32) uint64 {
    v7 := []int{127, 42}
    var v8 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 836
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"hello", "a", ""}
    var v3 int16 = -1
    v4 := []int{1, 2147483647, 2147483647}
    var v5 uint8 = 255
    var v6 string = "hello"
    var v9 uint = 4294967295
    switch "hello" {
    case "":
        _ = len(v4)
    default:
        v10 := GenS1{F0: v9}
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v5 = 1
        v9 = (v9 << (((v9 >> ((v9) & 63))) & 63))
        if len(v2) > 0 {
            v2[0] = v6
        }
        _ = fi11
    }
    var v12 uint = (v9 - 255)
    v3 += int16(v12)
    v13 := GenS1{F0: v12}
    var v14 int8 = 58
    v15 := GenS2{F0: ((v6 + v6) + v6)}
    if true {
        fmt.Println(v12, v12)
    } else {
        v16 := GenS2{F0: v6}
        for fi17 := 0; fi17 < 13; fi17++ {
            v3 -= -32768
            _ = fi17
        }
        _ = v16
    }
    _ = len(v1)
    for fi18 := 0; fi18 < 13; fi18++ {
        if true {
            for fi19 := 0; fi19 < 19; fi19++ {
                var v20 bool = (!false)
                var v21 int = fi19
                _ = fi19
                _ = v20
                _ = v21
            }
            v5 = v5
        } else {
            var v22 int64 = -1
            v23 := v1[fi18]
            _ = v22
            _ = v23
        }
        v14 = (-(v14))
        v1[42] = ((553 < 42) || (676 != 1))
        _ = fi18
    }
    var v24 int16 = (v3 & v3)
    if false {
        fmt.Println(v6, v24, v3)
    } else {
        switch 9223372036854775807 {
        case 128:
            var v25 int = 9223372036854775807
            _ = v25
        case 42:
            v26 := GenS3{F0: int8(v24), F1: "foo bar", F2: (v5 % (((v5 / (((^(v5))) | 1))) | 1))}
            _ = v26
        case 889:
            v27 := GenS4{F0: "hello"}
            _ = v27
        default:
            fmt.Println(v24, v14, v3)
        }
        var v28 string = (v6 + "")
        _ = v28
    }
    v29 := GenS2{F0: "foo bar"}
    var v30 int16 = (v24 >> ((v12) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
