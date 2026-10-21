package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 string; F2 string }
func genFn0() uint64 {
    var v4 uint8 = 228
    for fi5 := 0; fi5 < 11; fi5++ {
        switch fi5 {
        case 42:
            var v6 int64 = -9223372036854775808
            _ = v6
        case 0:
            for fi7 := 0; fi7 < 10; fi7++ {
                fmt.Println(fi5, v4)
                v4 += 30
                _ = fi7
            }
        case 128:
            v4 = v4
        }
        _ = fi5
    }
    _ = v4
    return 18446744073709551615
}

func main() {
    v1 := []string{("foo bar" + (("a" + ("\n\t" + ("hello" + "hello"))) + ("foo bar" + "foo bar"))), "\n\t", "a", ""}
    var v2 uint8 = 0
    v3 := make(map[uint32]string)
    v2 -= (v2 >> ((2) & 63))
    if false {
        v2 = v2
    }
    fmt.Println(v2)
    fmt.Println(v2)
    fmt.Println(v2)
    var v8 int16 = -32768
    v8 = (v8 / ((v8) | 1))
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    v9 := GenS1{F0: (v8 | -628), F1: 18446744073709551615}
    var v10 uint = 93
    v11 := GenS2{F0: 412}
    var v12 uint16 = 2
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := GenS1{F0: (-(v8)), F1: 1}
        _ = fi13
        _ = v14
    }
    v10 = v10
    v2 -= 110
    if true {
        if (((v2 + v2) < (v2 / (((v2 / ((v2) | 1))) | 1))) || false) {
            var v15 uint16 = (v12 / ((2) | 1))
            for fi16 := 0; fi16 < 12; fi16++ {
                v17 := GenS3{F0: (genFn0() >= 2)}
                if (!((1e10 * (0.1 - 1e10)) < (float32(-2.25) / 1.0))) {
                    var v18 int16 = v8
                    _ = v18
                } else {
                    v3[24] = ("\n\t" + ("\n\t" + ("" + (("a" + "a") + "foo bar"))))
                    if (!(108 < 42)) {
                        switch (^(fi16)) {
                        case 4294967295:
                            v10 += v10
                        case -1:
                            v12 = (v15 >> ((v10) & 63))
                        case -107:
                            v15 = 0
                        }
                    }
                    v3[4294967295] = ""
                }
                v12 += (v15 - v12)
                _ = fi16
                _ = v17
            }
            var v19 uint8 = v2
            _ = v15
            _ = v19
        }
        switch 4294967295 {
        case 194:
            var v20 uint16 = v12
            _ = v20
        case 521:
            v2 = v2
        case 967:
            switch 42 {
            case 42:
                for fi21 := 0; fi21 < 10; fi21++ {
                    v22 := GenS2{F0: (v12 / ((110) | 1))}
                    if true {
                        var v23 int8 = -128
                        _ = v23
                    }
                    _ = fi21
                    _ = v22
                }
            }
        default:
            var v24 uint64 = 42
            _ = v24
        }
        v2 -= v2
    } else {
        for fi25 := 0; fi25 < 19; fi25++ {
            v26 := GenS4{F0: "", F1: "\n\t", F2: ""}
            var v27 string = ""
            _ = fi25
            _ = v26
            _ = v27
        }
        if (v8 == v8) {
            v10 = 686
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        } else {
            v10 -= (v10 & uint(v12))
        }
        v10 += (v10 | 692)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
