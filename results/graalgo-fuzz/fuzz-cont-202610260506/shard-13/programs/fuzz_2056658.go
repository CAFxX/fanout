package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    v3 := []int{188, 42}
    for fi4 := 0; fi4 < 20; fi4++ {
        var v5 uint32 = 2
        _ = fi4
        _ = v5
    }
    if (true || (1 > 32767)) {
        if ((((!(!false)) || (-32768 <= 1000)) && false) && false) {
            _ = len(v3)
        } else {
            switch "" {
            case "":
                for fi6 := 0; fi6 < 6; fi6++ {
                    for fi7 := 0; fi7 < 11; fi7++ {
                        var v8 string = "foo bar"
                        _ = fi7
                        _ = v8
                    }
                    _ = fi6
                }
            case "\n\t":
                var v9 uint = 2
                _ = v9
            case "hello":
                var v10 float32 = 1e10
                _ = v10
            }
        }
        for fi11 := 0; fi11 < 3; fi11++ {
            _ = len(v2)
            _ = fi11
        }
    }
    var v12 uint32 = 827
    if (32767 <= 42) {
        if true {
            for fi13 := 0; fi13 < 6; fi13++ {
                v12 += 2147483647
                v12 = (v12 % ((v12) | 1))
                _ = fi13
            }
        }
        var v14 bool = true
        for fi15 := 0; fi15 < 15; fi15++ {
            fmt.Println(v12, v12, v12)
            _ = fi15
        }
        _ = v14
    } else {
        fmt.Println(v12)
        v12 += (v12 ^ 2147483648)
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        v17 := GenS1{F0: -32768, F1: 386, F2: fi16}
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 int8 = 0
            _ = fi18
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    v20 := GenS1{F0: -32768, F1: 42, F2: 32767}
    _ = len(v3)
    v12 += (^(v12))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v3
}
