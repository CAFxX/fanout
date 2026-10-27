package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 bool; F1 int8; F2 int8 }
func genFn0(p0 int16) string {
    if false {
        if (-2 <= -60) {
            v5 := []int{4294967295, 1, 9223372036854775807, -9223372036854775808}
            _ = v5
        }
        if false {
            for fi6 := 0; fi6 < 4; fi6++ {
                p0 = p0
                switch "hello" {
                case "hello":
                    var v7 int32 = 2147483647
                    _ = v7
                default:
                    for fi8 := 0; fi8 < 10; fi8++ {
                        var v9 float32 = 0.0
                        v10 := GenS1{F0: 0, F1: -78, F2: 1}
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                }
                v11 := []int{fi6, fi6, (fi6 / (((fi6 << ((268) & 63))) | 1))}
                _ = fi6
                _ = v11
            }
        } else {
            p0 = (p0 & p0)
            v12 := []int{336, -1}
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 uint8 = 255
        var v15 uint32 = uint32(v14)
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        p0 += (p0 / ((42) | 1))
        p0 += p0
        v17 := []int{211}
        _ = fi16
        _ = v17
    }
    switch "\n\t" {
    case "\n\t":
        p0 = (p0 - p0)
    case "foo bar":
        var v18 float64 = ((3.14159 + (((3.14159 / (0.1 / 0.1)) + 0.0) / 3.14159)) * (float64(1.5) + 0.1))
        _ = v18
    case "a":
        v19 := GenS2{F0: -2147483648, F1: ((160 <= 0) || false)}
        _ = v19
    default:
        var v20 int16 = -32768
        _ = v20
    }
    v21 := GenS3{F0: false, F1: -124, F2: 1}
    _ = p0
    _ = v21
    return ""
}

func main() {
    var v1 string = ""
    v2 := []string{"\n\t", "a"}
    var v3 float32 = 1e10
    v4 := make(map[string]int)
    var v22 int32 = 1
    if len(v2) > 0 {
        v2[0] = "a"
    }
    if true {
        var v23 uint32 = 1
        var v24 int = 0
        _ = v23
        _ = v24
    } else {
        if false {
            var v25 bool = true
            var v26 uint = 4294967295
            _ = v25
            _ = v26
        } else {
            v27 := GenS3{F0: (!true), F1: 127, F2: 0}
            v1 = (("foo bar" + "\n\t") + v1)
            _ = v27
        }
        switch 42 {
        case 32767:
            fmt.Println(v3, v22, v1)
        case 42:
            v1 = "\n\t"
        case 127:
            v28 := GenS2{F0: -1, F1: (-2147483648 > 0)}
            _ = v28
        }
    }
    if (!false) {
        if (!(float32(1e10) > 1e10)) {
            if ((v22 << ((4294967295) & 63)) == -313) {
                var v29 uint8 = 247
                _ = v29
            }
        } else {
            v3 += v3
        }
    }
    v2 = append(v2, "a")
    var v30 int16 = -32768
    var v31 int32 = 0
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
