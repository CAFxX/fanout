package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int64 }
type GenS2 struct { F0 string; F1 uint16; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int64; F1 int16 }

func main() {
    var v1 int = -872
    var v2 uint64 = 4294967295
    var v3 bool = (42 < 1)
    var v4 int32 = 1
    v5 := []int{(v1 - (v1 * 4294967295)), (^(v1)), (v1 % (((^(v1))) | 1))}
    var v6 int16 = -32768
    v7 := GenS1{F0: "foo bar", F1: 902, F2: -1000}
    var v8 uint8 = 188
    switch 2147483647 {
    case 32768:
        var v9 int16 = v6
        _ = v9
    default:
        v5 = append(v5, v1)
    }
    if (!("foo bar" != ("\n\t" + "foo bar"))) {
        if (!v3) {
            fmt.Println(v6, v1, v1)
            v1 -= 763
            for fi10 := 0; fi10 < 7; fi10++ {
                v5 = append(v5, int(v4))
                var v11 uint8 = v8
                switch (v1 >> ((65535) & 63)) {
                case -500:
                    v4 += 42
                default:
                    fmt.Println(v3, v2, v8)
                }
                _ = fi10
                _ = v11
            }
        } else {
            v12 := GenS2{F0: ("" + (("a" + "\n\t") + ((("\n\t" + ("" + ("" + "\n\t"))) + (("foo bar" + "") + (("a" + ("" + "")) + ("a" + "")))) + (("hello" + "foo bar") + "")))), F1: 255, F2: -9223372036854775808}
            var v13 uint8 = 1
            var v14 int16 = v6
            _ = v12
            _ = v13
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 8; fi15++ {
            v8 = (v8 / (((v8 - 0)) | 1))
            switch -616 {
            case 42:
                for fi16 := 0; fi16 < 1; fi16++ {
                    switch (-(fi16)) {
                    case -986:
                        var v17 uint64 = v2
                        _ = v17
                    case 32768:
                        v4 = v4
                    }
                    if true {
                        v6 = (-(v6))
                        v18 := GenS3{F0: -128}
                        _ = v18
                    } else {
                        var v19 uint = 18446744073709551615
                        fmt.Println(fi15)
                        fmt.Println(v4, v1)
                        _ = v19
                    }
                    v20 := GenS4{F0: (v4 % ((v4) | 1)), F1: int32(v8)}
                    _ = fi16
                    _ = v20
                }
            case 71:
                v21 := GenS5{F0: int64(v2), F1: v6}
                _ = v21
            case -485:
                if v3 {
                    var v22 float64 = 1.5
                    _ = v22
                }
            }
            v23 := GenS3{F0: 127}
            _ = fi15
            _ = v23
        }
    }
    v24 := GenS1{F0: ("" + "hello"), F1: 2, F2: -1}
    v3 = v3
    var v25 uint = 42
    fmt.Println(v8)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
