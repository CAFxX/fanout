package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint8 }

func main() {
    var v1 uint16 = 42
    var v2 int8 = -109
    var v3 float64 = 1e10
    var v4 uint64 = 42
    fmt.Println(v4, v1, v4)
    var v5 uint = 849
    var v6 string = ""
    v7 := []int{-9223372036854775808, 9223372036854775807}
    v5 += (v5 - (v5 / ((v5) | 1)))
    if true {
        v2 += (v2 / ((1) | 1))
        if true {
            var v8 string = (v6 + v6)
            var v9 string = (((v8 + v8) + v8) + v6)
            _ = v8
            _ = v9
        }
    } else {
        fmt.Println(v5, v1, v6)
        v4 -= (v4 + 69)
    }
    if false {
        if (v3 <= v3) {
            switch 4294967295 {
            case 754:
                if true {
                    if ((1e10 >= 1.5) && false) {
                        fmt.Println(v6, v2)
                    }
                } else {
                    switch -1 {
                    case -801:
                        fmt.Println(v2, v1)
                    case 334:
                        for fi10 := 0; fi10 < 7; fi10++ {
                            var v11 uint32 = 42
                            if false {
                                fmt.Println(fi10, v5)
                            }
                            _ = fi10
                            _ = v11
                        }
                    }
                }
            }
        }
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        v13 := GenS1{F0: -133, F1: 42, F2: 255}
        _ = fi12
        _ = v13
    }
    var v14 uint = (v5 >> (((v5 + (v5 / ((737) | 1)))) & 63))
    for fi15 := 0; fi15 < 3; fi15++ {
        _ = len(v7)
        for fi16 := 0; fi16 < 2; fi16++ {
            v4 += 2
            fmt.Println(v4, v2, v4)
            if true {
                v3 = v3
                if true {
                    switch fi16 {
                    case -144:
                        var v17 float64 = 3.14159
                        _ = v17
                    }
                }
            } else {
                var v18 float32 = (1e10 + 3.14159)
                fmt.Println(v5, fi15)
                if (-809 >= 445) {
                    fmt.Println(v1, fi16)
                    if true {
                        v4 = v4
                    }
                    fmt.Println(fi16)
                }
                _ = v18
            }
            _ = fi16
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
