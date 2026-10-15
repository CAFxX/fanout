package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 int32; F1 uint8 }

func main() {
    var v1 uint = 1
    v2 := make(map[string]uint32)
    v3 := make(map[string]uint32)
    var v4 uint8 = uint8(v1)
    v5 := []string{("hello" + ("hello" + "")), ("\n\t" + "foo bar"), "foo bar"}
    v6 := []uint32{958, 2, 244, 0}
    v1 = (v1 << ((12) & 63))
    _ = len(v5)
    v7 := GenS1{F0: 498, F1: 2, F2: (v4 | (^(v4)))}
    v8 := GenS2{F0: (float32(((0.0 - -2.25) + -2.25)) > 0.0), F1: 236}
    switch "foo bar" {
    case "foo bar":
        var v9 uint32 = 2147483647
        _ = v9
    case "\n\t":
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := GenS3{F0: 913}
            var v12 uint8 = (v4 + (^(v4)))
            for fi13 := 0; fi13 < 1; fi13++ {
                if (!false) {
                    var v14 float64 = -2.25
                    _ = v14
                } else {
                    v6 = append(v6, 726)
                    if true {
                        var v15 bool = ("hello" >= ("" + (("\n\t" + ("\n\t" + "a")) + ("a" + "\n\t"))))
                        if len(v6) > 0 {
                            v6[0] = 1
                        }
                        _ = v15
                    }
                }
                _ = fi13
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    case "a":
        fmt.Println(v1, v4)
    default:
        var v16 uint8 = (^(v4))
        _ = v16
    }
    var v17 uint32 = 2
    if false {
        v18 := v3["hello"]
        v4 = uint8(v17)
        _ = v18
    }
    fmt.Println(v1)
    v19 := GenS1{F0: 32767, F1: 18446744073709551615, F2: v4}
    for fi20 := 0; fi20 < 20; fi20++ {
        if (3.14159 <= -2.25) {
            var v21 int16 = 32767
            _ = v21
        } else {
            v22 := GenS3{F0: -981}
            switch (("a" + "") + "hello") {
            case "":
                if (!(128 >= fi20)) {
                    var v23 uint64 = 743
                    v5 = append(v5, "\n\t")
                    var v24 float32 = -2.25
                    _ = v23
                    _ = v24
                }
            case "\n\t":
                if true {
                    for fi25 := 0; fi25 < 5; fi25++ {
                        if len(v6) > 0 {
                            v6[0] = v17
                        }
                        switch "a" {
                        case "foo bar":
                            v4 += v4
                        default:
                            switch 4294967295 {
                            case 128:
                                switch ("\n\t" + "a") {
                                case "hello":
                                    var v26 uint32 = (v17 | v17)
                                    _ = v26
                                case "a":
                                    v1 = (v1 & (v1 % (((^(v1))) | 1)))
                                case "foo bar":
                                    var v27 bool = ((!(!(fi20 != fi25))) && true)
                                    _ = v27
                                default:
                                    fmt.Println(fi20, fi20)
                                }
                            default:
                                if (false && (false && (0 <= 42))) {
                                    var v28 bool = (!("hello" >= ("foo bar" + ("foo bar" + ("a" + ("" + "foo bar"))))))
                                    v17 -= 4294967295
                                    var v29 int32 = -312
                                    _ = v28
                                    _ = v29
                                } else {
                                    switch int(v4) {
                                    case 314:
                                        if true {
                                            if len(v6) > 0 {
                                                v6[0] = v17
                                            }
                                            var v30 int8 = -1
                                            _ = v30
                                        } else {
                                            switch ("foo bar" + "") {
                                            case "foo bar":
                                                fmt.Println(v1)
                                            case "\n\t":
                                                v1 -= (^(v1))
                                            default:
                                                if len(v6) > 0 {
                                                    v6[0] = v17
                                                }
                                            }
                                            v31 := GenS4{F0: fi25, F1: (^(fi20))}
                                            _ = v31
                                        }
                                    default:
                                        var v32 bool = true
                                        _ = v32
                                    }
                                    _ = len(v6)
                                }
                            }
                        }
                        v4 = 128
                        v33 := v3[("hello" + "foo bar")]
                        v4 = (^(v4))
                        _ = fi25
                        _ = v33
                    }
                    v34 := GenS4{F0: (fi20 | -816), F1: (fi20 + fi20)}
                    v35 := GenS5{F0: 2147483647, F1: (^(v4))}
                    _ = v34
                    _ = v35
                }
            case "foo bar":
                var v36 uint8 = v4
                _ = v36
            default:
                var v37 int = fi20
                _ = v37
            }
            v17 = 65535
            _ = v22
        }
        _ = fi20
    }
    v1 = (v1 % ((uint(v17)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
