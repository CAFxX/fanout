package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 bool }
type GenS2 struct { F0 int; F1 uint; F2 int8 }
type GenS3 struct { F0 string; F1 uint16; F2 uint }
type GenS4 struct { F0 uint32; F1 int64 }

func main() {
    var v1 float32 = 3.14159
    var v2 string = (("" + "hello") + "\n\t")
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    var v5 uint64 = 4
    v6 := []int{}
    switch v2 {
    case "a":
        v2 = ((v2 + v2) + v2)
    case "":
        v7 := GenS1{F0: 2, F1: 820, F2: false}
        _ = v7
    case "hello":
        var v8 int = 127
        _ = v8
    }
    v5 = (v5 * v5)
    if true {
        var v9 int = -684
        _ = v9
    } else {
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 float32 = -2.25
            v12 := v3[(^(fi10))]
            _ = fi10
            _ = v11
            _ = v12
        }
        if (0 < 1) {
            switch "hello" {
            case "foo bar":
                var v13 int = -9223372036854775808
                _ = v13
            case "a":
                var v14 int8 = 1
                _ = v14
            }
        } else {
            var v15 uint16 = uint16(v5)
            _ = v15
        }
        switch v2 {
        case "hello":
            v2 = "hello"
        default:
            fmt.Println(v2, v5)
        }
    }
    switch ((v2 + v2) + v2) {
    case "foo bar":
        var v16 uint8 = 140
        _ = v16
    default:
        fmt.Println(v1, v5)
    }
    var v17 int64 = 703
    v1 = v1
    if false {
        v18 := GenS1{F0: 255, F1: 2, F2: false}
        switch (v2 + v2) {
        case "\n\t":
            for fi19 := 0; fi19 < 9; fi19++ {
                var v20 uint8 = 128
                v21 := GenS2{F0: fi19, F1: 255, F2: -99}
                _ = fi19
                _ = v20
                _ = v21
            }
        case "":
            if ((((v2 + v2) + v2) + v2) > v2) {
                for fi22 := 0; fi22 < 9; fi22++ {
                    var v23 string = (("foo bar" + ((v2 + (v2 + "")) + v2)) + v2)
                    _ = fi22
                    _ = v23
                }
            }
        case "hello":
            switch 2147483647 {
            case 127:
                v17 -= 713
            }
        }
        _ = v18
    } else {
        switch v2 {
        case "foo bar":
            var v24 int64 = v17
            _ = v24
        default:
            var v25 uint32 = 42
            _ = v25
        }
        v26 := GenS1{F0: 128, F1: 195, F2: true}
        _ = v26
    }
    if false {
        for fi27 := 0; fi27 < 12; fi27++ {
            for fi28 := 0; fi28 < 14; fi28++ {
                fmt.Println(fi28, v2)
                if (255 != 128) {
                    v29 := GenS3{F0: v2, F1: 2, F2: 651}
                    switch v2 {
                    case "\n\t":
                        var v30 bool = (!(42 > 174))
                        _ = v30
                    }
                    var v31 float64 = (3.14159 * 1e10)
                    if true {
                        var v32 int64 = -499
                        v17 -= (v17 - v17)
                        v31 = (v31 - (v31 / float64((v1 * v1))))
                        _ = len(v6)
                        v17 += (v17 >> ((1) & 63))
                        v17 -= (v32 >> ((18446744073709551615) & 63))
                        _ = v32
                    }
                    var v33 string = v2
                    _ = v29
                    _ = v31
                    _ = v33
                }
                if true {
                    v34 := GenS1{F0: 2, F1: 32767, F2: (42 != -899)}
                    _ = len(v6)
                    switch -78 {
                    case -9223372036854775808:
                        v1 += v1
                    default:
                        var v35 int = fi28
                        _ = v35
                    }
                    v6 = append(v6, 42)
                    v36 := GenS4{F0: 42, F1: 9223372036854775807}
                    _ = v34
                    _ = v36
                } else {
                    var v37 int32 = 42
                    var v38 int32 = -2147483648
                    _ = v37
                    _ = v38
                }
                _ = fi28
            }
            _ = fi27
        }
        var v39 float32 = (v1 - 3.14159)
        var v40 int32 = 1
        _ = v39
        _ = v40
    } else {
        var v41 bool = false
        if len(v6) > 0 {
            v6[0] = 32768
        }
        _ = v41
    }
    if true {
        v42 := v3[2147483647]
        _ = v42
    }
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
