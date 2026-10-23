package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 string }
type GenS2 struct { F0 int16; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint; F1 int64 }

func main() {
    var v1 int16 = 32767
    var v2 uint16 = 408
    v3 := make(map[int]uint32)
    var v4 uint64 = 353
    v5 := make(map[string]string)
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(fi6)
        v4 = (v4 ^ (v4 | 4294967295))
        _ = fi6
    }
    if ((true && (!(((0.0 / 1e10) / 1.0) >= (1.5 * ((3.14159 - 0.1) + 3.14159))))) || false) {
        if (2 <= 91) {
            _ = len(v3)
            var v7 int16 = v1
            _ = v7
        } else {
            var v8 int = 127
            _ = v8
        }
        var v9 int16 = 1
        _ = v9
    }
    v10 := GenS1{F0: (v4 + (v4 * (v4 - (v4 ^ v4)))), F1: (v1 << ((42) & 63)), F2: "a"}
    switch 42 {
    case 42:
        switch 42 {
        case -1:
            if (3.14159 != (0.0 - 3.14159)) {
                fmt.Println(v1, v4)
                switch (("foo bar" + "\n\t") + ("a" + "a")) {
                case "foo bar":
                    if (true && false) {
                        v11 := []int{}
                        _ = v11
                    }
                default:
                    var v12 int8 = 100
                    _ = v12
                }
                v5["\n\t"] = ""
            } else {
                for fi13 := 0; fi13 < 16; fi13++ {
                    fmt.Println(v1, fi13, v2)
                    _ = fi13
                }
                if true {
                    v14 := GenS1{F0: v4, F1: (v1 % (((-(v1))) | 1)), F2: ("foo bar" + ("hello" + "a"))}
                    for fi15 := 0; fi15 < 1; fi15++ {
                        if false {
                            switch "\n\t" {
                            case "foo bar":
                                v16 := GenS1{F0: 803, F1: 1000, F2: "a"}
                                _ = v16
                            case "":
                                v17 := []int{fi15, -9223372036854775808}
                                _ = v17
                            case "hello":
                                v18 := []int{fi15, (fi15 - (fi15 + 4294967295)), (^(fi15)), 1}
                                _ = v18
                            }
                            switch fi15 {
                            case 541:
                                if false {
                                    var v19 int32 = -1
                                    v20 := []int{fi15, (fi15 + fi15)}
                                    _ = v19
                                    _ = v20
                                } else {
                                    if (v4 < 840) {
                                        v21 := GenS1{F0: (^(v4)), F1: (^(v1)), F2: "a"}
                                        var v22 bool = false
                                        if ((0.0 + 0.0) <= ((0.0 * 1.5) - float64(0.1))) {
                                            var v23 int8 = int8(fi15)
                                            _ = v23
                                        } else {
                                            var v24 uint16 = (v2 - v2)
                                            switch 2147483647 {
                                            case -9223372036854775808:
                                                v22 = v22
                                            default:
                                                switch fi15 {
                                                case 9223372036854775807:
                                                    switch fi15 {
                                                    case 32768:
                                                        switch (fi15 * (fi15 - (fi15 - fi15))) {
                                                        case 1:
                                                            switch (fi15 >> ((65535) & 63)) {
                                                            case 128:
                                                                fmt.Println(v4, v24)
                                                            }
                                                        case -387:
                                                            for fi25 := 0; fi25 < 11; fi25++ {
                                                                _ = len(v3)
                                                                v24 -= (v2 ^ (^(v2)))
                                                                _ = fi25
                                                            }
                                                        }
                                                    case -778:
                                                        fmt.Println(v22)
                                                    default:
                                                        v22 = false
                                                    }
                                                case 307:
                                                    v26 := []int{}
                                                    _ = v26
                                                }
                                            }
                                            _ = v24
                                        }
                                        _ = v21
                                        _ = v22
                                    }
                                    for fi27 := 0; fi27 < 11; fi27++ {
                                        fmt.Println(fi27)
                                        v28 := []int{(-(fi15)), (fi15 ^ fi27)}
                                        _ = fi27
                                        _ = v28
                                    }
                                }
                            }
                        } else {
                            v29 := v5["hello"]
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 14; fi30++ {
                            var v31 uint64 = (v4 % ((v4) | 1))
                            var v32 float64 = 1.5
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        _ = fi15
                    }
                    for fi33 := 0; fi33 < 7; fi33++ {
                        fmt.Println(fi33, v1, fi33)
                        _ = fi33
                    }
                    _ = v14
                }
                v4 = v4
            }
        }
    }
    switch "foo bar" {
    case "a":
        if true {
            var v34 string = (((("foo bar" + (("a" + ((("\n\t" + (("hello" + "foo bar") + "hello")) + "") + "")) + "hello")) + ((("hello" + "foo bar") + "") + "foo bar")) + "") + "\n\t")
            var v35 int8 = -128
            _ = v34
            _ = v35
        }
    case "hello":
        v36 := v5["\n\t"]
        _ = v36
    }
    v2 = v2
    v37 := []int{-2147483648, 2147483647, 127}
    for fi38 := 0; fi38 < 17; fi38++ {
        delete(v3, fi38)
        v4 = (^(v4))
        _ = fi38
    }
    fmt.Println(v1)
    switch "hello" {
    case "foo bar":
        if (true || true) {
            v39 := GenS1{F0: (v4 ^ (v4 * v4)), F1: v1, F2: "hello"}
            _ = v39
        } else {
            fmt.Println(v2, v2)
        }
    case "hello":
        v1 -= (v1 >> ((590) & 63))
    case "":
        switch -791 {
        case 509:
            if (true && ((2 != 9) && false)) {
                switch -944 {
                case 57:
                    var v40 int32 = 879
                    _ = v40
                }
            } else {
                v2 += v2
                v2 = 255
                v1 -= v1
            }
        case 42:
            for fi41 := 0; fi41 < 9; fi41++ {
                switch int(v2) {
                case 1:
                    v1 += v1
                case 32767:
                    v42 := GenS2{F0: (v1 - (v1 - (v1 + v1))), F1: (v4 + (v4 | v4)), F2: -9223372036854775808}
                    _ = v42
                }
                for fi43 := 0; fi43 < 9; fi43++ {
                    v1 += 32767
                    v44 := GenS2{F0: v1, F1: v4, F2: 42}
                    _ = fi43
                    _ = v44
                }
                if true {
                    if false {
                        v45 := v5[("a" + "")]
                        var v46 float32 = 0.0
                        _ = len(v37)
                        _ = v45
                        _ = v46
                    } else {
                        for fi47 := 0; fi47 < 3; fi47++ {
                            var v48 int16 = 0
                            _ = fi47
                            _ = v48
                        }
                        v37 = append(v37, (fi41 ^ int(v4)))
                    }
                    var v49 bool = ((v4 - uint64(v2)) > (v4 | 4294967295))
                    _ = v49
                } else {
                    v50 := GenS3{F0: 358, F1: -1}
                    var v51 float32 = -2.25
                    var v52 uint16 = v2
                    _ = v50
                    _ = v51
                    _ = v52
                }
                _ = fi41
            }
        case 2147483647:
            v4 = (v4 % ((v4) | 1))
        default:
            var v53 uint16 = (v2 & (v2 >> ((304) & 63)))
            _ = v53
        }
    default:
        switch "hello" {
        case "":
            if len(v37) >= 2 {
                v37 = v37[1:len(v37)]
            }
        case "hello":
            v54 := v5["hello"]
            _ = v54
        }
    }
    _ = len(v37)
    v2 -= 255
    fmt.Println(v4, v2)
    for fi55 := 0; fi55 < 1; fi55++ {
        _ = len(v3)
        switch (("a" + "\n\t") + "foo bar") {
        case "":
            var v56 uint64 = v4
            _ = v56
        }
        if (42 == 754) {
            var v57 float64 = (1e10 / 1.0)
            v57 += float64(1.5)
            if false {
                v2 = v2
                var v58 uint32 = 2147483648
                _ = v58
            }
            _ = v57
        } else {
            v59 := GenS1{F0: v4, F1: int16(fi55), F2: ("a" + "\n\t")}
            _ = v59
        }
        _ = fi55
    }
    var v60 string = ("a" + "a")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v60
}
