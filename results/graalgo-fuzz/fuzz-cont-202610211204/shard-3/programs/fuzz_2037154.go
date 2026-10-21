package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS3 struct { F0 int; F1 uint16; F2 string }
type GenS4 struct { F0 uint; F1 int64; F2 bool }
func genFn0() int32 {
    var v7 int16 = 42
    var v8 uint8 = 42
    v9 := GenS1{F0: v8}
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS1{F0: 127}
        _ = fi10
        _ = v11
    }
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = -412
    var v3 int64 = 1
    var v4 bool = true
    var v5 uint32 = 4294967295
    v6 := []int{32767, 4294967295, 2147483647}
    v12 := GenS2{F0: 778, F1: 175, F2: v5}
    v4 = v4
    var v13 uint16 = 1
    if true {
        switch 128 {
        case 4294967295:
            var v14 uint16 = (v13 | v13)
            _ = v14
        case -9223372036854775808:
            _ = len(v1)
        case -676:
            v1[v5] = v4
        }
        v1[780] = true
        _ = len(v6)
    }
    var v15 bool = (v4 && ((v13 | (v13 / ((v13) | 1))) == (v13 | 0)))
    switch 128 {
    case 128:
        v5 = (^(v5))
    default:
        var v16 uint64 = 675
        _ = v16
    }
    v17 := GenS2{F0: (v13 | (v13 >> ((52) & 63))), F1: 229, F2: (v5 & 0)}
    for fi18 := 0; fi18 < 2; fi18++ {
        v5 -= 65535
        _ = fi18
    }
    var v19 uint = 255
    v4 = v15
    for fi20 := 0; fi20 < 5; fi20++ {
        fmt.Println(v13, v2)
        v21 := GenS1{F0: 128}
        _ = fi20
        _ = v21
    }
    if true {
        if v15 {
            fmt.Println(v4, v15, v19)
            fmt.Println(v2)
            for fi22 := 0; fi22 < 6; fi22++ {
                v23 := GenS3{F0: fi22, F1: 0, F2: ("hello" + "a")}
                for fi24 := 0; fi24 < 7; fi24++ {
                    var v25 int16 = 1
                    if (!false) {
                        fmt.Println(v2, v4, v4)
                        if len(v6) > 0 {
                            v6[0] = fi24
                        }
                    } else {
                        v26 := GenS2{F0: v13, F1: (v25 + (v25 >> ((v19) & 63))), F2: (v5 ^ (v5 << ((42) & 63)))}
                        _ = v26
                    }
                    v19 -= 0
                    v27 := GenS1{F0: 0}
                    if false {
                        switch (("foo bar" + "a") + "hello") {
                        case "\n\t":
                            fmt.Println(v19, v2)
                        }
                        v3 -= 353
                    } else {
                        v6 = append(v6, (-(fi24)))
                        v19 += 690
                    }
                    switch "hello" {
                    case "a":
                        v13 += 65535
                    }
                    v13 += 42
                    _ = fi24
                    _ = v25
                    _ = v27
                }
                if v15 {
                    fmt.Println(v15, v5, v5)
                    _ = len(v1)
                }
                _ = fi22
                _ = v23
            }
        } else {
            v28 := GenS4{F0: v19, F1: v3, F2: ((v5 >> ((v19) & 63)) <= (v5 % ((940) | 1)))}
            _ = v28
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v29 := GenS4{F0: v19, F1: (v3 >> ((v19) & 63)), F2: v15}
    fmt.Println(v2, v5)
    v2 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v17
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
