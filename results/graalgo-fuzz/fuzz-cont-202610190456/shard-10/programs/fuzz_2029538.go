package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 string = ("" + ("" + "\n\t"))
    var v2 uint8 = 255
    var v3 int64 = 847
    v4 := make(map[int]int)
    if (487 < 984) {
        if true {
            if true {
                var v5 uint = 0
                if false {
                    var v6 int8 = 42
                    _ = v6
                } else {
                    for fi7 := 0; fi7 < 4; fi7++ {
                        v5 = 887
                        var v8 bool = true
                        v5 += v5
                        _ = fi7
                        _ = v8
                    }
                }
                _ = v5
            }
        }
        delete(v4, 128)
    } else {
        for fi9 := 0; fi9 < 20; fi9++ {
            for fi10 := 0; fi10 < 15; fi10++ {
                v11 := GenS1{F0: 814}
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 13; fi12++ {
                switch (fi12 - fi12) {
                case -707:
                    if true {
                        _ = len(v4)
                    }
                case 127:
                    v3 = v3
                default:
                    switch "hello" {
                    case "a":
                        var v13 uint32 = 65535
                        _ = v13
                    case "foo bar":
                        var v14 int = fi12
                        _ = v14
                    case "hello":
                        v1 = (v1 + "hello")
                    default:
                        if (!true) {
                            switch v1 {
                            case "":
                                var v15 float32 = 0.1
                                _ = v15
                            default:
                                var v16 string = ""
                                _ = v16
                            }
                        }
                    }
                }
                _ = fi12
            }
            _ = fi9
        }
        v17 := GenS1{F0: 32767}
        _ = len(v4)
        _ = v17
    }
    fmt.Println(v1)
    var v18 int = 9223372036854775807
    if (v3 > (v3 + (^(v3)))) {
        v3 = -1
        v19 := GenS1{F0: 843}
        v18 += v18
        _ = v19
    } else {
        if true {
            v2 = v2
            v3 += (v3 + -1)
        }
        v3 -= -154
        for fi20 := 0; fi20 < 12; fi20++ {
            fmt.Println(v18, v1, v1)
            _ = fi20
        }
    }
    var v21 float32 = 0.1
    for fi22 := 0; fi22 < 2; fi22++ {
        switch -583 {
        case 464:
            if false {
                v21 -= 0.0
                switch v18 {
                case 32767:
                    var v23 int8 = 127
                    _ = v23
                case -9223372036854775808:
                    switch ("foo bar" + (v1 + "hello")) {
                    case "foo bar":
                        v3 = v3
                    }
                case 128:
                    if (!false) {
                        v24 := GenS2{F0: 255, F1: 833, F2: 998}
                        _ = v24
                    }
                default:
                    var v25 int64 = v3
                    _ = v25
                }
                v26 := GenS3{F0: 1}
                _ = v26
            }
        case 611:
            v1 = v1
        case 2147483647:
            for fi27 := 0; fi27 < 18; fi27++ {
                v28 := []int{(-(fi22)), 486, v18}
                var v29 uint32 = 1
                v21 += -2.25
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        _ = fi22
    }
    v3 -= 42
    var v30 uint64 = 4294967295
    var v31 int64 = -1
    switch (v1 + (v1 + "")) {
    case "a":
        var v32 int64 = (v3 << ((1) & 63))
        _ = v32
    case "foo bar":
        if true {
            v2 -= (v2 ^ v2)
            v31 += (^(v31))
            v18 = 32767
        }
    }
    v2 = (v2 / ((v2) | 1))
    var v33 int = v18
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v4
}
