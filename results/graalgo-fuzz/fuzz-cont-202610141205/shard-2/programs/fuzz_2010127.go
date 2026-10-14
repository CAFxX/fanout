package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int32, p1 uint8) int {
    p1 = p1
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 uint32 = 75
        v7 := GenS1{F0: (p0 / (((p0 | (-(p0)))) | 1))}
        fmt.Println(fi5, p1)
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 bool = (33 >= 0)
    var v9 int16 = 42
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 int16 = -61
    v2 := []string{"a", "a"}
    v3 := make(map[int]string)
    v4 := []string{"a", "\n\t", ("foo bar" + "foo bar"), (("a" + (("a" + "a") + ("foo bar" + "foo bar"))) + "hello")}
    _ = len(v2)
    fmt.Println(v1)
    var v10 int8 = -1
    for fi11 := 0; fi11 < 12; fi11++ {
        switch (("" + "\n\t") + "foo bar") {
        case "a":
            var v12 float64 = 1e10
            _ = v12
        case "foo bar":
            for fi13 := 0; fi13 < 20; fi13++ {
                v14 := GenS1{F0: -623}
                var v15 int32 = -1
                _ = fi13
                _ = v14
                _ = v15
            }
        case "":
            v10 += v10
        }
        _ = fi11
    }
    fmt.Println(v10, v10)
    switch "\n\t" {
    case "foo bar":
        v16 := v3[4294967295]
        _ = v16
    case "a":
        v10 = v10
    case "hello":
        switch 9223372036854775807 {
        case 401:
            v10 -= v10
        case 42:
            for fi17 := 0; fi17 < 10; fi17++ {
                fmt.Println(v10, v1)
                _ = fi17
            }
        }
    }
    if (false || true) {
        v1 = -589
        delete(v3, 1)
        v10 = v10
    }
    v18 := GenS1{F0: 42}
    var v19 uint = 513
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
