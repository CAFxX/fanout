package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() bool {
    v5 := []int{1, 128, 9223372036854775807, 32768}
    var v6 int32 = -363
    _ = len(v5)
    switch ("foo bar" + (("foo bar" + "\n\t") + "\n\t")) {
    case "":
        if true {
            v7 := GenS1{F0: (("" + "hello") + ("" + ("hello" + ((("foo bar" + "a") + "foo bar") + ("foo bar" + ("foo bar" + ("hello" + "")))))))}
            v6 = 0
            fmt.Println(v6)
            _ = v7
        }
    case "a":
        var v8 float32 = -2.25
        _ = v8
    case "foo bar":
        var v9 int = 535
        _ = v9
    default:
        var v10 int = 530
        _ = v10
    }
    v6 = (v6 << ((42) & 63))
    _ = v5
    _ = v6
    return false
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[uint32]string)
    var v3 string = ("foo bar" + ("\n\t" + "\n\t"))
    v4 := make(map[uint32]uint32)
    if (-20 == 42) {
        v3 = v3
        if (2 == 812) {
            for fi11 := 0; fi11 < 17; fi11++ {
                v3 = ("a" + v3)
                v12 := GenS1{F0: v3}
                var v13 uint = 65535
                _ = fi11
                _ = v12
                _ = v13
            }
            v14 := GenS1{F0: v3}
            _ = v14
        } else {
            v3 = (v3 + "a")
            if (!false) {
                var v15 string = ((v3 + v3) + ((v3 + ((v3 + v3) + v3)) + v3))
                _ = v15
            }
        }
        v1 -= 127
    } else {
        if true {
            v3 = ("foo bar" + (v3 + ("" + v3)))
            if (960 != 42) {
                fmt.Println(v3)
                var v16 bool = true
                if false {
                    var v17 float32 = float32((float64(3.14159) + (((3.14159 / ((((-2.25 + 1e10) + float64(1.5)) - 0.1) + float64(((1.5 + 0.1) * 1.5)))) + ((float64(float32((1.5 * (3.14159 + 1.5)))) * 0.1) - 3.14159)) * -2.25)))
                    var v18 uint16 = 675
                    fmt.Println(v17, v17)
                    _ = v17
                    _ = v18
                }
                _ = v16
            }
        }
        fmt.Println(v3, v3)
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        for fi20 := 0; fi20 < 20; fi20++ {
            fmt.Println(v1, v3)
            v1 -= (v1 - v1)
            _ = fi20
        }
        _ = fi19
    }
    v21 := v4[795]
    v1 = v1
    var v22 int64 = 392
    v23 := []int{-1, 0, 42}
    fmt.Println(v22)
    switch -137 {
    case 1:
        v23 = append(v23, 32767)
    }
    delete(v2, 563)
    if true {
        if (32767 == 0) {
            v24 := GenS1{F0: (v3 + v3)}
            _ = v24
        } else {
            if false {
                var v25 int16 = 42
                _ = v25
            } else {
                for fi26 := 0; fi26 < 4; fi26++ {
                    var v27 uint8 = v1
                    _ = fi26
                    _ = v27
                }
                v28 := GenS1{F0: v3}
                _ = v28
            }
            var v29 float64 = 0.1
            _ = v29
        }
    }
    fmt.Println(v1, v3)
    var v30 uint8 = (v1 % (((v1 ^ 229)) | 1))
    for fi31 := 0; fi31 < 8; fi31++ {
        v1 = v1
        var v32 uint8 = (^(v30))
        if (0 > -1) {
            if false {
                for fi33 := 0; fi33 < 8; fi33++ {
                    switch int(v22) {
                    case 127:
                        v30 += v1
                    case -2147483648:
                        switch v3 {
                        case "":
                            if false {
                                fmt.Println(v3, v22)
                                fmt.Println(v1)
                            } else {
                                v32 = (v30 / ((v30) | 1))
                                if len(v23) >= 2 {
                                    v23 = v23[1:len(v23)]
                                }
                                v22 -= v22
                            }
                        default:
                            switch fi31 {
                            case 820:
                                if true {
                                    fmt.Println(v1, fi33)
                                    var v34 int64 = (v22 << ((0) & 63))
                                    _ = v34
                                } else {
                                    var v35 uint = 1
                                    v32 -= (^(v1))
                                    if len(v23) > 0 {
                                        v23[0] = fi33
                                    }
                                    _ = v35
                                }
                            case 0:
                                v2[(v21 ^ v21)] = (((v3 + v3) + v3) + ("\n\t" + v3))
                            }
                        }
                    default:
                        v30 = 172
                    }
                    fmt.Println(v32, v1, v22)
                    _ = fi33
                }
            }
        } else {
            fmt.Println(v22, v30, v32)
            var v36 float64 = ((0.0 * ((0.0 + 3.14159) / (1e10 - -2.25))) - 1e10)
            if true {
                v4[v21] = v21
            }
            _ = v36
        }
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v4
}
