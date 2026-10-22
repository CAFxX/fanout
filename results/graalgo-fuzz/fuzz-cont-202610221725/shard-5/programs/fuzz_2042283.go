package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint }
func genFn0(p0 uint32) int {
    p0 = p0
    p0 = (p0 ^ (p0 >> ((47) & 63)))
    var v4 int32 = 42
    _ = p0
    _ = v4
    return 32768
}

func main() {
    var v1 int = -417
    var v2 string = (((("\n\t" + "a") + "\n\t") + "a") + "a")
    v3 := make(map[uint32]int)
    fmt.Println(v1)
    if (255 != 128) {
        v1 += v1
        for fi5 := 0; fi5 < 15; fi5++ {
            for fi6 := 0; fi6 < 5; fi6++ {
                v7 := GenS1{F0: 2, F1: 4294967295, F2: 0}
                var v8 float32 = 1e10
                fmt.Println(fi6)
                _ = fi6
                _ = v7
                _ = v8
            }
            v2 = (v2 + v2)
            v1 = (^(fi5))
            _ = fi5
        }
        var v9 uint = 42
        _ = v9
    }
    v1 = v1
    var v10 int32 = 2147483647
    v11 := GenS1{F0: 0, F1: 42, F2: 1}
    var v12 uint8 = 2
    v10 += v10
    var v13 int16 = 0
    switch (v1 ^ genFn0(721)) {
    case 2147483647:
        if (1e10 > 3.14159) {
            var v14 int16 = (v13 ^ (v13 << ((1) & 63)))
            if ((false && (!(!(1 >= -469)))) && false) {
                var v15 uint32 = 731
                var v16 uint8 = v12
                _ = v15
                _ = v16
            }
            if (v13 < v14) {
                var v17 float32 = ((1e10 + 0.0) / 0.1)
                _ = v17
            } else {
                v18 := []int{(v1 + (v1 | genFn0(786))), -1}
                var v19 float32 = (((3.14159 - 3.14159) * -2.25) + 0.0)
                for fi20 := 0; fi20 < 5; fi20++ {
                    switch v2 {
                    case "foo bar":
                        v21 := GenS1{F0: 42, F1: 652, F2: 42}
                        _ = v21
                    case "a":
                        for fi22 := 0; fi22 < 15; fi22++ {
                            var v23 int8 = 42
                            var v24 int = v1
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                    case "\n\t":
                        var v25 int8 = 100
                        _ = v25
                    }
                    fmt.Println(v2, fi20)
                    v26 := GenS1{F0: 301, F1: 2, F2: 482}
                    _ = fi20
                    _ = v26
                }
                _ = v18
                _ = v19
            }
            _ = v14
        }
    case -526:
        for fi27 := 0; fi27 < 6; fi27++ {
            v28 := GenS1{F0: 42, F1: 42, F2: 556}
            v12 = v12
            _ = fi27
            _ = v28
        }
    case -946:
        for fi29 := 0; fi29 < 16; fi29++ {
            for fi30 := 0; fi30 < 3; fi30++ {
                fmt.Println(v12, v1)
                switch v2 {
                case "":
                    v12 = v12
                case "\n\t":
                    var v31 uint16 = 245
                    _ = v31
                }
                var v32 int32 = int32(v13)
                _ = fi30
                _ = v32
            }
            v33 := GenS1{F0: 65535, F1: 0, F2: 255}
            _ = fi29
            _ = v33
        }
    }
    var v34 int8 = -68
    v35 := GenS2{F0: 1}
    var v36 float64 = 1e10
    v37 := GenS3{F0: 42, F1: 2}
    switch v2 {
    case "a":
        fmt.Println(v10)
    default:
        if (false && ((v36 * v36) == v36)) {
            var v38 uint8 = (v12 & v12)
            v13 -= v13
            if false {
                fmt.Println(v2, v2)
            } else {
                if (0.1 > v36) {
                    switch v1 {
                    case -2147483648:
                        var v39 float32 = (1e10 + float32(((float64(0.1) + 0.1) + v36)))
                        _ = v39
                    case 343:
                        var v40 int = v1
                        _ = v40
                    case 32767:
                        var v41 bool = false
                        _ = v41
                    }
                }
                v34 -= (v34 * v34)
                if false {
                    switch (v1 / ((v1) | 1)) {
                    case -9223372036854775808:
                        v42 := v3[76]
                        _ = v42
                    }
                    v43 := GenS2{F0: 298}
                    _ = v43
                }
            }
            _ = v38
        } else {
            delete(v3, 42)
        }
    }
    switch (v1 - 4294967295) {
    case -932:
        if ((v36 <= float64((3.14159 / 1.0))) && false) {
            var v44 uint = 4294967295
            var v45 int = v1
            _ = v44
            _ = v45
        }
    case 42:
        v10 -= -1
    default:
        fmt.Println(v36)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
