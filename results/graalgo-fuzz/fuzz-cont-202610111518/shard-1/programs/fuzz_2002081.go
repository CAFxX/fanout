package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() int64 {
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int8 = -1
        _ = fi7
        _ = v8
    }
    v9 := []int{403, 2147483647}
    var v10 bool = (!true)
    v10 = (83 > -961)
    _ = v10
    _ = v9
    return -1
}
func genFn1(p0 int32, p1 uint64) uint {
    p1 = (p1 * (^(p1)))
    v11 := []int{190}
    for fi12 := 0; fi12 < 4; fi12++ {
        p1 = 0
        if len(v11) > 0 {
            v11[0] = fi12
        }
        _ = fi12
    }
    var v13 int = 9223372036854775807
    _ = p0
    _ = p1
    _ = v11
    _ = v13
    return 4294967295
}

func main() {
    v1 := []uint32{}
    var v2 int64 = -814
    v3 := make(map[int]int)
    v4 := make(map[int]string)
    v5 := make(map[string]uint32)
    var v6 uint16 = 65535
    switch "" {
    case "hello":
        _ = len(v4)
    case "":
        var v14 float64 = 0.1
        _ = v14
    case "\n\t":
        v2 += -582
    }
    var v15 uint32 = 485
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := v4[0]
        _ = fi16
        _ = v17
    }
    if (true || (!(!true))) {
        switch -449 {
        case 932:
            for fi18 := 0; fi18 < 13; fi18++ {
                for fi19 := 0; fi19 < 2; fi19++ {
                    if true {
                        var v20 int8 = -128
                        v2 += (^(v2))
                        v6 = v6
                        _ = v20
                    }
                    _ = fi19
                }
                for fi21 := 0; fi21 < 5; fi21++ {
                    fmt.Println(v2)
                    var v22 float32 = (-2.25 / -2.25)
                    v22 -= (1.5 / (v22 * v22))
                    if true {
                        var v23 int32 = -801
                        v24 := GenS1{F0: 42}
                        _ = v23
                        _ = v24
                    } else {
                        v4[(fi21 * -203)] = ("" + "foo bar")
                        v2 -= (v2 % ((genFn0()) | 1))
                    }
                    _ = fi21
                    _ = v22
                }
                for fi25 := 0; fi25 < 8; fi25++ {
                    var v26 int32 = 901
                    fmt.Println(fi18, fi18)
                    _ = fi25
                    _ = v26
                }
                _ = fi18
            }
        }
        v2 = v2
        v6 += uint16(v15)
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        v2 -= (v2 - v2)
        fmt.Println(v6)
        switch ("" + "\n\t") {
        case "a":
            v2 = v2
        case "hello":
            v15 -= (v15 - v15)
        }
        _ = fi27
    }
    switch "" {
    case "\n\t":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v6 = v6
    if len(v1) > 0 {
        v1[0] = 1
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
