package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint8, p1 string) int32 {
    p1 = (p1 + p1)
    if (2 == 255) {
        v5 := []int{}
        v6 := GenS1{F0: -2147483648, F1: 0, F2: 1}
        for fi7 := 0; fi7 < 16; fi7++ {
            p0 = 128
            if (!(3.14159 >= 3.14159)) {
                p1 = p1
            } else {
                switch 39 {
                case -9223372036854775808:
                    var v8 int16 = 32767
                    _ = v8
                default:
                    var v9 int16 = 1
                    _ = v9
                }
                switch "foo bar" {
                case "foo bar":
                    switch ("\n\t" + p1) {
                    case "hello":
                        if false {
                            var v10 float32 = (1e10 + (0.0 - 0.0))
                            var v11 int = (fi7 & int(p0))
                            v12 := GenS1{F0: 0, F1: 0, F2: 2147483648}
                            _ = v10
                            _ = v11
                            _ = v12
                        } else {
                            for fi13 := 0; fi13 < 20; fi13++ {
                                var v14 int16 = -32768
                                p1 = p1
                                fmt.Println(fi13, p1, fi7)
                                switch "hello" {
                                case "":
                                    fmt.Println(v14, p1, p1)
                                case "\n\t":
                                    p1 = "a"
                                }
                                _ = fi13
                                _ = v14
                            }
                            p1 = "foo bar"
                        }
                    default:
                        _ = len(v5)
                    }
                case "a":
                    if (849 == 2) {
                        p1 = ("" + "foo bar")
                    } else {
                        var v15 int16 = -32768
                        _ = v15
                    }
                default:
                    switch (p1 + p1) {
                    case "a":
                        v16 := GenS2{F0: 1, F1: (p0 - p0)}
                        _ = v16
                    case "foo bar":
                        if len(v5) > 0 {
                            v5[0] = fi7
                        }
                    default:
                        var v17 uint64 = 1
                        _ = v17
                    }
                }
            }
            _ = fi7
        }
        _ = v5
        _ = v6
    }
    switch (p1 + (p1 + (p1 + (p1 + "hello")))) {
    case "\n\t":
        fmt.Println(p1, p0)
    case "foo bar":
        var v18 string = "\n\t"
        _ = v18
    case "":
        v19 := []int{}
        _ = v19
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 int64 = -1
    var v2 float64 = ((3.14159 / -2.25) * (1.5 + 0.0))
    v3 := []uint32{2147483647, 0, 119, 174}
    var v4 int64 = v1
    var v20 int32 = -715
    var v21 int32 = v20
    var v22 float64 = ((-2.25 + (v2 + 1.5)) * v2)
    v23 := GenS2{F0: v4, F1: 2}
    v21 += (v20 ^ v21)
    v4 = v1
    v21 -= -1
    var v24 bool = false
    switch 127 {
    case 42:
        for fi25 := 0; fi25 < 6; fi25++ {
            for fi26 := 0; fi26 < 19; fi26++ {
                _ = len(v3)
                fmt.Println(v22, v24)
                var v27 uint16 = 255
                var v28 int32 = -2147483648
                _ = fi26
                _ = v27
                _ = v28
            }
            var v29 uint32 = 2147483647
            _ = fi25
            _ = v29
        }
    case 32768:
        switch -293 {
        case 2147483647:
            v21 -= genFn0(255, "\n\t")
        case 574:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            for fi30 := 0; fi30 < 2; fi30++ {
                if v24 {
                    v31 := GenS1{F0: -382, F1: v21, F2: 65535}
                    v21 += genFn0(25, ("" + "\n\t"))
                    _ = v31
                } else {
                    _ = len(v3)
                    v32 := GenS2{F0: (v1 / ((-851) | 1)), F1: 42}
                    _ = v32
                }
                v33 := GenS3{F0: (((v21 & (v21 >> ((1) & 63))) == (v21 + (-(v20)))) || v24)}
                _ = fi30
                _ = v33
            }
        }
    default:
        for fi34 := 0; fi34 < 4; fi34++ {
            switch "foo bar" {
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            for fi35 := 0; fi35 < 16; fi35++ {
                v36 := GenS3{F0: (!(v4 <= 1))}
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
    }
    v4 -= -9223372036854775808
    for fi37 := 0; fi37 < 10; fi37++ {
        v4 += v1
        var v38 int = fi37
        if false {
            var v39 uint64 = 18446744073709551615
            var v40 uint32 = 2147483648
            _ = v39
            _ = v40
        }
        _ = fi37
        _ = v38
    }
    if v24 {
        fmt.Println(v21, v1, v24)
        v41 := GenS1{F0: v20, F1: v20, F2: 835}
        _ = v41
    }
    v20 = -1
    for fi42 := 0; fi42 < 7; fi42++ {
        var v43 uint = 1
        _ = fi42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
