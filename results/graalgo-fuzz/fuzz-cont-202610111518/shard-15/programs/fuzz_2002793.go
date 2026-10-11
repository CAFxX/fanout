package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 string; F1 uint16; F2 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 2
    var v3 int16 = -986
    var v4 string = ""
    if true {
        delete(v1, 158)
    } else {
        delete(v1, 42)
        if true {
            switch v4 {
            case "foo bar":
                v5 := GenS1{F0: true}
                _ = v5
            case "\n\t":
                fmt.Println(v3, v2)
            default:
                if (32767 >= 344) {
                    var v6 uint8 = 255
                    _ = v6
                }
            }
            var v7 int64 = 202
            fmt.Println(v2, v2)
            _ = v7
        } else {
            fmt.Println(v3, v3)
            var v8 bool = true
            _ = v8
        }
        for fi9 := 0; fi9 < 8; fi9++ {
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 int8 = -1
                _ = fi10
                _ = v11
            }
            var v12 float32 = (-2.25 * -2.25)
            fmt.Println(fi9)
            _ = fi9
            _ = v12
        }
    }
    var v13 int32 = 2147483647
    for fi14 := 0; fi14 < 3; fi14++ {
        v13 -= 2147483647
        v1[2147483647] = v4
        _ = fi14
    }
    if false {
        v15 := []int{42, -2147483648, 42}
        v16 := GenS2{F0: 419}
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 bool = (((v2 & 140) < (v2 - v2)) && false)
            for fi19 := 0; fi19 < 18; fi19++ {
                switch ("\n\t" + v4) {
                case "a":
                    if true {
                        v4 = (("" + (v4 + v4)) + "\n\t")
                    } else {
                        v13 = (v13 >> ((884) & 63))
                    }
                case "":
                    var v20 bool = true
                    _ = v20
                case "foo bar":
                    v13 = v13
                    fmt.Println(fi17, v4)
                    v18 = false
                default:
                    var v21 uint8 = 2
                    _ = v21
                }
                if (127 != v2) {
                    v22 := GenS3{F0: 659, F1: 483}
                    _ = v22
                } else {
                    var v23 int16 = (v3 | (v3 ^ v3))
                    _ = v23
                }
                v18 = (1e10 != 1e10)
                fmt.Println(v18)
                var v24 int16 = v3
                _ = fi19
                _ = v24
            }
            _ = fi17
            _ = v18
        }
        _ = v15
        _ = v16
    } else {
        var v25 bool = (int16(v13) >= v3)
        for fi26 := 0; fi26 < 15; fi26++ {
            v27 := []int{433, fi26, (^(fi26))}
            var v28 float64 = 1e10
            for fi29 := 0; fi29 < 2; fi29++ {
                if v25 {
                    fmt.Println(fi29, v25, v2)
                    switch "\n\t" {
                    case "":
                        if v25 {
                            var v30 int64 = -614
                            _ = v30
                        }
                    }
                } else {
                    v31 := GenS2{F0: (v13 % (((v13 / (((v13 | 762)) | 1))) | 1))}
                    _ = v31
                }
                var v32 bool = false
                _ = fi29
                _ = v32
            }
            _ = fi26
            _ = v27
            _ = v28
        }
        v33 := v1[42]
        _ = v25
        _ = v33
    }
    if (!false) {
        fmt.Println(v2, v3)
        var v34 int16 = v3
        _ = v34
    } else {
        if (v13 >= v13) {
            switch 698 {
            case -674:
                for fi35 := 0; fi35 < 20; fi35++ {
                    v36 := v1[476]
                    for fi37 := 0; fi37 < 6; fi37++ {
                        var v38 uint16 = 65535
                        delete(v1, uint32(v2))
                        var v39 float64 = -2.25
                        var v40 int64 = -1
                        v41 := GenS4{F0: ("\n\t" + (v36 + v36)), F1: (v38 | v38), F2: (v40 & int64(v38))}
                        _ = fi37
                        _ = v38
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                    for fi42 := 0; fi42 < 18; fi42++ {
                        if ((("a" > ("a" + ((v4 + (v4 + "hello")) + v4))) && (!false)) && (!true)) {
                            switch int(v13) {
                            case 966:
                                v3 = v3
                            default:
                                v43 := GenS1{F0: (!(!false))}
                                _ = v43
                            }
                            var v44 uint64 = 42
                            v45 := GenS5{F0: 1}
                            _ = v44
                            _ = v45
                        } else {
                            fmt.Println(v36)
                            v13 -= 2147483647
                        }
                        _ = fi42
                    }
                    _ = fi35
                    _ = v36
                }
            case 356:
                v46 := GenS3{F0: 19, F1: 0}
                _ = v46
            }
            fmt.Println(v3, v2, v13)
        } else {
            v13 = (-(v13))
        }
        switch v4 {
        case "hello":
            _ = len(v1)
        case "foo bar":
            var v47 uint64 = 958
            _ = v47
        }
        v48 := GenS4{F0: (("" + v4) + "\n\t"), F1: 495, F2: 42}
        _ = v48
    }
    if false {
        v49 := GenS6{F0: int16(v2)}
        v50 := []int{-1}
        _ = v49
        _ = v50
    }
    v51 := v1[2]
    v52 := GenS4{F0: (v4 + "hello"), F1: 104, F2: -9223372036854775808}
    switch ((v4 + "hello") + (v51 + (v4 + "foo bar"))) {
    case "foo bar":
        var v53 uint16 = 1
        _ = v53
    case "a":
        v3 = v3
    case "\n\t":
        delete(v1, 735)
    }
    v54 := GenS3{F0: 255, F1: 488}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v54
}
