package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int8 = 107
    v2 := []int{}
    var v3 bool = true
    fmt.Println(v3, v1)
    var v4 int32 = 1
    v1 = (v1 - v1)
    _ = len(v2)
    if len(v2) > 0 {
        v2[0] = -1
    }
    var v5 int = -36
    for fi6 := 0; fi6 < 13; fi6++ {
        fmt.Println(v4)
        _ = fi6
    }
    v1 -= v1
    v4 = v4
    if (!(!(!false))) {
        fmt.Println(v4, v5, v1)
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: v4, F1: 32767, F2: 4294967295}
            _ = fi7
            _ = v8
        }
        fmt.Println(v5, v1)
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int = -1
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := GenS1{F0: (v4 - v4), F1: -561, F2: 18446744073709551615}
            if (!(255 > 985)) {
                switch 32768 {
                case -846:
                    var v13 uint8 = 2
                    _ = v13
                default:
                    switch (("foo bar" + (("foo bar" + ("hello" + "")) + ("hello" + ""))) + "hello") {
                    case "foo bar":
                        v14 := GenS2{F0: 42}
                        _ = v14
                    }
                }
                if len(v2) > 0 {
                    v2[0] = int(v1)
                }
                if len(v2) > 0 {
                    v2[0] = fi9
                }
                fmt.Println(v3)
                var v15 int64 = -1
                _ = v15
            } else {
                fmt.Println(v5, fi11, v4)
                var v16 uint = 255
                _ = v16
            }
            _ = fi11
            _ = v12
        }
        v4 = (v4 >> ((0) & 63))
        _ = fi9
        _ = v10
    }
    var v17 uint = 99
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 uint32 = 48
        _ = fi18
        _ = v19
    }
    v5 = int(v17)
    v2 = append(v2, 32767)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
