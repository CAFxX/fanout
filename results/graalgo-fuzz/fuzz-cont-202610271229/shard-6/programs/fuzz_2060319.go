package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
func genFn0() int {
    if (0 == -1) {
        v7 := []int{-2147483648}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 8; fi8++ {
            fmt.Println(fi8)
            _ = fi8
        }
        switch "hello" {
        case "":
            for fi9 := 0; fi9 < 1; fi9++ {
                switch "foo bar" {
                case "hello":
                    var v10 uint32 = 4294967295
                    _ = v10
                case "foo bar":
                    var v11 int16 = 42
                    _ = v11
                case "a":
                    fmt.Println(fi9)
                }
                _ = fi9
            }
        }
    }
    var v12 float64 = 3.14159
    if false {
        switch -9223372036854775808 {
        case 42:
            fmt.Println(v12)
        case 2147483647:
            var v13 int64 = 48
            _ = v13
        default:
            var v14 int64 = -1
            _ = v14
        }
        v12 += 0.0
    }
    var v15 uint32 = 65535
    if (!false) {
        v16 := []int{1, 32767, 0, -1}
        _ = v16
    } else {
        var v17 uint = 18446744073709551615
        _ = v17
    }
    _ = v12
    _ = v15
    return 32768
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    var v3 int16 = 1000
    v4 := []uint32{2}
    var v5 uint64 = 145
    var v6 uint8 = 2
    switch ("hello" + "a") {
    case "hello":
        _ = len(v4)
    case "\n\t":
        if (!(-1 >= 1)) {
            v18 := GenS1{F0: v3}
            var v19 float32 = 3.14159
            _ = v18
            _ = v19
        } else {
            _ = len(v4)
            var v20 uint8 = (v6 << ((0) & 63))
            fmt.Println(v5)
            _ = v20
        }
    default:
        delete(v2, "")
    }
    if (-546 != 42) {
        if (true && (false || (!(255 < 958)))) {
            var v21 string = "\n\t"
            fmt.Println(v3)
            v5 = v5
            _ = v21
        }
        if false {
            var v22 uint32 = 463
            _ = len(v4)
            v5 += v5
            _ = v22
        } else {
            for fi23 := 0; fi23 < 7; fi23++ {
                v24 := GenS1{F0: v3}
                _ = len(v4)
                var v25 uint8 = v6
                _ = fi23
                _ = v24
                _ = v25
            }
            v26 := GenS1{F0: (v3 >> ((0) & 63))}
            _ = v26
        }
        v5 -= uint64(v6)
    } else {
        v27 := GenS1{F0: v3}
        _ = v27
    }
    v28 := GenS2{F0: -725}
    v29 := GenS2{F0: -1}
    var v30 uint8 = (v6 / ((42) | 1))
    v31 := GenS3{F0: 1}
    v3 = (v3 / ((v3) | 1))
    for fi32 := 0; fi32 < 2; fi32++ {
        switch "" {
        case "a":
            var v33 float32 = (1e10 * (0.1 / 1.5))
            _ = v33
        default:
            v30 -= v30
        }
        var v34 int = (fi32 & -2147483648)
        var v35 bool = false
        _ = fi32
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        fmt.Println(v5, v30, fi36)
        _ = fi36
    }
    var v37 float64 = 1.5
    for fi38 := 0; fi38 < 4; fi38++ {
        var v39 int = fi38
        var v40 uint32 = 2147483647
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println(v6, v6, v6)
    v3 += (v3 << ((2) & 63))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
