package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int64; F1 uint16; F2 int32 }
func genFn0() int32 {
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := []int{}
        v8 := GenS1{F0: 0, F1: 0}
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS1{F0: 42, F1: 15}
            var v11 uint32 = 298
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v12 int = -189
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := []int{}
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 string = ""
            _ = fi15
            _ = v16
        }
        v17 := GenS1{F0: 310, F1: 551}
        _ = fi13
        _ = v14
        _ = v17
    }
    v18 := []int{v12, (v12 >> ((0) & 63)), v12, (v12 * v12)}
    _ = v12
    _ = v18
    return 1
}
func genFn1() int16 {
    v19 := []int{0, 688}
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := GenS2{F0: "\n\t", F1: 1}
        for fi22 := 0; fi22 < 3; fi22++ {
            if len(v19) > 0 {
                v19[0] = 2147483647
            }
            _ = fi22
        }
        _ = fi20
        _ = v21
    }
    _ = v19
    return 32767
}

func main() {
    var v1 uint32 = 421
    var v2 int16 = -1
    var v3 uint32 = 89
    v4 := []string{}
    var v5 int16 = -1
    v2 = (v5 + v2)
    if false {
        v4 = append(v4, ("a" + ("\n\t" + "a")))
    } else {
        v5 = (v2 * (v2 % ((v5) | 1)))
        v4 = append(v4, ("hello" + "\n\t"))
        for fi23 := 0; fi23 < 19; fi23++ {
            var v24 bool = true
            _ = fi23
            _ = v24
        }
    }
    if (42 != -67) {
        var v25 int64 = 537
        if false {
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
            var v26 float64 = 1e10
            for fi27 := 0; fi27 < 4; fi27++ {
                fmt.Println(v1, v26)
                v28 := GenS1{F0: v25, F1: 65535}
                _ = fi27
                _ = v28
            }
            _ = v26
        } else {
            v1 = v1
        }
        var v29 float32 = 3.14159
        _ = v25
        _ = v29
    } else {
        var v30 float32 = (-2.25 - float32(float64(3.14159)))
        if ((122 <= 0) && (18446744073709551615 == 0)) {
            var v31 uint = 558
            if true {
                var v32 string = (("\n\t" + ("hello" + "")) + ("\n\t" + "hello"))
                if false {
                    if true {
                        fmt.Println(v32)
                    }
                    var v33 int32 = -1
                    _ = v33
                }
                var v34 float64 = ((1e10 / -2.25) - 3.14159)
                _ = v32
                _ = v34
            }
            _ = v31
        }
        _ = v30
    }
    var v35 uint64 = 18446744073709551615
    var v36 float64 = (3.14159 - ((0.0 * 0.0) - 0.1))
    if (812 != 4294967295) {
        v35 += (v35 * 4294967295)
        v2 -= v5
        if true {
            switch "" {
            case "\n\t":
                switch -55 {
                case 604:
                    var v37 int8 = 100
                    _ = v37
                default:
                    v1 += v3
                }
            case "a":
                v2 = (v2 >> ((42) & 63))
            case "foo bar":
                fmt.Println(v1, v35)
            }
            if ((v35 << ((255) & 63)) > (v35 | (v35 % ((v35) | 1)))) {
                _ = len(v4)
            }
        }
    } else {
        fmt.Println(v3, v1)
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        v35 = v35
    }
    v5 = v2
    v38 := GenS3{F0: 984, F1: 32767, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v38
    _ = v4
    _ = v5
}
