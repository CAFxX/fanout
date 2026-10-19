package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int8; F2 bool }
type GenS3 struct { F0 int8 }
func genFn0(p0 int16) int64 {
    p0 -= -32768
    var v7 uint16 = 42
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int16 = 42
        p0 = p0
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v7
    return 1
}

func main() {
    var v1 string = ("" + "hello")
    var v2 int8 = -1
    var v3 uint64 = 42
    var v4 uint32 = uint32(v3)
    var v5 uint16 = 65535
    var v6 float64 = float64(3.14159)
    var v10 uint16 = v5
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS1{F0: (v3 * (v3 | (v3 * (v3 + v3))))}
        v13 := GenS1{F0: v3}
        fmt.Println(v5, v6, fi11)
        _ = fi11
        _ = v12
        _ = v13
    }
    v3 = (v3 * 347)
    fmt.Println(v1, v3, v1)
    v14 := GenS2{F0: 519, F1: v2, F2: (v4 >= (v4 >> ((990) & 63)))}
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 uint64 = 42
        switch fi15 {
        case 42:
            v17 := GenS2{F0: 42, F1: (v2 / (((v2 & -78)) | 1)), F2: false}
            _ = v17
        default:
            if false {
                var v18 int8 = (v2 >> ((65535) & 63))
                for fi19 := 0; fi19 < 4; fi19++ {
                    switch v1 {
                    case "\n\t":
                        v20 := []int{fi19}
                        _ = v20
                    case "":
                        fmt.Println(fi19, v18, v6)
                    case "a":
                        if true {
                            v21 := []int{fi15, int(v4), 2147483647}
                            _ = v21
                        }
                    default:
                        v22 := []int{(fi15 & fi15)}
                        _ = v22
                    }
                    v6 += (v6 * 1e10)
                    _ = fi19
                }
                _ = v18
            } else {
                v1 = v1
                var v23 uint = uint(v5)
                var v24 string = (v1 + v1)
                _ = v23
                _ = v24
            }
        }
        if true {
            fmt.Println(v10, v1)
            if (942 != 682) {
                fmt.Println(v16, v2)
            } else {
                var v25 int32 = -574
                _ = v25
            }
        }
        _ = fi15
        _ = v16
    }
    if (!(139 == 128)) {
        v6 += float64((0.0 / (1e10 * (0.0 + 3.14159))))
        fmt.Println(v10, v2, v4)
        for fi26 := 0; fi26 < 10; fi26++ {
            fmt.Println(v4)
            switch 32767 {
            case 127:
                v5 -= (v10 % ((v10) | 1))
            default:
                fmt.Println(v3, v4, v10)
            }
            if (307 >= 65535) {
                if (!((v3 - 244) < v3)) {
                    switch v1 {
                    case "\n\t":
                        switch (-(fi26)) {
                        case 128:
                            var v27 float64 = (v6 / (v6 / v6))
                            _ = v27
                        case 2147483647:
                            switch (v1 + "a") {
                            case "hello":
                                v2 -= v2
                            case "foo bar":
                                v10 = (^(v10))
                            case "\n\t":
                                if true {
                                    v28 := []int{fi26, (fi26 & (fi26 >> ((uint(v3)) & 63))), fi26}
                                    _ = v28
                                } else {
                                    v4 = (v4 + v4)
                                    v29 := []int{(fi26 - fi26), (^(fi26))}
                                    var v30 int64 = genFn0(243)
                                    _ = v29
                                    _ = v30
                                }
                            default:
                                v1 = (v1 + (v1 + v1))
                            }
                        case 4294967295:
                            for fi31 := 0; fi31 < 18; fi31++ {
                                v32 := []int{(fi31 & (fi31 - fi26)), 495, fi31}
                                _ = fi31
                                _ = v32
                            }
                        default:
                            v33 := []int{}
                            _ = v33
                        }
                    case "a":
                        var v34 uint16 = (v10 / ((v10) | 1))
                        _ = v34
                    default:
                        v35 := GenS1{F0: v3}
                        _ = v35
                    }
                }
            }
            _ = fi26
        }
    } else {
        var v36 float32 = (0.0 * (0.1 * (float32(1.5) - float32(((v6 / v6) * -2.25)))))
        v10 = v10
        _ = v36
    }
    switch v1 {
    case "\n\t":
        var v37 float64 = (((1e10 * 0.1) / 1.5) - v6)
        _ = v37
    case "hello":
        v38 := GenS1{F0: 1}
        _ = v38
    case "a":
        fmt.Println(v3, v1, v2)
    }
    var v39 float64 = (v6 / v6)
    v3 -= (^(v3))
    v40 := GenS3{F0: v2}
    for fi41 := 0; fi41 < 2; fi41++ {
        var v42 int8 = (^(v2))
        for fi43 := 0; fi43 < 16; fi43++ {
            var v44 int8 = v42
            if false {
                v2 -= 127
                v10 = 179
                var v45 float32 = 1e10
                _ = v45
            }
            v46 := []int{(^(fi41)), -719, 420, fi43}
            v46 = append(v46, fi41)
            _ = fi43
            _ = v44
            _ = v46
        }
        _ = fi41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
