package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint; F1 uint; F2 int8 }

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    var v3 int8 = 52
    var v4 uint64 = 568
    v5 := GenS1{F0: (87 <= 0), F1: (2147483647 < -2147483648)}
    v6 := GenS1{F0: (862 == 1), F1: (((-1 < 42) && true) && false)}
    var v7 uint32 = 2147483648
    var v8 int32 = 42
    fmt.Println(v7, v4)
    v1[""] = ("foo bar" + (("\n\t" + "foo bar") + "\n\t"))
    var v9 int32 = (v8 >> ((1) & 63))
    var v10 uint64 = (^(v4))
    v4 = 1
    _ = len(v2)
    v11 := GenS1{F0: (!(true || false)), F1: ((v10 != (^(v10))) || (!(false && (!(1 != -465)))))}
    for fi12 := 0; fi12 < 2; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 uint16 = 876
            v15 := GenS2{F0: 1, F1: 4294967295, F2: -128}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
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
