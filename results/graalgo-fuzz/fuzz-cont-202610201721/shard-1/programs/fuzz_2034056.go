package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int8 }

func main() {
    var v1 uint = 0
    v2 := []int{}
    v3 := make(map[string]bool)
    var v4 int16 = -1
    v4 -= 32767
    fmt.Println(v1, v1)
    fmt.Println(v1, v4)
    v1 -= 42
    if (18446744073709551615 >= uint64(v1)) {
        var v5 uint8 = 0
        var v6 int32 = 1
        if ((!(466 >= 65535)) || false) {
            v6 = v6
            fmt.Println(v1)
            switch "foo bar" {
            case "foo bar":
                for fi7 := 0; fi7 < 19; fi7++ {
                    v4 -= 42
                    var v8 uint16 = 928
                    switch ("foo bar" + (("a" + "hello") + ("hello" + ""))) {
                    case "":
                        var v9 string = "foo bar"
                        _ = v9
                    case "\n\t":
                        switch "foo bar" {
                        case "hello":
                            var v10 int = (fi7 << ((v1) & 63))
                            _ = v10
                        case "\n\t":
                            if len(v2) > 0 {
                                v2[0] = -606
                            }
                        default:
                            v5 -= (v5 - (^(v5)))
                        }
                    case "hello":
                        for fi11 := 0; fi11 < 12; fi11++ {
                            var v12 uint64 = 4294967295
                            _ = fi11
                            _ = v12
                        }
                    }
                    _ = fi7
                    _ = v8
                }
            }
        }
        _ = v5
        _ = v6
    } else {
        var v13 float64 = 1.5
        _ = v13
    }
    if (!(!true)) {
        var v14 int16 = v4
        _ = v14
    } else {
        for fi15 := 0; fi15 < 15; fi15++ {
            v1 -= v1
            v1 -= v1
            _ = fi15
        }
        v4 += v4
        var v16 bool = true
        _ = v16
    }
    var v17 int8 = 127
    v18 := GenS1{F0: 128, F1: 0, F2: (v17 & v17)}
    if false {
        var v19 uint64 = 838
        delete(v3, "\n\t")
        _ = v19
    }
    v4 += (-(v4))
    var v20 uint8 = 114
    var v21 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
