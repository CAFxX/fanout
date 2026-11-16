package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint64) string {
    v4 := []int{}
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := GenS1{F0: fi5, F1: 9223372036854775807}
        _ = fi5
        _ = v6
    }
    var v7 uint32 = 0
    v4 = append(v4, -355)
    var v8 string = ("a" + ("foo bar" + (("" + "\n\t") + (("\n\t" + "\n\t") + "\n\t"))))
    _ = p0
    _ = v4
    _ = v7
    _ = v8
    return ((v8 + (v8 + "\n\t")) + ("" + (v8 + "hello")))
}
func genFn1(p0 int32) int {
    v9 := GenS1{F0: -2147483648, F1: 511}
    v10 := GenS2{F0: 988}
    var v11 uint = 1
    v12 := GenS1{F0: 42, F1: 0}
    fmt.Println(p0)
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return -2147483648
}

func main() {
    var v1 uint16 = 0
    var v2 int8 = 42
    var v3 uint32 = 2147483647
    for fi13 := 0; fi13 < 12; fi13++ {
        v14 := GenS1{F0: (fi13 << ((1) & 63)), F1: -43}
        v15 := []int{int(v2), fi13, fi13, (fi13 * (fi13 * 4294967295))}
        _ = fi13
        _ = v14
        _ = v15
    }
    switch -836 {
    case -2147483648:
        var v16 float32 = 1e10
        _ = v16
    }
    if true {
        v1 = (v1 / (((v1 + (v1 << ((1) & 63)))) | 1))
        fmt.Println(v2, v2)
        v17 := GenS1{F0: genFn1(0), F1: -1}
        _ = v17
    }
    v18 := GenS2{F0: 608}
    v19 := GenS2{F0: 2}
    v1 -= (v1 % ((v1) | 1))
    var v20 bool = (0.0 != 1.5)
    var v21 string = ((("" + (("foo bar" + "a") + "\n\t")) + ("foo bar" + "a")) + ((("" + "") + "hello") + ("foo bar" + "hello")))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
