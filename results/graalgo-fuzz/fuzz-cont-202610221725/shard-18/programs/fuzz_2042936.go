package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
func genFn0() int32 {
    var v5 int8 = -26
    var v6 float64 = 0.0
    if false {
        v5 += (v5 ^ -128)
        if false {
            v7 := GenS1{F0: -1, F1: false}
            _ = v7
        } else {
            var v8 int32 = 417
            _ = v8
        }
        v9 := []int{4294967295, 280, 128, 4294967295}
        _ = v9
    }
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int64 = -1
    var v2 uint32 = 4294967295
    var v3 uint8 = 2
    v4 := make(map[string]uint32)
    for fi10 := 0; fi10 < 5; fi10++ {
        v2 += v2
        _ = fi10
    }
    v11 := []int{32767}
    if true {
        v12 := GenS1{F0: 4294967295, F1: ((true && false) && ((v1 - -1) != 0))}
        v3 += 2
        if (0 < -900) {
            v2 += uint32(v3)
        } else {
            switch 2147483647 {
            case -1:
                var v13 int = 2147483647
                _ = v13
            case -219:
                for fi14 := 0; fi14 < 9; fi14++ {
                    switch (fi14 / ((533) | 1)) {
                    case 42:
                        var v15 int = fi14
                        _ = v15
                    case 32767:
                        var v16 uint8 = 190
                        _ = v16
                    default:
                        var v17 float32 = 0.1
                        _ = v17
                    }
                    v18 := GenS1{F0: fi14, F1: (int16(fi14) > -858)}
                    if (true || (false || (669 < -491))) {
                        v19 := GenS1{F0: fi14, F1: false}
                        _ = v19
                    } else {
                        v1 += -542
                        switch "" {
                        case "foo bar":
                            var v20 uint16 = 0
                            _ = v20
                        }
                        if false {
                            v1 += v1
                        }
                    }
                    _ = fi14
                    _ = v18
                }
            case -28:
                if true {
                    v3 = (v3 % (((v3 & v3)) | 1))
                    fmt.Println(v3)
                    if (!(false && false)) {
                        if len(v11) > 0 {
                            v11[0] = 521
                        }
                        var v21 int64 = 1
                        switch ("" + ((("" + ("hello" + "a")) + "hello") + ("a" + ("\n\t" + ("foo bar" + "hello"))))) {
                        case "a":
                            v21 = v1
                        }
                        _ = v21
                    } else {
                        for fi22 := 0; fi22 < 4; fi22++ {
                            v2 = 58
                            v2 = v2
                            v23 := GenS1{F0: 997, F1: false}
                            _ = fi22
                            _ = v23
                        }
                        v1 = -1
                    }
                } else {
                    var v24 uint = 1
                    _ = len(v11)
                    _ = len(v4)
                    _ = v24
                }
            default:
                v25 := GenS1{F0: 9223372036854775807, F1: true}
                _ = v25
            }
            v3 = uint8(v2)
        }
        _ = v12
    } else {
        v1 = (v1 * 42)
        for fi26 := 0; fi26 < 18; fi26++ {
            var v27 int16 = 0
            v3 = 80
            _ = fi26
            _ = v27
        }
        v28 := GenS1{F0: 9223372036854775807, F1: true}
        _ = v28
    }
    fmt.Println(v2, v1, v2)
    v1 = (v1 & v1)
    var v29 uint32 = v2
    v11 = append(v11, 32767)
    v1 += (v1 * -771)
    delete(v4, "hello")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
