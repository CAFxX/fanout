package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() uint {
    var v7 string = "foo bar"
    for fi8 := 0; fi8 < 6; fi8++ {
        fmt.Println(fi8, fi8)
        fmt.Println(fi8)
        var v9 float64 = 1.5
        _ = fi8
        _ = v9
    }
    var v10 int32 = -2147483648
    fmt.Println(v10, v7)
    _ = v10
    _ = v7
    return 75
}
func genFn1(p0 int32) int16 {
    var v11 bool = (!((!(genFn0() == 0)) || false))
    v12 := []int{9223372036854775807, 1, 9223372036854775807, -1}
    var v13 uint16 = 65535
    fmt.Println(v11, p0)
    v14 := GenS1{F0: -32768}
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return -1
}

func main() {
    var v1 int8 = -40
    v2 := []uint32{}
    var v3 float64 = -2.25
    v4 := make(map[string]uint32)
    v5 := []string{"\n\t", ("foo bar" + "a")}
    var v6 uint64 = 1
    delete(v4, "")
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    var v15 string = ""
    v4["\n\t"] = 2147483647
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    switch ((v15 + ((v15 + ("a" + "hello")) + (v15 + ((v15 + ("hello" + v15)) + v15)))) + "foo bar") {
    case "":
        fmt.Println(v3)
    }
    v6 = 1
    v16 := GenS1{F0: genFn1(208)}
    var v17 int32 = -2147483648
    v6 += (v6 + 2)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
