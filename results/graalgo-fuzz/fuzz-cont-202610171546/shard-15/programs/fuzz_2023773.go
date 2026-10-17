package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
func genFn0(p0 int64) int32 {
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := []int{fi6, (fi6 | fi6), (fi6 / (((fi6 / (((fi6 / (((-(fi6))) | 1))) | 1))) | 1))}
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return -424
}

func main() {
    var v1 int32 = 42
    v2 := []int{}
    v3 := []int{int(v1), 32768, 650}
    var v4 int64 = 42
    v5 := make(map[string]string)
    v8 := GenS1{F0: 32767, F1: (("foo bar" + "\n\t") + "foo bar")}
    fmt.Println(v1, v4)
    v4 = (v4 << ((4294967295) & 63))
    v9 := v5[""]
    fmt.Println(v1, v1)
    switch (((((v9 + v9) + v9) + (((v9 + v9) + "a") + v9)) + (v9 + "foo bar")) + "") {
    case "hello":
        switch "" {
        case "\n\t":
            fmt.Println(v4, v4, v9)
        default:
            var v10 int64 = -9223372036854775808
            _ = v10
        }
    case "foo bar":
        var v11 float64 = (1e10 / -2.25)
        _ = v11
    case "a":
        var v12 string = (v9 + (v9 + v9))
        _ = v12
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch -9223372036854775808 {
    case 32768:
        v3 = append(v3, 9223372036854775807)
    case 9223372036854775807:
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 uint16 = 42
            if len(v3) > 0 {
                v3[0] = (fi13 << ((17) & 63))
            }
            v9 = v9
            _ = fi13
            _ = v14
        }
    }
    switch v9 {
    case "hello":
        v15 := GenS1{F0: 825, F1: "foo bar"}
        _ = v15
    }
    v9 = (v9 + ((v9 + (v9 + v9)) + v9))
    v9 = v9
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
