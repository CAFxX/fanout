package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 int64 }
type GenS8 struct { F0 int16; F1 int32 }

func main() {
    var v1 string = ("foo bar" + "")
    v2 := []uint32{964, 65535, 42}
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    v5 := make(map[string]int)
    v6 := make(map[uint32]uint32)
    v7 := GenS1{F0: 1, F1: -922, F2: 65535}
    v1 = v1
    if (2147483647 == 2147483647) {
        fmt.Println(v1)
        v6[65535] = 42
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS2{F0: 125, F1: fi8}
            fmt.Println(v1)
            _ = fi8
            _ = v9
        }
    } else {
        v1 = "foo bar"
        var v10 float32 = float32((0.0 - -2.25))
        _ = v10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        if (!false) {
            v12 := GenS3{F0: 128, F1: 2}
            if (18446744073709551615 == 760) {
                v1 = ("foo bar" + v1)
                v13 := GenS1{F0: -2147483648, F1: -1, F2: 697}
                v14 := GenS4{F0: 2, F1: 0, F2: 1}
                _ = v13
                _ = v14
            } else {
                for fi15 := 0; fi15 < 10; fi15++ {
                    if (float64(1.5) == -2.25) {
                        if (485 <= 750) {
                            fmt.Println(fi11, fi11)
                        }
                        if (int8(fi11) != -107) {
                            var v16 string = "foo bar"
                            _ = v16
                        }
                        var v17 bool = (469 < 230)
                        _ = v17
                    }
                    v1 = v1
                    fmt.Println(fi15, fi11, fi15)
                    var v18 float32 = -2.25
                    var v19 uint64 = 308
                    _ = fi15
                    _ = v18
                    _ = v19
                }
                var v20 string = "foo bar"
                _ = v20
            }
            switch v1 {
            case "":
                v1 = v1
            default:
                var v21 uint32 = 602
                _ = v21
            }
            _ = v12
        }
        v22 := GenS5{F0: (((v1 + "a") + (v1 + v1)) + (("\n\t" + "a") + "a"))}
        for fi23 := 0; fi23 < 18; fi23++ {
            var v24 uint = 18446744073709551615
            var v25 float32 = -2.25
            v26 := v3[fi23]
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi11
        _ = v22
    }
    if false {
        v1 = ((((((v1 + (v1 + (v1 + ("\n\t" + (v1 + v1))))) + ("" + (("foo bar" + (v1 + ((v1 + "\n\t") + (v1 + (v1 + (((v1 + "") + "hello") + "")))))) + (v1 + "foo bar")))) + v1) + v1) + ("hello" + "a")) + v1)
        v27 := GenS5{F0: (v1 + (v1 + (("hello" + v1) + v1)))}
        v1 = ""
        _ = v27
    }
    v28 := GenS6{F0: 2}
    var v29 bool = true
    _ = len(v2)
    if true {
        var v30 uint64 = 4294967295
        v31 := GenS7{F0: (2 == 2), F1: 9223372036854775807}
        _ = v30
        _ = v31
    } else {
        v29 = (v29 && v29)
        _ = len(v2)
        var v32 int8 = 100
        _ = v32
    }
    fmt.Println(v1, v1)
    var v33 int = 207
    v34 := GenS4{F0: 255, F1: -1, F2: 0}
    var v35 int64 = -860
    switch v1 {
    case "a":
        if true {
            delete(v4, 239)
        }
    }
    for fi36 := 0; fi36 < 6; fi36++ {
        for fi37 := 0; fi37 < 10; fi37++ {
            var v38 float64 = float64(3.14159)
            _ = fi37
            _ = v38
        }
        for fi39 := 0; fi39 < 15; fi39++ {
            v29 = v29
            var v40 uint8 = 145
            fmt.Println(fi36, v35)
            _ = fi39
            _ = v40
        }
        var v41 float64 = 0.0
        _ = fi36
        _ = v41
    }
    if true {
        switch (v33 << ((1) & 63)) {
        case 615:
            var v42 string = ((("" + v1) + v1) + v1)
            _ = v42
        default:
            var v43 bool = v29
            _ = v43
        }
    } else {
        v44 := GenS8{F0: -32768, F1: 586}
        var v45 int16 = 0
        switch v33 {
        case 127:
            v29 = v29
        case 42:
            v29 = true
        }
        _ = v44
        _ = v45
    }
    fmt.Println(v35, v29)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
