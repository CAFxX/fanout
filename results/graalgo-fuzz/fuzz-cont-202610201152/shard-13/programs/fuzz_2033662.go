package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
func genFn0(p0 uint64, p1 bool, p2 string) int16 {
    v7 := GenS1{F0: 1, F1: -9223372036854775808}
    var v8 float32 = (0.1 / 3.14159)
    if (42 >= 42) {
        if (((1 <= 29) || (p0 < 523)) || p1) {
            p1 = p1
            if p1 {
                fmt.Println(p0, v8)
                fmt.Println(p2)
                var v9 uint64 = (p0 * p0)
                _ = v9
            }
        } else {
            var v10 bool = true
            _ = v10
        }
        v11 := GenS1{F0: -2147483648, F1: 0}
        var v12 int8 = -123
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"a", (("hello" + "foo bar") + ("\n\t" + ""))}
    v3 := make(map[string]uint32)
    v4 := []string{("a" + (((("foo bar" + "") + ("foo bar" + "hello")) + "hello") + "")), ("" + "\n\t")}
    var v5 int16 = -32768
    v6 := []int{}
    fmt.Println(v5)
    v5 = (v5 << ((2) & 63))
    _ = len(v1)
    v13 := GenS1{F0: -2147483648, F1: 9223372036854775807}
    var v14 int64 = 0
    var v15 int8 = 0
    for fi16 := 0; fi16 < 2; fi16++ {
        v3[""] = 1
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 uint16 = 32767
            _ = fi17
            _ = v18
        }
        v19 := GenS1{F0: -1, F1: 42}
        _ = fi16
        _ = v19
    }
    v5 += (v5 % ((genFn0(0, ((("" + ("hello" + (("foo bar" + (("hello" + "hello") + ((("a" + "") + ("\n\t" + "\n\t")) + ("a" + "a")))) + "hello"))) + ("foo bar" + "")) > "a"), "hello")) | 1))
    if (false && true) {
        switch "" {
        case "hello":
            v5 = (v5 >> ((0) & 63))
        case "a":
            if false {
                var v20 string = "hello"
                _ = v20
            }
        case "\n\t":
            v2 = append(v2, "a")
        default:
            var v21 uint = 131
            _ = v21
        }
    } else {
        var v22 uint64 = 0
        _ = v22
    }
    var v23 int64 = int64(v5)
    var v24 float64 = 1e10
    var v25 uint32 = 2
    switch 9223372036854775807 {
    case 4294967295:
        var v26 int = 42
        _ = v26
    case 41:
        _ = len(v3)
    case 9223372036854775807:
        for fi27 := 0; fi27 < 14; fi27++ {
            v5 = (v5 >> ((863) & 63))
            if (true && false) {
                _ = len(v1)
            } else {
                v25 = v25
            }
            _ = fi27
        }
    default:
        for fi28 := 0; fi28 < 11; fi28++ {
            fmt.Println(v25)
            v29 := GenS1{F0: 1, F1: v23}
            _ = fi28
            _ = v29
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
