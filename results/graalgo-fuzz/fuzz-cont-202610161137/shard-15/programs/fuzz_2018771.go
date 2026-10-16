package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
func genFn0() uint8 {
    if (!true) {
        var v7 float64 = 1.5
        _ = v7
    }
    var v8 int64 = 9223372036854775807
    _ = v8
    return 127
}
func genFn1(p0 int32, p1 uint, p2 uint32) int64 {
    var v9 string = ("\n\t" + "")
    if false {
        v10 := GenS1{F0: 122}
        var v11 int = 4294967295
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}

func main() {
    var v1 int64 = 14
    v2 := make(map[uint32]int)
    var v3 float64 = float64(0.0)
    var v4 uint64 = 449
    var v5 uint32 = uint32(v4)
    var v6 float32 = ((((-2.25 + 1.5) - float32((v3 / v3))) - -2.25) / -2.25)
    v3 += float64((v6 * (v6 * v6)))
    for fi12 := 0; fi12 < 8; fi12++ {
        fmt.Println(v4)
        v13 := []int{fi12}
        for fi14 := 0; fi14 < 14; fi14++ {
            v1 -= 0
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    var v15 bool = false
    fmt.Println(v6, v4)
    if (((!v15) && false) && (128 == 249)) {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 uint16 = 42
            _ = fi16
            _ = v17
        }
    }
    var v18 uint16 = uint16(v4)
    for fi19 := 0; fi19 < 10; fi19++ {
        switch fi19 {
        case 4294967295:
            v4 += (v4 % ((v4) | 1))
        case -9223372036854775808:
            switch fi19 {
            case -2147483648:
                var v20 int = fi19
                _ = v20
            case -1:
                fmt.Println(v15, v18, v4)
            default:
                switch (("a" + "") + "foo bar") {
                case "hello":
                    var v21 int8 = -128
                    _ = v21
                case "a":
                    if v15 {
                        var v22 int = (fi19 >> ((4294967295) & 63))
                        if (1 != 1) {
                            var v23 int32 = -1
                            _ = v23
                        } else {
                            fmt.Println(v15, v18, v3)
                            if false {
                                switch "hello" {
                                case "a":
                                    v15 = false
                                case "hello":
                                    fmt.Println(v1, fi19)
                                case "\n\t":
                                    fmt.Println(v4)
                                }
                            } else {
                                fmt.Println(v15)
                                for fi24 := 0; fi24 < 6; fi24++ {
                                    if v15 {
                                        v25 := GenS2{F0: int16(v5)}
                                        v26 := []int{}
                                        v5 = v5
                                        _ = v25
                                        _ = v26
                                    } else {
                                        fmt.Println(v15)
                                    }
                                    _ = fi24
                                }
                                var v27 uint32 = (v5 / ((v5) | 1))
                                _ = v27
                            }
                        }
                        for fi28 := 0; fi28 < 1; fi28++ {
                            var v29 int16 = 42
                            _ = fi28
                            _ = v29
                        }
                        _ = v22
                    }
                case "foo bar":
                    v30 := v2[uint32(v4)]
                    _ = v30
                default:
                    v31 := GenS2{F0: 1}
                    _ = v31
                }
            }
        case -703:
            v2[(^(v5))] = fi19
        default:
            if v15 {
                var v32 string = "hello"
                v5 = 0
                for fi33 := 0; fi33 < 15; fi33++ {
                    v6 = v6
                    fmt.Println(v1)
                    var v34 uint8 = genFn0()
                    _ = fi33
                    _ = v34
                }
                _ = v32
            }
        }
        var v35 int = (fi19 >> ((259) & 63))
        _ = fi19
        _ = v35
    }
    delete(v2, (v5 * (v5 + 2147483648)))
    v36 := GenS3{F0: 216}
    v37 := GenS4{F0: 2}
    switch -2147483648 {
    case 42:
        if (-40 < -1) {
            v5 -= (v5 | v5)
        }
    case 128:
        v38 := []int{-2147483648, 924, -2147483648, -2147483648}
        _ = v38
    }
    v4 += (v4 | v4)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
