package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int; F1 uint; F2 uint8 }
func genFn0(p0 string, p1 int16) int64 {
    v4 := []int{127}
    var v5 bool = false
    var v6 uint64 = 4294967295
    v4 = append(v4, -9223372036854775808)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 int32 = 42
    var v2 int64 = -130
    var v3 uint16 = 138
    v2 -= int64(v1)
    v7 := GenS1{F0: 65535}
    var v8 float32 = 0.1
    if (!false) {
        fmt.Println(v2, v2, v8)
    } else {
        v9 := []int{}
        if true {
            if len(v9) > 0 {
                v9[0] = 32768
            }
        }
        if true {
            if ((1 > 32) && true) {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                var v10 uint16 = v3
                for fi11 := 0; fi11 < 12; fi11++ {
                    v12 := GenS2{F0: 2, F1: 2}
                    for fi13 := 0; fi13 < 14; fi13++ {
                        fmt.Println(fi13)
                        v10 = 1
                        _ = fi13
                    }
                    var v14 uint8 = 224
                    _ = fi11
                    _ = v12
                    _ = v14
                }
                _ = v10
            } else {
                var v15 uint = 65535
                _ = v15
            }
            for fi16 := 0; fi16 < 18; fi16++ {
                if false {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
                switch 32768 {
                case -519:
                    for fi17 := 0; fi17 < 20; fi17++ {
                        v8 -= ((((v8 - 1.5) - ((v8 / 1.5) + float32(3.14159))) / 1.0) / v8)
                        _ = fi17
                    }
                case -466:
                    switch (fi16 >> ((255) & 63)) {
                    case 928:
                        var v18 uint64 = 42
                        _ = v18
                    case 32767:
                        switch 128 {
                        case 127:
                            _ = len(v9)
                        case 128:
                            v19 := GenS1{F0: 0}
                            _ = v19
                        }
                    }
                case 646:
                    v1 += (^(v1))
                default:
                    var v20 uint16 = v3
                    _ = v20
                }
                _ = fi16
            }
            var v21 uint = 42
            _ = v21
        }
        _ = v9
    }
    switch "\n\t" {
    case "\n\t":
        for fi22 := 0; fi22 < 16; fi22++ {
            fmt.Println(v8, fi22, v1)
            _ = fi22
        }
    default:
        if (296 == 728) {
            v23 := GenS1{F0: 65535}
            switch ("foo bar" + "") {
            case "foo bar":
                var v24 string = "a"
                _ = v24
            case "a":
                v25 := []int{127, -2147483648, 0, 127}
                _ = v25
            case "hello":
                for fi26 := 0; fi26 < 10; fi26++ {
                    fmt.Println(fi26, v2)
                    for fi27 := 0; fi27 < 15; fi27++ {
                        v2 += -731
                        _ = fi27
                    }
                    _ = fi26
                }
            }
            _ = v23
        } else {
            v1 = 2147483647
        }
    }
    v28 := GenS3{F0: 9223372036854775807, F1: 4294967295, F2: 128}
    var v29 int8 = 1
    var v30 int64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v7
    _ = v8
}
