package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 int }
func genFn0(p0 int16, p1 uint8, p2 uint32) uint16 {
    p2 -= p2
    p0 -= p0
    var v7 int64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32767
}

func main() {
    var v1 string = ("" + "foo bar")
    var v2 int = 32768
    var v3 uint = 2
    var v4 int8 = 100
    v5 := make(map[string]uint32)
    var v6 int16 = 1000
    for fi8 := 0; fi8 < 19; fi8++ {
        if (!(-1 != -814)) {
            var v9 bool = true
            _ = v9
        }
        _ = fi8
    }
    if false {
        if (2 > 1) {
            v10 := []int{v2}
            switch v1 {
            case "foo bar":
                _ = len(v10)
            case "hello":
                if (-867 <= v2) {
                    var v11 uint32 = 96
                    if true {
                        var v12 int16 = v6
                        var v13 string = (v1 + "")
                        _ = v12
                        _ = v13
                    } else {
                        if len(v10) > 0 {
                            v10[0] = (v2 * (v2 >> ((v3) & 63)))
                        }
                        var v14 int16 = (v6 | v6)
                        for fi15 := 0; fi15 < 18; fi15++ {
                            delete(v5, "foo bar")
                            v16 := GenS1{F0: v14, F1: (v3 >> ((42) & 63)), F2: 195}
                            _ = fi15
                            _ = v16
                        }
                        _ = v14
                    }
                    _ = len(v10)
                    _ = v11
                }
            }
            _ = v10
        } else {
            v17 := GenS2{F0: 65535, F1: 64}
            var v18 int16 = (v6 ^ int16(v4))
            _ = v17
            _ = v18
        }
    } else {
        for fi19 := 0; fi19 < 7; fi19++ {
            v2 = fi19
            _ = fi19
        }
    }
    v20 := []int{(v2 >> (((v3 | (^(v3)))) & 63)), 2147483647, 4294967295, (v2 - (v2 / ((v2) | 1)))}
    var v21 int64 = -1
    fmt.Println(v3)
    v22 := v5["a"]
    if false {
        v23 := GenS2{F0: 843, F1: (v22 ^ v22)}
        var v24 string = ((v1 + (v1 + v1)) + (v1 + "a"))
        _ = v23
        _ = v24
    } else {
        if ((v6 & (v6 << (((v3 * v3)) & 63))) < v6) {
            for fi25 := 0; fi25 < 12; fi25++ {
                v26 := GenS3{F0: 209}
                v27 := GenS4{F0: 32767}
                _ = fi25
                _ = v26
                _ = v27
            }
            if ((-2147483648 != 388) && (!(true && (v21 == v21)))) {
                v28 := GenS5{F0: -32768, F1: v2}
                v29 := GenS5{F0: (v6 - v6), F1: 9223372036854775807}
                _ = v28
                _ = v29
            } else {
                switch v2 {
                case 32768:
                    v3 = v3
                case 128:
                    var v30 int16 = 0
                    _ = v30
                default:
                    v31 := GenS2{F0: 2, F1: 2147483647}
                    _ = v31
                }
            }
        } else {
            v32 := GenS2{F0: 0, F1: (v22 ^ (v22 & (^(v22))))}
            for fi33 := 0; fi33 < 19; fi33++ {
                var v34 float32 = 3.14159
                _ = fi33
                _ = v34
            }
            var v35 uint16 = 559
            _ = v32
            _ = v35
        }
    }
    v4 = int8(v6)
    fmt.Println(v4, v4, v22)
    for fi36 := 0; fi36 < 3; fi36++ {
        _ = len(v5)
        var v37 bool = true
        _ = fi36
        _ = v37
    }
    switch 0 {
    case 42:
        var v38 uint8 = 42
        _ = v38
    }
    v21 += -705
    fmt.Println(v22, v6, v4)
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
