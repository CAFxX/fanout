package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int32; F1 bool; F2 int32 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint16; F1 int32 }

func main() {
    var v1 int = 32767
    v2 := []int{}
    var v3 uint = 4294967295
    var v4 uint16 = 353
    v5 := []string{("hello" + "\n\t")}
    var v6 int32 = -391
    if (false && (true && (("hello" + "\n\t") < "\n\t"))) {
        switch v1 {
        case -1:
            v3 += v3
        case 9223372036854775807:
            if true {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v4 -= 633
            } else {
                v7 := GenS1{F0: v4, F1: (v1 >> ((v3) & 63))}
                _ = v7
            }
        case 1:
            switch 128 {
            case 1:
                var v8 bool = false
                _ = v8
            case 128:
                switch (v1 & (v1 ^ v1)) {
                case 264:
                    v9 := GenS1{F0: v4, F1: (v1 / (((v1 & v1)) | 1))}
                    _ = v9
                case 9223372036854775807:
                    switch 128 {
                    case -904:
                        fmt.Println(v1, v6)
                    }
                }
            case -1:
                var v10 uint = (v3 & (^(v3)))
                _ = v10
            }
        default:
            v3 += (^(v3))
        }
    }
    v3 = 215
    if (false || (v4 <= (^(v4)))) {
        v1 = (v1 / (((v1 % (((v1 / ((-1) | 1))) | 1))) | 1))
    } else {
        v5 = append(v5, ("a" + "\n\t"))
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        if (false && true) {
            if (!(!false)) {
                for fi12 := 0; fi12 < 17; fi12++ {
                    v6 = -1
                    _ = fi12
                }
                v13 := GenS2{F0: v6, F1: ("foo bar" <= ("a" + (("a" + "a") + ""))), F2: v6}
                _ = v13
            } else {
                v4 = (v4 - v4)
            }
            v1 = (-(v1))
            v14 := GenS3{F0: true, F1: 41, F2: 215}
            _ = v14
        } else {
            if false {
                if true {
                    v5 = append(v5, (("foo bar" + ("foo bar" + "foo bar")) + "hello"))
                }
                if len(v2) > 0 {
                    v2[0] = v1
                }
            } else {
                if false {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        if len(v2) > 0 {
                            v2[0] = 9223372036854775807
                        }
                        v6 = 1
                        _ = fi15
                    }
                }
            }
            v16 := GenS3{F0: true, F1: 0, F2: 261}
            _ = v16
        }
        v17 := GenS4{F0: true, F1: 488}
        v1 = (v1 << ((v3) & 63))
        _ = fi11
        _ = v17
    }
    v6 = -461
    v18 := GenS4{F0: (false || true), F1: 985}
    switch v1 {
    case -463:
        for fi19 := 0; fi19 < 5; fi19++ {
            v4 = (v4 * v4)
            _ = fi19
        }
    case 32768:
        var v20 int8 = 0
        _ = v20
    default:
        fmt.Println(v6, v4)
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 float64 = float64(1.5)
        _ = fi21
        _ = v22
    }
    var v23 uint16 = v4
    v23 -= (v4 ^ (v4 >> (((^(v3))) & 63)))
    var v24 int = 128
    v25 := GenS3{F0: (("" + ("\n\t" + "\n\t")) > "foo bar"), F1: 127, F2: 293}
    if (2147483647 >= v6) {
        fmt.Println(v1, v1, v3)
    } else {
        for fi26 := 0; fi26 < 14; fi26++ {
            var v27 float64 = 0.1
            v28 := GenS5{F0: v4, F1: (-(v6))}
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = len(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
