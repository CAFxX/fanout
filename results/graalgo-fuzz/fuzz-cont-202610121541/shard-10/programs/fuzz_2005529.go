package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 float64 = float64((float32(1.5) + 3.14159))
    v2 := make(map[int]uint32)
    var v3 string = "hello"
    v4 := make(map[int]int)
    v5 := make(map[string]int)
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 int = fi6
        v1 -= v1
        if true {
            var v8 int8 = 56
            if (2147483647 <= -1) {
                fmt.Println(v3)
                switch (^(v7)) {
                case -530:
                    var v9 bool = ((0 == 2) || (v1 <= (v1 + (v1 / v1))))
                    _ = v9
                case 516:
                    for fi10 := 0; fi10 < 10; fi10++ {
                        v8 = v8
                        v3 = ("hello" + "")
                        var v11 uint8 = 1
                        var v12 int = (-(v7))
                        v11 = (^(v11))
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                }
                for fi13 := 0; fi13 < 9; fi13++ {
                    v14 := []int{v7, 32767, (v7 >> ((148) & 63))}
                    if (1 >= -1) {
                        if false {
                            v15 := GenS1{F0: -1, F1: 255}
                            _ = v15
                        } else {
                            if true {
                                fmt.Println(fi13)
                                var v16 int = (-(v7))
                                v17 := GenS1{F0: 42, F1: 160}
                                _ = v16
                                _ = v17
                            }
                            v18 := v2[fi13]
                            v19 := GenS2{F0: 1}
                            var v20 int8 = v8
                            if ((-2147483648 >= 2147483647) || true) {
                                if (!(32767 >= 2)) {
                                    v3 = (("" + v3) + (v3 + v3))
                                    fmt.Println(v20, v20)
                                    switch -360 {
                                    case -847:
                                        v21 := GenS2{F0: 255}
                                        _ = v21
                                    case 32768:
                                        v22 := GenS2{F0: 42}
                                        _ = v22
                                    case 9223372036854775807:
                                        v1 += (1e10 - (v1 / v1))
                                    }
                                }
                            }
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        var v23 int16 = 933
                        _ = v23
                    } else {
                        var v24 int32 = 0
                        v7 = int(v24)
                        var v25 uint32 = 4294967295
                        _ = v24
                        _ = v25
                    }
                    var v26 uint16 = 2
                    _ = fi13
                    _ = v14
                    _ = v26
                }
            }
            _ = v8
        } else {
            for fi27 := 0; fi27 < 3; fi27++ {
                v28 := []int{(fi6 - fi27), (fi27 >> ((0) & 63)), (fi27 - (fi6 ^ fi6))}
                _ = fi27
                _ = v28
            }
        }
        _ = fi6
        _ = v7
    }
    delete(v2, 32767)
    v29 := GenS1{F0: 1, F1: 1}
    _ = len(v5)
    var v30 int64 = 0
    if (true && ((1e10 - 3.14159) != 3.14159)) {
        delete(v5, v3)
        for fi31 := 0; fi31 < 15; fi31++ {
            v32 := GenS1{F0: (v30 | (v30 | (v30 % (((v30 % ((v30) | 1))) | 1)))), F1: 127}
            v33 := []int{(fi31 + (fi31 + fi31)), fi31, int(v30)}
            _ = len(v5)
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    v1 = 1.5
    v30 += v30
    v34 := GenS3{F0: -91}
    if (!false) {
        for fi35 := 0; fi35 < 10; fi35++ {
            _ = len(v5)
            _ = fi35
        }
        for fi36 := 0; fi36 < 8; fi36++ {
            var v37 int8 = int8(v30)
            _ = fi36
            _ = v37
        }
        v30 -= v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
}
