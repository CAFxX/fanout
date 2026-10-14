package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 1
    var v3 int64 = 0
    var v4 uint32 = 2147483647
    v5 := []uint32{v4, v4}
    var v6 int64 = (v3 % ((v3) | 1))
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 string = ("" + ("hello" + ("foo bar" + ("a" + "\n\t"))))
        _ = fi7
        _ = v8
    }
    if ("\n\t" < "a") {
        var v9 float64 = 0.1
        _ = v9
    } else {
        if false {
            v2 -= v2
            switch 32768 {
            case 32767:
                var v10 uint = 4294967295
                _ = v10
            case -221:
                var v11 uint16 = 1
                _ = v11
            default:
                for fi12 := 0; fi12 < 12; fi12++ {
                    _ = len(v5)
                    v6 = v6
                    _ = fi12
                }
            }
            var v13 string = "\n\t"
            _ = v13
        } else {
            v2 -= (-(v2))
        }
        v14 := GenS1{F0: 1}
        if len(v5) > 0 {
            v5[0] = v4
        }
        _ = v14
    }
    var v15 int16 = -385
    var v16 int = 127
    var v17 int8 = (v2 | 1)
    var v18 bool = (0.1 == 3.14159)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        v20 := GenS1{F0: (v17 >> ((255) & 63))}
        fmt.Println(fi19)
        _ = fi19
        _ = v20
    }
    var v21 int64 = v3
    var v22 int16 = (v15 ^ v15)
    v1[2147483647] = v16
    if (v4 == (^(v4))) {
        var v23 int = (-(v16))
        var v24 uint8 = 37
        if v18 {
            var v25 uint16 = 65535
            if v18 {
                switch "" {
                case "a":
                    v1[v4] = (-(v16))
                case "hello":
                    var v26 uint32 = 287
                    _ = v26
                case "\n\t":
                    switch (v16 | (v23 - v16)) {
                    case 497:
                        if true {
                            v27 := GenS1{F0: v17}
                            _ = v27
                        }
                    case -2:
                        for fi28 := 0; fi28 < 11; fi28++ {
                            fmt.Println(v18, v22)
                            fmt.Println(v24, v24, v16)
                            for fi29 := 0; fi29 < 4; fi29++ {
                                var v30 uint16 = v25
                                if (v18 && false) {
                                    var v31 float32 = 1.5
                                    switch "\n\t" {
                                    case "\n\t":
                                        v32 := GenS2{F0: v22, F1: 65535}
                                        _ = v32
                                    case "":
                                        fmt.Println(v24, v3)
                                    case "foo bar":
                                        var v33 uint64 = 18446744073709551615
                                        _ = v33
                                    default:
                                        switch (fi29 ^ (^(v16))) {
                                        case 127:
                                            var v34 int = int(v30)
                                            var v35 int8 = -62
                                            _ = v34
                                            _ = v35
                                        case -2147483648:
                                            var v36 uint32 = (v4 >> ((42) & 63))
                                            _ = v36
                                        }
                                        v37 := GenS2{F0: 0, F1: 42}
                                        v15 = -1
                                        _ = v37
                                    }
                                    v23 -= v16
                                    _ = v31
                                }
                                v38 := GenS3{F0: 2147483647}
                                _ = fi29
                                _ = v30
                                _ = v38
                            }
                            _ = fi28
                        }
                    }
                default:
                    v39 := v1[v4]
                    _ = v39
                }
            }
            fmt.Println(v24, v4)
            _ = v25
        }
        _ = v23
        _ = v24
    }
    v15 = (v22 | (v22 >> ((18446744073709551615) & 63)))
    if (v21 >= (v21 / ((0) | 1))) {
        var v40 bool = true
        fmt.Println(v22, v16)
        var v41 uint = 0
        _ = v40
        _ = v41
    } else {
        fmt.Println(v6, v6, v21)
    }
    if v18 {
        v18 = v18
    } else {
        switch ("a" + ("a" + "a")) {
        case "foo bar":
            var v42 bool = v18
            _ = v42
        default:
            if v18 {
                switch (v16 ^ 0) {
                case 9223372036854775807:
                    fmt.Println(v4, v22, v2)
                case -1:
                    _ = len(v1)
                case -396:
                    if (65535 < 0) {
                        fmt.Println(v17, v6, v2)
                    } else {
                        fmt.Println(v4)
                        v21 -= (v21 >> ((442) & 63))
                        v43 := GenS1{F0: v2}
                        _ = v43
                    }
                default:
                    if true {
                        v16 += 32768
                        v2 = v17
                        fmt.Println(v16, v2)
                    }
                }
            }
        }
        fmt.Println(v3)
    }
    fmt.Println(v22, v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
