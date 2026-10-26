package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 string = ("a" + "\n\t")
    v2 := []string{"hello", "\n\t", "foo bar", ((v1 + v1) + "foo bar")}
    v3 := []uint32{}
    var v4 int8 = 42
    var v5 int = 32768
    var v6 int = -749
    if (("foo bar" < v1) && true) {
        if false {
            v5 = v5
            v7 := GenS1{F0: v5}
            var v8 int64 = -9223372036854775808
            _ = v7
            _ = v8
        }
        if len(v2) > 0 {
            v2[0] = v1
        }
        v5 -= 0
    }
    v1 = v1
    var v9 int16 = 727
    v1 = v1
    var v10 int8 = v4
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v11 int8 = (v4 | v4)
    switch (^(v6)) {
    case -2147483648:
        var v12 float32 = (1e10 - 1e10)
        _ = v12
    case 32767:
        var v13 uint32 = 2147483648
        _ = v13
    case 473:
        var v14 int32 = 1
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
