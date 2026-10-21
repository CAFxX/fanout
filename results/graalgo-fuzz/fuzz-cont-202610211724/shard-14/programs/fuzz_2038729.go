package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }

func main() {
    var v1 float64 = float64(1e10)
    var v2 string = "\n\t"
    var v3 float32 = 0.1
    for fi4 := 0; fi4 < 10; fi4++ {
        v5 := GenS1{F0: 4294967295, F1: 2}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        switch "a" {
        case "hello":
            for fi7 := 0; fi7 < 8; fi7++ {
                v8 := []int{32767}
                var v9 int16 = 870
                _ = fi7
                _ = v8
                _ = v9
            }
        case "a":
            v3 -= v3
        default:
            var v10 bool = false
            _ = v10
        }
        _ = fi6
    }
    v2 = "a"
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 int = fi11
        v2 = ("foo bar" + v2)
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 0, F1: 42}
    var v14 int16 = 2
    v3 = 0.1
    if true {
        fmt.Println(v14, v1, v14)
        v15 := []int{}
        v16 := GenS1{F0: 4294967295, F1: 418}
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 10; fi17++ {
            v18 := []int{929, fi17, (fi17 ^ 9223372036854775807)}
            v18 = append(v18, fi17)
            v1 -= 1e10
            _ = fi17
            _ = v18
        }
    }
    var v19 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
}
