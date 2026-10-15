package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string }

func main() {
    var v1 float64 = float64(3.14159)
    var v2 int64 = 1
    var v3 int16 = 1000
    var v4 string = "\n\t"
    var v5 int16 = 1000
    var v6 int32 = -2147483648
    var v7 uint32 = 929
    var v8 float32 = ((0.1 + 1e10) * (1.5 / 3.14159))
    var v9 string = v4
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS1{F0: (v7 + v7), F1: (v7 >= (^(v7))), F2: (v9 + (v9 + v9))}
        if (!true) {
            for fi12 := 0; fi12 < 13; fi12++ {
                v2 -= (v2 ^ (v2 - (v2 % ((v2) | 1))))
                _ = fi12
            }
            v5 -= (v5 + (v5 / (((^(v5))) | 1)))
        }
        _ = fi10
        _ = v11
    }
    var v13 int = -402
    for fi14 := 0; fi14 < 2; fi14++ {
        v1 = v1
        fmt.Println(v9)
        v15 := GenS1{F0: 1, F1: (-122 <= 38), F2: v4}
        _ = fi14
        _ = v15
    }
    if (2 > 0) {
        fmt.Println(v3)
    } else {
        var v16 uint = 0
        _ = v16
    }
    var v17 uint8 = 1
    if (float64((v8 * v8)) < -2.25) {
        v18 := GenS2{F0: v7}
        switch v4 {
        case "hello":
            switch (-(v13)) {
            case 128:
                switch "" {
                case "hello":
                    for fi19 := 0; fi19 < 1; fi19++ {
                        switch fi19 {
                        case 42:
                            v6 += 42
                        case 32768:
                            v8 += v8
                        }
                        v5 = (^(v3))
                        if true {
                            var v20 int16 = (v3 << ((0) & 63))
                            if (!(((v17 - v17) != (v17 << ((805) & 63))) || false)) {
                                if (!true) {
                                    fmt.Println(fi19)
                                    if ((100 < -1) || true) {
                                        v20 = (v5 ^ (v5 & (v3 | (v3 >> ((932) & 63)))))
                                        for fi21 := 0; fi21 < 19; fi21++ {
                                            var v22 uint = uint(v7)
                                            var v23 int64 = int64(v22)
                                            _ = fi21
                                            _ = v22
                                            _ = v23
                                        }
                                    }
                                    v24 := GenS3{F0: int8(v17)}
                                    _ = v24
                                }
                            } else {
                                for fi25 := 0; fi25 < 16; fi25++ {
                                    if false {
                                        v17 = (v17 & (v17 >> ((4294967295) & 63)))
                                        var v26 float32 = 0.0
                                        _ = v26
                                    } else {
                                        v27 := []int{v13, v13, (fi19 ^ (v13 & fi25)), (fi19 << ((0) & 63))}
                                        v28 := GenS3{F0: -128}
                                        var v29 uint64 = 459
                                        v7 -= v7
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                    }
                                    if false {
                                        fmt.Println(v13, v17)
                                        v30 := GenS3{F0: 127}
                                        v17 += 128
                                        _ = v30
                                    }
                                    v31 := []int{(-(v13)), fi19}
                                    _ = fi25
                                    _ = v31
                                }
                                v5 += 32767
                            }
                            var v32 string = (v9 + v4)
                            _ = v20
                            _ = v32
                        }
                        _ = fi19
                    }
                case "":
                    var v33 float32 = float32(1e10)
                    _ = v33
                }
            case 4294967295:
                v7 += v7
            case 127:
                v6 -= int32(v13)
            }
        case "a":
            if ((((v1 / 1e10) / v1) > v1) && true) {
                fmt.Println(v9, v4, v9)
                v17 += 215
            }
        case "":
            v34 := GenS4{F0: v9}
            _ = v34
        }
        _ = v18
    } else {
        var v35 uint = 255
        _ = v35
    }
    if (!true) {
        var v36 int16 = -844
        v37 := []int{v13, (v13 << ((42) & 63))}
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
