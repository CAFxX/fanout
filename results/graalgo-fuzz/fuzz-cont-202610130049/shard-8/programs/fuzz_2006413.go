package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int64) uint64 {
    switch ("\n\t" + "foo bar") {
    case "a":
        for fi6 := 0; fi6 < 13; fi6++ {
            var v7 uint32 = 65535
            v8 := GenS1{F0: -78}
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 int16 = 376
                p0 += (p0 << ((0) & 63))
                p0 -= p0
                _ = fi9
                _ = v10
            }
            _ = fi6
            _ = v7
            _ = v8
        }
    default:
        for fi11 := 0; fi11 < 9; fi11++ {
            if false {
                p0 = -1
                v12 := []int{fi11}
                _ = v12
            } else {
                var v13 float64 = ((float64(3.14159) - 0.0) - 0.0)
                if true {
                    for fi14 := 0; fi14 < 5; fi14++ {
                        var v15 uint16 = 42
                        switch "a" {
                        case "a":
                            v16 := GenS2{F0: 32767, F1: 240}
                            _ = v16
                        }
                        _ = fi14
                        _ = v15
                    }
                    v13 = 1e10
                }
                var v17 int8 = -1
                _ = v13
                _ = v17
            }
            _ = fi11
        }
    }
    v18 := []int{-1, -9223372036854775808, -1, 32767}
    p0 = p0
    for fi19 := 0; fi19 < 8; fi19++ {
        p0 = p0
        _ = fi19
    }
    _ = p0
    _ = v18
    return 51
}

func main() {
    v1 := []uint32{9, 0, 2, 65535}
    var v2 string = ""
    var v3 float64 = float64(-2.25)
    v4 := []uint32{}
    var v5 float32 = 1.5
    for fi20 := 0; fi20 < 14; fi20++ {
        var v21 bool = (1 >= 206)
        var v22 int32 = 47
        _ = fi20
        _ = v21
        _ = v22
    }
    if (!(115 != 1)) {
        v1 = append(v1, 430)
        var v23 uint = 65535
        _ = v23
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v24 bool = (!(((!(v5 != 1e10)) || true) && false))
    var v25 int8 = 1
    if false {
        var v26 int8 = v25
        _ = v26
    } else {
        if (((42 < -725) || false) || (!(!v24))) {
            if (!true) {
                switch -35 {
                case 127:
                    v24 = (!(332 != 1))
                case -1:
                    var v27 uint64 = 2
                    _ = v27
                default:
                    var v28 int16 = -890
                    _ = v28
                }
            } else {
                var v29 uint32 = 2147483648
                fmt.Println(v24, v25)
                _ = v29
            }
            v2 = ""
            var v30 int32 = 1
            _ = v30
        }
        var v31 uint8 = 81
        if v24 {
            _ = len(v1)
        } else {
            for fi32 := 0; fi32 < 18; fi32++ {
                var v33 int8 = (v25 | 42)
                for fi34 := 0; fi34 < 9; fi34++ {
                    fmt.Println(fi34, v2, v25)
                    fmt.Println(v25)
                    switch "foo bar" {
                    case "":
                        _ = len(v1)
                    }
                    var v35 int64 = 0
                    _ = fi34
                    _ = v35
                }
                switch v2 {
                case "hello":
                    for fi36 := 0; fi36 < 20; fi36++ {
                        v37 := GenS1{F0: v33}
                        _ = fi36
                        _ = v37
                    }
                default:
                    if (1000 > 1000) {
                        var v38 uint = 0
                        switch v2 {
                        case "":
                            var v39 float32 = v5
                            _ = v39
                        case "a":
                            var v40 uint16 = 2
                            _ = v40
                        default:
                            v41 := GenS3{F0: 1}
                            _ = v41
                        }
                        _ = v38
                    }
                }
                _ = fi32
                _ = v33
            }
            v42 := GenS2{F0: 32767, F1: 65535}
            if len(v4) > 0 {
                v4[0] = 0
            }
            _ = v42
        }
        _ = v31
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi43 := 0; fi43 < 9; fi43++ {
        v5 -= v5
        _ = fi43
    }
    var v44 float32 = (1e10 * v5)
    var v45 float32 = 0.0
    switch (((v2 + (("foo bar" + v2) + ((("\n\t" + (v2 + "\n\t")) + "a") + v2))) + "\n\t") + (v2 + (v2 + (v2 + "a")))) {
    case "\n\t":
        var v46 int8 = (v25 * v25)
        _ = v46
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "hello":
        if v24 {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    v47 := GenS1{F0: 1}
    v2 = (v2 + "hello")
    if (-532 > 879) {
        var v48 int16 = -994
        v49 := GenS4{F0: 26}
        _ = v48
        _ = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v47
    _ = v5
}
