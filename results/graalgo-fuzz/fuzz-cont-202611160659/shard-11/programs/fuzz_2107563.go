package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 int32 }
func genFn0(p0 int64, p1 bool) string {
    var v5 int16 = 32767
    p0 += p0
    v6 := GenS1{F0: -114, F1: "foo bar", F2: v5}
    switch "foo bar" {
    case "a":
        fmt.Println(p0, v5)
    case "":
        var v7 uint16 = 255
        _ = v7
    case "foo bar":
        var v8 string = ("" + "foo bar")
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return "foo bar"
}
func genFn1() uint {
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 bool = false
        _ = fi9
        _ = v10
    }
    var v11 string = "foo bar"
    _ = v11
    return 65535
}

func main() {
    v1 := []int{127, 0, -2147483648, 4294967295}
    var v2 string = "\n\t"
    v3 := []uint32{405, 4294967295}
    var v4 int = 4294967295
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 -= -2147483648
    fmt.Println(v2, v2)
    v2 = "\n\t"
    v12 := GenS1{F0: 42, F1: "", F2: 1000}
    _ = len(v1)
    v4 = (v4 - v4)
    if false {
        var v13 int32 = 0
        v14 := GenS1{F0: (v13 - v13), F1: v2, F2: 1000}
        _ = v13
        _ = v14
    } else {
        var v15 int16 = -204
        _ = v15
    }
    if (0.0 != 0.1) {
        var v16 uint32 = 2147483648
        _ = v16
    } else {
        var v17 uint16 = 2
        var v18 int8 = 127
        v18 += (v18 + int8(v17))
        _ = v17
        _ = v18
    }
    v4 = -856
    switch (v2 + "a") {
    case "":
        if false {
            v2 = "foo bar"
        }
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v19 := GenS2{F0: 1000, F1: 2147483648, F2: 298}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
