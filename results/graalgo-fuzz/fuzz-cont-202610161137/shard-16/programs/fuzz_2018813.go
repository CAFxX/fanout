package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint64; F2 string }

func main() {
    var v1 uint = 65535
    var v2 int8 = 42
    var v3 string = "foo bar"
    var v4 int = -1
    var v5 int = (v4 * (v4 - (v4 * v4)))
    v6 := GenS1{F0: v2, F1: 0}
    v2 = v2
    var v7 float32 = 1.5
    if false {
        v3 = v3
    }
    v8 := GenS1{F0: int8(v5), F1: v2}
    switch v5 {
    case -9223372036854775808:
        var v9 uint64 = 4294967295
        _ = v9
    case 4294967295:
        v10 := GenS1{F0: v2, F1: -1}
        _ = v10
    case 0:
        var v11 uint64 = 4294967295
        _ = v11
    }
    v12 := GenS1{F0: (-(v2)), F1: (v2 / ((v2) | 1))}
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS1{F0: (v2 >> ((0) & 63)), F1: (v2 + v2)}
        v3 = "\n\t"
        for fi15 := 0; fi15 < 7; fi15++ {
            v1 = v1
            var v16 uint = (v1 >> (((^(v1))) & 63))
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v3, v3)
    if (!false) {
        switch v5 {
        case 1:
            if (2147483647 > 0) {
                for fi17 := 0; fi17 < 5; fi17++ {
                    for fi18 := 0; fi18 < 5; fi18++ {
                        if true {
                            v3 = ""
                            if false {
                                switch (v3 + (v3 + v3)) {
                                case "\n\t":
                                    v19 := GenS2{F0: 1, F1: 133, F2: ""}
                                    _ = v19
                                case "":
                                    var v20 uint8 = 255
                                    _ = v20
                                }
                                v21 := []int{(fi17 * -819), (v4 << (((v1 ^ (v1 - v1))) & 63))}
                                _ = v21
                            }
                        } else {
                            var v22 uint32 = 2147483647
                            v2 = v2
                            switch v3 {
                            case "\n\t":
                                var v23 string = "hello"
                                _ = v23
                            }
                            _ = v22
                        }
                        _ = fi18
                    }
                    if ((!false) || (646 >= 2)) {
                        v4 = (^(v5))
                    }
                    fmt.Println(v3, fi17, v2)
                    _ = fi17
                }
                switch (^(v5)) {
                case 267:
                    v3 = "hello"
                case 128:
                    if (259 <= 1) {
                        switch v5 {
                        case -2147483648:
                            fmt.Println(v4, v5)
                        case 9223372036854775807:
                            var v24 float64 = 3.14159
                            _ = v24
                        }
                        v5 += (^(v5))
                        v4 = (-(v4))
                    }
                }
            }
        }
        v2 += (-(v2))
    } else {
        v25 := []int{(v5 >> (((^(v1))) & 63)), (v5 * v5), 651}
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
