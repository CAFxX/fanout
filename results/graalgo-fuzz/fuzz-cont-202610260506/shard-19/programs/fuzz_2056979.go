package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint64 }

func main() {
    var v1 uint32 = 990
    var v2 int8 = -24
    var v3 uint16 = 0
    var v4 uint16 = v3
    var v5 uint16 = v3
    var v6 float32 = (-2.25 * 1e10)
    var v7 int = 42
    if (true || (!(!true))) {
        switch "" {
        case "foo bar":
            switch (v7 & v7) {
            case -889:
                v5 = uint16(v1)
            case -9223372036854775808:
                v8 := GenS1{F0: 81}
                _ = v8
            case 298:
                v9 := GenS1{F0: 42}
                _ = v9
            default:
                fmt.Println(v7, v2)
            }
        case "\n\t":
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 int = (-(v7))
                fmt.Println(v1)
                var v12 int32 = 439
                _ = fi10
                _ = v11
                _ = v12
            }
        }
    } else {
        v2 += 96
        v5 = (v3 - v5)
        for fi13 := 0; fi13 < 6; fi13++ {
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 int64 = -764
                v1 -= v1
                v4 = (^(v4))
                _ = fi14
                _ = v15
            }
            v16 := []int{(fi13 / (((v7 % ((4294967295) | 1))) | 1)), v7}
            var v17 float32 = v6
            _ = fi13
            _ = v16
            _ = v17
        }
    }
    var v18 int16 = -728
    for fi19 := 0; fi19 < 3; fi19++ {
        switch 32767 {
        case -126:
            var v20 uint16 = 2
            _ = v20
        case -2147483648:
            var v21 float64 = 0.0
            _ = v21
        default:
            switch ("\n\t" + (("hello" + "foo bar") + ("" + "a"))) {
            case "\n\t":
                switch v7 {
                case 127:
                    if (v2 <= (v2 ^ v2)) {
                        fmt.Println(v18)
                    }
                }
            case "hello":
                var v22 int16 = -32768
                _ = v22
            }
        }
        var v23 float32 = 1e10
        _ = fi19
        _ = v23
    }
    v24 := []int{v7, -984, (-(v7)), v7}
    if true {
        var v25 bool = (!true)
        var v26 int8 = v2
        _ = v25
        _ = v26
    } else {
        switch v7 {
        case 225:
            for fi27 := 0; fi27 < 5; fi27++ {
                v28 := GenS2{F0: (fi27 % (((^(v7))) | 1)), F1: 2}
                v2 = (-(v2))
                v7 = (fi27 ^ fi27)
                _ = fi27
                _ = v28
            }
        case 0:
            switch (v7 * 32767) {
            case 4294967295:
                v18 = (v18 + (v18 ^ (^(v18))))
            default:
                var v29 int16 = v18
                _ = v29
            }
        case 127:
            for fi30 := 0; fi30 < 5; fi30++ {
                if (42 == 18446744073709551615) {
                    for fi31 := 0; fi31 < 7; fi31++ {
                        if (-2147483648 <= -2147483648) {
                            v2 = 100
                        }
                        switch v7 {
                        case -461:
                            if len(v24) > 0 {
                                v24[0] = (fi31 << ((0) & 63))
                            }
                        }
                        _ = fi31
                    }
                }
                var v32 uint32 = 332
                var v33 uint8 = 128
                _ = fi30
                _ = v32
                _ = v33
            }
        }
    }
    var v34 string = "a"
    var v35 int8 = (v2 * (v2 << ((65535) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
