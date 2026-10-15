package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int8; F2 uint }
type GenS4 struct { F0 string; F1 uint8 }
func genFn0(p0 uint32, p1 uint32, p2 int16) uint32 {
    fmt.Println(p1, p0, p2)
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := GenS1{F0: (!false), F1: (p2 << ((255) & 63)), F2: -32768}
            fmt.Println(fi6, p2, fi6)
            v8 := GenS2{F0: 18446744073709551615}
            _ = fi6
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 int = (fi9 >> ((0) & 63))
            var v11 int32 = 1
            _ = fi9
            _ = v10
            _ = v11
        }
        if false {
            p2 += (p2 ^ (-(p2)))
            if (998 < 42) {
                for fi12 := 0; fi12 < 2; fi12++ {
                    var v13 int64 = 528
                    _ = fi12
                    _ = v13
                }
                for fi14 := 0; fi14 < 18; fi14++ {
                    var v15 uint8 = 15
                    _ = fi14
                    _ = v15
                }
            } else {
                switch ("\n\t" + "a") {
                case "a":
                    v16 := []int{fi5}
                    _ = v16
                case "\n\t":
                    var v17 uint8 = 127
                    _ = v17
                }
            }
            if (!(1.5 >= 3.14159)) {
                p1 = 458
            }
        } else {
            if (-1 != 666) {
                var v18 int32 = 2147483647
                if (float64(0.1) > 3.14159) {
                    p2 += (-(p2))
                    v19 := []int{fi5, 4294967295, (^(fi5)), (-(fi5))}
                    _ = v19
                }
                _ = v18
            } else {
                var v20 int = fi5
                _ = v20
            }
        }
        _ = fi5
    }
    var v21 uint = 255
    var v22 float32 = 3.14159
    var v23 string = ("" + (("hello" + "foo bar") + "\n\t"))
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    _ = v23
    return (p1 + (p1 | 2))
}

func main() {
    var v1 uint64 = 742
    v2 := make(map[uint32]int)
    var v3 int64 = -1
    var v4 float64 = float64(float32((0.0 - -2.25)))
    var v24 uint64 = (v1 & v1)
    var v25 int = 841
    fmt.Println(v1, v4, v4)
    if (255 == 255) {
        var v26 uint32 = 1
        v24 = v24
        _ = v26
    } else {
        _ = len(v2)
        v27 := GenS1{F0: (false && true), F1: 1, F2: -57}
        _ = v27
    }
    v1 += v1
    delete(v2, 42)
    var v28 string = "foo bar"
    switch -383 {
    case 32767:
        v29 := GenS1{F0: false, F1: 42, F2: 1}
        _ = v29
    }
    if (v4 == (v4 / v4)) {
        for fi30 := 0; fi30 < 14; fi30++ {
            v4 = float64(0.0)
            if true {
                for fi31 := 0; fi31 < 9; fi31++ {
                    v32 := GenS3{F0: true, F1: -1, F2: 0}
                    if (-649 > 42) {
                        var v33 int32 = 42
                        v34 := GenS3{F0: (v3 > (v3 | v3)), F1: -68, F2: 18446744073709551615}
                        var v35 uint64 = (v24 | v1)
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                    v36 := GenS3{F0: false, F1: 127, F2: 4294967295}
                    _ = fi31
                    _ = v32
                    _ = v36
                }
                fmt.Println(fi30, v1)
            }
            v37 := GenS4{F0: ("foo bar" + ""), F1: 0}
            _ = fi30
            _ = v37
        }
        v24 += (v1 | v1)
    }
    v38 := []int{(v25 | int(v24)), (v25 - -9223372036854775808)}
    fmt.Println(v4, v25, v3)
    var v39 int64 = 930
    var v40 string = (v28 + v28)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
