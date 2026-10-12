package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 int32; F2 string }
func genFn0() uint {
    if false {
        var v5 uint32 = 6
        _ = v5
    }
    var v6 float32 = float32((1e10 + 0.1))
    v6 -= 3.14159
    v6 = v6
    v7 := GenS1{F0: "", F1: -50, F2: 2}
    _ = v6
    _ = v7
    return 2
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 1
    v3 := make(map[int]bool)
    var v4 int16 = 93
    v8 := GenS2{F0: 119, F1: 255}
    v9 := GenS2{F0: 42, F1: 2}
    v3[2147483647] = false
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 int = 0
        var v12 uint = 4294967295
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS3{F0: 2, F1: 88}
    v4 += (v4 % (((^(v4))) | 1))
    if ((v4 << ((255) & 63)) <= v4) {
        for fi14 := 0; fi14 < 2; fi14++ {
            fmt.Println(v4, v2)
            for fi15 := 0; fi15 < 18; fi15++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi15
            }
            if ((1.5 * (3.14159 * 3.14159)) <= (0.1 * -2.25)) {
                for fi16 := 0; fi16 < 9; fi16++ {
                    var v17 uint = 2
                    if true {
                        switch fi16 {
                        case 128:
                            fmt.Println(v2, fi14, v4)
                        }
                        var v18 bool = true
                        v19 := GenS4{F0: 42}
                        v17 -= 685
                        _ = v18
                        _ = v19
                    } else {
                        if (-231 >= 1) {
                            if (-2147483648 == 922) {
                                v4 -= v4
                                v17 += (v17 ^ v17)
                            } else {
                                v20 := v3[32768]
                                _ = v20
                            }
                            switch (("" + ("\n\t" + "foo bar")) + "hello") {
                            case "":
                                v17 += 1
                            case "a":
                                v2 = 42
                            case "hello":
                                var v21 int16 = (v4 >> ((v17) & 63))
                                _ = v21
                            }
                        } else {
                            if false {
                                fmt.Println(v2, v2)
                            }
                        }
                    }
                    _ = fi16
                    _ = v17
                }
                v22 := GenS5{F0: "hello", F1: 906, F2: ("hello" + "a")}
                if true {
                    var v23 int32 = 2147483647
                    v2 -= (v2 >> ((293) & 63))
                    _ = v23
                } else {
                    fmt.Println(fi14, v2)
                    fmt.Println(fi14, fi14)
                    fmt.Println(fi14, fi14, fi14)
                }
                _ = v22
            } else {
                v24 := GenS3{F0: 2, F1: 42}
                v4 = int16(v2)
                _ = v24
            }
            _ = fi14
        }
        fmt.Println(v2)
    }
    delete(v3, -625)
    v4 -= (v4 % ((1000) | 1))
    var v25 uint64 = 728
    var v26 int64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
