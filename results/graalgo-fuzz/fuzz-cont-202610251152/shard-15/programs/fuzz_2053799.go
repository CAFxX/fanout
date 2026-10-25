package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int32, p1 int32) uint32 {
    p0 -= p1
    switch "a" {
    case "\n\t":
        fmt.Println(p0, p1)
    case "":
        var v5 uint64 = 0
        _ = v5
    }
    if (!true) {
        if (!("foo bar" < "foo bar")) {
            if true {
                p0 += -255
                var v6 uint = 255
                _ = v6
            }
        }
        p1 += p0
        p1 = 224
    } else {
        var v7 float64 = (1.5 * float64((0.1 - 1.5)))
        _ = v7
    }
    var v8 string = "foo bar"
    _ = p0
    _ = p1
    _ = v8
    return 2147483648
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 2
    var v3 int32 = -1
    v4 := []string{}
    switch "" {
    case "foo bar":
        var v9 uint16 = (^(v2))
        _ = v9
    default:
        var v10 int64 = -1
        _ = v10
    }
    v2 = (v2 + v2)
    var v11 int16 = 42
    v3 = (v3 & (v3 ^ 2147483647))
    if true {
        v11 -= v11
        var v12 uint16 = 42
        _ = v12
    } else {
        if (-154 != -1) {
            if true {
                var v13 int = 127
                var v14 float32 = (0.1 - 0.1)
                _ = v13
                _ = v14
            } else {
                v15 := GenS1{F0: 1}
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
                _ = v15
            }
            v11 += 0
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "\n\t" {
    case "":
        v2 = v2
    case "\n\t":
        for fi16 := 0; fi16 < 6; fi16++ {
            for fi17 := 0; fi17 < 2; fi17++ {
                v2 -= v2
                fmt.Println(fi16)
                v2 += v2
                _ = fi17
            }
            fmt.Println(v2, v11)
            switch "foo bar" {
            case "\n\t":
                if true {
                    var v18 float32 = float32(((((0.0 - (0.0 - 3.14159)) - 3.14159) / 1.5) / float64(0.1)))
                    v1[genFn0((v3 >> ((661) & 63)), (v3 & (v3 & int32(v11))))] = fi16
                    _ = v18
                } else {
                    var v19 uint = 42
                    fmt.Println(v19, v11, v2)
                    fmt.Println(fi16, v3, fi16)
                    _ = v19
                }
            case "":
                v1[genFn0(-769, (v3 >> ((314) & 63)))] = fi16
            default:
                v11 -= v11
            }
            _ = fi16
        }
    case "a":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v4 = append(v4, "hello")
    var v20 bool = (true || true)
    v4 = append(v4, "hello")
    v11 += v11
    var v21 float64 = 1.5
    for fi22 := 0; fi22 < 6; fi22++ {
        v2 -= (v2 * v2)
        if v20 {
            v2 -= 765
            fmt.Println(v2, v3, v2)
        }
        v20 = true
        _ = fi22
    }
    var v23 float64 = v21
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
}
