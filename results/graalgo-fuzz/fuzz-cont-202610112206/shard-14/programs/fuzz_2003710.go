package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint; F1 int16; F2 string }
type GenS3 struct { F0 bool; F1 bool }
func genFn0(p0 bool, p1 int64) int16 {
    var v5 uint16 = 2
    fmt.Println(p0)
    fmt.Println(p1, p1, v5)
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1() uint {
    var v6 int16 = -295
    var v7 bool = (false && ((!(2 >= 42)) && true))
    v8 := GenS1{F0: 117, F1: v6, F2: 64}
    fmt.Println(v6, v6)
    switch "foo bar" {
    case "hello":
        v6 = genFn0(((9223372036854775807 >= -206) && (false && ((!("\n\t" < "")) || v7))), 0)
    case "":
        for fi9 := 0; fi9 < 11; fi9++ {
            switch (fi9 >> ((482) & 63)) {
            case 128:
                switch fi9 {
                case 989:
                    for fi10 := 0; fi10 < 18; fi10++ {
                        v11 := []int{fi10, (fi9 - 32768)}
                        _ = len(v11)
                        v7 = (0 <= 2)
                        _ = fi10
                        _ = v11
                    }
                default:
                    var v12 string = "a"
                    _ = v12
                }
            case -2147483648:
                for fi13 := 0; fi13 < 5; fi13++ {
                    var v14 int = 1
                    _ = fi13
                    _ = v14
                }
            }
            if v7 {
                v6 = int16(fi9)
            }
            fmt.Println(v6)
            _ = fi9
        }
    }
    _ = v6
    _ = v7
    _ = v8
    return 65535
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    if false {
        switch -478 {
        case 128:
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 float32 = (0.0 + 0.1)
                _ = fi15
                _ = v16
            }
        case 1:
            var v17 int16 = 32767
            _ = v17
        case -2147483648:
            var v18 int32 = -2147483648
            _ = v18
        default:
            var v19 float32 = 1e10
            _ = v19
        }
    }
    v20 := GenS2{F0: 42, F1: 702, F2: "a"}
    var v21 uint8 = 255
    fmt.Println(v21)
    v21 = 66
    for fi22 := 0; fi22 < 1; fi22++ {
        for fi23 := 0; fi23 < 14; fi23++ {
            var v24 string = ""
            v3[fi23] = (^(fi23))
            v25 := GenS2{F0: 65535, F1: 995, F2: ""}
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi22
    }
    delete(v3, 2147483647)
    _ = len(v2)
    switch 0 {
    case 32768:
        var v26 string = "hello"
        _ = v26
    case 849:
        for fi27 := 0; fi27 < 15; fi27++ {
            v28 := GenS3{F0: (((127 == (v21 % ((v21) | 1))) || (false || (110 < (v21 * 0)))) && true), F1: (1e10 != 1e10)}
            fmt.Println(fi27, v21)
            _ = fi27
            _ = v28
        }
    default:
        fmt.Println(v21)
    }
    v29 := []int{-9223372036854775808, 647}
    v30 := GenS2{F0: 2, F1: genFn0(true, 0), F2: ("hello" + "foo bar")}
    var v31 uint32 = 65535
    var v32 int = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
