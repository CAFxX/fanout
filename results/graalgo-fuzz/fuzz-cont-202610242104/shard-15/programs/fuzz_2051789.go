package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint8; F2 string }
type GenS3 struct { F0 int; F1 bool; F2 uint }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS5 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS6 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 bool }

func main() {
    var v1 bool = (false || (1 <= 901))
    v2 := make(map[uint32]string)
    var v3 float64 = (1e10 - 0.0)
    if (v1 || (14 == 1)) {
        var v4 int8 = -57
        _ = v4
    }
    if ((0.0 != float32(3.14159)) || false) {
        if ((3.14159 / float32(v3)) > 0.0) {
            v5 := GenS1{F0: 0}
            v3 += float64(0.1)
            v6 := GenS2{F0: "foo bar", F1: 1, F2: ("" + "\n\t")}
            _ = v5
            _ = v6
        } else {
            switch ("hello" + "") {
            case "hello":
                var v7 int = -1
                _ = v7
            }
            v8 := GenS3{F0: 32768, F1: v1, F2: 24}
            var v9 uint8 = 0
            _ = v8
            _ = v9
        }
        var v10 int8 = 1
        _ = v10
    } else {
        var v11 uint16 = 1
        _ = v11
    }
    switch -2147483648 {
    case 575:
        v3 -= 0.1
    case 522:
        v12 := []int{119, 2147483647, 32767, 2147483647}
        _ = v12
    case -471:
        if true {
            v13 := GenS4{F0: 1, F1: 2147483648, F2: 409}
            var v14 float64 = 1.5
            v15 := GenS3{F0: -21, F1: v1, F2: 514}
            _ = v13
            _ = v14
            _ = v15
        }
    }
    v16 := GenS5{F0: 486, F1: 0, F2: 149}
    fmt.Println(v3)
    fmt.Println(v3, v3)
    var v17 bool = false
    fmt.Println(v3, v3, v1)
    fmt.Println(v3)
    if true {
        switch "a" {
        case "a":
            v18 := GenS6{F0: 321, F1: 228, F2: 72}
            _ = v18
        case "\n\t":
            v17 = false
        case "hello":
            if v17 {
                var v19 float64 = (3.14159 + v3)
                _ = v19
            }
        default:
            if v17 {
                fmt.Println(v17, v17, v3)
                switch "hello" {
                case "a":
                    var v20 uint64 = 602
                    _ = v20
                case "foo bar":
                    v21 := GenS6{F0: 660, F1: 127, F2: 255}
                    _ = v21
                case "\n\t":
                    v22 := GenS7{F0: 1}
                    _ = v22
                }
                for fi23 := 0; fi23 < 12; fi23++ {
                    for fi24 := 0; fi24 < 10; fi24++ {
                        var v25 uint32 = 105
                        if (694 > 42) {
                            v26 := v2[1]
                            fmt.Println(v26)
                            _ = v26
                        } else {
                            var v27 int8 = -91
                            v1 = (!v17)
                            _ = v27
                        }
                        v28 := GenS8{F0: -581}
                        _ = fi24
                        _ = v25
                        _ = v28
                    }
                    v17 = v1
                    _ = fi23
                }
            } else {
                v3 += -2.25
                v29 := GenS9{F0: ("\n\t" >= "hello")}
                var v30 bool = v1
                _ = v29
                _ = v30
            }
        }
        switch -508 {
        case -2147483648:
            v31 := GenS2{F0: "\n\t", F1: 29, F2: (("\n\t" + ("" + (("" + (("hello" + "hello") + ("a" + ("" + "\n\t")))) + ("a" + "")))) + "a")}
            _ = v31
        case -438:
            v17 = v17
        }
        if (v1 || (2147483647 < 1)) {
            if (892 == 65535) {
                v32 := v2[2]
                if true {
                    var v33 float32 = 0.0
                    var v34 uint32 = 234
                    _ = v33
                    _ = v34
                } else {
                    var v35 uint8 = 98
                    v32 = v32
                    _ = v35
                }
                _ = v32
            }
            var v36 bool = (("a" + ("" + ("" + ("\n\t" + ("\n\t" + "foo bar"))))) >= "hello")
            for fi37 := 0; fi37 < 13; fi37++ {
                var v38 uint64 = 380
                switch "a" {
                case "a":
                    v39 := GenS5{F0: v38, F1: 1, F2: 2}
                    _ = v39
                case "\n\t":
                    v1 = true
                case "foo bar":
                    if (0.0 != float32(-2.25)) {
                        var v40 uint32 = 426
                        _ = v40
                    } else {
                        v41 := GenS3{F0: (fi37 & (fi37 & (fi37 + int(v38)))), F1: (127 <= 127), F2: 0}
                        _ = v41
                    }
                default:
                    v1 = v36
                }
                _ = fi37
                _ = v38
            }
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
