package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint64, p1 int16, p2 uint32) bool {
    var v6 int32 = -2147483648
    p2 = 2
    v6 += v6
    var v7 int = 2147483647
    if false {
        v8 := []int{}
        for fi9 := 0; fi9 < 8; fi9++ {
            p1 -= p1
            _ = fi9
        }
        _ = v8
    } else {
        if true {
            for fi10 := 0; fi10 < 18; fi10++ {
                v6 = 2147483647
                var v11 int8 = 1
                _ = fi10
                _ = v11
            }
            switch v7 {
            case -42:
                for fi12 := 0; fi12 < 14; fi12++ {
                    v13 := GenS1{F0: p2, F1: p1}
                    fmt.Println(p2, p0)
                    _ = fi12
                    _ = v13
                }
            case -918:
                if true {
                    var v14 int16 = -1
                    var v15 bool = false
                    _ = v14
                    _ = v15
                }
            default:
                p0 = 2
            }
            var v16 int32 = (v6 * 1)
            _ = v16
        } else {
            var v17 int64 = 42
            v17 = v17
            for fi18 := 0; fi18 < 20; fi18++ {
                p1 = p1
                switch 128 {
                case -2147483648:
                    var v19 uint = 48
                    _ = v19
                }
                _ = fi18
            }
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return false
}
func genFn1(p0 bool, p1 int32, p2 int) int64 {
    switch "foo bar" {
    case "foo bar":
        var v20 uint8 = 28
        _ = v20
    default:
        for fi21 := 0; fi21 < 5; fi21++ {
            p0 = (!(p2 < fi21))
            v22 := []int{p2}
            p2 -= -462
            _ = fi21
            _ = v22
        }
    }
    if (p0 || (42 < 18446744073709551615)) {
        var v23 uint32 = 2147483647
        fmt.Println(p0, p1, p0)
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 bool = true
    var v2 int = 4294967295
    var v3 uint64 = 4294967295
    v4 := make(map[string]bool)
    var v5 uint = 2
    var v24 int = v2
    v25 := GenS1{F0: 0, F1: int16(v3)}
    v26 := v4["a"]
    fmt.Println(v24, v2, v26)
    v5 += v5
    if v26 {
        switch 53 {
        case -9223372036854775808:
            var v27 int8 = 100
            _ = v27
        case 391:
            var v28 int32 = 0
            _ = v28
        }
    } else {
        for fi29 := 0; fi29 < 11; fi29++ {
            var v30 uint32 = 42
            _ = fi29
            _ = v30
        }
    }
    v2 = v24
    var v31 float32 = (1e10 * -2.25)
    switch 2147483647 {
    case 4294967295:
        if (65535 == 255) {
            v32 := GenS2{F0: 1}
            _ = v32
        }
    default:
        for fi33 := 0; fi33 < 18; fi33++ {
            v34 := []int{(fi33 << (((v5 * v5)) & 63)), (v24 ^ (v2 + 0))}
            var v35 bool = (174 == 127)
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    var v36 int8 = 0
    v26 = (!(-1 == 32767))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v36
    _ = v4
    _ = v5
}
