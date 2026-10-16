package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint16 }

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 42
    v3 := []int{1, 4294967295}
    var v4 uint16 = 1
    fmt.Println(v4)
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: 65535}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 bool = (false && (false && true))
        _ = fi7
        _ = v8
    }
    fmt.Println(v4)
    delete(v1, "a")
    var v9 string = (("\n\t" + "") + "foo bar")
    v9 = "hello"
    v10 := GenS1{F0: 2}
    for fi11 := 0; fi11 < 10; fi11++ {
        delete(v1, v9)
        for fi12 := 0; fi12 < 18; fi12++ {
            switch "\n\t" {
            case "a":
                var v13 uint8 = 255
                _ = v13
            case "foo bar":
                v14 := GenS2{F0: 274, F1: 127}
                _ = v14
            case "\n\t":
                v9 = ""
            default:
                v15 := GenS1{F0: 0}
                _ = v15
            }
            v9 = v9
            _ = fi12
        }
        _ = fi11
    }
    var v16 int32 = 0
    v2 = v2
    v9 = v9
    if (!(190 <= 42)) {
        var v17 int16 = -818
        switch -929 {
        case 97:
            switch 978 {
            case 2147483647:
                v4 += v4
            case 32768:
                v2 = -9223372036854775808
            case -2147483648:
                v17 += -881
            default:
                var v18 uint64 = 315
                _ = v18
            }
        case 9223372036854775807:
            for fi19 := 0; fi19 < 3; fi19++ {
                var v20 int32 = (v16 / (((v16 | 42)) | 1))
                var v21 int32 = int32(v4)
                _ = fi19
                _ = v20
                _ = v21
            }
        default:
            var v22 uint = 27
            _ = v22
        }
        _ = v17
    } else {
        v4 = 1
        var v23 int64 = 9223372036854775807
        if ((v16 * v16) <= v16) {
            v16 += (v16 / (((v16 ^ -536)) | 1))
            var v24 int32 = v16
            _ = v24
        } else {
            for fi25 := 0; fi25 < 5; fi25++ {
                for fi26 := 0; fi26 < 8; fi26++ {
                    v27 := GenS1{F0: 65535}
                    _ = fi26
                    _ = v27
                }
                _ = fi25
            }
            v28 := GenS1{F0: 65535}
            switch (("foo bar" + ("\n\t" + v9)) + v9) {
            case "foo bar":
                v4 = v4
            }
            _ = v28
        }
        _ = v23
    }
    v16 = 0
    v29 := GenS2{F0: 18446744073709551615, F1: 1}
    switch "" {
    case "a":
        v4 = (v4 / (((v4 << ((1) & 63))) | 1))
    case "hello":
        v30 := GenS2{F0: 123, F1: 6}
        _ = v30
    case "":
        if false {
            v31 := GenS3{F0: v2, F1: 816}
            _ = v31
        } else {
            v2 = (v2 / ((v2) | 1))
            v9 = v9
            v16 -= v16
        }
    }
    switch "" {
    case "hello":
        v3 = append(v3, -1)
    default:
        fmt.Println(v2, v16, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v9
}
