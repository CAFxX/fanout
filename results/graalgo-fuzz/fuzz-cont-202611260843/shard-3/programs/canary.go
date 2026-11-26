package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 string; F1 int16; F2 bool }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 int; F1 int }
func genFn0(p0 int32, p1 uint64, p2 bool) uint8 {
    for fi5 := 0; fi5 < 10; fi5++ {
        p0 -= p0
        _ = fi5
    }
    switch "\n\t" {
    case "\n\t":
        v6 := GenS1{F0: 42, F1: 255, F2: 2}
        _ = v6
    }
    var v7 uint64 = p1
    v8 := GenS1{F0: -102, F1: 127, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 string, p1 uint) uint32 {
    var v9 uint32 = 0
    p1 = p1
    _ = p0
    _ = p1
    _ = v9
    return (^(v9))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "foo bar"
    var v3 uint8 = 255
    var v4 uint = 12
    fmt.Println(v2)
    v2 = (("hello" + "foo bar") + ("" + "a"))
    if true {
        switch "foo bar" {
        case "\n\t":
            v10 := []int{128, -491, 128, -9223372036854775808}
            _ = v10
        case "hello":
            for fi11 := 0; fi11 < 19; fi11++ {
                if ((((float32(0.0) + float32(3.14159)) < 1.5) && false) && true) {
                    if false {
                        switch "hello" {
                        case "foo bar":
                            v2 = v2
                        default:
                            switch v2 {
                            case "hello":
                                v12 := GenS2{F0: (("\n\t" + v2) + v2), F1: -128}
                                _ = v12
                            case "a":
                                v13 := GenS3{F0: "foo bar", F1: -590, F2: true}
                                _ = v13
                            case "foo bar":
                                switch (v2 + (v2 + ("\n\t" + v2))) {
                                case "\n\t":
                                    var v14 uint8 = (v3 - (v3 ^ (v3 % (((v3 | (v3 + v3))) | 1))))
                                    _ = v14
                                }
                            }
                        }
                        v4 = v4
                    }
                    v4 = (v4 | v4)
                } else {
                    fmt.Println(v2, v2, v3)
                    for fi15 := 0; fi15 < 9; fi15++ {
                        if true {
                            v16 := GenS4{F0: 255, F1: 2147483648}
                            v4 += v4
                            var v17 uint16 = 255
                            v3 += (v3 >> ((v4) & 63))
                            var v18 int = (fi15 ^ (fi11 << (((v4 / ((v4) | 1))) & 63)))
                            _ = v16
                            _ = v17
                            _ = v18
                        } else {
                            if (true || true) {
                                var v19 uint32 = uint32(v3)
                                if (false || true) {
                                    delete(v1, (^(v19)))
                                } else {
                                    var v20 uint = 4294967295
                                    v4 = v20
                                    var v21 uint = v4
                                    _ = v20
                                    _ = v21
                                }
                                if false {
                                    v2 = v2
                                }
                                _ = v19
                            } else {
                                v22 := []int{fi11, fi15}
                                _ = v22
                            }
                            v4 -= v4
                            var v23 int8 = -128
                            _ = v23
                        }
                        _ = fi15
                    }
                }
                var v24 uint8 = (v3 << (((v4 - 278)) & 63))
                switch v2 {
                case "":
                    v1[65535] = v2
                case "foo bar":
                    if true {
                        fmt.Println(v4, v3)
                    } else {
                        var v25 bool = (!true)
                        if (v25 && v25) {
                            switch ((v2 + "\n\t") + ((v2 + v2) + v2)) {
                            case "hello":
                                if (184 >= 2) {
                                    fmt.Println(v3, v24)
                                    var v26 int8 = 42
                                    _ = v26
                                }
                            case "a":
                                fmt.Println(v3, v24)
                            default:
                                v4 = v4
                            }
                            v27 := GenS5{F0: (v24 | (^(v3)))}
                            if ((^(v3)) == 1) {
                                v24 -= genFn0(-2147483648, 0, v25)
                            } else {
                                var v28 string = ("" + "foo bar")
                                if (!false) {
                                    var v29 uint = v4
                                    _ = v29
                                }
                                var v30 uint8 = v3
                                _ = v28
                                _ = v30
                            }
                            _ = v27
                        } else {
                            fmt.Println(v2, v24)
                            for fi31 := 0; fi31 < 6; fi31++ {
                                fmt.Println(v25, fi31, v25)
                                v25 = v25
                                v1[2147483647] = ""
                                _ = fi31
                            }
                            v32 := GenS6{F0: 892}
                            _ = v32
                        }
                        _ = v25
                    }
                }
                _ = fi11
                _ = v24
            }
        case "foo bar":
            v3 = v3
        }
        var v33 float32 = ((0.1 * float32(1.5)) - 0.0)
        _ = v33
    } else {
        v3 -= (v3 & 255)
        v2 = v2
    }
    v4 = v4
    if true {
        var v34 float64 = -2.25
        v3 -= 1
        _ = v34
    } else {
        v35 := GenS7{F0: 65535, F1: 855}
        _ = v35
    }
    var v36 int8 = -1
    v36 = (v36 / (((v36 ^ v36)) | 1))
    if (!false) {
        var v37 int16 = -237
        v3 = v3
        v38 := GenS8{F0: 127, F1: 32767}
        _ = v37
        _ = v38
    } else {
        v4 = (v4 + v4)
    }
    v2 = v2
    var v39 int16 = 1000
    var v40 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v40
}
