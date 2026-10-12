package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint64; F1 int8 }

func main() {
    var v1 uint = 593
    var v2 string = "foo bar"
    var v3 uint32 = 187
    var v4 uint64 = 335
    fmt.Println(v4, v1, v1)
    v4 = (v4 & 0)
    v5 := GenS1{F0: 128, F1: 128}
    if true {
        fmt.Println(v2, v1)
        v1 -= v1
        v2 = v2
    } else {
        switch ("\n\t" + (((v2 + v2) + ((v2 + "\n\t") + v2)) + v2)) {
        case "":
            v4 = v4
        case "hello":
            var v6 uint64 = 2
            _ = v6
        default:
            for fi7 := 0; fi7 < 13; fi7++ {
                v8 := GenS1{F0: 21, F1: fi7}
                v9 := GenS2{F0: 0, F1: 954}
                var v10 uint64 = v4
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
        }
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        v1 = 65535
        _ = fi11
    }
    v2 = (v2 + "")
    v12 := []int{42, 0, 32768}
    v13 := GenS3{F0: (v4 & (v4 * v4)), F1: -1}
    v1 += 255
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
