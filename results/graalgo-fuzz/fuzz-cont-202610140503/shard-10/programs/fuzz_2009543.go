package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int }

func main() {
    v1 := []uint32{65535}
    var v2 uint16 = 1
    var v3 int32 = 42
    v4 := make(map[int]bool)
    if (!(!(-2.25 != (3.14159 * 0.0)))) {
        v2 += (v2 | (v2 ^ (^(v2))))
        _ = len(v1)
    }
    switch 4294967295 {
    case 10:
        v2 = (v2 & v2)
    case -9223372036854775808:
        if true {
            for fi5 := 0; fi5 < 14; fi5++ {
                if true {
                    switch ("" + ("" + "\n\t")) {
                    case "hello":
                        var v6 int32 = (-(v3))
                        _ = v6
                    case "\n\t":
                        for fi7 := 0; fi7 < 20; fi7++ {
                            v2 -= (v2 << ((4294967295) & 63))
                            _ = fi7
                        }
                    }
                } else {
                    var v8 float64 = 1.5
                    switch ((("a" + ("" + ("\n\t" + "a"))) + "foo bar") + "hello") {
                    case "hello":
                        var v9 uint16 = v2
                        _ = v9
                    case "\n\t":
                        fmt.Println(v8, v2, fi5)
                    case "foo bar":
                        var v10 int16 = -223
                        _ = v10
                    }
                    v2 = 32767
                    _ = v8
                }
                v3 = v3
                _ = fi5
            }
        }
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        if false {
            var v12 float32 = (((0.0 + 1e10) * float32(-2.25)) + 1e10)
            _ = v12
        } else {
            fmt.Println(fi11, fi11)
            v2 = (^(v2))
        }
        var v13 uint16 = 32767
        _ = fi11
        _ = v13
    }
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        var v14 float64 = 0.0
        _ = v14
    case "hello":
        if (-1 != -84) {
            v3 = v3
        } else {
            v3 += 42
            if (!false) {
                v3 = (v3 ^ v3)
            } else {
                var v15 int8 = 127
                _ = v15
            }
        }
    case "":
        v2 = 65535
    }
    v16 := GenS1{F0: 2, F1: v2, F2: 9223372036854775807}
    v2 += (v2 << ((1) & 63))
    var v17 bool = (true && true)
    v18 := GenS1{F0: 0, F1: (v2 | v2), F2: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
