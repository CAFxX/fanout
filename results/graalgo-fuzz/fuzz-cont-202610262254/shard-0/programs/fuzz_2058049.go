package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint32 }

func main() {
    var v1 uint32 = 767
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    switch 42 {
    case 895:
        v4 := []int{974, 127, 4294967295}
        _ = v4
    case 32768:
        switch -977 {
        case 9223372036854775807:
            v5 := v2[""]
            _ = v5
        case -411:
            var v6 int8 = 100
            _ = v6
        case -1:
            switch 32768 {
            case 42:
                v3["a"] = v1
            case 127:
                for fi7 := 0; fi7 < 19; fi7++ {
                    v8 := []int{(fi7 % ((fi7) | 1))}
                    v1 += v1
                    for fi9 := 0; fi9 < 19; fi9++ {
                        v8 = append(v8, fi7)
                        var v10 int = fi7
                        _ = fi9
                        _ = v10
                    }
                    _ = fi7
                    _ = v8
                }
            case -952:
                delete(v2, "a")
            }
        }
    case 128:
        v1 += 2
    }
    if false {
        v11 := GenS1{F0: 0, F1: v1}
        v1 += v1
        _ = v11
    } else {
        v1 += 826
    }
    v12 := GenS2{F0: 4294967295, F1: 0, F2: v1}
    var v13 string = "a"
    var v14 uint8 = 2
    var v15 int32 = 1
    var v16 string = v13
    delete(v2, (v16 + ("\n\t" + "foo bar")))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
