package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 int = 4294967295
    v2 := make(map[int]int)
    var v3 bool = true
    v4 := make(map[int]int)
    v1 = (^(v1))
    v5 := []int{0, -228}
    _ = len(v2)
    fmt.Println(v1)
    v1 = v1
    if v3 {
        _ = len(v5)
        switch v1 {
        case -1:
            v6 := GenS1{F0: 978, F1: 515}
            _ = v6
        case -434:
            v7 := GenS1{F0: 18446744073709551615, F1: 42}
            _ = v7
        case 42:
            var v8 uint32 = 2147483648
            _ = v8
        default:
            var v9 int64 = 0
            _ = v9
        }
    }
    v1 = (v1 << ((583) & 63))
    v3 = true
    switch ("\n\t" + (("hello" + "hello") + "\n\t")) {
    case "foo bar":
        var v10 uint64 = 0
        _ = v10
    default:
        _ = len(v4)
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        v4[32768] = (^(v1))
        var v12 float64 = 0.0
        v5 = append(v5, fi11)
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 4294967295, F1: 65535}
    v1 -= 201
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
