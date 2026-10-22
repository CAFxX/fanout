package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint8; F1 bool; F2 uint32 }
func genFn0() uint {
    var v7 string = ""
    v7 = ("hello" + ((v7 + v7) + ("" + ("hello" + "a"))))
    fmt.Println(v7)
    _ = v7
    return 65535
}
func genFn1() int64 {
    var v8 int64 = -784
    v9 := []int{127, 13, 2147483647, -594}
    _ = v8
    _ = v9
    return v8
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -1
    var v3 uint8 = 146
    var v4 uint16 = 453
    var v5 float32 = 0.0
    var v6 int = 1
    var v10 int16 = 830
    var v11 int8 = (v2 << ((genFn0()) & 63))
    v3 -= (v3 % ((82) | 1))
    v6 = v6
    v12 := []int{v6}
    if false {
        fmt.Println(v10, v4)
        var v13 string = "\n\t"
        _ = v13
    }
    if (true || true) {
        var v14 int64 = -4
        _ = v14
    } else {
        v15 := GenS1{F0: 4294967295}
        for fi16 := 0; fi16 < 12; fi16++ {
            if (((v5 / v5) * v5) <= v5) {
                var v17 bool = ((true || true) && ("" <= ""))
                var v18 int32 = 1
                _ = v17
                _ = v18
            } else {
                if true {
                    fmt.Println(v11, v3)
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        if ((-9223372036854775808 != 629) && false) {
                            if len(v12) > 0 {
                                v12[0] = v6
                            }
                            v10 = v10
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                        } else {
                            for fi19 := 0; fi19 < 17; fi19++ {
                                v20 := GenS1{F0: uint32(v3)}
                                v2 += v11
                                _ = fi19
                                _ = v20
                            }
                            fmt.Println(v4, fi16, v3)
                        }
                    case "foo bar":
                        var v21 int8 = 58
                        _ = v21
                    }
                    v6 -= v6
                }
            }
            v1[2] = ("\n\t" + "a")
            _ = fi16
        }
        if ((v6 + 723) != (-(v6))) {
            fmt.Println(v4)
            var v22 int8 = -35
            v23 := v1[1]
            _ = v22
            _ = v23
        }
        _ = v15
    }
    v24 := GenS2{F0: 4294967295, F1: (v10 & (v10 ^ v10))}
    v11 -= (v11 ^ 66)
    v10 = -189
    for fi25 := 0; fi25 < 7; fi25++ {
        v4 += v4
        _ = fi25
    }
    if true {
        v26 := GenS3{F0: v11, F1: 255}
        if false {
            var v27 bool = true
            _ = v27
        }
        _ = v26
    }
    switch v6 {
    case -557:
        if (v5 >= v5) {
            v28 := GenS4{F0: (v3 - (v3 / (((^(v3))) | 1))), F1: true, F2: 1}
            v1[304] = ""
            _ = v28
        }
    }
    var v29 float32 = 0.0
    v30 := GenS2{F0: 783, F1: (v10 >> ((4294967295) & 63))}
    v11 = (v11 + (-(v2)))
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
