package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 string = "foo bar"
    var v2 int64 = -1
    var v3 int = 42
    var v4 int8 = 0
    var v5 string = v1
    v5 = v5
    v2 = 9223372036854775807
    if true {
        v5 = ((v1 + v1) + ("" + ""))
        for fi6 := 0; fi6 < 17; fi6++ {
            v7 := GenS1{F0: 2}
            _ = fi6
            _ = v7
        }
    } else {
        v8 := []int{-9223372036854775808, -1, v3}
        _ = v8
    }
    fmt.Println(v2, v1)
    switch (((("" + "a") + (v1 + "\n\t")) + v5) + v5) {
    case "foo bar":
        for fi9 := 0; fi9 < 14; fi9++ {
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 uint32 = 65535
                if true {
                    v4 -= (-(v4))
                    fmt.Println(v3)
                }
                _ = fi10
                _ = v11
            }
            switch ((("hello" + (v1 + "foo bar")) + v5) + v5) {
            case "foo bar":
                v1 = "hello"
            }
            _ = fi9
        }
    }
    v12 := GenS2{F0: (399 != 692)}
    v1 = v5
    var v13 uint32 = 20
    v14 := []int{(-(v3))}
    for fi15 := 0; fi15 < 11; fi15++ {
        if true {
            if true {
                fmt.Println(v5)
                var v16 uint = 1
                var v17 uint64 = 42
                _ = v16
                _ = v17
            } else {
                if (true || false) {
                    if (255 == 42) {
                        if true {
                            switch v1 {
                            case "":
                                v18 := GenS3{F0: 128, F1: v5}
                                _ = v18
                            case "foo bar":
                                for fi19 := 0; fi19 < 2; fi19++ {
                                    var v20 uint = 4294967295
                                    v2 -= v2
                                    v4 = int8(fi15)
                                    _ = fi19
                                    _ = v20
                                }
                            case "hello":
                                v14 = append(v14, 127)
                            default:
                                switch (v3 - fi15) {
                                case 32767:
                                    for fi21 := 0; fi21 < 8; fi21++ {
                                        fmt.Println(v1)
                                        _ = len(v14)
                                        _ = fi21
                                    }
                                case -9223372036854775808:
                                    fmt.Println(v13, v1)
                                case 128:
                                    v1 = v1
                                default:
                                    switch (("" + (("a" + v1) + "")) + "") {
                                    case "\n\t":
                                        for fi22 := 0; fi22 < 8; fi22++ {
                                            if true {
                                                var v23 int8 = -1
                                                _ = v23
                                            }
                                            _ = fi22
                                        }
                                    case "a":
                                        if (((false || (!false)) || true) || ((("\n\t" + v1) < ("foo bar" + ((v1 + v5) + "foo bar"))) || (float32(float64((0.1 * -2.25))) != 1e10))) {
                                            for fi24 := 0; fi24 < 14; fi24++ {
                                                v4 += (v4 & v4)
                                                _ = fi24
                                            }
                                            v25 := GenS2{F0: false}
                                            _ = v25
                                        } else {
                                            v1 = "\n\t"
                                        }
                                    }
                                }
                            }
                            var v26 uint8 = 1
                            v26 += v26
                            _ = v26
                        }
                    }
                }
                var v27 int8 = v4
                var v28 float64 = 1e10
                _ = v27
                _ = v28
            }
            var v29 int64 = -9223372036854775808
            _ = v29
        }
        for fi30 := 0; fi30 < 12; fi30++ {
            v31 := GenS4{F0: 65535, F1: 1}
            _ = fi30
            _ = v31
        }
        var v32 uint = 42
        _ = fi15
        _ = v32
    }
    _ = len(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
