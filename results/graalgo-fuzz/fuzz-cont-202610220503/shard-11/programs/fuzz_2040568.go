package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint32 = 42
    var v2 uint8 = uint8(v1)
    v3 := []string{"hello", "a", "\n\t", (("" + "hello") + "hello")}
    var v4 int32 = 2147483647
    var v5 int = 4294967295
    v6 := []int{v5, v5}
    v6 = append(v6, v5)
    var v7 int = v5
    fmt.Println(v2, v4, v2)
    v5 = (v5 % (((-(v5))) | 1))
    v5 = (v7 ^ (-(v7)))
    if true {
        v7 -= v7
        if (127 <= 127) {
            var v8 int32 = v4
            v6 = append(v6, -384)
            for fi9 := 0; fi9 < 7; fi9++ {
                _ = len(v3)
                _ = fi9
            }
            _ = v8
        }
    } else {
        v1 = v1
        v7 = (^(v7))
        for fi10 := 0; fi10 < 20; fi10++ {
            fmt.Println(v1)
            _ = fi10
        }
    }
    v11 := GenS1{F0: true}
    v2 = (v2 ^ v2)
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            fmt.Println(v7)
            switch 127 {
            case 1:
                v14 := GenS1{F0: true}
                _ = v14
            case 128:
                var v15 uint = 633
                _ = v15
            case 127:
                switch "a" {
                case "hello":
                    fmt.Println(fi12, v2)
                default:
                    fmt.Println(fi12)
                }
            default:
                v4 += v4
            }
            _ = fi13
        }
        _ = fi12
    }
    v1 = 2147483648
    var v16 int32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
