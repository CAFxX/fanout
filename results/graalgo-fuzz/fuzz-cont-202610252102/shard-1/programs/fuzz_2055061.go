package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint }
type GenS2 struct { F0 uint32; F1 uint }
func genFn0(p0 int) uint64 {
    var v7 int64 = 42
    var v8 int16 = -281
    p0 = p0
    v9 := []int{p0, 0, 971}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1() int {
    var v10 uint = 255
    var v11 uint64 = 2
    _ = v10
    _ = v11
    return 9223372036854775807
}

func main() {
    var v1 float64 = -2.25
    var v2 int = 1
    var v3 uint64 = 0
    var v4 bool = (1 != 349)
    var v5 float32 = -2.25
    var v6 uint = 255
    v12 := GenS1{F0: (v6 & v6), F1: v6, F2: v6}
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int8 = -128
        fmt.Println(fi13, v6, v14)
        v15 := []int{fi13, v2, (v2 / (((fi13 | v2)) | 1)), v2}
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 float64 = v1
        v18 := GenS1{F0: (^(v6)), F1: 42, F2: (v6 - 541)}
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 int32 = 2147483647
    v4 = (!(v6 >= 2))
    v2 -= v2
    for fi20 := 0; fi20 < 10; fi20++ {
        v3 = v3
        v4 = (!((^(v3)) > 2))
        _ = fi20
    }
    var v21 uint32 = 2147483648
    v19 = (v19 % ((v19) | 1))
    if (((-128 <= 47) && false) && (-1 == (v19 & (-(v19))))) {
        if false {
            v22 := GenS1{F0: 391, F1: v6, F2: v6}
            _ = v22
        }
    } else {
        for fi23 := 0; fi23 < 15; fi23++ {
            v24 := GenS1{F0: (v6 * v6), F1: (v6 >> (((^(v6))) & 63)), F2: 1}
            if (!false) {
                if true {
                    fmt.Println(v3)
                    switch genFn1() {
                    case -71:
                        if (117 < 67) {
                            for fi25 := 0; fi25 < 17; fi25++ {
                                fmt.Println(v19)
                                fmt.Println(v19, v6)
                                switch (^(fi23)) {
                                case 281:
                                    var v26 string = ("\n\t" + ("foo bar" + "foo bar"))
                                    _ = v26
                                case 42:
                                    fmt.Println(v4, v3, v2)
                                }
                                _ = fi25
                            }
                            v27 := GenS2{F0: (v21 ^ (v21 & v21)), F1: v6}
                            _ = v27
                        }
                    case 127:
                        v4 = ((!(v3 != v3)) || false)
                    }
                }
                v28 := GenS1{F0: v6, F1: uint(v21), F2: v6}
                v29 := []int{fi23, (-(v2)), (^(v2)), int(v6)}
                _ = v28
                _ = v29
            } else {
                var v30 uint64 = genFn0(fi23)
                _ = v30
            }
            _ = fi23
            _ = v24
        }
        v31 := []int{v2}
        _ = v31
    }
    fmt.Println(v3, v5, v1)
    v3 -= uint64(v6)
    fmt.Println(v3, v6, v2)
    v4 = true
    for fi32 := 0; fi32 < 3; fi32++ {
        v6 = uint(v21)
        v1 = v1
        _ = fi32
    }
    if ((v21 << (((^(v6))) & 63)) >= (v21 % ((234) | 1))) {
        v33 := []int{}
        v34 := GenS1{F0: uint(v21), F1: 0, F2: v6}
        if v4 {
            var v35 uint32 = (v21 / (((v21 - uint32(v3))) | 1))
            _ = v35
        } else {
            var v36 int64 = -1
            _ = v36
        }
        _ = v33
        _ = v34
    } else {
        var v37 uint = (v6 | (v6 + 65535))
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
