package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int; F2 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 uint32 }
func genFn0(p0 string) int32 {
    if (1 >= 1) {
        var v4 uint = 1
        _ = v4
    }
    v5 := GenS1{F0: -1}
    _ = p0
    _ = v5
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{("foo bar" + "a"), "a", "", ("" + (("a" + "") + ("\n\t" + "\n\t")))}
    v3 := make(map[string]int)
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS2{F0: 595, F1: 1}
        _ = fi6
        _ = v7
    }
    v2 = append(v2, (("\n\t" + "foo bar") + "a"))
    for fi8 := 0; fi8 < 8; fi8++ {
        v3[""] = (fi8 + fi8)
        if (!(!(-459 == 32767))) {
            v9 := GenS3{F0: 1, F1: (fi8 & -429), F2: 127}
            _ = v9
        } else {
            switch ("foo bar" + "hello") {
            case "foo bar":
                for fi10 := 0; fi10 < 10; fi10++ {
                    var v11 uint8 = 128
                    _ = fi10
                    _ = v11
                }
            case "hello":
                v12 := GenS1{F0: 1}
                _ = v12
            case "\n\t":
                if true {
                    v13 := GenS4{F0: 42}
                    if false {
                        if (0.0 < 0.1) {
                            var v14 float32 = (((-2.25 + (((float32(3.14159) - ((1e10 + 0.0) * 1.5)) + -2.25) * 3.14159)) + float32(1.5)) - (1.5 - ((float32(3.14159) + 0.1) * 0.0)))
                            fmt.Println(v14, fi8)
                            _ = v14
                        } else {
                            fmt.Println(fi8)
                            for fi15 := 0; fi15 < 1; fi15++ {
                                var v16 int16 = -648
                                var v17 uint8 = 1
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                        }
                        fmt.Println(fi8)
                    } else {
                        var v18 uint = 568
                        var v19 uint16 = 663
                        var v20 uint64 = 0
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    _ = v13
                }
            }
        }
        _ = fi8
    }
    switch ("foo bar" + "") {
    case "\n\t":
        var v21 int64 = 1
        _ = v21
    default:
        if false {
            var v22 float64 = -2.25
            delete(v3, "")
            if false {
                if true {
                    v22 = v22
                    if true {
                        v22 += (v22 - ((1.5 * (v22 * -2.25)) - (-2.25 + v22)))
                    } else {
                        v23 := v1[2]
                        _ = v23
                    }
                }
                v22 = 1.5
            } else {
                var v24 float64 = v22
                v24 = 1.5
                v24 -= v22
                _ = v24
            }
            _ = v22
        }
    }
    var v25 int32 = -336
    v26 := v1[2]
    _ = len(v1)
    v27 := GenS5{F0: genFn0((v26 + "foo bar")), F1: 2147483648}
    _ = len(v2)
    fmt.Println(v25, v25)
    if false {
        var v28 uint64 = 0
        if false {
            var v29 int64 = 9223372036854775807
            _ = v29
        }
        _ = v28
    }
    var v30 int = 2147483647
    var v31 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
}
