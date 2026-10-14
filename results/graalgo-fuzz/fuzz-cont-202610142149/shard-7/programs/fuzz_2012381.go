package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 bool }
type GenS2 struct { F0 uint8 }
func genFn0() uint {
    var v5 uint8 = 255
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 float64 = 0.0
        v8 := GenS1{F0: 37, F1: 526, F2: true}
        _ = fi6
        _ = v7
        _ = v8
    }
    if ((2147483647 >= 1) || (!false)) {
        var v9 int8 = 1
        _ = v9
    } else {
        v5 = (v5 % ((2) | 1))
        v5 += (v5 | v5)
        for fi10 := 0; fi10 < 17; fi10++ {
            for fi11 := 0; fi11 < 6; fi11++ {
                v12 := GenS1{F0: (^(fi11)), F1: 2, F2: (!true)}
                fmt.Println(fi10)
                v5 = 4
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    }
    fmt.Println(v5)
    if (((1000 > -342) || (18446744073709551615 == 0)) || true) {
        v13 := GenS1{F0: -9223372036854775808, F1: 573, F2: true}
        fmt.Println(v5)
        for fi14 := 0; fi14 < 6; fi14++ {
            v15 := []int{fi14}
            v5 = (v5 << ((873) & 63))
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    _ = v5
    return 789
}
func genFn1(p0 uint32) int64 {
    fmt.Println(p0)
    if true {
        var v16 int32 = 2147483647
        _ = v16
    }
    _ = p0
    return -708
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    var v3 uint64 = 4294967295
    v4 := make(map[int]bool)
    v17 := GenS1{F0: -35, F1: 0, F2: true}
    v18 := GenS1{F0: 1, F1: genFn0(), F2: (!(0.1 <= float64(0.1)))}
    v3 += (v3 % (((v3 / (((v3 % ((4294967295) | 1))) | 1))) | 1))
    v19 := []int{}
    var v20 uint16 = 2
    var v21 uint16 = (v20 | v20)
    for fi22 := 0; fi22 < 6; fi22++ {
        switch "" {
        case "hello":
            if (true && true) {
                fmt.Println(v20, v3, fi22)
                v3 = 285
                switch fi22 {
                case -892:
                    var v23 int16 = -1
                    _ = v23
                case 4294967295:
                    switch "hello" {
                    case "foo bar":
                        v24 := GenS2{F0: 1}
                        _ = v24
                    }
                }
            } else {
                switch (fi22 >> ((307) & 63)) {
                case 1:
                    _ = len(v1)
                case -2147483648:
                    var v25 uint8 = uint8(v3)
                    _ = v25
                default:
                    var v26 int32 = -764
                    _ = v26
                }
                v21 = v20
                v3 -= v3
            }
        }
        _ = fi22
    }
    var v27 int64 = genFn1(107)
    _ = len(v19)
    var v28 uint8 = 11
    v21 += v20
    var v29 uint64 = (v3 | (^(v3)))
    var v30 uint16 = v20
    for fi31 := 0; fi31 < 10; fi31++ {
        switch 0 {
        case 984:
            switch "foo bar" {
            case "foo bar":
                v2[(("\n\t" + "hello") + "hello")] = false
            }
        default:
            var v32 bool = (233 != 18446744073709551615)
            _ = v32
        }
        v33 := GenS2{F0: 2}
        _ = fi31
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
