package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint8, p1 uint8, p2 uint8) uint {
    p2 -= p0
    p0 = 178
    if (-2.25 < 3.14159) {
        fmt.Println(p1, p2)
    } else {
        var v5 uint = 65535
        switch -2147483648 {
        case -1:
            var v6 int = 151
            _ = v6
        case 32768:
            for fi7 := 0; fi7 < 20; fi7++ {
                v8 := []int{}
                _ = fi7
                _ = v8
            }
        default:
            var v9 int = -9223372036854775808
            _ = v9
        }
        _ = v5
    }
    v10 := []int{-2147483648}
    p2 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    var v4 float64 = ((-2.25 + 3.14159) - ((float64(0.1) / 1.0) * 1e10))
    fmt.Println(v4)
    var v11 string = "hello"
    v11 = "hello"
    if true {
        for fi12 := 0; fi12 < 4; fi12++ {
            switch fi12 {
            case 1:
                v11 = v11
            case 0:
                var v13 string = (v11 + "a")
                _ = v13
            }
            _ = fi12
        }
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        for fi15 := 0; fi15 < 3; fi15++ {
            switch v11 {
            case "":
                if (!(42 == -1)) {
                    var v16 uint64 = 223
                    _ = v16
                } else {
                    delete(v3, "\n\t")
                }
            }
            v11 = ("hello" + (v11 + "a"))
            _ = fi15
        }
        v17 := GenS1{F0: 169, F1: ("hello" + "foo bar")}
        _ = fi14
        _ = v17
    }
    v4 += 0.1
    v18 := GenS2{F0: -344}
    v19 := []int{-287, 127, 127, 128}
    var v20 int = -2147483648
    v20 = 32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
