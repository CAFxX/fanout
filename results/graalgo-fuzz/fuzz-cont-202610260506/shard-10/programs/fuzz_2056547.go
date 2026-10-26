package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int16; F1 int; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint8; F2 int }
type GenS6 struct { F0 int64; F1 uint32 }
func genFn0(p0 int8, p1 int8, p2 int) uint {
    switch 4294967295 {
    case -9223372036854775808:
        v5 := GenS1{F0: 0, F1: 98, F2: 1}
        _ = v5
    case -184:
        v6 := GenS1{F0: 255, F1: 2147483648, F2: 122}
        _ = v6
    case 2147483647:
        p1 = (p0 << ((65535) & 63))
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}
func genFn1(p0 uint64) int64 {
    fmt.Println(p0)
    var v7 bool = (2 <= 42)
    v7 = false
    fmt.Println(p0, p0)
    switch "a" {
    case "a":
        var v8 uint32 = 2147483648
        _ = v8
    case "":
        v7 = (!(false || (0 == -36)))
    case "\n\t":
        switch (("\n\t" + (("\n\t" + "") + "\n\t")) + "hello") {
        case "a":
            if v7 {
                var v9 uint8 = 2
                switch "hello" {
                case "":
                    var v10 bool = false
                    _ = v10
                case "a":
                    v9 = 1
                }
                if true {
                    if (-51 != 62) {
                        p0 += 57
                    }
                    var v11 float64 = 1e10
                    _ = v11
                } else {
                    if v7 {
                        var v12 int = 589
                        var v13 int = v12
                        _ = v12
                        _ = v13
                    } else {
                        v9 += (v9 + (v9 & (v9 << ((0) & 63))))
                        fmt.Println(p0, v7, v7)
                    }
                    var v14 float64 = (1.5 - float64(1e10))
                    for fi15 := 0; fi15 < 10; fi15++ {
                        switch "\n\t" {
                        case "a":
                            for fi16 := 0; fi16 < 12; fi16++ {
                                fmt.Println(v14, p0, fi16)
                                if ("a" == "hello") {
                                    v17 := []int{fi15}
                                    v7 = (!false)
                                    _ = v17
                                }
                                fmt.Println(fi15)
                                v18 := GenS2{F0: 1000, F1: fi15, F2: 1}
                                v14 -= (float64(0.1) - (v14 - ((3.14159 * ((v14 - (((-2.25 * ((v14 / v14) * 1e10)) + float64((0.0 / (0.0 + (float32(((1.5 + v14) - (1.5 / 0.1))) * 1.5))))) + -2.25)) + v14)) * (0.1 / ((v14 - v14) + (v14 + float64(0.1)))))))
                                var v19 int64 = 9223372036854775807
                                _ = fi16
                                _ = v18
                                _ = v19
                            }
                        case "foo bar":
                            var v20 string = "hello"
                            _ = v20
                        case "":
                            v21 := []int{(fi15 & fi15)}
                            _ = v21
                        default:
                            fmt.Println(fi15, v7)
                        }
                        if (42 != 0) {
                            v9 += v9
                        }
                        _ = fi15
                    }
                    _ = v14
                }
                _ = v9
            } else {
                v22 := GenS2{F0: 673, F1: 9223372036854775807, F2: 313}
                if (1 >= 42) {
                    switch 2147483647 {
                    case -589:
                        if true {
                            v23 := GenS1{F0: 127, F1: 0, F2: 2}
                            _ = v23
                        } else {
                            var v24 int16 = 32767
                            p0 -= 42
                            _ = v24
                        }
                    case 988:
                        v25 := []int{-102, -2147483648}
                        _ = v25
                    case 32767:
                        var v26 uint = 1
                        _ = v26
                    }
                    v27 := []int{127, 2147483647}
                    p0 = 18446744073709551615
                    _ = v27
                } else {
                    v7 = (!(2147483647 < int32(p0)))
                    v28 := GenS1{F0: 0, F1: 4294967295, F2: 128}
                    switch ((("foo bar" + "foo bar") + "a") + "") {
                    case "a":
                        v29 := GenS1{F0: 10, F1: 0, F2: 42}
                        _ = v29
                    case "foo bar":
                        if v7 {
                            v30 := GenS3{F0: (false && (33 >= 127))}
                            _ = v30
                        } else {
                            v31 := []int{2147483647, 561, 32767, -468}
                            v32 := GenS2{F0: 1000, F1: 127, F2: 2}
                            v7 = (!(!false))
                            _ = v31
                            _ = v32
                        }
                    }
                    _ = v28
                }
                fmt.Println(v7, p0)
                _ = v22
            }
        }
    default:
        v33 := []int{-711, 1, 4294967295, -987}
        _ = v33
    }
    _ = p0
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 int32 = -305
    var v2 string = "foo bar"
    var v3 float32 = ((0.0 + -2.25) / ((-2.25 * (-2.25 - 1.5)) * 1e10))
    var v4 uint16 = 1
    switch ("hello" + v2) {
    case "\n\t":
        for fi34 := 0; fi34 < 18; fi34++ {
            v1 -= 0
            v3 = v3
            v35 := []int{}
            _ = fi34
            _ = v35
        }
    case "a":
        v36 := []int{-52, 32767, 127, 980}
        _ = v36
    case "foo bar":
        switch v2 {
        case "a":
            var v37 int16 = 42
            _ = v37
        case "foo bar":
            v4 += (v4 & (v4 | 749))
        case "\n\t":
            if (!(1 != (v1 / ((v1) | 1)))) {
                v4 += v4
                v38 := []int{998, 582, 32767, -2147483648}
                if (!(!(!(!(!(!(219 >= 870))))))) {
                    if false {
                        var v39 int = -256
                        v40 := GenS2{F0: 1, F1: (v39 / (((^(v39))) | 1)), F2: 2147483647}
                        _ = v39
                        _ = v40
                    } else {
                        if ((v2 + v2) > v2) {
                            v1 = (v1 % (((^(v1))) | 1))
                        } else {
                            for fi41 := 0; fi41 < 12; fi41++ {
                                if true {
                                    switch v2 {
                                    case "a":
                                        var v42 uint16 = (^(v4))
                                        _ = v42
                                    case "foo bar":
                                        var v43 uint16 = (v4 << ((65535) & 63))
                                        _ = v43
                                    default:
                                        v44 := GenS4{F0: 18446744073709551615, F1: 295}
                                        _ = v44
                                    }
                                    if len(v38) >= 2 {
                                        v38 = v38[1:len(v38)]
                                    }
                                }
                                _ = fi41
                            }
                        }
                        v45 := GenS3{F0: false}
                        _ = v45
                    }
                    if true {
                        var v46 bool = (true && (true && (!true)))
                        _ = v46
                    } else {
                        if (!false) {
                            for fi47 := 0; fi47 < 19; fi47++ {
                                if len(v38) >= 2 {
                                    v38 = v38[1:len(v38)]
                                }
                                fmt.Println(v3, v4, fi47)
                                _ = fi47
                            }
                            v48 := GenS5{F0: 0, F1: 61, F2: 32768}
                            var v49 uint16 = (v4 / ((v4) | 1))
                            _ = v48
                            _ = v49
                        }
                        v3 -= -2.25
                    }
                } else {
                    if true {
                        v1 += (v1 & (v1 - v1))
                        v4 = v4
                    }
                }
                _ = v38
            } else {
                fmt.Println(v3)
                if true {
                    fmt.Println(v1, v1, v1)
                    v4 -= 1000
                }
                v50 := []int{}
                _ = v50
            }
        default:
            v1 = -1
        }
    default:
        switch 9223372036854775807 {
        case -9223372036854775808:
            var v51 int8 = -35
            _ = v51
        case 32768:
            if (false || (v4 > (^(v4)))) {
                var v52 uint8 = 42
                fmt.Println(v1, v4)
                _ = v52
            }
        default:
            v53 := GenS6{F0: genFn1(931), F1: 898}
            _ = v53
        }
    }
    var v54 uint16 = v4
    fmt.Println(v54, v2, v54)
    v55 := []int{983, 127}
    v3 = 1e10
    fmt.Println(v54)
    if false {
        var v56 uint16 = (v54 & v4)
        v2 = (((v2 + "") + ("foo bar" + v2)) + v2)
        _ = v56
    }
    for fi57 := 0; fi57 < 3; fi57++ {
        var v58 bool = false
        if v58 {
            v3 = 1.5
            v1 += v1
        }
        _ = fi57
        _ = v58
    }
    v1 -= v1
    fmt.Println(v4, v4, v54)
    switch (v2 + v2) {
    case "\n\t":
        var v59 uint = 0
        _ = v59
    case "foo bar":
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v55:", len(v55), v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v55
}
