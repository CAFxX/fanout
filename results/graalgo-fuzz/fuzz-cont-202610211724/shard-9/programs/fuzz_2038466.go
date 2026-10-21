package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 float32 = 1.5
    var v2 int32 = 42
    var v3 string = "\n\t"
    v4 := []int{-1, 42, 9223372036854775807}
    var v5 int8 = 107
    v3 = "hello"
    _ = len(v4)
    fmt.Println(v2, v1)
    v5 -= 93
    var v6 int64 = -894
    if true {
        v7 := GenS1{F0: ("hello" + ("foo bar" + (v3 + "foo bar")))}
        _ = v7
    } else {
        switch -554 {
        case 0:
            for fi8 := 0; fi8 < 17; fi8++ {
                switch v3 {
                case "hello":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "a":
                    for fi9 := 0; fi9 < 15; fi9++ {
                        fmt.Println(v6, v1)
                        if false {
                            var v10 int32 = -1
                            var v11 bool = false
                            v11 = v11
                            _ = v10
                            _ = v11
                        }
                        _ = fi9
                    }
                }
                switch v3 {
                case "a":
                    v2 = v2
                case "foo bar":
                    fmt.Println(fi8, fi8)
                default:
                    if len(v4) > 0 {
                        v4[0] = (fi8 + (^(fi8)))
                    }
                }
                _ = fi8
            }
        case -9223372036854775808:
            if false {
                var v12 int = 761
                switch (v12 ^ v12) {
                case -893:
                    v12 += 32767
                case 9223372036854775807:
                    v4 = append(v4, 1)
                default:
                    switch "" {
                    case "":
                        v13 := GenS1{F0: v3}
                        _ = v13
                    case "\n\t":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case "hello":
                        fmt.Println(v12, v2, v1)
                    default:
                        for fi14 := 0; fi14 < 14; fi14++ {
                            var v15 bool = true
                            var v16 int = v12
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    }
                }
                v17 := GenS1{F0: "a"}
                _ = v12
                _ = v17
            }
        default:
            v6 = (v6 >> ((255) & 63))
        }
        if true {
            for fi18 := 0; fi18 < 1; fi18++ {
                for fi19 := 0; fi19 < 10; fi19++ {
                    _ = len(v4)
                    _ = fi19
                }
                _ = fi18
            }
        }
        if false {
            fmt.Println(v1)
            var v20 uint8 = 42
            _ = v20
        }
    }
    v6 += (v6 ^ v6)
    v21 := GenS2{F0: v6, F1: 0}
    var v22 uint = 2
    switch v3 {
    case "foo bar":
        switch v3 {
        case "hello":
            var v23 int32 = (v2 / ((v2) | 1))
            _ = v23
        case "a":
            if true {
                v24 := GenS2{F0: -216, F1: 42}
                v4 = append(v4, -1)
                _ = v24
            }
        case "foo bar":
            fmt.Println(v6)
        }
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        var v26 int16 = 1000
        _ = fi25
        _ = v26
    }
    v27 := GenS3{F0: 42}
    v22 -= 42
    var v28 uint32 = 42
    var v29 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
