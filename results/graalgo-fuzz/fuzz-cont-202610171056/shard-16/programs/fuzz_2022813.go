package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 string; F2 uint32 }
type GenS4 struct { F0 int16; F1 int32; F2 string }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 uint64; F1 string }
func genFn0() uint16 {
    v5 := GenS1{F0: 20}
    var v6 int8 = 127
    var v7 int = 0
    v8 := GenS2{F0: -128}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint = 1
    v2 := []string{"", ("hello" + ""), ""}
    var v3 int = -79
    var v4 float64 = 1.5
    var v9 uint32 = 971
    var v10 int16 = -691
    var v11 int16 = v10
    var v12 bool = false
    _ = len(v2)
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v11 -= 0
    } else {
        v9 -= (v9 + (v9 & v9))
        v11 = v10
        switch 4294967295 {
        case 42:
            fmt.Println(v9, v1)
        default:
            if v12 {
                v4 += (v4 / (((0.0 + v4) - 0.1) * 0.0))
                v12 = (-9223372036854775808 < 1)
                v13 := GenS3{F0: 42, F1: ("\n\t" + (("" + "a") + "")), F2: v9}
                _ = v13
            }
        }
    }
    var v14 float64 = v4
    for fi15 := 0; fi15 < 5; fi15++ {
        v11 = (v10 << (((v1 << (((^(v1))) & 63))) & 63))
        var v16 uint32 = v9
        var v17 float32 = (3.14159 - 1e10)
        _ = fi15
        _ = v16
        _ = v17
    }
    switch (v3 & (v3 * (v3 & (-(v3))))) {
    case -184:
        if ((v9 < (v9 | v9)) && (v12 && ("foo bar" < "\n\t"))) {
            if v12 {
                switch 127 {
                case -909:
                    fmt.Println(v4)
                case -9223372036854775808:
                    if ((v14 - ((float64(1e10) / 3.14159) - 0.0)) > float64((3.14159 / (3.14159 / 1.0)))) {
                        var v18 float32 = ((1.5 + 0.0) / -2.25)
                        if (false && ((v12 || (!(0 > 42))) && (v14 != v4))) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            for fi19 := 0; fi19 < 13; fi19++ {
                                v20 := GenS4{F0: (-(v10)), F1: -2147483648, F2: ("\n\t" + "foo bar")}
                                v3 += 681
                                _ = fi19
                                _ = v20
                            }
                        } else {
                            v18 += v18
                            v21 := GenS3{F0: -2147483648, F1: "", F2: v9}
                            fmt.Println(v18)
                            _ = v21
                        }
                        var v22 bool = ((869 <= v9) || (32768 <= (v3 / ((v3) | 1))))
                        _ = v18
                        _ = v22
                    } else {
                        v23 := GenS5{F0: ((230 > 128) || (!((!((v9 != (v9 ^ (v9 * 4294967295))) || (v1 == (v1 - v1)))) && true))), F1: (v3 | v3)}
                        for fi24 := 0; fi24 < 20; fi24++ {
                            if v12 {
                                v4 = (v14 / (v4 / (v4 * (0.0 / v4))))
                            } else {
                                switch fi24 {
                                case 32767:
                                    if v12 {
                                        v10 += v10
                                        v9 += (^(v9))
                                        var v25 uint32 = v9
                                        _ = v25
                                    } else {
                                        for fi26 := 0; fi26 < 11; fi26++ {
                                            var v27 int = (fi24 << ((42) & 63))
                                            _ = fi26
                                            _ = v27
                                        }
                                    }
                                }
                                switch "a" {
                                case "hello":
                                    v11 = (v10 - v11)
                                }
                                if (!((-(v3)) < (^(v3)))) {
                                    if ((v9 < 42) && (0 < genFn0())) {
                                        var v28 float64 = 1.5
                                        v29 := GenS6{F0: 42, F1: (v9 | 65535)}
                                        switch v3 {
                                        case -55:
                                            v12 = (1e10 < v28)
                                        case 1:
                                            v30 := GenS2{F0: 100}
                                            _ = v30
                                        }
                                        _ = v28
                                        _ = v29
                                    } else {
                                        var v31 uint32 = (v9 & (v9 & (^(v9))))
                                        _ = v31
                                    }
                                }
                            }
                            v2 = append(v2, "")
                            _ = fi24
                        }
                        _ = v23
                    }
                case 4294967295:
                    v9 = 2147483647
                }
                v2 = append(v2, ("\n\t" + ""))
            } else {
                if (v12 || ((v3 + int(v11)) == v3)) {
                    v12 = v12
                }
                v3 = 4294967295
                v1 += 1
            }
            fmt.Println(v3, v14, v12)
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v1)
        }
    }
    if len(v2) > 0 {
        v2[0] = "a"
    }
    switch "\n\t" {
    case "hello":
        v1 = v1
    }
    v32 := GenS7{F0: 0, F1: "foo bar"}
    v33 := GenS6{F0: 2147483648, F1: v9}
    var v34 bool = (float32(float64(3.14159)) <= 0.0)
    var v35 uint = 255
    v36 := GenS1{F0: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v9
}
