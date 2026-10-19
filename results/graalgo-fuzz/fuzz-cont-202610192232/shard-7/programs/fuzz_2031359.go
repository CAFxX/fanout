package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint8; F1 int16 }

func main() {
    var v1 uint16 = 2
    var v2 uint64 = 18446744073709551615
    var v3 uint64 = 18446744073709551615
    v4 := make(map[string]uint32)
    var v5 int16 = 42
    for fi6 := 0; fi6 < 9; fi6++ {
        v3 = v3
        _ = fi6
    }
    switch ("\n\t" + "hello") {
    case "\n\t":
        v7 := []int{}
        _ = v7
    case "":
        fmt.Println(v5, v5)
    }
    if (v5 != int16(v1)) {
        var v8 int8 = 0
        _ = v8
    }
    if true {
        v3 -= (^(v2))
    } else {
        var v9 bool = true
        _ = v9
    }
    v5 -= v5
    if false {
        v5 += 42
        v10 := GenS1{F0: 1, F1: 9223372036854775807}
        for fi11 := 0; fi11 < 15; fi11++ {
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := v4["hello"]
                _ = fi12
                _ = v13
            }
            var v14 int16 = (v5 + v5)
            var v15 bool = ((549 < 4294967295) || true)
            _ = fi11
            _ = v14
            _ = v15
        }
        _ = v10
    } else {
        fmt.Println(v3)
        fmt.Println(v5)
        v1 += 65535
    }
    switch (("\n\t" + "") + ("a" + ("a" + "foo bar"))) {
    case "foo bar":
        if (1 <= -264) {
            for fi16 := 0; fi16 < 16; fi16++ {
                v4[""] = 789
                v17 := []int{(fi16 / ((1) | 1))}
                v17 = append(v17, (fi16 * 138))
                _ = fi16
                _ = v17
            }
        } else {
            var v18 float64 = 0.0
            _ = v18
        }
    }
    var v19 int = 128
    v3 -= 1
    var v20 uint = 463
    fmt.Println(v1, v3)
    fmt.Println(v5)
    if ((204 == 127) || (!((false || true) && (float32(3.14159) >= 3.14159)))) {
        switch ("a" + "hello") {
        case "foo bar":
            for fi21 := 0; fi21 < 4; fi21++ {
                v1 += (v1 | v1)
                _ = fi21
            }
        case "":
            v22 := v4[""]
            _ = v22
        case "a":
            if (true && ((0.0 / 1.0) < 3.14159)) {
                switch (v19 << (((v20 << (((v20 << ((v20) & 63))) & 63))) & 63)) {
                case 9223372036854775807:
                    fmt.Println(v3)
                case -9223372036854775808:
                    var v23 uint64 = 631
                    _ = v23
                case 445:
                    for fi24 := 0; fi24 < 16; fi24++ {
                        v1 += v1
                        _ = fi24
                    }
                default:
                    var v25 float32 = 0.1
                    _ = v25
                }
            } else {
                var v26 int = -265
                _ = v26
            }
        default:
            var v27 uint64 = uint64(v20)
            _ = v27
        }
    }
    v28 := []int{}
    if (false || true) {
        fmt.Println(v2, v2)
        v2 -= v2
        v29 := GenS2{F0: 183, F1: v5}
        _ = v29
    }
    var v30 uint64 = v3
    var v31 int32 = 1
    fmt.Println(v31, v2, v31)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
