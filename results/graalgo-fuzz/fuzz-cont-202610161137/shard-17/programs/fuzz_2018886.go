package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 string; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint; F1 bool; F2 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 int8; F2 string }
type GenS6 struct { F0 int8; F1 int32; F2 int8 }
type GenS7 struct { F0 uint32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]bool)
    var v3 uint32 = 42
    var v4 int64 = 42
    v5 := GenS1{F0: 0, F1: 1}
    if true {
        for fi6 := 0; fi6 < 15; fi6++ {
            if (!(!false)) {
                for fi7 := 0; fi7 < 18; fi7++ {
                    fmt.Println(v3, fi7)
                    var v8 int = int(v4)
                    if false {
                        delete(v1, ("" + "a"))
                        v8 = v8
                    }
                    _ = fi7
                    _ = v8
                }
                switch fi6 {
                case -1:
                    var v9 int64 = (v4 >> ((0) & 63))
                    _ = v9
                case 32768:
                    for fi10 := 0; fi10 < 18; fi10++ {
                        v11 := v2[(v3 << ((763) & 63))]
                        v3 = 4294967295
                        _ = fi10
                        _ = v11
                    }
                }
            }
            _ = fi6
        }
        if false {
            fmt.Println(v4, v3)
            for fi12 := 0; fi12 < 10; fi12++ {
                fmt.Println(v3, v3, v4)
                var v13 float64 = -2.25
                _ = fi12
                _ = v13
            }
            v3 += v3
        } else {
            _ = len(v1)
            v14 := GenS2{F0: (("\n\t" + ((("hello" + "hello") + "a") + "a")) + "hello"), F1: 0, F2: 502}
            var v15 uint64 = 1
            _ = v14
            _ = v15
        }
    } else {
        switch "a" {
        case "hello":
            if (42 != v4) {
                var v16 uint32 = (v3 * 0)
                v16 = v16
                _ = v16
            } else {
                if false {
                    fmt.Println(v4)
                    var v17 int32 = 696
                    _ = v17
                } else {
                    var v18 uint = 255
                    if false {
                        if (true && true) {
                            if true {
                                var v19 uint32 = v3
                                _ = v19
                            } else {
                                fmt.Println(v3, v4, v3)
                                v20 := GenS3{F0: v18, F1: false, F2: ("" + "foo bar")}
                                v21 := GenS4{F0: "a"}
                                _ = v20
                                _ = v21
                            }
                            if (0 == 32767) {
                                v22 := []int{}
                                var v23 uint32 = (v3 + 42)
                                _ = v22
                                _ = v23
                            } else {
                                v24 := GenS5{F0: 42, F1: 100, F2: "a"}
                                if (!false) {
                                    var v25 uint32 = (v3 * (v3 ^ 550))
                                    v26 := v2[v25]
                                    fmt.Println(v25, v4)
                                    _ = v25
                                    _ = v26
                                }
                                _ = v24
                            }
                        } else {
                            switch ("a" + "\n\t") {
                            case "":
                                v27 := GenS1{F0: 234, F1: -128}
                                _ = v27
                            }
                            var v28 uint16 = 255
                            var v29 uint8 = 128
                            _ = v28
                            _ = v29
                        }
                    }
                    var v30 uint64 = 0
                    _ = v18
                    _ = v30
                }
                v31 := GenS5{F0: 0, F1: 100, F2: "foo bar"}
                _ = v31
            }
        }
    }
    fmt.Println(v3, v4)
    switch -159 {
    case -2147483648:
        v32 := v1["foo bar"]
        _ = v32
    case 1:
        var v33 uint8 = 2
        _ = v33
    case 4294967295:
        fmt.Println(v3, v4)
    default:
        switch 42 {
        case -1:
            _ = len(v2)
        }
    }
    for fi34 := 0; fi34 < 20; fi34++ {
        v3 -= v3
        _ = fi34
    }
    v35 := GenS6{F0: 100, F1: 42, F2: -1}
    v3 -= v3
    for fi36 := 0; fi36 < 3; fi36++ {
        v1[("\n\t" + "hello")] = (v3 & v3)
        v3 -= 866
        _ = fi36
    }
    v37 := GenS1{F0: 1, F1: 100}
    v3 += v3
    v38 := GenS7{F0: (v3 % ((v3) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
