package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 string; F2 int32 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := []int{}
    var v2 int = 9223372036854775807
    v3 := make(map[string]bool)
    var v4 int8 = 21
    if (!false) {
        if true {
            fmt.Println(v2, v2)
        }
    }
    v4 = v4
    v4 = v4
    fmt.Println(v2, v2)
    for fi5 := 0; fi5 < 14; fi5++ {
        if (!(true && (!(243 > 317)))) {
            for fi6 := 0; fi6 < 20; fi6++ {
                fmt.Println(v4)
                _ = fi6
            }
            var v7 string = "\n\t"
            _ = v7
        }
        _ = fi5
    }
    if (true || (1 >= 128)) {
        for fi8 := 0; fi8 < 16; fi8++ {
            v4 = (^(v4))
            _ = fi8
        }
    } else {
        switch ("hello" + "foo bar") {
        case "foo bar":
            var v9 string = "\n\t"
            _ = v9
        }
        if false {
            fmt.Println(v2, v4)
            v4 = -53
            v10 := GenS1{F0: 128}
            _ = v10
        }
    }
    var v11 bool = false
    fmt.Println(v11, v4, v4)
    if (v11 || (!false)) {
        v12 := GenS2{F0: 128, F1: "a", F2: 42}
        _ = v12
    } else {
        v4 = (^(v4))
        v13 := GenS3{F0: 959}
        _ = v13
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
