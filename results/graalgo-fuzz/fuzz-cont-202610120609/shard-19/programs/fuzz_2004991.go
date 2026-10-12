package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 int8; F1 int8; F2 bool }
type GenS6 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint, p1 uint) uint16 {
    for fi6 := 0; fi6 < 2; fi6++ {
        switch fi6 {
        case 425:
            for fi7 := 0; fi7 < 7; fi7++ {
                var v8 string = (("a" + "a") + "foo bar")
                v9 := []int{(fi6 / (((fi7 % (((fi6 & 270)) | 1))) | 1))}
                v10 := GenS1{F0: -1, F1: 1, F2: 0}
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
        default:
            v11 := GenS2{F0: 2147483648, F1: 255, F2: 547}
            _ = v11
        }
        p1 += p0
        _ = fi6
    }
    p0 = 12
    var v12 uint32 = 2147483648
    v12 += v12
    if (!false) {
        v13 := GenS3{F0: -66, F1: -111}
        if (true && true) {
            v14 := []int{-9223372036854775808}
            switch 42 {
            case 4294967295:
                v12 -= (^(v12))
            case 1:
                switch ("hello" + "") {
                case "foo bar":
                    var v15 int16 = 42
                    _ = v15
                }
            default:
                var v16 int16 = 1000
                _ = v16
            }
            p0 -= (p1 & p0)
            _ = v14
        } else {
            if false {
                switch -9223372036854775808 {
                case 42:
                    fmt.Println(v12, v12, p0)
                case -133:
                    p1 = p1
                case -584:
                    if ((v12 % ((v12) | 1)) == (^(v12))) {
                        fmt.Println(p1, p0, p0)
                    }
                }
                for fi17 := 0; fi17 < 6; fi17++ {
                    p1 += 2
                    var v18 uint = 1
                    v19 := []int{502}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                if true {
                    switch ((("a" + "") + "") + "foo bar") {
                    case "\n\t":
                        for fi20 := 0; fi20 < 9; fi20++ {
                            v21 := []int{fi20}
                            _ = fi20
                            _ = v21
                        }
                    default:
                        fmt.Println(p1, p0)
                    }
                    v12 += v12
                    switch 42 {
                    case 1:
                        v22 := GenS4{F0: 0, F1: -368}
                        _ = v22
                    case 4294967295:
                        v23 := []int{371, -2147483648, 42, 1}
                        _ = v23
                    default:
                        var v24 uint8 = 127
                        _ = v24
                    }
                } else {
                    p1 += p1
                    for fi25 := 0; fi25 < 11; fi25++ {
                        p0 = (p1 << ((p1) & 63))
                        _ = fi25
                    }
                    var v26 float64 = (3.14159 + -2.25)
                    _ = v26
                }
            } else {
                p0 -= p1
                var v27 int32 = 1
                _ = v27
            }
        }
        _ = v13
    } else {
        for fi28 := 0; fi28 < 4; fi28++ {
            v12 += 65535
            switch ("\n\t" + "hello") {
            case "":
                var v29 bool = (!(72 == 34))
                _ = v29
            case "foo bar":
                for fi30 := 0; fi30 < 17; fi30++ {
                    switch 127 {
                    case -446:
                        v31 := GenS5{F0: 100, F1: 1, F2: (!true)}
                        _ = v31
                    }
                    v12 = v12
                    fmt.Println(fi28, p0, fi30)
                    _ = fi30
                }
            case "a":
                fmt.Println(p1, p0, v12)
            }
            for fi32 := 0; fi32 < 10; fi32++ {
                fmt.Println(fi28, v12, v12)
                fmt.Println(v12, p1, v12)
                v12 = v12
                _ = fi32
            }
            _ = fi28
        }
        var v33 uint16 = 2
        var v34 int8 = -128
        _ = v33
        _ = v34
    }
    _ = p0
    _ = p1
    _ = v12
    return 255
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    var v4 uint16 = 42
    v5 := make(map[string]uint32)
    switch "a" {
    case "a":
        if false {
            var v35 int64 = -9223372036854775808
            v35 = v35
            _ = v35
        } else {
            fmt.Println(v4)
            var v36 string = v1
            _ = v36
        }
    }
    v37 := []int{-1, 1}
    fmt.Println(v1)
    fmt.Println(v1, v4)
    var v38 uint16 = (v4 + 753)
    v1 = ("" + v1)
    var v39 uint = 65535
    for fi40 := 0; fi40 < 5; fi40++ {
        if (((127 > fi40) && false) && (!(!(false && true)))) {
            v38 = (^(v38))
            var v41 float64 = 3.14159
            _ = v41
        } else {
            var v42 int = -423
            if len(v37) > 0 {
                v37[0] = 128
            }
            var v43 float32 = 1e10
            _ = v42
            _ = v43
        }
        v37 = append(v37, (fi40 | 32767))
        v1 = (v1 + (v1 + ((v1 + v1) + ("a" + v1))))
        _ = fi40
    }
    for fi44 := 0; fi44 < 13; fi44++ {
        var v45 int = (fi44 ^ fi44)
        _ = fi44
        _ = v45
    }
    v46 := GenS6{F0: -542, F1: 676}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v46
    _ = v5
}
