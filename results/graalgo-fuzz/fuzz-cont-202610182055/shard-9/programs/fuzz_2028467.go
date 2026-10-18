package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int16; F2 string }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int; F1 string; F2 uint16 }
type GenS6 struct { F0 uint; F1 uint64; F2 uint64 }
func genFn0() uint8 {
    var v4 int16 = 1
    v5 := GenS1{F0: (!false), F1: -9223372036854775808}
    v4 = 829
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int8) int32 {
    v6 := GenS2{F0: 1}
    var v7 uint32 = 347
    var v8 int16 = -465
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int32 = -1
    var v2 int = 42
    var v3 uint = 65535
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := GenS3{F0: (!(false || true)), F1: 1000, F2: "\n\t"}
            fmt.Println(fi10, v3)
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := []int{(v2 - fi12), (fi12 + 190), (v2 % ((32767) | 1))}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v15 := GenS4{F0: -914, F1: 1}
        v16 := GenS2{F0: 990}
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: (!(!true)), F1: (^(v2))}
    v2 += v2
    v2 = 127
    if true {
        for fi18 := 0; fi18 < 3; fi18++ {
            for fi19 := 0; fi19 < 4; fi19++ {
                v20 := []int{127, (fi19 * -2147483648), (-(fi19))}
                v21 := GenS5{F0: (fi18 & (fi18 << (((^(v3))) & 63))), F1: ("\n\t" + "foo bar"), F2: 465}
                fmt.Println(v3)
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi18
        }
        if (!true) {
            if false {
                for fi22 := 0; fi22 < 1; fi22++ {
                    fmt.Println(v3)
                    for fi23 := 0; fi23 < 18; fi23++ {
                        fmt.Println(fi22)
                        _ = fi23
                    }
                    _ = fi22
                }
            } else {
                v3 -= v3
            }
            v24 := []int{v2, v2, (v2 % ((-2147483648) | 1))}
            v25 := GenS6{F0: v3, F1: 1, F2: 4294967295}
            _ = v24
            _ = v25
        } else {
            var v26 string = "\n\t"
            v3 -= v3
            _ = v26
        }
    }
    var v27 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v3
}
