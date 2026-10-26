package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int64; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 uint; F2 int16 }
func genFn0(p0 uint) int32 {
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: (("hello" + "\n\t") + ""), F1: 963}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: "", F1: 42}
    var v10 int64 = 42
    _ = p0
    _ = v10
    _ = v9
    return -1
}

func main() {
    var v1 uint32 = 127
    var v2 uint16 = 42
    v3 := make(map[int]int)
    var v4 uint16 = v2
    var v5 uint8 = 255
    v6 := []uint32{}
    for fi11 := 0; fi11 < 14; fi11++ {
        if (v2 == v2) {
            var v12 uint32 = v1
            v13 := GenS1{F0: ("a" + ""), F1: -204}
            _ = v12
            _ = v13
        }
        _ = fi11
    }
    var v14 int32 = -1
    if true {
        for fi15 := 0; fi15 < 4; fi15++ {
            for fi16 := 0; fi16 < 20; fi16++ {
                if true {
                    v4 = 939
                }
                _ = len(v6)
                switch ("" + ("\n\t" + "")) {
                case "\n\t":
                    v4 = 255
                }
                _ = fi16
            }
            var v17 int = fi15
            fmt.Println(fi15, v17, v5)
            _ = fi15
            _ = v17
        }
        v18 := GenS2{F0: -96, F1: 2, F2: -438}
        var v19 uint32 = (v1 + (v1 << ((955) & 63)))
        _ = v18
        _ = v19
    }
    v5 += v5
    _ = len(v6)
    v20 := GenS3{F0: v5}
    v14 += v14
    var v21 int64 = 902
    v22 := GenS4{F0: v14, F1: 2, F2: 32767}
    if (v2 > v4) {
        v23 := GenS1{F0: "foo bar", F1: v21}
        if (985 != 936) {
            v24 := GenS1{F0: "a", F1: (v21 >> ((65535) & 63))}
            _ = len(v3)
            _ = v24
        } else {
            delete(v3, 42)
            var v25 int = -2147483648
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 int16 = 1
                _ = fi26
                _ = v27
            }
            _ = v25
        }
        for fi28 := 0; fi28 < 8; fi28++ {
            var v29 int8 = int8(v21)
            _ = fi28
            _ = v29
        }
        _ = v23
    } else {
        var v30 int32 = int32(v21)
        _ = v30
    }
    _ = len(v6)
    v2 = v4
    v14 -= (v14 << ((42) & 63))
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v14)
    case "":
        switch "hello" {
        case "\n\t":
            v31 := GenS3{F0: v5}
            _ = v31
        }
    case "hello":
        fmt.Println(v2)
    }
    if (0 <= (v14 | (v14 >> ((4294967295) & 63)))) {
        v2 += (^(v2))
    } else {
        switch "a" {
        case "\n\t":
            delete(v3, 32767)
        case "a":
            v14 = int32(v21)
        case "foo bar":
            v32 := GenS4{F0: (v14 / ((-181) | 1)), F1: 18446744073709551615, F2: 6}
            _ = v32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
