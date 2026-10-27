package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 bool }
type GenS2 struct { F0 uint16; F1 uint8; F2 int }
type GenS3 struct { F0 string; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint; F1 uint; F2 int16 }

func main() {
    var v1 uint16 = 255
    var v2 int32 = 880
    v3 := make(map[uint32]bool)
    var v4 string = "hello"
    var v5 int16 = -1
    var v6 uint32 = 42
    var v7 float64 = 0.0
    var v8 uint16 = (v1 >> ((0) & 63))
    switch -9223372036854775808 {
    case -843:
        if false {
            v9 := []int{-459}
            for fi10 := 0; fi10 < 18; fi10++ {
                fmt.Println(v4, v5, v2)
                v11 := GenS1{F0: "foo bar", F1: v5, F2: ((((1.5 * -2.25) == (1e10 * -2.25)) || (!false)) && true)}
                _ = fi10
                _ = v11
            }
            _ = v9
        } else {
            v12 := GenS1{F0: ("hello" + v4), F1: v5, F2: true}
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int8 = -1
        for fi15 := 0; fi15 < 16; fi15++ {
            v5 = v5
            fmt.Println(v4)
            v16 := GenS1{F0: v4, F1: v5, F2: (v7 >= v7)}
            var v17 string = (v4 + v4)
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi13
        _ = v14
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        fmt.Println(v4)
        _ = fi18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v5 -= 1
        var v20 float64 = ((0.1 / 0.1) + (((3.14159 * (v7 - v7)) - v7) - 1e10))
        var v21 bool = false
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := GenS2{F0: (^(v8)), F1: 255, F2: 9223372036854775807}
    v7 += v7
    for fi23 := 0; fi23 < 12; fi23++ {
        fmt.Println(fi23, v6)
        _ = fi23
    }
    if false {
        var v24 uint32 = v6
        _ = v24
    } else {
        for fi25 := 0; fi25 < 6; fi25++ {
            for fi26 := 0; fi26 < 20; fi26++ {
                var v27 float32 = (-2.25 + float32((0.1 / 1.0)))
                v1 += 42
                _ = fi26
                _ = v27
            }
            if true {
                switch (fi25 % ((852) | 1)) {
                case -1:
                    for fi28 := 0; fi28 < 1; fi28++ {
                        if (42 == 104) {
                            var v29 int32 = 237
                            if false {
                                v30 := GenS1{F0: v4, F1: v5, F2: (-1 == (fi28 << ((65535) & 63)))}
                                v31 := GenS1{F0: "\n\t", F1: v5, F2: true}
                                if (204 != v5) {
                                    v32 := GenS1{F0: "a", F1: (v5 << ((255) & 63)), F2: false}
                                    _ = v32
                                }
                                v3[(v6 / ((v6) | 1))] = ((false || (40 < 0)) || ((v4 + "foo bar") >= ("foo bar" + (v4 + v4))))
                                if (false || true) {
                                    var v33 bool = (!true)
                                    switch (v4 + (v4 + (("" + v4) + "\n\t"))) {
                                    case "":
                                        v34 := []int{821, fi25}
                                        _ = v34
                                    }
                                    _ = v33
                                } else {
                                    var v35 float64 = (v7 + (float64((((1.5 + 1.5) - (((float32(float64(((((float32(float64(1.5)) - 3.14159) * 0.1) - 0.0) + 0.0))) + (3.14159 * -2.25)) / ((3.14159 * 1e10) + 1.5)) / (1e10 * 3.14159))) + float32(3.14159))) - v7))
                                    if (-1 != 1) {
                                        switch "foo bar" {
                                        case "hello":
                                            delete(v3, (^(v6)))
                                        }
                                        v36 := GenS2{F0: v8, F1: 2, F2: (-(fi25))}
                                        _ = v36
                                    } else {
                                        v7 = v35
                                        var v37 int8 = 7
                                        _ = v37
                                    }
                                    _ = v35
                                }
                                fmt.Println(v6)
                                _ = v30
                                _ = v31
                            }
                            v29 += v2
                            _ = v29
                        }
                        v5 = -566
                        if true {
                            v6 -= 548
                            v38 := []int{fi28}
                            v2 += v2
                            _ = v38
                        } else {
                            var v39 float32 = ((1e10 - (0.1 - 1.5)) / -2.25)
                            _ = v39
                        }
                        _ = fi28
                    }
                default:
                    var v40 uint16 = v8
                    _ = v40
                }
                if false {
                    v1 -= (v8 & v1)
                }
                if (((v6 & v6) == v6) || ("a" <= (v4 + ((v4 + (("foo bar" + (v4 + "\n\t")) + (v4 + (v4 + ("hello" + v4))))) + v4)))) {
                    var v41 uint16 = (v1 & v1)
                    v8 = (v41 % (((v8 << ((567) & 63))) | 1))
                    _ = v41
                }
            }
            v4 = "a"
            _ = fi25
        }
    }
    var v42 uint16 = v1
    v7 = float64(0.1)
    if (0 <= 127) {
        v43 := GenS3{F0: ((v4 + v4) + v4), F1: 128, F2: -1}
        var v44 int32 = 1
        v44 -= v44
        _ = v43
        _ = v44
    }
    for fi45 := 0; fi45 < 18; fi45++ {
        v46 := GenS2{F0: (v42 - (v1 & v8)), F1: 2, F2: (fi45 | fi45)}
        _ = fi45
        _ = v46
    }
    for fi47 := 0; fi47 < 3; fi47++ {
        v5 = v5
        if false {
            if true {
                switch (fi47 / ((fi47) | 1)) {
                case -1:
                    if true {
                        v48 := GenS4{F0: 373, F1: 65535, F2: int16(v6)}
                        var v49 bool = (!true)
                        for fi50 := 0; fi50 < 13; fi50++ {
                            if (fi47 <= (fi47 / ((fi50) | 1))) {
                                v49 = v49
                                v51 := GenS2{F0: (v42 << ((4294967295) & 63)), F1: 1, F2: fi47}
                                _ = v51
                            } else {
                                var v52 int32 = (-(v2))
                                v53 := GenS4{F0: 179, F1: 65535, F2: (v5 ^ 1000)}
                                _ = v52
                                _ = v53
                            }
                            v54 := []int{fi50, fi50, fi47, (fi47 / (((^(fi50))) | 1))}
                            var v55 string = v4
                            _ = fi50
                            _ = v54
                            _ = v55
                        }
                        _ = v48
                        _ = v49
                    }
                case 4294967295:
                    v8 = (v1 * v42)
                }
            } else {
                v7 = v7
            }
            v7 -= v7
        }
        _ = fi47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
