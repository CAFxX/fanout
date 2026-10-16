package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 uint }
func genFn0() uint {
    v4 := GenS1{F0: 18446744073709551615, F1: -910, F2: 1}
    var v5 uint32 = 454
    v6 := []int{32768, -9223372036854775808}
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 42
    var v3 int64 = 1
    for fi7 := 0; fi7 < 1; fi7++ {
        v3 -= 1
        _ = fi7
    }
    var v8 uint32 = 481
    v9 := GenS1{F0: genFn0(), F1: -309, F2: 42}
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            v8 = v8
            fmt.Println(v8, fi11, v8)
            var v12 float64 = 1.5
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    v13 := GenS1{F0: 878, F1: 1000, F2: 4294967295}
    for fi14 := 0; fi14 < 19; fi14++ {
        switch (-(fi14)) {
        case 2147483647:
            for fi15 := 0; fi15 < 10; fi15++ {
                v2 = (v2 << ((genFn0()) & 63))
                var v16 int32 = -386
                v17 := GenS2{F0: (v16 / ((1) | 1)), F1: (v8 >> ((255) & 63))}
                _ = fi15
                _ = v16
                _ = v17
            }
        case -932:
            v3 = v3
        case -881:
            if (!false) {
                var v18 string = (("a" + ("a" + "")) + "")
                fmt.Println(fi14, v3, v8)
                for fi19 := 0; fi19 < 5; fi19++ {
                    var v20 bool = (!false)
                    var v21 int = (fi14 - (fi14 * 4294967295))
                    v20 = true
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = v18
            } else {
                var v22 uint = 2
                if false {
                    v23 := []int{(fi14 | -2147483648)}
                    if (1e10 >= (3.14159 + ((0.0 - float64(float32(float64((3.14159 + 3.14159))))) / ((1.5 / (1e10 * ((-2.25 - (((0.0 - ((3.14159 - (3.14159 / ((0.1 / 1.5) - 3.14159))) * float64((1e10 - ((-2.25 / 3.14159) - -2.25))))) * 1.5) + -2.25)) / 1e10))) * (((((float64((1.5 + 0.0)) - 0.0) - (1e10 + float64(-2.25))) + 1.5) * -2.25) + (-2.25 / 0.1)))))) {
                        v8 = v8
                        v24 := GenS2{F0: int32(fi14), F1: v8}
                        for fi25 := 0; fi25 < 6; fi25++ {
                            fmt.Println(v3, fi14)
                            if false {
                                v3 = v3
                                var v26 uint = genFn0()
                                _ = v26
                            } else {
                                var v27 uint32 = v8
                                _ = v27
                            }
                            var v28 int = fi25
                            var v29 int16 = 32767
                            var v30 uint8 = 127
                            fmt.Println(fi14, fi25, fi14)
                            _ = len(v23)
                            var v31 float32 = float32(0.1)
                            v31 -= v31
                            _ = fi25
                            _ = v28
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        _ = v24
                    } else {
                        v32 := GenS3{F0: true}
                        switch "\n\t" {
                        case "\n\t":
                            _ = len(v1)
                        case "foo bar":
                            v3 += v3
                        }
                        _ = v32
                    }
                    var v33 int8 = 42
                    _ = v23
                    _ = v33
                }
                v34 := GenS2{F0: -733, F1: 42}
                _ = v22
                _ = v34
            }
        default:
            fmt.Println(fi14)
        }
        v35 := GenS3{F0: (42 >= -840)}
        delete(v1, v8)
        _ = fi14
        _ = v35
    }
    v36 := v1[(v8 & v8)]
    switch 2147483647 {
    case -1:
        var v37 uint64 = 0
        _ = v37
    case 32768:
        fmt.Println(v8)
    }
    fmt.Println(v36, v3, v8)
    v3 = (v3 ^ v3)
    var v38 uint = genFn0()
    v39 := GenS4{F0: 65535, F1: (^(v38))}
    v3 += (v3 + v3)
    v3 = (^(v3))
    v40 := []int{}
    var v41 int16 = -742
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v39
    _ = v40
    _ = v41
    _ = v8
    _ = v9
}
