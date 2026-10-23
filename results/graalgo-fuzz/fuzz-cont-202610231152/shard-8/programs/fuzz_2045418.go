package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 int64, p1 int, p2 int16) int64 {
    p0 = (p0 & 0)
    var v7 int16 = p2
    v8 := []int{p1, 457, (p1 | -9223372036854775808)}
    var v9 int64 = -905
    _ = len(v8)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return v9
}

func main() {
    var v1 bool = (true && true)
    v2 := make(map[uint32]bool)
    v3 := []uint32{}
    var v4 int = -46
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]string)
    fmt.Println(v4)
    v10 := GenS1{F0: "", F1: 77}
    if v1 {
        var v11 int16 = -32768
        if (!(2147483648 >= 65535)) {
            v2[2147483648] = v1
            v2[90] = (v1 && v1)
        }
        v12 := GenS1{F0: ("foo bar" + "foo bar"), F1: 255}
        _ = v11
        _ = v12
    }
    v4 = v4
    fmt.Println(v4, v4)
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 float32 = ((1.5 * 1e10) / 1e10)
        v14 = v14
        _ = fi13
        _ = v14
    }
    if v1 {
        for fi15 := 0; fi15 < 20; fi15++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = len(v6)
            _ = fi15
        }
        v16 := GenS1{F0: "hello", F1: 119}
        if (42 <= 229) {
            fmt.Println(v4)
        } else {
            v4 += v4
            v4 = v4
            for fi17 := 0; fi17 < 3; fi17++ {
                for fi18 := 0; fi18 < 6; fi18++ {
                    var v19 int64 = 42
                    if (!v1) {
                        var v20 int16 = 42
                        _ = v20
                    } else {
                        if false {
                            v4 -= -578
                        }
                        if v1 {
                            v2[2147483648] = v1
                            v19 = (v19 % (((v19 % (((^(v19))) | 1))) | 1))
                            v3 = append(v3, 851)
                        } else {
                            var v21 int8 = 100
                            _ = v21
                        }
                    }
                    v22 := GenS1{F0: "", F1: 61}
                    _ = fi18
                    _ = v19
                    _ = v22
                }
                _ = fi17
            }
        }
        _ = v16
    }
    var v23 string = ("\n\t" + "a")
    var v24 uint = 255
    switch v23 {
    case "":
        if v1 {
            fmt.Println(v4)
            if (!(!(-9223372036854775808 != (-(v4))))) {
                var v25 int64 = 9223372036854775807
                for fi26 := 0; fi26 < 12; fi26++ {
                    for fi27 := 0; fi27 < 13; fi27++ {
                        v28 := GenS2{F0: false}
                        _ = fi27
                        _ = v28
                    }
                    v29 := GenS1{F0: v23, F1: 1}
                    _ = fi26
                    _ = v29
                }
                _ = v25
            } else {
                for fi30 := 0; fi30 < 19; fi30++ {
                    fmt.Println(v23)
                    v4 += (v4 + 32768)
                    _ = fi30
                }
                v1 = v1
                v31 := GenS1{F0: v23, F1: 19}
                _ = v31
            }
            delete(v5, 2147483647)
        } else {
            if len(v3) > 0 {
                v3[0] = 428
            }
        }
    case "\n\t":
        for fi32 := 0; fi32 < 16; fi32++ {
            for fi33 := 0; fi33 < 13; fi33++ {
                v4 = fi33
                _ = fi33
            }
            _ = fi32
        }
    case "foo bar":
        if (128 < 228) {
            if (1000 >= -700) {
                delete(v6, v4)
                var v34 uint16 = 32767
                _ = v34
            }
        }
    default:
        var v35 uint32 = 42
        _ = v35
    }
    var v36 uint16 = 1
    v4 += 32767
    v1 = v1
    fmt.Println(v1, v23, v4)
    switch "hello" {
    case "hello":
        var v37 int8 = 100
        _ = v37
    case "a":
        v38 := v2[2147483647]
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
