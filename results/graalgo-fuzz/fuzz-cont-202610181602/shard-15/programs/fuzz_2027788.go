package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int8; F1 uint8; F2 int }
type GenS3 struct { F0 uint16 }
func genFn0(p0 string) int {
    v7 := GenS1{F0: 2, F1: -332}
    p0 = p0
    p0 = (((p0 + p0) + p0) + p0)
    _ = p0
    _ = v7
    return 876
}

func main() {
    v1 := []string{(("hello" + "\n\t") + "a"), "\n\t", ""}
    var v2 bool = (9223372036854775807 == -1)
    var v3 int16 = 32767
    var v4 float64 = (0.0 / 1.0)
    v5 := make(map[string]int)
    var v6 float64 = v4
    if (!v2) {
        v8 := GenS2{F0: 0, F1: 42, F2: 127}
        _ = v8
    }
    if len(v1) > 0 {
        v1[0] = (("\n\t" + "foo bar") + "foo bar")
    }
    v9 := GenS3{F0: 65535}
    if v2 {
        v4 -= 3.14159
    }
    v2 = v2
    v6 += 0.1
    if (!v2) {
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 int8 = 100
            _ = fi10
            _ = v11
        }
        fmt.Println(v4, v2)
        if (true && (18446744073709551615 != 2)) {
            fmt.Println(v2, v3)
            for fi12 := 0; fi12 < 15; fi12++ {
                if (42 != -1) {
                    v13 := GenS3{F0: 1}
                    var v14 int8 = 58
                    _ = v13
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 19; fi15++ {
                        fmt.Println(v4, v4, v6)
                        _ = fi15
                    }
                }
                var v16 bool = false
                var v17 float32 = -2.25
                _ = fi12
                _ = v16
                _ = v17
            }
        }
    } else {
        fmt.Println(v2)
        var v18 int32 = 619
        _ = v18
    }
    switch -392 {
    case 32767:
        switch "\n\t" {
        case "a":
            for fi19 := 0; fi19 < 11; fi19++ {
                fmt.Println(fi19, v6)
                v2 = (527 <= -2147483648)
                _ = fi19
            }
        default:
            var v20 int8 = 127
            _ = v20
        }
    case -114:
        var v21 bool = true
        _ = v21
    }
    if len(v1) > 0 {
        v1[0] = ("a" + ("" + "foo bar"))
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v22 bool = v2
    switch -618 {
    case -2147483648:
        for fi23 := 0; fi23 < 16; fi23++ {
            for fi24 := 0; fi24 < 16; fi24++ {
                delete(v5, "\n\t")
                v4 = 0.1
                _ = fi24
            }
            _ = fi23
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
