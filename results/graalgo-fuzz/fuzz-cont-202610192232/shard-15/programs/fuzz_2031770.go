package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 string }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 uint32 }
func genFn0() int64 {
    v4 := GenS1{F0: 1, F1: 0, F2: "foo bar"}
    v5 := GenS1{F0: 42, F1: 134, F2: ("" + "\n\t")}
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 uint = 1
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]uint32)
    var v6 int64 = genFn0()
    var v7 uint8 = 0
    var v8 string = "foo bar"
    v9 := GenS1{F0: -120, F1: 1, F2: (((v8 + (v8 + (v8 + "hello"))) + "") + ("foo bar" + "a"))}
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := []int{(fi10 << (((v1 | v1)) & 63)), fi10, (fi10 - 127), 0}
        var v12 uint16 = 799
        _ = fi10
        _ = v11
        _ = v12
    }
    v3[32768] = 493
    for fi13 := 0; fi13 < 10; fi13++ {
        if true {
            v7 = (v7 % ((uint8(v1)) | 1))
        } else {
            switch (v8 + ("hello" + ("foo bar" + ("\n\t" + v8)))) {
            case "foo bar":
                var v14 uint16 = 2
                _ = v14
            case "hello":
                v8 = (v8 + (v8 + ("a" + v8)))
            case "a":
                v15 := []int{fi13}
                _ = v15
            }
            var v16 int8 = 31
            _ = v16
        }
        v2 += int64(v1)
        switch (^(fi13)) {
        case -981:
            if (2147483648 == 1) {
                if (!(!(-2.25 == (1e10 + 3.14159)))) {
                    if (false && true) {
                        switch "a" {
                        case "hello":
                            v17 := GenS2{F0: 735, F1: v8}
                            _ = v17
                        default:
                            v18 := GenS3{F0: 2147483648}
                            _ = v18
                        }
                        for fi19 := 0; fi19 < 19; fi19++ {
                            switch v8 {
                            case "a":
                                v20 := []int{fi13}
                                v1 += (v1 & v1)
                                var v21 float64 = (0.1 * (0.1 * -2.25))
                                _ = v20
                                _ = v21
                            case "hello":
                                v8 = v8
                            }
                            if ((395 < 42) && (0.0 != float32(0.0))) {
                                v22 := GenS4{F0: 97, F1: (v7 / (((v7 << ((v1) & 63))) | 1)), F2: 341}
                                _ = v22
                            }
                            _ = fi19
                        }
                        for fi23 := 0; fi23 < 20; fi23++ {
                            switch "" {
                            case "":
                                v24 := GenS4{F0: -1, F1: (v7 + 150), F2: 459}
                                _ = v24
                            case "foo bar":
                                v25 := GenS3{F0: 4294967295}
                                _ = v25
                            case "hello":
                                v2 -= genFn0()
                            }
                            fmt.Println(v7)
                            v2 = genFn0()
                            _ = fi23
                        }
                    } else {
                        if (42 > 764) {
                            fmt.Println(v1, v1)
                        }
                    }
                    v26 := GenS1{F0: -1, F1: 42, F2: v8}
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 4; fi27++ {
                        v1 -= 2
                        _ = fi27
                    }
                    var v28 int64 = (v2 << (((v1 + (v1 << ((255) & 63)))) & 63))
                    v1 = v1
                    _ = v28
                }
            }
        case 32768:
            v6 = v2
        case 713:
            if false {
                switch 1 {
                case 547:
                    v29 := GenS2{F0: (^(v1)), F1: "foo bar"}
                    _ = v29
                default:
                    v30 := GenS1{F0: -91, F1: -2, F2: ("hello" + v8)}
                    _ = v30
                }
                v31 := []int{32767, 4294967295}
                v32 := GenS5{F0: (fi13 | 875), F1: "hello"}
                _ = v31
                _ = v32
            } else {
                if true {
                    var v33 int16 = -1
                    fmt.Println(v6)
                    var v34 int64 = (v6 * v2)
                    _ = v33
                    _ = v34
                }
                v7 += (v7 - v7)
                var v35 int = (fi13 % ((fi13) | 1))
                _ = v35
            }
        default:
            fmt.Println(v7, v2)
        }
        _ = fi13
    }
    if ((!false) && true) {
        for fi36 := 0; fi36 < 3; fi36++ {
            if (true && (!(32767 != -1))) {
                v37 := GenS6{F0: 396}
                if false {
                    var v38 int = (fi36 - fi36)
                    _ = v38
                } else {
                    switch -539 {
                    case 128:
                        var v39 int16 = -32768
                        _ = v39
                    default:
                        v40 := GenS2{F0: 2, F1: v8}
                        _ = v40
                    }
                    v41 := []int{(fi36 & (fi36 % ((fi36) | 1))), 966, fi36}
                    v42 := GenS6{F0: 0}
                    _ = v41
                    _ = v42
                }
                _ = v37
            }
            v2 = genFn0()
            v43 := []int{fi36, fi36}
            _ = fi36
            _ = v43
        }
        _ = len(v3)
    } else {
        v8 = "foo bar"
        if ((0.1 - (0.0 * 3.14159)) <= -2.25) {
            if (!true) {
                v1 = (v1 % ((42) | 1))
                var v44 uint8 = 1
                if false {
                    var v45 uint16 = 1
                    _ = v45
                }
                _ = v44
            }
        }
        v7 = (v7 ^ 127)
    }
    v46 := GenS5{F0: 127, F1: v8}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v46
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
