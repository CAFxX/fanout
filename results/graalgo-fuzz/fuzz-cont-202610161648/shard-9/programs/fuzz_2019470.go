package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int16 }
type GenS2 struct { F0 int64; F1 string }

func main() {
    var v1 string = ("" + "foo bar")
    var v2 int = -2147483648
    var v3 int32 = 1
    var v4 int32 = v3
    var v5 uint64 = 0
    v2 += v2
    var v6 bool = (42 > -67)
    switch (v1 + "hello") {
    case "foo bar":
        for fi7 := 0; fi7 < 3; fi7++ {
            v2 = fi7
            _ = fi7
        }
    }
    if false {
        fmt.Println(v6)
        v8 := []int{}
        _ = v8
    }
    v9 := GenS1{F0: v1, F1: -128, F2: -32768}
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS1{F0: (v1 + v1), F1: 127, F2: -611}
        v3 -= -2147483648
        var v12 uint16 = uint16(v5)
        _ = fi10
        _ = v11
        _ = v12
    }
    if (v1 == ("foo bar" + "\n\t")) {
        v13 := []int{v2, (v2 << ((18446744073709551615) & 63)), v2}
        for fi14 := 0; fi14 < 11; fi14++ {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            switch ("hello" + "\n\t") {
            case "a":
                var v15 int32 = v3
                _ = v15
            case "hello":
                switch "foo bar" {
                case "foo bar":
                    if true {
                        switch fi14 {
                        case 32767:
                            var v16 bool = (!(42 == 1))
                            _ = v16
                        case -192:
                            v4 += (v3 ^ int32(fi14))
                        case 280:
                            if len(v13) > 0 {
                                v13[0] = v2
                            }
                        default:
                            var v17 uint = 4294967295
                            _ = v17
                        }
                    } else {
                        for fi18 := 0; fi18 < 1; fi18++ {
                            var v19 uint8 = 0
                            v20 := GenS1{F0: "a", F1: -2, F2: 230}
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    }
                case "hello":
                    fmt.Println(v3)
                case "a":
                    var v21 bool = (42 < -1)
                    _ = v21
                default:
                    var v22 float64 = 1e10
                    _ = v22
                }
            case "":
                var v23 uint32 = 4294967295
                _ = v23
            default:
                fmt.Println(v4)
            }
            _ = fi14
        }
        _ = v13
    } else {
        v6 = true
        v2 = (v2 | 0)
    }
    switch "a" {
    case "":
        fmt.Println(v5)
    case "a":
        switch v2 {
        case 0:
            v24 := GenS1{F0: ("\n\t" + "\n\t"), F1: 1, F2: 1}
            _ = v24
        case 42:
            v2 += (v2 & 32767)
        case 32768:
            v25 := GenS2{F0: -283, F1: v1}
            _ = v25
        default:
            var v26 float64 = 1e10
            _ = v26
        }
    }
    if v6 {
        v6 = true
        if v6 {
            v27 := GenS2{F0: 1, F1: "a"}
            _ = v27
        }
    }
    var v28 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
