package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 bool; F2 int64 }
type GenS3 struct { F0 int16; F1 uint32 }

func main() {
    var v1 int64 = 203
    var v2 uint = 406
    v3 := []int{0, 1, 272, 0}
    fmt.Println(v2)
    var v4 uint16 = 65535
    switch ("foo bar" + "") {
    case "\n\t":
        switch "" {
        case "a":
            if true {
                for fi5 := 0; fi5 < 11; fi5++ {
                    var v6 uint32 = 42
                    fmt.Println(v4)
                    switch (fi5 / ((42) | 1)) {
                    case -2147483648:
                        var v7 int64 = 42
                        _ = v7
                    }
                    _ = fi5
                    _ = v6
                }
            }
        case "\n\t":
            v8 := GenS1{F0: (v1 << ((v2) & 63))}
            _ = v8
        }
    case "foo bar":
        var v9 float32 = 0.1
        _ = v9
    default:
        v10 := GenS2{F0: "foo bar", F1: false, F2: 1}
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            v2 = (v2 / ((18446744073709551615) | 1))
            fmt.Println(fi12, v2, fi12)
            switch (fi12 >> ((v2) & 63)) {
            case -422:
                v1 -= (v1 | 982)
            default:
                fmt.Println(fi12, fi12)
            }
            _ = fi12
        }
        if (((127 != 255) && ((v1 - v1) >= 42)) || true) {
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 int16 = -32768
                _ = fi13
                _ = v14
            }
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (true && (v1 != v1)) {
                v4 = v4
                v15 := GenS2{F0: "a", F1: true, F2: v1}
                _ = v15
            }
            var v16 uint16 = (^(v4))
            _ = v16
        }
        v17 := GenS3{F0: 0, F1: 42}
        _ = fi11
        _ = v17
    }
    var v18 bool = true
    v19 := GenS1{F0: v1}
    var v20 uint8 = 255
    v2 -= 202
    var v21 uint64 = 1
    v22 := GenS2{F0: "", F1: (v18 && v18), F2: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
