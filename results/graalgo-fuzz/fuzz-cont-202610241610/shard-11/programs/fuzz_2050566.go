package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS2 struct { F0 string; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint64; F1 bool; F2 int32 }
type GenS4 struct { F0 uint32; F1 bool }

func main() {
    var v1 int = -3
    var v2 int32 = -2147483648
    var v3 string = "foo bar"
    var v4 uint8 = 255
    v1 = (v1 ^ v1)
    v5 := GenS1{F0: 637, F1: v4, F2: 9223372036854775807}
    var v6 string = ((("" + v3) + v3) + v3)
    if true {
        fmt.Println(v1)
        v7 := []int{1, (-(v1)), int(v4), 128}
        _ = v7
    }
    v6 = (v3 + "a")
    if false {
        v8 := []int{(-(v1)), (-(v1))}
        _ = v8
    } else {
        v3 = ((v3 + v6) + "\n\t")
        fmt.Println(v3, v1)
    }
    v9 := GenS1{F0: 1, F1: (v4 * 66), F2: 616}
    for fi10 := 0; fi10 < 2; fi10++ {
        v3 = v6
        v11 := []int{4294967295, v1}
        if (!true) {
            switch "\n\t" {
            case "\n\t":
                var v12 int64 = -9223372036854775808
                _ = v12
            }
            for fi13 := 0; fi13 < 15; fi13++ {
                _ = len(v11)
                if false {
                    v14 := GenS1{F0: 42, F1: (v4 << ((65535) & 63)), F2: 42}
                    _ = v14
                } else {
                    if (116 == 42) {
                        fmt.Println(fi13, fi13)
                    } else {
                        if len(v11) > 0 {
                            v11[0] = (fi10 << ((42) & 63))
                        }
                        var v15 uint32 = 2
                        _ = v15
                    }
                    if (0 >= 0) {
                        var v16 int8 = 100
                        switch v6 {
                        case "a":
                            if true {
                                switch (fi13 / (((-(fi10))) | 1)) {
                                case 32767:
                                    v2 = v2
                                    var v17 int64 = -194
                                    if (936 != 2147483647) {
                                        if len(v11) > 0 {
                                            v11[0] = (v1 / ((v1) | 1))
                                        }
                                        switch (v3 + v6) {
                                        case "foo bar":
                                            var v18 float64 = 3.14159
                                            _ = v18
                                        }
                                    }
                                    fmt.Println(v6)
                                    var v19 int16 = 42
                                    var v20 uint = 2
                                    var v21 uint8 = 1
                                    _ = v17
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                case -2147483648:
                                    v22 := GenS2{F0: v6, F1: (-(v2)), F2: 1}
                                    _ = v22
                                default:
                                    v16 += v16
                                    v6 = ("\n\t" + (v3 + v3))
                                }
                                var v23 uint8 = (v4 * (v4 << ((42) & 63)))
                                v3 = v3
                                _ = v23
                            } else {
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(v16, v16)
                                    fmt.Println(fi13, v1, v1)
                                    if false {
                                        var v24 uint16 = 255
                                        v25 := GenS3{F0: 4294967295, F1: (0.1 > 0.1), F2: (v2 << ((4294967295) & 63))}
                                        _ = v24
                                        _ = v25
                                    }
                                }
                                var v26 int32 = -2147483648
                                _ = v26
                            }
                            if false {
                                fmt.Println(v4, v2, v2)
                                if false {
                                    fmt.Println(v16)
                                }
                            } else {
                                v6 = v6
                            }
                            v4 += 8
                            v27 := GenS2{F0: "a", F1: -2147483648, F2: 692}
                            v11 = append(v11, -2147483648)
                            _ = v27
                        case "foo bar":
                            v6 = "a"
                        case "\n\t":
                            var v28 int16 = 0
                            _ = v28
                        }
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        _ = v16
                    } else {
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        v29 := GenS4{F0: uint32(v4), F1: (false && (((1.5 * float32(float64(-2.25))) > 0.1) || true))}
                        _ = v29
                    }
                    v1 += (fi10 ^ 955)
                    var v30 int = (fi10 % ((fi10) | 1))
                    var v31 int32 = (v2 - v2)
                    v11 = append(v11, int(v4))
                    _ = v30
                    _ = v31
                }
                _ = fi13
            }
        }
        _ = fi10
        _ = v11
    }
    v1 += (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
