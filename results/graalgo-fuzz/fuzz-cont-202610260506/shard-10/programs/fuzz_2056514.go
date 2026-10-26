package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint8 }

func main() {
    var v1 uint32 = 151
    v2 := make(map[string]uint32)
    var v3 string = ("foo bar" + "\n\t")
    v4 := make(map[string]string)
    var v5 uint16 = 0
    var v6 float32 = 3.14159
    v7 := v4[v3]
    v8 := []int{42}
    if true {
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := GenS1{F0: v3, F1: 857, F2: 160}
            for fi11 := 0; fi11 < 4; fi11++ {
                if (v1 >= v1) {
                    v3 = ""
                } else {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                }
                if ((850 > 42) && false) {
                    v1 -= (v1 | v1)
                } else {
                    v6 -= v6
                }
                _ = fi11
            }
            if (fi9 != (fi9 << ((70) & 63))) {
                v3 = ("hello" + (v7 + (v3 + (v3 + v7))))
                switch fi9 {
                case 128:
                    v6 = float32(-2.25)
                case 432:
                    var v12 int8 = -1
                    _ = v12
                case -2147483648:
                    for fi13 := 0; fi13 < 10; fi13++ {
                        var v14 uint64 = 18446744073709551615
                        _ = fi13
                        _ = v14
                    }
                default:
                    if (!true) {
                        var v15 int = fi9
                        v16 := GenS1{F0: ("hello" + v7), F1: 4294967295, F2: 2}
                        _ = v15
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 13; fi17++ {
                            var v18 uint16 = v5
                            if (v7 > v7) {
                                var v19 uint = 255
                                switch fi17 {
                                case -1:
                                    v7 = v7
                                }
                                _ = v19
                            }
                            _ = fi17
                            _ = v18
                        }
                        fmt.Println(v5)
                    }
                }
            }
            _ = fi9
            _ = v10
        }
        var v20 uint64 = 18446744073709551615
        var v21 uint16 = (v5 ^ (v5 / (((v5 - (v5 / ((v5) | 1)))) | 1)))
        _ = v20
        _ = v21
    }
    var v22 bool = (false || false)
    v7 = v3
    var v23 int16 = -32768
    if len(v8) > 0 {
        v8[0] = 32767
    }
    v7 = v3
    _ = len(v2)
    var v24 string = (v3 + ((v3 + (v3 + "")) + v3))
    switch int(v5) {
    case 9223372036854775807:
        for fi25 := 0; fi25 < 3; fi25++ {
            if v22 {
                if false {
                    _ = len(v8)
                    for fi26 := 0; fi26 < 13; fi26++ {
                        var v27 int32 = 42
                        var v28 uint16 = v5
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        v29 := GenS1{F0: "", F1: 1, F2: 255}
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    fmt.Println(v1, v7, v6)
                } else {
                    var v30 float32 = 1.5
                    _ = v30
                }
                v31 := v2[v24]
                switch (fi25 / ((fi25) | 1)) {
                case 4294967295:
                    var v32 int64 = 42
                    _ = v32
                }
                _ = v31
            } else {
                var v33 int32 = 1
                _ = v33
            }
            var v34 int16 = 0
            _ = fi25
            _ = v34
        }
    case -743:
        for fi35 := 0; fi35 < 6; fi35++ {
            var v36 float64 = 0.0
            v4["foo bar"] = "hello"
            var v37 uint = 255
            _ = fi35
            _ = v36
            _ = v37
        }
    case 127:
        var v38 int64 = -305
        _ = v38
    default:
        var v39 int64 = 42
        _ = v39
    }
    _ = len(v4)
    v40 := GenS1{F0: v3, F1: 4294967295, F2: 2}
    for fi41 := 0; fi41 < 9; fi41++ {
        v6 = v6
        fmt.Println(v7)
        _ = fi41
    }
    var v42 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
