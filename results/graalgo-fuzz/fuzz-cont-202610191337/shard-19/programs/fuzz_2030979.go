package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int16; F2 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 int32 }
func genFn0(p0 string, p1 bool) int {
    v4 := GenS1{F0: 2, F1: 2}
    if p1 {
        for fi5 := 0; fi5 < 2; fi5++ {
            for fi6 := 0; fi6 < 3; fi6++ {
                if p1 {
                    p1 = (128 <= 1)
                    v7 := []int{fi6, (fi5 + fi5), fi6, fi5}
                    p0 = "foo bar"
                    _ = v7
                } else {
                    v8 := GenS2{F0: 340, F1: -32768, F2: 32767}
                    _ = v8
                }
                var v9 int = fi6
                _ = fi6
                _ = v9
            }
            _ = fi5
        }
    }
    var v10 uint16 = 834
    _ = p0
    _ = p1
    _ = v10
    _ = v4
    return 32767
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]uint32)
    var v3 float32 = float32(1.5)
    var v11 bool = false
    var v12 float32 = ((v3 + v3) + -2.25)
    if ("foo bar" > "a") {
        v13 := GenS3{F0: ""}
        if ((!true) && false) {
            var v14 int64 = 494
            for fi15 := 0; fi15 < 12; fi15++ {
                v3 = v12
                var v16 uint64 = 1
                _ = fi15
                _ = v16
            }
            v17 := GenS4{F0: 255, F1: 61}
            _ = v14
            _ = v17
        } else {
            var v18 uint16 = 966
            v19 := v2[2147483648]
            _ = v18
            _ = v19
        }
        v11 = v11
        _ = v13
    }
    v20 := GenS2{F0: 42, F1: 1, F2: -1}
    var v21 uint8 = 255
    v22 := []int{127, -9223372036854775808, genFn0("", v11)}
    v23 := v2[395]
    var v24 int8 = 1
    v25 := GenS2{F0: 42, F1: -119, F2: int16(v1)}
    v23 = (v23 & 629)
    v26 := GenS5{F0: (^(v23)), F1: int32(v23)}
    if v11 {
        if (65535 > 0) {
            v11 = false
            v27 := GenS3{F0: "\n\t"}
            _ = v27
        }
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        v12 = 1e10
    } else {
        switch "hello" {
        case "hello":
            for fi28 := 0; fi28 < 1; fi28++ {
                var v29 int16 = 32767
                v23 += (^(v23))
                v30 := GenS2{F0: (v23 / (((^(v23))) | 1)), F1: v29, F2: int16(v1)}
                _ = fi28
                _ = v29
                _ = v30
            }
        default:
            v31 := GenS6{F0: -1}
            _ = v31
        }
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        switch (((((("hello" + "foo bar") + (("\n\t" + "a") + (("" + ("a" + "hello")) + ("" + "foo bar")))) + (("\n\t" + "hello") + "a")) + "\n\t") + "") + "\n\t") {
        case "foo bar":
            var v32 bool = ((471 != 1) || (128 < 2147483647))
            _ = v32
        }
    }
    var v33 int = -605
    v1 = (v1 ^ v1)
    v11 = (v12 >= v12)
    switch (("foo bar" + "\n\t") + "foo bar") {
    case "\n\t":
        v1 = v1
    case "a":
        _ = len(v2)
    case "hello":
        var v34 uint16 = 255
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v33
}
