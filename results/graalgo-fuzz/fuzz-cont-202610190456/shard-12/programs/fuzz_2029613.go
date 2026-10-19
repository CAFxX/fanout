package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint16 }
func genFn0() uint64 {
    var v7 int16 = -485
    fmt.Println(v7)
    var v8 uint8 = 2
    _ = v7
    _ = v8
    return 503
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 1
    v3 := make(map[uint32]uint32)
    var v4 int8 = 127
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]uint32)
    switch -2147483648 {
    case 860:
        var v9 uint8 = v2
        _ = v9
    case 2147483647:
        for fi10 := 0; fi10 < 11; fi10++ {
            fmt.Println(v4, v2)
            for fi11 := 0; fi11 < 14; fi11++ {
                fmt.Println(v2, fi11, v2)
                v12 := []int{fi11, (fi11 + (fi10 + fi10)), (fi10 ^ (fi11 ^ (fi11 * 300)))}
                var v13 string = "foo bar"
                _ = fi11
                _ = v12
                _ = v13
            }
            fmt.Println(fi10)
            _ = fi10
        }
    default:
        if false {
            v14 := GenS1{F0: false, F1: (v2 | (v2 % (((v2 & (v2 % ((108) | 1)))) | 1))), F2: 65535}
            if true {
                fmt.Println(v2, v4)
                var v15 uint8 = v2
                _ = v15
            } else {
                if (false && true) {
                    v2 -= v2
                } else {
                    for fi16 := 0; fi16 < 2; fi16++ {
                        v17 := GenS1{F0: (2147483647 <= -851), F1: v2, F2: 1}
                        if false {
                            v18 := GenS1{F0: true, F1: 1, F2: 585}
                            var v19 uint16 = 198
                            var v20 float64 = -2.25
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        _ = fi16
                        _ = v17
                    }
                }
                var v21 float64 = (-2.25 - 3.14159)
                if (2 >= 585) {
                    if false {
                        fmt.Println(v2)
                        fmt.Println(v21)
                    } else {
                        for fi22 := 0; fi22 < 15; fi22++ {
                            var v23 int = fi22
                            if false {
                                v24 := GenS1{F0: false, F1: v2, F2: 32767}
                                var v25 uint16 = 42
                                v4 = v4
                                _ = v24
                                _ = v25
                            }
                            for fi26 := 0; fi26 < 3; fi26++ {
                                if (1 > genFn0()) {
                                    v4 = 100
                                    fmt.Println(v4, v4)
                                    v23 -= v23
                                    v27 := []int{fi22, (-(v23))}
                                    switch ("" + ("\n\t" + "")) {
                                    case "":
                                        var v28 uint32 = 589
                                        _ = v28
                                    default:
                                        v29 := GenS1{F0: ("hello" > "foo bar"), F1: v2, F2: 255}
                                        _ = v29
                                    }
                                    _ = len(v5)
                                    _ = v27
                                }
                                v2 += 0
                                _ = fi26
                            }
                            _ = fi22
                            _ = v23
                        }
                        for fi30 := 0; fi30 < 6; fi30++ {
                            v21 -= v21
                            _ = fi30
                        }
                    }
                    v21 = v21
                } else {
                    var v31 uint = 42
                    fmt.Println(v31, v31)
                    if (0.1 != 1.5) {
                        v2 -= (v2 & v2)
                    } else {
                        v32 := GenS1{F0: true, F1: (v2 << (((v31 / (((v31 & v31)) | 1))) & 63)), F2: 802}
                        v2 = (v2 << ((255) & 63))
                        _ = v32
                    }
                    _ = v31
                }
                _ = v21
            }
            _ = v14
        }
    }
    var v33 uint16 = 42
    switch 32768 {
    case -280:
        var v34 uint8 = v2
        _ = v34
    case -1:
        v35 := GenS1{F0: false, F1: 1, F2: (^(v33))}
        _ = v35
    case -521:
        var v36 int16 = -1
        _ = v36
    }
    v37 := []int{42, -1, 632}
    var v38 string = "hello"
    var v39 int32 = 0
    var v40 int16 = 32767
    if (1 < 42) {
        if len(v37) >= 2 {
            v37 = v37[1:len(v37)]
        }
        var v41 uint64 = 0
        _ = v41
    }
    var v42 int32 = 0
    for fi43 := 0; fi43 < 16; fi43++ {
        for fi44 := 0; fi44 < 1; fi44++ {
            var v45 uint8 = 192
            fmt.Println(v4)
            _ = fi44
            _ = v45
        }
        for fi46 := 0; fi46 < 4; fi46++ {
            var v47 float32 = 0.1
            _ = fi46
            _ = v47
        }
        if len(v37) > 0 {
            v37[0] = fi43
        }
        _ = fi43
    }
    var v48 string = ("a" + ((((v38 + v38) + "a") + v38) + v38))
    if (-128 > 127) {
        var v49 uint32 = 42
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v48
    _ = v5
    _ = v6
}
