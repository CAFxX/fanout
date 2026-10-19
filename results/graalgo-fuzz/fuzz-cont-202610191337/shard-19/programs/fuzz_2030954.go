package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint }
type GenS4 struct { F0 int32; F1 int8; F2 int32 }
type GenS5 struct { F0 uint32; F1 uint; F2 uint8 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 3.14159
    v3 := make(map[int]string)
    delete(v1, 32768)
    for fi4 := 0; fi4 < 17; fi4++ {
        switch "hello" {
        case "\n\t":
            v5 := []int{}
            _ = v5
        case "hello":
            for fi6 := 0; fi6 < 13; fi6++ {
                switch (^(fi6)) {
                case -9223372036854775808:
                    v2 = float32(-2.25)
                case 128:
                    v7 := GenS1{F0: 2}
                    _ = v7
                case 9223372036854775807:
                    switch (-(fi4)) {
                    case 9223372036854775807:
                        var v8 float64 = 1.5
                        _ = v8
                    case 1:
                        v2 += v2
                    }
                }
                v9 := GenS1{F0: 255}
                _ = fi6
                _ = v9
            }
        case "foo bar":
            if false {
                if (!(!(true && false))) {
                    v10 := GenS2{F0: 0}
                    _ = v10
                } else {
                    v2 = ((float32((float64(v2) + 1e10)) * (v2 * (float32(float64(-2.25)) - float32(((1.5 / (-2.25 * 1e10)) - 1e10))))) + 1e10)
                    _ = len(v3)
                    switch "hello" {
                    case "a":
                        var v11 uint16 = 32767
                        _ = v11
                    case "\n\t":
                        v2 = 1e10
                    case "":
                        _ = len(v1)
                    }
                }
            } else {
                delete(v1, (fi4 - (fi4 + fi4)))
                v12 := GenS3{F0: 2, F1: 288, F2: 65535}
                var v13 int8 = 117
                _ = v12
                _ = v13
            }
        }
        var v14 int8 = 81
        _ = fi4
        _ = v14
    }
    v15 := GenS2{F0: 16}
    v2 = ((1e10 / (v2 - ((v2 - ((v2 / (v2 / 0.1)) + v2)) / 1e10))) + 0.0)
    if false {
        switch ("\n\t" + "foo bar") {
        case "hello":
            v16 := v1[1]
            _ = v16
        case "\n\t":
            v2 += v2
        default:
            _ = len(v3)
        }
        var v17 float64 = (-2.25 - 0.0)
        _ = v17
    }
    if false {
        var v18 int32 = 1
        _ = v18
    }
    if false {
        var v19 float64 = 1e10
        _ = v19
    } else {
        var v20 int32 = 2147483647
        if false {
            v2 = 0.1
            if false {
                if (2 < 0) {
                    if (108 != 2) {
                        v20 = v20
                    }
                    if ((!(v2 < (3.14159 / v2))) && false) {
                        v20 += v20
                        for fi21 := 0; fi21 < 1; fi21++ {
                            v22 := GenS1{F0: 127}
                            switch "\n\t" {
                            case "a":
                                v2 = 0.1
                            }
                            _ = fi21
                            _ = v22
                        }
                        v20 += (v20 + (^(v20)))
                    }
                    v2 += ((1e10 - v2) - -2.25)
                }
                v20 = v20
            }
            v23 := GenS4{F0: (v20 & 637), F1: -128, F2: -2147483648}
            _ = v23
        } else {
            v2 = 1e10
            var v24 uint16 = 421
            v24 = (v24 << ((20) & 63))
            _ = v24
        }
        _ = v20
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        _ = len(v3)
        for fi26 := 0; fi26 < 17; fi26++ {
            var v27 uint32 = 197
            v28 := GenS1{F0: 0}
            switch fi25 {
            case -2147483648:
                fmt.Println(fi25, fi25, v27)
            default:
                v2 = v2
            }
            _ = fi26
            _ = v27
            _ = v28
        }
        v29 := []int{fi25, (fi25 - (^(fi25)))}
        _ = fi25
        _ = v29
    }
    var v30 uint = 0
    fmt.Println(v2, v30)
    v2 = float32(3.14159)
    var v31 int8 = 1
    v2 = v2
    for fi32 := 0; fi32 < 11; fi32++ {
        for fi33 := 0; fi33 < 10; fi33++ {
            v34 := []int{fi32, fi32, 480, -360}
            var v35 uint64 = 1
            if false {
                fmt.Println(v2)
            } else {
                v36 := GenS5{F0: 42, F1: (v30 + (v30 << ((v30) & 63))), F2: 36}
                var v37 uint = v30
                var v38 int = fi33
                _ = v36
                _ = v37
                _ = v38
            }
            _ = fi33
            _ = v34
            _ = v35
        }
        if true {
            v39 := GenS3{F0: 1, F1: 0, F2: 2}
            _ = v39
        } else {
            for fi40 := 0; fi40 < 11; fi40++ {
                var v41 bool = false
                switch fi32 {
                case 0:
                    if (100 == (-(v31))) {
                        v2 -= (v2 / v2)
                        var v42 int = fi40
                        _ = v42
                    }
                default:
                    var v43 int = fi40
                    _ = v43
                }
                v30 = v30
                _ = fi40
                _ = v41
            }
            for fi44 := 0; fi44 < 4; fi44++ {
                v45 := []int{}
                var v46 uint = (v30 & v30)
                _ = fi44
                _ = v45
                _ = v46
            }
        }
        fmt.Println(v31)
        _ = fi32
    }
    switch ("foo bar" + ((("hello" + ("a" + "")) + "a") + "foo bar")) {
    case "a":
        for fi47 := 0; fi47 < 9; fi47++ {
            if false {
                if true {
                    _ = len(v1)
                } else {
                    for fi48 := 0; fi48 < 12; fi48++ {
                        fmt.Println(v2)
                        if (!false) {
                            v49 := GenS5{F0: 948, F1: v30, F2: 255}
                            _ = v49
                        }
                        var v50 float32 = (((v2 * (1e10 * v2)) - ((v2 * ((v2 * v2) / (((v2 - 1e10) * v2) + (v2 - v2)))) - (1e10 / float32(3.14159)))) * v2)
                        _ = fi48
                        _ = v50
                    }
                    if false {
                        v30 += (^(v30))
                    } else {
                        fmt.Println(v30, v30, v2)
                        var v51 int16 = -366
                        v2 += v2
                        _ = v51
                    }
                }
            }
            v2 = v2
            var v52 float32 = (1.5 / v2)
            _ = fi47
            _ = v52
        }
    case "foo bar":
        if true {
            v2 = (v2 / (v2 - v2))
            v31 = (v31 + v31)
        } else {
            switch ("hello" + ("a" + ((("a" + ("\n\t" + "hello")) + ("\n\t" + "hello")) + "foo bar"))) {
            case "foo bar":
                for fi53 := 0; fi53 < 15; fi53++ {
                    var v54 string = ""
                    _ = fi53
                    _ = v54
                }
            default:
                v55 := GenS2{F0: 39}
                _ = v55
            }
        }
    case "\n\t":
        v2 += v2
    }
    for fi56 := 0; fi56 < 7; fi56++ {
        var v57 uint16 = 65535
        _ = fi56
        _ = v57
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
}
