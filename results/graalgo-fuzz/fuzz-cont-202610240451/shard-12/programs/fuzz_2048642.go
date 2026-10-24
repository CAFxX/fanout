package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int8 }
type GenS2 struct { F0 bool }

func main() {
    v1 := []string{"\n\t", ""}
    v2 := make(map[string]string)
    var v3 string = "a"
    var v4 uint64 = 1
    var v5 int = 0
    v6 := make(map[int]bool)
    v2[v3] = "\n\t"
    for fi7 := 0; fi7 < 11; fi7++ {
        v5 += -2147483648
        v3 = ((((v3 + ("foo bar" + v3)) + "foo bar") + v3) + "")
        _ = fi7
    }
    v8 := v2["a"]
    var v9 int32 = 0
    v10 := GenS1{F0: true, F1: 2147483648, F2: 45}
    v8 = "foo bar"
    v11 := GenS1{F0: (true || true), F1: 383, F2: 23}
    if false {
        v12 := GenS2{F0: (2147483647 != (v9 / (((v9 * 518)) | 1)))}
        if ((true || (!false)) || true) {
            if (-1 <= -4) {
                var v13 int8 = 1
                _ = v13
            } else {
                if (32767 < 902) {
                    v8 = ("hello" + "\n\t")
                    var v14 string = v8
                    if false {
                        var v15 int16 = 0
                        if false {
                            for fi16 := 0; fi16 < 7; fi16++ {
                                var v17 bool = false
                                v4 = (v4 >> ((65535) & 63))
                                v18 := GenS2{F0: (0 != 0)}
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                            v5 += 42
                        }
                        var v19 bool = true
                        _ = v15
                        _ = v19
                    }
                    _ = v14
                }
            }
            delete(v2, "a")
        }
        v5 += v5
        _ = v12
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
