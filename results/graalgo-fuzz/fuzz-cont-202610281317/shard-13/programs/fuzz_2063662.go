package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool }
func genFn0() uint32 {
    var v4 uint16 = 42
    v5 := GenS1{F0: 605}
    _ = v4
    _ = v5
    return 114
}
func genFn1() uint64 {
    var v6 int32 = -2147483648
    v7 := GenS2{F0: 42, F1: 2}
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := []uint32{65535}
    var v2 uint16 = 32767
    var v3 float64 = float64(0.0)
    v2 -= 669
    v3 = v3
    v3 = 0.1
    _ = len(v1)
    var v8 int32 = 42
    if false {
        v9 := GenS1{F0: -1}
        var v10 float32 = (0.0 * 1.5)
        var v11 int = 82
        _ = v10
        _ = v11
        _ = v9
    } else {
        if true {
            switch -1 {
            case 128:
                v8 = -2147483648
            case 1:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 32767:
                if false {
                    var v12 float64 = (-2.25 * (v3 - v3))
                    var v13 int = 32768
                    fmt.Println(v3, v8, v13)
                    _ = v12
                    _ = v13
                } else {
                    v3 = v3
                }
            }
            v2 = (v2 | v2)
        }
        fmt.Println(v3, v8, v2)
    }
    var v14 float64 = 1e10
    if ((127 == 1) && true) {
        v15 := GenS3{F0: v2}
        for fi16 := 0; fi16 < 19; fi16++ {
            for fi17 := 0; fi17 < 8; fi17++ {
                if true {
                    switch "\n\t" {
                    case "foo bar":
                        var v18 int16 = 0
                        _ = v18
                    }
                } else {
                    var v19 bool = ((true || (true && false)) && (false && false))
                    v3 -= v14
                    if (true || (!false)) {
                        switch (("a" + "foo bar") + "a") {
                        case "a":
                            var v20 int32 = v8
                            _ = v20
                        case "foo bar":
                            fmt.Println(v14, fi16)
                        default:
                            fmt.Println(fi16, v19)
                        }
                    }
                    var v21 uint = 2
                    switch ("" + "a") {
                    case "":
                        var v22 string = ""
                        _ = v22
                    case "foo bar":
                        var v23 uint = 2
                        _ = v23
                    case "a":
                        var v24 float64 = v3
                        _ = v24
                    }
                    v19 = (v19 || (38 == -21))
                    _ = v19
                    _ = v21
                }
                _ = fi17
            }
            _ = fi16
        }
        _ = v15
    } else {
        if ((v2 | v2) <= v2) {
            v2 += (v2 - v2)
        }
        v25 := GenS4{F0: (true || true)}
        for fi26 := 0; fi26 < 10; fi26++ {
            switch fi26 {
            case 1:
                if len(v1) > 0 {
                    v1[0] = 217
                }
            default:
                fmt.Println(v14, fi26, v14)
            }
            _ = fi26
        }
        _ = v25
    }
    var v27 uint32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v8
}
