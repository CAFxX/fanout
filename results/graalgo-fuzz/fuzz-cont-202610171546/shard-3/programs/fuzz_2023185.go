package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int32) uint16 {
    p0 -= (p0 - (^(p0)))
    var v5 uint32 = 2147483647
    var v6 uint16 = 42
    fmt.Println(v6)
    _ = p0
    _ = v5
    _ = v6
    return 255
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]string)
    v3 := []uint32{2147483648}
    v4 := []uint32{42, 4294967295}
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: 2, F1: "\n\t"}
        delete(v1, fi7)
        var v9 uint64 = 2
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint64 = 4294967295
    v11 := GenS1{F0: 1, F1: ((("hello" + "a") + "hello") + "a")}
    fmt.Println(v10)
    if false {
        for fi12 := 0; fi12 < 11; fi12++ {
            if true {
                var v13 uint8 = 0
                _ = v13
            } else {
                v10 -= (v10 | v10)
                delete(v1, fi12)
                for fi14 := 0; fi14 < 2; fi14++ {
                    if true {
                        v10 = v10
                        switch (^(fi14)) {
                        case -2147483648:
                            fmt.Println(fi12, v10, v10)
                        case 754:
                            switch "" {
                            case "hello":
                                var v15 uint64 = (v10 ^ v10)
                                _ = v15
                            case "":
                                v16 := GenS2{F0: 2}
                                _ = v16
                            case "a":
                                fmt.Println(v10, fi12, v10)
                            }
                        default:
                            v10 += (v10 - v10)
                        }
                        v17 := GenS1{F0: 255, F1: "\n\t"}
                        _ = v17
                    } else {
                        _ = len(v1)
                    }
                    v10 = v10
                    fmt.Println(fi12)
                    _ = fi14
                }
            }
            for fi18 := 0; fi18 < 13; fi18++ {
                v10 = 2
                var v19 uint8 = 0
                _ = fi18
                _ = v19
            }
            _ = fi12
        }
    }
    var v20 bool = (!(false && (1.5 < 1.5)))
    var v21 uint8 = 85
    fmt.Println(v20, v20)
    v22 := GenS1{F0: v21, F1: ("hello" + "")}
    v10 -= (v10 << ((4294967295) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
