package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int16; F1 int32; F2 uint16 }

func main() {
    var v1 uint32 = 2147483648
    var v2 uint8 = 42
    var v3 uint8 = (v2 << ((745) & 63))
    var v4 string = "foo bar"
    var v5 int = -152
    v2 -= v3
    v1 += v1
    v6 := []int{(v5 & (v5 | (v5 % ((-1) | 1)))), v5, v5, int(v1)}
    v5 = v5
    v3 += (v3 >> ((784) & 63))
    switch v5 {
    case 347:
        var v7 uint64 = 651
        _ = v7
    case 4294967295:
        var v8 int32 = 2147483647
        _ = v8
    default:
        if (true || false) {
            if (492 < (v1 % (((v1 % (((v1 / ((v1) | 1))) | 1))) | 1))) {
                v9 := GenS1{F0: 65535, F1: -9223372036854775808}
                if true {
                    v2 -= v3
                    var v10 uint = 271
                    var v11 bool = false
                    _ = v10
                    _ = v11
                }
                _ = v9
            } else {
                v12 := GenS1{F0: 349, F1: 906}
                var v13 int32 = -423
                switch v4 {
                case "hello":
                    v14 := GenS2{F0: 0, F1: (^(v13)), F2: 65535}
                    _ = v14
                case "\n\t":
                    for fi15 := 0; fi15 < 12; fi15++ {
                        v2 += 80
                        switch v5 {
                        case 357:
                            var v16 uint8 = v3
                            _ = v16
                        case -821:
                            if (true && false) {
                                v1 = (^(v1))
                            }
                        }
                        _ = fi15
                    }
                }
                _ = v12
                _ = v13
            }
            for fi17 := 0; fi17 < 20; fi17++ {
                switch fi17 {
                case 42:
                    var v18 int8 = 8
                    _ = v18
                case -9223372036854775808:
                    fmt.Println(v3, v5, v1)
                }
                _ = fi17
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
