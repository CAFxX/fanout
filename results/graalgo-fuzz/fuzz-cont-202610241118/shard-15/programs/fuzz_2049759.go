package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int16; F2 int }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint64, p1 uint32) uint64 {
    for fi5 := 0; fi5 < 11; fi5++ {
        if false {
            for fi6 := 0; fi6 < 6; fi6++ {
                if ("\n\t" > "") {
                    var v7 bool = true
                    _ = v7
                }
                _ = fi6
            }
            var v8 string = ""
            _ = v8
        }
        _ = fi5
    }
    p0 += p0
    var v9 bool = true
    var v10 int16 = -1
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (^(p0))
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 2
    var v3 int32 = -2147483648
    var v4 uint32 = 65535
    var v11 int = 9223372036854775807
    _ = len(v1)
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := []int{}
        _ = fi12
        _ = v13
    }
    if false {
        v14 := GenS1{F0: 255}
        var v15 uint32 = v4
        v3 -= (v3 * (v3 | 1))
        _ = v14
        _ = v15
    }
    var v16 int32 = v3
    var v17 uint32 = (^(v4))
    v18 := GenS2{F0: 127, F1: -32768, F2: (v11 - (v11 - 42))}
    v19 := GenS3{F0: 42, F1: -630}
    for fi20 := 0; fi20 < 5; fi20++ {
        if true {
            for fi21 := 0; fi21 < 2; fi21++ {
                fmt.Println(v16)
                v22 := []int{fi20, v11, fi21, fi20}
                if len(v22) > 0 {
                    v22[0] = -2147483648
                }
                _ = fi21
                _ = v22
            }
            switch "" {
            case "":
                delete(v1, fi20)
            case "a":
                v23 := GenS4{F0: "a"}
                _ = v23
            }
        }
        _ = fi20
    }
    v24 := GenS5{F0: 42}
    for fi25 := 0; fi25 < 15; fi25++ {
        if true {
            var v26 int64 = -648
            _ = v26
        } else {
            v3 = v3
        }
        for fi27 := 0; fi27 < 14; fi27++ {
            var v28 uint64 = 1
            if (!true) {
                var v29 int8 = -1
                _ = v29
            } else {
                var v30 int = 13
                fmt.Println(v16, v30, v16)
                if false {
                    v28 = 0
                } else {
                    v11 += (fi25 + (v30 - (fi25 + -938)))
                    v31 := GenS5{F0: 127}
                    _ = v31
                }
                if (true || (2 == (v17 * 679))) {
                    v4 = (v17 - v17)
                    v32 := GenS3{F0: -297, F1: v3}
                    _ = v32
                } else {
                    v17 -= v17
                    v17 = v17
                    switch "foo bar" {
                    case "foo bar":
                        var v33 string = "hello"
                        _ = v33
                    }
                }
                _ = v30
            }
            var v34 uint64 = v28
            _ = fi27
            _ = v28
            _ = v34
        }
        v16 = v3
        _ = fi25
    }
    for fi35 := 0; fi35 < 1; fi35++ {
        fmt.Println(v3, v16)
        _ = fi35
    }
    v2 = v2
    var v36 uint32 = 742
    switch "" {
    case "hello":
        v17 += (v4 >> ((65535) & 63))
    case "foo bar":
        v37 := GenS2{F0: 0, F1: -1, F2: (v11 & v11)}
        _ = v37
    case "a":
        v16 -= 0
    default:
        v36 += 800
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v36
    _ = v4
}
