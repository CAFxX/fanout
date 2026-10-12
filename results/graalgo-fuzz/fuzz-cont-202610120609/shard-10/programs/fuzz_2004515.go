package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int; F1 int; F2 int }
type GenS4 struct { F0 string }
func genFn0(p0 int8, p1 bool) uint32 {
    v4 := []int{}
    var v5 uint8 = 64
    if (127 == 42) {
        var v6 bool = (2 > 0)
        fmt.Println(p1)
        if (v6 || true) {
            p1 = (-2.25 == 0.0)
            for fi7 := 0; fi7 < 20; fi7++ {
                v8 := GenS1{F0: (v5 << ((732) & 63)), F1: 127, F2: -9223372036854775808}
                _ = fi7
                _ = v8
            }
        } else {
            v5 -= v5
            for fi9 := 0; fi9 < 11; fi9++ {
                for fi10 := 0; fi10 < 6; fi10++ {
                    switch (fi10 % ((271) | 1)) {
                    case 1:
                        fmt.Println(v5)
                        var v11 int8 = (p0 - (p0 >> ((255) & 63)))
                        _ = v11
                    }
                    fmt.Println(v5, v6)
                    if (!((1.5 / 1e10) == 3.14159)) {
                        switch fi9 {
                        case -23:
                            var v12 int64 = 1
                            _ = v12
                        default:
                            if v6 {
                                var v13 bool = (1 >= 32767)
                                _ = v13
                            }
                        }
                    }
                    _ = fi10
                }
                _ = fi9
            }
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 int32 = 2147483647
                _ = fi14
                _ = v15
            }
        }
        _ = v6
    }
    var v16 uint64 = 610
    v16 += 1
    _ = p0
    _ = p1
    _ = v16
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1() int {
    var v17 float32 = (0.1 * float32(1.5))
    v18 := GenS2{F0: 255, F1: 61}
    _ = v17
    _ = v18
    return 32768
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{"hello", "foo bar", "a"}
    var v3 int8 = -1
    v3 = (v3 >> ((42) & 63))
    for fi19 := 0; fi19 < 15; fi19++ {
        _ = len(v2)
        v3 += (v3 << ((42) & 63))
        _ = fi19
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 int8 = (-(v3))
        var v22 uint64 = 18446744073709551615
        _ = fi20
        _ = v21
        _ = v22
    }
    if (("hello" + "") != ("\n\t" + (("a" + "foo bar") + ""))) {
        if ((343 != 1) || true) {
            fmt.Println(v3)
            switch 936 {
            case 367:
                v23 := GenS3{F0: -2147483648, F1: 32768, F2: 949}
                _ = v23
            case 32767:
                v1[("hello" + "\n\t")] = "hello"
            default:
                v2 = append(v2, ("" + ("" + "foo bar")))
            }
        } else {
            fmt.Println(v3)
            if true {
                v3 = (v3 / (((v3 + v3)) | 1))
                var v24 int = 131
                var v25 float32 = 3.14159
                _ = v24
                _ = v25
            }
        }
    }
    v3 -= v3
    v26 := GenS2{F0: 128, F1: -364}
    if (42 < 427) {
        var v27 int64 = 1
        v28 := GenS4{F0: "foo bar"}
        _ = v27
        _ = v28
    }
    var v29 int8 = 42
    switch "foo bar" {
    case "":
        var v30 bool = false
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
}
