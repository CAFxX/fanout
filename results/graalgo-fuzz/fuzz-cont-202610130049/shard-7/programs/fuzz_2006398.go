package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 bool }

func main() {
    var v1 uint32 = 65535
    var v2 int8 = 42
    var v3 uint32 = 762
    var v4 uint64 = 2
    v5 := []string{}
    var v6 bool = (false || true)
    switch "\n\t" {
    case "":
        v7 := GenS1{F0: 9223372036854775807, F1: (9223372036854775807 < 783)}
        _ = v7
    case "foo bar":
        if ((((3.14159 >= ((1.5 - 0.1) - ((0.1 + ((0.1 - 1e10) + float32(1e10))) / float32(3.14159)))) && (false || (v2 >= v2))) && v6) || (18446744073709551615 > 1)) {
            var v8 uint8 = 88
            _ = v8
        } else {
            var v9 uint = 18446744073709551615
            if v6 {
                for fi10 := 0; fi10 < 8; fi10++ {
                    var v11 int32 = 42
                    var v12 bool = v6
                    v6 = false
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
            v9 = v9
            _ = v9
        }
    default:
        for fi13 := 0; fi13 < 18; fi13++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if ((3.14159 + float32(-2.25)) > (float32(-2.25) * 0.1)) {
                if len(v5) > 0 {
                    v5[0] = ""
                }
                if len(v5) > 0 {
                    v5[0] = ""
                }
            }
            _ = fi13
        }
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        _ = len(v5)
        if (-198 < 1) {
            var v15 int16 = 1
            _ = v15
        } else {
            var v16 uint = 0
            var v17 int32 = 178
            var v18 bool = (false && v6)
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi14
    }
    fmt.Println(v6, v4, v1)
    switch "\n\t" {
    case "foo bar":
        for fi19 := 0; fi19 < 17; fi19++ {
            v6 = true
            var v20 string = "\n\t"
            for fi21 := 0; fi21 < 13; fi21++ {
                v4 -= v4
                fmt.Println(v4)
                _ = fi21
            }
            _ = fi19
            _ = v20
        }
    }
    var v22 int = 0
    var v23 int32 = 962
    v2 -= (v2 * (v2 - (v2 | (v2 - 0))))
    var v24 float32 = ((((float32(0.0) - 0.0) * 0.1) / 0.1) - 0.1)
    v22 = (-(v22))
    if (v6 || (1e10 > 3.14159)) {
        for fi25 := 0; fi25 < 15; fi25++ {
            fmt.Println(v22, v6, v22)
            _ = fi25
        }
    }
    fmt.Println(v23, v3)
    for fi26 := 0; fi26 < 1; fi26++ {
        v4 += uint64(v1)
        if (!true) {
            if ((v4 % (((v4 + v4)) | 1)) >= (v4 % (((^(v4))) | 1))) {
                v4 = 0
                var v27 int8 = int8(v4)
                _ = v27
            }
            switch v22 {
            case 42:
                _ = len(v5)
            case 879:
                for fi28 := 0; fi28 < 8; fi28++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v29 uint64 = 0
                    _ = fi28
                    _ = v29
                }
            case -9223372036854775808:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                var v30 bool = v6
                _ = v30
            }
        } else {
            v22 -= 238
            v4 -= v4
            v31 := GenS2{F0: 20}
            _ = v31
        }
        fmt.Println(v24, v22, v2)
        _ = fi26
    }
    if false {
        fmt.Println(v22)
        var v32 float32 = v24
        fmt.Println(v32, v22, v24)
        _ = v32
    } else {
        v33 := GenS3{F0: "hello", F1: false}
        var v34 float64 = (1e10 + -2.25)
        fmt.Println(v2)
        _ = v33
        _ = v34
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
