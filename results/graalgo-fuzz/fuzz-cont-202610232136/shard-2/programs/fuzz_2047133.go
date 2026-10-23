package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int32, p1 uint64, p2 uint) string {
    if (!(false && ((^(p0)) > p0))) {
        v7 := GenS1{F0: 32767}
        p0 = 1
        var v8 uint8 = 1
        _ = v7
        _ = v8
    } else {
        p1 -= p1
        switch ("foo bar" + ("\n\t" + "")) {
        case "a":
            var v9 int16 = 1000
            _ = v9
        case "\n\t":
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := []int{fi10, (fi10 + (fi10 | 4294967295)), fi10, (^(fi10))}
                v12 := GenS2{F0: int16(p0)}
                v11 = append(v11, 32768)
                _ = fi10
                _ = v11
                _ = v12
            }
        }
    }
    var v13 float64 = (3.14159 * (0.1 / 3.14159))
    p2 -= (^(p2))
    switch -1 {
    case -891:
        p2 -= p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return "a"
}

func main() {
    var v1 int16 = 32767
    var v2 int8 = 0
    var v3 uint64 = 42
    v4 := make(map[uint32]int)
    v5 := make(map[string]bool)
    var v6 int8 = int8(v1)
    v6 = 127
    v14 := GenS3{F0: 0}
    var v15 int32 = 1
    var v16 string = "\n\t"
    var v17 int = -2147483648
    v6 -= v6
    var v18 uint32 = 1
    v19 := []int{}
    var v20 int = v17
    switch v16 {
    case "foo bar":
        var v21 uint64 = (v3 + 18446744073709551615)
        _ = v21
    case "":
        v6 = v2
    default:
        var v22 int16 = 0
        _ = v22
    }
    var v23 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
