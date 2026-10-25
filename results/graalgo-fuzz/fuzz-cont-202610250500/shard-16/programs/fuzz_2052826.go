package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 string }
func genFn0(p0 uint64, p1 bool, p2 string) bool {
    if ((0 < 128) || (10 <= 100)) {
        v6 := GenS1{F0: 0, F1: -454, F2: p2}
        var v7 uint16 = 2
        v8 := []int{-911, 0}
        _ = v6
        _ = v7
        _ = v8
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return ((0 == 0) && p1)
}
func genFn1(p0 bool) string {
    v9 := []int{42, 1}
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            switch (((("" + "foo bar") + "a") + (("\n\t" + (((("\n\t" + ("\n\t" + "foo bar")) + ((("a" + ("foo bar" + "\n\t")) + "") + "a")) + "a") + "foo bar")) + "a")) + "") {
            case "hello":
                if len(v9) > 0 {
                    v9[0] = (^(fi10))
                }
            case "\n\t":
                if p0 {
                    p0 = true
                }
            case "a":
                p0 = false
            default:
                p0 = p0
            }
            p0 = true
            var v12 uint64 = 645
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    switch 128 {
    case 32767:
        var v13 int32 = 2147483647
        _ = v13
    case 691:
        if true {
            if (0 >= 1) {
                fmt.Println(p0)
                v14 := GenS1{F0: -1, F1: 1, F2: ""}
                p0 = (((255 == 255) && false) && (18446744073709551615 >= 1))
                _ = v14
            }
        }
    }
    _ = p0
    _ = v9
    return "hello"
}

func main() {
    var v1 float32 = 0.0
    var v2 uint8 = 42
    var v3 int8 = 66
    var v4 uint16 = 728
    var v5 int = 32767
    var v15 int32 = int32(v3)
    v5 = 2147483647
    v16 := []int{}
    v1 += (v1 * 3.14159)
    switch ("a" + ("foo bar" + ("\n\t" + ""))) {
    case "foo bar":
        v15 = v15
    }
    fmt.Println(v1, v15, v2)
    v17 := GenS1{F0: 205, F1: 0, F2: ""}
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    fmt.Println(v3, v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
