package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint; F1 int; F2 uint8 }
type GenS5 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int8 }
func genFn0(p0 int) int {
    switch "hello" {
    case "foo bar":
        var v6 int64 = int64(p0)
        _ = v6
    default:
        var v7 uint8 = 2
        _ = v7
    }
    fmt.Println(p0)
    switch (p0 | (p0 / ((p0) | 1))) {
    case 128:
        v8 := GenS1{F0: 255}
        _ = v8
    case 42:
        var v9 int16 = 42
        _ = v9
    }
    var v10 uint32 = 4294967295
    var v11 int32 = 0
    _ = p0
    _ = v10
    _ = v11
    return (-(p0))
}
func genFn1() uint64 {
    var v12 uint32 = 1
    if true {
        v13 := []int{}
        _ = v13
    } else {
        v14 := []int{genFn0(986), -452, -2147483648, -9223372036854775808}
        var v15 uint8 = 1
        var v16 uint64 = 984
        _ = v14
        _ = v15
        _ = v16
    }
    v17 := GenS2{F0: 739, F1: 171, F2: 0}
    _ = v12
    _ = v17
    return 4294967295
}

func main() {
    var v1 uint = 47
    var v2 uint8 = 194
    var v3 string = ("" + "hello")
    v4 := make(map[string]int)
    var v5 int = 4294967295
    switch v3 {
    case "a":
        var v18 uint16 = 97
        _ = v18
    case "\n\t":
        v19 := GenS3{F0: (v1 << ((v1) & 63)), F1: genFn0((-(v5)))}
        _ = v19
    }
    fmt.Println(v3)
    var v20 bool = true
    v21 := []int{}
    var v22 int8 = 1
    v22 = (v22 << ((v1) & 63))
    switch v5 {
    case 328:
        v22 -= v22
    case -9223372036854775808:
        if ((!false) || (!v20)) {
            fmt.Println(v22)
            var v23 string = ("hello" + (("a" + "") + "a"))
            _ = v23
        }
    default:
        v20 = false
    }
    var v24 bool = (!v20)
    var v25 uint = 42
    switch (("" + (v3 + (v3 + "hello"))) + v3) {
    case "\n\t":
        var v26 int64 = 1
        _ = v26
    case "a":
        _ = len(v21)
    case "foo bar":
        for fi27 := 0; fi27 < 10; fi27++ {
            v20 = false
            v24 = (257 <= 0)
            var v28 int64 = 292
            _ = fi27
            _ = v28
        }
    default:
        if true {
            if v20 {
                v29 := GenS4{F0: 65535, F1: genFn0(v5), F2: v2}
                _ = len(v21)
                switch "a" {
                case "\n\t":
                    v4["hello"] = v5
                case "foo bar":
                    var v30 int32 = 568
                    _ = v30
                default:
                    switch "a" {
                    case "":
                        var v31 int8 = (v22 - v22)
                        _ = v31
                    case "a":
                        v32 := GenS4{F0: (^(v1)), F1: 42, F2: v2}
                        _ = v32
                    case "hello":
                        v2 = (v2 % (((v2 / ((161) | 1))) | 1))
                    default:
                        v33 := GenS5{F0: 65535, F1: 1, F2: 32767}
                        _ = v33
                    }
                }
                _ = v29
            }
        } else {
            v2 = (v2 >> (((v1 - 0)) & 63))
            v22 = v22
            var v34 int32 = 589
            _ = v34
        }
    }
    v35 := GenS3{F0: v25, F1: (v5 << ((v25) & 63))}
    for fi36 := 0; fi36 < 7; fi36++ {
        v37 := GenS3{F0: v25, F1: (fi36 - v5)}
        v3 = v3
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 1; fi38++ {
        for fi39 := 0; fi39 < 1; fi39++ {
            if (0 != int16(v25)) {
                var v40 uint64 = 4294967295
                _ = v40
            } else {
                switch int(v1) {
                case -2147483648:
                    fmt.Println(v3, v2)
                default:
                    v24 = ((true && (int(v22) <= (fi38 >> (((^(v1))) & 63)))) && ((v2 % (((v2 % ((v2) | 1))) | 1)) == 42))
                }
                fmt.Println(v3, v5, v25)
            }
            v3 = (v3 + v3)
            fmt.Println(v22)
            _ = fi39
        }
        if ((-(v22)) <= v22) {
            v24 = (v25 >= v1)
        } else {
            if false {
                switch genFn0(fi38) {
                case 0:
                    for fi41 := 0; fi41 < 15; fi41++ {
                        v3 = (((v3 + "foo bar") + ("a" + "a")) + "hello")
                        v24 = false
                        v42 := GenS3{F0: (v25 & (v25 & v1)), F1: 32768}
                        if len(v21) > 0 {
                            v21[0] = (v5 | (v5 ^ (v5 ^ 95)))
                        }
                        _ = len(v4)
                        _ = fi41
                        _ = v42
                    }
                }
                switch (fi38 | 1) {
                case 1:
                    if (!(v3 < v3)) {
                        v22 = (^(v22))
                        switch v3 {
                        case "hello":
                            switch v3 {
                            case "a":
                                _ = len(v4)
                            case "\n\t":
                                v2 = v2
                            }
                        case "a":
                            v43 := GenS2{F0: (v25 % (((v25 ^ 0)) | 1)), F1: 255, F2: 748}
                            _ = v43
                        case "\n\t":
                            v21 = append(v21, fi38)
                        default:
                            for fi44 := 0; fi44 < 9; fi44++ {
                                if true {
                                    v45 := GenS6{F0: uint64(v2)}
                                    var v46 uint16 = 65535
                                    if false {
                                        v1 = (v25 + (v1 % ((v25) | 1)))
                                        fmt.Println(v3, v22, v20)
                                        v47 := GenS3{F0: 42, F1: (fi44 + (v5 % ((fi44) | 1)))}
                                        _ = v47
                                    } else {
                                        if (1e10 <= (1.5 / (1e10 + 3.14159))) {
                                            v3 = v3
                                            var v48 uint64 = 1
                                            _ = v48
                                        } else {
                                            var v49 int = fi38
                                            switch 2147483647 {
                                            case 32768:
                                                switch 4294967295 {
                                                case 4294967295:
                                                    var v50 string = (v3 + "hello")
                                                    _ = v50
                                                case -9223372036854775808:
                                                    fmt.Println(v1, fi44, v3)
                                                case 32768:
                                                    v51 := GenS7{F0: v22}
                                                    _ = v51
                                                default:
                                                    v1 -= v25
                                                }
                                            case 1:
                                                if len(v21) >= 2 {
                                                    v21 = v21[1:len(v21)]
                                                }
                                            default:
                                                var v52 uint64 = uint64(v1)
                                                _ = v52
                                            }
                                            if (v24 || false) {
                                                if (-2147483648 > -732) {
                                                    _ = len(v21)
                                                } else {
                                                    var v53 string = v3
                                                    switch (^(v5)) {
                                                    case 4294967295:
                                                        var v54 string = "a"
                                                        _ = v54
                                                    case 32768:
                                                        fmt.Println(v24, v20, v2)
                                                    default:
                                                        fmt.Println(v22, v5)
                                                        v55 := GenS1{F0: v2}
                                                        _ = v55
                                                    }
                                                    _ = v53
                                                }
                                            } else {
                                                v56 := GenS4{F0: v25, F1: (^(fi38)), F2: (v2 >> ((v25) & 63))}
                                                _ = v56
                                            }
                                            _ = v49
                                        }
                                        v20 = (float32(0.0) >= 3.14159)
                                    }
                                    _ = v45
                                    _ = v46
                                } else {
                                    var v57 uint16 = 491
                                    _ = v57
                                }
                                delete(v4, ((v3 + "\n\t") + "foo bar"))
                                _ = fi44
                            }
                        }
                    } else {
                        _ = len(v21)
                    }
                case 32767:
                    v22 += v22
                default:
                    var v58 int = fi38
                    _ = v58
                }
                v3 = "hello"
            }
        }
        _ = fi38
    }
    v5 += 94
    fmt.Println(v22)
    switch "hello" {
    case "":
        var v59 int = (v5 | 0)
        _ = v59
    case "\n\t":
        for fi60 := 0; fi60 < 1; fi60++ {
            v22 = 42
            _ = fi60
        }
    case "foo bar":
        v2 = (^(v2))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
