package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int = 127
    var v2 int32 = 189
    var v3 uint16 = 0
    var v4 string = "a"
    switch v4 {
    case "a":
        var v5 uint32 = 42
        _ = v5
    case "hello":
        v3 += (v3 ^ (v3 + v3))
    }
    fmt.Println(v2)
    v6 := GenS1{F0: v4}
    v1 -= v1
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 bool = true
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 uint64 = 0
        switch "\n\t" {
        case "":
            if true {
                v3 = v3
                var v11 float32 = float32(1.5)
                fmt.Println(v2, v2)
                _ = v11
            }
        case "a":
            if (0 < 942) {
                v12 := GenS1{F0: "foo bar"}
                v13 := []int{v1, int(v3), fi9, (fi9 << ((42) & 63))}
                var v14 uint8 = uint8(v10)
                _ = v12
                _ = v13
                _ = v14
            } else {
                v15 := GenS1{F0: (v4 + v4)}
                _ = v15
            }
        case "foo bar":
            var v16 uint64 = v10
            _ = v16
        }
        _ = fi9
        _ = v10
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        v2 = (v2 / ((v2) | 1))
        for fi18 := 0; fi18 < 3; fi18++ {
            fmt.Println(v3)
            v19 := []int{(v1 % ((832) | 1)), fi17, 127}
            v4 = v4
            fmt.Println(fi17, v1, v3)
            _ = fi18
            _ = v19
        }
        v1 = 127
        _ = fi17
    }
    v20 := []int{-1}
    var v21 uint64 = 18446744073709551615
    var v22 int16 = 916
    var v23 float32 = -2.25
    if false {
        var v24 string = (v4 + v4)
        _ = v24
    }
    v21 += (^(v21))
    v20 = append(v20, 0)
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v6
}
