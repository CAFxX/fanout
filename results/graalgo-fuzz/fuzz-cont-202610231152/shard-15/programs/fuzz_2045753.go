package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint16; F2 bool }
type GenS4 struct { F0 uint16; F1 int16; F2 uint }

func main() {
    var v1 bool = true
    var v2 uint16 = 2
    var v3 bool = ((v1 && (82 == 18446744073709551615)) || (!true))
    var v4 uint = 255
    var v5 int16 = 32767
    v6 := make(map[uint32]string)
    v7 := GenS1{F0: -9223372036854775808}
    v8 := GenS2{F0: 42}
    fmt.Println(v4)
    var v9 string = "a"
    var v10 string = v9
    v5 = v5
    var v11 uint32 = 2147483647
    var v12 uint64 = 166
    v13 := GenS1{F0: 349}
    if v1 {
        fmt.Println(v10, v2, v9)
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 string = ""
        _ = fi14
        _ = v15
    }
    v16 := GenS3{F0: 1, F1: (v2 + (v2 | (v2 * 255))), F2: (v3 || ((v3 || ((585 >= 4294967295) && ((v11 << ((4294967295) & 63)) <= (v11 / ((v11) | 1))))) || false))}
    var v17 int16 = (v5 + -1)
    if (v2 == (^(v2))) {
        for fi18 := 0; fi18 < 11; fi18++ {
            if (v17 < v5) {
                v19 := []int{(^(fi18)), fi18, fi18}
                v3 = ((fi18 * 778) < (fi18 | fi18))
                _ = v19
            }
            if false {
                var v20 uint32 = (v11 - 4294967295)
                v21 := GenS2{F0: -1}
                _ = v20
                _ = v21
            }
            _ = fi18
        }
        v17 = (^(v5))
        v10 = ("a" + v10)
    } else {
        if v3 {
            for fi22 := 0; fi22 < 14; fi22++ {
                if (((42 >= -1) || (!(v2 != (^(v2))))) || false) {
                    v23 := []int{}
                    v11 = uint32(v12)
                    v2 = (v2 % ((875) | 1))
                    _ = v23
                }
                if ((v5 << ((v4) & 63)) < (v17 - (-(v17)))) {
                    for fi24 := 0; fi24 < 1; fi24++ {
                        switch fi24 {
                        case 795:
                            switch v10 {
                            case "":
                                var v25 uint = (v4 << ((v4) & 63))
                                _ = v25
                            case "foo bar":
                                fmt.Println(v5, v9, v11)
                            }
                            v26 := GenS1{F0: int(v2)}
                            _ = v26
                        default:
                            if v1 {
                                v2 -= 42
                            }
                        }
                        var v27 int = 42
                        var v28 int64 = -327
                        _ = fi24
                        _ = v27
                        _ = v28
                    }
                    v9 = ("a" + (v9 + ((v9 + "hello") + v9)))
                }
                v29 := GenS4{F0: (v2 & 261), F1: 1, F2: (v4 >> ((1) & 63))}
                _ = fi22
                _ = v29
            }
            var v30 float64 = (0.1 - 0.1)
            v6[0] = (("hello" + (v10 + "")) + v9)
            _ = v30
        } else {
            var v31 float64 = float64(3.14159)
            _ = v31
        }
        var v32 bool = false
        _ = v32
    }
    v33 := v6[v11]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
