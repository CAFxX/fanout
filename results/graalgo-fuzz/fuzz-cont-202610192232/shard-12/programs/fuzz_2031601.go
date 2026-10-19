package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 uint8; F2 bool }
type GenS4 struct { F0 uint8; F1 int32 }

func main() {
    var v1 uint16 = 0
    var v2 uint64 = 18446744073709551615
    v3 := []int{32768, 42}
    var v4 uint32 = 4294967295
    v5 := []string{}
    var v6 uint64 = v2
    for fi7 := 0; fi7 < 12; fi7++ {
        _ = len(v5)
        _ = fi7
    }
    var v8 int8 = 127
    if true {
        v2 = v2
        v8 = v8
    } else {
        if len(v5) > 0 {
            v5[0] = ""
        }
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        fmt.Println(v2, v1)
        if true {
            v10 := GenS1{F0: 184, F1: -1}
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 int16 = -119
                switch (("" + ("hello" + "")) + ("a" + "hello")) {
                case "foo bar":
                    var v13 int = 495
                    _ = v13
                case "hello":
                    if (1 >= 239) {
                        var v14 uint16 = v1
                        _ = v14
                    } else {
                        v15 := GenS2{F0: 2147483647}
                        switch "foo bar" {
                        case "foo bar":
                            v16 := GenS3{F0: 236, F1: 11, F2: (fi9 <= (fi11 * (fi11 / ((fi11) | 1))))}
                            _ = v16
                        }
                        var v17 float32 = 0.1
                        _ = v15
                        _ = v17
                    }
                default:
                    fmt.Println(fi11, fi11)
                }
                v18 := GenS3{F0: -1, F1: 255, F2: (912 > -2147483648)}
                _ = fi11
                _ = v12
                _ = v18
            }
            if (!true) {
                switch int(v1) {
                case -40:
                    v3 = append(v3, fi9)
                case 2147483647:
                    var v19 string = "a"
                    _ = v19
                case 42:
                    v6 = 118
                default:
                    v20 := GenS2{F0: v4}
                    _ = v20
                }
                if (false || true) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    if len(v3) > 0 {
                        v3[0] = 32768
                    }
                    var v21 int64 = 42
                    _ = len(v5)
                    _ = v21
                }
            } else {
                fmt.Println(fi9, v1, v6)
            }
            _ = v10
        } else {
            fmt.Println(v8)
            v4 = (v4 | (v4 ^ (^(v4))))
        }
        switch "hello" {
        case "foo bar":
            var v22 uint64 = 42
            _ = v22
        }
        _ = fi9
    }
    v1 += v1
    v23 := GenS3{F0: -20, F1: 2, F2: ((!(v4 <= 906)) && (!false))}
    if (float32((float64((-2.25 + (3.14159 / 1.5))) / ((3.14159 + (1e10 * -2.25)) / -2.25))) != 0.1) {
        v6 = (^(v6))
    }
    v8 -= v8
    v24 := GenS2{F0: (^(v4))}
    v8 += (v8 >> ((255) & 63))
    for fi25 := 0; fi25 < 4; fi25++ {
        var v26 bool = (0 > 696)
        for fi27 := 0; fi27 < 20; fi27++ {
            switch "hello" {
            case "\n\t":
                v5 = append(v5, "")
            }
            _ = fi27
        }
        _ = fi25
        _ = v26
    }
    if true {
        switch "\n\t" {
        case "foo bar":
            v3 = append(v3, 42)
        case "\n\t":
            for fi28 := 0; fi28 < 10; fi28++ {
                v29 := GenS3{F0: 141, F1: 249, F2: ((32767 > 0) || (148 <= 1))}
                _ = fi28
                _ = v29
            }
        case "":
            var v30 uint32 = (^(v4))
            _ = v30
        default:
            v4 += 2
        }
        if (451 < -1) {
            v31 := GenS4{F0: 17, F1: -1}
            switch "hello" {
            case "":
                switch 1 {
                case -9223372036854775808:
                    switch 386 {
                    case 127:
                        if len(v5) > 0 {
                            v5[0] = ("a" + (("a" + "\n\t") + ""))
                        }
                    }
                case -1:
                    if true {
                        fmt.Println(v6, v2, v2)
                    }
                }
            case "foo bar":
                switch "\n\t" {
                case "":
                    var v32 int8 = v8
                    _ = v32
                default:
                    v6 += (v6 ^ 394)
                }
            case "hello":
                var v33 int = -2147483648
                _ = v33
            }
            _ = v31
        }
    }
    v34 := GenS1{F0: -1, F1: -861}
    v2 += (v2 % ((v2) | 1))
    for fi35 := 0; fi35 < 8; fi35++ {
        v8 = 127
        v2 = 0
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
