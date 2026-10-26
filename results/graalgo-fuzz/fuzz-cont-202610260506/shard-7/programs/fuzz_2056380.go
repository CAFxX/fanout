package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 uint16 = 42
    v2 := make(map[uint32]int)
    var v3 float64 = -2.25
    if (!("hello" < "foo bar")) {
        fmt.Println(v3, v1)
    }
    for fi4 := 0; fi4 < 6; fi4++ {
        for fi5 := 0; fi5 < 16; fi5++ {
            v3 = 1e10
            v6 := v2[265]
            _ = fi5
            _ = v6
        }
        if ("foo bar" > "\n\t") {
            var v7 float32 = (0.1 - (-2.25 - (float32(v3) - (-2.25 * 3.14159))))
            _ = v7
        } else {
            v1 -= v1
        }
        v2[2147483647] = fi4
        _ = fi4
    }
    var v8 uint = 2
    switch 9223372036854775807 {
    case 32768:
        if true {
            for fi9 := 0; fi9 < 14; fi9++ {
                for fi10 := 0; fi10 < 2; fi10++ {
                    fmt.Println(fi9)
                    _ = fi10
                }
                switch -169 {
                case 128:
                    var v11 float64 = v3
                    _ = v11
                }
                _ = fi9
            }
        }
    case 1:
        v12 := []int{}
        _ = v12
    case -9223372036854775808:
        var v13 int16 = 1000
        _ = v13
    }
    if false {
        v8 += (v8 * v8)
        v8 -= (v8 - v8)
    }
    v8 += v8
    switch ("a" + (((((("foo bar" + (("a" + "\n\t") + (("\n\t" + (("foo bar" + (("a" + ("a" + (("a" + "\n\t") + "foo bar"))) + "hello")) + (((("" + ("\n\t" + "\n\t")) + "") + "hello") + ""))) + "hello"))) + "\n\t") + ("\n\t" + "a")) + "hello") + "") + "")) {
    case "":
        v1 -= (^(v1))
    default:
        fmt.Println(v8, v1, v3)
    }
    if (-94 >= 100) {
        for fi14 := 0; fi14 < 5; fi14++ {
            delete(v2, 994)
            var v15 int = (-(fi14))
            var v16 bool = false
            _ = fi14
            _ = v15
            _ = v16
        }
        v8 -= 477
        v8 = (^(v8))
    } else {
        v17 := GenS1{F0: -610}
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        v3 = float64((float32(1.5) * 0.1))
        for fi19 := 0; fi19 < 2; fi19++ {
            v20 := GenS2{F0: 42, F1: 0, F2: 152}
            if true {
                var v21 int16 = -1
                var v22 bool = (-1 <= -1)
                _ = v21
                _ = v22
            }
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    if true {
        if true {
            fmt.Println(v1, v1)
        }
        var v23 uint16 = (^(v1))
        if true {
            if false {
                for fi24 := 0; fi24 < 9; fi24++ {
                    v8 = 2
                    v1 = v1
                    for fi25 := 0; fi25 < 17; fi25++ {
                        var v26 float64 = 1e10
                        _ = fi25
                        _ = v26
                    }
                    _ = fi24
                }
            }
            switch (("a" + ("" + "")) + ("foo bar" + ("a" + "foo bar"))) {
            case "a":
                v27 := GenS3{F0: 994, F1: 1}
                _ = v27
            case "foo bar":
                v23 += 2
            case "\n\t":
                var v28 uint32 = 57
                _ = v28
            default:
                switch -486 {
                case -2147483648:
                    switch -1 {
                    case 4294967295:
                        for fi29 := 0; fi29 < 17; fi29++ {
                            fmt.Println(v3, v23, v1)
                            v1 = (v23 >> ((v8) & 63))
                            v3 -= float64(1.5)
                            _ = fi29
                        }
                    case -9223372036854775808:
                        for fi30 := 0; fi30 < 17; fi30++ {
                            v3 = float64(((3.14159 * -2.25) * float32(v3)))
                            v31 := GenS1{F0: -816}
                            if (!((((v3 - v3) / (v3 - (v3 - ((v3 - v3) * (3.14159 / float64(3.14159)))))) / (((float64(0.0) - v3) * 0.1) * ((-2.25 - v3) + (float64(0.0) + 0.1)))) > 1.5)) {
                                for fi32 := 0; fi32 < 11; fi32++ {
                                    switch "foo bar" {
                                    case "":
                                        var v33 int32 = 42
                                        _ = v33
                                    case "\n\t":
                                        v34 := GenS1{F0: 727}
                                        _ = v34
                                    }
                                    var v35 string = "hello"
                                    _ = fi32
                                    _ = v35
                                }
                                v36 := GenS2{F0: 32767, F1: 42, F2: 124}
                                var v37 string = "a"
                                _ = v36
                                _ = v37
                            } else {
                                fmt.Println(v3, fi30, fi30)
                                var v38 uint8 = 144
                                _ = v38
                            }
                            _ = fi30
                            _ = v31
                        }
                    default:
                        var v39 uint = (v8 | (v8 ^ 1))
                        _ = v39
                    }
                }
            }
        }
        _ = v23
    } else {
        for fi40 := 0; fi40 < 16; fi40++ {
            _ = len(v2)
            _ = fi40
        }
    }
    var v41 int8 = 100
    v42 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v8
}
