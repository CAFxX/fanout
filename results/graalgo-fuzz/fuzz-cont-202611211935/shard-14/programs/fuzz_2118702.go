package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 bool) uint {
    p0 = (-1 <= -946)
    fmt.Println(p0)
    if ((65535 <= 952) && ((1 > 1) || (-61 > 119))) {
        p0 = ((!(-2147483648 > 42)) || p0)
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: 4294967295}
        if (255 > 32767) {
            var v9 int16 = 42
            v10 := []int{-2147483648, fi7}
            var v11 uint16 = 36
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    p0 = ((0.1 - 3.14159) <= float64(0.0))
    _ = p0
    return 4294967295
}
func genFn1(p0 bool) int8 {
    fmt.Println(p0)
    var v12 int64 = 319
    _ = p0
    _ = v12
    return 74
}

func main() {
    var v1 float32 = -2.25
    v2 := []uint32{2147483648, 415, 721, 207}
    v3 := make(map[uint32]int)
    var v4 uint64 = 0
    var v5 uint8 = 0
    var v6 float64 = -2.25
    for fi13 := 0; fi13 < 1; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 string = "\n\t"
            if true {
                _ = len(v2)
            }
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    v5 = 1
    _ = len(v2)
    v4 = 1
    switch ("foo bar" + ((("a" + "\n\t") + "\n\t") + "")) {
    case "\n\t":
        for fi16 := 0; fi16 < 3; fi16++ {
            v17 := GenS2{F0: 42}
            _ = fi16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 int = fi18
            if (42 > 1) {
                if false {
                    var v21 int32 = -2147483648
                    fmt.Println(v6, v4)
                    _ = v21
                } else {
                    if true {
                        if ((!true) && false) {
                            v1 -= v1
                        }
                    }
                    var v22 uint = 2
                    if (float32(-2.25) > (((v1 - (0.1 - v1)) / 1.5) - 0.0)) {
                        v5 = (v5 - (v5 ^ uint8(v22)))
                    }
                    _ = v22
                }
            } else {
                var v23 float32 = float32(-2.25)
                _ = v23
            }
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    var v24 uint16 = 42
    fmt.Println(v5, v5, v24)
    v25 := GenS2{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
