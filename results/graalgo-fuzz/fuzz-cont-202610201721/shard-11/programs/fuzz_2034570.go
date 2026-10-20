package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int }
func genFn0(p0 uint8, p1 int8) string {
    fmt.Println(p0, p1)
    v4 := []int{}
    p1 = p1
    switch 2147483647 {
    case 42:
        v5 := GenS1{F0: "a", F1: (p1 % ((int8(p0)) | 1))}
        _ = v5
    case -949:
        for fi6 := 0; fi6 < 12; fi6++ {
            fmt.Println(p1, p1)
            _ = fi6
        }
    case 4294967295:
        for fi7 := 0; fi7 < 5; fi7++ {
            if (true || (255 < 395)) {
                if len(v4) > 0 {
                    v4[0] = -1
                }
                fmt.Println(p1)
            }
            switch fi7 {
            case -149:
                if true {
                    switch fi7 {
                    case -9223372036854775808:
                        var v8 bool = false
                        _ = v8
                    case 127:
                        v9 := GenS2{F0: 504}
                        _ = v9
                    case 0:
                        var v10 string = "foo bar"
                        _ = v10
                    }
                    for fi11 := 0; fi11 < 7; fi11++ {
                        v12 := GenS2{F0: 1}
                        _ = fi11
                        _ = v12
                    }
                    var v13 int64 = int64(p0)
                    _ = v13
                }
            default:
                v14 := GenS1{F0: "foo bar", F1: p1}
                _ = v14
            }
            _ = fi7
        }
    default:
        p1 = (p1 | (p1 + -1))
    }
    _ = p0
    _ = p1
    _ = v4
    return "hello"
}

func main() {
    var v1 float64 = ((3.14159 * 0.1) - 0.0)
    var v2 bool = true
    var v3 uint64 = 42
    for fi15 := 0; fi15 < 17; fi15++ {
        v3 = v3
        _ = fi15
    }
    v16 := []int{32768, -2147483648, -601, -9223372036854775808}
    switch 531 {
    case -9223372036854775808:
        var v17 int32 = 0
        _ = v17
    case 32768:
        v18 := GenS3{F0: 128}
        _ = v18
    case 54:
        var v19 int64 = 472
        _ = v19
    default:
        fmt.Println(v3, v1)
    }
    switch ("foo bar" + ("\n\t" + "\n\t")) {
    case "hello":
        v1 = v1
    default:
        for fi20 := 0; fi20 < 4; fi20++ {
            if false {
                for fi21 := 0; fi21 < 8; fi21++ {
                    switch (fi21 | 2147483647) {
                    case -709:
                        v3 = v3
                    default:
                        v2 = false
                    }
                    v3 += 18446744073709551615
                    _ = fi21
                }
            }
            _ = fi20
        }
    }
    var v22 uint16 = 2
    v3 = 0
    v23 := GenS3{F0: -934}
    v22 = 278
    switch (("hello" + "foo bar") + "") {
    case "a":
        fmt.Println(v2, v1, v22)
    default:
        v3 -= uint64(v22)
    }
    v1 += 3.14159
    v24 := GenS1{F0: "a", F1: -128}
    v2 = (!(!true))
    var v25 int32 = 0
    v22 = (v22 >> ((503) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
