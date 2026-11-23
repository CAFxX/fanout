package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int16 = 575
    v2 := make(map[uint32]uint32)
    v3 := []string{(("hello" + ("" + "")) + "hello"), ("foo bar" + "\n\t"), ("hello" + "a")}
    var v4 bool = false
    if (984 == 785) {
        v5 := GenS1{F0: 88}
        _ = v5
    } else {
        fmt.Println(v4, v1)
        for fi6 := 0; fi6 < 7; fi6++ {
            v4 = (42 == v1)
            for fi7 := 0; fi7 < 16; fi7++ {
                v4 = true
                fmt.Println(fi6)
                v3 = append(v3, ("a" + "foo bar"))
                _ = fi7
            }
            var v8 int64 = -800
            _ = fi6
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        switch 9223372036854775807 {
        case 4294967295:
            switch "hello" {
            case "hello":
                v10 := GenS1{F0: -128}
                _ = v10
            case "\n\t":
                switch "" {
                case "":
                    switch "" {
                    case "":
                        if v4 {
                            v1 += v1
                            if true {
                                var v11 uint16 = 1
                                _ = v11
                            }
                            fmt.Println(v4, v1)
                        }
                    case "\n\t":
                        switch "foo bar" {
                        case "":
                            for fi12 := 0; fi12 < 17; fi12++ {
                                v1 = (v1 + (v1 & (^(v1))))
                                v4 = ((42 <= -128) || (-622 != -521))
                                _ = fi12
                            }
                        case "hello":
                            for fi13 := 0; fi13 < 6; fi13++ {
                                v14 := GenS1{F0: 42}
                                _ = len(v2)
                                v1 += -32768
                                _ = fi13
                                _ = v14
                            }
                        case "foo bar":
                            if v4 {
                                var v15 float64 = ((1e10 + float64((0.0 / 1e10))) / 1e10)
                                var v16 float32 = (-2.25 * 1e10)
                                _ = v15
                                _ = v16
                            }
                        }
                    case "a":
                        var v17 bool = (42 >= -2147483648)
                        _ = v17
                    default:
                        if true {
                            v4 = (42 <= 1)
                        } else {
                            var v18 uint8 = 74
                            for fi19 := 0; fi19 < 11; fi19++ {
                                var v20 uint = 18446744073709551615
                                _ = fi19
                                _ = v20
                            }
                            var v21 uint32 = 831
                            _ = v18
                            _ = v21
                        }
                    }
                case "foo bar":
                    v4 = (1000 < (v1 / ((v1) | 1)))
                case "hello":
                    delete(v2, 914)
                }
            default:
                var v22 uint8 = 128
                _ = v22
            }
        case 42:
            for fi23 := 0; fi23 < 20; fi23++ {
                if (v1 >= v1) {
                    var v24 int16 = 1
                    var v25 float64 = 1.5
                    _ = v24
                    _ = v25
                }
                v1 -= 1
                _ = fi23
            }
        }
        v26 := GenS1{F0: -96}
        if true {
            var v27 float64 = 0.0
            _ = v27
        }
        _ = fi9
        _ = v26
    }
    v1 += v1
    var v28 int64 = -607
    v29 := GenS1{F0: -1}
    v4 = (2147483647 < 2147483647)
    fmt.Println(v4, v28)
    var v30 int32 = 1
    switch 32768 {
    case -1:
        v30 -= v30
    default:
        v28 -= v28
    }
    if (!true) {
        for fi31 := 0; fi31 < 13; fi31++ {
            fmt.Println(fi31, v28)
            _ = fi31
        }
        switch -275 {
        case 2147483647:
            for fi32 := 0; fi32 < 5; fi32++ {
                var v33 int64 = v28
                _ = fi32
                _ = v33
            }
        case -2147483648:
            var v34 int32 = (^(v30))
            _ = v34
        default:
            fmt.Println(v1)
        }
    }
    var v35 uint8 = 127
    for fi36 := 0; fi36 < 14; fi36++ {
        v28 = 412
        if false {
            v1 = (v1 << ((885) & 63))
            fmt.Println(fi36)
        }
        _ = fi36
    }
    var v37 int16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v37
    _ = v4
}
