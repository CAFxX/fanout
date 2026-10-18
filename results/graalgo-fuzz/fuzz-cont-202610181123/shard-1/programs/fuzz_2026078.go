package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int16 }

func main() {
    var v1 float32 = float32(3.14159)
    var v2 int16 = -699
    var v3 int16 = v2
    v4 := make(map[string]int)
    var v5 int8 = -128
    v6 := []int{}
    v7 := GenS1{F0: 255, F1: ("hello" + ("\n\t" + "foo bar")), F2: 860}
    var v8 uint32 = 710
    v3 += 373
    v9 := GenS1{F0: 4294967295, F1: "a", F2: v2}
    var v10 int32 = 244
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v11 := GenS1{F0: 42, F1: "foo bar", F2: 32767}
    switch "hello" {
    case "foo bar":
        var v12 string = ("a" + "\n\t")
        _ = v12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 bool = (v5 == v5)
        for fi15 := 0; fi15 < 16; fi15++ {
            switch "hello" {
            case "hello":
                if (1.5 != float64(0.1)) {
                    var v16 uint16 = 65535
                    fmt.Println(fi15, v1, v14)
                    var v17 string = ("" + (("" + "") + "hello"))
                    _ = v16
                    _ = v17
                }
            default:
                var v18 bool = v14
                _ = v18
            }
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
