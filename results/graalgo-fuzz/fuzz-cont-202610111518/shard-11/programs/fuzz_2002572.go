package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 832
    var v3 string = "hello"
    v4 := make(map[string]bool)
    _ = len(v4)
    v5 := GenS1{F0: -280}
    fmt.Println(v2)
    var v6 int = 128
    v7 := GenS2{F0: v3}
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            switch v3 {
            case "foo bar":
                fmt.Println(fi9, fi9)
            case "":
                v3 = v3
            case "hello":
                if ((true || true) || false) {
                    fmt.Println(v2)
                }
            default:
                v10 := []int{}
                _ = v10
            }
            switch 4294967295 {
            case 699:
                v11 := GenS2{F0: ""}
                _ = v11
            case 4294967295:
                v3 = v3
            default:
                fmt.Println(v2, v2)
            }
            _ = fi9
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 int = 830
        fmt.Println(v13)
        v6 += (fi12 % ((fi12) | 1))
        _ = fi12
        _ = v13
    }
    var v14 string = (("" + v3) + "hello")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
