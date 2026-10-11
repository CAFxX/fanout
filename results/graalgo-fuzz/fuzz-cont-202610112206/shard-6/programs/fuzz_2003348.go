package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int16, p1 uint, p2 uint16) int16 {
    switch 473 {
    case -1:
        for fi4 := 0; fi4 < 18; fi4++ {
            var v5 float32 = 1e10
            var v6 uint32 = 444
            v7 := []int{}
            _ = fi4
            _ = v5
            _ = v6
            _ = v7
        }
    }
    if false {
        if false {
            if (1 == -1) {
                if true {
                    var v8 uint8 = 42
                    v8 = (v8 << (((p1 - p1)) & 63))
                    p1 += 1
                    _ = v8
                }
                var v9 uint32 = 1
                _ = v9
            }
        }
        v10 := []int{1, 2147483647}
        _ = len(v10)
        _ = v10
    }
    var v11 uint16 = (p2 >> ((65535) & 63))
    var v12 int32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return p0
}

func main() {
    var v1 uint32 = 263
    v2 := make(map[uint32]uint32)
    var v3 int = -2147483648
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 int8 = 127
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            switch 1 {
            case 334:
                v3 = 2147483647
            case 42:
                switch "\n\t" {
                case "\n\t":
                    v17 := v2[(^(v1))]
                    _ = v17
                case "foo bar":
                    fmt.Println(v3, v1)
                }
            }
            if true {
                v1 = v1
                var v18 int64 = -1
                if false {
                    switch 32767 {
                    case -274:
                        v19 := []int{32767, 553, -2147483648, (v3 | 839)}
                        if (!((!(32767 == 567)) && ("" < "hello"))) {
                            var v20 uint64 = 526
                            switch (-(fi15)) {
                            case 2147483647:
                                v21 := v2[(v1 + v1)]
                                _ = v21
                            }
                            v18 = (v18 ^ -1)
                            var v22 int16 = 1000
                            _ = v20
                            _ = v22
                        }
                        var v23 uint8 = 216
                        _ = v19
                        _ = v23
                    case 42:
                        v1 -= (v1 & v1)
                    }
                    var v24 float32 = -2.25
                    v25 := []int{(fi15 & fi15), v3, fi15}
                    _ = v24
                    _ = v25
                } else {
                    v26 := GenS1{F0: 4294967295, F1: int16(v1)}
                    _ = v26
                }
                _ = v18
            } else {
                var v27 int64 = 9223372036854775807
                _ = v27
            }
            _ = fi16
        }
        if false {
            v28 := GenS1{F0: 0, F1: 1}
            _ = v28
        }
        _ = fi15
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        v30 := []int{v3, 4294967295, fi29}
        for fi31 := 0; fi31 < 7; fi31++ {
            v1 -= v1
            _ = fi31
        }
        _ = fi29
        _ = v30
    }
    v1 -= 42
    v32 := []int{v3}
    v3 += v3
    var v33 uint16 = 855
    v1 = v1
    var v34 int64 = -9223372036854775808
    fmt.Println(v3)
    v35 := GenS2{F0: v34}
    var v36 uint16 = v33
    fmt.Println(v1)
    var v37 float32 = ((0.1 * float32(float64(float32((0.1 - (1.5 / ((3.14159 / float64(-2.25)) * 0.1))))))) / ((-2.25 - float32(((1e10 - 3.14159) / (1e10 + 1.5)))) - float32((3.14159 * -2.25))))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
