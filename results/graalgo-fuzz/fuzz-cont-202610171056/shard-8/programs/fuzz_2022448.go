package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int32 }
type GenS2 struct { F0 string; F1 string; F2 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint; F1 int; F2 string }

func main() {
    var v1 uint16 = 42
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    var v4 bool = (false || (!true))
    var v5 int = 605
    var v6 uint16 = (v1 >> ((42) & 63))
    if (!(v4 && (!true))) {
        v1 = 2
        v7 := GenS1{F0: 682, F1: "foo bar", F2: -1}
        _ = v7
    } else {
        v8 := []int{(v5 >> ((878) & 63))}
        var v9 uint = 1
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 599, F1: "\n\t", F2: -1}
    var v11 uint16 = v1
    if v4 {
        v5 -= v5
        v6 = v6
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 int16 = -1
            for fi14 := 0; fi14 < 4; fi14++ {
                var v15 uint64 = 521
                fmt.Println(v1, v13, fi12)
                if v4 {
                    var v16 float32 = 0.1
                    if v4 {
                        var v17 uint32 = 979
                        v4 = true
                        var v18 float32 = v16
                        _ = v17
                        _ = v18
                    } else {
                        var v19 uint32 = 42
                        var v20 uint64 = (v15 - v15)
                        _ = v19
                        _ = v20
                    }
                    v21 := GenS1{F0: 43, F1: ("a" + ("foo bar" + ("\n\t" + "foo bar"))), F2: 42}
                    v22 := GenS2{F0: ("" + ""), F1: "hello", F2: 4294967295}
                    var v23 int = -632
                    _ = v16
                    _ = v21
                    _ = v22
                    _ = v23
                } else {
                    var v24 float64 = -2.25
                    _ = v24
                }
                _ = fi14
                _ = v15
            }
            _ = fi12
            _ = v13
        }
    } else {
        switch "foo bar" {
        case "hello":
            if false {
                fmt.Println(v4)
                var v25 int64 = 1
                switch v5 {
                case 9223372036854775807:
                    var v26 uint64 = 81
                    _ = v26
                case -250:
                    v11 -= (v1 + 42)
                }
                _ = v25
            } else {
                v5 += v5
                fmt.Println(v5, v11, v5)
            }
        default:
            v27 := GenS3{F0: (v5 >> ((0) & 63))}
            _ = v27
        }
    }
    v4 = v4
    fmt.Println(v11, v11, v4)
    v28 := GenS4{F0: 42, F1: uint8(v11)}
    var v29 int8 = 100
    fmt.Println(v5)
    v6 = 255
    var v30 bool = (2 < (^(v11)))
    for fi31 := 0; fi31 < 16; fi31++ {
        switch fi31 {
        case 9223372036854775807:
            if (true || v4) {
                switch (fi31 >> ((65535) & 63)) {
                case -1:
                    v32 := GenS3{F0: (-(v5))}
                    _ = v32
                case 662:
                    v29 -= (v29 >> ((42) & 63))
                default:
                    v33 := GenS5{F0: 42, F1: (^(fi31)), F2: ((("a" + "\n\t") + "\n\t") + "foo bar")}
                    _ = v33
                }
            } else {
                for fi34 := 0; fi34 < 7; fi34++ {
                    v35 := GenS2{F0: ("a" + ("" + ("hello" + "hello"))), F1: (("hello" + "hello") + "hello"), F2: 2}
                    _ = fi34
                    _ = v35
                }
                v1 += v6
                v36 := GenS3{F0: v5}
                _ = v36
            }
        case -9223372036854775808:
            fmt.Println(v29, v29)
        default:
            var v37 int = v5
            _ = v37
        }
        _ = fi31
    }
    for fi38 := 0; fi38 < 13; fi38++ {
        var v39 bool = false
        v39 = (4294967295 <= 2)
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
