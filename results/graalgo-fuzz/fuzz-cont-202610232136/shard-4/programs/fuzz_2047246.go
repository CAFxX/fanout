package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int32; F1 uint32; F2 string }
type GenS3 struct { F0 uint32; F1 uint32; F2 int16 }

func main() {
    var v1 float32 = 3.14159
    v2 := []uint32{336, 713, 336}
    v3 := make(map[string]string)
    v4 := make(map[int]uint32)
    v5 := []uint32{4294967295, 777}
    switch -1 {
    case 0:
        v1 -= v1
    }
    var v6 bool = (false || (("a" + "") <= "hello"))
    v1 = 3.14159
    v7 := GenS1{F0: 0, F1: (true || true)}
    var v8 float32 = 3.14159
    var v9 uint64 = 619
    var v10 uint = 635
    v11 := GenS2{F0: -89, F1: 42, F2: "foo bar"}
    if len(v5) > 0 {
        v5[0] = 72
    }
    if v6 {
        if (0.0 < (float32((1.5 / 0.1)) / v8)) {
            v12 := GenS1{F0: 30, F1: ((2147483647 != 2) && v6)}
            fmt.Println(v1)
            _ = v12
        }
        v13 := GenS2{F0: int32(v10), F1: 886, F2: "foo bar"}
        _ = v13
    }
    v6 = true
    v14 := GenS3{F0: 65535, F1: 42, F2: -32768}
    var v15 int = 2147483647
    switch v15 {
    case -1:
        v16 := GenS1{F0: 127, F1: v6}
        _ = v16
    case 395:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 9223372036854775807:
        var v17 int8 = -1
        _ = v17
    default:
        for fi18 := 0; fi18 < 3; fi18++ {
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 uint32 = 953
                v21 := GenS3{F0: (v20 >> (((v10 & v10)) & 63)), F1: (v20 & (v20 / ((42) | 1))), F2: 32767}
                _ = fi19
                _ = v20
                _ = v21
            }
            v6 = v6
            _ = fi18
        }
    }
    if (true || ("foo bar" <= "hello")) {
        if v6 {
            switch v15 {
            case -16:
                v22 := GenS3{F0: 2147483647, F1: 2, F2: -1}
                _ = v22
            case -1:
                var v23 int = (-(v15))
                _ = v23
            default:
                _ = len(v5)
            }
            if (!false) {
                fmt.Println(v6)
                fmt.Println(v10, v10)
                fmt.Println(v1)
            } else {
                v2 = append(v2, 2147483647)
                switch "a" {
                case "":
                    v24 := GenS3{F0: 921, F1: 921, F2: 889}
                    _ = v24
                case "foo bar":
                    v6 = false
                case "hello":
                    var v25 uint = v10
                    _ = v25
                }
            }
            v15 = v15
        } else {
            if true {
                delete(v3, "a")
            } else {
                for fi26 := 0; fi26 < 9; fi26++ {
                    switch 524 {
                    case 0:
                        v27 := GenS2{F0: 42, F1: 0, F2: ""}
                        _ = v27
                    case -9223372036854775808:
                        fmt.Println(fi26, v6, fi26)
                    case 42:
                        v6 = ((!(v6 || ((v6 && (float32(1e10) == 1e10)) && (false || false)))) || (!v6))
                    }
                    _ = fi26
                }
                v9 += (v9 | (^(v9)))
            }
            v9 = v9
        }
        v15 += (v15 / (((v15 >> ((v10) & 63))) | 1))
    } else {
        for fi28 := 0; fi28 < 1; fi28++ {
            var v29 int = 32768
            _ = fi28
            _ = v29
        }
        var v30 bool = v6
        switch "\n\t" {
        case "foo bar":
            if true {
                v10 = 42
                var v31 uint8 = 1
                if (v31 <= v31) {
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                    if (932 < (v9 >> (((v10 * uint(v9))) & 63))) {
                        v8 = v1
                        var v32 uint32 = 867
                        var v33 int8 = 100
                        _ = v32
                        _ = v33
                    }
                    var v34 uint64 = 18446744073709551615
                    _ = v34
                }
                _ = v31
            }
        case "":
            delete(v3, "")
        case "hello":
            for fi35 := 0; fi35 < 3; fi35++ {
                var v36 string = ""
                if v30 {
                    var v37 float32 = 1e10
                    _ = v37
                }
                var v38 int8 = 127
                _ = fi35
                _ = v36
                _ = v38
            }
        default:
            v6 = (0.1 >= float64((v1 / 1.5)))
        }
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
