package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint64; F1 int; F2 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int; F2 uint16 }
type GenS5 struct { F0 uint32; F1 int; F2 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8; F1 uint32; F2 int }

func main() {
    v1 := []int{}
    v2 := []uint32{}
    v3 := []uint32{1}
    var v4 uint16 = 65535
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 uint64 = 4294967295
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int8 = 66
        _ = fi7
        _ = v8
    }
    fmt.Println(v4)
    var v9 uint32 = 2
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v10 := GenS1{F0: v9, F1: false}
    v9 -= (v9 - (^(v9)))
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            if true {
                v4 -= (^(v4))
                v4 = v4
                fmt.Println(v4)
            }
            var v13 uint8 = 127
            switch (fi12 - (fi12 ^ fi11)) {
            case 42:
                if true {
                    if false {
                        switch 9223372036854775807 {
                        case -2147483648:
                            var v14 float64 = 3.14159
                            _ = v14
                        default:
                            _ = len(v1)
                        }
                        v13 = v13
                    } else {
                        fmt.Println(v4, fi11, fi11)
                        fmt.Println(v4, v9, v13)
                    }
                    v15 := GenS2{F0: 0, F1: fi12, F2: 879}
                    switch ("a" + "hello") {
                    case "hello":
                        var v16 float32 = ((float32((float64(0.0) * 1.5)) - -2.25) + 1.5)
                        _ = v16
                    case "a":
                        v4 = 2
                        fmt.Println(fi12, fi12)
                    default:
                        _ = len(v2)
                    }
                    var v17 uint8 = 255
                    _ = v15
                    _ = v17
                } else {
                    v18 := GenS2{F0: uint64(v4), F1: 32768, F2: 930}
                    _ = v18
                }
            case 2147483647:
                v19 := GenS3{F0: uint8(v9)}
                _ = v19
            case -1:
                if false {
                    v2 = append(v2, (v9 - (v9 - (v9 - v9))))
                    v9 -= 4294967295
                }
            default:
                var v20 uint = 460
                _ = v20
            }
            _ = fi12
            _ = v13
        }
        for fi21 := 0; fi21 < 7; fi21++ {
            v4 = (v4 ^ v4)
            if true {
                var v22 uint64 = 42
                var v23 int32 = -2147483648
                var v24 bool = (!(((-2.25 >= 1.5) && ((0.1 * 1.5) != 0.1)) || (2147483648 >= 1)))
                _ = v22
                _ = v23
                _ = v24
            }
            v25 := GenS4{F0: -1, F1: (fi21 ^ fi11), F2: (v4 & 32767)}
            _ = fi21
            _ = v25
        }
        if true {
            var v26 uint32 = v9
            _ = v26
        } else {
            switch (("a" + "") + ((("" + "\n\t") + ("" + ("" + ""))) + ("a" + "a"))) {
            case "":
                if (2147483647 != (fi11 + fi11)) {
                    if len(v1) > 0 {
                        v1[0] = (-(fi11))
                    }
                    for fi27 := 0; fi27 < 9; fi27++ {
                        fmt.Println(fi11, v9, fi27)
                        _ = fi27
                    }
                    fmt.Println(v9, fi11)
                } else {
                    var v28 uint32 = (v9 ^ 4294967295)
                    if true {
                        var v29 int = 641
                        var v30 int = fi11
                        _ = v29
                        _ = v30
                    } else {
                        v28 += v28
                    }
                    _ = v28
                }
            }
            for fi31 := 0; fi31 < 11; fi31++ {
                v4 -= (v4 ^ v4)
                v32 := GenS1{F0: v9, F1: true}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v33 uint = 1
                _ = fi31
                _ = v32
                _ = v33
            }
            if true {
                var v34 uint8 = 2
                _ = v34
            } else {
                if true {
                    v35 := GenS5{F0: (^(v9)), F1: -614, F2: 100}
                    switch "" {
                    case "":
                        v1 = append(v1, fi11)
                    case "foo bar":
                        switch "a" {
                        case "\n\t":
                            for fi36 := 0; fi36 < 7; fi36++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v37 uint8 = 45
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v38 int = (fi36 << ((2) & 63))
                                if true {
                                    if true {
                                        v38 -= v38
                                        var v39 uint16 = v4
                                        _ = v39
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = v38
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = v38
                                    }
                                } else {
                                    var v40 uint = 679
                                    _ = v40
                                }
                                _ = fi36
                                _ = v37
                                _ = v38
                            }
                        case "foo bar":
                            v9 += (v9 & v9)
                        }
                    }
                    for fi41 := 0; fi41 < 12; fi41++ {
                        var v42 float64 = float64(0.1)
                        v43 := GenS6{F0: ((("hello" + ((("a" + "") + "\n\t") + "a")) + "hello") + "foo bar")}
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                    _ = v35
                }
                v4 -= (^(v4))
            }
        }
        _ = fi11
    }
    v4 = 42
    v9 -= (v9 & v9)
    v44 := GenS7{F0: -128, F1: v9, F2: -1}
    if false {
        for fi45 := 0; fi45 < 6; fi45++ {
            var v46 string = ("foo bar" + "")
            v47 := GenS3{F0: 128}
            _ = fi45
            _ = v46
            _ = v47
        }
        var v48 int32 = 227
        _ = len(v3)
        _ = v48
    } else {
        v4 = v4
    }
    var v49 uint16 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v49
    _ = v9
}
