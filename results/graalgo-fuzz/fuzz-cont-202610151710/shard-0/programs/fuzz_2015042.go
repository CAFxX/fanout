package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = true
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    var v5 int8 = 127
    v6 := make(map[int]string)
    switch 42 {
    case -2147483648:
        var v7 uint32 = 1
        _ = v7
    case -1:
        var v8 uint = 886
        _ = v8
    case 32767:
        var v9 uint64 = 220
        _ = v9
    default:
        var v10 uint16 = 523
        _ = v10
    }
    fmt.Println(v2, v2)
    v2 = (0 <= 9223372036854775807)
    v5 -= 127
    v2 = (!v2)
    switch "" {
    case "a":
        var v11 bool = true
        _ = v11
    case "":
        if true {
            for fi12 := 0; fi12 < 17; fi12++ {
                for fi13 := 0; fi13 < 5; fi13++ {
                    switch fi13 {
                    case 459:
                        v5 += 42
                    default:
                        var v14 float32 = 1e10
                        _ = v14
                    }
                    var v15 uint16 = 1
                    var v16 uint = uint(v15)
                    _ = fi13
                    _ = v15
                    _ = v16
                }
                _ = fi12
            }
        } else {
            switch ("foo bar" + ("\n\t" + ((("\n\t" + ("\n\t" + "hello")) + ("hello" + "")) + ""))) {
            case "\n\t":
                var v17 int32 = 42
                _ = v17
            }
            var v18 string = "foo bar"
            _ = v18
        }
    case "foo bar":
        v19 := []int{}
        _ = v19
    }
    var v20 string = ""
    v21 := GenS1{F0: v2, F1: v2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
