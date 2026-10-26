package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint; F2 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 uint8; F2 bool }
func genFn0() uint64 {
    var v6 int = 933
    var v7 int32 = 2147483647
    var v8 uint8 = 208
    _ = v6
    _ = v7
    _ = v8
    return 228
}

func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]uint32)
    v3 := []int{2147483647, -2147483648}
    v4 := make(map[string]uint32)
    var v5 uint8 = 255
    v9 := GenS1{F0: (v5 << ((4294967295) & 63))}
    v10 := GenS1{F0: v5}
    var v11 int64 = -9223372036854775808
    switch (("\n\t" + (("" + ("a" + "hello")) + "a")) + ("foo bar" + "\n\t")) {
    case "hello":
        v12 := GenS1{F0: (v5 | (v5 - v5))}
        _ = v12
    case "":
        var v13 uint8 = v5
        _ = v13
    default:
        var v14 uint16 = 255
        _ = v14
    }
    v15 := GenS2{F0: 7, F1: v1, F2: 452}
    v11 = int64(v1)
    v16 := GenS1{F0: 128}
    if (0 < 4294967295) {
        fmt.Println(v1)
    } else {
        v17 := GenS3{F0: -1, F1: 2, F2: 42}
        if (false && true) {
            if true {
                switch 213 {
                case 32767:
                    switch "a" {
                    case "":
                        var v18 int64 = 1
                        _ = v18
                    default:
                        _ = len(v2)
                    }
                case 633:
                    v11 += -1
                }
                delete(v4, "hello")
                fmt.Println(v5)
            }
        }
        v19 := GenS2{F0: genFn0(), F1: v1, F2: 4294967295}
        _ = v17
        _ = v19
    }
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 bool = (int16(fi20) < 1)
            var v22 uint16 = 2
            _ = fi20
            _ = v21
            _ = v22
        }
        v23 := v4["foo bar"]
        _ = v23
    }
    delete(v4, ("hello" + "a"))
    v11 += -1
    for fi24 := 0; fi24 < 9; fi24++ {
        var v25 bool = ((v5 | v5) < v5)
        switch (fi24 | fi24) {
        case 354:
            var v26 uint16 = 2
            _ = v26
        case -898:
            v11 -= (v11 & v11)
        case -1:
            v1 -= v1
        }
        _ = len(v3)
        _ = fi24
        _ = v25
    }
    if (-174 != 32767) {
        v11 += v11
        if false {
            v27 := GenS4{F0: -1}
            _ = v27
        }
    }
    if (v5 == (v5 | 127)) {
        v28 := GenS5{F0: 1, F1: (v5 % (((^(v5))) | 1)), F2: (true && true)}
        v11 = 285
        fmt.Println(v11, v5, v11)
        _ = v28
    } else {
        var v29 uint32 = 2147483647
        fmt.Println(v11, v5)
        _ = v29
    }
    v11 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
