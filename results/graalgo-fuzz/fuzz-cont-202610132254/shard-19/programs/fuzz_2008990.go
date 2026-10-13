package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := []string{"a", ((((("a" + "foo bar") + (("foo bar" + (("foo bar" + "foo bar") + ((("a" + ("" + "\n\t")) + (("\n\t" + ("hello" + "hello")) + "hello")) + "hello"))) + "a")) + "") + "\n\t") + "a")}
    v2 := make(map[string]string)
    var v3 int32 = 0
    v4 := []string{}
    v5 := make(map[string]int)
    var v6 int16 = 1000
    v7 := GenS1{F0: true, F1: 4294967295}
    v3 = -2147483648
    var v8 bool = (3.14159 >= 1.5)
    if (42 <= v6) {
        var v9 string = "foo bar"
        if ((v3 + -2147483648) <= (v3 - v3)) {
            var v10 string = (("hello" + (v9 + "foo bar")) + v9)
            _ = v10
        }
        _ = v9
    } else {
        v6 -= (v6 >> ((18446744073709551615) & 63))
        var v11 float64 = (-2.25 * float64(((3.14159 * -2.25) + float32((1e10 + (1.5 / (1.5 / 0.1)))))))
        _ = v11
    }
    if v8 {
        v12 := GenS1{F0: (730 != 4294967295), F1: -2147483648}
        var v13 uint16 = 42
        var v14 bool = true
        _ = v12
        _ = v13
        _ = v14
    }
    _ = len(v5)
    v6 = 0
    switch -1 {
    case -1:
        var v15 uint32 = 1
        _ = v15
    default:
        _ = len(v4)
    }
    _ = len(v1)
    if (!(v8 || ((100 > 0) && (2 >= 65535)))) {
        fmt.Println(v6)
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 uint = 0
            v6 += (v6 % ((-624) | 1))
            fmt.Println(fi16, v8)
            _ = fi16
            _ = v17
        }
    } else {
        var v18 string = "a"
        _ = v18
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        v2["hello"] = "a"
        v8 = (true || (!true))
        if v8 {
            var v20 uint8 = 128
            switch (fi19 ^ (fi19 >> ((42) & 63))) {
            case 128:
                delete(v2, "\n\t")
            case -2147483648:
                var v21 uint32 = 42
                _ = v21
            }
            _ = v20
        } else {
            delete(v2, "a")
        }
        _ = fi19
    }
    var v22 float32 = float32((1e10 / 1.0))
    for fi23 := 0; fi23 < 3; fi23++ {
        v24 := GenS2{F0: 298}
        _ = fi23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
