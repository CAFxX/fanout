package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int32 }
func genFn0() int32 {
    v6 := GenS1{F0: "\n\t", F1: 65535, F2: -1}
    var v7 float64 = 3.14159
    var v8 int8 = 127
    v9 := []int{831, 4294967295, 42, 2147483647}
    var v10 int16 = 1
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 805
}

func main() {
    var v1 uint64 = 42
    var v2 uint16 = 261
    var v3 int16 = 0
    v4 := make(map[int]bool)
    var v5 int8 = 127
    fmt.Println(v3, v3)
    fmt.Println(v2, v3)
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 bool = false
        _ = fi11
        _ = v12
    }
    var v13 int32 = 0
    delete(v4, int(v1))
    if false {
        switch ("hello" + ("\n\t" + "foo bar")) {
        case "":
            v14 := []int{32767, -945, 127}
            _ = v14
        case "\n\t":
            var v15 uint64 = (v1 | v1)
            _ = v15
        case "foo bar":
            v2 += 982
        }
        switch 32767 {
        case 128:
            var v16 uint8 = 42
            _ = v16
        }
        v17 := GenS1{F0: "\n\t", F1: 1, F2: v13}
        _ = v17
    } else {
        v13 = (v13 - v13)
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        if true {
            var v19 uint16 = 2
            for fi20 := 0; fi20 < 11; fi20++ {
                delete(v4, (fi18 % ((fi18) | 1)))
                _ = fi20
            }
            _ = v19
        }
        fmt.Println(v3)
        _ = fi18
    }
    v21 := v4[32768]
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
