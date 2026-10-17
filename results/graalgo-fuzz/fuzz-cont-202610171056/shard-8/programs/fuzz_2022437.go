package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint8 }

func main() {
    var v1 uint32 = 65535
    var v2 float64 = -2.25
    v3 := make(map[uint32]string)
    var v4 uint32 = v1
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 uint8 = 255
        for fi7 := 0; fi7 < 5; fi7++ {
            if true {
                v8 := GenS1{F0: 1, F1: -1}
                _ = v8
            } else {
                switch "a" {
                case "foo bar":
                    fmt.Println(fi7, v4)
                case "":
                    v9 := []int{fi7}
                    _ = v9
                case "a":
                    if (fi5 >= fi5) {
                        fmt.Println(fi5, v2)
                        _ = len(v3)
                    } else {
                        v6 = (^(v6))
                        delete(v3, v4)
                        v2 = (-2.25 - (((v2 + float64(0.0)) * 0.1) + ((v2 * (1e10 / 1e10)) * ((v2 / -2.25) + ((1e10 - ((1.5 * ((1e10 / v2) / v2)) + v2)) + (v2 * 1e10))))))
                    }
                }
                v10 := GenS1{F0: 65535, F1: 1}
                fmt.Println(v6)
                _ = v10
            }
            v3[(^(v4))] = "hello"
            _ = fi7
        }
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := GenS2{F0: v6, F1: 42, F2: 42}
            v3[(^(v1))] = "hello"
            _ = fi11
            _ = v12
        }
        _ = fi5
        _ = v6
    }
    if ((32767 == 32767) && (4294967295 > 2)) {
        var v13 string = "\n\t"
        _ = v13
    } else {
        v1 = v4
        v3[v4] = "\n\t"
        if (v2 <= -2.25) {
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 uint16 = 2
                switch fi14 {
                case 987:
                    var v16 bool = false
                    _ = v16
                case 127:
                    for fi17 := 0; fi17 < 13; fi17++ {
                        if (((0.1 / (-2.25 * 1.5)) * 1e10) == 1e10) {
                            if false {
                                if (true && (0 != 0)) {
                                    if (!false) {
                                        var v18 int8 = 100
                                        v15 = (v15 + (v15 >> ((65) & 63)))
                                        v1 += (^(v1))
                                        _ = v18
                                    }
                                    switch ("hello" + ("\n\t" + "\n\t")) {
                                    case "hello":
                                        v15 = v15
                                    case "foo bar":
                                        v1 = 0
                                    case "a":
                                        v19 := GenS2{F0: 128, F1: 284, F2: 42}
                                        _ = v19
                                    }
                                    v15 = 1
                                } else {
                                    var v20 bool = (0 < 32767)
                                    _ = v20
                                }
                                v21 := GenS3{F0: 1, F1: (^(v4))}
                                _ = v21
                            } else {
                                v22 := GenS4{F0: (v15 << ((2) & 63)), F1: 897, F2: 42}
                                _ = v22
                            }
                        } else {
                            v15 = (^(v15))
                        }
                        _ = fi17
                    }
                case 2147483647:
                    for fi23 := 0; fi23 < 6; fi23++ {
                        var v24 uint64 = 0
                        v2 = float64(3.14159)
                        _ = fi23
                        _ = v24
                    }
                }
                _ = fi14
                _ = v15
            }
        }
    }
    var v25 uint32 = v1
    if (484 < 588) {
        var v26 string = "foo bar"
        _ = v26
    }
    if ((1 > 792) || (v1 != v25)) {
        v27 := GenS2{F0: 4, F1: 42, F2: 32767}
        _ = v27
    } else {
        var v28 string = ("\n\t" + "\n\t")
        v29 := []int{127, -1, 2147483647, 32768}
        _ = v28
        _ = v29
    }
    v2 = ((v2 * (v2 + (v2 / (v2 * v2)))) / v2)
    if true {
        var v30 uint8 = 2
        _ = v30
    } else {
        v31 := []int{}
        v2 += 0.0
        _ = v31
    }
    v32 := v3[(v4 % ((65535) | 1))]
    v3[v1] = "foo bar"
    fmt.Println(v25)
    v32 = ("a" + ((v32 + ("a" + "foo bar")) + v32))
    if (!(true || (!(42 < 127)))) {
        fmt.Println(v32, v4, v25)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v4
}
