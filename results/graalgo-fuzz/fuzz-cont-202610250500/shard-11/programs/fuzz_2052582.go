package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint32) uint32 {
    var v4 int = 515
    if ((77 <= 42) && (4294967295 == (p0 / ((p0) | 1)))) {
        fmt.Println(v4)
    }
    v5 := []int{v4}
    _ = len(v5)
    _ = p0
    _ = v4
    _ = v5
    return (p0 / (((^(p0))) | 1))
}
func genFn1(p0 uint64) uint64 {
    p0 += 42
    if false {
        var v6 uint8 = 0
        _ = v6
    }
    _ = p0
    return (p0 * p0)
}

func main() {
    var v1 int64 = 0
    v2 := []string{"\n\t", "foo bar", "a", "hello"}
    var v3 uint16 = 433
    v7 := GenS1{F0: 18446744073709551615, F1: uint(v3)}
    v3 = v3
    var v8 bool = false
    v8 = (32767 <= 128)
    v9 := GenS1{F0: 18446744073709551615, F1: 4294967295}
    fmt.Println(v3, v1)
    var v10 int32 = 1
    v3 += v3
    if false {
        v11 := GenS2{F0: 42}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            switch (fi12 << ((729) & 63)) {
            case 25:
                v13 := GenS2{F0: 55}
                _ = v13
            }
            v8 = ((!v8) || (225 < 1))
            v10 += int32(fi12)
            _ = fi12
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v14 := GenS1{F0: 42, F1: 2}
    v2 = append(v2, ("\n\t" + "foo bar"))
    v1 -= (v1 << ((0) & 63))
    if (uint32(v3) == 118) {
        _ = len(v2)
        v2 = append(v2, ("a" + ""))
        var v15 float64 = 0.0
        _ = v15
    }
    v1 = (v1 & int64(v3))
    fmt.Println(v3, v10)
    for fi16 := 0; fi16 < 16; fi16++ {
        fmt.Println(v1)
        switch "" {
        case "a":
            fmt.Println(v1, v1)
        case "hello":
            if (!v8) {
                for fi17 := 0; fi17 < 16; fi17++ {
                    v8 = false
                    v2 = append(v2, ("a" + ("a" + "\n\t")))
                    var v18 float32 = float32(float64(1e10))
                    _ = fi17
                    _ = v18
                }
                for fi19 := 0; fi19 < 7; fi19++ {
                    var v20 uint64 = 2
                    v21 := GenS2{F0: 116}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                for fi22 := 0; fi22 < 17; fi22++ {
                    if (!(42 >= 107)) {
                        if v8 {
                            fmt.Println(fi22)
                        }
                    } else {
                        v1 += (v1 >> ((519) & 63))
                    }
                    var v23 uint = 1
                    var v24 uint64 = 1
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            } else {
                var v25 float32 = 3.14159
                fmt.Println(v25, v8, v10)
                _ = v25
            }
        }
        _ = fi16
    }
    fmt.Println(v10, v1, v3)
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
