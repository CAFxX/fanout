package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int; F2 bool }
type GenS4 struct { F0 bool }

func main() {
    var v1 uint = 456
    var v2 int8 = -84
    v3 := []int{}
    var v4 uint = 0
    v4 -= v1
    if true {
        _ = len(v3)
        if (2147483648 >= 42) {
            for fi5 := 0; fi5 < 12; fi5++ {
                for fi6 := 0; fi6 < 16; fi6++ {
                    v3 = append(v3, (fi5 << ((2) & 63)))
                    _ = fi6
                }
                _ = fi5
            }
        } else {
            for fi7 := 0; fi7 < 2; fi7++ {
                fmt.Println(v1)
                fmt.Println(v4, fi7, v4)
                v8 := GenS1{F0: 127}
                _ = fi7
                _ = v8
            }
        }
        fmt.Println(v2, v1)
    } else {
        var v9 float32 = -2.25
        switch 42 {
        case 128:
            for fi10 := 0; fi10 < 10; fi10++ {
                if (v4 < v4) {
                    if false {
                        if (true && false) {
                            v2 -= (v2 ^ v2)
                        } else {
                            if (2147483647 == 343) {
                                v9 += v9
                                for fi11 := 0; fi11 < 12; fi11++ {
                                    fmt.Println(v1, fi10, v4)
                                    _ = fi11
                                }
                            }
                        }
                        v1 -= v1
                        for fi12 := 0; fi12 < 13; fi12++ {
                            v2 = (v2 << ((726) & 63))
                            switch "foo bar" {
                            case "\n\t":
                                if false {
                                    switch "hello" {
                                    case "\n\t":
                                        var v13 uint16 = uint16(v4)
                                        _ = v13
                                    }
                                    var v14 int8 = (-(v2))
                                    _ = v14
                                } else {
                                    if false {
                                        var v15 float32 = float32(1e10)
                                        _ = v15
                                    }
                                }
                            case "":
                                v1 = 1
                            case "foo bar":
                                if ((false || (!(((fi12 | (fi12 / ((32768) | 1))) <= fi10) || (!false)))) || (0.0 > (-2.25 * 1e10))) {
                                    if (!(32767 < 1)) {
                                        if (204 > 2) {
                                            var v16 string = ""
                                            _ = v16
                                        } else {
                                            if (-42 > 42) {
                                                var v17 uint8 = 56
                                                v4 += (^(v1))
                                                var v18 int8 = v2
                                                var v19 uint64 = 1
                                                var v20 int32 = 1
                                                if true {
                                                    var v21 uint8 = (v17 >> (((v4 - (v1 ^ 4294967295))) & 63))
                                                    v18 = 1
                                                    _ = v21
                                                } else {
                                                    if ((v17 - (v17 + v17)) > v17) {
                                                        switch (fi12 / ((fi12) | 1)) {
                                                        case -484:
                                                            v9 -= 0.1
                                                        case 2147483647:
                                                            switch "foo bar" {
                                                            case "":
                                                                var v22 uint16 = 2
                                                                _ = v22
                                                            default:
                                                                v3 = append(v3, (-(fi12)))
                                                                if len(v3) > 0 {
                                                                    v3[0] = (fi10 & fi12)
                                                                }
                                                                var v23 uint = (v1 & v4)
                                                                _ = v23
                                                            }
                                                            v17 -= (v17 >> ((v1) & 63))
                                                        case 176:
                                                            v24 := GenS1{F0: (v17 - (^(v17)))}
                                                            _ = v24
                                                        }
                                                    } else {
                                                        v25 := GenS2{F0: 42}
                                                        if ((v17 ^ 2) < v17) {
                                                            v4 += v1
                                                        } else {
                                                            var v26 string = "a"
                                                            fmt.Println(v4, v26, v17)
                                                            _ = v26
                                                        }
                                                        var v27 uint64 = v19
                                                        _ = v25
                                                        _ = v27
                                                    }
                                                }
                                                v3 = append(v3, (^(fi12)))
                                                _ = v17
                                                _ = v18
                                                _ = v19
                                                _ = v20
                                            }
                                            v2 = -128
                                        }
                                        v4 = 42
                                        fmt.Println(v9, fi12)
                                    } else {
                                        v28 := GenS1{F0: 42}
                                        _ = v28
                                    }
                                    var v29 float32 = (v9 * -2.25)
                                    _ = v29
                                } else {
                                    v2 = -1
                                    var v30 int8 = (v2 & v2)
                                    _ = v30
                                }
                            default:
                                v31 := GenS3{F0: 1, F1: -300, F2: (!(7 < 255))}
                                _ = v31
                            }
                            _ = fi12
                        }
                    }
                    var v32 uint8 = 131
                    _ = v32
                }
                var v33 uint = v4
                _ = fi10
                _ = v33
            }
        }
        switch "" {
        case "":
            v34 := GenS2{F0: 0}
            _ = v34
        default:
            v1 -= (v1 * v4)
        }
        _ = v9
    }
    v35 := GenS4{F0: (("" + ("\n\t" + "\n\t")) < "\n\t")}
    if true {
        switch "hello" {
        case "foo bar":
            for fi36 := 0; fi36 < 16; fi36++ {
                for fi37 := 0; fi37 < 3; fi37++ {
                    v38 := GenS4{F0: false}
                    v3 = append(v3, 42)
                    v4 -= v1
                    _ = fi37
                    _ = v38
                }
                _ = len(v3)
                _ = fi36
            }
        case "\n\t":
            var v39 uint8 = 128
            _ = v39
        }
        for fi40 := 0; fi40 < 1; fi40++ {
            var v41 string = ("" + ("foo bar" + ("hello" + ("\n\t" + "foo bar"))))
            v42 := GenS3{F0: 2, F1: (fi40 * (fi40 >> ((v4) & 63))), F2: false}
            var v43 float32 = 3.14159
            _ = fi40
            _ = v41
            _ = v42
            _ = v43
        }
        var v44 uint64 = 237
        _ = v44
    }
    v45 := GenS2{F0: v2}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v46 := GenS3{F0: 19, F1: -74, F2: false}
    var v47 uint64 = 0
    v48 := GenS2{F0: v2}
    switch (((("foo bar" + "a") + "foo bar") + "a") + "hello") {
    case "\n\t":
        switch 156 {
        case 1:
            _ = len(v3)
        default:
            switch ("a" + "hello") {
            case "hello":
                v4 = (v1 + (^(v4)))
            }
        }
    }
    v4 -= 65535
    switch 42 {
    case 0:
        v1 -= 255
    case 675:
        var v49 uint16 = 255
        _ = v49
    }
    var v50 float64 = ((1.5 - -2.25) / 3.14159)
    var v51 string = "hello"
    var v52 float64 = v50
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v50
    _ = v51
    _ = v52
}
