package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint }
type GenS3 struct { F0 bool; F1 bool }

func main() {
    var v1 int32 = -2147483648
    var v2 uint16 = 32767
    var v3 int = 128
    if true {
        if true {
            v4 := []int{-1, (v3 % ((v3) | 1)), (^(v3))}
            _ = v4
        } else {
            v5 := GenS1{F0: false}
            switch v3 {
            case 32767:
                var v6 int32 = (v1 & v1)
                _ = v6
            case 2147483647:
                var v7 uint32 = 2147483647
                _ = v7
            case -165:
                var v8 float32 = ((1e10 + ((1.5 / 0.1) - 0.0)) - 1e10)
                _ = v8
            }
            v9 := []int{}
            _ = v5
            _ = v9
        }
        v10 := []int{-966, -1, (v3 * (v3 << ((2) & 63)))}
        var v11 bool = true
        _ = v10
        _ = v11
    }
    switch 128 {
    case 32768:
        for fi12 := 0; fi12 < 17; fi12++ {
            if false {
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 uint8 = 127
                    switch (v3 / (((fi13 & (-(fi12)))) | 1)) {
                    case 32768:
                        v2 = v2
                    default:
                        if false {
                            if (true || ((^(fi12)) < (-(v3)))) {
                                v3 = fi12
                                v15 := GenS2{F0: 2147483648, F1: uint64(v14), F2: 596}
                                var v16 bool = false
                                switch "a" {
                                case "hello":
                                    v17 := GenS2{F0: 4294967295, F1: 1, F2: 65535}
                                    _ = v17
                                default:
                                    v16 = false
                                }
                                v2 = 1
                                var v18 int8 = -50
                                _ = v15
                                _ = v16
                                _ = v18
                            }
                        }
                    }
                    if (!false) {
                        if false {
                            v19 := GenS1{F0: false}
                            v14 = (^(v14))
                            if (!false) {
                                var v20 uint32 = 1
                                var v21 uint = 65535
                                if false {
                                    fmt.Println(v21)
                                    v22 := []int{fi13, v3}
                                    _ = v22
                                } else {
                                    v23 := GenS3{F0: false, F1: (32767 > v2)}
                                    v2 += 65535
                                    v24 := []int{}
                                    _ = v23
                                    _ = v24
                                }
                                _ = v20
                                _ = v21
                            }
                            _ = v19
                        }
                        if false {
                            var v25 uint16 = v2
                            _ = v25
                        }
                        fmt.Println(fi13, v14)
                        var v26 uint32 = uint32(v14)
                        v27 := []int{(fi13 ^ (v3 - 42)), (fi12 ^ 4294967295), -780, fi12}
                        if len(v27) > 0 {
                            v27[0] = (fi13 << ((607) & 63))
                        }
                        switch (fi12 << ((256) & 63)) {
                        case -2147483648:
                            if len(v27) >= 2 {
                                v27 = v27[1:len(v27)]
                            }
                        case 2147483647:
                            var v28 int16 = -1
                            _ = v28
                        case -1:
                            v27 = append(v27, fi13)
                        }
                        _ = v26
                        _ = v27
                    }
                    _ = fi13
                    _ = v14
                }
            } else {
                v29 := []int{fi12, -9223372036854775808}
                v1 = int32(v2)
                var v30 uint = 255
                _ = v29
                _ = v30
            }
            _ = fi12
        }
    default:
        var v31 uint = 18446744073709551615
        _ = v31
    }
    v32 := GenS1{F0: (42 != uint8(v2))}
    fmt.Println(v3, v3)
    var v33 string = ("foo bar" + "foo bar")
    v2 = 42
    if true {
        if false {
            v33 = "a"
            var v34 uint64 = 18446744073709551615
            v33 = v33
            _ = v34
        }
        fmt.Println(v1, v33, v2)
    }
    var v35 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v35
}
