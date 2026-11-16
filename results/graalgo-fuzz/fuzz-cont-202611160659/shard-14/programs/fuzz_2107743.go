package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 bool; F2 int8 }
type GenS3 struct { F0 int8; F1 int16; F2 int64 }
type GenS4 struct { F0 int }

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 659
    v3 := []int{-2147483648, 128, -7}
    var v4 float32 = 3.14159
    v5 := make(map[string]string)
    var v6 int32 = 42
    delete(v5, "")
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int8 = -23
        if false {
            for fi9 := 0; fi9 < 7; fi9++ {
                v10 := GenS1{F0: 320}
                v11 := GenS2{F0: 128, F1: false, F2: (-(v8))}
                _ = len(v3)
                _ = fi9
                _ = v10
                _ = v11
            }
            v1[(-(fi7))] = ((v8 % ((-43) | 1)) <= v8)
        }
        v12 := GenS2{F0: 170, F1: true, F2: v8}
        _ = fi7
        _ = v12
        _ = v8
    }
    v5["\n\t"] = (("" + "a") + "a")
    fmt.Println(v4, v6, v6)
    for fi13 := 0; fi13 < 7; fi13++ {
        v4 = v4
        _ = fi13
    }
    v2 = v2
    if (-7 >= 14) {
        for fi14 := 0; fi14 < 20; fi14++ {
            if false {
                var v15 bool = false
                v16 := GenS3{F0: 0, F1: 1000, F2: v2}
                v2 -= 1
                _ = v15
                _ = v16
            }
            v2 -= 42
            for fi17 := 0; fi17 < 5; fi17++ {
                fmt.Println(fi17)
                var v18 bool = false
                _ = fi17
                _ = v18
            }
            _ = fi14
        }
        var v19 uint = 2
        v2 = v2
        _ = v19
    }
    v5[""] = ("" + "a")
    var v20 string = ("a" + "")
    for fi21 := 0; fi21 < 17; fi21++ {
        v22 := GenS4{F0: (fi21 % (((fi21 % ((fi21) | 1))) | 1))}
        var v23 int = fi21
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
