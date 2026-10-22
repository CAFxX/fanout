package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
func genFn0() uint16 {
    for fi4 := 0; fi4 < 19; fi4++ {
        v5 := []int{}
        var v6 uint = 0
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int16 = -283
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: 1, F1: 100}
        if true {
            if (42 >= 2147483647) {
                fmt.Println(v7, v7)
                v7 += 32767
                v10 := []int{(fi8 >> ((255) & 63)), fi8, fi8}
                _ = v10
            }
            fmt.Println(fi8, fi8)
        }
        _ = fi8
        _ = v9
    }
    _ = v7
    return 554
}
func genFn1(p0 int16) uint16 {
    var v11 string = ("a" + "\n\t")
    var v12 int16 = p0
    _ = p0
    _ = v11
    _ = v12
    return genFn0()
}

func main() {
    var v1 float64 = 0.0
    var v2 bool = true
    v3 := make(map[string]uint32)
    delete(v3, ("hello" + "a"))
    delete(v3, "\n\t")
    fmt.Println(v2)
    if (128 >= 255) {
        var v13 bool = (!v2)
        var v14 string = "\n\t"
        _ = v13
        _ = v14
    } else {
        var v15 int = -9223372036854775808
        _ = v15
    }
    var v16 uint32 = 42
    if false {
        for fi17 := 0; fi17 < 14; fi17++ {
            v16 = v16
            v2 = (v2 && v2)
            v18 := []int{1}
            _ = fi17
            _ = v18
        }
    } else {
        for fi19 := 0; fi19 < 13; fi19++ {
            fmt.Println(fi19)
            _ = fi19
        }
    }
    fmt.Println(v16)
    v20 := GenS1{F0: -9223372036854775808, F1: int8(v16)}
    if true {
        if (float64(3.14159) < v1) {
            var v21 int8 = 29
            v22 := GenS1{F0: 511, F1: (v21 >> ((429) & 63))}
            _ = v21
            _ = v22
        } else {
            v1 = ((1.5 * (v1 * v1)) * v1)
            if ((1e10 < float64(1.5)) || v2) {
                var v23 uint8 = 255
                v24 := GenS1{F0: -9223372036854775808, F1: 100}
                var v25 int8 = 0
                _ = v23
                _ = v24
                _ = v25
            } else {
                v16 = (v16 + (v16 ^ (^(v16))))
                var v26 int = 0
                if (!false) {
                    v26 -= -2147483648
                } else {
                    if (0 < -34) {
                        _ = len(v3)
                    } else {
                        var v27 uint16 = 65535
                        _ = v27
                    }
                }
                _ = v26
            }
            fmt.Println(v1)
        }
        var v28 int = int(v16)
        v2 = false
        _ = v28
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        _ = len(v3)
        _ = fi29
    }
    v2 = v2
    for fi30 := 0; fi30 < 16; fi30++ {
        fmt.Println(fi30, v16)
        v2 = (3.14159 == 3.14159)
        if v2 {
            var v31 float64 = v1
            switch fi30 {
            case 9223372036854775807:
                delete(v3, "a")
            case 42:
                var v32 uint8 = 12
                _ = v32
            }
            _ = v31
        } else {
            v33 := []int{(fi30 | (fi30 << ((65535) & 63))), 2147483647, fi30}
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
            fmt.Println(v16, v1)
            _ = v33
        }
        _ = fi30
    }
    var v34 float32 = 1.5
    v1 -= float64((v34 / v34))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v34
}
