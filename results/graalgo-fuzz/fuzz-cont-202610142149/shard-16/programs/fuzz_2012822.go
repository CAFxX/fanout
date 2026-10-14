package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int8 = 1
    var v2 string = "a"
    var v3 int = 1
    v4 := make(map[string]bool)
    var v5 uint = 18446744073709551615
    v1 = -1
    v6 := GenS1{F0: 44, F1: 65535}
    fmt.Println(v5, v3)
    v1 -= 108
    var v7 uint32 = 0
    var v8 uint32 = v7
    switch v3 {
    case 804:
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 uint64 = 18446744073709551615
            v8 += (v7 % ((42) | 1))
            _ = fi9
            _ = v10
        }
    case 32767:
        v3 = v3
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        if ((3.14159 - 1e10) <= -2.25) {
            v3 = -9223372036854775808
            var v12 uint32 = v8
            for fi13 := 0; fi13 < 14; fi13++ {
                fmt.Println(v8, v3)
                _ = fi13
            }
            _ = v12
        } else {
            for fi14 := 0; fi14 < 1; fi14++ {
                if (!(false || false)) {
                    v1 = v1
                } else {
                    v8 = v7
                    if (1 == 42) {
                        switch "foo bar" {
                        case "\n\t":
                            v15 := GenS1{F0: 127, F1: 837}
                            _ = v15
                        }
                    }
                    v3 -= (fi14 << ((2) & 63))
                }
                if ((!(uint32(v5) >= (v8 & (^(v7))))) || true) {
                    var v16 uint64 = 581
                    fmt.Println(v3, v3)
                    _ = v16
                } else {
                    var v17 uint8 = 224
                    v18 := GenS2{F0: 1}
                    var v19 string = v2
                    _ = v17
                    _ = v18
                    _ = v19
                }
                switch ("\n\t" + (v2 + v2)) {
                case "\n\t":
                    v8 = v7
                case "hello":
                    var v20 int8 = 1
                    _ = v20
                default:
                    if false {
                        switch (("hello" + v2) + "") {
                        case "":
                            switch v2 {
                            case "":
                                var v21 bool = (-2147483648 > fi14)
                                _ = v21
                            case "a":
                                var v22 string = (("foo bar" + v2) + ((((((v2 + "hello") + "foo bar") + (v2 + v2)) + (v2 + v2)) + v2) + ((v2 + "") + v2)))
                                _ = v22
                            case "hello":
                                var v23 uint = v5
                                _ = v23
                            }
                            v5 = v5
                            var v24 uint16 = 2
                            var v25 int64 = 1
                            var v26 int16 = -1
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    } else {
                        v8 += 306
                        v27 := []int{-1, v3, fi14}
                        _ = v27
                    }
                }
                _ = fi14
            }
        }
        _ = fi11
    }
    v8 -= v8
    if false {
        var v28 uint16 = 32767
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
