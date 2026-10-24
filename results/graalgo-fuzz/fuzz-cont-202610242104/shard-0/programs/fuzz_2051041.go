package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int64; F2 int64 }

func main() {
    v1 := []string{"foo bar", ("a" + (("\n\t" + ("hello" + "a")) + "hello"))}
    var v2 uint32 = 65535
    var v3 string = ""
    var v4 uint = 1
    v5 := GenS1{F0: true}
    switch (v3 + ("a" + ((((("foo bar" + "foo bar") + (v3 + ("a" + (((v3 + v3) + (("hello" + "hello") + "hello")) + v3)))) + v3) + "foo bar") + "foo bar"))) {
    case "hello":
        var v6 uint64 = 0
        _ = v6
    case "a":
        fmt.Println(v3)
    case "foo bar":
        v3 = v3
    default:
        v2 -= 2
    }
    v7 := GenS2{F0: 0, F1: -276, F2: -9223372036854775808}
    for fi8 := 0; fi8 < 17; fi8++ {
        switch (fi8 + 128) {
        case -2147483648:
            var v9 int = fi8
            _ = v9
        case 4294967295:
            var v10 uint16 = 232
            _ = v10
        case 1:
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 string = ((("hello" + "") + "\n\t") + v3)
                _ = fi11
                _ = v12
            }
        default:
            v4 -= 308
        }
        v3 = v3
        _ = fi8
    }
    if true {
        v4 += (v4 & v4)
        v3 = "foo bar"
    }
    v1 = append(v1, v3)
    if (935 > 1) {
        for fi13 := 0; fi13 < 13; fi13++ {
            _ = len(v1)
            _ = fi13
        }
        if len(v1) > 0 {
            v1[0] = v3
        }
        if len(v1) > 0 {
            v1[0] = ((v3 + "hello") + (("hello" + v3) + ""))
        }
    }
    var v14 float32 = 3.14159
    var v15 uint16 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
