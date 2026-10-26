package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint; F2 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 uint8; F1 bool }

func main() {
    v1 := []uint32{352, 4294967295, 65535, 42}
    v2 := make(map[uint32]bool)
    var v3 uint32 = 4294967295
    var v4 uint = 255
    v3 = (v3 << ((v4) & 63))
    if false {
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 int = 685
            var v7 uint8 = 127
            _ = fi5
            _ = v6
            _ = v7
        }
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 int = fi8
            for fi10 := 0; fi10 < 15; fi10++ {
                fmt.Println(fi10, fi8, v9)
                if (9 > 128) {
                    v9 += (fi8 % (((fi10 - 4294967295)) | 1))
                    if (!(2 == uint8(v4))) {
                        v9 = (fi8 << (((v4 / (((v4 >> ((v4) & 63))) | 1))) & 63))
                        var v11 float64 = 3.14159
                        v12 := GenS1{F0: (v4 ^ v4), F1: (v4 * v4), F2: 0}
                        _ = v11
                        _ = v12
                    } else {
                        var v13 uint = v4
                        var v14 uint32 = v3
                        v15 := GenS1{F0: (v4 + (^(v4))), F1: 4294967295, F2: 42}
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    fmt.Println(fi10, v3)
                } else {
                    if false {
                        v16 := GenS2{F0: (("" + ("" + "hello")) + "hello")}
                        v4 += (v4 % (((v4 >> (((v4 >> ((45) & 63))) & 63))) | 1))
                        _ = v16
                    } else {
                        var v17 bool = false
                        var v18 uint64 = 2
                        v18 -= v18
                        if (uint16(v4) >= uint16(v4)) {
                            if ((v3 > 2147483647) && (!(-2.25 == 0.0))) {
                                if true {
                                    var v19 float64 = -2.25
                                    if len(v1) > 0 {
                                        v1[0] = v3
                                    }
                                    _ = v19
                                }
                            }
                            var v20 float64 = -2.25
                            _ = v20
                        }
                        _ = v17
                        _ = v18
                    }
                    v21 := GenS3{F0: 758, F1: 975, F2: ("\n\t" + "foo bar")}
                    _ = v21
                }
                fmt.Println(v4, fi10)
                if (65535 <= v4) {
                    switch (("\n\t" + ("foo bar" + ("foo bar" + "hello"))) + ("hello" + "foo bar")) {
                    case "foo bar":
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                    }
                }
                _ = fi10
            }
            if false {
                v9 = (fi8 / (((^(fi8))) | 1))
                switch "\n\t" {
                case "hello":
                    var v22 uint = uint(v3)
                    _ = v22
                }
            } else {
                var v23 uint8 = 53
                if true {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch ("a" + "a") {
                    case "hello":
                        if (false && false) {
                            v3 = (v3 % ((v3) | 1))
                            v24 := GenS4{F0: 0}
                            _ = len(v2)
                            _ = v24
                        }
                    }
                } else {
                    v3 = 26
                }
                if (!(((false || true) || ((0.0 * 1.5) > 1.5)) || (((!(-882 <= 1)) && false) && (!(false && false))))) {
                    v25 := GenS2{F0: ("\n\t" + "\n\t")}
                    _ = v25
                }
                _ = v23
            }
            _ = fi8
            _ = v9
        }
        var v26 float32 = 3.14159
        _ = v26
    }
    var v27 bool = true
    v27 = (v4 != 2)
    v27 = (1 != 0)
    v3 = 2147483647
    var v28 uint8 = 127
    if (v27 && (0 <= 978)) {
        v27 = ((!v27) && false)
    } else {
        if true {
            var v29 uint8 = v28
            v3 += v3
            for fi30 := 0; fi30 < 19; fi30++ {
                if (!v27) {
                    v31 := GenS5{F0: v29, F1: (v28 + (v28 + v28))}
                    fmt.Println(v27, v3, v4)
                    v2[v3] = (((1.5 / 1.5) * (1e10 + -2.25)) == (-2.25 - 1.5))
                    _ = v31
                }
                _ = fi30
            }
            _ = v29
        }
    }
    if false {
        if (!("" != "a")) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v1 = append(v1, (v3 + (^(v3))))
            v32 := GenS5{F0: 42, F1: (^(v28))}
            _ = v32
        }
        v3 += 2
        var v33 uint64 = 2
        _ = v33
    } else {
        fmt.Println(v4, v3)
    }
    v1 = append(v1, v3)
    switch 9223372036854775807 {
    case -433:
        fmt.Println(v3)
    }
    if false {
        fmt.Println(v27, v3, v4)
        if false {
            v34 := GenS6{F0: v28, F1: false}
            _ = v34
        }
        if true {
            fmt.Println(v28)
        } else {
            fmt.Println(v3, v3)
        }
    }
    fmt.Println(v4)
    var v35 uint64 = 551
    var v36 uint16 = 65535
    var v37 uint32 = (^(v3))
    v2[v37] = (v27 && v27)
    switch "" {
    case "a":
        v38 := GenS2{F0: ("foo bar" + ((("foo bar" + "\n\t") + ("" + "\n\t")) + "foo bar"))}
        _ = v38
    default:
        var v39 bool = (42 != 42)
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
