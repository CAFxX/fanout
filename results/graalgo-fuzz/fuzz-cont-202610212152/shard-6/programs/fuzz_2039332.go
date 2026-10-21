package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint = 2
    var v2 uint16 = 0
    var v3 int8 = 106
    var v4 uint64 = 0
    v5 := make(map[string]string)
    for fi6 := 0; fi6 < 2; fi6++ {
        v3 = 100
        _ = fi6
    }
    switch ("" + "") {
    case "a":
        if ((4294967295 != 128) || (32768 == 804)) {
            v2 += v2
            fmt.Println(v4, v3, v3)
            for fi7 := 0; fi7 < 2; fi7++ {
                fmt.Println(v2)
                v2 = v2
                v8 := GenS1{F0: 2147483648, F1: -567}
                _ = fi7
                _ = v8
            }
        }
    case "hello":
        switch -9223372036854775808 {
        case 4294967295:
            for fi9 := 0; fi9 < 9; fi9++ {
                switch (fi9 ^ (fi9 ^ fi9)) {
                case 0:
                    v3 -= (v3 >> ((v1) & 63))
                case 9223372036854775807:
                    if true {
                        v10 := GenS1{F0: 1, F1: -834}
                        _ = v10
                    }
                }
                if false {
                    var v11 uint64 = 0
                    _ = v11
                }
                var v12 string = "a"
                _ = fi9
                _ = v12
            }
        }
    case "\n\t":
        v13 := GenS2{F0: 42}
        _ = v13
    default:
        v3 = 44
    }
    v14 := []int{-2147483648, 0}
    for fi15 := 0; fi15 < 13; fi15++ {
        v14 = append(v14, fi15)
        fmt.Println(v1, v3, v3)
        switch fi15 {
        case -2147483648:
            if false {
                _ = len(v5)
            } else {
                for fi16 := 0; fi16 < 2; fi16++ {
                    v17 := GenS2{F0: 42}
                    var v18 int64 = 9223372036854775807
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
        case 679:
            var v19 uint64 = 719
            _ = v19
        default:
            v2 += 201
        }
        _ = fi15
    }
    v20 := GenS3{F0: (v2 + (v2 << ((v1) & 63)))}
    var v21 int16 = 1
    v1 += (v1 * 18446744073709551615)
    var v22 uint8 = 1
    switch 9223372036854775807 {
    case 32768:
        v5[""] = "foo bar"
    }
    if len(v14) > 0 {
        v14[0] = 127
    }
    var v23 uint64 = 18446744073709551615
    if true {
        switch -9223372036854775808 {
        case 128:
            var v24 uint8 = 2
            _ = v24
        default:
            v22 = uint8(v23)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
