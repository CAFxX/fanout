package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int }
func genFn0() uint16 {
    var v7 uint32 = 65535
    var v8 uint16 = 65535
    for fi9 := 0; fi9 < 18; fi9++ {
        v8 = v8
        switch ("foo bar" + "\n\t") {
        case "foo bar":
            var v10 int8 = 100
            _ = v10
        case "hello":
            var v11 uint8 = uint8(v7)
            _ = v11
        case "\n\t":
            switch 9223372036854775807 {
            case 569:
                v12 := []int{fi9}
                _ = v12
            }
        default:
            switch fi9 {
            case 128:
                var v13 int32 = 1
                _ = v13
            case 32768:
                v14 := GenS1{F0: (v7 - v7)}
                _ = v14
            default:
                v15 := GenS2{F0: (100 != 1), F1: fi9}
                _ = v15
            }
        }
        v16 := []int{(fi9 | -264), (fi9 % ((0) | 1)), fi9, fi9}
        _ = fi9
        _ = v16
    }
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint16 = 585
    var v2 uint16 = (v1 + (^(v1)))
    v3 := []string{"", ("" + "foo bar"), ("\n\t" + "\n\t")}
    v4 := make(map[int]string)
    v5 := []string{("foo bar" + "\n\t"), ("" + ""), ("a" + ("foo bar" + ("hello" + "hello")))}
    var v6 int = -53
    v17 := GenS2{F0: false, F1: (v6 % ((-243) | 1))}
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    var v18 uint16 = 65535
    v4[128] = "hello"
    for fi19 := 0; fi19 < 15; fi19++ {
        v1 -= (v2 | v1)
        v6 += fi19
        _ = fi19
    }
    delete(v4, 626)
    v18 = v2
    v6 -= 0
    v18 += v1
    var v20 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
