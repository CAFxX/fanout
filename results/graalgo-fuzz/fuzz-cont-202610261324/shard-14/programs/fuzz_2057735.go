package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint16; F1 uint; F2 int32 }
type GenS4 struct { F0 int; F1 int64; F2 uint32 }
type GenS5 struct { F0 uint32; F1 uint; F2 bool }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 int8; F1 int16 }

func main() {
    v1 := []uint32{65535}
    var v2 bool = false
    var v3 uint8 = 0
    var v4 int16 = -632
    var v5 float32 = 1.5
    if v2 {
        fmt.Println(v4, v5)
    }
    v2 = (!(v2 || (-1 == -2147483648)))
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(v2, v4, fi6)
        _ = fi6
    }
    if (-43 > 1) {
        v3 = 127
        var v7 uint = 4294967295
        _ = len(v1)
        _ = v7
    }
    switch ("hello" + ("\n\t" + ("hello" + ((("a" + "foo bar") + "hello") + ("foo bar" + "foo bar"))))) {
    case "":
        if v2 {
            var v8 int32 = int32(v3)
            v4 -= 1000
            v9 := GenS1{F0: 1, F1: -876}
            _ = v8
            _ = v9
        }
    case "hello":
        v10 := GenS2{F0: uint16(v3), F1: "foo bar"}
        _ = v10
    case "a":
        switch "foo bar" {
        case "":
            var v11 int = 2147483647
            _ = v11
        case "hello":
            v12 := GenS3{F0: 42, F1: 42, F2: 42}
            _ = v12
        default:
            v3 += (v3 << ((uint(v3)) & 63))
        }
    }
    if (v2 && (v2 && (v2 || v2))) {
        if false {
            fmt.Println(v2)
            v13 := GenS4{F0: -9223372036854775808, F1: 1, F2: 387}
            var v14 uint8 = v3
            _ = v13
            _ = v14
        } else {
            var v15 int = 9223372036854775807
            _ = v15
        }
        v16 := GenS5{F0: 0, F1: 973, F2: true}
        _ = v16
    } else {
        var v17 bool = (401 > 2147483648)
        if v17 {
            switch -9223372036854775808 {
            case -9223372036854775808:
                if true {
                    fmt.Println(v3)
                    var v18 bool = false
                    if (4294967295 == 42) {
                        fmt.Println(v17, v3)
                        if v2 {
                            if (105 != 2147483647) {
                                for fi19 := 0; fi19 < 16; fi19++ {
                                    switch "hello" {
                                    case "":
                                        if len(v1) > 0 {
                                            v1[0] = 1
                                        }
                                    default:
                                        switch "foo bar" {
                                        case "foo bar":
                                            if (!(1.5 >= (1e10 * 3.14159))) {
                                                for fi20 := 0; fi20 < 1; fi20++ {
                                                    fmt.Println(v5)
                                                    v21 := GenS6{F0: 9223372036854775807, F1: v17}
                                                    _ = fi20
                                                    _ = v21
                                                }
                                            }
                                        case "":
                                            v22 := GenS7{F0: 100, F1: v4}
                                            _ = v22
                                        case "hello":
                                            fmt.Println(v2, v18)
                                        default:
                                            switch (fi19 * (fi19 % ((-9223372036854775808) | 1))) {
                                            case -1:
                                                fmt.Println(v3)
                                            case 337:
                                                if (262 >= 0) {
                                                    fmt.Println(v4, v17)
                                                }
                                            }
                                        }
                                    }
                                    if v2 {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        for fi23 := 0; fi23 < 7; fi23++ {
                                            var v24 int16 = v4
                                            _ = fi23
                                            _ = v24
                                        }
                                        var v25 uint64 = 42
                                        _ = v25
                                    } else {
                                        fmt.Println(v2)
                                        var v26 int8 = -120
                                        _ = v26
                                    }
                                    fmt.Println(v17, v4)
                                    _ = fi19
                                }
                                var v27 int16 = v4
                                _ = v27
                            } else {
                                for fi28 := 0; fi28 < 16; fi28++ {
                                    var v29 int32 = -1
                                    _ = fi28
                                    _ = v29
                                }
                                var v30 string = ""
                                switch 4294967295 {
                                case 1:
                                    var v31 int32 = 1
                                    _ = v31
                                case 9223372036854775807:
                                    v4 = (v4 & -1)
                                case 2147483647:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = v30
                            }
                            for fi32 := 0; fi32 < 5; fi32++ {
                                if (0 <= 2147483648) {
                                    var v33 int16 = 42
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    fmt.Println(v2, v3, fi32)
                                    _ = v33
                                }
                                v4 = 575
                                v17 = v2
                                _ = fi32
                            }
                        }
                    }
                    _ = v18
                }
            case 0:
                for fi34 := 0; fi34 < 15; fi34++ {
                    var v35 uint8 = v3
                    _ = fi34
                    _ = v35
                }
            }
        } else {
            v17 = true
        }
        v17 = v2
        _ = v17
    }
    var v36 bool = true
    v3 += 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
