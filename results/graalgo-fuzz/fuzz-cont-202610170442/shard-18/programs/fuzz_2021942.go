package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 string; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int32; F2 int }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 uint8 }
func genFn0() uint {
    var v4 uint32 = 2147483647
    v4 = (v4 << ((42) & 63))
    _ = v4
    return 255
}

func main() {
    var v1 int64 = 847
    var v2 bool = (true && (!false))
    var v3 uint = 2
    fmt.Println(v1)
    for fi5 := 0; fi5 < 14; fi5++ {
        fmt.Println(fi5, fi5)
        _ = fi5
    }
    fmt.Println(v1, v3)
    if v2 {
        v6 := []int{2147483647, -9223372036854775808, 9223372036854775807}
        if (!((!v2) && (!v2))) {
            if v2 {
                switch -9223372036854775808 {
                case 1:
                    v7 := GenS1{F0: -1}
                    _ = v7
                case 127:
                    if len(v6) > 0 {
                        v6[0] = 160
                    }
                case 4294967295:
                    v8 := GenS2{F0: v3, F1: "foo bar", F2: 1}
                    _ = v8
                default:
                    v2 = true
                }
                var v9 uint32 = 0
                _ = v9
            }
        }
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS3{F0: 0, F1: 42, F2: fi10}
            _ = fi10
            _ = v11
        }
        _ = v6
    }
    var v12 uint16 = 1
    var v13 float32 = ((1e10 - -2.25) / 1.5)
    v14 := GenS4{F0: false, F1: (v3 ^ (v3 / (((v3 - v3)) | 1)))}
    var v15 uint = 2
    if ((!(0.1 == (3.14159 * v13))) && true) {
        for fi16 := 0; fi16 < 16; fi16++ {
            v15 += (v3 * v3)
            v17 := GenS3{F0: 255, F1: 0, F2: (fi16 << ((v15) & 63))}
            _ = fi16
            _ = v17
        }
        var v18 uint16 = 65535
        switch 9223372036854775807 {
        case -1:
            v3 = v15
        case -82:
            fmt.Println(v2)
        case 128:
            fmt.Println(v18, v2)
        }
        _ = v18
    } else {
        v12 += v12
        switch "foo bar" {
        case "foo bar":
            if ((^(v15)) < 94) {
                fmt.Println(v2, v2, v3)
                v19 := []int{}
                if true {
                    v12 += (v12 - (v12 >> ((981) & 63)))
                    switch -413 {
                    case 32767:
                        if (("a" < (((("a" + "foo bar") + (("a" + "a") + "foo bar")) + "\n\t") + "")) && (!v2)) {
                            var v20 int = -157
                            _ = v20
                        }
                    case 113:
                        var v21 string = "a"
                        _ = v21
                    }
                }
                _ = v19
            } else {
                v22 := GenS5{F0: -1}
                switch "a" {
                case "a":
                    fmt.Println(v12, v12, v12)
                case "foo bar":
                    if v2 {
                        if true {
                            switch 127 {
                            case 878:
                                for fi23 := 0; fi23 < 20; fi23++ {
                                    v24 := []int{(fi23 | (fi23 - -9223372036854775808)), 32768, fi23}
                                    if v2 {
                                        v15 += (v3 ^ (v15 ^ v15))
                                        var v25 uint = 921
                                        _ = v25
                                    }
                                    _ = fi23
                                    _ = v24
                                }
                            case -2147483648:
                                var v26 float32 = 1.5
                                _ = v26
                            case -392:
                                for fi27 := 0; fi27 < 8; fi27++ {
                                    for fi28 := 0; fi28 < 12; fi28++ {
                                        v12 -= 65535
                                        v1 += (^(v1))
                                        v15 = (v15 + v15)
                                        _ = fi28
                                    }
                                    fmt.Println(v1, v15, v15)
                                    _ = fi27
                                }
                            }
                            v29 := GenS6{F0: 127, F1: 255}
                            if (false && (v2 || v2)) {
                                if (-1 < -347) {
                                    switch 42 {
                                    case 581:
                                        var v30 uint8 = 42
                                        _ = v30
                                    default:
                                        var v31 int = 491
                                        _ = v31
                                    }
                                    var v32 int32 = 0
                                    _ = v32
                                } else {
                                    if true {
                                        var v33 int8 = -114
                                        v2 = v2
                                        _ = v33
                                    }
                                    if (1.5 == v13) {
                                        var v34 float64 = 3.14159
                                        if (1 != 0) {
                                            v35 := []int{}
                                            fmt.Println(v13, v34)
                                            var v36 int8 = -1
                                            _ = v35
                                            _ = v36
                                        }
                                        _ = v34
                                    }
                                }
                                v37 := []int{-127, -1}
                                v37 = append(v37, -2147483648)
                                _ = v37
                            } else {
                                if false {
                                    fmt.Println(v13)
                                    v3 += genFn0()
                                    fmt.Println(v2, v15, v2)
                                }
                                if v2 {
                                    var v38 int32 = 2147483647
                                    _ = v38
                                }
                            }
                            _ = v29
                        } else {
                            v39 := GenS1{F0: -32768}
                            _ = v39
                        }
                        v40 := GenS6{F0: 53, F1: 27}
                        _ = v40
                    } else {
                        v41 := GenS3{F0: 2, F1: 703, F2: -1}
                        switch "" {
                        case "foo bar":
                            fmt.Println(v13, v13, v12)
                        case "\n\t":
                            v42 := []int{-717, -134, -9223372036854775808}
                            _ = v42
                        }
                        _ = v41
                    }
                case "\n\t":
                    v43 := []int{-489, 127}
                    _ = v43
                }
                _ = v22
            }
        case "hello":
            switch ("\n\t" + "hello") {
            case "":
                switch 2147483647 {
                case 32767:
                    var v44 float64 = 1e10
                    _ = v44
                case -32:
                    if (215 > 0) {
                        for fi45 := 0; fi45 < 6; fi45++ {
                            v13 = v13
                            for fi46 := 0; fi46 < 7; fi46++ {
                                var v47 int32 = -2147483648
                                v3 = v15
                                v1 = v1
                                var v48 uint64 = 49
                                var v49 int64 = (v1 & v1)
                                _ = fi46
                                _ = v47
                                _ = v48
                                _ = v49
                            }
                            _ = fi45
                        }
                    }
                default:
                    var v50 float32 = v13
                    _ = v50
                }
            case "\n\t":
                for fi51 := 0; fi51 < 20; fi51++ {
                    v52 := GenS1{F0: -1}
                    _ = fi51
                    _ = v52
                }
            case "foo bar":
                var v53 uint = genFn0()
                _ = v53
            }
        default:
            v54 := []int{2147483647, -1}
            _ = v54
        }
        v3 = 65535
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
