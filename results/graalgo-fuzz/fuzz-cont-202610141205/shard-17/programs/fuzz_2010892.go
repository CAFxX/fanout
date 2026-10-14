package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint8; F2 bool }
type GenS4 struct { F0 uint16; F1 int32; F2 int }
type GenS5 struct { F0 int }

func main() {
    var v1 uint8 = 255
    var v2 bool = false
    v3 := make(map[uint32]int)
    switch 4294967295 {
    case 9223372036854775807:
        for fi4 := 0; fi4 < 1; fi4++ {
            switch fi4 {
            case 127:
                var v5 bool = true
                _ = v5
            case 32767:
                v6 := GenS1{F0: 1000}
                _ = v6
            default:
                if v2 {
                    switch "a" {
                    case "hello":
                        v7 := GenS1{F0: 377}
                        _ = v7
                    case "a":
                        v8 := GenS1{F0: 42}
                        _ = v8
                    case "":
                        if v2 {
                            for fi9 := 0; fi9 < 2; fi9++ {
                                var v10 uint16 = 2
                                v11 := GenS2{F0: 2}
                                _ = fi9
                                _ = v10
                                _ = v11
                            }
                            v2 = v2
                            var v12 uint16 = 1
                            _ = v12
                        }
                    }
                    v13 := GenS1{F0: 32767}
                    _ = v13
                }
            }
            _ = fi4
        }
    case 187:
        v14 := GenS1{F0: -32768}
        _ = v14
    case 2147483647:
        switch "hello" {
        case "\n\t":
            var v15 int8 = 5
            _ = v15
        case "":
            v16 := GenS2{F0: 2}
            _ = v16
        }
    }
    switch ("" + (("" + (("a" + ("\n\t" + (("a" + "a") + ("\n\t" + "foo bar")))) + ("foo bar" + (("" + "foo bar") + ("\n\t" + "foo bar"))))) + "\n\t")) {
    case "hello":
        if (1 < 0) {
            for fi17 := 0; fi17 < 11; fi17++ {
                v18 := GenS3{F0: 42, F1: (^(v1)), F2: v2}
                if false {
                    var v19 uint64 = 2
                    var v20 uint = 65535
                    _ = v19
                    _ = v20
                }
                if (259 > 18446744073709551615) {
                    v1 += (^(v1))
                    var v21 uint16 = 255
                    for fi22 := 0; fi22 < 11; fi22++ {
                        switch 0 {
                        case 128:
                            var v23 float64 = 3.14159
                            _ = v23
                        case -888:
                            var v24 int16 = 742
                            _ = v24
                        }
                        _ = fi22
                    }
                    _ = v21
                }
                _ = fi17
                _ = v18
            }
            if false {
                switch 42 {
                case 2147483647:
                    v25 := []int{109, 128, 0}
                    _ = v25
                }
                v1 = v1
                v26 := []int{}
                _ = v26
            } else {
                v27 := []int{2147483647, -812}
                v2 = v2
                var v28 int = 4294967295
                _ = v27
                _ = v28
            }
        } else {
            v2 = false
            var v29 uint = 814
            var v30 uint8 = v1
            _ = v29
            _ = v30
        }
    case "\n\t":
        var v31 uint64 = 989
        _ = v31
    case "":
        v32 := GenS4{F0: 334, F1: -2147483648, F2: 0}
        _ = v32
    default:
        fmt.Println(v2)
    }
    switch 32767 {
    case 42:
        v1 = 128
    case 2147483647:
        var v33 uint = 121
        _ = v33
    case 0:
        fmt.Println(v2, v2)
    }
    var v34 uint64 = 654
    switch -972 {
    case 32768:
        var v35 int16 = 1
        _ = v35
    case 0:
        var v36 uint16 = 2
        _ = v36
    }
    delete(v3, 2147483647)
    if (!true) {
        if (!("\n\t" != "foo bar")) {
            if v2 {
                v37 := GenS5{F0: -1}
                _ = v37
            }
            v38 := []int{109, -2147483648, 32767, 1}
            _ = v38
        } else {
            for fi39 := 0; fi39 < 11; fi39++ {
                var v40 uint8 = 127
                var v41 int32 = 2147483647
                var v42 bool = ("hello" == "foo bar")
                _ = fi39
                _ = v40
                _ = v41
                _ = v42
            }
        }
        for fi43 := 0; fi43 < 9; fi43++ {
            v34 = (^(v34))
            var v44 float64 = (-2.25 - 1.5)
            if (((130 < 0) && v2) || v2) {
                var v45 float64 = ((1.5 - 1e10) + v44)
                _ = v45
            }
            _ = fi43
            _ = v44
        }
    }
    v2 = v2
    for fi46 := 0; fi46 < 12; fi46++ {
        var v47 bool = v2
        switch fi46 {
        case 1:
            var v48 uint = 596
            _ = v48
        case 32768:
            var v49 int8 = 127
            _ = v49
        case -2147483648:
            var v50 uint64 = (v34 >> ((42) & 63))
            _ = v50
        default:
            fmt.Println(fi46, fi46)
        }
        _ = fi46
        _ = v47
    }
    v1 = v1
    _ = len(v3)
    var v51 float32 = 1.5
    switch "" {
    case "foo bar":
        var v52 uint = 158
        _ = v52
    case "\n\t":
        var v53 int32 = 42
        _ = v53
    }
    v2 = (383 > 2)
    for fi54 := 0; fi54 < 3; fi54++ {
        for fi55 := 0; fi55 < 9; fi55++ {
            v56 := []int{(fi55 | (fi54 & (fi55 & 602)))}
            _ = fi55
            _ = v56
        }
        v57 := []int{}
        switch "\n\t" {
        case "a":
            var v58 uint8 = v1
            _ = v58
        default:
            _ = len(v3)
        }
        _ = fi54
        _ = v57
    }
    v2 = (255 < 65535)
    for fi59 := 0; fi59 < 3; fi59++ {
        for fi60 := 0; fi60 < 12; fi60++ {
            fmt.Println(fi60, v34)
            var v61 int32 = -2147483648
            var v62 int16 = 0
            _ = fi60
            _ = v61
            _ = v62
        }
        _ = fi59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v51
}
