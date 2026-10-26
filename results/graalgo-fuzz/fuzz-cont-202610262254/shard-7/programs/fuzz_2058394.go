package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint16; F1 string }

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 42
    var v3 uint64 = 2
    v2 -= v2
    if true {
        fmt.Println(v2, v2)
        var v4 uint64 = v3
        var v5 uint8 = (v2 & v2)
        _ = v4
        _ = v5
    } else {
        v6 := v1[32767]
        v7 := []int{9223372036854775807}
        _ = v6
        _ = v7
    }
    v3 = v3
    v1[-9223372036854775808] = 529
    v8 := GenS1{F0: 42}
    v3 = (v3 - 18446744073709551615)
    if (1e10 >= (1.5 + 1.5)) {
        v2 += (v2 + v2)
        if ("hello" < "hello") {
            var v9 bool = (945 > 32768)
            _ = len(v1)
            fmt.Println(v3, v9)
            _ = v9
        }
        for fi10 := 0; fi10 < 1; fi10++ {
            v2 += (v2 - v2)
            fmt.Println(v2, v3, v2)
            v11 := []int{(fi10 * (fi10 % (((-(fi10))) | 1)))}
            _ = fi10
            _ = v11
        }
    }
    var v12 int64 = 0
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS2{F0: 2, F1: (v3 + 1)}
        fmt.Println(v3, v3)
        if false {
            switch (fi13 ^ fi13) {
            case 682:
                v15 := GenS2{F0: (v2 ^ (v2 | (v2 % ((v2) | 1)))), F1: (v3 | v3)}
                _ = v15
            case 127:
                for fi16 := 0; fi16 < 18; fi16++ {
                    var v17 float64 = 3.14159
                    fmt.Println(fi13, v12)
                    _ = fi16
                    _ = v17
                }
            }
        } else {
            for fi18 := 0; fi18 < 10; fi18++ {
                if true {
                    fmt.Println(fi13)
                }
                v1[fi18] = 17
                v19 := GenS1{F0: 9223372036854775807}
                _ = fi18
                _ = v19
            }
        }
        _ = fi13
        _ = v14
    }
    v20 := GenS3{F0: 65535, F1: "\n\t"}
    v2 = v2
    v21 := GenS3{F0: 233, F1: ("foo bar" + "a")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v8
}
