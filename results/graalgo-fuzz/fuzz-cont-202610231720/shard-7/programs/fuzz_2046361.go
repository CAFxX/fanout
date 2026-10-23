package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int16; F2 uint }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0(p0 uint16, p1 uint8) int64 {
    var v4 uint16 = 255
    for fi5 := 0; fi5 < 2; fi5++ {
        if false {
            v6 := []int{fi5}
            p1 -= (p1 & (^(p1)))
            _ = v6
        }
        for fi7 := 0; fi7 < 11; fi7++ {
            if (1 > 1) {
                p1 = p1
                switch (fi5 >> ((0) & 63)) {
                case 127:
                    fmt.Println(p0, p1, fi7)
                }
                fmt.Println(p1, p0)
            } else {
                var v8 uint16 = (p0 % ((p0) | 1))
                if true {
                    v4 += (p0 % ((p0) | 1))
                    var v9 uint64 = 2
                    var v10 int16 = 304
                    _ = v10
                    _ = v9
                } else {
                    var v11 int = fi5
                    v8 = v8
                    _ = v11
                }
                _ = v8
            }
            if false {
                fmt.Println(p1, fi5)
                switch "foo bar" {
                case "hello":
                    switch (^(fi7)) {
                    case 32768:
                        p0 = (^(v4))
                        switch (("\n\t" + "a") + "foo bar") {
                        case "a":
                            switch fi5 {
                            case -225:
                                p0 = p0
                            case -718:
                                var v12 string = ("" + "\n\t")
                                _ = v12
                            default:
                                p1 = p1
                            }
                        case "":
                            var v13 int32 = 1
                            _ = v13
                        case "\n\t":
                            var v14 uint = 42
                            _ = v14
                        }
                    default:
                        var v15 float64 = -2.25
                        _ = v15
                    }
                default:
                    v16 := GenS1{F0: -1}
                    _ = v16
                }
                v17 := []int{(fi7 % ((fi5) | 1)), fi5, -1}
                _ = v17
            }
            _ = fi7
        }
        _ = fi5
    }
    if true {
        if false {
            v18 := GenS2{F0: 42, F1: 1, F2: 0}
            var v19 int16 = 47
            _ = v18
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        v21 := []int{(fi20 << ((255) & 63))}
        var v22 bool = false
        _ = fi20
        _ = v21
        _ = v22
    }
    _ = p0
    _ = p1
    _ = v4
    return 105
}

func main() {
    var v1 uint32 = 692
    v2 := []uint32{(v1 / (((v1 + (v1 ^ 1))) | 1)), (v1 >> ((4294967295) & 63)), (v1 << ((0) & 63))}
    var v3 float64 = (((((1e10 * (-2.25 * float64(0.1))) + 0.1) / (0.0 + (1e10 * float64(1.5)))) + ((float64(0.0) + 1.5) - 1e10)) - 0.0)
    var v23 uint8 = 128
    var v24 float32 = -2.25
    var v25 uint = 553
    v24 = v24
    v26 := GenS3{F0: 1, F1: 18446744073709551615}
    var v27 string = ""
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        v29 := GenS2{F0: v25, F1: -32768, F2: v25}
        v27 = "\n\t"
        v3 = (v3 + v3)
        _ = fi28
        _ = v29
    }
    v30 := GenS2{F0: (v25 + (v25 % ((v25) | 1))), F1: 42, F2: v25}
    switch "a" {
    case "":
        var v31 int16 = -1
        _ = v31
    case "hello":
        switch 9223372036854775807 {
        case 329:
            v2 = append(v2, v1)
        default:
            if (117 < 42) {
                v25 = (v25 - v25)
            }
        }
    case "a":
        switch 32767 {
        case 42:
            if false {
                var v32 uint8 = v23
                var v33 string = ((v27 + v27) + ("" + v27))
                _ = v32
                _ = v33
            } else {
                if len(v2) > 0 {
                    v2[0] = v1
                }
                var v34 float32 = v24
                _ = v34
            }
        }
    }
    var v35 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v35
}
