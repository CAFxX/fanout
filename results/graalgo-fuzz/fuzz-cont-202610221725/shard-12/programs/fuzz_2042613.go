package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int64) uint8 {
    if true {
        v5 := []int{9223372036854775807, 128, 128, 4294967295}
        _ = v5
    } else {
        v6 := []int{-150, 986}
        _ = v6
    }
    var v7 uint32 = 65535
    p0 -= 0
    _ = p0
    _ = v7
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]int)
    switch (("\n\t" + "hello") + "a") {
    case "\n\t":
        if (!(true || (771 > 42))) {
            if true {
                v8 := GenS1{F0: 1, F1: -602, F2: 127}
                var v9 bool = (939 == 42)
                for fi10 := 0; fi10 < 12; fi10++ {
                    if true {
                        if (709 > 365) {
                            for fi11 := 0; fi11 < 6; fi11++ {
                                v9 = (((42 >= 2) && (1 != 127)) && (v9 || v9))
                                _ = fi11
                            }
                            if v9 {
                                if false {
                                    v9 = v9
                                    var v12 uint32 = 2147483648
                                    _ = v12
                                }
                                fmt.Println(fi10, fi10)
                            }
                        }
                    } else {
                        v13 := GenS1{F0: -9223372036854775808, F1: 42, F2: 100}
                        for fi14 := 0; fi14 < 6; fi14++ {
                            v9 = v9
                            v15 := GenS2{F0: 710}
                            var v16 float64 = 1e10
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        _ = v13
                    }
                    var v17 float64 = 1.5
                    _ = fi10
                    _ = v17
                }
                _ = v8
                _ = v9
            } else {
                v18 := GenS3{F0: 1}
                _ = v18
            }
        } else {
            var v19 int32 = -1
            var v20 uint16 = 873
            _ = v19
            _ = v20
        }
    case "foo bar":
        var v21 int8 = 0
        _ = v21
    case "a":
        if false {
            var v22 uint32 = 42
            for fi23 := 0; fi23 < 20; fi23++ {
                fmt.Println(v22, v22)
                fmt.Println(v22, v22)
                _ = fi23
            }
            fmt.Println(v22)
            _ = v22
        } else {
            if ((1.5 + float64(0.1)) > float64(0.1)) {
                var v24 int64 = -9223372036854775808
                v25 := GenS1{F0: 1, F1: (^(v24)), F2: 45}
                v24 += (v24 / ((0) | 1))
                _ = v24
                _ = v25
            } else {
                var v26 int32 = 0
                var v27 string = ("foo bar" + "hello")
                _ = v26
                _ = v27
            }
            var v28 uint32 = 65535
            _ = v28
        }
    }
    var v29 int8 = 28
    v30 := GenS2{F0: -1}
    var v31 float64 = float64(0.1)
    _ = len(v3)
    fmt.Println(v31, v31)
    var v32 uint32 = 42
    v32 -= v32
    v4[v32] = -1
    for fi33 := 0; fi33 < 17; fi33++ {
        v34 := GenS4{F0: 1, F1: 42}
        switch "foo bar" {
        case "\n\t":
            var v35 int = (fi33 | 32767)
            _ = v35
        case "":
            v29 -= (v29 << ((31) & 63))
        }
        if ((!false) || true) {
            var v36 int64 = -842
            if true {
                v36 = -9223372036854775808
            } else {
                var v37 int16 = 1000
                _ = v37
            }
            v38 := GenS1{F0: 61, F1: 42, F2: (v29 + v29)}
            _ = v36
            _ = v38
        } else {
            var v39 uint32 = (v32 >> ((65535) & 63))
            _ = v39
        }
        _ = fi33
        _ = v34
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        var v41 uint64 = 75
        var v42 uint16 = 320
        _ = fi40
        _ = v41
        _ = v42
    }
    switch 1 {
    case 2147483647:
        delete(v2, v32)
    default:
        v43 := []int{4294967295, 42}
        _ = v43
    }
    for fi44 := 0; fi44 < 15; fi44++ {
        for fi45 := 0; fi45 < 10; fi45++ {
            delete(v3, (fi44 % ((fi45) | 1)))
            _ = fi45
        }
        v46 := GenS5{F0: 42}
        var v47 string = "a"
        _ = fi44
        _ = v46
        _ = v47
    }
    v31 -= v31
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
