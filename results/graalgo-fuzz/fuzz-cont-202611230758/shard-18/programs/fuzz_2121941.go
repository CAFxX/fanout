package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int8 }
type GenS2 struct { F0 string; F1 uint8; F2 uint64 }

func main() {
    var v1 int = 127
    var v2 uint32 = 0
    var v3 float64 = 1.5
    var v4 uint8 = 128
    var v5 int32 = 625
    v2 = (v2 / (((^(v2))) | 1))
    if (!false) {
        var v6 int = v1
        fmt.Println(v1)
        _ = v6
    }
    if (v5 >= -2147483648) {
        switch "\n\t" {
        case "foo bar":
            v1 = v1
        default:
            v7 := GenS1{F0: ("\n\t" + "a"), F1: 42, F2: -91}
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        switch "a" {
        case "a":
            var v9 uint64 = 1
            _ = v9
        }
        _ = fi8
    }
    switch v1 {
    case 32768:
        if (true && (-9223372036854775808 >= v1)) {
            fmt.Println(v4)
            var v10 int = (v1 & v1)
            v11 := []int{(v10 % ((v10) | 1))}
            _ = v10
            _ = v11
        }
    case 0:
        v4 = v4
    case -1:
        fmt.Println(v5, v1)
    }
    switch "hello" {
    case "foo bar":
        v12 := []int{}
        _ = v12
    case "\n\t":
        for fi13 := 0; fi13 < 8; fi13++ {
            switch "" {
            case "a":
                for fi14 := 0; fi14 < 16; fi14++ {
                    fmt.Println(v4)
                    _ = fi14
                }
            case "\n\t":
                var v15 int64 = -1
                _ = v15
            default:
                for fi16 := 0; fi16 < 1; fi16++ {
                    if false {
                        v17 := GenS1{F0: "\n\t", F1: -9223372036854775808, F2: 127}
                        _ = v17
                    }
                    _ = fi16
                }
            }
            var v18 int16 = 1
            switch (fi13 & (-(fi13))) {
            case -2147483648:
                v19 := []int{int(v5), 9223372036854775807}
                _ = v19
            case 624:
                v3 -= 0.1
            case 128:
                if ((115 > 100) && (("a" + ("a" + "\n\t")) == "\n\t")) {
                    if true {
                        v20 := GenS2{F0: ("\n\t" + "a"), F1: v4, F2: 2}
                        for fi21 := 0; fi21 < 1; fi21++ {
                            if (v2 >= (v2 + 42)) {
                                v22 := GenS2{F0: "\n\t", F1: v4, F2: 360}
                                _ = v22
                            }
                            v23 := []int{}
                            _ = fi21
                            _ = v23
                        }
                        _ = v20
                    }
                    var v24 int64 = 9223372036854775807
                    switch "foo bar" {
                    case "hello":
                        switch "" {
                        case "foo bar":
                            v25 := []int{fi13, (v1 / (((v1 / (((v1 % (((^(v1))) | 1))) | 1))) | 1)), (v1 % ((4294967295) | 1)), -1}
                            _ = v25
                        case "\n\t":
                            v26 := []int{}
                            _ = v26
                        }
                    case "\n\t":
                        v27 := GenS2{F0: "hello", F1: v4, F2: 2}
                        _ = v27
                    }
                    _ = v24
                } else {
                    v2 += v2
                }
            default:
                switch fi13 {
                case -888:
                    var v28 float64 = (((v3 * v3) + v3) * v3)
                    _ = v28
                case -748:
                    v5 -= v5
                default:
                    v2 = v2
                }
            }
            _ = fi13
            _ = v18
        }
    case "a":
        v29 := []int{}
        _ = v29
    }
    var v30 uint16 = 1
    v4 += (v4 << ((232) & 63))
    switch "\n\t" {
    case "":
        v31 := GenS2{F0: "foo bar", F1: (v4 & (v4 & v4)), F2: 1}
        _ = v31
    }
    if (!true) {
        if true {
            v1 -= v1
        }
        v30 = (v30 & (v30 + 255))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
