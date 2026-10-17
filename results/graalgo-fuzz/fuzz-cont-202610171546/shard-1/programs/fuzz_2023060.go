package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8; F1 int8 }

func main() {
    var v1 int32 = -459
    v2 := make(map[string]string)
    var v3 int64 = -1
    var v4 uint32 = 0
    switch (("hello" + (("\n\t" + ("foo bar" + ("a" + "\n\t"))) + ("" + "hello"))) + "\n\t") {
    case "foo bar":
        v3 = v3
    default:
        if (373 <= -265) {
            v5 := GenS1{F0: 901}
            delete(v2, "foo bar")
            if (!(32767 > 1)) {
                if true {
                    v6 := GenS1{F0: 65535}
                    v7 := GenS2{F0: v1}
                    fmt.Println(v3, v1)
                    _ = v6
                    _ = v7
                }
                if (42 != 0) {
                    var v8 int8 = -82
                    var v9 int = 128
                    _ = v8
                    _ = v9
                }
                v3 = v3
            }
            _ = v5
        }
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := GenS1{F0: 1}
        _ = fi10
        _ = v11
    }
    var v12 int32 = (v1 * (^(v1)))
    if false {
        v13 := GenS1{F0: 98}
        v3 -= 662
        _ = v13
    } else {
        v14 := GenS3{F0: -128, F1: 65535}
        _ = v14
    }
    fmt.Println(v12, v3)
    var v15 uint64 = 0
    if ((-1 <= 0) || true) {
        v16 := GenS4{F0: (v3 * v3)}
        fmt.Println(v1, v15)
        if true {
            for fi17 := 0; fi17 < 10; fi17++ {
                v18 := GenS2{F0: (^(v1))}
                _ = fi17
                _ = v18
            }
        } else {
            v12 -= 2147483647
        }
        _ = v16
    } else {
        if (false || (v4 == v4)) {
            v3 = -9223372036854775808
            fmt.Println(v4, v15, v15)
        }
    }
    if (!(((((1.5 / (1e10 - ((0.0 + (3.14159 - 0.0)) + 1.5))) / 1.0) - 3.14159) / 0.1) != float64(3.14159))) {
        var v19 uint32 = (v4 % (((v4 + (v4 | (^(v4))))) | 1))
        fmt.Println(v3, v1)
        fmt.Println(v4, v19)
        _ = v19
    }
    if true {
        fmt.Println(v3, v1)
        if (1e10 <= ((0.0 * 1e10) + 1.5)) {
            v20 := v2[("foo bar" + "hello")]
            var v21 uint8 = 127
            _ = v20
            _ = v21
        } else {
            v22 := GenS2{F0: -549}
            v23 := GenS5{F0: (((0.0 + 1e10) > -2.25) && false)}
            if (162 < 42) {
                var v24 float32 = (1e10 * ((1.5 / 0.1) * 1.5))
                for fi25 := 0; fi25 < 20; fi25++ {
                    switch (-(fi25)) {
                    case 2147483647:
                        v26 := v2["\n\t"]
                        _ = v26
                    case 128:
                        for fi27 := 0; fi27 < 13; fi27++ {
                            var v28 uint = uint(v15)
                            v15 += 785
                            _ = fi27
                            _ = v28
                        }
                    case 0:
                        var v29 string = ""
                        _ = v29
                    default:
                        switch "hello" {
                        case "a":
                            var v30 uint64 = 2
                            _ = v30
                        }
                    }
                    _ = fi25
                }
                var v31 int16 = 0
                _ = v24
                _ = v31
            }
            _ = v22
            _ = v23
        }
    } else {
        switch 32768 {
        case 707:
            v4 -= 0
        case 1:
            v1 = v1
        case 32768:
            v32 := GenS6{F0: int8(v1), F1: 0}
            _ = v32
        }
    }
    delete(v2, "")
    v33 := GenS1{F0: 0}
    fmt.Println(v3)
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
