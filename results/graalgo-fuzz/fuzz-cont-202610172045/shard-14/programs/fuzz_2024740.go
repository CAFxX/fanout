package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 bool; F1 uint8; F2 int32 }
type GenS3 struct { F0 int8; F1 uint; F2 int32 }
func genFn0(p0 uint64, p1 int16, p2 uint8) uint16 {
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            p1 = 32767
            var v9 uint8 = p2
            var v10 uint16 = 1
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 string = "foo bar"
        _ = fi7
        _ = v11
    }
    switch 127 {
    case 32767:
        for fi12 := 0; fi12 < 12; fi12++ {
            switch "foo bar" {
            case "":
                p1 = (p1 ^ 0)
            case "hello":
                p1 = p1
            case "a":
                p0 -= 1
            }
            var v13 string = ("hello" + "\n\t")
            _ = fi12
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 int64 = 1
    var v2 string = ("hello" + "\n\t")
    var v3 string = "hello"
    var v4 float64 = 1.5
    var v5 uint16 = 1
    var v6 float32 = 3.14159
    if ((127 < 0) || (220 < 128)) {
        var v14 uint16 = 196
        v6 += 1.5
        _ = v14
    }
    var v15 bool = (false && true)
    v16 := GenS1{F0: 42, F1: -699}
    v4 -= 1.5
    v3 = v2
    fmt.Println(v2, v5)
    var v17 int16 = 326
    v18 := []int{2147483647}
    if true {
        var v19 uint64 = 2
        _ = v19
    } else {
        switch int(v1) {
        case -9223372036854775808:
            if (-1 == (^(v17))) {
                fmt.Println(v5, v15, v2)
            }
        case 0:
            var v20 uint32 = 0
            _ = v20
        }
        _ = len(v18)
        if v15 {
            v21 := GenS2{F0: v15, F1: 2, F2: 655}
            switch -836 {
            case 9223372036854775807:
                v4 = v4
            case -9223372036854775808:
                var v22 int32 = -390
                _ = v22
            case -2147483648:
                for fi23 := 0; fi23 < 16; fi23++ {
                    if true {
                        for fi24 := 0; fi24 < 18; fi24++ {
                            fmt.Println(v1)
                            var v25 float32 = (v6 - v6)
                            _ = fi24
                            _ = v25
                        }
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                        var v26 int64 = v1
                        _ = v26
                    } else {
                        var v27 int32 = 2147483647
                        v28 := GenS3{F0: 127, F1: 712, F2: v27}
                        var v29 bool = v15
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    _ = fi23
                }
            }
            _ = v21
        }
    }
    var v30 int8 = 0
    if v15 {
        v1 -= (v1 & 220)
    }
    if v15 {
        if v15 {
            v31 := GenS1{F0: 1, F1: 2147483647}
            v30 = v30
            var v32 uint16 = (v5 + v5)
            _ = v31
            _ = v32
        }
        v17 += v17
        var v33 uint8 = 255
        _ = v33
    }
    v34 := GenS3{F0: -122, F1: 1, F2: 0}
    var v35 float32 = float32(-2.25)
    v1 += (v1 << ((18446744073709551615) & 63))
    var v36 int = 127
    v15 = ((v36 >> ((673) & 63)) == v36)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
