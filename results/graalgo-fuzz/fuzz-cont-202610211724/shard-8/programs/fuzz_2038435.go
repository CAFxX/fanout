package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int16; F1 bool; F2 uint8 }

func main() {
    v1 := []int{503, -889, -9223372036854775808, -390}
    v2 := make(map[string]bool)
    v3 := []uint32{965}
    var v4 uint64 = 57
    v5 := make(map[uint32]string)
    fmt.Println(v4)
    for fi6 := 0; fi6 < 16; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            v4 = (v4 >> ((4294967295) & 63))
            var v8 int = fi7
            v8 -= (fi6 / ((fi6) | 1))
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v9 int32 = 942
    v4 += v4
    switch (("foo bar" + "\n\t") + "a") {
    case "foo bar":
        if ((!false) && (!(!((18446744073709551615 != 5) && (!(!(false && ((65535 != 65535) && false)))))))) {
            for fi10 := 0; fi10 < 8; fi10++ {
                v11 := GenS1{F0: 1000, F1: uint32(v4)}
                _ = fi10
                _ = v11
            }
            var v12 float64 = (3.14159 / 3.14159)
            if (4294967295 > 0) {
                if true {
                    switch 2147483647 {
                    case 127:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case -2147483648:
                        if true {
                            switch ("\n\t" + "\n\t") {
                            case "hello":
                                if true {
                                    var v13 int8 = 59
                                    v14 := GenS1{F0: 0, F1: 2147483648}
                                    if (4294967295 < v4) {
                                        _ = len(v5)
                                    }
                                    _ = v13
                                    _ = v14
                                } else {
                                    v12 = (1e10 + (v12 * 0.1))
                                }
                            case "\n\t":
                                v15 := GenS1{F0: -364, F1: 4294967295}
                                _ = v15
                            }
                        } else {
                            if (-90 <= v9) {
                                v16 := GenS1{F0: 0, F1: 2147483647}
                                _ = v16
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            v17 := GenS2{F0: 111, F1: 42}
                            _ = v17
                        }
                    case 9223372036854775807:
                        v12 -= (v12 - v12)
                    }
                    for fi18 := 0; fi18 < 15; fi18++ {
                        fmt.Println(fi18, v12, v4)
                        _ = fi18
                    }
                }
                v19 := v2[((("\n\t" + ("foo bar" + "\n\t")) + (("hello" + "foo bar") + "hello")) + "")]
                v12 = (1e10 - (float64(1.5) - v12))
                _ = v19
            }
            _ = v12
        } else {
            for fi20 := 0; fi20 < 5; fi20++ {
                v4 = v4
                v4 += v4
                v4 = (v4 - (v4 | v4))
                _ = fi20
            }
            switch "" {
            case "\n\t":
                var v21 int64 = -1
                _ = v21
            case "foo bar":
                var v22 string = ((("hello" + "a") + (((("a" + "") + "") + "\n\t") + ("a" + ("" + "foo bar")))) + "a")
                _ = v22
            case "":
                var v23 uint32 = 0
                _ = v23
            }
        }
    case "hello":
        var v24 float32 = float32(1.5)
        _ = v24
    case "a":
        var v25 uint16 = 264
        _ = v25
    }
    v9 += v9
    v26 := GenS2{F0: 87, F1: 32767}
    for fi27 := 0; fi27 < 7; fi27++ {
        var v28 int8 = 1
        var v29 float32 = float32((float64(0.0) + 1e10))
        switch ("\n\t" + "\n\t") {
        case "":
            if true {
                if ((9223372036854775807 == (fi27 ^ (fi27 - fi27))) || true) {
                    var v30 float32 = float32(0.0)
                    var v31 int32 = v9
                    _ = v30
                    _ = v31
                }
            }
        default:
            fmt.Println(v9)
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    var v32 uint = 609
    v32 = 2
    for fi33 := 0; fi33 < 19; fi33++ {
        switch 42 {
        case 4294967295:
            switch fi33 {
            case -215:
                var v34 uint32 = 763
                _ = v34
            case -2147483648:
                fmt.Println(v9)
            case -227:
                fmt.Println(v4)
            default:
                delete(v2, "")
            }
        default:
            v35 := GenS3{F0: -340, F1: (v9 != (v9 ^ 1)), F2: 212}
            _ = v35
        }
        _ = fi33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v9
}
