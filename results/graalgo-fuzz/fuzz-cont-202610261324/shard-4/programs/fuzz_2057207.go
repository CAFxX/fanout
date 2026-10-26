package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 bool; F2 string }
type GenS4 struct { F0 int8; F1 uint8; F2 int8 }
func genFn0(p0 uint16, p1 uint16, p2 int16) string {
    var v5 float32 = (((1e10 * float32((-2.25 - 1e10))) - (3.14159 * 0.1)) * 1e10)
    v6 := GenS1{F0: 2, F1: true, F2: 64}
    switch 4294967295 {
    case 761:
        var v7 int = 4294967295
        _ = v7
    case 0:
        fmt.Println(v5, v5)
    case -774:
        if (0 >= 42) {
            for fi8 := 0; fi8 < 6; fi8++ {
                if true {
                    for fi9 := 0; fi9 < 20; fi9++ {
                        v10 := []int{}
                        if (false && (!false)) {
                            fmt.Println(p1)
                        }
                        var v11 uint64 = 2
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    var v12 uint64 = 4294967295
                    v13 := []int{}
                    _ = v12
                    _ = v13
                } else {
                    var v14 uint16 = p1
                    var v15 int64 = -1
                    if true {
                        v16 := []int{int(p1)}
                        _ = v16
                    } else {
                        v17 := GenS1{F0: 2, F1: (258 == 978), F2: 127}
                        _ = v17
                    }
                    _ = v14
                    _ = v15
                }
                if (true || (true || (0 > 2147483647))) {
                    if (false || true) {
                        if ((v5 + ((((((1e10 / v5) * 1e10) - (v5 - ((v5 + 0.0) / (v5 - float32(0.1))))) - v5) / (v5 + 1.5)) + float32(0.0))) > (1e10 + (((v5 * v5) + v5) / v5))) {
                            v18 := GenS1{F0: 2, F1: true, F2: -94}
                            _ = v18
                        }
                        var v19 bool = ((!(true && (!true))) || (1 <= 2))
                        p0 = 0
                        _ = v19
                    } else {
                        var v20 uint8 = 127
                        _ = v20
                    }
                    fmt.Println(fi8, p2, v5)
                    v21 := GenS1{F0: 18446744073709551615, F1: ((!true) || (true || (p2 != p2))), F2: 127}
                    _ = v21
                }
                _ = fi8
            }
        }
    default:
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := GenS1{F0: 42, F1: true, F2: 0}
            var v24 int16 = (-(p2))
            switch -225 {
            case 4294967295:
                var v25 float32 = (v5 * v5)
                _ = v25
            case 151:
                v26 := GenS1{F0: 3, F1: true, F2: 42}
                _ = v26
            }
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (("hello" + "hello") + "\n\t")
}
func genFn1() int32 {
    for fi27 := 0; fi27 < 6; fi27++ {
        v28 := GenS2{F0: 0}
        var v29 uint64 = 42
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 uint16 = 255
    var v31 uint32 = 1
    _ = v30
    _ = v31
    return 2147483647
}

func main() {
    var v1 bool = (1 >= 401)
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 42
    var v4 float64 = 1.5
    fmt.Println(v1)
    var v32 int16 = 0
    fmt.Println(v1, v4)
    if v1 {
        for fi33 := 0; fi33 < 7; fi33++ {
            fmt.Println(v1)
            var v34 int32 = 42
            _ = fi33
            _ = v34
        }
        v3 = 42
        v35 := GenS3{F0: 871, F1: (753 <= 149), F2: ("" + "\n\t")}
        _ = v35
    }
    v3 += 42
    v36 := GenS1{F0: 18446744073709551615, F1: (!v1), F2: -1}
    v37 := []int{int(v32), -410}
    v2[(v3 | 0)] = (^(v3))
    for fi38 := 0; fi38 < 1; fi38++ {
        var v39 string = ""
        _ = fi38
        _ = v39
    }
    for fi40 := 0; fi40 < 1; fi40++ {
        var v41 uint = 670
        for fi42 := 0; fi42 < 1; fi42++ {
            fmt.Println(v1, v32, v1)
            var v43 uint = (v41 / (((v41 - (^(v41)))) | 1))
            _ = fi42
            _ = v43
        }
        _ = fi40
        _ = v41
    }
    fmt.Println(v4)
    v44 := GenS4{F0: -1, F1: 0, F2: int8(v3)}
    if v1 {
        switch 127 {
        case 380:
            var v45 uint8 = 140
            _ = v45
        }
        v4 += (v4 - -2.25)
        for fi46 := 0; fi46 < 19; fi46++ {
            var v47 int32 = 0
            _ = fi46
            _ = v47
        }
    } else {
        for fi48 := 0; fi48 < 19; fi48++ {
            var v49 int16 = -980
            _ = fi48
            _ = v49
        }
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v37
    _ = v4
    _ = v44
}
