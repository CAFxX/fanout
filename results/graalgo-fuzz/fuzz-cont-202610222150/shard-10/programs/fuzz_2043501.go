package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 uint16; F2 int64 }
type GenS5 struct { F0 uint64; F1 int; F2 uint }

func main() {
    var v1 string = ""
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    if (-490 >= 9223372036854775807) {
        switch 0 {
        case 2147483647:
            v4 := GenS1{F0: 0}
            _ = v4
        }
    } else {
        v1 = v1
    }
    v1 = v1
    var v5 int64 = -1
    var v6 uint8 = 1
    v7 := GenS1{F0: v6}
    switch "\n\t" {
    case "hello":
        v1 = ("a" + "a")
    case "":
        var v8 float64 = float64(1e10)
        _ = v8
    case "a":
        for fi9 := 0; fi9 < 9; fi9++ {
            for fi10 := 0; fi10 < 14; fi10++ {
                fmt.Println(fi9, fi9, v6)
                fmt.Println(fi10)
                _ = fi10
            }
            var v11 int64 = v5
            var v12 int64 = v5
            _ = fi9
            _ = v11
            _ = v12
        }
    default:
        fmt.Println(v6)
    }
    fmt.Println(v6, v1, v1)
    fmt.Println(v1)
    switch 1 {
    case 4294967295:
        for fi13 := 0; fi13 < 17; fi13++ {
            v14 := GenS2{F0: v1, F1: fi13}
            var v15 uint8 = v6
            _ = fi13
            _ = v14
            _ = v15
        }
    case 128:
        switch (v1 + (v1 + v1)) {
        case "\n\t":
            var v16 float64 = float64(-2.25)
            _ = v16
        case "":
            if (328 == 2147483648) {
                v6 -= (v6 - v6)
            } else {
                v17 := GenS3{F0: v5}
                _ = v17
            }
        }
    default:
        v5 = v5
    }
    v1 = v1
    v5 -= -691
    var v18 string = v1
    v5 = v5
    v6 = 2
    if true {
        v1 = v1
        if false {
            for fi19 := 0; fi19 < 3; fi19++ {
                fmt.Println(fi19)
                _ = fi19
            }
            for fi20 := 0; fi20 < 9; fi20++ {
                v3[2147483647] = "a"
                _ = fi20
            }
            v21 := []int{4294967295, -2147483648}
            _ = v21
        }
        for fi22 := 0; fi22 < 6; fi22++ {
            for fi23 := 0; fi23 < 3; fi23++ {
                v18 = "\n\t"
                v24 := GenS2{F0: ("hello" + (("" + "a") + v1)), F1: fi23}
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    } else {
        v25 := []int{32767}
        v5 += v5
        switch 9223372036854775807 {
        case 235:
            for fi26 := 0; fi26 < 6; fi26++ {
                fmt.Println(v18, v1, v6)
                _ = fi26
            }
        case 32767:
            if false {
                switch -41 {
                case 2147483647:
                    delete(v2, -249)
                case 127:
                    v6 = 254
                case 128:
                    if (2147483647 != -2147483648) {
                        var v27 uint64 = 0
                        var v28 uint32 = 4294967295
                        _ = v27
                        _ = v28
                    }
                }
            } else {
                if true {
                    v18 = v1
                    fmt.Println(v5, v5)
                    v29 := GenS4{F0: false, F1: 706, F2: v5}
                    _ = v29
                } else {
                    if true {
                        for fi30 := 0; fi30 < 1; fi30++ {
                            var v31 float32 = 0.1
                            if (54 > 42) {
                                v32 := GenS5{F0: 727, F1: (fi30 - (fi30 << ((4294967295) & 63))), F2: 914}
                                _ = v32
                            }
                            var v33 uint16 = 42
                            _ = fi30
                            _ = v31
                            _ = v33
                        }
                        fmt.Println(v5, v6, v5)
                        for fi34 := 0; fi34 < 19; fi34++ {
                            if len(v25) >= 2 {
                                v25 = v25[1:len(v25)]
                            }
                            var v35 float32 = -2.25
                            var v36 uint16 = 2
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                    }
                }
                if (1 == 127) {
                    for fi37 := 0; fi37 < 6; fi37++ {
                        delete(v3, fi37)
                        _ = fi37
                    }
                    v3[204] = v18
                }
                switch v18 {
                case "\n\t":
                    v2[4294967295] = 32767
                default:
                    switch 716 {
                    case 9223372036854775807:
                        _ = len(v25)
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            var v38 string = ("\n\t" + (("a" + "a") + ((v18 + (v18 + v1)) + v1)))
                            _ = v38
                        case "foo bar":
                            if false {
                                v25 = append(v25, 2147483647)
                                var v39 int = 908
                                _ = v39
                            }
                        default:
                            var v40 float64 = 0.0
                            _ = v40
                        }
                    }
                }
            }
        }
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
