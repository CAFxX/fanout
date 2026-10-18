package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
func genFn0(p0 int32, p1 bool, p2 uint16) string {
    var v5 int8 = -128
    v6 := []int{4294967295, 128, 4294967295}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch ("" + "") {
    case "foo bar":
        v5 += (v5 << ((18446744073709551615) & 63))
    case "\n\t":
        _ = len(v6)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return "hello"
}
func genFn1(p0 uint32) bool {
    var v7 string = ""
    if true {
        var v8 int64 = 42
        var v9 string = (v7 + "a")
        _ = v8
        _ = v9
    } else {
        p0 -= 0
    }
    v10 := []int{}
    _ = p0
    _ = v10
    _ = v7
    return (-32768 <= -32768)
}

func main() {
    var v1 uint32 = 2
    var v2 int = 352
    var v3 float32 = 0.1
    var v4 int32 = 0
    for fi11 := 0; fi11 < 19; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            v13 := []int{}
            fmt.Println(v3, fi12)
            var v14 uint8 = 1
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 string = "hello"
        v4 += v4
        _ = fi11
        _ = v15
    }
    var v16 int64 = 42
    for fi17 := 0; fi17 < 15; fi17++ {
        if (!false) {
            fmt.Println(fi17)
            v2 = (^(v2))
        }
        _ = fi17
    }
    v18 := []int{v2}
    v19 := GenS1{F0: v1, F1: "hello"}
    _ = len(v18)
    v3 += 0.0
    if (!(42 <= 32767)) {
        v3 -= -2.25
        _ = len(v18)
        v20 := GenS1{F0: v1, F1: ""}
        _ = v20
    } else {
        var v21 float32 = (v3 * 1.5)
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
