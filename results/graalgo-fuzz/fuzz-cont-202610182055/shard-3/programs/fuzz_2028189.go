package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int8 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []uint32{1}
    v2 := make(map[int]string)
    var v3 int32 = -2147483648
    v4 := make(map[int]int)
    var v5 int16 = -1
    v5 = (v5 << ((571) & 63))
    var v6 int64 = 42
    v2[0] = (("a" + ("hello" + "")) + "hello")
    var v7 int8 = int8(v5)
    if false {
        var v8 int = 4294967295
        _ = v8
    } else {
        for fi9 := 0; fi9 < 14; fi9++ {
            if false {
                if false {
                    _ = len(v1)
                }
                var v10 int16 = v5
                _ = v10
            }
            v11 := GenS1{F0: 32767, F1: (fi9 / (((fi9 | -2147483648)) | 1)), F2: 42}
            var v12 string = ""
            _ = fi9
            _ = v11
            _ = v12
        }
        v7 = v7
        v6 = -844
    }
    v5 = (v5 * 1000)
    v13 := v2[1]
    v6 = (v6 << ((31) & 63))
    v6 = 859
    if true {
        var v14 string = "\n\t"
        v6 += (v6 >> ((4294967295) & 63))
        _ = v14
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v15 uint16 = 904
    v16 := GenS1{F0: v5, F1: -58, F2: (v7 * v7)}
    var v17 int8 = (v7 - v7)
    switch ((v13 + ((("hello" + "") + v13) + v13)) + "a") {
    case "hello":
        var v18 uint64 = 4294967295
        _ = v18
    case "\n\t":
        fmt.Println(v15, v5, v15)
    }
    if (true && (0 == 182)) {
        v19 := GenS1{F0: v5, F1: -1, F2: 107}
        v20 := GenS2{F0: (v6 % (((v6 + (v6 << ((2) & 63)))) | 1))}
        if (v3 > v3) {
            var v21 uint32 = 2
            for fi22 := 0; fi22 < 20; fi22++ {
                v23 := GenS1{F0: (v5 & (v5 & -32768)), F1: fi22, F2: (v17 % (((v7 / ((int8(fi22)) | 1))) | 1))}
                _ = fi22
                _ = v23
            }
            _ = v21
        }
        _ = v19
        _ = v20
    } else {
        var v24 bool = false
        v5 += (v5 + (-(v5)))
        _ = v24
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        v15 -= v15
        _ = fi25
    }
    fmt.Println(v7, v15, v17)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
