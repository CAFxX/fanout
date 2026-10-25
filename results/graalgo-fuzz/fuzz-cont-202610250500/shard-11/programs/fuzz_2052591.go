package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint8 }
type GenS2 struct { F0 string; F1 uint32; F2 int64 }
type GenS3 struct { F0 bool }
func genFn0() uint8 {
    var v6 int64 = 42
    v7 := GenS1{F0: 1, F1: -191, F2: 0}
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    var v3 uint16 = 826
    var v4 float64 = -2.25
    var v5 bool = (true && (32768 >= 32768))
    if ((2147483647 >= 42) && ((1.5 + 1.5) >= 1.5)) {
        for fi8 := 0; fi8 < 6; fi8++ {
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 int32 = 803
                _ = fi9
                _ = v10
            }
            v5 = v5
            _ = fi8
        }
        if false {
            var v11 string = "\n\t"
            for fi12 := 0; fi12 < 3; fi12++ {
                fmt.Println(v11)
                _ = fi12
            }
            var v13 float64 = v4
            _ = v11
            _ = v13
        } else {
            v14 := v2[508]
            var v15 bool = (1 <= 1)
            for fi16 := 0; fi16 < 5; fi16++ {
                for fi17 := 0; fi17 < 19; fi17++ {
                    v4 += v4
                    _ = fi17
                }
                var v18 uint64 = 18446744073709551615
                _ = fi16
                _ = v18
            }
            _ = v14
            _ = v15
        }
        fmt.Println(v3, v5)
    } else {
        fmt.Println(v5, v5)
        v2[uint32(v3)] = "\n\t"
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        v4 = (1.5 + 0.0)
        _ = fi19
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        switch (fi20 + fi20) {
        case 42:
            v21 := []int{(fi20 / ((fi20) | 1)), fi20, fi20, fi20}
            _ = v21
        case 2:
            v22 := GenS1{F0: 65535, F1: fi20, F2: 128}
            _ = v22
        case 2147483647:
            v23 := GenS1{F0: 65535, F1: 1, F2: genFn0()}
            _ = v23
        default:
            if false {
                v3 -= 961
            } else {
                var v24 bool = v5
                _ = v24
            }
        }
        if ((v4 == (1e10 + (float64(1e10) + (float64(0.0) / float64(1.5))))) && false) {
            v4 += (-2.25 * 1e10)
        }
        _ = fi20
    }
    v4 = v4
    for fi25 := 0; fi25 < 16; fi25++ {
        var v26 float32 = 3.14159
        fmt.Println(v3, v5)
        _ = fi25
        _ = v26
    }
    v27 := GenS2{F0: "", F1: 868, F2: 370}
    v3 = 0
    v3 += (v3 << ((4294967295) & 63))
    for fi28 := 0; fi28 < 10; fi28++ {
        for fi29 := 0; fi29 < 20; fi29++ {
            v3 = v3
            fmt.Println(v5)
            _ = fi29
        }
        v30 := v2[4294967295]
        _ = fi28
        _ = v30
    }
    if (!v5) {
        v1[4294967295] = ""
    } else {
        fmt.Println(v4, v5, v3)
        v4 += (v4 - (((float64(-2.25) + -2.25) * v4) + (float64(-2.25) + v4)))
        v31 := GenS1{F0: 709, F1: -1, F2: 0}
        _ = v31
    }
    switch "a" {
    case "a":
        v3 = v3
    case "":
        if v5 {
            var v32 uint64 = 42
            var v33 int = 145
            _ = v32
            _ = v33
        } else {
            for fi34 := 0; fi34 < 9; fi34++ {
                var v35 bool = (((0.1 / ((0.0 / 3.14159) * v4)) * v4) == v4)
                switch "foo bar" {
                case "foo bar":
                    switch fi34 {
                    case 4294967295:
                        var v36 uint = 2
                        _ = v36
                    case -2147483648:
                        var v37 uint8 = genFn0()
                        _ = v37
                    }
                case "\n\t":
                    if v5 {
                        if v5 {
                            v38 := []int{}
                            var v39 uint = 952
                            _ = v38
                            _ = v39
                        } else {
                            fmt.Println(v4, v3)
                            var v40 float32 = (((1.5 * -2.25) / 1.5) - 1.5)
                            _ = v40
                        }
                        v41 := v1[0]
                        _ = v41
                    }
                }
                v42 := GenS3{F0: (!v35)}
                _ = fi34
                _ = v35
                _ = v42
            }
            var v43 float32 = (0.1 + (3.14159 * 1.5))
            _ = v43
        }
    case "\n\t":
        if v5 {
            _ = len(v2)
            v44 := GenS1{F0: 2, F1: 1, F2: 2}
            _ = v44
        } else {
            var v45 float32 = 0.1
            _ = v45
        }
    }
    var v46 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v46
    _ = v5
}
