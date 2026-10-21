package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int16 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]string)
    var v3 float64 = 0.0
    v3 -= 0.0
    v4 := v2[733]
    if (!false) {
        if false {
            switch v4 {
            case "":
                for fi5 := 0; fi5 < 18; fi5++ {
                    v6 := GenS1{F0: true}
                    v7 := GenS1{F0: (1e10 <= 3.14159)}
                    _ = fi5
                    _ = v6
                    _ = v7
                }
            case "\n\t":
                for fi8 := 0; fi8 < 3; fi8++ {
                    for fi9 := 0; fi9 < 2; fi9++ {
                        v1[v4] = ((("" + "foo bar") + (v4 + v4)) + "hello")
                        if true {
                            v10 := GenS2{F0: v4, F1: -1}
                            if (875 == 1) {
                                var v11 uint32 = 0
                                _ = v11
                            }
                            v4 = v4
                            if ((0 < 0) && false) {
                                delete(v1, "\n\t")
                                v3 = v3
                            } else {
                                if (795 < 567) {
                                    v4 = v4
                                } else {
                                    v12 := GenS1{F0: (127 >= (fi8 >> ((1) & 63)))}
                                    _ = v12
                                }
                                v13 := GenS1{F0: false}
                                v14 := v1[(("foo bar" + v4) + ("\n\t" + v4))]
                                var v15 uint8 = 128
                                v16 := []int{fi8, (fi9 % ((fi8) | 1)), fi9}
                                _ = v13
                                _ = v14
                                _ = v15
                                _ = v16
                            }
                            _ = v10
                        } else {
                            switch v4 {
                            case "foo bar":
                                fmt.Println(v4, v4, v3)
                            case "hello":
                                if (42 <= 0) {
                                    v3 += -2.25
                                    v3 = 3.14159
                                } else {
                                    v3 = (-2.25 / v3)
                                    var v17 uint32 = 4294967295
                                    _ = v17
                                }
                            case "\n\t":
                                v18 := GenS3{F0: 4294967295, F1: 1}
                                _ = v18
                            }
                        }
                        _ = fi9
                    }
                    v19 := []int{(^(fi8)), fi8, 235}
                    v4 = v4
                    _ = fi8
                    _ = v19
                }
            }
            v3 -= v3
            v4 = (v4 + (v4 + (((v4 + "a") + (v4 + ("\n\t" + "a"))) + (v4 + v4))))
        } else {
            _ = len(v2)
            for fi20 := 0; fi20 < 1; fi20++ {
                switch (fi20 * (fi20 - 521)) {
                case 127:
                    v21 := GenS4{F0: 255, F1: (1 > 4294967295)}
                    _ = v21
                }
                var v22 int16 = 42
                _ = fi20
                _ = v22
            }
        }
    } else {
        var v23 float64 = ((float64(0.0) / ((float64(-2.25) / 1.0) * (float64(0.1) - float64(-2.25)))) + 0.0)
        v24 := GenS5{F0: 32767}
        _ = v23
        _ = v24
    }
    v25 := v1[("hello" + ("foo bar" + (v4 + (v4 + ("a" + ((v4 + v4) + (((v4 + "a") + "foo bar") + "\n\t")))))))]
    switch -547 {
    case 32767:
        if (((4294967295 >= -2147483648) || ("a" == "foo bar")) || (1 > 1)) {
            v3 -= v3
        } else {
            var v26 bool = (18446744073709551615 <= 1)
            _ = v26
        }
    case 4294967295:
        fmt.Println(v3, v4)
    case 0:
        switch v25 {
        case "hello":
            if (42 != 0) {
                v25 = ""
            }
        default:
            v27 := GenS2{F0: (("\n\t" + (v25 + v25)) + (v25 + v4)), F1: 42}
            _ = v27
        }
    }
    if false {
        v3 += 0.1
        delete(v1, ((v4 + ("" + "a")) + ("foo bar" + "foo bar")))
        var v28 int16 = 1
        _ = v28
    } else {
        v25 = "foo bar"
    }
    var v29 uint64 = 1
    switch (v4 + v4) {
    case "\n\t":
        v25 = v25
    case "foo bar":
        v29 = (^(v29))
    case "a":
        for fi30 := 0; fi30 < 20; fi30++ {
            v31 := v1["a"]
            _ = fi30
            _ = v31
        }
    }
    v32 := []int{-9223372036854775808, 960}
    fmt.Println(v25, v25, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v32
    _ = v4
}
