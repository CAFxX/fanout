package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int32, p1 uint64) int64 {
    var v7 uint64 = p1
    v8 := GenS1{F0: 127}
    var v9 uint = 442
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1() int16 {
    var v10 int16 = 0
    v11 := GenS1{F0: 42}
    v12 := GenS1{F0: 255}
    if false {
        var v13 uint16 = 223
        _ = v13
    } else {
        v10 -= (-(v10))
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v10 -= -32768
        _ = fi14
    }
    _ = v10
    _ = v11
    _ = v12
    return (v10 % ((v10) | 1))
}

func main() {
    var v1 float64 = float64((0.0 + -2.25))
    var v2 int32 = 42
    var v3 bool = true
    var v4 uint16 = 2
    v5 := make(map[string]string)
    var v6 int64 = 0
    v15 := GenS1{F0: 218}
    fmt.Println(v6, v4, v2)
    var v16 float32 = 3.14159
    if (!(!false)) {
        fmt.Println(v3)
    } else {
        v1 += 1.5
        v16 = v16
        v17 := GenS2{F0: "", F1: (0 != 100)}
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        for fi19 := 0; fi19 < 16; fi19++ {
            if false {
                switch "a" {
                case "":
                    var v20 string = (("" + "\n\t") + ("a" + "hello"))
                    _ = v20
                }
                if v3 {
                    switch "" {
                    case "foo bar":
                        var v21 int = fi18
                        _ = v21
                    case "hello":
                        v22 := GenS1{F0: 128}
                        _ = v22
                    case "":
                        if (v2 == -697) {
                            var v23 int8 = -1
                            var v24 bool = false
                            _ = v23
                            _ = v24
                        }
                    default:
                        if ((v2 | v2) <= v2) {
                            if true {
                                var v25 int64 = -1
                                fmt.Println(v3, v25, v16)
                                _ = v25
                            }
                        } else {
                            switch fi18 {
                            case 0:
                                v16 += float32(0.0)
                            case 632:
                                var v26 int8 = -52
                                _ = v26
                            case 1:
                                v27 := GenS3{F0: 2147483647}
                                _ = v27
                            default:
                                v28 := GenS4{F0: 364, F1: 992}
                                _ = v28
                            }
                            if true {
                                fmt.Println(v4)
                                v4 = (v4 << ((338) & 63))
                                if v3 {
                                    v6 = (v6 >> ((739) & 63))
                                } else {
                                    v29 := GenS3{F0: 2}
                                    if (v3 && false) {
                                        fmt.Println(v16)
                                        var v30 bool = (2147483647 == 2147483647)
                                        var v31 int16 = 137
                                        _ = v30
                                        _ = v31
                                    }
                                    var v32 int32 = v2
                                    _ = v29
                                    _ = v32
                                }
                            } else {
                                v33 := GenS2{F0: "\n\t", F1: false}
                                v3 = true
                                switch "" {
                                case "":
                                    v16 = v16
                                case "\n\t":
                                    if false {
                                        var v34 uint = 4294967295
                                        v16 = v16
                                        v35 := GenS2{F0: "foo bar", F1: v3}
                                        var v36 int64 = 828
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    }
                                case "a":
                                    var v37 uint8 = 103
                                    _ = v37
                                default:
                                    var v38 string = ""
                                    _ = v38
                                }
                                _ = v33
                            }
                            fmt.Println(v4, v2, v2)
                        }
                    }
                } else {
                    if (-444 >= -32768) {
                        var v39 uint16 = (v4 << ((4294967295) & 63))
                        _ = v39
                    }
                }
                var v40 int32 = 1
                _ = v40
            }
            var v41 float32 = v16
            _ = fi19
            _ = v41
        }
        v42 := GenS5{F0: 255}
        _ = fi18
        _ = v42
    }
    var v43 string = "\n\t"
    for fi44 := 0; fi44 < 9; fi44++ {
        v1 = 3.14159
        _ = fi44
    }
    for fi45 := 0; fi45 < 9; fi45++ {
        var v46 string = ""
        _ = fi45
        _ = v46
    }
    var v47 uint16 = 929
    var v48 int16 = -402
    var v49 uint = 0
    v43 = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
