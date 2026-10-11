package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int16 }
func genFn0(p0 uint16, p1 uint16, p2 uint8) int32 {
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 int16 = 312
        if (true || (0.0 >= 3.14159)) {
            var v6 int64 = -849
            p2 = 127
            var v7 int16 = -32768
            _ = v6
            _ = v7
        } else {
            p1 -= p0
            var v8 uint32 = 591
            switch (fi4 ^ (fi4 * 42)) {
            case -2147483648:
                switch "a" {
                case "\n\t":
                    var v9 int = (fi4 | fi4)
                    _ = v9
                default:
                    var v10 uint = 522
                    _ = v10
                }
            }
            _ = v8
        }
        _ = fi4
        _ = v5
    }
    v11 := GenS1{F0: 1000, F1: 0, F2: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 2147483647
}
func genFn1() int16 {
    switch 2147483647 {
    case -1:
        var v12 int32 = 92
        _ = v12
    case 42:
        var v13 uint32 = 2147483647
        _ = v13
    case -344:
        v14 := []int{9223372036854775807}
        _ = v14
    }
    var v15 string = "a"
    _ = v15
    return 32767
}

func main() {
    var v1 bool = false
    v2 := []int{32767, -578}
    var v3 uint64 = 0
    for fi16 := 0; fi16 < 10; fi16++ {
        fmt.Println(fi16, v3)
        if len(v2) > 0 {
            v2[0] = 9223372036854775807
        }
        _ = fi16
    }
    if (32767 > 65535) {
        v3 = (v3 / (((v3 >> ((18446744073709551615) & 63))) | 1))
    }
    v17 := GenS1{F0: -359, F1: 42, F2: 886}
    switch "hello" {
    case "foo bar":
        v18 := GenS1{F0: 857, F1: -885, F2: -32768}
        _ = v18
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "":
        var v19 uint32 = 2147483647
        _ = v19
    }
    var v20 uint32 = 4294967295
    var v21 uint16 = 2
    switch 4294967295 {
    case 42:
        if len(v2) > 0 {
            v2[0] = -1
        }
    case -732:
        switch 4294967295 {
        case 1:
            v22 := GenS1{F0: -1, F1: -1, F2: 0}
            _ = v22
        }
    }
    _ = len(v2)
    var v23 int64 = -1
    fmt.Println(v20, v21, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
}
