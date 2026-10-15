package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 int; F2 uint16 }

func main() {
    var v1 bool = false
    v2 := []string{}
    var v3 int32 = 368
    var v4 int8 = 99
    if true {
        fmt.Println(v3)
    } else {
        v5 := GenS1{F0: 2, F1: 0, F2: "a"}
        switch ("foo bar" + "hello") {
        case "":
            var v6 uint16 = 605
            _ = v6
        case "foo bar":
            v3 -= (v3 << ((42) & 63))
        default:
            for fi7 := 0; fi7 < 13; fi7++ {
                switch 4294967295 {
                case 128:
                    v4 = (v4 ^ (v4 % (((v4 & 103)) | 1)))
                }
                _ = fi7
            }
        }
        v3 = -1
        _ = v5
    }
    v1 = ((42 == -1) || true)
    var v8 uint32 = 2
    v2 = append(v2, "a")
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 string = ""
        _ = fi9
        _ = v10
    }
    var v11 int32 = (v3 ^ v3)
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int16 = 1000
        _ = fi12
        _ = v13
    }
    v1 = v1
    var v14 int = -2147483648
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v4, v11)
    case "hello":
        switch (("hello" + "") + "foo bar") {
        case "\n\t":
            switch "hello" {
            case "hello":
                var v15 uint32 = 65535
                _ = v15
            default:
                v16 := GenS1{F0: 42, F1: 255, F2: "foo bar"}
                _ = v16
            }
        default:
            var v17 uint16 = uint16(v8)
            _ = v17
        }
    default:
        if (2147483647 != v11) {
            for fi18 := 0; fi18 < 19; fi18++ {
                var v19 float32 = 0.1
                v20 := GenS2{F0: (!v1)}
                for fi21 := 0; fi21 < 19; fi21++ {
                    var v22 uint32 = 42
                    _ = len(v2)
                    _ = len(v2)
                    _ = fi21
                    _ = v22
                }
                _ = fi18
                _ = v19
                _ = v20
            }
            for fi23 := 0; fi23 < 3; fi23++ {
                v11 -= v3
                for fi24 := 0; fi24 < 8; fi24++ {
                    var v25 uint16 = 255
                    _ = fi24
                    _ = v25
                }
                v11 = v3
                _ = fi23
            }
            if (4294967295 != 4294967295) {
                var v26 uint16 = 255
                for fi27 := 0; fi27 < 18; fi27++ {
                    v11 = v3
                    var v28 uint64 = 1
                    for fi29 := 0; fi29 < 7; fi29++ {
                        switch "foo bar" {
                        case "hello":
                            var v30 float64 = (0.0 * 1e10)
                            _ = v30
                        case "foo bar":
                            switch (fi27 ^ fi27) {
                            case 128:
                                if v1 {
                                    v31 := GenS3{F0: v28}
                                    _ = v31
                                }
                            case 2147483647:
                                v32 := GenS4{F0: "\n\t", F1: -1, F2: (v26 / ((v26) | 1))}
                                _ = v32
                            case -1:
                                v8 = (v8 + 2)
                            default:
                                _ = len(v2)
                            }
                            v33 := GenS3{F0: (v28 * v28)}
                            _ = v33
                        default:
                            v34 := GenS2{F0: v1}
                            _ = v34
                        }
                        switch int(v4) {
                        case 9223372036854775807:
                            switch int(v28) {
                            case 9223372036854775807:
                                switch "hello" {
                                case "foo bar":
                                    var v35 int16 = 0
                                    _ = v35
                                }
                            case -1:
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                            case -208:
                                var v36 int8 = (v4 / (((-(v4))) | 1))
                                _ = v36
                            }
                        case 0:
                            v4 += (v4 << ((255) & 63))
                        case 32768:
                            var v37 uint64 = v28
                            _ = v37
                        }
                        v3 = (v11 - v11)
                        _ = fi29
                    }
                    _ = fi27
                    _ = v28
                }
                _ = v26
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch "a" {
            case "foo bar":
                fmt.Println(v14)
            case "hello":
                fmt.Println(v14)
            }
        }
    }
    v1 = (!((((true || v1) || ((!v1) || (255 <= 990))) || (1 != 3)) || (42 <= -9223372036854775808)))
    var v38 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v8
}
