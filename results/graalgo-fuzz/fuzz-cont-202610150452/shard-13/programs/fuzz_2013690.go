package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32 }
func genFn0() uint32 {
    var v4 int8 = 100
    v5 := GenS1{F0: 1, F1: ("\n\t" + "\n\t")}
    var v6 float32 = float32(0.1)
    v7 := []int{625, 127}
    switch 4294967295 {
    case -867:
        fmt.Println(v6)
    case -2147483648:
        v6 += float32(1e10)
    }
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int16) uint16 {
    v8 := []int{0}
    switch 32767 {
    case 42:
        v9 := GenS2{F0: 1}
        _ = v9
    }
    var v10 string = ("hello" + ((("a" + ("" + ("" + "a"))) + "foo bar") + "foo bar"))
    _ = p0
    _ = v10
    _ = v8
    return 195
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    var v3 string = ("foo bar" + ("hello" + ("\n\t" + "")))
    var v11 bool = (128 != 380)
    if (v11 && (127 == 42)) {
        fmt.Println(v11, v3)
        _ = len(v2)
        v12 := []int{-2147483648}
        _ = v12
    } else {
        switch -302 {
        case 9223372036854775807:
            v11 = true
        case 42:
            v3 = v3
        case 944:
            var v13 string = ("hello" + v3)
            _ = v13
        default:
            var v14 float32 = (0.0 + -2.25)
            _ = v14
        }
        if (-62 < 0) {
            var v15 uint = 4294967295
            switch -293 {
            case 4294967295:
                if true {
                    v15 += v15
                    for fi16 := 0; fi16 < 15; fi16++ {
                        switch "hello" {
                        case "a":
                            for fi17 := 0; fi17 < 3; fi17++ {
                                v15 = v15
                                switch ((v3 + ("foo bar" + (("" + "foo bar") + v3))) + v3) {
                                case "foo bar":
                                    var v18 uint64 = 646
                                    _ = v18
                                case "\n\t":
                                    v19 := GenS3{F0: genFn1(-32768), F1: genFn0()}
                                    _ = v19
                                default:
                                    var v20 bool = (2 == 65535)
                                    _ = v20
                                }
                                _ = fi17
                            }
                        case "hello":
                            fmt.Println(v3)
                        case "foo bar":
                            v11 = false
                        }
                        _ = fi16
                    }
                }
            case 9223372036854775807:
                v3 = (((("" + "") + ("foo bar" + (v3 + ""))) + "foo bar") + "a")
            }
            v1[128] = 879
            _ = v15
        }
        v21 := GenS4{F0: v11}
        _ = v21
    }
    v22 := GenS5{F0: 46}
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 int64 = 724
        if (0.0 < -2.25) {
            var v25 uint = 649
            _ = v25
        }
        fmt.Println(v24, v11)
        _ = fi23
        _ = v24
    }
    if false {
        fmt.Println(v11, v11)
        v26 := v1[116]
        for fi27 := 0; fi27 < 6; fi27++ {
            var v28 float32 = float32(-2.25)
            if (!v11) {
                delete(v1, (fi27 << ((2) & 63)))
            }
            _ = fi27
            _ = v28
        }
        _ = v26
    } else {
        fmt.Println(v11)
        if (784 < 65535) {
            for fi29 := 0; fi29 < 2; fi29++ {
                v11 = v11
                for fi30 := 0; fi30 < 13; fi30++ {
                    v3 = ((v3 + v3) + "\n\t")
                    _ = fi30
                }
                _ = fi29
            }
        }
    }
    v31 := []int{-1, 32767, -925}
    for fi32 := 0; fi32 < 20; fi32++ {
        v33 := GenS4{F0: v11}
        _ = fi32
        _ = v33
    }
    fmt.Println(v11, v11)
    v34 := GenS5{F0: 42}
    if (false && (!true)) {
        fmt.Println(v11, v11)
        v11 = true
    }
    if (-356 >= 0) {
        for fi35 := 0; fi35 < 3; fi35++ {
            var v36 string = v3
            v11 = v11
            _ = fi35
            _ = v36
        }
        v3 = v3
    }
    if (2 > 1) {
        var v37 uint16 = 32767
        fmt.Println(v3, v3, v37)
        _ = v37
    } else {
        var v38 string = ("\n\t" + "foo bar")
        if (42 <= -1) {
            var v39 float64 = (1.5 / 1e10)
            _ = v39
        } else {
            v38 = "\n\t"
            v40 := GenS6{F0: 51}
            _ = v40
        }
        var v41 uint8 = 1
        _ = v38
        _ = v41
    }
    v42 := GenS1{F0: 1, F1: "\n\t"}
    v3 = v3
    v3 = v3
    v43 := GenS7{F0: 672}
    if len(v31) > 0 {
        v31[0] = -2147483648
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v34
    _ = v42
    _ = v43
}
