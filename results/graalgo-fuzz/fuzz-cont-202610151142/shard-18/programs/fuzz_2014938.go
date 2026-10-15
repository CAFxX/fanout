package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int64 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 bool; F1 uint32 }

func main() {
    var v1 uint16 = 32767
    var v2 float32 = 3.14159
    var v3 int16 = 1000
    var v4 float32 = 0.0
    if (42 <= 1) {
        v1 += v1
        for fi5 := 0; fi5 < 4; fi5++ {
            if (!false) {
                for fi6 := 0; fi6 < 19; fi6++ {
                    v1 += (v1 * v1)
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v2, v4)
                    case "\n\t":
                        if (v3 > v3) {
                            v7 := GenS1{F0: 228, F1: false, F2: -840}
                            v8 := []int{127}
                            _ = v7
                            _ = v8
                        } else {
                            fmt.Println(v4)
                            switch fi5 {
                            case 32767:
                                v4 += -2.25
                                v1 = (^(v1))
                                var v9 int = 2147483647
                                _ = v9
                            }
                            v10 := []int{}
                            v1 = (v1 / ((42) | 1))
                            _ = v10
                        }
                    default:
                        v4 = v2
                    }
                    _ = fi6
                }
            }
            _ = fi5
        }
        if ((v3 >> ((1) & 63)) < (^(v3))) {
            if true {
                v2 = v2
                if false {
                    fmt.Println(v3)
                    if true {
                        var v11 int64 = -1
                        _ = v11
                    }
                } else {
                    var v12 int32 = -2147483648
                    _ = v12
                }
            }
            if false {
                fmt.Println(v2, v1, v2)
                for fi13 := 0; fi13 < 19; fi13++ {
                    if (!true) {
                        v14 := GenS2{F0: 0, F1: (v3 & int16(v1))}
                        var v15 float32 = (3.14159 * -2.25)
                        v16 := []int{fi13, fi13, (fi13 / ((190) | 1))}
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    fmt.Println(fi13, v3)
                    var v17 uint16 = (v1 % (((v1 << ((4294967295) & 63))) | 1))
                    _ = fi13
                    _ = v17
                }
                v4 = 3.14159
            }
        } else {
            if (!true) {
                var v18 int16 = (v3 | 1000)
                fmt.Println(v2, v18)
                _ = v18
            } else {
                for fi19 := 0; fi19 < 1; fi19++ {
                    if (-2147483648 != 1) {
                        for fi20 := 0; fi20 < 16; fi20++ {
                            v4 -= (float32(3.14159) / (v2 / v4))
                            fmt.Println(fi20, v3, v3)
                            _ = fi20
                        }
                        var v21 string = "foo bar"
                        _ = v21
                    }
                    var v22 string = ""
                    _ = fi19
                    _ = v22
                }
            }
            v3 += v3
        }
    }
    v2 += v4
    if (false || true) {
        switch "\n\t" {
        case "foo bar":
            var v23 int8 = 1
            _ = v23
        }
    }
    switch 1 {
    case 4294967295:
        v2 += 0.0
    case 0:
        var v24 uint32 = 2147483647
        _ = v24
    }
    v3 = 1
    var v25 float32 = 0.0
    v26 := GenS3{F0: (((1e10 * v2) - (((-2.25 / v4) + (((float32(-2.25) - 3.14159) * ((float32(1.5) + 1.5) + float32((0.0 / -2.25)))) * (3.14159 - 3.14159))) / (v4 + ((v2 + 3.14159) * 1e10)))) == (float32(3.14159) - (v25 / float32(-2.25)))), F1: 4294967295}
    fmt.Println(v3, v3)
    v2 -= v4
    v27 := GenS3{F0: true, F1: 661}
    switch 1 {
    case 1:
        for fi28 := 0; fi28 < 5; fi28++ {
            v29 := []int{}
            var v30 uint16 = (v1 << ((4294967295) & 63))
            _ = fi28
            _ = v29
            _ = v30
        }
    default:
        switch 1 {
        case 32767:
            v4 += v4
        }
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        var v32 uint = 331
        var v33 uint32 = 740
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 uint8 = 255
    var v35 int32 = 0
    fmt.Println(v3, v3)
    var v36 int32 = v35
    v3 -= (v3 ^ v3)
    if true {
        var v37 float64 = 1e10
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
