package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 bool; F2 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32; F1 bool }

func main() {
    var v1 string = ((("a" + "hello") + "foo bar") + "foo bar")
    var v2 bool = false
    var v3 uint8 = 127
    v4 := make(map[uint32]bool)
    var v5 int = 127
    var v6 int32 = 0
    var v7 int32 = (v6 - v6)
    var v8 bool = (!(1e10 < 1e10))
    fmt.Println(v1)
    if (v8 && false) {
        v5 = v5
        for fi9 := 0; fi9 < 3; fi9++ {
            if false {
                delete(v4, 42)
                v1 = ("foo bar" + "\n\t")
            } else {
                v10 := GenS1{F0: 0, F1: 1, F2: 252}
                v1 = v1
                _ = v10
            }
            v5 += (fi9 << ((18446744073709551615) & 63))
            _ = fi9
        }
        v6 = (v7 - 0)
    } else {
        var v11 uint32 = 1
        _ = v11
    }
    if (2147483648 <= 1) {
        v12 := GenS2{F0: 765}
        for fi13 := 0; fi13 < 17; fi13++ {
            v14 := GenS3{F0: -12}
            if (0 >= 975) {
                delete(v4, 416)
                v15 := GenS1{F0: 4294967295, F1: (-(v7)), F2: 708}
                _ = v15
            } else {
                var v16 int32 = 2147483647
                v17 := GenS2{F0: 2}
                v18 := []int{(fi13 >> ((284) & 63)), 127}
                _ = v16
                _ = v17
                _ = v18
            }
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    switch -906 {
    case 127:
        v19 := GenS3{F0: -1}
        _ = v19
    case -349:
        if (2147483647 >= 183) {
            var v20 int32 = -215
            v5 += v5
            v3 = 1
            _ = v20
        } else {
            var v21 string = (v1 + v1)
            v22 := GenS1{F0: 1, F1: -1, F2: 255}
            _ = v21
            _ = v22
        }
    case -9223372036854775808:
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 uint = 23
            v3 = 4
            _ = fi23
            _ = v24
        }
    default:
        v25 := GenS4{F0: v8, F1: v2, F2: 2}
        _ = v25
    }
    if false {
        v26 := []int{v5, v5, 32768, v5}
        _ = v26
    }
    v27 := []int{}
    var v28 uint16 = 65535
    v28 -= v28
    if false {
        for fi29 := 0; fi29 < 1; fi29++ {
            v30 := GenS5{F0: (^(v7))}
            _ = fi29
            _ = v30
        }
        if len(v27) >= 2 {
            v27 = v27[1:len(v27)]
        }
    }
    delete(v4, 596)
    fmt.Println(v2, v5)
    v5 = v5
    for fi31 := 0; fi31 < 14; fi31++ {
        v32 := GenS6{F0: 509, F1: true}
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
