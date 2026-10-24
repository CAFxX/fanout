package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint16; F1 string; F2 uint32 }
func genFn0(p0 int16, p1 uint16) bool {
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint64 = 1
        var v8 int64 = 716
        switch "hello" {
        case "a":
            switch fi6 {
            case -9223372036854775808:
                v9 := []int{32768, (^(fi6))}
                _ = v9
            }
        case "\n\t":
            v10 := []int{(-(fi6)), fi6}
            _ = v10
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v11 string = "a"
    _ = p0
    _ = p1
    _ = v11
    return true
}
func genFn1(p0 int, p1 int) int64 {
    v12 := []int{}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    _ = p0
    _ = p1
    _ = v12
    return 597
}

func main() {
    var v1 uint32 = 42
    v2 := []string{"hello", "a"}
    var v3 uint64 = 18446744073709551615
    var v4 uint64 = v3
    v5 := make(map[int]uint32)
    _ = len(v2)
    v2 = append(v2, (("foo bar" + ("" + ((("foo bar" + ("hello" + ("hello" + ("a" + ("a" + ("a" + "a")))))) + "a") + (("foo bar" + "hello") + "foo bar")))) + "a"))
    var v13 int32 = 0
    for fi14 := 0; fi14 < 1; fi14++ {
        v2 = append(v2, "")
        _ = fi14
    }
    v15 := GenS1{F0: 255, F1: -9223372036854775808}
    v1 = v1
    v16 := GenS2{F0: 0, F1: "", F2: (v1 << ((2) & 63))}
    for fi17 := 0; fi17 < 9; fi17++ {
        v1 = v1
        _ = fi17
    }
    var v18 bool = (255 != 1)
    v19 := GenS1{F0: 63, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
