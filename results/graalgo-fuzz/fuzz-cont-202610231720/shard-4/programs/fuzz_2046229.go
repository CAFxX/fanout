package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint64 }
func genFn0(p0 int64, p1 int32, p2 string) int64 {
    v7 := GenS1{F0: 839, F1: -128}
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p0 * p0)
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = -1
    var v3 int8 = 127
    v4 := []int{(^(v2)), v2, (-(v2))}
    v5 := []string{"", (("hello" + "\n\t") + (("foo bar" + "a") + "a")), "foo bar"}
    var v6 bool = false
    if (2147483648 > 242) {
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS2{F0: 42, F1: 389, F2: 252}
            v11 := GenS1{F0: 65535, F1: 42}
            _ = fi9
            _ = v10
            _ = v11
        }
        switch v2 {
        case 32767:
            var v12 int64 = 42
            _ = v12
        case 127:
            v6 = (v3 > 100)
        case 675:
            fmt.Println(v3)
        }
        switch v2 {
        case 127:
            var v13 bool = v6
            _ = v13
        case -2147483648:
            v14 := GenS2{F0: 523, F1: 65535, F2: 2}
            _ = v14
        case 9223372036854775807:
            var v15 uint16 = 32767
            _ = v15
        default:
            v6 = true
        }
    } else {
        fmt.Println(v2, v3, v2)
    }
    fmt.Println(v2, v3, v6)
    for fi16 := 0; fi16 < 11; fi16++ {
        switch v2 {
        case 9223372036854775807:
            var v17 uint32 = 42
            _ = v17
        case 32768:
            switch 844 {
            case -1:
                if (!(true && (92 <= 0))) {
                    v2 = fi16
                } else {
                    if v6 {
                        for fi18 := 0; fi18 < 19; fi18++ {
                            var v19 int8 = v3
                            _ = fi18
                            _ = v19
                        }
                        v1[1] = 678
                    } else {
                        _ = len(v1)
                        var v20 string = ("" + "")
                        if true {
                            var v21 bool = (-9223372036854775808 == int64(fi16))
                            _ = v21
                        } else {
                            var v22 int32 = -529
                            _ = v22
                        }
                        _ = v20
                    }
                    var v23 uint64 = 18446744073709551615
                    var v24 uint = 0
                    _ = v23
                    _ = v24
                }
            case 32767:
                var v25 uint32 = 1
                _ = v25
            case 127:
                var v26 bool = (v6 || v6)
                _ = v26
            }
        case 1:
            v3 += (v3 >> ((0) & 63))
        default:
            if ((^(fi16)) >= v2) {
                v27 := GenS1{F0: 65535, F1: v3}
                _ = v27
            }
        }
        _ = fi16
    }
    v2 = (-(v2))
    var v28 int8 = 64
    v28 = -128
    delete(v1, 65535)
    v6 = ((((32767 < 1) && (((v28 + v3) > 127) || (-1 >= -430))) || (int16(v28) < 42)) || (65535 != 0))
    v29 := GenS2{F0: -32768, F1: 0, F2: 1}
    v30 := GenS1{F0: 255, F1: -79}
    var v31 int8 = (v3 ^ 127)
    v2 -= 2147483647
    delete(v1, 473)
    switch (v2 << ((4294967295) & 63)) {
    case 2147483647:
        var v32 float32 = (1.5 * 0.1)
        _ = v32
    }
    for fi33 := 0; fi33 < 4; fi33++ {
        var v34 uint32 = 4294967295
        for fi35 := 0; fi35 < 8; fi35++ {
            var v36 int64 = 42
            switch "" {
            case "\n\t":
                v3 = 111
            case "foo bar":
                var v37 bool = false
                _ = v37
            }
            fmt.Println(v36)
            _ = fi35
            _ = v36
        }
        v34 += (v34 % (((v34 + (v34 - 42))) | 1))
        _ = fi33
        _ = v34
    }
    delete(v1, 4294967295)
    fmt.Println(v2, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
