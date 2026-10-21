package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int; F1 bool; F2 uint8 }
type GenS4 struct { F0 int; F1 int32; F2 int8 }
type GenS5 struct { F0 int32; F1 string; F2 bool }
func genFn0(p0 int16) string {
    var v7 int32 = -2147483648
    v8 := GenS1{F0: -2147483648}
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    _ = v7
    _ = v8
    return "a"
}

func main() {
    var v1 uint8 = 9
    var v2 bool = false
    v3 := []string{"", "", "\n\t", ("\n\t" + "hello")}
    var v4 float32 = (0.1 - 0.0)
    var v5 uint16 = 42
    v6 := make(map[uint32]uint32)
    v2 = ((uint64(v5) == 18446744073709551615) || (744 == 2147483648))
    if (((float64(3.14159) * (((((3.14159 + float64((0.1 / (float32((0.1 * 1.5)) + 3.14159)))) * (1e10 * -2.25)) * (float64(float32(3.14159)) - 0.0)) / -2.25) - 0.0)) - -2.25) != 1e10) {
        fmt.Println(v4)
        var v9 int16 = 32767
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        _ = v9
    }
    if v2 {
        if (128 < 128) {
            var v10 string = "foo bar"
            _ = v10
        }
        v11 := GenS1{F0: 0}
        for fi12 := 0; fi12 < 5; fi12++ {
            v2 = false
            var v13 string = "\n\t"
            var v14 int = fi12
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v11
    } else {
        var v15 string = "foo bar"
        fmt.Println(v4, v5)
        if v2 {
            if v2 {
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 int32 = 0
                    v1 -= (^(v1))
                    _ = fi16
                    _ = v17
                }
                v1 = (^(v1))
            } else {
                v1 = uint8(v5)
                v1 -= v1
                v18 := GenS2{F0: 2147483648, F1: v1}
                _ = v18
            }
        } else {
            for fi19 := 0; fi19 < 11; fi19++ {
                fmt.Println(v4)
                if (int(v1) != (^(fi19))) {
                    switch fi19 {
                    case 4294967295:
                        v20 := GenS3{F0: (fi19 ^ fi19), F1: (60 >= 1), F2: (v1 / ((v1) | 1))}
                        _ = v20
                    }
                    v21 := GenS4{F0: (fi19 % ((fi19) | 1)), F1: -412, F2: 127}
                    _ = v21
                } else {
                    var v22 uint32 = 805
                    for fi23 := 0; fi23 < 3; fi23++ {
                        if ((v5 + (v5 - v5)) > (v5 - (v5 - (v5 >> ((255) & 63))))) {
                            v24 := GenS5{F0: int32(v5), F1: "", F2: (91 != -52)}
                            _ = v24
                        } else {
                            if (!(v2 || (32767 <= -32768))) {
                                v2 = true
                                var v25 int64 = -644
                                if false {
                                    var v26 int = fi19
                                    _ = v26
                                }
                                _ = v25
                            }
                        }
                        v2 = v2
                        _ = fi23
                    }
                    var v27 uint8 = 255
                    _ = v22
                    _ = v27
                }
                _ = fi19
            }
            v3 = append(v3, v15)
        }
        _ = v15
    }
    if (32767 <= (v5 + 32767)) {
        v1 = v1
        v5 += (v5 - 2)
        fmt.Println(v2)
    }
    fmt.Println(v5)
    var v28 bool = (!((0 >= 791) && (!false)))
    v4 -= (float32(-2.25) / ((1e10 / (float32(float64(((((1.5 / 1.0) * (v4 - v4)) + v4) * v4))) + (((v4 - (1e10 - v4)) / (1.5 / (0.0 / v4))) + float32(3.14159)))) * ((v4 + (3.14159 + (float32(-2.25) / float32(0.1)))) * 1.5)))
    var v29 bool = v28
    var v30 uint64 = 0
    switch ("foo bar" + "a") {
    case "":
        var v31 int = -534
        _ = v31
    case "foo bar":
        v29 = (!(!v29))
    case "\n\t":
        for fi32 := 0; fi32 < 20; fi32++ {
            for fi33 := 0; fi33 < 1; fi33++ {
                v34 := GenS4{F0: 399, F1: -1, F2: -79}
                v3 = append(v3, "a")
                _ = fi33
                _ = v34
            }
            if v29 {
                if (-1 <= 1) {
                    if v29 {
                        var v35 uint8 = v1
                        v2 = (fi32 == int(v30))
                        _ = v35
                    }
                } else {
                    var v36 int8 = 42
                    if len(v3) > 0 {
                        v3[0] = (("" + "\n\t") + "\n\t")
                    }
                    _ = v36
                }
            }
            v28 = v29
            _ = fi32
        }
    }
    var v37 int8 = -1
    v38 := v6[2147483647]
    for fi39 := 0; fi39 < 18; fi39++ {
        v1 += v1
        var v40 bool = v29
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 6; fi41++ {
        if v28 {
            var v42 int = 2147483647
            _ = v42
        }
        v38 = v38
        if v29 {
            if len(v3) > 0 {
                v3[0] = ""
            }
        }
        _ = fi41
    }
    fmt.Println(v4, v29, v29)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
