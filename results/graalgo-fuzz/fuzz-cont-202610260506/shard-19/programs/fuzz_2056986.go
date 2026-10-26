package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 bool }
type GenS2 struct { F0 uint }
func genFn0(p0 int16) uint64 {
    if true {
        v7 := []int{842}
        switch (("\n\t" + ("hello" + "foo bar")) + ("foo bar" + ("foo bar" + "\n\t"))) {
        case "\n\t":
            p0 += (p0 & (p0 / (((p0 + p0)) | 1)))
        case "a":
            var v8 float32 = 3.14159
            _ = v8
        case "foo bar":
            if true {
                for fi9 := 0; fi9 < 12; fi9++ {
                    switch fi9 {
                    case 1:
                        p0 = p0
                    }
                    _ = fi9
                }
            }
        default:
            p0 = (p0 % ((p0) | 1))
        }
        _ = v7
    } else {
        p0 = (p0 & p0)
        switch "" {
        case "hello":
            if (false || (!false)) {
                fmt.Println(p0)
            }
        case "\n\t":
            v10 := GenS1{F0: 9223372036854775807, F1: 613, F2: false}
            _ = v10
        case "a":
            var v11 bool = true
            _ = v11
        default:
            switch 127 {
            case 9223372036854775807:
                p0 -= p0
            case -668:
                for fi12 := 0; fi12 < 7; fi12++ {
                    if ("foo bar" != "foo bar") {
                        if false {
                            v13 := GenS1{F0: fi12, F1: fi12, F2: ((("hello" + "\n\t") + "a") <= "\n\t")}
                            v14 := GenS1{F0: 128, F1: (^(fi12)), F2: (255 <= 991)}
                            fmt.Println(p0)
                            _ = v13
                            _ = v14
                        }
                        p0 -= p0
                        var v15 bool = true
                        _ = v15
                    } else {
                        for fi16 := 0; fi16 < 20; fi16++ {
                            v17 := GenS2{F0: 255}
                            v18 := []int{}
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                    fmt.Println(p0)
                    var v19 uint16 = 2
                    _ = fi12
                    _ = v19
                }
            case 2147483647:
                fmt.Println(p0)
            }
        }
        p0 = 1
    }
    if false {
        v20 := []int{}
        _ = v20
    }
    var v21 int32 = -414
    var v22 int16 = -32768
    fmt.Println(v22, v21, v21)
    _ = p0
    _ = v21
    _ = v22
    return 708
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -539
    var v3 int64 = (v2 * v2)
    v4 := []int{32767}
    v5 := make(map[string]uint32)
    var v6 int32 = -1
    var v23 uint16 = 1
    v6 += -1
    v24 := v5["hello"]
    switch "hello" {
    case "":
        v25 := GenS2{F0: 440}
        _ = v25
    }
    v3 -= (v2 + (v3 % (((v2 | int64(v23))) | 1)))
    v26 := GenS1{F0: 4294967295, F1: -480, F2: ((255 >= (v23 << ((4294967295) & 63))) || (false || true))}
    var v27 float32 = 1.5
    fmt.Println(v27, v6)
    switch (("\n\t" + "\n\t") + (("a" + "foo bar") + "\n\t")) {
    case "\n\t":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "foo bar":
        v28 := GenS2{F0: 65535}
        _ = v28
    default:
        var v29 int64 = v3
        _ = v29
    }
    fmt.Println(v23, v27, v2)
    fmt.Println(v3)
    switch 638 {
    case 528:
        for fi30 := 0; fi30 < 11; fi30++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if true {
                for fi31 := 0; fi31 < 1; fi31++ {
                    v6 += (v6 + v6)
                    v3 -= v2
                    _ = fi31
                }
            }
            _ = fi30
        }
    default:
        v6 -= v6
    }
    if false {
        v24 -= (v24 + (v24 << ((631) & 63)))
        for fi32 := 0; fi32 < 16; fi32++ {
            fmt.Println(v3, v2)
            _ = fi32
        }
    }
    switch 74 {
    case 226:
        for fi33 := 0; fi33 < 12; fi33++ {
            var v34 float64 = float64(v27)
            v23 += (^(v23))
            _ = fi33
            _ = v34
        }
    }
    var v35 uint8 = 121
    var v36 int64 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
