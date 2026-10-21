package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int16 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 int8 = 127
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    var v4 int8 = (^(v1))
    v5 := []uint32{514, 2}
    v6 := make(map[string]string)
    var v7 int8 = (v1 & 87)
    for fi8 := 0; fi8 < 11; fi8++ {
        v5 = append(v5, 1)
        var v9 uint = 18446744073709551615
        _ = fi8
        _ = v9
    }
    var v10 uint16 = 32767
    if (-2147483648 != 505) {
        if true {
            if true {
                v11 := GenS1{F0: 848, F1: 644, F2: -1}
                _ = v11
            } else {
                var v12 int64 = 9223372036854775807
                _ = v12
            }
        }
        v1 = -15
        switch 4294967295 {
        case 9223372036854775807:
            if (-1 > -59) {
                _ = len(v2)
                v5 = append(v5, 42)
            } else {
                v4 -= -1
                for fi13 := 0; fi13 < 14; fi13++ {
                    switch ("hello" + ("" + (("" + ("\n\t" + "\n\t")) + ("foo bar" + ("\n\t" + ("" + "hello")))))) {
                    case "hello":
                        v14 := GenS2{F0: ((v1 + v7) < v4), F1: 128}
                        _ = v14
                    case "\n\t":
                        v7 = 100
                    case "":
                        v10 += v10
                    default:
                        v15 := GenS3{F0: 127, F1: "\n\t"}
                        _ = v15
                    }
                    v1 = (^(v7))
                    _ = fi13
                }
                v16 := GenS4{F0: 4294967295}
                _ = v16
            }
        default:
            v1 = v4
        }
    } else {
        switch 2147483647 {
        case 32768:
            fmt.Println(v7, v7)
        case -198:
            var v17 uint = 4294967295
            _ = v17
        }
        fmt.Println(v1)
    }
    var v18 string = ("a" + "a")
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 uint32 = 65535
        var v21 string = v18
        _ = fi19
        _ = v20
        _ = v21
    }
    switch v18 {
    case "\n\t":
        switch -1 {
        case 263:
            var v22 uint8 = 42
            _ = v22
        case 42:
            v7 = (v1 * (v1 % ((v1) | 1)))
        default:
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 bool = true
                _ = fi23
                _ = v24
            }
        }
    case "a":
        if len(v5) > 0 {
            v5[0] = 0
        }
    }
    switch -9223372036854775808 {
    case 127:
        var v25 uint8 = 0
        _ = v25
    case -942:
        var v26 int32 = 699
        _ = v26
    case 42:
        v4 += (v4 & v1)
    default:
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 uint64 = 42
            if len(v5) > 0 {
                v5[0] = 472
            }
            v29 := GenS1{F0: 1, F1: uint64(v10), F2: -1}
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    if false {
        switch v18 {
        case "\n\t":
            var v30 float32 = 0.0
            _ = v30
        default:
            switch v18 {
            case "\n\t":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                v31 := v6[(v18 + "")]
                _ = v31
            }
        }
        if len(v5) > 0 {
            v5[0] = 65535
        }
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        var v33 int64 = 0
        fmt.Println(v10, v1)
        var v34 int8 = v4
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 bool = true
    var v36 uint = 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
