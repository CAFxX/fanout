package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 bool; F2 uint }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 bool }

func main() {
    var v1 uint8 = 229
    v2 := make(map[string]uint32)
    var v3 uint64 = 1
    v4 := make(map[int]string)
    v5 := make(map[uint32]uint32)
    var v6 int64 = -717
    v1 = v1
    switch -9223372036854775808 {
    case -2147483648:
        v3 = v3
    case 42:
        if true {
            fmt.Println(v1, v6)
            v6 += (v6 - v6)
        }
    case -389:
        var v7 uint = 0
        _ = v7
    }
    var v8 float32 = (0.1 / 1.0)
    switch 128 {
    case 1:
        v9 := []int{}
        _ = v9
    case 0:
        if true {
            var v10 int32 = 1
            v6 = v6
            v11 := GenS1{F0: 1, F1: (v1 + (v1 >> ((62) & 63))), F2: 42}
            _ = v10
            _ = v11
        }
    case -646:
        for fi12 := 0; fi12 < 13; fi12++ {
            if false {
                switch "" {
                case "foo bar":
                    var v13 int = 441
                    _ = v13
                case "hello":
                    v1 = v1
                case "\n\t":
                    var v14 int32 = -792
                    _ = v14
                }
            } else {
                if false {
                    switch fi12 {
                    case -407:
                        var v15 uint64 = v3
                        _ = v15
                    case 2147483647:
                        var v16 int8 = -63
                        _ = v16
                    default:
                        v17 := v5[42]
                        _ = v17
                    }
                    var v18 uint = 4294967295
                    fmt.Println(v3, v8, v3)
                    _ = v18
                } else {
                    if (((65535 < 65535) && false) && (-81 != 305)) {
                        v2["\n\t"] = 2147483648
                        v19 := GenS2{F0: false}
                        var v20 uint8 = v1
                        _ = v19
                        _ = v20
                    }
                    v6 = v6
                    fmt.Println(v8, v3)
                }
                for fi21 := 0; fi21 < 13; fi21++ {
                    v22 := GenS1{F0: 73, F1: 1, F2: -2147483648}
                    _ = fi21
                    _ = v22
                }
                if ((v1 >= 127) && (!false)) {
                    var v23 string = (("foo bar" + "foo bar") + ("hello" + "a"))
                    var v24 float64 = (1.5 * 0.0)
                    v25 := GenS2{F0: (false || false)}
                    _ = v23
                    _ = v24
                    _ = v25
                }
            }
            v26 := GenS3{F0: true, F1: true, F2: 42}
            _ = fi12
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        if false {
            v28 := GenS3{F0: false, F1: false, F2: 4294967295}
            v29 := []int{(fi27 << ((65535) & 63)), fi27}
            delete(v4, 679)
            _ = v28
            _ = v29
        } else {
            switch ("" + "") {
            case "hello":
                if false {
                    fmt.Println(v8)
                } else {
                    var v30 int8 = 0
                    _ = v30
                }
            case "foo bar":
                if (42 == 225) {
                    var v31 string = "\n\t"
                    switch (fi27 - (fi27 ^ (fi27 / ((0) | 1)))) {
                    case 4294967295:
                        switch "hello" {
                        case "a":
                            v8 = 0.1
                        case "\n\t":
                            v32 := GenS4{F0: 790, F1: -527}
                            _ = v32
                        }
                    case 42:
                        v33 := []int{}
                        _ = v33
                    case -774:
                        for fi34 := 0; fi34 < 17; fi34++ {
                            var v35 uint16 = 470
                            if false {
                                var v36 float32 = 3.14159
                                fmt.Println(v3)
                                fmt.Println(fi34)
                                _ = v36
                            }
                            switch -1 {
                            case -1:
                                var v37 float32 = (((3.14159 / v8) + 1e10) - v8)
                                _ = v37
                            default:
                                var v38 int = fi27
                                _ = v38
                            }
                            _ = fi34
                            _ = v35
                        }
                    default:
                        fmt.Println(v6, v8)
                    }
                    var v39 int8 = 127
                    _ = v31
                    _ = v39
                } else {
                    v40 := GenS5{F0: (float32(float64((v8 + 1.5))) == v8)}
                    if (792 > 872) {
                        v41 := GenS3{F0: true, F1: false, F2: 255}
                        _ = v41
                    }
                    _ = v40
                }
            case "":
                if false {
                    _ = len(v5)
                    var v42 uint = 255
                    switch "" {
                    case "foo bar":
                        _ = len(v5)
                    }
                    _ = v42
                }
            default:
                v43 := GenS5{F0: true}
                _ = v43
            }
            for fi44 := 0; fi44 < 10; fi44++ {
                fmt.Println(v1, fi27)
                v1 += 127
                _ = fi44
            }
            fmt.Println(v3, fi27)
        }
        _ = fi27
    }
    fmt.Println(v1, v1)
    v45 := v4[0]
    v45 = v45
    var v46 uint64 = v3
    v47 := v4[0]
    v45 = ("a" + v47)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v8
}
