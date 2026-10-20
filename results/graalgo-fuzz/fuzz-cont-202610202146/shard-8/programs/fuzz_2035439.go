package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }

func main() {
    var v1 int32 = 42
    var v2 int16 = -83
    var v3 uint = 0
    var v4 uint64 = 450
    var v5 float64 = 3.14159
    var v6 uint16 = 878
    var v7 uint64 = 42
    v8 := GenS1{F0: 42, F1: false}
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 uint8 = 1
        _ = fi9
        _ = v10
    }
    if (((!true) || (1 <= 42)) && (0 < uint32(v3))) {
        var v11 uint64 = 506
        fmt.Println(v3, v7, v7)
        _ = v11
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := GenS1{F0: -1, F1: (2 >= (^(v6)))}
            for fi14 := 0; fi14 < 18; fi14++ {
                v2 += 1000
                _ = fi14
            }
            _ = fi12
            _ = v13
        }
        var v15 string = (("\n\t" + "a") + "hello")
        _ = v15
    }
    v3 += uint(v6)
    v16 := []int{}
    for fi17 := 0; fi17 < 9; fi17++ {
        if true {
            for fi18 := 0; fi18 < 3; fi18++ {
                var v19 bool = false
                v1 += (v1 - 1)
                var v20 string = "a"
                _ = fi18
                _ = v19
                _ = v20
            }
            for fi21 := 0; fi21 < 12; fi21++ {
                v22 := GenS1{F0: -128, F1: true}
                _ = fi21
                _ = v22
            }
            for fi23 := 0; fi23 < 19; fi23++ {
                v1 -= (-(v1))
                switch fi23 {
                case -2147483648:
                    if len(v16) > 0 {
                        v16[0] = (fi23 - (fi17 >> (((v3 & (v3 - (v3 | (v3 & v3))))) & 63)))
                    }
                    var v24 int = 32767
                    v1 -= int32(v7)
                    _ = v24
                case 993:
                    var v25 int8 = 100
                    _ = v25
                }
                _ = fi23
            }
        } else {
            v3 += v3
            var v26 bool = false
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            _ = v26
        }
        var v27 uint8 = 196
        if false {
            switch fi17 {
            case 127:
                if (("a" >= "\n\t") && (v5 > ((float64(1.5) * v5) * v5))) {
                    v28 := GenS1{F0: 101, F1: (65535 > uint32(v4))}
                    if ((v2 - v2) < 221) {
                        var v29 uint64 = v7
                        var v30 float32 = 0.0
                        _ = v29
                        _ = v30
                    } else {
                        var v31 uint16 = v6
                        var v32 uint = (v3 & v3)
                        var v33 int = -2147483648
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    _ = v28
                } else {
                    _ = len(v16)
                }
            case 1:
                var v34 int32 = 1
                _ = v34
            }
            for fi35 := 0; fi35 < 3; fi35++ {
                _ = len(v16)
                v16 = append(v16, fi17)
                _ = fi35
            }
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
        _ = fi17
        _ = v27
    }
    fmt.Println(v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
