package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 uint64 = 498
    v2 := make(map[uint32]bool)
    var v3 bool = true
    v4 := []int{-985, -2147483648, 32767, 9223372036854775807}
    for fi5 := 0; fi5 < 18; fi5++ {
        if len(v4) > 0 {
            v4[0] = 42
        }
        _ = fi5
    }
    if true {
        var v6 int64 = 710
        v1 = (^(v1))
        _ = v6
    } else {
        v7 := GenS1{F0: 42, F1: 255, F2: -682}
        var v8 int8 = -117
        if v3 {
            if (true && true) {
                fmt.Println(v8, v8, v8)
            } else {
                v9 := v2[2]
                _ = v9
            }
            v4 = append(v4, 42)
        } else {
            for fi10 := 0; fi10 < 9; fi10++ {
                if (!(0.1 > 3.14159)) {
                    var v11 int64 = -878
                    fmt.Println(v11, v11, fi10)
                    v8 += (v8 + v8)
                    _ = v11
                }
                _ = fi10
            }
        }
        _ = v7
        _ = v8
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        v3 = v3
        if (!(!(42 == -1))) {
            var v13 float64 = 0.1
            var v14 uint64 = (v1 % ((884) | 1))
            fmt.Println(v13)
            _ = v13
            _ = v14
        } else {
            v3 = ("a" >= "hello")
        }
        _ = fi12
    }
    var v15 uint8 = 1
    v16 := v2[65535]
    switch ("a" + "\n\t") {
    case "a":
        v17 := GenS2{F0: 0}
        _ = v17
    case "foo bar":
        fmt.Println(v16, v3)
    case "\n\t":
        v18 := GenS2{F0: -983}
        _ = v18
    }
    if v3 {
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 string = ("a" + "")
            _ = fi19
            _ = v20
        }
        v2[42] = (3.14159 == (float32(1e10) + 1.5))
    } else {
        if v3 {
            var v21 uint64 = 42
            var v22 int = 360
            var v23 uint32 = 65535
            _ = v21
            _ = v22
            _ = v23
        } else {
            fmt.Println(v15, v1)
            var v24 int32 = -776
            v15 = v15
            _ = v24
        }
        for fi25 := 0; fi25 < 15; fi25++ {
            v16 = (((!v16) && false) || v16)
            var v26 bool = v3
            _ = fi25
            _ = v26
        }
    }
    var v27 float64 = (-2.25 + 3.14159)
    var v28 float32 = 0.1
    for fi29 := 0; fi29 < 9; fi29++ {
        switch "hello" {
        case "":
            if (((0.1 * v28) - v28) >= v28) {
                for fi30 := 0; fi30 < 19; fi30++ {
                    v28 = (-2.25 * (v28 * v28))
                    var v31 bool = (1 != 32767)
                    _ = fi30
                    _ = v31
                }
                for fi32 := 0; fi32 < 1; fi32++ {
                    switch "hello" {
                    case "hello":
                        var v33 int8 = 17
                        _ = v33
                    default:
                        var v34 string = ("hello" + "foo bar")
                        _ = v34
                    }
                    var v35 int = 42
                    v36 := GenS2{F0: 42}
                    _ = len(v2)
                    _ = len(v2)
                    v37 := GenS3{F0: v35, F1: v35, F2: 1}
                    _ = fi32
                    _ = v35
                    _ = v36
                    _ = v37
                }
            }
        default:
            if ((255 > 763) || v16) {
                for fi38 := 0; fi38 < 9; fi38++ {
                    v39 := GenS4{F0: 0, F1: 986}
                    fmt.Println(fi38)
                    var v40 uint16 = 0
                    var v41 uint64 = v1
                    v4 = append(v4, (fi29 << ((18446744073709551615) & 63)))
                    _ = fi38
                    _ = v39
                    _ = v40
                    _ = v41
                }
            }
        }
        if v3 {
            var v42 uint64 = v1
            _ = v42
        }
        v27 += 3.14159
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
