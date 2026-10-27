package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }
func genFn0(p0 string, p1 uint32) uint {
    v7 := []int{-811}
    var v8 string = p0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return uint(p1)
}
func genFn1() uint {
    var v9 int = 128
    v9 += v9
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := []uint32{65535, 664, 2147483648, 52}
    v2 := []string{"", "foo bar", "", "hello"}
    v3 := []int{}
    var v4 int32 = -1
    v5 := make(map[int]uint32)
    var v6 string = "a"
    switch 32768 {
    case -435:
        v4 = 566
    }
    var v10 uint8 = 0
    switch -245 {
    case 9223372036854775807:
        v6 = (v6 + v6)
    }
    v10 -= (v10 % (((v10 - v10)) | 1))
    for fi11 := 0; fi11 < 1; fi11++ {
        fmt.Println(v4, v10)
        _ = fi11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := GenS1{F0: v10}
        if false {
            v10 = v10
            for fi14 := 0; fi14 < 4; fi14++ {
                var v15 int16 = 1
                v4 -= int32(fi12)
                fmt.Println(v10, v6)
                _ = fi14
                _ = v15
            }
            delete(v5, 127)
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi12
        _ = v13
    }
    var v16 int32 = v4
    fmt.Println(v4, v10, v6)
    v4 = -72
    var v17 uint32 = 0
    if true {
        v16 -= (v16 >> ((868) & 63))
        var v18 bool = false
        _ = v18
    } else {
        v6 = v6
        for fi19 := 0; fi19 < 6; fi19++ {
            v6 = v6
            if false {
                if ((0.0 + 0.1) > 1.5) {
                    if (true || (!true)) {
                        var v20 int64 = -9223372036854775808
                        v4 = v4
                        v17 -= v17
                        _ = v20
                    }
                    if (!true) {
                        var v21 uint8 = (v10 % ((v10) | 1))
                        var v22 int16 = -266
                        var v23 int8 = -1
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        switch (v6 + v6) {
                        case "hello":
                            fmt.Println(fi19)
                        case "foo bar":
                            if true {
                                var v24 bool = (731 != 18446744073709551615)
                                switch "\n\t" {
                                case "":
                                    _ = len(v5)
                                default:
                                    if (v24 && v24) {
                                        switch fi19 {
                                        case -1:
                                            var v25 uint8 = (^(v10))
                                            _ = v25
                                        }
                                        v17 -= (^(v17))
                                    }
                                }
                                if v24 {
                                    v26 := GenS1{F0: v10}
                                    _ = v26
                                } else {
                                    var v27 bool = v24
                                    _ = v27
                                }
                                _ = v24
                            }
                        case "\n\t":
                            switch v6 {
                            case "":
                                _ = len(v3)
                            case "a":
                                v28 := GenS1{F0: v10}
                                _ = v28
                            case "foo bar":
                                fmt.Println(fi19, v17, v17)
                            default:
                                var v29 uint64 = 1
                                _ = v29
                            }
                        default:
                            _ = len(v5)
                        }
                        if true {
                            var v30 bool = (!(!false))
                            if false {
                                fmt.Println(v6, v10)
                                var v31 uint64 = 18446744073709551615
                                _ = v31
                            } else {
                                fmt.Println(v4)
                                fmt.Println(v10, v17)
                                v17 = (v17 & v17)
                            }
                            _ = v30
                        } else {
                            v17 += (v17 | (^(v17)))
                            v32 := GenS2{F0: ((v6 + v6) + (v6 + "foo bar"))}
                            _ = v32
                        }
                        v33 := GenS2{F0: v6}
                        _ = v33
                    }
                }
                v6 = "a"
            } else {
                fmt.Println(v4)
                for fi34 := 0; fi34 < 3; fi34++ {
                    v35 := GenS3{F0: 112}
                    _ = fi34
                    _ = v35
                }
            }
            _ = fi19
        }
    }
    if false {
        v16 = (v16 & v16)
        var v36 int32 = (v16 ^ (-(v4)))
        var v37 uint32 = 42
        _ = v36
        _ = v37
    } else {
        var v38 uint16 = 314
        var v39 uint = 0
        _ = v38
        _ = v39
    }
    fmt.Println(v10, v10, v16)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
