package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 uint; F1 bool }

func main() {
    v1 := make(map[string]int)
    var v2 int = 128
    var v3 uint64 = 995
    v4 := make(map[int]uint32)
    v2 = v2
    fmt.Println(v3)
    for fi5 := 0; fi5 < 17; fi5++ {
        switch "a" {
        case "a":
            if false {
                for fi6 := 0; fi6 < 12; fi6++ {
                    if true {
                        v7 := GenS1{F0: 32767, F1: -518}
                        _ = v7
                    } else {
                        if false {
                            v8 := GenS1{F0: int16(fi5), F1: -9223372036854775808}
                            _ = v8
                        } else {
                            var v9 int64 = -1
                            var v10 int8 = -128
                            v1["hello"] = (fi6 / (((v2 / (((^(fi5))) | 1))) | 1))
                            v10 = 100
                            _ = v10
                            _ = v9
                        }
                    }
                    switch "foo bar" {
                    case "\n\t":
                        var v11 uint16 = 65535
                        _ = v11
                    }
                    switch "a" {
                    case "hello":
                        var v12 int64 = 9223372036854775807
                        _ = v12
                    }
                    _ = fi6
                }
            }
        case "\n\t":
            _ = len(v4)
        case "":
            switch "foo bar" {
            case "foo bar":
                if true {
                    v13 := GenS1{F0: -32768, F1: -1}
                    if (-350 != -698) {
                        fmt.Println(v3)
                        if ((0 < 165) && false) {
                            v14 := []int{v2, (-(fi5)), v2, v2}
                            _ = v14
                        }
                    }
                    switch 32768 {
                    case 4294967295:
                        var v15 uint64 = v3
                        _ = v15
                    case -2147483648:
                        v3 -= v3
                    }
                    _ = v13
                }
            case "a":
                v16 := v1["\n\t"]
                _ = v16
            case "\n\t":
                v2 = (v2 & v2)
            default:
                switch -9223372036854775808 {
                case 4294967295:
                    if (65 > 128) {
                        fmt.Println(v3, v2)
                        var v17 int64 = 926
                        fmt.Println(fi5, v2, v3)
                        _ = v17
                    }
                }
            }
        default:
            _ = len(v1)
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            fmt.Println(fi5, fi18)
            var v19 int16 = 32767
            _ = fi18
            _ = v19
        }
        _ = fi5
    }
    if (!((v3 * 4294967295) < (^(v3)))) {
        switch v2 {
        case 42:
            if (1.5 < float64(float32((float64(1.5) / 0.1)))) {
                for fi20 := 0; fi20 < 13; fi20++ {
                    v21 := GenS1{F0: 0, F1: 42}
                    _ = fi20
                    _ = v21
                }
                fmt.Println(v3, v3)
                v22 := GenS1{F0: -32768, F1: 367}
                _ = v22
            }
        default:
            v23 := GenS2{F0: "hello"}
            _ = v23
        }
        v24 := []int{v2}
        _ = v24
    } else {
        if true {
            if ((-2147483648 == 1) && true) {
                v3 += v3
                for fi25 := 0; fi25 < 1; fi25++ {
                    if (978 != 42) {
                        v26 := GenS2{F0: "hello"}
                        var v27 int8 = 0
                        v27 -= (v27 << ((65535) & 63))
                        _ = v26
                        _ = v27
                    } else {
                        v28 := GenS3{F0: false, F1: 0}
                        _ = v28
                    }
                    _ = fi25
                }
                var v29 float64 = 3.14159
                _ = v29
            }
        } else {
            switch 1 {
            case 463:
                switch "\n\t" {
                case "":
                    v2 -= (v2 ^ (v2 ^ int(v3)))
                case "\n\t":
                    var v30 int8 = 74
                    _ = v30
                default:
                    if false {
                        if ((uint8(v3) >= 1) || (!false)) {
                            v31 := GenS4{F0: 0, F1: (9223372036854775807 >= 47)}
                            var v32 uint8 = 255
                            v2 -= 32767
                            _ = v31
                            _ = v32
                        } else {
                            var v33 float32 = ((float32(-2.25) - 0.0) - 1.5)
                            if (377 <= 309) {
                                if false {
                                    v2 = v2
                                    v34 := GenS3{F0: false, F1: 0}
                                    _ = v34
                                }
                                v2 = v2
                                fmt.Println(v3, v3, v2)
                            }
                            v2 = 168
                            _ = v33
                        }
                        var v35 int16 = -1
                        _ = v35
                    } else {
                        v36 := []int{v2, (-(v2))}
                        _ = v36
                    }
                }
            }
            var v37 uint64 = (v3 >> ((0) & 63))
            _ = v37
        }
        var v38 float64 = 0.0
        var v39 bool = (42 >= (v3 * v3))
        _ = v38
        _ = v39
    }
    v40 := []int{(v2 << ((42) & 63)), v2, (v2 & 1)}
    var v41 float32 = float32(-2.25)
    var v42 uint32 = 139
    for fi43 := 0; fi43 < 7; fi43++ {
        v3 -= v3
        var v44 uint32 = 65535
        fmt.Println(v42, v3, v42)
        _ = fi43
        _ = v44
    }
    var v45 uint8 = 42
    _ = len(v1)
    delete(v1, "a")
    v46 := GenS4{F0: uint(v45), F1: true}
    for fi47 := 0; fi47 < 20; fi47++ {
        fmt.Println(v2, v42, v2)
        _ = fi47
    }
    if (v42 > (v42 * 4294967295)) {
        v42 -= (v42 * v42)
        delete(v1, "")
        _ = len(v1)
    } else {
        v48 := v4[v2]
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v45
    _ = v46
}
