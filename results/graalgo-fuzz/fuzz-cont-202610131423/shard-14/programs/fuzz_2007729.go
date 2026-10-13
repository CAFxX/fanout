package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 string; F1 int8; F2 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string; F1 uint8; F2 uint }
func genFn0() int {
    v5 := GenS1{F0: 1}
    var v6 bool = false
    var v7 uint64 = 42
    if v6 {
        var v8 int16 = -1
        if false {
            var v9 float64 = 0.0
            var v10 uint = 0
            _ = v10
            _ = v9
        } else {
            v6 = v6
            var v11 bool = (0 == -565)
            v12 := []int{-424, 1, 2147483647}
            _ = v11
            _ = v12
        }
        _ = v8
    }
    switch 2147483647 {
    case 42:
        fmt.Println(v7)
    case 9223372036854775807:
        v13 := GenS2{F0: "foo bar", F1: true}
        _ = v13
    case 32768:
        if v6 {
            v6 = (42 <= 2)
            v6 = v6
            fmt.Println(v7, v7)
        }
    default:
        v6 = (-9223372036854775808 == 0)
    }
    _ = v5
    _ = v6
    _ = v7
    return 127
}

func main() {
    v1 := []uint32{}
    var v2 int64 = 42
    v3 := make(map[string]uint32)
    var v4 uint8 = 128
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v14 int16 = -32768
    if ("foo bar" == "") {
        var v15 int8 = 42
        switch -2147483648 {
        case 42:
            v16 := GenS3{F0: 255, F1: 2}
            _ = v16
        case -542:
            for fi17 := 0; fi17 < 2; fi17++ {
                for fi18 := 0; fi18 < 7; fi18++ {
                    v14 = 1000
                    var v19 uint = 674
                    var v20 uint32 = uint32(v4)
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                switch ("hello" + "foo bar") {
                case "foo bar":
                    var v21 int8 = 127
                    _ = v21
                default:
                    v3[""] = 4294967295
                }
                v2 -= v2
                _ = fi17
            }
        default:
            if false {
                for fi22 := 0; fi22 < 13; fi22++ {
                    _ = len(v1)
                    _ = fi22
                }
                var v23 uint64 = 0
                v24 := GenS4{F0: ((("" + "\n\t") + "") + "foo bar"), F1: (v15 + 0), F2: 255}
                _ = v23
                _ = v24
            }
        }
        _ = v15
    } else {
        fmt.Println(v4)
        for fi25 := 0; fi25 < 15; fi25++ {
            var v26 int = (fi25 ^ (fi25 + (fi25 / (((fi25 & (fi25 ^ (fi25 + (fi25 - fi25))))) | 1))))
            _ = fi25
            _ = v26
        }
        switch 0 {
        case 705:
            var v27 uint16 = 255
            _ = v27
        }
    }
    if false {
        for fi28 := 0; fi28 < 15; fi28++ {
            fmt.Println(fi28, v2, v14)
            _ = fi28
        }
        v14 = 243
        var v29 uint64 = 52
        _ = v29
    } else {
        v30 := GenS5{F0: -2147483648}
        _ = v30
    }
    v31 := GenS1{F0: 128}
    fmt.Println(v2, v4)
    if false {
        v32 := GenS2{F0: "a", F1: true}
        if (((v2 < (v2 + (v2 / ((-1) | 1)))) && (!false)) && true) {
            fmt.Println(v14, v4, v14)
            v14 = (v14 & (v14 ^ v14))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            if (true && true) {
                var v33 uint64 = 652
                _ = v33
            }
            if (234 < 1) {
                v4 = v4
                v34 := GenS6{F0: (!(-38 <= 248))}
                _ = v34
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = v32
    } else {
        var v35 int8 = 0
        _ = v35
    }
    var v36 string = ((("\n\t" + "a") + "") + (("\n\t" + ("hello" + ("\n\t" + "\n\t"))) + ((("" + "a") + "") + "hello")))
    var v37 uint = 663
    v38 := GenS7{F0: (v36 + v36), F1: (v4 << (((v37 & 333)) & 63)), F2: 1}
    var v39 int = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
