package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
func genFn0(p0 uint, p1 bool, p2 uint16) int16 {
    var v4 uint32 = 65535
    p0 = (p0 + (^(p0)))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -32768
}
func genFn1(p0 bool, p1 bool) uint {
    p0 = false
    for fi5 := 0; fi5 < 20; fi5++ {
        p0 = p1
        switch fi5 {
        case 649:
            v6 := []int{(fi5 - fi5), fi5, fi5, fi5}
            _ = v6
        case 162:
            fmt.Println(p1, p0)
        case -899:
            p0 = (0 != 984)
        }
        p0 = (127 > -128)
        _ = fi5
    }
    fmt.Println(p0)
    p0 = p1
    p0 = (false || p1)
    _ = p0
    _ = p1
    return 65535
}

func main() {
    v1 := []string{"\n\t", "\n\t", ""}
    v2 := make(map[int]string)
    var v3 int = 9223372036854775807
    var v7 uint64 = 2
    delete(v2, 409)
    var v8 int = (v3 + v3)
    switch (("hello" + "\n\t") + "a") {
    case "\n\t":
        for fi9 := 0; fi9 < 4; fi9++ {
            for fi10 := 0; fi10 < 18; fi10++ {
                if true {
                    v7 = (^(v7))
                    fmt.Println(v7, fi10, v7)
                    v7 -= v7
                } else {
                    v8 -= fi10
                    var v11 int = (-(v3))
                    _ = v11
                }
                _ = fi10
            }
            if (65535 > 2147483648) {
                var v12 int64 = -108
                var v13 float32 = 1.5
                fmt.Println(v7)
                _ = v12
                _ = v13
            }
            _ = fi9
        }
    default:
        if ((true && (!(!true))) && (!false)) {
            v8 = (v3 << ((42) & 63))
        }
    }
    fmt.Println(v8, v3)
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 bool = true
        _ = fi14
        _ = v15
    }
    var v16 int32 = 145
    if false {
        fmt.Println(v8, v16, v16)
    } else {
        v8 -= -537
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        if (0 != -880) {
            var v18 uint = 280
            _ = v18
        } else {
            v3 += v8
            _ = len(v2)
        }
        v1 = append(v1, "a")
        _ = fi17
    }
    v8 = (v3 | v3)
    v16 = v16
    v19 := GenS1{F0: false, F1: 2147483647}
    if (128 != 140) {
        for fi20 := 0; fi20 < 14; fi20++ {
            v16 = (v16 ^ v16)
            switch ("foo bar" + "a") {
            case "hello":
                if true {
                    fmt.Println(v3)
                    if true {
                        var v21 uint = 13
                        fmt.Println(v8, v21, v3)
                        v8 = 9223372036854775807
                        _ = v21
                    }
                    var v22 string = "\n\t"
                    _ = v22
                }
            case "foo bar":
                v1 = append(v1, "\n\t")
            default:
                var v23 float32 = float32(-2.25)
                _ = v23
            }
            v24 := v2[v3]
            _ = fi20
            _ = v24
        }
        switch 32768 {
        case 31:
            _ = len(v2)
        case 32768:
            fmt.Println(v7, v3, v8)
        }
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v8)
        case "a":
            var v25 uint64 = v7
            _ = v25
        }
    } else {
        var v26 uint = 18446744073709551615
        var v27 uint = v26
        _ = v26
        _ = v27
    }
    delete(v2, v8)
    var v28 int32 = v16
    for fi29 := 0; fi29 < 8; fi29++ {
        v30 := GenS1{F0: true, F1: v28}
        v7 = (v7 / ((v7) | 1))
        _ = fi29
        _ = v30
    }
    var v31 uint32 = 4294967295
    v2[v8] = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v7
    _ = v8
}
