package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 int = 2147483647
    v2 := []string{}
    var v3 int8 = -88
    v4 := make(map[int]int)
    v5 := make(map[uint32]int)
    v6 := []string{("hello" + ("" + ("\n\t" + "a"))), (("\n\t" + "a") + ""), (("\n\t" + ("a" + "a")) + "")}
    if false {
        v7 := GenS1{F0: -9223372036854775808, F1: false}
        _ = v7
    } else {
        fmt.Println(v1)
        if true {
            if false {
                switch "foo bar" {
                case "hello":
                    switch ("a" + "hello") {
                    case "":
                        var v8 string = ("a" + "hello")
                        _ = v8
                    }
                case "foo bar":
                    var v9 uint8 = 183
                    _ = v9
                case "":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if (1 != 546) {
                    if true {
                        v1 = (v1 + 419)
                    }
                    for fi10 := 0; fi10 < 4; fi10++ {
                        var v11 uint64 = 2
                        if (!(5 > 0)) {
                            for fi12 := 0; fi12 < 15; fi12++ {
                                v4[(v1 + v1)] = fi12
                                var v13 bool = true
                                if (!(1 > 0)) {
                                    v1 = v1
                                    v11 -= (v11 >> ((42) & 63))
                                    var v14 float32 = 1e10
                                    v15 := v4[v1]
                                    var v16 uint16 = 2
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                }
                                _ = fi12
                                _ = v13
                            }
                            v11 += v11
                            var v17 bool = (!((v11 + v11) <= v11))
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 1; fi18++ {
                            v19 := GenS1{F0: 0, F1: true}
                            _ = fi18
                            _ = v19
                        }
                        _ = fi10
                        _ = v11
                    }
                    for fi20 := 0; fi20 < 2; fi20++ {
                        v6 = append(v6, ("hello" + "hello"))
                        _ = fi20
                    }
                } else {
                    v3 -= (v3 >> ((0) & 63))
                }
            }
            var v21 float32 = ((-2.25 * 0.1) - 1.5)
            _ = len(v4)
            _ = v21
        } else {
            v22 := GenS2{F0: "", F1: 194}
            _ = v22
        }
        v1 = v1
    }
    var v23 int16 = 1000
    v23 = v23
    v23 += v23
    var v24 uint32 = 344
    if ((("\n\t" + ("a" + "hello")) != "") || false) {
        v25 := GenS3{F0: 42, F1: 128, F2: 42}
        fmt.Println(v23, v3, v3)
        _ = v25
    }
    v3 = (v3 - int8(v24))
    switch "\n\t" {
    case "hello":
        _ = len(v6)
    case "":
        switch (^(v1)) {
        case 249:
            v26 := GenS1{F0: 746, F1: false}
            _ = v26
        case 277:
            var v27 int32 = 714
            _ = v27
        }
    case "a":
        if (!(32767 <= 65535)) {
            switch 2147483647 {
            case 416:
                for fi28 := 0; fi28 < 9; fi28++ {
                    for fi29 := 0; fi29 < 2; fi29++ {
                        if (65535 == 2) {
                            fmt.Println(v3)
                            v1 -= (v1 & (fi29 % ((int(v24)) | 1)))
                            v3 += (v3 % ((-10) | 1))
                        } else {
                            v24 += v24
                            if (!(640 == 420)) {
                                var v30 bool = false
                                v3 -= v3
                                v31 := GenS2{F0: "", F1: 0}
                                _ = v30
                                _ = v31
                            }
                            if false {
                                v23 += int16(v24)
                                var v32 float64 = float64(1e10)
                                v23 += (v23 ^ -650)
                                _ = v32
                            } else {
                                var v33 float64 = 0.0
                                v4[(v1 & fi28)] = v1
                                _ = v33
                            }
                        }
                        _ = fi29
                    }
                    switch "hello" {
                    case "\n\t":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "":
                        fmt.Println(v23, v23)
                    }
                    _ = fi28
                }
            case 32767:
                _ = len(v2)
            }
            fmt.Println(v1, v24, v3)
            var v34 int16 = -1
            _ = v34
        }
    }
    var v35 uint = 4294967295
    if true {
        var v36 uint32 = 699
        for fi37 := 0; fi37 < 16; fi37++ {
            v38 := GenS4{F0: true}
            if false {
                var v39 uint16 = 0
                _ = v39
            } else {
                delete(v5, v24)
                if (true || true) {
                    if (v36 == (v36 >> (((v35 - v35)) & 63))) {
                        switch "foo bar" {
                        case "\n\t":
                            if ((0.0 / 3.14159) > -2.25) {
                                v1 -= (v1 ^ -9223372036854775808)
                            } else {
                                for fi40 := 0; fi40 < 17; fi40++ {
                                    var v41 uint16 = 32767
                                    _ = fi40
                                    _ = v41
                                }
                            }
                        default:
                            var v42 uint64 = 844
                            _ = v42
                        }
                        v43 := GenS4{F0: (!true)}
                        _ = v43
                    }
                    for fi44 := 0; fi44 < 9; fi44++ {
                        var v45 int32 = -891
                        _ = fi44
                        _ = v45
                    }
                }
            }
            _ = fi37
            _ = v38
        }
        if (128 != 129) {
            v46 := GenS5{F0: 688}
            for fi47 := 0; fi47 < 11; fi47++ {
                var v48 float64 = float64((0.1 - 1.5))
                for fi49 := 0; fi49 < 9; fi49++ {
                    var v50 int64 = -1
                    if false {
                        _ = len(v4)
                        var v51 uint16 = 1
                        fmt.Println(v36, v24, v50)
                        v4[-208] = (fi47 ^ fi49)
                        _ = len(v6)
                        _ = v51
                    }
                    v1 += (fi49 - -284)
                    switch "hello" {
                    case "":
                        var v52 uint8 = 0
                        _ = v52
                    case "foo bar":
                        if len(v2) > 0 {
                            v2[0] = ("\n\t" + "a")
                        }
                    case "a":
                        v3 -= 42
                    }
                    _ = fi49
                    _ = v50
                }
                v53 := GenS3{F0: 2147483647, F1: 127, F2: 332}
                _ = fi47
                _ = v48
                _ = v53
            }
            _ = v46
        } else {
            delete(v5, 2147483647)
            v3 = v3
            v2 = append(v2, "foo bar")
        }
        _ = v36
    } else {
        if ((-1 == 9223372036854775807) && (96 < 2)) {
            _ = len(v5)
            v24 = v24
            var v54 int64 = -317
            _ = v54
        }
        for fi55 := 0; fi55 < 2; fi55++ {
            for fi56 := 0; fi56 < 6; fi56++ {
                v35 += 2
                v1 -= fi56
                if false {
                    v24 = (v24 + v24)
                    var v57 uint8 = 128
                    _ = v57
                }
                _ = fi56
            }
            _ = len(v2)
            _ = fi55
        }
        v58 := GenS1{F0: -565, F1: false}
        _ = v58
    }
    v23 = (^(v23))
    v59 := GenS3{F0: -2147483648, F1: 59, F2: 32767}
    var v60 bool = (v1 > (v1 % (((v1 * v1)) | 1)))
    var v61 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v61
}
