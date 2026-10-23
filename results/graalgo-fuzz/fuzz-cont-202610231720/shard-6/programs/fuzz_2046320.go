package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int64; F2 uint }

func main() {
    var v1 string = ("\n\t" + "foo bar")
    v2 := make(map[uint32]bool)
    var v3 uint8 = 255
    v4 := []int{-69, -2147483648, -9223372036854775808}
    var v5 int8 = -90
    v6 := make(map[string]bool)
    v3 = (v3 ^ (v3 >> ((0) & 63)))
    v5 = 100
    switch (((v1 + v1) + (v1 + (v1 + (v1 + v1)))) + "a") {
    case "foo bar":
        if true {
            var v7 int8 = (v5 - (v5 << ((42) & 63)))
            v8 := GenS1{F0: 18446744073709551615}
            if false {
                switch v1 {
                case "foo bar":
                    fmt.Println(v5, v7, v5)
                }
                if false {
                    v1 = (v1 + ((("\n\t" + v1) + "") + ((v1 + v1) + "a")))
                    v5 -= -128
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = v7
            _ = v8
        }
    }
    var v9 int = -410
    v1 = v1
    v10 := GenS2{F0: v5, F1: 1, F2: 42}
    v9 = 2147483647
    var v11 bool = (0 > 703)
    fmt.Println(v1, v11)
    v11 = false
    switch (v9 >> ((269) & 63)) {
    case -1:
        switch v9 {
        case 2147483647:
            v11 = v11
        case -2147483648:
            var v12 uint16 = 32767
            _ = v12
        default:
            if v11 {
                for fi13 := 0; fi13 < 20; fi13++ {
                    v1 = v1
                    for fi14 := 0; fi14 < 4; fi14++ {
                        v3 = 1
                        v3 = (v3 ^ v3)
                        v3 += (^(v3))
                        _ = fi14
                    }
                    v9 = fi13
                    _ = fi13
                }
                delete(v6, (v1 + v1))
                fmt.Println(v9)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
