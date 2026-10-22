package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int16; F2 int }
type GenS3 struct { F0 uint; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS6 struct { F0 uint }
func genFn0() int64 {
    var v7 float64 = 0.1
    v8 := []int{}
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS1{F0: 0}
        for fi11 := 0; fi11 < 1; fi11++ {
            v7 = ((v7 * v7) - 0.1)
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    var v12 float32 = 0.1
    var v13 uint = 255
    _ = v12
    _ = v13
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 1e10
    var v3 int = 4294967295
    var v4 uint8 = 2
    var v5 int16 = 566
    v6 := make(map[string]bool)
    var v14 uint8 = (v4 >> ((255) & 63))
    v15 := GenS2{F0: ("\n\t" + ""), F1: 712, F2: 42}
    var v16 string = "hello"
    if (!(false && (-120 != 127))) {
        v17 := []int{-146}
        var v18 uint8 = (v14 + v4)
        for fi19 := 0; fi19 < 9; fi19++ {
            for fi20 := 0; fi20 < 19; fi20++ {
                switch (v3 / ((32767) | 1)) {
                case -225:
                    if false {
                        var v21 int = v3
                        v14 += v14
                        _ = v21
                    }
                case -9223372036854775808:
                    if true {
                        switch (fi19 - (^(v3))) {
                        case 9:
                            var v22 string = v16
                            _ = v22
                        case -16:
                            v16 = ((((v16 + ("hello" + "")) + "a") + (v16 + "foo bar")) + "hello")
                        case 42:
                            var v23 uint16 = 0
                            _ = v23
                        default:
                            v24 := GenS3{F0: 781, F1: 199, F2: v18}
                            _ = v24
                        }
                    }
                case -216:
                    v5 -= (v5 % ((v5) | 1))
                }
                var v25 int16 = (v5 - (v5 / (((v5 + 32767)) | 1)))
                var v26 float32 = (0.1 + v2)
                _ = fi20
                _ = v25
                _ = v26
            }
            _ = fi19
        }
        _ = v17
        _ = v18
    }
    var v27 uint = 42
    v28 := GenS4{F0: 311, F1: v5}
    v29 := v6["hello"]
    v30 := GenS5{F0: 417, F1: 2147483647, F2: 0}
    v31 := GenS2{F0: v16, F1: v5, F2: 148}
    v27 += v27
    fmt.Println(v5, v29, v4)
    switch "foo bar" {
    case "hello":
        var v32 bool = (32767 >= 561)
        _ = v32
    case "foo bar":
        v33 := []int{(v3 << ((2) & 63)), (v3 + v3), 128, (v3 + (v3 / ((v3) | 1)))}
        _ = v33
    case "":
        v34 := GenS1{F0: (v27 / ((1) | 1))}
        _ = v34
    default:
        for fi35 := 0; fi35 < 1; fi35++ {
            v36 := GenS4{F0: 0, F1: (v5 + v5)}
            for fi37 := 0; fi37 < 9; fi37++ {
                fmt.Println(v3)
                v38 := []int{fi35, int(v5)}
                v39 := GenS6{F0: (^(v27))}
                _ = fi37
                _ = v38
                _ = v39
            }
            v4 -= (v4 >> ((v27) & 63))
            _ = fi35
            _ = v36
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
