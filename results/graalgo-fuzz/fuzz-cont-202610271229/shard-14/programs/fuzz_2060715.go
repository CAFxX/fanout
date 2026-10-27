package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int16 = 0
    var v2 int32 = -2147483648
    v3 := []int{470}
    var v4 int8 = 42
    var v5 uint32 = 1
    v6 := []string{"foo bar", ("\n\t" + "a"), ""}
    switch 127 {
    case 132:
        var v7 int32 = (-(v2))
        _ = v7
    case 1:
        var v8 int = -723
        _ = v8
    case 42:
        for fi9 := 0; fi9 < 2; fi9++ {
            for fi10 := 0; fi10 < 8; fi10++ {
                if true {
                    var v11 string = ""
                    if len(v3) > 0 {
                        v3[0] = (fi10 | int(v2))
                    }
                    v4 = (v4 % (((v4 * 1)) | 1))
                    _ = v11
                }
                _ = len(v3)
                var v12 int8 = 100
                _ = fi10
                _ = v12
            }
            for fi13 := 0; fi13 < 13; fi13++ {
                v2 = (v2 << ((42) & 63))
                _ = fi13
            }
            for fi14 := 0; fi14 < 20; fi14++ {
                fmt.Println(v4, fi9)
                _ = fi14
            }
            _ = fi9
        }
    }
    if (false && (!false)) {
        v1 = (v1 / (((v1 & -173)) | 1))
    }
    var v15 int = 9223372036854775807
    v4 += (-(v4))
    v5 -= (v5 << ((0) & 63))
    var v16 uint = 1
    switch v15 {
    case -987:
        v1 -= (v1 ^ 0)
    case -9223372036854775808:
        var v17 uint = (v16 ^ v16)
        _ = v17
    case 128:
        var v18 int32 = -1
        _ = v18
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        v20 := GenS1{F0: 580}
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 uint32 = (^(v5))
            fmt.Println(v15, v5, v16)
            _ = fi21
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v5 = (v5 | 65535)
    if true {
        fmt.Println(v2, v16, v5)
    } else {
        v16 = 65535
        fmt.Println(v5)
        switch "hello" {
        case "a":
            for fi23 := 0; fi23 < 14; fi23++ {
                v2 = (v2 >> ((v16) & 63))
                _ = fi23
            }
        default:
            v4 = v4
        }
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        v4 = v4
        if (1.5 > 1e10) {
            var v25 int16 = v1
            _ = v25
        } else {
            if (!false) {
                v1 -= v1
                var v26 uint64 = 1
                _ = v26
            }
            if len(v3) > 0 {
                v3[0] = (^(fi24))
            }
            for fi27 := 0; fi27 < 12; fi27++ {
                v2 = (v2 | (v2 + v2))
                v5 += (v5 << ((2) & 63))
                var v28 uint64 = 18446744073709551615
                _ = fi27
                _ = v28
            }
        }
        _ = fi24
    }
    if (false && (!false)) {
        var v29 uint64 = 18446744073709551615
        _ = v29
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi30 := 0; fi30 < 19; fi30++ {
            _ = len(v6)
            _ = fi30
        }
    }
    v31 := GenS1{F0: 1}
    var v32 int8 = (v4 + (v4 + -1))
    for fi33 := 0; fi33 < 7; fi33++ {
        switch fi33 {
        case -9223372036854775808:
            v2 = (-(v2))
        default:
            if (42 >= 530) {
                var v34 int8 = (v32 % ((v4) | 1))
                for fi35 := 0; fi35 < 16; fi35++ {
                    var v36 float32 = ((((3.14159 / 1.5) + ((float32((float64((((1.5 + 1.5) - ((1.5 - ((-2.25 / (float32((float64(3.14159) - -2.25)) * 1.5)) - 0.1)) - 0.1)) * -2.25)) + ((3.14159 / -2.25) * 0.0))) + 3.14159) * 3.14159)) - ((-2.25 + (1e10 - 1e10)) - 3.14159)) / -2.25)
                    fmt.Println(v2)
                    fmt.Println(v36)
                    _ = fi35
                    _ = v36
                }
                _ = v34
            }
        }
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
