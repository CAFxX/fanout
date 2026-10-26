package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint8, p1 int16) uint8 {
    var v5 uint16 = 42
    var v6 uint16 = v5
    v7 := []int{32768}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (p0 + 128)
}

func main() {
    var v1 bool = (!true)
    var v2 uint = 42
    v3 := []string{"hello", ("" + "\n\t"), "a", ""}
    var v4 int16 = 1
    var v8 bool = v1
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 uint = v2
        _ = fi9
        _ = v10
    }
    v3 = append(v3, "")
    if v8 {
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 string = (("foo bar" + ((("a" + ("" + "a")) + "") + "foo bar")) + ("foo bar" + ("\n\t" + ((("foo bar" + "hello") + "\n\t") + (("a" + "foo bar") + "")))))
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 19; fi13++ {
            for fi14 := 0; fi14 < 17; fi14++ {
                v2 = 4294967295
                v8 = (4294967295 > 1)
                _ = fi14
            }
            _ = fi13
        }
    }
    fmt.Println(v2, v1, v8)
    v8 = v8
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 uint8 = 42
        for fi17 := 0; fi17 < 15; fi17++ {
            v8 = v8
            var v18 uint = 0
            _ = fi17
            _ = v18
        }
        v1 = false
        _ = fi15
        _ = v16
    }
    if ("" != ((("hello" + ("hello" + ("foo bar" + "\n\t"))) + "a") + "a")) {
        v19 := GenS1{F0: 0, F1: -32768}
        _ = v19
    }
    fmt.Println(v8, v1, v4)
    if v8 {
        v4 = -548
        var v20 uint64 = 18446744073709551615
        for fi21 := 0; fi21 < 18; fi21++ {
            v3 = append(v3, "\n\t")
            v4 = 0
            fmt.Println(v8, fi21, v8)
            _ = fi21
        }
        _ = v20
    }
    var v22 bool = (!v8)
    var v23 uint16 = 593
    fmt.Println(v22)
    var v24 int16 = v4
    v23 = (^(v23))
    if v22 {
        v2 = 1
        v4 -= 1000
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v8
}
