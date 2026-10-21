package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int }

func main() {
    var v1 uint = 882
    v2 := []string{("a" + "\n\t"), "\n\t", ""}
    v3 := make(map[string]uint32)
    var v4 int8 = -128
    if true {
        for fi5 := 0; fi5 < 3; fi5++ {
            delete(v3, "\n\t")
            v4 -= (v4 >> ((v1) & 63))
            var v6 uint = v1
            _ = fi5
            _ = v6
        }
    } else {
        if (!true) {
            var v7 string = ("\n\t" + "\n\t")
            _ = v7
        }
        _ = len(v3)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        if ((true && (1 <= 65535)) && false) {
            _ = len(v3)
            _ = len(v3)
        }
        _ = fi8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int64 = 1
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 221, F1: 2147483647, F2: 1}
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 int8 = v4
        _ = fi12
        _ = v13
    }
    var v14 int32 = -2147483648
    v14 += (^(v14))
    if false {
        var v15 uint64 = 567
        _ = v15
    } else {
        if (true || ("foo bar" != "hello")) {
            var v16 uint8 = 28
            _ = v16
        }
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 uint = v1
            v1 = v1
            var v19 uint = v18
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
