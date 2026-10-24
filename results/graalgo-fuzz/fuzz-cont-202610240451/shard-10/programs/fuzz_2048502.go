package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 int; F2 uint8 }
type GenS6 struct { F0 int8; F1 uint8; F2 int }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int64 }
func genFn0() uint8 {
    if true {
        for fi6 := 0; fi6 < 13; fi6++ {
            var v7 int16 = 1
            v7 += 32767
            var v8 int64 = -1
            _ = fi6
            _ = v7
            _ = v8
        }
        if (true || (-1 <= -503)) {
            var v9 string = ""
            v9 = v9
            switch 42 {
            case -418:
                switch -2147483648 {
                case -2147483648:
                    v9 = (((v9 + "a") + (v9 + v9)) + "foo bar")
                case -1:
                    v10 := []int{1, 32768, 42, -177}
                    _ = v10
                default:
                    v11 := []int{}
                    _ = v11
                }
            case 260:
                v9 = v9
            case -583:
                v9 = (v9 + v9)
            }
            _ = v9
        } else {
            v12 := GenS1{F0: 1, F1: 419, F2: (737 < 2)}
            var v13 int8 = -1
            v13 = (v13 % ((42) | 1))
            _ = v12
            _ = v13
        }
    }
    if (!((!true) || false)) {
        var v14 uint16 = 809
        switch ("a" + "\n\t") {
        case "":
            v15 := GenS2{F0: 32767}
            _ = v15
        }
        for fi16 := 0; fi16 < 12; fi16++ {
            v14 += v14
            switch "a" {
            case "\n\t":
                v17 := GenS3{F0: "a", F1: 897}
                _ = v17
            case "":
                v14 = 32767
            case "a":
                var v18 uint64 = 2
                _ = v18
            default:
                v14 = (v14 / (((v14 * v14)) | 1))
            }
            v14 -= (v14 | 42)
            _ = fi16
        }
        _ = v14
    }
    var v19 int = 0
    for fi20 := 0; fi20 < 18; fi20++ {
        v21 := []int{}
        _ = fi20
        _ = v21
    }
    if true {
        switch ("hello" + (("a" + ((("a" + ((((("" + "foo bar") + "") + "hello") + "\n\t") + "\n\t")) + "\n\t") + "foo bar")) + "hello")) {
        case "a":
            var v22 uint = 65535
            _ = v22
        case "":
            var v23 float32 = (-2.25 / 1.5)
            _ = v23
        case "\n\t":
            if false {
                fmt.Println(v19)
                fmt.Println(v19)
            }
        }
        if (v19 <= -977) {
            if false {
                for fi24 := 0; fi24 < 17; fi24++ {
                    switch ("" + "\n\t") {
                    case "hello":
                        for fi25 := 0; fi25 < 16; fi25++ {
                            v26 := GenS4{F0: 0}
                            _ = fi25
                            _ = v26
                        }
                    case "foo bar":
                        for fi27 := 0; fi27 < 12; fi27++ {
                            v19 = (fi24 << ((0) & 63))
                            var v28 int8 = -21
                            _ = fi27
                            _ = v28
                        }
                    default:
                        var v29 int64 = 42
                        _ = v29
                    }
                    var v30 int8 = 127
                    _ = fi24
                    _ = v30
                }
                if (("hello" < "\n\t") && ("" >= "\n\t")) {
                    for fi31 := 0; fi31 < 13; fi31++ {
                        switch 340 {
                        case 128:
                            v19 += v19
                        default:
                            var v32 int8 = -1
                            _ = v32
                        }
                        _ = fi31
                    }
                }
                var v33 int8 = 32
                _ = v33
            }
            if (false || false) {
                v34 := []int{}
                if false {
                    var v35 float64 = 0.0
                    var v36 int64 = -9223372036854775808
                    _ = v35
                    _ = v36
                } else {
                    for fi37 := 0; fi37 < 3; fi37++ {
                        var v38 string = "hello"
                        _ = fi37
                        _ = v38
                    }
                    var v39 uint8 = 32
                    v39 -= v39
                    _ = v39
                }
                _ = v34
            }
        }
        var v40 string = ((("" + ("a" + "")) + "a") + ("\n\t" + ("" + ("a" + "hello"))))
        _ = v40
    } else {
        var v41 uint64 = 4294967295
        _ = v41
    }
    _ = v19
    return 1
}

func main() {
    var v1 uint16 = 42
    var v2 int16 = -105
    v3 := []string{"foo bar", "a", ""}
    var v4 uint64 = 2
    v5 := []uint32{851}
    switch 809 {
    case -1:
        switch -2147483648 {
        case -205:
            v42 := GenS5{F0: -1, F1: -174, F2: 128}
            _ = v42
        default:
            fmt.Println(v4, v4)
        }
    case 4294967295:
        v4 = (v4 + (v4 >> ((42) & 63)))
    case 2147483647:
        v43 := GenS6{F0: -128, F1: 0, F2: 0}
        _ = v43
    default:
        for fi44 := 0; fi44 < 8; fi44++ {
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
            v4 -= (v4 * v4)
            if (!((0.1 * 1.5) == 0.0)) {
                _ = len(v5)
                v45 := GenS1{F0: 42, F1: 2147483647, F2: true}
                _ = v45
            }
            _ = fi44
        }
    }
    var v46 int32 = 42
    fmt.Println(v46, v1)
    v3 = append(v3, "")
    v4 -= 0
    _ = len(v3)
    v2 = v2
    for fi47 := 0; fi47 < 19; fi47++ {
        v48 := GenS7{F0: v2}
        v49 := GenS8{F0: 9223372036854775807}
        _ = fi47
        _ = v48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
}
