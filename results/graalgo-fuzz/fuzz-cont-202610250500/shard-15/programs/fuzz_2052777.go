package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int16; F2 int8 }

func main() {
    var v1 int8 = -40
    v2 := []uint32{}
    var v3 float32 = 0.0
    var v4 uint32 = 2147483648
    v5 := []uint32{2147483647, v4}
    var v6 int = -749
    if (!(4294967295 < 4294967295)) {
        var v7 bool = false
        if ((!v7) || v7) {
            v1 = 100
            if (((v4 / ((v4) | 1)) > v4) || (42 < 4294967295)) {
                for fi8 := 0; fi8 < 17; fi8++ {
                    var v9 int16 = -1
                    v4 += v4
                    _ = fi8
                    _ = v9
                }
            }
            for fi10 := 0; fi10 < 19; fi10++ {
                switch (^(fi10)) {
                case 128:
                    if v7 {
                        for fi11 := 0; fi11 < 12; fi11++ {
                            v12 := GenS1{F0: (42 == 18446744073709551615)}
                            if v7 {
                                fmt.Println(v7, v6)
                                if (true && (v7 || (0 >= 42))) {
                                    v4 -= 838
                                    var v13 bool = (3.14159 != 1.5)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v13
                                } else {
                                    v3 = (v3 / (float32(1e10) + (v3 + -2.25)))
                                    var v14 int64 = -329
                                    _ = v14
                                }
                            }
                            _ = fi11
                            _ = v12
                        }
                    }
                default:
                    v15 := GenS2{F0: "", F1: 65535}
                    _ = v15
                }
                v16 := GenS3{F0: 127, F1: 32767, F2: v1}
                _ = fi10
                _ = v16
            }
        }
        _ = v7
    } else {
        fmt.Println(v6)
    }
    v17 := GenS2{F0: "a", F1: 0}
    var v18 int16 = -32768
    v1 += v1
    fmt.Println(v18)
    var v19 int16 = 1000
    v19 = (v19 + v19)
    if false {
        _ = len(v2)
    }
    switch (-(v6)) {
    case 32768:
        v20 := GenS2{F0: "foo bar", F1: 42}
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
