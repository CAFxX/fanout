package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 3.14159
    var v3 uint16 = 255
    if false {
        if true {
            for fi4 := 0; fi4 < 1; fi4++ {
                fmt.Println(fi4)
                v2 += (v2 * v2)
                v2 = 1e10
                _ = fi4
            }
            var v5 int16 = -973
            _ = v5
        }
        for fi6 := 0; fi6 < 8; fi6++ {
            v3 -= 723
            v7 := []int{-9223372036854775808}
            var v8 uint64 = 333
            _ = fi6
            _ = v7
            _ = v8
        }
        switch (("hello" + "hello") + "a") {
        case "":
            _ = len(v1)
        }
    } else {
        v3 = v3
        fmt.Println(v3, v2)
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: 1}
        _ = fi9
        _ = v10
    }
    v3 -= (v3 >> ((42) & 63))
    var v11 float64 = -2.25
    v12 := v1[uint32(v3)]
    switch 9223372036854775807 {
    case -863:
        switch "" {
        case "foo bar":
            v2 += v11
        case "":
            v13 := []int{}
            _ = v13
        default:
            for fi14 := 0; fi14 < 20; fi14++ {
                v15 := []int{fi14, (fi14 >> ((42) & 63)), (fi14 ^ 42)}
                for fi16 := 0; fi16 < 14; fi16++ {
                    v17 := GenS2{F0: 189}
                    if (true || false) {
                        var v18 string = "hello"
                        if false {
                            v19 := GenS2{F0: 43}
                            switch fi16 {
                            case -2147483648:
                                v20 := GenS2{F0: 255}
                                _ = v20
                            case 0:
                                var v21 float64 = (0.1 / v2)
                                _ = v21
                            case 42:
                                v3 = 32767
                            default:
                                v22 := GenS3{F0: 2147483647, F1: 127}
                                _ = v22
                            }
                            switch (v12 + ("\n\t" + v12)) {
                            case "foo bar":
                                var v23 bool = true
                                _ = v23
                            }
                            _ = v19
                        }
                        _ = v18
                    }
                    var v24 uint = 450
                    _ = fi16
                    _ = v17
                    _ = v24
                }
                var v25 int16 = -1
                _ = fi14
                _ = v15
                _ = v25
            }
        }
    default:
        if false {
            var v26 uint8 = 1
            for fi27 := 0; fi27 < 19; fi27++ {
                v28 := []int{}
                var v29 float32 = (-2.25 - ((1e10 + 1e10) - float32(v2)))
                _ = fi27
                _ = v28
                _ = v29
            }
            if false {
                v3 += v3
            }
            _ = v26
        } else {
            for fi30 := 0; fi30 < 18; fi30++ {
                v3 -= v3
                for fi31 := 0; fi31 < 16; fi31++ {
                    if false {
                        v12 = "foo bar"
                        v3 = 65535
                        var v32 int32 = -2147483648
                        _ = v32
                    }
                    switch (fi31 >> ((42) & 63)) {
                    case -2147483648:
                        fmt.Println(v12, fi31)
                    case -9223372036854775808:
                        var v33 bool = (394 < -2147483648)
                        _ = v33
                    }
                    var v34 int32 = -1
                    _ = fi31
                    _ = v34
                }
                for fi35 := 0; fi35 < 12; fi35++ {
                    v3 = (v3 * (v3 & (v3 & v3)))
                    _ = fi35
                }
                _ = fi30
            }
            var v36 float32 = (-2.25 + 3.14159)
            _ = v36
        }
    }
    if false {
        v37 := v1[1]
        fmt.Println(v11, v37, v37)
        var v38 float32 = float32(1e10)
        _ = v37
        _ = v38
    } else {
        fmt.Println(v2)
        fmt.Println(v12, v11, v12)
    }
    v39 := GenS2{F0: 128}
    fmt.Println(v2, v12, v2)
    v40 := []int{671, 4294967295, 9223372036854775807, 127}
    v41 := v1[uint32(v3)]
    v42 := GenS3{F0: -751, F1: 114}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v42
}
