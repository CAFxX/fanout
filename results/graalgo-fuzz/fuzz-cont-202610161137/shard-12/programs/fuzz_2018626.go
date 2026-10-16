package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
func genFn0(p0 int64, p1 uint32) string {
    if (true || false) {
        if (p0 <= -9223372036854775808) {
            p1 += p1
        } else {
            fmt.Println(p0, p1)
            fmt.Println(p1)
            p1 -= p1
        }
        v6 := []int{int(p1)}
        _ = v6
    }
    p1 += 4294967295
    _ = p0
    _ = p1
    return ("" + "")
}

func main() {
    v1 := []string{"", "", ("\n\t" + (("foo bar" + ("hello" + "\n\t")) + "\n\t"))}
    v2 := []uint32{}
    var v3 bool = false
    v4 := []uint32{65535, 2147483648}
    var v5 uint = 868
    fmt.Println(v5)
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int8 = -77
        _ = fi7
        _ = v8
    }
    _ = len(v4)
    v3 = ((!(!(-363 > 42))) && (("foo bar" + (("a" + "\n\t") + ("" + "foo bar"))) < ("" + ("hello" + "\n\t"))))
    var v9 uint = 65535
    switch "a" {
    case "a":
        v9 -= (v9 % ((4294967295) | 1))
    case "hello":
        v5 = (v9 * (v5 / (((v9 - (v9 - 365))) | 1)))
    case "":
        var v10 float64 = (float64(3.14159) - 1.5)
        _ = v10
    }
    var v11 float32 = 1e10
    v2 = append(v2, 2147483648)
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: 128, F1: 4294967295}
        fmt.Println(v9, v3, v3)
        _ = fi12
        _ = v13
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
