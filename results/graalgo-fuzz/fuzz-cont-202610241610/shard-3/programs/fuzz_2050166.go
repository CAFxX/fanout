package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint64 }

func main() {
    var v1 int = 444
    var v2 uint = 65535
    var v3 uint32 = 872
    var v4 int16 = 1
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]bool)
    v2 = 405
    v7 := GenS1{F0: 1}
    var v8 uint8 = 47
    v2 -= 65535
    fmt.Println(v4, v2)
    var v9 int = (^(v1))
    if (-1 > -285) {
        if false {
            for fi10 := 0; fi10 < 20; fi10++ {
                v11 := GenS2{F0: (v4 >> ((v2) & 63)), F1: 555}
                switch "\n\t" {
                case "\n\t":
                    v1 -= (-(v9))
                case "a":
                    switch fi10 {
                    case 366:
                        v1 -= int(v4)
                    }
                case "":
                    if ((((v8 >> (((v2 - (v2 | 0))) & 63)) <= 128) && (!(((-2147483648 == -2147483648) || false) || true))) && ("\n\t" >= ("" + "a"))) {
                        v9 += (^(v1))
                        fmt.Println(v1, v1, v8)
                    }
                default:
                    for fi12 := 0; fi12 < 20; fi12++ {
                        var v13 uint64 = 4294967295
                        v14 := v5[v3]
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                }
                _ = fi10
                _ = v11
            }
            for fi15 := 0; fi15 < 10; fi15++ {
                v16 := GenS1{F0: 32767}
                _ = fi15
                _ = v16
            }
            if (false && (true || true)) {
                var v17 float32 = (((0.1 + 1.5) + float32(1e10)) / (1e10 - -2.25))
                v18 := GenS2{F0: (v4 ^ v4), F1: 0}
                _ = v17
                _ = v18
            } else {
                v19 := []int{(v9 ^ v1), (v9 & (-(v9))), (v1 << ((1) & 63))}
                v2 = 166
                _ = v19
            }
        } else {
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 bool = ((((v2 + 409) == v2) || (false || (float64(0.0) < 0.0))) && (!(v3 < 65535)))
                if v21 {
                    if ((!(!false)) && v21) {
                        var v22 uint8 = (v8 / ((v8) | 1))
                        v6["hello"] = v21
                        v4 += 150
                        _ = v22
                    } else {
                        switch ("" + "hello") {
                        case "":
                            v23 := GenS1{F0: v4}
                            _ = v23
                        case "\n\t":
                            fmt.Println(v4, v1)
                        }
                    }
                    v8 = v8
                }
                _ = fi20
                _ = v21
            }
        }
        v8 = (v8 & 2)
    }
    switch (("foo bar" + "a") + "") {
    case "":
        switch "\n\t" {
        case "foo bar":
            v2 = (v2 << ((795) & 63))
        default:
            var v24 uint8 = v8
            _ = v24
        }
    case "\n\t":
        var v25 string = "\n\t"
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
