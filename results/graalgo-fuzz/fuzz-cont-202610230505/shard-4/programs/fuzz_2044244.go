package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint64) uint16 {
    p0 = (p0 & p0)
    var v5 float32 = float32((((1.5 + 3.14159) + (0.0 / 1e10)) * 1e10))
    var v6 uint = 802
    if true {
        v7 := GenS1{F0: ("\n\t" + "\n\t")}
        var v8 int32 = 0
        switch "\n\t" {
        case "\n\t":
            v9 := GenS2{F0: 376, F1: (!(-156 < 1))}
            _ = v9
        case "hello":
            var v10 int32 = v8
            _ = v10
        default:
            if (((752 >= 1) || true) || ((false && (!false)) || true)) {
                if (true && ((65535 >= uint16(v6)) && (false && (-25 < 1)))) {
                    for fi11 := 0; fi11 < 6; fi11++ {
                        v5 = (v5 * v5)
                        var v12 uint32 = 2147483647
                        _ = fi11
                        _ = v12
                    }
                }
                switch ("\n\t" + "\n\t") {
                case "foo bar":
                    p0 = (p0 % (((p0 - 140)) | 1))
                case "hello":
                    for fi13 := 0; fi13 < 8; fi13++ {
                        v14 := GenS2{F0: -440, F1: true}
                        v6 -= v6
                        v6 = (v6 + v6)
                        _ = fi13
                        _ = v14
                    }
                }
                var v15 string = ""
                _ = v15
            }
        }
        _ = v7
        _ = v8
    }
    v16 := GenS3{F0: 2147483647}
    _ = p0
    _ = v16
    _ = v5
    _ = v6
    return 32767
}
func genFn1(p0 int64) int16 {
    if false {
        var v17 string = ""
        v18 := GenS1{F0: "\n\t"}
        _ = v17
        _ = v18
    } else {
        p0 = p0
        var v19 int16 = 130
        v20 := []int{683, -663, 128, 32768}
        _ = v19
        _ = v20
    }
    p0 -= -9223372036854775808
    _ = p0
    return 293
}

func main() {
    var v1 int64 = 42
    var v2 uint8 = 5
    v3 := []string{(("hello" + "\n\t") + ""), ("\n\t" + "a")}
    v4 := make(map[string]bool)
    var v21 int16 = 32767
    var v22 float32 = 3.14159
    for fi23 := 0; fi23 < 4; fi23++ {
        v24 := v4["foo bar"]
        for fi25 := 0; fi25 < 18; fi25++ {
            var v26 uint16 = 255
            fmt.Println(v22, fi23)
            v27 := GenS3{F0: 447}
            _ = fi25
            _ = v26
            _ = v27
        }
        v22 = float32(1.5)
        _ = fi23
        _ = v24
    }
    v28 := GenS1{F0: "hello"}
    var v29 bool = (1.5 == float64((float32(1.5) + float32(((float64((v22 + v22)) / (0.0 + ((1.5 * 0.1) - float64(v22)))) / 0.1)))))
    switch 171 {
    case 2147483647:
        var v30 float32 = -2.25
        _ = v30
    case -740:
        for fi31 := 0; fi31 < 9; fi31++ {
            v4[("foo bar" + "a")] = (fi31 < 9223372036854775807)
            _ = fi31
        }
    }
    fmt.Println(v1, v21)
    v32 := GenS1{F0: ""}
    var v33 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
