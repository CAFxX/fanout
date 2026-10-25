package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 uint; F2 bool }

func main() {
    var v1 uint = 65535
    var v2 string = "hello"
    v3 := []string{"foo bar", v2}
    v4 := []string{"a", (v2 + "foo bar"), v2}
    v1 -= (^(v1))
    switch 1 {
    case 2147483647:
        if (!false) {
            v2 = "hello"
            v2 = ("\n\t" + (((v2 + v2) + "") + v2))
            switch 9223372036854775807 {
            case 32768:
                fmt.Println(v2)
            case 1:
                v5 := GenS1{F0: 100}
                _ = v5
            case -893:
                if (255 == 128) {
                    _ = len(v4)
                    for fi6 := 0; fi6 < 20; fi6++ {
                        var v7 uint = 65535
                        v1 = (v1 << ((v1) & 63))
                        _ = fi6
                        _ = v7
                    }
                }
            }
        } else {
            for fi8 := 0; fi8 < 10; fi8++ {
                if true {
                    v9 := GenS1{F0: 100}
                    fmt.Println(v1, fi8)
                    switch "a" {
                    case "hello":
                        switch 0 {
                        case 390:
                            if (0 < -1) {
                                v10 := GenS2{F0: 2147483647, F1: 1}
                                switch "foo bar" {
                                case "foo bar":
                                    v2 = "a"
                                case "":
                                    v2 = ""
                                case "a":
                                    v1 = 867
                                default:
                                    for fi11 := 0; fi11 < 16; fi11++ {
                                        v12 := GenS3{F0: 32767}
                                        _ = fi11
                                        _ = v12
                                    }
                                }
                                v4 = append(v4, v2)
                                _ = v10
                            }
                        case -1:
                            v2 = v2
                        case 0:
                            var v13 uint64 = 2
                            _ = v13
                        }
                    case "\n\t":
                        _ = len(v4)
                    case "foo bar":
                        if len(v4) > 0 {
                            v4[0] = ((v2 + v2) + "")
                        }
                    }
                    _ = v9
                } else {
                    var v14 uint = (v1 / (((v1 - 714)) | 1))
                    _ = v14
                }
                _ = len(v3)
                switch fi8 {
                case 32768:
                    v15 := GenS4{F0: -497, F1: v1, F2: (((fi8 & (fi8 / ((9223372036854775807) | 1))) == 127) || false)}
                    _ = v15
                default:
                    _ = len(v4)
                }
                _ = fi8
            }
        }
    case -528:
        var v16 float32 = 0.1
        _ = v16
    case 1:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v17 := GenS1{F0: 24}
    v18 := GenS4{F0: 512, F1: v1, F2: (3.14159 == 1e10)}
    if (false && (false || false)) {
        var v19 uint = v1
        _ = v19
    }
    v1 = (v1 & v1)
    var v20 uint8 = 128
    var v21 uint = (v1 << ((v1) & 63))
    if false {
        if len(v3) > 0 {
            v3[0] = ""
        }
        for fi22 := 0; fi22 < 19; fi22++ {
            fmt.Println(fi22, v2)
            _ = fi22
        }
        fmt.Println(v2, v1, v1)
    } else {
        v23 := GenS3{F0: 1}
        v24 := GenS1{F0: -128}
        _ = v23
        _ = v24
    }
    fmt.Println(v21, v1, v20)
    fmt.Println(v21)
    switch 794 {
    case 27:
        v3 = append(v3, (v2 + v2))
    case -459:
        v21 = (v1 << ((v21) & 63))
    case -2147483648:
        for fi25 := 0; fi25 < 5; fi25++ {
            fmt.Println(v21, v1, v21)
            _ = fi25
        }
    }
    var v26 uint16 = 750
    v2 = ("a" + v2)
    v27 := GenS3{F0: v26}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
