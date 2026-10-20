package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 bool; F1 bool; F2 int16 }

func main() {
    var v1 uint16 = 255
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    var v4 uint16 = v1
    v5 := GenS1{F0: -1, F1: 9223372036854775807}
    v1 += v1
    var v6 int16 = 1000
    switch "" {
    case "":
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 bool = (true || false)
            _ = fi7
            _ = v8
        }
    case "hello":
        for fi9 := 0; fi9 < 20; fi9++ {
            v1 -= v4
            _ = fi9
        }
    }
    v6 = (v6 << ((18446744073709551615) & 63))
    if false {
        if (!(2147483647 > 42)) {
            for fi10 := 0; fi10 < 20; fi10++ {
                switch "" {
                case "":
                    v4 = (^(v1))
                case "a":
                    if true {
                        for fi11 := 0; fi11 < 1; fi11++ {
                            var v12 uint8 = uint8(v4)
                            v1 += 2
                            v13 := GenS1{F0: 1, F1: 939}
                            if (!(2147483647 >= -364)) {
                                var v14 string = (((("a" + ("" + "\n\t")) + "") + "hello") + "\n\t")
                                _ = v14
                            } else {
                                var v15 int16 = (v6 << ((955) & 63))
                                var v16 int64 = 9223372036854775807
                                v17 := []int{(fi11 << ((2) & 63)), (fi11 % (((fi10 | (fi11 >> ((2) & 63)))) | 1)), (fi11 + fi10), fi11}
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        var v18 uint = 0
                        var v19 uint = (v18 >> ((2) & 63))
                        _ = v18
                        _ = v19
                    } else {
                        v20 := v2[2147483648]
                        delete(v3, 0)
                        var v21 uint32 = 498
                        _ = v20
                        _ = v21
                    }
                }
                _ = fi10
            }
            var v22 int16 = v6
            _ = v22
        } else {
            var v23 int8 = 0
            _ = v23
        }
    } else {
        fmt.Println(v1, v4, v4)
    }
    v2[2] = ("" + "")
    var v24 int = -795
    for fi25 := 0; fi25 < 20; fi25++ {
        v26 := GenS1{F0: 9223372036854775807, F1: 9223372036854775807}
        _ = fi25
        _ = v26
    }
    v1 += v4
    if (457 != 533) {
        v27 := GenS1{F0: -368, F1: -1}
        for fi28 := 0; fi28 < 17; fi28++ {
            for fi29 := 0; fi29 < 9; fi29++ {
                v30 := GenS2{F0: (2 >= 2), F1: ((false || false) && ("hello" != ("a" + "hello"))), F2: (v6 & v6)}
                _ = fi29
                _ = v30
            }
            fmt.Println(v24)
            _ = fi28
        }
        _ = v27
    } else {
        var v31 float64 = 1.5
        var v32 uint64 = 4294967295
        _ = v31
        _ = v32
    }
    v33 := []int{}
    var v34 string = "\n\t"
    v4 = (v1 * (^(v1)))
    fmt.Println(v4, v6, v4)
    v35 := GenS1{F0: 868, F1: 1}
    if false {
        var v36 int64 = -729
        if false {
            v24 += v24
            if (false || false) {
                v37 := GenS1{F0: v36, F1: (v36 >> ((206) & 63))}
                _ = v37
            } else {
                v1 = (v4 * (v4 - (v1 * (v4 % ((36) | 1)))))
                fmt.Println(v6, v1, v6)
            }
        }
        _ = v36
    }
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
