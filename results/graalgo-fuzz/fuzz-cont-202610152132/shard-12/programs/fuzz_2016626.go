package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS9 struct { F0 int; F1 int64 }
func genFn0() int8 {
    for fi4 := 0; fi4 < 12; fi4++ {
        fmt.Println(fi4)
        var v5 int32 = -1
        _ = fi4
        _ = v5
    }
    switch ("\n\t" + "\n\t") {
    case "":
        var v6 uint16 = 0
        _ = v6
    case "a":
        var v7 int64 = -191
        _ = v7
    }
    return 6
}
func genFn1() uint8 {
    if (!false) {
        v8 := GenS1{F0: 0, F1: 1000}
        _ = v8
    }
    v9 := []int{9223372036854775807, -230, 598, -173}
    v10 := GenS2{F0: 164}
    _ = v10
    _ = v9
    return 255
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[string]int)
    var v3 uint16 = 91
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := GenS3{F0: -2147483648, F1: 65535}
            _ = fi12
            _ = v13
        }
        fmt.Println(fi11, v1, fi11)
        fmt.Println(v1, v1, v1)
        _ = fi11
    }
    v14 := GenS1{F0: 1, F1: 42}
    for fi15 := 0; fi15 < 9; fi15++ {
        switch (-(fi15)) {
        case 0:
            v16 := GenS4{F0: 416, F1: -602}
            _ = v16
        case 32768:
            v3 = v3
        }
        v17 := GenS5{F0: (v3 * 32767)}
        _ = fi15
        _ = v17
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        fmt.Println(v3)
        _ = fi18
    }
    if true {
        var v19 uint16 = (v3 | (v3 / (((^(v3))) | 1)))
        var v20 uint32 = 2
        switch "" {
        case "foo bar":
            if (!true) {
                var v21 uint = 1
                v22 := v2[""]
                _ = v21
                _ = v22
            }
        default:
            v23 := GenS5{F0: v3}
            _ = v23
        }
        _ = v19
        _ = v20
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        if (38 < 1) {
            v1 -= 1
        } else {
            for fi25 := 0; fi25 < 4; fi25++ {
                fmt.Println(fi25, v1, v1)
                v3 -= v3
                _ = fi25
            }
            for fi26 := 0; fi26 < 11; fi26++ {
                if false {
                    v27 := GenS4{F0: 18446744073709551615, F1: -2147483648}
                    var v28 float32 = (-2.25 - -2.25)
                    _ = v27
                    _ = v28
                } else {
                    var v29 int32 = 249
                    var v30 bool = false
                    v29 = (v29 * (v29 / ((v29) | 1)))
                    var v31 uint32 = 2147483648
                    _ = v29
                    _ = v30
                    _ = v31
                }
                _ = fi26
            }
            v1 = (v1 / ((-966) | 1))
        }
        v1 = v1
        if (false && false) {
            for fi32 := 0; fi32 < 6; fi32++ {
                switch -712 {
                case 1:
                    if false {
                        fmt.Println(fi32)
                    } else {
                        v33 := GenS3{F0: 0, F1: 724}
                        _ = len(v2)
                        var v34 uint64 = 4294967295
                        _ = v33
                        _ = v34
                    }
                default:
                    var v35 float32 = float32(1e10)
                    _ = v35
                }
                v1 += (-(v1))
                v36 := v2[(("\n\t" + "a") + "hello")]
                v37 := GenS6{F0: 42, F1: genFn0()}
                v38 := GenS7{F0: (fi32 % (((-(fi24))) | 1))}
                _ = fi32
                _ = v36
                _ = v37
                _ = v38
            }
            delete(v2, "a")
        }
        _ = fi24
    }
    var v39 int8 = 127
    v40 := GenS8{F0: 2, F1: 65535, F2: 887}
    v39 -= 42
    var v41 float32 = 1.5
    v42 := v2[(("" + "a") + "a")]
    v43 := GenS9{F0: (v42 * (^(v42))), F1: -9223372036854775808}
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
