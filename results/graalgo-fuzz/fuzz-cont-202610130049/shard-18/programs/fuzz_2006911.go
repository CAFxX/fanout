package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 string }
type GenS2 struct { F0 uint8; F1 int32; F2 string }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int8, p1 string) int16 {
    switch p1 {
    case "":
        if false {
            fmt.Println(p1)
            if (!false) {
                p1 = p1
                p0 += 78
                v5 := GenS1{F0: 620, F1: false, F2: "a"}
                _ = v5
            }
        } else {
            var v6 uint64 = 0
            if true {
                fmt.Println(v6)
                if (!(true || true)) {
                    var v7 int64 = 1
                    _ = v7
                }
                if true {
                    v6 = (v6 >> ((42) & 63))
                    v6 -= v6
                    var v8 bool = true
                    _ = v8
                } else {
                    var v9 uint = 436
                    _ = v9
                }
            } else {
                v10 := GenS1{F0: -1, F1: (445 > -2147483648), F2: ((p1 + p1) + p1)}
                var v11 uint = 0
                _ = v10
                _ = v11
            }
            var v12 int = 0
            _ = v12
            _ = v6
        }
    case "a":
        p1 = ((p1 + p1) + "hello")
    default:
        for fi13 := 0; fi13 < 12; fi13++ {
            v14 := []int{2147483647, (fi13 ^ (fi13 & (fi13 + fi13))), -1}
            _ = fi13
            _ = v14
        }
    }
    p0 += p0
    _ = p0
    _ = p1
    return -1
}

func main() {
    v1 := []int{32768, 2147483647, 1}
    var v2 int = -930
    var v3 uint16 = 2
    var v4 int = v2
    v15 := GenS1{F0: 0, F1: (3.14159 == float32(-2.25)), F2: ("hello" + (((("a" + "") + "hello") + ("" + "")) + ""))}
    v16 := GenS1{F0: -2147483648, F1: true, F2: "a"}
    var v17 uint16 = 696
    v18 := GenS2{F0: 255, F1: 2147483647, F2: (("" + "a") + ("" + ("\n\t" + "")))}
    v19 := GenS2{F0: 127, F1: 42, F2: (("\n\t" + ("foo bar" + "hello")) + "hello")}
    v20 := GenS1{F0: -681, F1: true, F2: "foo bar"}
    if ((210 != 92) || false) {
        fmt.Println(v2, v4)
        v2 -= v4
        var v21 int64 = 9223372036854775807
        _ = v21
    } else {
        var v22 float64 = -2.25
        if (42 < -1) {
            var v23 uint64 = 192
            if false {
                v24 := GenS2{F0: 128, F1: -2147483648, F2: ""}
                var v25 string = "foo bar"
                var v26 float64 = v22
                _ = v24
                _ = v25
                _ = v26
            }
            _ = v23
        } else {
            var v27 int16 = -32768
            _ = v27
        }
        switch "foo bar" {
        case "a":
            if len(v1) > 0 {
                v1[0] = (v2 & (v4 * v2))
            }
        case "foo bar":
            for fi28 := 0; fi28 < 20; fi28++ {
                var v29 uint = 18446744073709551615
                for fi30 := 0; fi30 < 10; fi30++ {
                    v1 = append(v1, 128)
                    v31 := GenS1{F0: 0, F1: true, F2: ("hello" + "foo bar")}
                    v2 = (fi30 << (((v29 % (((v29 / ((uint(v3)) | 1))) | 1))) & 63))
                    _ = fi30
                    _ = v31
                }
                _ = fi28
                _ = v29
            }
        }
        _ = v22
    }
    v32 := GenS3{F0: 0, F1: 0}
    v33 := GenS4{F0: 255}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = v4
    }
    v4 += v2
    fmt.Println(v3, v2)
    v1 = append(v1, -513)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch ("\n\t" + (("foo bar" + "\n\t") + "hello")) {
    case "hello":
        for fi34 := 0; fi34 < 2; fi34++ {
            var v35 float64 = -2.25
            _ = fi34
            _ = v35
        }
    default:
        fmt.Println(v3, v3)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
