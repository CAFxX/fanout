package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint }
func genFn0(p0 int, p1 int16, p2 string) uint {
    v7 := GenS1{F0: (0 == 147)}
    var v8 int32 = 2147483647
    for fi9 := 0; fi9 < 4; fi9++ {
        p1 = (p1 >> ((883) & 63))
        v8 = 0
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := []uint32{}
    v2 := make(map[string]uint32)
    v3 := []string{}
    var v4 uint8 = 127
    v5 := make(map[string]uint32)
    var v6 float64 = 3.14159
    _ = len(v3)
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 int32 = 0
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        if false {
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 bool = (((0.0 - 3.14159) * float32(v6)) >= (3.14159 * (1e10 + 0.0)))
                _ = fi13
                _ = v14
            }
            switch (fi12 & 9223372036854775807) {
            case 4294967295:
                fmt.Println(v6)
            case 157:
                switch ("hello" + ((("\n\t" + "\n\t") + "foo bar") + "a")) {
                case "foo bar":
                    var v15 uint64 = 2
                    _ = v15
                case "a":
                    var v16 int8 = 1
                    _ = v16
                }
            case 32768:
                v17 := GenS1{F0: false}
                _ = v17
            default:
                v4 = (v4 | (v4 * (^(v4))))
            }
            for fi18 := 0; fi18 < 8; fi18++ {
                v4 += v4
                var v19 uint16 = 42
                switch "\n\t" {
                case "a":
                    var v20 uint8 = v4
                    _ = v20
                }
                v21 := GenS2{F0: 886, F1: 0, F2: 2}
                _ = fi18
                _ = v19
                _ = v21
            }
        } else {
            var v22 uint8 = (^(v4))
            var v23 int = (fi12 >> ((genFn0(fi12, 32767, (((("foo bar" + "\n\t") + "foo bar") + ((("hello" + "foo bar") + (("hello" + "foo bar") + "\n\t")) + ("foo bar" + "hello"))) + ("\n\t" + "a")))) & 63))
            for fi24 := 0; fi24 < 1; fi24++ {
                var v25 bool = false
                _ = fi24
                _ = v25
            }
            _ = v22
            _ = v23
        }
        v26 := GenS2{F0: 415, F1: 1, F2: v4}
        _ = fi12
        _ = v26
    }
    var v27 uint32 = 2147483647
    var v28 float32 = float32(1e10)
    var v29 string = "\n\t"
    for fi30 := 0; fi30 < 10; fi30++ {
        fmt.Println(v27, v27, v29)
        if false {
            fmt.Println(v6, v29)
            var v31 uint8 = (v4 - (^(v4)))
            _ = v31
        } else {
            v6 = ((0.0 + v6) - float64((((0.0 + v28) + v28) / v28)))
            v32 := GenS2{F0: 4294967295, F1: v27, F2: (v4 % (((v4 - 255)) | 1))}
            _ = v32
        }
        _ = fi30
    }
    switch 1 {
    case -823:
        for fi33 := 0; fi33 < 18; fi33++ {
            v29 = v29
            v6 = (0.0 * (float64(v28) / v6))
            v34 := GenS3{F0: v27, F1: v4, F2: 4294967295}
            _ = fi33
            _ = v34
        }
    default:
        fmt.Println(v27, v4)
    }
    if false {
        if len(v3) > 0 {
            v3[0] = ((("hello" + v29) + v29) + v29)
        }
    } else {
        if ((((3.14159 - v6) * (-2.25 / 0.1)) * float64(v28)) != v6) {
            var v35 bool = true
            if (1 < 880) {
                var v36 uint8 = 42
                _ = v36
            } else {
                delete(v2, ((("hello" + v29) + ("a" + "\n\t")) + (v29 + "a")))
                switch (v29 + (v29 + "hello")) {
                case "a":
                    if ((v6 >= 0.1) && (!(449 > 9223372036854775807))) {
                        var v37 float32 = v28
                        v5[(v29 + (v29 + "foo bar"))] = v27
                        v35 = ((!(!v35)) || v35)
                        _ = v37
                    }
                }
            }
            _ = v35
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
