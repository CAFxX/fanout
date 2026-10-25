package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := make(map[int]int)
    var v2 int = 42
    v3 := make(map[uint32]int)
    var v4 string = "hello"
    if (false || (!false)) {
        v2 = v2
        for fi5 := 0; fi5 < 14; fi5++ {
            v2 = fi5
            var v6 int16 = 674
            var v7 string = (v4 + ("a" + v4))
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := GenS1{F0: "hello"}
        _ = v8
    } else {
        switch v2 {
        case 272:
            var v9 uint16 = 65535
            _ = v9
        case 966:
            for fi10 := 0; fi10 < 17; fi10++ {
                v2 = v2
                fmt.Println(v2, fi10)
                var v11 int = v2
                _ = fi10
                _ = v11
            }
        }
        var v12 int64 = -909
        var v13 string = ((((((((v4 + "foo bar") + v4) + "hello") + (v4 + v4)) + v4) + (v4 + v4)) + v4) + v4)
        _ = v12
        _ = v13
    }
    v2 = 4294967295
    _ = len(v1)
    if (0 < 1) {
        v14 := GenS1{F0: (v4 + (v4 + (v4 + (("\n\t" + v4) + (v4 + (v4 + v4))))))}
        _ = v14
    }
    fmt.Println(v2)
    v15 := []int{v2}
    switch v4 {
    case "":
        v16 := GenS1{F0: (v4 + "hello")}
        _ = v16
    case "\n\t":
        v4 = ""
    case "a":
        v15 = append(v15, -1)
    }
    fmt.Println(v4, v4)
    var v17 bool = (!false)
    fmt.Println(v4, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
