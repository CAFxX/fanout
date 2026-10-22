package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int32 }
func genFn0(p0 int) int64 {
    p0 += 2147483647
    v5 := GenS1{F0: 18446744073709551615, F1: 144, F2: 1}
    var v6 int32 = 0
    v7 := []int{}
    fmt.Println(v6, v6)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -1
}

func main() {
    v1 := []string{("a" + "a"), (("" + ("foo bar" + "")) + ("a" + "")), ("a" + "foo bar"), ("hello" + "foo bar")}
    var v2 uint = 0
    v3 := []int{}
    v4 := []int{164, -2147483648, 127, -9223372036854775808}
    for fi8 := 0; fi8 < 15; fi8++ {
        if false {
            v2 += (v2 % (((v2 * (v2 % ((v2) | 1)))) | 1))
            var v9 uint8 = 42
            _ = v9
        }
        _ = fi8
    }
    var v10 bool = (-1 != -2147483648)
    v3 = append(v3, -878)
    fmt.Println(v2, v10)
    if (1 > -24) {
        var v11 int16 = 0
        var v12 float32 = 0.1
        _ = v11
        _ = v12
    }
    v10 = (!true)
    fmt.Println(v2, v2)
    var v13 int = 32767
    var v14 bool = v10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
