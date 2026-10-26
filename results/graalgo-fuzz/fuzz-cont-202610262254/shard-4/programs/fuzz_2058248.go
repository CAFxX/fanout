package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint }
func genFn0() bool {
    v6 := GenS1{F0: 65535, F1: "a", F2: -115}
    var v7 int8 = -128
    _ = v6
    _ = v7
    return false
}
func genFn1() uint32 {
    if false {
        var v8 float64 = 3.14159
        _ = v8
    } else {
        v9 := []int{-434, -338, 4294967295}
        _ = v9
    }
    switch 32768 {
    case 4294967295:
        var v10 uint32 = 4294967295
        _ = v10
    case 32767:
        v11 := []int{1}
        _ = v11
    case 127:
        v12 := []int{-951, -806, -9223372036854775808}
        _ = v12
    }
    v13 := []int{128, 55}
    var v14 uint = 332
    _ = v13
    _ = v14
    return 1
}

func main() {
    var v1 uint8 = 0
    v2 := []int{-288, 2147483647}
    var v3 uint = 65535
    var v4 float64 = 1e10
    var v5 int8 = -128
    var v15 uint8 = v1
    switch (("\n\t" + "a") + "a") {
    case "foo bar":
        if true {
            for fi16 := 0; fi16 < 17; fi16++ {
                for fi17 := 0; fi17 < 14; fi17++ {
                    _ = len(v2)
                    _ = fi17
                }
                _ = fi16
            }
            switch -2147483648 {
            case 128:
                v5 -= v5
            case 32768:
                if ((!(false || (true || false))) && (937 > -1)) {
                    for fi18 := 0; fi18 < 13; fi18++ {
                        var v19 uint16 = 2
                        _ = fi18
                        _ = v19
                    }
                    var v20 uint32 = 42
                    _ = v20
                }
            default:
                if false {
                    for fi21 := 0; fi21 < 17; fi21++ {
                        if false {
                            fmt.Println(v5, fi21, v5)
                        } else {
                            var v22 string = "foo bar"
                            var v23 int32 = 2147483647
                            v24 := GenS1{F0: 2, F1: "foo bar", F2: int8(fi21)}
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        if (true && (3.14159 != 3.14159)) {
                            var v25 uint32 = 2147483648
                            if false {
                                var v26 uint8 = (v15 ^ 0)
                                _ = v26
                            }
                            switch "a" {
                            case "foo bar":
                                if (42 >= 0) {
                                    v4 += v4
                                    var v27 bool = false
                                    var v28 int64 = -1
                                    _ = v27
                                    _ = v28
                                }
                            case "":
                                fmt.Println(v3, fi21, fi21)
                            }
                            _ = v25
                        } else {
                            var v29 int = -617
                            var v30 uint8 = v1
                            for fi31 := 0; fi31 < 1; fi31++ {
                                v32 := GenS1{F0: 0, F1: "foo bar", F2: -128}
                                v29 = fi21
                                _ = fi31
                                _ = v32
                            }
                            _ = v29
                            _ = v30
                        }
                        _ = fi21
                    }
                    switch "foo bar" {
                    case "":
                        var v33 uint8 = 244
                        _ = v33
                    case "hello":
                        var v34 int16 = 607
                        _ = v34
                    case "foo bar":
                        v35 := GenS1{F0: 32767, F1: ("a" + ""), F2: (v5 - -1)}
                        _ = v35
                    }
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
        }
    default:
        switch "\n\t" {
        case "a":
            var v36 uint = (^(v3))
            _ = v36
        case "foo bar":
            switch 128 {
            case 1:
                v4 = v4
            case 128:
                v1 = v15
            case -1:
                v37 := GenS1{F0: 456, F1: (("\n\t" + "a") + "foo bar"), F2: (v5 / (((v5 % ((v5) | 1))) | 1))}
                _ = v37
            }
        case "hello":
            v38 := GenS1{F0: 2, F1: "", F2: v5}
            _ = v38
        default:
            var v39 uint16 = 41
            _ = v39
        }
    }
    if (!true) {
        v2 = append(v2, -998)
        var v40 uint16 = uint16(v15)
        v1 = (v1 >> ((v3) & 63))
        _ = v40
    } else {
        var v41 int64 = 0
        _ = v41
    }
    for fi42 := 0; fi42 < 10; fi42++ {
        switch "foo bar" {
        case "a":
            fmt.Println(v5)
        case "":
            var v43 int32 = 2147483647
            _ = v43
        }
        for fi44 := 0; fi44 < 1; fi44++ {
            switch (fi44 << (((^(v3))) & 63)) {
            case 127:
                v45 := GenS2{F0: 65535, F1: 42, F2: 42}
                _ = v45
            }
            v2 = append(v2, (fi44 - (fi44 << ((v3) & 63))))
            _ = fi44
        }
        for fi46 := 0; fi46 < 2; fi46++ {
            fmt.Println(fi46, v5, fi42)
            fmt.Println(fi42)
            fmt.Println(v1, v15)
            _ = fi46
        }
        _ = fi42
    }
    if len(v2) > 0 {
        v2[0] = -601
    }
    _ = len(v2)
    fmt.Println(v4, v4)
    v47 := GenS1{F0: 255, F1: "foo bar", F2: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
