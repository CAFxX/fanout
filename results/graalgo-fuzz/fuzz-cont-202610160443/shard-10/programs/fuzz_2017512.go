package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 int16; F2 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int }

func main() {
    var v1 uint32 = 65535
    var v2 string = ("hello" + "foo bar")
    var v3 string = v2
    var v4 uint8 = 247
    var v5 uint8 = v4
    var v6 float32 = 3.14159
    var v7 int = int(v1)
    fmt.Println(v6)
    v7 = v7
    var v8 int32 = 2147483647
    v9 := GenS1{F0: 1}
    v8 -= int32(v7)
    v1 = (^(v1))
    v4 = (v4 << ((18446744073709551615) & 63))
    v10 := []int{128, -2147483648, v7, 32768}
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := GenS1{F0: (^(v8))}
        v2 = v2
        _ = fi11
        _ = v12
    }
    if true {
        v2 = "foo bar"
        v7 -= 32767
        if false {
            v5 = 34
            switch (v7 & v7) {
            case 0:
                var v13 int8 = -128
                _ = v13
            case 2147483647:
                var v14 bool = ((((v1 * (^(v1))) < 1) && (!false)) || (2147483647 >= v1))
                _ = v14
            case -2147483648:
                var v15 uint8 = v4
                _ = v15
            default:
                if false {
                    v16 := GenS1{F0: 1}
                    for fi17 := 0; fi17 < 17; fi17++ {
                        v3 = v3
                        v3 = "foo bar"
                        _ = fi17
                    }
                    _ = v16
                } else {
                    if false {
                        v5 = 2
                        v18 := GenS2{F0: 255}
                        for fi19 := 0; fi19 < 16; fi19++ {
                            var v20 float32 = 1.5
                            _ = fi19
                            _ = v20
                        }
                        _ = v18
                    } else {
                        v21 := GenS3{F0: (^(v8))}
                        for fi22 := 0; fi22 < 13; fi22++ {
                            switch -64 {
                            case -1:
                                if true {
                                    if ((-(v8)) != v8) {
                                        v7 = 260
                                        fmt.Println(v4, v5)
                                        v23 := GenS3{F0: (v8 % (((v8 % (((v8 + v8)) | 1))) | 1))}
                                        _ = v23
                                    }
                                    var v24 int32 = v8
                                    if (false && (2 < uint16(v4))) {
                                        var v25 uint32 = v1
                                        v26 := GenS4{F0: (^(v5)), F1: 1, F2: 42}
                                        v2 = ("\n\t" + "a")
                                        _ = v25
                                        _ = v26
                                    } else {
                                        v24 = v8
                                        fmt.Println(v5, v5)
                                    }
                                    _ = v24
                                } else {
                                    var v27 uint = 701
                                    switch v7 {
                                    case -605:
                                        switch v3 {
                                        case "":
                                            fmt.Println(v4, v7, v3)
                                        case "a":
                                            var v28 uint64 = 16
                                            _ = v28
                                        default:
                                            v29 := GenS3{F0: v8}
                                            _ = v29
                                        }
                                    }
                                    _ = v27
                                }
                            case 1:
                                for fi30 := 0; fi30 < 8; fi30++ {
                                    var v31 string = v2
                                    v2 = v2
                                    _ = fi30
                                    _ = v31
                                }
                            case 2147483647:
                                var v32 int32 = (v8 * (v8 % ((v8) | 1)))
                                _ = v32
                            }
                            _ = fi22
                        }
                        _ = v21
                    }
                    for fi33 := 0; fi33 < 7; fi33++ {
                        var v34 int32 = v8
                        for fi35 := 0; fi35 < 11; fi35++ {
                            v36 := GenS4{F0: uint8(v1), F1: 1, F2: -1}
                            var v37 bool = true
                            v34 -= v8
                            var v38 uint8 = v5
                            v39 := GenS2{F0: 42}
                            var v40 int8 = 90
                            if (v37 && (true && false)) {
                                v41 := GenS5{F0: "foo bar"}
                                switch "\n\t" {
                                case "foo bar":
                                    v8 = -2147483648
                                case "hello":
                                    switch "a" {
                                    case "\n\t":
                                        if len(v10) > 0 {
                                            v10[0] = 128
                                        }
                                    case "":
                                        fmt.Println(v4, v3)
                                    }
                                case "":
                                    v5 = (v4 << ((865) & 63))
                                }
                                var v42 int16 = -21
                                _ = v41
                                _ = v42
                            } else {
                                v2 = v2
                                v40 = v40
                                var v43 uint16 = 430
                                _ = v43
                            }
                            _ = fi35
                            _ = v36
                            _ = v37
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                        fmt.Println(v4, v3, v8)
                        _ = fi33
                        _ = v34
                    }
                }
            }
        } else {
            var v44 uint8 = 127
            var v45 int32 = (v8 ^ v8)
            _ = v44
            _ = v45
        }
    }
    v3 = ((v3 + "") + ((v3 + "") + "foo bar"))
    v7 = 0
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    fmt.Println(v4, v1, v2)
    var v46 uint = 733
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v47 := GenS6{F0: v7}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
