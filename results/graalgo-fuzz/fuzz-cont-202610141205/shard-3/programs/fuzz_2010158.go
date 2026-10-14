package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }

func main() {
    v1 := []int{444, 444, -1, 128}
    var v2 uint16 = 32767
    var v3 int32 = -2147483648
    var v4 int16 = -64
    var v5 int32 = v3
    var v6 int = 2147483647
    if (2 < v2) {
        var v7 string = "\n\t"
        _ = v7
    } else {
        v2 += 1
        _ = len(v1)
    }
    fmt.Println(v3, v4)
    var v8 int16 = v4
    if (2 <= 198) {
        v8 = (-(v8))
        if false {
            _ = len(v1)
        }
    }
    switch ("\n\t" + (("a" + ("\n\t" + "")) + "a")) {
    case "\n\t":
        var v9 int16 = v8
        _ = v9
    case "a":
        switch int(v2) {
        case 2147483647:
            switch v6 {
            case 825:
                if (0 != -492) {
                    var v10 uint16 = v2
                    _ = v10
                }
            case 1:
                var v11 int32 = (-(v5))
                _ = v11
            default:
                if true {
                    fmt.Println(v3)
                }
            }
        default:
            v5 -= (v5 | (v5 - 2147483647))
        }
    default:
        v12 := GenS1{F0: 78, F1: true}
        _ = v12
    }
    v3 = 4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
