package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint64 }
type GenS2 struct { F0 int; F1 int32; F2 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 int8; F2 uint64 }
func genFn0(p0 int64) uint64 {
    if false {
        var v7 uint32 = 2
        v7 = v7
        _ = v7
    } else {
        p0 -= (p0 / (((p0 ^ p0)) | 1))
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        switch -996 {
        case 9223372036854775807:
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := GenS1{F0: 90, F1: 0, F2: 4294967295}
                _ = fi9
                _ = v10
            }
        case -2147483648:
            p0 += p0
        default:
            for fi11 := 0; fi11 < 5; fi11++ {
                v12 := []int{(fi11 & fi8), (fi8 / ((fi11) | 1)), 392}
                var v13 int8 = 127
                switch ("hello" + "hello") {
                case "":
                    v12 = append(v12, (-(fi8)))
                case "\n\t":
                    switch fi8 {
                    case -92:
                        var v14 int8 = 100
                        _ = v14
                    }
                case "foo bar":
                    v13 = (v13 & (v13 << ((18446744073709551615) & 63)))
                }
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        v15 := GenS1{F0: 0, F1: (-(fi8)), F2: 4294967295}
        var v16 uint32 = 2
        _ = fi8
        _ = v15
        _ = v16
    }
    var v17 float32 = (1.5 + (float32(((3.14159 / -2.25) / 1.5)) / 3.14159))
    _ = p0
    _ = v17
    return 4294967295
}
func genFn1() uint16 {
    var v18 string = ""
    v18 = (("hello" + ("foo bar" + v18)) + ("a" + "a"))
    fmt.Println(v18)
    v18 = v18
    fmt.Println(v18)
    _ = v18
    return 0
}

func main() {
    var v1 uint = 976
    var v2 uint64 = 2
    var v3 int16 = -1
    var v4 int32 = -1
    v5 := make(map[int]bool)
    v6 := make(map[string]bool)
    var v19 int32 = (v4 | v4)
    for fi20 := 0; fi20 < 11; fi20++ {
        v21 := GenS1{F0: 127, F1: 9223372036854775807, F2: 58}
        _ = fi20
        _ = v21
    }
    v22 := []int{42, 2147483647, 32768}
    v2 = v2
    var v23 uint32 = 1
    for fi24 := 0; fi24 < 19; fi24++ {
        v2 = 2
        for fi25 := 0; fi25 < 4; fi25++ {
            if true {
                fmt.Println(v23)
                _ = len(v22)
                v26 := v5[-1]
                _ = v26
            } else {
                delete(v5, -2147483648)
            }
            _ = fi25
        }
        _ = fi24
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        var v28 int16 = (v3 >> ((uint(v23)) & 63))
        v29 := GenS2{F0: fi27, F1: v19, F2: genFn0(-369)}
        v4 -= (v19 | 42)
        _ = fi27
        _ = v28
        _ = v29
    }
    v23 += 2
    switch "foo bar" {
    case "hello":
        for fi30 := 0; fi30 < 17; fi30++ {
            v31 := GenS3{F0: (("foo bar" + "hello") + "")}
            _ = fi30
            _ = v31
        }
    case "a":
        v23 -= v23
    case "":
        for fi32 := 0; fi32 < 4; fi32++ {
            switch "hello" {
            case "":
                for fi33 := 0; fi33 < 4; fi33++ {
                    v34 := GenS4{F0: v3, F1: 96, F2: v2}
                    fmt.Println(v1)
                    _ = fi33
                    _ = v34
                }
            case "a":
                v35 := GenS1{F0: -78, F1: (-(fi32)), F2: 42}
                _ = v35
            }
            fmt.Println(v3, v4)
            _ = fi32
        }
    }
    fmt.Println(v23, v3)
    var v36 uint8 = 42
    for fi37 := 0; fi37 < 11; fi37++ {
        v36 = (v36 | 0)
        v3 = v3
        _ = fi37
    }
    fmt.Println(v23, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
