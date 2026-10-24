package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
func genFn0(p0 uint64) uint {
    var v5 uint32 = 2147483647
    var v6 float64 = 0.0
    v7 := GenS1{F0: -759, F1: -1}
    var v8 int16 = -32768
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 514
}
func genFn1(p0 int8) uint32 {
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint64 = 2
        v11 := GenS1{F0: 42, F1: -1}
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        switch "hello" {
        case "\n\t":
            for fi13 := 0; fi13 < 8; fi13++ {
                var v14 int32 = 0
                p0 += (p0 ^ (p0 >> ((genFn0(18446744073709551615)) & 63)))
                _ = fi13
                _ = v14
            }
        default:
            fmt.Println(fi12, p0)
        }
        v15 := GenS1{F0: 0, F1: -32768}
        var v16 uint = 42
        _ = fi12
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        p0 = p0
        for fi18 := 0; fi18 < 15; fi18++ {
            fmt.Println(fi18, fi18, fi17)
            _ = fi18
        }
        _ = fi17
    }
    fmt.Println(p0)
    var v19 uint16 = 255
    _ = p0
    _ = v19
    return 65535
}

func main() {
    v1 := []string{"\n\t", "\n\t"}
    v2 := make(map[uint32]string)
    var v3 int64 = 749
    var v4 float64 = ((0.1 * 1e10) - float64(0.0))
    var v20 uint8 = 255
    v2[uint32(v20)] = "hello"
    v2[4294967295] = (("" + "hello") + ((("" + ((("foo bar" + "hello") + (("" + ("" + ("hello" + "a"))) + ("hello" + ("\n\t" + "foo bar")))) + "a")) + "\n\t") + ""))
    var v21 uint16 = 32767
    var v22 int = int(v3)
    for fi23 := 0; fi23 < 5; fi23++ {
        var v24 int = 32767
        if (!(!(1 > 466))) {
            var v25 int32 = 2147483647
            if (!true) {
                v25 -= (v25 ^ v25)
            }
            v3 += (v3 >> ((genFn0(431)) & 63))
            _ = v25
        } else {
            var v26 float32 = (0.1 - (-2.25 + float32((v4 + 0.0))))
            v2[uint32(v20)] = ("\n\t" + (("a" + ("hello" + "foo bar")) + (("\n\t" + "foo bar") + "a")))
            var v27 uint16 = v21
            _ = v26
            _ = v27
        }
        _ = fi23
        _ = v24
    }
    v20 = (^(v20))
    var v28 uint8 = 2
    v28 += v20
    switch "\n\t" {
    case "\n\t":
        v3 -= v3
    case "hello":
        v21 += 1
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v4
}
