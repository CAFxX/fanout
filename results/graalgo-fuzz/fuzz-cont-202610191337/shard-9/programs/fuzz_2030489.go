package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 string; F2 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 int16 }
func genFn0() bool {
    var v4 int32 = -1
    v5 := []int{}
    var v6 uint32 = 2147483648
    switch "a" {
    case "foo bar":
        v7 := GenS1{F0: 42}
        _ = v7
    case "":
        if (v4 <= 0) {
            var v8 int64 = 0
            if true {
                var v9 int8 = -1
                v5 = append(v5, -434)
                _ = v9
            } else {
                for fi10 := 0; fi10 < 8; fi10++ {
                    var v11 uint8 = 127
                    _ = fi10
                    _ = v11
                }
                v4 -= v4
                var v12 bool = (false || false)
                _ = v12
            }
            _ = v8
        } else {
            v6 += (v6 + v6)
            if len(v5) > 0 {
                v5[0] = 32767
            }
            if (949 <= -817) {
                switch ("" + "a") {
                case "foo bar":
                    v6 = v6
                case "a":
                    for fi13 := 0; fi13 < 12; fi13++ {
                        if true {
                            switch "foo bar" {
                            case "\n\t":
                                var v14 int64 = 42
                                _ = v14
                            case "hello":
                                v4 -= 764
                            default:
                                for fi15 := 0; fi15 < 6; fi15++ {
                                    var v16 string = ("" + "foo bar")
                                    _ = fi15
                                    _ = v16
                                }
                            }
                            if false {
                                var v17 float32 = 0.0
                                for fi18 := 0; fi18 < 8; fi18++ {
                                    fmt.Println(fi18, v17)
                                    _ = fi18
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = v17
                            }
                            if false {
                                v19 := GenS1{F0: 0}
                                if len(v5) > 0 {
                                    v5[0] = (fi13 - fi13)
                                }
                                v20 := GenS1{F0: 65535}
                                _ = v19
                                _ = v20
                            } else {
                                for fi21 := 0; fi21 < 18; fi21++ {
                                    v6 = v6
                                    v4 += v4
                                    fmt.Println(fi13, v6)
                                    v22 := GenS1{F0: 2}
                                    _ = fi21
                                    _ = v22
                                }
                            }
                        } else {
                            v6 = v6
                            for fi23 := 0; fi23 < 18; fi23++ {
                                fmt.Println(v6)
                                _ = fi23
                            }
                            var v24 string = (("foo bar" + "foo bar") + "hello")
                            _ = v24
                        }
                        _ = fi13
                    }
                default:
                    if ("a" != "a") {
                        if (65535 <= 255) {
                            var v25 string = "a"
                            v25 = "foo bar"
                            var v26 uint8 = 128
                            _ = v25
                            _ = v26
                        } else {
                            for fi27 := 0; fi27 < 18; fi27++ {
                                fmt.Println(v6, v6, v4)
                                var v28 uint64 = 4294967295
                                _ = fi27
                                _ = v28
                            }
                        }
                    } else {
                        if true {
                            switch 127 {
                            case 32768:
                                if false {
                                    var v29 uint32 = v6
                                    _ = v29
                                } else {
                                    for fi30 := 0; fi30 < 4; fi30++ {
                                        var v31 int32 = v4
                                        var v32 float32 = 1e10
                                        v6 = v6
                                        _ = fi30
                                        _ = v31
                                        _ = v32
                                    }
                                }
                            case -2147483648:
                                var v33 uint64 = 1
                                _ = v33
                            case 835:
                                var v34 float32 = (0.0 / ((float32(float64(1.5)) / 1.0) / 0.1))
                                _ = v34
                            }
                            if (1 <= 1) {
                                var v35 float32 = 0.1
                                _ = v35
                            }
                            switch -134 {
                            case 32768:
                                switch ("a" + ("\n\t" + "foo bar")) {
                                case "a":
                                    v6 -= (v6 & 42)
                                case "foo bar":
                                    v36 := GenS1{F0: 0}
                                    _ = v36
                                case "hello":
                                    v6 = v6
                                }
                            }
                        }
                        fmt.Println(v4, v6)
                        var v37 int32 = v4
                        _ = v37
                    }
                }
            }
        }
    default:
        if len(v5) > 0 {
            v5[0] = -276
        }
    }
    _ = v4
    _ = v5
    _ = v6
    return true
}

func main() {
    v1 := []uint32{1}
    v2 := make(map[uint32]bool)
    v3 := []int{-9223372036854775808, 32768, 4294967295, 127}
    v38 := GenS2{F0: "a", F1: "foo bar", F2: 2147483647}
    var v39 string = ("foo bar" + ("\n\t" + "\n\t"))
    v2[2147483648] = (!(527 <= 0))
    v40 := GenS3{F0: -694}
    switch 32768 {
    case 42:
        fmt.Println(v39)
    }
    v39 = v39
    _ = len(v2)
    v39 = v39
    switch "a" {
    case "foo bar":
        v39 = ((("foo bar" + "a") + v39) + ((v39 + "hello") + (v39 + "a")))
    case "":
        v41 := v2[0]
        _ = v41
    case "a":
        var v42 uint = 42
        _ = v42
    default:
        var v43 int16 = 499
        _ = v43
    }
    _ = len(v3)
    var v44 bool = false
    for fi45 := 0; fi45 < 6; fi45++ {
        v46 := GenS4{F0: 0, F1: 0}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v44 = (0.1 != 1e10)
        _ = fi45
        _ = v46
    }
    v39 = ""
    for fi47 := 0; fi47 < 17; fi47++ {
        for fi48 := 0; fi48 < 13; fi48++ {
            if (!(v39 <= v39)) {
                v39 = v39
                if len(v3) > 0 {
                    v3[0] = (fi47 << ((1) & 63))
                }
            }
            _ = fi48
        }
        _ = fi47
    }
    var v49 uint16 = 668
    var v50 uint = 65535
    var v51 float64 = 1.5
    v44 = v44
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v44:", v44)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v40
    _ = v44
    _ = v49
    _ = v50
    _ = v51
}
