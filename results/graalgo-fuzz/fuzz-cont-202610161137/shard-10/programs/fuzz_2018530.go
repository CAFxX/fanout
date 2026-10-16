package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }

func main() {
    v1 := []int{32767}
    var v2 bool = (190 != 586)
    v3 := []string{"a", "foo bar"}
    _ = len(v3)
    if (-1 != 127) {
        fmt.Println(v2)
        if (4294967295 <= 1) {
            v2 = ((("\n\t" > "") && (!v2)) && (!false))
        }
    } else {
        if true {
            var v4 int16 = 32767
            v2 = ((!((3.14159 / 1e10) < 0.0)) && (3.14159 == 1e10))
            _ = v4
        }
    }
    switch 32767 {
    case 830:
        v2 = (2 == 42)
    case -9223372036854775808:
        v2 = (-1 != 1000)
    }
    switch 2147483647 {
    case 4294967295:
        _ = len(v1)
    default:
        if false {
            var v5 uint64 = 1
            _ = v5
        }
    }
    if false {
        switch "\n\t" {
        case "":
            var v6 float32 = 0.1
            _ = v6
        case "hello":
            var v7 float32 = 1e10
            _ = v7
        default:
            v2 = ("\n\t" <= "")
        }
        var v8 int16 = 0
        _ = v8
    } else {
        v3 = append(v3, ("hello" + ("" + (("foo bar" + "") + ("a" + "\n\t")))))
    }
    v2 = ("hello" > "foo bar")
    if (2147483647 == 42) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v2)
        fmt.Println(v2)
    } else {
        v2 = true
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS1{F0: 127, F1: -128}
            _ = fi9
            _ = v10
        }
        if v2 {
            var v11 uint = 18446744073709551615
            _ = v11
        }
    }
    v2 = v2
    if (42 < 2147483647) {
        v2 = (!v2)
        v2 = (3.14159 != (-2.25 / 0.1))
        _ = len(v1)
    }
    var v12 float32 = (3.14159 - 1e10)
    fmt.Println(v2)
    var v13 int32 = -56
    _ = len(v3)
    v12 -= 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
