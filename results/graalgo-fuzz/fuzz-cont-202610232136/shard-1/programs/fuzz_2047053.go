package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int8, p1 string) int16 {
    p1 = "\n\t"
    if (!(!false)) {
        var v6 int64 = -9223372036854775808
        v6 = v6
        p1 = ("hello" + ("a" + "a"))
        _ = v6
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]string)
    var v3 int32 = 42
    var v4 uint32 = 2147483647
    v5 := make(map[int]int)
    fmt.Println(v4)
    var v7 float32 = v1
    fmt.Println(v1, v4)
    var v8 uint64 = 1
    switch "a" {
    case "hello":
        var v9 uint = 0
        _ = v9
    case "foo bar":
        var v10 int = -9223372036854775808
        _ = v10
    case "a":
        fmt.Println(v8)
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        switch -60 {
        case -1:
            if true {
                v1 = (-2.25 + v1)
                v12 := []int{2147483647, fi11}
                var v13 uint = 0
                _ = v12
                _ = v13
            }
        case -9223372036854775808:
            var v14 uint16 = 255
            _ = v14
        case -389:
            v15 := GenS1{F0: 0, F1: uint16(v4)}
            _ = v15
        default:
            for fi16 := 0; fi16 < 4; fi16++ {
                var v17 uint = 716
                _ = fi16
                _ = v17
            }
        }
        var v18 string = ("" + "a")
        _ = fi11
        _ = v18
    }
    if true {
        var v19 int32 = v3
        var v20 uint8 = 255
        for fi21 := 0; fi21 < 6; fi21++ {
            v22 := v5[fi21]
            _ = fi21
            _ = v22
        }
        _ = v19
        _ = v20
    } else {
        switch "a" {
        case "":
            for fi23 := 0; fi23 < 8; fi23++ {
                for fi24 := 0; fi24 < 12; fi24++ {
                    var v25 float64 = (-2.25 + (3.14159 / 1.0))
                    fmt.Println(v1)
                    v26 := []int{fi23, 128, (fi24 & 0), 0}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                v4 = (v4 + 0)
                var v27 uint8 = 1
                _ = fi23
                _ = v27
            }
        }
    }
    v5[532] = 42
    v28 := []int{919, 32767, 2147483647, 2147483647}
    switch -261 {
    case 971:
        var v29 int16 = 0
        _ = v29
    case 127:
        fmt.Println(v7, v1, v3)
    case 4294967295:
        fmt.Println(v1, v7)
    }
    switch "foo bar" {
    case "\n\t":
        var v30 int16 = genFn0(127, "\n\t")
        _ = v30
    case "hello":
        var v31 uint = 0
        _ = v31
    }
    var v32 int16 = 1000
    if (57 >= 112) {
        for fi33 := 0; fi33 < 1; fi33++ {
            for fi34 := 0; fi34 < 12; fi34++ {
                var v35 float32 = v1
                var v36 bool = (true && ((3.14159 - ((0.1 / v7) + (v35 - v7))) > 1.5))
                fmt.Println(v32)
                _ = fi34
                _ = v35
                _ = v36
            }
            _ = fi33
        }
        if true {
            switch "\n\t" {
            case "":
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
            case "a":
                v37 := GenS1{F0: v8, F1: 230}
                _ = v37
            case "foo bar":
                v38 := GenS1{F0: v8, F1: 42}
                _ = v38
            default:
                for fi39 := 0; fi39 < 12; fi39++ {
                    var v40 float64 = (0.1 * 1.5)
                    _ = fi39
                    _ = v40
                }
            }
            for fi41 := 0; fi41 < 11; fi41++ {
                fmt.Println(v32)
                _ = fi41
            }
        }
        fmt.Println(v1, v7)
    }
    v3 = v3
    v42 := GenS1{F0: (v8 % ((0) | 1)), F1: 1}
    var v43 uint64 = (v8 / ((18446744073709551615) | 1))
    v44 := GenS1{F0: 42, F1: 0}
    delete(v2, (v4 << ((995) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v7
    _ = v8
}
