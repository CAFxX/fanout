package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 string }
type GenS2 struct { F0 uint32 }
func genFn0() int64 {
    v5 := []int{0, 403, 1, 254}
    var v6 uint8 = 233
    v6 = (v6 * (v6 ^ (v6 * v6)))
    if len(v5) > 0 {
        v5[0] = 0
    }
    var v7 float32 = 0.0
    _ = v5
    _ = v6
    _ = v7
    return 426
}

func main() {
    var v1 int16 = 0
    var v2 uint64 = 0
    v3 := make(map[uint32]bool)
    v4 := make(map[int]int)
    fmt.Println(v1)
    _ = len(v4)
    switch "" {
    case "\n\t":
        fmt.Println(v1, v2)
    }
    var v8 uint = 4294967295
    switch "\n\t" {
    case "a":
        var v9 int32 = 2147483647
        _ = v9
    case "\n\t":
        var v10 float64 = ((float64(float32(1.5)) + 0.1) - -2.25)
        _ = v10
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 float64 = 1e10
        _ = fi11
        _ = v12
    }
    _ = len(v4)
    for fi13 := 0; fi13 < 7; fi13++ {
        v1 = (v1 >> ((v8) & 63))
        v14 := []int{0, fi13}
        v15 := GenS1{F0: 127, F1: v2, F2: "hello"}
        _ = fi13
        _ = v14
        _ = v15
    }
    switch -1 {
    case 4294967295:
        for fi16 := 0; fi16 < 14; fi16++ {
            for fi17 := 0; fi17 < 20; fi17++ {
                if (!(!((1 >= 42) && (-2147483648 != 2147483647)))) {
                    var v18 uint8 = uint8(v2)
                    _ = v18
                }
                v19 := GenS1{F0: 192, F1: 42, F2: "a"}
                fmt.Println(v8, fi16, fi16)
                if true {
                    v20 := GenS1{F0: 1, F1: 0, F2: ("hello" + "\n\t")}
                    if true {
                        if false {
                            if ((1e10 / ((3.14159 + 0.0) - 1e10)) < (0.0 - 0.1)) {
                                v21 := []int{(^(fi17))}
                                var v22 int16 = -32768
                                var v23 int32 = -1
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                        } else {
                            var v24 float64 = -2.25
                            _ = v24
                        }
                    } else {
                        fmt.Println(v8)
                        var v25 string = ("\n\t" + "a")
                        _ = v25
                    }
                    _ = v20
                }
                v8 -= 255
                var v26 uint16 = 0
                v27 := GenS1{F0: 192, F1: (v2 - v2), F2: ""}
                _ = fi17
                _ = v19
                _ = v26
                _ = v27
            }
            _ = fi16
        }
    default:
        for fi28 := 0; fi28 < 4; fi28++ {
            v2 += v2
            if (true || (-2147483648 > -457)) {
                v29 := []int{fi28, (fi28 & (-(fi28))), (fi28 | fi28)}
                switch "\n\t" {
                case "a":
                    var v30 int = (^(fi28))
                    _ = v30
                case "hello":
                    v8 += uint(v2)
                default:
                    for fi31 := 0; fi31 < 20; fi31++ {
                        _ = len(v29)
                        _ = fi31
                    }
                }
                v2 += (v2 % ((v2) | 1))
                _ = v29
            }
            var v32 bool = (120 != -1)
            _ = fi28
            _ = v32
        }
    }
    v1 += (v1 >> ((v8) & 63))
    for fi33 := 0; fi33 < 9; fi33++ {
        if false {
            var v34 uint8 = 205
            for fi35 := 0; fi35 < 9; fi35++ {
                _ = len(v3)
                var v36 bool = ((true && (2147483647 <= -744)) || false)
                _ = fi35
                _ = v36
            }
            fmt.Println(v8, v2)
            _ = v34
        }
        _ = fi33
    }
    for fi37 := 0; fi37 < 6; fi37++ {
        v38 := GenS2{F0: 4294967295}
        var v39 uint16 = 32767
        v39 = (v39 & 32767)
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 uint16 = 2
    v41 := GenS1{F0: 42, F1: v2, F2: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v8
}
