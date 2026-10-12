package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int64 }
func genFn0(p0 uint32) int16 {
    switch 1 {
    case 11:
        var v7 int16 = 32767
        _ = v7
    default:
        v8 := []int{}
        _ = v8
    }
    p0 = p0
    var v9 uint64 = 42
    _ = p0
    _ = v9
    return 42
}
func genFn1(p0 uint, p1 uint8, p2 uint32) string {
    v10 := GenS1{F0: ("hello" + "")}
    var v11 int = -711
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := []int{(fi12 % ((v11) | 1))}
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return ("a" + (("foo bar" + "hello") + "\n\t"))
}

func main() {
    v1 := []uint32{791, 321, 4294967295}
    v2 := make(map[uint32]uint32)
    var v3 uint = 255
    v4 := make(map[uint32]uint32)
    var v5 int16 = 972
    v6 := []int{142, 0, -406}
    var v14 int8 = -1
    v14 = 127
    if true {
        v3 -= (v3 >> (((v3 ^ 0)) & 63))
        fmt.Println(v3, v3, v3)
    } else {
        var v15 uint8 = 1
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 uint16 = 65535
            for fi18 := 0; fi18 < 1; fi18++ {
                if (-432 != -9223372036854775808) {
                    if (!((0 != -1) && ((^(v14)) <= v14))) {
                        v17 = (v17 + (v17 + v17))
                        v19 := GenS2{F0: (v5 | (v5 / (((v5 & 1000)) | 1)))}
                        _ = v19
                    }
                    v1 = append(v1, 2147483647)
                    var v20 uint16 = uint16(v15)
                    _ = v20
                } else {
                    fmt.Println(v5)
                }
                v3 = 0
                switch "hello" {
                case "a":
                    v15 -= 1
                case "foo bar":
                    v21 := GenS3{F0: ("a" + ("hello" + "hello"))}
                    _ = v21
                default:
                    switch ("foo bar" + "") {
                    case "hello":
                        if true {
                            if (839 <= fi18) {
                                if true {
                                    switch "hello" {
                                    case "\n\t":
                                        var v22 float32 = float32(-2.25)
                                        _ = v22
                                    case "hello":
                                        var v23 uint8 = v15
                                        _ = v23
                                    case "foo bar":
                                        fmt.Println(v17, v15, v17)
                                    default:
                                        fmt.Println(v15, v15, fi16)
                                    }
                                } else {
                                    v2[666] = 2
                                    v15 = v15
                                }
                            } else {
                                var v24 int = -360
                                fmt.Println(v3, v15, fi18)
                                var v25 uint16 = v17
                                _ = v24
                                _ = v25
                            }
                        } else {
                            var v26 uint = v3
                            if true {
                                v27 := v2[65535]
                                fmt.Println(v14)
                                var v28 string = "foo bar"
                                if true {
                                    var v29 uint8 = v15
                                    _ = v29
                                }
                                _ = v27
                                _ = v28
                            }
                            _ = v26
                        }
                    default:
                        if true {
                            var v30 float64 = 0.0
                            _ = v30
                        } else {
                            fmt.Println(v17, v3)
                            v14 -= -128
                        }
                    }
                }
                _ = fi18
            }
            var v31 uint = 809
            _ = fi16
            _ = v17
            _ = v31
        }
        v15 = (v15 + 42)
        _ = v15
    }
    if false {
        for fi32 := 0; fi32 < 17; fi32++ {
            if true {
                for fi33 := 0; fi33 < 4; fi33++ {
                    v3 += 123
                    v34 := GenS4{F0: v3}
                    _ = fi33
                    _ = v34
                }
            }
            var v35 uint = v3
            v36 := GenS5{F0: 4294967295}
            _ = fi32
            _ = v35
            _ = v36
        }
        switch "\n\t" {
        case "":
            fmt.Println(v3)
        case "foo bar":
            if (65535 != 386) {
                v5 = -1
                v37 := GenS6{F0: (!true)}
                for fi38 := 0; fi38 < 7; fi38++ {
                    var v39 uint16 = 42
                    switch 32768 {
                    case -698:
                        v40 := GenS5{F0: 459}
                        _ = v40
                    case 2147483647:
                        v3 = (v3 + v3)
                    case 32768:
                        v39 -= v39
                    }
                    _ = fi38
                    _ = v39
                }
                _ = v37
            }
        case "hello":
            var v41 int = 128
            _ = v41
        default:
            for fi42 := 0; fi42 < 17; fi42++ {
                for fi43 := 0; fi43 < 15; fi43++ {
                    v44 := GenS5{F0: 0}
                    switch "hello" {
                    case "foo bar":
                        var v45 string = ((("a" + "foo bar") + "") + "hello")
                        _ = v45
                    }
                    _ = fi43
                    _ = v44
                }
                _ = fi42
            }
        }
        if true {
            for fi46 := 0; fi46 < 8; fi46++ {
                v47 := GenS5{F0: 2}
                _ = fi46
                _ = v47
            }
            v48 := GenS2{F0: -32768}
            _ = v48
        } else {
            var v49 int64 = -9223372036854775808
            _ = len(v6)
            switch 32767 {
            case -934:
                _ = len(v6)
            case 397:
                v50 := GenS2{F0: 0}
                _ = v50
            case 0:
                v49 += 1
            }
            _ = v49
        }
    } else {
        if false {
            if (2147483647 < 24) {
                fmt.Println(v14)
                for fi51 := 0; fi51 < 17; fi51++ {
                    for fi52 := 0; fi52 < 19; fi52++ {
                        var v53 uint = v3
                        if (!true) {
                            var v54 uint = (^(v3))
                            _ = v54
                        } else {
                            var v55 bool = false
                            _ = v55
                        }
                        switch "hello" {
                        case "\n\t":
                            if false {
                                switch "" {
                                case "\n\t":
                                    fmt.Println(v5, v14, v53)
                                case "hello":
                                    var v56 float64 = 1.5
                                    _ = v56
                                }
                                var v57 uint16 = 0
                                _ = v57
                            } else {
                                switch fi51 {
                                case 127:
                                    _ = len(v6)
                                case -9223372036854775808:
                                    v1 = append(v1, 630)
                                case -970:
                                    var v58 float64 = 3.14159
                                    _ = v58
                                default:
                                    v14 -= (v14 | 0)
                                }
                                _ = len(v1)
                                var v59 float64 = 3.14159
                                _ = v59
                            }
                        }
                        _ = fi52
                        _ = v53
                    }
                    _ = fi51
                }
                var v60 uint32 = 0
                _ = v60
            } else {
                fmt.Println(v5)
                fmt.Println(v3, v14, v14)
            }
            var v61 float32 = 0.1
            fmt.Println(v5, v5, v61)
            _ = v61
        }
        if true {
            v14 -= 0
            v14 = 0
            for fi62 := 0; fi62 < 15; fi62++ {
                var v63 uint16 = 0
                v14 = int8(v3)
                _ = fi62
                _ = v63
            }
        }
        for fi64 := 0; fi64 < 18; fi64++ {
            var v65 uint32 = 64
            fmt.Println(v3, v5)
            _ = fi64
            _ = v65
        }
    }
    for fi66 := 0; fi66 < 2; fi66++ {
        v14 -= (v14 % (((v14 + (^(v14)))) | 1))
        var v67 int32 = 712
        v2[2147483648] = 2
        _ = fi66
        _ = v67
    }
    v5 -= genFn0(1)
    switch 127 {
    case 32767:
        if (true && true) {
            var v68 int16 = 625
            _ = v68
        } else {
            fmt.Println(v5, v5, v14)
        }
    case 651:
        fmt.Println(v3, v3)
    default:
        for fi69 := 0; fi69 < 20; fi69++ {
            v6 = append(v6, (fi69 << ((v3) & 63)))
            fmt.Println(v3, v5)
            _ = fi69
        }
    }
    v70 := GenS5{F0: 537}
    if (!true) {
        if false {
            if ((4294967295 > 0) && (v3 != (^(v3)))) {
                var v71 uint16 = 1
                v72 := GenS7{F0: -232}
                _ = v71
                _ = v72
            }
            var v73 string = "foo bar"
            _ = v73
        }
        v74 := v2[562]
        _ = v74
    }
    v5 -= v5
    var v75 bool = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v75
}
