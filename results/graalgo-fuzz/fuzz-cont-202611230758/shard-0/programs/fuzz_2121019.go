package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 float64 = (1.5 / 3.14159)
    var v2 int32 = 548
    var v3 int64 = 42
    v4 := make(map[string]int)
    v5 := GenS1{F0: 2}
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: 42}
        v2 += v2
        v3 = (v3 >> ((255) & 63))
        _ = fi6
        _ = v7
    }
    fmt.Println(v3)
    switch 32767 {
    case -58:
        switch "" {
        case "\n\t":
            v8 := []int{2147483647, 693}
            _ = v8
        case "foo bar":
            for fi9 := 0; fi9 < 7; fi9++ {
                v2 += v2
                if false {
                    var v10 bool = (fi9 >= (-(fi9)))
                    v11 := GenS1{F0: 65535}
                    for fi12 := 0; fi12 < 1; fi12++ {
                        v10 = (767 <= 65535)
                        v13 := GenS1{F0: 343}
                        _ = fi12
                        _ = v13
                    }
                    _ = v10
                    _ = v11
                }
                switch (fi9 >> ((65535) & 63)) {
                case 1000:
                    var v14 int8 = -78
                    _ = v14
                case 4294967295:
                    var v15 uint32 = 306
                    _ = v15
                case -33:
                    v2 = (v2 & 1)
                default:
                    if true {
                        if ((32767 > 1) && (v2 != (-(v2)))) {
                            v1 += v1
                        } else {
                            var v16 float32 = (3.14159 / 1.0)
                            _ = v16
                        }
                    }
                }
                _ = fi9
            }
        default:
            if false {
                fmt.Println(v3, v2)
                v1 = float64((1e10 - 1.5))
                if true {
                    v17 := GenS2{F0: 65535, F1: 533}
                    v18 := GenS3{F0: 262, F1: -997, F2: 2}
                    var v19 int16 = 42
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    switch (("hello" + "") + "\n\t") {
                    case "hello":
                        for fi20 := 0; fi20 < 18; fi20++ {
                            if false {
                                v21 := v4["hello"]
                                _ = v21
                            } else {
                                fmt.Println(v3, fi20)
                            }
                            _ = fi20
                        }
                    case "\n\t":
                        var v22 uint8 = 255
                        _ = v22
                    }
                }
            } else {
                var v23 string = ("a" + "a")
                var v24 uint8 = 1
                _ = v23
                _ = v24
            }
        }
    }
    if false {
        var v25 int8 = 127
        _ = v25
    } else {
        _ = len(v4)
        var v26 int = 127
        _ = v26
    }
    if true {
        v4["a"] = 127
    } else {
        fmt.Println(v1)
    }
    switch 9223372036854775807 {
    case 127:
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 int32 = v2
            _ = fi27
            _ = v28
        }
    case 1:
        if (!true) {
            var v29 uint64 = 928
            _ = v29
        }
    default:
        v30 := []int{}
        _ = v30
    }
    v31 := GenS2{F0: 32767, F1: 0}
    var v32 uint32 = 320
    v33 := GenS4{F0: 1}
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
