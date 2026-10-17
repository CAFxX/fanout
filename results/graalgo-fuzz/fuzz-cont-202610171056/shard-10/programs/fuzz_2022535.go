package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS2 struct { F0 string; F1 int; F2 int16 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint }
type GenS4 struct { F0 int32 }

func main() {
    var v1 uint16 = 32767
    var v2 int16 = 1000
    v3 := make(map[int]bool)
    var v4 float32 = 1e10
    var v5 float64 = ((1.5 - float64((((float32(float64((0.1 / ((1.5 - (float32(3.14159) - v4)) + v4)))) + (1e10 * 0.1)) + (((v4 + 1e10) - v4) * ((v4 * ((v4 + v4) - (v4 + 0.0))) - float32(3.14159)))) * float32(3.14159)))) - ((3.14159 / (float64((-2.25 / 3.14159)) - (0.0 * 1.5))) + float64(((v4 / (((v4 * 0.0) / v4) * float32(float64(v4)))) * (float32(3.14159) / ((-2.25 / (3.14159 / v4)) - 0.1))))))
    var v6 uint16 = v1
    v7 := GenS1{F0: v6, F1: 2147483647, F2: (v6 << ((255) & 63))}
    var v8 int32 = 2147483647
    v9 := []int{0, 822}
    if (-1 >= 4294967295) {
        var v10 uint16 = 42
        v11 := GenS1{F0: 32767, F1: 797, F2: (v1 * (^(v1)))}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = v10
        _ = v11
    }
    var v12 uint64 = 18446744073709551615
    v13 := GenS2{F0: "a", F1: -953, F2: 0}
    if false {
        for fi14 := 0; fi14 < 19; fi14++ {
            fmt.Println(v5)
            var v15 uint64 = v12
            _ = fi14
            _ = v15
        }
    } else {
        v9 = append(v9, 1)
        var v16 float64 = 0.1
        v17 := GenS1{F0: 32767, F1: (v8 ^ v8), F2: (v1 * v6)}
        _ = v16
        _ = v17
    }
    switch -808 {
    case 128:
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 bool = (true && true)
            if v19 {
                var v20 int32 = v8
                for fi21 := 0; fi21 < 20; fi21++ {
                    v9 = append(v9, int(v6))
                    _ = fi21
                }
                for fi22 := 0; fi22 < 15; fi22++ {
                    var v23 bool = ((-270 >= int32(v2)) && (false && true))
                    switch -256 {
                    case 543:
                        v2 -= v2
                    case 128:
                        if (-814 <= 1) {
                            v23 = (!(42 == 153))
                            v24 := GenS1{F0: v1, F1: (v8 + v8), F2: v6}
                            switch 9223372036854775807 {
                            case 4294967295:
                                v25 := GenS1{F0: 955, F1: v20, F2: v1}
                                _ = v25
                            case 1:
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                            case 790:
                                var v26 int8 = -1
                                _ = v26
                            }
                            _ = v24
                        } else {
                            _ = len(v9)
                        }
                    case -2147483648:
                        if (("hello" + ("" + "foo bar")) == ("a" + "")) {
                            var v27 int32 = v20
                            if v19 {
                                if true {
                                    fmt.Println(v2)
                                    v28 := GenS2{F0: "", F1: (fi22 % ((1) | 1)), F2: v2}
                                    _ = v28
                                }
                                switch "hello" {
                                case "foo bar":
                                    var v29 bool = false
                                    _ = v29
                                default:
                                    var v30 int8 = -1
                                    _ = v30
                                }
                                v1 = (v6 - 255)
                            }
                            _ = v27
                        }
                    }
                    _ = fi22
                    _ = v23
                }
                _ = v20
            } else {
                for fi31 := 0; fi31 < 10; fi31++ {
                    var v32 float32 = v4
                    var v33 uint64 = 1
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                v8 = (v8 ^ int32(v6))
                if (!v19) {
                    v2 = (v2 % ((v2) | 1))
                    var v34 int = fi18
                    _ = v34
                } else {
                    _ = len(v9)
                    fmt.Println(fi18)
                    _ = len(v3)
                }
            }
            _ = fi18
            _ = v19
        }
    default:
        switch 127 {
        case -711:
            v5 = (float64(((v4 / (3.14159 * (0.0 + ((v4 * v4) / v4)))) - v4)) / ((v5 + v5) / -2.25))
        case 121:
            switch -9223372036854775808 {
            case -138:
                v5 = float64((v4 / (1.5 / (1e10 * (v4 + 3.14159)))))
            case 2147483647:
                var v35 uint = 255
                _ = v35
            case -127:
                v1 = (^(v6))
            default:
                v36 := GenS3{F0: v6, F1: 42, F2: 4294967295}
                _ = v36
            }
        case 0:
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
    }
    if false {
        switch 127 {
        case 32768:
            v3[-2147483648] = (1e10 >= ((v5 - v5) - v5))
        }
    } else {
        v37 := GenS4{F0: -1}
        var v38 uint = 65535
        if false {
            v39 := GenS1{F0: 54, F1: (-(v8)), F2: (v1 - v1)}
            _ = v39
        }
        _ = v37
        _ = v38
    }
    if (!(65535 != 255)) {
        if ((1.5 + v5) != (v5 / 1e10)) {
            switch int(v12) {
            case 128:
                v40 := GenS4{F0: v8}
                _ = v40
            default:
                var v41 uint8 = 237
                _ = v41
            }
            v6 = (v1 - v1)
            var v42 int = 1
            _ = v42
        } else {
            var v43 bool = ((694 >= -9223372036854775808) || (0 > -896))
            _ = v43
        }
        fmt.Println(v1, v2)
    }
    var v44 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
