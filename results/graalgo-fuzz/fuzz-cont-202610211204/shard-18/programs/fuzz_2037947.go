package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int32; F1 uint8; F2 int8 }

func main() {
    var v1 int8 = 127
    v2 := make(map[string]string)
    var v3 int8 = (^(v1))
    for fi4 := 0; fi4 < 10; fi4++ {
        for fi5 := 0; fi5 < 17; fi5++ {
            if true {
                var v6 uint8 = 138
                switch fi4 {
                case 1:
                    var v7 uint16 = 884
                    _ = v7
                default:
                    if true {
                        v6 -= v6
                        fmt.Println(v6)
                        if true {
                            switch fi5 {
                            case -2147483648:
                                v1 = -12
                            }
                        } else {
                            var v8 int8 = (v1 - v1)
                            var v9 uint16 = 685
                            if true {
                                if ((1e10 + -2.25) > 0.1) {
                                    var v10 string = "\n\t"
                                    _ = v10
                                }
                                v8 = v1
                                var v11 int64 = 579
                                _ = v11
                            } else {
                                if (uint64(v9) >= 124) {
                                    v12 := []int{1, 128, fi4}
                                    var v13 uint64 = 4294967295
                                    _ = v12
                                    _ = v13
                                }
                                v14 := GenS1{F0: 42, F1: 2147483648}
                                v1 += (-(v3))
                                _ = v14
                            }
                            _ = v8
                            _ = v9
                        }
                    }
                }
                _ = v6
            }
            delete(v2, ("\n\t" + "\n\t"))
            var v15 int8 = v3
            _ = fi5
            _ = v15
        }
        v3 += 0
        _ = fi4
    }
    var v16 uint8 = 203
    fmt.Println(v3, v1)
    if (false && true) {
        var v17 uint64 = 18446744073709551615
        _ = v17
    } else {
        for fi18 := 0; fi18 < 13; fi18++ {
            v19 := []int{fi18, (fi18 >> ((0) & 63)), -259, 2147483647}
            if (v16 < 250) {
                if (int8(fi18) > 115) {
                    var v20 bool = true
                    _ = v20
                }
                if true {
                    switch (fi18 | (fi18 - 2147483647)) {
                    case 32767:
                        for fi21 := 0; fi21 < 4; fi21++ {
                            switch (("" + "foo bar") + "") {
                            case "foo bar":
                                var v22 uint8 = v16
                                _ = v22
                            default:
                                v3 = (v1 << ((682) & 63))
                            }
                            _ = fi21
                        }
                    default:
                        for fi23 := 0; fi23 < 6; fi23++ {
                            v24 := GenS1{F0: 1, F1: 2}
                            var v25 int64 = -9223372036854775808
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                    }
                    for fi26 := 0; fi26 < 8; fi26++ {
                        if false {
                            if (0.1 != (3.14159 / ((1.5 - float32(-2.25)) / 3.14159))) {
                                v27 := GenS2{F0: 543, F1: (!(false && (v16 == v16)))}
                                fmt.Println(fi26, fi26, fi18)
                                _ = v27
                            }
                            v28 := v2["\n\t"]
                            switch (-(fi26)) {
                            case 32767:
                                if true {
                                    var v29 uint8 = (^(v16))
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    _ = v29
                                }
                            }
                            _ = v28
                        }
                        _ = fi26
                    }
                    var v30 uint64 = 4294967295
                    _ = v30
                } else {
                    v2["a"] = "\n\t"
                }
            }
            v1 += (^(v3))
            _ = fi18
            _ = v19
        }
        _ = len(v2)
    }
    v31 := GenS3{F0: -504, F1: v16, F2: v3}
    if (!(false || true)) {
        var v32 uint64 = 327
        var v33 int32 = 2147483647
        var v34 uint16 = 255
        _ = v32
        _ = v33
        _ = v34
    }
    v35 := []int{32768, -668}
    v36 := GenS1{F0: 550, F1: 2147483647}
    if true {
        for fi37 := 0; fi37 < 15; fi37++ {
            v1 = v3
            _ = fi37
        }
    } else {
        v38 := GenS2{F0: 2, F1: true}
        switch 181 {
        case 2147483647:
            v39 := GenS3{F0: 320, F1: v16, F2: (^(v1))}
            _ = v39
        default:
            var v40 uint = 350
            _ = v40
        }
        _ = v38
    }
    for fi41 := 0; fi41 < 20; fi41++ {
        fmt.Println(fi41, v3)
        _ = fi41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
}
