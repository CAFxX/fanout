package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint64 }
type GenS2 struct { F0 int8; F1 int32; F2 int32 }
type GenS3 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint8, p1 uint8, p2 uint) int {
    p0 -= p0
    p2 += p2
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 float32 = ((1.5 * (1.5 - (0.1 - (0.1 * (3.14159 + 1e10))))) / 1.5)
        p0 = p1
        _ = fi4
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return 207
}
func genFn1(p0 uint, p1 int8, p2 int32) int16 {
    var v6 string = ""
    if (4294967295 != 214) {
        if true {
            if (-2.25 <= ((0.0 * 1.5) / 3.14159)) {
                v6 = v6
                v7 := []int{-1, genFn0(107, 179, (^(p0))), -709}
                _ = v7
            } else {
                fmt.Println(p2, v6)
                var v8 int64 = 1
                fmt.Println(p2, p1)
                _ = v8
            }
            var v9 uint16 = 0
            if ((v6 + "hello") < (v6 + v6)) {
                v10 := []int{genFn0(95, 255, (p0 >> ((p0) & 63)))}
                v9 = v9
                _ = v10
            } else {
                fmt.Println(p2, p2, p0)
                p2 += p2
                if ((((1.5 / (1.5 * 3.14159)) * (1e10 / (3.14159 - 0.0))) / 1.0) < (1e10 * (1e10 - (1e10 - 0.1)))) {
                    v9 = 32767
                    p0 += 4294967295
                } else {
                    p2 = -337
                    fmt.Println(v6, p1, v9)
                    p1 = 1
                }
            }
            _ = v9
        }
    }
    fmt.Println(v6, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 764
}

func main() {
    var v1 uint64 = 1
    var v2 int64 = -1
    var v3 uint = 65535
    var v11 bool = ((v2 ^ v2) > -999)
    var v12 int16 = 32767
    v12 = genFn1(255, int8(v12), 1)
    for fi13 := 0; fi13 < 5; fi13++ {
        for fi14 := 0; fi14 < 17; fi14++ {
            if (v11 || true) {
                v15 := GenS1{F0: v12, F1: (fi13 >> ((883) & 63)), F2: 42}
                _ = v15
            }
            var v16 bool = true
            var v17 int64 = 955
            _ = fi14
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    fmt.Println(v1)
    fmt.Println(v3, v11)
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 int64 = 5
        fmt.Println(v19)
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        v21 := []int{fi20, (^(fi20)), fi20, (^(fi20))}
        v3 = v3
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        v23 := []int{}
        v12 = v12
        _ = fi22
        _ = v23
    }
    if ((-2.25 >= 1.5) || v11) {
        var v24 float64 = float64(1e10)
        _ = v24
    } else {
        if (3.14159 <= -2.25) {
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 float64 = 1e10
                if (!((v12 * (v12 | v12)) >= v12)) {
                    v3 += 2
                    fmt.Println(v1)
                }
                _ = fi25
                _ = v26
            }
        } else {
            if (!("\n\t" > "")) {
                v3 = (v3 - 18446744073709551615)
                if false {
                    var v27 int64 = 0
                    _ = v27
                } else {
                    v28 := GenS2{F0: -128, F1: -1, F2: -1}
                    _ = v28
                }
            } else {
                v29 := GenS2{F0: 45, F1: 0, F2: 674}
                v30 := []int{127, 4294967295, -11, 4294967295}
                v31 := GenS3{F0: (v12 + (^(v12))), F1: v1}
                _ = v29
                _ = v30
                _ = v31
            }
            for fi32 := 0; fi32 < 2; fi32++ {
                v33 := GenS2{F0: -1, F1: -476, F2: int32(v3)}
                _ = fi32
                _ = v33
            }
            v34 := []int{32768, genFn0(2, 222, (v3 << ((18446744073709551615) & 63))), genFn0(42, 2, v3), 128}
            _ = v34
        }
    }
    fmt.Println(v3)
    v12 = v12
    var v35 string = ("" + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v35
}
