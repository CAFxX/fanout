package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int32, p1 string, p2 int) uint16 {
    var v5 uint = 803
    v6 := []int{(^(p2))}
    p2 -= (p2 << (((v5 ^ v5)) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 65535
}
func genFn1() string {
    v7 := []int{184, 1}
    var v8 int32 = 0
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"hello", "foo bar", "hello"}
    var v3 uint8 = 1
    v4 := make(map[uint32]bool)
    for fi9 := 0; fi9 < 20; fi9++ {
        if true {
            v10 := GenS1{F0: "", F1: 4294967295, F2: 4294967295}
            if false {
                var v11 uint32 = 2147483648
                _ = v11
            } else {
                if false {
                    v3 += v3
                }
                v3 -= (^(v3))
                var v12 int8 = 0
                _ = v12
            }
            var v13 int16 = 1000
            _ = v10
            _ = v13
        }
        v3 = (v3 / ((v3) | 1))
        _ = fi9
    }
    v3 = (v3 - v3)
    v3 = (v3 / ((v3) | 1))
    v4[2147483648] = false
    v14 := GenS1{F0: "hello", F1: 42, F2: 1}
    for fi15 := 0; fi15 < 10; fi15++ {
        v3 += v3
        v3 -= v3
        fmt.Println(v3)
        _ = fi15
    }
    fmt.Println(v3)
    v16 := GenS2{F0: 1}
    var v17 string = ("a" + "a")
    if (1000 < 42) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v17 = ""
    } else {
        var v18 string = v17
        var v19 int8 = 1
        v17 = ((v18 + (v18 + "foo bar")) + ("hello" + ""))
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        v17 = (((((v17 + (v17 + (v17 + "a"))) + v17) + v17) + v17) + "\n\t")
        for fi21 := 0; fi21 < 8; fi21++ {
            v3 -= (^(v3))
            fmt.Println(fi20, v17, fi21)
            var v22 float32 = float32(1.5)
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    var v23 float32 = -2.25
    switch 2147483647 {
    case -230:
        var v24 uint8 = (v3 * (v3 >> ((519) & 63)))
        _ = v24
    case 416:
        var v25 uint = 65535
        _ = v25
    case 475:
        if len(v2) > 0 {
            v2[0] = ("" + (v17 + ""))
        }
    }
    _ = len(v2)
    if (92 == 101) {
        var v26 float64 = 1.5
        var v27 string = "foo bar"
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        v3 += (v3 * v3)
        v17 = v17
        fmt.Println(v23, fi28)
        _ = fi28
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
