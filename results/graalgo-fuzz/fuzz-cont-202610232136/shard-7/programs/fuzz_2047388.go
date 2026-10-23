package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 bool }
type GenS2 struct { F0 int16; F1 int16 }

func main() {
    var v1 string = "hello"
    var v2 float64 = -2.25
    var v3 float32 = 0.1
    var v4 int8 = 117
    var v5 float64 = ((((1.5 - v2) + (1e10 / 3.14159)) - v2) + 1.5)
    v6 := []uint32{}
    v3 = -2.25
    v2 -= 0.0
    if true {
        v7 := GenS1{F0: -2147483648, F1: v4, F2: true}
        fmt.Println(v4, v1, v5)
        _ = len(v6)
        _ = v7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        v2 += ((0.0 + v2) / (v5 + (3.14159 * v5)))
        v5 = v2
        v6 = append(v6, 4294967295)
        _ = fi8
    }
    v5 = (v2 / v5)
    var v9 uint64 = 1
    if true {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    } else {
        switch 150 {
        case 2147483647:
            var v10 int = 537
            _ = v10
        case 1:
            v2 = v2
        case 42:
            fmt.Println(v3, v3)
        }
        var v11 string = v1
        for fi12 := 0; fi12 < 15; fi12++ {
            switch v11 {
            case "\n\t":
                var v13 float64 = 1e10
                _ = v13
            default:
                v11 = ("foo bar" + v11)
            }
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 uint8 = 0
                _ = fi14
                _ = v15
            }
            v2 = v2
            _ = fi12
        }
        _ = v11
    }
    fmt.Println(v4, v3)
    fmt.Println(v9, v4)
    var v16 bool = true
    if len(v6) > 0 {
        v6[0] = 42
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        v16 = false
        _ = fi17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 uint8 = 42
        fmt.Println(v16)
        var v20 uint16 = uint16(v19)
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 float64 = 3.14159
            var v24 string = v1
            v25 := GenS2{F0: 1000, F1: 42}
            v26 := GenS2{F0: 1, F1: 0}
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
            _ = v26
        }
        fmt.Println(v9)
        for fi27 := 0; fi27 < 9; fi27++ {
            if v16 {
                var v28 int32 = -1
                var v29 int32 = v28
                _ = v28
                _ = v29
            } else {
                v2 = v2
            }
            v9 = (v9 | 342)
            _ = fi27
        }
        _ = fi21
    }
    v30 := GenS1{F0: -641, F1: v4, F2: (1 < -636)}
    switch ((("hello" + "a") + v1) + "\n\t") {
    case "hello":
        var v31 float32 = v3
        _ = v31
    case "":
        v32 := GenS2{F0: 614, F1: -32768}
        _ = v32
    default:
        for fi33 := 0; fi33 < 11; fi33++ {
            for fi34 := 0; fi34 < 19; fi34++ {
                fmt.Println(v1, fi34, v3)
                if false {
                    if true {
                        var v35 bool = false
                        _ = v35
                    } else {
                        if true {
                            v2 = v2
                        } else {
                            var v36 int32 = -1
                            v5 -= 0.0
                            _ = v36
                        }
                        if v16 {
                            switch v1 {
                            case "hello":
                                var v37 uint = 4294967295
                                _ = v37
                            default:
                                var v38 uint8 = 85
                                _ = v38
                            }
                            v1 = v1
                            v3 += float32(0.1)
                            v3 -= v3
                            switch 1 {
                            case 127:
                                v1 = ((v1 + v1) + ("hello" + "foo bar"))
                            }
                        } else {
                            if len(v6) > 0 {
                                v6[0] = 14
                            }
                            if v16 {
                                v3 = v3
                            }
                        }
                    }
                } else {
                    var v39 int16 = -1
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = v39
                }
                v9 = 1
                var v40 uint64 = (v9 ^ (v9 % ((510) | 1)))
                v41 := GenS1{F0: 170, F1: (v4 % (((v4 ^ 5)) | 1)), F2: v16}
                var v42 bool = false
                _ = fi34
                _ = v40
                _ = v41
                _ = v42
            }
            _ = fi33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
