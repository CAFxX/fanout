package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint8; F1 int16 }
func genFn0(p0 string, p1 uint8) uint32 {
    for fi5 := 0; fi5 < 6; fi5++ {
        p0 = p0
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 int8 = 6
            p0 = p0
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    fmt.Println(p1, p0)
    v8 := []int{9223372036854775807}
    p1 = p1
    _ = p0
    _ = p1
    _ = v8
    return 0
}
func genFn1() uint16 {
    var v9 uint16 = 42
    v9 = v9
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 uint16 = (^(v9))
        _ = fi10
        _ = v11
    }
    v12 := []int{}
    var v13 int8 = 100
    _ = v12
    _ = v13
    _ = v9
    return 2
}

func main() {
    v1 := []uint32{0, 65535, 250, 65535}
    var v2 uint = 42
    var v3 int8 = -4
    var v4 uint16 = 6
    v2 += v2
    var v14 int16 = 42
    v15 := GenS1{F0: (false || (18446744073709551615 >= 2)), F1: v14}
    v14 += (v14 ^ 62)
    var v16 uint64 = 1
    if false {
        var v17 uint16 = uint16(v16)
        _ = v17
    }
    v1 = append(v1, uint32(v2))
    var v18 string = "foo bar"
    if true {
        switch "\n\t" {
        case "":
            switch v18 {
            case "":
                if false {
                    v3 = v3
                    var v19 float32 = 0.0
                    var v20 int16 = -857
                    _ = v19
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 2; fi21++ {
                        for fi22 := 0; fi22 < 5; fi22++ {
                            v18 = ((v18 + v18) + "a")
                            _ = fi22
                        }
                        if len(v1) > 0 {
                            v1[0] = 42
                        }
                        _ = fi21
                    }
                }
            case "foo bar":
                fmt.Println(v4, v3, v3)
            case "hello":
                if false {
                    if true {
                        var v23 int = 32768
                        _ = v23
                    } else {
                        var v24 uint32 = 65535
                        switch v18 {
                        case "foo bar":
                            for fi25 := 0; fi25 < 3; fi25++ {
                                var v26 int32 = -1
                                var v27 float64 = 1.5
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        case "a":
                            for fi28 := 0; fi28 < 17; fi28++ {
                                var v29 int64 = -142
                                _ = fi28
                                _ = v29
                            }
                        default:
                            v30 := GenS1{F0: true, F1: -804}
                            _ = v30
                        }
                        fmt.Println(v24)
                        _ = v24
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v31 int = 477
                    _ = v31
                } else {
                    switch "hello" {
                    case "hello":
                        var v32 int16 = (v14 % ((int16(v16)) | 1))
                        _ = v32
                    case "\n\t":
                        var v33 bool = false
                        _ = v33
                    case "a":
                        for fi34 := 0; fi34 < 8; fi34++ {
                            if (!true) {
                                v35 := GenS2{F0: v18, F1: (v2 >> (((v2 / (((v2 + 4294967295)) | 1))) & 63))}
                                _ = v35
                            } else {
                                var v36 int8 = v3
                                for fi37 := 0; fi37 < 2; fi37++ {
                                    if (((v4 | (v4 >> ((0) & 63))) < v4) || true) {
                                        switch -2147483648 {
                                        case 0:
                                            var v38 float64 = ((float64(-2.25) / -2.25) / -2.25)
                                            _ = v38
                                        case -2147483648:
                                            var v39 uint64 = 0
                                            _ = v39
                                        }
                                    } else {
                                        var v40 uint16 = 32767
                                        fmt.Println(v2)
                                        _ = v40
                                    }
                                    _ = len(v1)
                                    _ = fi37
                                }
                                v2 += v2
                                _ = v36
                            }
                            _ = fi34
                        }
                    }
                    if (!true) {
                        v14 = (v14 | v14)
                        fmt.Println(v14, v4)
                    } else {
                        var v41 uint = v2
                        _ = v41
                    }
                    v14 = (v14 + v14)
                }
            }
        case "a":
            fmt.Println(v14, v4, v2)
        case "foo bar":
            var v42 int = 441
            _ = v42
        default:
            v1 = append(v1, 42)
        }
        if true {
            fmt.Println(v4)
        }
        v1 = append(v1, 65535)
    }
    if len(v1) > 0 {
        v1[0] = uint32(v2)
    }
    fmt.Println(v18)
    v1 = append(v1, uint32(v2))
    for fi43 := 0; fi43 < 1; fi43++ {
        fmt.Println(v16, v4)
        _ = fi43
    }
    v44 := GenS3{F0: 42, F1: int16(v2)}
    var v45 int64 = 1
    v46 := GenS2{F0: v18, F1: uint(v4)}
    switch "\n\t" {
    case "hello":
        for fi47 := 0; fi47 < 1; fi47++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v4 = genFn1()
            _ = fi47
        }
    default:
        var v48 int8 = (v3 / ((v3) | 1))
        _ = v48
    }
    v1 = append(v1, 736)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
}
