package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 bool; F2 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 uint64; F2 int32 }
func genFn0(p0 int, p1 string, p2 int8) uint64 {
    p0 = p0
    p2 += p2
    p1 = p1
    var v4 uint64 = 0
    switch p1 {
    case "hello":
        p1 = p1
    default:
        v5 := GenS1{F0: 2147483647}
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return v4
}
func genFn1(p0 uint) uint8 {
    v6 := []int{2147483647, 32767, 0, 9223372036854775807}
    v7 := GenS1{F0: 921}
    _ = p0
    _ = v6
    _ = v7
    return 255
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 18446744073709551615
    var v3 int16 = 1
    fmt.Println(v2)
    v8 := []int{32768, 1, 42}
    v2 -= genFn0(2147483647, "\n\t", -128)
    _ = len(v1)
    for fi9 := 0; fi9 < 16; fi9++ {
        switch "a" {
        case "foo bar":
            for fi10 := 0; fi10 < 4; fi10++ {
                v11 := GenS2{F0: 32767}
                _ = fi10
                _ = v11
            }
        case "":
            v12 := GenS3{F0: 42, F1: false, F2: -110}
            _ = v12
        case "a":
            var v13 bool = true
            _ = v13
        }
        fmt.Println(v2)
        for fi14 := 0; fi14 < 10; fi14++ {
            if true {
                v15 := GenS1{F0: 296}
                v16 := GenS2{F0: 65535}
                v17 := GenS1{F0: 1}
                _ = v15
                _ = v16
                _ = v17
            } else {
                var v18 uint = 554
                switch (fi9 & fi9) {
                case 963:
                    v19 := GenS4{F0: -1}
                    _ = v19
                case 4294967295:
                    fmt.Println(v18, v2, v3)
                default:
                    fmt.Println(v2)
                }
                _ = v18
            }
            _ = len(v8)
            fmt.Println(fi14)
            _ = fi14
        }
        _ = fi9
    }
    v20 := GenS5{F0: (42 == 128), F1: v2, F2: -438}
    v3 = v3
    for fi21 := 0; fi21 < 12; fi21++ {
        if true {
            var v22 float32 = float32(((0.0 / (3.14159 * 3.14159)) / 3.14159))
            var v23 string = ""
            v22 = v22
            _ = v22
            _ = v23
        }
        for fi24 := 0; fi24 < 15; fi24++ {
            fmt.Println(v3, fi24, fi21)
            _ = fi24
        }
        _ = fi21
    }
    v3 -= -275
    v3 = -1
    _ = len(v8)
    if (!(true && true)) {
        if (((859 >= 1) || true) && (!true)) {
            v1["\n\t"] = (!false)
            v3 = v3
        } else {
            var v25 int64 = 42
            v3 -= -878
            delete(v1, "foo bar")
            _ = v25
        }
        switch 32768 {
        case -562:
            switch "a" {
            case "foo bar":
                v26 := GenS5{F0: (127 != 100), F1: (^(v2)), F2: -1}
                _ = v26
            case "a":
                _ = len(v1)
            case "\n\t":
                v2 -= (v2 ^ (v2 & 0))
            default:
                for fi27 := 0; fi27 < 17; fi27++ {
                    var v28 bool = false
                    _ = fi27
                    _ = v28
                }
            }
        default:
            for fi29 := 0; fi29 < 13; fi29++ {
                var v30 string = ""
                var v31 uint = 350
                var v32 uint64 = v2
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
        }
        var v33 int32 = 1
        _ = v33
    } else {
        fmt.Println(v3)
    }
    fmt.Println(v3, v2)
    v3 = -1
    if true {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    } else {
        v34 := GenS4{F0: -128}
        v2 += (v2 | (v2 - (v2 ^ (v2 | genFn0(-899, "", -1)))))
        _ = v34
    }
    fmt.Println(v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v8
}
