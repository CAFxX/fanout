package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint16; F2 int16 }
type GenS3 struct { F0 int }

func main() {
    var v1 uint8 = 255
    var v2 uint16 = 32767
    var v3 int64 = -9223372036854775808
    v1 += (v1 % ((42) | 1))
    fmt.Println(v2)
    for fi4 := 0; fi4 < 14; fi4++ {
        var v5 bool = true
        v5 = true
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 string = (("\n\t" + ("\n\t" + ("\n\t" + "\n\t"))) + ("\n\t" + ""))
            fmt.Println(v2)
            _ = fi6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    v8 := []int{-2147483648, -925}
    if true {
        if false {
            var v9 bool = false
            v10 := GenS1{F0: (v3 > -1)}
            _ = v10
            _ = v9
        }
    }
    v11 := GenS2{F0: 0, F1: 50, F2: -411}
    v3 = (v3 % (((v3 * (v3 % ((1) | 1)))) | 1))
    if false {
        for fi12 := 0; fi12 < 7; fi12++ {
            v2 = (v2 ^ v2)
            for fi13 := 0; fi13 < 8; fi13++ {
                fmt.Println(v1, v2, v1)
                _ = fi13
            }
            _ = fi12
        }
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 uint = 2
            _ = fi14
            _ = v15
        }
        v2 = (v2 / (((v2 >> ((42) & 63))) | 1))
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        if (fi16 == int(v1)) {
            switch "a" {
            case "hello":
                fmt.Println(fi16)
            }
            switch ("" + "") {
            case "hello":
                if false {
                    v8 = append(v8, fi16)
                }
            }
            for fi17 := 0; fi17 < 12; fi17++ {
                if true {
                    var v18 string = "a"
                    v1 -= (v1 | uint8(v2))
                    _ = v18
                } else {
                    if true {
                        if (1 >= 42) {
                            fmt.Println(fi17)
                            fmt.Println(v1, fi17)
                        } else {
                            var v19 uint32 = 981
                            _ = v19
                        }
                        if (255 > 65535) {
                            var v20 int16 = 1000
                            v21 := GenS3{F0: (-(fi16))}
                            if false {
                                v2 = 963
                                var v22 uint8 = v1
                                _ = v22
                            } else {
                                fmt.Println(v1)
                                if true {
                                    v23 := GenS2{F0: v20, F1: v2, F2: (v20 << ((1) & 63))}
                                    _ = v23
                                }
                            }
                            _ = v20
                            _ = v21
                        } else {
                            switch (-(fi17)) {
                            case 42:
                                switch fi16 {
                                case 128:
                                    v2 += (v2 * v2)
                                default:
                                    v2 += v2
                                }
                            case 96:
                                var v24 uint8 = 64
                                _ = v24
                            case 9223372036854775807:
                                var v25 int64 = 266
                                _ = v25
                            }
                        }
                    } else {
                        if (-2147483648 >= 2147483647) {
                            var v26 uint32 = 65535
                            _ = v26
                        } else {
                            if (true || false) {
                                v1 -= (^(v1))
                                if true {
                                    v2 = 304
                                    if false {
                                        if true {
                                            if ((4294967295 > 255) && true) {
                                                var v27 int64 = (^(v3))
                                                _ = v27
                                            }
                                            v3 -= (v3 * v3)
                                        } else {
                                            v3 = v3
                                            var v28 uint16 = 32767
                                            _ = v28
                                        }
                                    }
                                }
                            }
                            var v29 string = "\n\t"
                            v30 := GenS3{F0: fi17}
                            _ = v29
                            _ = v30
                        }
                        var v31 int16 = 1000
                        switch ("hello" + "hello") {
                        case "foo bar":
                            v2 -= (^(v2))
                        case "":
                            var v32 float32 = 0.1
                            _ = v32
                        case "hello":
                            if (!(v3 > (v3 - v3))) {
                                var v33 int32 = -2147483648
                                _ = v33
                            }
                        }
                        _ = v31
                    }
                    if false {
                        if true {
                            v34 := GenS2{F0: 32767, F1: v2, F2: 1}
                            var v35 int64 = v3
                            _ = v34
                            _ = v35
                        }
                    } else {
                        v36 := GenS2{F0: -32768, F1: 42, F2: 594}
                        v1 += v1
                        _ = v36
                    }
                }
                _ = fi17
            }
        }
        v37 := GenS2{F0: 0, F1: v2, F2: 32767}
        v1 += 128
        _ = fi16
        _ = v37
    }
    var v38 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v38
    _ = v8
}
