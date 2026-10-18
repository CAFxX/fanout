package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }

func main() {
    var v1 int32 = -2147483648
    var v2 int16 = -387
    var v3 string = ("" + "")
    v4 := []uint32{1, 685}
    var v5 uint32 = 4294967295
    v6 := GenS1{F0: (true && false), F1: 100}
    if (!false) {
        v3 = ("\n\t" + v3)
    } else {
        v2 = -1
        var v7 int32 = v1
        _ = v7
    }
    var v8 float32 = 1.5
    v8 = ((v8 + (((v8 * v8) / -2.25) + (1.5 * v8))) * float32(0.1))
    var v9 uint16 = 42
    var v10 int64 = 858
    switch 32767 {
    case -2147483648:
        v9 = (v9 >> ((1) & 63))
    case -984:
        v4 = append(v4, (^(v5)))
    case 229:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        if (!true) {
            _ = len(v4)
            var v11 uint64 = 1
            v12 := GenS1{F0: true, F1: -1}
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        if (580 < 981) {
            fmt.Println(v1, v5)
        } else {
            v14 := GenS1{F0: ((false || (128 > 128)) && false), F1: -1}
            _ = v14
        }
        switch fi13 {
        case -652:
            _ = len(v4)
        case 16:
            var v15 uint32 = v5
            _ = v15
        case 128:
            v4 = append(v4, (v5 << ((554) & 63)))
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = len(v4)
        _ = fi13
    }
    if len(v4) > 0 {
        v4[0] = v5
    }
    if false {
        fmt.Println(v10, v10, v1)
        if (true || (v2 >= v2)) {
            if false {
                if ((1.5 / float64(3.14159)) != 0.1) {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        for fi17 := 0; fi17 < 5; fi17++ {
                            switch "foo bar" {
                            case "":
                                var v18 int8 = 91
                                _ = v18
                            case "hello":
                                v19 := GenS1{F0: (-2.25 != 0.1), F1: 42}
                                _ = len(v4)
                                var v20 uint = 2
                                _ = v19
                                _ = v20
                            }
                            var v21 bool = false
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi17
                            _ = v21
                        }
                        v2 = int16(v5)
                        switch (^(fi16)) {
                        case 32767:
                            var v22 bool = true
                            _ = v22
                        case 888:
                            v9 -= 838
                        case 42:
                            v1 -= (v1 << ((4294967295) & 63))
                        }
                        _ = fi16
                    }
                    var v23 uint16 = (^(v9))
                    _ = v23
                } else {
                    var v24 uint64 = 624
                    _ = v24
                }
                fmt.Println(v5)
                v9 += 0
            } else {
                if len(v4) > 0 {
                    v4[0] = v5
                }
            }
            var v25 uint64 = 18446744073709551615
            _ = v25
        }
    }
    var v26 uint8 = 127
    if false {
        v2 = (v2 >> ((2) & 63))
    } else {
        v27 := GenS1{F0: (2 != 654), F1: -32}
        fmt.Println(v1, v3, v26)
        _ = v27
    }
    v26 -= 255
    fmt.Println(v8)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
