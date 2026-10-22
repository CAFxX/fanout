package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }

func main() {
    var v1 int8 = 96
    var v2 int64 = -9223372036854775808
    var v3 uint32 = 2
    v4 := make(map[int]uint32)
    v5 := make(map[string]uint32)
    var v6 int16 = 0
    v5[("hello" + "a")] = v3
    fmt.Println(v1, v3)
    if true {
        fmt.Println(v3, v2)
    }
    var v7 uint64 = 621
    var v8 float64 = 0.0
    fmt.Println(v3, v2)
    if false {
        var v9 uint8 = 18
        _ = v9
    }
    if ((v7 >= (v7 ^ (v7 ^ v7))) || ((!true) || (v3 <= v3))) {
        if (32767 > 836) {
            switch ("\n\t" + "a") {
            case "hello":
                for fi10 := 0; fi10 < 9; fi10++ {
                    v11 := GenS1{F0: 738, F1: (fi10 ^ 206)}
                    _ = fi10
                    _ = v11
                }
            case "a":
                v12 := GenS1{F0: 18446744073709551615, F1: 9223372036854775807}
                _ = v12
            case "\n\t":
                fmt.Println(v3, v7, v2)
            }
        } else {
            fmt.Println(v8)
            v2 = (^(v2))
        }
    }
    switch ((("" + "hello") + "hello") + "\n\t") {
    case "foo bar":
        for fi13 := 0; fi13 < 12; fi13++ {
            v14 := []int{(fi13 | (-(fi13))), fi13, 4294967295}
            var v15 int32 = -515
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    v4[4294967295] = (v3 + (^(v3)))
    for fi16 := 0; fi16 < 7; fi16++ {
        delete(v5, ("foo bar" + ("foo bar" + "foo bar")))
        v17 := []int{}
        _ = fi16
        _ = v17
    }
    var v18 string = "foo bar"
    var v19 uint8 = 216
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
