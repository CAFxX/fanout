package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 string; F1 uint32; F2 int64 }
func genFn0(p0 int, p1 uint32, p2 uint8) int {
    for fi6 := 0; fi6 < 12; fi6++ {
        if true {
            if (18446744073709551615 > 615) {
                p0 = (p0 / (((fi6 >> ((42) & 63))) | 1))
                p2 = p2
                for fi7 := 0; fi7 < 20; fi7++ {
                    var v8 float64 = float64(-2.25)
                    _ = fi7
                    _ = v8
                }
            }
            p0 += 4294967295
            p0 += 491
        }
        _ = fi6
    }
    v9 := []int{p0, 0, p0, -1}
    p2 += (p2 >> ((18446744073709551615) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 644
}

func main() {
    var v1 string = "hello"
    var v2 uint64 = 0
    v3 := []string{"hello"}
    var v4 int32 = 42
    var v5 string = (v1 + v1)
    if true {
        var v10 int8 = -62
        v1 = (v5 + "foo bar")
        _ = v10
    } else {
        v11 := GenS1{F0: 128, F1: 300}
        for fi12 := 0; fi12 < 7; fi12++ {
            switch v5 {
            case "\n\t":
                switch fi12 {
                case -225:
                    v1 = "hello"
                case 32768:
                    for fi13 := 0; fi13 < 20; fi13++ {
                        fmt.Println(fi12)
                        _ = fi13
                    }
                case 4294967295:
                    var v14 int8 = -128
                    _ = v14
                }
            }
            v5 = ""
            _ = fi12
        }
        _ = v11
    }
    if false {
        var v15 bool = true
        var v16 int64 = -1
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 15; fi17++ {
            v5 = "foo bar"
            _ = fi17
        }
        var v18 uint8 = 127
        v2 -= v2
        _ = v18
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v19 uint8 = 72
    var v20 int32 = 1
    if true {
        for fi21 := 0; fi21 < 16; fi21++ {
            fmt.Println(v1, v1, v20)
            fmt.Println(v5, v2, v4)
            _ = fi21
        }
        var v22 string = ("foo bar" + ((((v5 + v5) + "a") + "hello") + "\n\t"))
        _ = v22
    }
    switch "" {
    case "":
        var v23 uint8 = (v19 << ((18446744073709551615) & 63))
        _ = v23
    case "foo bar":
        var v24 int64 = 1
        _ = v24
    case "hello":
        fmt.Println(v20)
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        _ = len(v3)
        for fi26 := 0; fi26 < 17; fi26++ {
            v27 := GenS1{F0: 1, F1: v4}
            if false {
                v5 = ((("" + v5) + (v5 + ("\n\t" + v1))) + v1)
                var v28 int64 = 175
                _ = v28
            } else {
                if false {
                    if ((fi25 << ((107) & 63)) < 9223372036854775807) {
                        fmt.Println(v20, v4, fi26)
                    }
                    v29 := GenS1{F0: v19, F1: -39}
                    _ = v29
                } else {
                    v2 = v2
                    v1 = v5
                    fmt.Println(v2)
                }
            }
            if len(v3) > 0 {
                v3[0] = "a"
            }
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    switch "a" {
    case "a":
        v30 := GenS1{F0: v19, F1: (^(v4))}
        _ = v30
    case "hello":
        v1 = v1
    case "":
        for fi31 := 0; fi31 < 1; fi31++ {
            v32 := GenS2{F0: false, F1: 0}
            v33 := GenS3{F0: v5, F1: 1, F2: 764}
            _ = fi31
            _ = v32
            _ = v33
        }
    default:
        var v34 int32 = -289
        _ = v34
    }
    switch "\n\t" {
    case "\n\t":
        for fi35 := 0; fi35 < 6; fi35++ {
            if ((2 != 65535) && false) {
                for fi36 := 0; fi36 < 14; fi36++ {
                    var v37 float32 = (0.1 / 1.0)
                    _ = fi36
                    _ = v37
                }
                v38 := GenS1{F0: v19, F1: 1}
                _ = v38
            }
            _ = fi35
        }
    case "hello":
        var v39 uint16 = 2
        _ = v39
    case "":
        for fi40 := 0; fi40 < 11; fi40++ {
            v2 = (^(v2))
            v20 -= v20
            var v41 uint32 = 4294967295
            _ = fi40
            _ = v41
        }
    default:
        for fi42 := 0; fi42 < 15; fi42++ {
            if len(v3) > 0 {
                v3[0] = ("\n\t" + ((v5 + ("\n\t" + "foo bar")) + (((("\n\t" + ("\n\t" + v1)) + v5) + "\n\t") + ("" + v1))))
            }
            fmt.Println(v20)
            var v43 uint16 = 84
            _ = fi42
            _ = v43
        }
    }
    var v44 int8 = -1
    v45 := GenS2{F0: (2 > 4294967295), F1: v44}
    var v46 int32 = (v4 << ((1) & 63))
    v44 -= 100
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
