package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 bool; F1 int16; F2 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16 }

func main() {
    var v1 uint = 2
    v2 := []string{("a" + "\n\t"), (("" + "hello") + "\n\t"), "foo bar"}
    v3 := []int{-282, 825, 343, 2147483647}
    v4 := []uint32{}
    v5 := make(map[string]bool)
    v6 := make(map[int]int)
    _ = len(v6)
    delete(v5, "hello")
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v7 uint32 = 221
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v7, v1)
        var v8 uint64 = uint64(v7)
        _ = v8
    } else {
        v7 += v7
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v7 = v7
    v9 := GenS1{F0: 40}
    switch -1 {
    case 9223372036854775807:
        if (true || false) {
            _ = len(v5)
        }
    default:
        if (2147483648 == v7) {
            v1 -= 540
        }
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        v1 = (^(v1))
        v1 += v1
        v1 = (v1 / ((v1) | 1))
        _ = fi10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        switch "\n\t" {
        case "\n\t":
            var v12 bool = true
            _ = v12
        case "a":
            v1 = v1
        case "hello":
            v13 := GenS2{F0: (v7 & (v7 ^ v7)), F1: 65535}
            _ = v13
        }
        switch (fi11 % ((fi11) | 1)) {
        case 32767:
            v1 += 987
        case 0:
            if ((627 >= -1) && ((-465 >= 1) || true)) {
                for fi14 := 0; fi14 < 10; fi14++ {
                    switch ("\n\t" + ("a" + ("foo bar" + "hello"))) {
                    case "foo bar":
                        v15 := GenS3{F0: 255}
                        _ = v15
                    }
                    fmt.Println(v1, v1)
                    v16 := GenS4{F0: 2, F1: "foo bar"}
                    v17 := GenS1{F0: (fi11 + (fi14 ^ 32768))}
                    _ = fi14
                    _ = v16
                    _ = v17
                }
                v4 = append(v4, v7)
            } else {
                for fi18 := 0; fi18 < 1; fi18++ {
                    if false {
                        if false {
                            if len(v4) > 0 {
                                v4[0] = (v7 << ((1) & 63))
                            }
                            v2 = append(v2, "foo bar")
                        }
                        switch "a" {
                        case "\n\t":
                            switch "foo bar" {
                            case "hello":
                                switch (fi18 & fi18) {
                                case 32767:
                                    v7 -= (v7 | 0)
                                case 127:
                                    v19 := GenS5{F0: (1.5 == 0.0), F1: -864, F2: 0}
                                    _ = v19
                                case 32768:
                                    var v20 int16 = int16(v1)
                                    _ = v20
                                default:
                                    if true {
                                        fmt.Println(v7, fi18)
                                    } else {
                                        if false {
                                            switch fi18 {
                                            case 32767:
                                                var v21 int16 = 1
                                                _ = v21
                                            case 128:
                                                var v22 string = ("a" + ("foo bar" + (("\n\t" + "a") + ("a" + "foo bar"))))
                                                _ = v22
                                            case -9223372036854775808:
                                                var v23 uint32 = (^(v7))
                                                _ = v23
                                            }
                                        } else {
                                            switch "\n\t" {
                                            case "a":
                                                var v24 float32 = float32((0.1 / 1.0))
                                                _ = v24
                                            case "hello":
                                                var v25 float64 = 3.14159
                                                _ = v25
                                            default:
                                                var v26 int32 = -2147483648
                                                _ = v26
                                            }
                                            fmt.Println(v1, v7, fi18)
                                        }
                                        var v27 bool = true
                                        _ = v27
                                    }
                                }
                                v28 := GenS6{F0: 127}
                                _ = v28
                            case "\n\t":
                                if true {
                                    _ = len(v6)
                                    var v29 int8 = -128
                                    var v30 uint8 = 42
                                    _ = v29
                                    _ = v30
                                } else {
                                    v31 := GenS7{F0: -551}
                                    _ = v31
                                }
                            case "foo bar":
                                var v32 int16 = 1000
                                _ = v32
                            }
                        }
                    } else {
                        switch ("hello" + ("\n\t" + ("hello" + ("foo bar" + ("\n\t" + ""))))) {
                        case "hello":
                            if true {
                                v7 += (v7 & v7)
                            }
                        }
                        var v33 int16 = 1
                        var v34 int32 = int32(v7)
                        _ = v33
                        _ = v34
                    }
                    v7 -= (^(v7))
                    _ = fi18
                }
                for fi35 := 0; fi35 < 2; fi35++ {
                    fmt.Println(fi11)
                    v36 := GenS3{F0: 42}
                    _ = fi35
                    _ = v36
                }
            }
        case 522:
            v1 += (v1 * (v1 % ((uint(v7)) | 1)))
        default:
            fmt.Println(v1, fi11)
        }
        var v37 bool = (4294967295 != 389)
        _ = fi11
        _ = v37
    }
    var v38 uint16 = 550
    for fi39 := 0; fi39 < 15; fi39++ {
        delete(v6, fi39)
        v7 += (v7 >> ((v1) & 63))
        _ = fi39
    }
    var v40 float64 = -2.25
    v5[(("\n\t" + ("hello" + ("\n\t" + "\n\t"))) + "\n\t")] = true
    fmt.Println(v40)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
