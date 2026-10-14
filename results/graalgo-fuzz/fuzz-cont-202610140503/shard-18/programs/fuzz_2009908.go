package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int16 }

func main() {
    var v1 uint32 = 65535
    var v2 int16 = -32768
    v3 := []uint32{140, (v1 * v1)}
    var v4 uint32 = (v1 + v1)
    var v5 uint = 2
    var v6 int64 = -9223372036854775808
    v1 = v4
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(v1)
        v8 := GenS1{F0: fi7}
        _ = fi7
        _ = v8
    }
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v9 float32 = (-2.25 * (1e10 / -2.25))
        var v10 int32 = 1
        _ = v10
        _ = v9
    }
    if true {
        switch 128 {
        case -624:
            v5 = uint(v4)
        case 1:
            var v11 int = 390
            _ = v11
        }
        for fi12 := 0; fi12 < 16; fi12++ {
            fmt.Println(v2, v2, fi12)
            if ((-108 == -103) || (!true)) {
                for fi13 := 0; fi13 < 2; fi13++ {
                    _ = len(v3)
                    var v14 uint16 = 65535
                    switch "\n\t" {
                    case "foo bar":
                        var v15 uint32 = v1
                        _ = v15
                    case "":
                        if true {
                            if (("\n\t" + "\n\t") >= "a") {
                                var v16 uint32 = 4294967295
                                fmt.Println(v1, v5, v4)
                                _ = v16
                            }
                        } else {
                            v2 = (v2 << (((v5 << ((v5) & 63))) & 63))
                        }
                    default:
                        v4 = v4
                    }
                    _ = fi13
                    _ = v14
                }
                var v17 int16 = v2
                switch (fi12 << ((42) & 63)) {
                case -2147483648:
                    var v18 uint64 = 18446744073709551615
                    _ = v18
                case -171:
                    switch "foo bar" {
                    case "a":
                        v19 := GenS2{F0: 330, F1: 83}
                        _ = v19
                    }
                }
                _ = v17
            } else {
                for fi20 := 0; fi20 < 7; fi20++ {
                    v4 = v4
                    var v21 float32 = 1.5
                    _ = fi20
                    _ = v21
                }
            }
            _ = fi12
        }
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := GenS1{F0: (fi22 ^ fi22)}
        var v24 int16 = (v2 << (((v5 & v5)) & 63))
        _ = fi22
        _ = v23
        _ = v24
    }
    if len(v3) > 0 {
        v3[0] = (v4 + (^(v4)))
    }
    var v25 uint16 = 42
    v4 -= v4
    var v26 int = 9223372036854775807
    if (!true) {
        var v27 uint16 = v25
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
