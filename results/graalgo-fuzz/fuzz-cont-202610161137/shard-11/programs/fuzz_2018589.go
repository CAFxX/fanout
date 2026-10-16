package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int16) uint {
    p0 = (p0 ^ (p0 * p0))
    if false {
        p0 -= (-(p0))
        switch "hello" {
        case "\n\t":
            switch 708 {
            case 78:
                fmt.Println(p0)
            case 32767:
                if false {
                    for fi5 := 0; fi5 < 14; fi5++ {
                        fmt.Println(p0, fi5)
                        _ = fi5
                    }
                } else {
                    if (true || (255 > 127)) {
                        switch ("foo bar" + ("" + "\n\t")) {
                        case "":
                            p0 = 32767
                        case "foo bar":
                            var v6 int = 127
                            _ = v6
                        case "\n\t":
                            var v7 uint64 = 4294967295
                            _ = v7
                        }
                        switch "foo bar" {
                        case "\n\t":
                            switch "hello" {
                            case "a":
                                p0 += p0
                            case "foo bar":
                                var v8 int = 1
                                _ = v8
                            case "\n\t":
                                v9 := GenS1{F0: 255, F1: 255, F2: 127}
                                _ = v9
                            default:
                                fmt.Println(p0)
                            }
                        case "a":
                            v10 := []int{32767, 42, 4294967295}
                            _ = v10
                        case "":
                            fmt.Println(p0)
                        }
                    }
                    p0 = (p0 - p0)
                    var v11 uint16 = 1
                    _ = v11
                }
            }
        case "foo bar":
            v12 := []int{-950, 164, 2147483647, 2147483647}
            _ = v12
        }
    }
    var v13 uint32 = 2147483647
    _ = p0
    _ = v13
    return 255
}
func genFn1(p0 uint, p1 uint8) int64 {
    var v14 float32 = 1e10
    v14 = float32(-2.25)
    _ = p0
    _ = p1
    _ = v14
    return 0
}

func main() {
    var v1 string = (("" + ("\n\t" + "\n\t")) + ("\n\t" + (("" + "hello") + "foo bar")))
    var v2 string = v1
    var v3 int64 = 302
    var v4 string = v1
    if true {
        v15 := GenS1{F0: 32767, F1: 126, F2: 0}
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := []int{}
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        var v18 float64 = 1.5
        v2 = ((v1 + ("foo bar" + ("\n\t" + v1))) + ("hello" + "foo bar"))
        if (false || ((105 <= 292) && true)) {
            v1 = "a"
        }
        _ = v18
    }
    var v19 int = -9223372036854775808
    v2 = ("hello" + v1)
    v20 := GenS1{F0: 811, F1: 7, F2: -128}
    if false {
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := []int{v19, 1}
            for fi23 := 0; fi23 < 14; fi23++ {
                v2 = v1
                v24 := GenS2{F0: -1}
                if true {
                    var v25 uint32 = 1
                    var v26 uint16 = 790
                    _ = v25
                    _ = v26
                }
                _ = fi23
                _ = v24
            }
            _ = fi21
            _ = v22
        }
        if true {
            for fi27 := 0; fi27 < 6; fi27++ {
                v28 := GenS2{F0: -22}
                v4 = v2
                v1 = v1
                _ = fi27
                _ = v28
            }
        }
        v29 := GenS1{F0: 379, F1: 197, F2: 127}
        _ = v29
    } else {
        switch "foo bar" {
        case "a":
            if (!(32767 <= 0)) {
                for fi30 := 0; fi30 < 19; fi30++ {
                    var v31 int8 = -113
                    _ = fi30
                    _ = v31
                }
                var v32 uint32 = 1
                v33 := []int{(v19 * (v19 | v19)), -657, v19}
                _ = v32
                _ = v33
            }
        case "\n\t":
            fmt.Println(v3, v2, v19)
        }
    }
    v4 = v1
    v34 := []int{2147483647, v19, v19}
    if len(v34) >= 2 {
        v34 = v34[1:len(v34)]
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        v36 := GenS3{F0: 2}
        var v37 uint8 = 131
        switch -726 {
        case -2147483648:
            if (false || (v37 > v37)) {
                if len(v34) >= 2 {
                    v34 = v34[1:len(v34)]
                }
            }
        default:
            var v38 float32 = 0.1
            _ = v38
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    var v39 uint = 18446744073709551615
    fmt.Println(v19)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v39
    _ = v4
}
