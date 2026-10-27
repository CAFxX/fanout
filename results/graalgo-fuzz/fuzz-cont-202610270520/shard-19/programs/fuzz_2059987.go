package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 uint; F2 uint32 }
func genFn0(p0 int32, p1 uint16) uint64 {
    if false {
        switch (("foo bar" + "a") + ("a" + "hello")) {
        case "foo bar":
            v6 := []int{-2147483648, -350}
            _ = v6
        }
        p0 = (p0 & p0)
        var v7 float32 = 3.14159
        _ = v7
    }
    var v8 bool = (65535 > 753)
    fmt.Println(v8, p0, p1)
    _ = p0
    _ = p1
    _ = v8
    return 0
}

func main() {
    v1 := []string{("" + ("\n\t" + "hello")), "a"}
    var v2 float64 = 0.1
    var v3 float64 = (v2 - float64(3.14159))
    var v4 int = 32767
    var v5 uint = 2
    if (!false) {
        if true {
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 int64 = 42
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi9
                _ = v10
            }
        }
        v5 -= v5
        v11 := GenS1{F0: 480}
        _ = v11
    }
    v1 = append(v1, "a")
    switch "a" {
    case "\n\t":
        var v12 string = ("hello" + "hello")
        _ = v12
    default:
        var v13 string = ""
        _ = v13
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2)
    var v14 float64 = v3
    var v15 bool = true
    if (-9223372036854775808 > 610) {
        v16 := GenS1{F0: (v4 * (^(v4)))}
        var v17 uint8 = 255
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v16
        _ = v17
    } else {
        v18 := GenS2{F0: -2147483648, F1: 539, F2: 4294967295}
        v15 = (1000 <= 218)
        _ = len(v1)
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
