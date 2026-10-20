package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int; F2 uint16 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 int16; F1 int8; F2 int8 }
type GenS5 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    var v3 uint16 = 255
    switch "a" {
    case "":
        v3 = 65535
    }
    fmt.Println(v3)
    if (true || (676 == -504)) {
        for fi4 := 0; fi4 < 16; fi4++ {
            switch "hello" {
            case "":
                v5 := GenS1{F0: -1}
                _ = v5
            default:
                v3 = v3
            }
            _ = fi4
        }
        v3 = 0
    }
    _ = len(v2)
    var v6 bool = true
    var v7 bool = (v6 && true)
    v8 := GenS2{F0: (v3 << ((0) & 63)), F1: 4294967295, F2: v3}
    switch 4294967295 {
    case 9223372036854775807:
        _ = len(v1)
    case -296:
        for fi9 := 0; fi9 < 9; fi9++ {
            delete(v1, fi9)
            v7 = (-633 != 0)
            _ = fi9
        }
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        if false {
            v11 := GenS3{F0: (((("" + "foo bar") + "hello") + "\n\t") + "\n\t"), F1: 18446744073709551615}
            _ = v11
        } else {
            var v12 int32 = 528
            fmt.Println(v3, v7)
            _ = v12
        }
        v13 := GenS3{F0: "foo bar", F1: 105}
        for fi14 := 0; fi14 < 2; fi14++ {
            if (0 == 848) {
                v2[(fi14 * 1)] = 0
            } else {
                v15 := GenS2{F0: (v3 * (v3 << ((1) & 63))), F1: (fi10 << ((2) & 63)), F2: v3}
                switch 481 {
                case -9223372036854775808:
                    switch fi10 {
                    case -333:
                        v16 := v1[fi10]
                        _ = v16
                    case 9223372036854775807:
                        var v17 float32 = float32(1.5)
                        _ = v17
                    case -975:
                        v7 = (620 >= v3)
                    default:
                        v18 := GenS4{F0: 42, F1: 127, F2: -128}
                        _ = v18
                    }
                    fmt.Println(v7)
                    if v7 {
                        var v19 uint16 = v3
                        v19 = v3
                        var v20 int32 = -133
                        fmt.Println(v19)
                        fmt.Println(fi10)
                        var v21 string = "\n\t"
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    fmt.Println(fi14, v6)
                    if (248 == 18446744073709551615) {
                        var v22 int8 = 127
                        v7 = (0 < 40)
                        _ = v22
                    } else {
                        if v6 {
                            fmt.Println(v7, v6)
                            if true {
                                fmt.Println(fi14)
                            }
                        } else {
                            v1[(fi10 << ((987) & 63))] = (1 >= 1)
                            var v23 uint = 65535
                            v7 = (!(false && true))
                            _ = v23
                        }
                        v24 := v1[128]
                        v25 := GenS5{F0: -72}
                        _ = v24
                        _ = v25
                    }
                }
                delete(v2, (fi14 >> ((538) & 63)))
                _ = v15
            }
            _ = fi14
        }
        _ = fi10
        _ = v13
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        fmt.Println(fi26)
        if false {
            var v27 uint64 = 2
            v27 -= 0
            for fi28 := 0; fi28 < 5; fi28++ {
                v6 = v7
                _ = fi28
            }
            _ = v27
        }
        _ = fi26
    }
    v6 = ((false || true) || (!((80 >= 127) && false)))
    v6 = (!(0 <= 1000))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
