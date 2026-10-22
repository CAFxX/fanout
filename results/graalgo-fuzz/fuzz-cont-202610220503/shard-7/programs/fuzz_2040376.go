package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 string }
type GenS2 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS3 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint8, p1 int, p2 int) string {
    if (127 >= -1) {
        v6 := []int{}
        v7 := GenS1{F0: 18446744073709551615, F1: "\n\t", F2: "hello"}
        _ = v6
        _ = v7
    }
    var v8 int32 = -262
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return ""
}

func main() {
    var v1 uint8 = 243
    var v2 string = ""
    var v3 int8 = 1
    var v4 int32 = -644
    v5 := []string{(v2 + "a"), ("a" + (v2 + "foo bar"))}
    var v9 uint16 = 497
    v10 := GenS1{F0: 18446744073709551615, F1: v2, F2: v2}
    var v11 string = v2
    v12 := GenS1{F0: 326, F1: v2, F2: (v11 + "foo bar")}
    v5 = append(v5, "hello")
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int64 = int64(v4)
        var v15 float32 = (0.1 * 0.1)
        var v16 uint32 = 315
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    if true {
        for fi17 := 0; fi17 < 20; fi17++ {
            _ = len(v5)
            _ = fi17
        }
        var v18 bool = true
        _ = v18
    }
    switch -2147483648 {
    case 127:
        var v19 int64 = 9223372036854775807
        _ = v19
    }
    var v20 uint64 = 646
    if (v2 < v11) {
        v11 = (v11 + "\n\t")
        if (-2.25 > 1e10) {
            switch 0 {
            case 9223372036854775807:
                v11 = (v2 + ((((v11 + (((v2 + ("foo bar" + v11)) + (v11 + (((v2 + v2) + ((v2 + "\n\t") + "foo bar")) + (v2 + "\n\t")))) + ((v2 + (v2 + "foo bar")) + v11))) + ((v11 + ("\n\t" + v2)) + "foo bar")) + v2) + (v11 + (v11 + v2))))
            }
            var v21 int16 = 32767
            var v22 int8 = v3
            _ = v21
            _ = v22
        }
    }
    var v23 int32 = (v4 % (((v4 - (-(v4)))) | 1))
    for fi24 := 0; fi24 < 14; fi24++ {
        switch v11 {
        case "":
            v25 := GenS2{F0: 4294967295, F1: (^(v9)), F2: -805}
            _ = v25
        case "hello":
            var v26 float64 = 0.0
            _ = v26
        default:
            if len(v5) > 0 {
                v5[0] = v11
            }
        }
        _ = fi24
    }
    v23 = v23
    v27 := GenS3{F0: 32767, F1: (v1 / ((213) | 1))}
    for fi28 := 0; fi28 < 11; fi28++ {
        var v29 int8 = v3
        var v30 string = v11
        var v31 uint8 = (v1 ^ v1)
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    v23 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
