package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint32; F1 bool; F2 bool }
func genFn0(p0 int8) uint8 {
    p0 = p0
    p0 = p0
    switch "" {
    case "":
        switch "" {
        case "hello":
            v7 := []int{32768, 1, -9223372036854775808, -955}
            _ = v7
        case "foo bar":
            p0 -= p0
        case "a":
            var v8 int8 = (p0 + p0)
            _ = v8
        default:
            switch 691 {
            case 847:
                var v9 float64 = float64(3.14159)
                _ = v9
            case 32767:
                v10 := GenS1{F0: -9223372036854775808}
                _ = v10
            default:
                for fi11 := 0; fi11 < 10; fi11++ {
                    p0 -= (p0 + p0)
                    if true {
                        var v12 float32 = 1.5
                        var v13 int32 = 2147483647
                        v14 := GenS1{F0: -615}
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    _ = fi11
                }
            }
        }
    case "\n\t":
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := GenS1{F0: fi15}
            _ = fi15
            _ = v16
        }
    default:
        switch ("" + "hello") {
        case "foo bar":
            v17 := []int{1}
            _ = v17
        case "\n\t":
            v18 := GenS1{F0: 2147483647}
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 uint = 2
        v20 = 65535
        _ = fi19
        _ = v20
    }
    if (true || true) {
        p0 -= (-(p0))
    }
    _ = p0
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    v3 := []string{"\n\t"}
    var v4 int64 = 1
    var v5 float64 = -2.25
    v6 := make(map[int]int)
    switch 252 {
    case 212:
        var v21 int16 = -1
        _ = v21
    default:
        if (false && (!false)) {
            v4 = (v4 + v4)
            switch -1 {
            case -281:
                _ = len(v3)
            case 42:
                v22 := GenS2{F0: ("hello" + "a"), F1: ((214 > 120) || true)}
                _ = v22
            case 0:
                var v23 uint8 = 45
                _ = v23
            }
        } else {
            switch "" {
            case "a":
                fmt.Println(v4, v4)
            case "":
                v24 := GenS1{F0: 32768}
                _ = v24
            case "\n\t":
                v4 = v4
            }
            var v25 int = 0
            _ = v25
        }
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    var v26 uint16 = 2
    var v27 float64 = 1e10
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v28 int32 = 42
    v29 := GenS2{F0: (("a" + ("\n\t" + "foo bar")) + ("hello" + "hello")), F1: true}
    v30 := GenS3{F0: 2147483648, F1: false, F2: ((0 != 32767) || true)}
    v26 += (v26 / (((v26 * v26)) | 1))
    var v31 string = ("\n\t" + "\n\t")
    switch 1 {
    case 42:
        switch 4294967295 {
        case 933:
            v27 -= v5
        }
    case 4294967295:
        v26 += v26
    case 2147483647:
        fmt.Println(v31, v31, v31)
    }
    fmt.Println(v27, v26, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
