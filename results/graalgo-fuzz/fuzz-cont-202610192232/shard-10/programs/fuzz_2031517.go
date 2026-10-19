package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 int32; F2 uint8 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]int)
    v3 := []string{}
    v4 := GenS1{F0: 255}
    delete(v2, "")
    var v5 uint8 = 208
    v1 -= v1
    var v6 int64 = 42
    v1 += -2147483648
    v7 := GenS2{F0: 2}
    var v8 uint32 = 65535
    switch ("\n\t" + "") {
    case "":
        if (true && (!(42 <= 17))) {
            v1 += 156
            var v9 int = int(v8)
            v1 -= (v1 % (((v1 | 2147483647)) | 1))
            _ = v9
        } else {
            for fi10 := 0; fi10 < 14; fi10++ {
                v5 = uint8(v8)
                var v11 int32 = 337
                _ = fi10
                _ = v11
            }
            if false {
                for fi12 := 0; fi12 < 16; fi12++ {
                    if (true || false) {
                        var v13 uint8 = v5
                        for fi14 := 0; fi14 < 12; fi14++ {
                            if false {
                                fmt.Println(v5)
                                v15 := GenS3{F0: (^(v8)), F1: v1, F2: v13}
                                _ = v15
                            } else {
                                v1 += (-(v1))
                                v1 += v1
                                switch "" {
                                case "\n\t":
                                    delete(v2, "\n\t")
                                default:
                                    v8 -= 0
                                }
                            }
                            var v16 uint32 = 2147483647
                            v5 -= 2
                            _ = fi14
                            _ = v16
                        }
                        v5 = 42
                        _ = v13
                    }
                    _ = fi12
                }
                for fi17 := 0; fi17 < 18; fi17++ {
                    v8 += (v8 - (v8 & (v8 ^ v8)))
                    _ = fi17
                }
            }
        }
    case "\n\t":
        v18 := v2["hello"]
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
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
