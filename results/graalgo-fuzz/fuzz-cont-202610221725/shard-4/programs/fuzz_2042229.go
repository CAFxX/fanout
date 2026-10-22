package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int; F2 uint8 }
func genFn0(p0 int64, p1 bool, p2 uint) uint16 {
    switch ("hello" + "a") {
    case "":
        var v4 int64 = p0
        _ = v4
    case "foo bar":
        p1 = p1
    default:
        v5 := GenS1{F0: 1, F1: ((32767 < 2) && p1), F2: 275}
        _ = v5
    }
    switch -741 {
    case -2147483648:
        p0 = p0
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        fmt.Println(p1)
        p1 = true
        _ = fi6
    }
    p2 = p2
    if (0 > 0) {
        for fi7 := 0; fi7 < 20; fi7++ {
            p1 = (!true)
            v8 := GenS1{F0: 65535, F1: false, F2: 1}
            var v9 uint8 = 216
            _ = fi7
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := GenS1{F0: 255, F1: p1, F2: 981}
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 965
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int64 = -1
    v3 := []int{}
    if (true || true) {
        fmt.Println(v1, v2)
        if len(v3) > 0 {
            v3[0] = -1
        }
        switch "\n\t" {
        case "hello":
            var v12 int64 = (v2 - v2)
            _ = v12
        default:
            fmt.Println(v1)
        }
    } else {
        if false {
            v13 := GenS2{F0: 65535}
            for fi14 := 0; fi14 < 17; fi14++ {
                v2 = v2
                if (!false) {
                    v15 := GenS2{F0: 1}
                    _ = v15
                } else {
                    var v16 float32 = ((3.14159 - ((1.5 / -2.25) - 0.0)) - (float32((-2.25 - (1e10 / 0.1))) / 3.14159))
                    if len(v3) > 0 {
                        v3[0] = (fi14 + fi14)
                    }
                    v2 = (v2 ^ v2)
                    _ = v16
                }
                _ = fi14
            }
            if len(v3) > 0 {
                v3[0] = 42
            }
            _ = v13
        } else {
            if false {
                fmt.Println(v2)
                var v17 int16 = 32767
                _ = v17
            }
            var v18 string = "a"
            v19 := GenS2{F0: 953}
            _ = v18
            _ = v19
        }
    }
    var v20 int16 = 32767
    v21 := GenS3{F0: (v2 | (-(v2))), F1: 455, F2: 127}
    var v22 int8 = -55
    for fi23 := 0; fi23 < 8; fi23++ {
        v2 = (^(v2))
        _ = fi23
    }
    v24 := GenS3{F0: (-(v2)), F1: -1, F2: 158}
    v1 += v1
    v2 -= -9223372036854775808
    switch ("foo bar" + ((("\n\t" + "a") + ("\n\t" + ("" + ("hello" + ("hello" + ("foo bar" + "a")))))) + ("a" + "foo bar"))) {
    case "\n\t":
        var v25 uint16 = 141
        _ = v25
    default:
        if true {
            var v26 int = -835
            _ = v26
        } else {
            v20 = (^(v20))
            for fi27 := 0; fi27 < 6; fi27++ {
                var v28 uint32 = 65535
                _ = fi27
                _ = v28
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
}
