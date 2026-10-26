package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool }

func main() {
    v1 := []uint32{}
    v2 := []int{922}
    v3 := make(map[string]int)
    var v4 int32 = 0
    var v5 int = 740
    v4 -= (v4 >> ((2) & 63))
    v6 := GenS1{F0: (v5 % ((42) | 1)), F1: 32767, F2: v5}
    switch -284 {
    case 0:
        var v7 uint64 = 4294967295
        _ = v7
    }
    var v8 uint16 = 65535
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v9 int8 = -76
    if ((v5 | v5) != v5) {
        v5 = v5
    }
    v10 := GenS1{F0: (^(v5)), F1: (v5 * (v5 | -2147483648)), F2: v5}
    fmt.Println(v8, v8, v5)
    switch v5 {
    case -1:
        for fi11 := 0; fi11 < 3; fi11++ {
            fmt.Println(v8)
            _ = fi11
        }
    }
    if (!(v9 <= (v9 - (v9 | int8(v5))))) {
        if false {
            if true {
                var v12 string = ""
                v8 = (v8 / (((v8 | 1)) | 1))
                v9 = (v9 >> ((4294967295) & 63))
                _ = v12
            }
            v9 += v9
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v13 int64 = -1
        _ = v13
    }
    fmt.Println(v4, v4)
    if (v4 > 431) {
        v9 = (v9 & v9)
    } else {
        v14 := GenS1{F0: v5, F1: (^(v5)), F2: (v5 - int(v8))}
        _ = v14
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        if true {
            v8 -= 65535
            v16 := GenS1{F0: 9223372036854775807, F1: (-(fi15)), F2: -685}
            _ = v16
        } else {
            if ((("a" + ("hello" + "\n\t")) + ("a" + "")) > "\n\t") {
                if (!false) {
                    v17 := GenS1{F0: (v5 - -184), F1: fi15, F2: (fi15 * 0)}
                    var v18 uint = 42
                    _ = v17
                    _ = v18
                } else {
                    if (("\n\t" + "\n\t") <= "\n\t") {
                        v9 = v9
                        switch 32768 {
                        case 1:
                            switch ((("" + "\n\t") + "foo bar") + (("" + ("" + "hello")) + "")) {
                            case "":
                                v19 := GenS2{F0: 105}
                                _ = v19
                            case "hello":
                                v9 -= -128
                            default:
                                fmt.Println(v9, fi15, fi15)
                            }
                        case -2147483648:
                            switch v5 {
                            case 600:
                                switch v5 {
                                case -2147483648:
                                    var v20 uint64 = 888
                                    _ = v20
                                case 32767:
                                    _ = len(v2)
                                case 550:
                                    var v21 uint8 = 127
                                    _ = v21
                                }
                            case -162:
                                fmt.Println(v9, v9)
                            case 2147483647:
                                for fi22 := 0; fi22 < 12; fi22++ {
                                    var v23 uint16 = (v8 * v8)
                                    v24 := GenS3{F0: (9223372036854775807 > 1)}
                                    var v25 int8 = v9
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                            }
                        }
                    } else {
                        var v26 int = v5
                        for fi27 := 0; fi27 < 11; fi27++ {
                            var v28 uint16 = (v8 - (v8 << ((358) & 63)))
                            _ = fi27
                            _ = v28
                        }
                        _ = v26
                    }
                    var v29 float64 = float64((3.14159 * 0.0))
                    delete(v3, "\n\t")
                    _ = v29
                }
                fmt.Println(v5, v5, v4)
                var v30 int64 = -1
                _ = v30
            }
        }
        v2 = append(v2, v5)
        for fi31 := 0; fi31 < 4; fi31++ {
            if (!false) {
                fmt.Println(v9, v5, fi15)
                v9 -= (v9 - v9)
                fmt.Println(v4, fi31)
                var v32 float32 = 0.1
                var v33 uint64 = 1
                v5 += 42
                _ = v32
                _ = v33
            } else {
                if (-419 != 1000) {
                    delete(v3, "hello")
                    fmt.Println(v4, fi15)
                    var v34 uint32 = 4294967295
                    fmt.Println(fi31, v34)
                    _ = v34
                }
            }
            _ = fi31
        }
        _ = fi15
    }
    v4 -= 42
    v35 := GenS3{F0: false}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
