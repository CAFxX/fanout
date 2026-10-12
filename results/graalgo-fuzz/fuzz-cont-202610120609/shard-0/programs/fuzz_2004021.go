package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint16 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]string)
    var v3 int16 = -154
    var v4 uint8 = 76
    var v5 uint = 816
    v6 := []int{107, 127, -730}
    v6 = append(v6, -242)
    var v7 float32 = -2.25
    switch ("\n\t" + "a") {
    case "a":
        if (!false) {
            switch -9223372036854775808 {
            case -2147483648:
                fmt.Println(v3, v4)
            case -356:
                _ = len(v6)
            case 618:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            default:
                if false {
                    var v8 uint64 = 18446744073709551615
                    var v9 float64 = 0.1
                    for fi10 := 0; fi10 < 4; fi10++ {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi10
                    }
                    _ = v8
                    _ = v9
                } else {
                    v5 = (v5 / ((255) | 1))
                    v11 := v1["foo bar"]
                    delete(v2, 0)
                    _ = v11
                }
            }
        }
    case "":
        v4 -= (v4 >> ((756) & 63))
    case "foo bar":
        v5 = v5
    default:
        var v12 bool = (1 <= 900)
        _ = v12
    }
    fmt.Println(v4, v4)
    for fi13 := 0; fi13 < 1; fi13++ {
        if (int8(fi13) == int8(v5)) {
            var v14 int16 = v3
            _ = v14
        } else {
            v3 -= 0
            if (127 <= -128) {
                v2[(fi13 & (fi13 | 127))] = "a"
                for fi15 := 0; fi15 < 7; fi15++ {
                    v16 := GenS1{F0: "\n\t", F1: v5}
                    if true {
                        var v17 uint8 = v4
                        var v18 int8 = 0
                        if (!(!((v17 & (v4 * 1)) != v4))) {
                            v19 := GenS2{F0: v5}
                            _ = v19
                        }
                        v7 = (1e10 * ((v7 - (3.14159 / float32(float64((float32(3.14159) - v7))))) / 0.1))
                        _ = v17
                        _ = v18
                    } else {
                        var v20 int8 = 6
                        _ = v20
                    }
                    v4 -= 1
                    _ = fi15
                    _ = v16
                }
                var v21 bool = ((v4 & 1) <= (v4 + v4))
                _ = v21
            } else {
                for fi22 := 0; fi22 < 3; fi22++ {
                    var v23 int = fi13
                    var v24 uint16 = 771
                    if (false || false) {
                        var v25 uint = (v5 >> (((v5 / (((v5 - v5)) | 1))) & 63))
                        if ((false && false) || (true && true)) {
                            var v26 int16 = 42
                            _ = v26
                        } else {
                            fmt.Println(v3)
                        }
                        v3 = v3
                        _ = v25
                    }
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
        }
        switch ("a" + "hello") {
        case "\n\t":
            if (42 <= 9223372036854775807) {
                fmt.Println(fi13)
            }
        case "hello":
            var v27 string = "\n\t"
            _ = v27
        case "foo bar":
            switch 32768 {
            case 0:
                var v28 uint = v5
                _ = v28
            }
        default:
            switch -2147483648 {
            case -394:
                var v29 uint16 = 247
                _ = v29
            }
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi13
    }
    v4 = (v4 << ((65535) & 63))
    var v30 uint = v5
    if true {
        switch "hello" {
        case "":
            switch "a" {
            case "":
                for fi31 := 0; fi31 < 3; fi31++ {
                    for fi32 := 0; fi32 < 18; fi32++ {
                        fmt.Println(v7, v3)
                        var v33 int8 = 73
                        if false {
                            v5 = (v30 ^ (v30 & 4294967295))
                            switch (-(fi32)) {
                            case 0:
                                switch fi32 {
                                case 4294967295:
                                    var v34 int16 = (v3 - -32768)
                                    _ = v34
                                case -9223372036854775808:
                                    var v35 uint16 = 2
                                    fmt.Println(v30, v30, v7)
                                    v7 -= 0.1
                                    _ = v35
                                case -1:
                                    var v36 uint32 = 861
                                    _ = v36
                                }
                                var v37 int = fi32
                                v3 = (v3 / ((42) | 1))
                                _ = v37
                            case 863:
                                v38 := GenS3{F0: -9223372036854775808, F1: 2}
                                _ = v38
                            }
                            var v39 int64 = 42
                            v5 = v30
                            v4 -= v4
                            var v40 uint = 1
                            v2[(fi31 | (fi32 * (-(fi32))))] = "hello"
                            _ = len(v6)
                            _ = v39
                            _ = v40
                        } else {
                            var v41 uint = (v5 + 2)
                            _ = v41
                        }
                        _ = fi32
                        _ = v33
                    }
                    switch ("\n\t" + "\n\t") {
                    case "foo bar":
                        var v42 uint32 = 65535
                        _ = v42
                    case "":
                        for fi43 := 0; fi43 < 15; fi43++ {
                            fmt.Println(fi43, v5, v30)
                            _ = fi43
                        }
                    default:
                        v4 -= 127
                    }
                    v44 := GenS2{F0: (v30 / ((v5) | 1))}
                    _ = fi31
                    _ = v44
                }
            }
        case "hello":
            var v45 int32 = -116
            _ = v45
        }
    } else {
        v3 -= 186
        var v46 bool = ("\n\t" >= "a")
        _ = v46
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
