package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint16 }
type GenS2 struct { F0 string }
func genFn0(p0 int, p1 uint16, p2 int32) int8 {
    var v7 int64 = 0
    v8 := []int{(p0 % (((p0 & (p0 >> ((4294967295) & 63)))) | 1)), (p0 << ((0) & 63))}
    switch p0 {
    case 127:
        p1 += (^(p1))
    case 727:
        fmt.Println(p2)
    case 9223372036854775807:
        p2 = (p2 + (p2 << ((18446744073709551615) & 63)))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 2
    var v3 int = 0
    var v4 bool = false
    var v5 int32 = 42
    v6 := []string{"hello", "foo bar", "a"}
    v9 := GenS1{F0: v3, F1: 1, F2: 32767}
    fmt.Println(v5, v2, v3)
    if (v2 > 2) {
        var v10 int = v3
        _ = v10
    } else {
        var v11 uint64 = 585
        _ = v11
    }
    v4 = false
    if len(v6) > 0 {
        v6[0] = "foo bar"
    }
    fmt.Println(v4, v4, v3)
    var v12 string = ("a" + ((("a" + "") + ("" + "\n\t")) + ""))
    v13 := GenS1{F0: v3, F1: 19, F2: 952}
    v5 = -2147483648
    var v14 int64 = int64(v3)
    _ = len(v1)
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 string = ((v12 + ("a" + "")) + "hello")
        if len(v6) > 0 {
            v6[0] = ""
        }
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: "foo bar"}
    v3 = v3
    switch (v12 + "") {
    case "foo bar":
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 float32 = 1e10
            switch (v12 + "\n\t") {
            case "hello":
                v19 -= v19
            default:
                v5 -= v5
            }
            _ = fi18
            _ = v19
        }
    case "a":
        for fi20 := 0; fi20 < 15; fi20++ {
            fmt.Println(v3, v12, v12)
            _ = fi20
        }
    case "":
        v6 = append(v6, v12)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
