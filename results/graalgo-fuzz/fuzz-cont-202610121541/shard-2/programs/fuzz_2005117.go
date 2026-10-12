package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
func genFn0(p0 int16, p1 uint64) uint8 {
    var v6 uint8 = 0
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := []int{fi7, fi7}
        var v9 string = ("a" + ("foo bar" + ("\n\t" + (("" + "hello") + ""))))
        if false {
            if false {
                v10 := GenS1{F0: (fi7 * (fi7 & fi7)), F1: (fi7 - (fi7 << ((2) & 63)))}
                _ = v10
            } else {
                fmt.Println(v9)
            }
            v6 -= 254
            var v11 uint64 = 1
            _ = v11
        } else {
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 int64 = -1
                _ = fi12
                _ = v13
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v14 uint8 = (^(v6))
    v14 = v14
    v15 := GenS1{F0: 32767, F1: 392}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v6
    return 127
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 1
    var v3 bool = (true && (15 == 183))
    var v4 bool = v3
    var v5 int16 = 1
    if v4 {
        if true {
            for fi16 := 0; fi16 < 11; fi16++ {
                for fi17 := 0; fi17 < 6; fi17++ {
                    v18 := GenS1{F0: fi17, F1: 128}
                    var v19 float32 = float32(3.14159)
                    v4 = true
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v5 = 32767
                v20 := GenS1{F0: -338, F1: fi16}
                _ = fi16
                _ = v20
            }
        }
        fmt.Println(v4, v4, v5)
        var v21 uint = 18446744073709551615
        _ = v21
    } else {
        if v4 {
            for fi22 := 0; fi22 < 4; fi22++ {
                switch (fi22 % ((-52) | 1)) {
                case 1:
                    v5 -= v5
                case 9223372036854775807:
                    var v23 int32 = -477
                    _ = v23
                default:
                    v3 = v4
                }
                v5 = v5
                _ = fi22
            }
            v24 := []int{-2147483648, -873, 680}
            _ = v24
        }
        for fi25 := 0; fi25 < 9; fi25++ {
            if true {
                if false {
                    v26 := GenS1{F0: (fi25 ^ (fi25 ^ -529)), F1: int(v2)}
                    _ = v26
                }
                v5 -= v5
            } else {
                v2 = 4294967295
            }
            var v27 uint64 = (v2 / ((v2) | 1))
            _ = fi25
            _ = v27
        }
    }
    v28 := GenS1{F0: -579, F1: 801}
    v29 := GenS1{F0: 2147483647, F1: 109}
    if (!v3) {
        var v30 float32 = 1e10
        fmt.Println(v4, v5)
        _ = v30
    } else {
        var v31 int16 = v5
        fmt.Println(v3)
        _ = v31
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        var v33 int32 = -413
        _ = fi32
        _ = v33
    }
    v34 := GenS1{F0: 368, F1: -1}
    v4 = v3
    for fi35 := 0; fi35 < 3; fi35++ {
        v5 = 767
        fmt.Println(v4)
        var v36 uint8 = 185
        _ = fi35
        _ = v36
    }
    v4 = v4
    v5 -= 780
    var v37 float64 = (3.14159 / (-2.25 - (1.5 + 3.14159)))
    for fi38 := 0; fi38 < 8; fi38++ {
        v5 = -32768
        _ = fi38
    }
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v37
    _ = v4
    _ = v5
}
