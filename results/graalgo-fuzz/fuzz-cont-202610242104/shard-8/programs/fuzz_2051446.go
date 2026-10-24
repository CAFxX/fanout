package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint32; F1 int64; F2 int32 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 uint16 = 2
    var v2 uint16 = v1
    var v3 uint32 = 0
    var v4 int = 42
    var v5 int16 = 0
    var v6 float32 = ((1.5 - 1.5) - 3.14159)
    fmt.Println(v5, v5, v6)
    fmt.Println(v5, v3, v1)
    if true {
        var v7 string = ""
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 string = ("a" + v7)
            v5 -= (v5 >> ((42) & 63))
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            switch v4 {
            case 127:
                fmt.Println(v2)
            }
            if true {
                switch int(v2) {
                case -134:
                    v6 -= v6
                case 9223372036854775807:
                    for fi11 := 0; fi11 < 10; fi11++ {
                        v7 = (((v7 + (v7 + "hello")) + (v7 + "foo bar")) + "foo bar")
                        var v12 uint32 = v3
                        _ = fi11
                        _ = v12
                    }
                case -370:
                    fmt.Println(v5)
                default:
                    for fi13 := 0; fi13 < 5; fi13++ {
                        v14 := []int{}
                        _ = fi13
                        _ = v14
                    }
                }
                var v15 int = (fi10 + (fi10 << ((uint(v1)) & 63)))
                _ = v15
            } else {
                if true {
                    if (!(!((!((!(363 != uint64(v1))) && (-2147483648 <= -2147483648))) && (!true)))) {
                        fmt.Println(v4, v1)
                    }
                    var v16 int16 = (v5 * v5)
                    switch 32768 {
                    case -8:
                        v1 += v2
                    case 32767:
                        var v17 uint32 = (v3 % ((2147483647) | 1))
                        _ = v17
                    default:
                        var v18 uint16 = 334
                        _ = v18
                    }
                    _ = v16
                } else {
                    for fi19 := 0; fi19 < 9; fi19++ {
                        switch v7 {
                        case "foo bar":
                            if false {
                                var v20 int8 = -96
                                v21 := []int{}
                                _ = v20
                                _ = v21
                            } else {
                                v6 += v6
                                v4 = 619
                                v22 := GenS1{F0: false}
                                _ = v22
                            }
                        case "\n\t":
                            switch v7 {
                            case "":
                                var v23 uint64 = 818
                                var v24 uint8 = 255
                                v25 := GenS1{F0: true}
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                        case "hello":
                            v26 := []int{32768, fi19, fi19}
                            _ = v26
                        default:
                            v27 := []int{(fi19 | (fi10 << ((18446744073709551615) & 63))), fi19, fi10}
                            _ = v27
                        }
                        v28 := []int{}
                        _ = fi19
                        _ = v28
                    }
                    v29 := GenS1{F0: false}
                    v30 := []int{fi10}
                    _ = v29
                    _ = v30
                }
                v7 = v7
                for fi31 := 0; fi31 < 10; fi31++ {
                    v32 := GenS1{F0: ((false && (!((v3 | (v3 / ((896) | 1))) < v3))) && true)}
                    _ = fi31
                    _ = v32
                }
            }
            if ((true && ((255 < 42) && ((1.5 + 0.1) == -2.25))) && false) {
                v1 = (v1 * 255)
                switch (v7 + (((v7 + v7) + (v7 + v7)) + v7)) {
                case "a":
                    var v33 float32 = -2.25
                    _ = v33
                case "\n\t":
                    switch (fi10 >> ((uint(v2)) & 63)) {
                    case -512:
                        for fi34 := 0; fi34 < 3; fi34++ {
                            fmt.Println(v5)
                            var v35 int = fi10
                            v36 := GenS2{F0: ((v7 + v7) + v7), F1: (v1 >> ((255) & 63)), F2: -1}
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                    default:
                        for fi37 := 0; fi37 < 6; fi37++ {
                            fmt.Println(v5, v4, v7)
                            v3 -= (v3 ^ 1)
                            switch 42 {
                            case 42:
                                v5 += v5
                                v38 := GenS3{F0: (^(v3)), F1: 193, F2: -1}
                                _ = v38
                            case -468:
                                v2 = v2
                            case 128:
                                switch ("\n\t" + ((v7 + v7) + v7)) {
                                case "hello":
                                    fmt.Println(v7)
                                case "":
                                    v39 := GenS1{F0: (0 < -1)}
                                    _ = v39
                                }
                            }
                            _ = fi37
                        }
                    }
                default:
                    var v40 uint8 = 0
                    _ = v40
                }
            }
            _ = fi10
        }
        _ = v7
    }
    v41 := GenS4{F0: 4294967295}
    for fi42 := 0; fi42 < 4; fi42++ {
        if true {
            var v43 float32 = 0.0
            _ = v43
        }
        _ = fi42
    }
    v44 := GenS4{F0: 2}
    v3 += (v3 & (^(v3)))
    v5 = (v5 << ((506) & 63))
    switch "foo bar" {
    case "hello":
        switch "foo bar" {
        case "foo bar":
            v45 := []int{0}
            _ = v45
        }
    case "foo bar":
        v46 := GenS2{F0: ("a" + "\n\t"), F1: 1, F2: 0}
        _ = v46
    case "a":
        var v47 uint16 = v1
        _ = v47
    }
    fmt.Println(v5, v5)
    switch v4 {
    case -389:
        var v48 int32 = 2147483647
        _ = v48
    case 127:
        v49 := GenS3{F0: 640, F1: 42, F2: 2147483647}
        _ = v49
    case 4294967295:
        v50 := []int{v4, (v4 - -2147483648), (v4 - v4)}
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v44
    _ = v5
    _ = v6
}
