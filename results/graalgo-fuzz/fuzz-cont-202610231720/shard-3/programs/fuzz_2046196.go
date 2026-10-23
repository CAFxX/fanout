package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int16 }
func genFn0(p0 int) string {
    v7 := GenS1{F0: 18446744073709551615, F1: ("foo bar" + ((("foo bar" + "") + (("hello" + "foo bar") + "hello")) + ("\n\t" + "hello"))), F2: 32767}
    fmt.Println(p0)
    switch 4294967295 {
    case 9223372036854775807:
        fmt.Println(p0)
    default:
        var v8 uint8 = 127
        _ = v8
    }
    fmt.Println(p0)
    v9 := []int{p0}
    _ = p0
    _ = v7
    _ = v9
    return "hello"
}

func main() {
    var v1 int = -570
    v2 := make(map[string]string)
    v3 := []uint32{451, 2147483647}
    var v4 string = ""
    var v5 float64 = 1.5
    var v6 uint = 716
    _ = len(v2)
    v4 = v4
    v5 = 0.1
    if (!true) {
        _ = len(v3)
    } else {
        if len(v3) > 0 {
            v3[0] = 0
        }
    }
    var v10 int64 = -526
    for fi11 := 0; fi11 < 15; fi11++ {
        v6 += (v6 * v6)
        _ = fi11
    }
    v2[""] = v4
    v6 = v6
    fmt.Println(v6, v10)
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
