package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint16 }
func genFn0() int16 {
    var v5 uint32 = 0
    fmt.Println(v5)
    _ = v5
    return 6
}
func genFn1(p0 uint, p1 int) int16 {
    var v6 int8 = -97
    var v7 bool = true
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int16 = 1
        if (!((^(v6)) == v6)) {
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := []int{fi8, -789, (^(fi8))}
                v7 = v7
                _ = fi10
                _ = v11
            }
            v6 = v6
            if (float32((float64(3.14159) - 1.5)) != 1.5) {
                p1 = p1
                p0 = p0
            } else {
                fmt.Println(fi8, v7, p0)
            }
        } else {
            v12 := []int{p1, p1, -767}
            _ = v12
        }
        v13 := []int{p1, 1, 4294967295}
        _ = fi8
        _ = v13
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 bool = false
    v2 := make(map[int]int)
    var v3 int64 = 590
    var v4 string = ("" + "a")
    for fi14 := 0; fi14 < 16; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 int32 = 1
            v16 += v16
            v4 = (v4 + "")
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    v17 := GenS1{F0: -32768, F1: 4294967295, F2: 364}
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 string = (v4 + v4)
        var v20 float64 = 3.14159
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: -157, F1: 72, F2: 2147483648}
    var v22 int = -9223372036854775808
    fmt.Println(v1)
    if false {
        v1 = v1
        if (float64(3.14159) < 0.1) {
            if v1 {
                v1 = (!(1 >= 771))
            }
        }
        var v23 uint = 18446744073709551615
        _ = v23
    } else {
        switch ("a" + (v4 + v4)) {
        case "foo bar":
            var v24 int32 = -887
            _ = v24
        case "hello":
            switch v4 {
            case "a":
                delete(v2, 127)
            }
        default:
            switch v22 {
            case 4294967295:
                v25 := GenS2{F0: 925, F1: ((v1 || (4294967295 >= 651)) && (255 <= 0)), F2: 42}
                _ = v25
            case 434:
                v26 := GenS2{F0: 42, F1: (((v3 ^ v3) <= (^(v3))) && (61 != 2147483647)), F2: 42}
                _ = v26
            case 2147483647:
                if (!v1) {
                    switch (v22 + v22) {
                    case -321:
                        v27 := GenS2{F0: 1, F1: v1, F2: 1}
                        _ = v27
                    case 307:
                        for fi28 := 0; fi28 < 3; fi28++ {
                            v22 = fi28
                            _ = fi28
                        }
                    case -9223372036854775808:
                        if v1 {
                            for fi29 := 0; fi29 < 16; fi29++ {
                                for fi30 := 0; fi30 < 19; fi30++ {
                                    v31 := []int{}
                                    _ = fi30
                                    _ = v31
                                }
                                var v32 float64 = 0.0
                                if v1 {
                                    v4 = ""
                                    for fi33 := 0; fi33 < 11; fi33++ {
                                        var v34 uint32 = 4294967295
                                        _ = fi33
                                        _ = v34
                                    }
                                }
                                _ = fi29
                                _ = v32
                            }
                            var v35 uint64 = 42
                            var v36 string = (v4 + (v4 + v4))
                            _ = v35
                            _ = v36
                        } else {
                            fmt.Println(v22, v3)
                        }
                    default:
                        v1 = v1
                    }
                    v3 = 42
                    fmt.Println(v1, v1)
                }
            }
        }
        v37 := []int{0, v22}
        v1 = v1
        _ = v37
    }
    var v38 uint8 = 225
    if false {
        v39 := GenS1{F0: 0, F1: 761, F2: 775}
        v40 := []int{-2147483648, (^(v22)), v22}
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v38
    _ = v4
}
