package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint8; F1 int64; F2 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 uint64; F2 bool }
type GenS5 struct { F0 string; F1 uint16; F2 uint8 }

func main() {
    var v1 int8 = 0
    var v2 int64 = -996
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]int)
    var v5 bool = (!false)
    if v5 {
        v2 = v2
        switch 0 {
        case 2147483647:
            var v6 uint16 = 709
            _ = v6
        case 128:
            if (((v2 != v2) || v5) && ((^(v3)) > v3)) {
                if (!v5) {
                    v4[2147483647] = 9223372036854775807
                    v2 = (v2 << ((255) & 63))
                }
                for fi7 := 0; fi7 < 16; fi7++ {
                    var v8 float64 = -2.25
                    _ = fi7
                    _ = v8
                }
                v9 := GenS1{F0: 2, F1: (!v5)}
                _ = v9
            }
        }
        fmt.Println(v2)
    }
    var v10 int = 0
    if v5 {
        switch 32768 {
        case -225:
            v3 -= 0
        case 9223372036854775807:
            v3 = v3
        default:
            fmt.Println(v10, v3, v1)
        }
    } else {
        fmt.Println(v10)
        if (false || v5) {
            for fi11 := 0; fi11 < 8; fi11++ {
                for fi12 := 0; fi12 < 19; fi12++ {
                    if true {
                        if ((v3 != (^(v3))) || (v1 != (^(v1)))) {
                            if (v5 && v5) {
                                var v13 int16 = 32767
                                _ = v13
                            }
                            v14 := []int{32768, (fi11 ^ fi12), (fi12 - (v10 - (v10 << ((42) & 63))))}
                            var v15 uint32 = 1
                            _ = v14
                            _ = v15
                        }
                    } else {
                        v16 := GenS2{F0: 68, F1: v2, F2: (v1 << ((42) & 63))}
                        var v17 int8 = 1
                        _ = len(v4)
                        _ = v16
                        _ = v17
                    }
                    _ = fi12
                }
                var v18 int32 = 2147483647
                _ = fi11
                _ = v18
            }
            var v19 uint16 = 515
            _ = v19
        } else {
            _ = len(v4)
        }
    }
    switch "\n\t" {
    case "":
        v1 = (-(v1))
    default:
        switch (^(v10)) {
        case 32767:
            for fi20 := 0; fi20 < 17; fi20++ {
                if (32767 > uint16(v3)) {
                    var v21 int = 32767
                    _ = v21
                }
                var v22 uint8 = 128
                if v5 {
                    v23 := []int{}
                    var v24 int64 = int64(fi20)
                    _ = v23
                    _ = v24
                } else {
                    var v25 uint32 = 0
                    _ = v25
                }
                _ = fi20
                _ = v22
            }
        case 1:
            v10 = (v10 / ((v10) | 1))
        case -2147483648:
            for fi26 := 0; fi26 < 5; fi26++ {
                var v27 bool = true
                fmt.Println(v27)
                for fi28 := 0; fi28 < 10; fi28++ {
                    v29 := []int{(fi28 * (fi28 << ((65535) & 63))), fi26, (^(fi26)), v10}
                    var v30 bool = ((4 > 224) || (v5 && v27))
                    var v31 int = -98
                    _ = fi28
                    _ = v29
                    _ = v30
                    _ = v31
                }
                _ = fi26
                _ = v27
            }
        }
    }
    var v32 int16 = -526
    switch ("" + ("\n\t" + "foo bar")) {
    case "":
        switch "foo bar" {
        case "hello":
            for fi33 := 0; fi33 < 18; fi33++ {
                v5 = (0.1 > 0.0)
                _ = fi33
            }
        case "a":
            for fi34 := 0; fi34 < 4; fi34++ {
                v2 += 428
                if v5 {
                    fmt.Println(v32)
                    switch 0 {
                    case 127:
                        delete(v4, 2147483648)
                    case 32768:
                        switch fi34 {
                        case -2147483648:
                            v35 := GenS3{F0: "hello"}
                            _ = v35
                        default:
                            v36 := GenS4{F0: 127, F1: v3, F2: false}
                            _ = v36
                        }
                    case 2147483647:
                        delete(v4, 2147483647)
                    default:
                        v5 = v5
                    }
                    var v37 uint16 = 42
                    _ = v37
                }
                _ = fi34
            }
        }
    case "hello":
        fmt.Println(v5, v2, v32)
    case "\n\t":
        v38 := []int{(v10 / ((v10) | 1)), (-(v10)), (^(v10))}
        _ = v38
    }
    var v39 int8 = (^(v1))
    v40 := GenS2{F0: 255, F1: v2, F2: (v1 ^ v1)}
    if (447 < 42) {
        var v41 uint = 2
        v2 = v2
        if (!v5) {
            v41 += (v41 & v41)
            v5 = false
        }
        _ = v41
    }
    v42 := GenS5{F0: ("a" + ("\n\t" + ("foo bar" + ((("hello" + "a") + "foo bar") + "a")))), F1: 355, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v5
}
