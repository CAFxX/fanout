package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint32 }

func main() {
    var v1 uint = 4294967295
    v2 := []string{"a", (("a" + ("" + "\n\t")) + "foo bar")}
    var v3 uint8 = 42
    v4 := make(map[int]int)
    var v5 string = "a"
    v6 := make(map[uint32]bool)
    delete(v6, 1)
    _ = len(v2)
    if (true && true) {
        v2 = append(v2, ("hello" + (v5 + (v5 + "foo bar"))))
    }
    v4[-1] = -1
    fmt.Println(v5)
    var v7 int8 = -128
    if true {
        var v8 float64 = (-2.25 * 3.14159)
        _ = v8
    } else {
        v9 := GenS1{F0: 255}
        v1 = 0
        _ = v9
    }
    _ = len(v2)
    var v10 string = "a"
    var v11 int8 = (v7 ^ -76)
    v12 := GenS2{F0: true, F1: 42}
    if true {
        for fi13 := 0; fi13 < 14; fi13++ {
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := GenS1{F0: 0}
                fmt.Println(v1, v10)
                _ = fi14
                _ = v15
            }
            v16 := GenS2{F0: (!((v5 + (((v10 + "\n\t") + ((v10 + "a") + v10)) + (("" + v10) + v5))) == "")), F1: 42}
            _ = fi13
            _ = v16
        }
        delete(v4, -986)
    }
    if (!(492 <= 0)) {
        v5 = (v5 + (v5 + (((v10 + v5) + v10) + v5)))
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := v4[fi17]
        switch v10 {
        case "foo bar":
            var v19 int = fi17
            _ = v19
        case "\n\t":
            if (true || (!true)) {
                var v20 int32 = 2147483647
                _ = v20
            } else {
                v6[42] = true
                var v21 int32 = 2147483647
                fmt.Println(v18, v1)
                _ = v21
            }
        case "hello":
            v7 = -123
        }
        var v22 uint32 = 521
        _ = fi17
        _ = v18
        _ = v22
    }
    switch 0 {
    case 1:
        v4[-9223372036854775808] = 127
    case -929:
        var v23 uint8 = v3
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
