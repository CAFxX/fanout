package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 int16 }
func genFn0() uint32 {
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 0, F1: 1, F2: (fi7 / ((fi7) | 1))}
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 int32 = 42
            v11 := []int{(-(fi7)), 853, fi7, -2147483648}
            var v12 uint32 = 517
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        var v13 int32 = -2147483648
        _ = fi7
        _ = v13
        _ = v8
    }
    v14 := []int{127, 243, 2147483647}
    var v15 int64 = -1
    for fi16 := 0; fi16 < 19; fi16++ {
        switch (fi16 % ((fi16) | 1)) {
        case 4294967295:
            switch "foo bar" {
            case "\n\t":
                fmt.Println(fi16, fi16)
            }
        case -9223372036854775808:
            v17 := GenS1{F0: 2, F1: 0, F2: fi16}
            _ = v17
        case 933:
            v15 -= (v15 + v15)
        }
        v18 := GenS2{F0: (293 == 535)}
        _ = fi16
        _ = v18
    }
    _ = v14
    _ = v15
    return 456
}

func main() {
    v1 := []string{"foo bar", ("" + "foo bar"), "hello"}
    v2 := make(map[int]uint32)
    var v3 uint32 = 1
    v4 := make(map[int]int)
    var v5 uint64 = 4294967295
    v6 := []int{79, 1}
    for fi19 := 0; fi19 < 16; fi19++ {
        if (0 == 104) {
            v3 -= 2147483647
            v3 -= v3
            if (1.5 >= 3.14159) {
                v20 := GenS2{F0: (-32768 == 37)}
                _ = v20
            } else {
                var v21 float64 = -2.25
                _ = v21
            }
        }
        _ = fi19
    }
    v3 = (v3 << ((596) & 63))
    switch "hello" {
    case "\n\t":
        fmt.Println(v5, v5)
    }
    var v22 uint16 = 79
    var v23 uint8 = 127
    v23 = v23
    for fi24 := 0; fi24 < 10; fi24++ {
        v23 = (v23 / (((v23 - 255)) | 1))
        var v25 int16 = -561
        _ = len(v6)
        _ = fi24
        _ = v25
    }
    v26 := GenS3{F0: 2147483647, F1: -32768}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
