package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint }

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]int)
    var v3 uint16 = 463
    var v4 int64 = -9223372036854775808
    var v5 int32 = -993
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int = fi6
            v8 = 0
            var v9 int32 = (-(v5))
            _ = fi7
            _ = v8
            _ = v9
        }
        v1 = v1
        _ = fi6
    }
    switch "a" {
    case "\n\t":
        var v10 uint64 = 653
        _ = v10
    case "":
        if (1e10 < 0.0) {
            v11 := GenS1{F0: 0, F1: -874}
            _ = v11
        } else {
            var v12 uint8 = 2
            for fi13 := 0; fi13 < 15; fi13++ {
                v5 -= -2147483648
                v14 := GenS1{F0: (v12 << ((142) & 63)), F1: 9223372036854775807}
                _ = fi13
                _ = v14
            }
            _ = v12
        }
    }
    var v15 uint64 = 822
    for fi16 := 0; fi16 < 13; fi16++ {
        v17 := []int{int(v5), (fi16 ^ int(v1)), (-(fi16))}
        if (v15 > (v15 ^ v15)) {
            for fi18 := 0; fi18 < 14; fi18++ {
                if true {
                    v19 := GenS2{F0: 127, F1: 255}
                    var v20 float32 = ((-2.25 - -2.25) / (1.5 - 1e10))
                    _ = v19
                    _ = v20
                } else {
                    fmt.Println(v1, v1)
                    if true {
                        switch (fi16 >> ((2) & 63)) {
                        case 759:
                            var v21 uint64 = (v15 ^ (v15 - v15))
                            _ = v21
                        }
                        v22 := v2[fi18]
                        _ = v22
                    }
                    var v23 string = "a"
                    _ = v23
                }
                _ = fi18
            }
            var v24 bool = ((v15 >> ((65535) & 63)) >= v15)
            for fi25 := 0; fi25 < 19; fi25++ {
                switch "\n\t" {
                case "foo bar":
                    if true {
                        if (42 == 1) {
                            _ = len(v2)
                            var v26 bool = (860 <= 42)
                            v26 = (2 >= 776)
                            _ = v26
                        } else {
                            v27 := GenS2{F0: 128, F1: 226}
                            v3 += v3
                            _ = v27
                        }
                        _ = len(v2)
                        v5 -= -2147483648
                    }
                case "a":
                    fmt.Println(fi25)
                case "hello":
                    v28 := GenS1{F0: 0, F1: -2147483648}
                    _ = v28
                }
                v29 := GenS3{F0: 4294967295}
                _ = fi25
                _ = v29
            }
            _ = v24
        } else {
            if true {
                v15 -= 0
            }
        }
        _ = fi16
        _ = v17
    }
    var v30 uint64 = 0
    switch "a" {
    case "hello":
        var v31 int8 = 100
        _ = v31
    case "\n\t":
        fmt.Println(v3, v5)
    }
    var v32 uint16 = (^(v3))
    v33 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
