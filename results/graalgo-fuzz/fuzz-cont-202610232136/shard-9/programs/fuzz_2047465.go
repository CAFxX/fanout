package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint8 }
func genFn0() string {
    var v5 int32 = 1
    v5 -= -2147483648
    var v6 float32 = 0.0
    _ = v5
    _ = v6
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = "foo bar"
    v3 := []string{"hello"}
    v4 := []int{2147483647}
    _ = len(v3)
    _ = len(v4)
    if (1 >= -667) {
        v7 := GenS1{F0: 42, F1: -32768}
        v1[1] = 32767
        switch "" {
        case "a":
            _ = len(v1)
        case "hello":
            v8 := GenS2{F0: 0}
            _ = v8
        default:
            var v9 string = (v2 + "")
            _ = v9
        }
        _ = v7
    }
    fmt.Println(v2)
    fmt.Println(v2)
    fmt.Println(v2)
    v2 = ((v2 + (v2 + ("hello" + v2))) + v2)
    if (false && ((false && false) && (-1 == -1))) {
        for fi10 := 0; fi10 < 19; fi10++ {
            switch fi10 {
            case 0:
                v3 = append(v3, v2)
            }
            _ = fi10
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
