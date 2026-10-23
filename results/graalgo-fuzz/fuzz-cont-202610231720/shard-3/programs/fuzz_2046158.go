package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 uint32; F2 int64 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 uint8; F2 int }
type GenS7 struct { F0 string; F1 uint64 }

func main() {
    var v1 int32 = -1
    var v2 int8 = -10
    var v3 int8 = v2
    v4 := make(map[string]int)
    v5 := make(map[string]bool)
    v6 := GenS1{F0: 42}
    v1 = v1
    var v7 int16 = -953
    v7 = 1
    v1 = (v1 + -648)
    var v8 int16 = int16(v1)
    if (255 == 117) {
        v9 := GenS2{F0: (v1 | v1), F1: 250, F2: -9223372036854775808}
        v7 = (v7 - (^(v7)))
        if false {
            var v10 uint16 = 438
            v11 := GenS1{F0: 9223372036854775807}
            if ((1 > 42) || (!((true && true) || true))) {
                for fi12 := 0; fi12 < 15; fi12++ {
                    v13 := GenS3{F0: 42, F1: (v1 | v1)}
                    _ = fi12
                    _ = v13
                }
            } else {
                var v14 uint16 = (v10 / ((v10) | 1))
                _ = v14
            }
            _ = v10
            _ = v11
        }
        _ = v9
    }
    if false {
        var v15 string = "a"
        var v16 uint8 = 2
        _ = v15
        _ = v16
    }
    v17 := GenS4{F0: 0, F1: 68}
    switch 32767 {
    case 32768:
        if false {
            v18 := GenS3{F0: -238, F1: v1}
            var v19 float64 = (3.14159 - float64((float32(-2.25) - 1.5)))
            _ = v18
            _ = v19
        } else {
            v2 = v2
            v8 = v8
        }
    case 1:
        v5["hello"] = (3.14159 >= 0.0)
    case 128:
        switch -324 {
        case -590:
            v1 = v1
        }
    }
    _ = len(v5)
    v20 := GenS5{F0: (!(594 < 1))}
    fmt.Println(v3)
    var v21 uint8 = 1
    switch -862 {
    case -954:
        var v22 float32 = (3.14159 - -2.25)
        _ = v22
    case 797:
        if false {
            for fi23 := 0; fi23 < 18; fi23++ {
                v24 := GenS6{F0: ((v3 * v3) <= (v3 & (v3 + 42))), F1: 2, F2: fi23}
                for fi25 := 0; fi25 < 3; fi25++ {
                    v26 := GenS6{F0: (213 <= 1), F1: (v21 << ((2) & 63)), F2: (fi25 * int(v3))}
                    var v27 uint = 42
                    switch "foo bar" {
                    case "a":
                        if false {
                            v1 -= v1
                            v4["\n\t"] = fi23
                            v28 := GenS2{F0: int32(fi23), F1: 236, F2: 0}
                            var v29 uint32 = 467
                            v2 += v3
                            v30 := GenS1{F0: 789}
                            _ = v28
                            _ = v29
                            _ = v30
                        } else {
                            switch (^(fi25)) {
                            case -163:
                                var v31 uint8 = 107
                                if true {
                                    var v32 uint8 = (v21 << ((v27) & 63))
                                    fmt.Println(v1)
                                    _ = v32
                                } else {
                                    v33 := GenS7{F0: ("\n\t" + ("hello" + "a")), F1: 2}
                                    v2 += v2
                                    _ = v33
                                }
                                _ = v31
                            default:
                                v34 := v5["\n\t"]
                                _ = v34
                            }
                            if false {
                                switch "a" {
                                case "\n\t":
                                    v21 += (v21 << ((4294967295) & 63))
                                case "foo bar":
                                    v35 := GenS6{F0: (("foo bar" == ("\n\t" + ("hello" + (("foo bar" + "hello") + (("hello" + "") + "\n\t"))))) && ((!false) && false)), F1: 127, F2: fi25}
                                    _ = v35
                                }
                                v36 := []int{32768}
                                _ = v36
                            }
                            fmt.Println(v21, fi25)
                        }
                    case "\n\t":
                        var v37 uint16 = 262
                        _ = v37
                    case "":
                        v27 -= 652
                    }
                    var v38 float32 = ((1.5 + (0.0 + (0.0 - 0.0))) - 0.1)
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v38
                }
                _ = fi23
                _ = v24
            }
            v39 := []int{32768, 1, 0}
            _ = len(v39)
            _ = v39
        }
    default:
        fmt.Println(v8, v1, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
