package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint16, p1 int16, p2 int32) int16 {
    switch "" {
    case "":
        var v5 uint64 = 911
        _ = v5
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int = fi6
        var v8 uint16 = (p0 + (p0 + 628))
        _ = fi6
        _ = v7
        _ = v8
    }
    if (p0 == 1) {
        var v9 string = "a"
        v10 := []int{128, -9223372036854775808}
        for fi11 := 0; fi11 < 16; fi11++ {
            if true {
                var v12 float32 = (1.5 - 0.0)
                _ = v12
            }
            p1 += p1
            _ = fi11
        }
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 * p1)
}
func genFn1() uint {
    v13 := []int{4294967295}
    if (0 != -289) {
        var v14 int32 = -339
        _ = v14
    }
    var v15 uint8 = 77
    v15 = 1
    _ = v13
    _ = v15
    return 255
}

func main() {
    var v1 uint16 = 32767
    var v2 int64 = 9223372036854775807
    var v3 int = 32767
    var v4 uint32 = 2147483648
    v3 = -2147483648
    v2 = 0
    v4 = v4
    v2 = (v2 * (v2 - (v2 + v2)))
    for fi16 := 0; fi16 < 20; fi16++ {
        v1 = v1
        v17 := GenS1{F0: ((("\n\t" + "\n\t") + "hello") + ("foo bar" + "a"))}
        v3 -= (^(fi16))
        _ = fi16
        _ = v17
    }
    fmt.Println(v3, v2)
    if true {
        for fi18 := 0; fi18 < 5; fi18++ {
            for fi19 := 0; fi19 < 4; fi19++ {
                v20 := GenS2{F0: 2, F1: 1}
                _ = fi19
                _ = v20
            }
            v4 -= (^(v4))
            switch (^(fi18)) {
            case 128:
                switch fi18 {
                case -133:
                    switch ("" + "") {
                    case "foo bar":
                        switch (v3 ^ int(v4)) {
                        case -1:
                            if true {
                                var v21 uint64 = 473
                                fmt.Println(v4, v3, v21)
                                v1 += (v1 % (((v1 / ((v1) | 1))) | 1))
                                _ = v21
                            }
                        }
                    }
                default:
                    var v22 uint32 = (v4 * 530)
                    _ = v22
                }
            case 4294967295:
                v23 := GenS3{F0: 1, F1: v1}
                _ = v23
            case 32768:
                var v24 int = v3
                _ = v24
            }
            _ = fi18
        }
        var v25 float64 = float64((3.14159 / ((0.1 - 3.14159) * float32(1e10))))
        _ = v25
    } else {
        var v26 int32 = 1
        fmt.Println(v2)
        _ = v26
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        for fi28 := 0; fi28 < 5; fi28++ {
            v29 := []int{fi27, (fi27 << ((1) & 63)), (fi27 * fi27)}
            if true {
                var v30 int = (v3 - -9223372036854775808)
                _ = v30
            }
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    v2 = 42
    var v31 uint32 = v4
    v32 := GenS3{F0: 42, F1: v1}
    var v33 float32 = (1e10 - 0.1)
    for fi34 := 0; fi34 < 3; fi34++ {
        if (float64(0.1) == (-2.25 / -2.25)) {
            var v35 float64 = ((((1.5 - 0.0) + -2.25) * 0.1) / (0.1 / 1.5))
            var v36 int8 = 42
            var v37 int32 = int32(v36)
            _ = v35
            _ = v36
            _ = v37
        } else {
            v3 = (fi34 & -9223372036854775808)
            v3 = fi34
        }
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
