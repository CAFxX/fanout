package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int64; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 uint = 320
    var v2 int16 = 1000
    var v3 string = "foo bar"
    v4 := make(map[int]int)
    var v5 int32 = 0
    v6 := make(map[string]string)
    v1 += (v1 % (((^(v1))) | 1))
    v7 := []int{42, 1, 128, -880}
    if (1 < 0) {
        v1 -= v1
        if true {
            var v8 int = 4294967295
            var v9 uint16 = 205
            for fi10 := 0; fi10 < 6; fi10++ {
                if (false || true) {
                    v11 := GenS1{F0: uint32(v1), F1: 544, F2: v9}
                    _ = v11
                }
                _ = fi10
            }
            _ = v8
            _ = v9
        } else {
            v12 := GenS1{F0: 465, F1: 2147483648, F2: 65535}
            _ = v12
        }
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if true {
        var v13 bool = true
        for fi14 := 0; fi14 < 15; fi14++ {
            if (((!((v2 - v2) > 0)) && ((255 == 2) || v13)) || v13) {
                for fi15 := 0; fi15 < 15; fi15++ {
                    v16 := v6[v3]
                    fmt.Println(v16, fi15)
                    _ = fi15
                    _ = v16
                }
                var v17 int16 = v2
                var v18 uint32 = 218
                _ = v17
                _ = v18
            }
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 bool = v13
                _ = fi19
                _ = v20
            }
            v21 := GenS2{F0: 4294967295, F1: 817}
            _ = fi14
            _ = v21
        }
        _ = v13
    }
    var v22 float32 = 0.1
    if true {
        if true {
            var v23 int16 = (-(v2))
            _ = v23
        } else {
            v24 := GenS1{F0: 42, F1: 2147483647, F2: 32767}
            _ = v24
        }
        for fi25 := 0; fi25 < 6; fi25++ {
            var v26 bool = false
            for fi27 := 0; fi27 < 8; fi27++ {
                v28 := GenS2{F0: 2, F1: 2147483647}
                v1 += v1
                _ = fi27
                _ = v28
            }
            var v29 int16 = v2
            _ = fi25
            _ = v26
            _ = v29
        }
        v30 := GenS2{F0: 2, F1: 4294967295}
        _ = v30
    }
    var v31 float32 = v22
    v32 := GenS3{F0: 9223372036854775807, F1: 48, F2: 4294967295}
    v33 := v6["foo bar"]
    var v34 uint16 = 1
    switch 0 {
    case -71:
        if true {
            v3 = v33
            for fi35 := 0; fi35 < 14; fi35++ {
                var v36 uint64 = 18446744073709551615
                if (((v3 + (("hello" + "foo bar") + "a")) + v3) > ((v3 + v33) + ((v33 + (("a" + "\n\t") + v3)) + "hello"))) {
                    v37 := GenS4{F0: 128, F1: 2147483647}
                    var v38 string = ("foo bar" + "\n\t")
                    v5 -= 1
                    _ = v37
                    _ = v38
                }
                for fi39 := 0; fi39 < 17; fi39++ {
                    var v40 float64 = (0.0 / 1.5)
                    v1 += v1
                    _ = fi39
                    _ = v40
                }
                _ = fi35
                _ = v36
            }
        }
    case 127:
        fmt.Println(v2, v5, v3)
    }
    var v41 uint16 = (v34 % ((v34) | 1))
    if (1 <= 100) {
        _ = len(v6)
        var v42 uint64 = 0
        if false {
            _ = len(v7)
            var v43 float32 = 3.14159
            switch ((("hello" + v33) + (v3 + "hello")) + "a") {
            case "foo bar":
                var v44 int32 = v5
                _ = v44
            default:
                v2 = (v2 & (-(v2)))
            }
            _ = v43
        } else {
            v41 = (^(v41))
        }
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
