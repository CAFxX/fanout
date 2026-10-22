package main

import (
    "fmt"
)

func genFn0() int {
    var v7 uint8 = 1
    switch "" {
    case "":
        var v8 string = ""
        _ = v8
    case "foo bar":
        v7 -= v7
    }
    v9 := []int{-9223372036854775808, 9223372036854775807, 2147483647}
    var v10 uint16 = 0
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = v10
    _ = v7
    _ = v9
    return 2147483647
}

func main() {
    v1 := []string{"foo bar", "", ("\n\t" + "hello")}
    v2 := make(map[int]bool)
    var v3 uint32 = 2
    v4 := []int{-834, -190, 32768, -483}
    var v5 bool = true
    var v6 int = 4294967295
    v3 = (v3 & v3)
    for fi11 := 0; fi11 < 10; fi11++ {
        switch "hello" {
        case "hello":
            if true {
                v6 = fi11
                var v12 string = "a"
                _ = v12
            }
        case "\n\t":
            _ = len(v1)
        default:
            var v13 float32 = 1e10
            _ = v13
        }
        _ = fi11
    }
    if len(v4) > 0 {
        v4[0] = (v6 + (v6 * genFn0()))
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        switch "foo bar" {
        case "hello":
            var v15 int = (fi14 + fi14)
            _ = v15
        case "foo bar":
            v6 = 982
        }
        _ = fi14
    }
    if (true || (-707 < -1)) {
        _ = len(v1)
        v6 = (v6 * (v6 << ((2) & 63)))
    } else {
        var v16 int32 = -1
        var v17 int8 = 1
        v17 = (v17 / (((v17 % (((v17 & 1)) | 1))) | 1))
        _ = v16
        _ = v17
    }
    v6 = genFn0()
    v18 := v2[genFn0()]
    fmt.Println(v18)
    var v19 int64 = int64(v6)
    v4 = append(v4, genFn0())
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
