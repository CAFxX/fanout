package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int8 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint16) int64 {
    var v7 int8 = 42
    for fi8 := 0; fi8 < 12; fi8++ {
        p0 = p0
        fmt.Println(v7, p0)
        v9 := GenS1{F0: 1, F1: -1, F2: v7}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v7
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = 0.0
    var v3 float32 = (0.1 + 1e10)
    var v4 uint32 = 933
    v5 := []uint32{}
    v6 := make(map[uint32]bool)
    _ = len(v5)
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 int64 = 617
        if len(v5) > 0 {
            v5[0] = (v4 >> ((275) & 63))
        }
        if true {
            var v12 uint = 350
            _ = v12
        } else {
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := GenS1{F0: 128, F1: (v11 >> ((206) & 63)), F2: -1}
                var v15 string = "hello"
                _ = fi13
                _ = v14
                _ = v15
            }
            v4 = v4
        }
        _ = fi10
        _ = v11
    }
    switch "hello" {
    case "\n\t":
        delete(v1, 2147483647)
    case "a":
        for fi16 := 0; fi16 < 15; fi16++ {
            if (!true) {
                for fi17 := 0; fi17 < 5; fi17++ {
                    v1[(fi17 % ((fi17) | 1))] = v4
                    if (false || (true && (true || ((v4 ^ 2147483647) <= (^(v4)))))) {
                        var v18 int = (fi16 >> ((0) & 63))
                        var v19 uint8 = 164
                        if (!true) {
                            v18 += (^(fi16))
                            var v20 bool = false
                            _ = v20
                        }
                        _ = v18
                        _ = v19
                    }
                    var v21 int8 = -1
                    _ = fi17
                    _ = v21
                }
            } else {
                v22 := GenS2{F0: 927}
                var v23 float64 = (v2 * 3.14159)
                _ = v22
                _ = v23
            }
            _ = fi16
        }
    case "":
        v5 = append(v5, v4)
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        if false {
            if true {
                v3 = 0.1
            } else {
                var v25 uint8 = 0
                _ = v25
            }
        } else {
            v26 := GenS2{F0: 0}
            fmt.Println(v3)
            v4 = v4
            _ = v26
        }
        _ = fi24
    }
    var v27 int64 = -9223372036854775808
    v28 := GenS2{F0: 18446744073709551615}
    if ((v3 > v3) && false) {
        v29 := GenS1{F0: 255, F1: (v27 >> ((65535) & 63)), F2: -128}
        _ = v29
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        var v31 uint8 = 39
        _ = fi30
        _ = v31
    }
    if (!false) {
        var v32 string = ("a" + "")
        for fi33 := 0; fi33 < 16; fi33++ {
            v32 = v32
            _ = fi33
        }
        _ = v32
    } else {
        var v34 string = ("hello" + ("foo bar" + ""))
        var v35 int64 = v27
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 20; fi36++ {
        v3 = v3
        switch 32768 {
        case -695:
            _ = len(v6)
        case 42:
            for fi37 := 0; fi37 < 19; fi37++ {
                v27 = v27
                _ = fi37
            }
        }
        _ = fi36
    }
    var v38 int16 = 32767
    if (128 >= 255) {
        v5 = append(v5, (v4 + 2147483648))
    }
    v2 = 1e10
    fmt.Println(v4, v27, v3)
    v39 := GenS1{F0: 214, F1: -9223372036854775808, F2: 1}
    v40 := v6[v4]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
