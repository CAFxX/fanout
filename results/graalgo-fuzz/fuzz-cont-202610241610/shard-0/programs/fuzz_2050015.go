package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int, p1 uint64, p2 uint8) uint32 {
    if (-282 <= 42) {
        fmt.Println(p1, p1, p2)
        for fi4 := 0; fi4 < 19; fi4++ {
            var v5 float64 = 0.1
            var v6 float32 = 1e10
            for fi7 := 0; fi7 < 13; fi7++ {
                v5 -= v5
                v8 := GenS1{F0: -32768, F1: 1, F2: 1000}
                v9 := []int{}
                var v10 int = 1000
                var v11 bool = (940 != 0)
                fmt.Println(v10, p2)
                p1 -= (^(p1))
                if v11 {
                    if len(v9) > 0 {
                        v9[0] = p0
                    }
                } else {
                    var v12 uint16 = 0
                    v13 := GenS2{F0: 2, F1: 0}
                    if len(v9) > 0 {
                        v9[0] = fi4
                    }
                    _ = v12
                    _ = v13
                }
                _ = fi7
                _ = v10
                _ = v11
                _ = v8
                _ = v9
            }
            _ = fi4
            _ = v5
            _ = v6
        }
        p0 += p0
    }
    v14 := GenS2{F0: 7, F1: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 65535
}
func genFn1() int {
    var v15 uint16 = 2
    fmt.Println(v15)
    _ = v15
    return -229
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    var v3 float32 = float32(3.14159)
    if (!((2147483648 >= 4294967295) || (!true))) {
        var v16 uint8 = 1
        v3 = (0.0 / v3)
        v16 -= 0
        _ = v16
    } else {
        v3 -= float32(0.0)
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        v3 -= float32(0.1)
        _ = fi17
    }
    switch -1 {
    case 32767:
        for fi18 := 0; fi18 < 8; fi18++ {
            v19 := []int{fi18, (fi18 - fi18), (fi18 ^ genFn1()), genFn1()}
            var v20 int16 = 32767
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 uint32 = 688
                v23 := GenS3{F0: 2}
                if (4294967295 >= (fi21 * -455)) {
                    var v24 float32 = v3
                    var v25 bool = (false || false)
                    _ = v24
                    _ = v25
                }
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi18
            _ = v19
            _ = v20
        }
    case 996:
        delete(v2, genFn0(42, 18446744073709551615, 255))
    default:
        v3 -= v3
    }
    v26 := GenS3{F0: 18446744073709551615}
    switch "foo bar" {
    case "a":
        v2[genFn0(42, 42, 28)] = -2147483648
    case "\n\t":
        v1[42] = 4294967295
    default:
        fmt.Println(v3)
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        if (("foo bar" + "\n\t") < ("foo bar" + ("a" + "a"))) {
            v3 += (0.1 * ((v3 * (1e10 - 0.0)) / 1e10))
            for fi28 := 0; fi28 < 19; fi28++ {
                var v29 uint8 = 86
                _ = fi28
                _ = v29
            }
        } else {
            if (v3 < 1.5) {
                var v30 float32 = (v3 - 3.14159)
                _ = v30
            }
        }
        v3 = (1e10 + float32(3.14159))
        for fi31 := 0; fi31 < 8; fi31++ {
            var v32 float64 = 1e10
            var v33 int = (fi27 | (fi27 << ((1) & 63)))
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi27
    }
    var v34 uint8 = 128
    v34 += v34
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v34
}
