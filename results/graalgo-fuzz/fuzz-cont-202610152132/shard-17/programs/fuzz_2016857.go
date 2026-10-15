package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint16 }

func main() {
    var v1 int16 = 0
    v2 := make(map[int]bool)
    v3 := []int{1, 42}
    var v4 int16 = v1
    v5 := make(map[string]int)
    switch ("a" + "hello") {
    case "foo bar":
        v4 -= v4
    case "a":
        var v6 int32 = 1
        _ = v6
    }
    v5["a"] = -1
    v4 = -342
    v1 += 1000
    v1 = (v1 + v4)
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(fi7, fi7, fi7)
        _ = fi7
    }
    switch "foo bar" {
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "":
        v3 = append(v3, 32767)
    default:
        fmt.Println(v4)
    }
    v8 := GenS1{F0: 65535, F1: "a", F2: 1}
    if (!((1 < 1) && true)) {
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := GenS2{F0: int32(fi9), F1: 42}
            v5[""] = (fi9 >> ((198) & 63))
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 int8 = 64
                if false {
                    var v13 int8 = (v12 << ((255) & 63))
                    _ = v13
                }
                if (!false) {
                    v5[("\n\t" + "a")] = fi9
                    var v14 int16 = v4
                    _ = v14
                } else {
                    v15 := GenS1{F0: 2147483648, F1: "", F2: 42}
                    _ = v15
                }
                _ = fi11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        var v16 uint64 = 18446744073709551615
        _ = v16
    } else {
        v3 = append(v3, 2147483647)
    }
    if (255 == 540) {
        v17 := v2[42]
        _ = v17
    } else {
        var v18 string = ("foo bar" + ("a" + ("a" + ("a" + "a"))))
        _ = v18
    }
    var v19 uint8 = 137
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
