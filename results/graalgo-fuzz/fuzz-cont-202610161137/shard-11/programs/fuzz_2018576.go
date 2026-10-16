package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() uint {
    v4 := GenS1{F0: 9223372036854775807}
    v5 := GenS1{F0: -153}
    v6 := []int{0, -851, 32767}
    var v7 string = "\n\t"
    var v8 string = ""
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 65535
}
func genFn1() uint16 {
    if false {
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 uint16 = 261
            _ = fi9
            _ = v10
        }
        var v11 uint16 = 32767
        _ = v11
    }
    var v12 float64 = 0.1
    fmt.Println(v12)
    _ = v12
    return 32767
}

func main() {
    v1 := make(map[string]int)
    var v2 string = ((("\n\t" + "a") + ((("hello" + ("\n\t" + "hello")) + ("a" + ("foo bar" + ""))) + ("" + ""))) + "hello")
    var v3 uint8 = 42
    v13 := []int{32768, 0, 1, -1}
    switch 2147483647 {
    case -84:
        v2 = "\n\t"
    case 32767:
        for fi14 := 0; fi14 < 14; fi14++ {
            fmt.Println(v2, v2, v2)
            switch "a" {
            case "hello":
                v3 = (^(v3))
            case "":
                switch (fi14 | (fi14 * fi14)) {
                case 1:
                    delete(v1, v2)
                case 0:
                    switch (v2 + ((v2 + (v2 + ("" + ""))) + v2)) {
                    case "\n\t":
                        fmt.Println(fi14, fi14, fi14)
                    }
                }
            }
            if len(v13) > 0 {
                v13[0] = 4294967295
            }
            _ = fi14
        }
    }
    var v15 string = v2
    _ = len(v1)
    fmt.Println(v2, v2)
    switch -563 {
    case 42:
        v15 = ("foo bar" + ("foo bar" + (("a" + "\n\t") + "")))
    case 1:
        var v16 int8 = 100
        _ = v16
    }
    var v17 bool = (-370 < 42)
    var v18 uint64 = 1
    if len(v13) > 0 {
        v13[0] = 127
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
