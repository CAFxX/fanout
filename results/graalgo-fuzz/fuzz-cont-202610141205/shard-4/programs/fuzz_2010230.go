package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint }
func genFn0(p0 string, p1 int16, p2 int8) int64 {
    fmt.Println(p1, p0)
    v7 := GenS1{F0: p2, F1: 42, F2: 739}
    v8 := []int{-1, 85, -48}
    _ = len(v8)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -9223372036854775808
}

func main() {
    var v1 uint64 = 277
    v2 := []string{"foo bar", "\n\t", "hello"}
    var v3 uint64 = v1
    var v4 int8 = 124
    var v5 int64 = 0
    v6 := make(map[string]bool)
    v9 := GenS2{F0: 487, F1: 18446744073709551615}
    var v10 int = 42
    v3 = v1
    v10 = 128
    var v11 int32 = -1
    var v12 uint64 = (v1 << ((481) & 63))
    switch (v10 << ((42) & 63)) {
    case 4294967295:
        var v13 uint = 255
        _ = v13
    case 0:
        switch v10 {
        case -832:
            switch (("\n\t" + "") + "\n\t") {
            case "a":
                var v14 uint = 255
                _ = v14
            case "\n\t":
                v4 -= (^(v4))
            }
        default:
            v1 -= v3
        }
    case -2147483648:
        var v15 float64 = 0.1
        _ = v15
    default:
        var v16 uint16 = 0
        _ = v16
    }
    fmt.Println(v4, v11, v3)
    switch (v10 % ((v10) | 1)) {
    case 42:
        for fi17 := 0; fi17 < 6; fi17++ {
            if (true || true) {
                v3 = 964
            }
            _ = fi17
        }
    case -1:
        for fi18 := 0; fi18 < 7; fi18++ {
            switch ("\n\t" + "hello") {
            case "hello":
                fmt.Println(v4, v1)
            case "foo bar":
                if (true && ((^(v11)) <= (^(v11)))) {
                    fmt.Println(v12, v12)
                } else {
                    v19 := v6[("\n\t" + "a")]
                    if (!false) {
                        for fi20 := 0; fi20 < 3; fi20++ {
                            v1 = v3
                            v10 += fi20
                            v5 -= genFn0((("a" + "hello") + "hello"), 702, v4)
                            _ = fi20
                        }
                    }
                    _ = v19
                }
            case "":
                fmt.Println(fi18)
            }
            v3 = (v1 >> ((42) & 63))
            _ = fi18
        }
    case -897:
        v21 := GenS1{F0: (v4 >> ((764) & 63)), F1: 42, F2: (v5 << ((4294967295) & 63))}
        _ = v21
    default:
        v11 += v11
    }
    if (1 == 42) {
        var v22 int8 = v4
        for fi23 := 0; fi23 < 1; fi23++ {
            _ = len(v6)
            for fi24 := 0; fi24 < 11; fi24++ {
                fmt.Println(fi24)
                var v25 bool = false
                delete(v6, "\n\t")
                _ = fi24
                _ = v25
            }
            fmt.Println(v5, v5)
            _ = fi23
        }
        _ = v22
    } else {
        var v26 bool = false
        var v27 float32 = float32(0.1)
        switch (v10 | (v10 << ((42) & 63))) {
        case 1:
            var v28 int = v10
            _ = v28
        }
        _ = v26
        _ = v27
    }
    var v29 int8 = 1
    var v30 uint64 = v3
    var v31 bool = true
    v3 = 400
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
