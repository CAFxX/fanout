package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 int16 = -32768
    v2 := []uint32{0, 65535}
    var v3 int = 321
    var v4 uint16 = 2
    var v5 int64 = 1
    var v6 int8 = 0
    var v7 uint = 379
    var v8 float32 = (((float32((0.1 * (3.14159 * 3.14159))) + -2.25) * 3.14159) / 3.14159)
    var v9 int16 = v1
    var v10 uint = uint(v4)
    if len(v2) > 0 {
        v2[0] = 1
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        _ = len(v2)
        _ = fi11
    }
    if false {
        v1 = (v9 << (((v10 / ((v7) | 1))) & 63))
    } else {
        if len(v2) > 0 {
            v2[0] = 1
        }
        if (65535 > 2) {
            if len(v2) > 0 {
                v2[0] = 0
            }
        } else {
            _ = len(v2)
            fmt.Println(v3, v3)
        }
    }
    var v12 float32 = v8
    for fi13 := 0; fi13 < 4; fi13++ {
        fmt.Println(fi13, v4, v5)
        var v14 string = "a"
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        if (!(false || (false && ((v9 / ((v9) | 1)) >= v1)))) {
            v8 -= -2.25
            if false {
                v16 := GenS1{F0: (("" + "a") + "a"), F1: (2 > uint8(v4))}
                _ = v16
            }
            if (((((("" + "\n\t") + ("foo bar" + ("hello" + ("" + (("" + "\n\t") + ""))))) + "hello") + "\n\t") + "a") < "\n\t") {
                fmt.Println(v4, v7)
                if false {
                    switch 42 {
                    case 42:
                        var v17 int16 = 1
                        _ = v17
                    default:
                        var v18 uint64 = 4294967295
                        _ = v18
                    }
                    if (true || false) {
                        v10 = 1
                    } else {
                        fmt.Println(fi15, v8, v10)
                        if false {
                            var v19 int32 = 0
                            _ = v19
                        } else {
                            var v20 uint8 = 42
                            _ = v20
                        }
                        if true {
                            switch "hello" {
                            case "hello":
                                var v21 uint64 = 0
                                _ = v21
                            case "foo bar":
                                if (((v10 / (((^(v10))) | 1)) == (^(v10))) && true) {
                                    v22 := GenS2{F0: 2147483648}
                                    v23 := GenS2{F0: 1}
                                    _ = v22
                                    _ = v23
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = 2
                                    }
                                    v12 += float32(0.0)
                                }
                            }
                        } else {
                            if len(v2) > 0 {
                                v2[0] = 2147483648
                            }
                        }
                    }
                }
                v5 -= (-(v5))
            }
        } else {
            if (false && false) {
                v9 = (v1 * v1)
                v24 := GenS2{F0: 2147483647}
                _ = v24
            } else {
                v9 += v1
                var v25 int16 = (v1 / ((int16(v5)) | 1))
                var v26 bool = false
                _ = v25
                _ = v26
            }
            var v27 int16 = 281
            fmt.Println(v27, fi15)
            _ = v27
        }
        if ((!(!false)) || ((^(v6)) < (v6 * (v6 % (((v6 * v6)) | 1))))) {
            for fi28 := 0; fi28 < 14; fi28++ {
                _ = len(v2)
                var v29 uint16 = 255
                v2 = append(v2, 42)
                _ = fi28
                _ = v29
            }
        }
        v6 = v6
        _ = fi15
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        if false {
            for fi30 := 0; fi30 < 6; fi30++ {
                var v31 uint64 = 933
                _ = fi30
                _ = v31
            }
            v32 := GenS3{F0: 2147483647}
            _ = v32
        }
    } else {
        v33 := GenS3{F0: 2147483647}
        _ = v33
    }
    v6 = (-(v6))
    if ("a" != ((("a" + ("foo bar" + ((("hello" + ("\n\t" + "hello")) + "hello") + ("\n\t" + "")))) + "") + "")) {
        v2 = append(v2, uint32(v4))
    } else {
        for fi34 := 0; fi34 < 7; fi34++ {
            for fi35 := 0; fi35 < 9; fi35++ {
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v7 = v7
                    if len(v2) > 0 {
                        v2[0] = 836
                    }
                } else {
                    var v36 string = "a"
                    var v37 uint32 = 2147483647
                    v8 = ((v8 / v8) / v8)
                    _ = v36
                    _ = v37
                }
                v6 -= v6
                var v38 uint32 = 926
                var v39 uint = 65535
                _ = fi35
                _ = v38
                _ = v39
            }
            v40 := GenS1{F0: ("a" + "foo bar"), F1: false}
            v41 := GenS4{F0: 482}
            _ = fi34
            _ = v40
            _ = v41
        }
        var v42 int8 = (v6 * 42)
        _ = v42
    }
    var v43 int8 = (v6 * (^(v6)))
    if false {
        if (v6 >= -77) {
            v12 += 0.0
        } else {
            var v44 uint8 = 130
            _ = v44
        }
        var v45 uint16 = v4
        switch "hello" {
        case "hello":
            fmt.Println(v12, v8)
        default:
            switch v3 {
            case 0:
                v6 = -26
            }
        }
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
