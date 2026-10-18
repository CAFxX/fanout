package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 bool }
type GenS7 struct { F0 int8; F1 int16; F2 string }

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]string)
    var v3 float64 = (1.5 / (0.1 - 1.5))
    var v4 uint32 = 42
    var v5 int16 = 32767
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            v4 = v4
            var v8 int8 = int8(v5)
            v9 := v2[v4]
            if len(v1) > 0 {
                v1[0] = v4
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        if (3.14159 >= 1.5) {
            for fi10 := 0; fi10 < 12; fi10++ {
                if true {
                    v1 = append(v1, (v4 << ((65535) & 63)))
                    if false {
                        var v11 int32 = int32(fi10)
                        var v12 float64 = v3
                        var v13 uint = 1
                        var v14 uint32 = v4
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    if true {
                        if ((0.0 / v3) <= (v3 / -2.25)) {
                            if false {
                                var v15 float64 = 3.14159
                                var v16 uint = 732
                                _ = v15
                                _ = v16
                            } else {
                                v17 := GenS1{F0: 65535}
                                fmt.Println(fi6)
                                _ = v17
                            }
                            switch (("\n\t" + (("" + "") + "hello")) + "hello") {
                            case "hello":
                                var v18 int32 = 42
                                _ = v18
                            case "\n\t":
                                fmt.Println(v4)
                            }
                        } else {
                            v19 := GenS2{F0: "a"}
                            v4 = v4
                            var v20 uint8 = 55
                            _ = v19
                            _ = v20
                        }
                    } else {
                        var v21 int64 = 0
                        if (255 >= 93) {
                            v22 := GenS1{F0: v4}
                            var v23 uint = 0
                            _ = v22
                            _ = v23
                        }
                        v5 = v5
                        _ = v21
                    }
                    if ((("hello" + "hello") + "\n\t") == "a") {
                        var v24 float32 = 1.5
                        if len(v1) > 0 {
                            v1[0] = 870
                        }
                        _ = v24
                    } else {
                        fmt.Println(v5, v4, v3)
                        v1 = append(v1, (^(v4)))
                        v3 -= (float64((1e10 - 1.5)) + v3)
                    }
                    var v25 float32 = ((0.0 / -2.25) - 3.14159)
                    _ = v25
                }
                var v26 float32 = 3.14159
                v5 -= (^(v5))
                _ = fi10
                _ = v26
            }
            if false {
                v3 = (0.0 / float64((-2.25 - float32(v3))))
                for fi27 := 0; fi27 < 9; fi27++ {
                    v4 += v4
                    fmt.Println(v4)
                    _ = fi27
                }
            } else {
                v3 += v3
            }
        } else {
            v5 = v5
            switch fi6 {
            case 848:
                v28 := GenS3{F0: 2}
                _ = v28
            default:
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
        }
        v29 := GenS1{F0: (v4 ^ (v4 % ((v4) | 1)))}
        _ = fi6
        _ = v29
    }
    if (!((("foo bar" + "a") > "\n\t") && (-114 == 42))) {
        v30 := GenS4{F0: "hello"}
        v31 := GenS5{F0: 128}
        _ = v30
        _ = v31
    } else {
        var v32 uint16 = 358
        v4 = (v4 - v4)
        fmt.Println(v5, v4)
        _ = v32
    }
    switch 218 {
    case 675:
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            v33 := GenS6{F0: v5, F1: true}
            _ = v33
        case "a":
            var v34 uint8 = 2
            _ = v34
        default:
            var v35 bool = (v4 != v4)
            _ = v35
        }
    case 4294967295:
        if ((("a" + "") > ("foo bar" + "")) && ((!(!(90 >= 127))) || (!(!false)))) {
            fmt.Println(v5, v3, v3)
        }
    case 127:
        var v36 uint64 = 0
        _ = v36
    default:
        v37 := GenS5{F0: int(v4)}
        _ = v37
    }
    for fi38 := 0; fi38 < 15; fi38++ {
        var v39 uint32 = v4
        fmt.Println(v39, v39, v5)
        for fi40 := 0; fi40 < 13; fi40++ {
            var v41 uint8 = 86
            _ = fi40
            _ = v41
        }
        _ = fi38
        _ = v39
    }
    v5 = v5
    v42 := GenS7{F0: 25, F1: 32767, F2: (("foo bar" + "foo bar") + "")}
    var v43 bool = true
    v5 = (v5 * 32767)
    for fi44 := 0; fi44 < 7; fi44++ {
        v45 := GenS1{F0: (^(v4))}
        _ = fi44
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
