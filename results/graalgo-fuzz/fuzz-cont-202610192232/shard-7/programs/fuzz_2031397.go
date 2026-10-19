package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 string }
type GenS2 struct { F0 uint8; F1 uint64 }
func genFn0(p0 int16) int8 {
    var v5 string = "hello"
    var v6 int = 32767
    _ = p0
    _ = v5
    _ = v6
    return -128
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    var v3 int32 = 1
    var v4 int16 = 1000
    fmt.Println(v3)
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int16 = (^(v4))
        v9 := GenS1{F0: 901, F1: 0, F2: ""}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 uint8 = 42
        var v12 int = (fi10 % ((9223372036854775807) | 1))
        v4 += 42
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615, F2: "hello"}
    v14 := v2[2147483647]
    var v15 int8 = 127
    var v16 uint8 = 0
    v17 := []int{32768}
    v18 := GenS1{F0: 42, F1: 4294967295, F2: ""}
    for fi19 := 0; fi19 < 8; fi19++ {
        if false {
            v3 += v3
            for fi20 := 0; fi20 < 14; fi20++ {
                v21 := GenS2{F0: (^(v16)), F1: 89}
                switch "" {
                case "\n\t":
                    fmt.Println(v16)
                }
                v14 = (1e10 != -2.25)
                if true {
                    var v22 int8 = (v15 % (((v15 + (-(v15)))) | 1))
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
            v4 -= (v4 + v4)
        }
        var v23 int32 = (-(v3))
        v24 := GenS2{F0: (v16 + (v16 % ((v16) | 1))), F1: 771}
        _ = fi19
        _ = v23
        _ = v24
    }
    switch 9223372036854775807 {
    case 32767:
        delete(v1, 759)
    case 9223372036854775807:
        v15 = 100
    case 42:
        var v25 bool = (v14 || v14)
        _ = v25
    default:
        delete(v1, 127)
    }
    var v26 uint16 = 1
    for fi27 := 0; fi27 < 17; fi27++ {
        for fi28 := 0; fi28 < 18; fi28++ {
            v17 = append(v17, (fi27 & fi27))
            fmt.Println(v16, v26, fi28)
            _ = fi28
        }
        for fi29 := 0; fi29 < 1; fi29++ {
            switch "hello" {
            case "foo bar":
                v30 := GenS1{F0: 2, F1: 2, F2: "hello"}
                _ = v30
            case "\n\t":
                v31 := GenS2{F0: (v16 - v16), F1: 0}
                _ = v31
            default:
                v14 = v14
            }
            delete(v1, (fi27 & fi29))
            _ = fi29
        }
        _ = fi27
    }
    v2[2147483648] = (-751 != -667)
    switch -701 {
    case 32767:
        v4 = (v4 % (((^(v4))) | 1))
    case -720:
        var v32 int32 = 0
        _ = v32
    default:
        switch "hello" {
        case "hello":
            for fi33 := 0; fi33 < 16; fi33++ {
                switch ("foo bar" + ("\n\t" + "\n\t")) {
                case "a":
                    v34 := GenS2{F0: v16, F1: 0}
                    _ = v34
                case "foo bar":
                    v15 = (-(v15))
                case "\n\t":
                    fmt.Println(v26, v26, v14)
                default:
                    v3 -= (v3 << ((18446744073709551615) & 63))
                }
                fmt.Println(v14)
                var v35 uint64 = 16
                _ = fi33
                _ = v35
            }
        }
    }
    var v36 int8 = int8(v4)
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    v37 := v2[4294967295]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
