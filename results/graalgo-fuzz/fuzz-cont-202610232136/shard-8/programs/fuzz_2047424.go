package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint32 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint = 0
    v2 := []int{}
    v3 := make(map[string]uint32)
    for fi4 := 0; fi4 < 8; fi4++ {
        v1 = (v1 << ((922) & 63))
        _ = fi4
    }
    if false {
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 int8 = 1
            var v7 int64 = 475
            v8 := GenS1{F0: 1, F1: fi5, F2: 710}
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
    } else {
        var v9 uint8 = 255
        _ = v9
    }
    if (1 >= 0) {
        switch 32767 {
        case 0:
            v10 := v3[("hello" + (("a" + "\n\t") + "\n\t"))]
            _ = v10
        }
        var v11 bool = false
        _ = v11
    } else {
        switch ("foo bar" + "") {
        case "foo bar":
            v1 += v1
        case "a":
            _ = len(v2)
        case "\n\t":
            if (!false) {
                var v12 uint16 = 144
                if (-2.25 != (3.14159 * ((((1.5 * 1e10) - float64(3.14159)) - (0.1 + ((3.14159 - 1.5) + ((float64(1e10) - 0.1) * 1e10)))) + 0.1))) {
                    v1 -= (v1 >> ((v1) & 63))
                    v13 := GenS1{F0: 89, F1: -903, F2: 1}
                    for fi14 := 0; fi14 < 7; fi14++ {
                        v15 := GenS1{F0: -1, F1: fi14, F2: 2147483647}
                        v1 = uint(v12)
                        _ = fi14
                        _ = v15
                    }
                    _ = v13
                }
                _ = v12
            } else {
                var v16 uint = 979
                v17 := GenS1{F0: 2147483647, F1: 94, F2: 284}
                v18 := GenS2{F0: 0}
                _ = v16
                _ = v17
                _ = v18
            }
        }
        switch 4294967295 {
        case -375:
            var v19 int32 = -249
            _ = v19
        case 127:
            v1 -= v1
        case -2147483648:
            v3[("a" + ("" + "a"))] = 65535
        default:
            if ((^(v1)) > (^(v1))) {
                var v20 uint16 = 206
                var v21 uint64 = uint64(v1)
                if (!false) {
                    var v22 uint = v1
                    _ = v22
                } else {
                    if (!true) {
                        for fi23 := 0; fi23 < 12; fi23++ {
                            for fi24 := 0; fi24 < 10; fi24++ {
                                if false {
                                    v21 -= v21
                                    v21 = (v21 - (^(v21)))
                                    v20 = v20
                                } else {
                                    if false {
                                        v1 = v1
                                        switch fi24 {
                                        case 820:
                                            v2 = append(v2, -2147483648)
                                        case 1:
                                            fmt.Println(v21)
                                        }
                                    }
                                }
                                if true {
                                    var v25 float64 = ((0.0 + (3.14159 - 1e10)) / 1.0)
                                    fmt.Println(fi24)
                                    _ = len(v2)
                                    _ = v25
                                } else {
                                    switch ((("a" + "") + "a") + "\n\t") {
                                    case "":
                                        var v26 uint = (v1 << (((^(v1))) & 63))
                                        _ = v26
                                    case "hello":
                                        if ((!(!false)) || false) {
                                            var v27 uint = v1
                                            var v28 uint32 = 82
                                            _ = v27
                                            _ = v28
                                        }
                                    case "a":
                                        if true {
                                            _ = len(v3)
                                            v29 := v3["a"]
                                            delete(v3, ("" + "foo bar"))
                                            _ = v29
                                        }
                                    }
                                    v30 := GenS1{F0: -372, F1: fi24, F2: 2147483647}
                                    _ = v30
                                }
                                v1 += 817
                                _ = fi24
                            }
                            switch "hello" {
                            case "hello":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case "foo bar":
                                v1 -= (v1 ^ v1)
                            }
                            _ = fi23
                        }
                    } else {
                        var v31 string = "foo bar"
                        _ = v31
                    }
                    var v32 uint32 = 2147483647
                    v33 := GenS1{F0: -2147483648, F1: 127, F2: 42}
                    _ = v32
                    _ = v33
                }
                _ = v20
                _ = v21
            }
        }
    }
    fmt.Println(v1)
    switch 416 {
    case 42:
        v1 = (v1 * v1)
    case -592:
        v1 -= v1
    default:
        v1 = (v1 - v1)
    }
    if true {
        fmt.Println(v1)
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        var v35 float64 = 3.14159
        _ = fi34
        _ = v35
    }
    v1 = v1
    v1 = 2
    v36 := GenS2{F0: -23}
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
}
