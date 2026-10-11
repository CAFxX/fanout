package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 string }

func main() {
    v1 := []int{648}
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    var v4 int32 = -1
    var v5 bool = ((0.0 <= float32(1e10)) && (-702 > 9223372036854775807))
    var v6 int32 = (v4 * (v4 * 556))
    var v7 uint = 4294967295
    v4 = (v4 | v4)
    v4 -= 726
    v6 = v6
    fmt.Println(v4, v6)
    var v8 int64 = 42
    v8 = (v8 / ((9223372036854775807) | 1))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v9 := GenS1{F0: v4, F1: 91, F2: ((("foo bar" + ("hello" + "\n\t")) + "hello") + ("\n\t" + "foo bar"))}
    fmt.Println(v5, v6)
    for fi10 := 0; fi10 < 12; fi10++ {
        if ("foo bar" > ("hello" + "\n\t")) {
            v5 = v5
            switch fi10 {
            case 128:
                v6 = v6
            case 85:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                for fi11 := 0; fi11 < 11; fi11++ {
                    var v12 uint8 = 195
                    if (1 == (v4 - (v6 ^ v4))) {
                        v13 := GenS1{F0: (v6 | v4), F1: 1, F2: "a"}
                        v12 -= (v12 - 1)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v13
                    }
                    fmt.Println(v4, v6, fi10)
                    _ = fi11
                    _ = v12
                }
            }
        } else {
            var v14 float32 = 3.14159
            _ = v14
        }
        v6 = (v4 % ((v6) | 1))
        _ = fi10
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
