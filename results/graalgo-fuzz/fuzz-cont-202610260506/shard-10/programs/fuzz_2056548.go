package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int; F2 uint8 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string; F1 string; F2 uint16 }
type GenS7 struct { F0 uint32; F1 string; F2 uint16 }
type GenS8 struct { F0 int32 }
func genFn0(p0 uint8) string {
    v5 := []int{}
    if false {
        v6 := GenS1{F0: 1}
        _ = v6
    }
    var v7 int16 = 42
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS2{F0: 32767}
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: -227}
    _ = p0
    _ = v10
    _ = v5
    _ = v7
    return "hello"
}

func main() {
    var v1 int32 = -339
    var v2 int16 = -310
    v3 := []int{-371, 2147483647}
    var v4 uint32 = 42
    var v11 int8 = -16
    var v12 bool = false
    if v12 {
        switch ((("foo bar" + "") + (("\n\t" + ("a" + "\n\t")) + "\n\t")) + "hello") {
        case "a":
            v13 := GenS2{F0: 0}
            _ = v13
        case "":
            if true {
                for fi14 := 0; fi14 < 3; fi14++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if v12 {
                        v4 = (v4 / (((v4 & v4)) | 1))
                        if (401 >= 261) {
                            v2 += v2
                            v12 = ((v4 & (v4 - (v4 << ((840) & 63)))) == (^(v4)))
                            for fi15 := 0; fi15 < 8; fi15++ {
                                if (!v12) {
                                    v2 += (v2 / (((v2 + 1)) | 1))
                                    switch "foo bar" {
                                    case "a":
                                        var v16 bool = v12
                                        _ = v16
                                    case "":
                                        var v17 uint8 = 1
                                        _ = v17
                                    }
                                    v4 = v4
                                } else {
                                    if false {
                                        v1 = (v1 | 487)
                                        var v18 uint32 = v4
                                        v4 = v18
                                        v19 := GenS1{F0: 0}
                                        _ = v18
                                        _ = v19
                                    }
                                }
                                v4 += 2147483648
                                _ = fi15
                            }
                        } else {
                            var v20 int32 = (v1 + v1)
                            fmt.Println(v1, v12, v1)
                            v21 := GenS3{F0: 28, F1: fi14, F2: 2}
                            _ = v20
                            _ = v21
                        }
                    }
                    _ = fi14
                }
                for fi22 := 0; fi22 < 20; fi22++ {
                    _ = len(v3)
                    v23 := GenS4{F0: 65535, F1: (808 == 18446744073709551615)}
                    v11 = 100
                    _ = fi22
                    _ = v23
                }
                switch ("a" + "") {
                case "foo bar":
                    fmt.Println(v11)
                case "hello":
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                default:
                    var v24 float32 = 3.14159
                    _ = v24
                }
            } else {
                v2 = v2
            }
        }
        fmt.Println(v12, v11, v4)
        var v25 float64 = ((-2.25 - 0.0) - 1.5)
        _ = v25
    } else {
        var v26 int64 = 42
        _ = v26
    }
    if v12 {
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
    }
    v4 += (v4 | v4)
    var v27 int16 = v2
    var v28 int64 = 9223372036854775807
    switch "\n\t" {
    case "\n\t":
        if true {
            v3 = append(v3, -9223372036854775808)
            for fi29 := 0; fi29 < 17; fi29++ {
                v30 := GenS4{F0: 65535, F1: true}
                if (v11 <= (v11 >> ((569) & 63))) {
                    v4 = (^(v4))
                    switch 4294967295 {
                    case 826:
                        v31 := GenS5{F0: 777}
                        _ = v31
                    case -360:
                        if (-519 <= v28) {
                            var v32 uint = 42
                            if true {
                                v33 := GenS6{F0: (("" + ("a" + "\n\t")) + (("" + "") + ("" + "hello"))), F1: "\n\t", F2: 65535}
                                _ = v33
                            }
                            _ = v32
                        } else {
                            v4 += v4
                        }
                    default:
                        var v34 int64 = 1
                        _ = v34
                    }
                }
                v35 := GenS3{F0: 173, F1: (^(fi29)), F2: 2}
                _ = fi29
                _ = v30
                _ = v35
            }
            v3 = append(v3, -859)
        }
    }
    var v36 int32 = (v1 << ((2) & 63))
    switch -9223372036854775808 {
    case 0:
        for fi37 := 0; fi37 < 3; fi37++ {
            fmt.Println(v2, fi37)
            var v38 int = fi37
            for fi39 := 0; fi39 < 5; fi39++ {
                v40 := GenS7{F0: v4, F1: "", F2: 1}
                var v41 uint32 = v4
                var v42 int8 = (v11 % (((v11 - 83)) | 1))
                _ = fi39
                _ = v40
                _ = v41
                _ = v42
            }
            _ = fi37
            _ = v38
        }
    default:
        if false {
            fmt.Println(v2)
            for fi43 := 0; fi43 < 3; fi43++ {
                switch "hello" {
                case "hello":
                    switch "" {
                    case "\n\t":
                        switch "" {
                        case "foo bar":
                            var v44 int32 = v36
                            _ = v44
                        case "":
                            if false {
                                var v45 string = "hello"
                                _ = v45
                            }
                        default:
                            switch (^(fi43)) {
                            case -1:
                                v46 := GenS8{F0: (-(v1))}
                                _ = v46
                            }
                        }
                    case "a":
                        v4 -= (v4 - 2147483647)
                    default:
                        var v47 int64 = -1
                        _ = v47
                    }
                default:
                    fmt.Println(v36, v27)
                }
                _ = fi43
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v4
}
