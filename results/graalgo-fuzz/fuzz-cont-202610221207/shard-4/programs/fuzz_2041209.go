package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint64 }
func genFn0(p0 uint32, p1 int) int64 {
    v5 := GenS1{F0: -2147483648}
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            p0 = 2147483648
            p1 = -199
            _ = fi7
        }
        if true {
            v8 := []int{127, fi6}
            p0 = p0
            v9 := GenS1{F0: int32(fi6)}
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 int32 = -661
                _ = fi10
                _ = v11
            }
        }
        v12 := []int{(fi6 - 4294967295), (-(p1)), p1, 9223372036854775807}
        _ = fi6
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v5
    return -9223372036854775808
}

func main() {
    var v1 uint = 255
    var v2 float32 = 1e10
    v3 := []uint32{1, 2}
    var v4 uint8 = 128
    v13 := GenS1{F0: 2147483647}
    fmt.Println(v2, v2, v2)
    v2 += v2
    for fi14 := 0; fi14 < 12; fi14++ {
        switch (fi14 ^ (^(fi14))) {
        case 1:
            if (!false) {
                v1 = (v1 * (^(v1)))
            } else {
                var v15 int32 = 1
                _ = v15
            }
        case 127:
            switch fi14 {
            case 213:
                v4 -= v4
            case 32767:
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if ((uint64(v4) <= 1) || false) {
                        for fi16 := 0; fi16 < 2; fi16++ {
                            var v17 int = fi14
                            if (42 == -381) {
                                v4 = v4
                                var v18 uint16 = 32767
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v18
                            }
                            _ = fi16
                            _ = v17
                        }
                        var v19 int16 = -913
                        _ = v19
                    }
                    v20 := GenS2{F0: 98, F1: (v4 << ((v1) & 63)), F2: 4294967295}
                    _ = v20
                }
            case 4294967295:
                fmt.Println(v1, v4, fi14)
            }
        }
        var v21 uint8 = v4
        _ = fi14
        _ = v21
    }
    v4 -= 255
    var v22 string = "a"
    for fi23 := 0; fi23 < 8; fi23++ {
        var v24 int32 = -2147483648
        _ = fi23
        _ = v24
    }
    fmt.Println(v22)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
