package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint16 }
func genFn0() bool {
    var v7 int8 = 127
    v8 := GenS1{F0: 1, F1: 0, F2: 42}
    var v9 float32 = ((0.1 - 1e10) + 1.5)
    v10 := GenS1{F0: 850, F1: 199, F2: 18446744073709551615}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return true
}
func genFn1() uint {
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := []int{}
        var v13 float64 = (1e10 - 3.14159)
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint = 65535
    v15 := GenS2{F0: 1, F1: -128}
    _ = v14
    _ = v15
    return 572
}

func main() {
    v1 := []uint32{0, 4294967295, 0, 760}
    v2 := make(map[string]int)
    var v3 int = -456
    var v4 string = ("\n\t" + "a")
    var v5 float32 = float32(0.1)
    var v6 uint8 = 2
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 int8 = -128
        for fi18 := 0; fi18 < 10; fi18++ {
            _ = len(v1)
            _ = fi18
        }
        _ = fi16
        _ = v17
    }
    var v19 float32 = 0.1
    var v20 float32 = v19
    for fi21 := 0; fi21 < 11; fi21++ {
        v5 -= v5
        v6 -= v6
        for fi22 := 0; fi22 < 20; fi22++ {
            switch fi22 {
            case -9223372036854775808:
                v23 := GenS3{F0: 115}
                _ = v23
            default:
                var v24 float64 = (float64(float32(1.5)) / 0.1)
                _ = v24
            }
            v5 = (v19 - -2.25)
            switch v4 {
            case "\n\t":
                if true {
                    var v25 uint32 = 42
                    _ = v25
                } else {
                    var v26 int8 = 1
                    _ = v26
                }
            default:
                v20 = -2.25
            }
            _ = fi22
        }
        _ = fi21
    }
    fmt.Println(v19, v6, v6)
    v1 = append(v1, 42)
    var v27 uint16 = 32767
    for fi28 := 0; fi28 < 1; fi28++ {
        if true {
            v19 = v5
            var v29 uint = 4294967295
            for fi30 := 0; fi30 < 16; fi30++ {
                v27 = v27
                v2["a"] = fi30
                v20 -= float32(1e10)
                _ = fi30
            }
            _ = v29
        }
        _ = fi28
    }
    v3 = -2147483648
    v31 := v2["a"]
    _ = len(v2)
    if false {
        v6 += v6
        if (!(-32768 < 497)) {
            v3 += v3
            if true {
                v27 = v27
                v19 = v20
            } else {
                var v32 int64 = -9223372036854775808
                switch v3 {
                case 42:
                    var v33 int16 = 1000
                    _ = v33
                case -524:
                    var v34 int16 = 1
                    _ = v34
                case -9223372036854775808:
                    v20 = v5
                }
                _ = v32
            }
            var v35 uint = 255
            _ = v35
        }
        v36 := GenS2{F0: 2, F1: 1}
        _ = v36
    } else {
        var v37 int8 = -1
        _ = v37
    }
    v27 = v27
    v3 += v31
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
