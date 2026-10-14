package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint64 }
type GenS2 struct { F0 string; F1 uint32; F2 int8 }
func genFn0(p0 uint64, p1 int16, p2 uint8) uint32 {
    v6 := []int{}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v7 := GenS1{F0: 65535, F1: "", F2: (p0 & p0)}
    v8 := GenS1{F0: 2, F1: "\n\t", F2: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int32 = 0
    v2 := []int{4294967295, 32768, int(v1)}
    v3 := make(map[int]bool)
    var v4 uint8 = 2
    v5 := make(map[uint32]int)
    fmt.Println(v1, v4)
    v1 = v1
    switch 128 {
    case 32767:
        v1 = (v1 * 1)
    case 32768:
        var v9 float32 = 1.5
        _ = v9
    case 99:
        switch "a" {
        case "a":
            if (255 == 255) {
                if false {
                    switch "hello" {
                    case "":
                        v10 := GenS1{F0: 168, F1: "", F2: 18446744073709551615}
                        _ = v10
                    case "\n\t":
                        v11 := GenS1{F0: 2, F1: "hello", F2: 78}
                        _ = v11
                    case "foo bar":
                        delete(v5, 2)
                    }
                    var v12 int16 = 783
                    _ = v12
                }
            } else {
                switch 127 {
                case -2147483648:
                    var v13 int64 = -237
                    _ = v13
                }
            }
        case "foo bar":
            v14 := GenS1{F0: 42, F1: ("a" + "\n\t"), F2: 4294967295}
            _ = v14
        case "hello":
            var v15 bool = false
            _ = v15
        }
    default:
        var v16 int32 = -2147483648
        _ = v16
    }
    v17 := GenS1{F0: 0, F1: "a", F2: 42}
    v18 := GenS1{F0: 2, F1: "hello", F2: 373}
    for fi19 := 0; fi19 < 2; fi19++ {
        v1 += 1
        v4 -= (v4 % ((2) | 1))
        v4 -= v4
        _ = fi19
    }
    _ = len(v2)
    var v20 uint64 = 4294967295
    v21 := GenS2{F0: "a", F1: 4294967295, F2: 127}
    fmt.Println(v20, v4)
    fmt.Println(v4, v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
