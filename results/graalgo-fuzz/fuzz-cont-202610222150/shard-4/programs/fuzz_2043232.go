package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8; F1 int64; F2 int }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    var v3 float64 = 0.1
    v4 := []string{}
    v5 := make(map[uint32]bool)
    var v6 uint8 = 1
    v5[2147483648] = (true && (1.5 > 1.5))
    var v7 uint16 = 1
    var v8 float64 = ((0.1 + v3) * (1.5 / 1e10))
    var v9 int16 = 1
    switch 32767 {
    case 2147483647:
        var v10 uint16 = uint16(v6)
        _ = v10
    }
    var v11 uint32 = 65535
    for fi12 := 0; fi12 < 17; fi12++ {
        _ = len(v4)
        for fi13 := 0; fi13 < 17; fi13++ {
            switch ("foo bar" + ("a" + "a")) {
            case "a":
                var v14 int8 = 127
                _ = v14
            case "foo bar":
                v15 := GenS1{F0: 42}
                _ = v15
            }
            _ = fi13
        }
        _ = fi12
    }
    v7 -= (v7 << ((411) & 63))
    for fi16 := 0; fi16 < 13; fi16++ {
        if true {
            v8 += (v3 / v8)
            if ((v9 * 0) <= v9) {
                var v17 float32 = -2.25
                switch fi16 {
                case 32767:
                    if (true || (v9 > -32768)) {
                        switch fi16 {
                        case 32767:
                            if (!(!false)) {
                                v18 := GenS1{F0: v9}
                                var v19 uint32 = 715
                                var v20 int8 = 127
                                _ = v18
                                _ = v19
                                _ = v20
                            } else {
                                v9 -= (^(v9))
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                    } else {
                        v21 := GenS2{F0: true, F1: 42}
                        v11 += (^(v11))
                        _ = v21
                    }
                case 113:
                    v22 := GenS3{F0: 100}
                    _ = v22
                case 4294967295:
                    v23 := GenS1{F0: v9}
                    _ = v23
                default:
                    var v24 int = fi16
                    _ = v24
                }
                if (!(!("\n\t" == ""))) {
                    v25 := GenS4{F0: (0 > -2147483648), F1: 42}
                    v26 := GenS3{F0: -73}
                    _ = v25
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 17; fi27++ {
                        if false {
                            fmt.Println(v8, fi27, fi16)
                            v28 := GenS4{F0: ((v9 == 42) && (-1 == 231)), F1: -2147483648}
                            _ = v28
                        }
                        _ = fi27
                    }
                }
                _ = v17
            } else {
                v29 := GenS4{F0: false, F1: 42}
                _ = v29
            }
        } else {
            if ("a" >= (("\n\t" + "hello") + ((("\n\t" + "") + "foo bar") + "foo bar"))) {
                v11 = 65535
                v30 := GenS3{F0: -60}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v30
            }
        }
        for fi31 := 0; fi31 < 14; fi31++ {
            v3 = ((float64(float32(0.0)) + float64((1e10 * 0.1))) - float64(1.5))
            if true {
                if true {
                    var v32 uint16 = (v7 + 42)
                    _ = v32
                }
                v3 += (3.14159 * 0.1)
                v9 += 32767
            } else {
                fmt.Println(v7, v7)
                var v33 int16 = v9
                v34 := GenS3{F0: 116}
                _ = v33
                _ = v34
            }
            fmt.Println(v3, v8, fi16)
            _ = len(v4)
            v35 := GenS5{F0: (0 == 258)}
            _ = fi31
            _ = v35
        }
        var v36 uint32 = (v11 & v11)
        _ = fi16
        _ = v36
    }
    v6 = (v6 + (v6 * (^(v6))))
    v37 := GenS6{F0: -128, F1: 1, F2: -447}
    v38 := GenS3{F0: 1}
    if ((v9 >> ((0) & 63)) < v9) {
        var v39 uint = 536
        fmt.Println(v7)
        _ = v39
    } else {
        var v40 bool = (v7 > 65535)
        v41 := GenS5{F0: false}
        switch ("\n\t" + "hello") {
        case "hello":
            fmt.Println(v3, v9, v7)
        case "a":
            var v42 uint32 = 439
            _ = v42
        default:
            _ = len(v5)
        }
        _ = v40
        _ = v41
    }
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
