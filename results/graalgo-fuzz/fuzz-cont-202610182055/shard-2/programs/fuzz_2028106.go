package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 bool }

func main() {
    var v1 uint8 = 217
    v2 := []uint32{4294967295, 2147483648, 4294967295, 4294967295}
    var v3 uint = 457
    var v4 uint64 = 18446744073709551615
    var v5 int64 = 0
    var v6 string = ("\n\t" + "\n\t")
    _ = len(v2)
    v3 += v3
    if true {
        v4 = 4294967295
        var v7 int16 = -1
        if true {
            for fi8 := 0; fi8 < 15; fi8++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi8
            }
        } else {
            var v9 int32 = 2147483647
            for fi10 := 0; fi10 < 13; fi10++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch (fi10 * 32768) {
                case 32767:
                    var v11 uint = (^(v3))
                    _ = v11
                case 0:
                    fmt.Println(v7)
                case 900:
                    v9 = (v9 / (((-(v9))) | 1))
                default:
                    v12 := GenS1{F0: (^(v1)), F1: "a"}
                    _ = v12
                }
                switch int(v5) {
                case 286:
                    for fi13 := 0; fi13 < 4; fi13++ {
                        switch v6 {
                        case "foo bar":
                            var v14 uint64 = 422
                            v3 -= v3
                            fmt.Println(v6)
                            _ = v14
                        case "hello":
                            if len(v2) > 0 {
                                v2[0] = 1
                            }
                        case "a":
                            if len(v2) > 0 {
                                v2[0] = 65535
                            }
                        }
                        var v15 uint8 = 92
                        _ = fi13
                        _ = v15
                    }
                case 32767:
                    switch (^(fi10)) {
                    case -1:
                        v16 := GenS2{F0: (v6 + "\n\t")}
                        _ = v16
                    default:
                        var v17 uint16 = 0
                        _ = v17
                    }
                }
                _ = fi10
            }
            _ = v9
        }
        _ = v7
    }
    var v18 int64 = (v5 ^ (v5 % ((v5) | 1)))
    v5 = -437
    v19 := GenS3{F0: (v5 >> (((v3 | 255)) & 63)), F1: (!false)}
    var v20 uint16 = 32767
    v3 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
