package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint32 }
func genFn0(p0 uint8, p1 uint16, p2 int) int {
    fmt.Println(p2)
    v5 := []int{p2, (p2 / ((131) | 1)), (p2 | p2), -2147483648}
    p0 += (p0 >> ((42) & 63))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v7 bool = (!(true || true))
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p2 - (p2 | (^(p2))))
}
func genFn1() uint {
    var v8 int8 = -18
    if true {
        v9 := []int{genFn0(151, 32767, -2147483648), 0, 960, 953}
        _ = v9
    } else {
        v10 := GenS1{F0: -2147483648}
        v8 -= v8
        _ = v10
    }
    v11 := []int{42, 32767}
    var v12 uint16 = 45
    fmt.Println(v12, v12)
    _ = v11
    _ = v12
    _ = v8
    return 4294967295
}

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]bool)
    var v3 uint64 = 18446744073709551615
    var v4 float64 = 3.14159
    fmt.Println(v1)
    if (-2.25 == (1.5 * (-2.25 + 1.5))) {
        v13 := GenS1{F0: 4294967295}
        _ = v13
    } else {
        if false {
            v14 := []int{42, 334}
            fmt.Println(v4, v3, v4)
            switch 4294967295 {
            case 1:
                for fi15 := 0; fi15 < 7; fi15++ {
                    v16 := GenS2{F0: 555}
                    _ = fi15
                    _ = v16
                }
            case -9223372036854775808:
                v17 := GenS3{F0: true, F1: (v3 % ((4294967295) | 1))}
                _ = v17
            default:
                v3 += (v3 | 79)
            }
            _ = v14
        } else {
            v2[2147483647] = false
        }
    }
    v1 -= (-(v1))
    v1 -= 1
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 int = -1
        v4 += v4
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: 2147483647}
    if (2 > 2) {
        v21 := GenS4{F0: -196}
        _ = v21
    }
    v22 := GenS3{F0: true, F1: (v3 % (((v3 / ((v3) | 1))) | 1))}
    _ = len(v2)
    for fi23 := 0; fi23 < 11; fi23++ {
        v24 := []int{fi23, (-(fi23)), (fi23 + fi23), (fi23 - genFn0(2, 1, (fi23 * -2147483648)))}
        _ = fi23
        _ = v24
    }
    v4 = 0.1
    var v25 int8 = 0
    for fi26 := 0; fi26 < 12; fi26++ {
        delete(v2, 65535)
        if false {
            if false {
                fmt.Println(v3)
            } else {
                if (true && (v4 <= ((3.14159 * 1e10) - 1e10))) {
                    for fi27 := 0; fi27 < 12; fi27++ {
                        v28 := GenS5{F0: v25}
                        _ = fi27
                        _ = v28
                    }
                    v3 = v3
                    if (fi26 == genFn0(176, 32767, (^(fi26)))) {
                        fmt.Println(v25, v1, v3)
                    } else {
                        for fi29 := 0; fi29 < 16; fi29++ {
                            v30 := GenS5{F0: v25}
                            _ = fi29
                            _ = v30
                        }
                    }
                }
            }
            switch "" {
            case "\n\t":
                if ((846 <= 696) || true) {
                    var v31 uint8 = 242
                    _ = v31
                }
            case "a":
                var v32 int8 = -128
                _ = v32
            case "hello":
                v25 += v25
            }
            v33 := GenS1{F0: (fi26 >> ((genFn1()) & 63))}
            _ = v33
        } else {
            var v34 int32 = int32(v25)
            var v35 uint16 = 32767
            for fi36 := 0; fi36 < 3; fi36++ {
                var v37 string = "a"
                _ = fi36
                _ = v37
            }
            _ = v34
            _ = v35
        }
        _ = fi26
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        for fi39 := 0; fi39 < 7; fi39++ {
            v1 += v1
            _ = fi39
        }
        var v40 bool = ((("\n\t" + ("hello" + "a")) < (("" + (("hello" + ("foo bar" + "")) + "")) + "foo bar")) || (4294967295 > 950))
        _ = fi38
        _ = v40
    }
    for fi41 := 0; fi41 < 14; fi41++ {
        if true {
            fmt.Println(v3)
            v4 -= v4
            var v42 int64 = 42
            _ = v42
        }
        if (2 == (v3 | (v3 ^ v3))) {
            v43 := GenS6{F0: -867}
            var v44 int = fi41
            _ = v43
            _ = v44
        }
        v45 := GenS7{F0: uint32(v3)}
        _ = fi41
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v25
    _ = v3
    _ = v4
}
