package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := []uint32{2147483648, 502}
    v2 := make(map[int]bool)
    v3 := []int{127}
    var v4 int16 = 42
    v5 := []int{}
    if len(v1) > 0 {
        v1[0] = 2147483648
    }
    var v6 uint = 255
    v4 -= (^(v4))
    switch (("\n\t" + ("" + "hello")) + "a") {
    case "":
        var v7 uint32 = 654
        _ = v7
    case "hello":
        v8 := GenS1{F0: 657}
        _ = v8
    }
    v1 = append(v1, 567)
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        for fi9 := 0; fi9 < 11; fi9++ {
            v4 = v4
            v10 := GenS1{F0: (fi9 % (((fi9 * fi9)) | 1))}
            v6 += v6
            _ = fi9
            _ = v10
        }
    default:
        v11 := GenS1{F0: 4294967295}
        _ = v11
    }
    var v12 float64 = (3.14159 + 1.5)
    if (627 < 65535) {
        var v13 int16 = (v4 | 42)
        _ = v13
    } else {
        if (true && ((v6 << ((v6) & 63)) >= (v6 ^ 0))) {
            if true {
                var v14 uint64 = 663
                _ = v14
            }
            v12 = (v12 + v12)
        }
        v2[445] = (!((3.14159 - 1e10) != v12))
        var v15 float64 = v12
        _ = v15
    }
    v4 = v4
    var v16 int8 = 0
    v6 += 2
    for fi17 := 0; fi17 < 18; fi17++ {
        v3 = append(v3, 128)
        var v18 int16 = v4
        _ = len(v2)
        _ = fi17
        _ = v18
    }
    switch "foo bar" {
    case "":
        fmt.Println(v4, v6, v16)
    case "\n\t":
        fmt.Println(v16, v16)
    case "hello":
        for fi19 := 0; fi19 < 17; fi19++ {
            v20 := GenS1{F0: 1}
            _ = fi19
            _ = v20
        }
    default:
        if (9223372036854775807 == 1) {
            v4 = (v4 << (((v6 | v6)) & 63))
            v12 -= 1e10
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
