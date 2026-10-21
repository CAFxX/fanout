package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int32; F1 uint; F2 int8 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    v3 := make(map[string]bool)
    v4 := []string{}
    v5 := GenS1{F0: 42, F1: 543}
    if (508 >= 4294967295) {
        var v6 int64 = 9223372036854775807
        var v7 float64 = 3.14159
        if true {
            fmt.Println(v6, v7)
            v2[0] = 823
        }
        _ = v6
        _ = v7
    } else {
        var v8 string = (("a" + "") + "")
        v9 := GenS1{F0: 2147483647, F1: 1}
        for fi10 := 0; fi10 < 17; fi10++ {
            if (true && false) {
                var v11 int8 = 100
                v4 = append(v4, v8)
                _ = v11
            }
            _ = fi10
        }
        _ = v8
        _ = v9
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v12 := GenS2{F0: 42, F1: 550, F2: 1}
    v13 := GenS2{F0: 2147483647, F1: 197, F2: -1}
    v14 := GenS3{F0: true, F1: 2147483648}
    if (false && (0.0 < -2.25)) {
        var v15 int16 = 1000
        var v16 int64 = -9223372036854775808
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 17; fi17++ {
            switch "foo bar" {
            case "\n\t":
                v4 = append(v4, "")
            }
            var v18 uint8 = 198
            v19 := GenS4{F0: 919}
            _ = fi17
            _ = v18
            _ = v19
        }
        switch 128 {
        case -1:
            _ = len(v1)
        default:
            if false {
                for fi20 := 0; fi20 < 18; fi20++ {
                    var v21 uint = 179
                    v21 = (v21 ^ (^(v21)))
                    switch ((("" + "a") + (((("" + ("\n\t" + "\n\t")) + "hello") + "hello") + ("\n\t" + "a"))) + ("\n\t" + "")) {
                    case "foo bar":
                        v21 = 0
                    default:
                        v4 = append(v4, ("\n\t" + "\n\t"))
                    }
                    _ = fi20
                    _ = v21
                }
            }
        }
        switch 46 {
        case -1:
            var v22 float64 = 0.1
            _ = v22
        }
    }
    if (100 < 42) {
        var v23 bool = (42 == 317)
        if true {
            for fi24 := 0; fi24 < 19; fi24++ {
                for fi25 := 0; fi25 < 17; fi25++ {
                    fmt.Println(fi25, v23)
                    var v26 string = "hello"
                    _ = fi25
                    _ = v26
                }
                for fi27 := 0; fi27 < 2; fi27++ {
                    switch (fi24 << ((1) & 63)) {
                    case 9223372036854775807:
                        switch fi27 {
                        case 302:
                            if (4294967295 >= 140) {
                                v28 := v2[2]
                                _ = v28
                            }
                        }
                    case 128:
                        var v29 uint8 = 47
                        _ = v29
                    }
                    v23 = (v23 || (201 != 240))
                    _ = fi27
                }
                _ = fi24
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        v23 = (255 != 127)
        _ = v23
    } else {
        var v30 uint16 = 122
        var v31 int = 9223372036854775807
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        var v33 uint = 917
        var v34 uint8 = 1
        fmt.Println(v33, v33)
        _ = fi32
        _ = v33
        _ = v34
    }
    v35 := GenS5{F0: 25}
    var v36 int16 = -1
    for fi37 := 0; fi37 < 1; fi37++ {
        var v38 uint16 = 0
        for fi39 := 0; fi39 < 12; fi39++ {
            v36 = -986
            v40 := v1[(fi39 << ((255) & 63))]
            v40 = ((fi37 - fi37) == (fi37 & (fi37 << ((4294967295) & 63))))
            _ = fi39
            _ = v40
        }
        _ = fi37
        _ = v38
    }
    var v41 uint8 = 1
    v36 = (-(v36))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v5
}
