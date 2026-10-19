package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int16 }

func main() {
    var v1 bool = false
    var v2 int32 = -1
    var v3 uint = 42
    var v4 float32 = 1e10
    var v5 bool = v1
    v6 := []int{128, 32767}
    if v5 {
        v2 = 20
    }
    var v7 int = 128
    for fi8 := 0; fi8 < 13; fi8++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v9 int64 = 0
        v6 = append(v6, (^(v7)))
        _ = fi8
        _ = v9
    }
    var v10 int8 = -1
    if (v5 || (-32768 > 951)) {
        v5 = v5
        var v11 float32 = (v4 + 3.14159)
        _ = len(v6)
        _ = v11
    } else {
        if v5 {
            if (v3 >= (^(v3))) {
                fmt.Println(v2)
            }
            var v12 string = "hello"
            if v1 {
                _ = len(v6)
                v1 = ((v10 % ((v10) | 1)) != v10)
            }
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        if false {
            v2 = v2
        } else {
            var v14 float32 = 1.5
            _ = v14
        }
        _ = fi13
    }
    var v15 int8 = v10
    for fi16 := 0; fi16 < 9; fi16++ {
        v7 -= -418
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 float64 = (3.14159 + 1.5)
            fmt.Println(v2)
            var v19 uint8 = 255
            _ = fi17
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 11; fi20++ {
            v2 -= (v2 * (^(v2)))
            switch "\n\t" {
            case "":
                var v21 uint16 = 633
                _ = v21
            case "foo bar":
                v4 += 1e10
            }
            _ = len(v6)
            _ = fi20
        }
        _ = fi16
    }
    v22 := GenS1{F0: v3, F1: 65535}
    switch v7 {
    case 127:
        var v23 bool = v5
        _ = v23
    default:
        if (!v5) {
            fmt.Println(v5, v15, v10)
            switch (("hello" + (((("" + ("\n\t" + "a")) + ("a" + "foo bar")) + "hello") + "hello")) + "a") {
            case "foo bar":
                fmt.Println(v10)
            case "\n\t":
                v24 := GenS2{F0: 925, F1: 0}
                _ = v24
            case "a":
                v5 = ("\n\t" > "")
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
