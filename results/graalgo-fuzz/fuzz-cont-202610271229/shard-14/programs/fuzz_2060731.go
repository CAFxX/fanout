package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint8 = 127
    var v2 string = "hello"
    v3 := make(map[uint32]bool)
    var v4 int64 = 0
    var v5 uint64 = 1
    var v6 string = v2
    v4 = 766
    var v7 float64 = float64((1e10 - -2.25))
    v4 -= (v4 * (v4 >> ((18446744073709551615) & 63)))
    v6 = (v6 + ("\n\t" + "foo bar"))
    var v8 uint64 = 693
    v9 := []int{}
    var v10 int = 1
    v11 := GenS1{F0: 1}
    var v12 uint64 = v5
    if true {
        delete(v3, 2147483648)
    } else {
        _ = len(v9)
        fmt.Println(v10, v7, v6)
        for fi13 := 0; fi13 < 16; fi13++ {
            for fi14 := 0; fi14 < 18; fi14++ {
                v8 -= v8
                v15 := GenS1{F0: 0}
                fmt.Println(v8)
                _ = fi14
                _ = v15
            }
            v3[2147483648] = (1.5 >= (1e10 - ((0.0 / -2.25) - 1.5)))
            _ = fi13
        }
    }
    switch v10 {
    case 32768:
        if ((v4 >> ((0) & 63)) >= (v4 & v4)) {
            for fi16 := 0; fi16 < 11; fi16++ {
                _ = len(v9)
                fmt.Println(v4)
                _ = fi16
            }
            switch (("\n\t" + (v6 + "foo bar")) + v2) {
            case "a":
                if (79 > 299) {
                    for fi17 := 0; fi17 < 12; fi17++ {
                        switch v10 {
                        case 9223372036854775807:
                            if (!(false && true)) {
                                for fi18 := 0; fi18 < 20; fi18++ {
                                    _ = len(v3)
                                    _ = fi18
                                }
                                v6 = v6
                            }
                        case 0:
                            switch 9223372036854775807 {
                            case 2147483647:
                                var v19 uint = 4294967295
                                _ = v19
                            }
                        case -55:
                            var v20 uint16 = 0
                            _ = v20
                        default:
                            v5 -= (v5 + (v5 % (((^(v8))) | 1)))
                        }
                        _ = fi17
                    }
                    v6 = "a"
                }
            case "hello":
                v3[2] = true
            }
        } else {
            v12 = 1
        }
    case 32767:
        v9 = append(v9, v10)
    case 127:
        v6 = (v6 + v2)
    default:
        var v21 uint64 = v5
        _ = v21
    }
    if (!((v5 + 2) >= (v12 * v5))) {
        if true {
            if true {
                switch "hello" {
                case "\n\t":
                    for fi22 := 0; fi22 < 3; fi22++ {
                        v8 = v12
                        _ = fi22
                    }
                }
                _ = len(v9)
                if ((v1 >> ((120) & 63)) != (v1 ^ 45)) {
                    var v23 int64 = v4
                    _ = v23
                }
            } else {
                v9 = append(v9, (-(v10)))
                v24 := v3[4294967295]
                _ = v24
            }
            v25 := GenS2{F0: 1}
            _ = v25
        } else {
            var v26 uint64 = v12
            var v27 float64 = (3.14159 / (((v7 + v7) * v7) + v7))
            _ = v26
            _ = v27
        }
        v28 := GenS1{F0: 65535}
        _ = v28
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        v1 -= 127
        for fi30 := 0; fi30 < 19; fi30++ {
            v7 = 1.5
            _ = fi30
        }
        _ = fi29
    }
    var v31 int32 = 42
    v32 := GenS3{F0: (v4 % ((v4) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
