package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint, p1 int16, p2 int8) uint64 {
    var v5 uint = (p0 | p0)
    switch 1 {
    case 1:
        v6 := []int{}
        _ = v6
    case 128:
        if true {
            switch -9223372036854775808 {
            case 0:
                if true {
                    var v7 int8 = (p2 << ((36) & 63))
                    if (true || true) {
                        fmt.Println(p0, p0, p2)
                    }
                    _ = v7
                } else {
                    var v8 int16 = p1
                    _ = v8
                }
            }
            if (true || (0.1 < float32(3.14159))) {
                switch "hello" {
                case "\n\t":
                    for fi9 := 0; fi9 < 14; fi9++ {
                        for fi10 := 0; fi10 < 13; fi10++ {
                            if (true && (65535 <= p0)) {
                                v5 += 451
                                var v11 int = (fi10 | (fi9 << ((v5) & 63)))
                                _ = v11
                            }
                            p1 = (p1 + (p1 | (p1 + p1)))
                            p2 -= (p2 * (p2 & 127))
                            var v12 uint32 = 2147483648
                            _ = fi10
                            _ = v12
                        }
                        p0 = (p0 - 639)
                        var v13 string = "foo bar"
                        _ = fi9
                        _ = v13
                    }
                case "foo bar":
                    var v14 int64 = -1
                    _ = v14
                }
                v15 := []int{}
                _ = v15
            } else {
                if ((true && (!(!true))) || ((p0 - v5) >= v5)) {
                    p1 = 1000
                } else {
                    if false {
                        fmt.Println(p0, v5)
                    }
                    fmt.Println(v5, v5)
                }
                p2 -= 57
            }
            v16 := []int{1, 9223372036854775807, -821}
            _ = v16
        }
    }
    v17 := []int{}
    switch -612 {
    case 1:
        var v18 uint16 = 255
        _ = v18
    case 2147483647:
        fmt.Println(p2)
    case 418:
        fmt.Println(p0, p2, p2)
    default:
        switch 32767 {
        case 32767:
            fmt.Println(p1, p2, p2)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v5
    return 0
}
func genFn1() int64 {
    v19 := GenS1{F0: 365}
    var v20 float64 = 0.0
    v21 := GenS1{F0: 67}
    _ = v19
    _ = v20
    _ = v21
    return -696
}

func main() {
    v1 := []uint32{0, 455, 2}
    var v2 bool = false
    v3 := make(map[string]int)
    var v4 uint8 = 1
    var v22 int64 = -9223372036854775808
    fmt.Println(v2)
    v23 := GenS1{F0: genFn0(65535, 393, -128)}
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 int64 = v22
        _ = fi24
        _ = v25
    }
    var v26 int16 = -491
    if v2 {
        v26 = v26
        v2 = (-1 > -1)
    }
    v27 := GenS1{F0: 42}
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    var v28 int64 = (v22 & 1)
    v1 = append(v1, 594)
    v22 = (v22 / (((v28 + v28)) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
