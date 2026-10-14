package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int16 = -32768
    var v2 int32 = -1
    v3 := make(map[uint32]int)
    v4 := []int{int(v2)}
    var v5 int32 = (^(v2))
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint32 = 1
        _ = fi6
        _ = v7
    }
    switch 0 {
    case 723:
        for fi8 := 0; fi8 < 2; fi8++ {
            if (42 < 42) {
                for fi9 := 0; fi9 < 1; fi9++ {
                    v5 = 352
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi9
                }
                v5 = (v5 << ((255) & 63))
            }
            _ = fi8
        }
    case -1:
        if (!true) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            v2 = v2
        }
    case 9223372036854775807:
        fmt.Println(v5)
    default:
        var v10 int = 32767
        _ = v10
    }
    v11 := GenS1{F0: -40, F1: v1}
    v2 = (v2 << ((296) & 63))
    if (!false) {
        v3[65535] = -86
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 bool = (!(3.14159 <= -2.25))
            _ = fi12
            _ = v13
        }
        if (!false) {
            if true {
                fmt.Println(v5)
                _ = len(v3)
            }
            v14 := GenS1{F0: -7, F1: v1}
            _ = v14
        } else {
            for fi15 := 0; fi15 < 14; fi15++ {
                v2 -= v5
                for fi16 := 0; fi16 < 19; fi16++ {
                    v17 := GenS1{F0: -1, F1: (v1 ^ v1)}
                    switch "\n\t" {
                    case "\n\t":
                        v2 = v2
                    case "":
                        if false {
                            if true {
                                var v18 uint64 = 4294967295
                                _ = v18
                            }
                            var v19 int = (fi15 + 1)
                            var v20 int32 = (v5 << ((148) & 63))
                            _ = v19
                            _ = v20
                        }
                    default:
                        switch "\n\t" {
                        case "a":
                            v2 += 1
                        case "foo bar":
                            switch 32768 {
                            case -2147483648:
                                var v21 int8 = 127
                                _ = v21
                            case 1:
                                v1 = -32768
                            case -908:
                                switch ("" + (("a" + (("foo bar" + "hello") + ("" + "\n\t"))) + "")) {
                                case "a":
                                    v22 := GenS2{F0: 127}
                                    _ = v22
                                case "foo bar":
                                    fmt.Println(fi15, fi16)
                                }
                            default:
                                v5 = (v5 & (v2 & v2))
                            }
                        }
                    }
                    _ = len(v4)
                    _ = fi16
                    _ = v17
                }
                _ = fi15
            }
            if true {
                delete(v3, 2)
                v1 = v1
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            switch 1 {
            case -1:
                for fi23 := 0; fi23 < 5; fi23++ {
                    var v24 int8 = 42
                    _ = fi23
                    _ = v24
                }
            case 42:
                var v25 uint32 = 651
                _ = v25
            case 128:
                var v26 uint32 = 42
                _ = v26
            default:
                var v27 int = -9223372036854775808
                _ = v27
            }
        }
    } else {
        var v28 string = "a"
        _ = v28
    }
    fmt.Println(v2)
    switch "a" {
    case "a":
        for fi29 := 0; fi29 < 17; fi29++ {
            var v30 uint64 = 1
            var v31 uint = 526
            v30 += 4294967295
            _ = fi29
            _ = v30
            _ = v31
        }
    case "":
        v3[2147483647] = 128
    case "foo bar":
        var v32 uint16 = 132
        _ = v32
    default:
        v33 := GenS2{F0: 0}
        _ = v33
    }
    switch ("hello" + "foo bar") {
    case "":
        fmt.Println(v2)
    case "a":
        var v34 int32 = (v2 * (v5 | (v5 << ((255) & 63))))
        _ = v34
    case "\n\t":
        fmt.Println(v1, v1)
    }
    var v35 bool = false
    var v36 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
