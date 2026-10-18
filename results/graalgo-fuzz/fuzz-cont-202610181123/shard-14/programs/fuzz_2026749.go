package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int32 }

func main() {
    var v1 int8 = 26
    v2 := make(map[uint32]bool)
    var v3 int16 = -1
    v4 := []string{"hello", "a"}
    v5 := make(map[uint32]string)
    v6 := []string{"", ("a" + "\n\t")}
    v7 := GenS1{F0: 32767, F1: 127, F2: 731}
    v1 -= (v1 ^ v1)
    switch 32767 {
    case 428:
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 string = "hello"
            var v10 uint64 = 93
            if len(v4) > 0 {
                v4[0] = v9
            }
            _ = fi8
            _ = v10
            _ = v9
        }
    case 2147483647:
        v11 := GenS1{F0: 125, F1: 42, F2: 42}
        _ = v11
    }
    switch "\n\t" {
    case "":
        v1 -= v1
    case "a":
        if (2 > 42) {
            var v12 uint = 829
            _ = v12
        }
    case "foo bar":
        switch (("" + (("a" + "\n\t") + "")) + "\n\t") {
        case "a":
            v2[0] = false
        }
    }
    v3 = (v3 + 0)
    v4 = append(v4, ("\n\t" + ("foo bar" + "a")))
    fmt.Println(v1, v3)
    v3 += (^(v3))
    switch "a" {
    case "hello":
        v1 -= v1
    case "":
        v1 = (v1 - v1)
    case "a":
        var v13 string = (("\n\t" + "a") + "foo bar")
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
