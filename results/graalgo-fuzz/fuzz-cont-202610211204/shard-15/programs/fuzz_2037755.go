package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint16 }

func main() {
    v1 := []string{"hello", "hello", ("hello" + "\n\t"), (((("" + "hello") + "a") + ("\n\t" + ("a" + ""))) + ("foo bar" + "foo bar"))}
    var v2 uint8 = 0
    var v3 int64 = -914
    v4 := make(map[int]bool)
    v5 := GenS1{F0: (v3 >> ((1) & 63))}
    v6 := GenS1{F0: 1}
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint16 = 959
        _ = fi7
        _ = v8
    }
    if (((!false) && (42 >= 42)) || false) {
        if len(v1) > 0 {
            v1[0] = (((("a" + "hello") + (("\n\t" + "foo bar") + "\n\t")) + "foo bar") + "foo bar")
        }
    } else {
        var v9 int = 831
        var v10 int16 = 569
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: -1}
    v12 := GenS1{F0: -279}
    v13 := GenS1{F0: v3}
    if (!(v3 < (v3 - (-(v3))))) {
        var v14 int8 = 8
        _ = v14
    }
    _ = len(v1)
    for fi15 := 0; fi15 < 6; fi15++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi15
    }
    v2 = v2
    if false {
        v16 := GenS2{F0: 9223372036854775807, F1: 165}
        v3 = (v3 + (v3 - (v3 + 1)))
        var v17 uint32 = 508
        _ = v16
        _ = v17
    }
    v3 += (v3 << ((18446744073709551615) & 63))
    if false {
        if (true && true) {
            var v18 uint64 = 4294967295
            fmt.Println(v18, v2)
            _ = v18
        }
        for fi19 := 0; fi19 < 15; fi19++ {
            if false {
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                v20 := v4[fi19]
                _ = v20
            }
            v21 := GenS2{F0: (^(fi19)), F1: 255}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi19
            _ = v21
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
