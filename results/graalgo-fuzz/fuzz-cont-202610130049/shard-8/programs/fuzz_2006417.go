package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint8; F2 bool }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int8 }

func main() {
    var v1 int = 32768
    v2 := make(map[uint32]uint32)
    var v3 int8 = -72
    var v4 int32 = -2147483648
    var v5 int64 = 981
    switch v1 {
    case -214:
        for fi6 := 0; fi6 < 9; fi6++ {
            if true {
                var v7 bool = (983 == 32767)
                _ = v7
            }
            _ = fi6
        }
    case 32768:
        v8 := GenS1{F0: "\n\t", F1: (v1 << ((471) & 63)), F2: 2147483648}
        _ = v8
    case 990:
        var v9 int16 = -1
        _ = v9
    }
    switch "foo bar" {
    case "\n\t":
        v3 += v3
    case "foo bar":
        v10 := GenS2{F0: 0, F1: 198, F2: (42 >= 2)}
        _ = v10
    }
    var v11 int16 = 606
    var v12 uint8 = 1
    v13 := []int{v1, 32767}
    switch (v1 | -1) {
    case -9223372036854775808:
        if false {
            for fi14 := 0; fi14 < 11; fi14++ {
                if ((v1 + (fi14 / ((v1) | 1))) <= (fi14 ^ (fi14 % ((v1) | 1)))) {
                    for fi15 := 0; fi15 < 20; fi15++ {
                        v16 := v2[42]
                        v5 = (v5 ^ (v5 << ((362) & 63)))
                        _ = fi15
                        _ = v16
                    }
                    switch v1 {
                    case 128:
                        v3 = int8(v4)
                    case 2147483647:
                        v1 -= 32768
                    case 32768:
                        var v17 uint8 = 128
                        _ = v17
                    default:
                        if len(v13) > 0 {
                            v13[0] = fi14
                        }
                    }
                } else {
                    v18 := GenS1{F0: "hello", F1: fi14, F2: 2}
                    _ = v18
                }
                v3 += (v3 << ((4294967295) & 63))
                if true {
                    v12 += (v12 << ((65535) & 63))
                    v12 = (v12 % (((v12 << ((42) & 63))) | 1))
                }
                _ = fi14
            }
        } else {
            switch ("hello" + "foo bar") {
            case "\n\t":
                v4 -= v4
            case "foo bar":
                v3 -= v3
            }
            if len(v13) > 0 {
                v13[0] = (v1 + v1)
            }
            var v19 uint32 = 2147483648
            _ = v19
        }
    case -1:
        if (505 < 1) {
            v20 := GenS3{F0: 2, F1: true}
            v5 = 0
            _ = v20
        } else {
            var v21 int = 686
            _ = v21
        }
    case -2147483648:
        if true {
            v5 = (-(v5))
            if true {
                fmt.Println(v3, v3)
                switch ("hello" + "\n\t") {
                case "a":
                    fmt.Println(v1)
                case "hello":
                    if false {
                        v22 := GenS3{F0: 42, F1: (32767 != 42)}
                        v23 := GenS2{F0: 0, F1: v12, F2: true}
                        _ = v22
                        _ = v23
                    }
                }
            }
            if false {
                var v24 uint32 = uint32(v12)
                _ = v24
            } else {
                if false {
                    switch "hello" {
                    case "":
                        v25 := GenS2{F0: 65535, F1: (v12 % (((v12 % (((v12 ^ (v12 / ((v12) | 1)))) | 1))) | 1)), F2: false}
                        _ = v25
                    case "foo bar":
                        v26 := GenS4{F0: v3}
                        _ = v26
                    case "a":
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    default:
                        var v27 uint8 = v12
                        _ = v27
                    }
                    var v28 int32 = (v4 + -2147483648)
                    var v29 int64 = 0
                    _ = v28
                    _ = v29
                } else {
                    v30 := GenS2{F0: 42, F1: v12, F2: (!false)}
                    _ = v30
                }
                fmt.Println(v12, v4)
                v5 = v5
            }
        }
    }
    switch (v1 + 32767) {
    case 725:
        for fi31 := 0; fi31 < 20; fi31++ {
            v32 := GenS2{F0: 2147483647, F1: v12, F2: false}
            v33 := GenS3{F0: 2147483647, F1: (!false)}
            _ = fi31
            _ = v32
            _ = v33
        }
    case -2147483648:
        if true {
            v3 = -69
            if len(v13) > 0 {
                v13[0] = v1
            }
        } else {
            if (-2.25 < 1.5) {
                delete(v2, 0)
            }
            switch "hello" {
            case "\n\t":
                switch -1 {
                case 32767:
                    for fi34 := 0; fi34 < 10; fi34++ {
                        fmt.Println(v12)
                        _ = fi34
                    }
                case 2147483647:
                    v4 += (v4 & v4)
                }
            default:
                for fi35 := 0; fi35 < 17; fi35++ {
                    for fi36 := 0; fi36 < 13; fi36++ {
                        v37 := GenS2{F0: 4294967295, F1: 128, F2: false}
                        v38 := GenS3{F0: 611, F1: (!false)}
                        v39 := GenS1{F0: "\n\t", F1: (fi35 - (-(fi35))), F2: 42}
                        _ = fi36
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    _ = fi35
                }
            }
            fmt.Println(v1)
        }
    case 4294967295:
        var v40 int16 = v11
        _ = v40
    }
    for fi41 := 0; fi41 < 14; fi41++ {
        v11 -= (v11 - -922)
        var v42 int8 = (^(v3))
        _ = fi41
        _ = v42
    }
    switch "foo bar" {
    case "\n\t":
        switch (v1 ^ v1) {
        case 0:
            if true {
                if ((v1 * (v1 << ((2) & 63))) >= v1) {
                    fmt.Println(v5)
                    _ = len(v13)
                    for fi43 := 0; fi43 < 9; fi43++ {
                        v44 := v2[62]
                        v13 = append(v13, v1)
                        if (0 == fi43) {
                            v4 -= 2147483647
                            switch -9223372036854775808 {
                            case -9223372036854775808:
                                var v45 uint16 = 65535
                                _ = v45
                            default:
                                var v46 int64 = v5
                                _ = v46
                            }
                            v11 += 1000
                        } else {
                            var v47 bool = true
                            fmt.Println(v4)
                            if (4294967295 < 0) {
                                if v47 {
                                    v11 += v11
                                    _ = len(v2)
                                } else {
                                    if v47 {
                                        fmt.Println(v1)
                                        var v48 int16 = v11
                                        if (!false) {
                                            fmt.Println(v11)
                                            fmt.Println(v12, fi43)
                                            fmt.Println(v11)
                                        }
                                        _ = v48
                                    } else {
                                        var v49 bool = false
                                        v44 = (^(v44))
                                        _ = v49
                                    }
                                    if (18446744073709551615 == 4294967295) {
                                        v4 += v4
                                    } else {
                                        _ = len(v2)
                                        v50 := GenS2{F0: (v44 / (((v44 + (v44 / ((v44) | 1)))) | 1)), F1: (v12 >> ((434) & 63)), F2: (v44 >= (v44 / (((^(v44))) | 1)))}
                                        v11 += v11
                                        _ = v50
                                    }
                                    switch v1 {
                                    case 0:
                                        v47 = v47
                                    default:
                                        var v51 bool = ((v11 >> ((0) & 63)) > -942)
                                        _ = v51
                                    }
                                }
                                v1 += 0
                            } else {
                                var v52 uint64 = 18446744073709551615
                                _ = v52
                            }
                            _ = v47
                        }
                        _ = fi43
                        _ = v44
                    }
                } else {
                    var v53 string = "hello"
                    _ = v53
                }
                if true {
                    v1 += v1
                }
            }
        case 127:
            for fi54 := 0; fi54 < 16; fi54++ {
                fmt.Println(v3, v5)
                fmt.Println(v3, v12)
                _ = fi54
            }
        case -718:
            _ = len(v2)
        }
    case "hello":
        _ = len(v2)
    }
    if len(v13) > 0 {
        v13[0] = 0
    }
    var v55 int32 = int32(v11)
    switch v1 {
    case -522:
        fmt.Println(v11)
    }
    v4 = -967
    fmt.Println(v3, v1)
    fmt.Println(v11)
    switch ("foo bar" + ((("foo bar" + "\n\t") + "foo bar") + "\n\t")) {
    case "\n\t":
        if len(v13) > 0 {
            v13[0] = (v1 >> ((1) & 63))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
}
