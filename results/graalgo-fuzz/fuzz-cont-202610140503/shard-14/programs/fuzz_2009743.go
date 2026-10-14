package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int = -9223372036854775808
    var v2 int8 = -1
    v3 := []string{"a", "foo bar", "\n\t"}
    v2 = v2
    var v4 uint8 = 2
    switch "hello" {
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = ("" + "")
        }
    case "":
        switch "foo bar" {
        case "":
            v4 = (v4 + v4)
        case "hello":
            if ((!false) && (1 != -1)) {
                v4 += (v4 * 127)
            } else {
                v5 := GenS1{F0: 519}
                v2 = (-(v2))
                v2 -= 0
                _ = v5
            }
        }
    }
    v4 = (^(v4))
    switch v1 {
    case 1:
        var v6 int16 = -412
        _ = v6
    }
    v3 = append(v3, (("" + ("foo bar" + "")) + "foo bar"))
    v2 -= (v2 - v2)
    var v7 int16 = 1000
    switch "a" {
    case "a":
        fmt.Println(v7, v4, v2)
    case "hello":
        var v8 uint = 2
        _ = v8
    }
    v2 -= v2
    switch ("hello" + (("a" + "") + "\n\t")) {
    case "a":
        if false {
            var v9 float32 = 0.0
            _ = v9
        }
    case "foo bar":
        v1 -= (v1 + (v1 + (v1 + (v1 ^ v1))))
    }
    v2 = v2
    v2 += v2
    fmt.Println(v4, v2)
    v1 -= (v1 - 32768)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
