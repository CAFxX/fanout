package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint16, p1 uint32) int {
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := []int{fi6}
        if true {
            v8 := GenS1{F0: fi6, F1: 100, F2: -996}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v9 uint8 = 1
            _ = v8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v10 int64 = -540
    _ = p0
    _ = p1
    _ = v10
    return 128
}
func genFn1(p0 string, p1 int32, p2 uint64) string {
    v11 := []int{127}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return (p0 + "hello")
}

func main() {
    var v1 int = 42
    var v2 int16 = 1000
    v3 := make(map[int]bool)
    var v4 uint32 = 2147483647
    var v5 int = v1
    v1 += v1
    v12 := GenS1{F0: (^(v5)), F1: -128, F2: -2147483648}
    for fi13 := 0; fi13 < 3; fi13++ {
        v5 = 667
        for fi14 := 0; fi14 < 3; fi14++ {
            v4 += v4
            v4 += 642
            _ = fi14
        }
        v2 = v2
        _ = fi13
    }
    v15 := GenS1{F0: 670, F1: 1, F2: 0}
    var v16 uint16 = 32767
    var v17 uint8 = 255
    for fi18 := 0; fi18 < 13; fi18++ {
        if true {
            if ((((fi18 - -9223372036854775808) != 2147483647) && (!(false && ((!false) && (false || true))))) && false) {
                for fi19 := 0; fi19 < 18; fi19++ {
                    v20 := v3[fi19]
                    v21 := v3[(v1 * fi18)]
                    v20 = (1 != 18446744073709551615)
                    v2 -= (v2 % ((v2) | 1))
                    v22 := GenS2{F0: 9223372036854775807}
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                if true {
                    var v23 string = "hello"
                    _ = v23
                }
            }
            fmt.Println(v2, v1)
            v16 -= v16
        } else {
            v4 += v4
            v24 := []int{0, fi18}
            _ = v24
        }
        _ = fi18
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 uint = 0
        v3[(v1 + (v5 + int(v2)))] = ((v4 | (v4 & v4)) < v4)
        switch (v5 + genFn0(v16, (v4 / ((800) | 1)))) {
        case 42:
            v4 = (v4 % (((v4 & (v4 & v4))) | 1))
        }
        _ = fi25
        _ = v26
    }
    var v27 uint = 605
    v1 += (v5 << ((65535) & 63))
    var v28 float64 = 0.1
    v28 = 3.14159
    for fi29 := 0; fi29 < 10; fi29++ {
        fmt.Println(fi29, v17)
        _ = fi29
    }
    var v30 uint8 = (v17 % ((v17) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
