package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS3 struct { F0 int; F1 bool; F2 int16 }
type GenS4 struct { F0 string; F1 int32; F2 int32 }
type GenS5 struct { F0 uint8; F1 bool }
func genFn0(p0 int8, p1 int8, p2 int16) int8 {
    v4 := GenS1{F0: -1}
    v5 := GenS1{F0: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 7
}

func main() {
    var v1 int64 = 707
    var v2 uint64 = 2
    v3 := []string{}
    if ((!false) || ((true && false) && false)) {
        var v6 int8 = genFn0(100, -128, 32767)
        _ = v6
    }
    v2 = v2
    if true {
        fmt.Println(v2, v2)
        var v7 float64 = 0.1
        _ = v7
    }
    v2 = (v2 / (((v2 & v2)) | 1))
    v1 -= v1
    fmt.Println(v1, v1)
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint = 177
        v1 += (v1 | (v1 % (((v1 >> (((v9 + v9)) & 63))) | 1)))
        var v10 uint32 = 952
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 = v1
    var v11 string = "hello"
    fmt.Println(v11, v2)
    var v12 bool = ((0 == 984) || (!(0 == 1)))
    v13 := GenS2{F0: 4294967295, F1: 303, F2: (v2 % ((v2) | 1))}
    if v12 {
        for fi14 := 0; fi14 < 4; fi14++ {
            switch "hello" {
            case "a":
                v11 = (v11 + "\n\t")
            case "foo bar":
                v2 = v2
            default:
                var v15 int32 = -2147483648
                _ = v15
            }
            _ = fi14
        }
        switch "hello" {
        case "":
            v3 = append(v3, "")
        case "foo bar":
            for fi16 := 0; fi16 < 18; fi16++ {
                switch int(v1) {
                case 995:
                    var v17 int64 = v1
                    _ = v17
                }
                v1 = -676
                for fi18 := 0; fi18 < 19; fi18++ {
                    v11 = "a"
                    _ = len(v3)
                    _ = fi18
                }
                _ = fi16
            }
        case "hello":
            for fi19 := 0; fi19 < 20; fi19++ {
                switch (fi19 >> ((42) & 63)) {
                case 42:
                    v20 := GenS3{F0: fi19, F1: (true && (-1 < 0)), F2: -375}
                    _ = v20
                case 2147483647:
                    v3 = append(v3, ("" + v11))
                case 127:
                    for fi21 := 0; fi21 < 18; fi21++ {
                        v11 = v11
                        var v22 int16 = 0
                        _ = fi21
                        _ = v22
                    }
                }
                _ = fi19
            }
        default:
            v2 -= 42
        }
        v23 := GenS4{F0: v11, F1: -708, F2: 0}
        _ = v23
    } else {
        if (!v12) {
            fmt.Println(v1, v12, v12)
            var v24 int64 = (v1 & v1)
            _ = len(v3)
            _ = v24
        }
        switch ("hello" + "") {
        case "a":
            v25 := GenS4{F0: "", F1: 42, F2: -1}
            _ = v25
        case "\n\t":
            var v26 int = -686
            _ = v26
        default:
            if (42 != -32768) {
                for fi27 := 0; fi27 < 17; fi27++ {
                    if v12 {
                        v1 = 677
                        v28 := GenS1{F0: -1}
                        _ = v28
                    }
                    v29 := GenS4{F0: v11, F1: 2147483647, F2: -324}
                    if true {
                        switch int(v1) {
                        case 333:
                            v11 = (v11 + "foo bar")
                        case 1:
                            v3 = append(v3, "foo bar")
                        }
                        v3 = append(v3, "a")
                    }
                    _ = fi27
                    _ = v29
                }
            } else {
                var v30 int8 = -1
                var v31 int = 42
                for fi32 := 0; fi32 < 7; fi32++ {
                    v31 = v31
                    if true {
                        v33 := GenS5{F0: 255, F1: (false || (!(42 > 32767)))}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v33
                    } else {
                        _ = len(v3)
                    }
                    _ = fi32
                }
                _ = v30
                _ = v31
            }
        }
    }
    var v34 int = 546
    fmt.Println(v11, v11)
    var v35 int16 = 1000
    for fi36 := 0; fi36 < 18; fi36++ {
        fmt.Println(v12)
        _ = fi36
    }
    if len(v3) > 0 {
        v3[0] = (v11 + "hello")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v35
}
