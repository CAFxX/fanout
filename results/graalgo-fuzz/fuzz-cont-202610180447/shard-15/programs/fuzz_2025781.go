package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
func genFn0(p0 int8, p1 uint8, p2 uint32) uint16 {
    v7 := []int{32767, 42, -213, 881}
    _ = len(v7)
    var v8 uint = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return uint16(p1)
}
func genFn1(p0 uint32, p1 bool) int32 {
    v9 := []int{127, 128, -1, 32768}
    if (p1 || ((float64((0.1 * (0.1 - 1e10))) / 3.14159) < (0.1 * -2.25))) {
        fmt.Println(p1)
        switch ("" + "a") {
        case "hello":
            switch 4294967295 {
            case 127:
                var v10 int64 = 651
                _ = v10
            case 2147483647:
                for fi11 := 0; fi11 < 17; fi11++ {
                    p1 = (-613 > -2147483648)
                    if len(v9) > 0 {
                        v9[0] = (fi11 | (fi11 / ((fi11) | 1)))
                    }
                    _ = fi11
                }
            }
        case "\n\t":
            v9 = append(v9, -163)
        }
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = v9
    return 0
}

func main() {
    v1 := []int{127, -9223372036854775808, -358}
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    v5 := make(map[string]string)
    var v6 uint8 = 128
    if (249 < 255) {
        v6 -= v6
    }
    var v12 string = "foo bar"
    fmt.Println(v6)
    v13 := GenS1{F0: v12}
    var v14 uint16 = genFn0(1, v6, 262)
    v12 = v12
    switch 450 {
    case 4294967295:
        for fi15 := 0; fi15 < 14; fi15++ {
            v16 := GenS2{F0: -1}
            fmt.Println(v6, fi15, v14)
            _ = len(v1)
            _ = fi15
            _ = v16
        }
    case 1:
        fmt.Println(v6)
    }
    delete(v5, v12)
    v17 := GenS1{F0: v12}
    v18 := GenS1{F0: "hello"}
    v14 = v14
    var v19 uint = 0
    if true {
        v4[400] = (v6 > v6)
        fmt.Println(v19, v19)
        v19 = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
