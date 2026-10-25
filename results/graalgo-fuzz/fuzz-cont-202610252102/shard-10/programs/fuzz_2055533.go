package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 string; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint16 }
func genFn0() uint32 {
    var v6 int8 = -70
    v7 := GenS1{F0: (v6 * (v6 ^ 127))}
    _ = v6
    _ = v7
    return 877
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "a"
    v3 := make(map[int]uint32)
    var v4 uint32 = 862
    var v5 float64 = 1e10
    var v8 int64 = -1
    fmt.Println(v2, v4, v8)
    v4 -= v4
    if true {
        v9 := GenS1{F0: 1}
        switch 127 {
        case -2147483648:
            switch 4294967295 {
            case -1:
                fmt.Println(v8, v4, v5)
            case 851:
                v5 = v5
            case 2147483647:
                v5 = v5
            default:
                if ((1.5 <= ((0.1 / 0.1) + 1e10)) && (v4 != v4)) {
                    v10 := GenS2{F0: v4, F1: v8}
                    for fi11 := 0; fi11 < 11; fi11++ {
                        v12 := []int{}
                        _ = fi11
                        _ = v12
                    }
                    v13 := GenS2{F0: 79, F1: v8}
                    _ = v10
                    _ = v13
                } else {
                    v8 = 42
                    if (true || (!true)) {
                        v14 := GenS1{F0: 127}
                        v15 := GenS1{F0: 127}
                        _ = v14
                        _ = v15
                    }
                    switch 9223372036854775807 {
                    case -2147483648:
                        var v16 int16 = 42
                        _ = v16
                    case -1:
                        switch 2147483647 {
                        case -348:
                            v17 := GenS3{F0: -1, F1: -1}
                            _ = v17
                        case -1:
                            v5 = (3.14159 - float64(1.5))
                        }
                    case 42:
                        var v18 float64 = v5
                        _ = v18
                    default:
                        v2 = v2
                    }
                }
            }
        }
        for fi19 := 0; fi19 < 6; fi19++ {
            v20 := []int{}
            _ = fi19
            _ = v20
        }
        _ = v9
    }
    var v21 uint16 = 1
    var v22 int16 = 1000
    for fi23 := 0; fi23 < 11; fi23++ {
        if true {
            v21 += 42
            for fi24 := 0; fi24 < 2; fi24++ {
                fmt.Println(v21, v4, v4)
                var v25 uint = 42
                v2 = ((((v2 + (v2 + ("" + (v2 + v2)))) + v2) + "hello") + v2)
                _ = fi24
                _ = v25
            }
            var v26 int = fi23
            _ = v26
        }
        var v27 uint = 351
        for fi28 := 0; fi28 < 3; fi28++ {
            _ = len(v3)
            var v29 uint = (v27 << ((v27) & 63))
            _ = fi28
            _ = v29
        }
        _ = fi23
        _ = v27
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        if false {
            v31 := GenS4{F0: "", F1: 2147483647, F2: v21}
            var v32 float32 = 1e10
            v33 := GenS2{F0: v4, F1: v8}
            _ = v31
            _ = v32
            _ = v33
        }
        if ((v2 + "a") > v2) {
            var v34 float32 = 1.5
            _ = v34
        }
        switch int(v21) {
        case 127:
            if (!true) {
                _ = len(v3)
                switch (fi30 * fi30) {
                case 128:
                    v21 += v21
                case -279:
                    v35 := GenS5{F0: v21}
                    _ = v35
                default:
                    v36 := GenS2{F0: (v4 | genFn0()), F1: v8}
                    _ = v36
                }
                var v37 bool = (0 > 699)
                _ = v37
            }
        case 128:
            fmt.Println(fi30, v4, fi30)
        case 413:
            delete(v3, fi30)
        default:
            v22 = v22
        }
        _ = fi30
    }
    var v38 int = -33
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v8
}
