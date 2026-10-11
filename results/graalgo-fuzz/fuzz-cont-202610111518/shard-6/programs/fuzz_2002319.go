package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int16; F1 uint8; F2 uint }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 int8; F1 uint32 }

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 42
    var v3 bool = ((18446744073709551615 == 238) && (((1e10 * 1.5) <= 1.5) && (!((0 >= 581) && true))))
    var v4 uint = 240
    var v5 float32 = (((3.14159 * 0.0) + (1.5 + 3.14159)) - float32(-2.25))
    var v6 int64 = 29
    if true {
        var v7 uint64 = 1
        var v8 float64 = -2.25
        v9 := GenS1{F0: 545, F1: v7, F2: 2147483648}
        _ = v7
        _ = v8
        _ = v9
    }
    fmt.Println(v5)
    for fi10 := 0; fi10 < 3; fi10++ {
        v3 = (2147483647 < -1)
        var v11 int32 = 0
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            v14 := []int{2147483647, int(v6), fi12, fi13}
            v1[("hello" + ("hello" + ""))] = (("\n\t" + ("foo bar" + "foo bar")) + "a")
            v5 = (float32(0.0) / 1.0)
            _ = fi13
            _ = v14
        }
        var v15 uint32 = 990
        _ = fi12
        _ = v15
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        if true {
            if v3 {
                var v17 uint16 = 0
                _ = v17
            }
            v18 := GenS2{F0: 107, F1: v4, F2: 441}
            _ = v18
        } else {
            for fi19 := 0; fi19 < 17; fi19++ {
                v20 := GenS3{F0: 2, F1: "a"}
                var v21 uint = 151
                v22 := []int{(fi19 + 4294967295), fi16, fi19, (fi16 % (((-(fi19))) | 1))}
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            v6 += int64(fi16)
        }
        fmt.Println(v3)
        _ = fi16
    }
    v1[""] = "hello"
    fmt.Println(v6)
    for fi23 := 0; fi23 < 16; fi23++ {
        v24 := []int{}
        _ = fi23
        _ = v24
    }
    v25 := []int{42, 32767}
    switch "foo bar" {
    case "":
        for fi26 := 0; fi26 < 14; fi26++ {
            for fi27 := 0; fi27 < 1; fi27++ {
                switch "foo bar" {
                case "\n\t":
                    var v28 uint16 = 42
                    _ = v28
                case "foo bar":
                    var v29 int64 = int64(v2)
                    _ = v29
                default:
                    switch fi26 {
                    case -2147483648:
                        var v30 int = -25
                        _ = v30
                    case 107:
                        if v3 {
                            v5 += v5
                            fmt.Println(v4, v5)
                        } else {
                            if v3 {
                                var v31 float32 = v5
                                fmt.Println(v4)
                                _ = v31
                            } else {
                                v25 = append(v25, fi27)
                            }
                            v4 -= 287
                        }
                    case 825:
                        v25 = append(v25, (fi27 ^ fi27))
                    }
                }
                v5 -= (1.5 / v5)
                _ = fi27
            }
            v3 = (-46 == 968)
            _ = fi26
        }
    }
    var v32 int16 = v2
    v25 = append(v25, -563)
    var v33 uint = (v4 << ((v4) & 63))
    v34 := GenS4{F0: (v2 | v32), F1: 0, F2: v33}
    var v35 int32 = 42
    var v36 float64 = ((1e10 + 3.14159) - 1.5)
    for fi37 := 0; fi37 < 20; fi37++ {
        var v38 string = ("" + "")
        v39 := GenS5{F0: 42, F1: (v2 | (v32 + v2))}
        _ = fi37
        _ = v38
        _ = v39
    }
    if (!(true && v3)) {
        v36 = (float64(v5) * v36)
        v4 -= v4
        v40 := GenS6{F0: -118, F1: 65535}
        _ = v40
    } else {
        for fi41 := 0; fi41 < 16; fi41++ {
            v42 := GenS2{F0: 128, F1: v33, F2: 591}
            switch -9223372036854775808 {
            case 128:
                for fi43 := 0; fi43 < 11; fi43++ {
                    var v44 int8 = 127
                    _ = len(v1)
                    _ = fi43
                    _ = v44
                }
            case 9223372036854775807:
                for fi45 := 0; fi45 < 20; fi45++ {
                    switch "foo bar" {
                    case "foo bar":
                        var v46 int16 = v32
                        _ = v46
                    case "":
                        v35 = v35
                    default:
                        var v47 int = 128
                        _ = v47
                    }
                    _ = fi45
                }
            }
            _ = fi41
            _ = v42
        }
        v32 += v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
