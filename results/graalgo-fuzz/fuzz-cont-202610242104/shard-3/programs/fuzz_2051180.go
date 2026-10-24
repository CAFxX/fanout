package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 string }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int64 = 42
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    v4 := GenS1{F0: 9223372036854775807, F1: false, F2: "foo bar"}
    v5 := GenS1{F0: -902, F1: (true && true), F2: "foo bar"}
    v1 += 9223372036854775807
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 uint16 = 65535
        v1 += (v1 - v1)
        _ = fi6
        _ = v7
    }
    if false {
        v8 := GenS1{F0: -646, F1: true, F2: "hello"}
        v9 := GenS2{F0: v1}
        _ = v8
        _ = v9
    } else {
        switch "" {
        case "\n\t":
            v1 += (-(v1))
        case "foo bar":
            var v10 int16 = 1
            _ = v10
        }
        if true {
            v11 := []int{0, -629, -1, -406}
            _ = v11
        } else {
            var v12 int32 = -1
            v13 := GenS2{F0: v1}
            _ = v12
            _ = v13
        }
    }
    v14 := []int{32768, -520, -64}
    switch -321 {
    case 4294967295:
        var v15 string = ("" + "foo bar")
        _ = v15
    case 32767:
        if false {
            for fi16 := 0; fi16 < 8; fi16++ {
                var v17 bool = true
                _ = fi16
                _ = v17
            }
        } else {
            var v18 int32 = 0
            _ = v18
        }
    }
    v19 := GenS1{F0: 367, F1: (((0.1 < 1.5) && ("hello" == (("a" + (("\n\t" + "hello") + "a")) + "hello"))) && (616 > -792)), F2: ("a" + ("foo bar" + "foo bar"))}
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
