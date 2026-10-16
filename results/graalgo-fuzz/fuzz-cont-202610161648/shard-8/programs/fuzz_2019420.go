package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 uint; F2 uint16 }

func main() {
    var v1 float32 = (0.1 / 1.0)
    v2 := make(map[uint32]int)
    var v3 bool = false
    if v3 {
        var v4 float32 = (1.5 / float32(1.5))
        if true {
            v5 := GenS1{F0: 42, F1: 113}
            for fi6 := 0; fi6 < 15; fi6++ {
                var v7 uint64 = 313
                v8 := []int{}
                v1 = float32(1e10)
                _ = fi6
                _ = v7
                _ = v8
            }
            fmt.Println(v1)
            _ = v5
        }
        _ = v4
    }
    if (-17 > -128) {
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 uint32 = 416
            _ = fi9
            _ = v10
        }
        if (127 != 0) {
            v11 := []int{-486, 887}
            _ = v11
        } else {
            for fi12 := 0; fi12 < 1; fi12++ {
                switch (fi12 & 1) {
                case 32767:
                    v3 = ((((((float64(1.5) * -2.25) + (((3.14159 / 1e10) / (1e10 - ((0.0 * -2.25) - float64((0.0 * v1))))) / (float64((v1 * float32((0.0 - 0.0)))) + 0.1))) * (float64(1.5) * float64(v1))) + 1e10) + 3.14159) != 3.14159)
                case -2147483648:
                    for fi13 := 0; fi13 < 8; fi13++ {
                        _ = len(v2)
                        _ = fi13
                    }
                case 32768:
                    var v14 int32 = 1
                    _ = v14
                }
                v15 := v2[365]
                _ = fi12
                _ = v15
            }
        }
        v1 += float32((float64(v1) * 3.14159))
    }
    if v3 {
        v16 := GenS2{F0: 128}
        for fi17 := 0; fi17 < 1; fi17++ {
            v2[920] = fi17
            v18 := GenS2{F0: 2}
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 7; fi19++ {
            for fi20 := 0; fi20 < 8; fi20++ {
                v1 = v1
                v3 = (1.5 >= v1)
                fmt.Println(fi20)
                _ = fi20
            }
            v3 = false
            _ = fi19
        }
        _ = v16
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        v1 -= v1
        _ = fi21
    }
    v22 := GenS3{F0: 18446744073709551615, F1: 4294967295, F2: 2}
    var v23 int16 = 486
    var v24 int = 42
    _ = len(v2)
    v3 = v3
    v25 := []int{v24, (v24 + v24), (-(v24)), (v24 * (v24 + (-(v24))))}
    var v26 uint16 = 884
    var v27 string = ("\n\t" + "hello")
    v23 += (v23 << ((42) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
