package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 bool = (true && (-9223372036854775808 > 9223372036854775807))
    var v2 uint8 = 255
    v3 := []int{-9223372036854775808}
    v4 := []string{"\n\t"}
    v5 := make(map[int]bool)
    var v6 int32 = 0
    for fi7 := 0; fi7 < 10; fi7++ {
        v2 += (v2 + v2)
        _ = fi7
    }
    fmt.Println(v6, v6)
    v8 := GenS1{F0: 127}
    v6 = (v6 - (^(v6)))
    switch "a" {
    case "a":
        var v9 int64 = 1
        _ = v9
    case "foo bar":
        v4 = append(v4, "\n\t")
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 string = (("a" + "") + "")
            var v13 int16 = 42
            _ = fi11
            _ = v12
            _ = v13
        }
        v14 := GenS1{F0: fi10}
        _ = fi10
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v1 = true
        _ = fi15
    }
    var v16 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
