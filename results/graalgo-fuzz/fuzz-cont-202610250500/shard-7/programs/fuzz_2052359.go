package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint32 }
type GenS2 struct { F0 uint64; F1 bool; F2 string }

func main() {
    var v1 bool = true
    var v2 int64 = 219
    v3 := make(map[uint32]uint32)
    var v4 int32 = 42
    if false {
        v5 := []int{32768, 9223372036854775807, 1, 1}
        _ = len(v5)
        v6 := GenS1{F0: 182, F1: 42, F2: 65535}
        _ = v5
        _ = v6
    }
    delete(v3, 880)
    var v7 uint8 = 255
    if v1 {
        v3[14] = 0
        v7 -= 1
    }
    v7 -= v7
    v4 = 2147483647
    var v8 float64 = (float64((3.14159 + 0.0)) + -2.25)
    _ = len(v3)
    v2 = v2
    v2 -= (-(v2))
    v9 := GenS1{F0: 18446744073709551615, F1: 32768, F2: 2147483647}
    if ((((((v8 - float64(1.5)) / v8) - 3.14159) - (0.1 - float64(float32(v8)))) / (float64((1.5 / float32(v8))) / (((0.0 * v8) + 1e10) / v8))) < v8) {
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := GenS1{F0: 462, F1: fi10, F2: 735}
            v12 := GenS1{F0: 947, F1: (-(fi10)), F2: 564}
            _ = fi10
            _ = v11
            _ = v12
        }
    } else {
        fmt.Println(v8, v1)
        if v1 {
            if v1 {
                fmt.Println(v8)
                if ((1 < 42) && (v8 > -2.25)) {
                    v13 := []int{633, -1, -641, 618}
                    v2 -= 741
                    _ = v13
                }
                v7 = v7
            } else {
                if ((v2 * (v2 - (v2 - 9223372036854775807))) > v2) {
                    if false {
                        for fi14 := 0; fi14 < 11; fi14++ {
                            var v15 uint64 = 935
                            var v16 int16 = int16(v7)
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        v7 -= (^(v7))
                    }
                    var v17 uint64 = 259
                    for fi18 := 0; fi18 < 4; fi18++ {
                        switch ("hello" + (((("\n\t" + "") + "foo bar") + "a") + ("a" + "\n\t"))) {
                        case "hello":
                            for fi19 := 0; fi19 < 9; fi19++ {
                                v20 := GenS2{F0: (v17 * 42), F1: (!v1), F2: "\n\t"}
                                v1 = (1e10 < ((0.0 - 0.0) / 1.5))
                                _ = fi19
                                _ = v20
                            }
                        case "foo bar":
                            if v1 {
                                fmt.Println(v4, v8)
                                fmt.Println(v8)
                            }
                        case "\n\t":
                            if ((!v1) && ((fi18 & 9223372036854775807) != fi18)) {
                                v1 = (0.0 >= 1e10)
                                v21 := []int{}
                                _ = v21
                            } else {
                                v22 := []int{497, (fi18 + 638), fi18}
                                v23 := GenS1{F0: 42, F1: fi18, F2: 2147483647}
                                v2 -= (v2 << ((42) & 63))
                                _ = v22
                                _ = v23
                            }
                        }
                        v24 := GenS2{F0: (v17 | 787), F1: (v1 || (1 >= 42)), F2: ""}
                        _ = fi18
                        _ = v24
                    }
                    _ = v17
                } else {
                    var v25 uint64 = 0
                    fmt.Println(v2, v8, v2)
                    var v26 bool = (!v1)
                    _ = v25
                    _ = v26
                }
                if true {
                    for fi27 := 0; fi27 < 15; fi27++ {
                        v7 = (v7 + (v7 << ((1) & 63)))
                        _ = fi27
                    }
                }
            }
        }
    }
    v7 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
