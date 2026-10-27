package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int64 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint32, p1 uint32) int32 {
    var v4 string = "foo bar"
    v5 := []int{}
    if (false && false) {
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := GenS1{F0: 42, F1: 1, F2: -9223372036854775808}
            p0 = 2147483648
            _ = fi6
            _ = v7
        }
        switch 84 {
        case -2147483648:
            fmt.Println(p0)
        case -1:
            var v8 uint8 = 127
            _ = v8
        }
        if len(v5) > 0 {
            v5[0] = 128
        }
    } else {
        var v9 int16 = 42
        if (v9 < v9) {
            if len(v5) > 0 {
                v5[0] = 32768
            }
            v9 -= v9
        }
        _ = v9
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v10 bool = (!true)
    _ = p0
    _ = p1
    _ = v10
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := []string{("a" + "a")}
    var v2 uint = 255
    var v3 uint32 = 65535
    v1 = append(v1, "a")
    v3 += v3
    if true {
        _ = len(v1)
        if false {
            var v11 uint16 = 255
            var v12 float64 = float64(1.5)
            _ = v11
            _ = v12
        }
    } else {
        switch 9223372036854775807 {
        case -9223372036854775808:
            v13 := GenS2{F0: (v3 / ((42) | 1)), F1: ((!(0 == 1)) && (!true))}
            _ = v13
        case 1:
            fmt.Println(v3, v2)
        case 25:
            v14 := GenS1{F0: 0, F1: 914, F2: -747}
            _ = v14
        default:
            var v15 int16 = 42
            _ = v15
        }
    }
    v2 -= 1
    v1 = append(v1, "hello")
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 string = ((("" + "") + "") + ("hello" + "a"))
        var v18 int64 = -1
        if (false && false) {
            fmt.Println(v17)
        } else {
            if ((fi16 & (fi16 << ((v2) & 63))) > fi16) {
                switch fi16 {
                case 0:
                    for fi19 := 0; fi19 < 17; fi19++ {
                        var v20 bool = (!(0 != 505))
                        v21 := GenS3{F0: (!(!true)), F1: 496}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                default:
                    var v22 uint = 4294967295
                    _ = v22
                }
                switch (("" + (("hello" + ((v17 + v17) + v17)) + v17)) + ("\n\t" + v17)) {
                case "":
                    if len(v1) > 0 {
                        v1[0] = v17
                    }
                }
            } else {
                var v23 int64 = v18
                fmt.Println(fi16, v2, fi16)
                _ = v23
            }
            var v24 float32 = 1.5
            if ((v2 << (((v2 >> ((v2) & 63))) & 63)) != v2) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                var v25 int = (fi16 / ((int(v18)) | 1))
                _ = v25
            }
            _ = v24
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    var v26 uint32 = uint32(v2)
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    fmt.Println(v3, v26)
    var v27 uint32 = 65535
    v27 -= 4294967295
    switch "" {
    case "foo bar":
        for fi28 := 0; fi28 < 6; fi28++ {
            for fi29 := 0; fi29 < 17; fi29++ {
                if false {
                    var v30 uint64 = 2
                    v31 := GenS3{F0: (false && false), F1: v30}
                    _ = v30
                    _ = v31
                } else {
                    var v32 uint32 = v3
                    var v33 int32 = 42
                    _ = v32
                    _ = v33
                }
                if (v2 > v2) {
                    switch 101 {
                    case -76:
                        fmt.Println(fi29, v26)
                    }
                }
                _ = fi29
            }
            _ = fi28
        }
    default:
        switch "foo bar" {
        case "hello":
            v34 := GenS1{F0: (v2 - (v2 * (v2 ^ v2))), F1: 0, F2: 9223372036854775807}
            _ = v34
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
