package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int = 2147483647
    var v2 uint32 = 4294967295
    var v3 uint = 4294967295
    var v4 int16 = -124
    var v5 uint64 = 2
    if true {
        v6 := []int{v1, (v1 * v1), 32767}
        var v7 string = "\n\t"
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: (^(v5)), F1: 127}
    var v9 bool = true
    v9 = true
    switch "a" {
    case "":
        switch 0 {
        case 9223372036854775807:
            if (!v9) {
                for fi10 := 0; fi10 < 10; fi10++ {
                    v11 := []int{}
                    if false {
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    } else {
                        _ = len(v11)
                    }
                    fmt.Println(v3, v2, v5)
                    _ = fi10
                    _ = v11
                }
                var v12 uint = uint(v5)
                if v9 {
                    v12 = v12
                } else {
                    v12 += (v12 % ((65535) | 1))
                    var v13 uint16 = 32767
                    _ = v13
                }
                _ = v12
            }
        case 32767:
            if ((-2.25 + -2.25) == (0.1 - (1e10 * 1.5))) {
                var v14 int64 = 9223372036854775807
                _ = v14
            }
        default:
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 float64 = 1.5
                v4 = (v4 & (v4 ^ 32767))
                for fi17 := 0; fi17 < 12; fi17++ {
                    var v18 uint64 = (v5 >> ((v3) & 63))
                    fmt.Println(v4, fi17, fi15)
                    if (v9 && false) {
                        var v19 int64 = 0
                        _ = v19
                    } else {
                        v2 = v2
                        v5 -= (v18 + (v5 * 558))
                    }
                    if v9 {
                        fmt.Println(fi15, v16)
                    } else {
                        fmt.Println(v2, v18, v1)
                        var v20 int = (v1 / ((fi17) | 1))
                        var v21 uint = uint(v5)
                        _ = v20
                        _ = v21
                    }
                    var v22 uint8 = 255
                    _ = fi17
                    _ = v18
                    _ = v22
                }
                _ = fi15
                _ = v16
            }
        }
    }
    v1 += (v1 << ((v3) & 63))
    v9 = ((float64((-2.25 - 3.14159)) * -2.25) != 0.0)
    var v23 string = "a"
    switch ("foo bar" + ("hello" + v23)) {
    case "":
        v9 = (42 < 255)
    }
    var v24 int8 = -49
    if v9 {
        v1 = v1
    } else {
        if ((197 >= 492) && v9) {
            if (v3 == uint(v5)) {
                v4 -= v4
            } else {
                var v25 int8 = (v24 & v24)
                _ = v25
            }
            v26 := GenS2{F0: v2}
            for fi27 := 0; fi27 < 7; fi27++ {
                v28 := []int{}
                var v29 int8 = v24
                var v30 uint8 = 189
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
            _ = v26
        }
        v3 -= 530
        var v31 int8 = v24
        _ = v31
    }
    var v32 int = v1
    v33 := []int{1, (v1 / ((9223372036854775807) | 1)), (v32 >> ((v3) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
