package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 int; F2 uint64 }
func genFn0(p0 int) bool {
    var v4 uint64 = 0
    var v5 bool = ((-1 >= 9223372036854775807) && (((!(-2147483648 < 2147483647)) || (32767 >= 42)) && (!true)))
    p0 = p0
    _ = p0
    _ = v4
    _ = v5
    return (!v5)
}
func genFn1(p0 string) bool {
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return true
}

func main() {
    var v1 float64 = (3.14159 / 1e10)
    v2 := []string{"foo bar", "\n\t", (("a" + "hello") + "foo bar")}
    var v3 int = 215
    for fi6 := 0; fi6 < 8; fi6++ {
        v3 = (fi6 >> ((65535) & 63))
        v7 := GenS1{F0: 32768, F1: 2}
        _ = fi6
        _ = v7
    }
    var v8 float32 = 0.0
    var v9 uint8 = 27
    var v10 uint16 = 255
    var v11 int8 = -1
    v12 := GenS2{F0: (^(v10)), F1: 346}
    if false {
        fmt.Println(v9, v3, v10)
        v13 := GenS1{F0: v3, F1: v9}
        var v14 uint8 = 253
        _ = v13
        _ = v14
    }
    if (546 >= 0) {
        if (!(!(uint64(v9) <= 2))) {
            if true {
                if false {
                    var v15 uint64 = 142
                    v16 := GenS3{F0: v15, F1: uint16(v9)}
                    _ = v15
                    _ = v16
                } else {
                    var v17 int8 = (v11 >> ((65535) & 63))
                    var v18 float64 = v1
                    _ = v17
                    _ = v18
                }
            }
            var v19 uint64 = 18446744073709551615
            _ = v19
        }
        _ = len(v2)
        var v20 int8 = (^(v11))
        _ = v20
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v21 uint32 = 2147483648
    for fi22 := 0; fi22 < 6; fi22++ {
        fmt.Println(v8, fi22)
        for fi23 := 0; fi23 < 10; fi23++ {
            v24 := GenS4{F0: 0, F1: "\n\t"}
            v21 = uint32(v9)
            if false {
                fmt.Println(v8, v11)
            }
            v8 = v8
            switch ((("a" + "\n\t") + "") + "hello") {
            case "":
                switch fi23 {
                case 0:
                    var v25 int16 = 0
                    _ = v25
                case 22:
                    v10 += v10
                case 127:
                    v26 := GenS3{F0: 0, F1: 865}
                    _ = v26
                }
            default:
                if true {
                    var v27 int16 = 1
                    switch "foo bar" {
                    case "":
                        _ = len(v2)
                        v28 := GenS1{F0: (fi22 << ((0) & 63)), F1: v9}
                        fmt.Println(v10, v27)
                        _ = v28
                    case "hello":
                        fmt.Println(v9, v3, v8)
                    case "a":
                        v29 := GenS5{F0: 42, F1: 16}
                        _ = v29
                    default:
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                    }
                    if true {
                        v30 := GenS6{F0: 42, F1: 2}
                        var v31 uint8 = v9
                        _ = v30
                        _ = v31
                    } else {
                        v1 = v1
                        if false {
                            fmt.Println(v10, v1)
                            _ = len(v2)
                        } else {
                            v32 := GenS7{F0: ("a" + "hello")}
                            v3 -= fi22
                            _ = v32
                        }
                        switch v3 {
                        case 0:
                            if true {
                                v1 += v1
                                v33 := GenS8{F0: (v10 < 175), F1: v3, F2: 702}
                                _ = v33
                            }
                        case 1:
                            if len(v2) > 0 {
                                v2[0] = (("hello" + "") + "")
                            }
                        case 667:
                            switch ("a" + "") {
                            case "hello":
                                v11 -= -1
                            }
                        }
                    }
                    _ = v27
                }
            }
            _ = fi23
            _ = v24
        }
        fmt.Println(v1, v9)
        _ = fi22
    }
    v2 = append(v2, ("\n\t" + (("foo bar" + "hello") + ("a" + ("a" + ("foo bar" + "\n\t"))))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v8
    _ = v9
}
