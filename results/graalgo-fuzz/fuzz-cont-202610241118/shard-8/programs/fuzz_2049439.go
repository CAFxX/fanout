package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int8 }

func main() {
    v1 := []int{-1, 42, -619}
    var v2 int16 = 32767
    var v3 string = (("a" + "hello") + "hello")
    var v4 uint32 = 782
    var v5 uint64 = 1
    v6 := []uint32{v4, 42, v4}
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            v2 -= -32768
            v4 = (v4 / ((v4) | 1))
            _ = fi8
        }
        _ = fi7
    }
    var v9 float64 = ((3.14159 * ((float64(-2.25) + 1.5) + (1e10 - float64(float32(0.0))))) * 1.5)
    v5 -= v5
    switch "hello" {
    case "":
        if false {
            var v10 int64 = 1
            var v11 uint = 461
            var v12 int8 = 42
            _ = v10
            _ = v11
            _ = v12
        }
    }
    var v13 float64 = v9
    fmt.Println(v13)
    v1 = append(v1, -2147483648)
    var v14 float32 = (-2.25 * 1e10)
    v3 = (v3 + v3)
    fmt.Println(v5)
    fmt.Println(v14, v5)
    v15 := GenS1{F0: int16(v4)}
    v9 += 1.5
    if (false || (!((false && false) && (58 >= 0)))) {
        for fi16 := 0; fi16 < 19; fi16++ {
            v17 := GenS1{F0: (v2 ^ 1000)}
            var v18 uint32 = 1
            if false {
                for fi19 := 0; fi19 < 12; fi19++ {
                    var v20 int8 = 42
                    v6 = append(v6, v4)
                    var v21 float64 = 3.14159
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                for fi22 := 0; fi22 < 20; fi22++ {
                    fmt.Println(v18, fi22, fi22)
                    fmt.Println(v2, v18, v13)
                    _ = fi22
                }
                if (true && true) {
                    _ = len(v6)
                    var v23 float32 = 0.0
                    var v24 string = "\n\t"
                    _ = v23
                    _ = v24
                } else {
                    var v25 uint16 = 255
                    v6 = append(v6, 2147483648)
                    _ = v25
                }
            } else {
                if ((v3 + v3) == v3) {
                    var v26 int = (fi16 >> ((1) & 63))
                    switch (v3 + (("" + ("" + v3)) + "foo bar")) {
                    case "a":
                        v5 += v5
                    case "\n\t":
                        for fi27 := 0; fi27 < 8; fi27++ {
                            var v28 float64 = 3.14159
                            _ = fi27
                            _ = v28
                        }
                    default:
                        _ = len(v1)
                    }
                    _ = v26
                } else {
                    if ((true || false) || (!(true || true))) {
                        v13 = (v13 / v13)
                        for fi29 := 0; fi29 < 10; fi29++ {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            fmt.Println(v5, v9)
                            if len(v6) > 0 {
                                v6[0] = v4
                            }
                            _ = fi29
                        }
                        if true {
                            var v30 int16 = (v2 << ((109) & 63))
                            _ = v30
                        } else {
                            for fi31 := 0; fi31 < 4; fi31++ {
                                if (!(457 == 2)) {
                                    var v32 uint8 = 67
                                    v32 = (v32 << ((1) & 63))
                                    if true {
                                        var v33 string = "hello"
                                        if true {
                                            var v34 float64 = v9
                                            switch v3 {
                                            case "hello":
                                                v35 := GenS1{F0: (-(v2))}
                                                _ = v35
                                            }
                                            _ = v34
                                        } else {
                                            v36 := GenS2{F0: 255, F1: v4, F2: -1}
                                            var v37 uint8 = v32
                                            _ = v36
                                            _ = v37
                                        }
                                        _ = v33
                                    } else {
                                        v2 -= (v2 & v2)
                                        var v38 uint32 = uint32(v32)
                                        v39 := GenS1{F0: (v2 >> ((65535) & 63))}
                                        _ = v38
                                        _ = v39
                                    }
                                    _ = v32
                                }
                                _ = fi31
                            }
                            v3 = v3
                            v40 := GenS1{F0: int16(v18)}
                            _ = v40
                        }
                    }
                    v3 = (("hello" + v3) + "")
                    v2 = v2
                }
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        v4 += 1
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
