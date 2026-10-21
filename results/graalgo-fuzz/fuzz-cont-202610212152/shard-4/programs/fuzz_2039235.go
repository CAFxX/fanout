package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint8; F2 int }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 116
    v3 := make(map[uint32]uint32)
    var v4 uint = 1
    var v5 int32 = -2147483648
    v2 = (v2 & v2)
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: 189, F1: (v2 ^ v2), F2: 751}
            _ = fi7
            _ = v8
        }
        var v9 uint32 = 2147483648
        for fi10 := 0; fi10 < 7; fi10++ {
            v5 = (-(v5))
            _ = fi10
        }
        _ = fi6
        _ = v9
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            fmt.Println(v5)
            if true {
                var v13 string = (((("foo bar" + "") + ("" + "a")) + ("a" + ("\n\t" + "a"))) + "\n\t")
                var v14 uint16 = 42
                v2 = v2
                switch "" {
                case "foo bar":
                    v15 := GenS1{F0: 18446744073709551615, F1: v2, F2: 541}
                    _ = v15
                }
                _ = v13
                _ = v14
            } else {
                fmt.Println(v5)
            }
            v16 := GenS2{F0: 683, F1: v2, F2: fi11}
            _ = fi12
            _ = v16
        }
        var v17 int8 = -81
        v2 = 42
        _ = fi11
        _ = v17
    }
    fmt.Println(v5, v5)
    v2 = 128
    v18 := GenS2{F0: 4294967295, F1: 128, F2: -9223372036854775808}
    var v19 int32 = -2147483648
    var v20 uint = v4
    v21 := []int{13, -415}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
