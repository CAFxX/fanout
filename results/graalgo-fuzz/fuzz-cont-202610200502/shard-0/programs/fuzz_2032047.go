package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 uint }
func genFn0(p0 uint8, p1 uint64, p2 uint16) bool {
    var v5 uint8 = (p0 >> ((1) & 63))
    v6 := GenS1{F0: (p0 ^ 255)}
    var v7 int32 = -337
    v8 := GenS1{F0: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (95 == 4294967295)
}
func genFn1() uint32 {
    var v9 uint = 0
    v9 = v9
    if false {
        if (false && true) {
            v9 = (v9 & (^(v9)))
        } else {
            var v10 int32 = 2147483647
            fmt.Println(v10, v10)
            var v11 uint64 = 1
            _ = v10
            _ = v11
        }
        v12 := []int{4294967295, 42, -406, 96}
        _ = v12
    } else {
        v9 -= v9
        if (uint8(v9) != 128) {
            for fi13 := 0; fi13 < 15; fi13++ {
                switch "\n\t" {
                case "\n\t":
                    v9 -= 42
                case "hello":
                    switch (fi13 | (^(fi13))) {
                    case 0:
                        if (true || (65535 == 65535)) {
                            v14 := GenS1{F0: 195}
                            v9 += 896
                            var v15 int8 = 96
                            _ = v14
                            _ = v15
                        } else {
                            fmt.Println(v9, v9)
                            for fi16 := 0; fi16 < 3; fi16++ {
                                v17 := []int{(fi13 - fi16)}
                                _ = fi16
                                _ = v17
                            }
                        }
                    case -1:
                        for fi18 := 0; fi18 < 9; fi18++ {
                            if true {
                                v9 = (^(v9))
                                v19 := GenS1{F0: 128}
                                _ = v19
                            }
                            v20 := []int{fi18, (fi18 ^ (fi13 - (fi18 & 474))), 42, 384}
                            _ = fi18
                            _ = v20
                        }
                    case 42:
                        var v21 int16 = 417
                        _ = v21
                    default:
                        v22 := GenS1{F0: 255}
                        _ = v22
                    }
                default:
                    fmt.Println(v9, v9)
                }
                v9 = (v9 ^ v9)
                v9 -= v9
                _ = fi13
            }
            v9 -= v9
            for fi23 := 0; fi23 < 3; fi23++ {
                var v24 string = ""
                _ = fi23
                _ = v24
            }
        } else {
            fmt.Println(v9)
            var v25 int32 = -2147483648
            v26 := GenS1{F0: 42}
            _ = v25
            _ = v26
        }
    }
    fmt.Println(v9)
    _ = v9
    return 42
}

func main() {
    v1 := []int{13}
    var v2 uint16 = 137
    var v3 uint64 = 1
    var v4 int32 = 251
    for fi27 := 0; fi27 < 9; fi27++ {
        switch (("foo bar" + "foo bar") + "\n\t") {
        case "":
            var v28 bool = (v4 != (v4 * v4))
            _ = v28
        default:
            fmt.Println(v2, fi27)
        }
        switch (^(fi27)) {
        case 1:
            fmt.Println(fi27, v3)
        case -9223372036854775808:
            for fi29 := 0; fi29 < 18; fi29++ {
                v30 := GenS2{F0: v2}
                _ = fi29
                _ = v30
            }
        default:
            v31 := GenS1{F0: 127}
            _ = v31
        }
        if (111 >= -10) {
            var v32 string = ("\n\t" + ("hello" + ("a" + "")))
            _ = v32
        } else {
            v2 += 232
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v3 += v3
        }
        _ = fi27
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v33 uint = 1
    if false {
        v34 := GenS3{F0: 100}
        v35 := GenS2{F0: uint16(v3)}
        var v36 uint = (v33 - (v33 << ((v33) & 63)))
        _ = v34
        _ = v35
        _ = v36
    } else {
        if true {
            if false {
                switch "\n\t" {
                case "\n\t":
                    var v37 uint16 = 206
                    _ = v37
                case "hello":
                    for fi38 := 0; fi38 < 18; fi38++ {
                        if len(v1) > 0 {
                            v1[0] = -71
                        }
                        _ = fi38
                    }
                }
            }
            for fi39 := 0; fi39 < 20; fi39++ {
                fmt.Println(v3, v33, fi39)
                _ = fi39
            }
        }
        v40 := GenS4{F0: -82, F1: v33}
        _ = v40
    }
    _ = len(v1)
    if (v3 > (v3 >> ((v33) & 63))) {
        switch ("" + "") {
        case "a":
            v4 -= 73
        case "hello":
            var v41 int16 = -32768
            _ = v41
        case "":
            var v42 int8 = -1
            _ = v42
        default:
            var v43 int16 = 273
            _ = v43
        }
    }
    var v44 float32 = 0.1
    var v45 bool = true
    v46 := GenS4{F0: -1, F1: v33}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v47 string = ""
    var v48 int16 = 32767
    if (!false) {
        switch v47 {
        case "":
            v44 = v44
        case "hello":
            _ = len(v1)
        case "a":
            fmt.Println(v48, v33, v3)
        }
    } else {
        if (42 > 76) {
            if ((!(2147483647 == (v4 % ((596) | 1)))) && (!(v45 && ((v2 | 32767) > (v2 ^ uint16(v33)))))) {
                var v49 int32 = (v4 | 2147483647)
                var v50 uint64 = 1
                _ = v49
                _ = v50
            } else {
                var v51 uint = 736
                _ = v51
            }
        } else {
            v48 += 1000
            var v52 uint32 = uint32(v3)
            _ = v52
        }
        v3 = v3
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
}
