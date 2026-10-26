package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 string; F1 int32; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint, p1 uint) uint {
    if false {
        var v4 uint16 = 1
        p0 = 2
        _ = v4
    }
    var v5 int32 = -561
    switch 32768 {
    case 4294967295:
        p1 += (p1 % ((p0) | 1))
    default:
        var v6 float32 = float32(0.0)
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v5
    return (^(p1))
}

func main() {
    var v1 uint16 = 1
    var v2 uint8 = 0
    v3 := make(map[int]bool)
    var v7 int32 = -2147483648
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 float32 = -2.25
        v1 = (v1 - 65535)
        v10 := []int{128, fi8, fi8}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 string = "a"
    fmt.Println(v2, v1, v11)
    var v12 string = ("\n\t" + v11)
    for fi13 := 0; fi13 < 18; fi13++ {
        v14 := []int{(fi13 + 0), -2147483648}
        for fi15 := 0; fi15 < 10; fi15++ {
            v16 := GenS1{F0: 128, F1: 255}
            var v17 int8 = 83
            fmt.Println(v17, v1, v2)
            v11 = v12
            var v18 bool = false
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        var v19 float32 = 0.0
        _ = fi13
        _ = v14
        _ = v19
    }
    switch -652 {
    case 662:
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 uint = 0
            for fi22 := 0; fi22 < 2; fi22++ {
                v23 := GenS1{F0: v2, F1: 42}
                v24 := GenS1{F0: v2, F1: v21}
                var v25 int32 = 2147483647
                switch fi20 {
                case 42:
                    v11 = v12
                    switch (fi22 + (fi22 % (((fi20 | -558)) | 1))) {
                    case 32768:
                        v26 := GenS2{F0: "", F1: (v7 / ((2147483647) | 1)), F2: (v1 / (((v1 & (v1 | (v1 >> ((v21) & 63))))) | 1))}
                        var v27 int32 = v25
                        _ = v26
                        _ = v27
                    case 42:
                        fmt.Println(v1, v21)
                    case 128:
                        var v28 int = int(v25)
                        _ = v28
                    }
                case 9223372036854775807:
                    var v29 string = "a"
                    _ = v29
                }
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            v30 := GenS1{F0: (v2 ^ (v2 << ((v21) & 63))), F1: 0}
            _ = fi20
            _ = v21
            _ = v30
        }
    default:
        v12 = ("hello" + v12)
    }
    v31 := GenS3{F0: 100, F1: (v1 >> ((255) & 63))}
    v11 = ("a" + v11)
    v32 := GenS2{F0: "", F1: (v7 >> ((255) & 63)), F2: v1}
    var v33 uint64 = 731
    switch v12 {
    case "":
        v34 := GenS1{F0: v2, F1: 65535}
        _ = v34
    }
    if true {
        var v35 int = 0
        switch v35 {
        case -9223372036854775808:
            var v36 int = (v35 << ((4294967295) & 63))
            _ = v36
        default:
            var v37 int32 = (v7 << ((42) & 63))
            _ = v37
        }
        _ = v35
    }
    _ = len(v3)
    var v38 string = v12
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v7
}
