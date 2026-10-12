package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 int8; F1 uint }

func main() {
    var v1 int32 = 118
    var v2 int8 = 127
    var v3 int64 = -9223372036854775808
    v4 := []uint32{42}
    var v5 uint16 = 32767
    v6 := make(map[uint32]string)
    v7 := GenS1{F0: ("\n\t" + ("" + "foo bar")), F1: "hello", F2: 255}
    v5 = (v5 & (v5 + (^(v5))))
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 uint16 = (^(v5))
            v11 := GenS1{F0: ("hello" + ("hello" + "foo bar")), F1: "foo bar", F2: 4294967295}
            var v12 uint8 = 0
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        var v13 bool = true
        var v14 int8 = v2
        _ = fi8
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        _ = len(v4)
        v1 = 1
        if true {
            v5 = v5
            fmt.Println(v3, v5)
            for fi16 := 0; fi16 < 20; fi16++ {
                v17 := GenS2{F0: 65535}
                switch (fi16 ^ 0) {
                case 0:
                    var v18 uint16 = 1
                    _ = v18
                case 1:
                    switch "\n\t" {
                    case "\n\t":
                        var v19 uint64 = 0
                        _ = v19
                    default:
                        fmt.Println(v3)
                    }
                }
                v20 := GenS2{F0: 42}
                _ = fi16
                _ = v17
                _ = v20
            }
        }
        _ = fi15
    }
    if true {
        var v21 uint = 4294967295
        _ = v21
    } else {
        switch ("foo bar" + ("" + "a")) {
        case "":
            v22 := GenS3{F0: (v2 << ((42) & 63))}
            _ = v22
        }
        _ = len(v4)
    }
    v23 := GenS2{F0: 122}
    v24 := GenS4{F0: 42}
    v1 = (v1 << ((931) & 63))
    fmt.Println(v2)
    v25 := GenS3{F0: -1}
    v3 += v3
    var v26 uint64 = 4294967295
    for fi27 := 0; fi27 < 10; fi27++ {
        fmt.Println(fi27, v1, v26)
        for fi28 := 0; fi28 < 20; fi28++ {
            v3 = (^(v3))
            fmt.Println(fi28, v1)
            if ((v2 % (((v2 ^ (v2 | v2))) | 1)) == (v2 - (v2 & (v2 >> ((255) & 63))))) {
                if (!false) {
                    fmt.Println(v2, fi27)
                } else {
                    fmt.Println(v5, v3, v1)
                    v29 := GenS3{F0: v2}
                    if true {
                        v30 := GenS5{F0: (^(v5)), F1: 0}
                        _ = v30
                    }
                    _ = v29
                }
            }
            _ = fi28
        }
        _ = fi27
    }
    var v31 int8 = -128
    v32 := GenS6{F0: (v2 >> ((42) & 63)), F1: 0}
    v33 := GenS4{F0: (^(v26))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
