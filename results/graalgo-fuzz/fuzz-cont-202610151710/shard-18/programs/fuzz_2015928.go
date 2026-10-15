package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int32; F2 bool }
type GenS4 struct { F0 int8; F1 uint; F2 uint64 }
type GenS5 struct { F0 uint8 }
func genFn0() int32 {
    v6 := []int{-9223372036854775808}
    v7 := GenS1{F0: 42, F1: 32767}
    var v8 int8 = 19
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v9 int64 = -9223372036854775808
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int64 = -668
    var v2 string = "\n\t"
    v3 := make(map[string]uint32)
    v4 := make(map[int]int)
    var v5 int32 = -450
    if true {
        fmt.Println(v5)
        v10 := []int{128, -1, -1}
        fmt.Println(v1)
        _ = v10
    }
    v1 -= v1
    var v11 int32 = -882
    v5 -= (v11 + 2147483647)
    var v12 uint16 = 32767
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 1; fi14++ {
            v15 := []int{fi14}
            var v16 int64 = 1
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    switch -1 {
    case 128:
        fmt.Println(v2, v2, v11)
    case 32767:
        v17 := GenS1{F0: 1, F1: v12}
        _ = v17
    default:
        for fi18 := 0; fi18 < 12; fi18++ {
            fmt.Println(v1, v2)
            var v19 int64 = v1
            _ = fi18
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        switch ((((v2 + ((((v2 + (v2 + "")) + v2) + (v2 + v2)) + "a")) + "\n\t") + "\n\t") + "foo bar") {
        case "foo bar":
            var v21 int64 = v1
            _ = v21
        default:
            var v22 int64 = (v1 * 0)
            _ = v22
        }
        v23 := GenS1{F0: 42, F1: 32767}
        _ = fi20
        _ = v23
    }
    var v24 float64 = -2.25
    v25 := GenS1{F0: -1, F1: (v12 / (((v12 + v12)) | 1))}
    if false {
        if ((-(v5)) != v11) {
            if false {
                switch ("foo bar" + ("\n\t" + (("foo bar" + v2) + ""))) {
                case "":
                    v24 = v24
                }
                var v26 int8 = 75
                _ = v26
            }
            v5 = 0
        }
    }
    fmt.Println(v5)
    if (!(1 == 65535)) {
        v1 = (v1 % (((v1 + (v1 - -9223372036854775808))) | 1))
    }
    v27 := GenS2{F0: v2}
    if (((-861 > 1) || true) && (true || ((255 < 2) || (((v11 | v5) != genFn0()) || (!(float32(v24) >= float32(0.0))))))) {
        fmt.Println(v1)
        if (-868 != -32768) {
            if true {
                for fi28 := 0; fi28 < 11; fi28++ {
                    switch v2 {
                    case "a":
                        v1 += (v1 / ((v1) | 1))
                    }
                    if (false && false) {
                        v29 := GenS3{F0: false, F1: genFn0(), F2: true}
                        _ = v29
                    }
                    var v30 int32 = v5
                    _ = fi28
                    _ = v30
                }
            } else {
                var v31 int16 = -730
                _ = v31
            }
            if false {
                v12 -= 0
                fmt.Println(v12, v12, v12)
            } else {
                switch -269 {
                case 127:
                    if true {
                        v1 = 1
                    } else {
                        if (1 < v12) {
                            if false {
                                for fi32 := 0; fi32 < 20; fi32++ {
                                    var v33 int = (fi32 * fi32)
                                    fmt.Println(v5)
                                    _ = fi32
                                    _ = v33
                                }
                            }
                        } else {
                            v34 := []int{-990, 32767}
                            fmt.Println(v12, v1, v5)
                            _ = v34
                        }
                    }
                case 128:
                    var v35 uint = 42
                    _ = v35
                }
                for fi36 := 0; fi36 < 6; fi36++ {
                    v37 := GenS4{F0: 42, F1: 65535, F2: 1}
                    _ = fi36
                    _ = v37
                }
            }
        } else {
            v38 := []int{-446, 1}
            if true {
                v1 -= v1
                for fi39 := 0; fi39 < 9; fi39++ {
                    v40 := GenS5{F0: 128}
                    _ = fi39
                    _ = v40
                }
                var v41 string = v2
                _ = v41
            } else {
                v38 = append(v38, -108)
                switch (v2 + v2) {
                case "foo bar":
                    fmt.Println(v1, v11, v12)
                case "a":
                    v24 -= (v24 * v24)
                }
            }
            _ = v38
        }
        for fi42 := 0; fi42 < 5; fi42++ {
            v43 := GenS4{F0: 42, F1: 0, F2: 4294967295}
            var v44 string = ("hello" + ((v2 + "") + v2))
            _ = fi42
            _ = v43
            _ = v44
        }
    }
    var v45 uint64 = 0
    for fi46 := 0; fi46 < 16; fi46++ {
        fmt.Println(v2)
        _ = fi46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}
