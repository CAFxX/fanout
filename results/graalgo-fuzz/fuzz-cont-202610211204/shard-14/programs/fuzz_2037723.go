package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint64 }
func genFn0(p0 int) uint16 {
    p0 = p0
    switch (p0 << ((975) & 63)) {
    case 751:
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 int16 = 126
            var v8 uint = 3
            v9 := GenS1{F0: 0, F1: -65, F2: uint64(v8)}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
    default:
        p0 -= p0
    }
    fmt.Println(p0)
    _ = p0
    return 65535
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    var v4 uint = 1
    v5 := make(map[uint32]string)
    v1 += v1
    for fi10 := 0; fi10 < 6; fi10++ {
        if false {
            for fi11 := 0; fi11 < 15; fi11++ {
                v1 = 127
                var v12 int32 = -1
                var v13 int64 = -1
                _ = fi11
                _ = v12
                _ = v13
            }
        } else {
            v14 := v2[(("foo bar" + "hello") + ("\n\t" + ""))]
            _ = v14
        }
        _ = fi10
    }
    v15 := GenS1{F0: -11, F1: 0, F2: 4294967295}
    var v16 int64 = 0
    v17 := v3["\n\t"]
    v4 = v4
    fmt.Println(v1, v1, v17)
    var v18 uint32 = 2
    var v19 int16 = 42
    for fi20 := 0; fi20 < 3; fi20++ {
        switch "a" {
        case "foo bar":
            for fi21 := 0; fi21 < 18; fi21++ {
                v1 += (v1 * uint8(v18))
                if ((v19 >> ((v4) & 63)) <= (v19 % ((0) | 1))) {
                    v22 := GenS1{F0: 0, F1: 100, F2: 4294967295}
                    v16 = (v16 | -9223372036854775808)
                    v1 += (v1 << ((0) & 63))
                    var v23 int8 = 101
                    if false {
                        v5[2147483648] = v17
                        var v24 int = (fi20 ^ int(v23))
                        var v25 float32 = ((1e10 + float32((float64(0.0) / ((0.1 * -2.25) + float64(1e10))))) / 1.0)
                        _ = v24
                        _ = v25
                    }
                    v3[(v17 + v17)] = v17
                    var v26 float64 = (1.5 - float64(float32(1e10)))
                    _ = v22
                    _ = v23
                    _ = v26
                }
                v2[""] = "foo bar"
                var v27 bool = true
                switch fi21 {
                case 128:
                    fmt.Println(fi21)
                case 4294967295:
                    fmt.Println(fi21, v1, v19)
                case 127:
                    var v28 int64 = (^(v16))
                    _ = v28
                }
                _ = fi21
                _ = v27
            }
        case "\n\t":
            if (v17 != v17) {
                var v29 uint8 = v1
                _ = v29
            } else {
                if false {
                    v18 = (v18 * (v18 | (v18 - (v18 % ((2147483647) | 1)))))
                    v19 += int16(v1)
                    var v30 uint16 = 1
                    _ = v30
                }
                switch v17 {
                case "foo bar":
                    v31 := []int{(fi20 * fi20), fi20, (fi20 & (^(fi20))), fi20}
                    _ = v31
                default:
                    v32 := GenS1{F0: 127, F1: int8(v19), F2: 0}
                    _ = v32
                }
                v33 := []int{fi20}
                _ = v33
            }
        case "a":
            var v34 int32 = int32(v1)
            _ = v34
        }
        _ = fi20
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        v36 := GenS1{F0: 127, F1: -1, F2: 18446744073709551615}
        for fi37 := 0; fi37 < 16; fi37++ {
            fmt.Println(v17)
            fmt.Println(v16, v19)
            if true {
                v4 += (^(v4))
            }
            _ = fi37
        }
        _ = fi35
        _ = v36
    }
    v17 = v17
    v4 = v4
    switch v17 {
    case "\n\t":
        v38 := []int{}
        _ = v38
    }
    var v39 uint32 = (v18 - (v18 % ((v18) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
