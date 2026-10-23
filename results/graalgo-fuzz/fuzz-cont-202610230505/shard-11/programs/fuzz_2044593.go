package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int32) string {
    if false {
        var v5 bool = (2 < 18446744073709551615)
        if true {
            for fi6 := 0; fi6 < 2; fi6++ {
                v7 := GenS1{F0: (false || (-128 < 42)), F1: -660, F2: 1}
                v8 := GenS1{F0: (v5 && ((false && (!(((v5 && false) && (!(1 > 1))) || (v5 && v5)))) && ((0.1 * 1e10) > ((((0.1 / (1e10 + (float32(0.0) - ((-2.25 * 3.14159) * -2.25)))) + -2.25) - 3.14159) - (1.5 / 0.1))))), F1: -9223372036854775808, F2: 65535}
                _ = fi6
                _ = v7
                _ = v8
            }
            var v9 float32 = (((float32((1e10 * float64(-2.25))) * 1e10) - ((0.0 + 1e10) / 0.1)) - 0.0)
            _ = v9
        } else {
            p0 = (p0 & p0)
        }
        _ = v5
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        if true {
            p0 = -726
        }
        switch (fi10 * 729) {
        case 32767:
            p0 += (p0 % (((p0 << ((42) & 63))) | 1))
        default:
            v11 := []int{}
            _ = v11
        }
        if (5 == 127) {
            v12 := []int{(fi10 + fi10), fi10}
            fmt.Println(fi10, fi10)
            _ = v12
        } else {
            if ((1e10 * -2.25) < ((1.5 / float64((-2.25 + float32(float64(((-2.25 / (3.14159 + (1e10 + 0.1))) / (3.14159 - 1e10))))))) / 1.5)) {
                if (217 >= 127) {
                    if true {
                        var v13 uint8 = 2
                        _ = v13
                    }
                } else {
                    var v14 uint64 = 18446744073709551615
                    _ = v14
                }
                var v15 float64 = 0.1
                _ = v15
            }
            var v16 int32 = p0
            _ = v16
        }
        _ = fi10
    }
    if false {
        var v17 int8 = 42
        _ = v17
    }
    fmt.Println(p0)
    _ = p0
    return "a"
}

func main() {
    var v1 int32 = -697
    var v2 int64 = -555
    v3 := make(map[uint32]uint32)
    v4 := []uint32{267, 2147483647}
    v1 = (v1 % (((v1 >> ((18446744073709551615) & 63))) | 1))
    fmt.Println(v1, v2)
    if false {
        v1 = 294
        if true {
            for fi18 := 0; fi18 < 5; fi18++ {
                switch "\n\t" {
                case "a":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                var v19 uint64 = 0
                _ = fi18
                _ = v19
            }
            _ = len(v3)
            for fi20 := 0; fi20 < 14; fi20++ {
                if true {
                    v1 = 0
                    v21 := GenS1{F0: true, F1: v2, F2: 2147483647}
                    _ = v21
                }
                _ = fi20
            }
        }
    } else {
        v22 := GenS2{F0: 42}
        var v23 string = ("hello" + ("hello" + ""))
        fmt.Println(v1, v2, v2)
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        v25 := GenS2{F0: 255}
        _ = fi24
        _ = v25
    }
    v1 = -2147483648
    var v26 uint32 = 4294967295
    for fi27 := 0; fi27 < 12; fi27++ {
        v26 += 889
        for fi28 := 0; fi28 < 20; fi28++ {
            v26 -= 634
            if false {
                var v29 uint32 = v26
                _ = v29
            } else {
                _ = len(v3)
            }
            if ((v1 * 990) != (v1 ^ (v1 + (^(v1))))) {
                v26 += (^(v26))
            }
            _ = fi28
        }
        var v30 int32 = (v1 / ((v1) | 1))
        _ = fi27
        _ = v30
    }
    var v31 string = "a"
    var v32 float64 = 0.1
    for fi33 := 0; fi33 < 20; fi33++ {
        v32 -= v32
        _ = fi33
    }
    var v34 uint64 = 1
    var v35 uint64 = v34
    if (false && true) {
        if (!(!true)) {
            v1 -= v1
        }
        if ((true && (4294967295 < 18446744073709551615)) && (32767 > 777)) {
            var v36 uint64 = v35
            _ = v36
        } else {
            switch -716 {
            case 128:
                var v37 uint = 42
                _ = v37
            }
        }
        switch (v31 + ("" + v31)) {
        case "":
            v26 = (v26 % (((v26 * v26)) | 1))
        }
    } else {
        var v38 float64 = (v32 * (-2.25 / float64(float32(v32))))
        _ = v38
    }
    v1 = 684
    var v39 uint32 = 1
    v2 = v2
    fmt.Println(v31, v39)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v39
    _ = v4
}
