package main

import (
    "fmt"
)


func main() {
    v1 := []string{("hello" + (((("" + "hello") + ("" + "")) + ("a" + "hello")) + "foo bar")), "a", "a", "a"}
    v2 := make(map[int]int)
    var v3 int32 = -1
    v4 := []string{("\n\t" + "\n\t")}
    var v5 uint8 = 255
    var v6 float64 = -2.25
    switch ("\n\t" + (("hello" + "foo bar") + "foo bar")) {
    case "":
        switch "foo bar" {
        case "foo bar":
            var v7 int64 = -9223372036854775808
            _ = v7
        case "hello":
            var v8 float32 = 0.0
            _ = v8
        default:
            fmt.Println(v6, v6, v6)
        }
    default:
        var v9 uint8 = v5
        _ = v9
    }
    var v10 float64 = (((3.14159 * v6) * v6) * 1.5)
    v10 = v6
    switch 145 {
    case 1:
        delete(v2, 9223372036854775807)
    case 42:
        if false {
            var v11 int8 = 1
            var v12 float64 = v10
            _ = v11
            _ = v12
        } else {
            v10 -= (v6 - v6)
        }
    case 128:
        v5 = (^(v5))
    }
    var v13 int16 = -1
    var v14 uint64 = 0
    delete(v2, 4294967295)
    switch ((("foo bar" + "") + "foo bar") + (("" + "foo bar") + "hello")) {
    case "":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if true {
        v2[-2147483648] = 2147483647
        _ = len(v2)
    } else {
        v5 -= (v5 % (((v5 + 1)) | 1))
        _ = len(v2)
    }
    v1 = append(v1, "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
