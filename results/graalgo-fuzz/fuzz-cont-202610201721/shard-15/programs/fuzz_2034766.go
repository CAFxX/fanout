package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := []string{"\n\t", ((("hello" + "foo bar") + "a") + "\n\t"), (("" + ("" + "a")) + ""), ("hello" + "foo bar")}
    var v2 int16 = -32768
    v3 := []int{-733}
    v4 := GenS1{F0: 2}
    switch "hello" {
    case "\n\t":
        v2 -= v2
    case "hello":
        v2 = v2
    }
    var v5 string = "a"
    for fi6 := 0; fi6 < 8; fi6++ {
        v2 += 0
        fmt.Println(v5, fi6)
        v7 := GenS1{F0: 1}
        _ = fi6
        _ = v7
    }
    v5 = v5
    if true {
        for fi8 := 0; fi8 < 14; fi8++ {
            v2 += v2
            var v9 bool = true
            _ = fi8
            _ = v9
        }
    }
    switch 1 {
    case 32768:
        for fi10 := 0; fi10 < 15; fi10++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi10
        }
    case -623:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        _ = len(v3)
        _ = fi11
    }
    if (!false) {
        v5 = ("foo bar" + "hello")
        v12 := GenS1{F0: 1}
        _ = v12
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
