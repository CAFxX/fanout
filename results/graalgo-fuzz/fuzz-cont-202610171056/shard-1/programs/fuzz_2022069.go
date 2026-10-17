package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint16; F2 int32 }
type GenS3 struct { F0 bool; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint8; F1 int32; F2 uint64 }
func genFn0(p0 int) uint16 {
    if (true || (!(178 < 128))) {
        for fi6 := 0; fi6 < 17; fi6++ {
            p0 -= (fi6 - (^(fi6)))
            switch p0 {
            case -506:
                var v7 uint32 = 2147483647
                _ = v7
            default:
                if (65535 == 41) {
                    switch "hello" {
                    case "":
                        if (573 != 42) {
                            v8 := GenS1{F0: 708}
                            _ = v8
                        }
                    case "\n\t":
                        fmt.Println(fi6, fi6)
                    default:
                        fmt.Println(fi6)
                    }
                    var v9 uint32 = 2147483647
                    _ = v9
                }
            }
            _ = fi6
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := []int{fi10, 127, p0}
            p0 = p0
            p0 -= fi10
            _ = fi10
            _ = v11
        }
    }
    p0 += p0
    _ = p0
    return 0
}
func genFn1(p0 int, p1 uint32, p2 int8) string {
    var v12 float64 = float64(-2.25)
    v13 := GenS2{F0: (p0 >> ((18446744073709551615) & 63)), F1: 2, F2: 346}
    for fi14 := 0; fi14 < 5; fi14++ {
        v12 = v12
        _ = fi14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        if (("\n\t" + "hello") > "hello") {
            v16 := GenS3{F0: true, F1: 2147483647, F2: p2}
            _ = v16
        } else {
            switch fi15 {
            case 32767:
                p2 = (p2 + p2)
            default:
                v17 := []int{fi15}
                _ = v17
            }
        }
        p1 = (p1 / ((4294967295) | 1))
        _ = fi15
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        for fi19 := 0; fi19 < 9; fi19++ {
            v20 := []int{p0, (fi18 / (((fi18 & 0)) | 1))}
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return "foo bar"
}

func main() {
    var v1 float64 = 0.0
    var v2 string = "hello"
    var v3 uint8 = 2
    var v4 int = 128
    v5 := make(map[string]uint32)
    var v21 uint64 = 0
    var v22 string = (v2 + "")
    for fi23 := 0; fi23 < 13; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            fmt.Println(v4)
            _ = fi24
        }
        v25 := v5[v2]
        for fi26 := 0; fi26 < 8; fi26++ {
            if (!false) {
                fmt.Println(v22)
                fmt.Println(v2, v1, v22)
            } else {
                v1 -= v1
                if (0 == 32767) {
                    var v27 bool = (790 <= 0)
                    v28 := GenS4{F0: v3, F1: (v3 - v3)}
                    var v29 int16 = -1
                    _ = v27
                    _ = v28
                    _ = v29
                } else {
                    v30 := GenS5{F0: 42, F1: v21}
                    switch (v4 - fi26) {
                    case 0:
                        v31 := []int{(fi23 % ((fi26) | 1))}
                        _ = v31
                    case 32768:
                        fmt.Println(v4, fi26)
                    default:
                        v21 += (v21 - 1)
                    }
                    if true {
                        v5["foo bar"] = 510
                        if false {
                            v32 := v5[v22]
                            _ = v32
                        }
                        delete(v5, (v2 + (("\n\t" + v2) + v22)))
                    }
                    var v33 uint8 = 137
                    switch v2 {
                    case "\n\t":
                        v21 = 609
                    case "hello":
                        if true {
                            fmt.Println(v1, v1)
                        }
                    default:
                        v2 = "foo bar"
                    }
                    v4 += 32767
                    _ = v30
                    _ = v33
                }
                v25 -= (v25 + (v25 - v25))
            }
            var v34 uint8 = (v3 + 66)
            var v35 int64 = 1
            _ = fi26
            _ = v34
            _ = v35
        }
        _ = fi23
        _ = v25
    }
    fmt.Println(v21, v3)
    if false {
        if (false && false) {
            v36 := v5["foo bar"]
            var v37 string = ((v22 + ("\n\t" + ("hello" + v2))) + (("foo bar" + ("hello" + v22)) + ""))
            switch ((v2 + "\n\t") + ("a" + ((v22 + (v2 + v37)) + "a"))) {
            case "":
                v38 := v5[(v2 + v2)]
                _ = v38
            default:
                switch 32767 {
                case 32767:
                    v39 := GenS6{F0: -14}
                    _ = v39
                case 32768:
                    switch "a" {
                    case "hello":
                        var v40 bool = true
                        _ = v40
                    }
                }
            }
            _ = v36
            _ = v37
        } else {
            v41 := v5[v2]
            v21 = (v21 + (^(v21)))
            _ = v41
        }
        v42 := GenS7{F0: 0}
        var v43 string = (("\n\t" + "\n\t") + (("foo bar" + v2) + ((v22 + "\n\t") + v2)))
        _ = v42
        _ = v43
    }
    if false {
        for fi44 := 0; fi44 < 12; fi44++ {
            if false {
                fmt.Println(v4, v2)
            }
            _ = fi44
        }
        fmt.Println(v2, v22, v1)
        var v45 float32 = -2.25
        _ = v45
    }
    _ = len(v5)
    fmt.Println(v1, v22)
    if (v21 >= v21) {
        v46 := GenS8{F0: (v3 - (v3 % ((v3) | 1))), F1: -862, F2: v21}
        if true {
            v2 = "\n\t"
        } else {
            var v47 int32 = -866
            _ = v47
        }
        _ = v46
    } else {
        fmt.Println(v3, v22)
        switch (v4 / ((127) | 1)) {
        case -1:
            delete(v5, v22)
        case -796:
            var v48 float64 = float64(float32(v1))
            _ = v48
        case 931:
            for fi49 := 0; fi49 < 17; fi49++ {
                v4 = (v4 - (v4 >> ((1) & 63)))
                _ = fi49
            }
        default:
            var v50 uint8 = (^(v3))
            _ = v50
        }
        v4 -= v4
    }
    delete(v5, v2)
    switch ((v2 + (("hello" + v2) + ((v22 + v2) + ""))) + "") {
    case "":
        v22 = v2
    case "a":
        v51 := GenS8{F0: (v3 & v3), F1: -1, F2: 1}
        _ = v51
    default:
        var v52 uint8 = (^(v3))
        _ = v52
    }
    var v53 uint8 = (v3 - (v3 >> ((65535) & 63)))
    v4 += v4
    var v54 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
}
