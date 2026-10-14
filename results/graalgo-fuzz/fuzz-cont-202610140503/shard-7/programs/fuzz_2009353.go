package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 uint64; F2 uint32 }
func genFn0(p0 int32, p1 uint32, p2 uint16) uint16 {
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 string = ""
        _ = fi4
        _ = v5
    }
    p2 = 2
    v6 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p2 & 42)
}

func main() {
    var v1 int8 = 0
    v2 := make(map[string]uint32)
    var v3 string = "\n\t"
    var v7 string = "hello"
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := GenS1{F0: 127}
        switch fi8 {
        case -9223372036854775808:
            v10 := v2[("" + (v3 + v3))]
            _ = v10
        case 48:
            var v11 uint64 = 18446744073709551615
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    var v12 int32 = 2147483647
    switch v7 {
    case "":
        v13 := GenS1{F0: 1}
        _ = v13
    default:
        var v14 float64 = float64(1.5)
        _ = v14
    }
    switch ((v3 + "") + ((v7 + v3) + v3)) {
    case "foo bar":
        v1 -= -128
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 uint8 = 7
        _ = fi15
        _ = v16
    }
    fmt.Println(v1, v12)
    v1 = v1
    var v17 uint8 = 42
    if ((0.1 * (1.5 + -2.25)) > -2.25) {
        v18 := []int{1, -2147483648, -210}
        v17 = v17
        _ = v18
    } else {
        v1 = v1
        v17 = v17
        switch 9223372036854775807 {
        case 2147483647:
            v3 = (v3 + "\n\t")
        case 32767:
            if (2 >= (v17 & (v17 * v17))) {
                for fi19 := 0; fi19 < 19; fi19++ {
                    v12 += (v12 << ((911) & 63))
                    for fi20 := 0; fi20 < 2; fi20++ {
                        v3 = v3
                        v7 = v7
                        if (((v1 - 127) > (v1 / ((84) | 1))) && false) {
                            switch "" {
                            case "":
                                v1 -= (v1 + v1)
                            case "a":
                                _ = len(v2)
                            }
                        } else {
                            if true {
                                var v21 int16 = -32768
                                switch v7 {
                                case "":
                                    v22 := GenS2{F0: (-(v12)), F1: (true && ((v12 % ((v12) | 1)) <= v12))}
                                    _ = v22
                                case "\n\t":
                                    if ((false && true) && false) {
                                        v17 += 128
                                    }
                                }
                                _ = v21
                            }
                        }
                        _ = fi20
                    }
                    _ = fi19
                }
            } else {
                if false {
                    v2[v7] = 42
                }
                v23 := []int{}
                _ = v23
            }
        case 572:
            v24 := GenS3{F0: uint16(v17)}
            _ = v24
        default:
            v3 = v3
        }
    }
    v12 = -257
    if (((!(false || true)) || false) && (207 == 146)) {
        for fi25 := 0; fi25 < 9; fi25++ {
            v3 = v3
            _ = fi25
        }
        var v26 int16 = 32767
        _ = v26
    } else {
        v27 := GenS4{F0: -673, F1: 0, F2: 351}
        v7 = v7
        _ = v27
    }
    if (!true) {
        var v28 int16 = -32768
        _ = v28
    } else {
        var v29 float32 = 1.5
        v1 = v1
        v17 += v17
        _ = v29
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        v17 = v17
        _ = len(v2)
        _ = fi30
    }
    if (!true) {
        delete(v2, "a")
        v31 := []int{2147483647, -593, -922, 255}
        for fi32 := 0; fi32 < 10; fi32++ {
            v1 = (v1 * v1)
            var v33 uint8 = v17
            if len(v31) >= 2 {
                v31 = v31[1:len(v31)]
            }
            _ = fi32
            _ = v33
        }
        _ = v31
    }
    v34 := GenS3{F0: 668}
    v7 = v7
    if ((!(!true)) && (1 == (v12 >> ((458) & 63)))) {
        var v35 int16 = -32768
        _ = len(v2)
        for fi36 := 0; fi36 < 18; fi36++ {
            var v37 uint32 = 464
            v12 = (v12 << ((42) & 63))
            _ = fi36
            _ = v37
        }
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v7
}
