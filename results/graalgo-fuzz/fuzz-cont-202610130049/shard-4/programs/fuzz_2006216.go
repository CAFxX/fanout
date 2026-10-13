package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int32; F1 bool; F2 string }

func main() {
    var v1 bool = (false && true)
    v2 := []uint32{4294967295, 4294967295}
    var v3 string = ((("foo bar" + "\n\t") + "\n\t") + "")
    v4 := make(map[string]uint32)
    var v5 uint64 = 652
    var v6 int32 = 42
    var v7 uint16 = 2
    v1 = v1
    var v8 float32 = ((3.14159 - 1e10) + 0.1)
    fmt.Println(v6, v8, v6)
    fmt.Println(v1, v7, v1)
    delete(v4, v3)
    v9 := GenS1{F0: (((-771 != -423) && (1000 != 42)) && (0 != 2147483648)), F1: "foo bar"}
    if (1 > 34) {
        v8 -= v8
        v5 -= 1
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = "foo bar"
    v3 = "hello"
    for fi10 := 0; fi10 < 19; fi10++ {
        if v1 {
            switch (fi10 - fi10) {
            case -9223372036854775808:
                var v11 float32 = (v8 / v8)
                _ = v11
            }
            v12 := GenS2{F0: 524, F1: (128 != 0), F2: ((v3 + v3) + (((v3 + v3) + v3) + (("\n\t" + v3) + "a")))}
            _ = v12
        } else {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v7)
            case "foo bar":
                for fi13 := 0; fi13 < 20; fi13++ {
                    var v14 int32 = (-(v6))
                    _ = fi13
                    _ = v14
                }
            case "":
                v15 := GenS1{F0: v1, F1: "hello"}
                _ = v15
            }
            fmt.Println(v8, v3, v8)
            v2 = append(v2, 788)
        }
        _ = fi10
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
