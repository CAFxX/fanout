package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS3 struct { F0 string; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int) int64 {
    v5 := GenS1{F0: 1}
    if (65535 >= 1) {
        var v6 uint = 0
        var v7 int16 = 1000
        _ = v6
        _ = v7
    }
    _ = p0
    _ = v5
    return -9223372036854775808
}
func genFn1(p0 bool) int32 {
    p0 = (42 != -1)
    v8 := []int{2147483647}
    if ((-62 >= 32767) && (3.14159 == 3.14159)) {
        if (p0 && p0) {
            fmt.Println(p0)
        }
    }
    if p0 {
        p0 = (100 > 1)
    }
    var v9 float32 = 1e10
    _ = p0
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int = -1
    v2 := []string{(("a" + "hello") + ("\n\t" + "")), "foo bar", "\n\t"}
    v3 := make(map[string]string)
    var v4 uint16 = 1
    delete(v3, "a")
    v1 = 196
    v10 := GenS1{F0: 9223372036854775807}
    v3["foo bar"] = ""
    v11 := GenS2{F0: -32768, F1: 0, F2: 4294967295}
    switch (("foo bar" + "foo bar") + "") {
    case "\n\t":
        var v12 int32 = 1
        _ = v12
    default:
        var v13 int16 = 1
        _ = v13
    }
    fmt.Println(v4, v4)
    for fi14 := 0; fi14 < 5; fi14++ {
        v1 = 923
        v15 := GenS3{F0: "", F1: 447, F2: 10}
        var v16 string = "hello"
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS4{F0: 719}
    _ = len(v3)
    if (!true) {
        if false {
            v18 := GenS3{F0: (("hello" + ("foo bar" + "foo bar")) + ("a" + "foo bar")), F1: 2, F2: 128}
            _ = v18
        }
    } else {
        var v19 float64 = 0.0
        v20 := GenS1{F0: -1}
        _ = v19
        _ = v20
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v21 uint8 = 128
        _ = len(v3)
        _ = v21
    }
    var v22 int8 = 47
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
