package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]int)
    v5 := make(map[uint32]uint32)
    var v6 uint64 = 1
    for fi7 := 0; fi7 < 12; fi7++ {
        if false {
            var v8 int = fi7
            v6 -= v6
            v9 := []int{(v8 + fi7), v8, fi7}
            _ = v8
            _ = v9
        } else {
            switch -192 {
            case -9223372036854775808:
                v10 := GenS1{F0: false}
                _ = v10
            case 494:
                fmt.Println(v6, fi7)
            default:
                var v11 string = "foo bar"
                _ = v11
            }
            fmt.Println(v6)
        }
        var v12 uint16 = 42
        v6 -= 1
        _ = fi7
        _ = v12
    }
    fmt.Println(v6)
    for fi13 := 0; fi13 < 16; fi13++ {
        v6 += 546
        var v14 int32 = -1
        v6 += (^(v6))
        _ = fi13
        _ = v14
    }
    v15 := []int{119}
    var v16 uint8 = 1
    _ = len(v3)
    v4[-2147483648] = 128
    if (!(-324 == -1)) {
        var v17 int64 = -9223372036854775808
        _ = len(v1)
        _ = v17
    } else {
        if false {
            if len(v15) > 0 {
                v15[0] = 1
            }
        } else {
            var v18 int32 = -1
            for fi19 := 0; fi19 < 20; fi19++ {
                switch (-(fi19)) {
                case 128:
                    v20 := v1[("" + ("\n\t" + (("\n\t" + "") + ("a" + "\n\t"))))]
                    _ = v20
                }
                v18 += (v18 * v18)
                var v21 int = fi19
                _ = fi19
                _ = v21
            }
            v18 -= v18
            _ = v18
        }
        var v22 uint = 686
        _ = v22
    }
    v6 -= v6
    var v23 float32 = ((0.1 - -2.25) + float32(1.5))
    switch (("a" + "a") + "") {
    case "foo bar":
        for fi24 := 0; fi24 < 18; fi24++ {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            _ = fi24
        }
    case "hello":
        if ((false || false) || (-2.25 == -2.25)) {
            v15 = append(v15, 1)
        }
    }
    fmt.Println(v6, v6, v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
