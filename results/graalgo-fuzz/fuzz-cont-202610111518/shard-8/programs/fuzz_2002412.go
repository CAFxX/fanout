package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint; F1 bool }

func main() {
    var v1 float64 = (-2.25 + -2.25)
    v2 := []uint32{}
    var v3 int32 = -2147483648
    switch "\n\t" {
    case "a":
        if true {
            var v4 uint32 = 1
            v4 += (v4 >> ((65535) & 63))
            _ = v4
        } else {
            v1 = 0.1
        }
    }
    var v5 uint32 = 2147483648
    switch 498 {
    case 1:
        if len(v2) > 0 {
            v2[0] = v5
        }
    }
    var v6 int64 = -99
    fmt.Println(v1, v5)
    if (true || false) {
        var v7 uint16 = 529
        _ = v7
    } else {
        var v8 int = -2147483648
        if (false || true) {
            for fi9 := 0; fi9 < 9; fi9++ {
                if false {
                    if false {
                        v5 += v5
                        var v10 int32 = 691
                        var v11 bool = (float32(v1) > float32((-2.25 + ((v1 + (v1 / v1)) + (3.14159 / 0.1)))))
                        _ = v10
                        _ = v11
                    } else {
                        if true {
                            var v12 bool = true
                            _ = v12
                        } else {
                            var v13 uint16 = 65535
                            _ = v13
                        }
                    }
                    v14 := GenS1{F0: 2}
                    fmt.Println(v3, v1)
                    _ = v14
                } else {
                    var v15 int16 = -32768
                    fmt.Println(v8, v6)
                    _ = v15
                }
                var v16 bool = true
                _ = fi9
                _ = v16
            }
            var v17 int64 = (v6 - (v6 & v6))
            v18 := GenS2{F0: v5, F1: "foo bar"}
            _ = v17
            _ = v18
        }
        _ = len(v2)
        _ = v8
    }
    var v19 int64 = (v6 | (v6 >> ((42) & 63)))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    fmt.Println(v3, v19, v19)
    var v20 uint8 = 142
    _ = len(v2)
    for fi21 := 0; fi21 < 20; fi21++ {
        v1 = -2.25
        var v22 int64 = v19
        _ = fi21
        _ = v22
    }
    v23 := GenS3{F0: 18446744073709551615, F1: true}
    for fi24 := 0; fi24 < 8; fi24++ {
        if (!true) {
            var v25 int16 = -1
            _ = v25
        } else {
            v5 = (v5 ^ 385)
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi24
    }
    var v26 int32 = v3
    var v27 uint64 = 933
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v5
    _ = v6
}
