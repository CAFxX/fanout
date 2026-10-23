package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint64) uint16 {
    switch "\n\t" {
    case "":
        for fi7 := 0; fi7 < 9; fi7++ {
            switch "hello" {
            case "hello":
                if true {
                    v8 := GenS1{F0: (fi7 % ((fi7) | 1))}
                    _ = v8
                }
            case "foo bar":
                switch 359 {
                case 178:
                    if false {
                        v9 := []int{}
                        _ = v9
                    } else {
                        var v10 int32 = -337
                        _ = v10
                    }
                case -1:
                    for fi11 := 0; fi11 < 5; fi11++ {
                        v12 := []int{fi11, fi7, (fi11 + -9223372036854775808)}
                        _ = fi11
                        _ = v12
                    }
                case -689:
                    p0 -= 107
                }
            }
            v13 := []int{(fi7 ^ (fi7 * (fi7 * fi7)))}
            var v14 uint32 = 42
            _ = fi7
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := []int{(fi15 ^ 127), -969}
        _ = fi15
        _ = v16
    }
    if true {
        for fi17 := 0; fi17 < 19; fi17++ {
            fmt.Println(p0, p0)
            _ = fi17
        }
        for fi18 := 0; fi18 < 2; fi18++ {
            if (false && (433 >= 145)) {
                p0 = (p0 * p0)
            }
            _ = fi18
        }
    } else {
        switch 2147483647 {
        case -2147483648:
            for fi19 := 0; fi19 < 11; fi19++ {
                v20 := []int{fi19}
                _ = fi19
                _ = v20
            }
        default:
            v21 := GenS1{F0: 32767}
            _ = v21
        }
        p0 += p0
    }
    v22 := []int{2147483647, 32768, -1, 127}
    var v23 int = -756
    _ = p0
    _ = v22
    _ = v23
    return uint16(p0)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 2
    var v3 int = -9223372036854775808
    v4 := make(map[uint32]uint32)
    var v5 string = ("\n\t" + "a")
    v6 := []int{(v3 ^ (v3 ^ v3)), v3}
    v24 := GenS2{F0: genFn0(718), F1: 2, F2: 361}
    v2 = v2
    var v25 uint64 = 4294967295
    var v26 int16 = -752
    if false {
        v5 = v5
        v27 := GenS2{F0: 390, F1: 1, F2: 797}
        var v28 int8 = -5
        _ = v27
        _ = v28
    } else {
        _ = len(v6)
    }
    v3 = (^(v3))
    v26 -= -745
    if (false && (((531 != 17) && true) && (v3 > v3))) {
        if (509 != 1) {
            v26 += (v26 + v26)
        } else {
            var v29 int64 = 0
            switch (v3 - v3) {
            case 32767:
                v30 := GenS3{F0: 2}
                _ = v30
            case -2147483648:
                v2 -= v2
            default:
                if false {
                    if len(v6) > 0 {
                        v6[0] = v3
                    }
                }
            }
            _ = v29
        }
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        fmt.Println(v5)
        _ = fi31
    }
    v5 = v5
    v26 += (v26 % ((v26) | 1))
    var v32 uint16 = 255
    fmt.Println(v32)
    v32 += 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
