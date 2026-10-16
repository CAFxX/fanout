package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int64; F1 bool; F2 string }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 int8; F2 string }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint8; F1 uint64; F2 bool }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{306, -2147483648, 9223372036854775807, -1}
    v3 := []uint32{264}
    for fi4 := 0; fi4 < 10; fi4++ {
        v5 := GenS1{F0: 0, F1: ((("hello" + "hello") + "foo bar") + "hello")}
        fmt.Println(fi4)
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 int = (fi6 ^ 0)
        for fi8 := 0; fi8 < 6; fi8++ {
            v7 = (fi6 | 0)
            var v9 int32 = 42
            _ = fi8
            _ = v9
        }
        if ((785 >= 2) && false) {
            for fi10 := 0; fi10 < 15; fi10++ {
                v7 = 32767
                fmt.Println(v7, fi10, fi10)
                v1[127] = true
                _ = fi10
            }
            v7 = 2147483647
            switch -251 {
            case -127:
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
            case -1:
                v11 := v1[32767]
                _ = v11
            default:
                var v12 uint8 = 245
                _ = v12
            }
        }
        _ = fi6
        _ = v7
    }
    v13 := GenS2{F0: -326, F1: false, F2: "\n\t"}
    var v14 uint32 = 1
    v15 := GenS1{F0: 1, F1: ("\n\t" + (("a" + (("hello" + "foo bar") + ("" + "foo bar"))) + "foo bar"))}
    switch -356 {
    case 42:
        fmt.Println(v14)
    default:
        if (0 < 437) {
            for fi16 := 0; fi16 < 8; fi16++ {
                if true {
                    v17 := GenS3{F0: 0, F1: 325}
                    _ = v17
                }
                for fi18 := 0; fi18 < 4; fi18++ {
                    switch (fi16 >> ((4294967295) & 63)) {
                    case 9223372036854775807:
                        if true {
                            if (((255 <= 1) || (2 <= 127)) || true) {
                                v14 -= 0
                                if false {
                                    v14 = (v14 + v14)
                                    v19 := GenS3{F0: 1, F1: 4294967295}
                                    _ = v19
                                } else {
                                    v14 = (v14 % ((648) | 1))
                                    if (255 == 0) {
                                        var v20 int32 = 2147483647
                                        if (!false) {
                                            var v21 float64 = 1.5
                                            var v22 uint8 = 133
                                            if true {
                                                v23 := GenS4{F0: (v20 << ((0) & 63))}
                                                v24 := GenS5{F0: -1, F1: -79, F2: "hello"}
                                                v20 = 0
                                                _ = v23
                                                _ = v24
                                            }
                                            _ = len(v1)
                                            var v25 uint = 928
                                            _ = v21
                                            _ = v22
                                            _ = v25
                                        } else {
                                            var v26 int64 = -318
                                            switch (("hello" + "\n\t") + "a") {
                                            case "a":
                                                v14 -= (^(v14))
                                            }
                                            delete(v1, fi18)
                                            v1[fi18] = ((3.14159 - (1.5 - 1.5)) >= float64((3.14159 / 1.0)))
                                            v14 = (v14 << ((255) & 63))
                                            _ = v26
                                        }
                                        fmt.Println(fi18, fi18)
                                        v20 = 42
                                        if false {
                                            if false {
                                                var v27 uint16 = 32767
                                                if false {
                                                    switch "\n\t" {
                                                    case "":
                                                        v20 = (v20 >> ((0) & 63))
                                                    case "hello":
                                                        if false {
                                                            var v28 int16 = -1
                                                            var v29 int32 = v20
                                                            _ = v28
                                                            _ = v29
                                                        } else {
                                                            var v30 int8 = 100
                                                            _ = v30
                                                        }
                                                    case "\n\t":
                                                        v1[-67] = (v27 == (v27 >> ((962) & 63)))
                                                    }
                                                    var v31 string = "hello"
                                                    _ = v31
                                                } else {
                                                    v14 -= v14
                                                    fmt.Println(fi18)
                                                    switch "foo bar" {
                                                    case "a":
                                                        var v32 uint = 42
                                                        _ = v32
                                                    default:
                                                        var v33 string = ("foo bar" + "\n\t")
                                                        _ = v33
                                                    }
                                                }
                                                _ = v27
                                            }
                                            v34 := GenS6{F0: 255, F1: 127}
                                            var v35 int32 = v20
                                            _ = v34
                                            _ = v35
                                        } else {
                                            fmt.Println(v14, v20)
                                            _ = len(v3)
                                            v36 := GenS7{F0: 164, F1: 2, F2: true}
                                            _ = v36
                                        }
                                        _ = v20
                                    } else {
                                        var v37 int32 = -1
                                        v37 += (-(v37))
                                        _ = v37
                                    }
                                    v38 := GenS1{F0: -32768, F1: ""}
                                    _ = v38
                                }
                                var v39 uint16 = 65535
                                _ = v39
                            }
                            v40 := GenS1{F0: 42, F1: "a"}
                            var v41 bool = ("hello" > "hello")
                            v14 -= (v14 ^ (v14 & (v14 + 1)))
                            v42 := v1[128]
                            _ = v40
                            _ = v41
                            _ = v42
                        } else {
                            fmt.Println(fi16)
                            var v43 int64 = 42
                            _ = v43
                        }
                    case 1:
                        v14 = (^(v14))
                    case -767:
                        var v44 bool = true
                        _ = v44
                    }
                    var v45 int8 = 1
                    _ = fi18
                    _ = v45
                }
                _ = fi16
            }
            switch "hello" {
            case "a":
                v46 := GenS4{F0: -2147483648}
                _ = v46
            }
        }
    }
    var v47 int32 = 1
    v14 += v14
    v48 := GenS3{F0: -9223372036854775808, F1: 790}
    var v49 string = "hello"
    if false {
        switch v49 {
        case "":
            v49 = "hello"
        default:
            fmt.Println(v49)
        }
        fmt.Println(v47)
    } else {
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v47
    _ = v48
    _ = v49
}
