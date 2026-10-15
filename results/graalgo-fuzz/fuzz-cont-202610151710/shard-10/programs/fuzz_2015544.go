package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int8 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint32; F1 uint32 }
func genFn0() uint8 {
    v5 := []int{-1, -2147483648}
    var v6 int16 = 42
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint16 = 2
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 float32 = 0.0
            if (2147483648 <= 65535) {
                fmt.Println(v8)
            } else {
                v5 = append(v5, (fi9 * (^(fi9))))
                var v11 int32 = 0
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    _ = v5
    _ = v6
    return 131
}
func genFn1(p0 uint32, p1 bool) int8 {
    v12 := []int{4294967295, -2147483648, -199, 756}
    v13 := GenS1{F0: (("foo bar" + "a") + ("\n\t" + "foo bar")), F1: 2, F2: 0}
    p0 = 4294967295
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return int8(p0)
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]bool)
    v4 := make(map[string]uint32)
    var v14 int16 = -702
    var v15 uint64 = 0
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 bool = false
        var v18 int = fi16
        delete(v3, (v18 - -1))
        _ = fi16
        _ = v17
        _ = v18
    }
    if (!(!false)) {
        var v19 uint64 = (^(v15))
        v15 -= v19
        v20 := []int{2147483647, 32767}
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 int8 = 38
            if true {
                v23 := GenS2{F0: (fi21 & fi21), F1: ""}
                _ = v23
            }
            _ = fi21
            _ = v22
        }
        for fi24 := 0; fi24 < 15; fi24++ {
            for fi25 := 0; fi25 < 18; fi25++ {
                var v26 uint16 = 685
                _ = fi25
                _ = v26
            }
            _ = fi24
        }
    }
    fmt.Println(v15, v14)
    v15 = v15
    v15 += (v15 / ((v15) | 1))
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        v15 = (v15 & (v15 << ((539) & 63)))
    case "a":
        var v27 uint32 = 2147483647
        _ = v27
    default:
        v28 := GenS3{F0: 65535, F1: 65535}
        _ = v28
    }
    var v29 int8 = 75
    for fi30 := 0; fi30 < 13; fi30++ {
        v29 -= (v29 & (v29 >> ((255) & 63)))
        fmt.Println(v15, v29, v14)
        fmt.Println(v15)
        _ = fi30
    }
    if true {
        var v31 int8 = (v29 ^ v29)
        _ = v31
    } else {
        var v32 string = "foo bar"
        v15 = v15
        _ = v32
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        var v34 uint8 = 255
        _ = fi33
        _ = v34
    }
    var v35 int32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v4
}
