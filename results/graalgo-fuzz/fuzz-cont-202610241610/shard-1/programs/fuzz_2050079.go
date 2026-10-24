package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 1
    var v3 uint16 = 289
    var v4 int16 = 1
    v5 := make(map[uint32]string)
    var v6 int8 = (v2 >> ((18446744073709551615) & 63))
    v4 = v4
    v7 := GenS1{F0: ((v3 > 42) && false), F1: 18446744073709551615, F2: 582}
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: true, F1: 209, F2: -1}
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: true, F1: 2, F2: 42}
    v11 := GenS1{F0: (v6 != v6), F1: 2, F2: 42}
    if (true && true) {
        v12 := []int{32767, 127, -152}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 6; fi13++ {
            if true {
                if (true || ((^(fi13)) == 165)) {
                    var v14 uint16 = 32767
                    for fi15 := 0; fi15 < 12; fi15++ {
                        var v16 int64 = 735
                        _ = fi15
                        _ = v16
                    }
                    _ = v14
                } else {
                    var v17 uint8 = 42
                    v3 = v3
                    v17 += (v17 >> ((4294967295) & 63))
                    _ = v17
                }
            }
            var v18 uint32 = 65535
            _ = fi13
            _ = v18
        }
    }
    var v19 string = "a"
    if (v6 <= v2) {
        if ((!(2 == 18446744073709551615)) || (-70 >= -11)) {
            var v20 float64 = float64((1e10 - ((-2.25 + -2.25) - (0.0 / 1e10))))
            _ = v20
        }
        v21 := GenS1{F0: (2 > 0), F1: 409, F2: -9223372036854775808}
        _ = v21
    } else {
        var v22 uint8 = 2
        v19 = ((("a" + ("\n\t" + v19)) + "") + (("foo bar" + v19) + v19))
        _ = v22
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        var v24 bool = false
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        v2 = (-(v6))
        _ = fi25
    }
    v3 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
