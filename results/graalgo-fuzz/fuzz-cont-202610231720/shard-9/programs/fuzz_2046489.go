package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int64 }
type GenS2 struct { F0 string; F1 uint16 }
func genFn0() bool {
    var v5 string = "hello"
    v5 = ""
    var v6 int32 = -1
    v7 := GenS1{F0: 255, F1: 1, F2: -9223372036854775808}
    _ = v5
    _ = v6
    _ = v7
    return (42 <= 1)
}

func main() {
    var v1 int32 = -767
    v2 := []int{-9223372036854775808}
    v3 := []uint32{0, 0}
    var v4 int = 32767
    fmt.Println(v4, v4)
    v8 := GenS1{F0: 38, F1: -32768, F2: 1}
    switch "foo bar" {
    case "":
        if len(v3) > 0 {
            v3[0] = 2
        }
    case "\n\t":
        v9 := GenS2{F0: (("foo bar" + "foo bar") + ((("hello" + (("a" + "hello") + "hello")) + "foo bar") + "a")), F1: 255}
        _ = v9
    default:
        _ = len(v3)
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        if (42 <= 255) {
            for fi11 := 0; fi11 < 9; fi11++ {
                fmt.Println(fi11, fi10)
                _ = fi11
            }
        } else {
            if len(v2) > 0 {
                v2[0] = v4
            }
            v2 = append(v2, (fi10 ^ 292))
        }
        for fi12 := 0; fi12 < 2; fi12++ {
            if true {
                v13 := GenS2{F0: "hello", F1: 255}
                _ = v13
            } else {
                v1 += v1
                var v14 uint8 = 127
                _ = v14
            }
            fmt.Println(fi10, v4, fi12)
            v4 -= v4
            v15 := GenS1{F0: 2, F1: 1000, F2: 42}
            _ = fi12
            _ = v15
        }
        if len(v2) > 0 {
            v2[0] = -1
        }
        _ = fi10
    }
    v4 -= v4
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 int8 = 100
        if true {
            v4 += v4
            fmt.Println(v4, v17, v4)
            switch "a" {
            case "a":
                fmt.Println(v4, v1)
            }
        } else {
            fmt.Println(v1)
            for fi18 := 0; fi18 < 11; fi18++ {
                var v19 float64 = (1e10 + (0.0 + float64(0.1)))
                v17 = 42
                _ = fi18
                _ = v19
            }
            var v20 uint64 = 649
            _ = v20
        }
        v4 = (fi16 | v4)
        _ = fi16
        _ = v17
    }
    if true {
        fmt.Println(v1, v4)
        var v21 int8 = 127
        _ = v21
    } else {
        v4 += v4
        if len(v2) > 0 {
            v2[0] = (v4 ^ 4294967295)
        }
    }
    if true {
        var v22 int8 = 82
        _ = v22
    } else {
        if true {
            var v23 int16 = 32767
            _ = v23
        } else {
            var v24 float64 = 1.5
            var v25 int = (v4 << ((65535) & 63))
            _ = v24
            _ = v25
        }
        v1 -= v1
    }
    if (-1 >= 42) {
        for fi26 := 0; fi26 < 4; fi26++ {
            var v27 float32 = 1.5
            _ = fi26
            _ = v27
        }
    }
    var v28 uint16 = 358
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v8
}
