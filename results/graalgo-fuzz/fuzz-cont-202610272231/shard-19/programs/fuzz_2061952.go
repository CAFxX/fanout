package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 bool; F1 int32; F2 int16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = false
    var v3 uint32 = 2
    var v4 uint64 = 2
    fmt.Println(v2, v3)
    var v5 string = "hello"
    v6 := GenS1{F0: -32768, F1: v2}
    if v2 {
        var v7 int = -1
        _ = v7
    }
    var v8 uint = 476
    for fi9 := 0; fi9 < 9; fi9++ {
        fmt.Println(v5, v8, fi9)
        var v10 int32 = -1
        _ = fi9
        _ = v10
    }
    if true {
        v11 := []int{127}
        _ = v11
    } else {
        if v2 {
            v12 := []int{-66, 127}
            _ = v12
        }
        var v13 bool = (2 > 0)
        _ = v13
    }
    if ((("hello" + "hello") + v5) == (v5 + (((v5 + (v5 + (v5 + "foo bar"))) + v5) + "a"))) {
        switch 127 {
        case 0:
            var v14 uint64 = 174
            _ = v14
        case -818:
            fmt.Println(v3, v4, v4)
        case 2147483647:
            fmt.Println(v4, v5)
        }
        fmt.Println(v5)
    } else {
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 int16 = -32768
            switch (^(fi15)) {
            case 128:
                for fi17 := 0; fi17 < 15; fi17++ {
                    v3 -= (v3 & v3)
                    v18 := []int{int(v3)}
                    fmt.Println(v16, v8, v4)
                    _ = fi17
                    _ = v18
                }
            default:
                v19 := GenS2{F0: v8, F1: (v4 - 4294967295), F2: (v4 | uint64(v8))}
                _ = v19
            }
            _ = fi15
            _ = v16
        }
        if false {
            fmt.Println(v3)
            v20 := GenS2{F0: v8, F1: v4, F2: v4}
            _ = v20
        }
        var v21 uint16 = 629
        _ = v21
    }
    var v22 int32 = 2147483647
    v4 = (v4 >> (((^(v8))) & 63))
    switch "foo bar" {
    case "foo bar":
        v23 := GenS3{F0: 190}
        _ = v23
    default:
        var v24 uint16 = 65535
        _ = v24
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        var v26 float64 = 0.0
        v5 = v5
        _ = fi25
        _ = v26
    }
    switch 423 {
    case -2147483648:
        for fi27 := 0; fi27 < 9; fi27++ {
            v22 = int32(fi27)
            for fi28 := 0; fi28 < 14; fi28++ {
                var v29 string = ("\n\t" + v5)
                switch 552 {
                case 32768:
                    var v30 uint32 = 940
                    _ = v30
                default:
                    if (int16(v3) >= -32768) {
                        v22 = int32(v4)
                        fmt.Println(fi28)
                        var v31 uint = v8
                        _ = v31
                    } else {
                        if (1 == (v4 << ((v8) & 63))) {
                            switch (v5 + "a") {
                            case "":
                                v4 = 42
                            case "a":
                                switch (fi27 >> (((^(v8))) & 63)) {
                                case 9223372036854775807:
                                    var v32 int32 = -897
                                    var v33 int64 = 0
                                    v29 = v29
                                    _ = v32
                                    _ = v33
                                case 127:
                                    var v34 uint = (v8 ^ 2)
                                    var v35 float64 = 0.1
                                    _ = v34
                                    _ = v35
                                }
                            case "hello":
                                fmt.Println(v4)
                            }
                        } else {
                            v36 := GenS2{F0: v8, F1: v4, F2: (v4 & (v4 << ((v8) & 63)))}
                            _ = v36
                        }
                        v37 := GenS3{F0: 128}
                        _ = v37
                    }
                }
                _ = fi28
                _ = v29
            }
            v38 := GenS4{F0: (v22 + (v22 << ((v8) & 63))), F1: -128}
            _ = fi27
            _ = v38
        }
    case -9223372036854775808:
        v4 += (v4 + v4)
    case -917:
        v39 := GenS2{F0: 4294967295, F1: v4, F2: (v4 / ((v4) | 1))}
        _ = v39
    }
    v40 := GenS2{F0: v8, F1: 2, F2: (^(v4))}
    var v41 uint16 = 771
    for fi42 := 0; fi42 < 2; fi42++ {
        v5 = ("\n\t" + ((v5 + v5) + "\n\t"))
        if (v2 && v2) {
            v43 := GenS3{F0: 2}
            var v44 int64 = 761
            delete(v1, v3)
            _ = v43
            _ = v44
        } else {
            var v45 string = v5
            var v46 uint16 = v41
            _ = v45
            _ = v46
        }
        delete(v1, (v3 / ((v3) | 1)))
        _ = fi42
    }
    switch ("\n\t" + "hello") {
    case "hello":
        var v47 uint = v8
        _ = v47
    case "a":
        v48 := []int{}
        _ = v48
    case "":
        fmt.Println(v41, v2)
    default:
        for fi49 := 0; fi49 < 9; fi49++ {
            v22 -= v22
            for fi50 := 0; fi50 < 1; fi50++ {
                v51 := []int{1, int(v3), int(v41), 9223372036854775807}
                switch "hello" {
                case "\n\t":
                    if len(v51) >= 2 {
                        v51 = v51[1:len(v51)]
                    }
                default:
                    switch ((("\n\t" + v5) + ("hello" + v5)) + "") {
                    case "":
                        v41 += 0
                    case "hello":
                        var v52 int64 = -9223372036854775808
                        _ = v52
                    default:
                        v53 := GenS5{F0: false, F1: (-(v22)), F2: 1}
                        _ = v53
                    }
                }
                _ = fi50
                _ = v51
            }
            _ = fi49
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v8
}
