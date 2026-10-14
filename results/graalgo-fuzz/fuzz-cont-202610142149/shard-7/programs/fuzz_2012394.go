package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int64; F1 int16 }
func genFn0(p0 uint32, p1 int32, p2 uint32) bool {
    for fi5 := 0; fi5 < 3; fi5++ {
        p0 -= p2
        p0 = (p2 ^ (p2 << ((1) & 63)))
        v6 := []int{429, 128, (fi5 * (fi5 | 128))}
        _ = fi5
        _ = v6
    }
    fmt.Println(p0, p1, p1)
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return false
}
func genFn1(p0 uint, p1 bool, p2 uint) uint16 {
    if p1 {
        fmt.Println(p2, p0)
        p2 -= (p0 - 2)
    }
    var v7 string = "foo bar"
    if false {
        var v8 string = v7
        var v9 string = v7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    var v3 bool = (1e10 >= (3.14159 * 0.1))
    var v4 bool = v3
    fmt.Println(v3, v3)
    v3 = v3
    switch "hello" {
    case "hello":
        v10 := GenS1{F0: 255, F1: 32768}
        _ = v10
    case "foo bar":
        var v11 int16 = -975
        _ = v11
    default:
        switch -9223372036854775808 {
        case 127:
            var v12 uint32 = 65535
            _ = v12
        case 128:
            switch 127 {
            case -1:
                var v13 int32 = 1
                _ = v13
            }
        case 0:
            var v14 uint64 = 0
            _ = v14
        default:
            switch 4294967295 {
            case 32767:
                v1[-540] = 364
            }
        }
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 uint = 61
        switch "foo bar" {
        case "hello":
            delete(v1, fi15)
        case "foo bar":
            v17 := GenS1{F0: (^(v16)), F1: -9223372036854775808}
            _ = v17
        case "":
            if v4 {
                var v18 float32 = -2.25
                _ = v18
            } else {
                var v19 uint8 = 2
                v16 = 0
                _ = v19
            }
        default:
            v20 := GenS2{F0: (v16 | v16), F1: -481}
            _ = v20
        }
        fmt.Println(v3, fi15)
        _ = fi15
        _ = v16
    }
    if v4 {
        var v21 uint = 1
        if (!v3) {
            fmt.Println(v3)
            var v22 string = "a"
            var v23 int8 = 127
            _ = v22
            _ = v23
        }
        v21 += (v21 ^ v21)
        _ = v21
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        switch "foo bar" {
        case "":
            if (-372 <= -714) {
                var v25 bool = v3
                _ = v25
            } else {
                var v26 float32 = float32(((1e10 + 0.1) / 3.14159))
                var v27 int8 = 100
                var v28 int8 = v27
                _ = v26
                _ = v27
                _ = v28
            }
        case "hello":
            var v29 int = 128
            _ = v29
        case "a":
            switch ((("\n\t" + (("a" + "foo bar") + ("" + "hello"))) + ("foo bar" + "a")) + "hello") {
            case "\n\t":
                v2["\n\t"] = fi24
            case "":
                v30 := []int{755, fi24, -441}
                _ = v30
            case "a":
                v4 = true
            default:
                if v3 {
                    v3 = (2147483648 == 4294967295)
                    v31 := GenS3{F0: 242, F1: -1}
                    _ = v31
                }
            }
        default:
            for fi32 := 0; fi32 < 5; fi32++ {
                fmt.Println(fi24, fi24, v4)
                v33 := []int{(fi32 * (fi24 | fi32)), -591}
                v34 := v2["foo bar"]
                _ = fi32
                _ = v33
                _ = v34
            }
        }
        _ = fi24
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        fmt.Println(fi35)
        v2[""] = 32767
        _ = fi35
    }
    for fi36 := 0; fi36 < 10; fi36++ {
        var v37 string = "\n\t"
        var v38 uint64 = 18446744073709551615
        fmt.Println(v38, v37)
        _ = fi36
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 9; fi39++ {
        for fi40 := 0; fi40 < 3; fi40++ {
            var v41 uint16 = genFn1(255, (("a" + "\n\t") > "foo bar"), 891)
            var v42 uint8 = 1
            _ = fi40
            _ = v41
            _ = v42
        }
        var v43 bool = ((734 >= 53) && v3)
        v44 := GenS1{F0: 0, F1: -2147483648}
        _ = fi39
        _ = v43
        _ = v44
    }
    v45 := GenS2{F0: 4294967295, F1: -9223372036854775808}
    var v46 uint32 = 192
    fmt.Println(v4)
    v3 = v4
    v46 -= v46
    v47 := []int{853}
    fmt.Println(v46)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v47:", len(v47), v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
}
