package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint32 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 string; F2 int }
type GenS5 struct { F0 string; F1 string; F2 uint32 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    var v3 float64 = 3.14159
    var v4 uint = 255
    var v5 string = ("foo bar" + ("a" + "hello"))
    v6 := GenS1{F0: 193, F1: 128, F2: 65535}
    v7 := GenS1{F0: 128, F1: 127, F2: 307}
    if true {
        var v8 string = (("a" + "a") + v5)
        _ = v8
    }
    if true {
        v9 := GenS2{F0: 32767, F1: 42}
        v1[9223372036854775807] = ("\n\t" + "hello")
        v5 = v5
        _ = v9
    } else {
        if true {
            var v10 string = v5
            if (983 < 1) {
                for fi11 := 0; fi11 < 1; fi11++ {
                    v12 := GenS3{F0: 0}
                    _ = fi11
                    _ = v12
                }
                if (-355 < 1) {
                    for fi13 := 0; fi13 < 1; fi13++ {
                        var v14 int32 = -466
                        v15 := GenS3{F0: 73}
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                }
            }
            v5 = (v10 + v5)
            _ = v10
        } else {
            v5 = "\n\t"
            v16 := []int{1}
            delete(v2, -2147483648)
            _ = v16
        }
        fmt.Println(v4, v4)
    }
    var v17 uint64 = 801
    v18 := GenS4{F0: 32767, F1: (v5 + v5), F2: 1}
    var v19 uint8 = 255
    switch 649 {
    case -1:
        var v20 uint8 = v19
        _ = v20
    case -387:
        switch 0 {
        case -1:
            v21 := []int{2147483647, -9223372036854775808, 0}
            _ = v21
        case 1:
            var v22 int32 = -1
            _ = v22
        default:
            for fi23 := 0; fi23 < 14; fi23++ {
                for fi24 := 0; fi24 < 5; fi24++ {
                    v4 += (v4 ^ 65535)
                    _ = len(v2)
                    _ = fi24
                }
                fmt.Println(v3, v5, v5)
                var v25 int32 = -1
                _ = fi23
                _ = v25
            }
        }
    case 4294967295:
        var v26 int32 = 42
        _ = v26
    }
    v17 = (v17 / (((v17 ^ (v17 >> ((v4) & 63)))) | 1))
    if (!true) {
        _ = len(v2)
    }
    if false {
        v5 = v5
        var v27 uint = 0
        if true {
            v28 := GenS5{F0: v5, F1: v5, F2: 399}
            v4 += 42
            _ = v28
        }
        _ = v27
    }
    var v29 int32 = 42
    fmt.Println(v3)
    var v30 uint8 = (v19 - uint8(v4))
    for fi31 := 0; fi31 < 3; fi31++ {
        _ = len(v2)
        v32 := GenS2{F0: -9223372036854775808, F1: 128}
        _ = fi31
        _ = v32
    }
    var v33 string = (v5 + v5)
    v34 := []int{461, 2147483647, 4294967295, 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
