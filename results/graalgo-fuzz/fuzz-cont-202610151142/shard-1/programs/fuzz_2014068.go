package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
func genFn0(p0 uint) uint16 {
    for fi6 := 0; fi6 < 20; fi6++ {
        switch (fi6 - fi6) {
        case 42:
            if (true || true) {
                fmt.Println(fi6, fi6)
            }
        case -873:
            if true {
                for fi7 := 0; fi7 < 4; fi7++ {
                    switch (fi7 * (fi6 << (((p0 << (((^(p0))) & 63))) & 63))) {
                    case 2147483647:
                        p0 += p0
                    default:
                        p0 += (p0 >> (((p0 * (^(p0)))) & 63))
                    }
                    fmt.Println(p0)
                    v8 := []int{32767, (fi6 >> (((p0 >> ((p0) & 63))) & 63)), 32768, (fi7 - fi6)}
                    _ = fi7
                    _ = v8
                }
                fmt.Println(fi6)
            } else {
                var v9 int = (fi6 ^ fi6)
                v10 := []int{-42, fi6, (fi6 & 9223372036854775807), int(p0)}
                p0 -= 0
                _ = v10
                _ = v9
            }
        }
        fmt.Println(fi6, fi6)
        v11 := GenS1{F0: ("foo bar" + ""), F1: (0.1 < float32(((-2.25 / (0.0 - -2.25)) + 1.5)))}
        _ = fi6
        _ = v11
    }
    switch 583 {
    case 4294967295:
        var v12 float32 = ((1.5 / 1e10) / 3.14159)
        _ = v12
    case 611:
        if true {
            var v13 bool = (!(false && (true && false)))
            var v14 int8 = 100
            _ = v13
            _ = v14
        }
    default:
        var v15 int64 = 275
        _ = v15
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        p0 = p0
        for fi17 := 0; fi17 < 1; fi17++ {
            p0 -= (^(p0))
            v18 := GenS1{F0: "hello", F1: false}
            _ = fi17
            _ = v18
        }
        var v19 bool = true
        _ = fi16
        _ = v19
    }
    _ = p0
    return 65535
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 127
    var v3 uint8 = 2
    var v4 int64 = 9223372036854775807
    var v5 int8 = v2
    v20 := []int{}
    for fi21 := 0; fi21 < 10; fi21++ {
        v1[""] = 1
        v22 := GenS1{F0: "foo bar", F1: false}
        _ = fi21
        _ = v22
    }
    v2 = 100
    v20 = append(v20, 42)
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    if false {
        var v23 float32 = 1e10
        _ = v23
    }
    fmt.Println(v4)
    if false {
        v24 := GenS1{F0: "\n\t", F1: false}
        v3 = (^(v3))
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
