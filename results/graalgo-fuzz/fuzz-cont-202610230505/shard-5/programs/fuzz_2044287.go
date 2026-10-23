package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 string, p1 int32) uint64 {
    for fi6 := 0; fi6 < 20; fi6++ {
        p0 = p0
        v7 := GenS1{F0: 32767}
        v8 := GenS1{F0: 410}
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(p0, p0)
    p1 = (p1 + 961)
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 uint32 = 718
    var v5 uint8 = 147
    if (!(-1 == -1)) {
        if true {
            if true {
                switch "\n\t" {
                case "hello":
                    var v9 float32 = 1.5
                    _ = v9
                case "foo bar":
                    _ = len(v2)
                case "":
                    v10 := v1[0]
                    _ = v10
                }
                var v11 int = 839
                _ = v11
            } else {
                var v12 int32 = 42
                var v13 uint = 42
                v5 += v5
                _ = v12
                _ = v13
            }
        }
        fmt.Println(v5, v5)
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := v2[(fi14 ^ fi14)]
            _ = fi14
            _ = v15
        }
    }
    var v16 uint32 = 1
    for fi17 := 0; fi17 < 4; fi17++ {
        _ = len(v3)
        var v18 int8 = -128
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: 1}
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 float32 = 1.5
        var v22 float32 = (v21 / v21)
        var v23 string = "foo bar"
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    if ((float32((0.0 / 3.14159)) <= (-2.25 * 0.1)) && true) {
        for fi24 := 0; fi24 < 5; fi24++ {
            fmt.Println(v5, v16)
            if false {
                switch "hello" {
                case "hello":
                    fmt.Println(fi24)
                case "\n\t":
                    var v25 uint64 = 1
                    _ = v25
                default:
                    v16 = (v4 >> ((65535) & 63))
                }
            }
            _ = fi24
        }
        for fi26 := 0; fi26 < 2; fi26++ {
            v27 := []int{fi26, fi26}
            _ = fi26
            _ = v27
        }
        v28 := []int{-654, 128, 32768, 128}
        _ = v28
    }
    var v29 bool = false
    v30 := GenS1{F0: 32767}
    v31 := []int{-305}
    switch 32768 {
    case 42:
        delete(v3, 2147483647)
    default:
        if len(v31) >= 2 {
            v31 = v31[1:len(v31)]
        }
    }
    v32 := v1[(v4 * v16)]
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
