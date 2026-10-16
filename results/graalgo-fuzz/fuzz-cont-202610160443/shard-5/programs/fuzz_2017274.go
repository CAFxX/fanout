package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int16; F1 uint64; F2 bool }
type GenS3 struct { F0 uint32; F1 int16 }
func genFn0() int16 {
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 int16 = 42
        if true {
            var v6 uint32 = 4294967295
            _ = v6
        }
        _ = fi4
        _ = v5
    }
    v7 := []int{}
    var v8 int32 = -1
    if (false && true) {
        if (!(0 <= 42)) {
            var v9 bool = (!(!("a" != "foo bar")))
            v9 = (1 == -66)
            _ = v9
        }
    } else {
        v10 := GenS1{F0: 2, F1: ((v8 | v8) > v8)}
        var v11 uint = 4294967295
        _ = v10
        _ = v11
    }
    v8 += v8
    _ = v7
    _ = v8
    return int16(v8)
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    var v3 uint16 = 1
    switch -2147483648 {
    case -264:
        for fi12 := 0; fi12 < 16; fi12++ {
            for fi13 := 0; fi13 < 15; fi13++ {
                v3 += v3
                _ = fi13
            }
            _ = fi12
        }
    case -861:
        v3 -= v3
    default:
        switch 9223372036854775807 {
        case 127:
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 float32 = ((1e10 * 3.14159) * 0.1)
                _ = fi14
                _ = v15
            }
        case -465:
            v16 := []int{}
            _ = v16
        }
    }
    v3 += v3
    v3 = (v3 << ((255) & 63))
    if true {
        v3 += v3
        v3 = 65535
        v17 := GenS1{F0: 0, F1: true}
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        if false {
            v19 := []int{fi18, 632, (fi18 << ((0) & 63))}
            for fi20 := 0; fi20 < 6; fi20++ {
                var v21 int32 = 2147483647
                if (252 == 387) {
                    if false {
                        switch fi18 {
                        case 32768:
                            v21 = v21
                        default:
                            if len(v19) > 0 {
                                v19[0] = (fi20 * (^(fi18)))
                            }
                        }
                        v21 -= v21
                    } else {
                        fmt.Println(fi20, fi18)
                        var v22 int32 = v21
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        v23 := GenS1{F0: 2, F1: false}
                        v22 -= (v21 << ((4294967295) & 63))
                        _ = v22
                        _ = v23
                    }
                    v24 := GenS1{F0: 93, F1: (1000 >= -101)}
                    v25 := GenS1{F0: 498, F1: ("hello" > ("\n\t" + ("hello" + ("foo bar" + "hello"))))}
                    _ = v24
                    _ = v25
                }
                v26 := GenS1{F0: 1, F1: (false && false)}
                if (-2.25 >= 1e10) {
                    v3 = 0
                    v27 := v1[42]
                    _ = v27
                } else {
                    v3 += 42
                    fmt.Println(fi18, fi20, fi20)
                    if true {
                        fmt.Println(v3)
                    } else {
                        var v28 uint16 = v3
                        v29 := GenS1{F0: 0, F1: (0 >= 2)}
                        _ = v28
                        _ = v29
                    }
                }
                fmt.Println(v21, fi20)
                v30 := GenS1{F0: 536, F1: (("\n\t" + "a") >= "")}
                _ = len(v1)
                _ = fi20
                _ = v21
                _ = v26
                _ = v30
            }
            _ = v19
        }
        v1[1] = 1
        _ = fi18
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v3)
    case "":
        for fi31 := 0; fi31 < 20; fi31++ {
            var v32 int64 = int64(v3)
            v3 += 255
            _ = fi31
            _ = v32
        }
    default:
        fmt.Println(v3)
    }
    switch -2147483648 {
    case 9223372036854775807:
        v33 := GenS2{F0: 32767, F1: 0, F2: ((-1 >= -99) && ((!false) || ((!false) && false)))}
        _ = v33
    }
    v3 = v3
    v3 = (v3 >> ((2) & 63))
    v3 = (^(v3))
    var v34 uint32 = 2
    v35 := GenS3{F0: v34, F1: -2}
    var v36 bool = true
    var v37 int32 = -1
    v36 = (-2.25 == (3.14159 + 3.14159))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
