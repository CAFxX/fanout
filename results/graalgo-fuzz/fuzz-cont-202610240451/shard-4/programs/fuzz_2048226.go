package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int16 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 string }
func genFn0(p0 uint16) string {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 17; fi4++ {
        p0 += (p0 % (((^(p0))) | 1))
        if true {
            fmt.Println(p0, p0)
            v5 := GenS1{F0: -2147483648, F1: -9223372036854775808, F2: 42}
            _ = v5
        }
        _ = fi4
    }
    p0 = (^(p0))
    fmt.Println(p0)
    v6 := []int{403, -821, 0}
    _ = p0
    _ = v6
    return "\n\t"
}

func main() {
    var v1 uint64 = 2
    var v2 int = 0
    v3 := make(map[int]string)
    v1 = v1
    v3[v2] = "hello"
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{0}
        v1 += v1
        for fi9 := 0; fi9 < 19; fi9++ {
            switch "hello" {
            case "foo bar":
                if false {
                    v2 = (fi9 << ((182) & 63))
                }
            default:
                var v10 int = (fi7 + fi7)
                _ = v10
            }
            var v11 uint = 65535
            _ = fi9
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    v1 -= 42
    var v12 uint16 = 32767
    var v13 uint16 = 276
    switch "" {
    case "foo bar":
        for fi14 := 0; fi14 < 9; fi14++ {
            if ((true && ((((205 > 127) || false) && (1 > 0)) && false)) || (!true)) {
                var v15 uint32 = 0
                v16 := GenS1{F0: 2147483647, F1: 513, F2: -1}
                for fi17 := 0; fi17 < 7; fi17++ {
                    var v18 uint = 919
                    var v19 int16 = 392
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = v15
                _ = v16
            }
            _ = fi14
        }
    case "":
        fmt.Println(v1)
    case "\n\t":
        for fi20 := 0; fi20 < 10; fi20++ {
            fmt.Println(v13)
            var v21 int = 32768
            var v22 uint = 255
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    if (159 != 2147483647) {
        var v23 bool = false
        v1 = (v1 | (^(v1)))
        _ = v23
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        for fi25 := 0; fi25 < 12; fi25++ {
            var v26 int64 = -71
            _ = fi25
            _ = v26
        }
        var v27 int = (fi24 - (fi24 | (fi24 + int(v1))))
        _ = fi24
        _ = v27
    }
    if false {
        var v28 float32 = (0.1 / -2.25)
        var v29 string = ""
        _ = v28
        _ = v29
    } else {
        var v30 int8 = 0
        _ = v30
    }
    var v31 float32 = 1e10
    if (true && (false && true)) {
        for fi32 := 0; fi32 < 2; fi32++ {
            delete(v3, fi32)
            for fi33 := 0; fi33 < 3; fi33++ {
                if ((!(true && true)) && false) {
                    switch "hello" {
                    case "\n\t":
                        var v34 int64 = -9223372036854775808
                        _ = v34
                    case "":
                        fmt.Println(v1)
                    default:
                        v13 = (v12 / (((v12 % (((v12 * 0)) | 1))) | 1))
                    }
                    v13 -= (v13 ^ v13)
                }
                _ = fi33
            }
            _ = fi32
        }
    }
    v35 := []int{(v2 * 1), v2}
    v36 := GenS2{F0: 42, F1: 0}
    v37 := GenS3{F0: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v35:", len(v35), v35)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v37
}
