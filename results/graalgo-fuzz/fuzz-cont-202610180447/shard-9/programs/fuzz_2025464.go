package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 bool; F2 string }

func main() {
    var v1 int64 = 1
    v2 := make(map[int]uint32)
    v3 := []string{"\n\t", "", "hello"}
    switch "a" {
    case "foo bar":
        var v4 int = 127
        _ = v4
    default:
        for fi5 := 0; fi5 < 11; fi5++ {
            var v6 int32 = -975
            v3 = append(v3, "a")
            v7 := GenS1{F0: 0}
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    var v8 string = ("\n\t" + "hello")
    switch -886 {
    case 32768:
        v9 := GenS1{F0: 2}
        _ = v9
    case 2147483647:
        v1 = (v1 << ((137) & 63))
    case 1:
        if (true || true) {
            v10 := GenS2{F0: (v1 - v1), F1: true, F2: (v8 + ("hello" + "hello"))}
            var v11 uint8 = 198
            _ = v10
            _ = v11
        }
    }
    _ = len(v3)
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            if true {
                if (fi12 > fi12) {
                    v1 += v1
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    v8 = v8
                    _ = len(v3)
                    var v14 bool = (1e10 < (1.5 - 1e10))
                    _ = v14
                }
                switch fi13 {
                case 2147483647:
                    v15 := v2[fi12]
                    _ = v15
                case 127:
                    v8 = v8
                default:
                    if true {
                        var v16 uint16 = 2
                        _ = v16
                    } else {
                        if (v1 == v1) {
                            v1 += (v1 | v1)
                        } else {
                            fmt.Println(v1, fi12)
                        }
                    }
                }
                v8 = ("" + (v8 + v8))
            } else {
                v17 := v2[fi12]
                v17 -= 2147483648
                _ = v17
            }
            _ = fi13
        }
        fmt.Println(v1, v1)
        _ = fi12
    }
    switch 0 {
    case 42:
        v18 := GenS2{F0: v1, F1: (-2147483648 >= 32768), F2: (((v8 + ((("foo bar" + "a") + ("a" + v8)) + "hello")) + (v8 + "\n\t")) + v8)}
        _ = v18
    case 775:
        v1 += 398
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        fmt.Println(v8)
        v1 = (v1 << ((0) & 63))
        _ = fi19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        v21 := GenS1{F0: 847}
        v22 := GenS2{F0: 9223372036854775807, F1: false, F2: v8}
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
