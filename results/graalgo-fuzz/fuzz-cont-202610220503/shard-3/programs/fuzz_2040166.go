package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int16 }
type GenS2 struct { F0 uint8 }
func genFn0() uint8 {
    var v6 uint = 715
    var v7 int8 = -1
    v8 := GenS1{F0: 326, F1: false, F2: -32768}
    _ = v6
    _ = v7
    _ = v8
    return 128
}
func genFn1(p0 string) int64 {
    if true {
        v9 := GenS1{F0: 65535, F1: true, F2: 1000}
        for fi10 := 0; fi10 < 13; fi10++ {
            if true {
                v11 := GenS1{F0: 2, F1: false, F2: 32767}
                _ = v11
            }
            if true {
                p0 = "hello"
                var v12 int = 940
                _ = v12
            }
            _ = fi10
        }
        v13 := []int{}
        _ = v13
        _ = v9
    }
    v14 := []int{1, -462}
    if len(v14) > 0 {
        v14[0] = 42
    }
    _ = p0
    _ = v14
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]int)
    v3 := []uint32{2147483648}
    v4 := []int{1, 821, 1}
    var v5 int32 = 918
    v15 := GenS2{F0: 127}
    var v16 uint8 = 1
    v5 -= (v5 / ((-492) | 1))
    fmt.Println(v16)
    for fi17 := 0; fi17 < 5; fi17++ {
        for fi18 := 0; fi18 < 14; fi18++ {
            v19 := GenS2{F0: v16}
            v16 = v16
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    if false {
        var v20 string = "a"
        v20 = "a"
        _ = v20
    }
    v21 := GenS2{F0: 127}
    switch ("" + "a") {
    case "":
        fmt.Println(v16, v5)
    case "a":
        v5 += (^(v5))
    case "\n\t":
        v16 = genFn0()
    }
    if (((!(0.0 == (3.14159 + 1e10))) || (genFn0() < 0)) && (2147483648 >= 2)) {
        if true {
            switch 1 {
            case 423:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
        v5 = 2147483647
    } else {
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 int8 = 42
            v23 -= 75
            switch 349 {
            case 128:
                fmt.Println(fi22, v16)
            default:
                if len(v4) > 0 {
                    v4[0] = fi22
                }
            }
            _ = fi22
            _ = v23
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v5 -= (v5 >> ((2) & 63))
    }
    switch -1 {
    case 128:
        v3 = append(v3, 65535)
    case 572:
        fmt.Println(v16)
    default:
        v24 := v1[-820]
        _ = v24
    }
    v16 += v16
    var v25 int32 = v5
    v26 := GenS1{F0: 255, F1: (538 > 127), F2: -617}
    v1[-2147483648] = -1
    v5 = (-(v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
