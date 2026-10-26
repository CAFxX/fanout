package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 string; F2 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 int8; F2 uint16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 string }
func genFn0(p0 int8, p1 int32, p2 bool) int8 {
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            switch ("hello" + "foo bar") {
            case "\n\t":
                if p2 {
                    var v7 int64 = -345
                    _ = v7
                } else {
                    var v8 uint32 = 536
                    v9 := []int{int(p0)}
                    switch (fi6 + (fi6 ^ fi5)) {
                    case 9223372036854775807:
                        p0 = p0
                    case 793:
                        fmt.Println(p0, v8)
                    }
                    _ = v8
                    _ = v9
                }
            default:
                v10 := []int{42, (fi5 | -859), (fi5 | (fi5 & fi6))}
                _ = v10
            }
            fmt.Println(p2)
            v11 := []int{fi6, fi5}
            if false {
                v12 := GenS1{F0: 65535}
                v13 := GenS2{F0: 864, F1: 18446744073709551615}
                _ = v12
                _ = v13
            } else {
                var v14 uint = 42
                v14 += 42
                _ = v14
            }
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = fi6
            _ = v11
        }
        v15 := GenS3{F0: (fi5 * fi5)}
        switch "foo bar" {
        case "hello":
            v16 := GenS1{F0: 150}
            _ = v16
        case "foo bar":
            for fi17 := 0; fi17 < 11; fi17++ {
                fmt.Println(p1)
                p1 += (p1 - -2147483648)
                if ((989 > 4294967295) || (!p2)) {
                    v18 := []int{}
                    fmt.Println(p0, fi5)
                    p1 = (p1 * -1)
                    _ = v18
                }
                _ = fi17
            }
        default:
            for fi19 := 0; fi19 < 12; fi19++ {
                p0 -= -72
                v20 := GenS4{F0: 1, F1: ("" + ""), F2: (fi5 % (((-(fi5))) | 1))}
                _ = fi19
                _ = v20
            }
        }
        _ = fi5
        _ = v15
    }
    p1 = 911
    v21 := GenS2{F0: 2, F1: 354}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return (p0 + p0)
}

func main() {
    var v1 int64 = -819
    var v2 int64 = (v1 | -51)
    var v3 int16 = 961
    v4 := []uint32{65535, 2, 2147483648, 0}
    var v22 int16 = v3
    switch ("hello" + "foo bar") {
    case "\n\t":
        for fi23 := 0; fi23 < 4; fi23++ {
            var v24 float32 = 1.5
            for fi25 := 0; fi25 < 13; fi25++ {
                if (v24 <= 1.5) {
                    var v26 uint32 = 2147483647
                    if (v22 >= (v3 & (v22 | (v22 >> ((964) & 63))))) {
                        var v27 int8 = int8(v2)
                        if ("" != (("foo bar" + (("\n\t" + "hello") + "\n\t")) + "hello")) {
                            if len(v4) > 0 {
                                v4[0] = 958
                            }
                        }
                        _ = v27
                    } else {
                        switch "a" {
                        case "foo bar":
                            switch fi23 {
                            case 611:
                                v28 := GenS1{F0: 0}
                                _ = v28
                            case 9223372036854775807:
                                var v29 int64 = (v2 << ((255) & 63))
                                _ = v29
                            }
                        case "":
                            switch ("\n\t" + "") {
                            case "\n\t":
                                var v30 uint16 = 2
                                _ = v30
                            case "hello":
                                var v31 uint64 = 42
                                _ = v31
                            case "foo bar":
                                if len(v4) > 0 {
                                    v4[0] = v26
                                }
                            }
                        case "\n\t":
                            if true {
                                v4 = append(v4, (v26 & v26))
                            }
                        default:
                            v22 = -32768
                        }
                        if (-2147483648 <= -1) {
                            v32 := GenS5{F0: 756}
                            fmt.Println(v3, v24)
                            _ = v32
                        } else {
                            var v33 float32 = v24
                            switch ((("\n\t" + ("a" + ("hello" + ""))) + "") + "a") {
                            case "foo bar":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case "hello":
                                var v34 bool = (0 == -128)
                                _ = v34
                            case "a":
                                fmt.Println(v3, v26)
                            }
                            fmt.Println(fi25)
                            if len(v4) > 0 {
                                v4[0] = (v26 + v26)
                            }
                            var v35 uint16 = 996
                            _ = v33
                            _ = v35
                        }
                        v36 := GenS6{F0: int16(v26), F1: 127, F2: 94}
                        fmt.Println(fi25)
                        _ = v36
                    }
                    _ = v26
                } else {
                    var v37 float64 = (1.5 + 0.1)
                    v37 -= (-2.25 * v37)
                    if len(v4) > 0 {
                        v4[0] = 2147483647
                    }
                    _ = v37
                }
                _ = fi25
            }
            _ = fi23
            _ = v24
        }
    default:
        v3 -= 0
    }
    var v38 uint16 = 42
    var v39 int8 = 100
    fmt.Println(v38)
    var v40 float64 = (1.5 / 0.1)
    var v41 int8 = v39
    v42 := GenS4{F0: 0, F1: "hello", F2: 32767}
    var v43 uint8 = 1
    var v44 int16 = (v22 | (v3 ^ 152))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v2, v1)
    for fi45 := 0; fi45 < 13; fi45++ {
        var v46 float32 = (3.14159 + 3.14159)
        v47 := GenS7{F0: 2}
        _ = fi45
        _ = v46
        _ = v47
    }
    v4 = append(v4, 2147483648)
    var v48 bool = (true && true)
    v43 = v43
    for fi49 := 0; fi49 < 8; fi49++ {
        v39 = (v41 >> ((42) & 63))
        for fi50 := 0; fi50 < 11; fi50++ {
            _ = len(v4)
            if (!(428 < 1)) {
                var v51 uint16 = v38
                v38 = (^(v38))
                if (65535 <= 1) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = len(v4)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = v51
            } else {
                _ = len(v4)
                v52 := GenS8{F0: "hello"}
                _ = v52
            }
            _ = fi50
        }
        _ = fi49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v48
}
