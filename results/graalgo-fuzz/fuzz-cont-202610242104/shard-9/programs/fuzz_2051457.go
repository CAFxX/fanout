package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int; F1 int16 }

func main() {
    var v1 float64 = -2.25
    v2 := []string{("foo bar" + "hello"), "\n\t", "a", ""}
    var v3 int32 = 2147483647
    var v4 int8 = -62
    v5 := make(map[string]string)
    v6 := GenS1{F0: (v3 / ((2147483647) | 1))}
    fmt.Println(v1, v4)
    if (255 >= 255) {
        v7 := v5[""]
        var v8 uint16 = 2
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := GenS2{F0: 1, F1: 4294967295}
            v1 = ((-2.25 - v1) * v1)
            for fi11 := 0; fi11 < 6; fi11++ {
                _ = len(v5)
                var v12 string = "a"
                v13 := GenS2{F0: -1, F1: 1}
                v3 -= v3
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi9
            _ = v10
        }
        if false {
            if (!false) {
                v1 += ((v1 + v1) * v1)
            }
            var v14 float32 = (1e10 - 1.5)
            _ = v14
        } else {
            switch int(v3) {
            case 127:
                switch "foo bar" {
                case "a":
                    v3 = (-(v3))
                case "":
                    v4 = 42
                default:
                    v4 = (v4 / (((v4 ^ v4)) | 1))
                }
            case 0:
                v3 = v3
            case -9223372036854775808:
                for fi15 := 0; fi15 < 3; fi15++ {
                    v1 += (v1 - v1)
                    v5[("" + (("hello" + ("\n\t" + "")) + "\n\t"))] = (("foo bar" + "a") + "\n\t")
                    _ = fi15
                }
            }
        }
    }
    v3 = v3
    var v16 float32 = (1.5 / (1e10 / (((float32(v1) / (0.0 / float32(v1))) / -2.25) / (3.14159 + 0.0))))
    v5["foo bar"] = "a"
    var v17 int = 656
    if false {
        var v18 bool = false
        var v19 int32 = 1
        _ = v18
        _ = v19
    }
    if (((v1 / (v1 * (v1 - -2.25))) + v1) == ((float64(3.14159) + 1.5) / (v1 / 0.1))) {
        var v20 uint64 = 2
        _ = v20
    } else {
        var v21 uint32 = 1
        fmt.Println(v3, v3, v3)
        _ = len(v2)
        _ = v21
    }
    if (!false) {
        v4 -= (v4 + (v4 + (v4 ^ 42)))
        v22 := GenS2{F0: 1, F1: 982}
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        _ = v22
    }
    if false {
        fmt.Println(v16, v17, v16)
        v1 -= v1
    } else {
        v23 := GenS3{F0: v17, F1: 42}
        _ = v23
    }
    var v24 uint16 = 234
    fmt.Println(v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
