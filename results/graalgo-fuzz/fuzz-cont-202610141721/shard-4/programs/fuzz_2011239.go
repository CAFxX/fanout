package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]uint32)
    v3 := []int{127, 9223372036854775807}
    v4 := []uint32{1}
    v5 := []uint32{2147483647, 42, 2147483648}
    var v6 string = ("a" + "\n\t")
    v7 := GenS1{F0: 65535, F1: 0}
    var v8 int64 = -417
    var v9 int16 = 1000
    _ = len(v4)
    var v10 uint16 = 32767
    v10 = (^(v10))
    _ = len(v2)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if ((true || (127 < 0)) || false) {
        if (0.1 != -2.25) {
            switch ((v6 + v6) + (("hello" + (v6 + ("\n\t" + v6))) + "a")) {
            case "hello":
                var v11 uint16 = 42
                _ = v11
            }
            v12 := GenS1{F0: 936, F1: 1}
            v13 := v2[323]
            _ = v12
            _ = v13
        }
        switch 32768 {
        case 32767:
            var v14 uint64 = 42
            _ = v14
        case 792:
            if true {
                var v15 bool = true
                _ = v15
            }
        default:
            v16 := GenS1{F0: 42, F1: 4294967295}
            _ = v16
        }
    }
    v17 := v2[2147483648]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
