package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint16 }

func main() {
    var v1 bool = true
    var v2 bool = v1
    v3 := make(map[string]uint32)
    var v4 uint16 = 65535
    switch 32768 {
    case 32767:
        for fi5 := 0; fi5 < 10; fi5++ {
            v2 = (0 <= 0)
            _ = fi5
        }
    case 0:
        if (-119 > 642) {
            v1 = v1
            v6 := GenS1{F0: -51}
            var v7 string = "foo bar"
            _ = v6
            _ = v7
        } else {
            v8 := GenS1{F0: -771}
            if false {
                fmt.Println(v2, v1, v1)
                v9 := v3["hello"]
                v4 -= (v4 + (^(v4)))
                _ = v9
            }
            _ = v8
        }
    case -3:
        var v10 uint8 = 127
        _ = v10
    default:
        v2 = (float64(0.0) < 1.5)
    }
    _ = len(v3)
    v2 = (!true)
    v2 = v1
    var v11 uint = 936
    if (((v4 / (((^(v4))) | 1)) > (^(v4))) && (2 != 42)) {
        v11 = v11
        var v12 float64 = (3.14159 + (0.0 * -2.25))
        _ = v12
    }
    var v13 uint = (v11 * (v11 << (((v11 & (v11 * (^(v11))))) & 63)))
    var v14 uint32 = 42
    v13 -= (v11 << ((v13) & 63))
    v15 := GenS1{F0: 1000}
    v2 = v1
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 uint16 = 1
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: 1000}
    v19 := GenS2{F0: 127, F1: 42, F2: v4}
    v20 := []int{-1, 32768, -991}
    switch "foo bar" {
    case "":
        if true {
            if v2 {
                fmt.Println(v2, v14)
                var v21 int32 = 1
                _ = v21
            } else {
                if (4294967295 < 42) {
                    var v22 uint8 = 42
                    _ = v22
                }
                fmt.Println(v13, v2, v11)
            }
            for fi23 := 0; fi23 < 20; fi23++ {
                switch ("hello" + "\n\t") {
                case "\n\t":
                    var v24 uint = (v13 >> (((^(v11))) & 63))
                    _ = v24
                case "foo bar":
                    for fi25 := 0; fi25 < 20; fi25++ {
                        if (true || v2) {
                            switch 127 {
                            case 32768:
                                v14 = v14
                            default:
                                if ((v4 * uint16(v13)) != (v4 & 32767)) {
                                    var v26 bool = v1
                                    switch ("foo bar" + "\n\t") {
                                    case "a":
                                        var v27 int = (fi23 * (-(fi25)))
                                        _ = v27
                                    default:
                                        var v28 uint64 = 42
                                        _ = v28
                                    }
                                    fmt.Println(v26, v4)
                                    v4 = 737
                                    delete(v3, ("hello" + "\n\t"))
                                    _ = v26
                                } else {
                                    v20 = append(v20, fi23)
                                }
                            }
                            var v29 uint32 = 2147483647
                            switch "\n\t" {
                            case "hello":
                                var v30 string = "\n\t"
                                _ = v30
                            case "foo bar":
                                v3["\n\t"] = v29
                            case "\n\t":
                                var v31 int16 = 270
                                _ = v31
                            }
                            _ = v29
                        }
                        if ((42 >= 127) || v1) {
                            v2 = (!(42 == 137))
                            var v32 int8 = 100
                            _ = v32
                        }
                        var v33 float32 = (1.5 / 1.5)
                        _ = fi25
                        _ = v33
                    }
                case "hello":
                    var v34 float32 = 3.14159
                    _ = v34
                default:
                    fmt.Println(v14)
                }
                _ = fi23
            }
            var v35 int8 = -128
            _ = v35
        } else {
            var v36 int = 32767
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
