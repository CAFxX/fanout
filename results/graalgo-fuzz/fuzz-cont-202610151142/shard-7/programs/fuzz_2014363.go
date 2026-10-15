package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint8 }
func genFn0() int64 {
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{fi6, (-(fi6))}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 int16 = -364
            switch fi9 {
            case 4294967295:
                v10 += v10
            case -222:
                var v11 int32 = 1
                _ = v11
            case 1:
                fmt.Println(fi9, fi9)
            }
            _ = fi9
            _ = v10
        }
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := GenS1{F0: 0, F1: 573}
            _ = fi12
            _ = v13
        }
        if (2147483648 == 2) {
            fmt.Println(fi8)
        } else {
            var v14 int32 = -2147483648
            _ = v14
        }
        _ = fi8
    }
    var v15 float64 = 1e10
    fmt.Println(v15)
    var v16 uint64 = 2
    _ = v15
    _ = v16
    return 0
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    v3 := []int{42, 9223372036854775807, -526, 492}
    v4 := make(map[uint32]string)
    v5 := []int{4294967295}
    var v17 uint32 = 4294967295
    var v18 bool = (!true)
    var v19 uint8 = 255
    if true {
        v20 := GenS1{F0: uint64(v17), F1: -349}
        switch -118 {
        case -392:
            v18 = (!(!v18))
        case 42:
            switch 9223372036854775807 {
            case 1:
                var v21 float32 = (1.5 - 1e10)
                _ = v21
            }
        }
        _ = v20
    } else {
        var v22 int16 = 1000
        var v23 uint64 = 42
        for fi24 := 0; fi24 < 5; fi24++ {
            if v18 {
                var v25 float32 = ((1e10 + ((0.0 / 1.5) * (float32(0.0) + 0.1))) / (-2.25 * 0.1))
                var v26 int16 = v22
                v3 = append(v3, 657)
                _ = v25
                _ = v26
            }
            v27 := GenS1{F0: (^(v23)), F1: genFn0()}
            _ = fi24
            _ = v27
        }
        _ = v22
        _ = v23
    }
    fmt.Println(v18, v17)
    v17 -= v17
    var v28 uint32 = (v17 * (v17 << ((42) & 63)))
    v18 = (v18 && ((!(uint32(v19) <= v17)) && (1 <= 667)))
    v1[0] = v18
    var v29 float64 = 1.5
    for fi30 := 0; fi30 < 9; fi30++ {
        if ((v19 - v19) <= v19) {
            var v31 int16 = -1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v32 bool = true
            _ = v31
            _ = v32
        } else {
            var v33 float64 = (v29 - 3.14159)
            _ = v33
        }
        v34 := GenS2{F0: 255}
        _ = fi30
        _ = v34
    }
    var v35 uint = 1
    var v36 uint16 = 635
    for fi37 := 0; fi37 < 11; fi37++ {
        var v38 float64 = (float64((((1.5 + 0.1) + ((-2.25 - 0.1) / 1.5)) / 1e10)) + 3.14159)
        if len(v5) > 0 {
            v5[0] = (fi37 + fi37)
        }
        v29 += float64(1e10)
        _ = fi37
        _ = v38
    }
    v19 = v19
    v35 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
