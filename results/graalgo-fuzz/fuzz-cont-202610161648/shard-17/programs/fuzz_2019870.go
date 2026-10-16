package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    var v3 uint16 = 2
    if true {
        for fi4 := 0; fi4 < 13; fi4++ {
            v3 = v3
            _ = fi4
        }
        v5 := v1[""]
        _ = v5
    }
    v6 := GenS1{F0: 1, F1: "\n\t"}
    v3 += 255
    v3 -= 151
    var v7 string = "foo bar"
    v3 = (v3 << ((18446744073709551615) & 63))
    v8 := GenS1{F0: -32768, F1: ((v7 + ("foo bar" + "hello")) + v7)}
    v9 := GenS2{F0: (v3 ^ (v3 << ((688) & 63))), F1: 4294967295, F2: 60}
    v1["hello"] = (v7 + v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
