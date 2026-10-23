package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 bool; F1 int16; F2 string }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 uint16 = 786
    v2 := make(map[uint32]int)
    var v3 int16 = 1
    var v4 uint16 = (v1 ^ v1)
    for fi5 := 0; fi5 < 14; fi5++ {
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 int16 = v3
            v4 = v4
            var v8 int = fi5
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := v2[42]
        _ = fi5
        _ = v9
    }
    var v10 uint8 = 95
    switch "a" {
    case "a":
        fmt.Println(v4, v10)
    }
    switch "foo bar" {
    case "\n\t":
        for fi11 := 0; fi11 < 6; fi11++ {
            if true {
                v12 := GenS1{F0: -2147483648, F1: v1}
                v10 += v10
                _ = v12
            }
            v13 := []int{(fi11 + (fi11 / ((9223372036854775807) | 1))), fi11, fi11}
            _ = fi11
            _ = v13
        }
    case "a":
        v4 -= v4
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        if true {
            var v15 int = (fi14 % ((fi14) | 1))
            for fi16 := 0; fi16 < 6; fi16++ {
                fmt.Println(fi14, v10)
                v17 := []int{(v15 | (v15 >> ((uint(v10)) & 63)))}
                if (!(-2147483648 >= 42)) {
                    delete(v2, 2147483647)
                    switch (("foo bar" + "hello") + "a") {
                    case "foo bar":
                        v4 = (^(v4))
                    case "hello":
                        v18 := GenS1{F0: 1, F1: 42}
                        _ = v18
                    case "":
                        switch "hello" {
                        case "foo bar":
                            var v19 int32 = int32(fi14)
                            _ = v19
                        }
                    default:
                        v4 += (v1 | (^(v1)))
                    }
                    v3 = v3
                }
                v20 := GenS1{F0: 1, F1: 65535}
                _ = fi16
                _ = v17
                _ = v20
            }
            _ = v15
        } else {
            if ((1.5 * ((1e10 + 1e10) * 0.1)) <= 1.5) {
                v21 := GenS1{F0: -1, F1: v1}
                fmt.Println(v3)
                _ = v21
            }
            v22 := []int{(fi14 * (fi14 / ((32767) | 1))), fi14, (fi14 & (fi14 >> ((183) & 63))), 1}
            _ = v22
        }
        if false {
            v23 := []int{(fi14 & int(v10)), int(v3), 1}
            v24 := GenS2{F0: (!(int32(v10) < int32(v10))), F1: v3, F2: "hello"}
            _ = v23
            _ = v24
        } else {
            v1 += v1
            v3 += 42
        }
        _ = fi14
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        v1 -= (v4 >> ((255) & 63))
        fmt.Println(v10, v3)
        _ = fi25
    }
    var v26 bool = false
    v1 = (v4 << ((351) & 63))
    fmt.Println(v26)
    _ = len(v2)
    v10 -= v10
    var v27 bool = (!v26)
    v10 = v10
    if (4294967295 < 0) {
        if (!v26) {
            v26 = true
            var v28 uint16 = 0
            for fi29 := 0; fi29 < 3; fi29++ {
                var v30 uint32 = 1
                switch 9223372036854775807 {
                case 32767:
                    var v31 uint64 = 2
                    _ = v31
                case 4294967295:
                    v32 := GenS2{F0: ((1e10 - (float64(-2.25) * 1.5)) <= 3.14159), F1: v3, F2: "\n\t"}
                    _ = v32
                default:
                    v10 -= (v10 >> ((4294967295) & 63))
                }
                _ = fi29
                _ = v30
            }
            _ = v28
        } else {
            var v33 uint64 = 1
            _ = v33
        }
        var v34 int16 = (^(v3))
        _ = v34
    } else {
        v35 := GenS3{F0: "a", F1: 5}
        v3 += 42
        fmt.Println(v27, v1)
        _ = v35
    }
    fmt.Println(v1, v1, v1)
    v36 := GenS4{F0: -25}
    var v37 uint = 342
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
