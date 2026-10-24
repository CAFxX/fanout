package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int8) int16 {
    var v7 float32 = (0.1 / (0.1 * -2.25))
    v8 := []int{4294967295, 127}
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: false}
        p0 -= p0
        switch (("foo bar" + "\n\t") + "hello") {
        case "":
            p0 = 0
        case "\n\t":
            var v11 uint8 = 42
            _ = v11
        default:
            fmt.Println(fi9, v7)
        }
        _ = fi9
        _ = v10
    }
    var v12 int32 = -684
    _ = p0
    _ = v12
    _ = v7
    _ = v8
    return 1000
}

func main() {
    v1 := []uint32{704}
    v2 := []string{"foo bar", ("foo bar" + ("\n\t" + "foo bar")), "a", "hello"}
    var v3 int32 = 42
    var v4 int32 = (v3 << ((255) & 63))
    var v5 uint8 = 0
    var v6 int32 = (v3 << ((0) & 63))
    v13 := GenS1{F0: true}
    if (!((v4 << ((255) & 63)) >= 2147483647)) {
        var v14 int8 = 1
        v6 = v6
        _ = v14
    } else {
        v4 = (v4 | 0)
    }
    if ("a" != "") {
        v3 += (-(v4))
        if true {
            if true {
                for fi15 := 0; fi15 < 3; fi15++ {
                    v5 += (^(v5))
                    v5 = (v5 % ((v5) | 1))
                    var v16 bool = (-1 != 9223372036854775807)
                    _ = fi15
                    _ = v16
                }
            }
            var v17 uint32 = 65535
            _ = v17
        }
    } else {
        var v18 uint32 = 2
        fmt.Println(v6, v18)
        _ = v18
    }
    v5 -= 175
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 int64 = -1
        _ = fi19
        _ = v20
    }
    var v21 float32 = ((1.5 / 1e10) / (3.14159 * 3.14159))
    v6 = v3
    var v22 string = "foo bar"
    v6 = 70
    v22 = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
