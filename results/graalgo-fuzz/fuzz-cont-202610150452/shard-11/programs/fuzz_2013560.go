package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int32; F2 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 uint; F2 string }
type GenS5 struct { F0 int64 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]bool)
    v3 := []uint32{617, 311, 814}
    v4 := make(map[string]int)
    v5 := []string{"foo bar", "hello"}
    var v6 uint16 = 65535
    if true {
        var v7 int = 601
        var v8 bool = false
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        if (!false) {
            fmt.Println(v6, v1, fi9)
        }
        fmt.Println(fi9)
        _ = fi9
    }
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS1{F0: -1}
        _ = fi10
        _ = v11
    }
    var v12 int16 = 1000
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS2{F0: 87, F1: 77, F2: 2}
        if true {
            fmt.Println(v1, v6, fi13)
            var v15 uint64 = 18446744073709551615
            switch 32768 {
            case 9223372036854775807:
                switch "\n\t" {
                case "foo bar":
                    var v16 uint16 = 65535
                    _ = v16
                case "\n\t":
                    if true {
                        v15 += (v15 - v15)
                        switch ("" + "hello") {
                        case "hello":
                            if (-73 < -1) {
                                v15 = v15
                            } else {
                                var v17 uint64 = (v15 + (v15 & v15))
                                for fi18 := 0; fi18 < 18; fi18++ {
                                    v15 += (v15 / ((v17) | 1))
                                    fmt.Println(fi13, fi18)
                                    var v19 string = "foo bar"
                                    _ = fi18
                                    _ = v19
                                }
                                v4[""] = 128
                                _ = v17
                            }
                        case "a":
                            for fi20 := 0; fi20 < 19; fi20++ {
                                v21 := GenS1{F0: -118}
                                switch -667 {
                                case -247:
                                    if true {
                                        var v22 uint64 = 830
                                        var v23 string = "\n\t"
                                        _ = v22
                                        _ = v23
                                    }
                                }
                                _ = fi20
                                _ = v21
                            }
                        default:
                            v24 := GenS3{F0: 100}
                            _ = v24
                        }
                    } else {
                        for fi25 := 0; fi25 < 12; fi25++ {
                            var v26 uint16 = 32767
                            v12 += 1
                            _ = fi25
                            _ = v26
                        }
                    }
                }
            default:
                v15 = (v15 >> ((uint(v6)) & 63))
            }
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v6)
    switch 513 {
    case 2147483647:
        for fi27 := 0; fi27 < 20; fi27++ {
            fmt.Println(v12, fi27, v12)
            v6 = (v6 << ((717) & 63))
            var v28 uint8 = (v1 << ((2) & 63))
            _ = fi27
            _ = v28
        }
    }
    var v29 int = 42
    var v30 uint32 = 987
    var v31 bool = (uint32(v6) != (v30 ^ 2147483647))
    v30 = 63
    v32 := GenS4{F0: 2, F1: 18446744073709551615, F2: "a"}
    switch "hello" {
    case "":
        v31 = v31
    case "\n\t":
        if (((1e10 / (3.14159 / ((0.1 + float64(-2.25)) - -2.25))) * 1e10) == float64(0.1)) {
            switch (v29 + -910) {
            case 32768:
                v33 := GenS5{F0: int64(v12)}
                _ = v33
            case -1:
                var v34 int = 0
                _ = v34
            }
            v12 -= (v12 / ((42) | 1))
            var v35 int = -1
            _ = v35
        } else {
            fmt.Println(v1)
            v3 = append(v3, 1)
            var v36 uint8 = v1
            _ = v36
        }
    default:
        for fi37 := 0; fi37 < 7; fi37++ {
            fmt.Println(v6, v30)
            for fi38 := 0; fi38 < 3; fi38++ {
                delete(v4, "foo bar")
                v30 = v30
                v39 := GenS3{F0: 79}
                _ = fi38
                _ = v39
            }
            _ = fi37
        }
    }
    var v40 int16 = (v12 | -298)
    for fi41 := 0; fi41 < 8; fi41++ {
        v42 := GenS1{F0: 736}
        _ = fi41
        _ = v42
    }
    v40 -= 42
    if ((v40 | 0) == (v40 - v12)) {
        if len(v5) > 0 {
            v5[0] = (("" + ("foo bar" + ("a" + ("a" + "a")))) + "a")
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
