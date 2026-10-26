package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 string; F1 int64; F2 uint32 }
type GenS5 struct { F0 int; F1 uint64 }

func main() {
    v1 := []string{"a", ("\n\t" + "hello")}
    var v2 int32 = -2147483648
    v3 := make(map[int]bool)
    v4 := []string{("foo bar" + "\n\t"), "\n\t"}
    var v5 int64 = 1
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: 65535, F1: 2, F2: 65535}
        switch "hello" {
        case "a":
            switch 9223372036854775807 {
            case 32768:
                var v8 uint16 = 42
                _ = v8
            case 2147483647:
                var v9 uint16 = 627
                _ = v9
            default:
                v2 = v2
            }
        }
        if false {
            v3[(^(fi6))] = false
        } else {
            var v10 uint64 = 17
            var v11 uint64 = (v10 - (^(v10)))
            _ = v10
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    v5 -= v5
    v12 := GenS2{F0: false, F1: 1}
    fmt.Println(v5, v2)
    _ = len(v1)
    var v13 int64 = (v5 | -1)
    if (!true) {
        var v14 uint = 315
        for fi15 := 0; fi15 < 17; fi15++ {
            _ = len(v4)
            switch "\n\t" {
            case "\n\t":
                v16 := GenS1{F0: 65535, F1: 869, F2: uint32(v14)}
                _ = v16
            }
            _ = fi15
        }
        _ = v14
    } else {
        v13 = v5
        v17 := GenS2{F0: ((!(!true)) && (1 >= 104)), F1: 90}
        _ = v17
    }
    v18 := GenS1{F0: 633, F1: 175, F2: 42}
    v13 = 42
    var v19 int64 = -1
    var v20 uint32 = 2147483648
    var v21 int8 = 1
    v22 := GenS3{F0: (127 >= 2), F1: v2}
    for fi23 := 0; fi23 < 10; fi23++ {
        switch "foo bar" {
        case "":
            v3[fi23] = true
        default:
            v20 = 42
        }
        v24 := v3[fi23]
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        if (v2 >= -1) {
            delete(v3, fi25)
            if (false && false) {
                for fi26 := 0; fi26 < 3; fi26++ {
                    if (!false) {
                        switch ("" + "") {
                        case "\n\t":
                            v27 := GenS4{F0: ("\n\t" + "\n\t"), F1: (^(v13)), F2: 42}
                            _ = v27
                        case "":
                            fmt.Println(fi26)
                        case "hello":
                            v20 = 65535
                        }
                        if ((!(!(!(true && (3.14159 == 0.0))))) || true) {
                            v19 -= v19
                            if false {
                                fmt.Println(v21, fi25, fi25)
                                fmt.Println(v20, v21, v21)
                                v28 := GenS3{F0: false, F1: int32(v21)}
                                if len(v1) > 0 {
                                    v1[0] = ("hello" + "\n\t")
                                }
                                fmt.Println(v19, v2, v20)
                                _ = v28
                            } else {
                                var v29 int32 = int32(v21)
                                var v30 uint64 = 792
                                if (1 <= 4294967295) {
                                    var v31 uint64 = (v30 - v30)
                                    v32 := GenS5{F0: (fi25 | (fi25 / ((fi25) | 1))), F1: 935}
                                    fmt.Println(v2)
                                    _ = v31
                                    _ = v32
                                }
                                _ = v29
                                _ = v30
                            }
                        } else {
                            var v33 int = fi25
                            fmt.Println(v20)
                            _ = v33
                        }
                    }
                    _ = fi26
                }
                fmt.Println(v20)
                v5 -= v5
            } else {
                v19 -= (-(v5))
            }
            var v34 int16 = 1
            _ = v34
        }
        v21 = (v21 - (v21 * v21))
        _ = len(v3)
        _ = fi25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
