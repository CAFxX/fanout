package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint64 = 0
    var v2 uint = 1
    var v3 float64 = float64(0.1)
    v4 := []uint32{623}
    v5 := make(map[uint32]int)
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: -289}
        for fi8 := 0; fi8 < 4; fi8++ {
            v1 -= (v1 ^ (v1 % (((v1 & (^(v1)))) | 1)))
            var v9 int64 = -1
            delete(v5, 0)
            var v10 int64 = v9
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    v1 = v1
    switch "a" {
    case "foo bar":
        var v11 int = 127
        _ = v11
    case "hello":
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 float64 = (v3 - ((v3 / v3) / (v3 * v3)))
            _ = fi12
            _ = v13
        }
    case "a":
        switch "foo bar" {
        case "hello":
            fmt.Println(v3)
        default:
            fmt.Println(v2)
        }
    }
    if (684 == 388) {
        switch -2147483648 {
        case 9223372036854775807:
            for fi14 := 0; fi14 < 4; fi14++ {
                if (fi14 <= fi14) {
                    v1 = (^(v1))
                    fmt.Println(v1)
                    v3 -= v3
                } else {
                    fmt.Println(fi14, v2)
                    var v15 uint32 = 42
                    _ = v15
                }
                _ = fi14
            }
        case 42:
            fmt.Println(v3, v3, v2)
        }
        var v16 uint64 = v1
        _ = v16
    }
    v3 = 0.1
    var v17 int32 = -829
    v18 := GenS2{F0: (-(v17))}
    for fi19 := 0; fi19 < 16; fi19++ {
        delete(v5, 2147483647)
        var v20 uint32 = 84
        var v21 int32 = v17
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println(v17)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v22 uint32 = 2
    delete(v5, 727)
    if (!true) {
        for fi23 := 0; fi23 < 15; fi23++ {
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 int32 = 268
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
        for fi26 := 0; fi26 < 5; fi26++ {
            if (v3 > (v3 - v3)) {
                if ((v3 / 1e10) == (float64(0.0) + v3)) {
                    v27 := v5[v22]
                    if (!(false || ((0.1 + 0.0) != float32((v3 * v3))))) {
                        v1 = v1
                        for fi28 := 0; fi28 < 7; fi28++ {
                            var v29 float32 = 1e10
                            var v30 uint16 = 0
                            if false {
                                v27 -= int(v22)
                                var v31 string = (("\n\t" + "") + "hello")
                                _ = v31
                            }
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                        for fi32 := 0; fi32 < 7; fi32++ {
                            v27 -= -2147483648
                            _ = fi32
                        }
                    }
                    var v33 uint = v2
                    _ = v27
                    _ = v33
                } else {
                    var v34 int8 = 1
                    _ = v34
                }
                if false {
                    var v35 float32 = 0.0
                    _ = v35
                }
            }
            v3 += 1e10
            for fi36 := 0; fi36 < 5; fi36++ {
                if false {
                    var v37 uint32 = v22
                    v2 -= v2
                    v22 += (v37 / ((4294967295) | 1))
                    _ = v37
                } else {
                    v5[v22] = int(v22)
                    delete(v5, 42)
                    v22 = (v22 + 1)
                }
                _ = fi36
            }
            _ = fi26
        }
    } else {
        if true {
            v4 = append(v4, (v22 % (((^(v22))) | 1)))
            if false {
                v2 = v2
                var v38 int = 67
                _ = v38
            }
        } else {
            v2 += v2
            var v39 string = "\n\t"
            _ = v39
        }
        var v40 float32 = float32((((1e10 + (3.14159 / v3)) / 1.5) / (v3 / v3)))
        if false {
            switch "hello" {
            case "foo bar":
                for fi41 := 0; fi41 < 9; fi41++ {
                    v42 := GenS2{F0: v17}
                    v17 -= (v17 << (((v2 << ((255) & 63))) & 63))
                    var v43 int8 = -33
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            case "hello":
                fmt.Println(v17, v22)
            }
        } else {
            v44 := GenS2{F0: -528}
            _ = v44
        }
        _ = v40
    }
    var v45 uint8 = 128
    v46 := GenS2{F0: v17}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
}
