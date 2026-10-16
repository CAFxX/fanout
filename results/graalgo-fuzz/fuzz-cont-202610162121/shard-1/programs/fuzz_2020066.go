package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int64 }

func main() {
    var v1 uint64 = 734
    var v2 int = 1
    v3 := []uint32{}
    var v4 uint = 2
    var v5 uint16 = 1
    var v6 bool = true
    for fi7 := 0; fi7 < 2; fi7++ {
        if false {
            fmt.Println(v5, v6, v6)
        }
        v6 = ((42 >= -47) && (2147483647 >= 65535))
        _ = fi7
    }
    switch "" {
    case "\n\t":
        v8 := GenS1{F0: "\n\t", F1: v4, F2: 732}
        _ = v8
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "foo bar":
        for fi9 := 0; fi9 < 13; fi9++ {
            v6 = v6
            _ = fi9
        }
    default:
        var v10 int16 = -222
        _ = v10
    }
    fmt.Println(v4, v2, v5)
    v5 += (^(v5))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v11 := GenS1{F0: "hello", F1: 306, F2: (v1 << ((v4) & 63))}
    var v12 float32 = (3.14159 - 0.1)
    var v13 float32 = float32((((float64(float32(3.14159)) / (-2.25 + float64(v12))) + (1e10 / 1.5)) + (-2.25 - 0.1)))
    _ = len(v3)
    v5 = 0
    v3 = append(v3, 2)
    if false {
        switch v2 {
        case 146:
            v3 = append(v3, 465)
        }
        v13 = 0.0
        var v14 uint8 = 42
        _ = v14
    } else {
        fmt.Println(v2, v5)
        var v15 int32 = 0
        if len(v3) > 0 {
            v3[0] = 42
        }
        _ = v15
    }
    switch "foo bar" {
    case "foo bar":
        v2 += (v2 << ((v4) & 63))
    case "a":
        if true {
            switch v2 {
            case 32767:
                switch (v2 ^ -49) {
                case 418:
                    switch ("hello" + (("a" + "hello") + "a")) {
                    case "\n\t":
                        var v16 bool = v6
                        _ = v16
                    }
                case 456:
                    switch (v2 ^ (^(v2))) {
                    case -52:
                        v3 = append(v3, 135)
                    }
                case 32767:
                    if (100 <= 99) {
                        var v17 uint = v4
                        if ((v6 && v6) || false) {
                            v18 := GenS2{F0: 939, F1: 9223372036854775807}
                            _ = v18
                        }
                        _ = v17
                    } else {
                        for fi19 := 0; fi19 < 18; fi19++ {
                            var v20 int32 = 42
                            _ = fi19
                            _ = v20
                        }
                        v21 := GenS2{F0: 823, F1: -375}
                        var v22 uint8 = 1
                        _ = v21
                        _ = v22
                    }
                default:
                    switch "hello" {
                    case "a":
                        v1 += v1
                    case "hello":
                        var v23 float64 = 0.0
                        _ = v23
                    default:
                        var v24 bool = ((v12 * v13) < ((-2.25 + v12) * 3.14159))
                        _ = v24
                    }
                }
            default:
                fmt.Println(v2, v4)
            }
            fmt.Println(v6)
        } else {
            var v25 float64 = 0.1
            var v26 float64 = 3.14159
            for fi27 := 0; fi27 < 5; fi27++ {
                for fi28 := 0; fi28 < 7; fi28++ {
                    switch fi27 {
                    case 257:
                        fmt.Println(v12, fi28, v12)
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    fmt.Println(v26)
                    v12 = v13
                    if true {
                        switch (fi28 & fi28) {
                        case -2147483648:
                            var v29 int64 = 144
                            if (!false) {
                                if (42 > 93) {
                                    switch "hello" {
                                    case "hello":
                                        v2 -= (fi27 * int(v4))
                                    case "foo bar":
                                        v4 += 18446744073709551615
                                    case "\n\t":
                                        var v30 int16 = 1000
                                        _ = v30
                                    }
                                    v31 := GenS1{F0: ("\n\t" + ("\n\t" + (("foo bar" + "a") + ("a" + ("" + ""))))), F1: (v4 + (v4 * (v4 % ((v4) | 1)))), F2: 0}
                                    v32 := GenS1{F0: "hello", F1: v4, F2: (^(v1))}
                                    _ = v31
                                    _ = v32
                                }
                            }
                            _ = v29
                        }
                    }
                    switch "a" {
                    case "foo bar":
                        var v33 uint = v4
                        _ = v33
                    case "a":
                        if false {
                            fmt.Println(v4, v26, v13)
                        }
                    }
                    _ = fi28
                }
                if len(v3) > 0 {
                    v3[0] = 718
                }
                v6 = true
                _ = fi27
            }
            _ = v25
            _ = v26
        }
    case "hello":
        for fi34 := 0; fi34 < 1; fi34++ {
            var v35 int8 = -1
            _ = fi34
            _ = v35
        }
    }
    v3 = append(v3, 42)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
