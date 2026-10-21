package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 string; F2 int16 }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"foo bar", ("a" + (("foo bar" + "a") + (("hello" + (("" + "foo bar") + (("foo bar" + ("" + ("\n\t" + "a"))) + "hello"))) + "a"))), "foo bar", ("a" + "\n\t")}
    var v3 bool = true
    v4 := []string{"\n\t", "a", "hello", ("hello" + "hello")}
    v3 = false
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi5 := 0; fi5 < 6; fi5++ {
        v1["a"] = v3
        _ = fi5
    }
    switch -1 {
    case 127:
        var v6 uint = 65535
        _ = v6
    case 0:
        for fi7 := 0; fi7 < 8; fi7++ {
            delete(v1, ("" + ("a" + "\n\t")))
            _ = fi7
        }
    case 676:
        fmt.Println(v3)
    default:
        v8 := GenS1{F0: "foo bar"}
        _ = v8
    }
    fmt.Println(v3)
    v3 = true
    var v9 int8 = 42
    v10 := GenS2{F0: 42, F1: ("a" + "hello"), F2: -32768}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
