package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS3 struct { F0 int32; F1 int8; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint; F2 uint8 }
func genFn0(p0 int64, p1 int32, p2 uint16) uint8 {
    fmt.Println(p2, p2, p0)
    var v4 uint = 18446744073709551615
    v5 := GenS1{F0: p2, F1: (^(v4)), F2: 42}
    if true {
        for fi6 := 0; fi6 < 17; fi6++ {
            fmt.Println(p1, p1)
            var v7 bool = (false || true)
            _ = fi6
            _ = v7
        }
        v4 -= 255
        switch 9223372036854775807 {
        case -676:
            var v8 int16 = 136
            _ = v8
        case 9223372036854775807:
            fmt.Println(p2, p1, p0)
        default:
            fmt.Println(v4, p1, p2)
        }
    } else {
        fmt.Println(p2, p0, p2)
        var v9 int32 = p1
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 128
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 2
    var v3 bool = false
    v2 -= v2
    fmt.Println(v3, v3)
    _ = len(v1)
    if (false || (42 > 109)) {
        switch "" {
        case "foo bar":
            _ = len(v1)
        case "hello":
            v3 = v3
        case "\n\t":
            if v3 {
                switch "" {
                case "hello":
                    var v10 int64 = 0
                    _ = v10
                default:
                    v11 := GenS2{F0: 2147483648, F1: 42, F2: 65535}
                    _ = v11
                }
                var v12 uint64 = 1
                var v13 bool = ("a" != "")
                _ = v12
                _ = v13
            }
        }
    }
    if false {
        if true {
            var v14 int8 = 42
            _ = v14
        }
    }
    var v15 float64 = -2.25
    switch 9223372036854775807 {
    case 60:
        if ((!false) && v3) {
            for fi16 := 0; fi16 < 1; fi16++ {
                for fi17 := 0; fi17 < 6; fi17++ {
                    v18 := v1[32767]
                    fmt.Println(v3)
                    switch (fi16 + 32768) {
                    case -253:
                        var v19 string = "hello"
                        _ = v19
                    case 128:
                        var v20 uint16 = v2
                        _ = v20
                    case 0:
                        v18 = true
                    }
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
        } else {
            v3 = (42 == 2147483647)
            for fi21 := 0; fi21 < 10; fi21++ {
                var v22 int64 = -9223372036854775808
                _ = fi21
                _ = v22
            }
        }
    }
    v23 := GenS3{F0: -48, F1: 100, F2: v2}
    v24 := GenS1{F0: (v2 << ((997) & 63)), F1: 255, F2: 4294967295}
    v2 += v2
    for fi25 := 0; fi25 < 2; fi25++ {
        fmt.Println(v15, fi25)
        v15 -= (v15 / ((v15 * ((1.5 + (1e10 - 1.5)) * float64(1.5))) * v15))
        _ = fi25
    }
    v26 := GenS4{F0: (v2 >= 990), F1: 327, F2: 160}
    v27 := []int{127, -176, 42, -9223372036854775808}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v3
}
