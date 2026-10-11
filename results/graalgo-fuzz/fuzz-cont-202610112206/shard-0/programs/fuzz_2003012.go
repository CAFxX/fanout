package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
func genFn0(p0 int16) bool {
    v6 := []int{1, -1, -1, -689}
    v7 := GenS1{F0: (^(p0))}
    _ = p0
    _ = v6
    _ = v7
    return ((!true) && false)
}
func genFn1() int32 {
    var v8 int8 = 127
    fmt.Println(v8)
    v9 := GenS1{F0: 0}
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 uint64 = 606
    var v2 float32 = -2.25
    var v3 int8 = -128
    v4 := make(map[string]bool)
    var v5 uint8 = 246
    var v10 int = -600
    var v11 bool = ((genFn1() < int32(v3)) && ((0.1 - 0.1) <= (1.5 * 0.0)))
    v10 -= v10
    var v12 int32 = -1
    v2 = (v2 - v2)
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            delete(v4, "")
            _ = fi14
        }
        _ = fi13
    }
    fmt.Println(v5, v3, v12)
    var v15 uint = 150
    if true {
        v16 := []int{(v10 << (((v15 ^ 18446744073709551615)) & 63)), v10, int(v5), (v10 | (-(v10)))}
        _ = v16
    }
    switch "" {
    case "hello":
        if (v1 != (v1 + 4294967295)) {
            var v17 bool = false
            _ = v17
        } else {
            if (v1 <= 18446744073709551615) {
                var v18 uint32 = 2147483648
                v15 = v15
                var v19 uint16 = 1
                _ = v18
                _ = v19
            } else {
                v10 -= (^(v10))
                if v11 {
                    switch (-(v10)) {
                    case 1:
                        if (!v11) {
                            var v20 uint = v15
                            _ = v20
                        } else {
                            v21 := []int{(v10 | v10)}
                            _ = v21
                        }
                    }
                } else {
                    for fi22 := 0; fi22 < 8; fi22++ {
                        if (v5 >= (v5 % (((v5 + v5)) | 1))) {
                            switch v10 {
                            case 9223372036854775807:
                                fmt.Println(v5, v1, v12)
                            case -76:
                                var v23 int16 = -32768
                                _ = v23
                            }
                        } else {
                            switch v10 {
                            case -343:
                                var v24 int16 = -32768
                                _ = v24
                            case 1:
                                v15 = v15
                            }
                            v3 = v3
                        }
                        _ = fi22
                    }
                    v3 -= v3
                }
                v2 = v2
            }
            if (float64(((-2.25 - (((v2 * 1.5) + v2) / (v2 / 1.0))) - (float32(((-2.25 + -2.25) - float64(float32(float64(v2))))) / (float32(1e10) * ((0.0 + ((v2 + float32(0.0)) - 1.5)) + v2))))) >= (3.14159 / -2.25)) {
                v12 += (v12 / ((v12) | 1))
                var v25 uint16 = 65535
                if false {
                    if (v15 < 676) {
                        v2 = float32((3.14159 - -2.25))
                    }
                }
                _ = v25
            } else {
                v26 := GenS1{F0: 32767}
                _ = v26
            }
            var v27 uint32 = 411
            _ = v27
        }
    default:
        if ((v11 || (v5 <= (v5 / ((140) | 1)))) && v11) {
            switch "hello" {
            case "foo bar":
                var v28 float64 = 3.14159
                _ = v28
            case "\n\t":
                switch ("" + ("foo bar" + "hello")) {
                case "foo bar":
                    if v11 {
                        var v29 uint16 = 460
                        _ = v29
                    }
                case "hello":
                    for fi30 := 0; fi30 < 9; fi30++ {
                        var v31 int32 = v12
                        _ = fi30
                        _ = v31
                    }
                default:
                    var v32 int8 = v3
                    _ = v32
                }
            }
            v33 := GenS1{F0: 138}
            _ = v33
        } else {
            if v11 {
                fmt.Println(v11, v2)
                var v34 int16 = 251
                if v11 {
                    var v35 float64 = -2.25
                    for fi36 := 0; fi36 < 9; fi36++ {
                        v10 = (fi36 << ((v15) & 63))
                        _ = fi36
                    }
                    _ = v35
                } else {
                    var v37 uint16 = 32767
                    _ = v37
                }
                _ = v34
            }
        }
    }
    if v11 {
        v38 := GenS2{F0: (int64(v1) <= -9223372036854775808)}
        _ = v38
    }
    var v39 uint8 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
