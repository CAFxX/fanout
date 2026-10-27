package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 0
    var v4 uint16 = 65535
    var v5 uint8 = (v3 / ((v3) | 1))
    v6 := []string{"hello", ("a" + ("hello" + "\n\t")), "hello", ("\n\t" + "\n\t")}
    delete(v2, 1)
    var v7 int = -9223372036854775808
    if ((157 >= 127) && false) {
        var v8 uint8 = 1
        switch ((("foo bar" + ("" + "")) + (((("hello" + (("foo bar" + "a") + "\n\t")) + "") + (("foo bar" + "foo bar") + "hello")) + "hello")) + "hello") {
        case "hello":
            v2[0] = 65535
        case "foo bar":
            fmt.Println(v4)
        case "a":
            v7 = (v7 - v7)
        default:
            if (!false) {
                var v9 int64 = -9223372036854775808
                _ = v9
            } else {
                v10 := GenS1{F0: 42}
                var v11 string = "hello"
                _ = v10
                _ = v11
            }
        }
        _ = v8
    } else {
        v2[2147483647] = 874
    }
    v3 += 42
    v3 = v3
    v12 := GenS1{F0: -117}
    if true {
        var v13 int8 = 100
        _ = v13
    }
    var v14 uint = 24
    switch ((("a" + "hello") + ("\n\t" + ("\n\t" + ("hello" + "a")))) + ("foo bar" + "\n\t")) {
    case "":
        var v15 int32 = -2147483648
        _ = v15
    }
    var v16 uint16 = (v4 - v4)
    var v17 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
