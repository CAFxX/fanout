package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }

func main() {
    var v1 int32 = -2147483648
    var v2 int = 4294967295
    v3 := make(map[string]int)
    var v4 int32 = -79
    fmt.Println(v1, v4, v4)
    if (-287 == 0) {
        for fi5 := 0; fi5 < 6; fi5++ {
            delete(v3, "hello")
            switch 127 {
            case 42:
                v3[""] = fi5
            case 1:
                fmt.Println(v1, v1)
            }
            switch v2 {
            case 1:
                delete(v3, ("\n\t" + (("hello" + ("" + "hello")) + "a")))
            case 2147483647:
                if (2 > 4294967295) {
                    fmt.Println(v2, fi5, v4)
                    v3[""] = (v2 << ((42) & 63))
                } else {
                    var v6 uint64 = 42
                    _ = v6
                }
            default:
                fmt.Println(v4, v2)
            }
            _ = fi5
        }
        v7 := GenS1{F0: 18446744073709551615, F1: v4}
        _ = v7
    }
    v4 -= (v1 >> ((42) & 63))
    var v8 uint64 = 18446744073709551615
    v8 = (v8 ^ (v8 ^ (v8 >> ((42) & 63))))
    switch ("hello" + "foo bar") {
    case "hello":
        v9 := GenS1{F0: 1, F1: v1}
        _ = v9
    default:
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 float64 = -2.25
            var v12 uint64 = v8
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    switch "a" {
    case "a":
        v4 += 42
    default:
        v3[("foo bar" + "")] = (^(v2))
    }
    if (("" > (("\n\t" + ((("hello" + "a") + "hello") + (("a" + "") + ("\n\t" + "foo bar")))) + "foo bar")) && true) {
        v4 -= v1
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 int = 128
            _ = fi13
            _ = v14
        }
        var v15 uint = 18446744073709551615
        _ = v15
    }
    v4 = -2147483648
    var v16 uint64 = (v8 + v8)
    var v17 int8 = int8(v1)
    var v18 uint = 550
    if (!false) {
        v4 = v1
    }
    if false {
        delete(v3, (("a" + "hello") + "foo bar"))
        fmt.Println(v16, v4, v18)
        var v19 int8 = v17
        _ = v19
    }
    fmt.Println(v16, v17, v18)
    v20 := GenS1{F0: 42, F1: (v4 << (((v18 << ((747) & 63))) & 63))}
    var v21 int64 = -9223372036854775808
    var v22 int8 = (v17 / ((v17) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v8
}
