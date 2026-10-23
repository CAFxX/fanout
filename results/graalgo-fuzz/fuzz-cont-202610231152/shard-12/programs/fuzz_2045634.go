package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int32 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 bool }
func genFn0() int64 {
    var v6 bool = (false || true)
    var v7 int16 = -32768
    fmt.Println(v6, v7)
    v8 := GenS1{F0: "foo bar", F1: 0, F2: -615}
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int8) uint16 {
    if (false && ((0 <= -1) && (43 != 1))) {
        var v9 int16 = 323
        for fi10 := 0; fi10 < 17; fi10++ {
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 int16 = -1
                if true {
                    v13 := GenS2{F0: 1, F1: 2147483648, F2: 1}
                    v14 := []int{}
                    _ = v13
                    _ = v14
                }
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
        _ = v9
    }
    if true {
        v15 := []int{4294967295, 42, -9223372036854775808, 42}
        fmt.Println(p0)
        v16 := GenS2{F0: 536, F1: 41, F2: 4294967295}
        _ = v15
        _ = v16
    }
    var v17 uint16 = 42
    _ = p0
    _ = v17
    return (v17 % ((118) | 1))
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 0
    var v3 int64 = 9223372036854775807
    var v4 float64 = 1.5
    var v5 string = ("foo bar" + "a")
    for fi18 := 0; fi18 < 11; fi18++ {
        switch ((v5 + "\n\t") + v5) {
        case "hello":
            for fi19 := 0; fi19 < 8; fi19++ {
                var v20 float32 = 0.0
                switch (fi19 % (((fi18 * (fi18 >> ((4294967295) & 63)))) | 1)) {
                case 128:
                    var v21 bool = (-1 >= 128)
                    _ = v21
                }
                v4 = v4
                _ = fi19
                _ = v20
            }
        case "a":
            if ((16 != 127) || ((42 == 521) || false)) {
                fmt.Println(v5, fi18, v4)
                v22 := []int{(fi18 >> ((510) & 63)), fi18}
                _ = v22
            } else {
                if (!true) {
                    if (2 >= 65535) {
                        v23 := []int{(fi18 * (-(fi18))), fi18, (fi18 / ((-2147483648) | 1)), fi18}
                        v24 := GenS2{F0: (v3 & v2), F1: 1, F2: 837}
                        _ = v23
                        _ = v24
                    } else {
                        fmt.Println(v2, v3, v4)
                    }
                } else {
                    var v25 string = v5
                    for fi26 := 0; fi26 < 17; fi26++ {
                        fmt.Println(fi26)
                        if (float64((((-2.25 + (3.14159 + 0.0)) / 0.1) * float32(v4))) >= v4) {
                            var v27 uint8 = 2
                            v28 := GenS1{F0: (v25 + v25), F1: 2147483648, F2: 75}
                            _ = v27
                            _ = v28
                        }
                        fmt.Println(fi26, v2, v5)
                        _ = fi26
                    }
                    _ = v25
                }
                for fi29 := 0; fi29 < 9; fi29++ {
                    switch fi18 {
                    case -9223372036854775808:
                        if (1 == 682) {
                            v2 += -9223372036854775808
                            v30 := GenS1{F0: v5, F1: 29, F2: 63}
                            fmt.Println(v2, fi18)
                            _ = v30
                        } else {
                            v3 = (v3 ^ (v3 & v2))
                        }
                    case 32767:
                        v31 := []int{1}
                        _ = v31
                    case -592:
                        fmt.Println(v4)
                    default:
                        if (722 != 42) {
                            var v32 uint8 = 42
                            v32 = 127
                            v3 -= (v3 << ((549) & 63))
                            _ = v32
                        }
                        if (!(89 >= 127)) {
                            var v33 uint8 = 133
                            var v34 uint8 = 127
                            v33 -= 128
                            _ = v33
                            _ = v34
                        }
                    }
                    _ = fi29
                }
            }
        }
        _ = fi18
    }
    v4 -= float64(((1e10 - 0.1) + 0.1))
    var v35 int = 32767
    for fi36 := 0; fi36 < 9; fi36++ {
        for fi37 := 0; fi37 < 15; fi37++ {
            v38 := GenS3{F0: (((!false) && true) && false)}
            _ = fi37
            _ = v38
        }
        v35 = 9223372036854775807
        _ = fi36
    }
    fmt.Println(v3, v35)
    var v39 uint64 = 139
    var v40 int64 = (v3 / ((v3) | 1))
    for fi41 := 0; fi41 < 13; fi41++ {
        var v42 float64 = v4
        _ = fi41
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
