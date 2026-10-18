package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int }
type GenS2 struct { F0 uint32; F1 uint16; F2 int32 }
func genFn0(p0 string) string {
    var v7 uint8 = 1
    v8 := []int{278}
    p0 = (p0 + (p0 + (p0 + ("hello" + "foo bar"))))
    _ = p0
    _ = v7
    _ = v8
    return (("a" + "hello") + ("hello" + ""))
}
func genFn1() int32 {
    v9 := []int{2147483647, 32767}
    var v10 int8 = 100
    _ = v10
    _ = v9
    return -1
}

func main() {
    var v1 uint32 = 2
    var v2 uint8 = 127
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    var v5 int8 = 0
    var v6 uint16 = 42
    v5 = v5
    var v11 int = 4294967295
    switch (v11 + (^(v11))) {
    case 838:
        var v12 uint16 = (^(v6))
        _ = v12
    case -2147483648:
        fmt.Println(v6)
    case 32767:
        v13 := GenS1{F0: 1, F1: (v11 > (v11 << ((263) & 63))), F2: 4294967295}
        _ = v13
    }
    fmt.Println(v2, v6, v5)
    v14 := []int{(v11 + (v11 + (v11 / (((^(v11))) | 1)))), int(v1), (v11 * -9223372036854775808), -9223372036854775808}
    v6 = (v6 * 2)
    _ = len(v14)
    switch "a" {
    case "":
        v2 += (^(v2))
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS2{F0: v1, F1: v6, F2: -1}
        v5 = (v5 | v5)
        v3["\n\t"] = true
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
