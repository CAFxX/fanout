package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint8, p1 bool) int64 {
    switch 32768 {
    case 32768:
        var v4 int16 = int16(p0)
        _ = v4
    case -1:
        v5 := []int{0}
        _ = v5
    default:
        if ((p0 - p0) > 47) {
            if p1 {
                fmt.Println(p0, p1)
                if (true && (!(true || false))) {
                    p0 += 2
                    p1 = (4294967295 == uint32(p0))
                    for fi6 := 0; fi6 < 12; fi6++ {
                        v7 := []int{(fi6 | -9223372036854775808)}
                        _ = fi6
                        _ = v7
                    }
                } else {
                    var v8 int32 = 42
                    _ = v8
                }
            }
        } else {
            var v9 uint = 0
            _ = v9
        }
    }
    v10 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = v10
    return 0
}
func genFn1() int64 {
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := []int{4294967295, (-(fi11)), 32767, 42}
        switch "" {
        case "\n\t":
            v13 := GenS2{F0: 2}
            _ = v13
        case "a":
            _ = len(v12)
        case "foo bar":
            _ = len(v12)
        default:
            v14 := GenS1{F0: 88}
            _ = v14
        }
        v15 := GenS1{F0: 42}
        _ = fi11
        _ = v12
        _ = v15
    }
    var v16 uint8 = 0
    switch "foo bar" {
    case "\n\t":
        var v17 uint64 = 404
        _ = v17
    }
    _ = v16
    return 42
}

func main() {
    var v1 int = 2147483647
    v2 := []int{v1, (-(v1)), v1, (-(v1))}
    v3 := []string{("a" + "a"), "\n\t", (("\n\t" + "a") + "foo bar"), ("hello" + "foo bar")}
    switch "a" {
    case "\n\t":
        for fi18 := 0; fi18 < 14; fi18++ {
            for fi19 := 0; fi19 < 1; fi19++ {
                v1 -= v1
                if (false || false) {
                    v20 := GenS1{F0: 1}
                    switch "hello" {
                    case "hello":
                        var v21 int16 = -168
                        _ = v21
                    case "":
                        v22 := GenS1{F0: 42}
                        _ = v22
                    default:
                        v1 -= (fi18 % (((v1 & -2147483648)) | 1))
                    }
                    var v23 uint = 18446744073709551615
                    _ = v20
                    _ = v23
                } else {
                    v1 = fi18
                    fmt.Println(fi18)
                    var v24 float32 = (float32(0.0) - 1e10)
                    _ = v24
                }
                _ = fi19
            }
            var v25 uint32 = 30
            _ = fi18
            _ = v25
        }
    case "foo bar":
        var v26 float32 = 1e10
        _ = v26
    }
    v1 -= v1
    v27 := GenS2{F0: 0}
    v2 = append(v2, (v1 ^ 0))
    v1 = 128
    for fi28 := 0; fi28 < 6; fi28++ {
        var v29 int32 = 42
        v30 := GenS2{F0: 337}
        _ = fi28
        _ = v29
        _ = v30
    }
    v1 += 598
    _ = len(v3)
    switch -167 {
    case -2147483648:
        v1 = (v1 ^ (v1 / (((v1 % ((v1) | 1))) | 1)))
    default:
        v31 := GenS2{F0: 2}
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
}
