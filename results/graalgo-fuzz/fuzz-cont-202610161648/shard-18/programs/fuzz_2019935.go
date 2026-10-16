package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int64; F1 int16; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 uint16; F1 int32; F2 int }
type GenS8 struct { F0 string }

func main() {
    var v1 int = -1
    v2 := make(map[string]string)
    var v3 string = "a"
    var v4 uint32 = 1
    v5 := make(map[string]string)
    var v6 int32 = 889
    fmt.Println(v1, v3)
    if ((false && true) && (9223372036854775807 == (v1 / (((-(v1))) | 1)))) {
        v6 = (^(v6))
        v6 += v6
        if (uint64(v4) < 2) {
            var v7 int64 = 42
            var v8 int8 = -7
            var v9 int8 = (v8 * v8)
            _ = v7
            _ = v8
            _ = v9
        }
    } else {
        var v10 uint64 = 4294967295
        for fi11 := 0; fi11 < 14; fi11++ {
            fmt.Println(v10, v10)
            v4 += v4
            _ = fi11
        }
        var v12 float32 = 1e10
        _ = v10
        _ = v12
    }
    switch 42 {
    case -1:
        v13 := GenS1{F0: v4, F1: v4}
        _ = v13
    case -102:
        v1 = 1
    }
    if (!(!false)) {
        v1 += 32767
    } else {
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 bool = false
            _ = fi14
            _ = v15
        }
        var v16 int = (v1 * (v1 | (v1 >> ((18446744073709551615) & 63))))
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 int8 = 42
        _ = fi17
        _ = v18
    }
    if (!((v6 > 574) && true)) {
        if false {
            v4 = 1
            if ((v6 << ((65535) & 63)) <= (-(v6))) {
                for fi19 := 0; fi19 < 3; fi19++ {
                    v3 = v3
                    fmt.Println(v4, v3)
                    _ = fi19
                }
                v20 := GenS2{F0: 9223372036854775807, F1: -1, F2: 346}
                switch v3 {
                case "foo bar":
                    var v21 int = (-(v1))
                    _ = v21
                case "":
                    var v22 string = (v3 + "hello")
                    _ = v22
                case "a":
                    v23 := []int{(v1 + (v1 - v1)), v1, v1}
                    _ = v23
                default:
                    var v24 bool = false
                    _ = v24
                }
                _ = v20
            } else {
                for fi25 := 0; fi25 < 12; fi25++ {
                    v26 := GenS1{F0: (v4 ^ (v4 ^ v4)), F1: v4}
                    switch (fi25 % (((fi25 | (fi25 + fi25))) | 1)) {
                    case 4294967295:
                        v4 = 309
                    case -957:
                        var v27 string = v3
                        _ = v27
                    }
                    _ = fi25
                    _ = v26
                }
                if true {
                    v28 := []int{(v1 + (v1 ^ v1)), v1, v1}
                    v29 := GenS3{F0: 18, F1: 255}
                    _ = v28
                    _ = v29
                }
                var v30 float32 = 1.5
                _ = v30
            }
            v31 := GenS1{F0: (^(v4)), F1: 0}
            _ = v31
        }
        v1 += (v1 << ((209) & 63))
        if (0 != 2) {
            v32 := GenS4{F0: 1}
            _ = v32
        }
    }
    v33 := GenS5{F0: 999, F1: -128}
    if (1 > -1) {
        v6 = v6
        for fi34 := 0; fi34 < 16; fi34++ {
            switch fi34 {
            case 571:
                var v35 float64 = float64(3.14159)
                _ = v35
            case 128:
                switch ("a" + (v3 + (v3 + "foo bar"))) {
                case "foo bar":
                    if ((868 == 0) && false) {
                        v1 += -9223372036854775808
                        fmt.Println(v1)
                        var v36 uint8 = 42
                        _ = v36
                    }
                default:
                    v37 := GenS2{F0: 42, F1: 188, F2: 2}
                    _ = v37
                }
            }
            v38 := v2[v3]
            _ = fi34
            _ = v38
        }
        v39 := v5[((v3 + "a") + ("\n\t" + v3))]
        _ = v39
    }
    switch v1 {
    case 164:
        for fi40 := 0; fi40 < 5; fi40++ {
            var v41 uint16 = 274
            _ = fi40
            _ = v41
        }
    case -9223372036854775808:
        v6 = v6
    }
    fmt.Println(v6, v1)
    var v42 uint16 = 42
    switch v3 {
    case "foo bar":
        var v43 int32 = -267
        _ = v43
    case "":
        v44 := GenS6{F0: 0, F1: (v6 | (v6 % ((-267) | 1)))}
        _ = v44
    }
    switch "a" {
    case "a":
        var v45 int32 = v6
        _ = v45
    case "":
        var v46 uint64 = 843
        _ = v46
    case "foo bar":
        v47 := GenS7{F0: uint16(v4), F1: (v6 % ((v6) | 1)), F2: (v1 & 4294967295)}
        _ = v47
    }
    if (!false) {
        fmt.Println(v6, v1)
    }
    v6 -= 0
    v48 := []int{(v1 & v1), v1}
    v49 := GenS8{F0: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v42
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
