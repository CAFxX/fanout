package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 bool; F2 int64 }
type GenS4 struct { F0 bool; F1 int8; F2 uint16 }
func genFn0(p0 uint16) string {
    switch 128 {
    case -2147483648:
        var v5 int64 = 430
        _ = v5
    case 128:
        switch 4294967295 {
        case 32767:
            switch 1 {
            case 252:
                var v6 int16 = 1
                _ = v6
            case 128:
                for fi7 := 0; fi7 < 2; fi7++ {
                    p0 = p0
                    var v8 string = ""
                    var v9 uint = 255
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            case -909:
                v10 := GenS1{F0: 4294967295, F1: 1}
                _ = v10
            default:
                for fi11 := 0; fi11 < 20; fi11++ {
                    var v12 bool = (false || false)
                    p0 -= p0
                    p0 -= p0
                    _ = fi11
                    _ = v12
                }
            }
        }
    }
    var v13 int64 = 0
    _ = p0
    _ = v13
    return ""
}
func genFn1() int32 {
    v14 := []int{9223372036854775807, 127}
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 int = 952
        for fi17 := 0; fi17 < 6; fi17++ {
            if len(v14) > 0 {
                v14[0] = fi17
            }
            _ = fi17
        }
        for fi18 := 0; fi18 < 8; fi18++ {
            if false {
                var v19 int16 = 0
                if ((526 >= 937) || false) {
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    if true {
                        v14 = append(v14, (v16 << ((42) & 63)))
                        var v20 bool = false
                        _ = v20
                    } else {
                        fmt.Println(fi18, v16)
                        v21 := GenS2{F0: 0}
                        v22 := GenS1{F0: 447, F1: -38}
                        _ = v21
                        _ = v22
                    }
                    switch "hello" {
                    case "":
                        v16 = (v16 - -566)
                    case "foo bar":
                        if len(v14) > 0 {
                            v14[0] = fi18
                        }
                    }
                } else {
                    v23 := GenS3{F0: "\n\t", F1: (((!true) && false) && false), F2: -1}
                    v14 = append(v14, (fi15 << ((18446744073709551615) & 63)))
                    _ = v23
                }
                _ = v19
            }
            _ = len(v14)
            if (!(-2.25 <= float64(1.5))) {
                if false {
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                }
                var v24 int16 = -1
                v25 := GenS2{F0: -184}
                _ = v24
                _ = v25
            } else {
                v16 = (fi15 * v16)
                var v26 float32 = float32(((-2.25 + float64(-2.25)) * 3.14159))
                _ = v26
            }
            fmt.Println(fi18, fi18, fi18)
            _ = fi18
        }
        _ = fi15
        _ = v16
    }
    var v27 uint16 = 66
    _ = v14
    _ = v27
    return -2147483648
}

func main() {
    var v1 string = "foo bar"
    var v2 uint64 = 1
    v3 := []uint32{2147483647, 42}
    var v4 float64 = 1e10
    v4 = (v4 / (v4 - 3.14159))
    for fi28 := 0; fi28 < 7; fi28++ {
        var v29 uint8 = 0
        _ = fi28
        _ = v29
    }
    var v30 uint32 = 2
    if true {
        var v31 int = 4294967295
        _ = v31
    } else {
        fmt.Println(v2)
        for fi32 := 0; fi32 < 14; fi32++ {
            if true {
                v33 := GenS4{F0: (v1 >= v1), F1: 47, F2: 32767}
                _ = len(v3)
                _ = v33
            }
            _ = fi32
        }
        v3 = append(v3, (v30 & v30))
    }
    v30 -= (v30 + (v30 >> ((65535) & 63)))
    fmt.Println(v1, v1)
    for fi34 := 0; fi34 < 12; fi34++ {
        v30 = (v30 * (v30 / (((v30 * v30)) | 1)))
        v35 := GenS3{F0: v1, F1: true, F2: -635}
        _ = fi34
        _ = v35
    }
    var v36 uint32 = (v30 % (((v30 * (v30 + 1))) | 1))
    v1 = "\n\t"
    switch "a" {
    case "foo bar":
        _ = len(v3)
    case "\n\t":
        v37 := GenS2{F0: -863}
        _ = v37
    }
    var v38 uint32 = (v36 / (((v30 << ((255) & 63))) | 1))
    if true {
        v39 := GenS3{F0: "\n\t", F1: ((((v1 + ("\n\t" + (v1 + v1))) + v1) + ((v1 + "a") + v1)) < v1), F2: 1}
        _ = v39
    }
    for fi40 := 0; fi40 < 5; fi40++ {
        if false {
            v41 := GenS4{F0: false, F1: 1, F2: 42}
            v4 -= ((float64((3.14159 - 1.5)) - v4) * (v4 + v4))
            _ = v41
        }
        _ = fi40
    }
    for fi42 := 0; fi42 < 20; fi42++ {
        fmt.Println(v36, v1)
        fmt.Println(v30, v36, v4)
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v38
    _ = v4
}
