package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 string }
func genFn0(p0 int8, p1 uint16, p2 string) uint8 {
    p0 -= (^(p0))
    if ((!true) && (0 == 1)) {
        p0 = (-(p0))
        for fi6 := 0; fi6 < 15; fi6++ {
            p2 = p2
            if false {
                var v7 int16 = 0
                p2 = (p2 + "hello")
                _ = v7
            } else {
                var v8 string = (p2 + p2)
                var v9 uint16 = (p1 ^ (^(p1)))
                v10 := []int{(^(fi6)), fi6, fi6, fi6}
                _ = v10
                _ = v8
                _ = v9
            }
            fmt.Println(p1, p2)
            _ = fi6
        }
        p0 = p0
    }
    switch (p2 + (((p2 + p2) + ((("" + "") + ("" + ("hello" + ((p2 + "hello") + "foo bar")))) + (p2 + ((("foo bar" + ("" + "a")) + p2) + "hello")))) + p2)) {
    case "foo bar":
        var v11 uint64 = 460
        _ = v11
    case "\n\t":
        switch ((p2 + p2) + (p2 + ("a" + p2))) {
        case "\n\t":
            v12 := []int{421}
            _ = v12
        case "foo bar":
            p0 = p0
        case "":
            v13 := []int{-103}
            _ = v13
        }
    default:
        var v14 uint = 18446744073709551615
        _ = v14
    }
    p0 = -128
    for fi15 := 0; fi15 < 5; fi15++ {
        fmt.Println(p0, p1)
        p0 = p0
        var v16 int64 = -927
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    var v1 int64 = -691
    var v2 int32 = -2147483648
    v3 := []string{("foo bar" + "foo bar"), "a", ""}
    var v4 uint16 = 65535
    var v5 uint = 2
    v1 -= (v1 << ((721) & 63))
    if true {
        if ((!((((3.14159 == 0.1) && (!false)) && false) && false)) || (255 == 49)) {
            fmt.Println(v4, v5)
            for fi17 := 0; fi17 < 15; fi17++ {
                v4 -= v4
                _ = len(v3)
                for fi18 := 0; fi18 < 10; fi18++ {
                    switch "" {
                    case "foo bar":
                        var v19 uint64 = 0
                        _ = v19
                    case "a":
                        v2 = -956
                    }
                    v1 = int64(fi18)
                    v20 := GenS1{F0: 12, F1: int32(fi17)}
                    _ = fi18
                    _ = v20
                }
                _ = fi17
            }
        }
        var v21 int16 = 1000
        var v22 uint = 87
        _ = v21
        _ = v22
    } else {
        if len(v3) > 0 {
            v3[0] = "a"
        }
    }
    switch 4294967295 {
    case -355:
        for fi23 := 0; fi23 < 19; fi23++ {
            v24 := GenS1{F0: 1, F1: (v2 & (^(v2)))}
            var v25 uint16 = 872
            _ = fi23
            _ = v24
            _ = v25
        }
    case 128:
        switch 32767 {
        case -9223372036854775808:
            for fi26 := 0; fi26 < 2; fi26++ {
                fmt.Println(fi26)
                _ = fi26
            }
        case 32767:
            v27 := GenS2{F0: (v1 - -1)}
            _ = v27
        default:
            var v28 uint = 259
            _ = v28
        }
    }
    v1 = (v1 * 339)
    var v29 bool = true
    for fi30 := 0; fi30 < 8; fi30++ {
        v31 := GenS1{F0: 910, F1: (v2 % ((v2) | 1))}
        v4 -= 42
        v3 = append(v3, "foo bar")
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 15; fi32++ {
        v4 = (v4 >> ((v5) & 63))
        v2 += (v2 >> ((284) & 63))
        for fi33 := 0; fi33 < 16; fi33++ {
            var v34 int8 = 102
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    v3 = append(v3, (("" + "a") + "foo bar"))
    v35 := GenS3{F0: -32768, F1: ""}
    v1 = (v1 - (v1 ^ 1))
    var v36 uint8 = 128
    fmt.Println(v36)
    var v37 float32 = (-2.25 * 0.0)
    v37 -= v37
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
